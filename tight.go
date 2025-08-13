package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var _ fyne.Theme = (*Tight)(nil)

type Tight struct{}

func NewTight() *Tight {
	return &Tight{}
}

/* Color implements fyne.Theme. */
func (Tight) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case ColorNameDiscrete:
		if variant == theme.VariantLight {
			return color.NRGBA{R: 0x77, G: 0x77, B: 0x77, A: 0xff}
		}
		return color.NRGBA{R: 0x77, G: 0x77, B: 0x77, A: 0xff}
	case theme.ColorNameForeground:
		if variant == theme.VariantLight {
			return color.NRGBA{R: 0x27, G: 0x27, B: 0x27, A: 0xff}
		}
		return color.NRGBA{R: 0xf0, G: 0xf1, B: 0xf1, A: 0xff}
	case ColorNameProminent:
		if variant == theme.VariantLight {
			return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xff}
		}
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

/* Font implements fyne.Theme. */
func (Tight) Font(style fyne.TextStyle) fyne.Resource {
	if style.Monospace {
		return theme.DefaultTheme().Font(style)
	}
	if style.Italic {
		if style.Bold {
			return theme.DefaultTheme().Font(style)
		}
		return theme.DefaultTheme().Font(style)
	}
	return theme.DefaultTheme().Font(style)
}

/* Icon implements fyne.Theme. */
func (Tight) Icon(name fyne.ThemeIconName) fyne.Resource {
	switch name {
	default:
		return theme.DefaultTheme().Icon(name)
	}
}

/* Size implements fyne.Theme. */
func (Tight) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameSeparatorThickness:
		return 1 // 1
	case theme.SizeNameCaptionText:
		return 11 // 11
	case theme.SizeNameInlineIcon:
		return 20 // 20
	case theme.SizeNameInnerPadding:
		return 4 // 8
	case theme.SizeNameLineSpacing:
		return 3 // 4
	case theme.SizeNamePadding:
		return 3 // 4
	case theme.SizeNameScrollBar:
		return 12 // 12
	case theme.SizeNameScrollBarSmall:
		return 3 // 3
	case theme.SizeNameText:
		return 14 // 14
	case theme.SizeNameHeadingText:
		return 24 // 24
	case theme.SizeNameSubHeadingText:
		return 18 // 18
	case theme.SizeNameInputBorder:
		return 1 // 1
	case theme.SizeNameInputRadius:
		return 6 // 5
	case theme.SizeNameSelectionRadius:
		return 4 // 3
	case theme.SizeNameScrollBarRadius:
		return 4 // 3
	case theme.SizeNameWindowButtonHeight:
		return 16 // 16
	case theme.SizeNameWindowButtonRadius:
		return 8 // 8
	case theme.SizeNameWindowButtonIcon:
		return 14 // 14
	case theme.SizeNameWindowTitleBarHeight:
		return 26 // 26
	case SizeNameMinorText:
		return 10
	default:
		return theme.DefaultTheme().Size(name)
	}
}
