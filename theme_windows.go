package theme

import (
	"golang.org/x/sys/windows/registry"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	regKey  = `Software\Microsoft\Windows\CurrentVersion\Themes\Personalize`
	regName = `AppsUseLightTheme`
)

// NewTheme creates a new theme with the given font.
// On Windows, the theme is determined by the system theme, i.e. light or dark.
func NewTheme(font text.Face) *Theme {
	k, _ := registry.OpenKey(registry.CURRENT_USER, regKey, registry.QUERY_VALUE)
	defer k.Close()
	v, _, _ := k.GetIntegerValue(regName)
	var styleState StyleState
	dark := v == 1
	if dark {
		styleState = darkStyleState
	} else {
		styleState = lightStyleState
	}
	return &Theme{
		ButtonStyle:      styleState,
		LabelStyle:       styleState,
		TextStyle:        styleState,
		ProgressBarStyle: styleState,
		TextInputStyle:   styleState,
		ContainerStyle:   styleState,
		CaretStyle:       styleState,
		GraphicStyle:     styleState,
		Font:             font,
	}
}
