package main

import (
	"fmt"
	"net"
	"strconv"
	"bufio"
	"os"
	"runtime"
	"time"
	)

const compIP = "129.241.187.144"
const sendPort = "20012"
const broadcastIP = "129.241.187.255"
const serverIP = "129.241.187.136"
const recievePort = "30000"
const fixSizePort = "34933"
const delimTermPort = "33546"
//const serverIP := 


//addr = new InternetAddress(serverIP, serverPort)
//sock = new Socket(tcp)
//sock.connect(addr)


func sendUDP() {
	fmt.Println("start send\n")
	addr, err := net.ResolveUDPAddr("udp", compIP + ":" + sendPort)
	if err != nil {
    	panic(err)
	}
	fmt.Println("resolve\n")
	sock, err := net.DialUDP("udp", nil, addr)
	if err != nil {
    	panic(err)
	}	
	fmt.Println("dial\n")
	melding := "sup"
	byteMelding := []byte(melding)
	for {
	sock.Write(byteMelding)
	}	
	fmt.Println("write\n")
	return
	
}

func listenUDP() {
	str := make([]byte, 1024)
	addr, err := net.ResolveUDPAddr("udp", ":" + recievePort)
	if err != nil {
    	panic(err)
	}
	fmt.Println("resolve 2")
	sock, err := net.ListenUDP("udp", addr)
	if err != nil {
    	panic(err)
	}
	fmt.Println("listen")
	//err = sock.SetReadDeadline(Now().Add(Second*2))
	if err != nil {
    	panic(err)
	}
	_, senderAddr, err := sock.ReadFromUDP(str[0:])
	if err != nil {
    	panic(err)
	}	
	fmt.Println("read")	
	fmt.Println("message: ")
	fmt.Println(string(str))
	fmt.Println("from addr: ")
	fmt.Println(senderAddr)

	

	/*
	if err != nil {
		fmt.println("Error: " + err.Error())
		return
	}
	*/
	return
}

func connTCP(port string, readFunc func(*bufio.Reader)) {
	TCPAddr, _ := net.ResolveTCPAddr("tcp", serverIP + ":" + port)	
	conn, err := net.DialTCP("tcp", nil, TCPAddr)
	//defer conn.Close()
    if err!= nil {
    	fmt.Fprintln(os.Stderr, "Tried to connect to: " + serverIP + ":" + port)
    	fmt.Fprintln(os.Stderr, " Connection error: " + err.Error())
       	return
    }
    reader := bufio.NewReaderSize(conn, 1024)
	
	readFunc(reader)

	conn.Write([]byte("HEI SERVER\x00"))

	readFunc(reader)

	return
}

func readFixSizeTCP(reader *bufio.Reader){
	str := make([]byte, 1024)
	n, err := reader.Read(str)    // returnerer lengden på beskjeden (n)
	if n != 1024 || err != nil {
		fmt.Fprintln(os.Stderr, "Message size: " + strconv.Itoa(n))
		fmt.Fprintln(os.Stderr, "Fixed size error: " + err.Error())
		return
	}
	fmt.Println("Melding: " + string(str))
	return
}

func readDelimTermTCP(reader *bufio.Reader){
	str, err := reader.ReadString('\000')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Delim error: " + err.Error())
		return
	}
	fmt.Println("Melding: " + str)
	
	return
}

/*
func writeDelimTermTCP(connection *TCPConn, msg string){
	connection.Write([]byte(msg + "\x00"))
}

func writeFixSizeTCP(connection *TCPConn, msg string){
	connection.Write([]byte(msg))
}
*/

func listenTCP() {


	return
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	go sendUDP()
	
	go listenUDP()	
	
	time.Sleep(10000*time.Millisecond)
	return
}
