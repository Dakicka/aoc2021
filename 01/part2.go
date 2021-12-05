package main

import (
	"bufio"
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

	var inputArray []int
	for scanner.Scan() {
		inputValue, _ := strconv.Atoi(scanner.Text())
		inputArray = append(inputArray, inputValue)
	}

	var numbersOfIncrease int = 0
	for i, s := range inputArray {
		if i <= 2 {
			continue
		}
		a := inputArray[i-3] + inputArray[i-2] + inputArray[i-1]
		b := inputArray[i-2] + inputArray[i-1] + s
		if a < b {
			numbersOfIncrease++
		}
	}
	println("Number of increases:", numbersOfIncrease)
}
