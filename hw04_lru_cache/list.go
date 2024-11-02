package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	length int
	head   *ListItem
	tail   *ListItem
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
	}

	newItem.Next = l.head
	l.head = newItem

	if l.tail == nil {
		l.tail = l.head
	}

	l.length++

	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
	}

	l.tail.Next = newItem
	newItem.Prev = l.tail
	l.tail = newItem

	if l.head == nil {
		l.head = l.tail
	}

	l.length++

	return newItem
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	if l.head == i {
		l.head = i.Next
	}

	if l.tail == i {
		l.tail = i.Prev
	}

	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == nil || i == l.head {
		return
	}

	l.Remove(i)

	i.Prev = nil
	i.Next = l.head
	l.head = i

	l.length++
}

func NewList() List {
	return new(list)
}
