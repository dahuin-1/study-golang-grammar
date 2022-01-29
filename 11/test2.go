package main

import "fmt"


func inputNums() []int  {
	var score int
	nums := make([]int, 5)
	
	for i := 0; i<5; i++{
		fmt.Scanln(&score)
		nums[i] = score
	}
	
	return nums
}


func calExam(arr []int) (sum int, over90 int, under70 int) {
	//sum, over90, under70 = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 90 {
			over90 ++
		}
		if arr[i] < 70 {
			under70 ++
		}
		sum = sum + arr[i]
	}
	
	return
}



func printResult (sum int, over90 int, under70 int) {
	var result bool = true
	if sum < 400 {
		result = false
		fmt.Println("총점이 400점 미만입니다.")
	} 
	if over90 < 2 {
		result = false
		fmt.Println("90점 이상 과목 수가 2개 미만입니다.")
	} 
	if under70 > 0 {
		result = false
		fmt.Println("70점 미만 과목이 있습니다.")
	}
	if result == false{
		fmt.Println("아이패드를 살 수 없습니다")
	}else{
		fmt.Println("아이패드를 살 수 있습니다")
	}
}

func main() {	
	scores := inputNums()
	printResult(calExam(scores))
}