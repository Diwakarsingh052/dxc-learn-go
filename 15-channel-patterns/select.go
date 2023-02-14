package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg = &sync.WaitGroup{}

	//wgChan keep track of if the go routine work is finished or not and we wil close the channel when work is done
	var wgChan = &sync.WaitGroup{}
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	c3 := make(chan string, 1)
	done := make(chan bool, 1)

	wgChan.Add(3)
	go func() {
		defer wgChan.Done()

		time.Sleep(3 * time.Second)
		c1 <- "one" // send

	}()

	go func() {
		defer wgChan.Done()

		time.Sleep(1 * time.Second)
		c2 <- "two"

	}()

	go func() {
		defer wgChan.Done()

		c3 <- "three"

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//waiting for go routines to finish
		wgChan.Wait()
		fmt.Println("closing")
		//closing the channel
		close(done)
	}()

	//fmt.Println(<-c1)
	//fmt.Println(<-c2)
	//fmt.Println(<-c3)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			// whichever case is not blocking exec that first
			//whichever case is ready first exec that.
			case x := <-c1: // recv over the channel
				fmt.Println("send the result in the further pipeline", x)
			case y := <-c2:
				fmt.Println(y)
			case z := <-c3:
				fmt.Println(z)
			case <-done: // this case will exec when channel is closed
				fmt.Println("it is closed")
				return
			}
		}
	}()

	fmt.Println("end of the main")
	wg.Wait()
}
