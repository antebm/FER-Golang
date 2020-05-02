package main

import "fmt"

func main() {
	numbers := [][]int{
		[]int{3,10,4,2},
		[]int{9,3,8,7},
		[]int{15,14,13,12},
	}
	solution := sretanBroj(numbers)
	fmt.Println(solution)
}

func sretanBroj (matrica [][]int) int {
	luckyNumber := false
	index := 0

	for i := 0; i < len(matrica); i++ {

		min := matrica[i][0]

		for j := 0; j < len(matrica[0]); j++ {
			if(matrica[i][j] < min){
				index = j
				min = matrica[i][j]
			}
		}

		for k := 0; k < len(matrica); k++ {
			if matrica[k][index] > min {
				luckyNumber = false
				break;
			}
			luckyNumber = true
		}
		if luckyNumber {
			return matrica[i][index]
		}
	}

	return 0
}