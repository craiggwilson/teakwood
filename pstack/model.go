package pstack

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teacomps/sizeutil"
)

func New(orientation Orientation, lengths []Length, children []tea.Model) Model {
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

	children []tea.Model
	lengths  []Length
	width    int
	height   int
}

func (m Model) Children() []tea.Model {
	return m.children
}

func (m Model) Height() int {
	return m.height
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

func (m *Model) SetChildren(v ...tea.Model) {
	m.children = v
}

func (m *Model) SetHeight(v int) {
	m.height = v
}

func (m *Model) SetLengths(v ...Length) {
	m.lengths = v
}

func (m *Model) SetOrientation(v Orientation) {
	m.orientation = v
}

func (m *Model) SetPosition(v lipgloss.Position) {
	m.position = v
}

func (m *Model) SetWidth(v int) {
	m.width = v
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	for i, child := range m.children {
		m.children[i], cmd = child.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	absLengths := m.computeAbsoluteLengths()

	views := make([]string, len(m.children))
	if len(absLengths) > 0 {
		for i, child := range m.children {
			s := lipgloss.NewStyle()
			switch m.orientation {
			case Vertical:
				s = s.MaxWidth(m.width).MaxHeight(absLengths[i])
			case Horizontal:
				s = s.MaxHeight(m.height).MaxWidth(absLengths[i])
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
		remaining = m.height
	case Horizontal:
		remaining = m.width
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

		// Reset the size back to the default. Ignore the cmds from this as we'll be
		// calling this same function on the child later and will use that one.
		m.children[i], _ = sizeutil.TrySetHeight(m.children[i], 0)
		m.children[i], _ = sizeutil.TrySetWidth(m.children[i], 0)

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

	for i := range m.children {
		switch m.orientation {
		case Vertical:
			m.children[i], _ = sizeutil.TrySetHeight(m.children[i], absLengths[i])
			m.children[i], _ = sizeutil.TrySetWidth(m.children[i], m.width)
		case Horizontal:
			m.children[i], _ = sizeutil.TrySetHeight(m.children[i], m.height)
			m.children[i], _ = sizeutil.TrySetWidth(m.children[i], absLengths[i])
		}
	}

	return absLengths
}
