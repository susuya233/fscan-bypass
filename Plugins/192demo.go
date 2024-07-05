package Plugins

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func isAlive(ip string) bool {
	cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
	err := cmd.Run()
	return err == nil
}
func ScanNetwork(filename string) error {
	networks := make(map[string]bool)
	var mu sync.Mutex
	var wg sync.WaitGroup
	ipCh := make(chan string, 1024)
	for i := 0; i < 256; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ip := range ipCh {
				if isAlive(ip) {
					network := ip[:len(ip)-1] + "0/24"
					mu.Lock()
					networks[network] = true
					mu.Unlock()
				}
			}
		}()
	}
	for i := 0; i < 256; i++ {
		ipCh <- fmt.Sprintf("192.168.%d.1", i) //kaishi
		ipCh <- fmt.Sprintf("192.168.%d.2", i)
		ipCh <- fmt.Sprintf("192.168.%d.3", i)
		ipCh <- fmt.Sprintf("192.168.%d.254", i)
	} //
	close(ipCh)
	wg.Wait()
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("ERROR（一般看到这个报错代表要不你编译有问题要不你机器有问题）: %v", err)
	}
	defer file.Close()
	for network := range networks {
		_, err := file.WriteString(network + "\n")
		if err != nil {
			return fmt.Errorf("ERROE（一般看到这个报错代表要不你编译有问题要不你机器有问题）: %v", err)
		}
	}
	return nil
}
