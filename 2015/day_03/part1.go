package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var debugMode *bool

func logDebug(v string) {
	if *debugMode {
		log.Println(v)
	}
}

func newLocation(x int, y int, d string) (int, int) {
	var new_x, new_y int
	new_x = x
	new_y = y
	switch d {
	case ">":
		new_x++
	case "<":
		new_x--
	case "^":
		new_y++
	case "v":
		new_y--
	}

	return new_x, new_y
}

func main() {
	// Reading the debug flag
	debugMode = flag.Bool("debug", false, "Display more debug logs")
	flag.Parse()

	// Reading the arg
	values := flag.Args()

	if len(values) != 1 {
		log.Fatal("Missing the file to open as argument")
		os.Exit(1)
	}

	logDebug("Debug mode activated")

	fil, err := os.Open(values[0])
	if err != nil {
		log.Fatal("Error reading the file.")
	}
	defer fil.Close()

	// Setting a cursor (where is Santa)
	var santa_x, santa_y int
	var index string
	santa_x = 0
	santa_y = 0

	// Preparing the MAP :)
	tracker := make(map[string]int)

	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		rawText := scanner.Text()
		for _, w := range rawText {

			santa_x, santa_y = newLocation(santa_x, santa_y, string(w))
			index = fmt.Sprintf("%d_%d", santa_x, santa_y)

			// Calculate an index (x_y)
			tracker[index]++
		}
	}
	fmt.Println("Number of houses visited:", len(tracker))
}
