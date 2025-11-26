package storage

import (
	"encoding/json"
	"go-demo1/3-struct/file"
	"go-demo1/3-struct/bins"
)

func SaveBins(binList []bins.Bin) error {
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
	err = json.Unmarshal(data, &bins.BinList)
	return err
}
