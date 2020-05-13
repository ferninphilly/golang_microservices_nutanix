# Lab 01: Concurrency and Parallelism in Golang

## GoRoutines and channels

So as we addressed in the lecture we're going to start by implementing some simple goroutines to demonstrate how they are working. 
Let's start with a very simple example. Open up a new .go file in this directory (answers in the "challengeanswers" section) and create the following function:

```
func badmovie(a string, b string) {
    fmt.Printf("%s is the film title and %s is the rotten tomatoes score!\n",a,b)
}
```

So..pretty basic, standard golang function, right? Excellent.
SO now we need to talk about how we can use this. In and of itself a simple `go badmovie` doesn't do much for us. As a matter of fact...let's run this program in playground and see what happens:

```
package main

import (
	"fmt"
)

func badmovie(a, b string) {
	fmt.Printf("%s is the film title and %s is the rotten tomatoes score!\n",a,b)
}

func main() {
	badmovie("The Room", "25%")
}
```

Okay! Very basic and straightforward, right? 
WELL...now let's start adding some twists. Add the "go" keyword in front of the function call:

```
package main

import (
	"fmt"
)

func badmovie(a, b string) {
	fmt.Printf("%s is the film title and %s is the rotten tomatoes score!",a,b)
  fmt.Println("You are tearing me apaaaht, Lisa!")
}

func main() {
	go badmovie("The Room", "25%")
  fmt.Println("You are tearing me apaaaht, Lisa!")
}
```

UMMM.....what happened? 
Why didn't the function run? 

So now we need to talk about the concepts around **blocking** and **non-blocking** functions that we went over in the lectures.

### CHALLENGE ONE: CREATE A BLOCKING FUNCTION THAT WILL STOP EVERYTHING AND ALLOW OUR GO ROUTINE TO RUN


#### Blocking

Okay- so initially let's start out with a poor man't blocking: let's __sleep__:

```
package main

import (
	"fmt"
	"time"
)

func badmovie(a, b string) {
	fmt.Printf("%s is the film title and %s is the rotten tomatoes score!\n",a,b)
}

func main() {
	go badmovie("The Room", "25%")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("You are tearing me apaaaht, Lisa!")
}

```

So do you see what happened here? 
GOROUTINES are the equivalent of you standing in an empty room and doing a thing. That "thing" is the main thread ...so maybe it's a "PRINT" statement, maybe it's a complex piece of mathematical calculation...whatever. 
The idea is that once you are done this thing you are to pack up and leave the room. 
It's over and you're done at that point.
So here you are standing in this room and doing a thing.

![guyinroom](./images/guyinroom.jpeg)

A **goroutine** is what happens when you have a second person suddenly appear in the room. The person is there to do something at the same time you are working on your complex mathematics equation or whatever. You order them to go off and do a thing with the command `go {nameofFunction}` and they run out of the room.

SO...what happened when they ran out of the room to do another thing? 
BASICALLY...you finished doing your thing, walked out of the room, and shut off the light. 
You're done. It's over. Main thread closed. 

Here's the key question though: __did you give your other person who you ordered out with "go" time to finish doing their thing??__

Hopefully this makes clearer what happened here:
You did your thing:

* Ran your "main" program
* Told your "other guy/gal" to do their thing
* Walked out and closed the door.

![twopeopleinoffice](./images/twopeopleinoffice.jpg)

SO...conceptually let's look at how we can mitigate this a bit. The most obvious way to do this is simple- let's take a nap! Using the `time.Sleep()` method we can do our complex maths/printing/whatever then head to our virtual "couch" and take a nap...thereby giving our "Igor" time to do his/her thing! 

![lazyworker](./images/lazyworker.jpeg)

So let's take a look at what that looks like:

```
package main

import (
	"fmt"
	"time"
)

func badmovie(a, b string) {
	fmt.Printf("%s is the film title and %s is the rotten tomatoes score!\n",a,b)
}

func main() {
	go badmovie("The Room", "25%")
	fmt.Println("You are tearing me apaaaht, Lisa!")
	time.Sleep(500 * time.Millisecond)

}
```

Now... please notice something here:

Look at the order of the code!

Had we been following the normal code order it would have run "badmovie" BEFORE it ran the print statement. But it didn't do that. Why not?

## CHALLENGE TWO: Reorder the code so that we show the print statement BEFORE we show the results of the goroutine code

## CHALLENGE THREE: Write a different program utilizing a goroutine. Work out sleep times. 

### GO PATTERNS ON CONCURRENCY


