package storage

import (
	"encoding/json"
	"go-demo1/3-struct/file"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

var BinList = []Bin{}

func SaveBins(binList []Bin) error {
	data, err := json.Marshal(binList)
	if err != nil {
		return err
	}
	const outPath = "bins.json"
	err = file.WriteFile(data, outPath)
	return err
}

func LoadBins(name string) error {
	data, err := file.ReadFile(name)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &BinList)
	return err
}
