package main

import (
	"log"
	"time"
)

// Ping pong sederhana
// Buat sebuah aplikasi sederhana dengan kriteria:
// 2 pemain ping pong saling mengembalikan bola
// Permainan berhenti setelah 1 detik
// Gunakan konkurensi dengan 1 channel

func main() {
	table := make(chan *ball)

	go player("ryan", table)
	go player("kido", table)
	go player("irfan", table)
	go player("aldi", table)

	table <- new(ball)
	time.Sleep(1 * time.Second)
	<-table
}

type ball struct {
	hits int
}

func player(name string, table chan *ball) {
	for {
		ball := <-table
		ball.hits++
		log.Println(name, "hits the ball ", ball.hits)
		time.Sleep(50 * time.Millisecond)
		table <- ball
	}
}
