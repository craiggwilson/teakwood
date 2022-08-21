package main

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
	"github.com/craiggwilson/teakwood/adapter"
	"github.com/craiggwilson/teakwood/examples"
	"github.com/craiggwilson/teakwood/frame"
	"github.com/craiggwilson/teakwood/items/tabs"
	"github.com/craiggwilson/teakwood/label"
	"github.com/craiggwilson/teakwood/named"
	"github.com/craiggwilson/teakwood/stack"
)

const rootName = "root"
const tabsName = "tabs"
const currentContent = "pages"
const helpName = "help"

type keyMap struct {
	Help     key.Binding
	NextPage key.Binding
	PrevPage key.Binding
	Quit     key.Binding
}

func (m *keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		m.Quit, m.Help,
	}
}

func (m *keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{m.Quit, m.Help},
		{m.PrevPage, m.NextPage},
	}
}

type document struct {
	title   string
	content string
}

func (d *document) View() string {
	return d.content
}

type documentTitles struct {
	documents *[]document
}

func (m documentTitles) Len() int {
	return len(*m.documents)
}

func (m documentTitles) Item(idx int) tea.Model {
	return label.New((*m.documents)[idx].title)
}

type mainModel struct {
	root   tea.Model
	keyMap *keyMap

	documents []document

	width  int
	height int
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch tmsg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(tmsg, m.keyMap.Quit):
			return m, tea.Quit
		case key.Matches(tmsg, m.keyMap.Help):
			cmds = append(cmds, named.Update(helpName, func(adapter adapter.Model[help.Model], msg tea.Msg) (tea.Model, tea.Cmd) {
				adapter.UpdateAdaptee(func(h help.Model) help.Model {
					h.ShowAll = !h.ShowAll
					return h
				})

				return adapter, nil
			}))
		case key.Matches(tmsg, m.keyMap.PrevPage):
			cmds = append(cmds, named.Update(tabsName, func(tabs tabs.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
				tabs.PrevTab()
				return tabs, named.Update(currentContent, func(a adapter.Model[document], msg tea.Msg) (tea.Model, tea.Cmd) {
					a.UpdateAdaptee(func(document) document {
						return m.documents[tabs.CurrentTab()]
					})
					return a, nil
				})
			}))
		case key.Matches(tmsg, m.keyMap.NextPage):
			cmds = append(cmds, named.Update(tabsName, func(tabs tabs.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
				tabs.NextTab()
				return tabs, named.Update(currentContent, func(a adapter.Model[document], msg tea.Msg) (tea.Model, tea.Cmd) {
					a.UpdateAdaptee(func(document) document {
						return m.documents[tabs.CurrentTab()]
					})
					return a, nil
				})
			}))
		}
	case tea.WindowSizeMsg:
		cmds = append(cmds, named.Update(rootName, func(st stack.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
			return st.UpdateBounds(teakwood.NewRectangle(0, 0, tmsg.Width, tmsg.Height)), nil
		}))
	}

	m.root, cmd = m.root.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
	return m.root.View()
}

func main() {
	km := keyMap{
		Help:     key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
		NextPage: key.NewBinding(key.WithKeys("k", "right"), key.WithHelp("k/right", "next page")),
		PrevPage: key.NewBinding(key.WithKeys("j", "left"), key.WithHelp("j/left", "prev page")),
		Quit:     key.NewBinding(key.WithKeys("q", "esc"), key.WithHelp("q/esc", "quit")),
	}

	helpAdapter := adapter.New(
		help.New(),
		adapter.WithUpdateBounds(func(m help.Model, bounds teakwood.Rectangle) help.Model {
			m.Width = bounds.Width
			return m
		}),
		adapter.WithView(func(m help.Model) string {
			return m.View(&km)
		}),
	)

	mdl := mainModel{
		keyMap: &km,
	}

	mdl.documents = []document{
		{title: "Page 1", content: "Page 1 Content"},
		{title: "Page 2", content: "Page 2 Content"},
	}

	mdl.root = named.New(rootName, stack.New(
		stack.Vertical,
		[]stack.Length{stack.Absolute(3), stack.Proportional(1), stack.Auto()},
		[]teakwood.Visual{
			frame.New(
				label.New("Header"),
				frame.WithStyle(lipgloss.NewStyle().BorderForeground(lipgloss.Color("1")).Border(lipgloss.NormalBorder(), true).Align(.5)),
			),
			stack.New(
				stack.Horizontal,
				[]stack.Length{stack.Absolute(30), stack.Proportional(6), stack.Proportional(1)},
				[]teakwood.Visual{
					frame.New(
						label.New("Left"),
						frame.WithStyle(lipgloss.NewStyle().BorderForeground(lipgloss.Color("2")).Border(lipgloss.NormalBorder(), true)),
					),
					frame.New(
						stack.New(
							stack.Vertical,
							[]stack.Length{stack.Auto(), stack.Proportional(1)},
							[]teakwood.Visual{
								named.New(tabsName, tabs.New(documentTitles{&mdl.documents})),
								named.New(currentContent, adapter.New(mdl.documents[0], adapter.WithView(func(d document) string {
									return d.View()
								}))),
							},
						),
						frame.WithStyle(lipgloss.NewStyle().Margin(0)),
					),
					frame.New(
						label.New("Right"),
						frame.WithStyle(lipgloss.NewStyle().BorderForeground(lipgloss.Color("4")).Border(lipgloss.NormalBorder(), true)),
					),
				},
			),
			frame.New(
				named.New(helpName, helpAdapter),
				frame.WithStyle(lipgloss.NewStyle().BorderForeground(lipgloss.Color("5")).Border(lipgloss.NormalBorder(), true)),
			),
		}),
	)

	examples.Run(mdl, tea.WithAltScreen())
}
