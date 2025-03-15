package main

import "fmt"

func main() {
	const usdToEur = 0.92
	const usdToRub = 85.19
	eurToRub := usdToRub / usdToEur
	amount := getUserInput()
	_ = convert(0, "", "")
	fmt.Println("Eur to Rub :", amount*eurToRub)
}

func getUserInput() float64 {
	var userInput float64
	fmt.Print("Enter the amount of EUR: ")
	fmt.Scan(&userInput)
	return userInput
}

func convert(amount float64, initialCurrency, desiredCurrency string) float64 {
	return 0
}
