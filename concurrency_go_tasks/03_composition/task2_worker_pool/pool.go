package pool

import (
	"sync"
)

// RunPool обрабатывает задачи параллельно в заданном количестве воркеров
// и возвращает сумму результатов.
func RunPool(jobs []int, workers int) int {

	if workers <= 0 {
		workers = 1
	}
	// TODO: реализовать пул воркеров и сбор результатов
	jobsChan := make(chan int, len(jobs))
	resultChan := make(chan int, len(jobs))

	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for job := range jobsChan {
				resultChan <- job
			}
		}()
	}

	for _, job := range jobs {
		jobsChan <- job
	}
	close(jobsChan)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	sum := 0
	for result := range resultChan {
		sum += result
	}
	return sum

}
