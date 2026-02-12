package pipelinectx

import "context"

// Run строит конвейер из двух стадий: удвоение и суммирование.
// Конвейер должен останавливаться, если ctx отменён.
// Возвращает итоговую сумму и ошибку контекста при отмене.
func Run(ctx context.Context, nums []int) (int, error) {
	// TODO: реализовать конвейер с остановкой по ctx
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}
	input := make(chan int)
	double := make(chan int)

	done := make(chan struct{})

	go func() {
		defer close(double)
		for {
			select {
			case <-ctx.Done():
				for range input {
				}
				return
			case val, ok := <-input:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				case double <- val * 2:

				}
			}
		}
	}()
	var sum int
	go func() {
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-double:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				default:
					sum += val
				}
			}
		}
	}()

	for _, num := range nums {
		select {
		case <-ctx.Done():
			close(input)
			<-done
			return sum, ctx.Err()
		case input <- num:
		}
	}
	close(input)

	<-done

	select {
	case <-ctx.Done():
		return sum, ctx.Err()

	default:
		return sum, nil
	}
}
