package network

import (
	"net"
	//"strings"
	//"time"
	"strconv"
	"fmt"
	//"os"
)

var broadcastAddr *net.UDPAddr //Private?

type messageUDP struct{ //Private?
	recieveAddr string
	data []byte
	length int // length of recieved data in bytes
}

// Bare for testing
func PrintUDPMessage(msg  messageUDP){ //Private?
	fmt.Printf("msg messageUDP: \n recieveAddr = %s \n data = %s \n length = %v \n", msg.recieveAddr, msg.data, msg.length)
}

func InitUDP(localListenPort, broadcastPort, messageSize int, sendCh, recieveCh chan messageUDP)(err error){
	fmt.Println("InitUDP running")
	
	broadcastAddr, err = net.ResolveUDPAddr("udp","255.255.255.255:"+strconv.Itoa(broadcastPort))
	// ? Dette error-opplegget eller panic?
	if err != nil{
		return err
	}
	
	// Find localaddress and sets up with choosen port
	tempConn, err := net.DialUDP("udp", nil, broadcastAddr)
	defer tempConn.Close()
	tempAddr := tempConn.LocalAddr()
	LocalAddress, err = net.ResolveUDPAddr("udp", tempAddr.String())
	LocalAddress.Port = localListenPort
	

	localListenConn, err := net.ListenUDP("udp", LocalAddress)
	if err != nil {
		return err
	}
		
	broadcastListenConn, err := net.ListenUDP("udp", broadcastAddr)
	if err != nil {
		localListenConn.Close()
		return err
	}
	
	go recieveServerUDP(localListenConn, broadcastListenConn, messageSize, recieveCh)
	go transmitServerUDP(localListenConn, broadcastListenConn, sendCh)

	return err
}

func recieveServerUDP(lConn, bConn *net.UDPConn, messageSize int, receiveCh chan messageUDP) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error in recieveServerUDP: %s \n Closing connection.", r)
			lConn.Close()
			bConn.Close()
		}
	}()

	bConnRecieveCh := make(chan messageUDP)
	lConnRecieveCh := make(chan messageUDP)

	go connectionReaderUDP(lConn, messageSize, lConnRecieveCh)
	go connectionReaderUDP(bConn, messageSize, bConnRecieveCh)

	for {
		select {
		case buffer := <-bConnRecieveCh:
			receiveCh <- buffer

		case buffer := <-lConnRecieveCh:
			receiveCh <- buffer
		}
	}
}

func transmitServerUDP(lConn, bConn *net.UDPConn, sendCh chan messageUDP) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error in transmitServerUDP: %s \n Closing connection.", r)
			lConn.Close()
			bConn.Close()
		}
	}()

	var err error
	var n int

	for {
		msg := <-sendCh // Venter på ny melding å sende
		
		if msg.recieveAddr == "broadcast" {
			n, err = lConn.WriteToUDP(msg.data, broadcastAddr)
		} else {
			recieveAddr, err := net.ResolveUDPAddr("udp", msg.recieveAddr)
			if err != nil {
				fmt.Printf("Error in transmitServerUDP: ResolveUDPAddr failed\n")
				panic(err)
			}
			n, err = lConn.WriteToUDP(msg.data, recieveAddr)
		}
		if err != nil || n < 0 {
			fmt.Printf("Error in transmitServerUDP \n")
			panic(err)
		}
		fmt.Printf("transmitServerUDP sent %s to %s\n", msg.Data, msg.recieveAddr)
	}
}



func connectionReaderUDP(conn *net.UDPConn, messageSize int, rCh chan messageUDP) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error in connectionReaderUDP: %s \n Closing connection.", r)
			conn.Close()
		}
	}()

	buffer := make([]byte, messageSize)

	for {
		//		fmt.Printf("connectionReaderUDP: Waiting on data from UDPConn\n")
		n, recieveAddr, err := conn.ReadFromUDP(buffer) // n number of bytes copied to buffer
		//		fmt.Printf("connectionReaderUDP: Received %s from %s \n", string(buf), raddr.String())
		if err != nil || n < 0 {
			panic(err)
		}
		rCh <- messageUDP{recieveAddr: recieveAddr.String(), data: buffer, length: n}
	}
}


/*//-------------------SKROT-----------------------------------------
func main(){
	
	
	
	//ifaces, _ := net.Interfaces()
	//// handle err
	//fmt.Println("ifaces:",ifaces)
	//for _, i := range ifaces {
		//addrs, _ := i.Addrs()
		//fmt.Println("addrs:",addrs)
		//// handle err
		//for _, addr := range addrs {
			//switch v := addr.(type) {
			//case *net.IPAddr:
				//// process IP address
				//fmt.Println("Min IP:",addr)
			//}
		//}

	//}

	//name, err := os.Hostname()
	//if err != nil {
		 //fmt.Printf("Oops: %v\n", err)
	//}
	//fmt.Println(name)
	
	addr, err := net.ResolveUDPAddr("udp", "199.225.136.255:27346")
	fmt.Println("LocalAddr","\n",addr)
	if err != nil {
    	panic(err)
	}
	fmt.Println("OK")
	//minAddr, err := localAddr()	
	//fmt.Println(localAddr())

}

/*

/*

func listenUDP() {
	buf := make([]byte, 1024)
	addr, err := ResolveUDPAddr("udp", ":27346")
	if err != nil {
    	panic(err)
	}
	sock, err := ListenUDP("udp", addr)
	if err != nil {
    	panic(err)
	}
	var msg Message
	for {
		length, senderAddr, err := sock.ReadFromUDP(buf)
		if err != nil {
    		panic(err)
		}	
		json.Unmarshal(buf[:length], &msg)
		var _ = senderAddr
		// send msg videre
	} 
}
	
	
func sendUDP(msg Message) {

	addr, err := ResolveUDPAddr("udp", "129.241.187.255:27346")
	if err != nil {
    	panic(err)
	}
	sock, err := DialUDP("udp", nil, addr)
	if err != nil {
    	panic(err)
	}	

	buf, err := json.Marshal(msg)
	if err != nil {
    	panic(err)
	}

	_, err = sock.Write(buf)
	if err != nil {
    	panic(err)
	}
}


*/
