package unibookBackend

import (
	"fmt"
	"net"
)

func FindOpenPort(start, end int) (int, error) {
	for port := start; port <= end; port++ {
		address := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", address)
		if err == nil {
			listener.Close()
			return port, nil
		}
	}
	return 0, fmt.Errorf("no open port found between %d and %d", start, end)
}
