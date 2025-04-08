package storage

import (
	"encoding/json"
	"time"

	"pferdina.com/3-struct/bins"
	"pferdina.com/3-struct/files"
)

const storageFileName = "storage.json"

type Storage struct {
	files     files.IFiles
	Bins      []bins.IBin `json:"bins"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

func (s *Storage) AddBin(bin bins.IBin) error {
	s.Bins = append(s.Bins, bin)
	s.UpdatedAt = time.Now()

	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	print(storageFileName, data)
	err = s.files.WriteFile(data, storageFileName)
	if err != nil {
		return err
	}
	return nil
}

func NewStorage() *Storage {
	files := files.NewFiles()
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
		Bins:      make([]bins.IBin, 0),
		UpdatedAt: time.Now(),
		files:     files.NewFiles(),
	}
}
