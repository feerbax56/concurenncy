package pipeline

import "sync"

// Run строит конвейер из трёх стадий: квадрат, умножение на 2 и суммирование.
func Run(nums []int) int {
	// TODO: реализовать конвейер обработки чисел
	kvadratCh := make(chan int)
	multipleCh := make(chan int)
	sumCh := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(kvadratCh)
		for _, n := range nums {
			kvadratCh <- n * n
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(multipleCh)

		for n := range kvadratCh {
			multipleCh <- n * 2
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(sumCh)

		sum := 0
		for num := range multipleCh {
			sum += num
		}

		sumCh <- sum
	}()
	return <-sumCh
}
