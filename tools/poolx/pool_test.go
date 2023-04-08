package poolx

import (
	"fmt"
	"testing"
	"time"
)

func TestNewFnPool(t *testing.T) {
	p := NewFnPool(fn, 10)
	err := p.Invoke([]byte("test"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	time.Sleep(1 * time.Second)
}

func fn(in []byte) {
	fmt.Println("函数打印", string(in))
}

func TestNewFnPool2(t *testing.T) {
	play := Player{Name: "Abe"}
	p := NewFnPool(play.Play, 10)
	err := p.Invoke([]byte("video game"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	time.Sleep(1 * time.Second)

}

type Player struct {
	Name string `json:"name"`
}

func (p *Player) Play(in []byte) {
	fmt.Println(p.Name, "play", string(in))
}
