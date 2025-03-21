// Package learning_go_chapter_10_add contains an Add method that meets the specification of the Exercise in chapter 10 of Learning Go
package learning_go_chapter_10_add

import "golang.org/x/exp/constraints"

// Number is a combination of the Integer and Float constraints from [constraint]
//
// [constraint]: https://golang.org/x/exp/constraints
type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds two [github.com/CliffWheadon/learning_go_chapter_10_add/Number]s together
//
// It has two parameters: a and b which are both [github.com/CliffWheadon/learning_go_chapter_10_add/Number]s. It returns a [github.com/CliffWheadon/learning_go_chapter_10_add/Number] representing the sum of a and b
//
// For more information on addition, see: [Math is Fun]
//
// [Math is Fun]: https://www.mathsisfun.com/numbers/addition.html
func Add[N Number](a, b N) N {
	return a + b
}
