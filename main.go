package main

//NOTE: this program will run indefinitely since it is meant to be always listening on the channel

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

//each alarm will have its own time that it goes off, plus an associated sound file. The alarm will
//most likely just send a filepath to the main function in the form of a string, so that the main
//function can do the heavy lifting in playing the file. Since this program isn't about testing the
//alarm time function nor the music playing, our main program will simply print the message
//associated with each alarm

type alarm struct {
	name    string
	message string
	running bool
}

//while the alarm "running" bool is set to true, this method sends on channel c
func (a *alarm) run(c chan string, wg *sync.WaitGroup) {
	for {
		if a.running != false {
			rand.Seed(time.Now().UnixNano())
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			c <- a.message
		} else {
			msg := fmt.Sprintf("%v is off", a.name)
			c <- msg
			wg.Done()
			return
		}
	}
}

func listen(c, c2 chan string) {
	for {
		val, ok := <-c
		if ok != true {
			c2 <- "All alarms are off"
			return
		}
		fmt.Println(val)
	}
}

//overly-elaborate mechanism for shutting of alarms while they're running
func shutoff(alarms ...*alarm) {
	for _, a := range alarms {
		time.Sleep(time.Duration(5) * time.Second)
		a.running = false
	}
}

func main() {
	c := make(chan string)
	c2 := make(chan string)

	a1 := &alarm{"alarm1", "This is the first alarm", true}
	a2 := &alarm{"alarm2", "This is the second alarm", true}
	a3 := &alarm{"alarm3", "This is the third alarm", true}

	wg.Add(3)

	go a1.run(c, wg)
	go a2.run(c, wg)
	go a3.run(c, wg)

	go listen(c, c2)

	go shutoff(a1, a2, a3)

	fmt.Println(<-c2)

	wg.Wait()
}
