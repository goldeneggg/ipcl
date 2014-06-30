package writer

import (
	//"fmt"
	"github.com/goldeneggg/ipcl/parser"
)

type defaultFormat struct {
	source_cidr string
	network     string
	mask        string
	min_address string
	max_address string
	board_cast  string
}

var vals = []struct {
	srcCIDR  string
	expected defaultFormat
}{
	{"192.168.1.0/24", defaultFormat{
		source_cidr: "source_cidr : 192.168.56.0/24",
		network:     "network     : 192.168.56.0",
		mask:        "mask        : 255.255.255.0",
		min_address: "min_address : 192.168.56.1",
		max_address: "max_address : 192.168.56.254",
		board_cast:  "broadcast   : 192.168.56.255"}},
}

func ExampleWriter_Write() {
	writer := NewWriter(false, false)

	cis := make([]parser.CIDRInfo, len(vals))
	for _, vt := range vals {
		ci, _ := parser.Parse(vt.srcCIDR)
		cis = append(cis, ci)
	}

	writer.Write(cis)
	//	// Output:
	//	// source_cidr : 192.168.56.0/24
	//	// network     : 192.168.56.0
	//	// mask        : 255.255.255.0
	//	// host_num    : 254
	//	// min_address : 192.168.56.1
	//	// max_address : 192.168.56.254
	//	// broadcast   : 192.168.56.255
}
