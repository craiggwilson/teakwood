package iter

import "fmt"

func makeEmptyError(name string) error {
	return fmt.Errorf("%s contains no elements", name)
}

func makeOutOfRangeError(name string) error {
	return fmt.Errorf("%s was out of range", name)
}
