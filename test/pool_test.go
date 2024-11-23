package test

import (
	"sync"
	"testing"
)

var result string
var strPool = sync.Pool{New: func() any {
	return "hello world"

}}

func goString() string {
	wg := &sync.WaitGroup{}
	var s string

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			var x = "hello world"
			s = x
		}(wg)
	}

	wg.Wait()
	return s
}

func goStringWithPool() string {
	wg := &sync.WaitGroup{}
	var s string

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			// get from pool
			var x = strPool.Get().(string)
			s = x
			strPool.Put(x)
		}(wg)
	}

	wg.Wait()
	return s
}

func BenchmarkGoString(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = goString()
	}
	result = s
}

func BenchmarkGoStringWithPool(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = goStringWithPool()
	}
	result = s
}
