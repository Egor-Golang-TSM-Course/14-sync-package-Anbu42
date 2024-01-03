package main

import (
	"fmt"
	"sync"
)

type LogBuffer struct {
	buffer []string
	mutex  sync.Mutex
}

func (l *LogBuffer) WriteLog(message string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.buffer = append(l.buffer, message)
}

func main() {
	logBuffer := LogBuffer{buffer: make([]string, 0), mutex: sync.Mutex{}}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			for j := 0; j < 3; j++ {
				logBuffer.WriteLog(fmt.Sprintf("Log from Goroutine %d, Message %d", index, j))
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("Logs:")
	for _, log := range logBuffer.buffer {
		fmt.Println(log)
	}
}
