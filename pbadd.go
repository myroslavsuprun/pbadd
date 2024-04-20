package pbadd

import "golang.org/x/exp/constraints"

// Number interface contains a generic type for the parameners of the following funcitons:
// [github.com/myroslavsuprun/pbadd.Add]
type Number interface {
	constraints.Float | constraints.Integer
}

// Add sums two numbers and returns the result.
//
// A reference [link].
//
// [link]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](f, s T) T {
	return f + s
}
