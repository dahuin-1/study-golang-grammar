package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan string)

	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			c := <- ch3 //수신자를 의미하는 "<- 채널이름"은 두 개의 값을 반환합니다. 첫 번째는 채널 데이터, 두 번째는 채널의 개폐 여부
			fmt.Printf("ch3의 수신 데이터 %s\n", c)
		}
	}()

	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			ch1 <- 10
		}
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			ch2 <- 20
		}
	}()

	go func ()  {
		for {
			select {
				case a := <-ch1:
				fmt.Printf("ch1의 수신데이터 %d\n", a)
				case b := <-ch2:
				fmt.Printf("ch2의 수신데이터 %d\n", b)
				case ch3 <- "groom":
			}
		}
	}()

	time.Sleep(5 * time.Second)

}