# CHALLENGES FOR LAB ONE

## Goroutines and channels

1. So we've been over goroutines and channels now. Let's try to create some challenges. We're going to create a "relay race" using channels. The idea is that we will have a "runner" function that will pass a "baton", announcing each runner by their number.
SO- we'll need a "runner" function that receives the channel which will pass the runner number. 
The "main thread" will be sleeping and waiting for the next runner.

2. Create two functions- `square` and `cube` that run as goroutines. BOTH should receive a channel of type `int` as an argument. The "main" thread will pass an integer into that channel and the goroutines will return the square and cube of the number. We will then add these two together in the main thread. 

3. Two challenges here: 
  * Run the below code. One statement will print. I want you to set this up so it presents BOTH statements in order.
  * Now set it up so that it does NOT deadlock at the end once it's done with both statements (HINT: look at for in range)

```
package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "Steven"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Seagal"

}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

```