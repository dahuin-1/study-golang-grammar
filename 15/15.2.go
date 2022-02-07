package main

import "fmt"

func average() float64{	
	var num int
	fmt.Scanln(&num)
	
	if num <= 0 {
		defer fmt.Println("잘못된 과목 수입니다.")
	}
	
	var score, total int
	
	for i := 0; i < num; i++ {
		fmt.Scanln(&score)
		
		if score < 0 {
			defer fmt.Println("잘못된 점수입니다.")
		}
		
		total += score
		
	}
	
	avg := float64(total / num)
	
	return avg
}


func main() {
	defer func() {
		if r := recover(); r != nil{
			fmt.Println(r)
			
			main()
		}		
	}()
	
	result := average()	
	fmt.Println(result)	
}

/*
package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil{
			fmt.Println(r)
			
			main()
		}		
	}()
	
	var num1, num2 int
	fmt.Scanln(&num1, &num2)
	
	result := num1 / num2

	fmt.Println(result)
}
package main

import "fmt"

func panicTest() {
	var a = [4]int{1,2,3,4}
	
	defer fmt.Println("Panic done")
	
	for i := 0; i < 10; i++ {
		fmt.Println(a[i])
	}		
}

func main() {
	panicTest()

	fmt.Println("Hello, world!")
}
*/