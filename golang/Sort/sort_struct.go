package main

import (
	"fmt"
	"sort"
)

type Item struct {
	name string
	cost int
}

type ItemList struct {
	sort.Interface
	items []Item
}

func New() ItemList {
	return ItemList{
		items: []Item{},
	}
}

func (t ItemList) Len() int {
	return len(t.items)
}

func (t ItemList) Less(i, j int) bool {
	return t.items[i].cost < t.items[j].cost
}

func (t ItemList) Swap(i, j int) {
	temp := t.items[i]
	t.items[i] = t.items[j]
	t.items[j] = temp
}

func main() {
	fmt.Println("vim-go")
}
