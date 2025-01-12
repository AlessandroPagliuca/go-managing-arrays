package main

import "fmt"

// alias for map
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2)

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Ale")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 2)

	courseRatings["go"] = 4.7
	courseRatings["angular"] = 4.8

	courseRatings.output()

	for index, value := range userNames {
		fmt.Println("index:", index)
		fmt.Println("value:", value)
	}

	for key, value := range courseRatings {
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}
}
