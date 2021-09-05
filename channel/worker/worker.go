package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	WORKER_COUNT = 100
	QUEUE_COUNT  = 1000000
)

func worker(wg *sync.WaitGroup, c chan int) {
	//id := rand.Intn(100)
	defer wg.Done()
	for {
		_, ok := <-c
		if !ok {
			return
		}
		//fmt.Printf("worker_id:%v, number:%v\n", id, number)
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	// あらかじめQUEUEの数だけバッファを作成しておく
	// 作成したほうが早い
	queue := make(chan int, QUEUE_COUNT)

	// worker実行
	// 並列に起動することで少しでも起動時間を短縮しようという試み
	for i := 0; i < int(math.Sqrt(WORKER_COUNT)); i++ {
		go func() {
			for j := 0; j < int(math.Sqrt(WORKER_COUNT)); j++ {
				wg.Add(1)
				go worker(&wg, queue)
			}
		}()
	}

	beforeTime := time.Now()

	// Queueにデータを格納
	var wgQueue sync.WaitGroup
	for j := 0; j < int(math.Sqrt(QUEUE_COUNT)); j++ {
		go func() {
			for k := 0; k < int(math.Sqrt(QUEUE_COUNT)); k++ {
				wgQueue.Add(1)
				queue <- rand.Intn(1000)
				defer wgQueue.Done()
			}
		}()
	}

	wgQueue.Wait()
	close(queue)
	wg.Wait()
	afterTime := time.Now()
	duration := afterTime.Sub(beforeTime)
	fmt.Printf("duration: %v, worker:%v , queue:%v\n", duration, WORKER_COUNT, QUEUE_COUNT)
}
