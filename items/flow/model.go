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
		source:           source,
		renderer:         renderer,
		currentItemIndex: -1,
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
	bounds              teakwood.Rectangle
	currentItemIndex    int
	groupOffset         int
	orientation         Orientation
	position            lipgloss.Position
	renderer            items.Renderer[T]
	selectedItemIndexes []int
	source              items.Source[T]
	styles              Styles
	wrapping            bool

	groupCache        []string
	groupCounts       []int
	groupDisplayCount int
	groupLengths      []int
	itemViewCache     []string
	renderedView      string
}

func (m Model[T]) CurrentIndex() int {
	return m.currentItemIndex
}

func (m Model[T]) GroupOffset() int {
	return m.groupOffset
}

func (m Model[T]) Init() tea.Cmd {
	return nil
}

func (m *Model[T]) MoveCurrentIndexDown() {
	if m.orientation == Vertical {
		if m.currentItemIndex+1 < m.source.Len() {
			m.currentItemIndex++
		}
	} else {
		currentGroupIndex := m.groupIndex(m.currentItemIndex)
		if currentGroupIndex >= 0 {
			nextCurrentIndex := m.currentItemIndex + m.groupCounts[currentGroupIndex]
			if nextCurrentIndex < m.source.Len() {
				m.currentItemIndex = nextCurrentIndex
			} else {
				m.currentItemIndex = m.source.Len() - 1
			}

			currentGroupIndex = m.groupIndex(m.currentItemIndex)

			if m.groupOffset+m.groupDisplayCount <= currentGroupIndex {
				m.groupOffset++
			}
		}
	}
}

func (m *Model[T]) MoveCurrentIndexLeft() {
	if m.orientation == Horizontal {
		if m.currentItemIndex > 0 {
			m.currentItemIndex--
		}
	} else {
		currentGroupIndex := m.groupIndex(m.currentItemIndex)
		if currentGroupIndex > 0 {
			nextCurrentIndex := m.currentItemIndex - m.groupCounts[currentGroupIndex-1]
			if nextCurrentIndex >= 0 {
				m.currentItemIndex = nextCurrentIndex
			} else {
				m.currentItemIndex = 0
			}
		} else {
			m.currentItemIndex = 0
		}

		currentGroupIndex = m.groupIndex(m.currentItemIndex)

		if currentGroupIndex < m.groupOffset {
			m.groupOffset--
			if m.groupOffset < 0 {
				m.groupOffset = 0
			}
		}
	}
}

func (m *Model[T]) MoveCurrentIndexRight() {
	if m.orientation == Horizontal {
		if m.currentItemIndex+1 < m.source.Len() {
			m.currentItemIndex++
		}
	} else {
		currentGroupIndex := m.groupIndex(m.currentItemIndex)
		if currentGroupIndex >= 0 {
			nextCurrentIndex := m.currentItemIndex + m.groupCounts[currentGroupIndex]
			if nextCurrentIndex < m.source.Len() {
				m.currentItemIndex = nextCurrentIndex
			} else {
				m.currentItemIndex = m.source.Len() - 1
			}

			currentGroupIndex = m.groupIndex(m.currentItemIndex)

			if m.groupOffset+m.groupDisplayCount <= currentGroupIndex {
				m.groupOffset++
			}
		}
	}
}

func (m *Model[T]) MoveCurrentIndexUp() {
	if m.orientation == Vertical {
		if m.currentItemIndex > 0 {
			m.currentItemIndex--
		}
	} else {
		currentGroupIndex := m.groupIndex(m.currentItemIndex)
		if currentGroupIndex > 0 {
			nextCurrentIndex := m.currentItemIndex - m.groupCounts[currentGroupIndex-1]
			if nextCurrentIndex >= 0 {
				m.currentItemIndex = nextCurrentIndex
			} else {
				m.currentItemIndex = 0
			}
		} else {
			m.currentItemIndex = 0
		}

		currentGroupIndex = m.groupIndex(m.currentItemIndex)

		if currentGroupIndex < m.groupOffset {
			m.groupOffset--
			if m.groupOffset < 0 {
				m.groupOffset = 0
			}
		}
	}
}

func (m Model[T]) NumGroups() int {
	return len(m.groupCounts)
}

func (m Model[T]) Orientation() Orientation {
	return m.orientation
}

func (m Model[T]) Position() lipgloss.Position {
	return m.position
}

func (m Model[T]) SelectedIndexes() []int {
	return m.selectedItemIndexes
}

func (m *Model[T]) SetCurrentIndex(currentIndex int) {
	switch {
	case currentIndex < 0 || currentIndex >= m.source.Len():
		m.currentItemIndex = -1
	default:
		m.currentItemIndex = currentIndex
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
	m.selectedItemIndexes = selectedIndexes
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

func (m *Model[T]) groupIndex(itemIndex int) int {
	currentTotal := 0
	for i, gc := range m.groupCounts {
		currentTotal += gc
		if itemIndex < currentTotal {
			return i
		}
	}

	return -1
}

func (m *Model[T]) renderView() {
	m.itemViewCache = m.itemViewCache[:0]
	m.groupCounts = m.groupCounts[:0]
	m.groupCache = m.groupCache[:0]
	m.groupLengths = m.groupLengths[:0]

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

	curWidth := groupStyle.GetHorizontalFrameSize()

	for i := 0; i < m.source.Len(); i++ {
		itemView := m.renderer.Render(m.source.Item(i))
		if m.currentItemIndex == i {
			itemView = m.styles.CurrentItem.Render(itemView)
		} else if contains(m.selectedItemIndexes, i) {
			itemView = m.styles.SelectedItem.Render(itemView)
		} else {
			itemView = m.styles.Item.Render(itemView)
		}

		w := lipgloss.Width(itemView)

		if !m.wrapping || m.bounds.Width == 0 || curWidth+w < m.bounds.Width || len(m.itemViewCache) == 0 {
			m.itemViewCache = append(m.itemViewCache, itemView)
			curWidth += w
		} else {
			group := groupStyle.Render(lipgloss.JoinHorizontal(m.position, m.itemViewCache...))
			m.groupCache = append(m.groupCache, group)
			m.groupCounts = append(m.groupCounts, len(m.itemViewCache))
			m.groupLengths = append(m.groupLengths, lipgloss.Height(group))

			m.itemViewCache = m.itemViewCache[:1]
			m.itemViewCache[0] = itemView
			curWidth = groupStyle.GetHorizontalFrameSize() + w
		}
	}

	if len(m.itemViewCache) > 0 {
		group := groupStyle.Render(lipgloss.JoinHorizontal(m.position, m.itemViewCache...))
		m.groupCache = append(m.groupCache, group)
		m.groupCounts = append(m.groupCounts, len(m.itemViewCache))
		m.groupLengths = append(m.groupLengths, lipgloss.Height(group))
	}

	availableHeight := m.bounds.Height
	curHeight := 0
	m.groupDisplayCount = 0
	for i := m.groupOffset; i < len(m.groupLengths); i++ {
		if curHeight+m.groupLengths[i] >= availableHeight {
			break
		}

		curHeight += m.groupLengths[i]
		m.groupDisplayCount++
	}

	m.renderedView = lipgloss.JoinVertical(lipgloss.Left, m.groupCache[m.groupOffset:m.groupOffset+m.groupDisplayCount]...)
}

func (m *Model[T]) renderVertical() {
	groupStyle := m.styles.Group

	curHeight := groupStyle.GetVerticalFrameSize()

	for i := 0; i < m.source.Len(); i++ {
		itemView := m.renderer.Render(m.source.Item(i))
		if m.currentItemIndex == i {
			itemView = m.styles.CurrentItem.Render(itemView)
		} else if contains(m.selectedItemIndexes, i) {
			itemView = m.styles.SelectedItem.Render(itemView)
		} else {
			itemView = m.styles.Item.Render(itemView)
		}

		h := lipgloss.Height(itemView)

		if !m.wrapping || m.bounds.Height == 0 || curHeight+h < m.bounds.Height {
			m.itemViewCache = append(m.itemViewCache, itemView)
			curHeight += h
		} else {
			group := groupStyle.Render(lipgloss.PlaceVertical(m.bounds.Height, groupStyle.GetAlign(), lipgloss.JoinVertical(m.position, m.itemViewCache...)))
			m.groupCache = append(m.groupCache, group)
			m.groupCounts = append(m.groupCounts, len(m.itemViewCache))
			m.groupLengths = append(m.groupLengths, lipgloss.Width(group))

			m.itemViewCache = m.itemViewCache[:1]
			m.itemViewCache[0] = itemView
			curHeight = groupStyle.GetVerticalFrameSize() + h
		}
	}

	if len(m.itemViewCache) > 0 {
		group := groupStyle.Render(lipgloss.PlaceVertical(m.bounds.Height, groupStyle.GetAlign(), lipgloss.JoinVertical(m.position, m.itemViewCache...)))
		m.groupCache = append(m.groupCache, group)
		m.groupCounts = append(m.groupCounts, len(m.itemViewCache))
		m.groupLengths = append(m.groupLengths, lipgloss.Width(group))
	}

	availableWidth := m.bounds.Width
	curWidth := 0
	m.groupDisplayCount = 0
	for i := m.groupOffset; i < len(m.groupLengths); i++ {
		if curWidth+m.groupLengths[i] >= availableWidth {
			break
		}

		curWidth += m.groupLengths[i]
		m.groupDisplayCount++
	}

	m.renderedView = lipgloss.JoinHorizontal(lipgloss.Top, m.groupCache[m.groupOffset:m.groupOffset+m.groupDisplayCount]...)
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
