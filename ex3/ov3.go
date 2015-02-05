package main

import(
	"fmt"
	"net"
	"runtime"
	"bufio"
	"os"
	"strconv"
	"sync"
	"time"
)

const server = "129.241.187.136"
const tcpPortFixed = "34933"
const tcpPortDelimited = "33546"
const tcpPortListening = "30000"

// Fixed TCP reader

func readAndPrintFixed(reader *bufio.Reader){
	str := make([]byte, 1024)

	n, err := reader.Read(str)
	if n != 1024 || err != nil {
		fmt.Fprintln(os.Stderr, "read size: " + strconv.Itoa(n))
		fmt.Fprintln(os.Stderr, "Fixed reader error: " + err.Error())
		return
	}
	fmt.Println(string(str[:n]))
}

// Delimited by \0 TCP reader

func readAndPrintDelim(reader *bufio.Reader){
	str, err := reader.ReadString('\000')
	if err != nil{
		fmt.Fprintln(os.Stderr, "Delim reader error: " + err.Error())
		return
	}
	fmt.Println("recieved: " + str)
}

// TCP connector
func conTCP(port string, callback func(*bufio.Reader)){
	conn, err := net.Dial("tcp", net.JoinHostPort(server, port))
	if err != nil{
		fmt.Fprintln(os.Stderr, "connection error on " + server + ":" + port)
		fmt.Fprintln(os.Stderr, "connection join error: " + err.Error())
		return
	}
	defer conn.Close()

	// Chat with server
	wg.Add(1)
	chatWithServer(conn, callback)

	// Initialize socket for listening
	addr, err := net.ResolveTCPAddr("tcp", ":"+tcpPortListening)
	if err != nil{
		fmt.Fprintln(os.Stderr, "Address resolution error: " + err.Error())
	}
	ln, err := net.ListenTCP("tcp", addr)
	if err != nil{
		fmt.Fprintln(os.Stderr, "Connection init error: " + err.Error())
	}
	ln.SetDeadline(time.Now().Add(2*time.Second))
	defer ln.Close()
	// Ask for reverse connection
	fmt.Fprintf(conn, "Connect to: "+server+":"+tcpPortListening+"\000")
	// Handle incoming connection
	deadline := time.Now().Add(2*time.Second)
	for time.Now().Before(deadline){
		conn2, err := ln.Accept()
		if err != nil{
			fmt.Fprintln(os.Stderr, "Socket acceptance error: " + err.Error())
			continue
		}
		defer conn2.Close()
		wg.Add(1)
		go chatWithServer(conn2, callback)
	}
	wg.Wait();
}

func chatWithServer(conn net.Conn, callback func(*bufio.Reader)){
	reader := bufio.NewReaderSize(conn, 1024)
	// Read welcome package
	callback(reader)
	for i := 0; i < 5; i++{
		fmt.Fprintf(conn, "Hello World: " + strconv.Itoa(i) + "\000")
		callback(reader)
	}	
	wg.Done()
}

var wg sync.WaitGroup

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("--- Fixed TCP communication ---");
	conTCP(tcpPortFixed, readAndPrintFixed)
	fmt.Println()
	fmt.Println("--- 0-delimited TCP communication ---");
	conTCP(tcpPortDelimited, readAndPrintDelim)
}

	


