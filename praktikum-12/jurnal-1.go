package main

import "fmt"

type Player struct {
	name string
	gol, assist int
}

const NMAX = 100
type Players [NMAX]Player

func main() {

	var players, sortResult Players
	var totalPlayer int

	playerInput(&players, &totalPlayer)

	sortPlayer(players, totalPlayer, &sortResult)

	printPlayer(sortResult, totalPlayer)
}

func printPlayer(sortResult Players, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Println(sortResult[i].name, sortResult[i].gol, sortResult[i].assist)
	}
}

func playerInput(players *Players, totalPlayer *int) {
	// Input player hanya 1 nama saja alias tidak ada nama belakang
	var i int
	var player Player
	fmt.Scanln(&*totalPlayer)
	for i = 0; i < *totalPlayer; i++ {
		fmt.Scanln(&player.name, &player.gol, &player.assist)
		players[i] = player
	}
}

func sortPlayer(players Players, totalPlayer int, sortPlayer *Players) {

	var i, pass int
	var temp Player

	*sortPlayer = players

	pass = 1
	for pass <= totalPlayer - 1 {
		i = pass
		temp = sortPlayer[pass]
		for i > 0 && (temp.gol > sortPlayer[i - 1].gol || temp.assist > sortPlayer[i - 1].assist) {
			sortPlayer[i] = sortPlayer[i - 1]
			i--
		}
		sortPlayer[i] = temp
		pass++
	}

}
