/*
사용자에게 슬라이스를 이용해 무작위로 정수를 입력 받고, 입력 받은 정수를 오름차순으로 정렬하고 출력
*/
package main

import "fmt"

func bubbleSort(nums []int) {
    var temp int
	//var i, j int
	var n = len(nums)
	
	for i := n-1; i>0; i-- {
		// 0 ~ (i-1)까지 반복
		for j := 0; j < i; j++ {
		  // j번째와 j+1번째의 요소가 크기 순이 아니면 교환
		  if nums[j] > nums[j+1] {
			temp = nums[j];
			nums[j] = nums[j+1];
			nums[j+1] = temp;
		  }
		}
	  }
}

func inputNums() (nums []int) {	
	var t, num int
	fmt.Scan(&t)
	nums = make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Scanln(&num)
		nums[i] = num
		//fmt.Scanln("%d", &nums[i])
	}
	return nums
}

func outputNums(nums []int) {
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d ", nums[i])
		//fmt.Fprintf(writer, "%d ", nums[i])
	}
}

func main() {	
	//writer := bufio.NewWriter(os.Stdout)
	nums := inputNums()
	bubbleSort(nums)
	outputNums(nums)
}