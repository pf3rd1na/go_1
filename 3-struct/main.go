package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func main() {
	binList := make([]Bin, 0)

	getUserInput(&binList)

	for _, bin := range binList {
		fmt.Println("------------")
		fmt.Println(bin.id)
		fmt.Println(bin.private)
		fmt.Println(bin.createdAt)
		fmt.Println(bin.name)
		fmt.Println("------------")
	}
}

func getUserInput(binList *[]Bin) {
	bin := Bin{}

	r := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter id: ")
	r.Scan()
	bin.id = r.Text()

	bin.private = getPrivateField(r)

	fmt.Println("Enter name: ")
	r.Scan()
	bin.name = r.Text()

	bin.createdAt = time.Now()

	*binList = append(*binList, bin)
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
