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

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
	t.Run("extra", func(t *testing.T) { // добавлены новые тесты
		l := NewList()
		l.PushFront(32)  // [32]
		l.PushBack(23)   // [32, 23]
		l.PushBack(34)   // [32, 23, 34]
		l.PushFront(100) // [100, 32, 23, 34]
		l.PushBack(93)   // [100, 32, 23, 34, 93]
		require.Equal(t, 5, l.Len())

		preLast := l.Back().Prev
		l.Remove(preLast) // [100, 32, 23, 93]
		require.Equal(t, 4, l.Len())
		require.Equal(t, 23, l.Back().Prev.Value.(int))

		for _, v := range [...]int{20, 110, 46} {
			if v >= l.Front().Value.(int) {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [110, 100, 32, 23, 93, 20, 46]
		arr := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			arr = append(arr, i.Value.(int))
		}
		require.Equal(t, []int{110, 100, 32, 23, 93, 20, 46}, arr)
	})
}
