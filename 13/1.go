package main

import "fmt"

type student struct {
	name string
	gender string
	data map[string]int
}

func newStudent() *student { //포인터 구조체를 반환함
	d := student{} //구조체 객체를 생성하고 초기화함
	d.data = map[string]int{}
	return &d //초기화 한 포인터 구조체를 반환함
}

func main() {
	var numOfStudent, numOfSubject, score int
	var name, gender, subject string
	
	fmt.Scanln(&numOfStudent, &numOfSubject)

	//s1 := newStruct()
    var s = make([]student, numOfStudent)
	
	for i := 0; i < numOfStudent; i++ {	
		fmt.Scanln(&name, &gender)
		s1 := newStudent() 

		s1.name = name
		s1.gender = gender
		
		for i := 0; i < numOfSubject; i++ {
			fmt.Scanln(&subject, &score)
            //info[subject] = score
			s1.data[subject] = score
		}
		s[i] = *s1
	}
	
	for i := 0; i < numOfStudent; i++ {
		fmt.Println("----------")
		fmt.Println(s[i].name, s[i].gender)
		
		for index, val := range s[i].data {
			fmt.Println(index, val)
		}
		
	}
	fmt.Println("----------")
}