package main

/*
** Call example:
**     go run main.go --part 1 --key abcdef
**     go run main.go --part 2 --key pqrstuv
 */

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
	FATAL
)

func writeLog(messagetype messageType, message string) {
	switch messagetype {
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

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {

	startTime := time.Now()

	part := flag.Int("part", 0, "Exercice part (should 1 or 2)")
	key := flag.String("key", "", "Key for calculation")
	flag.Parse()

	// Reading the arg
	if *part == 0 && *key == "" {
		flag.Usage()
		os.Exit(1)
	}

	var toTest string

	var pattern string

	if *part == 1 {
		pattern = "00000"
	} else if *part == 2 {
		pattern = "000000"
	} else {
		writeLog(FATAL, "Nope... Only 2 parts")
		os.Exit(2)
	}

	var counter int = -1
	for {
		counter++
		toTest = fmt.Sprintf("%v%d", *key, counter)

		if strings.HasPrefix(GetMD5Hash(toTest), pattern) {
			writeLog(INFO, fmt.Sprintf("Key %v meet the requirement: %v", toTest, GetMD5Hash(toTest)))
			writeLog(INFO, fmt.Sprintf("Result in decimal is: %d", counter))
			break
		}
		if counter == 10000000 {
			writeLog(FATAL, "Circuit breaker reached.")
			os.Exit(1)
		}

	}

	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	writeLog(INFO, fmt.Sprintf("Elapsed time: %v", elapsed))

}
