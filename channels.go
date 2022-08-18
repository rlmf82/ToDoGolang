package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main2() {
	channel1 := make(chan int, 2)
	channel2 := make(chan struct{})

	go func() {
		channel1 <- 1
		channel2 <- struct{}{}
	}()

	go Select(channel1, channel2)
	select {}
}

func Select(channel1 chan int, channel2 chan struct{}) {
	for {
		time.Sleep(time.Second)

		select {
		case <-channel1:
			{
				fmt.Println("received")
				fmt.Println(<-channel1)
			}
		case <-channel2:
			{
				fmt.Println("exit")
				os.Exit(0)
			}
		}
	}

}

func makeUnbufferedChannels() {
	channel := make(chan int)

	go func() {
		channel <- 1
	}()

	go func() {
		fmt.Println(<-channel)
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("end")
}

func waitingGroups() {
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func() {
		fmt.Println("func 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("func 2")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("End")
}
