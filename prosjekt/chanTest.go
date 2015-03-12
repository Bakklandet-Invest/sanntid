package main

import(
	//"strings"
	"fmt"
	//"time"
)


func testFunc(ch chan string){
	ch <- "fuck"
	return
}

func chanTest() chan string{
	ch := make(chan string)
	testFunc(ch)
	return ch
}

func main(){
	ch := chanTest()
	text := make([]byte,1024)
	text <- ch
	fmt.Println(text)
}