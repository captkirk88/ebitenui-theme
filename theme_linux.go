package theme

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// NewTheme creates a new theme with the given font.
// On Linux the theme is determined by the system theme, i.e. light or dark.
// On Linux, the theme is defaulted to dark if the system theme is undetermined.
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
