package main

import "fmt"

func main() {
	//if else
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	if no := 21; no%2 == 0 {
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}

	//switch case
	/*
		rank by score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Very Good"
		score 10 	=> "Excellent"
		otherwise 	=> "Invalid Score"
	*/

	/*
		score := 6
		switch score {
		case 0:
			fmt.Println("Terrible")
		case 1:
			fmt.Println("Terrible")
		case 2:
			fmt.Println("Terrible")
		case 3:
			fmt.Println("Terrible")
		case 4:
			fmt.Println("Not Bad")
		case 5:
			fmt.Println("Not Bad")
		case 6:
			fmt.Println("Not Bad")
		case 7:
			fmt.Println("Not Bad")
		case 8:
			fmt.Println("Very Good")
		case 9:
			fmt.Println("Very Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid Score")

		}
	*/

	/*
		score := 6
		switch score {
		case 0, 1, 2, 3:
			fmt.Println("Terrible")
		case 4, 5, 6, 7:
			fmt.Println("Not Bad")
		case 8, 9:
			fmt.Println("Very Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid Score")
		}
	*/

	/*
		switch score := 6; score {
		case 0, 1, 2, 3:
			fmt.Println("Terrible")
		case 4, 5, 6, 7:
			fmt.Println("Not Bad")
		case 8, 9:
			fmt.Println("Very Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid Score")
		}
	*/

	switch no := 7; {
	case no%2 == 0:
		fmt.Printf("%d is divisible by 2\n", no)
	case no%3 == 0:
		fmt.Printf("%d is divisible by 3\n", no)
	case no%4 == 0:
		fmt.Printf("%d is divisible by 4\n", no)
	case no%5 == 0:
		fmt.Printf("%d is divisible by 5\n", no)
	case no%6 == 0:
		fmt.Printf("%d is divisible by 6\n", no)
	case no%7 == 0:
		fmt.Printf("%d is divisible by 7\n", no)
	case no%8 == 0:
		fmt.Printf("%d is divisible by 8\n", no)
	case no%9 == 0:
		fmt.Printf("%d is divisible by 9\n", no)
	}

	fmt.Println("Switch-Case with fallthrough")
	switch n := 3; n {
	case 0:
		fmt.Println("n = 0")
	case 1:
		fmt.Println("n <= 1")
		fallthrough
	case 2:
		fmt.Println("n <= 2")
		fallthrough
	case 3:
		fmt.Println("n <= 3")
		fallthrough
	case 4:
		fmt.Println("n <= 4")
		fallthrough
	case 5:
		fmt.Println("n <= 5")
		fallthrough
	case 6:
		fmt.Println("n <= 6")

	case 7:
		fmt.Println("n <= 7")
		fallthrough
	case 8:
		fmt.Println("n <= 8")
		fallthrough
	case 9:
		fmt.Println("n <= 9")
		fallthrough
	case 10:
		fmt.Println("n <= 10")
	}

	fmt.Println("Switch Case with fallthrough (applied)")
	switch plan := "Pro"; plan {
	case "Super":
		fmt.Println("[Super] - for a family of 4")
		fallthrough
	case "Premium":
		fmt.Println("[Premium] - Download")
		fmt.Println("[Premium] - Playlist")
		fallthrough
	case "Pro":
		fmt.Println("[Pro] - No ads!")
		fallthrough
	case "Free":
		fmt.Println("[Free] - Listen to songs")
	}

	fmt.Println("for construct")
	fmt.Println("for [v1.0]")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("for [v2.0](while)")
	/*
		sum := 1
		for sum <= 100 {
			sum += sum
		}
	*/
	for sum := 1; sum <= 100; {
		sum += sum
		fmt.Println("sum = ", sum)
	}

	fmt.Println("for [v3.0] (infinite)")
	numSum := 1
	/*
		for {
			numSum += numSum
			if numSum > 100 {
				break
			}
		}
	*/
	for {
		numSum += numSum
		if numSum < 100 {
			continue
		}
		break
	}
	fmt.Println("numSum =", numSum)

	fmt.Println("for [v4.0] (using labels)")

LOOP:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				fmt.Println("=============================")
				// break
				continue LOOP
			}
		}
	}
}
