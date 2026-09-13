// Copyright (c) 2026 Nawakarit
// Copyright (c) 2026 Naawaakrit
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License v3.0.
package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type MyTheme struct{}

func (m MyTheme) Color(name fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	if v == theme.VariantDark {
		switch name {
		case theme.ColorNameBackground:
			return color.NRGBA{11, 18, 32, 255}
		case theme.ColorNameForeground:
			return color.NRGBA{244, 247, 251, 255}
		case theme.ColorNameButton:
			return color.NRGBA{28, 42, 61, 255}
		case theme.ColorNamePressed:
			return color.NRGBA{255, 181, 71, 255}
		case theme.ColorNameHover:
			return color.NRGBA{41, 65, 93, 255}
		case theme.ColorNameDisabledButton:
			return color.NRGBA{37, 48, 64, 255}
		case theme.ColorNameDisabled:
			return color.NRGBA{130, 145, 164, 255}
		case theme.ColorNameFocus:
			return color.NRGBA{34, 211, 238, 100}
		case theme.ColorNamePrimary:
			return color.NRGBA{34, 211, 238, 255}
		case theme.ColorNameInputBackground:
			return color.NRGBA{22, 34, 53, 255}
		case theme.ColorNamePlaceHolder:
			return color.NRGBA{159, 176, 197, 255}
		case theme.ColorNameMenuBackground:
			return color.NRGBA{23, 38, 58, 255}
		case theme.ColorNameOverlayBackground:
			return color.NRGBA{16, 26, 43, 255}
		case theme.ColorNameShadow:
			return color.NRGBA{0, 0, 0, 150}
		case theme.ColorNameError:
			return color.NRGBA{255, 93, 115, 255}
		case theme.ColorNameSuccess:
			return color.NRGBA{66, 211, 146, 255}
		case theme.ColorNameWarning:
			return color.NRGBA{255, 181, 71, 255}
		}
	} else {
		switch name {
		case theme.ColorNameBackground:
			return color.NRGBA{246, 249, 252, 255}
		case theme.ColorNameForeground:
			return color.NRGBA{23, 32, 51, 255}
		case theme.ColorNameButton:
			return color.NRGBA{229, 238, 248, 255}
		case theme.ColorNamePressed:
			return color.NRGBA{255, 181, 71, 255}
		case theme.ColorNameHover:
			return color.NRGBA{215, 245, 250, 255}
		case theme.ColorNameDisabledButton:
			return color.NRGBA{220, 226, 233, 255}
		case theme.ColorNameDisabled:
			return color.NRGBA{138, 150, 163, 255}
		case theme.ColorNameFocus:
			return color.NRGBA{0, 124, 145, 80}
		case theme.ColorNamePrimary:
			return color.NRGBA{0, 124, 145, 255}
		case theme.ColorNameInputBackground:
			return color.NRGBA{255, 255, 255, 255}
		case theme.ColorNamePlaceHolder:
			return color.NRGBA{113, 128, 150, 255}
		case theme.ColorNameMenuBackground:
			return color.NRGBA{255, 255, 255, 255}
		case theme.ColorNameOverlayBackground:
			return color.NRGBA{255, 255, 255, 255}
		case theme.ColorNameShadow:
			return color.NRGBA{23, 32, 51, 55}
		case theme.ColorNameError:
			return color.NRGBA{207, 52, 69, 255}
		case theme.ColorNameSuccess:
			return color.NRGBA{23, 145, 93, 255}
		case theme.ColorNameWarning:
			return color.NRGBA{191, 118, 0, 255}
		}
	}
	return theme.DefaultTheme().Color(name, v)
}

func (m MyTheme) Font(s fyne.TextStyle) fyne.Resource {
	return myFont
}

func (m MyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (m MyTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNamePadding:
		return 4
	case theme.SizeNameSeparatorThickness:
		return 1
	case theme.SizeNameInlineIcon:
		return 19
	case theme.SizeNameScrollBar:
		return 12
	case theme.SizeNameScrollBarSmall:
		return 3
	case theme.SizeNameText:
		return 14
	case theme.SizeNameHeadingText:
		return 20
	case theme.SizeNameSubHeadingText:
		return 16
	case theme.SizeNameCaptionText:
		return 12
	case theme.SizeNameInputBorder:
		return 1
	}
	return theme.DefaultTheme().Size(name)
}
