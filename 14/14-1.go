package main

import (
	"fmt"
	"math"
)

type geometry interface { //인터페이스 선언 Reat와 Circle 메도스의 area를 모두 포함
	area() float64
	volume() float64
}

type cylinder struct { //원기둥
	r float64
	h float64
}

type cuboid struct { //직육면체
	w float64
	h float64
	l float64
}

func (cy cylinder) area() float64 { //원기둥 겉넓이
	return ((math.Pi * cy.r * cy.r) * 2) + (cy.h * (2 * math.Pi * cy.r))
}

func (cy cylinder) volume() float64 { //원기둥 부피
	return (math.Pi * cy.r * cy.r) * cy.h
}

func (cu cuboid) area() float64 {
	return (2 * cu.w * cu.h) + (2 * cu.w * cu.l) + (2 * cu.h * cu.l)
}

func (cu cuboid) volume() float64 {
	return cu.h * cu.l * cu.w
}

func main() {
	cy1 := cylinder{10, 10}
	cy2 := cylinder{4.2, 15.6}
	cu1 := cuboid{}
	cu2 := cuboid{}
	
	printMeasure(cy1, cy2, cu1, cu2)	
}

func printMeasure(m ...geometry) {
	for _, val := range m {
		fmt.Println(val.area())
		fmt.Println(val.volume())
	}
}