package main

import (
	"fmt"
	"sort"
)

func main() {

	myMap := make(map[string]int)

	// adding elements to the map
	myMap["Saklain"] = 27
	myMap["Tonmoy"] = 29
	myMap["Borsha"] = 25

	// by default, Map takes the key if there is only one variable
	for item := range myMap {
		fmt.Println(item)
	}

	fmt.Println("#################")

	// iterate over a map using key value pairs
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	// remove item from the map
	delete(myMap, "Tonmoy")

	fmt.Println("#################")

	fmt.Println(myMap)

	fmt.Println("#################")

	// Sorting Map Keys
	keys := []string{}

	for item := range myMap {
		keys = append(keys, item)
	}

	fmt.Println(keys)

	sort.Strings(keys)

	fmt.Println(keys)

	for _, value := range keys {
		fmt.Println(value, myMap[value])
	}

	fmt.Println("#################")
	
	// Sorting Map Values
	values := []int{}
	for _, value := range myMap {
		values = append(values, value)
	}

	fmt.Println(values)

	sort.Ints(values)

	fmt.Println(values)

	for _, value := range values {
		fmt.Println(value)
	}

}
