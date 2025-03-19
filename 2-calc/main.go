package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const AVG = "AVG"
const SUM = "SUM"
const MED = "MED"

var validOperations = []string{AVG, SUM, MED}

func main() {
	operation, numbers, err := getUserInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	switch operation {
	case AVG:
		fmt.Printf("Average: %.2f\n", average(numbers))
	case SUM:
		fmt.Printf("Sum: %.2f\n", sum(numbers))
	case MED:
		fmt.Printf("Median: %.2f\n", median(numbers))
	}
}

func getUserInput() (string, []float64, error) {
	var operation string
	fmt.Printf("Enter the operation from %v: ", validOperations)
	fmt.Scan(&operation)
	for !isOperationValid(operation) {
		fmt.Printf("Invalid operation. Please enter the operation from %v: ", validOperations)
		fmt.Scan(&operation)
	}
	operation = strings.ToUpper(operation)

	fmt.Print("Enter the numbers separated by a comma: ")
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the numbers separated by a comma: ")
	reader.Scan()
	rawNumbers := reader.Text()
	numbers, err := parseNumbers(rawNumbers)
	if err != nil {
		fmt.Println(err)
		return "", nil, err
	}
	return operation, numbers, nil
}

func isOperationValid(operation string) bool {
	for _, valid := range validOperations {
		if strings.EqualFold(operation, valid) {
			return true
		}
	}
	return false
}

func parseNumbers(rawNumbers string) ([]float64, error) {
	var numbers []float64
	for _, rawNumber := range strings.Split(rawNumbers, ",") {
		number, err := strconv.ParseFloat(strings.TrimSpace(rawNumber), 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %v", rawNumber)
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

func average(numbers []float64) float64 {
	return sum(numbers) / float64(len(numbers))
}

func sum(numbers []float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func median(numbers []float64) float64 {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	length := len(numbers)
	if length%2 == 0 {
		return (numbers[length/2-1] + numbers[length/2]) / 2
	}
	return numbers[length/2]
}
