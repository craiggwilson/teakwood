package util

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
)

type Widget struct {
	Bounds  teakwood.Rectangle
	Fill    bool
	Kind    string
	Name    string
	Styler  teakwood.Styler
	Visible bool
}

// ContentBounds gets the bounds for the content area. This is the size
// with the margins, padding, or borders.
func (w *Widget) ContentBounds() teakwood.Rectangle {
	s := w.Style(w.Bounds.Size())
	offsets := teakwood.OffsetsFromStyle(s)
	return w.Bounds.Offset(offsets)
}

func (w *Widget) Render(content string) string {
	if !w.Visible {
		return ""
	}

	return w.Style(w.Bounds.Size()).Render(content)
}

// Style gets the style that is applied to this element.
func (w *Widget) Style(size teakwood.Size) lipgloss.Style {
	s, _ := w.Styler.Style(w.Kind)
	if w.Name != "" {
		if nameStyle, ok := w.Styler.Style("#" + w.Name); ok {
			s = nameStyle.Inherit(s)
		}
	}

	if !size.IsZero() {
		s = s.MaxWidth(size.Width).MaxHeight(size.Height)
	}

	if w.Fill {
		s = s.Width(size.Width - s.GetHorizontalFrameSize()).
			Height(size.Height - s.GetVerticalFrameSize())
	}

	return s
}
