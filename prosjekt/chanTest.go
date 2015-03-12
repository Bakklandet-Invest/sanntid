package main

import(
	"strings"
	"fmt"
	//"time"
	//"bufio"
	//"os"
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
		text = "hei"
	//chanReader := bufio.NewReader(os.Stdin)
	ch := chanTest()
	//chanReader <- ch
		text = <- ch
	//text, _ := chanReader.ReadString()
	fmt.Println(text)
}