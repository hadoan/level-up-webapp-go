package main

import "fmt"

func mainHelloWord() {
	myString := "Hello World from myString"
	fmt.Println(myString)

	myInt := 2
	myFloat := 1.5
	myOtherFloat := float64(myInt) * myFloat
	fmt.Println(myOtherFloat)

	mySlice := []int{1, 2, 3, 4, 5}
	mySlice2 := mySlice[0:3]
	mySlice3 := mySlice[1:4]

	fmt.Println(mySlice, mySlice2, mySlice3, mySlice3[2])

	animals := []string{"Cat", "Dog", "Emu", "Warthog"}
	for i, animal := range animals {
		fmt.Println(animal, "is at index", i)
	}

	starWarsYears := map[string]int{
		"A New Hope":              1977,
		"The Empire Strikes back": 1980,
		"Return of the Jedi":      1983,
		"Attack of the Clones":    2002,
		"Revenge of the Sith":     2005,
	}

	fmt.Println("A New Hope ", starWarsYears["A New Hope"])
	// starWarsYears["A New Hope"] = 2010
	// fmt.Println("A New Hope ", starWarsYears["A New Hope"])

	for title, year := range starWarsYears {
		fmt.Println(title, "was released in", year)
	}

	myFuncString, myFuncInt := oneParamTwoReturns(9)
	fmt.Println(myFuncString)
	fmt.Println(myFuncInt)

	refArgument := 1
	copyRefFromCSharp(&refArgument)
	fmt.Println(refArgument)
}

func oneParamTwoReturns(myInt int) (string, int) {
	return fmt.Sprintf("int: %d", myInt), myInt + 1
}

func copyRefFromCSharp(refArgument *int) {
	*refArgument = *refArgument + 1
}
