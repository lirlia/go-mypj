// fib_test.go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	r := &result{
		value: map[int]int{},
	}

	if r == nil {
		t.Errorf("failed NewUser()")
	}

	var v int
	var err error
	v, err = r.fib(0)
	assert.Equal(t, v, 0)
	assert.Equal(t, err, nil)

	v, err = r.fib(1)
	assert.Equal(t, v, 1)
	assert.Equal(t, err, nil)

	v, err = r.fib(2)
	assert.Equal(t, v, 1)
	assert.Equal(t, err, nil)

	v, err = r.fib(50)
	assert.Equal(t, v, 12586269025)
	assert.Equal(t, err, nil)

	v, err = r.fib(-1)
	assert.Equal(t, v, 0)
	assert.EqualError(t, err, "最小値は0です")
}
