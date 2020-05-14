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

### CHALLENGE TWO

#### Reorder the code so that we show the print statement BEFORE we show the results of the goroutine code

### CHALLENGE THREE

#### Write a different program utilizing a goroutine. Work out sleep times

## GO PATTERNS ON CONCURRENCY

Okay- so obviously if we want to utilize goroutines we'll need to work out a pattern in order to properly utilize these goroutines. 
Let's add another twist in here and see what happens when we we do multiple go routines in a single "main" call (reference is [here](https://www.youtube.com/watch?v=u5k_arVcqR8) ifyou also like terrible movies)

```
package main

import (
	"fmt"
	"time"
)

func badlinereading(line string) {
	for i := 0; i < 3; i++ {
		fmt.Println(line, ":", i)
	}
}

func main() {

	badlinereading("oh God")
	go badlinereading("oh Man")

	go func(msg string) {
		fmt.Println(msg)
	}("Oh NO")

	time.Sleep(time.Second)
	fmt.Println("FIN")
}
```

![toughguys](./images/toughguys.jpg)

Okay- so what is it you all anticipate seeing come out of this? 
What order do you think we'll see the prints come out?
Why do you think it came out like that? 

![why](./images/why.png)

Now let's look at what happens when we create two go routines. To keep it interesting let's see if we can create a scenario where we can create to goroutines that are alternating between each other. Something like this: 

```

package main

import (
	"fmt"
	"time"
)

func badlinereading(line string) {
	for i := 0; i < 3; i++ {
		time.Sleep(101 * time.Millisecond)
		fmt.Println(line, ":", i)
	}
}

func badlinereadingtwo(otherline string) {
	for i := 0; i < 3; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println(otherline, ":", i)
	}
}

func main() {

	go badlinereading("oh God")
	go badlinereadingtwo("oh Man")

	go func(msg string) {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(msg)
	}("Oh NO")

	time.Sleep(time.Second)
	fmt.Println("FIN")
}
```

Do you see how things are mixing up with GOROUTINES? 
Do you see how they are fundamentally **asynchronous**? Where we don't have a TON of control over the order that things happen? 

The best we can do here is fundamentally mess around with the time.Sleep() to try to manipulate this. 
OBVIOUSLY this is not ideal....I mean...what's the point of utilizing these goroutines if we constantly have to slow things down in order to do things? 

### CHALLENGE FOUR

#### Can you get the goroutines set up to alternate using timing?

So here's another way to show that:

```
package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func main() {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}
```

Here's what's going on:

![explanation](./images/explanation.png)

## Channels

So obviously we need a good way to manage these goroutines. 
It's a pain in the backside using SLEEP all the time...so let's instead look at using communications channels. This is in keeping with the philosophy behind golang of **Do not communicate by sharing memory; instead, share memory by communicating.**

Let's start by looking at the way channels are set up. The <- operator specifies the channel direction, send or receive. If no direction is given, the channel is bi-directional.
Basically this is how goroutines communicate with each other. Let's start with a basic example (from go-by-example):

```
// _Channels_ are the pipes that connect concurrent
// goroutines. You can send values into channels from one
// goroutine and receive those values into another
// goroutine.

package main

import "fmt"

func main() {

	// Create a new channel with `make(chan val-type)`.
	// Channels are typed by the values they convey.
	messages := make(chan string)

	// _Send_ a value into a channel using the `channel <-`
	// syntax. Here we send `"ping"`  to the `messages`
	// channel we made above, from a new goroutine.
	go func() { messages <- "ping" }()

	// The `<-channel` syntax _receives_ a value from the
	// channel. Here we'll receive the `"ping"` message
	// we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)
}
```
Okay- so simple, right? We are: 

* Creating a channel (`make` command)
* Using a goroutine we're going to send a message to that command (`ping`)
* A variable (`msg`) is used to receive that data and...
* We print that variable.

So...the big question here is: **why did the goroutine have time to finish?**

As we've seen in previous examples- GOROUTINES are creating a second person, asking them to do something, watching them leave the room, and then continuing to do stuff on the "main" thread. 

IF you finish stuff on the "main" thread then you end up finishing the system, leaving the room, and turning out the lights before the goroutine (fake person) has time to finish doing it's thing. 

BUT...what if you had.... a telephone? AND...what if you COULDN'T leave the room until that phone rang? What if you could say to your fake "goroutine" person "call me when you're done doing your thing. I won't leave the room until you're done".
That's channels. 
Channels are **blocking**...meaning that they won't let the main thread finish doing it's thing until the channel has finished SO...we end up waiting for a call. Let's demonstrate this with a **deadlock** below:

```
package main

import "fmt"

func main() {
  fmt.Println("Okay let's go!")
  c:= make(chan string)
  c <- "Lisa"
  fmt.Println("You're tearing me apaaaht!")
}
```

So what happened there? 
Well- we sent the word "Lisa" to our channel which took that word and..... kind of just hung out. 
It really isn't doing anything WHILST simultaneously blocking the main thread from running.
SO...not doing much but stopping other people from working. 
It's basically like, well...

![theoffice](./images/theoffice.jpg)

So okay fine...so let's allow it to _do_ something.

### CHALLENGE FIVE

#### "FIX" the channels issue so that we no longer have deadlock! (HINT: try passing the channel into the function)

So one way to do this is to simply empty the channel. It doesn't need to be emptied anywhere in particular..you just need to empty is so...

```
package main

import "fmt"

func tearmeapart(c chan string){ 
	<-c
}

func main() {
  fmt.Println("Okay let's go!")
  c:= make(chan string)
  go tearmeapart(c)
  c <- "Lisa"	
  fmt.Println("You're tearing me apaaaht!")
}
```

See? We need a way to get that data in and fortunately a good **goroutine** is just the thing! The goroutines will act as consumers for the channels because, again, channels are how goroutines communicate.
In the above example we're sending the person out of the room carrying the telephone (the channel) and then sending him/her a message ("Lisa"). Since there is someone to answer the phone we can get back to work and continue on! 
If no one answers the phone and you have to sit there and let it ring waiting for someone to answer we get....**deadlock!**

![phoneringing](./images/phoneringing.jpeg)

Please NOTE that if you don't utilize a goroutine (i.e: send someone out of the room) then you have no need of the phone and nothing will happen!

### CHALLENGE SIX

#### Create a function that passes an int and a string through a channel to two separate goroutines. Set them up so they do a call and repeat of "Samurai" followed by "Cop" and repeat it three times! So the output should be "Samurai", "Cop", "Samurai", "Cop", etc...

## Buffered Channels

Another thing we can consider with channels is that we can set the size. This is the equivalent of "call waiting" on our example; basically we can send calls in and have them on "hold" waiting for people to answer them (making channels act like a queue). 
Making buffered channels is relatively simple. To make a buffered channel we simply do: `make(chan {type}, {size of buffer})`

So let's run a quick and very basic example here:

```
package main

import "fmt"

func main() {

    messages := make(chan string, 2)

    messages <- "Troll"
    messages <- "Part 2"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
```

ALSO notice that we do __not__ have to dump our message into a variable to print it.
SO- to quote an article: 

***Buffered channels are useful when you know how many goroutines you have launched, want to limit the number of goroutines you will launch, or want to limit the amount of work that is queued up.***

Another big consideration with buffered channels is that unlike unbuffered channels- buffered channels __do not block until they are full__. Let's do a quick example of that here:

```
package main

import (  
    "fmt"
    "time"
)

func write(ch chan int) {  
    for i := 0; i < 5; i++ {
        ch <- i
        fmt.Println("successfully wrote", i, "to ch")
    }
    close(ch)
}
func main() {  
    ch := make(chan int, 2)
    go write(ch)
    time.Sleep(2 * time.Second)
    for v := range ch {
        fmt.Println("read value", v,"from ch")
        time.Sleep(2 * time.Second)

    }
}
```
I didn't get cute with the movie references here because I want everyone to really understand the pattern that is happening (you can run that in the GO playground). 
**NOW** please go back and remove the "2" from the `make(chan int,2)` line (line 16 on the go playground) and try that again. What happens now? 
Now run this (I turned the loop up to 10 and the buffer up to 4):

```
package main

import (  
    "fmt"
    "time"
)

func write(ch chan int) {  
    for i := 0; i < 10; i++ {
        ch <- i
        fmt.Println("successfully wrote", i, "to ch")
    }
    close(ch)
}
func main() {  
    ch := make(chan int,4)
    go write(ch)
    time.Sleep(2 * time.Second)
    for v := range ch {
        fmt.Println("read value", v,"from ch")
        time.Sleep(2 * time.Second)

    }
}
```

Do you see what is happening with the buffered channels? They BLOCK when you have only a single value channel but then when you have a 4 value buffer they cease blocking while they fill up and then slowly drain out! 

Okay! SO...we've got the first part of the basic understanding done. In the next lab we're going to be addressing one of the biggest challenges around managing these things. 

Try to imagine managing multiple lines of these goroutines simultaneously would be, as you can imagine, a bit of a nightmare. 

![twilight](./images/twighlight.jpeg)

At this point please go to the **challenges** folder and try out some challenges before we move on to concurrency patterns in Lab 2.