package widget

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/craiggwilson/teakwood"
	"github.com/craiggwilson/teakwood/util/collections/list"
	"github.com/craiggwilson/teakwood/util/collections/set"
)

func New(name string, renderFn func(lipgloss.Style) string) *Widget {
	return &Widget{
		name:          name,
		classes:       set.NewMap[string](),
		pseudoClasses: set.NewMap[string](":enabled"),
		children:      list.NewSlice[teakwood.Visual](),
		renderFn:      renderFn,
	}
}

type Widget struct {
	name string

	id            string
	bounds        teakwood.Rectangle
	classes       *set.Map[string]
	pseudoClasses *set.Map[string]
	style         lipgloss.Style
	stylesheet    teakwood.StyleSheet

	canFocus bool
	enabled  bool
	focused  bool
	hovered  bool
	visible  bool

	parent   teakwood.Visual
	children *list.Slice[teakwood.Visual]

	renderFn func(lipgloss.Style) string
}

func (w *Widget) AddChildren(children ...teakwood.Visual) {
	for _, c := range children {
		w.children.Add(c)
		c.SetParent(w)
	}
}

func (w *Widget) AddClasses(classes ...string) {
	w.classes.Add(classes...)
}

func (w *Widget) AddPseudoClasses(classes ...string) {
	w.pseudoClasses.Add(classes...)
}

func (w *Widget) Ancestors() list.ReadOnly[teakwood.Visual] {
	lst := list.NewSlice[teakwood.Visual]()
	p := w.parent
	for p != nil {
		lst.Add(p)
		p = p.Parent()
	}

	lst.Reverse()
	return lst
}

func (w *Widget) Bounds() teakwood.Rectangle {
	return w.bounds
}

func (w *Widget) CanFocus() bool {
	return w.canFocus
}

func (w *Widget) Children() list.ReadOnly[teakwood.Visual] {
	return w.children
}

func (w *Widget) Classes() set.ReadOnly[string] {
	return w.classes
}

func (w *Widget) Enabled() bool {
	return w.enabled
}

func (w *Widget) Focused() bool {
	return w.focused
}

func (w *Widget) Hovered() bool {
	return w.hovered
}

func (w *Widget) ID() string {
	return w.id
}

func (w *Widget) Init() tea.Cmd {
	return nil
}

func (w *Widget) Name() string {
	return w.name
}

func (w *Widget) Parent() teakwood.Visual {
	return w.parent
}

func (w *Widget) PseudoClasses() set.ReadOnly[string] {
	return w.pseudoClasses
}

func (w *Widget) RemoveChildren(children ...teakwood.Visual) {
	for _, c := range children {
		if idx, ok := list.IndexOf[teakwood.Visual](w.children, c); ok {
			w.children.RemoveAt(idx)
		}
	}
}

func (w *Widget) RemoveClasses(classes ...string) {
	w.classes.Remove(classes...)
}

func (w *Widget) RemovePseudoClasses(classes ...string) {
	w.pseudoClasses.Remove(classes...)
}

func (w *Widget) SetBounds(bounds teakwood.Rectangle) {
	w.bounds = bounds
}

func (w *Widget) SetCanFocus(canFocus bool) {
	w.canFocus = canFocus
	if !w.canFocus {
		w.SetFocused(false)
	}
}

func (w *Widget) SetEnabled(enabled bool) {
	w.enabled = enabled
	if w.enabled {
		w.pseudoClasses.Add(":enabled")
		w.pseudoClasses.Remove(":disabled")
	} else {
		w.pseudoClasses.Add(":disabled")
		w.pseudoClasses.Remove(":enabled")
	}
}

func (w *Widget) SetFocused(focused bool) {
	if !w.canFocus {
		return
	}

	w.focused = focused
	if w.focused {
		w.pseudoClasses.Add(":focus")
	} else {
		w.pseudoClasses.Remove(":focus")
	}
}

func (w *Widget) SetID(id string) {
	w.id = id
}

func (w *Widget) SetParent(parent teakwood.Visual) {
	w.parent = parent
}

func (w *Widget) SetStyleSheet(stylesheet teakwood.StyleSheet) {
	w.stylesheet = stylesheet
}

func (w *Widget) SetVisible(visible bool) {
	w.visible = visible
}

func (w *Widget) Siblings() list.ReadOnly[teakwood.Visual] {
	lst := list.NewSlice[teakwood.Visual]()
	if w.parent == nil {
		return lst
	}

	it := w.parent.Children().Iter()

	for c, ok := it.Next(); ok; c, ok = it.Next() {
		if teakwood.Visual(w) != c {
			lst.Add(c)
		}
	}

	return lst
}

func (w *Widget) Style() lipgloss.Style {
	return w.style
}

func (w *Widget) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch tmsg := msg.(type) {
	case tea.MouseMsg:
		w.hovered = w.bounds.Contains(tmsg.X, tmsg.Y)
		if w.hovered {
			w.pseudoClasses.Add(":hover")
		} else {
			w.pseudoClasses.Remove(":hover")
		}

		if w.hovered && tmsg.Type == tea.MouseLeft {
			w.pseudoClasses.Add(":active")
		} else if tmsg.Type == tea.MouseRelease || !w.hovered {
			w.pseudoClasses.Remove(":active")
		}
	}

	// recompute style

	return w, nil
}

func (w *Widget) View() string {
	if !w.visible {
		return ""
	}

	return w.renderFn(w.style)
}

func (w *Widget) Visible() bool {
	return w.visible
}
