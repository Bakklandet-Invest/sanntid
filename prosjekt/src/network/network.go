package network

import (
	"fmt"
	"time"
	"encoding/json"
)

var RecieveChan = make(chan messageUDP) //Public
var sendChan = make(chan messageUDP)

const NotifyInterval = 20*time.Second

func Init(){
	const localListenPort = 14369
	const broadcastListenPort = 14370
	const messageSize = 1024

	err := InitUDP(localListenPort, broadcastListenPort, messageSize, sendChan, RecieveChan)
	if err != nil {
		fmt.Printf("InitUDP error: %v \n", err)
	}

	go aliveNotifier()
	go sendMessages()
}

func ParseMessage(msgUDP messageUDP) Message {
	fmt.Printf("Before parse: %s from %s\n", string(msgUDP.data), msgUDP.recieveAddr)

	var msg Message
	if err := json.Unmarshal(msgUDP.data[:msgUDP.length], &msg); err != nil {
		fmt.Printf("json.Unmarshal error: %s\n", err)
	}

	msg.Addr = msgUDP.recieveAddr
	fmt.Printf("After Unmarshal: %s\n", msg.Addr)
	return msg
}

func aliveNotifier() {
	alive := Message{Content: Alive, Addr: "broadcast", Floor: -1, Button: -1}
	for {
		MessageChan <- alive
		time.Sleep(NotifyInterval)
	}
}

func sendMessages() {
	for {
		fmt.Println("sendMessages waiting on MessageChan")
		msg := <-MessageChan

		PrintMessage(msg)
		raddr := msg.Addr
		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			fmt.Print("json.Marshal error: %s \n", err)
		}		
		
		sendChan <- messageUDP{recieveAddr: raddr, data: jsonMsg, length: len(jsonMsg)}

		time.Sleep(time.Millisecond)
	}
}

func PrintMessage(msg Message) {
	fmt.Printf("\n-----Message start-----\n")
	switch msg.Content {
	case Alive:
		fmt.Println("I'm alive")
	case NewOrder:
		fmt.Println("New order")
	case CompletedOrder:
		fmt.Println("Completed order")
	case Info:
		fmt.Println("Elevator info:")
	default:
		fmt.Println("Invalid message type\n")
	}
	fmt.Printf("Floor: %d\n", msg.Floor)
	fmt.Printf("Button: %d\n", msg.Button)
	fmt.Printf("Elev now in floor %v with direction %v\n", msg.ElevInfo.Floor, msg.ElevInfo.Dir)
	fmt.Println("-----Message end-------\n")
}


