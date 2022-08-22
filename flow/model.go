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
	Vertical Orientation = iota
	Horizontal
)

type Model struct {
	bounds      teakwood.Rectangle
	items       []tea.Model
	styles      Styles
	orientation Orientation
	position    lipgloss.Position
	wrapping    bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Orientation() Orientation {
	return m.orientation
}

func (m Model) Position() lipgloss.Position {
	return m.position
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
		return m.viewVertical()
	case Horizontal:
		return m.viewHorizontal()
	default:
		panic(fmt.Sprintf("unknown orientation: %v", m.orientation))
	}
}

func (m Model) Wrapping() bool {
	return m.wrapping
}

func (m Model) viewHorizontal() string {
	groupStyle := m.styles.Group.Copy().Width(m.bounds.Width)

	var rows []string
	views := make([]string, 0, len(m.items))
	curWidth := 0
	for i := 0; i < len(m.items); i++ {
		view := m.styles.Item.Render(m.items[i].View())
		viewWidth := lipgloss.Width(view)
		if !m.wrapping || m.bounds.Width == 0 || curWidth+viewWidth < m.bounds.Width {
			views = append(views, view)
			curWidth += viewWidth
		} else {
			rows = append(rows, groupStyle.Render(lipgloss.JoinHorizontal(m.position, views...)))
			views = views[:1]
			views[0] = view
			curWidth = viewWidth
		}
	}

	if len(views) > 0 {
		rows = append(rows, groupStyle.Render(lipgloss.JoinHorizontal(m.position, views...)))
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func (m Model) viewVertical() string {
	groupStyle := m.styles.Group

	var cols []string
	views := make([]string, 0, len(m.items))
	curHeight := 0
	for i := 0; i < len(m.items); i++ {
		view := m.styles.Item.Render(m.items[i].View())
		viewHeight := lipgloss.Height(view)

		if !m.wrapping || m.bounds.Height == 0 || curHeight+viewHeight < m.bounds.Height {
			views = append(views, view)
			curHeight += viewHeight
		} else {
			col := groupStyle.Render(lipgloss.JoinVertical(m.position, views...))
			cols = append(cols, lipgloss.PlaceVertical(m.bounds.Height, groupStyle.GetAlign(), col))
			views = views[:1]
			views[0] = view
			curHeight = viewHeight
		}
	}

	if len(views) > 0 {
		col := groupStyle.Render(lipgloss.JoinVertical(m.position, views...))
		cols = append(cols, lipgloss.PlaceVertical(m.bounds.Height, groupStyle.GetAlign(), col))
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
