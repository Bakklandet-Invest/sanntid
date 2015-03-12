package main

import(
	//"strings"
	"fmt"
	"time"
	//"bufio"
	//"os"
	"runtime"
	"net"
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

func findLocalIP() string {
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return
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

	
}