package main

import (
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := &Stack{Limit: 5}

	// 一ついれて取り出す
	s.Push("test")
	assert.Equal(t, s.Pop(), "test")

	// 存在しないいときに取り出す
	assert.Equal(t, s.Pop(), "")

	// 2ついれて2つ取り出す
	s.Push("test1")
	s.Push("test2")
	assert.Equal(t, s.Pop(), "test2")
	assert.Equal(t, s.Pop(), "test1")
	assert.Equal(t, s.Pop(), "")

	// 3ついれて2つ取り出して再度入れて取り出す
	s.Push("test1")
	s.Push("test2")
	s.Push("test3")
	assert.Equal(t, s.Pop(), "test3")
	assert.Equal(t, s.Pop(), "test2")
	s.Push("test4")
	assert.Equal(t, s.Pop(), "test4")
	assert.Equal(t, s.Pop(), "test1")
	assert.Equal(t, s.Pop(), "")

	// Stackがないときに連続でPopしても問題ないか
	assert.Equal(t, s.Pop(), "")
	assert.Equal(t, s.Pop(), "")

	// limit関連(limit 5)
	s.Push("test1")
	s.Push("test2")
	s.Push("test3")
	s.Push("test4")
	s.Push("test5")
	s.Push("test6")
	assert.Equal(t, s.Pop(), "test6")
	assert.Equal(t, s.Pop(), "test5")
	assert.Equal(t, s.Pop(), "test4")
	assert.Equal(t, s.Pop(), "test3")
	assert.Equal(t, s.Pop(), "test2")
	assert.Equal(t, s.Pop(), "")

	// limit変更
	s.Limit = 2
	s.Push("test1")
	s.Push("test2")
	s.Push("test3")
	assert.Equal(t, s.Pop(), "test3")
	assert.Equal(t, s.Pop(), "test2")
	assert.Equal(t, s.Pop(), "")

	// go routine対応 stack
	s.Limit = 100
	var wg sync.WaitGroup
	for i := 0; i < s.Limit; i++ {
		val := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Push(strconv.Itoa(val))
		}()
	}

	wg.Wait()
	assert.Equal(t, len(s.Value), s.Limit)

	// go routine対応 pop
	s = &Stack{Limit: 100}

	assert.Equal(t, len(s.Value), 0)

	for i := 0; i < s.Limit; i++ {
		s.Push(strconv.Itoa(i))
	}
	assert.Equal(t, len(s.Value), s.Limit)

	for i := 0; i < s.Limit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Pop()
		}()
	}

	wg.Wait()
	assert.Equal(t, len(s.Value), 0)

}
