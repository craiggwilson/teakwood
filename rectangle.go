package teakwood

import "github.com/charmbracelet/lipgloss"

func NewRectangle(x, y, width, height int) Rectangle {
	return Rectangle{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func OffsetsFromStyle(style lipgloss.Style) Rectangle {
	return NewRectangle(
		style.GetMarginTop()+style.GetPaddingTop()+style.GetBorderTopWidth(),
		style.GetMarginBottom()+style.GetPaddingBottom()+style.GetBorderBottomSize(),
		style.GetHorizontalFrameSize(),
		style.GetVerticalFrameSize(),
	)
}

type Rectangle struct {
	X      int
	Y      int
	Width  int
	Height int
}

func (r *Rectangle) Bottom() int {
	return r.Y + r.Height
}

func (r *Rectangle) Left() int {
	return r.X
}

func (r *Rectangle) Right() int {
	return r.X + r.Width
}

func (r *Rectangle) Size() Size {
	return NewSize(r.Width, r.Height)
}

func (r *Rectangle) Top() int {
	return r.Y
}

func (r *Rectangle) Contains(x, y int) bool {
	return x >= r.Left() && x <= r.Right() && y >= r.Top() && y <= r.Bottom()
}

func (r Rectangle) Offset(other Rectangle) Rectangle {
	return NewRectangle(r.X+other.X, r.Y+other.Y, r.Width-other.Width, r.Height-other.Height)
}
