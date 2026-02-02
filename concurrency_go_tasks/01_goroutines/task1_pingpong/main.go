package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// PingPong должен запускать две горутины "ping" и "pong",
// которые поочередно выводят строки пять раз каждая.
// Реализуйте синхронизацию через каналы и ожидание завершения.
func PingPong(w io.Writer) {
	// TODO: реализовать обмен сообщениями между горутинами
	pingChan := make(chan struct{})
	pongChan := make(chan struct{})
	doneChan := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			<-pingChan

			fmt.Println("ping")

			pongChan <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			<-pongChan
			fmt.Println("pong")

			if i < 4 {
				pingChan <- struct{}{}
			}

			close(doneChan)

		}
	}()
	pingChan <- struct{}{}

	<-doneChan
	wg.Wait()
}

func main() {
	PingPong(os.Stdout)
}
