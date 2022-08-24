package stack

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/craiggwilson/teakwood"
)

func NewAbsoluteItem(size int, visual teakwood.Visual) Item {
	return NewItem(Absolute(size), visual)
}

func NewAutoItem(visual teakwood.Visual) Item {
	return NewItem(Auto(), visual)
}

func NewProportionalItem(proportion int, visual teakwood.Visual) Item {
	return NewItem(Proportional(proportion), visual)
}

func NewItem(length Length, visual teakwood.Visual) Item {
	return Item{
		Length: length,
		Visual: visual,
	}
}

type Item struct {
	Length Length
	Visual teakwood.Visual
}

func (i Item) init() tea.Cmd {
	return i.Visual.Init()
}

func (i Item) update(msg tea.Msg) (Item, tea.Cmd) {
	mdl, cmd := i.Visual.Update(msg)
	i.Visual = mdl.(teakwood.Visual)
	return i, cmd
}

func (i Item) updateBounds(bounds teakwood.Rectangle) Item {
	i.Visual = i.Visual.UpdateBounds(bounds)
	return i
}

func (i Item) view() string {
	return i.Visual.View()
}
