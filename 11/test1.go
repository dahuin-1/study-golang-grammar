package main

import "fmt"

func bubbleSort(list []int) {
	var temp int
	n := len(list)
	for i := n-1; i > 0; i-- {
		// 0 ~ (i-1)까지 반복
		for j := 0; j<i; j++ {
		  // j번째와 j+1번째의 요소가 크기 순이 아니면 교환
		  if(list[j] > list[j+1]){
			temp = list[j];
			list[j] = list[j+1];
			list[j+1] = temp;
		  }
		}
	  }
}

func inputNums()(nums []int) {	
	var t, n int
	fmt.Scanln(&t)
	nums = make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Scanln(&n)
		nums[i] = n
	}
	return nums
}

func outputNums(list []int) {
	for i := range list {
		fmt.Printf("%d ", list[i])
	}

}

func main() {	
	nums := inputNums()
	bubbleSort(nums)
	outputNums(nums)
}