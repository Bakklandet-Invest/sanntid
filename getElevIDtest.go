package main



import (
	//"llist"
	"strconv"
	"net"
	"fmt"
	//"time"
	)
/*
type Hei struct {
	balle string
	out chan Hei
}

func initHei() *Hei {
	h := new(Hei)
	h.balle = "hei"
	h.out = make(chan Hei)
	return h	
}

func (h *Hei) yes() {
	h.out <- *h
}
*/
func findElevID() int {
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                IDstr := ipnet.IP.String()
                IDstr = IDstr[12:]
                ID, _ := strconv.Atoi(IDstr)
                return ID
            }
        }
    }
    return 0
}

func main() {
	fmt.Println(findElevID())
	
	/*
	h := initHei()
	go h.yes()
	time.Sleep(time.Second)
	l := <- h.out
	fmt.Println(l.balle)
	*/
	
	return
}
