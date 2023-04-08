package ipx

import (
	"fmt"
	"math/big"
	"net"
	"net/http"
	"regexp"
)

// RemoteIp 获取本机的 ip 地址
func RemoteIp(req *http.Request) string {
	remoteAddr := req.Header.Get("Remote_addr")
	if remoteAddr == "" {
		if ip := req.Header.Get("ipv4"); ip != "" {
			remoteAddr = ip
		} else if ip = req.Header.Get("XForwardedFor"); ip != "" {
			remoteAddr = ip
		} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
			remoteAddr = ip
		} else {
			remoteAddr = req.Header.Get("X-Real-Ip")
		}
	}

	if remoteAddr == "::1" || remoteAddr == "" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

// HadExist ipList 是用 | 分隔的字符串
func HadExist(ipList string, ip string) (ok bool, err error) {

	ok, err = regexp.MatchString(ipList, ip)

	return
}

// ToInt ip 转 10进制
func ToInt(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()

}

// ToString 10进制 转 ip
func ToString(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

// JudgeIp 校验ip地址
func JudgeIp(ip string) (ok bool) {
	return net.ParseIP(ip) != nil
}

// HasLocalIPAddr 检测 IP 地址字符串是否是内网地址
func HasLocalIPAddr(ip string) bool {
	return HasLocalIP(net.ParseIP(ip))
}

// HasLocalIP 检测 IP 地址是否是内网地址
func HasLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}

// 获取本机地址
func GetClientIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}

	return ""
}
