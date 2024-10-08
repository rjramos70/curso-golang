package main

import "fmt"

func main() {
	// valores defaults dos tipos
	var a int
	var b float64
	var c bool
	var d string
	var e *int // ponteiro de int

	fmt.Printf("%v %v %v %v %q %v", a, b, c, d, d, e)

	// Output: 0 0 false  "" <nil>

}
