package main

import "fmt"

type student struct {
	name string
	gender string
	data map[string]int
}

func newStudent() *student {
	d := student {}
	d.data = map[string]int{}
	return &d
}

func main() {
	var numOfStudent, numOfSubject, score int
	var subjectName, name, gender string
	
	fmt.Scanln(&numOfStudent, &numOfSubject)
	
	var studentList = make([]student, numOfStudent)
	
	for i := 0; i < numOfStudent; i++  {	
		fmt.Scanln(&name, &gender)

		s1 := newStudent() //student 객체 생성

		s1.name = name
		s1.gender = gender
		
		for j := 0; j < numOfSubject; j++ {
			fmt.Scanln(&subjectName, &score)
            
			s1.data[subjectName] = score
		}
		studentList[i] = *s1
	}
	
	for i := 0; i < numOfStudent; i++  {
		fmt.Println("----------")
		fmt.Println(studentList[i].name, studentList[i].gender)
		
		for index, val := range studentList[i].data {
			fmt.Println(index, val)
		}
		
	}
	fmt.Println("----------")
}