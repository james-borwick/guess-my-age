package main

import "fmt"

var ageGuess float64

func main() {

	for ageGuess != 31 {

		fmt.Print("Guess my age: ")
		fmt.Scanf("%f", &ageGuess)

		switch {
		case ageGuess > 28 && ageGuess < 34:
			fmt.Println("Very close. Try again...")
		case ageGuess < 13:
			fmt.Println("Way too young. Try again...")
		case ageGuess < 20:
			fmt.Println("I'm older than a teenager. Try again...")
		case ageGuess < 30:
			fmt.Println("I'm not in my twenties anymore. Try again...")
		case ageGuess < 40:
			fmt.Println("Getting close. Try again...")
		case ageGuess < 50:
			fmt.Println("I'm not that old yet. Try again...")
		case ageGuess > 49:
			fmt.Println("That's way too old. Try again...")
		}
	}
	fmt.Println("Correct! Well done.")
}
