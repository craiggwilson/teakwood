package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood/examples"
	"github.com/craiggwilson/teakwood/frame"
	"github.com/craiggwilson/teakwood/label"
	"github.com/craiggwilson/teakwood/named"
	"github.com/craiggwilson/teakwood/stack"
)

const rootName = "root"

type mainModel struct {
	root tea.Model
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch tmsg := msg.(type) {
	case tea.KeyMsg:
		switch tmsg.String() {
		case "tab":
			cmds = append(cmds, named.Update(rootName, func(st stack.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
				newOrientation := stack.Vertical
				if st.Orientation() == stack.Vertical {
					newOrientation = stack.Horizontal
				}
				st.SetOrientation(newOrientation)
				return st, nil
			}))
		case "up":
			cmds = append(cmds, named.Update(rootName, func(st stack.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
				st.SetSpacing(st.Spacing() + 1)
				return st, nil
			}))
		case "down":
			cmds = append(cmds, named.Update(rootName, func(st stack.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
				if st.Spacing() == 0 {
					return st, nil
				}
				st.SetSpacing(st.Spacing() - 1)
				return st, nil
			}))
		case "q", "esc":
			return m, tea.Quit
		}
	}

	m.root, cmd = m.root.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
	return m.root.View() + "\n"
}

func main() {
	lbl1 := frame.New(lipgloss.NewStyle().BorderForeground(lipgloss.Color("1")).Border(lipgloss.NormalBorder(), true), label.New("Label 1"))
	lbl2 := frame.New(lipgloss.NewStyle().BorderForeground(lipgloss.Color("2")).Border(lipgloss.NormalBorder(), true), label.New("Label 2"))
	lbl3 := frame.New(lipgloss.NewStyle().BorderForeground(lipgloss.Color("3")).Border(lipgloss.NormalBorder(), true), label.New("Label 3"))

	mdl := mainModel{
		root: named.New(rootName, stack.New(stack.Vertical, lbl1, lbl2, lbl3)),
	}

	examples.Run(mdl, tea.WithAltScreen())
}
