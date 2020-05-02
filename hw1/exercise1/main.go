package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 1, 2, 3}
	solution := counter(numbers)
	fmt.Println(solution)
}

func counter(numbers []int) map[int]int {
	counted := make(map[int]int)
	for _, number := range numbers {
		if _, ok := counted[number]; ok {
			counted[number]++
		}else {
			counted[number] = 1
		}
	}
	return counted
}