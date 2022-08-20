package adapter

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/craiggwilson/teacomps"
)

type Opt[T any] func(*Model[T])

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

func WithUpdateBounds[T any](f func(T, teacomps.Rectangle) T) Opt[T] {
	return func(m *Model[T]) {
		m.updateBounds = f
	}
}

func WithView[T any](f func(T) string) Opt[T] {
	return func(m *Model[T]) {
		m.view = f
	}
}
