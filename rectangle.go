package teacomps

func NewRectangle(x, y, width, height int) Rectangle {
	return Rectangle{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
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

func (r *Rectangle) Top() int {
	return r.Y
}
