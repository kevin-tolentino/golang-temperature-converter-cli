package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var originUnit string
var originValue float64

var shouldConvertAgain string

var err error

var errInvalidArguments = errors.New("Invalid arguments!")
var errReadingInput = errors.New("Error reading input")

func main() {
	if len(os.Args) != 2 { //check if the arguments are not 2
		printError(errInvalidArguments) // invoke to prompt that arguments are invalid
	}

	originUnit = strings.ToUpper(os.Args[1]) // call method for consistency sake of user input and assign to originUnit
	//fmt.Println(" line 26!!Origin Unit:", originUnit)

	for {
		if originUnit != "C" || originUnit != "F" { //why doesn't this work?
			//if !(originUnit == "C" || originUnit == "F") { // my own implementation to check that argument is C or F for unit
			printError(errInvalidArguments) // invoke error message if not so
		}
//Start with the positive case first (think of it like a unit test)
		fmt.Print("What is the current temperature in " + originUnit + " ? ")
		_, err = fmt.Scanln(&originValue) // read the input and
		//& tells the runner to pass by ref rather than value
		//fmt.Println(" line 35!!Origin Value:", originValue)

		if err != nil { //if not a float64 type
			printError(errReadingInput)
		}

		if originUnit == "C" { // if the originUnit was originally C unit
			convertToFahrenheit(originValue) //convert the value to fahrenheit
		} else {
			convertToCelsius(originValue) // if not C unit, convert value to celsius
		}

		fmt.Print("Would you like to convert another temperature ? (y/n) ")

		_, err = fmt.Scanln(&shouldConvertAgain) // read line

		if err != nil { //if any error, print error
			printError(errReadingInput) //
		}

		if strings.ToUpper(strings.TrimSpace(shouldConvertAgain)) != "Y" { // check if string y is used
			fmt.Println("Good bye!")
			break
		}
	}
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func convertToCelsius(value float64) {
	convertedValue := (value - 32) * 5 / 9
	fmt.Printf("%v F = %.0f C\n", value, convertedValue)
}

func convertToFahrenheit(value float64) {
	convertedValue := (value * 9 / 5) + 32
	fmt.Printf("%v C = %.0f F\n", value, convertedValue)
}
