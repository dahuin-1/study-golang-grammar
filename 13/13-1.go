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
//생성자함수
func newStudent() *student { //포인터 구조체를 반환함
	s := student{} //구조체 객체를 생성하고 초기화함
	s.data = map[string]int{}
	return &s //초기화 한 포인터 구조체를 반환함
}

func main() {
	var studentNum, subjectNum, score int
	var subjectName, name, gender string
	
	fmt.Scanln(&studentNum, &subjectNum)
	
	var studentList = make([]student, studentNum)
	
	for i :=0; i < studentNum; i++ { 
		fmt.Scanln(&name, &gender)
		s1 := newStudent() //생성자 호출을 통한 student 객체 생성

		s1.name = name
		s1.gender = gender
		
		for j :=0; j < subjectNum; j++ {
			fmt.Scanln(&subjectName, &score)

			s1.data[subjectName] = score
		}
		studentList[i] = *s1
	}
	
	for i :=0; i < studentNum; i++ { 
		fmt.Println("----------")
		fmt.Println(studentList[i].name, studentList[i].gender)
		
		for index, val := range studentList[i].data {
			fmt.Println(index, val)
		}
		
	}
	fmt.Println("----------")
	
}