package main

import "fmt"

func average() float64{	
	var num int
	fmt.Scanln(&num)
	
	
	
	var score, total int
	
	for i := 0; i < num; i++ {
		fmt.Scanln(&score)
		
		
		
		
	}
	
	avg := 
	
	return avg
}


func main() {
	
	
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
*/