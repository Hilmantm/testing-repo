package main

import "fmt"

const PARTAI_MAX = 1000000
type Partai struct {
	nama string
	suara int
}
type TabPartai [PARTAI_MAX]Partai

func main() {

	var tabPartai, resultSort TabPartai
	var n int

	partaiInput(&tabPartai, &n)

	sortPartai(tabPartai, n, &resultSort)

	printPartai(resultSort, n)

}

func printPartai(resultSort TabPartai, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Println(resultSort[i].nama,"(",resultSort[i].suara, ")")
	}
}

func sortPartai(tabPartai TabPartai, n int, resultSort *TabPartai) {

	var i, pass int
	var temp Partai

	*resultSort = tabPartai

	pass = 1
	for pass <= n - 1 {
		i = pass
		temp = resultSort[pass]
		for i > 0 && temp.suara > resultSort[i - 1].suara {
			resultSort[i] = resultSort[i - 1]
			i--
		}
		resultSort[i] = temp
		pass++
	}

}

func partaiInput(tabPartai *TabPartai, n *int) {
	var i, foundIndex int
	var partai Partai

	i = 0
	fmt.Scanln(&partai.nama)
	for i < PARTAI_MAX && partai.nama != "-1" {
		foundIndex = posisi(&*tabPartai, &*n, partai.nama)
		if foundIndex != -1 {
			tabPartai[foundIndex].suara += 1
		} else {
			tabPartai[*n].nama = partai.nama
			tabPartai[*n].suara = 1
			*n++
		}
		i++
		fmt.Scanln(&partai.nama)
	}
}

func posisi(tabPartai *TabPartai, n *int, nama string) int {
	var i, result int
	i = 0
	result = -1
	for i < *n && result == -1 {
		if tabPartai[i].nama == nama {
			result = i
		}
		i++
	}
	return result
}
