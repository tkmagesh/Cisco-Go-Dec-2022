package main

import "fmt"

func main() {
	// var nos []int = make([]int, 5)
	// var nos []int
	var nos []int = make([]int, 0 /* initialized */, 5 /* allocated */)
	nos = append(nos, 10)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 20)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 30)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 40)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 50)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 60)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)

	fmt.Println("Slicing")
	fmt.Println("nos[1:4] = ", nos[1:4])
	fmt.Println("nos[1:] = ", nos[1:])
	fmt.Println("nos[:4] = ", nos[:4])

	nosCopy := make([]int, len(nos))
	copy(nosCopy, nos)

	/*
		nosCopy := nos
	*/
	nosCopy[0] = 10000
	fmt.Println(nos, nosCopy)
}
