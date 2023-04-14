package channels

import (
	"context"
	"sync"

	"github.com/sirupsen/logrus"
)

// Collects data from list of channels and returns one channel of data
//   - done - signal channel to stop processing
//   - channels - array or enumeration of channels to collect
func Collect[T any](ctx context.Context, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup

	output := make(chan T)

	// func to collect data from given channel to output channel
	collect := func(ch <-chan T) {
		defer wg.Done()

		for item := range ch {
			select {
			case <-ctx.Done():
				return
			case output <- item:
			}
		}
	}

	// collecting data from channels to output
	wg.Add(len(channels))
	for _, ch := range channels {
		go collect(ch)
	}

	// Wait for collection to complete
	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

// Takes channel of items, run given function for each item from channel and returns
func Map[T any, R any](ctx context.Context, input <-chan T, fn func(context.Context, T) <-chan R) []<-chan R {
	outputs := make([]<-chan R, 0)

	for item := range input {
		select {
		case <-ctx.Done():
			logrus.Debugf("channels.Map(): cancelled %s", ctx.Err().Error())
			return outputs
		default:
			outputs = append(outputs, fn(ctx, item))
		}
	}

	return outputs
}
