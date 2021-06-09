package main

import "fmt"

const NMAX = 20
type tabInt [NMAX]int

func main() {

	var numbers, resultSort tabInt
	var n int

	numbers = tabInt{2, 4, 1, 0, 10, 40, 12, 21, 22, 40, 3, 8, 5}
	n = 13

	insertionSort(numbers, n, &resultSort)

	fmt.Println(resultSort)

}

func insertionSort(numbers tabInt, n int, result *tabInt) {

	var pass, i, temp int

	*result = numbers

	pass = 1
	for pass <= n - 1 {
		i = pass
		temp = result[pass]
		for i > 0 && temp > result[i - 1] {
			result[i] = result[i - 1]
			i--
		}
		result[i] = temp
		pass++
	}

}
