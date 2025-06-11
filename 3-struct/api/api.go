package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"pferdina.com/3-struct/storage"
)

const jsonbinBaseURL = "https://api.jsonbin.io/v3/b"

func CreateBin(storage *storage.Storage, binName string) {
	fmt.Println("Saving bin:", binName)

	bin, err := storage.GetBinByName(binName)
	if err != nil {
		fmt.Println("Bin not found")
		return
	}

	fmt.Println(bin)
}

func CreateBinAPI(apiKey string, binData interface{}) (string, error) {
	data, err := json.Marshal(binData)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", jsonbinBaseURL, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to create bin: %s", string(body))
	}
	var result struct {
		Record struct {
			Id string `json:"id"`
		} `json:"record"`
		Metadata struct {
			Id string `json:"id"`
		} `json:"metadata"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	if result.Metadata.Id != "" {
		return result.Metadata.Id, nil
	}
	return result.Record.Id, nil
}

func UpdateBinAPI(apiKey, binID string, binData interface{}) error {
	data, err := json.Marshal(binData)
	if err != nil {
		return err
	}
	client := &http.Client{}
	url := jsonbinBaseURL + "/" + binID
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to update bin: %s", string(body))
	}
	return nil
}

func DeleteBinAPI(apiKey, binID string) error {
	client := &http.Client{}
	url := jsonbinBaseURL + "/" + binID
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Master-Key", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to delete bin: %s", string(body))
	}
	return nil
}

func GetBinAPI(apiKey, binID string) ([]byte, error) {
	client := &http.Client{}
	url := jsonbinBaseURL + "/" + binID
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Master-Key", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get bin: %s", string(body))
	}
	return io.ReadAll(resp.Body)
}
