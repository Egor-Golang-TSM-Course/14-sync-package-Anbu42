package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

type WebCounter struct {
	visits sync.Map
}

func (w *WebCounter) Increment(url string) {
	key, _ := w.visits.LoadOrStore(url, int32(0))
	counter := key.(int32)

	w.visits.Swap(url, atomic.AddInt32(&counter, 1))
}

func (w *WebCounter) GetVisits(url string) int {
	key, _ := w.visits.Load(url)
	if value, ok := key.(int32); ok {
		return int(atomic.LoadInt32(&value))
	}
	return 0
}

func (w *WebCounter) Print() {
	fmt.Println("Total Visits:")
	w.visits.Range(func(key, value any) bool {
		url := key.(string)
		visits := w.GetVisits(url)
		fmt.Printf("%s: %d visits\n", url, visits)
		return true
	})
}

func main() {
	webCounter := new(WebCounter)

	numGoroutines := 10
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(index int) {
			defer wg.Done()

			url := fmt.Sprintf("http://example.com/page%d", rand.Intn(3))
			webCounter.Increment(url)
		}(i)
	}

	wg.Wait()

	webCounter.Print()
}
