package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PrintMessage(message string) {
    fmt.Println(message)
}


func ReadArray(file *os.File) ([]int, []int, error) {
	scanner := bufio.NewScanner(file)
	var firstArray, secondArray []int

	for scanner.Scan() {
		// Split the line into parts
		line := scanner.Text()
		parts := strings.Fields(line)

		// Parse the numbers and append them to respective slices
		if len(parts) == 2 {
			num1, err1 := strconv.Atoi(parts[0])
			num2, err2 := strconv.Atoi(parts[1])

			if err1 == nil && err2 == nil {
				firstArray = append(firstArray, num1)
				secondArray = append(secondArray, num2)
			} else {
				fmt.Println("Error parsing numbers:", parts)
			}
		} else {
			fmt.Println("Unexpected line format:", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, nil, err 
	}

	return firstArray, secondArray, nil
}