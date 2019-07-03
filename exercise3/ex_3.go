package exercise3

import "fmt"

var x = 42
var y = "James Bond"
var z = true

// Ex3 is a complete solution on jedi lvl 1
func Ex3() {
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)

	fmt.Println(s)
}
