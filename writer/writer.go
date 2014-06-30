package writer

import (
	"bytes"
	"fmt"
	"github.com/goldeneggg/ipcl/parser"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	Out     io.Writer = os.Stdout
	fpf               = fmt.Fprintf
	headers           = []string{"source_cidr",
		"network",
		"mask",
		"host_num",
		"min_address",
		"max_address",
		"broadcast"}
)

type Writer interface {
	Write(cidrs []parser.CIDRInfo)
}

type DefaultWriter struct {
	w io.Writer
}

type SepWriter struct {
	*DefaultWriter
	sep string
}

func (dw *DefaultWriter) Write(cidrs []parser.CIDRInfo) {
	for _, cidr := range cidrs {
		dw.writeSingle(cidr)
	}
}

func (dw *DefaultWriter) writeSingle(cidr parser.CIDRInfo) {
	fpf(dw.w, "%s : %s\n", headers[0], cidr.SrcCIDR)
	fpf(dw.w, "%s     : %s\n", headers[1], cidr.Network)
	fpf(dw.w, "%s        : %s\n", headers[2], mask2string(cidr.Mask))
	fpf(dw.w, "%s    : %d\n", headers[3], cidr.HostNum)
	fpf(dw.w, "%s : %s\n", headers[4], cidr.Min)
	fpf(dw.w, "%s : %s\n", headers[5], cidr.Max)
	fpf(dw.w, "%s   : %s\n", headers[6], cidr.Broadcast)
	fpf(dw.w, "\n")
}

func (sw *SepWriter) Write(cidrs []parser.CIDRInfo) {
	sw.writeHeader()
	for _, cidr := range cidrs {
		sw.writeLine(cidr)
	}
}

func (sw *SepWriter) writeHeader() {
	fpf(sw.w, "%s\n", strings.Join(headers, sw.sep))
}

func (sw *SepWriter) writeLine(cidr parser.CIDRInfo) {
	s := []string{cidr.SrcCIDR,
		cidr.Network.String(),
		mask2string(cidr.Mask),
		strconv.Itoa(cidr.HostNum),
		cidr.Min.String(),
		cidr.Max.String(),
		cidr.Broadcast.String()}
	fpf(sw.w, "%s\n", strings.Join(s, sw.sep))
}

func NewWriter(isCsv bool, isTsv bool) Writer {
	defWriter := &DefaultWriter{Out}
	if isCsv {
		return &SepWriter{defWriter, ","}
	} else if isTsv {
		return &SepWriter{defWriter, "\t"}
	} else {
		return defWriter
	}
}

func mask2string(mask []byte) string {
	var buf bytes.Buffer
	for i, m := range mask {
		buf.WriteString(itod(uint(m)))
		if i < len(mask)-1 {
			buf.WriteString(".")
		}
	}

	return buf.String()
}

func itod(i uint) string {
	if i == 0 {
		return "0"
	}

	// Assemble decimal in reverse order.
	var b [32]byte
	bp := len(b)
	for ; i > 0; i /= 10 {
		bp--
		b[bp] = byte(i%10) + '0'
	}

	return string(b[bp:])
}
