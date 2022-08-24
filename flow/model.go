package flow

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
)

func New(opts ...Opt) Model {
	var m Model

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

type Model struct {
	bounds      teakwood.Rectangle
	items       []tea.Model
	styles      Styles
	orientation Orientation
	position    lipgloss.Position
	wrapping    bool

	itemViews []string
}

func (m *Model) AddItem(page tea.Model) {
	m.items = append(m.items, page)
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) InsertItem(index int, page tea.Model) {
	m.items = append(m.items[:index+1], m.items[index:]...)
	m.items[index] = page
}

func (m Model) Items() []tea.Model {
	return m.items
}

func (m Model) Len() int {
	return len(m.items)
}

func (m Model) Orientation() Orientation {
	return m.orientation
}

func (m Model) Position() lipgloss.Position {
	return m.position
}

func (m *Model) RemoveItem(index int) {
	m.items = append(m.items[:index], m.items[index+1:]...)
}

func (m *Model) SetOrientation(orientation Orientation) {
	m.orientation = orientation
}

func (m *Model) SetPosition(position lipgloss.Position) {
	m.position = position
}

func (m *Model) SetItems(items ...tea.Model) {
	m.items = items
}

func (m *Model) SetWrapping(wrapping bool) {
	m.wrapping = wrapping
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	for i := 0; i < len(m.items); i++ {
		m.items[i], cmd = m.items[i].Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) UpdateBounds(bounds teakwood.Rectangle) teakwood.Visual {
	m.bounds = bounds
	return m
}

func (m Model) View() string {
	switch m.orientation {
	case Vertical:
		return m.renderVertical()
	case Horizontal:
		return m.renderHorizontal()
	default:
		panic(fmt.Sprintf("unknown orientation: %v", m.orientation))
	}
}

func (m Model) Wrapping() bool {
	return m.wrapping
}

func (m *Model) renderHorizontal() string {
	if cap(m.itemViews) != cap(m.items) {
		m.itemViews = make([]string, 0, len(m.items))
	}
	m.itemViews = m.itemViews[:0]

	groupStyle := m.styles.Group.Copy().Width(m.bounds.Width)

	var rows []string

	curWidth := groupStyle.GetHorizontalFrameSize()
	for i := range m.items {
		view := m.styles.Item.Render(m.items[i].View())
		w := lipgloss.Width(view)

		if !m.wrapping || m.bounds.Width == 0 || curWidth+w < m.bounds.Width {
			m.itemViews = append(m.itemViews, view)
			curWidth += w
		} else {
			rowView := groupStyle.Render(lipgloss.JoinHorizontal(m.position, m.itemViews...))
			rows = append(rows, rowView)

			m.itemViews = m.itemViews[:1]
			m.itemViews[0] = view
			curWidth = groupStyle.GetHorizontalFrameSize() + w
		}
	}

	if len(m.itemViews) > 0 {
		rows = append(rows, groupStyle.Render(lipgloss.JoinHorizontal(m.position, m.itemViews...)))
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func (m *Model) renderVertical() string {
	if cap(m.itemViews) != cap(m.items) {
		m.itemViews = make([]string, 0, len(m.items))
	}
	m.itemViews = m.itemViews[:0]

	groupStyle := m.styles.Group

	var cols []string

	curHeight := groupStyle.GetVerticalFrameSize()
	for i := 0; i < len(m.items); i++ {
		view := m.styles.Item.Render(m.items[i].View())
		h := lipgloss.Height(view)

		if !m.wrapping || m.bounds.Height == 0 || curHeight+h < m.bounds.Height {
			m.itemViews = append(m.itemViews, view)
			curHeight += h
		} else {
			colView := groupStyle.Render(lipgloss.PlaceVertical(m.bounds.Height, groupStyle.GetAlign(), lipgloss.JoinVertical(m.position, m.itemViews...)))
			cols = append(cols, colView)

			m.itemViews = m.itemViews[:1]
			m.itemViews[0] = view
			curHeight = groupStyle.GetVerticalFrameSize() + h
		}
	}

	if len(m.itemViews) > 0 {
		colView := groupStyle.Render(lipgloss.PlaceVertical(m.bounds.Height, groupStyle.GetAlign(), lipgloss.JoinVertical(m.position, m.itemViews...)))
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
