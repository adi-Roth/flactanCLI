package system

import (
	"net"
	"time"
)

// CheckInternet verifies if the system has internet access
func CheckInternet() bool {
	timeout := 2 * time.Second
	conn, err := net.DialTimeout("tcp", "1.1.1.1:80", timeout)
	if err != nil {
		return false
	}
	_ = conn.Close()
	return true
}
