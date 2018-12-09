package main

import "fmt"

func main() {
	game(9, 25)
}

const initPos = 1

func game(players, maxMarbles int) (winner int) {
	board := []int{0}
	pos := 0
	p := 1

	for m := 0; m < maxMarbles; m++ {
		if p == players {
			p = 1
		}

		pos, board = move(m+1, pos, board)
		fmt.Printf("[%d] %v\n", p, board)
		pos++
		p++
	}
	//fmt.Println(board)

	return 0
}

func move(val, current int, slice []int) (int, []int) {
	if len(slice) == 1 {
		slice = append(slice, val)
		return 1, slice
	}
	if len(slice) == current {
		pos := 1
		return pos, insert(val, 1, 2, slice)
	}

	nextPos := current + 1
	return nextPos, insert(val, nextPos, nextPos+1, slice)
}

func insert(val, start, end int, slice []int) []int {
	slice = append(slice, 0)
	copy(slice[end:], slice[start:])
	slice[start] = val

	return slice
}
