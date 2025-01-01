package theme

import (
	"image/color"

	"golang.org/x/sys/windows/registry"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	regKey  = `Software\Microsoft\Windows\CurrentVersion\Themes\Personalize`
	regName = `AppsUseLightTheme`
)

// NewTheme creates a new theme with the given font.
// On Windows, the theme is determined by the system theme.
func NewTheme(font text.Face) *Theme {
	k, _ := registry.OpenKey(registry.CURRENT_USER, regKey, registry.QUERY_VALUE)
	defer k.Close()
	v, _, _ := k.GetIntegerValue(regName)
	var styleState StyleState
	dark := v == 1
	if dark {
		styleState = StyleState{
			Idle: Style{
				BackgroundColor: color.RGBA{0x0f, 0x0f, 0x0f, 0xff},
				ForegroundColor: color.RGBA{0x00, 0x00, 0x00, 0xff},
			},
			Disabled: &Style{
				BackgroundColor: color.RGBA{0x88, 0x88, 0x88, 0xff},
				ForegroundColor: color.RGBA{0x00, 0x00, 0x00, 0xff},
			},
			Hover: &Style{
				BackgroundColor: color.RGBA{0x3f, 0x3f, 0x3f, 0xff},
				ForegroundColor: color.RGBA{0x3f, 0x3f, 0x3f, 0xff},
			},
			Pressed: &Style{
				BackgroundColor: color.RGBA{0x64, 0x64, 0x64, 0xff},
				ForegroundColor: color.RGBA{0xff, 0xff, 0xff, 0xff},
			},
		}
	} else {
		styleState = StyleState{
			Idle: Style{
				BackgroundColor: color.RGBA{0xff, 0xff, 0xff, 0xff},
				ForegroundColor: color.RGBA{0xff, 0xff, 0xff, 0xff},
			},
			Disabled: &Style{
				BackgroundColor: color.RGBA{0x88, 0x88, 0x88, 0xff},
				ForegroundColor: color.RGBA{0xff, 0xff, 0xff, 0xff},
			},
			Hover: &Style{
				BackgroundColor: color.RGBA{0x3f, 0x3f, 0x3f, 0xff},
				ForegroundColor: color.RGBA{0xff, 0xff, 0xff, 0xff},
			},
			Pressed: &Style{
				BackgroundColor: color.RGBA{0x64, 0x64, 0x64, 0xff},
				ForegroundColor: color.RGBA{0xff, 0xff, 0xff, 0xff},
			},
		}
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
