package main

import "fmt"

func main() {
	//1
	hobbyArray := [3]string{"Gym", "Football", "Kickbox"}
	fmt.Println(hobbyArray)

	//2
	fmt.Println(hobbyArray[0])
	newHobbyArray := hobbyArray[1:3]
	fmt.Println(newHobbyArray)

	//3
	sliceOneWay := hobbyArray[:2]
	//sliceSecondWay := hobbyArray[0:2]
	fmt.Println(sliceOneWay)
	//fmt.Println(sliceSecondWay)

	//4
	reSlice := sliceOneWay[1:3]
	fmt.Println(reSlice)

	//5
	courseGoals := []string{"Learn GO", "Have Fun"}
	fmt.Println(courseGoals)

	//6
	courseGoals[1] = "Have way more fun"
	courseGoals = append(courseGoals, "Build Projects")
	fmt.Println(courseGoals)

	//7
	type Product struct {
		title string
		id    string
		price float64
	}

	products := []Product{
		{title: "Peanuts", id: "1", price: 3.99},
		{title: "Crackers", id: "2", price: 5.99},
	}
	fmt.Println(products)

	newProduct := Product{title: "Cheese", id: "3", price: 7.99}

	products = append(products, newProduct)
	fmt.Println(products)
}

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.

// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list

// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)

// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.

// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)

// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array

// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
