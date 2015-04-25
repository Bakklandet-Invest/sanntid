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
/* ---HUSKELISTE
	ElevInfo må pakkes i ElevatorRun
	* ConnTimer tar kun inn network.ConnectionUDP, og vet ikke ID'en
	* -> fikset at den sletter nå
*/ /* */
func main(){
	updateOutChan := make(chan network.Message)
	updateInChan := make(chan network.Message)//Trenger ikke egen case, messageHandler tar ansvar
	extOrderChan := make(chan ButtonSignal)
	newOrderChan := make(chan ButtonSignal)
	fromMasterChan := make(chan ButtonSignal)//Kanalen som går inn til elev og gir oppdrag
	completeOrderChan := make(chan ButtonSignal)
	
	// Holder orden på master/slave-rollen	
	terminateChan := make(chan bool)
	terminatedChan := make(chan bool)
	checkMasterChan := make(chan bool)	
	go checkMaster(checkMasterChan<-)
	
	go networkHandler()
	go ELEVATORRUN(terminateChan, terminatedChan)
	go Slave()
	
}

func Slave(){		
	//go networkHandler()
	var masterAddr string //evt lagre sin egen slik at meldingene sendes til seg selv
	// og ikke forsvinner i vente på en master
	for{
		select{
			case master := <- checkMasterChan:
				if master == myID{
					terminateChan<-true
					<-terminatedChan
					go Master()
					return
				} //PROBLEM UNDER- i init fasen, vil ikke masterADdr eksistere
				masterAddr = liftsOnline(master)
			case update := <- updateOutChan:
				updateMsg := network.Message{Content: network.Info, Addr: "broadcast", ElevInfo: update}
				network.MessageChan <- updateMsg
			case extOrdButtonSignal := <- extOrderChan:
				extOrdMsg := network.Message{Content: network.NewOrder, Addr: <addresse til master>, Floor: extOrdButtonSignal.Floor, Button: extOrdButtonSignal.Button}
				network.MessageChan <- extOrdMsg //Få pakket forunftig beskjed her først da, med recieverAddr til MASTER
			// PROBLEM - lagre extOrd lokalt i tilfelle intet svar, kjøre da for å være sikker??
	
			case orderButtonSignal := <- newOrderChan:
				fromMasterChan <- orderButtonSignal
			case completeOrder := <- completeOrderChan
				complOrdMsg := network.Message{Content: network.CompletedOrder, Addr: <addresse til master>, Floor: extOrdButtonSignal.Floor, Button: extOrdButtonSignal.Button}
				network.MessageChan <- complOrdMsg
		}	
	}
}

func Master(){

	for{
		select{
			case master := <- checkMasterChan:
				if master != myID{
					terminateChan<-true
					<-terminatedChan
					go Slave()
					return
				}
			case update := <- updateOutChan:
				updateMsg := network.Message{Content: network.Info, Addr: "broadcast", ElevInfo: update}
				network.MessageChan <- updateMsg
			case extOrdButtonSignal := <- extOrderChan:
				extOrdMsg := network.Message{Content: network.NewOrder, Addr: <addresse til master>, Floor: extOrdButtonSignal.Floor, Button: extOrdButtonSignal.Button}
				network.MessageChan <- extOrdMsg //Få pakket forunftig beskjed her først da, med recieverAddr til MASTER
			case orderButtonSignal := <- newOrderChan: //Finner beste heis til å ta jobben og sender til den
				//returnerer heis best egnet for jobbet
				heisID = costfunction(orderButtonSignal) //order inneholder opp/ned+etasje
		// --- LAGRE ordren i uncompleteOrders og slette igjen når completed er mottatt??
				if heisID == myID{
					fromMasterChan<-orderButtonSignal
				}else{
					addrOrderReciever := liftsOnline(heisID).Addr
					newOrderMsg := network.Message{Content: network.NewOrder, Addr: addrOrderReciever, Floor: orderButtonSignal.Floor, Button: orderButtonSignal.Button}
					network.MessageChan <- newOrderMsg
				}
			case completeOrder := <- completeOrderChan:
				// slette fra uncompleteOrders (PS; SYNKE uncompleteOrders til alle for backup?) (no orders lost)
				
		}
	}
}


/*
func main(){
	var kanal = make(chan int)
	
	go networkHandler()
	
	<-kanal
}*/

func networkHandler(){
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

//  msg.Addr byttet med id i liftsonline(key)
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
			// KANAL skrive orderen til canalen
			/*tempButtonSignal := ButtonSignal{Floor: msg.Floor, Button: msg.Button} 
			newOrderChan<-tempButtonSignal*/
			newOrderChan<-ButtonSignal{Floor: msg.Floor, Button: msg.Button} 
			//cost :=  69 //costfunksjonen inn her
			
			//costMsg := network.Message{Content: network.Cost, Floor: msg.Floor, Button: msg.Button, Cost: cost}
			//network.MessageChan <- costMsg
		case network.CompletedOrder:
			//Slette order - denne casen er kun for master??
			completeOrderChan <- ButtonSignal{Floor: msg.Floor, Button: msg.Button}
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

//endret slik at den sletter vha idkey, ikke addr key
func deleteLift(addr string){
	id := network.FindID(addr)
	delete(liftsOnline, id)
}
