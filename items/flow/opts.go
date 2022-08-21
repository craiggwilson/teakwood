package flow

import "github.com/charmbracelet/lipgloss"

type Opt func(*Model)

func WithStyles(styles Styles) Opt {
	return func(m *Model) {
		m.styles = styles
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

func WithWrapping(wrapping bool) Opt {
	return func(m *Model) {
		m.wrapping = wrapping
	}
}
