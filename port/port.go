package port

import (
	"net"
	"strconv"
	"time"
)

// a struct of the scan results
// Port: port number
// State: "Open" or "Closed"
type ScanResult struct {
	Port  string
	State string
}

// a function to scan the ports
// this function takes the protocol, hostname and port number and
// returns a ScanResult type
func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: strconv.Itoa(port) + string("/") + protocol}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed"
		return result
	}
	defer conn.Close()
	result.State = "Open"
	return result
}

// this function scans udp and tcp ports for a specified range
// and returns an array of type ScanResult
func InitialScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 0; i < 10; i++ {
		results = append(results, ScanPort("udp", hostname, i))
	}

	for i := 0; i <= 10; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
	}
	return results
}
