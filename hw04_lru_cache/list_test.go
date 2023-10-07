package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})
	t.Run("MoveToFront", func(t *testing.T) {
		l := NewList()
		l.PushFront(10)
		l.PushFront(20)
		l.MoveToFront(l.Back())
		require.Equal(t, []int{10, 20}, convertListToSlice(l))
		require.Equal(t, []int{20, 10}, convertListToReverseSlice(l))
	})
	t.Run("Front", func(t *testing.T) {
		l := NewList()

		l.PushFront(30)
		require.Equal(t, 30, l.Front().Value)
		require.Nil(t, l.Front().Next)
		require.Nil(t, l.Front().Prev)

		l.PushFront(20)
		require.Equal(t, 20, l.Front().Value)
		require.Nil(t, l.Front().Prev)
		require.Nil(t, l.Front().Next.Next)

		l.PushFront(10)
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 20, l.Front().Next.Value)
		require.Equal(t, 30, l.Front().Next.Next.Value)
		require.Nil(t, l.Front().Next.Next.Next)

		require.Equal(t, []int{10, 20, 30}, convertListToSlice(l))
		require.Equal(t, []int{30, 20, 10}, convertListToReverseSlice(l))
	})
	t.Run("Back", func(t *testing.T) {
		l := NewList()

		l.PushBack(10)
		require.Equal(t, 10, l.Back().Value)
		require.Nil(t, l.Back().Next)
		require.Nil(t, l.Back().Prev)

		l.PushBack(20)
		require.Equal(t, 20, l.Back().Value)
		require.Nil(t, l.Back().Next)
		require.Nil(t, l.Back().Prev.Prev)

		l.PushBack(30)
		require.Equal(t, 30, l.Back().Value)
		require.Equal(t, 20, l.Back().Prev.Value)
		require.Nil(t, l.Back().Next)
		require.Nil(t, l.Back().Prev.Prev.Prev)

		require.Equal(t, []int{10, 20, 30}, convertListToSlice(l))
		require.Equal(t, []int{30, 20, 10}, convertListToReverseSlice(l))
	})

	t.Run("Len", func(t *testing.T) {
		l := NewList()
		require.Equal(t, 0, l.Len())

		l.PushBack(20)
		require.Equal(t, 1, l.Len())

		l.PushFront(30)
		require.Equal(t, 2, l.Len())

		l.Remove(l.Front())
		require.Equal(t, 1, l.Len())

		l.Remove(l.Back())
		require.Equal(t, 0, l.Len())
	})

	t.Run("Remove", func(t *testing.T) {
		l := NewList()
		l.PushBack(10)
		l.PushBack(20)
		l.PushFront(30)
		l.PushFront(40)

		l.Remove(l.Front())
		require.Equal(t, 30, l.Front().Value)
		l.Remove(l.Back())
		require.Equal(t, 10, l.Back().Value)

		l = NewList()
		l.PushBack(10)
		l.PushBack(20)
		l.PushFront(30)
		l.PushFront(40)

		l.Remove(l.Front())
		require.Equal(t, 30, l.Front().Value)
		l.Remove(l.Back())
		require.Equal(t, 10, l.Back().Value)
		l.Remove(l.Front())
		require.Equal(t, 10, l.Front().Value)
		l.Remove(l.Back())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, convertListToSlice(l))
		require.Equal(t, []int{50, 30, 10, 40, 60, 80, 70}, convertListToReverseSlice(l))
	})
}

func convertListToSlice(l List) []int {
	elems := make([]int, 0, l.Len())
	for i := l.Front(); i != nil; i = i.Next {
		elems = append(elems, i.Value.(int))
	}
	return elems
}

func convertListToReverseSlice(l List) []int {
	elems := make([]int, 0, l.Len())
	for i := l.Back(); i != nil; i = i.Prev {
		elems = append(elems, i.Value.(int))
	}
	return elems
}
