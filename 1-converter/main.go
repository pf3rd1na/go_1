package main

import (
	"errors"
	"fmt"
	"strings"
)

const USD = "USD"
const EUR = "EUR"
const RUB = "RUB"

type CrossRates map[string]map[string]float64

func main() {
	crossRates := CrossRates{
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
	currentCurrency, amount, targetCurrency := getUserInput(&crossRates)
	result, err := convert(amount, currentCurrency, targetCurrency, &crossRates)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.2f %s = %.2f %s\n", amount, currentCurrency, result, targetCurrency)
}

func getUserInput(crossRates *CrossRates) (string, float64, string) {
	currentCurrency := getValidCurrency(crossRates)
	amount := getValidAmount()
	targetCurrency := getValidDesiredCurrency(currentCurrency, crossRates)

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

func getValidCurrency(crossRates *CrossRates) string {
	var inputCurrency string
	fmt.Printf("Enter the current currency from %v: ", getValidCurrencies(crossRates))
	fmt.Scan(&inputCurrency)
	for !isCurrencyValid(inputCurrency, &CrossRates{}) {
		fmt.Printf("Invalid currency. Please enter the current currency from %v: ", getValidCurrencies(crossRates))
		fmt.Scan(&inputCurrency)
	}
	inputCurrency = strings.ToUpper(inputCurrency)
	return inputCurrency
}

func getValidDesiredCurrency(currentCurrency string, crossRates *CrossRates) string {
	desiredCurrency := getValidCurrency(crossRates)
	for desiredCurrency == currentCurrency {
		fmt.Println("The desired currency cannot be the same as the current currency.")
		desiredCurrency = getValidCurrency(crossRates)
	}
	return desiredCurrency
}

func isCurrencyValid(currency string, crossRates *CrossRates) bool {
	for _, valid := range getValidCurrencies(crossRates) {
		if strings.EqualFold(currency, valid) {
			return true
		}
	}
	return false
}

func convert(amount float64, initialCurrency, desiredCurrency string, crossRates *CrossRates) (float64, error) {
	value, exists := (*crossRates)[initialCurrency][desiredCurrency]
	if !exists {
		return 0, errors.New("invalid currency pair")
	}
	return value * amount, nil
}

func getValidCurrencies(crossRates *CrossRates) []string {
	var validCurrencies []string
	for currency := range *crossRates {
		validCurrencies = append(validCurrencies, currency)
	}
	return validCurrencies
}
