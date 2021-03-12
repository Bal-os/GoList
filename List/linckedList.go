package List

import "errors"

type node struct {
	next *node
	prev *node

	value interface{}
}

type LinckedList struct {
	first *node
	last  *node

	size int
}

func NewLinckedList() *LinckedList {
	return new(LinckedList)
}

func (list *LinckedList) GetFirst() (interface{}, error) {
	if list.first != nil {
		return list.first.value, nil
	} else {
		return nil, errors.New("is empty")
	}
}

func (list *LinckedList) GetLast() (interface{}, error) {
	if list.first != nil {
		return list.last.value, nil
	} else {
		return nil, errors.New("is empty")
	}
}

func (list *LinckedList) Add(ell interface{}) {
	var node = new(node)
	node.value = ell
	if list.first != nil {
		node.prev = list.last
		list.last.next = node
		list.last = node
	} else {
		list.first = node
		list.last = node
	}
	list.size++
}

func (list *LinckedList) AddAll(ell ...interface{}) {
	for _, value := range ell {
		list.Add(value)
	}
}

func (list *LinckedList) GetByIndex(index int) (interface{}, error) {
	if index > list.size || index < 0 {
		return nil, errors.New("out of range")
	}
	currNode := list.first
	for i := 0; i < index; i++ {
		currNode = currNode.next
	}
	return currNode.value, nil
}
