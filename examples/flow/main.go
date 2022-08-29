package main

import (
	"fmt"
	"log"
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
		case "=":
			m.content.Add("Item " + strconv.Itoa(m.content.Len()+1))
			m.filteredContent.ReapplyFilter()
		case "-":
			if m.content.Len() > 0 {
				m.content.RemoveAt(m.content.Len() - 1)
				m.filteredContent.ReapplyFilter()
			}
		case "enter":
			cmds = append(cmds, named.Update(rootName, func(f flow.Model[items.FilteredItem[string]], msg tea.Msg) (tea.Model, tea.Cmd) {
				ci := f.CurrentIndex()
				if ci < 0 || ci >= m.filteredContent.Len() {
					return f, nil
				}

				sindexes := f.SelectedIndexes()
				contains := false
				var i int
				var si int
				for i, si = range sindexes {
					if si == ci {
						contains = true
						break
					}
				}

				if contains {
					log.Println("Removing", "Selected", sindexes, "Current", ci)
					sindexes = append(sindexes[:i], sindexes[i+1:]...)

				} else {
					log.Println("Adding", "Selected", sindexes, "Current", ci)
					sindexes = append(sindexes, ci)
				}

				f.SetSelectedIndexes(sindexes...)
				return f, nil
			}))
		case "down":
			cmds = append(cmds, named.Update(rootName, func(f flow.Model[items.FilteredItem[string]], msg tea.Msg) (tea.Model, tea.Cmd) {
				f.MoveCurrentIndexDown()
				return f, nil
			}))
		case "left":
			cmds = append(cmds, named.Update(rootName, func(f flow.Model[items.FilteredItem[string]], msg tea.Msg) (tea.Model, tea.Cmd) {
				f.MoveCurrentIndexLeft()
				return f, nil
			}))
		case "right":
			cmds = append(cmds, named.Update(rootName, func(f flow.Model[items.FilteredItem[string]], msg tea.Msg) (tea.Model, tea.Cmd) {
				f.MoveCurrentIndexRight()
				return f, nil
			}))
		case "up":
			cmds = append(cmds, named.Update(rootName, func(f flow.Model[items.FilteredItem[string]], msg tea.Msg) (tea.Model, tea.Cmd) {
				f.MoveCurrentIndexUp()
				return f, nil
			}))

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
	itemStrings := make([]string, 3)
	for i := range itemStrings {
		itemStrings[i] = fmt.Sprintf("Item %d", i+1)
	}
	content := items.NewStrings(itemStrings...)
	filteredContent := items.NewFiltered[string](content, items.FuzzyStringFilter())
	filteredRenderer := items.RenderFunc[items.FilteredItem[string]](func(fi items.FilteredItem[string]) string {
		return content.Render(fi.Item)
	})

	mdl := mainModel{
		content:         content,
		filteredContent: filteredContent,
		root: named.New(rootName, flow.New[items.FilteredItem[string]](filteredContent, filteredRenderer,
			flow.WithHorizontalAlignment[items.FilteredItem[string]](lipgloss.Center),
			flow.WithVerticalAlignment[items.FilteredItem[string]](lipgloss.Center),
			flow.WithStyles[items.FilteredItem[string]](flow.Styles{
				Item:         lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true),
				Group:        lipgloss.NewStyle().Align(lipgloss.Center),
				SelectedItem: lipgloss.NewStyle().Border(lipgloss.DoubleBorder(), true),
				CurrentItem:  lipgloss.NewStyle().Border(lipgloss.DoubleBorder(), true).BorderForeground(lipgloss.Color("1")),
			}),
		)),
	}

	examples.Run(mdl, tea.WithAltScreen(), tea.WithMouseAllMotion())
}
