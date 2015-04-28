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
const broadcastIP = "129.241.187.255"

func broadcastLocalIP(){
	bcConn := setupSenderUDP()
	melding := findLocalIP()

	byteMelding := []byte(melding)
	bcConn.Write(byteMelding)

	return
}

func setupSenderUDP() *net.UDPConn {
	addr, _ := net.ResolveUDPAddr("udp", broadcastIP + ":" + writePort)
	sock, _ := net.DialUDP("udp", nil, addr)

	return sock
}


func acceptTCPConn() *net.TCPConn {
	TCPAddr, _ := net.ResolveTCPAddr("tcp", ":" + TCPPort)
	listener, _ := net.ListenTCP("tcp", TCPAddr)
	incConn, err := listener.AcceptTCP()
	if err != nil {
		fmt.Println("error accepting connection")
	}

	return incConn
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
    return "0.0.0.0"
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

	broadcastLocalIP()

	Conn := acceptTCPConn()
	
	var message string = "Initialized"
	sendMsgReader := bufio.NewReader(os.Stdin)
	receiveMsgReader := bufio.NewReader(Conn)
	
	go read(receiveMsgReader)
	time.Sleep(1*time.Millisecond)
	go write(Conn, message, sendMsgReader)
	
	
	
	<- ch

	//Conn.Close()
	
}

