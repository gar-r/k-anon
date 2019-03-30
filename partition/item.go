package partition

import "fmt"

type Item struct {
	item interface{}
}

func NewItem(item interface{}) *Item {
	return &Item{item: item}
}

func (p *Item) Contains(item interface{}) bool {
	return p.item == item
}

func (p *Item) ContainsPartition(other Partition) bool {
	return false
}

func (p *Item) Equals(other Partition) bool {
	q, success := other.(*Item)
	if !success {
		return false
	}
	return p.item == q.item
}

func (p *Item) String() string {
	return fmt.Sprintf("%v", p.item)
}

func (p *Item) GetItem() interface{} {
	return p.item
}
