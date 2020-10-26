package main

import "fmt"

func main() {
	numbers := [...]int{1, 4, 10, 6, 2}

	// descending order
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] > numbers[i] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
				continue
			}
		}
	}
	fmt.Println(numbers)

	//ascending order

	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[i] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
				continue
			}
		}
	}
	fmt.Println(numbers)
}
