package main

import (
	"fmt"
	"time"
)

// Definindo o ID  e o data que vai receber os dados através de canal
func worker(workerID int, data chan int) {

	for x:= range data {

		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)

	}

}

/*
func teste (ping chan string) {

	for x := range ping {
	fmt.Printf("%s\n", x)
	time.Sleep(time.Duration(2) * time.Second)
	}

}
*/
//T1
func main() {

	channel := make(chan int)
	// Toda vez que usar "go" antes da função, estamos gerando uma nova thread

	for i := 0; i < 1000; i++ { //Cada thread ocupa apenas 2K de espaço
	go worker(i, channel)
	 } //T2
	go worker(2, channel) //T3
	go worker(3, channel) //T4

	for i := 0; i < 10000; i++ {

		channel <- i

	}

/*
channel := make(chan string)

go teste(channel)

for i := 0;;i++ {
	if i % 2 == 0 {
	channel <- "Ping"
	} else {
		channel <- "Pong"
	}
}
*/

}