package main

import "fmt"

type student struct {
	/* 
	학생 이름 정보를 담는 string형 변수, 
	학생 성별 정보를 담는 string형 변수, 
	key 값으로 과목 이름, value 값으로 점수를 담는 맵형 컬렉션
	*/
	name 	string
	gender 	string
	data 	map[string]int
}

func newStudent() *student { //포인터 구조체를 반환함
	d := student{} //구조체 객체를 생성하고 초기화함
	d.data = map[string]int{}
	return &d //초기화 한 포인터 구조체를 반환함
}

func main() {
	var studentNum, subjetNum int
	var subjetName, name, gender string
	
	fmt.Scanln(&studentNum, &subjetNum)
	
	var s = make([]student, numOfStudent)
	
	for i := range studentNum { {	
		fmt.Scanln(&name, &gender)
		s1 := newStudent() //생성자 호출을 통한 student 객체 생성

		s1.name = name
		s1.gender = gender
		
		for j := range subjetNum {
			fmt.Scanln(&subjetName, &score)

			s1.data[subjectName] = score
		}
		s[i] = *s1
	}
	
	for i := range studentNum {
		fmt.Println("----------")
		
		
		for index, val := range s[i].score {
			fmt.Println(index, val)
		}
		
	}
	fmt.Println("----------")
}