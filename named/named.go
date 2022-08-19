package named

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/craiggwilson/teacomps/sizeutil"
)

func New(name string, model tea.Model) Model {
	return Model{name: name, model: model}
}

type Model struct {
	name  string
	model tea.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Name() string {
	return m.name
}

func (m *Model) SetHeight(v int) {
	m.model, _ = sizeutil.TrySetHeight(m.model, v)
}

func (m *Model) SetWidth(v int) {
	m.model, _ = sizeutil.TrySetWidth(m.model, v)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch tmsg := msg.(type) {
	case UpdateMsg:
		if tmsg.Name == m.name {
			m.model, cmd = tmsg.Update(m.model, msg)
			cmds = append(cmds, cmd)
		}
	}

	m.model, cmd = m.model.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.model.View()
}

type UpdateFunc func(tea.Model, tea.Msg) (tea.Model, tea.Cmd)

func Update[T any](name string, updater func(T, tea.Msg) (tea.Model, tea.Cmd)) tea.Cmd {
	return func() tea.Msg {
		return UpdateMsg{Name: name, Update: func(m tea.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
			return updater(m.(T), msg)
		}}
	}
}

type UpdateMsg struct {
	Name   string
	Update UpdateFunc
}
