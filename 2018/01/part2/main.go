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

	nums := []int{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		var n int
		fmt.Sscanf(s.Text(), "%d", &n)
		nums = append(nums, n)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	var sum int
	var reaches = make(map[int]bool)

	for {
		for _, n := range nums {
			sum += n
			if reaches[sum] {
				return sum
			}
			reaches[sum] = true
		}
	}
}
