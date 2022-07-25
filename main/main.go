package main

import (
	"fmt"
	"port-scanner/port"
)

func main() {
	fmt.Println("Scanning Ports")
	results := port.InitialScan("localhost")
	fmt.Println(results)
}
