package items

import "github.com/sahilm/fuzzy"

func NewFiltered[T any](source Source[T], filterFunc FilterFunc[T]) *Filtered[T] {
	return &Filtered[T]{
		source:     source,
		filterFunc: filterFunc,
	}
}

type Filtered[T any] struct {
	source     Source[T]
	filterFunc FilterFunc[T]

	filter        string
	filteredItems []FilteredItem[T]
}

func (s *Filtered[T]) ApplyFilter(filter string) {
	s.filter = filter
	s.filteredItems = s.filterFunc(filter, s.source)
}

func (s *Filtered[T]) Filter() string {
	return s.filter
}

func (s *Filtered[T]) Item(index int) FilteredItem[T] {
	if len(s.filter) == 0 {
		return FilteredItem[T]{
			Item:           s.source.Item(index),
			Score:          index,
			MatchedIndexes: nil,
			originalIndex:  index,
		}
	}

	return s.filteredItems[index]
}

func (s *Filtered[T]) Len() int {
	if len(s.filter) > 0 {
		return len(s.filteredItems)
	}

	return s.source.Len()
}

func (s *Filtered[T]) ReapplyFilter() {
	s.ApplyFilter(s.filter)
}

type FilteredItem[T any] struct {
	Item           T
	Score          int
	MatchedIndexes []int

	originalIndex int
}

type FilterFunc[T any] func(filter string, source Source[T]) []FilteredItem[T]

func FuzzyFilter[T any](stringer func(T) string) FilterFunc[T] {
	return func(filter string, source Source[T]) []FilteredItem[T] {
		fuzzySrc := fuzzySource[T]{source: source, stringer: stringer}
		matches := fuzzy.FindFrom(filter, &fuzzySrc)

		results := make([]FilteredItem[T], len(matches))
		for i, match := range matches {
			results[i] = FilteredItem[T]{
				Item:           source.Item(match.Index),
				Score:          match.Score,
				MatchedIndexes: match.MatchedIndexes,
				originalIndex:  match.Index,
			}
		}

		return results
	}
}

func FuzzyStringFilter() FilterFunc[string] {
	return FuzzyFilter(func(s string) string { return s })
}

type fuzzySource[T any] struct {
	source   Source[T]
	stringer func(T) string
}

func (fs *fuzzySource[T]) Len() int {
	return fs.source.Len()
}

func (fs *fuzzySource[T]) String(index int) string {
	return fs.stringer(fs.source.Item(index))
}
