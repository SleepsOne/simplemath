package simplemath

import "golang.org/x/exp/constraints"

// Number is a generic type that covers integers and floats.
type Number interface {
	constraints.Integer | constraints.Float
}

// Add returns the sum of two numbers of type Number.
// Learn more about addition here: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
