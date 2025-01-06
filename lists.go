package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	hobbies := [3]string{"programming", "reading", "gaming"}
	fmt.Println(hobbies)

	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(cap(mainHobbies))

	courseGoals := []string{"learn go", "learn python"}
	fmt.Println(courseGoals)

	courseGoals = append(courseGoals, "learn java")
	fmt.Println(courseGoals)

	products := []Product{
		{
			"product 1",
			"1",
			10.0,
		},
		{
			"product 2",
			"2",
			20.0,
		},
	}

	fmt.Println(products)
	newProduct := Product{
		"product 3",
		"3",
		30.0,
	}
	products = append(products, newProduct)

	fmt.Println(products)
}
