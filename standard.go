package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var _ fyne.Theme = (*Standard)(nil)

type Standard struct{}

func NewStandard() *Standard {
	return &Standard{}
}

/* Color implements fyne.Theme. */
func (Standard) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
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
func (Standard) Font(style fyne.TextStyle) fyne.Resource {
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
func (Standard) Icon(name fyne.ThemeIconName) fyne.Resource {
	switch name {
	default:
		return theme.DefaultTheme().Icon(name)
	}
}

/* Size implements fyne.Theme. */
func (Standard) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameInnerPadding:
		return 2
	case theme.SizeNamePadding:
		return 4
	case theme.SizeNameCaptionText:
		return 11
	case SizeNameMinorText:
		return 10
	default:
		return theme.DefaultTheme().Size(name)
	}
}
