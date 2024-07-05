// scan172.go
package Plugins

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

// isAlive checks if an IP address is alive by sending a ping request.
func isAlive1(ip string) bool {
	cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
	err := cmd.Run()
	return err == nil
}

// ScanNetwork172 scans possible gateway addresses in the 172.16.x.x to 172.31.x.x network and writes alive networks to a file.
func ScanNetwork172(filename string) error {
	networks := make(map[string]bool)
	var mu sync.Mutex
	var wg sync.WaitGroup

	ipCh := make(chan string, 1024)

	// Worker goroutines
	for i := 0; i < 256; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ip := range ipCh {
				if isAlive1(ip) {
					network := ip[:len(ip)-2] + "0/24"
					mu.Lock()
					networks[network] = true
					mu.Unlock()
				}
			}
		}()
	}

	// Produce possible gateway addresses
	for i := 16; i <= 31; i++ {
		for j := 0; j < 256; j++ {
			ipCh <- fmt.Sprintf("172.%d.%d.1", i, j)
		}
	}
	close(ipCh)

	// Wait for all goroutines to finish
	wg.Wait()

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	for network := range networks {
		_, err := file.WriteString(network + "\n")
		if err != nil {
			return fmt.Errorf("error writing to file: %v", err)
		}
	}

	return nil
}
