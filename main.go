package main

import (
	"fmt"
	"os"

	"advent2024.lsakada.fr/day1"
)

func main() {
    day1.PrintMessage("hello sakada")
	getFileData("day1_txt")
}

//The function return nothing at first
//It will later return some generic data to be used every day
func getFileData(fileName string) {
	path := "./input/" + fileName
	file, err := os.Open(path)

	if err != nil {
		fmt.Printf("getFileData - Error opening file: %v\n", err)
	}

	defer file.Close()
}