package adapter

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/craiggwilson/teakwood"
)

func New[T any](adaptee T, opts ...Opt[T]) Model[T] {
	m := Model[T]{
		adaptee: adaptee,
	}

	for _, opt := range opts {
		opt(&m)
	}

	return m
}

type Model[T any] struct {
	adaptee T

	init         func(T) tea.Cmd
	update       func(T, tea.Msg) (T, tea.Cmd)
	updateBounds func(T, teakwood.Rectangle) T
	view         func(T) string
}

func (m Model[T]) Adaptee() T {
	return m.adaptee
}

func (m Model[T]) Init() tea.Cmd {
	if m.init != nil {
		return m.init(m.adaptee)
	}

	if initer, ok := any(m.adaptee).(initer); ok {
		return initer.Init()
	}

	return nil
}

func (m Model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.update != nil {
		m.adaptee, cmd = m.update(m.adaptee, msg)
	} else if updater, ok := any(m.adaptee).(updater[T]); ok {
		m.adaptee, cmd = updater.Update(msg)
	}

	return m, cmd
}

func (m Model[T]) UpdateBounds(bounds teakwood.Rectangle) teakwood.Visual {
	if m.updateBounds != nil {
		m.adaptee = m.updateBounds(m.adaptee, bounds)
	} else if bu, ok := any(m.adaptee).(boundsUpdater); ok {
		newAdaptee := bu.UpdateBounds(bounds)
		m.adaptee = newAdaptee.(T)
	}

	return m
}

func (m *Model[T]) UpdateAdaptee(v func(T) T) {
	m.adaptee = v(m.adaptee)
}

func (m Model[T]) View() string {
	if m.view != nil {
		return m.view(m.adaptee)
	}

	if viewer, ok := any(m.adaptee).(viewer); ok {
		return viewer.View()
	}

	return ""
}

type boundsUpdater interface {
	UpdateBounds(teakwood.Rectangle) teakwood.Visual
}

type initer interface {
	Init() tea.Cmd
}

type updater[T any] interface {
	Update(tea.Msg) (T, tea.Cmd)
}

type viewer interface {
	View() string
}
