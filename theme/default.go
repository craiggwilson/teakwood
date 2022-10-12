package theme

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
)

func Default() teakwood.Styler {
	return &theme{
		styles: map[string]lipgloss.Style{
			"body": lipgloss.NewStyle(),
		},
	}
}

type theme struct {
	styles map[string]lipgloss.Style
}

func (t *theme) Style(name string) (lipgloss.Style, bool) {
	s, ok := t.styles[name]
	return s, ok
}
