package main

import "fmt"

func main() {
	const usdToEur = 0.92
	const usdToRub = 85.19
	eurToRub := usdToRub / usdToEur
	fmt.Println("Eur to Rub :", eurToRub)
}
