package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, s := range nums {
		if s%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}

}
