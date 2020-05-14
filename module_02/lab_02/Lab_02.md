# Lab 02: Concurrency Patterns in Golang

## WaitGroups

Hopefully from the last lab we got some idea of how painful it can be managing multiple goroutines simultaneously. 
We can do it using things like time.Sleep and other blocking methods but that quickly becomes overwhelming and painful.
Fortunately there's the **sync** package which helps with that. 

Let's take a look at the following code:

```
// To wait for multiple goroutines to finish, we can
// use a *wait group*.

package main

import (
	"fmt"
	"sync"
	"time"
)

// This is the function we'll run in every goroutine.
// Note that a WaitGroup must be passed to functions by
// pointer.
func worker(id int, wg *sync.WaitGroup) {
	// On return, notify the WaitGroup that we're done.
	defer wg.Done()

	fmt.Printf("Creating Critter number %d\n", id)

	// Sleep to simulate an expensive task.
	time.Sleep(time.Second)
	fmt.Printf("Critter number %d done\n", id)
}

func main() {

	// This WaitGroup is used to wait for all the
	// goroutines launched here to finish.
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup
	// counter for each.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Block until the WaitGroup counter goes back to 0;
	// all the workers notified they're done.
	wg.Wait()
}
```

So what's happening here? Well- let's start with the "main()" section. 
This line `var wg sync.WaitGroup` basically creates a "group" of helpers. To go back to our "guy hanging out in a room" analogy- the "waitgroup" is kind of like a bunch of uniforms that you created. SO...with a bunch of uniforms the next thing we need is...workers! Let's create FIVE of them...which we're going to do here:

```
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
```

So now we have five workers hanging out. We assign each of them in the `worker` function. Where would this be useful? WELL...imagine if you had to go to multiple servers in order to read files or, say, `apt-get install` some stuff.  We could pass each server IP address into the function and deploy workers as necessary! 

NOW...the last section basically says "okay...we need to wait until all of these goroutines are finished"...which is a MUCH more elegant answer than using `time.Sleep` or an awkward blocking operation. Instead we simply use the command `wg.Wait()` and lo and behold...we have our blocking operation.

![blocking](./images/blocking.jpeg)