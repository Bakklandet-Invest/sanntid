package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"runtime"
	"strings"
	"time"
)


const writePort = "20006"
const TCPPort = "20123"

func listenForIP() string {
	str := make([]byte, 1024)
	listenConn := SetupListenUDP()
	
	time.Sleep(200*time.Millisecond)
	_, _, err := listenConn.ReadFromUDP(str[:])

	if err != nil {
		fmt.Println("ReadFromUDP error")
	}

	return string(str)[:15]
}

func connect(targetAddr string) *net.TCPConn {
	TCPAddr, _ := net.ResolveTCPAddr("tcp", targetAddr + ":" + TCPPort)	
	conn, err := net.DialTCP("tcp", nil, TCPAddr)
	//defer conn.Close()
    if err!= nil {
    	fmt.Fprintln(os.Stderr, "Tried to connect to: " + targetAddr)
    	fmt.Fprintln(os.Stderr, " Connection error: " + err.Error())
       	return conn
    }
	return conn
}


func SetupListenUDP() *net.UDPConn {
	addr, _ := net.ResolveUDPAddr("udp", ":" + writePort)
	sock, _ := net.ListenUDP("udp", addr)
	return sock
}






func write(connection *net.TCPConn, msg string, reader *bufio.Reader){
	for {
    	fmt.Print("Enter text: ")
    	msg, _ = reader.ReadString('\n')

    	connection.Write([]byte(msg[:len(msg)-1] + "\000"))
		
		time.Sleep(1*time.Millisecond)
		
		if strings.ToLower(msg) == "disconnect\n" {break}	
	}
	return
}

func read(reader *bufio.Reader){
	for {
		str, err := reader.ReadString('\000')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Delim error: " + err.Error())
			return
		}
		fmt.Println("Melding: " + str)
		if strings.ToLower(str) == "disconnect\000" {break}
	}
	return
}

func main() {
	
	ch := make(chan string)
	
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	conIP := listenForIP()
	
	Conn := connect(conIP)
	
	var message string = "Initialized"
	sendMsgReader := bufio.NewReader(os.Stdin)
	receiveMsgReader := bufio.NewReader(Conn)
	
	go read(receiveMsgReader)
	time.Sleep(1*time.Millisecond)
	go write(Conn, message, sendMsgReader)
	
	
	<- ch

	//Conn.Close()
	
}
