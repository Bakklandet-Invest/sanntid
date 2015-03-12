package main

import(
	//"strings"
	"fmt"
	//"time"
	"bufio"
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
	chanReader := bufio.NewReader(os.Stdin)
	ch := chanTest()
	text := make([]byte,1024)
	chanReader <- ch
	text, _ := chanReader.ReadString()
	fmt.Println(text)
}