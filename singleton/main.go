package main

import (
	"git_test/singleton/singleton"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			_ = singleton.GetSingleInstance()
		}(wg)
	}

	wg.Wait()
}
