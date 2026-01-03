// refreshing mind about pointers, intro to pointers
package main

import "fmt"

func main() {
	age := 32 // stored in memory somewhere, and automatically created the address for the stored value

	var agePointer *int // pointer type
	agePointer = &age

	fmt.Println("Age: ", *agePointer)

	EditAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func EditAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}

// Point of this? --> There is no copy of 32 being created. We create it once, gets stored in memory once, create a pointer that points to that address, and then we pass a pointer in the function. No copies being made, exactly one variable 32.
// This is worth doing ONLY for lagre & complex values. Here where we do it with int is not worth at all beacuse its a super small value in memory, and we shouldn't avoid passing integer in functions.
