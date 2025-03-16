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
	var currentCurrency string
	fmt.Printf("Enter the current currency from %v: ", validCurrencies)
	fmt.Scan(&currentCurrency)
	for !isCurrencyValid(currentCurrency) {
		fmt.Printf("Invalid currency. Please enter the current currency from %v: ", validCurrencies)
		fmt.Scan(&currentCurrency)
	}
	currentCurrency = strings.ToUpper(currentCurrency)

	var amount float64
	fmt.Print("Enter the amount: ")
	fmt.Scan(&amount)
	for amount <= 0 {
		fmt.Print("Invalid amount. Please enter the amount: ")
		fmt.Scan(&amount)
	}

	var targetCurrency string
	fmt.Printf("Enter the desired currency from %v: ", validCurrencies)
	fmt.Scan(&targetCurrency)
	for !isCurrencyValid(targetCurrency) || targetCurrency == currentCurrency {
		fmt.Printf("Invalid currency. Please enter the desired currency from %v: ", validCurrencies)
		fmt.Scan(&targetCurrency)
	}
	targetCurrency = strings.ToUpper(targetCurrency)

	return currentCurrency, amount, targetCurrency
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
