package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) { // новые тесты
		c := NewCache(3)
		newC := c.Set("aaa", 10)
		require.False(t, newC)
		newC = c.Set("bbb", 15)
		require.False(t, newC)
		newC = c.Set("ccc", 20)
		require.False(t, newC)
		newC = c.Set("ddd", 30)
		require.False(t, newC)

		v, ok := c.Get("aaa")
		require.False(t, ok)
		require.Nil(t, v)
	})

	t.Run("purge with changes", func(t *testing.T) { // новые тесты
		c := NewCache(3)
		newC := c.Set("aaa", 10)
		require.False(t, newC)
		newC = c.Set("bbb", 15)
		require.False(t, newC)
		newC = c.Set("ccc", 20)
		require.False(t, newC)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 10, val)
		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 15, val)

		newC = c.Set("ddd", 40)
		require.False(t, newC)

		v, ok := c.Get("ссс")
		require.False(t, ok)
		require.Nil(t, v)
	})

	t.Run("Clear function", func(t *testing.T) { // новый тест для проверки метода clear()
		c := NewCache(3)
		newC := c.Set("aaa", 10)
		require.False(t, newC)
		newC = c.Set("bbb", 15)
		require.False(t, newC)
		newC = c.Set("ccc", 20)
		require.False(t, newC)

		c.Clear()
		v, ok := c.Get("aaa")
		require.Nil(t, v)
		require.False(t, ok)
		v, ok = c.Get("bbb")
		require.Nil(t, v)
		require.False(t, ok)
		v, ok = c.Get("ccc")
		require.Nil(t, v)
		require.False(t, ok)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
