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
	"github.com/craiggwilson/teakwood/label"
	"github.com/craiggwilson/teakwood/named"
	"github.com/craiggwilson/teakwood/stack"
	"github.com/craiggwilson/teakwood/tabs"
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
				doc := m.documents[tabs.CurrentTab()]
				return tabs, named.Update(currentContent, func(f frame.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
					f.SetContent(label.New(doc.content))
					return f, nil
				})
			}))
		case key.Matches(tmsg, m.keyMap.NextPage):
			cmds = append(cmds, named.Update(tabsName, func(tabs tabs.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
				tabs.NextTab()
				doc := m.documents[tabs.CurrentTab()]
				return tabs, named.Update(currentContent, func(f frame.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
					f.SetContent(label.New(doc.content))
					return f, nil
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

	tabItems := []tea.Model{
		label.New(mdl.documents[0].title),
		label.New(mdl.documents[1].title),
	}

	mdl.root = named.New(
		rootName,
		stack.New(
			stack.WithOrientation(stack.Vertical),
			stack.WithItems(
				stack.NewItem(
					stack.Absolute(3),
					frame.New(
						label.New("Header"),
						frame.WithStyle(lipgloss.NewStyle().BorderForeground(lipgloss.Color("1")).Border(lipgloss.NormalBorder(), true).Align(.5)),
					),
				),
				stack.NewItem(
					stack.Proportional(1),
					stack.New(
						stack.WithItems(
							stack.NewItem(
								stack.Absolute(30),
								frame.New(
									label.New("Left"),
									frame.WithStyle(lipgloss.NewStyle().BorderForeground(lipgloss.Color("2")).Border(lipgloss.NormalBorder(), true)),
								),
							),
							stack.NewItem(
								stack.Proportional(6),
								stack.New(
									stack.WithOrientation(stack.Vertical),
									stack.WithItems(
										stack.NewAutoItem(
											named.New(tabsName, tabs.New(tabs.WithItems(tabItems...))),
										),
										stack.NewItem(
											stack.Proportional(1),
											named.New(currentContent, frame.New(label.New(mdl.documents[0].content))),
										),
									),
								),
							),
							stack.NewItem(
								stack.Proportional(1),
								frame.New(
									label.New("Right"),
									frame.WithStyle(lipgloss.NewStyle().BorderForeground(lipgloss.Color("4")).Border(lipgloss.NormalBorder(), true)),
								),
							),
						),
					),
				),
				stack.NewAutoItem(
					frame.New(
						named.New(helpName, helpAdapter),
						frame.WithStyle(lipgloss.NewStyle().BorderForeground(lipgloss.Color("5")).Border(lipgloss.NormalBorder(), true)),
					),
				),
			),
		),
	)

	examples.Run(mdl, tea.WithAltScreen(), tea.WithMouseAllMotion())
}
