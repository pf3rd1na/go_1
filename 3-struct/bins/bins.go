package bins

import (
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func (bin *Bin) PrintBin() {
	fmt.Println("------------")
	fmt.Println(bin.id)
	fmt.Println(bin.private)
	fmt.Println(bin.createdAt)
	fmt.Println(bin.name)
	fmt.Println("------------")
}

func NewBin(id string, private bool, name string) *Bin {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}
