package stack

import "github.com/charmbracelet/lipgloss"

type Opt func(*Model)

func WithItems(items ...Item) Opt {
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
