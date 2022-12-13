package main

import "fmt"

func main() {
	// var productRanks map[string]int
	// var productRanks map[string]int = make(map[string]int)
	// var productRanks map[string]int = map[string]int{}
	// productRanks := map[string]int{}
	// productRanks := map[string]int{"Pencil": 1, "Marker": 2}
	productRanks := map[string]int{
		"Pencil": 1,
		"Marker": 2,
	}
	productRanks["Pen"] = 4
	fmt.Println(productRanks)
	fmt.Println("len(productRanks) =", len(productRanks))

	fmt.Println("Iterating a map")
	for key, value := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, value)
	}

	keyToCheck := "Notebook"
	if val, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, val)
	} else {
		fmt.Printf("Key %q does not exist\n", keyToCheck)
	}

	keyToRemove := "Notebook"
	delete(productRanks, keyToRemove)
	fmt.Println("After removing Notebook :", productRanks)
}
