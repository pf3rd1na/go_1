package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"

	"pferdina.com/3-struct/api"
	"pferdina.com/3-struct/bins"
	"pferdina.com/3-struct/config"
	"pferdina.com/3-struct/storage"
)

func main() {
	api.GetEnv()
	cfg := config.NewConfig()
	storage := storage.NewStorage()

	// Command-line flags
	createFlag := flag.Bool("create", false, "Create a new bin")
	updateFlag := flag.Bool("update", false, "Update a bin")
	deleteFlag := flag.Bool("delete", false, "Delete a bin")
	getFlag := flag.Bool("get", false, "Get a bin")
	listFlag := flag.Bool("list", false, "List bins")
	fileFlag := flag.String("file", "", "File to read/write bin data")
	nameFlag := flag.String("name", "", "Name of the bin")
	idFlag := flag.String("id", "", "ID of the bin")
	flag.Parse()

	switch {
	case *createFlag:
		if *fileFlag == "" || *nameFlag == "" {
			fmt.Println("--file and --name are required for --create")
			return
		}
		data, err := os.ReadFile(*fileFlag)
		if err != nil {
			fmt.Println("Failed to read file:", err)
			return
		}
		bin := &bins.Bin{}
		err = json.Unmarshal(data, bin)
		if err != nil {
			fmt.Println("Failed to parse JSON:", err)
			return
		}
		bin.Name = *nameFlag
		// Create in API
		apiID, err := api.CreateBinAPI(cfg.Key, bin)
		if err != nil {
			fmt.Println("API create failed:", err)
			return
		}
		bin.Id = apiID
		err = storage.AddBin(*bin)
		if err != nil {
			fmt.Println("Local storage add failed:", err)
			return
		}
		fmt.Println("Bin created with ID:", apiID)
	case *updateFlag:
		if *fileFlag == "" || *idFlag == "" {
			fmt.Println("--file and --id are required for --update")
			return
		}
		data, err := os.ReadFile(*fileFlag)
		if err != nil {
			fmt.Println("Failed to read file:", err)
			return
		}
		bin := &bins.Bin{}
		err = json.Unmarshal(data, bin)
		if err != nil {
			fmt.Println("Failed to parse JSON:", err)
			return
		}
		bin.Id = *idFlag
		err = api.UpdateBinAPI(cfg.Key, *idFlag, bin)
		if err != nil {
			fmt.Println("API update failed:", err)
			return
		}
		err = storage.UpdateBin(*idFlag, *bin)
		if err != nil {
			fmt.Println("Local storage update failed:", err)
			return
		}
		fmt.Println("Bin updated with ID:", *idFlag)
	case *deleteFlag:
		if *idFlag == "" {
			fmt.Println("--id is required for --delete")
			return
		}
		err := api.DeleteBinAPI(cfg.Key, *idFlag)
		if err != nil {
			fmt.Println("API delete failed:", err)
			return
		}
		err = storage.DeleteBin(*idFlag)
		if err != nil {
			fmt.Println("Local storage delete failed:", err)
			return
		}
		fmt.Println("Bin deleted with ID:", *idFlag)
	case *getFlag:
		if *idFlag == "" {
			fmt.Println("--id is required for --get")
			return
		}
		// Try local first
		var found bool
		for _, bin := range storage.Bins {
			if bin.Id == *idFlag {
				fmt.Println("Local bin:")
				bin.PrintBin()
				found = true
				break
			}
		}
		// Always try API as well
		apiData, err := api.GetBinAPI(cfg.Key, *idFlag)
		if err != nil {
			fmt.Println("API get failed:", err)
		} else {
			fmt.Println("API bin:")
			fmt.Println(string(apiData))
		}
		if !found {
			fmt.Println("Bin not found locally.")
		}
	case *listFlag:
		bins := storage.ListBins()
		fmt.Println("Local bins:")
		for _, bin := range bins {
			fmt.Printf("ID: %s, Name: %s\n", bin.Id, bin.Name)
		}
	default:
		fmt.Println("No valid command provided. Use --create, --update, --delete, --get, or --list.")
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
	storage.AddBin(*bin)
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
