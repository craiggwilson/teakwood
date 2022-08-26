package flow

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
	"github.com/craiggwilson/teakwood/items"
)

func New[T any](source items.Source[T], renderer items.Renderer[T], opts ...Opt[T]) Model[T] {
	m := Model[T]{
		source:   source,
		renderer: renderer,
	}

	for _, opt := range opts {
		opt(&m)
	}

	return m
}

type Orientation int

const (
	Horizontal Orientation = iota
	Vertical
)

type Model[T any] struct {
	bounds      teakwood.Rectangle
	orientation Orientation
	position    lipgloss.Position
	renderer    items.Renderer[T]
	source      items.Source[T]
	styles      Styles
	wrapping    bool

	viewCache []string
}

func (m Model[T]) Init() tea.Cmd {
	return nil
}

func (m Model[T]) Orientation() Orientation {
	return m.orientation
}

func (m Model[T]) Position() lipgloss.Position {
	return m.position
}

func (m *Model[T]) SetOrientation(orientation Orientation) {
	m.orientation = orientation
}

func (m *Model[T]) SetPosition(position lipgloss.Position) {
	m.position = position
}

func (m *Model[T]) SetRenderer(renderer items.Renderer[T]) {
	m.renderer = renderer
}

func (m *Model[T]) SetSource(source items.Source[T]) {
	m.source = source
}

func (m *Model[T]) SetWrapping(wrapping bool) {
	m.wrapping = wrapping
}

func (m Model[T]) Source() items.Source[T] {
	return m.source
}

func (m Model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model[T]) UpdateBounds(bounds teakwood.Rectangle) teakwood.Visual {
	m.bounds = bounds
	return m
}

func (m Model[T]) View() string {
	if cap(m.viewCache) != m.source.Len() {
		m.viewCache = make([]string, 0, m.source.Len())
	}
	m.viewCache = m.viewCache[:0]

	switch m.orientation {
	case Vertical:
		return m.renderVertical()
	case Horizontal:
		return m.renderHorizontal()
	default:
		panic(fmt.Sprintf("unknown orientation: %v", m.orientation))
	}
}

func (m Model[T]) Wrapping() bool {
	return m.wrapping
}

func (m *Model[T]) renderHorizontal() string {
	groupStyle := m.styles.Group.Copy().Width(m.bounds.Width)

	var rows []string

	curWidth := groupStyle.GetHorizontalFrameSize()
	for i := 0; i < m.source.Len(); i++ {
		view := m.styles.Item.Render(m.renderer.Render(m.source.Item(i)))
		w := lipgloss.Width(view)

		if !m.wrapping || m.bounds.Width == 0 || curWidth+w < m.bounds.Width {
			m.viewCache = append(m.viewCache, view)
			curWidth += w
		} else {
			rowView := groupStyle.Render(lipgloss.JoinHorizontal(m.position, m.viewCache...))
			rows = append(rows, rowView)

			m.viewCache = m.viewCache[:1]
			m.viewCache[0] = view
			curWidth = groupStyle.GetHorizontalFrameSize() + w
		}
	}

	if len(m.viewCache) > 0 {
		rows = append(rows, groupStyle.Render(lipgloss.JoinHorizontal(m.position, m.viewCache...)))
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func (m *Model[T]) renderVertical() string {
	groupStyle := m.styles.Group

	var cols []string

	curHeight := groupStyle.GetVerticalFrameSize()
	for i := 0; i < m.source.Len(); i++ {
		view := m.styles.Item.Render(m.renderer.Render(m.source.Item(i)))
		h := lipgloss.Height(view)

		if !m.wrapping || m.bounds.Height == 0 || curHeight+h < m.bounds.Height {
			m.viewCache = append(m.viewCache, view)
			curHeight += h
		} else {
			colView := groupStyle.Render(lipgloss.PlaceVertical(m.bounds.Height, groupStyle.GetAlign(), lipgloss.JoinVertical(m.position, m.viewCache...)))
			cols = append(cols, colView)

			m.viewCache = m.viewCache[:1]
			m.viewCache[0] = view
			curHeight = groupStyle.GetVerticalFrameSize() + h
		}
	}

	if len(m.viewCache) > 0 {
		colView := groupStyle.Render(lipgloss.PlaceVertical(m.bounds.Height, groupStyle.GetAlign(), lipgloss.JoinVertical(m.position, m.viewCache...)))
		cols = append(cols, colView)
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, cols...)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
