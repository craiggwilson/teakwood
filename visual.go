package teakwood

import tea "github.com/charmbracelet/bubbletea"

type Visual interface {
	Init(Styler) tea.Cmd
	Measure(size Size) Size
	Update(tea.Msg, Rectangle) (Visual, tea.Cmd)
	View() string
}

type Focuser interface {
	Focused() bool
	SetFocus(bool)
}

type Visibler interface {
	Visible() bool
	SetVisible(bool)
}
