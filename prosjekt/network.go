package main

import (
	"fmt"
	"net"
	"strconv"
	"bufio"
	"os"
	"runtime"
)

const string selfIP = "129.241.187.144"
const string targIP = "129.241.187.136"

const string sendPort = "20013"
const string recievePort = "30000"
const string fixSizePort = "34933"
const string delimTermPort = "33546"

func connect(targetIP string, targetPort string) {
	TCPAddr, _ := net.ResolveTCPAddr("tcp", targetIP + ":" + targetPort)	
	conn, err := net.DialTCP("tcp", nil, TCPAddr)
	//defer conn.Close()
    if err!= nil {
    	fmt.Fprintln(os.Stderr, "Tried to connect to: " + targetIP + ":" + targetPort)
    	fmt.Fprintln(os.Stderr, " Connection error: " + err.Error())
       	return
    }
    //reader := bufio.NewReaderSize(conn, 1024)	
	return *TCPConn
}


func disconnect(connection *TCPConn) {
	connection.Close()
}


func write(connection *TCPConn, msg string){
	connection.Write([]byte(msg + "\0"))
	return
}


func read(reader *bufio.Reader){
	str, err := reader.ReadString('\0')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Delim error: " + err.Error())
		return
	}
	fmt.Println("Melding: " + str)
	return
}


func main() {
	var msg string = "Initialized"
	runtime.GOMAXPROCS(runtime.NumCPU())
	writeConn := connect(targIP, writePort)
	readConn := connect(targIP, readPort)

	while(string.ToLower(msg) != "disconnect"){
		fmt.Println("Say: ")
	}




}



/*
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
*/