package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// read input file and print each line
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = input.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(input)

	
	for scanner.Scan() {
		inputValue := scanner.Text()
		var direction string
		var directionValue int
		fmt.Sscanf(inputValue, %s %d &direction, &directionValue)
	}

	var numbersOfIncrease int = 0
	for i, s := range inputArray {
		if i == 0 {
			continue
		} else if s > inputArray[i-1] {
			numbersOfIncrease++
		}
	}
	println("Number of increases:", numbersOfIncrease)
}
