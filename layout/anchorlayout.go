package layout

import "github.com/ebitenui/ebitenui/widget"

type AnchorLayoutStyle struct {
	HorizontalPosition widget.AnchorLayoutPosition
	VerticalPosition   widget.AnchorLayoutPosition
	StretchHorizontal  bool
	StretchVertical    bool
	MinWidth           int
	MinHeight          int
}

func (l *AnchorLayout) Button() widget.ButtonOpt {

	return widget.ButtonOpts.WidgetOpts(
		widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: l.ButtonLayout.HorizontalPosition,
			VerticalPosition:   l.ButtonLayout.VerticalPosition,
			StretchHorizontal:  l.ButtonLayout.StretchHorizontal,
			StretchVertical:    l.ButtonLayout.StretchVertical,
		}),
		widget.WidgetOpts.MinSize(l.ButtonLayout.MinWidth, l.ButtonLayout.MinHeight))
}

func (l *AnchorLayout) Label() widget.TextOpt {
	return widget.TextOpts.WidgetOpts(
		widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: l.LabelLayout.HorizontalPosition,
			VerticalPosition:   l.LabelLayout.VerticalPosition,
			StretchHorizontal:  l.LabelLayout.StretchHorizontal,
			StretchVertical:    l.LabelLayout.StretchVertical,
		}),
		widget.WidgetOpts.MinSize(l.LabelLayout.MinWidth, l.LabelLayout.MinHeight))
}

func (l *AnchorLayout) Graphic() widget.GraphicOpt {
	return widget.GraphicOpts.WidgetOpts(
		widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: l.GraphicLayout.HorizontalPosition,
			VerticalPosition:   l.GraphicLayout.VerticalPosition,
			StretchHorizontal:  l.GraphicLayout.StretchHorizontal,
			StretchVertical:    l.GraphicLayout.StretchVertical,
		}),
		widget.WidgetOpts.MinSize(l.GraphicLayout.MinWidth, l.GraphicLayout.MinHeight))
}

func (l *AnchorLayout) ProgressBar() widget.ProgressBarOpt {
	return widget.ProgressBarOpts.WidgetOpts(
		widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: l.ProgressBarLayout.HorizontalPosition,
			VerticalPosition:   l.ProgressBarLayout.VerticalPosition,
			StretchHorizontal:  l.ProgressBarLayout.StretchHorizontal,
			StretchVertical:    l.ProgressBarLayout.StretchVertical,
		}),
		widget.WidgetOpts.MinSize(l.ProgressBarLayout.MinWidth, l.ProgressBarLayout.MinHeight))
}

func (l *AnchorLayout) Text() widget.TextOpt {
	return widget.TextOpts.WidgetOpts(
		widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: l.TextLayout.HorizontalPosition,
			VerticalPosition:   l.TextLayout.VerticalPosition,
			StretchHorizontal:  l.TextLayout.StretchHorizontal,
			StretchVertical:    l.TextLayout.StretchVertical,
		}),
		widget.WidgetOpts.MinSize(l.TextLayout.MinWidth, l.TextLayout.MinHeight))
}

func (l *AnchorLayout) TextInput() widget.TextInputOpt {
	return widget.TextInputOpts.WidgetOpts(
		widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: l.TextInputLayout.HorizontalPosition,
			VerticalPosition:   l.TextInputLayout.VerticalPosition,
			StretchHorizontal:  l.TextInputLayout.StretchHorizontal,
			StretchVertical:    l.TextInputLayout.StretchVertical,
		}),
		widget.WidgetOpts.MinSize(l.TextInputLayout.MinWidth, l.TextInputLayout.MinHeight))
}

func (l *AnchorLayout) Container() widget.ContainerOpt {
	return widget.ContainerOpts.WidgetOpts(
		widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: l.ContainerLayout.HorizontalPosition,
			VerticalPosition:   l.ContainerLayout.VerticalPosition,
			StretchHorizontal:  l.ContainerLayout.StretchHorizontal,
			StretchVertical:    l.ContainerLayout.StretchVertical,
		}),
		widget.WidgetOpts.MinSize(l.ContainerLayout.MinWidth, l.ContainerLayout.MinHeight))
}
