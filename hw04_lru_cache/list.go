package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(v interface{}) *listItem
	PushBack(v interface{}) *listItem
	Remove(i *listItem)
	MoveToFront(i *listItem)
}

type listItem struct {
	Value      interface{}
	Prev, Next *listItem
}

type list struct {
	head, tail *listItem
	length     int
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *listItem {
	return l.head
}

func (l *list) Back() *listItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *listItem {
	item := listItem{Value: v}
	if l.length == 0 {
		l.tail = &item
	} else {
		item.Prev = l.head
		l.head.Next = &item
	}
	l.head = &item
	l.length++
	return &item
}

func (l *list) PushBack(v interface{}) *listItem {
	item := listItem{Value: v}
	if l.length == 0 {
		l.head = &item
	} else {
		item.Next = l.tail
		l.tail.Prev = &item
	}
	l.tail = &item
	l.length++
	return &item
}

func (l *list) Remove(i *listItem) {
	switch {
	case i.Next == nil && i.Prev == nil:
		l.head, l.tail = nil, nil
	case i.Next == nil && i.Prev != nil:
		i.Prev.Next = nil
		l.head = i.Prev
	case i.Next != nil && i.Prev == nil:
		i.Next.Prev = nil
		l.tail = i.Next
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	i.Prev, i.Next = nil, nil
	l.length--
}

func (l *list) MoveToFront(i *listItem) {
	if l.Len() == 1 {
		return
	}
	l.Remove(i)
	l.length++
	i.Prev = l.head
	i.Prev.Next = i
	l.head = i
}

func NewList() List {
	return &list{}
}
