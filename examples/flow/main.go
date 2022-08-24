package main

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
	"github.com/craiggwilson/teakwood/examples"
	"github.com/craiggwilson/teakwood/flow"
	"github.com/craiggwilson/teakwood/named"
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
			cmds = append(cmds, named.Update(rootName, func(f flow.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
				newOrientation := flow.Vertical
				if f.Orientation() == flow.Vertical {
					newOrientation = flow.Horizontal
				}
				f.SetOrientation(newOrientation)
				return f, nil
			}))
		case "up":
			cmds = append(cmds, named.Update(rootName, func(f flow.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
				f.AddItem(hoverLabelModel{text: "Item " + strconv.Itoa(f.Len()+1)})
				return f, nil
			}))
		case "down":
			cmds = append(cmds, named.Update(rootName, func(f flow.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
				if f.Len() > 0 {
					f.RemoveItem(f.Len() - 1)
				}
				return f, nil
			}))
		case "q", "esc":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		cmds = append(cmds, named.Update(rootName, func(f flow.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
			return f.UpdateBounds(teakwood.NewRectangle(0, 0, tmsg.Width-1, tmsg.Height-1)), nil
		}))
	}

	m.root, cmd = m.root.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
	return m.root.View() + "\n"
}

type hoverLabelModel struct {
	text  string
	hover bool

	bounds teakwood.Rectangle
}

func (m hoverLabelModel) Init() tea.Cmd {
	return nil
}

func (m hoverLabelModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch tmsg := msg.(type) {
	case tea.MouseMsg:
		switch tmsg.Type {
		case tea.MouseMotion:
			m.hover = m.bounds.Contains(tmsg.X, tmsg.Y)
		}
	}

	return m, nil
}

func (m hoverLabelModel) UpdateBounds(bounds teakwood.Rectangle) teakwood.Visual {
	m.bounds = bounds
	return m
}

func (m hoverLabelModel) View() string {
	return lipgloss.NewStyle().Underline(m.hover).Render(m.text)
}

func main() {
	mdl := mainModel{
		root: named.New(rootName, flow.New(
			flow.WithItems(
				hoverLabelModel{text: "Item 1"},
				hoverLabelModel{text: "Item 2"},
				hoverLabelModel{text: "Item 3"},
			),
			flow.WithWrapping(true),
			flow.WithPosition(0),
			flow.WithStyles(flow.Styles{
				Item:  lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true),
				Group: lipgloss.NewStyle().Align(lipgloss.Center),
			}),
		)),
	}

	examples.Run(mdl, tea.WithAltScreen(), tea.WithMouseAllMotion())
}
