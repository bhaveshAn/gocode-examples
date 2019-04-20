package main

import (
    "fmt"
    "errors"
    "math"
)

// Structure Definition
type Person struct {
	name string
	age int
}

func main() {
	x := 10
	y := 9
	sum := x + y
	fmt.Println("The type and val are : ")
	fmt.Printf("%T and val %d\n", x, x)
	fmt.Println(sum)

	if sum > 20 {
		fmt.Println("sum is greater than 20")
	} else {
		fmt.Println("sum is smaller than 20")
	}

	// Array Definition

	a := [5]int {1,2,3,4,5}
	var b [5]int
	a[0] = 10
	c := []int{10, 9, 8, 7, 6, 5}
	c = append(c, 4)
	fmt.Println(a, b, c)

	// Dictionary Definition


	vertices := make(map[string]int)
	vertices["a"] = 1
	vertices["b"] = 2
	vertices["c"] = 3
	delete(vertices, "b")
	fmt.Println(vertices)

	for i:= 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 5
	for i < 10{
		fmt.Println(i)
		i++
	}

	for i, val := range a{
		fmt.Println(i, val)
	}
	fmt.Println(prod(2, 5, 2.2))

	val, error := sqrt(-4)
	if error != nil{
		fmt.Println(error)
	} else {
		fmt.Println(val)
	}

	// Structure Declaration

	p := Person{name: "Bhavesh", age: 22}
	fmt.Println(p)

    // Dealing with Pointers

	d := 12.22
	inc(&d)
	fmt.Println(d)
}

// Functions

func prod(x, y int, z float64) int {
	return x * y * int(z)	
}

// Functions and errors

func sqrt(x float64) (float64, error) {

	if x < 0{
		return 0, errors.New("Negative Number")
	} else {
		return math.Sqrt(x), nil
	}
}

// Pointer in a function

func inc(x *float64) {
	*x++
}