package imgui

type Position byte

const (
	Left Position = iota
	Right
	Top
	Bottom
)

var (
	Unassigned = &Widget{}
	Inactive   = &Widget{}
	Partial    = &Widget{}
)

const (
	KindRootContainer WidgetKind = "RootContainer"
	KindContainer     WidgetKind = "Container"
)

const (
	EventsAuto Events = iota
	EventsIgnore
)

type WidgetKind string
type Events byte

type Widget struct {
	kind   WidgetKind
	parent *Widget
	x      int
	y      int
	width  int
	height int

	offset struct {
		x int
		y int
	}

	clip struct {
		x      int
		y      int
		width  int
		height int
	}
}

func (w *Widget) absX() int {
	x := w.x
	if w.parent != nil {
		x += w.parent.absX() + w.parent.offset.x
	}

	return x
}

func (w *Widget) absY() int {
	y := w.y
	if w.parent != nil {
		y += w.parent.absY() + w.parent.offset.y
	}

	return y
}

type Input struct {
	Mouse struct {
		X      int
		Y      int
		IsDown bool
	}
}

type DrawCmd struct {
	Kind     WidgetKind
	Parent   *DrawCmd
	X        int
	Y        int
	Width    int
	Height   int
	IsHot    bool
	IsHover  bool
	IsActive bool
	IsBlur   bool
	IsFocus  bool
	Data     interface{}

	Clip struct {
		X      int
		Y      int
		Width  int
		Height int
	}
}

type DrawFunc func(cmd *DrawCmd)

type IMGUI struct {
	Input    Input
	Warm     *Widget
	Hot      *Widget
	Hover    *Widget
	Active   *Widget
	Blur     *Widget
	Focus    *Widget
	IsHot    bool
	IsHover  bool
	IsActive bool
	IsBlur   bool
	IsFocus  bool
	Events   Events
	DrawFunc DrawFunc

	// offset is just a helper struct that widgets can use to keep track
	// of things like mouse offsets between updates
	offset struct {
		x int
		y int
	}

	drawCmdData  interface{}
	drawBuffer   []*DrawCmd
	previousHot  *Widget
	containers   []*Widget
	freeWidgets  []*Widget
	usedWidgets  []*Widget
	freeDrawCmds []*DrawCmd
	usedDrawCmds []*DrawCmd
}

func New(drawFunc DrawFunc) *IMGUI {
	return &IMGUI{
		DrawFunc:    drawFunc,
		Warm:        Unassigned,
		Hot:         Unassigned,
		Hover:       Unassigned,
		Active:      Unassigned,
		Blur:        Unassigned,
		Focus:       Unassigned,
		previousHot: Unassigned,
	}
}

func (ui *IMGUI) getWidget() *Widget {
	var widget *Widget

	if len(ui.freeWidgets) == 0 {
		widget = &Widget{}
	} else {
		l := len(ui.freeWidgets) - 1
		widget = ui.freeWidgets[l]

		ui.freeWidgets[l] = nil
		ui.freeWidgets = ui.freeWidgets[:l]
	}

	ui.usedWidgets = append(ui.usedWidgets, widget)

	return widget
}

func (ui *IMGUI) getDrawCmd() *DrawCmd {
	var drawCmd *DrawCmd

	if len(ui.freeDrawCmds) == 0 {
		drawCmd = &DrawCmd{}
	} else {
		l := len(ui.freeDrawCmds) - 1
		drawCmd = ui.freeDrawCmds[l]

		ui.freeDrawCmds[l] = nil
		ui.freeDrawCmds = ui.freeDrawCmds[:l]
	}

	ui.usedDrawCmds = append(ui.usedDrawCmds, drawCmd)

	return drawCmd
}

func (ui *IMGUI) resetUsed() {
	// We go through this list backwards so that widgets on the next frame
	// can be assigned the same pointers
	for i := len(ui.usedWidgets) - 1; i >= 0; i-- {
		ui.freeWidgets = append(ui.freeWidgets, ui.usedWidgets[i])
	}

	ui.usedWidgets = ui.usedWidgets[:0]

	// We go through this list backwards so that widgets on the next frame
	// can be assigned the same pointers
	for i := len(ui.usedDrawCmds) - 1; i >= 0; i-- {
		ui.freeDrawCmds = append(ui.freeDrawCmds, ui.usedDrawCmds[i])
	}

	ui.usedDrawCmds = ui.usedDrawCmds[:0]

	ui.drawBuffer = ui.drawBuffer[:0]
}

func (ui *IMGUI) BeginUI(x, y, width, height int) {
	ui.resetUsed()

	ui.containers = ui.containers[:0]

	if ui.Input.Mouse.IsDown && ui.Active == Partial {
		ui.Active = ui.Hover
	}

	ui.Warm = Unassigned
	ui.Blur = Unassigned
	ui.Focus = Unassigned
	ui.Events = EventsAuto

	if ui.previousHot != ui.Hot {
		ui.Blur = ui.previousHot
		ui.Focus = ui.Hot
	}

	ui.beginContainer(KindRootContainer, x, y, width, height, 0, 0)
}

func (ui *IMGUI) EndUI() {
	ui.EndContainer()

	if len(ui.containers) > 0 {
		panic("not enough calls to end container")
	}

	ui.previousHot = ui.Hot
	ui.Hot = ui.Warm

	if !ui.Input.Mouse.IsDown {
		ui.Active = Unassigned
		ui.Hover = Unassigned
	} else if ui.Active == Unassigned {
		ui.Active = Inactive
	}

	if ui.Input.Mouse.IsDown {
		if ui.Active != Unassigned && ui.Active != Inactive {
			ui.Hover = ui.Active

			if ui.Active != ui.Hot {
				ui.Active = Partial
			}
		}
	} else {
		ui.Hover = ui.Hot
	}

	ui.IsHot = ui.Hot != Unassigned
	ui.IsHover = ui.Hover != Unassigned
	ui.IsActive = ui.Active != Unassigned
	ui.IsBlur = ui.Blur != Unassigned
	ui.IsFocus = ui.Focus != Unassigned
}

func (ui *IMGUI) beginContainer(kind WidgetKind, x, y, width, height, offsetX, offsetY int) {
	container := ui.newWidget(kind, x, y, width, height)

	container.offset.x = offsetX
	container.offset.y = offsetY

	if container.parent != nil {
		if container.x+container.width > container.parent.width {
			container.width = container.parent.width - container.x
		}

		if container.y+container.height > container.parent.height {
			container.height = container.parent.height - container.y
		}
	}

	ui.containers = append(ui.containers, container)

	ui.draw(container, nil)
}

func (ui *IMGUI) BeginContainer(x, y, width, height, offsetX, offsetY int) {
	ui.beginContainer(KindContainer, x, y, width, height, offsetX, offsetY)
}

func (ui *IMGUI) EndContainer() {
	l := len(ui.containers) - 1
	if l < 0 {
		panic("too many calls to end container")
	}

	ui.containers[l] = nil
	ui.containers = ui.containers[:l]
}

func (ui *IMGUI) newWidget(kind WidgetKind, x, y, width, height int) *Widget {
	w := ui.getWidget()

	w.kind = kind
	w.parent = nil
	w.x = x
	w.y = y
	w.width = width
	w.height = height
	w.offset.x = 0
	w.offset.y = 0
	w.clip.x = w.x
	w.clip.y = w.y
	w.clip.width = w.width
	w.clip.height = w.height

	if len(ui.containers) > 0 {
		w.parent = ui.containers[len(ui.containers)-1]

		w.clip.x = w.absX()
		w.clip.y = w.absY()
		parentX := w.parent.clip.x
		parentY := w.parent.clip.y
		parentRight := parentX + w.parent.clip.width
		parentBottom := parentY + w.parent.clip.height

		if parentX > w.clip.x {
			w.clip.width += w.clip.x - parentX
			w.clip.x = parentX
		}

		if parentY > w.clip.y {
			w.clip.height += w.clip.y - parentY
			w.clip.y = parentY
		}

		if w.clip.x+w.clip.width > parentRight {
			w.clip.width = parentRight - w.clip.x
		}

		if w.clip.y+w.clip.height > parentBottom {
			w.clip.height = parentBottom - w.clip.y
		}
	}

	return w
}

func (ui *IMGUI) doButtonLogic(w *Widget) bool {
	if ui.Events == EventsIgnore {
		return false
	}

	if ui.RegionHitRect(w.clip.x, w.clip.y, w.clip.width, w.clip.height) {
		ui.Warm = w
	}

	triggered := ui.Hot == w && ui.Active == Unassigned && ui.Input.Mouse.IsDown
	if triggered {
		ui.Active = w
	}

	return triggered
}

func (ui *IMGUI) RegionHitRect(x, y, width, height int) bool {
	return ui.Input.Mouse.X >= x && ui.Input.Mouse.X < x+width && ui.Input.Mouse.Y >= y && ui.Input.Mouse.Y < y+height
}

func (ui *IMGUI) DrawData(data interface{}) {
	ui.drawCmdData = data
}

func (ui *IMGUI) draw(w *Widget, drawCmd *DrawCmd) {
	if w.clip.width <= 0 || w.clip.height <= 0 {
		return
	}

	if drawCmd == nil {
		drawCmd = ui.getDrawCmd()
	}

	drawCmd.Kind = w.kind
	drawCmd.Clip.X = w.clip.x
	drawCmd.Clip.Y = w.clip.y
	drawCmd.Clip.Width = w.clip.width
	drawCmd.Clip.Height = w.clip.height
	drawCmd.Data = ui.drawCmdData

	ui.drawCmdData = nil

	ui.drawBuffer = append(ui.drawBuffer, drawCmd)
}

func (ui *IMGUI) Draw() {
	if ui.DrawFunc == nil {
		return
	}

	for _, cmd := range ui.drawBuffer {
		ui.DrawFunc(cmd)
	}
}
