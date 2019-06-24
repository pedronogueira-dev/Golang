package main

import "fmt"

func main() {
	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range sl {
		var res string
		if isEven(num) {
			res = "even"
		} else {
			res = "odd"
		}

		fmt.Printf("%v is %v\n", num, res)
	}
}
