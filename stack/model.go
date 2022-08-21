package stack

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
)

func New(orientation Orientation, lengths []Length, children []teakwood.Visual) Model {
	return Model{
		orientation: orientation,
		lengths:     lengths,
		children:    children,
	}
}

type Orientation int

const (
	Vertical Orientation = iota
	Horizontal
)

type Model struct {
	orientation Orientation
	position    lipgloss.Position

	children []teakwood.Visual
	lengths  []Length

	bounds teakwood.Rectangle
}

func (m Model) Children() []teakwood.Visual {
	return m.children
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	for _, child := range m.children {
		cmds = append(cmds, child.Init())
	}
	return tea.Batch(cmds...)
}

func (m Model) Lengths() []Length {
	return m.lengths
}

func (m Model) Orientation() Orientation {
	return m.orientation
}

func (m *Model) SetChildren(children ...teakwood.Visual) {
	m.children = children
}

func (m *Model) SetLengths(lengths ...Length) {
	m.lengths = lengths
}

func (m *Model) SetOrientation(orientation Orientation) {
	m.orientation = orientation
}

func (m *Model) SetPosition(position lipgloss.Position) {
	m.position = position
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	for i, child := range m.children {
		newChild, cmd := child.Update(msg)
		m.children[i] = newChild.(teakwood.Visual)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) UpdateBounds(bounds teakwood.Rectangle) teakwood.Visual {
	m.bounds = bounds
	return m
}

func (m Model) View() string {
	absLengths := m.computeAbsoluteLengths()

	views := make([]string, len(m.children))
	if len(absLengths) > 0 {
		for i, child := range m.children {
			s := lipgloss.NewStyle()
			switch m.orientation {
			case Vertical:
				s = s.MaxWidth(m.bounds.Width).MaxHeight(absLengths[i])
			case Horizontal:
				s = s.MaxHeight(m.bounds.Height).MaxWidth(absLengths[i])
			}

			views[i] = s.Render(child.View())
		}
	} else {
		for i, child := range m.children {
			views[i] = child.View()
		}
	}

	switch m.orientation {
	case Vertical:
		return lipgloss.JoinVertical(m.position, views...)
	case Horizontal:
		return lipgloss.JoinHorizontal(m.position, views...)
	default:
		panic(fmt.Sprintf("unknown orientation: %v", m.orientation))
	}
}

func (m *Model) computeAbsoluteLengths() []int {
	absLengths := make([]int, len(m.lengths))

	var remaining int
	switch m.orientation {
	case Vertical:
		remaining = m.bounds.Height
	case Horizontal:
		remaining = m.bounds.Width
	}

	proportionalCount := 0
	// 1) account for absolute views.
	for i, length := range m.lengths {
		if length.IsAuto() {
			continue
		}

		switch length.Unit {
		case UnitAbsolute:
			remaining -= length.Value
			absLengths[i] = length.Value
		case UnitProportional:
			proportionalCount += length.Value
		}
	}

	// 2) account for auto views.
	for i, length := range m.lengths {
		if !length.IsAuto() {
			continue
		}

		// Temporarily, we'll clear the bounds to indicate we'd like these to be autosized. They'll
		// get set back later.
		m.children[i] = m.children[i].UpdateBounds(teakwood.Rectangle{})

		w, h := lipgloss.Size(m.children[i].View())
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
	for i, length := range m.lengths {
		if length.IsAuto() {
			continue
		}

		switch length.Unit {
		case UnitProportional:
			abs := length.Value * remaining / proportionalCount
			remaining -= abs
			proportionalCount -= length.Value
			absLengths[i] = abs
		}
	}

	curX := m.bounds.X
	curY := m.bounds.Y
	for i := range m.children {
		switch m.orientation {
		case Vertical:
			childBounds := teakwood.NewRectangle(curX, curY, m.bounds.Width, absLengths[i])
			m.children[i] = m.children[i].UpdateBounds(childBounds)
			curY += absLengths[i]
		case Horizontal:
			childBounds := teakwood.NewRectangle(curX, curY, absLengths[i], m.bounds.Height)
			m.children[i] = m.children[i].UpdateBounds(childBounds)
			curX += absLengths[i]
		}
	}

	return absLengths
}
