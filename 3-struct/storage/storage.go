package storage

import (
	"encoding/json"
	"fmt"
	"go-demo1/3-struct/bins"
)

type Store interface {
	Read() ([]byte, error)
	Write([]byte) error
}

type Storage struct {
	BinList []bins.Bin
	store Store
}

func NewStorage(store Store) *Storage {
	file, err := store.Read()
	if err != nil {
		return &Storage{
			BinList: []bins.Bin{},
			store: store,
		}
	}
	var binList []bins.Bin
	err = json.Unmarshal(file, &binList)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return &Storage{
			BinList: []bins.Bin{},
		store: store,
		}
	}
	return &Storage{
		BinList: binList,
		store: store,
	}
}

func (storage *Storage) saveBins() error {
	data, err := json.Marshal(storage.BinList)
	if err != nil {
		return err
	}
	err = storage.store.Write(data)
	return err
}

func (storage *Storage) AddBin(bin bins.Bin) error {
	storage.BinList = append(storage.BinList, bin)
	return storage.saveBins()
}

func (storage *Storage) ShowBinList() {
	for _, bin := range storage.BinList {
		fmt.Printf("ID: %s, Name: %s, Private: %t, CreatedAt: %s\n", bin.Id, bin.Name, bin.Private, bin.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}