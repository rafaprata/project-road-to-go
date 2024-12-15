package main

import (
	"fmt"
	"math"
)

func readInput() string {
	var input string
	fmt.Println("Enter an up to 8 digit binary number:")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	verifyInput(input)
	return input
}

func verifyInput(input string) bool {
	valid := true
	switch {
	case inputLen(input):
		fmt.Println("Error: The binary number is greater than 8 digits")
		valid = false
	case inputChar(input):
		fmt.Println("Error: Input is not a binary number")
		valid = false
	}
	return valid
}

func inputLen(input string) bool {
	return len(input) > 8
}

func inputChar(input string) bool {
	var charNotValid bool
	for _, character := range input {
		switch {
		case character == '0' || character == '1':
			charNotValid = false
		default:
			charNotValid = true
			return charNotValid
		}
	}
	return charNotValid
}

func completeInput(input string) string {
	for len(input) < 8 {
		input = "0" + input
	}
	return input
}

func convertToDec(binaryNumber string) uint8 {
	var decNumber, intChar uint8
	for k, character := range binaryNumber {
		index := 7 - k
		switch {
		case character == '1':
			intChar = 1
		default:
			intChar = 0
		}
		decNumber += intChar * uint8(math.Pow(2, float64(index)))
	}
	return decNumber
}

func main() {
	binaryInput := readInput()
	binaryInput = completeInput(binaryInput)
	decNumber := convertToDec(binaryInput)
	fmt.Println(decNumber)
}
