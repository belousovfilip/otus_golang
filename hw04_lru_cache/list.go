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
	len   int
	front *ListItem
	back  *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	if l.front == nil {
		return l.back
	}
	return l.front
}

func (l *list) Back() *ListItem {
	if l.back == nil {
		return l.front
	}
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	var item *ListItem
	if l.Len() == 0 {
		item = NewListItem(v, nil, nil)
		l.front = item
	}
	if l.Len() == 1 {
		l.back = l.Front()
		item = NewListItem(v, nil, l.back)
		item.Next = l.back
		l.back.Prev = item
	}
	if l.Len() > 1 {
		front := l.Front()
		item = NewListItem(v, nil, front)
		item.Next = front
		front.Prev = item
	}
	l.front = item
	l.len++
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	var item *ListItem
	if l.Len() == 0 {
		item = NewListItem(v, nil, nil)
		l.back = item
	}
	if l.Len() == 1 {
		l.front = l.Back()
		item = NewListItem(v, l.front, nil)
		l.front.Next = item
		l.back = item
	}
	if l.Len() > 1 {
		item = NewListItem(v, l.back, nil)
		l.back.Next = item
		l.back = item
	}
	l.len++
	return item
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if l.front == i {
		l.front = i.Next
	}
	if l.back == i {
		l.back = i.Prev
	}
	l.len--
	i.Next = nil
	i.Prev = nil
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return &list{}
}

func NewListItem(value interface{}, prev, next *ListItem) *ListItem {
	return &ListItem{value, next, prev}
}
