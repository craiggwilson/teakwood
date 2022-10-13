package label

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
	"github.com/craiggwilson/teakwood/util"
)

const StyleKey = "label"

func New(text string, opts ...Opt) *Model {
	m := Model{
		text: text,
		widget: &util.Widget{
			Kind:    StyleKey,
			Visible: true,
		},
	}

	for _, opt := range opts {
		opt(&m)
	}

	return &m
}

type Model struct {
	widget *util.Widget

	text string
}

func (m *Model) Init(styler teakwood.Styler) tea.Cmd {
	m.widget.Styler = styler
	return nil
}

func (m *Model) Measure(size teakwood.Size) teakwood.Size {
	if !m.widget.Visible {
		return teakwood.Size{}
	}

	s := m.widget.Style(size)
	render := s.Render(m.text)
	return teakwood.NewSize(lipgloss.Size(render))
}

func (m *Model) SetText(text string) {
	m.text = text
}

func (m *Model) SetVisible(visible bool) {
	m.widget.Visible = visible
}

func (m *Model) Update(msg tea.Msg, bounds teakwood.Rectangle) (teakwood.Visual, tea.Cmd) {
	m.widget.Bounds = bounds
	return m, nil
}

func (m *Model) View() string {
	return m.widget.Render(m.text)
}

func (m *Model) Visible() bool {
	return m.widget.Visible
}

func (m *Model) Widget2() {

}
