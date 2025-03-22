package files

import (
	"errors"
	"os"
	"strings"
)

const jsonExtension = ".json"

func WriteFile(data []byte, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(data)
	return nil
}

func ReadFile(name string) ([]byte, error) {
	if !checkFileExtenstion(name, jsonExtension) {
		return nil, errors.New("file must be a json file")
	}
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func checkFileExtenstion(name string, ext string) bool {
	return strings.HasSuffix(name, ext)
}
