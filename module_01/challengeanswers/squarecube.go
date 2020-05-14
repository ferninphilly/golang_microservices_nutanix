package main

import "fmt"

func square(c chan int) {
	fmt.Println("Reading [square]")
	num := <-c
	c <- num * num
}

func cube(c chan int) {
	fmt.Println("Reading [cube]")
	num := <-c
	c <- num * num * num
}

func main() {
	fmt.Println("Let's get ready to square and cube!!")

	squareChan := make(chan int)
	cubeChan := make(chan int)

	go square(squareChan)
	go cube(cubeChan)

	testNum := 3
	fmt.Println("Here's the number I'm sending to the channels! ", testNum)

	squareChan <- testNum

	fmt.Println("Okay- resuming main thread")
	fmt.Printf("Sent %d to [square] channel", testNum)

	cubeChan <- testNum

	fmt.Println("Resuming Main thread")
	fmt.Printf("Sent %d to [cube] channel", testNum)

	fmt.Println("Now [main] is going to read from the channel!")

	squareVal, cubeVal := <-squareChan, <-cubeChan
	sum := squareVal + cubeVal

	fmt.Println("Here is the answer! ", sum)
	fmt.Println("Okay- [main] thread is done!")

}
