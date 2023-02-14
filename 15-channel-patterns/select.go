package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg = &sync.WaitGroup{}
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	c3 := make(chan string, 1)
	done := make(chan bool, 3)
	count := 0
	wg.Add(3)
	go func() {
		defer wg.Done()

		time.Sleep(3 * time.Second)
		c1 <- "one" // send
		done <- true
		count++
	}()

	go func() {
		defer wg.Done()

		time.Sleep(1 * time.Second)
		c2 <- "two"
		done <- true
		count++
	}()

	go func() {
		defer wg.Done()

		c3 <- "three"
		done <- true
		count++
	}()

	//fmt.Println(<-c1)
	//fmt.Println(<-c2)
	//fmt.Println(<-c3)

	for count != 3 {
		select {
		// whichever case is not blocking exec that first
		//whichever case is ready first exec that.
		case x := <-c1: // recv over the channel
			fmt.Println("send the result in the further pipeline", x)
		case y := <-c2:
			fmt.Println(y)
		case z := <-c3:
			fmt.Println(z)
		case <-done:
			fmt.Println(count)
			if count == 3 {
				//check = !check
			}
		}
	}

	fmt.Println("end of the main")
	wg.Wait()
}
