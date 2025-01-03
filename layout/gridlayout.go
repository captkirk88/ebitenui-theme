package layout

import "github.com/ebitenui/ebitenui/widget"

type GridLayoutStyle struct {
	HorizontalPosition                       widget.GridLayoutPosition
	VerticalPosition                         widget.GridLayoutPosition
	MinWidth, MinHeight, MaxWidth, MaxHeight int
}

func (l *GridLayout) Button() widget.ButtonOpt {
	return widget.ButtonOpts.WidgetOpts(
		widget.WidgetOpts.LayoutData(widget.GridLayoutData{
			HorizontalPosition: l.ButtonLayout.HorizontalPosition,
			VerticalPosition:   l.ButtonLayout.VerticalPosition,
			MaxWidth:           l.ButtonLayout.MaxWidth,
			MaxHeight:          l.ButtonLayout.MaxHeight,
		}),
		widget.WidgetOpts.MinSize(l.ButtonLayout.MinWidth, l.ButtonLayout.MinHeight))
}

func (l *GridLayout) Label() widget.TextOpt {
	return widget.TextOpts.WidgetOpts(
		widget.WidgetOpts.LayoutData(widget.GridLayoutData{
			HorizontalPosition: l.LabelLayout.HorizontalPosition,
			VerticalPosition:   l.LabelLayout.VerticalPosition,
			MaxWidth:           l.LabelLayout.MaxWidth,
			MaxHeight:          l.LabelLayout.MaxHeight,
		}),
		widget.WidgetOpts.MinSize(l.LabelLayout.MinWidth, l.LabelLayout.MinHeight))
}

func (l *GridLayout) Graphic() widget.GraphicOpt {
	return widget.GraphicOpts.WidgetOpts(
		widget.WidgetOpts.LayoutData(widget.GridLayoutData{
			HorizontalPosition: l.GraphicLayout.HorizontalPosition,
			VerticalPosition:   l.GraphicLayout.VerticalPosition,
			MaxWidth:           l.GraphicLayout.MaxWidth,
			MaxHeight:          l.GraphicLayout.MaxHeight,
		}),
		widget.WidgetOpts.MinSize(l.GraphicLayout.MinWidth, l.GraphicLayout.MinHeight))
}

func (l *GridLayout) ProgressBar() widget.ProgressBarOpt {
	return widget.ProgressBarOpts.WidgetOpts(
		widget.WidgetOpts.LayoutData(widget.GridLayoutData{
			HorizontalPosition: l.ProgressBarLayout.HorizontalPosition,
			VerticalPosition:   l.ProgressBarLayout.VerticalPosition,
			MaxWidth:           l.ProgressBarLayout.MaxWidth,
			MaxHeight:          l.ProgressBarLayout.MaxHeight,
		}),
		widget.WidgetOpts.MinSize(l.ProgressBarLayout.MinWidth, l.ProgressBarLayout.MinHeight))
}
