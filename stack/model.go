package stack

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
	items       []Item
	orientation Orientation
	position    lipgloss.Position

	itemViews []string
}

func (m Model) Items() []Item {
	return m.items
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	for _, child := range m.items {
		cmds = append(cmds, child.init())
	}
	return tea.Batch(cmds...)
}

func (m Model) Orientation() Orientation {
	return m.orientation
}

func (m *Model) SetItems(items ...Item) {
	m.items = items
	m.renderItems()
}

func (m *Model) SetOrientation(orientation Orientation) {
	m.orientation = orientation
	m.renderItems()
}

func (m *Model) SetPosition(position lipgloss.Position) {
	m.position = position
	m.renderItems()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	for i, item := range m.items {
		m.items[i], cmd = item.update(msg)
		cmds = append(cmds, cmd)
	}

	m.renderItems()

	return m, tea.Batch(cmds...)
}

func (m Model) UpdateBounds(bounds teakwood.Rectangle) teakwood.Visual {
	m.bounds = bounds
	m.renderItems()
	return m
}

func (m Model) View() string {
	switch m.orientation {
	case Vertical:
		return lipgloss.JoinVertical(m.position, m.itemViews...)
	case Horizontal:
		return lipgloss.JoinHorizontal(m.position, m.itemViews...)
	default:
		panic(fmt.Sprintf("unknown orientation: %v", m.orientation))
	}
}

func (m *Model) renderItems() {
	absLengths := make([]int, len(m.items))
	if len(m.itemViews) != len(m.items) {
		m.itemViews = make([]string, len(m.items))
	}

	var remaining int
	switch m.orientation {
	case Vertical:
		remaining = m.bounds.Height
	case Horizontal:
		remaining = m.bounds.Width
	}

	var proportionalCount int
	// 1) account for absolute views.
	for i, item := range m.items {
		if item.Length.IsAuto() {
			continue
		}

		switch item.Length.Unit {
		case UnitAbsolute:
			remaining -= item.Length.Value
			absLengths[i] = item.Length.Value
		case UnitProportional:
			proportionalCount += item.Length.Value
		}
	}

	// 2) account for auto views.
	for i, item := range m.items {
		if !item.Length.IsAuto() {
			continue
		}

		// Temporarily, we'll clear the bounds to indicate we'd like these to be autosized. They'll
		// get set back later.
		m.items[i] = m.items[i].updateBounds(teakwood.Rectangle{})

		w, h := lipgloss.Size(m.items[i].view())
		switch m.orientation {
		case Vertical:
			remaining -= h
			absLengths[i] = h
		case Horizontal:
			remaining -= w
			absLengths[i] = w
		}
	}

	// 3) account for proportional views.
	for i, item := range m.items {
		if item.Length.IsAuto() {
			continue
		}

		switch item.Length.Unit {
		case UnitProportional:
			abs := item.Length.Value * remaining / proportionalCount
			remaining -= abs
			proportionalCount -= item.Length.Value
			absLengths[i] = abs
		}
	}

	curX := m.bounds.X
	curY := m.bounds.Y
	for i := range m.items {
		var itemStyle lipgloss.Style

		switch m.orientation {
		case Vertical:
			childBounds := teakwood.NewRectangle(curX, curY, m.bounds.Width, absLengths[i])
			m.items[i] = m.items[i].updateBounds(childBounds)
			curY += absLengths[i]
			itemStyle = lipgloss.NewStyle().MaxWidth(m.bounds.Width).MaxHeight(absLengths[i])
		case Horizontal:
			childBounds := teakwood.NewRectangle(curX, curY, absLengths[i], m.bounds.Height)
			m.items[i] = m.items[i].updateBounds(childBounds)
			curX += absLengths[i]
			itemStyle = lipgloss.NewStyle().MaxHeight(m.bounds.Height).MaxWidth(absLengths[i])
		}

		m.itemViews[i] = itemStyle.Render(m.items[i].view())
	}
}
