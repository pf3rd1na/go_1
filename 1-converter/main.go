package main

import (
	"errors"
	"fmt"
	"strings"
)

const USD = "USD"
const EUR = "EUR"
const RUB = "RUB"

var crossRates = map[string]map[string]float64{
	USD: {
		EUR: 0.92,
		RUB: 85.19,
	},
	EUR: {
		USD: 1.09,
		RUB: 92.61,
	},
	RUB: {
		USD: 0.012,
		EUR: 0.011,
	},
}

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
	fmt.Printf("Enter the current currency from %v: ", getValidCurrencies())
	fmt.Scan(&currentCurrency)
	for !isCurrencyValid(currentCurrency) {
		fmt.Printf("Invalid currency. Please enter the current currency from %v: ", getValidCurrencies())
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
	fmt.Printf("Enter the desired currency from %v: ", getValidCurrencies())
	fmt.Scan(&targetCurrency)
	for !isCurrencyValid(targetCurrency) || targetCurrency == currentCurrency {
		fmt.Printf("Invalid currency. Please enter the desired currency from %v: ", getValidCurrencies())
		fmt.Scan(&targetCurrency)
	}
	targetCurrency = strings.ToUpper(targetCurrency)

	return currentCurrency, amount, targetCurrency
}

func isCurrencyValid(currency string) bool {
	for _, valid := range getValidCurrencies() {
		if strings.EqualFold(currency, valid) {
			return true
		}
	}
	return false
}

func convert(amount float64, initialCurrency, desiredCurrency string) (float64, error) {
	value, exists := crossRates[initialCurrency][desiredCurrency]
	if !exists {
		return 0, errors.New("invalid currency pair")
	}
	return value * amount, nil
}

func getValidCurrencies() []string {
	var validCurrencies []string
	for currency := range crossRates {
		validCurrencies = append(validCurrencies, currency)
	}
	return validCurrencies
}
