package flow

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Opt func(*Model)

func WithItems(items ...tea.Model) Opt {
	return func(m *Model) {
		m.items = items
	}
}

func WithOrientation(orientation Orientation) Opt {
	return func(m *Model) {
		m.orientation = orientation
	}
}

func WithPosition(position lipgloss.Position) Opt {
	return func(m *Model) {
		m.position = position
	}
}

func WithStyles(styles Styles) Opt {
	return func(m *Model) {
		m.styles = styles
	}
}

func WithWrapping(wrapping bool) Opt {
	return func(m *Model) {
		m.wrapping = wrapping
	}
}
