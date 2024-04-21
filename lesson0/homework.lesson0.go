package main

import (
	"fmt"
)

const year = uint16(2025)

var yearOfBirth uint16
var currentYear uint16
var currentAge uint16

func main() {
	greetUser()
	getUserData()
	ageCalculation(yearOfBirth, currentYear)

	// Вивід результату: Програма має вивести вік користувача з повідомленням, наприклад: "Your age: 25 years old".
	fmt.Printf("Your age is %v. in main function \n", currentAge)
	fmt.Printf("Thank you. Good bye! \n")
}

func greetUser() {
	fmt.Println("################")
	fmt.Printf("Hello! weclome to age calculator\n")
}

// 1. Ввід даних: Програма має запитати у користувача його рік народження та поточний рік.
func getUserData() (uint16, uint16) {
	fmt.Printf("Enter year of your birth\n")
	fmt.Scan(&yearOfBirth)
	if yearOfBirth > year {
		fmt.Printf("Are you form future? \n")
		fmt.Printf("Birth year should be less than %v. \n", year)
		fmt.Scan(&yearOfBirth)
	}

	fmt.Printf("Enter current year\n")
	fmt.Scan(&currentYear)

	// Обробка помилок: Програма має перевіряти, що введені дані є числами і що рік народження не більший за поточний рік.
	if currentYear < yearOfBirth {
		fmt.Printf("Current year greater than your dat of birth \n")
		fmt.Printf("Current year should be greater than %v. \n", yearOfBirth)
		fmt.Scan(&currentYear)
	}
	return yearOfBirth, currentYear
}

// 2.  Обрахунок віку: Програма має використати отримані дані для розрахунку віку.
func ageCalculation(yearOfBirth uint16, currentYear uint16) uint16 {
	fmt.Println("################")
	fmt.Printf("Year of your birth: %v.\n", yearOfBirth)
	fmt.Printf("Current year: %v.\n", currentYear)
	currentAge = currentYear - yearOfBirth

	// Вивід результату: Програма має вивести вік користувача з повідомленням, наприклад: "Your age: 25 years old".
	fmt.Printf("Your age: %v years old.\n", currentAge)
	fmt.Println("################")
	return currentAge

}
