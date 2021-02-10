package algs4

import (
	"github.com/huggin/algo/linklist"
)

// Bag implements a bag data type by linklist
type Bag struct {
	*linklist.LinkList
}

// NewBag ...
func NewBag() *Bag {
	return &Bag{linklist.NewLinkList()}
}
