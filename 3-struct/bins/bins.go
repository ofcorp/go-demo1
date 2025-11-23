package bins

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

var BinList = []Bin{}

func CreateBin(id string, private bool, name string) Bin {
	bin := Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
	BinList = append(BinList, bin)
	return bin
}