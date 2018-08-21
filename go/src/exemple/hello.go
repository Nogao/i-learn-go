package main

import (
	"fmt"	
)

func main() {

// first test, when you tap that, save, go on your Terminal, cd'yourFile' and write: go run hello.go
	fmt.Println("Hello, world")


// Now, as you can see, we use "x := 5" and not "var x = 5"
	x := 5
	y := 7
	sum := x + y

	fmt.Println(sum)

// in your terminal: go run hello.go
}