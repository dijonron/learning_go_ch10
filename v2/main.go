// Package ch10_exercise implements a generic Add function for the solution to Chapter 10 of Learning Go, 2nd Edition.
package ch10_exercise

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add is a function that takes two numbers and returns their sum.
// See [Add] for implementation details.
//
// [Add]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
