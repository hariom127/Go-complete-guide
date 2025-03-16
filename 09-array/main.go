package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// *** without using make function ****
	// userNames := []string{}

	// ---- using make we do not need to recreate array when append or put a value till capacity of array
	//-----------> make(typeofArray, initial num of value, capacity of array)
	userNames := make([]string, 2, 5)

	//*** Have 2 empty space ***
	//*** Have capacity of 5 element ***

	fmt.Println(userNames) // print 2 empty value

	userNames = append(userNames, "Hariom")
	userNames = append(userNames, "Jone")

	fmt.Println(userNames) // print 4 values => 2 empty 2 names

	// put value on 0 index
	userNames[0] = "Hero"
	fmt.Println(userNames)

	//===== make function with map ====
	// *** till 5 element go do not allocate the memory to this map
	// ** using floatMap type alias here
	coursesRating := make(floatMap, 5)

	coursesRating["go"] = 4.7
	coursesRating["react"] = 4.5
	coursesRating["node"] = 4.6
	coursesRating["ai"] = 4.8

	coursesRating.output()
	// fmt.Println(coursesRating)

	//************loops**********
	// loop on slice or array
	for index, value := range userNames {
		fmt.Println("index: ", index)
		fmt.Println("value: ", value)
	}

	// loop on map
	for key, value := range coursesRating {
		fmt.Println("key: ", key)
		fmt.Println("value: ", value)
	}

}
