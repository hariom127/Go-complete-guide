package list

import "fmt"

func main() {

	//---array with dynamic limit---
	prices := []float64{10.67, 56.67, 34.34}

	fmt.Println(prices)

	//--- we can not add new value in array out of range directly ---
	// prices[3] = 50.55

	// ---- Add new value which is out of range using append method ---
	prices = append(prices, 50.55, 45.55)

	fmt.Println(prices)

	//--- create array type of struct---
	type product struct {
		id    string
		title string
	}

	productList := []product{
		{"1", "product-1"},
		{"2", "product-2"},
		{"3", "product-3"},
	}
	fmt.Println(productList)

	//append new struct in array

	newProduct := product{"4", "product-4"}

	productList = append(productList, newProduct)
	fmt.Println(productList)

	// append one slice to another
	featurePrices := []float64{23.45, 66.44, 67.34}

	discountedPrice := append(prices, featurePrices...)
	fmt.Println(discountedPrice)

}

//----Array basic------
// func main() {
// 	productName := [4]string{"product-A"}
// 	pricing := [4]float64{34.4, 56.6, 66.6, 67.0}
// 	fmt.Println(pricing)

// 	productName[2] = "product-C"
// 	fmt.Println(productName)
// 	fmt.Println(productName[2])

// 	// slice startIndex include endIndex exclude
// 	featurePrice := pricing[1:3]
// 	// featurePrice2 := pricing[1:]
// 	// featurePrice3 := pricing[:3]

// 	fmt.Println(featurePrice)
// 	// fmt.Println(featurePrice2)
// 	// fmt.Println(featurePrice3)

// 	// slice refer the same memory location of the parent array
// 	// which means if we change sliced array element its overright the parent too
// 	featurePrice[0] = 1000.99
// 	fmt.Println(featurePrice)
// 	fmt.Println(pricing)
// 	fmt.Println(len(featurePrice), cap(featurePrice))

// }
