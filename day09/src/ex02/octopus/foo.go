package octopus

import (
	"sync"
)

func Multiplaxe(any ...<-chan any) <-chan any {
	wg := sync.WaitGroup{}
	wg.Add(len(any))

	out := make(chan interface{})
	for _, c := range any {
		go func() {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
