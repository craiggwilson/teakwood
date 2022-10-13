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

func (s *Size) IsZero() bool {
	return s.Width == 0 && s.Height == 0
}
