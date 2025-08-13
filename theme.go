package theme

import (
	"time"

	"fyne.io/fyne/v2"
)

/* package theme contains my themes */

const (
	ColorNameGray      fyne.ThemeColorName = "gray"
	ColorNameDiscrete  fyne.ThemeColorName = "discrete"
	ColorNameProminent fyne.ThemeColorName = "prominent"
	ColorNameImposing  fyne.ThemeColorName = "imposing"
	ColorNameScreaming fyne.ThemeColorName = "screaming"
	ColorNamePrimary   fyne.ThemeColorName = "primary"
	ColorNameSecondary fyne.ThemeColorName = "secondary"
	ColorNameTertiary  fyne.ThemeColorName = "tertiary"
	ColorNameMessage   fyne.ThemeColorName = "message"
	ColorNameWarning   fyne.ThemeColorName = "warning"
	ColorNameError     fyne.ThemeColorName = "error"
)

const (
	SizeNameMajorText     fyne.ThemeSizeName = "major"
	SizeNameMinorText     fyne.ThemeSizeName = "minor"
	SizeNamePrimaryText   fyne.ThemeSizeName = "primary"
	SizeNameSecondaryText fyne.ThemeSizeName = "secondary"
)

func ThemeVariantWatcher(a fyne.App, t fyne.Theme) {
	p := a.Settings().ThemeVariant()
	fyne.DoAndWait(func() {
		a.Settings().SetTheme(t)
	})
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for range ticker.C {
		v := a.Settings().ThemeVariant()
		if v != p {
			p = v
			fyne.DoAndWait(func() { a.Settings().SetTheme(t) })
		}
	}
}
