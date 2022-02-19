package imgui

const (
	KindHScrollbar      WidgetKind = "HScrollbar"
	KindHScrollbarThumb WidgetKind = "HScrollbarThumb"
	KindVScrollbar      WidgetKind = "VScrollbar"
	KindVScrollbarThumb WidgetKind = "VScrollbarThumb"
)

func (ui *IMGUI) HScrollbar(x, y, width, height, thumbWidth int, min, max float64, value *float64) bool {
	val := *value - min
	normMax := max - min

	if val < 0.0 {
		val = 0.0
	} else if val > normMax {
		val = normMax
	}

	bar := ui.newWidget(KindHScrollbar, x, y, width, height)

	ui.doButtonLogic(bar)

	thumbWidthHalf := int(float64(thumbWidth) / 2)
	realX := x + thumbWidthHalf

	maxThumbX := width - thumbWidth
	thumbX := realX + int(float64(maxThumbX)*(val/normMax)) - thumbWidthHalf

	thumb := ui.newWidget(KindHScrollbarThumb, thumbX, y, thumbWidth, height)

	if ui.doButtonLogic(thumb) {
		ui.offset.x = ui.Input.Mouse.X - thumbX - thumbWidthHalf
	}

	ui.IsHot = ui.Active == bar || ui.Active == thumb
	ui.IsHover = ui.Hover == bar || ui.Hover == thumb
	ui.IsActive = ui.Active == bar || ui.Active == thumb
	ui.IsBlur = ui.Blur == bar || ui.Blur == thumb
	ui.IsFocus = ui.Focus == bar || ui.Focus == thumb

	barDrawCmd := ui.getDrawCmd()

	barDrawCmd.X = bar.absX()
	barDrawCmd.Y = bar.absY()
	barDrawCmd.Width = bar.width
	barDrawCmd.Height = bar.height
	barDrawCmd.IsHot = ui.Hot == bar
	barDrawCmd.IsHover = ui.Hover == bar
	barDrawCmd.IsActive = ui.Active == bar
	barDrawCmd.IsBlur = ui.Blur == bar
	barDrawCmd.IsFocus = ui.Focus == bar

	ui.draw(bar, barDrawCmd)

	thumbDrawCmd := ui.getDrawCmd()

	thumbDrawCmd.Parent = barDrawCmd
	thumbDrawCmd.X = thumb.absX()
	thumbDrawCmd.Y = thumb.absY()
	thumbDrawCmd.Width = thumb.width
	thumbDrawCmd.Height = thumb.height
	thumbDrawCmd.IsHot = ui.Hot == thumb
	thumbDrawCmd.IsHover = ui.Hover == thumb
	thumbDrawCmd.IsActive = ui.Active == thumb
	thumbDrawCmd.IsBlur = ui.Blur == thumb
	thumbDrawCmd.IsFocus = ui.Focus == thumb

	ui.draw(thumb, thumbDrawCmd)

	if ui.IsActive {
		newX := ui.Input.Mouse.X - realX
		if ui.Active == thumb {
			newX -= ui.offset.x
		}
		if newX < 0 {
			newX = 0
		} else if newX > maxThumbX {
			newX = maxThumbX
		}

		newValue := float64(newX) * normMax / float64(maxThumbX)
		if newValue != val {
			*value = newValue + min

			return true
		}
	}

	return false
}

func (ui *IMGUI) VScrollbar(x, y, width, height, thumbHeight int, min, max float64, value *float64) bool {
	val := *value - min
	normMax := max - min

	if val < 0.0 {
		val = 0.0
	} else if val > normMax {
		val = normMax
	}

	bar := ui.newWidget(KindVScrollbar, x, y, width, height)

	ui.doButtonLogic(bar)

	thumbHeightHalf := int(float64(thumbHeight) / 2)
	realY := y + thumbHeightHalf

	maxThumbY := height - thumbHeight
	thumbY := realY + int(float64(maxThumbY)*(val/normMax)) - thumbHeightHalf

	thumb := ui.newWidget(KindVScrollbarThumb, x, thumbY, width, thumbHeight)

	if ui.doButtonLogic(thumb) {
		ui.offset.y = ui.Input.Mouse.Y - thumbY - thumbHeightHalf
	}

	ui.IsHot = ui.Active == bar || ui.Active == thumb
	ui.IsHover = ui.Hover == bar || ui.Hover == thumb
	ui.IsActive = ui.Active == bar || ui.Active == thumb
	ui.IsBlur = ui.Blur == bar || ui.Blur == thumb
	ui.IsFocus = ui.Focus == bar || ui.Focus == thumb

	barDrawCmd := ui.getDrawCmd()

	barDrawCmd.X = bar.absX()
	barDrawCmd.Y = bar.absY()
	barDrawCmd.Width = bar.width
	barDrawCmd.Height = bar.height
	barDrawCmd.IsHot = ui.Hot == bar
	barDrawCmd.IsHover = ui.Hover == bar
	barDrawCmd.IsActive = ui.Active == bar
	barDrawCmd.IsBlur = ui.Blur == bar
	barDrawCmd.IsFocus = ui.Focus == bar

	ui.draw(bar, barDrawCmd)

	thumbDrawCmd := ui.getDrawCmd()

	thumbDrawCmd.Parent = barDrawCmd
	thumbDrawCmd.X = thumb.absX()
	thumbDrawCmd.Y = thumb.absY()
	thumbDrawCmd.Width = thumb.width
	thumbDrawCmd.Height = thumb.height
	thumbDrawCmd.IsHot = ui.Hot == thumb
	thumbDrawCmd.IsHover = ui.Hover == thumb
	thumbDrawCmd.IsActive = ui.Active == thumb
	thumbDrawCmd.IsBlur = ui.Blur == thumb
	thumbDrawCmd.IsFocus = ui.Focus == thumb

	ui.draw(thumb, thumbDrawCmd)

	if ui.IsActive {
		newY := ui.Input.Mouse.Y - realY
		if ui.Active == thumb {
			newY -= ui.offset.y
		}
		if newY < 0 {
			newY = 0
		} else if newY > maxThumbY {
			newY = maxThumbY
		}

		newValue := float64(newY) * normMax / float64(maxThumbY)
		if newValue != val {
			*value = newValue + min

			return true
		}
	}

	return false
}
