package main

import (
	"log"
	"math/rand"
	"time"
)

// Dengan solusi sebelumnya, tambahkan fitur:
// Goroutine untuk wasit
// Wasit akan mengambil bola ketika salah satu pemain
// tidak dapat mengembalikan bola ke lawan
// Permainan selesai, wasit menentukan pemenang

func main() {
	table := make(chan *ball)
	done := make(chan *ball)

	go player("ryan", table, done)
	go player("kido", table, done)

	referree(table, done)
}

type ball struct {
	hits       int
	lastPlayer string
}

func referree(table chan *ball, done chan *ball) {
	table <- new(ball)

	ball := <-done
	log.Println("Winner is", ball.lastPlayer)
}

func player(name string, table chan *ball, done chan *ball) {
	for {
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)

		select {
		case ball := <-table:
			v := r.Intn(1000)
			if v%11 == 0 {
				log.Println(name, "drop the ball")
				done <- ball
				return
			}
			ball.hits++
			ball.lastPlayer = name
			log.Println(name, "hits the ball", ball.hits)
			time.Sleep(50 * time.Millisecond)
			table <- ball
		case <-time.After(2 * time.Second):
			// to stop goroutine to avoid memory leak
			return
		}

	}
}
