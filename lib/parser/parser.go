package parser

import (
	"bytes"
	"fmt"
	"math/big"
	"net"
	"strings"
)

const (
	TYPE_IPV4 = "ipv4"
	TYPE_IPV6 = "ipv6"
)

type CIDRInfo struct {
	t         string
	ipNet     *net.IPNet
	Network   net.IP     // []byte
	Mask      net.IPMask // []byte
	ones      int
	bits      int
	SrcCIDR   string
	HostNum   int
	Min       net.IP // []byte
	Max       net.IP // []byte
	Broadcast net.IP // []byte
}

func Parse(srcCIDR string) (CIDRInfo, error) {
	var cidr CIDRInfo

	ip, ipNet, eParse := net.ParseCIDR(srcCIDR)
	if eParse != nil {
		return cidr, eParse
	}

	t, eType := getType(ip)
	if eType != nil {
		return cidr, eType
	}

	cidr.t = t
	cidr.SrcCIDR = srcCIDR
	cidr.ipNet = ipNet
	cidr.Network = ipNet.IP
	cidr.Mask = ipNet.Mask
	cidr.ones, cidr.bits = cidr.Mask.Size()

	if cidr.t == TYPE_IPV4 {
		// calculate
		cidr.calcIPv4()
	}

	return cidr, nil
}

func getType(ip net.IP) (string, error) {
	if ip4 := ip.To4(); len(ip4) == net.IPv4len {
		return TYPE_IPV4, nil
	} else if ip6 := ip.To16(); len(ip6) == net.IPv6len {
		// TODO ipv6
		return "", fmt.Errorf("ip %+v is not IPv4\n", ip)
	} else {
		return "", fmt.Errorf("ip %+v is not IPv4\n", ip)
	}
}

func (cidr *CIDRInfo) calcIPv4() {
	// Host Num
	cidr.calcHostNumV4()

	// Min, Max, Broadcast
	if cidr.HostNum > 1 {
		cidr.calcAddressesV4()
	}
}

func (cidr *CIDRInfo) calcHostNumV4() {
	var hostNum int
	if cidr.ones == cidr.bits {
		hostNum = 1
	} else if cidr.ones == cidr.bits-1 {
		hostNum = 2
	} else {
		hostNum = int(new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(cidr.bits-cidr.ones)), nil).Int64()) - 2
	}
	cidr.HostNum = hostNum
}

func (cidr *CIDRInfo) calcAddressesV4() {
	// Broadcast Address
	cidr.calcBroadCast()

	// Min IP Address
	if cidr.ones == 31 {
		// XXX
		cidr.Min = cidr.Network
	} else {
		minIp := makeIpV4(cidr.Network)
		minIp[3] += 1
		cidr.Min = minIp
	}

	// Max IP Address
	if cidr.ones == 31 {
		// XXX
		cidr.Max = cidr.Broadcast
	} else {
		maxIp := makeIpV4(cidr.Broadcast)
		maxIp[3] -= 1
		cidr.Max = maxIp
	}
}

func (cidr *CIDRInfo) calcBroadCast() {
	var buf bytes.Buffer
	for _, octet := range cidr.Network {
		buf.WriteString(byte2binstr(octet))
	}

	bcBinStr := buf.String()[:cidr.ones] + strings.Repeat("1", cidr.bits-cidr.ones)

	broad := make([]byte, 4)
	for i, f, t := 0, 0, 8; i < 4; i, f, t = i+1, f+8, t+8 {
		broad[i] = binstr2byte(bcBinStr[f:t])
	}

	cidr.Broadcast = makeIpV4(broad)
}

func byte2binstr(b byte) string {
	return fmt.Sprintf("%08b", b)
}

func binstr2byte(binstr string) byte {
	var i, j int
	var bv byte
	for i, j = len(binstr)-1, 0; i >= 0; i, j = i-1, j+1 {
		if binstr[i] != '0' {
			bv |= (1 << uint(j))
		}
	}

	return bv
}

func binstr2hexstr(binstr string) string {
	return fmt.Sprintf("%02x", binstr2byte(binstr))
}

func makeIpV4(src []byte) net.IP {
	dst := make(net.IP, net.IPv4len)
	copy(dst, src)

	return dst
}

func (cidr *CIDRInfo) Contains(srcIP string) bool {
	return cidr.ipNet.Contains(net.ParseIP(srcIP))
}
