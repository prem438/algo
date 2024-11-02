package main

import (
	"sort"
	"testing"
)

func TestItemList(t *testing.T) {
	itemList := New()

	for i := 'a' ; i <= 'z' ; i++{
		item := Item{
			name: string(i),
			cost: int(i),
		}
		itemList.items = append(itemList.items, item)
	}
	t.Log(itemList)
	//sort.Sort(itemList)

	t.Log(itemList)
	result := sort.Reverse(itemList)
	sort.Sort(result)
	t.Log(result)
}

func Append(q []int) []int{
	q = append(q, 10)
	return q
}

func TestAppend(t *testing.T){
	p := []int{1}
	q := p
	p = Append(p)
	t.Log(p)
	t.Log(q)
}

func MapAppend(q *map[int]bool){
	(*q)[10] = true
}
func TestMapAppend(t *testing.T){
	p := map[int]bool{
		1: true,
	} 
	MapAppend(&p)
	t.Log(p)
}