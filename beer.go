package main

import (
	"fmt"
)

func main() {
	for i := 99; i > 0; i-- {
		fmt.Printf("%d bottle%s of beer on the wall, %d bottle%s of beer.\n", i, plural(i), i, plural(i))
		fmt.Printf("Take one down and pass it around, ")

		if i-1 > 0 {
			fmt.Printf("%d bottle%s of beer on the wall.\n\n", i-1, plural(i-1))
		} else {
			fmt.Println("no more bottles of beer on the wall.\n")
		}
	}

	fmt.Println("No more bottles of beer on the wall, no more bottles of beer.")
	fmt.Println("Go to the store and buy some more, 99 bottles of beer on the wall.")
}

func plural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}
