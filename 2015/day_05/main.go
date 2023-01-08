package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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

func atLeastThreeVowels(source string) bool {

	writeLog(DEBUG, fmt.Sprintf("3 vowels test on %v", source))

	regx := regexp.MustCompile("a|e|i|o|u")
	matches1 := regx.FindAllStringIndex(source, -1)

	if len(matches1) >= 3 {
		writeLog(DEBUG, fmt.Sprintf(" ok "))
		return true
	}

	writeLog(DEBUG, fmt.Sprintf(" KO "))
	return false

}

func repetedChar(source string) bool {
	writeLog(DEBUG, fmt.Sprintf("Repeated char on %v", source))
	repeated := 0
	var previous rune
	for _, l := range source {
		if l == previous {
			repeated = repeated + 1
		}
		previous = l
	}
	if repeated > 0 {
		writeLog(DEBUG, "OK")
		return true
	}
	writeLog(DEBUG, "KO")
	return false
}

func noForbiddenPairs(source string) bool {
	writeLog(DEBUG, fmt.Sprintf("Forbidden pairs in %v", source))
	regxnot := regexp.MustCompile("ab|cd|pq|xy")
	matches := regxnot.FindAllStringIndex(source, -1)
	if len(matches) == 0 {
		writeLog(DEBUG, "OK")
		return true
	}
	writeLog(DEBUG, "KO")
	return false
}

func repetedCharWithGap(source string) bool {
	writeLog(DEBUG, fmt.Sprintf("Repeated char with gap on %v", source))
	repeated := 0
	var previous, beforeprevious rune
	for _, l := range source {
		if l == beforeprevious {
			repeated = repeated + 1
		}
		beforeprevious = previous
		previous = l
	}
	if repeated > 0 {
		writeLog(DEBUG, "OK")
		return true
	}
	writeLog(DEBUG, "KO")
	return false
}

func hasMoreThanOnePair(source string) bool {
	writeLog(DEBUG, fmt.Sprintf("More than a pair %v", source))

	for i := range source {
		if i > 0 {
			p := fmt.Sprintf("%c%c", source[i-1], source[i])
			if strings.Count(source, p) > 1 {
				writeLog(DEBUG, "OK")
				return true
			}
		}
	}

	writeLog(DEBUG, "KO")
	return false
}

func main() {

	startTime := time.Now()

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

	writeLog(DEBUG, fmt.Sprintf("File to be opened: %v", *input))

	fil, err := os.Open(*input)
	if err != nil {
		log.Fatal("Error reading the file.")
	}

	niceWords1 := []string{}
	naughtyWords1 := []string{}

	niceWords2 := []string{}
	naughtyWords2 := []string{}

	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		word := scanner.Text()
		// Applying- ruleset #1
		if atLeastThreeVowels(word) && repetedChar(word) && noForbiddenPairs(word) {
			writeLog(DEBUG, fmt.Sprintf("%v is a NICE word", word))
			niceWords1 = append(niceWords1, word)
		} else {
			writeLog(DEBUG, fmt.Sprintf("%v is a NAUGHTY word", word))
			naughtyWords1 = append(naughtyWords1, word)
		}

		//Applying ruleset #2
		if repetedCharWithGap(word) && hasMoreThanOnePair(word) {
			writeLog(DEBUG, fmt.Sprintf("%v is a NICE word", word))
			niceWords2 = append(niceWords2, word)
		} else {
			writeLog(DEBUG, fmt.Sprintf("%v is a NAUGHTY word", word))
			naughtyWords2 = append(naughtyWords2, word)
		}
	}

	writeLog(DEBUG, fmt.Sprintf("%v", niceWords1))
	writeLog(INFO, fmt.Sprintf("Nice words - Ruleset #1: %v", len(niceWords1)))
	writeLog(DEBUG, fmt.Sprintf("%v", naughtyWords1))
	writeLog(INFO, fmt.Sprintf("Naughty words - Ruleset #1: %v", len(naughtyWords1)))

	writeLog(DEBUG, fmt.Sprintf("%v", niceWords2))
	writeLog(INFO, fmt.Sprintf("Nice words - Ruleset #2: %v", len(niceWords2)))
	writeLog(DEBUG, fmt.Sprintf("%v", naughtyWords2))
	writeLog(INFO, fmt.Sprintf("Naughty words - Ruleset #2: %v", len(naughtyWords2)))

	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	writeLog(INFO, fmt.Sprintf("Elapsed time: %v", elapsed))

}
