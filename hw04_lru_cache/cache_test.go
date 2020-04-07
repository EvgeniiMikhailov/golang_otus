package hw04_lru_cache //nolint:golint,stylecheck

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

	t.Run("write and read single element", func(t *testing.T) {
		c := NewCache(10)

		wasInCache := c.Set("a", 1)
		require.False(t, wasInCache)

		value, ok := c.Get("a")
		require.True(t, ok)
		require.Equal(t, 1, value)
	})

	t.Run("eviction first element", func(t *testing.T) {
		c := NewCache(2)

		wasInCache := c.Set("a", 1)
		require.False(t, wasInCache)

		wasInCache = c.Set("b", 2)
		require.False(t, wasInCache)

		wasInCache = c.Set("c", 3)
		require.False(t, wasInCache)

		value, ok := c.Get("a")
		require.False(t, ok)
		require.Nil(t, value)

		value, ok = c.Get("b")
		require.True(t, ok)
		require.Equal(t, 2, value)

		value, ok = c.Get("c")
		require.True(t, ok)
		require.Equal(t, 3, value)
	})

	t.Run("eviction second element", func(t *testing.T) {
		c := NewCache(2)

		wasInCache := c.Set("a", 1)
		require.False(t, wasInCache)

		wasInCache = c.Set("b", 2)
		require.False(t, wasInCache)

		value, ok := c.Get("a")
		require.True(t, ok)
		require.Equal(t, 1, value)

		wasInCache = c.Set("c", 3)
		require.False(t, wasInCache)

		value, ok = c.Get("b")
		require.False(t, ok)
		require.Nil(t, value)

		value, ok = c.Get("c")
		require.True(t, ok)
		require.Equal(t, 3, value)
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

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(10)

		wasInCache := c.Set("a", 1)
		require.False(t, wasInCache)

		value, ok := c.Get("a")
		require.True(t, ok)
		require.Equal(t, 1, value)

		c.Clear()
		value, ok = c.Get("a")
		require.False(t, ok)
		require.Nil(t, value)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove if task with asterisk completed

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
