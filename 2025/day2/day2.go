package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"encoding/csv"
	"strconv"
	"reflect"
)

var invalids []int

func sumWithForLoop(numbers []int) int {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}

func reverseInt(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
            s[i], s[j] = s[j], s[i]
    }
}

func splitToDigits(n int) []int{
    var ret []int
    
    for n !=0 {
        ret = append(ret, n % 10)
        n /= 10
    }
    
    reverseInt(ret)
    
    return ret
}


func read_file(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	for _, record := range records[0] {
		ranges := strings.Split(record, "-")
		if len(ranges) < 2 {
			continue 
		}
		start, err := strconv.Atoi(ranges[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(ranges[1])
		if err != nil {
			log.Fatal(err)
		}

		// Count the length to split in
		    // Initialize a counter for the digits
		count := 0
		copy_end := end
		// Count the digits
		for copy_end != 0 {
		    copy_end /= 10
		    count++
		}
		count /= 2

		fmt.Println("Iterating from", start, "to", end)
		for i := start; i < end; i++ {
			digits := splitToDigits(i)//[j:len(end)])
			if len(digits) < 2 { continue }
			for j := 1 ; j < count + 1; j++ {
				var digit_groups [][]int
				for z := 0; z < len(digits); z += j {
					endIdx := z + j
					if endIdx > len(digits) {
						endIdx = len(digits)
					}
					// Append the chunk to our groups
					digit_groups = append(digit_groups, digits[z:endIdx])
				}
				dummy := true
				for l := 0; l < len(digit_groups)-1; l++ {
					if !reflect.DeepEqual(digit_groups[l], digit_groups[l+1]) {
						//fmt.Println("found equal", digit_groups[l], "and", digit_groups[l+1])

						dummy = false
					}
				}
				if dummy == true {
					fmt.Println("Found invalid number: ", i)
					invalids = append(invalids, i)
					break
				}
				//fmt.Printf("Number: %d | Split Size: %d | Result: %v\n", i, j, digit_groups)
				
			}
		}
	}
}

func main(){
	read_file("input.txt")
	fmt.Println(invalids)
	fmt.Println("Sum:", sumWithForLoop(invalids))
}
