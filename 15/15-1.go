package main

import "fmt"

func main() {
	var nameList []string
	var name string
	
	for {
		fmt.Scanln(&name)
		if name == "0" {
			break
		} else {
			nameList = append(nameList, name)
		}
	}
	
	for i := 0; i < len(nameList); i++ {
		defer fmt.Println(nameList[i])
	}	
}