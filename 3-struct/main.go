package main

import (
	"go-demo1/3-struct/bins"
	"go-demo1/3-struct/api"
	"go-demo1/3-struct/file"
	"go-demo1/3-struct/storage"
)

func main() {
	bins.CreateBin("123456", true, "Personal Bin")
	storage.SaveBins(bins.BinList)
	storage.LoadBins("bins.json")
}


