package theme

import (
	"os/exec"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// NewTheme creates a new Theme based on the current macOS appearance (light or dark).
func NewTheme(font text.Face) *Theme {
	appearance, err := getMacOSAppearance()
	if err != nil {
		appearance = "Dark" // Default to Dark if the appearance cannot be determined
	}

	var styleState StyleState
	if appearance == "Dark" {
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

// getMacOSAppearance returns the current macOS appearance (light or dark).
func getMacOSAppearance() (string, error) {
	cmd := exec.Command("defaults", "read", "-g", "AppleInterfaceStyle")
	output, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			return "Light", nil // Default to Light if the setting is not found
		}
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// loadDarkTheme loads the dark theme.
func loadDarkTheme() *Theme {
	// ...implementation for loading dark theme...
	return &Theme{}
}

// loadLightTheme loads the light theme.
func loadLightTheme() *Theme {
	// ...implementation for loading light theme...
	return &Theme{}
}
