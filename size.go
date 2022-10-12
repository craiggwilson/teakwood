package teakwood

func NewSize(width, height int) Size {
	return Size{
		Width:  width,
		Height: height,
	}
}

type Size struct {
	Width  int
	Height int
}
