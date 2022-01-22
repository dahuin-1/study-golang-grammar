package main

import "fmt"

type mapStruct struct {
	data map[int]string
}

func newStruct() *mapStruct { //포인터 구조체를 반환함
	d := mapStruct{} //구조체 객체를 생성하고 초기화함
	d.data = map[int]string{} //data 필드의 맵을 초기화함
	return &d //초기화 한 포인터 구조체를 반환함
}

func main() {
	s1 := newStruct() // 생성자 호출해서 맵을 초기화 -> 바로 data 필드에 값을 저장
	s1.data[1] = "hello"
	s1.data[10] = "world"

	fmt.Println(s1)

	s2 := mapStruct{} // 구조체만 생성
	s2.data = map[int]string{} //data 필드에 값을 저장하기 위해 선언이 필요한 맵은 따로 초기화
	s2.data[1] = "hello"
	s2.data[10] = "world"

	fmt.Println(s2)
}

//&{map[1:hello 10:world]} 출력
//{map[1:hello 10:world]} 출력