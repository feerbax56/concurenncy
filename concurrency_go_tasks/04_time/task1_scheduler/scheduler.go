package scheduler

import (
	"sync"
	"time"
)

// Every запускает f каждые d и возвращает функцию для остановки.
func Every(d time.Duration, f func()) (stop func()) {
	// TODO: периодический вызов функции с возможностью остановки
	var stoped bool
	var mu sync.Mutex

	stopCh := make(chan struct{})

	go func() {
		ticker := time.NewTicker(d)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				mu.Lock()
				if stoped {
					mu.Unlock()
					return
				}
				mu.Unlock()
				f()
			case <-stopCh:
				mu.Lock()
				stoped = true
				mu.Unlock()
				return
			}
		}
	}()
	return func() {
		mu.Lock()
		defer mu.Unlock()
		if !stoped {
			stoped = true
			close(stopCh)
		}
	}
}
