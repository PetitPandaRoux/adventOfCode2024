package main

import (
	"fmt"
	"os"
	"sort"

	"advent2024.lsakada.fr/day1"
)

func main() {
	getFileData("day_1_full.txt")
}

//It will later return some generic data to be used every day
func getFileData(fileName string) {
	path := "./input/" + fileName
	file, err := os.Open(path)

	if err != nil {
		fmt.Printf("getFileData - Error opening file: %v\n", err)
	}

	defer file.Close()

	firstArray, secondArray, err := day1.ReadArray(file)

	sumOfScore := computeSimilaryScore(firstArray, secondArray)

	if err != nil {
		fmt.Printf("ReadArray - Error parsing file to array : %v\n", err)
	}

	fmt.Println(firstArray)
	fmt.Println(secondArray)
	fmt.Println(sumOfScore)
}

func computeSimilaryScore(firstArray []int, secondArray []int) (int) {
	sort.Ints(firstArray)
	sort.Ints(secondArray)

	scoreMemory := map[int]int {}
	sumOfScore := 0

	for i := 0; i < len(firstArray); i += 1 {
		if scoreMemory[firstArray[i]] != 0 {
			sumOfScore += scoreMemory[firstArray[i]]
			continue
		} 

		score := 0

		for j := 0; j < len(secondArray); j += 1 {
			if secondArray[j] == firstArray[i] {
				score += 1
			}

			if secondArray[j] > firstArray[i] {
				sumOfScore += score * firstArray[i]
				scoreMemory[firstArray[i]] = score * firstArray[i]
				break
			}
		}
	}
	fmt.Println(scoreMemory)

	return sumOfScore
}

func sumArrayDiff(firstArray []int, secondArray []int) (int) {
	sort.Ints(firstArray)
	sort.Ints(secondArray)

	sumOfDiff := 0

	if len(firstArray) != len(secondArray) {
		fmt.Printf(("Two arrays does not have the same length"))
		return 0
	}

	for i := 0; i < len(firstArray); i += 1 {
		sumOfDiff += AbsDiffInt(firstArray[i], secondArray[i])
	}
	return sumOfDiff
}