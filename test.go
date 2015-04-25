package main

import (
    ."fmt"
    ."time"
)

var (
    shutdown = make(chan struct{})
    sup2 = make(chan int, 1)
    sup1 = make(chan int, 1)
    done     = make(chan int)
    pingChan = make(chan int) 
)

func henge() {
	for{
		select{
		case <- shutdown:
			done <- 123421
			return
		case <- sup1:
			Println("chillern og henger")
			go timeoutFunc(sup2)
			sup2 <- 123
			sup2 <- 321
			Println("ferdig med å henge")
		}
	}
}

func timeoutFunc(channel chan int) {
	Sleep(5*Second)
	select{
		case <- channel:
			Println("TIMEOUT")
	}
}

func lese() {
	go henge()
	for {
		sup1 <- 123
		
		select{
		case <- shutdown:
			done <- 11111111
			return
		case <- sup2:
			Println("lese leser")
		}
		Sleep(1500*Millisecond)
	}
}

func goroutine(nr int) {
	const n = 5

    // Start up the goroutines...
    for i := 1; i <= n; i++ {
        i := i+nr
        go func() {
            select {
            case <-shutdown:
                done <- i
            }
        }()
    }
}

func timeout(){
	for {
		Sleep(Second)
		<- pingChan
		Println("PING TIMED OUT")
	}
}
	

func Ping() {
	pingChan <- 1
	Println("PINGED")
	return
}


func Pinger() {
	for {
	go Ping()
	Sleep(400*Millisecond)
	}
}

func getPinged() {
	for{
		select{
		case <- pingChan:
			Println("got pinged")
		}
		Sleep(2*Second)
	}
}

func main() {
    
    const n = 5

    // Start up the goroutines...
    for i := 0; i < n; i++ {
        go goroutine(i*5)
    }
    go lese()

    //Sleep(2 * Second)

    // Close the channel. All goroutines will immediately "unblock".
    close(shutdown)

    for i := 0; i < 27; i++ {
        Println("routine", <-done, "has exited!")
    }
    go timeout()
    go Pinger()
    go getPinged()

 	Sleep(20*Second)
}
