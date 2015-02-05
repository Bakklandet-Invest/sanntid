// Go 1.2
// go run helloworld_go.go
package main
import (
	. "fmt" // Using '.' to avoid prefixing functions with their package names
		// This is probably not a good idea for large projects...
	"runtime"
	"time"
)

var c = make(chan int, 1)



func someGoroutine1() {
	for n := 0; n < 1000001; n++ {
		j := <- c
        	i = j + 1
		c <- i
		
    	}
}
func someGoroutine2() {
	for m := 0; m < 999; m++ {
		j := <- c
        	i = j - 1
		c <- i
	
    	}
}

var i = 0

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
					// Try doing the exercise both with and without it!

	c <- i	
	
	go someGoroutine1() // This spawns someGoroutine() as a goroutine
	
	go someGoroutine2()	// We have no way to wait for the completion of a goroutine (without 				additional syncronization of some sort)
	// g := <- c			
			// We'll come back to using channels in Exercise 2. For now: Sleep.
	time.Sleep(1000*time.Millisecond)
	Println(i)
	



}
