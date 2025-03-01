package system

import (
	"net/http"
	"time"
)

// CheckInternet verifies if the system has internet access
func CheckInternet() bool {
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	_, err := client.Get("https://www.google.com")
	return err == nil
}
