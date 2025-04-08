package bins

import (
	"fmt"
	"time"
)

type IBin interface {
	PrintBin()
}

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func (bin *Bin) PrintBin() {
	fmt.Println("------------")
	fmt.Println(bin.Id)
	fmt.Println(bin.Private)
	fmt.Println(bin.CreatedAt)
	fmt.Println(bin.Name)
	fmt.Println("------------")
}

func NewBin(id string, private bool, name string) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}
