package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool)
	waitGroup := &sync.WaitGroup{}

	waitGroup.Add(1)
	go func(done chan bool, waitGroup *sync.WaitGroup) {

		defer waitGroup.Done()
		n := []int{1, 2, 4, 2, 3, 5, 2, 3, 1, 3}
		for x := range n {
			time.Sleep(time.Duration(n[x]) * time.Second)
			fmt.Println(n[x])
			select {
			case _ = <-done:
				return

			default:
			}
		}
	}(done, waitGroup)

	<-quit
	done <- true

	waitGroup.Wait()
	fmt.Println("Service stopped")
}
