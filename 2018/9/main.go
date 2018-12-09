package main

import (
	"flag"
	"fmt"
)

func main() {
	players := flag.Int("players", 9, "Select the elf to play")
	marbles := flag.Int("marbles", 25, "Select the max marbles to play")
	flag.Parse()

	w, p := Game(*players, *marbles)
	fmt.Printf("The winner is player %d, with %d points\n", w, p)
}

type Marble struct {
	val      int
	next     *Marble
	previous *Marble
}

func Game(players, maxMarbles int) (winner, points int) {
	nextMarble := 0
	var leaderBoard = make([]int, players+1)

	current := &Marble{val: nextMarble}
	nextMarble++

	current.next = &Marble{val: nextMarble, next: current, previous: current}
	current.previous = current.next
	nextMarble++

	player := 2
	for ; nextMarble <= maxMarbles; nextMarble++ {
		if player > players {
			player = 1
		}

		var points int
		points, current = move(current, nextMarble)
		leaderBoard[player] = leaderBoard[player] + points
		player++
	}

	return calculateWinner(leaderBoard)
}

func move(marble *Marble, nextMarble int) (points int, m *Marble) {
	if (nextMarble % 23) == 0 {
		for i := 0; i < 7; i++ {
			marble = marble.previous
		}

		marbleForRemoval := marble
		marbleForRemoval.next.previous = marbleForRemoval.previous
		marbleForRemoval.previous.next = marbleForRemoval.next
		points = nextMarble + marble.val
		marble = marbleForRemoval.next

		return points, marble
	}

	newMarble := &Marble{
		val:      nextMarble,
		next:     marble.next.next,
		previous: marble.next,
	}
	newMarble.previous.next = newMarble
	newMarble.next.previous = newMarble
	marble = newMarble
	return points, marble
}

func calculateWinner(leaderBoard []int) (winner, points int) {
	var b int
	for player, p := range leaderBoard {
		if p > b {
			b = p
			winner = player
		}
	}

	return winner, b
}
