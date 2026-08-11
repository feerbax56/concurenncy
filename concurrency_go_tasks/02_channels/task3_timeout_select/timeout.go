package timeout

import (
	"context"
	"errors"
	"time"
)

var ErrTimeout = errors.New("work timeout")   // возвращается если работа заняла больше 100 мс
var ErrCanceled = errors.New("work canceled") // возвращается при отмене контекста

// Work выполняет длительную задачу и возвращает ошибку,
// если она заняла больше 100 мс или контекст был отменён.
func Work(ctx context.Context) error {
	// TODO: реализовать через select и time.After
	ch := make(chan error, 1)
	go func() {
		time.Sleep(150 * time.Millisecond)
		ch <- nil
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(100 * time.Millisecond):
		ch <- ErrTimeout
	case <-ctx.Done():
		ch <- ErrCanceled
	}
	return <-ch
}
