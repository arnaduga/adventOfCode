package main

// https://adventofcode.com/2015/day/1

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args[1:]) != 1 {
		log.Fatal("Missing the arg filename")
	}

	arg := os.Args[1]
	fil, err := os.Open(arg)

	if err != nil {
		log.Fatal("Error reading the file.")
	}

	scanner := bufio.NewScanner(fil)

	var f int = 0
	var firstTime bool = false

	for scanner.Scan() {
		readLine := scanner.Text()
		for i := range readLine {
			c := fmt.Sprintf("%c", readLine[i])
			if c == "(" {
				f = f + 1
			} else {
				f = f - 1
			}
			if f < 0 && !firstTime {
				firstTime = true
				fmt.Printf("Santa reached the basement at step %v\n", i+1)
			}
		}
	}

	fmt.Printf("I found Santa in the floor %v", f)

}
