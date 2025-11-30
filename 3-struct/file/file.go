package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type FileStore struct{
	filename string
}

func NewFileStore(filename string) *FileStore {
	return &FileStore{filename: filename}	
}

func (store *FileStore) Read() ([]byte, error) {
	if filepath.Ext(store.filename) != ".json" {
		return nil, errors.New("это не .json файл")
	}

	data, err := os.ReadFile(store.filename)
	if err != nil {
		return nil, err
	}

	if !json.Valid(data) {
		return nil, errors.New("файл не содержит корректный JSON")
	}

	return data, nil
}

func (store *FileStore) Write(content []byte) error {
	file, err := os.Create(store.filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		return err
	}
	fmt.Println("Запись успешна")
	return nil
}
