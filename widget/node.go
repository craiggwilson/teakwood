package widget

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/craiggwilson/teakwood/util/collections/list"
	"github.com/craiggwilson/teakwood/util/collections/set"
)

type Styler interface {
	Style(string) (lipgloss.Style, bool)
}

type Widget struct {
	name   string
	styler Styler

	id            string
	classes       *set.Map[string]
	pseudoClasses *set.Map[string]
	canFocus      bool
	focused       bool
	style         lipgloss.Style
	visible       bool

	parent   *Widget
	children *list.Slice[*Widget]

	render func() string
}

func (n *Widget) Ancestors() list.ReadOnly[*Widget] {
	lst := list.NewSlice[*Widget]()
	p := n.parent
	for p != nil {
		lst.Add(p)
		p = p.parent
	}

	lst.Reverse()
	return lst
}

func (n *Widget) Children() list.ReadOnly[*Widget] {
	return n.children
}

func (n *Widget) Classes() set.ReadOnly[string] {
	return n.classes
}

func (n *Widget) Focused() bool {
	return n.focused
}

func (n *Widget) ID() string {
	return n.id
}

func (n *Widget) Name() string {
	return n.name
}

func (n *Widget) Parent() *Widget {
	return n.parent
}

func (n *Widget) PseudoClasses() set.ReadOnly[string] {
	return n.pseudoClasses
}

func (n *Widget) Siblings() list.ReadOnly[*Widget] {
	lst := list.NewSlice[*Widget]()
	if n.parent == nil {
		return lst
	}

	it := n.parent.children.Iter()

	for c, ok := it.Next(); ok; c, ok = it.Next() {
		if n != c {
			lst.Add(c)
		}
	}

	return lst
}

func (n *Widget) Style() lipgloss.Style {
	return n.style
}

func (n *Widget) Visible() bool {
	return n.visible
}
