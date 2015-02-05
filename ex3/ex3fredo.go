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
const sendPort = "20013"
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
	addr, _ := net.ResolveUDPAddr("udp", broadcastIP + ":" + sendPort)
	sock, _ := net.DialUDP("udp", nil, addr)
	melding := "sup"
	byteMelding := []byte(melding)
	sock.Write(byteMelding)
	return
	
}

func listenUDP() {
	str := make([]byte, 1024)
	addr, _ := net.ResolveUDPAddr("udp", ":" + recievePort)
	sock, _ := net.ListenUDP("udp", addr)
	
	_, senderAddr, _ := sock.ReadFromUDP(str[:])
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

func listenTCP() {


	return
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go connTCP(fixSizePort, readFixSizeTCP)	
	
	go connTCP(delimTermPort, readDelimTermTCP)	
	time.Sleep(1000*time.Millisecond)
	return
}
