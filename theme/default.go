package theme

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
	"github.com/craiggwilson/teakwood/frame"
	"github.com/craiggwilson/teakwood/label"
)

func Default() teakwood.Styler {
	return defaultTheme()
}

func defaultTheme() *theme {
	return &theme{
		styles: map[string]lipgloss.Style{
			"body":         lipgloss.NewStyle(),
			frame.StyleKey: lipgloss.NewStyle().AlignHorizontal(lipgloss.Center),
			label.StyleKey: lipgloss.NewStyle(),
		},
	}
}

type theme struct {
	styles map[string]lipgloss.Style
}

func (t *theme) Style(key string) (lipgloss.Style, bool) {
	if s, ok := t.styles[key]; ok {
		return s.Copy(), true
	}
	return lipgloss.Style{}, false
}
