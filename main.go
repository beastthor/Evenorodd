package main

import "fmt"

func main() {

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, number := range numbers {

		//if a number divided by 2 ,remainder is 0 then that number is even or else it is odd

		if number%2 == 0 {
			fmt.Println(i, number, "is even")
		} else {
			fmt.Println(i, number, "is odd")
		}

	}
}
