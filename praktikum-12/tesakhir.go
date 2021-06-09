package main

import (
	"fmt"
	"math"
)

const MAX_PROVINCE = 100
type Province struct {
	nama string
	populasi int
	tumbuh float64
}
type Provinces struct {
	table [MAX_PROVINCE]Province
	nProvince int
}s

func main() {

	var provinces Provinces
	var name1, name2 string
	var findResult Province

	fmt.Println("Data pertumbuhan provinsi:")
	inputProvince(&provinces)

	fmt.Print("Nama Provinsi? ")
	fmt.Scanln(&name1)

	findResult = findProvince(provinces, name1)
	fmt.Println(findResult.nama, findResult.populasi, findResult.tumbuh)

	fmt.Print("Prediksi populasi tahun depan provinsi? ")
	fmt.Scanln(&name2)
	fmt.Println("Populasi Provinsi", name2, "tahun depan:", prediction(provinces, name2))

	fmt.Println("Urutan data pertumbuhan provinsi berdasarkan populasi terurut menurun:")
	sortProvince(&provinces)
	printProvinces(provinces)
}

func printProvinces(provinces Provinces) {
	var i int
	for i = 0; i < provinces.nProvince; i++ {
		fmt.Println(provinces.table[i].nama, provinces.table[i].populasi, provinces.table[i].tumbuh)
	}
}

func inputProvince(provinces *Provinces) {
	var i int
	var province Province
	fmt.Scanln(&province.nama, &province.populasi, &province.tumbuh)
	for i < MAX_PROVINCE && province.nama != "NONE" && province.populasi != 0 && province.tumbuh != 0.0 {
		provinces.table[provinces.nProvince] = province
		provinces.nProvince += 1
		i++
		fmt.Scanln(&province.nama, &province.populasi, &province.tumbuh)
	}
}

func findProvince(provinces Provinces, name string) Province {
	var result Province
	var i, resultIndex int

	// menggunakan sequential search
	resultIndex = -1
	for i < provinces.nProvince && resultIndex == -1 {
		if provinces.table[i].nama == name {
			result = provinces.table[i]
			resultIndex = i
		}
		i++
	}

	return result
}

func prediction(provinces Provinces, name string) int {
	var province Province
	var resultMultiplication float64
	province = findProvince(provinces, name)

	// Rumus = populasi * pertumbuhan (pertumbuhan dibulatkan keatas / ceiling)
	resultMultiplication = float64(province.populasi) * province.tumbuh
	return province.populasi + int(math.Ceil(resultMultiplication))
}

func sortProvince(provinces *Provinces) {
	var pass, i int
	var temp Province

	// sorting menggunakan insertion sort
	pass = 1
	for pass <= provinces.nProvince - 1 {
		i = pass
		temp = provinces.table[pass]
		for i > 0 && temp.populasi > provinces.table[i - 1].populasi {
			provinces.table[i] = provinces.table[i - 1]
			i--
		}
		provinces.table[i] = temp
		pass++
	}
}
