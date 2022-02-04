package main

import "fmt"

var gravity = 9.8

type object struct {
	m 	float32
	v 	float32
	h 	float32
	ke 	float32
	pe 	float32
	me 	float32
}

func (o object) getKe() float32 { //value receiver
	return o.m * o.v * o.v / 2 //k운동에너지
}

func (o object) getPe() float32 { //value receiver
	return o.m * float32(gravity) * o.h //p위치에너지
}

func main() {
	//물체의 개수, 질량, 속도, 높이
	//var mass, velocity, height float32
	var numOfObject int
	
	fmt.Scanln(&numOfObject)	
	
	objectList := make([]object, numOfObject)

	for i := 0; i < numOfObject; i++  {	
		fmt.Scanln(&objectList[i].m, &objectList[i].v, &objectList[i].h)

		objectList[i].ke = objectList[i].getKe()
		objectList[i].pe = objectList[i].getPe()
		objectList[i].me = objectList[i].ke + objectList[i].pe
	}

	for i := 0; i < numOfObject; i++ {
		fmt.Println(objectList[i].ke, objectList[i].pe, objectList[i].me)
	}
}