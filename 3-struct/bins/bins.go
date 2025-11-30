package bins

import (
	"fmt"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

func createBin(id string, private bool, name string) Bin {
	bin := Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
	return bin
}

func AddNewBin() Bin {
	var id string
	var name string
	var private bool
	fmt.Print("Введите ID Bin: ")
	fmt.Scan(&id)
	fmt.Print("Введите имя Bin: ")
	fmt.Scan(&name)
	fmt.Print("Bin приватный? (true/false): ")
	fmt.Scan(&private)
	bin := createBin(id, private, name)
	return bin
}