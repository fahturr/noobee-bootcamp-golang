package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name string
	Hit  int
}

const BreakPoint int = 11

func play(player Player, ball, result chan int, done chan bool, turn chan Player) {
	for {
		b := <-ball

		b++
		player.Hit++
		power := rand.Intn(100-1) + 1

		fmt.Printf("(Bola %02d) Pukulan-%02d %s (Power %d) \n", b, player.Hit, player.Name, power)
		time.Sleep(650 * time.Millisecond)

		if power%BreakPoint == 0 {
			done <- true
			turn <- player
			result <- b
			break
		}

		ball <- b
	}
}

func finish(player chan Player, done chan bool, result chan int) {
	if <-done {
		p := <-player
		r := <-result

		fmt.Println("\n======================")
		fmt.Printf("%s Kalah pada Pukulan ke-%02d & Bola ke-%02d \n", p.Name, p.Hit, r)
		fmt.Println("======================")

		return
	}
}

func main() {
	lastTurn := make(chan Player)
	ball := make(chan int)
	done := make(chan bool)
	result := make(chan int)

	players := listPlayer()

	for _, p := range players {
		go play(p, ball, result, done, lastTurn)
	}

	ball <- 0

	finish(lastTurn, done, result)
}

func listPlayer() []Player {
	return []Player{
		{
			Name: "Harrryyy",
			Hit:  0,
		},
		{
			Name: "Maggurie",
			Hit:  0,
		},
	}
}
