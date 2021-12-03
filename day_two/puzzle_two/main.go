package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(filepath.Join(cwd, "movement_input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	horizontal := 0
	depth := 0
	aim := 0
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		input := strings.Split(reader.Text(), " ")
		value, err := strconv.Atoi(input[1])
		if err != nil {
			log.Fatal(err)
		}
		switch input[0] {
		case "forward":
			horizontal += value
			depth += aim * value
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}
	fmt.Printf("The current position of the submarine is %d.\n", horizontal * depth)
}