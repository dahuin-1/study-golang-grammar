package main

import "fmt"

/*타입문 작성은 선택입니다*/

func calCoin() func(a int, b int) int { 
	return func(a int, b int) int { //클로저
		return a * b
	}
}

func main() {
	var coin10, coin50, coin100, coin500 int
	fmt.Scan(&coin10, &coin50, &coin100, &coin500)
	
  if coin10 < 0 || coin50 < 0 || coin100 < 0 || coin500 < 0{
		fmt.Println("잘못된입력입니다.")
		return
	}
	
	f := calCoin()
	
	add10 := f(10, coin10)
	add50 := f(50, coin50)
	add100 := f(100, coin100)
	add500 := f(500, coin500)
	
	totalmoney := add10 + add50 + add100 + add500
	
	fmt.Println(totalmoney)	
}