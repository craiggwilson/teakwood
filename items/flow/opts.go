package flow

import (
	"github.com/charmbracelet/lipgloss"
)

type Opt[T any] func(*Model[T])

func WithCurrentIndex[T any](currentIndex int) Opt[T] {
	return func(m *Model[T]) {
		m.currentItemIndex = currentIndex
	}
}

func WithGroupJoinPosition[T any](position lipgloss.Position) Opt[T] {
	return func(m *Model[T]) {
		m.groupJoinPosition = position
	}
}

func WithHorizontalAlignment[T any](position lipgloss.Position) Opt[T] {
	return func(m *Model[T]) {
		m.horizontalAlignment = position
	}
}

func WithMaxItemsPerGroup[T any](max int) Opt[T] {
	return func(m *Model[T]) {
		m.maxItemsPerGroup = max
	}
}

func WithOrientation[T any](orientation Orientation) Opt[T] {
	return func(m *Model[T]) {
		m.orientation = orientation
	}
}

func WithSelectedIndexes[T any](selectedIndexes ...int) Opt[T] {
	return func(m *Model[T]) {
		m.selectedItemIndexes = selectedIndexes
	}
}

func WithStyles[T any](styles Styles) Opt[T] {
	return func(m *Model[T]) {
		m.styles = styles
	}
}

func WithVerticalAlignment[T any](position lipgloss.Position) Opt[T] {
	return func(m *Model[T]) {
		m.verticalAlignment = position
	}
}
