package bins

import "time"

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

var BinList = []Bin{}

func CreateBin(id string, private bool, name string) Bin {
	bin := Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
	BinList = append(BinList, bin)
	return bin
}