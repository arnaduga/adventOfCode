package main

// Good answer: 1606483

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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

	var area int = 0
	var ribbon int = 0

	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {

		dim := strings.Split(scanner.Text(), "x")
		var dimInt []int

		for i := range dim {
			t, _ := strconv.Atoi(dim[i])
			dimInt = append(dimInt, t)
		}

		sort.Ints(dimInt)

		s1 := dimInt[0] * dimInt[1]
		s2 := dimInt[1] * dimInt[2]
		s3 := dimInt[0] * dimInt[2]
		thisGift := dimInt[0]*dimInt[1] + 2*s1 + 2*s2 + 2*s3
		area += thisGift

		t1_rib := dimInt[0]*2 + dimInt[1]*2
		t2_rib := dimInt[0] * dimInt[1] * dimInt[2]
		ribbon += t1_rib + t2_rib
	}

	fmt.Printf("Square foot wrapping paper: %v\n", area)
	fmt.Printf("Ribbon length: %v\n", ribbon)
}
