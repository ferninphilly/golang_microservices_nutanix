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
	var i int
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	for i = 0; i < 2; i++ {
		select {
		case s1 := <-output1:
			fmt.Println(s1)
		case s2 := <-output2:
			fmt.Println(s2)
		}
	}
}
