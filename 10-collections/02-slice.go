package main

import "fmt"

func main() {
	nos := []int{3, 1, 4, 2, 5}

	fmt.Println(nos)

	fmt.Println("Iterating an slice (using indexer)")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating an slice (using range)")
	for i, val := range nos {
		fmt.Printf("nos[%d] = %d\n", i, val)
	}

	newList := nos
	newList[0] = 1000
	fmt.Printf("nos = %v, newList = %v\n", nos, newList)

	//
	// nos = append(nos, 9)
	// nos = append(nos, 9, 11, 14, 15)
	newNos := []int{13, 16, 12, 15}
	nos = append(nos, newNos...)
	fmt.Println("After adding to nos")
	fmt.Printf("len(nos) = %v, len(newList) = %v\n", len(nos), len(newList))
	fmt.Printf("nos = %v, newList = %v\n", nos, newList)
}
