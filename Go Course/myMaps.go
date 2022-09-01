package main

import "fmt"

func executeMaps() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	family := make(map[string]string)
	family["father"] = "Rafael"
	family["mother"] = "Flavia"
	family["daughter"] = "Alice"

	fmt.Println(colors)
	fmt.Println(family)

	delete(colors, "red")
	fmt.Println(colors)

	//iterate
	printMap(colors)
	fmt.Println("==================================")
	printMap(family)
	remove(family, "father")
	fmt.Println("==================================")
	printMap(family)
	fmt.Println("==================================")
}

func remove(m map[string]string, key string) {
	delete(m, key)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Printf("The key %s is value %s", key, value)
		fmt.Println()
	}
}
