package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Request struct {
	ID       int
	Payload  string
	RespChan chan<- Response
}

type Response struct {
	RequestID int
	Result    string
}

func Handler(requestCh <-chan Request, wg *sync.WaitGroup) {
	defer wg.Done()
	for req := range requestCh {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		result := fmt.Sprintf("Processed: %s", req.Payload)

		req.RespChan <- Response{RequestID: req.ID, Result: result}
	}
}

func main() {
	requestCh := make(chan Request, 5)

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Handler(requestCh, &wg)
	}

	for i := 1; i <= 5; i++ {
		payload := fmt.Sprintf("Request #%d", i)
		respChan := make(chan Response, 1)

		request := Request{
			ID:       i,
			Payload:  payload,
			RespChan: respChan,
		}

		requestCh <- request

		select {
		case response := <-respChan:
			fmt.Printf("Received response for request #%d: %s\n", response.RequestID, response.Result)
		case <-time.After(500 * time.Millisecond):
			fmt.Printf("Timeout for request #%d\n", i)
		}
	}

	close(requestCh)

	wg.Wait()

}
