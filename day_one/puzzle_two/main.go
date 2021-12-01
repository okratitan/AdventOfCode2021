package main

import (
	"fmt"
	"io/ioutil"
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
	reader, err := ioutil.ReadFile(filepath.Join(cwd, "depth_input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	values := strings.Split(string(reader), "\n")
	previous := 0
	increases := 0
	totalValues := len(values)
	for i, _ := range values {
		if i + 2 >= totalValues {
			break
		}
		if i == 0 {
			for j := i; j < i + 3; j++ {
				val, err := strconv.Atoi(values[j])
				if err != nil {
					log.Fatal(err)
				}
				previous += val
			}
			continue
		}
		popVal, err := strconv.Atoi(values[i-1])
		if err != nil {
			log.Fatal(err)
		}
		pushVal, err := strconv.Atoi(values[i+2])
		if err != nil {
			log.Fatal(err)
		}
		current := previous - popVal + pushVal
		if current > previous {
			increases++
		}
		previous = current
	}
	fmt.Printf("The depth of the ocean floor using a 3 value sliding window increases %d times.\n", increases)
}
