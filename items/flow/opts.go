package flow

import (
	"github.com/charmbracelet/lipgloss"
)

type Opt[T any] func(*Model[T])

func WithOrientation[T any](orientation Orientation) Opt[T] {
	return func(m *Model[T]) {
		m.orientation = orientation
	}
}

func WithPosition[T any](position lipgloss.Position) Opt[T] {
	return func(m *Model[T]) {
		m.position = position
	}
}

func WithStyles[T any](styles Styles) Opt[T] {
	return func(m *Model[T]) {
		m.styles = styles
	}
}

func WithWrapping[T any](wrapping bool) Opt[T] {
	return func(m *Model[T]) {
		m.wrapping = wrapping
	}
}