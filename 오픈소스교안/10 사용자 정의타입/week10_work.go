package main

import "fmt"

//type game[number]name map[int]string
//type game[int]string map[int]string
/*type game struct{
	game map[int]string
	 단일변수타입만 보다가 두개 받아야 하는 맵을 사용자 정의 타입으로 만들려다가 많이 헤맴..
}*/

type game map[int]string //사용자 정의 타입 정의

// func (g game) Find(number int) {
// 	for k, v := range g {
// 		if number == k {
// 			fmt.Printf("찾으시는 번호는 참가자 %s입니다.", v)
// 			break
// 		}
// 	}
// 	fmt.Println("찾는 참가자가 없습니다")
// } //찾는 번호를 매개변수로 받아 맵을 반복해서 돌려 찾는 번호와 같으면
//찾는 번호의 v값을 print해주고 break로 반복문을 빠져나오는 메소드
func (g game) Find(number int) string {
	for k, v := range g {
		if number == k {
			return v
		}
	}
	return "찾는 참가자가 없습니다"
} //리턴타입이 있는 메서드로도 만들어봄

func main() {
	m1 := game{456: "성기훈", 218: "박해수", 67: "강새벽", 1: "오일남", 199: "알리", 101: "아이오아이"}
	// m1.Find(1) 리턴 없는 메서드
	//fmt.Println(m1.Find(1)) 오일남 나옴
	// fmt.Printf("찾으시는 번호는 참가자 %s입니다.", m1.Find(1))
	//fmt.Println(m1.Find(6)) 참가자 없다고 나옴
	var n1 int
	fmt.Print("참가자 번호 입력 : ")
	fmt.Scanln(&n1)
	fmt.Printf("찾으시는 번호는 참가자 %s입니다.", m1.Find(n1))
}
