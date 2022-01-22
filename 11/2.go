package main

import "fmt"


func inputNums() []int {
	var score int
	//fmt.Scanln(&score)
	nums := make([]int, 5)
	
	for i := 0; i<5; i++{
		fmt.Scanln(&score)
		nums[i] = score
	}
	return nums
}


func calExam(arr []int) (totalScore int, overCount int, underCount int) {
	overCount = 0;
	underCount = 0;
	for i := 0; i < len(arr); i++ {
		totalScore += arr[i]
		if arr[i] >= 90 {
			overCount++;
		}
		if arr[i] < 70 {
			underCount++;
		}
	}
	return
}



func printResult (totalScore int, overCount int, underCount int) {
	var result bool = true
	if totalScore < 400 {
		fmt.Println("총점이 400점 미만입니다.")
		result = false
	} 
	if overCount < 2 {
		fmt.Println("90점 이상 과목 수가 2개 미만입니다.")
		result = false
	} 
	if underCount > 0 {
		fmt.Println("70점 미만 과목이 있습니다.")
		result = false
	}
	if result == false {
		fmt.Println("아이패드를 살 수 없습니다.")
	}else{
		fmt.Println("아이패드를 살 수 있습니다.")
	}
}

func main() {	
	arr := inputNums()
	//totalScore, overCount, underCount := calExam(arr)
	//printResult(totalScore, overCount, underCount)
	printResult(calExam(arr))
}