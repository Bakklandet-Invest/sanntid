package main

import(
	"network"
	"time"
	"fmt"
	"strconv"
)

var liftsOnline = make(map[string]network.ConnectionUDP)
var disconnElevChan = make(chan network.ConnectionUDP)

func writemap() {
	for {
		fmt.Println("Size av liftsOnline:",len(liftsOnline))
		fmt.Println("liftsOnline:",liftsOnline)
		fmt.Println("Master har ID:",selectMaster(liftsOnline))
		time.Sleep(time.Second*10)
	}
}

func selectMaster(lifts map[string]network.ConnectionUDP)string {
	master := "-1"
	for key := range lifts{
		m, _ := strconv.Atoi(master)
		k, _ := strconv.Atoi(key)
		if k > m{
			master = key
		}
	}
	return master
}

//func masterAlive()bool{
//}

func main(){
	network.Init()
	go writemap()

	for{
		select{
		case msg := <-network.RecieveChan:
			messageHandler(network.ParseMessage(msg))
		case conn := <- disconnElevChan:
			deleteLift(conn.Addr)
		}
			
			
	}
}



func messageHandler(msg network.Message) {
	id := network.FindID(msg.Addr)
	switch msg.Content{
		case network.Alive:
			if conn, alive := liftsOnline[id]; alive{
				conn.Timer.Reset(network.ResetConnTime)
			} else{
				newConn := network.ConnectionUDP{msg.Addr, time.NewTimer(network.ResetConnTime)}
				
				liftsOnline[id] = newConn
				go connTimer(newConn)
			}
		case network.NewOrder:
			//kanskje skrive ut noe?
			
			//cost :=  69 //costfunksjonen inn her
			
			//costMsg := network.Message{Content: network.Cost, Floor: msg.Floor, Button: msg.Button, Cost: cost}
			//network.MessageChan <- costMsg
		case network.CompletedOrder:
			//Slette order
		case network.Info:
			// LAGRER INFO OM HEISEN M/ TILHØRENDE ID, hvor?
		}		
}

func connTimer(conn network.ConnectionUDP){
	for{
	<-conn.Timer.C
	disconnElevChan <- conn
	}
}

func deleteLift(addr string){
	delete(liftsOnline, addr)
}
