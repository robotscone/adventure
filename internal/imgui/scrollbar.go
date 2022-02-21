package imgui

const (
	KindHScrollbar      WidgetKind = "HScrollbar"
	KindHScrollbarThumb WidgetKind = "HScrollbarThumb"
	KindVScrollbar      WidgetKind = "VScrollbar"
	KindVScrollbarThumb WidgetKind = "VScrollbarThumb"
)

func (ui *IMGUI) HScrollbar(x, y, width, height, thumbWidth int, min, max float64, input *float64) bool {
	anchor := Left
	if max < min {
		anchor = Right
		min, max = max, min
	}

	value := *input - min
	rng := max - min

	if value < 0.0 {
		value = 0.0
	} else if value > rng {
		value = rng
	}

	bar := ui.newWidget(KindHScrollbar, x, y, width, height)

	ui.doButtonLogic(bar)

	thumbHalf := int(float64(thumbWidth) / 2)
	x += thumbHalf
	width -= thumbWidth

	position := int(float64(width) * (value / rng))
	thumbPosition := x
	if anchor == Left {
		thumbPosition += position
	} else {
		thumbPosition += width - position
	}

	thumb := ui.newWidget(KindHScrollbarThumb, thumbPosition-thumbHalf, y, thumbWidth, height)

	if ui.doButtonLogic(thumb) {
		ui.offset.x = ui.Input.Mouse.X - thumbPosition
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
		newPosition := ui.Input.Mouse.X - x
		if ui.Active == thumb {
			newPosition -= ui.offset.x
		}
		if newPosition < 0 {
			newPosition = 0
		} else if newPosition > width {
			newPosition = width
		}

		newValue := rng * (float64(newPosition) / float64(width))
		if anchor == Right {
			newValue = rng - newValue
		}

		if newValue != value {
			*input = min + newValue

			return true
		}
	}

	return false
}

func (ui *IMGUI) VScrollbar(x, y, width, height, thumbHeight int, min, max float64, input *float64) bool {
	anchor := Top
	if max < min {
		anchor = Bottom
		min, max = max, min
	}

	value := *input - min
	rng := max - min

	if value < 0.0 {
		value = 0.0
	} else if value > rng {
		value = rng
	}

	bar := ui.newWidget(KindVScrollbar, x, y, width, height)

	ui.doButtonLogic(bar)

	thumbHalf := int(float64(thumbHeight) / 2)
	y += thumbHalf
	height -= thumbHeight

	position := int(float64(height) * (value / rng))
	thumbPosition := y
	if anchor == Top {
		thumbPosition += position
	} else {
		thumbPosition += height - position
	}

	thumb := ui.newWidget(KindVScrollbarThumb, x, thumbPosition-thumbHalf, width, thumbHeight)

	if ui.doButtonLogic(thumb) {
		ui.offset.y = ui.Input.Mouse.Y - thumbPosition
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
		newPosition := ui.Input.Mouse.Y - y
		if ui.Active == thumb {
			newPosition -= ui.offset.y
		}
		if newPosition < 0 {
			newPosition = 0
		} else if newPosition > height {
			newPosition = height
		}

		newValue := rng * (float64(newPosition) / float64(height))
		if anchor == Bottom {
			newValue = rng - newValue
		}

		if newValue != value {
			*input = min + newValue

			return true
		}
	}

	return false
}
