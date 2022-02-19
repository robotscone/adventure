package imgui

const (
	KindHProgressBar       WidgetKind = "HProgressBar"
	KindHProgressIndicator WidgetKind = "HProgressIndicator"
	KindVProgressBar       WidgetKind = "VProgressBar"
	KindVProgressIndicator WidgetKind = "VProgressIndicator"
)

func (ui *IMGUI) HProgressBar(x, y, width, height int, min, max, value float64, anchor Position) bool {
	value = value - min
	normMax := max - min

	if value < 0.0 {
		value = 0.0
	} else if value > normMax {
		value = normMax
	}

	container := ui.newWidget(KindHProgressBar, x, y, width, height)

	indicatorWidth := int(float64(width) * (value / normMax))
	if anchor == Right {
		x = x + width - indicatorWidth
	}

	indicator := ui.newWidget(KindHProgressIndicator, x, y, indicatorWidth, height)

	ui.doButtonLogic(container)
	ui.doButtonLogic(indicator)

	ui.IsHot = ui.Active == container || ui.Active == indicator
	ui.IsHover = ui.Hover == container || ui.Hover == indicator
	ui.IsActive = ui.Active == container || ui.Active == indicator
	ui.IsBlur = ui.Blur == container || ui.Blur == indicator
	ui.IsFocus = ui.Focus == container || ui.Focus == indicator

	barDrawCmd := ui.getDrawCmd()

	barDrawCmd.X = container.absX()
	barDrawCmd.Y = container.absY()
	barDrawCmd.Width = container.width
	barDrawCmd.Height = container.height
	barDrawCmd.IsHot = ui.Hot == container
	barDrawCmd.IsHover = ui.Hover == container
	barDrawCmd.IsActive = ui.Active == container
	barDrawCmd.IsBlur = ui.Blur == container
	barDrawCmd.IsFocus = ui.Focus == container

	data := ui.drawCmdData

	ui.draw(container, barDrawCmd)

	indicatorDrawCmd := ui.getDrawCmd()

	indicatorDrawCmd.Parent = barDrawCmd
	indicatorDrawCmd.X = indicator.absX()
	indicatorDrawCmd.Y = indicator.absY()
	indicatorDrawCmd.Width = indicator.width
	indicatorDrawCmd.Height = indicator.height
	indicatorDrawCmd.IsHot = ui.Hot == indicator
	indicatorDrawCmd.IsHover = ui.Hover == indicator
	indicatorDrawCmd.IsActive = ui.Active == indicator
	indicatorDrawCmd.IsBlur = ui.Blur == indicator
	indicatorDrawCmd.IsFocus = ui.Focus == indicator

	ui.drawCmdData = data

	ui.draw(indicator, indicatorDrawCmd)

	return !ui.Input.Mouse.IsDown && ui.IsHot && ui.IsActive
}

func (ui *IMGUI) VProgressBar(x, y, width, height int, min, max, value float64, anchor Position) bool {
	value = value - min
	normMax := max - min

	if value < 0.0 {
		value = 0.0
	} else if value > normMax {
		value = normMax
	}

	container := ui.newWidget(KindVProgressBar, x, y, width, height)

	indicatorHeight := int(float64(height) * (value / normMax))
	if anchor == Bottom {
		y = y + height - indicatorHeight
	}

	indicator := ui.newWidget(KindVProgressIndicator, x, y, width, indicatorHeight)

	ui.doButtonLogic(container)
	ui.doButtonLogic(indicator)

	ui.IsHot = ui.Active == container || ui.Active == indicator
	ui.IsHover = ui.Hover == container || ui.Hover == indicator
	ui.IsActive = ui.Active == container || ui.Active == indicator
	ui.IsBlur = ui.Blur == container || ui.Blur == indicator
	ui.IsFocus = ui.Focus == container || ui.Focus == indicator

	barDrawCmd := ui.getDrawCmd()

	barDrawCmd.X = container.absX()
	barDrawCmd.Y = container.absY()
	barDrawCmd.Width = container.width
	barDrawCmd.Height = container.height
	barDrawCmd.IsHot = ui.Hot == container
	barDrawCmd.IsHover = ui.Hover == container
	barDrawCmd.IsActive = ui.Active == container
	barDrawCmd.IsBlur = ui.Blur == container
	barDrawCmd.IsFocus = ui.Focus == container

	data := ui.drawCmdData

	ui.draw(container, barDrawCmd)

	indicatorDrawCmd := ui.getDrawCmd()

	indicatorDrawCmd.Parent = barDrawCmd
	indicatorDrawCmd.X = indicator.absX()
	indicatorDrawCmd.Y = indicator.absY()
	indicatorDrawCmd.Width = indicator.width
	indicatorDrawCmd.Height = indicator.height
	indicatorDrawCmd.IsHot = ui.Hot == indicator
	indicatorDrawCmd.IsHover = ui.Hover == indicator
	indicatorDrawCmd.IsActive = ui.Active == indicator
	indicatorDrawCmd.IsBlur = ui.Blur == indicator
	indicatorDrawCmd.IsFocus = ui.Focus == indicator

	ui.drawCmdData = data

	ui.draw(indicator, indicatorDrawCmd)

	return !ui.Input.Mouse.IsDown && ui.IsHot && ui.IsActive
}
