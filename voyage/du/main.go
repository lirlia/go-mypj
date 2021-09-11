package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

func GetFileByteSize(dir string) (int, error) {
	fi, err := os.Stat(dir)
	if err != nil {
		return 0, err
	}

	return int(fi.Size()), nil
}

// 再起的にチェックしてディレクトリは以下のファイル一覧をだす
func GetFileListInDirectory(dir string) ([]string, error) {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			p, err := GetFileListInDirectory(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}
			paths = append(paths, p...)
			continue
		}

		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths, err
}

func GetFileListInDirectoryRoutine(dir string) ([]string, error) {
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, file := range files {
		wg.Add(1)
		f := file
		go func() {
			if f.IsDir() {
				p, _ := GetFileListInDirectoryRoutine(filepath.Join(dir, f.Name()))
				mutex.Lock()
				paths = append(paths, p...)
				mutex.Unlock()
			} else {
				mutex.Lock()
				paths = append(paths, filepath.Join(dir, f.Name()))
				mutex.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return paths, err
}

func DirsByteSize(dirs []string) (int, error) {
	var size, s int
	var err error
	var files []string

	for _, dir := range dirs {
		files, err = GetFileListInDirectory(dir)
		if err != nil {
			return 0, err
		}

		for _, file := range files {
			s, err = GetFileByteSize(file)
			if err != nil {
				return 0, err
			}

			size += s
			//fmt.Println(file)
		}
	}
	return size, err
}

func DirsByteSizeDirGoRoutine(dirs []string) (int, error) {
	var size, s int
	var err error
	var files []string
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	e := make(chan error)

	for _, dir := range dirs {
		d := dir

		wg.Add(1)
		go func() {
			defer wg.Done()
			files, err = GetFileListInDirectory(d)
			if err != nil {
				e <- err
			}

			for _, file := range files {
				s, err = GetFileByteSize(file)
				if err != nil {
					e <- err
				}

				mutex.Lock()
				size += s
				fmt.Println(size)
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	return size, <-e
}

func DirsByteSizeFileSearchGoRoutine(dirs []string) (int, error) {
	var size, s int
	var err error
	var files []string

	for _, dir := range dirs {
		files, err = GetFileListInDirectoryRoutine(dir)
		if err != nil {
			return 0, err
		}

		for _, file := range files {
			s, err = GetFileByteSize(file)
			if err != nil {
				return 0, err
			}

			size += s
			//fmt.Println(file)
		}
	}
	return size, err
}
