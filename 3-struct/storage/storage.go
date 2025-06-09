package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"pferdina.com/3-struct/bins"
	"pferdina.com/3-struct/files"
)

const storageFileName = "storage.json"

type Storage struct {
	files     files.IFiles
	Bins      []bins.Bin `json:"bins"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

func (s *Storage) AddBin(bin bins.Bin) error {
	s.Bins = append(s.Bins, bin)
	s.UpdatedAt = time.Now()

	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	err = s.files.WriteFile(data, storageFileName)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetBinByName(name string) (*bins.Bin, error) {
	for _, bin := range s.Bins {
		if bin.Name == name {
			return &bin, nil
		}
	}
	return nil, errors.New("bin not found with name: " + name)
}

func (s *Storage) UpdateBin(id string, updatedBin bins.Bin) error {
	for i, bin := range s.Bins {
		if bin.Id == id {
			s.Bins[i] = updatedBin
			s.UpdatedAt = time.Now()
			data, err := json.Marshal(s)
			if err != nil {
				return err
			}
			return s.files.WriteFile(data, storageFileName)
		}
	}
	return errors.New("bin not found with id: " + id)
}

func (s *Storage) DeleteBin(id string) error {
	for i, bin := range s.Bins {
		if bin.Id == id {
			s.Bins = append(s.Bins[:i], s.Bins[i+1:]...)
			s.UpdatedAt = time.Now()
			data, err := json.Marshal(s)
			if err != nil {
				return err
			}
			return s.files.WriteFile(data, storageFileName)
		}
	}
	return errors.New("bin not found with id: " + id)
}

func (s *Storage) ListBins() []bins.Bin {
	return s.Bins
}

func NewStorage() *Storage {
	files := files.NewFiles()
	data, err := files.ReadFile(storageFileName)
	if err != nil {
		fmt.Println("Error reading storage file:", err)
		return newEmptyStorage()
	}
	var storage Storage
	err = json.Unmarshal(data, &storage)
	if err != nil {
		fmt.Println("Error unmarshalling storage data:", err)
		return newEmptyStorage()
	}
	storage.files = files
	return &storage
}

func newEmptyStorage() *Storage {
	return &Storage{
		Bins:      make([]bins.Bin, 0),
		UpdatedAt: time.Now(),
		files:     files.NewFiles(),
	}
}
