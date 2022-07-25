package port

import (
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	Port  string
	State string
}

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
