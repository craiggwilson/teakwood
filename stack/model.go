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
}

func (m *Model) SetOrientation(orientation Orientation) {
	m.orientation = orientation
}

func (m *Model) SetPosition(position lipgloss.Position) {
	m.position = position
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	for i, item := range m.items {
		m.items[i], cmd = item.update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	m.renderItems()
	return m.renderView()
}

func (m Model) ViewWithBounds(bounds teakwood.Rectangle) string {
	m.renderItemsWithBounds(bounds)
	return m.renderView()
}

func (m Model) renderView() string {
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
	if len(m.itemViews) != len(m.items) {
		m.itemViews = make([]string, len(m.items))
	}

	for i, item := range m.items {
		m.itemViews[i] = item.view()
	}
}

func (m *Model) renderItemsWithBounds(bounds teakwood.Rectangle) {
	absLengths := make([]int, len(m.items))
	if len(m.itemViews) != len(m.items) {
		m.itemViews = make([]string, len(m.items))
	}

	var remaining int
	switch m.orientation {
	case Vertical:
		remaining = bounds.Height
	case Horizontal:
		remaining = bounds.Width
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

	curX := bounds.X
	curY := bounds.Y
	for i := range m.items {
		var itemStyle lipgloss.Style

		switch m.orientation {
		case Vertical:
			childBounds := teakwood.NewRectangle(curX, curY, bounds.Width, absLengths[i])
			m.itemViews[i] = itemStyle.Render(m.items[i].viewWithBounds(childBounds))
			curY += absLengths[i]
			itemStyle = lipgloss.NewStyle().MaxWidth(bounds.Width).MaxHeight(absLengths[i])
		case Horizontal:
			childBounds := teakwood.NewRectangle(curX, curY, absLengths[i], bounds.Height)
			m.itemViews[i] = itemStyle.Render(m.items[i].viewWithBounds(childBounds))
			curX += absLengths[i]
			itemStyle = lipgloss.NewStyle().MaxHeight(bounds.Height).MaxWidth(absLengths[i])
		}
	}
}
