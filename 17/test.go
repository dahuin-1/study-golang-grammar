package main

import "fmt"

func testGo() {
	fmt.Println("Hello goorm!")
}

func main() {
	go testGo() //고루틴으로 비동기 실행
	fmt.Scanln()
}