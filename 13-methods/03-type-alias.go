package main

import "fmt"

type MyString string // type alias

func (s MyString) Length() int {
	return len(s)
}

func main() {
	str := MyString("Mollit aliqua ad incididunt reprehenderit aliqua velit proident. Enim exercitation ipsum aute aliqua minim ex. Sit incididunt exercitation Lorem minim sint et esse. Non veniam anim officia ad commodo est.")
	fmt.Println(str.Length())
}
