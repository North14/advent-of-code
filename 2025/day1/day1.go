package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
)

var dial int = 50
var count int = 0

func rotate_dial(rot string) {
	d := rot[0]
	c, err := strconv.Atoi(rot[1:])

	if err != nil {
		log.Printf("Skipping invalid line '%s': %v", rot, err)
		return
	}

	if d == 'R' { // 76 = L, 82 = R
		dial += c
	} else if d == 'L' {
		dial -= c
	} else {
		log.Fatal("Unknown input, not L or R")
	}
}

func read_file(filename string) *bufio.Scanner {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil
	}
	for scanner.Scan() {
		rotate_dial(scanner.Text())
		// fmt.Println(dial)
		if dial % 100 == 0 {
			count++
		}
	}
	return scanner
}

func main(){
	read_file("input.txt")
	fmt.Println("Count is:", count)
}
