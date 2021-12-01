package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(filepath.Join(cwd, "depth_input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	previous := ""
	increases := 0
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		current := reader.Text()
		if strings.Compare(previous, current) == -1 {
			increases++
		}
		previous = current
	}
	fmt.Printf("The depth of the ocean floor increases %d times.\n", increases)
}
