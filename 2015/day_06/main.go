package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type messageType int

var debugMode *bool

const (
	DEBUG messageType = 0 + iota
	INFO
	WARNING
	ERROR
	FATAL
)

func writeLog(messagetype messageType, message string) {
	switch messagetype {
	case DEBUG:
		if *debugMode {
			logger := log.New(os.Stdout, "[DEBUG]   ", log.Ldate|log.Ltime|log.Lmicroseconds)
			logger.Println(message)
		}
	case INFO:
		logger := log.New(os.Stdout, "[INFO]    ", log.Ldate|log.Ltime|log.Lmicroseconds)
		logger.Println(message)
	case WARNING:
		logger := log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.Lmicroseconds)
		logger.Println(message)
	case ERROR:
		logger := log.New(os.Stderr, "[ERROR]   ", log.Ldate|log.Ltime|log.Lmicroseconds)
		logger.Println(message)
	case FATAL:
		logger := log.New(os.Stderr, "[FATAL]   ", log.Ldate|log.Ltime|log.Lmicroseconds)
		logger.Println(message)
	}
}

func extractAction(source string) (string, int, int, int, int) {
	var action string
	var x1, y1, x2, y2 int = 0, 0, 0, 0

	if strings.Contains(source, "toggle") {
		action = "toggle"
		fmt.Sscanf(source, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)
	} else {
		fmt.Sscanf(source, "turn %s %d,%d through %d,%d", &action, &x1, &y1, &x2, &y2)
	}
	return action, x1, y1, x2, y2
}

func applyInstruction(arr *[1000][1000]int, a string, x1 int, y1 int, x2 int, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {

			if a == "on" {
				arr[x][y] = 1
			}
			if a == "off" {
				arr[x][y] = 0
			}
			if a == "toggle" {
				if arr[x][y] == 0 {
					arr[x][y] = 1
				} else {
					arr[x][y] = 0
				}
			}

		}
	}
}

func applyInstruction2(arr *[1000][1000]int, a string, x1 int, y1 int, x2 int, y2 int) {

	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {

			if a == "on" {
				arr[x][y] = arr[x][y] + 1
			}
			if a == "off" {
				arr[x][y] = arr[x][y] - 1
				if arr[x][y] < 0 {
					arr[x][y] = 0
				}
			}

			if a == "toggle" {
				arr[x][y] = arr[x][y] + 2
			}

		}
	}
}

func countLightsOn(arr *[1000][1000]int) int {
	var c int = 0
	for x := range arr {
		for y := range arr[x] {
			if arr[x][y] == 1 {
				c = c + 1
			}
		}
	}
	return c
}

func countLightsBrightness(arr *[1000][1000]int) int {
	var c int = 0
	for x := range arr {
		for y := range arr[x] {
			c = c + arr[x][y]
		}
	}
	return c
}

func main() {

	startTime := time.Now()
	writeLog(INFO, fmt.Sprintf("Start time: %v", startTime))

	debugMode = flag.Bool("debug", false, "Display debug log if true")
	var input *string
	input = flag.String("input", "", "Input file")
	flag.Parse()
	writeLog(DEBUG, *input)

	// Reading the arg
	if *input == "" {
		flag.Usage()
		os.Exit(1)
	}

	//Init the matrix
	var matrix [1000][1000]int
	var brightness [1000][1000]int
	writeLog(INFO, "Initializing matrix")
	for x := range matrix {
		for y := range matrix[x] {
			matrix[x][y] = 0
			brightness[x][y] = 0
		}
	}

	writeLog(DEBUG, fmt.Sprintf("File to be opened: %v", *input))

	fil, err := os.Open(*input)
	if err != nil {
		log.Fatal("Error reading the file.")
	}

	action := "none"
	x1 := -1
	y1 := -1
	x2 := -1
	y2 := -1

	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		word := scanner.Text()
		action, x1, y1, x2, y2 = extractAction(word)
		applyInstruction(&matrix, action, x1, y1, x2, y2)
		applyInstruction2(&brightness, action, x1, y1, x2, y2)
	}
	writeLog(INFO, fmt.Sprintf("Number of lights ON: %d", countLightsOn(&matrix)))
	writeLog(INFO, fmt.Sprintf("Total brightness: %d", countLightsBrightness(&brightness)))

	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	writeLog(INFO, fmt.Sprintf("End time: %v", endTime))
	writeLog(INFO, fmt.Sprintf("Elapsed time: %v", elapsed))

}
