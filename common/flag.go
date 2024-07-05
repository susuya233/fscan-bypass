package common

import (
	"flag"
)

func Banner() {
	banner := `
   ___                              _    
  / _ \     ___  ___ _ __ __ _  ___| | __ 
 / /_\/____/ __|/ __| '__/ _` + "`" + ` |/ __| |/ /
/ /_\\_____\__ \ (__| | | (_| | (__|   <    
\____/     |___/\___|_|  \__,_|\___|_|\_\   
                     
					 
`
	print(banner)
}

func Flag(Info *HostInfo) {
	Banner()
	flag.BoolVar(&Route, "route", false, "根据路由表测试路由连通")
	flag.BoolVar(&Route, "routeall", false, "测试全部路由可能的链路的连通")
	flag.StringVar(&Info.Host, "h", "", "要扫描的主机的IP地址,例如: 192.168.11.11 | 192.168.11.11-255 | 192.168.11.11,192.168.11.12")
	flag.StringVar(&NoHosts, "hn", "", "扫描时跳过的IP地址,例如: -hn 192.168.1.1/24")
	flag.StringVar(&Ports, "p", DefaultPorts, "需要扫描的目标端口,例如: 22 | 1-65535 | 22,80,3306")
	flag.StringVar(&PortAdd, "pa", "", "添加默认扫描端口,例如:-pa 3389")
	flag.StringVar(&UserAdd, "usera", "", "添加默认用户,例如:-usera user")
	flag.StringVar(&PassAdd, "pwda", "", "添加默认密码,例如:-pwda password")
	flag.StringVar(&NoPorts, "pn", "", "扫描时跳过的端口,例如: -pn 445")
	flag.StringVar(&Command, "c", "", "ssh命令执行 (ssh|wmiexec)")
	flag.StringVar(&SshKey, "sshkey", "", "sshsshkey文件 (id_rsa)")
	flag.StringVar(&Domain, "domain", "", "smb爆破模块时,设置域名")
	flag.StringVar(&Username, "user", "", "指定爆破时的用户名")
	flag.StringVar(&Password, "pwd", "", "指定爆破时的密码	")
	flag.Int64Var(&Timeout, "time", 3, "设置超时时间")
	flag.StringVar(&Scantype, "m", "all", "设置扫描类型,as: -m ssh")
	flag.StringVar(&Path, "path", "", "fcgi、smb、romote 文件路径")
	flag.IntVar(&Threads, "t", 700, "线程数量")
	flag.IntVar(&LiveTop, "top", 10, "show live len top")
	flag.StringVar(&HostFile, "hf", "", "从文件获取目标ip,例如: -hf ip.txt")
	flag.StringVar(&Userfile, "userf", "", "用户名文件")
	flag.StringVar(&Passfile, "pwdf", "", "密码文件")
	flag.StringVar(&PortFile, "portf", "", "端口列表文件")
	flag.StringVar(&PocPath, "pocpath", "", "POC目录")
	flag.StringVar(&RedisFile, "rf", "", "redis未授权,文件写入sshkey文件(例如: -rf id_rsa.pub)")
	flag.StringVar(&RedisShell, "rs", "", "redis计划任务反弹shell的ip端口(as: -rs 192.168.1.1:6666)")
	flag.BoolVar(&NoPoc, "nopoc", false, "跳过webPOC扫描")
	flag.BoolVar(&IsBrute, "nobr", false, "跳过sql、ftp、ssh等的密码爆破")
	flag.IntVar(&BruteThread, "br", 1, "跳过存活探测")
	flag.BoolVar(&NoPing, "np", false, "不进行ping监测")
	flag.BoolVar(&Ping, "ping", false, "使用ping代替icmp进行存活探测")
	flag.StringVar(&Outputfile, "o", "result.txt", "输出文件")
	flag.BoolVar(&TmpSave, "no", false, "扫描结果不保存到文件中	")
	flag.Int64Var(&WaitTime, "debug", 30, "未响应,打印当前进度时常(默认30s)")
	flag.BoolVar(&Silent, "silent", false, "静默扫描")
	flag.BoolVar(&Nocolor, "nocolor", false, "不进行颜色输出")
	flag.BoolVar(&PocFull, "full", false, "poc全扫描,例如:shiro 100key")
	flag.StringVar(&URL, "u", "", "url")
	flag.StringVar(&UrlFile, "uf", "", "url文件")
	flag.StringVar(&Pocinfo.PocName, "pocname", "", "指定poc文件名称,可模糊指定, -pocname weblogic")
	flag.StringVar(&Proxy, "proxy", "", "设置代理服务器,例如: -proxy http://127.0.0.1:8080")
	flag.StringVar(&Socks5Proxy, "socks5", "", "设置socks5代理用于tcp连接, 连接超时将不起作用.")
	flag.StringVar(&Cookie, "cookie", "", "设置cookie,例如:-cookie rememberMe=login")
	flag.Int64Var(&WebTimeout, "wt", 5, "设置web请求超时时间(默认5s)")
	flag.BoolVar(&DnsLog, "dns", false, "使用dnslogPOC")
	flag.IntVar(&PocNum, "num", 20, "POC速率(默认20)")
	flag.StringVar(&SC, "sc", "", "ms17 shellcode,例如 -sc add")
	flag.BoolVar(&IsWmi, "wmi", false, "开启wmi")
	flag.StringVar(&Hash, "hash", "", "hash")
	flag.BoolVar(&Noredistest, "noredis", false, "不测试redis")
	flag.BoolVar(&JsonOutput, "json", false, "json格式输出")
	flag.Parse()
}
