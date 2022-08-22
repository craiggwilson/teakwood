package tabs

import tea "github.com/charmbracelet/bubbletea"

type Opt func(*Model)

func WithItems(items ...tea.Model) Opt {
	return func(m *Model) {
		m.items = items
	}
}

func WithStyles(styles Styles) Opt {
	return func(m *Model) {
		m.styles = styles
	}
}
