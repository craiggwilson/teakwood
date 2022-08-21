package frame

import "github.com/charmbracelet/lipgloss"

type Opt func(*Model)

func WithStyle(style lipgloss.Style) Opt {
	return func(m *Model) {
		m.style = style
	}
}
