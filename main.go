package main

import (
	"SusufscPPP/Plugins"
	"SusufscPPP/common"
	"fmt"
	"os"
	"time"
)

const timeout = 500 * time.Millisecond

func main() {

	start := time.Now()
	var Info common.HostInfo
	var is172, is10, is192 bool

	common.Flag(&Info)
	if common.Route || common.Routeall {
		fmt.Println(common.Route)
		is172, is10, is192, _ = Plugins.Subip()
		fmt.Printf("is172: %v, is10: %v, is192: %v\n", is172, is10, is192)
		if is192 || common.Routeall {
			fmt.Println("正在扫描 192.168.*.*")
			err := Plugins.ScanNetwork("ip192all.txt")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("扫描结束结果保存至 ip192all.txt")
		}
		if is172 || common.Routeall {
			fmt.Println("正在扫描 172.16.*.*")
			err := Plugins.ScanNetwork172("ipall172.txt")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("扫描结束结果保存至 ipall172.txt")
		}
		if is10 || common.Routeall {
			fmt.Println("正在扫描 10.*.*.*")
			err := Plugins.ScanNetwork10("ipall10.txt")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("扫描结束结果保存至 ipall10.txt")
		}
		os.Exit(1)
	}

	common.Parse(&Info)
	Plugins.Scan(Info)
	fmt.Printf("[*] 扫描结束,耗时: %s\n", time.Since(start))
}
