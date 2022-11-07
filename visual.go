package teakwood

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/craiggwilson/teakwood/util/collections/list"
	"github.com/craiggwilson/teakwood/util/collections/set"
)

type Visual interface {
	tea.Model

	AddClasses(classes ...string)
	Ancestors() list.ReadOnly[Visual]
	Bounds() Rectangle
	CanFocus() bool
	Children() list.ReadOnly[Visual]
	Classes() set.ReadOnly[string]
	Enabled() bool
	Focused() bool
	Hovered() bool
	ID() string
	Name() string
	Parent() Visual
	PseudoClasses() set.ReadOnly[string]

	SetBounds(Rectangle)
	SetEnabled(bool)
	SetFocused(bool)
	SetID(string)
	SetParent(Visual)
	SetStyleSheet(StyleSheet)
	SetVisible(bool)
}
