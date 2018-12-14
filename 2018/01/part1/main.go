package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	sum := calculate("input.txt")
	fmt.Println(sum)
}

func calculate(input string) int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error reading file %v", err)
	}

	var sum int
	s := bufio.NewScanner(f)
	for s.Scan() {
		var n int
		fmt.Sscanf(s.Text(), "%d", &n)
		sum += n
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	return sum

}
