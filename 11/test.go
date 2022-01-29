package main

import "fmt"

var gravity = 9.8

type cal func(float32, float32) float32
//calMechEnergy() 함수에 매개변수로서 함수를 사용하기 위해 type문으로 main 내 두 함수의 원형을 정의

func calMechEnergy(f cal, x float32, y float32) float32 {
    result := f(x, y)
    return result
}

func main() {
	var m, v, h float32
	fmt.Scanln(&m, &v, &h)
	

	kinEnergy := func(m float32, v float32) float32{ //운동 에너지를 계산하는 익명 함수를 kinEnergy 변수에 초기화
		return (m * v * v) / 2
	}
	potEnergy := func(m float32, h float32) float32{
		return m * float32(gravity) * h
	}
	
	ke := calMechEnergy(kinEnergy, m, v)
	pe := calMechEnergy(potEnergy, m, v)
	fmt.Println(ke, pe, ke+pe)
}