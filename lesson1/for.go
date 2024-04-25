package main

import "fmt"

func main() {
	for {
		fmt.Print("Введіть щось: ")
		var input string
		_, _ = fmt.Scanln(&input)
		fmt.Printf("Ви ввели: %s\n", input)
	}
}
