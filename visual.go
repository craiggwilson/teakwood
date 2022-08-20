package teacomps

import tea "github.com/charmbracelet/bubbletea"

type Visual interface {
	tea.Model

	UpdateBounds(Rectangle) Visual
}
