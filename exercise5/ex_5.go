package exercise5

import "fmt"

type ownType int

var x ownType
var y int

// Ex5 is a complete solution on jedi lvl 1
func Ex5() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
