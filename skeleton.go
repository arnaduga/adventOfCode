package main

import (
	"bufio"
	"flag"
	"log"
	"os"
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

var debugMode *bool

func logDebug(v string) {
	if *debugMode {
		log.Println(v)
	}
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

	//logDebug("Hello")

	fil, err := os.Open(values[0])
	if err != nil {
		log.Fatal("Error reading the file.")
	}
	defer fil.Close()

	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}

}
