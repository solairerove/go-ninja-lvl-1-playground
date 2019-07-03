package exercise4

import "fmt"

type ownType int

var x ownType

// Ex4 is a complete solution on jedi lvl 1
func Ex4() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
}
