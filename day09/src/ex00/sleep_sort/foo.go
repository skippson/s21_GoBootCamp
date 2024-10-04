package sleep_sort

import (
	"sync"
	"time"
)

func SleepSort(slice []int) chan int {
	size := len(slice)
	ch := make(chan int, size)

	wg := sync.WaitGroup{}
	wg.Add(size)
	
	for _, val := range slice{
		go func(){
			defer wg.Done()
			time.Sleep(time.Duration(val) * time.Second)
			ch <- val
		}()
	}

	go func(){
		wg.Wait()
		close(ch)
	}()

	return ch
}
