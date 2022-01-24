package main

import (
	"fmt"
    "log"
)

func errorChek(score int) (int, error) {
    if score >= 0 {
        return score, nil
    }
    return 0, fmt.Errorf("시험 점수를 양의 정수로 입력하세요.")
}

func main() {
    var score int
    fmt.Scanln(&score)
    
    s, err := errorChek(score)
    
    if err != nil {
        log.Panic(err)
    }
    fmt.Printf("시험 점수는 %d점입니다.",s)
}