package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {	
	reader := bufio.NewReader(os.Stdin)
	var text string = "Initialized"
	
	
	for {
		fmt.Print("Enter text: ")
		text, _ = reader.ReadString('\n')
				
		fmt.Println(text[:len(text)-1])

		if strings.ToLower(text) == "disconnect\n" {
			fmt.Println("disconnecting..")
			break
		}
	}
	
}

