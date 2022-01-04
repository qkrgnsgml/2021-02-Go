package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) (bool, error) {
	if n < 2 {
		return false, fmt.Errorf("%d 는(은) 소수가 아닙니다~", n)
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}

	return true, nil // 소수 판정 값, 정상데이터
}

func primeRange(a int, b int) (float64, []int) {
	avg := 0.0
	sum := 0.0
	c := []int{}
	if a > b {
		temp := a
		a = b
		b = temp
	}

	for i := a; i <= b; i++ {
		p, err := isPrime(i)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if p {
			c = append(c, i)
			sum += float64(i)
		}
	}
	avg = sum / float64(len(c))
	return avg, c
}

func main() {
	sl := []int{}
	avg := 0.0
	var n1, n2 int
	fmt.Print("정수 2개 입력 : ")
	_, err := fmt.Scanln(&n1, &n2)
	if err != nil {
		log.Fatal(err)
	}
	avg, sl = primeRange(n1, n2)
	fmt.Println("소수 목록 : ", sl)
	fmt.Printf("평균 : %.2f", avg)
}
