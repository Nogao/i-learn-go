package main

import "fmt"

func main() {
	vertices := make(map[string]int)

	vertices["Wheels on car"] = 4
	vertices["Wheels on motorcycle"] = 2
	vertices["Wheels on truck"] = 12

	fmt.Println(vertices)
	fmt.Println(vertices["Wheels on car"])

	delete(vertices, "Wheels on motorcycle")

	fmt.Println(vertices)
}