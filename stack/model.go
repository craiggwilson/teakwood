package stack

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func New(o Orientation, children ...tea.Model) Model {
	return Model{
		orientation: o,
		children:    children,
	}
}

type Orientation int

const (
	Vertical Orientation = iota
	Horizontal
)

type Model struct {
	children    []tea.Model
	orientation Orientation
	position    lipgloss.Position
	spacing     int
}

func (m Model) Children() []tea.Model {
	return m.children
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

func (m *Model) SetChildren(children ...tea.Model) {
	m.children = children
}

func (m *Model) SetOrientation(orientation Orientation) {
	m.orientation = orientation
}

func (m *Model) SetPosition(position lipgloss.Position) {
	m.position = position
}

func (m *Model) SetSpacing(spacaing int) {
	m.spacing = spacaing
}

func (m Model) Spacing() int {
	return m.spacing
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
	var spacer string
	switch m.orientation {
	case Vertical:
		spacer = strings.Repeat("\n", m.spacing)
	case Horizontal:
		spacer = strings.Repeat(" ", m.spacing)
	}
	views := make([]string, len(m.children))
	for i, child := range m.children {
		if i > 0 {
			views[i-1] += spacer
		}
		views[i] = child.View()
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
