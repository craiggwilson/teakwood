package frame

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/craiggwilson/teakwood"
	"github.com/craiggwilson/teakwood/util"
)

const StyleKey = "frame"

func New(content teakwood.Visual, opts ...Opt) *Model {
	m := Model{
		content: content,
		widget: &util.Widget{
			Fill:    true,
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
	widget  *util.Widget
	content teakwood.Visual
}

func (m *Model) Init(styler teakwood.StyleSheet) tea.Cmd {
	m.widget.Styler = styler
	return m.content.Init(styler)
}

func (m *Model) Measure(size teakwood.Size) teakwood.Size {
	s := m.widget.Style(size)
	result := teakwood.NewSize(
		util.Min(size.Width, s.GetWidth()),
		util.Min(size.Height, s.GetHeight()),
	)

	return result
}

func (m *Model) SetContent(content teakwood.Visual) {
	m.content = content
}

func (m *Model) Update(msg tea.Msg, bounds teakwood.Rectangle) (teakwood.Visual, tea.Cmd) {
	m.widget.Bounds = bounds

	var cmd tea.Cmd
	m.content, cmd = m.content.Update(msg, m.widget.ContentBounds())
	return m, cmd
}

func (m *Model) View() string {
	return m.widget.Render(m.content.View())
}
