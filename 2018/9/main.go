package main

import (
	"flag"
	"fmt"
)

func main() {
	debugMode := flag.Bool("debug", false, "Print the movements and leaderboard on console")
	players := flag.Int("players", 9, "Select the elf to play")
	marbles := flag.Int("marbles", 25, "Select the max marbles to play")
	flag.Parse()

	w, p := Game(*players, *marbles, *debugMode)
	fmt.Printf("The winner is player %d, with %d points\n", w, p)
}

func Game(players, maxMarbles int, debug bool) (winner, points int) {
	var leaderBoard = make(map[int]int)

	board := []int{0}
	pos := 0
	p := 1

	for m := 0; m < maxMarbles; m++ {
		if p > players {
			p = 1
		}

		var points int
		pos, points, board = move(m+1, pos, board)
		leaderBoard[p] = leaderBoard[p] + points
		if debug {
			fmt.Printf("[%d] %v\n", p, board)
		}

		pos++
		p++
	}

	if debug {
		fmt.Println("=== LEADERBOARD ===")
		for player, points := range leaderBoard {
			fmt.Printf("[%d] %d\n", player, points)
		}
	}

	return calculateWinner(leaderBoard)
}

func calculateWinner(leaderBoard map[int]int) (winner, points int) {
	var b int
	for player, p := range leaderBoard {
		if p > b {
			b = p
			winner = player
		}
	}

	return winner, b
}

func move(val, current int, slice []int) (nextPos, points int, board []int) {
	if len(slice) == 1 {
		slice = append(slice, val)
		return 1, points, slice
	}

	if len(slice) == current {
		pos := 1
		return pos, points, insert(val, 1, 2, slice)
	}

	nextPos = current + 1
	if (val % 23) == 0 {
		pos := current - 8
		if pos <= 0 {
			pos = pos + len(slice)
		}
		points = val + slice[pos]
		return pos, points, remove(pos, pos+1, slice)
	}
	return nextPos, points, insert(val, nextPos, nextPos+1, slice)
}

func insert(val, start, end int, slice []int) []int {
	slice = append(slice, 0)
	copy(slice[end:], slice[start:])
	slice[start] = val

	return slice
}

func remove(start, end int, slice []int) []int {
	return append(slice[:start], slice[end:]...)
}
