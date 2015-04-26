package filehandler

import(
	"encoding/gob"
	"os"
	"network"
)

//var file string

//func Init(){
	//file = "backup.txt"
	//f, err := os.Create(file)
	//if err != nil {
		//panic("Can't create backup.txt")
	//}
//}

func SaveBackup(/*file string,*/ LiftsOnlineInfo map[string]network.ElevatorInfo) {
	f, err := os.Create("backup.txt")
	if err != nil {
		panic("Can't create backup.txt")
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	if err := enc.Encode(LiftsOnlineInfo); err != nil {
		panic("Can't encode LiftsOnlineInfo")
    }
}

func LoadBackup(/*file string*/) (LiftsOnlineInfo map[string]network.ElevatorInfo) {
	f, err := os.Open("backup.txt")
	if err != nil {
			panic("Cant' open backup.txt")
	}
	defer f.Close()

	dec:= gob.NewDecoder(f)
	if err := dec.Decode(&LiftsOnlineInfo); err != nil {
			panic("Can't decode from backup.txt")
	}
	return LiftsOnlineInfo
}
