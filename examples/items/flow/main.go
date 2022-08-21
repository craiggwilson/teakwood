package main

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
	"github.com/craiggwilson/teakwood/examples"
	"github.com/craiggwilson/teakwood/items"
	"github.com/craiggwilson/teakwood/items/flow"
	"github.com/craiggwilson/teakwood/named"
)

const rootName = "root"

type mainModel struct {
	root  tea.Model
	items items.StringsSource
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
			m.items = append(m.items, "Item "+strconv.Itoa(len(m.items)+1))
			cmds = append(cmds, named.Update(rootName, func(f flow.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
				f.SetItemsSource(m.items)
				return f, nil
			}))
		case "down":
			if len(m.items) > 0 {
				m.items = m.items[:len(m.items)-1]
				cmds = append(cmds, named.Update(rootName, func(f flow.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
					f.SetItemsSource(m.items)
					return f, nil
				}))
			}
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

func main() {
	items := items.StringsSource{"Item 1", "Item 2", "Item 3"}

	mdl := mainModel{
		items: items,
		root: named.New(rootName, flow.New(items,
			flow.WithWrapping(true),
			flow.WithPosition(0),
			flow.WithStyles(flow.Styles{
				Item:  lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true),
				Group: lipgloss.NewStyle().Align(lipgloss.Center),
			}),
		)),
	}

	examples.Run(mdl, tea.WithAltScreen())
}
