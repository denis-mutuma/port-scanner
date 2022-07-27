package main

import (
	"fmt"

	"github.com/denis-mutuma/port-scanner/port"
)

// the entry point of the program
func main() {
	fmt.Println("Scanning Ports")
	results := port.InitialScan("localhost")
	fmt.Println(results)
}
