package array

import "fmt"

type Array struct {
	size  int
	count int
	items []int
}

func (a *Array) init(size int) {
	a.size = size
	a.items = make([]int, a.size)
	a.count = 0
}

func (a *Array) getIndex(value int) int {
	for i, v := range a.items {
		if v == value {
			return i
		}
	}
	return -1
}

func (a *Array) expand() {
	oldArray := a.items

	a.size = a.size * 2
	a.items = make([]int, a.size)

	for i, item := range oldArray {
		a.items[i] = item
	}
}

func (a *Array) insert(item int) {
	if a.count == a.size {
		a.expand()
	}

	a.items[a.count] = item
	a.count++
}

func (a *Array) removeAt(index int) {
	if index < 0 || index >= a.count {
		return
	}

	for i := index; i+1 < a.size; i++ {
		a.items[i] = a.items[i+1]
	}

	a.count--
}

func (a *Array) print() {
	for _, value := range a.items {
		fmt.Println(value)
	}
}

func (a *Array) swap(i, j int) {
	if i < 0 || j < 0 {
		return
	}

	if i >= a.count || j >= a.count {
		return
	}

	item1 := &a.items[i]
	item2 := &a.items[j]

	temp := *item1
	*item1 = *item2
	*item2 = temp
}

func (a *Array) invert() {
	for i := 0; i < (a.count / 2); i++ {
		a.swap(i, (a.count - 1 - i))
	}
}
