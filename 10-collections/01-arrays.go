/* arrays */
package main

import "fmt"

func main() {
	// var nos [5]int
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos [5]int = [5]int{3, 1, 4}
	// var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating an array (using indexer)")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating an array (using range)")
	for i, val := range nos {
		fmt.Printf("nos[%d] = %d\n", i, val)
	}

	newList := nos
	newList[0] = 1000
	fmt.Printf("nos = %v, newList = %v\n", nos, newList)

	updateToTens(&nos)
	fmt.Println(nos)

}

func updateToTens(list *[5]int) {
	for i := 0; i < 5; i++ {
		list[i] = list[i] * 10
	}
}
