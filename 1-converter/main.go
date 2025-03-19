package main

import (
	"errors"
	"fmt"
	"strings"
)

const USD = "USD"
const EUR = "EUR"
const RUB = "RUB"

var validCurrencies = []string{USD, EUR, RUB}

func main() {
	currentCurrency, amount, targetCurrency := getUserInput()
	result, err := convert(amount, currentCurrency, targetCurrency)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.2f %s = %.2f %s\n", amount, currentCurrency, result, targetCurrency)
}

func getUserInput() (string, float64, string) {
	currentCurrency := getValidCurrency()
	amount := getValidAmount()
	targetCurrency := getValidDesiredCurrency(currentCurrency)

	return currentCurrency, amount, targetCurrency
}

func getValidAmount() float64 {
	var amount float64
	fmt.Print("Enter the amount: ")
	fmt.Scan(&amount)
	for amount <= 0 {
		fmt.Print("Invalid amount. Please enter the amount: ")
		fmt.Scan(&amount)
	}
	return amount
}

func getValidCurrency() string {
	var inputCurrency string
	fmt.Printf("Enter the current currency from %v: ", validCurrencies)
	fmt.Scan(&inputCurrency)
	for !isCurrencyValid(inputCurrency) {
		fmt.Printf("Invalid currency. Please enter the current currency from %v: ", validCurrencies)
		fmt.Scan(&inputCurrency)
	}
	inputCurrency = strings.ToUpper(inputCurrency)
	return inputCurrency
}

func getValidDesiredCurrency(currentCurrency string) string {
	desiredCurrency := getValidCurrency()
	for desiredCurrency == currentCurrency {
		fmt.Println("The desired currency cannot be the same as the current currency.")
		desiredCurrency = getValidCurrency()
	}
	return desiredCurrency
}

func isCurrencyValid(currency string) bool {
	for _, valid := range validCurrencies {
		if strings.EqualFold(currency, valid) {
			return true
		}
	}
	return false
}

func convert(amount float64, initialCurrency, desiredCurrency string) (float64, error) {
	const usdToEur = 0.92
	const usdToRub = 85.19
	eurToRub := usdToRub / usdToEur
	switch initialCurrency {
	case USD:
		switch desiredCurrency {
		case EUR:
			return amount * usdToEur, nil
		case RUB:
			return amount * usdToRub, nil
		}
	case EUR:
		switch desiredCurrency {
		case USD:
			return amount / usdToEur, nil
		case RUB:
			return amount * eurToRub, nil
		}
	case RUB:
		switch desiredCurrency {
		case USD:
			return amount / usdToRub, nil
		case EUR:
			return amount / eurToRub, nil
		}
	}
	return -1, errors.New("something went wrong")
}
