package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"pferdina.com/3-struct/bins"
)

func main() {
	binList := make([]bins.Bin, 0)

	getUserInput(&binList)

	for _, bin := range binList {
		bin.PrintBin()
	}
}

func getUserInput(binList *[]bins.Bin) {
	r := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter id: ")
	r.Scan()
	id := r.Text()

	private := getPrivateField(r)

	fmt.Println("Enter name: ")
	r.Scan()
	name := r.Text()

	bin := bins.NewBin(id, private, name)
	*binList = append(*binList, *bin)
}

func getPrivateField(r *bufio.Scanner) bool {
	fmt.Println("Enter private: ")
	r.Scan()
	private, err := strconv.ParseBool(r.Text())
	for err != nil {
		fmt.Println("Enter valid bool (true/false): ")
		r.Scan()
		private, err = strconv.ParseBool(r.Text())
	}
	return private
}
