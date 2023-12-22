package imgui

const KindButton WidgetKind = "Button"

func (ui *IMGUI) Button(x, y, width, height int) bool {
	button := ui.newWidget(KindButton, x, y, width, height)

	ui.doButtonLogic(button)

	ui.IsHot = ui.Hot == button
	ui.IsHover = ui.Hover == button
	ui.IsActive = ui.Active == button
	ui.IsBlur = ui.Blur == button
	ui.IsFocus = ui.Focus == button

	drawCmd := ui.getDrawCmd()

	drawCmd.X = button.absX()
	drawCmd.Y = button.absY()
	drawCmd.Width = button.width
	drawCmd.Height = button.height
	drawCmd.IsHot = ui.IsHot
	drawCmd.IsHover = ui.IsHover
	drawCmd.IsActive = ui.IsActive
	drawCmd.IsBlur = ui.IsBlur
	drawCmd.IsFocus = ui.IsFocus

	ui.draw(button, drawCmd)

	return !ui.Input.Mouse.IsDown && ui.IsHot && ui.IsActive
}

func (ui *IMGUI) ImageButton(x, y, width, height int, spriteKey string) bool {
	ui.DrawData(spriteKey)

	return ui.Button(x, y, width, height)
}
