package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

func hello(n int, w *sync.WaitGroup) {
	defer w.Done() // 끝났음을 전달

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)
	fmt.Println(n)
}

func main() {
	wait := new(sync.WaitGroup) //wait group 생성
	wait.Add(100) // 100개의 고루틴을 기다림
	for i := 0; i < 100; i++ {
		go hello(i, wait)
	}
	wait.Wait()
}

