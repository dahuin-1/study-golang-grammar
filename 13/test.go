package main

import "fmt"

var gravity = 9.8

/*
질량(m), 속도(v), 높이(h), 운동 에너지(ke), 위치 에너지(pe), 역학적 에너지(me)를 필드 값으로 가지는 떨어지는 물체의 구조체를 선언합니다. 
필드 값 자료형은 모두 float32입니다.
*/
type object struct {
	m 	float32
	v 	float32
	h 	float32
	ke 	float32
	pe 	float32
	me 	float32
}

func (o object) getKe() float32 { //위치에너지
	return o.m * o.v * o.v / 2
}

func (o object) getPe() float32 { //위치에너지
	return o.m * o.h * float32(gravity)
}


func main() {
	var numOfObject int
	//사용자에게 입력받을 물체의 개수, 질량, 속도, 높이 변수를 선언합
	//var m, v, h float32
	fmt.Scanln(&numOfObject)	

	objectList := make([]object, numOfObject)
	
	for i := 0; i < numOfObject; i++ {	
		fmt.Scanln(&objectList[i].m, &objectList[i].v, &objectList[i].h)

		objectList[i].ke = objectList[i].getKe()
		objectList[i].pe = objectList[i].getPe()
		objectList[i].me = objectList[i].ke + objectList[i].pe
	}
	
	for i := 0; i < numOfObject; i++ {
		fmt.Println(objectList[i].ke, objectList[i].pe, objectList[i].me)
	}
	
}

