package theme

import (
	"image/color"

	eimage "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// Theme represents the theme data for EbitenUI.
type Theme struct {
	ButtonStyle      StyleState `yaml:"button"`
	LabelStyle       StyleState `yaml:"label"`
	TextStyle        StyleState `yaml:"text"`
	ProgressBarStyle StyleState `yaml:"progressbar"`
	TextInputStyle   StyleState `yaml:"textinput"`
	ContainerStyle   StyleState `yaml:"container"`
	CaretStyle       StyleState `yaml:"caret"`
	GraphicStyle     StyleState `yaml:"graphic"`
	Font             text.Face  `yaml:"-"`
}

// Style represents the style data for a widget.
type Style struct {
	BackgroundColor color.RGBA        `yaml:"bg_color"`
	ForegroundColor color.RGBA        `yaml:"fg_color"`
	Image           *eimage.NineSlice `yaml:"image"`
}

// StyleState represents the style state data for a widget.
type StyleState struct {
	Idle       Style                  `yaml:"idle"`
	Disabled   *Style                 `yaml:"disabled"`
	Hover      *Style                 `yaml:"hover"`
	Pressed    *Style                 `yaml:"pressed"`
	CustomData map[string]interface{} `yaml:"custom_data"`
	Font       *text.Face             `yaml:"-"`
}

// Button returns the button options for the theme.
func (th *Theme) Button(text string) []widget.ButtonOpt {
	th.ButtonStyle.valid(th)
	return []widget.ButtonOpt{
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:     th.ButtonStyle.Idle.Image,
			Disabled: th.ButtonStyle.Disabled.Image,
			Hover:    th.ButtonStyle.Hover.Image,
			Pressed:  th.ButtonStyle.Pressed.Image,
		}),
		widget.ButtonOpts.TextColor(&widget.ButtonTextColor{
			Idle:     th.ButtonStyle.Idle.ForegroundColor,
			Disabled: th.ButtonStyle.Disabled.ForegroundColor,
			Hover:    th.ButtonStyle.Hover.ForegroundColor,
			Pressed:  th.ButtonStyle.Pressed.ForegroundColor,
		}),
		widget.ButtonOpts.TextFace(*th.ButtonStyle.Font),
		widget.ButtonOpts.TextLabel(text),
	}
}

// Label returns the label options for the theme.
func (th *Theme) Label(text string) []widget.LabelOpt {
	th.LabelStyle.valid(th)
	return []widget.LabelOpt{
		widget.LabelOpts.TextOpts(
			widget.TextOpts.TextLabel(text),
			widget.TextOpts.TextColor(th.LabelStyle.Idle.ForegroundColor),
			widget.TextOpts.TextFace(*th.LabelStyle.Font),
		),
	}
}

// Text returns the text options for the theme.
func (th *Theme) Text(text string) []widget.TextOpt {
	th.LabelStyle.valid(th)
	return []widget.TextOpt{
		widget.TextOpts.TextColor(th.TextStyle.Idle.ForegroundColor),
		widget.TextOpts.TextFace(*th.TextStyle.Font),
		widget.TextOpts.TextLabel(text),
	}
}

// ProgressBar returns the progress bar options for the theme.
func (th *Theme) ProgressBar() []widget.ProgressBarOpt {
	th.ProgressBarStyle.valid(th)
	return []widget.ProgressBarOpt{
		widget.ProgressBarOpts.Images(
			&widget.ProgressBarImage{
				Idle:     th.ProgressBarStyle.Idle.Image,
				Disabled: th.ProgressBarStyle.Idle.Image,
				Hover:    th.ProgressBarStyle.Idle.Image,
			},
			&widget.ProgressBarImage{
				Idle:     th.ProgressBarStyle.Idle.Image,
				Disabled: th.ProgressBarStyle.Disabled.Image,
				Hover:    th.ProgressBarStyle.Idle.Image,
			},
		),
		widget.ProgressBarOpts.Direction(th.ProgressBarStyle.CustomData["direction"].(widget.Direction)),
		widget.ProgressBarOpts.Inverted(th.ProgressBarStyle.CustomData["inverted"].(bool)),
	}
}

// Container returns the container options for the theme.
func (th *Theme) Container() []widget.ContainerOpt {
	th.ContainerStyle.valid(th)
	return []widget.ContainerOpt{
		widget.ContainerOpts.BackgroundImage(th.ContainerStyle.Idle.Image),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	}
}

// Caret returns the caret options for the theme.
func (th *Theme) Caret() []widget.CaretOpt {
	th.CaretStyle.valid(th)
	return []widget.CaretOpt{
		widget.CaretOpts.Color(th.CaretStyle.Idle.ForegroundColor),
	}
}

// TextInput returns the text input options for the theme.
func (th *Theme) TextInput() []widget.TextInputOpt {
	th.TextInputStyle.valid(th)
	return []widget.TextInputOpt{
		widget.TextInputOpts.Color(&widget.TextInputColor{
			Idle:          th.TextInputStyle.Idle.ForegroundColor,
			Disabled:      th.TextInputStyle.Disabled.ForegroundColor,
			Caret:         th.CaretStyle.Idle.ForegroundColor,
			DisabledCaret: th.CaretStyle.Disabled.ForegroundColor,
		}),
		widget.TextInputOpts.Face(th.Font),
		widget.TextInputOpts.Image(&widget.TextInputImage{
			Idle:     th.TextInputStyle.Idle.Image,
			Disabled: th.TextInputStyle.Disabled.Image,
		}),
		widget.TextInputOpts.AllowDuplicateSubmit(th.TextInputStyle.CustomData["allow_duplicate_submit"].(bool)),
		widget.TextInputOpts.CaretOpts(th.Caret()...),
	}
}

func (th *Theme) Graphic(img *ebiten.Image) []widget.GraphicOpt {
	th.GraphicStyle.valid(th)
	return []widget.GraphicOpt{
		widget.GraphicOpts.Image(img),
	}
}

func (s *Style) valid() *Style {
	if s.Image == nil {
		s.Image = eimage.NewNineSliceColor(s.BackgroundColor)
	}
	return s
}

func (s *StyleState) valid(th *Theme) *StyleState {
	if s.Font == nil {
		s.Font = &th.Font
	}
	s.Idle.valid()
	if s.Disabled == nil {
		s.Disabled = &s.Idle
	} else {
		s.Disabled.valid()
	}
	if s.Hover == nil {
		s.Hover = &s.Idle
	} else {
		s.Hover.valid()
	}
	if s.Pressed == nil {
		s.Pressed = &s.Idle
	} else {
		s.Pressed.valid()
	}
	return s
}
