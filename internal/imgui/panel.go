package imgui

const KindPanel WidgetKind = "Panel"

func (ui *IMGUI) Panel(x, y, width, height int) bool {
	panel := ui.newWidget(KindPanel, x, y, width, height)

	ui.doButtonLogic(panel)

	ui.IsHot = ui.Hot == panel
	ui.IsHover = ui.Hover == panel
	ui.IsActive = ui.Active == panel
	ui.IsBlur = ui.Blur == panel
	ui.IsFocus = ui.Focus == panel

	drawCmd := ui.getDrawCmd()

	drawCmd.X = panel.absX()
	drawCmd.Y = panel.absY()
	drawCmd.Width = panel.width
	drawCmd.Height = panel.height
	drawCmd.IsHot = ui.IsHot
	drawCmd.IsHover = ui.IsHover
	drawCmd.IsActive = ui.IsActive
	drawCmd.IsBlur = ui.IsBlur
	drawCmd.IsFocus = ui.IsFocus

	ui.draw(panel, drawCmd)

	return !ui.Input.Mouse.IsDown && ui.IsHot && ui.IsActive
}
