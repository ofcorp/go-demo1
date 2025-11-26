package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(name string) ([]byte, error) {
	if filepath.Ext(name) != ".json" {
		return nil, errors.New("это не .json файл")
	}

	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	if !json.Valid(data) {
		return nil, errors.New("файл не содержит корректный JSON")
	}

	return data, nil
}

func WriteFile(content []byte, name string) error{
	file, err := os.Create(name)
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
