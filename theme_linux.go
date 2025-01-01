package theme

import (
	"fmt"
	"image/color"
	"os/exec"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// NewTheme creates a new theme with the given font.
// On Linux, the theme is defaulted to dark.
func NewTheme(font text.Face) *Theme {
	var styleState StyleState
	var dark bool
	if isGNOME_Dark() {
		dark = true
	} else if isKDE_Dark() {
		dark = true
	} else if isXFCE_Dark() {
		dark = true
	}
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

func isXFCE_Dark() bool {
	cmd := exec.Command("xfconf-query", "-c", "xsettings", "-p", "/Net/ThemeName")
	output, err := cmd.Output()

	if err == nil {
		// Trim the output to get the theme name
		theme := strings.TrimSpace(string(output))
		fmt.Println("Current theme:", theme)

		// Check if the theme name contains 'dark'
		if strings.Contains(strings.ToLower(theme), "dark") {
			return true
		}
	}
	return false
}

func isKDE_Dark() bool {
	cmd := exec.Command("lookandfeeltool", "--list")
	output, err := cmd.Output()

	if err == nil {
		// Split the output into lines
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			// Check if the line contains an asterisk (*) indicating the current theme
			if strings.Contains(line, "*") {
				theme := strings.TrimSpace(line)
				// Check if the theme name contains 'dark'
				if strings.Contains(strings.ToLower(theme), "dark") {
					return true
				}
				break
			}
		}
	}
	return false
}

func isGNOME_Dark() bool {
	cmd := exec.Command("gsettings", "get", "org.gnome.desktop.interface", "color-scheme")
	output, err := cmd.Output()

	if err == nil {
		// Trim the output and check if it is 'prefer-dark'
		theme := strings.TrimSpace(string(output))
		if theme == "'prefer-dark'" {
			return true
		}
	}
	return false
}
