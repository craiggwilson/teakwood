package main

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teacomps/adapter"
	"github.com/craiggwilson/teacomps/examples"
	"github.com/craiggwilson/teacomps/frame"
	"github.com/craiggwilson/teacomps/label"
	"github.com/craiggwilson/teacomps/named"
	"github.com/craiggwilson/teacomps/pages"
	"github.com/craiggwilson/teacomps/pstack"
	"github.com/craiggwilson/teacomps/tabs"
)

const rootName = "root"
const tabsName = "tabs"
const pagesName = "pages"
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

type mainModel struct {
	root   tea.Model
	keyMap *keyMap

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
				return tabs, named.Update(pagesName, func(pages pages.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
					pages.SetCurrentPage(tabs.CurrentTab())
					return pages, nil
				})
			}))
		case key.Matches(tmsg, m.keyMap.NextPage):
			cmds = append(cmds, named.Update(tabsName, func(tabs tabs.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
				tabs.NextTab()
				return tabs, named.Update(pagesName, func(pages pages.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
					pages.SetCurrentPage(tabs.CurrentTab())
					return pages, nil
				})
			}))
		}
	case tea.WindowSizeMsg:
		cmds = append(cmds, named.Update(rootName, func(st pstack.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
			st.SetWidth(tmsg.Width)
			st.SetHeight(tmsg.Height)
			return st, nil
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
		adapter.WithView(func(m help.Model) string {
			return m.View(&km)
		}),
	)

	tab1 := label.New("Page 1")
	tab2 := label.New("Page 2")
	page1 := label.New("Page 1 Content")
	page2 := label.New("Page 2 Content")

	mdl := mainModel{
		keyMap: &km,
		root: named.New(rootName, pstack.New(
			pstack.Vertical,
			[]pstack.Length{pstack.Absolute(3), pstack.Proportional(1), pstack.Auto()},
			[]tea.Model{
				frame.New(
					lipgloss.NewStyle().BorderForeground(lipgloss.Color("1")).Border(lipgloss.NormalBorder(), true).Align(.5),
					label.New("Header"),
				),
				pstack.New(
					pstack.Horizontal,
					[]pstack.Length{pstack.Absolute(30), pstack.Proportional(6), pstack.Proportional(1)},
					[]tea.Model{
						frame.New(
							lipgloss.NewStyle().BorderForeground(lipgloss.Color("2")).Border(lipgloss.NormalBorder(), true),
							label.New("Left"),
						),
						frame.New(
							lipgloss.NewStyle().Margin(0),
							pstack.New(
								pstack.Vertical,
								[]pstack.Length{pstack.Auto(), pstack.Proportional(1)},
								[]tea.Model{
									named.New(tabsName, tabs.New(tab1, tab2)),
									named.New(pagesName, pages.New(page1, page2)),
								},
							),
						),
						frame.New(
							lipgloss.NewStyle().BorderForeground(lipgloss.Color("4")).Border(lipgloss.NormalBorder(), true),
							label.New("Right"),
						),
					},
				),
				frame.New(
					lipgloss.NewStyle().BorderForeground(lipgloss.Color("5")).Border(lipgloss.NormalBorder(), true),
					named.New(helpName, helpAdapter),
				),
			}),
		),
	}

	examples.Run(mdl, tea.WithAltScreen())
}
