package parser

import (
	"net"
	"reflect"
	"testing"
)

var validateTests = []struct {
	srcCIDR  string
	expected CIDRInfo
}{
	{"192.168.1.0/1", CIDRInfo{t: "ipv4",
		HostNum:   2147483646,
		Min:       net.IPv4(128, 0, 0, 1).To4(),
		Max:       net.IPv4(255, 255, 255, 254).To4(),
		Broadcast: net.IPv4(255, 255, 255, 255).To4()}},
	{"192.168.1.0/2", CIDRInfo{t: "ipv4",
		HostNum:   1073741822,
		Min:       net.IPv4(192, 0, 0, 1).To4(),
		Max:       net.IPv4(255, 255, 255, 254).To4(),
		Broadcast: net.IPv4(255, 255, 255, 255).To4()}},
	{"192.168.1.0/3", CIDRInfo{t: "ipv4",
		HostNum:   536870910,
		Min:       net.IPv4(192, 0, 0, 1).To4(),
		Max:       net.IPv4(223, 255, 255, 254).To4(),
		Broadcast: net.IPv4(223, 255, 255, 255).To4()}},
	{"192.168.1.0/4", CIDRInfo{t: "ipv4",
		HostNum:   268435454,
		Min:       net.IPv4(192, 0, 0, 1).To4(),
		Max:       net.IPv4(207, 255, 255, 254).To4(),
		Broadcast: net.IPv4(207, 255, 255, 255).To4()}},
	{"192.168.1.0/5", CIDRInfo{t: "ipv4",
		HostNum:   134217726,
		Min:       net.IPv4(192, 0, 0, 1).To4(),
		Max:       net.IPv4(199, 255, 255, 254).To4(),
		Broadcast: net.IPv4(199, 255, 255, 255).To4()}},
	{"192.168.1.0/6", CIDRInfo{t: "ipv4",
		HostNum:   67108862,
		Min:       net.IPv4(192, 0, 0, 1).To4(),
		Max:       net.IPv4(195, 255, 255, 254).To4(),
		Broadcast: net.IPv4(195, 255, 255, 255).To4()}},
	{"192.168.1.0/7", CIDRInfo{t: "ipv4",
		HostNum:   33554430,
		Min:       net.IPv4(192, 0, 0, 1).To4(),
		Max:       net.IPv4(193, 255, 255, 254).To4(),
		Broadcast: net.IPv4(193, 255, 255, 255).To4()}},
	{"192.168.1.0/8", CIDRInfo{t: "ipv4",
		HostNum:   16777214,
		Min:       net.IPv4(192, 0, 0, 1).To4(),
		Max:       net.IPv4(192, 255, 255, 254).To4(),
		Broadcast: net.IPv4(192, 255, 255, 255).To4()}},
	{"192.168.1.0/9", CIDRInfo{t: "ipv4",
		HostNum:   8388606,
		Min:       net.IPv4(192, 128, 0, 1).To4(),
		Max:       net.IPv4(192, 255, 255, 254).To4(),
		Broadcast: net.IPv4(192, 255, 255, 255).To4()}},
	{"192.168.1.0/10", CIDRInfo{t: "ipv4",
		HostNum:   4194302,
		Min:       net.IPv4(192, 128, 0, 1).To4(),
		Max:       net.IPv4(192, 191, 255, 254).To4(),
		Broadcast: net.IPv4(192, 191, 255, 255).To4()}},
	{"192.168.1.0/11", CIDRInfo{t: "ipv4",
		HostNum:   2097150,
		Min:       net.IPv4(192, 160, 0, 1).To4(),
		Max:       net.IPv4(192, 191, 255, 254).To4(),
		Broadcast: net.IPv4(192, 191, 255, 255).To4()}},
	{"192.168.1.0/15", CIDRInfo{t: "ipv4",
		HostNum:   131070,
		Min:       net.IPv4(192, 168, 0, 1).To4(),
		Max:       net.IPv4(192, 169, 255, 254).To4(),
		Broadcast: net.IPv4(192, 169, 255, 255).To4()}},
	{"192.168.1.0/16", CIDRInfo{t: "ipv4",
		HostNum:   65534,
		Min:       net.IPv4(192, 168, 0, 1).To4(),
		Max:       net.IPv4(192, 168, 255, 254).To4(),
		Broadcast: net.IPv4(192, 168, 255, 255).To4()}},
	{"192.168.1.0/17", CIDRInfo{t: "ipv4",
		HostNum:   32766,
		Min:       net.IPv4(192, 168, 0, 1).To4(),
		Max:       net.IPv4(192, 168, 127, 254).To4(),
		Broadcast: net.IPv4(192, 168, 127, 255).To4()}},
	{"192.168.1.0/18", CIDRInfo{t: "ipv4",
		HostNum:   16382,
		Min:       net.IPv4(192, 168, 0, 1).To4(),
		Max:       net.IPv4(192, 168, 63, 254).To4(),
		Broadcast: net.IPv4(192, 168, 63, 255).To4()}},
	{"192.168.1.0/19", CIDRInfo{t: "ipv4",
		HostNum:   8190,
		Min:       net.IPv4(192, 168, 0, 1).To4(),
		Max:       net.IPv4(192, 168, 31, 254).To4(),
		Broadcast: net.IPv4(192, 168, 31, 255).To4()}},
	{"192.168.1.0/20", CIDRInfo{t: "ipv4",
		HostNum:   4094,
		Min:       net.IPv4(192, 168, 0, 1).To4(),
		Max:       net.IPv4(192, 168, 15, 254).To4(),
		Broadcast: net.IPv4(192, 168, 15, 255).To4()}},
	{"192.168.1.0/21", CIDRInfo{t: "ipv4",
		HostNum:   2046,
		Min:       net.IPv4(192, 168, 0, 1).To4(),
		Max:       net.IPv4(192, 168, 7, 254).To4(),
		Broadcast: net.IPv4(192, 168, 7, 255).To4()}},
	{"192.168.1.0/22", CIDRInfo{t: "ipv4",
		HostNum:   1022,
		Min:       net.IPv4(192, 168, 0, 1).To4(),
		Max:       net.IPv4(192, 168, 3, 254).To4(),
		Broadcast: net.IPv4(192, 168, 3, 255).To4()}},
	{"192.168.1.0/23", CIDRInfo{t: "ipv4",
		HostNum:   510,
		Min:       net.IPv4(192, 168, 0, 1).To4(),
		Max:       net.IPv4(192, 168, 1, 254).To4(),
		Broadcast: net.IPv4(192, 168, 1, 255).To4()}},
	{"192.168.1.0/24", CIDRInfo{t: "ipv4",
		HostNum:   254,
		Min:       net.IPv4(192, 168, 1, 1).To4(),
		Max:       net.IPv4(192, 168, 1, 254).To4(),
		Broadcast: net.IPv4(192, 168, 1, 255).To4()}},
	{"192.168.1.0/25", CIDRInfo{t: "ipv4",
		HostNum:   126,
		Min:       net.IPv4(192, 168, 1, 1).To4(),
		Max:       net.IPv4(192, 168, 1, 126).To4(),
		Broadcast: net.IPv4(192, 168, 1, 127).To4()}},
	{"192.168.1.0/26", CIDRInfo{t: "ipv4",
		HostNum:   62,
		Min:       net.IPv4(192, 168, 1, 1).To4(),
		Max:       net.IPv4(192, 168, 1, 62).To4(),
		Broadcast: net.IPv4(192, 168, 1, 63).To4()}},
	{"192.168.1.0/27", CIDRInfo{t: "ipv4",
		HostNum:   30,
		Min:       net.IPv4(192, 168, 1, 1).To4(),
		Max:       net.IPv4(192, 168, 1, 30).To4(),
		Broadcast: net.IPv4(192, 168, 1, 31).To4()}},
	{"192.168.1.0/28", CIDRInfo{t: "ipv4",
		HostNum:   14,
		Min:       net.IPv4(192, 168, 1, 1).To4(),
		Max:       net.IPv4(192, 168, 1, 14).To4(),
		Broadcast: net.IPv4(192, 168, 1, 15).To4()}},
	{"192.168.1.0/29", CIDRInfo{t: "ipv4",
		HostNum:   6,
		Min:       net.IPv4(192, 168, 1, 1).To4(),
		Max:       net.IPv4(192, 168, 1, 6).To4(),
		Broadcast: net.IPv4(192, 168, 1, 7).To4()}},
	{"192.168.1.0/30", CIDRInfo{t: "ipv4",
		HostNum:   2,
		Min:       net.IPv4(192, 168, 1, 1).To4(),
		Max:       net.IPv4(192, 168, 1, 2).To4(),
		Broadcast: net.IPv4(192, 168, 1, 3).To4()}},
	{"192.168.1.0/31", CIDRInfo{t: "ipv4",
		HostNum:   2,
		Min:       net.IPv4(192, 168, 1, 0).To4(),
		Max:       net.IPv4(192, 168, 1, 1).To4(),
		Broadcast: net.IPv4(192, 168, 1, 1).To4()}},
	{"192.168.1.0/32", CIDRInfo{t: "ipv4",
		HostNum:   1,
		Min:       nil,
		Max:       nil,
		Broadcast: nil}},
	{"10.0.0.0/1", CIDRInfo{t: "ipv4",
		HostNum:   2147483646,
		Min:       net.IPv4(0, 0, 0, 1).To4(),
		Max:       net.IPv4(127, 255, 255, 254).To4(),
		Broadcast: net.IPv4(127, 255, 255, 255).To4()}},
	{"10.0.0.0/2", CIDRInfo{t: "ipv4",
		HostNum:   1073741822,
		Min:       net.IPv4(0, 0, 0, 1).To4(),
		Max:       net.IPv4(63, 255, 255, 254).To4(),
		Broadcast: net.IPv4(63, 255, 255, 255).To4()}},
	{"10.0.0.0/3", CIDRInfo{t: "ipv4",
		HostNum:   536870910,
		Min:       net.IPv4(0, 0, 0, 1).To4(),
		Max:       net.IPv4(31, 255, 255, 254).To4(),
		Broadcast: net.IPv4(31, 255, 255, 255).To4()}},
	{"10.0.0.0/4", CIDRInfo{t: "ipv4",
		HostNum:   268435454,
		Min:       net.IPv4(0, 0, 0, 1).To4(),
		Max:       net.IPv4(15, 255, 255, 254).To4(),
		Broadcast: net.IPv4(15, 255, 255, 255).To4()}},
	{"10.0.0.0/5", CIDRInfo{t: "ipv4",
		HostNum:   134217726,
		Min:       net.IPv4(8, 0, 0, 1).To4(),
		Max:       net.IPv4(15, 255, 255, 254).To4(),
		Broadcast: net.IPv4(15, 255, 255, 255).To4()}},
	{"10.0.0.0/6", CIDRInfo{t: "ipv4",
		HostNum:   67108862,
		Min:       net.IPv4(8, 0, 0, 1).To4(),
		Max:       net.IPv4(11, 255, 255, 254).To4(),
		Broadcast: net.IPv4(11, 255, 255, 255).To4()}},
	{"10.0.0.0/7", CIDRInfo{t: "ipv4",
		HostNum:   33554430,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(11, 255, 255, 254).To4(),
		Broadcast: net.IPv4(11, 255, 255, 255).To4()}},
	{"10.0.0.0/8", CIDRInfo{t: "ipv4",
		HostNum:   16777214,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 255, 255, 254).To4(),
		Broadcast: net.IPv4(10, 255, 255, 255).To4()}},
	{"10.0.0.0/9", CIDRInfo{t: "ipv4",
		HostNum:   8388606,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 127, 255, 254).To4(),
		Broadcast: net.IPv4(10, 127, 255, 255).To4()}},
	{"10.0.0.0/10", CIDRInfo{t: "ipv4",
		HostNum:   4194302,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 63, 255, 254).To4(),
		Broadcast: net.IPv4(10, 63, 255, 255).To4()}},
	{"10.0.0.0/11", CIDRInfo{t: "ipv4",
		HostNum:   2097150,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 31, 255, 254).To4(),
		Broadcast: net.IPv4(10, 31, 255, 255).To4()}},
	{"10.0.0.0/12", CIDRInfo{t: "ipv4",
		HostNum:   1048574,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 15, 255, 254).To4(),
		Broadcast: net.IPv4(10, 15, 255, 255).To4()}},
	{"10.0.0.0/13", CIDRInfo{t: "ipv4",
		HostNum:   524286,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 7, 255, 254).To4(),
		Broadcast: net.IPv4(10, 7, 255, 255).To4()}},
	{"10.0.0.0/14", CIDRInfo{t: "ipv4",
		HostNum:   262142,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 3, 255, 254).To4(),
		Broadcast: net.IPv4(10, 3, 255, 255).To4()}},
	{"10.0.0.0/15", CIDRInfo{t: "ipv4",
		HostNum:   131070,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 1, 255, 254).To4(),
		Broadcast: net.IPv4(10, 1, 255, 255).To4()}},
	{"10.0.0.0/16", CIDRInfo{t: "ipv4",
		HostNum:   65534,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 0, 255, 254).To4(),
		Broadcast: net.IPv4(10, 0, 255, 255).To4()}},
	{"10.0.0.0/17", CIDRInfo{t: "ipv4",
		HostNum:   32766,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 0, 127, 254).To4(),
		Broadcast: net.IPv4(10, 0, 127, 255).To4()}},
	{"10.0.0.0/18", CIDRInfo{t: "ipv4",
		HostNum:   16382,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 0, 63, 254).To4(),
		Broadcast: net.IPv4(10, 0, 63, 255).To4()}},
	{"10.0.0.0/19", CIDRInfo{t: "ipv4",
		HostNum:   8190,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 0, 31, 254).To4(),
		Broadcast: net.IPv4(10, 0, 31, 255).To4()}},
	{"10.0.0.0/20", CIDRInfo{t: "ipv4",
		HostNum:   4094,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 0, 15, 254).To4(),
		Broadcast: net.IPv4(10, 0, 15, 255).To4()}},
	{"10.0.0.0/21", CIDRInfo{t: "ipv4",
		HostNum:   2046,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 0, 7, 254).To4(),
		Broadcast: net.IPv4(10, 0, 7, 255).To4()}},
	{"10.0.0.0/22", CIDRInfo{t: "ipv4",
		HostNum:   1022,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 0, 3, 254).To4(),
		Broadcast: net.IPv4(10, 0, 3, 255).To4()}},
	{"10.0.0.0/23", CIDRInfo{t: "ipv4",
		HostNum:   510,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 0, 1, 254).To4(),
		Broadcast: net.IPv4(10, 0, 1, 255).To4()}},
	{"10.0.0.0/24", CIDRInfo{t: "ipv4",
		HostNum:   254,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 0, 0, 254).To4(),
		Broadcast: net.IPv4(10, 0, 0, 255).To4()}},
	{"10.0.0.0/25", CIDRInfo{t: "ipv4",
		HostNum:   126,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 0, 0, 126).To4(),
		Broadcast: net.IPv4(10, 0, 0, 127).To4()}},
	{"10.0.0.0/26", CIDRInfo{t: "ipv4",
		HostNum:   62,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 0, 0, 62).To4(),
		Broadcast: net.IPv4(10, 0, 0, 63).To4()}},
	{"10.0.0.0/27", CIDRInfo{t: "ipv4",
		HostNum:   30,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 0, 0, 30).To4(),
		Broadcast: net.IPv4(10, 0, 0, 31).To4()}},
	{"10.0.0.0/28", CIDRInfo{t: "ipv4",
		HostNum:   14,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 0, 0, 14).To4(),
		Broadcast: net.IPv4(10, 0, 0, 15).To4()}},
	{"10.0.0.0/29", CIDRInfo{t: "ipv4",
		HostNum:   6,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 0, 0, 6).To4(),
		Broadcast: net.IPv4(10, 0, 0, 7).To4()}},
	{"10.0.0.0/30", CIDRInfo{t: "ipv4",
		HostNum:   2,
		Min:       net.IPv4(10, 0, 0, 1).To4(),
		Max:       net.IPv4(10, 0, 0, 2).To4(),
		Broadcast: net.IPv4(10, 0, 0, 3).To4()}},
	{"10.0.0.0/31", CIDRInfo{t: "ipv4",
		HostNum:   2,
		Min:       net.IPv4(10, 0, 0, 0).To4(),
		Max:       net.IPv4(10, 0, 0, 1).To4(),
		Broadcast: net.IPv4(10, 0, 0, 1).To4()}},
	{"10.10.0.0/32", CIDRInfo{t: "ipv4",
		HostNum:   1,
		Min:       nil,
		Max:       nil,
		Broadcast: nil}},
}

func TestParse(t *testing.T) {
	for _, vt := range validateTests {
		ci, e := Parse(vt.srcCIDR)
		if e != nil {
			t.Errorf("Parse error: %#v", e)
		}

		if ci.t != vt.expected.t {
			t.Errorf("Parse(%v) error t, actual: %s, expected: %s", vt.srcCIDR, ci.t, vt.expected.t)
		}
		if ci.HostNum != vt.expected.HostNum {
			t.Errorf("Parse(%v) error HostNum, actual: %d, expected: %d", vt.srcCIDR, ci.HostNum, vt.expected.HostNum)
		}
		if !reflect.DeepEqual(ci.Min, vt.expected.Min) {
			t.Errorf("Parse(%v) error Min, actual: %s, expected: %s", vt.srcCIDR, ci.Min, vt.expected.Min)
		}
		if !reflect.DeepEqual(ci.Broadcast, vt.expected.Broadcast) {
			t.Errorf("Parse(%v) error Broadcast, actual: %s, expected: %s", vt.srcCIDR, ci.Broadcast, vt.expected.Broadcast)
		}
		if !reflect.DeepEqual(ci.Max, vt.expected.Max) {
			t.Errorf("Parse(%v) error Max, actual: %s, expected: %s", vt.srcCIDR, ci.Max, vt.expected.Max)
		}
	}

}
