package main

import "testing"

type game struct {
	players int
	marbles int
}

var flagtests = []struct {
	in  string
	g   game
	out int
}{
	{"9 players; last marble is worth 25 points: high score is 32", game{9, 25}, 32},
	{"10 players; last marble is worth 1618 points: high score is 8317", game{10, 1618}, 8317},
	{"13 players; last marble is worth 7999 points: high score is 146373", game{13, 7999}, 146373},
	{"17 players; last marble is worth 1104 points: high score is 2764", game{17, 1104}, 2764},
	{"21 players; last marble is worth 6111 points: high score is 54718", game{21, 6111}, 54718},
	{"30 players; last marble is worth 5807 points: high score is 37305", game{30, 5807}, 37305},
	{"486 players; last marble is worth 70833 points: high score is 373597", game{486, 70833}, 373597},
	{"486 players; last marble is worth 7083300 points: high score is 373597", game{486, 7083300}, 2954067253},
}

func TestGame(t *testing.T) {
	for _, tt := range flagtests {
		t.Run(tt.in, func(t *testing.T) {
			_, p := Game(tt.g.players, tt.g.marbles)
			if p != tt.out {
				t.Errorf("got %d, want %d", p, tt.out)
			}
		})
	}
}
