package pstack

func Absolute(v int) Length {
	return Length{
		Value: v,
		Unit:  UnitAbsolute,
	}
}

func Auto() Length {
	return Length{
		Value: 0,
		Unit:  UnitAbsolute,
	}
}

func Proportional(v int) Length {
	return Length{
		Value: v,
		Unit:  UnitProportional,
	}
}

type Length struct {
	Value int
	Unit  Unit
}

func (l Length) IsAuto() bool {
	return l.Value == 0
}

type Unit int

const (
	UnitAbsolute Unit = iota
	UnitProportional
)
