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
	currentItemIndex    int
	groupJoinPosition   lipgloss.Position
	horizontalAlignment lipgloss.Position
	maxItemsPerGroup    int
	offset              int
	orientation         Orientation
	renderer            items.Renderer[T]
	selectedItemIndexes []int
	source              items.Source[T]
	styles              Styles
	verticalAlignment   lipgloss.Position

	visibleCount int
}

func (m Model[T]) CurrentIndex() int {
	return m.currentItemIndex
}

func (m Model[T]) GroupJoinPosition() lipgloss.Position {
	return m.groupJoinPosition
}

func (m Model[T]) HorizontalAlignment() lipgloss.Position {
	return m.horizontalAlignment
}

func (m Model[T]) Init() tea.Cmd {
	return nil
}

func (m Model[T]) MaxItemsPerGroup() int {
	return m.maxItemsPerGroup
}

func (m *Model[T]) MoveCurrentIndexDown() {
	if m.orientation == Vertical {
		if m.currentItemIndex+1 < m.source.Len() {
			m.currentItemIndex++
		}
	} else {
		// currentGroupIndex := m.groupIndex(m.currentItemIndex)
		// if currentGroupIndex >= 0 {
		// 	nextCurrentIndex := m.currentItemIndex + m.groupCounts[currentGroupIndex]
		// 	if nextCurrentIndex < m.source.Len() {
		// 		m.currentItemIndex = nextCurrentIndex
		// 	} else {
		// 		m.currentItemIndex = m.source.Len() - 1
		// 	}

		// 	currentGroupIndex = m.groupIndex(m.currentItemIndex)

		// 	if m.offset+m.visibleCount <= currentGroupIndex {
		// 		m.offset++
		// 	}
		// }
	}
}

func (m *Model[T]) MoveCurrentIndexLeft() {
	if m.orientation == Horizontal {
		if m.currentItemIndex > 0 {
			m.currentItemIndex--
		}
	} else {
		// currentGroupIndex := m.groupIndex(m.currentItemIndex)
		// if currentGroupIndex > 0 {
		// 	nextCurrentIndex := m.currentItemIndex - m.groupCounts[currentGroupIndex-1]
		// 	if nextCurrentIndex >= 0 {
		// 		m.currentItemIndex = nextCurrentIndex
		// 	} else {
		// 		m.currentItemIndex = 0
		// 	}
		// } else {
		// 	m.currentItemIndex = 0
		// }

		// currentGroupIndex = m.groupIndex(m.currentItemIndex)

		// if currentGroupIndex < m.offset {
		// 	m.offset--
		// 	if m.offset < 0 {
		// 		m.offset = 0
		// 	}
		// }
	}
}

func (m *Model[T]) MoveCurrentIndexRight() {
	if m.orientation == Horizontal {
		if m.currentItemIndex+1 < m.source.Len() {
			m.currentItemIndex++
		}
	} else {
		// currentGroupIndex := m.groupIndex(m.currentItemIndex)
		// if currentGroupIndex >= 0 {
		// 	nextCurrentIndex := m.currentItemIndex + m.groupCounts[currentGroupIndex]
		// 	if nextCurrentIndex < m.source.Len() {
		// 		m.currentItemIndex = nextCurrentIndex
		// 	} else {
		// 		m.currentItemIndex = m.source.Len() - 1
		// 	}

		// 	currentGroupIndex = m.groupIndex(m.currentItemIndex)

		// 	if m.offset+m.visibleCount <= currentGroupIndex {
		// 		m.offset++
		// 	}
		// }
	}
}

func (m *Model[T]) MoveCurrentIndexUp() {
	if m.orientation == Vertical {
		if m.currentItemIndex > 0 {
			m.currentItemIndex--
		}
	} else {
		// currentGroupIndex := m.groupIndex(m.currentItemIndex)
		// if currentGroupIndex > 0 {
		// 	nextCurrentIndex := m.currentItemIndex - m.groupCounts[currentGroupIndex-1]
		// 	if nextCurrentIndex >= 0 {
		// 		m.currentItemIndex = nextCurrentIndex
		// 	} else {
		// 		m.currentItemIndex = 0
		// 	}
		// } else {
		// 	m.currentItemIndex = 0
		// }

		// currentGroupIndex = m.groupIndex(m.currentItemIndex)

		// if currentGroupIndex < m.offset {
		// 	m.offset--
		// 	if m.offset < 0 {
		// 		m.offset = 0
		// 	}
		// }
	}
}

func (m Model[T]) Offset() int {
	return m.offset
}

func (m Model[T]) Orientation() Orientation {
	return m.orientation
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

func (m *Model[T]) SetGroupJoinPosition(position lipgloss.Position) {
	m.groupJoinPosition = position
}

func (m *Model[T]) SetHorizontalAlignment(position lipgloss.Position) {
	m.horizontalAlignment = position
}

func (m *Model[T]) SetOrientation(orientation Orientation) {
	m.orientation = orientation
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

func (m *Model[T]) SetVerticalAlignment(position lipgloss.Position) {
	m.verticalAlignment = position
}

func (m Model[T]) Source() items.Source[T] {
	return m.source
}

func (m Model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model[T]) VerticalAlignment() lipgloss.Position {
	return m.verticalAlignment
}

func (m Model[T]) View() string {
	return m.renderView(teakwood.Rectangle{})
}

func (m Model[T]) ViewWithBounds(bounds teakwood.Rectangle) string {
	return m.renderView(bounds)
}

func (m *Model[T]) renderView(bounds teakwood.Rectangle) string {
	switch m.orientation {
	case Horizontal:
		return m.renderHorizontal(bounds)
	case Vertical:
		return m.renderVertical(bounds)
	default:
		panic(fmt.Sprintf("unknown orientation: %v", m.orientation))
	}
}

func (m *Model[T]) renderHorizontal(bounds teakwood.Rectangle) string {
	groupStyle := m.styles.Group
	if groupStyle.GetAlign() != lipgloss.Left {
		groupStyle = groupStyle.Copy().Width(bounds.Width)
	}

	curWidth := groupStyle.GetHorizontalFrameSize()

	itemViewCache := make([]string, 0, max(5, m.maxItemsPerGroup))
	var groupCache []string
	var groupLengths []int

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

		if (m.maxItemsPerGroup == 0 || len(itemViewCache) < m.maxItemsPerGroup) && (bounds.Width == 0 || curWidth+w < bounds.Width || len(itemViewCache) == 0) {
			itemViewCache = append(itemViewCache, itemView)
			curWidth += w
		} else {
			group := groupStyle.Render(lipgloss.JoinHorizontal(m.groupJoinPosition, itemViewCache...))
			groupCache = append(groupCache, group)
			groupLengths = append(groupLengths, lipgloss.Height(group))

			itemViewCache = itemViewCache[:1]
			itemViewCache[0] = itemView
			curWidth = groupStyle.GetHorizontalFrameSize() + w
		}
	}

	if len(itemViewCache) > 0 {
		group := groupStyle.Render(lipgloss.JoinHorizontal(m.groupJoinPosition, itemViewCache...))
		groupCache = append(groupCache, group)
		groupLengths = append(groupLengths, lipgloss.Height(group))
	}

	availableHeight := bounds.Height
	curHeight := 0
	m.visibleCount = 0
	for i := m.offset; i < len(groupLengths); i++ {
		if curHeight+groupLengths[i] >= availableHeight {
			break
		}

		curHeight += groupLengths[i]
		m.visibleCount++
	}

	renderedView := lipgloss.JoinVertical(lipgloss.Left, groupCache[m.offset:m.offset+m.visibleCount]...)
	renderedView = lipgloss.PlaceHorizontal(bounds.Width, m.horizontalAlignment, renderedView)
	return lipgloss.PlaceVertical(bounds.Height, m.verticalAlignment, renderedView)
}

func (m *Model[T]) renderVertical(bounds teakwood.Rectangle) string {
	groupStyle := m.styles.Group

	curHeight := groupStyle.GetVerticalFrameSize()

	itemViewCache := make([]string, 0, max(5, m.maxItemsPerGroup))
	var groupCache []string
	var groupLengths []int

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

		if (m.maxItemsPerGroup == 0 || len(itemViewCache) < m.maxItemsPerGroup) && (bounds.Height == 0 || curHeight+h < bounds.Height || len(itemViewCache) == 0) {
			itemViewCache = append(itemViewCache, itemView)
			curHeight += h
		} else {
			groupItems := lipgloss.JoinVertical(m.groupJoinPosition, itemViewCache...)
			if groupStyle.GetAlign() != lipgloss.Top {
				groupItems = lipgloss.PlaceVertical(bounds.Height, groupStyle.GetAlign(), groupItems)
			}
			group := groupStyle.Render(groupItems)
			groupCache = append(groupCache, group)
			groupLengths = append(groupLengths, lipgloss.Width(group))

			itemViewCache = itemViewCache[:1]
			itemViewCache[0] = itemView
			curHeight = groupStyle.GetVerticalFrameSize() + h
		}
	}

	if len(itemViewCache) > 0 {
		groupItems := lipgloss.JoinVertical(m.groupJoinPosition, itemViewCache...)
		if groupStyle.GetAlign() != lipgloss.Top {
			groupItems = lipgloss.PlaceVertical(bounds.Height, groupStyle.GetAlign(), groupItems)
		}
		group := groupStyle.Render(groupItems)
		groupCache = append(groupCache, group)
		groupLengths = append(groupLengths, lipgloss.Width(group))
	}

	availableWidth := bounds.Width
	curWidth := 0
	m.visibleCount = 0
	for i := m.offset; i < len(groupLengths); i++ {
		if curWidth+groupLengths[i] >= availableWidth {
			break
		}

		curWidth += groupLengths[i]
		m.visibleCount++
	}

	renderedView := lipgloss.JoinHorizontal(lipgloss.Top, groupCache[m.offset:m.offset+m.visibleCount]...)
	renderedView = lipgloss.PlaceHorizontal(bounds.Width, m.horizontalAlignment, renderedView)
	return lipgloss.PlaceVertical(bounds.Height, m.verticalAlignment, renderedView)
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
