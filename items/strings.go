package items

func NewStrings(items ...string) *Strings {
	return &Strings{
		Slice: Slice[string]{
			items: items,
		},
	}
}

type Strings struct {
	Slice[string]
}

func (s *Strings) Render(item string) string {
	return item
}
