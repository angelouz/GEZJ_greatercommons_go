package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		"James",
		"Bond",
		32,
	}

	p2 := person{
		"Miss",
		"Moneypenny",
		28,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last, p2.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
