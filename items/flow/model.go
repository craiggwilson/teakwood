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
		source:       source,
		renderer:     renderer,
		currentIndex: -1,
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
	bounds          teakwood.Rectangle
	orientation     Orientation
	position        lipgloss.Position
	renderer        items.Renderer[T]
	currentIndex    int
	selectedIndexes []int
	source          items.Source[T]
	styles          Styles
	wrapping        bool

	viewCache    []string
	renderedView string
	groupCounts  []int
}

func (m Model[T]) CurrentIndex() int {
	return m.currentIndex
}

func (m Model[T]) Init() tea.Cmd {
	return nil
}

func (m *Model[T]) MoveCurrentIndexDown() {
	if m.orientation == Vertical {
		if m.currentIndex+1 < m.source.Len() {
			m.currentIndex++
		}
	} else {
		currentGroupIndex := m.currentGroupIndex()
		if currentGroupIndex >= 0 {
			nextCurrentIndex := m.currentIndex + m.groupCounts[currentGroupIndex]
			if nextCurrentIndex < m.source.Len() {
				m.currentIndex = nextCurrentIndex
			} else {
				m.currentIndex = m.source.Len() - 1
			}
		}
	}
}

func (m *Model[T]) MoveCurrentIndexLeft() {
	if m.orientation == Horizontal {
		if m.currentIndex > 0 {
			m.currentIndex--
		}
	} else {
		currentGroupIndex := m.currentGroupIndex()
		if currentGroupIndex > 0 {
			nextCurrentIndex := m.currentIndex - m.groupCounts[currentGroupIndex-1]
			if nextCurrentIndex >= 0 {
				m.currentIndex = nextCurrentIndex
			} else {
				m.currentIndex = 0
			}
		} else {
			m.currentIndex = 0
		}
	}
}

func (m *Model[T]) MoveCurrentIndexRight() {
	if m.orientation == Horizontal {
		if m.currentIndex+1 < m.source.Len() {
			m.currentIndex++
		}
	} else {
		currentGroupIndex := m.currentGroupIndex()
		if currentGroupIndex >= 0 {
			nextCurrentIndex := m.currentIndex + m.groupCounts[currentGroupIndex]
			if nextCurrentIndex < m.source.Len() {
				m.currentIndex = nextCurrentIndex
			} else {
				m.currentIndex = m.source.Len() - 1
			}
		}
	}
}

func (m *Model[T]) MoveCurrentIndexUp() {
	if m.orientation == Vertical {
		if m.currentIndex > 0 {
			m.currentIndex--
		}
	} else {
		currentGroupIndex := m.currentGroupIndex()
		if currentGroupIndex > 0 {
			nextCurrentIndex := m.currentIndex - m.groupCounts[currentGroupIndex-1]
			if nextCurrentIndex >= 0 {
				m.currentIndex = nextCurrentIndex
			} else {
				m.currentIndex = 0
			}
		} else {
			m.currentIndex = 0
		}
	}
}

func (m Model[T]) Orientation() Orientation {
	return m.orientation
}

func (m Model[T]) Position() lipgloss.Position {
	return m.position
}

func (m Model[T]) SelectedIndexes() []int {
	return m.selectedIndexes
}

func (m *Model[T]) SetCurrentIndex(currentIndex int) {
	switch {
	case currentIndex < 0 || currentIndex >= m.source.Len():
		m.currentIndex = -1
	default:
		m.currentIndex = currentIndex
	}
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

func (m *Model[T]) SetSelectedIndexes(selectedIndexes ...int) {
	m.selectedIndexes = selectedIndexes
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
	m.renderView()
	return m, nil
}

func (m Model[T]) UpdateBounds(bounds teakwood.Rectangle) teakwood.Visual {
	m.bounds = bounds
	return m
}

func (m Model[T]) View() string {
	return m.renderedView
}

func (m Model[T]) Wrapping() bool {
	return m.wrapping
}

func (m *Model[T]) currentGroupIndex() int {
	currentTotal := 0
	for i, gc := range m.groupCounts {
		currentTotal += gc
		if m.currentIndex < currentTotal {
			return i
		}
	}

	return -1
}

func (m *Model[T]) renderView() {
	if cap(m.viewCache) != m.source.Len() {
		m.viewCache = make([]string, 0, m.source.Len())
	}
	m.viewCache = m.viewCache[:0]
	m.groupCounts = m.groupCounts[:0]

	switch m.orientation {
	case Vertical:
		m.renderVertical()
	case Horizontal:
		m.renderHorizontal()
	default:
		panic(fmt.Sprintf("unknown orientation: %v", m.orientation))
	}
}

func (m *Model[T]) renderHorizontal() {
	groupStyle := m.styles.Group.Copy().Width(m.bounds.Width)

	var rows []string

	curWidth := groupStyle.GetHorizontalFrameSize()

	for i := 0; i < m.source.Len(); i++ {
		view := m.renderer.Render(m.source.Item(i))
		if m.currentIndex == i {
			view = m.styles.CurrentItem.Render(view)
		} else if contains(m.selectedIndexes, i) {
			view = m.styles.SelectedItem.Render(view)
		} else {
			view = m.styles.Item.Render(view)
		}

		w := lipgloss.Width(view)

		if !m.wrapping || m.bounds.Width == 0 || curWidth+w < m.bounds.Width {
			m.viewCache = append(m.viewCache, view)
			curWidth += w
		} else {
			m.groupCounts = append(m.groupCounts, len(m.viewCache))
			rowView := groupStyle.Render(lipgloss.JoinHorizontal(m.position, m.viewCache...))
			rows = append(rows, rowView)

			m.viewCache = m.viewCache[:1]
			m.viewCache[0] = view
			curWidth = groupStyle.GetHorizontalFrameSize() + w
		}
	}

	if len(m.viewCache) > 0 {
		m.groupCounts = append(m.groupCounts, len(m.viewCache))
		rows = append(rows, groupStyle.Render(lipgloss.JoinHorizontal(m.position, m.viewCache...)))
	}

	m.renderedView = lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func (m *Model[T]) renderVertical() {
	groupStyle := m.styles.Group

	var cols []string

	curHeight := groupStyle.GetVerticalFrameSize()
	for i := 0; i < m.source.Len(); i++ {
		view := m.renderer.Render(m.source.Item(i))
		if m.currentIndex == i {
			view = m.styles.CurrentItem.Render(view)
		} else if contains(m.selectedIndexes, i) {
			view = m.styles.SelectedItem.Render(view)
		} else {
			view = m.styles.Item.Render(view)
		}

		h := lipgloss.Height(view)

		if !m.wrapping || m.bounds.Height == 0 || curHeight+h < m.bounds.Height {
			m.viewCache = append(m.viewCache, view)
			curHeight += h
		} else {
			m.groupCounts = append(m.groupCounts, len(m.viewCache))
			colView := groupStyle.Render(lipgloss.PlaceVertical(m.bounds.Height, groupStyle.GetAlign(), lipgloss.JoinVertical(m.position, m.viewCache...)))
			cols = append(cols, colView)

			m.viewCache = m.viewCache[:1]
			m.viewCache[0] = view
			curHeight = groupStyle.GetVerticalFrameSize() + h
		}
	}

	if len(m.viewCache) > 0 {
		m.groupCounts = append(m.groupCounts, len(m.viewCache))
		colView := groupStyle.Render(lipgloss.PlaceVertical(m.bounds.Height, groupStyle.GetAlign(), lipgloss.JoinVertical(m.position, m.viewCache...)))
		cols = append(cols, colView)
	}

	m.renderedView = lipgloss.JoinHorizontal(lipgloss.Top, cols...)
}

func contains(is []int, test int) bool {
	for _, i := range is {
		if i == test {
			return true
		}
	}

	return false
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
