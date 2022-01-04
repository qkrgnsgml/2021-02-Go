package main

import (
	"fmt"
	"gowork/chap11/gaget"
)

type Player interface {
	Play(string)
	Stop()
}

// func playList(device gaget.TapePlayer , songs []string){
func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player Player) {
	player.Play("데모트랙")
	player.Stop()
	// player.Record()  player 인터페이스에 Record() 가 없음

	// recorder := gaget.TapeRecoder(player)  인터페이스는 형 변환 문법이 기존 언어와 다르다.
	// recorder.Record()

	recorder := player.(gaget.TapeRecoder) // 인터페이스 형 변환 문법
	recorder.Record()
}

func main() {
	TryOut(gaget.TapeRecoder{})

	player := gaget.TapePlayer{}
	// player := gaget.TapeRecoder{}
	fmt.Println(player)
	//mixtape := []string{"난 알아요", "거짓말", "으르렁"}
	//playList(player, mixtape)
}
