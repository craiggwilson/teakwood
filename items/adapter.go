package items

func NewAdapter[TFrom any, TTo any](source Source[TFrom], adaptFunc func(TFrom) TTo) *Adapter[TFrom, TTo] {
	return &Adapter[TFrom, TTo]{
		source:    source,
		adaptFunc: adaptFunc,
	}
}

type Adapter[TFrom any, TTo any] struct {
	source    Source[TFrom]
	adaptFunc func(TFrom) TTo
}

func (a *Adapter[TFrom, TTo]) Item(index int) TTo {
	from := a.source.Item(index)
	return a.adaptFunc(from)
}

func (a *Adapter[TFrom, TTo]) Len() int {
	return a.source.Len()
}
