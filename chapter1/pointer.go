package main

import "fmt"

func mainPointer() {
	fruit := "banana"
	giveMePear(&fruit)
	fmt.Println(fruit)
}

func giveMePear(fruit *string) {
	*fruit = "pear"
}
