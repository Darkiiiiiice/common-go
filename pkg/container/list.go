package container

import (
	"reflect"
)

type ArrayList struct {
	nodes []arrayListNode

	length   int
	capacity int
}

type arrayListNode struct {
	value interface{}
}

func InitArrayList(capacity int) *ArrayList {
	var list = new(ArrayList)
	list.length = 0
	list.capacity = capacity
	list.nodes = make([]arrayListNode, 0, capacity)
	return list
}

func (l *ArrayList) Length() int {
	return l.length
}

func (l *ArrayList) Clear() {
	l.nodes = make([]arrayListNode, 0)
	l.length = len(l.nodes)
	l.capacity = cap(l.nodes)
}

func (l *ArrayList) IsEmpty() bool {
	return l.length == 0
}

func (l *ArrayList) IsNotEmpty() bool {
	return l.length != 0
}

func (l *ArrayList) Get(index int) interface{} {
	if index < 0 || index >= l.length {
		return nil
	}
	return l.nodes[index].value
}

func (l *ArrayList) IndexOf(value interface{}) int {
	var index = -1
	for i, v := range l.nodes {
		if reflect.DeepEqual(v.value, value) {
			index = i
			break
		}
	}

	return index
}

func (l *ArrayList) Insert(index int, value interface{}) {
	if index < 0 {
		index = 0
	}

	l.nodes = append(l.nodes, arrayListNode{value: value})
	for i := l.length; i > index; i-- {
		l.nodes[i] = l.nodes[i-1]
	}

	if index < l.length {
		l.nodes[index] = arrayListNode{value: value}
	}
	l.length = len(l.nodes)
	l.capacity = cap(l.nodes)
}

func (l *ArrayList) Append(value interface{}) {
	l.nodes = append(l.nodes, arrayListNode{value: value})
	l.length = len(l.nodes)
	l.capacity = cap(l.nodes)
}

func (l *ArrayList) Delete(index int) {
	if index < 0 {
		return
	}

	for i := index; i < l.length-1; i++ {
		l.nodes[i] = l.nodes[i+1]
	}
	if index < l.length {
		l.nodes = l.nodes[:l.length-1]
	}

	l.length = len(l.nodes)
	l.capacity = cap(l.nodes)
}
