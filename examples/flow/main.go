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
	root tea.Model

	content         *items.Strings
	filteredContent *items.Filtered[string]
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
			cmds = append(cmds, named.Update(rootName, func(f flow.Model[items.FilteredItem[string]], msg tea.Msg) (tea.Model, tea.Cmd) {
				newOrientation := flow.Vertical
				if f.Orientation() == flow.Vertical {
					newOrientation = flow.Horizontal
				}
				f.SetOrientation(newOrientation)
				return f, nil
			}))
		case "up":
			m.content.Add("Item " + strconv.Itoa(m.content.Len()+1))
			m.filteredContent.ReapplyFilter()
		case "down":
			if m.content.Len() > 0 {
				m.content.RemoveAt(m.content.Len() - 1)
				m.filteredContent.ReapplyFilter()
			}
		case "f":
			filter := ""
			if m.content.Len() == m.filteredContent.Len() {
				filter = "1"
			}

			m.filteredContent.ApplyFilter(filter)

		case "q", "esc":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		cmds = append(cmds, named.Update(rootName, func(f flow.Model[items.FilteredItem[string]], msg tea.Msg) (tea.Model, tea.Cmd) {
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
	content := items.NewStrings("Item 1", "Item 2", "Item 3")
	filteredContent := items.NewFiltered[string](content, items.FuzzyStringFilter())
	filteredRenderer := items.RenderFunc[items.FilteredItem[string]](func(fi items.FilteredItem[string]) string {
		return content.Render(fi.Item)
	})

	mdl := mainModel{
		content:         content,
		filteredContent: filteredContent,
		root: named.New(rootName, flow.New[items.FilteredItem[string]](filteredContent, filteredRenderer,
			flow.WithWrapping[items.FilteredItem[string]](true),
			flow.WithPosition[items.FilteredItem[string]](0),
			flow.WithStyles[items.FilteredItem[string]](flow.Styles{
				Item:  lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true),
				Group: lipgloss.NewStyle().Align(lipgloss.Center),
			}),
		)),
	}

	examples.Run(mdl, tea.WithAltScreen(), tea.WithMouseAllMotion())
}
