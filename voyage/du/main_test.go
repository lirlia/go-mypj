package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDu(t *testing.T) {

	var size int
	var err error
	var dirs []string

	dirs = []string{"test"}
	size, err = DirsByteSize(dirs)
	assert.Equal(t, err, nil)
	assert.Equal(t, size, 1024+2048)

	dirs = []string{"test", "test1"}
	size, err = DirsByteSize(dirs)
	assert.Equal(t, err, nil)
	assert.Equal(t, size, 1024+2048+1024+2048)

	dirs = []string{"test", "test10"}
	size, err = DirsByteSize(dirs)
	assert.Equal(t, err, nil)
	assert.Equal(t, size, 2048*5+1024)

	dirs = []string{"test2"}
	size, err = DirsByteSize(dirs)
	assert.NotEqual(t, err, nil)
	assert.Equal(t, size, 0)

	dirs = []string{"test2", "test3"}
	size, err = DirsByteSize(dirs)
	assert.NotEqual(t, err, nil)
	assert.Equal(t, size, 0)

	dirs = []string{"/Users/"}
	start := time.Now()
	size, _ = DirsByteSize(dirs) // 49.889813秒
	//size, _ = DirsByteSizeDirGoRoutine(dirs)
	//size, _ = DirsByteSizeFileSearchGoRoutine(dirs) // 2.407401秒
	end := time.Now()
	fmt.Printf("普通: %f秒\n", (end.Sub(start)).Seconds())
	fmt.Println(size)

	start = time.Now()
	size, _ = DirsByteSizeFileSearchGoRoutine(dirs) // 2.407401秒
	end = time.Now()
	fmt.Printf("高速: %f秒\n", (end.Sub(start)).Seconds())
	fmt.Println(size)

}
