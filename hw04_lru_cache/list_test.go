package hw04_lru_cache //nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, l.Len(), 0)
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("push front two elements", func(t *testing.T) {
		l := NewList()

		l.PushFront(0) // [0]
		l.PushFront(1) // [1, 0]
		require.Equal(t, l.Len(), 2)
		require.Equal(t, 1, l.Front().Value)
		require.Equal(t, 0, l.Back().Value)
	})

	t.Run("push back two elements", func(t *testing.T) {
		l := NewList()

		l.PushBack(0) // [0]
		l.PushBack(1) // [0, 1]
		require.Equal(t, l.Len(), 2)
		require.Equal(t, 0, l.Front().Value)
		require.Equal(t, 1, l.Back().Value)
	})

	t.Run("check middle of free elements", func(t *testing.T) {
		l := NewList()

		l.PushBack(0) // [0]
		l.PushBack(1) // [0, 1]
		l.PushBack(2) // [0, 1, 2]
		require.Equal(t, l.Len(), 3)
		require.Equal(t, l.Front().Prev.Value, l.Back().Next.Value)
	})

	t.Run("remove front", func(t *testing.T) {
		l := NewList()

		l.PushBack(0) // [0]
		l.PushBack(1) // [0, 1]

		l.Remove(l.Front()) // [1]
		require.Equal(t, l.Len(), 1)
		require.Equal(t, l.Front().Value, l.Back().Value)
		require.Equal(t, 1, l.Front().Value)
	})

	t.Run("remove single item", func(t *testing.T) {
		l := NewList()

		l.PushFront(0) // [0]

		l.Remove(l.Back()) // []
		require.Equal(t, l.Len(), 0)
	})

	t.Run("remove back", func(t *testing.T) {
		l := NewList()

		l.PushBack(0) // [0]
		l.PushBack(1) // [0, 1]

		l.Remove(l.Back()) // [0]
		require.Equal(t, l.Len(), 1)
		require.Equal(t, l.Front().Value, l.Back().Value)
		require.Equal(t, 0, l.Front().Value)
	})

	t.Run("remove middle", func(t *testing.T) {
		l := NewList()

		l.PushBack(0)           // [0]
		l.PushBack(1)           // [0, 1]
		l.PushBack(2)           // [0, 1, 2]
		l.Remove(l.Back().Next) // [0, 2]
		require.Equal(t, l.Len(), 2)
		require.Equal(t, 0, l.Front().Value)
		require.Equal(t, 2, l.Back().Value)
	})

	t.Run("remove middle of free elements", func(t *testing.T) {
		l := NewList()

		l.PushBack(0)            // [0]
		l.PushBack(1)            // [0, 1]
		l.PushBack(2)            // [0, 1, 2]
		l.Remove(l.Front().Prev) // [0, 2]
		require.Equal(t, l.Len(), 2)
		require.Equal(t, 0, l.Front().Value)
		require.Equal(t, 2, l.Back().Value)
	})

	t.Run("move to front back item", func(t *testing.T) {
		l := NewList()

		l.PushBack(0)           // [0]
		l.PushBack(1)           // [0, 1]
		l.PushBack(2)           // [0, 1, 2]
		l.MoveToFront(l.Back()) // [2, 0, 1]
		require.Equal(t, l.Len(), 3)
		require.Equal(t, 2, l.Front().Value)
		require.Equal(t, 0, l.Front().Prev.Value)
		require.Equal(t, 1, l.Front().Prev.Prev.Value)
		require.Equal(t, 2, l.Back().Next.Next.Value)
		require.Equal(t, 0, l.Back().Next.Value)
		require.Equal(t, 1, l.Back().Value)
	})

	t.Run("move to front middle item", func(t *testing.T) {
		l := NewList()

		l.PushBack(0)                // [0]
		l.PushBack(1)                // [0, 1]
		l.PushBack(2)                // [0, 1, 2]
		l.MoveToFront(l.Back().Next) // [1, 0, 2]
		require.Equal(t, l.Len(), 3)
		require.Equal(t, 1, l.Front().Value)
		require.Equal(t, 0, l.Front().Prev.Value)
		require.Equal(t, 2, l.Front().Prev.Prev.Value)
		require.Equal(t, 1, l.Back().Next.Next.Value)
		require.Equal(t, 0, l.Back().Next.Value)
		require.Equal(t, 2, l.Back().Value)
	})

	t.Run("move to front front item", func(t *testing.T) {
		l := NewList()

		l.PushBack(0)            // [0]
		l.PushBack(1)            // [0, 1]
		l.PushBack(2)            // [0, 1, 2]
		l.MoveToFront(l.Front()) // [0, 1, 2]
		require.Equal(t, l.Len(), 3)
		require.Equal(t, 0, l.Front().Value)
		require.Equal(t, 1, l.Front().Prev.Value)
		require.Equal(t, 2, l.Front().Prev.Prev.Value)
		require.Equal(t, 0, l.Back().Next.Next.Value)
		require.Equal(t, 1, l.Back().Next.Value)
		require.Equal(t, 2, l.Back().Value)
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, l.Len(), 3)

		middle := l.Back().Next // 20
		l.Remove(middle)        // [10, 30]
		require.Equal(t, l.Len(), 2)

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, l.Len(), 7)
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Back(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{50, 30, 10, 40, 60, 80, 70}, elems)
	})
}
