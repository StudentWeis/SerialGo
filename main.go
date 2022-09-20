// A Serial Assistant tool writed by Go.
package main

import (
	"flag"
	"fmt"
	"log"

	"go.bug.st/serial"
)

// Port Configuration.
var confPort = flag.String("c", "COM6", "Serial port like COM6.")
var confBaud = flag.Int("b", 115200, "BaudRate.")

// Function.
var confList = flag.Bool("l", false, "List the ports.")

func main() {
	flag.Parse()
	if *confList {
		getPortList()
	} else {
		mode := &serial.Mode{
			BaudRate: *confBaud,
		}
		port, err := serial.Open(*confPort, mode)
		if err != nil {
			fmt.Printf("[error] Can't open port.")
		}
		err = port.SetMode(mode)
		if err != nil {
			fmt.Printf("[error] Can't set mode.")
		}
		buff := make([]byte, 100)
		for {
			n, err := port.Read(buff)
			if err != nil {
				log.Fatal(err)
				break
			}
			if n == 0 {
				fmt.Println("\nEOF")
				break
			}
			fmt.Printf("%v", string(buff[:n]))
		}
	}
}

// Get the list of Serial Ports.
func getPortList() {
	ports, err := serial.GetPortsList()
	if err != nil {
		fmt.Printf("[error] Can't get list.")
	}
	if len(ports) == 0 {
		fmt.Printf("[warning] No Port.")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
}
