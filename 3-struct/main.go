package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"pferdina.com/3-struct/bins"
	"pferdina.com/3-struct/storage"
)

func main() {
	storage := storage.NewStorage()

	for {
		action := getMenu()
		switch action {
		case 1:
			getUserInput(storage)
		case 5:
			os.Exit(0)
		}
	}
}

func getMenu() int {
	fmt.Println("1. Add bin")
	fmt.Println("5. Exit")
	var action int
	fmt.Scan(&action)
	return action
}

func getUserInput(storage *storage.Storage) {
	id := promtInput("Enter id: ")
	private := getPrivateField()
	name := promtInput("Enter name: ")

	bin := bins.NewBin(id, private, name)
	fmt.Println("Bin created", bin)
	storage.AddBin(bin)
}

func promtInput(message string) string {
	fmt.Print(message)
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	return r.Text()
}

func getPrivateField() bool {
	input := promtInput("Enter private (true/false): ")
	private, err := strconv.ParseBool(input)
	for err != nil {
		input = promtInput("Enter valid bool (true/false): ")
		private, err = strconv.ParseBool(input)
	}
	return private
}
