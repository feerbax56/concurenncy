package fibonacci

// Fib возвращает канал, из которого можно читать первые n чисел Фибоначчи.
func Fib(n int) <-chan int {
	ch := make(chan int)

	go func(n int) {
		defer close(ch)
		if n <= 0 {
			close(ch)
			return
		}
		ch <- 0
		if n == 1 {
			close(ch)
			return
		}
		ch <- 1
		if n == 2 {
			close(ch)
			return
		}
		a, b := 0, 1

		for i := 2; i < n; i++ {
			a, b = b, a+b
			ch <- b
		}
	}(n)

	// TODO: отправить последовательность Фибоначчи в канал
	return ch
}
