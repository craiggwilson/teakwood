package frame

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/craiggwilson/teakwood"
)

const styleKey = "frame"

func New(content teakwood.Visual, opts ...Opt) Model {
	m := Model{
		content: content,
	}

	for _, opt := range opts {
		opt(&m)
	}

	return m
}

type Model struct {
	bounds  teakwood.Rectangle
	content teakwood.Visual
	name    string
	styler  teakwood.Styler
}

func (m *Model) Init(styler teakwood.Styler) tea.Cmd {
	m.styler = styler
	return m.content.Init(styler)
}

func (m *Model) Measure(size teakwood.Size) teakwood.Size {
	s := m.getStyle(size)
	offsets := teakwood.OffsetsFromStyle(s)
	offsetsSize := offsets.Size()

	contentMeasure := m.content.Measure(offsetsSize)
	contentMeasure.Height += offsetsSize.Height
	contentMeasure.Width += offsetsSize.Width
	return contentMeasure
}

func (m *Model) SetContent(content teakwood.Visual) {
	m.content = content
}

func (m *Model) Update(msg tea.Msg, bounds teakwood.Rectangle) (teakwood.Visual, tea.Cmd) {
	m.bounds = bounds

	s := m.getStyle(m.bounds.Size())
	offsets := teakwood.OffsetsFromStyle(s)

	var cmd tea.Cmd
	m.content, cmd = m.content.Update(msg, bounds.Offset(offsets))
	return m, cmd
}

func (m *Model) View() string {
	s := m.getStyle(m.bounds.Size())
	return s.Render(m.content.View())
}

func (m *Model) getStyle(size teakwood.Size) lipgloss.Style {
	s, ok := m.styler.Style("#" + m.name)
	if !ok {
		s, _ = m.styler.Style(styleKey)
	}

	return s.MaxWidth(size.Width).MaxHeight(size.Height)
}
