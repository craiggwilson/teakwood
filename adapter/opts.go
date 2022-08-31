package adapter

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/craiggwilson/teakwood"
)

type Opt[T any] func(*Model[T])

func WithBoundsViewer[T any](f func(T, teakwood.Rectangle) string) Opt[T] {
	return func(m *Model[T]) {
		m.viewWithBounds = f
	}
}

func WithInit[T any](f func(T) tea.Cmd) Opt[T] {
	return func(m *Model[T]) {
		m.init = f
	}
}

func WithUpdate[T any](f func(T, tea.Msg) (T, tea.Cmd)) Opt[T] {
	return func(m *Model[T]) {
		m.update = f
	}
}

func WithView[T any](f func(T) string) Opt[T] {
	return func(m *Model[T]) {
		m.view = f
	}
}
