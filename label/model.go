package label

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
)

const styleKey = "label"

func New(text string, opts ...Opt) *Model {
	m := Model{
		text:    text,
		visible: true,
	}

	for _, opt := range opts {
		opt(&m)
	}

	return &m
}

type Model struct {
	bounds  teakwood.Rectangle
	name    string
	styler  teakwood.Styler
	text    string
	visible bool
}

func (m *Model) Init(styler teakwood.Styler) tea.Cmd {
	m.styler = styler
	return nil
}

func (m *Model) Measure(size teakwood.Size) teakwood.Size {
	if !m.visible {
		return teakwood.NewSize(0, 0)
	}

	s := m.getStyle(size)
	render := s.Render(m.text)
	width, height := lipgloss.Size(render)
	return teakwood.NewSize(width, height)
}

func (m *Model) SetText(text string) {
	m.text = text
}

func (m *Model) SetVisible(visible bool) {
	m.visible = visible
}

func (m *Model) Update(msg tea.Msg, bounds teakwood.Rectangle) (teakwood.Visual, tea.Cmd) {
	m.bounds = bounds
	return m, nil
}

func (m *Model) View() string {
	if !m.visible {
		return ""
	}

	return m.getStyle(m.bounds.Size()).Render(m.text)
}

func (m *Model) Visible() bool {
	return m.visible
}

func (m *Model) getStyle(size teakwood.Size) lipgloss.Style {
	s, ok := m.styler.Style("#" + m.name)
	if !ok {
		s, _ = m.styler.Style(styleKey)
	}

	return s.MaxWidth(size.Width).MaxHeight(size.Height)
}
