package main

import (
	"fmt"
	"io"
	"os"
)


func PrintInput(file *os.File) {
    b, err := io.ReadAll(file)
    checkError(err)
    fmt.Print(b)
}


func checkError(e error) {
    if e != nil {
        panic(e)
    }
}

func AbsDiffInt(x, y int) int {
    if x < y {
       return y - x
    }
    return x - y
}