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

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			<-pingChan

			fmt.Fprintln(w, "ping")

			pongChan <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		pingChan <- struct{}{}
		for i := 0; i < 5; i++ {
			<-pongChan
			fmt.Fprintln(w, "pong")

			if i < 4 {
				pingChan <- struct{}{}
			}

		}
	}()
	wg.Wait()
	close(pingChan)
	close(pongChan)
}

func main() {
	PingPong(os.Stdout)
}
