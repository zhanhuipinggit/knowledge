package main

import (
	"fmt"
	"sync"
)

func main() {
	dogChan := make(chan struct{})
	catChan := make(chan struct{})
	fishChan := make(chan struct{})
	done := make(chan struct{})

	wait := sync.WaitGroup{}
	wait.Add(3)
	go func() {
		wait.Done()
		for i := 0; i < 100; i++ {
			<-dogChan
			fmt.Println("dog")
			catChan <- struct{}{}
		}
	}()

	go func() {
		wait.Done()
		for i := 0; i < 100; i++ {
			<-catChan
			fmt.Println("cat")
			fishChan <- struct{}{}
		}
	}()

	go func() {
		wait.Done()
		for i := 0; i < 100; i++ {
			<-fishChan
			fmt.Println("fish")
			if i == 99 {
				close(done)
				return
			}
			dogChan <- struct{}{}
		}
	}()
	dogChan <- struct{}{}
	<-done
	wait.Wait()

}
