package main

import "fmt"

func main() {

	websites := map[string]string{
		"google":   "https://www.google.com",
		"facebook": "https://www.facebook.com",
	}
	websites["x"] = "https://www.x.com"

	fmt.Println(websites)
	fmt.Println(websites["google"])

	delete(websites, "x")
	fmt.Println(websites)
}
