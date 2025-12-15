package main

import (
	"fmt"
	"os"
	"log"
	"strings"
)

// solution inspired by
// https://github.com/tiennm99/adventofcode/blob/main/2025/day4/day4.go

var neighbours = []Vector {
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

type Vector struct {
	X, Y int
}

func accessable(input []string, v Vector) bool {
	if input[v.X][v.Y] != '@' {
		return false
	}
	count := 0
	for _, neighbour := range neighbours {
		x := v.X + neighbour.X
		y := v.Y + neighbour.Y
		if x < 0 || x >= len(input) || y < 0 || y >= len(input[x]) {
			continue
		}
		if input[x][y] == '@' {
			count++
		}
	}
	return count < 4
}


func counting(input []string) int {
	count := 0
	for i, l := range input {
		for j := range l {
			if accessable(input, Vector{i, j}) {
				count++
			}
		}
	}

	return count
}

func readFile(filename string) int {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	
	return counting(strings.Split(string(file), "\n"))
}

func main(){
	score := readFile("input.txt")
	fmt.Println("Score is:", score)
}
