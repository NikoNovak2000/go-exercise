package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"
	userNames[1] = "Biggie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.9

	courseRatings.output()

	for index, name := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Name:", name)
	}

	for key, value := range courseRatings {
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}
}
