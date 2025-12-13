package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"unicode"
	// "strconv"
	// "strings"
)

var count int

func stringToDigits(s string) []int {
	digits := make([]int, 0, len(s))
	for _, r := range s {
		if unicode.IsDigit(r) {
			digits = append(digits, int(r-'0'))
		}
	}
	return digits
}

// func xyToJoltage(x,y string) int {
// 	str, err := strconv.Atoi(strings.Join(x, y))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return str
// }

func read_file(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		digits := stringToDigits(scanner.Text())
		if len(digits) < 2 { break }
		x, y := digits[0], digits[1]
		for i := 2; i <= len(digits) - 2; i++ {
			if x < digits[i] {
				x = digits[i]
				y = 0
			} else if y < digits[i] {
				y = digits[i]
			}
		}
		if y < digits[len(digits)-1] { y = digits[len(digits)-1] }
		//fmt.Println("Input string", digits)
		fmt.Println("Got digits", x, y)
		count += x*10 + y
		fmt.Println("Current", count)
	}
}

func main(){
	read_file("input.txt")
	fmt.Println("Counted joltage: ", count)
}
