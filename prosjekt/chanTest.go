package main

import(
	//"strings"
	"fmt"
	//"time"
	//"bufio"
	//"os"
	"runtime"
)




func testFunc(ch chan string){
	ch <- "fuck"
	time.Sleep(10*time.Second)
	return
}

func chanTest() chan string{
	ch := make(chan string)
	go testFunc(ch)
	return ch
}

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	text := "hei"
	//chanReader := bufio.NewReader(os.Stdin)
	ch := chanTest()
	//chanReader <- ch
		text = <- ch
	//text, _ := chanReader.ReadString()
	fmt.Println(text)

	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
    	if ipv4 := addr.To4(); ipv4 != nil {
        	fmt.Println("IPv4: ", ipv4)
    }   
}
}