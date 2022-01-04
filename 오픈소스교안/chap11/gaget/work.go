package gaget

import "fmt"

type TapePlayer struct {
	Batteries int
	sor       string
}

func (t TapePlayer) Play(song string) {
	fmt.Println(song, " 재생 중...")
}

func (t TapePlayer) Stop() {
	fmt.Println("중지")
}

type TapeRecoder struct {
	Microphones int
}

func (t TapeRecoder) Play(song string) {
	fmt.Println(song, " 재생 중...")
}

func (t TapeRecoder) Record() {
	fmt.Println("녹화 중...")
}

func (t TapeRecoder) Stop() {
	fmt.Println("중지")
}
