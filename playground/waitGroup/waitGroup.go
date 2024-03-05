package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	goWaitGroup()

	goChannel()
}

/*
	Ref: https://stackoverflow.com/questions/65213707/where-to-put-wg-add

Rules for using sync.WaitGroup:
 1. Call WaitGroup.Add() in the "original" goroutine (that starts a new) before the go statement
 2. Recommended to call WaitGroup.Done() deferred, so it gets called even if the goroutine panics
 3. When passing WaitGroup to functions rather than use a package level variable, using pointer to it.
    Else the WaitGroup (which is a struct) would be copied and the Done() method called on the copy wouldn't be observed on the original
 4. Be careful "WaitGroup is reuse before previous Wait has returned", which means calling wg.Add() aging after wg.Done() before wg.Wait()
*/
func goWaitGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doWork(&wg, i+1) // Rule No.3
	}

	wg.Wait()
	fmt.Println("Everything is done!")
}

func doWork(wg *sync.WaitGroup, i int) {
	defer wg.Done() // Rule No.2
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Work %d is Done...\n", i)
}

/*
To be really idiomatic, most "Bang" channels (channels that serve only to send a signal) should have the type chan struct{} instead of
chan bool. Also, channels use sync underneath thus using sync should be more performant. WaitGroup helps when you have to block wait for
many goroutines to return. It's simpler when you might spawn a hundred of them in a for loop.

Now, I wouldn't even say that channels are more idiomatic. Channels being a signature feature of the Go language shouldn't mean that
it is idiomatic to use them whenever possible. What is idiomatic in Go is to use the simplest and easiest to understand solution: here,
the WaitGroup convey both the meaning (your main function is Waiting for workers to be done) and the mechanic (the workers notify when
they are Done).
*/
func goChannel() {
	family := []string{"Linus", "Sabrina", "Shane", "Zoe"}
	done := make(chan string)

	for _, name := range family {
		go func(name string) {
			time.Sleep(1 * time.Second)
			done <- name
		}(name)
	}

	for range family {
		fmt.Println(<-done)
	}
}
