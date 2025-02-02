package main

func main() {
	age := 35

	var agePointer *int

	agePointer = &age

	println("agePointer:", agePointer)
	println("agePointerToValue:", *agePointer)

	getAdultAge(agePointer)
	println(age)

}

func getAdultAge(age *int) {
	// return *age - 18
	*age = *age - 18 // over write the age value using pointer
}
