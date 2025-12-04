package main

import (
	"fmt"
	"math/rand"
)

type player interface {
	kickBall()
}

type footballPlayer struct {
	stamina int
	power   int
}

func main() {
	team := make([]player, 11)
	for i := 0; i < len(team); i++ {
		team[i] = footballPlayer{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}
	for i := 0; i < len(team); i++ {
		team[i].kickBall()
	}
}

func (f footballPlayer) kickBall() {
	shot := f.stamina + f.power
	fmt.Printf("I'm kicking the ball with %v shot\n", shot)
}
