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
	"github.com/craiggwilson/teakwood/items"
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

type mainModel struct {
	root   teakwood.Visual
	keyMap *keyMap

	documents items.Source[document]

	bounds teakwood.Rectangle
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
			cmds = append(cmds, named.Update(tabsName, func(tabs tabs.Model[document], msg tea.Msg) (tea.Model, tea.Cmd) {
				tabs.PrevTab()
				doc := m.documents.Item(tabs.CurrentTab())
				return tabs, named.Update(currentContent, func(f frame.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
					f.SetContent(label.New(doc.content))
					return f, nil
				})
			}))
		case key.Matches(tmsg, m.keyMap.NextPage):
			cmds = append(cmds, named.Update(tabsName, func(tabs tabs.Model[document], msg tea.Msg) (tea.Model, tea.Cmd) {
				tabs.NextTab()
				doc := m.documents.Item(tabs.CurrentTab())
				return tabs, named.Update(currentContent, func(f frame.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
					f.SetContent(label.New(doc.content))
					return f, nil
				})
			}))
		}
	case tea.WindowSizeMsg:
		m.bounds = teakwood.NewRectangle(0, 0, tmsg.Width, tmsg.Height)
	}

	v, cmd := m.root.Update(msg)
	m.root = v.(teakwood.Visual)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
	return m.root.ViewWithBounds(m.bounds)
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
			m.Width = 0
			return m.View(&km)
		}),
		adapter.WithBoundsViewer(func(m help.Model, bounds teakwood.Rectangle) string {
			m.Width = bounds.Width
			return m.View(&km)
		}),
	)

	mdl := mainModel{
		keyMap: &km,
	}

	mdl.documents = items.NewSlice(
		document{title: "Page 1", content: "Page 1 Content"},
		document{title: "Page 2", content: "Page 2 Content"},
	)

	documentTabRenderer := items.RenderFunc[document](func(item document) string {
		return item.title
	})

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
											named.New(tabsName, tabs.New[document](mdl.documents, documentTabRenderer)),
										),
										stack.NewItem(
											stack.Proportional(1),
											named.New(currentContent, frame.New(label.New(mdl.documents.Item(0).content))),
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
