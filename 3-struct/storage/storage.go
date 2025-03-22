package storage

import (
	"encoding/json"
	"time"

	"pferdina.com/3-struct/bins"
	"pferdina.com/3-struct/files"
)

const storageFileName = "storage.json"

type Storage struct {
	Bins      []bins.Bin `json:"bins"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

func (s *Storage) AddBin(bin *bins.Bin) error {
	s.Bins = append(s.Bins, *bin)
	s.UpdatedAt = time.Now()

	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	err = files.WriteFile(data, storageFileName)
	if err != nil {
		return err
	}
	return nil
}

func NewStorage() *Storage {
	data, err := files.ReadFile(storageFileName)
	if err != nil {
		return newEmptyStorage()
	}
	var storage Storage
	err = json.Unmarshal(data, &storage)
	if err != nil {
		return newEmptyStorage()
	}
	return &storage
}

func newEmptyStorage() *Storage {
	return &Storage{
		Bins:      make([]bins.Bin, 0),
		UpdatedAt: time.Now(),
	}
}
