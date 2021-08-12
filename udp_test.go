package main

import (
	"testing"
)

func TestUdp(t *testing.T) {
	//hexdump(UDP(dport=123)/Raw(load="abc"))
	var udp Udp = []byte{0x00, 0x35, 0x00, 0x7B, 0x00, 0x0B, 0x00, 0x00, 0x61, 0x62, 0x63}

	t.Log("[UDP] SrcPort:", udp.SrcPort())
	t.Log("[UDP] DstPort:", udp.DstPort())
	t.Log("[UDP] Length:", udp.Length())
	t.Log("[UDP] Checksum:", udp.Checksum())
	t.Logf("[UDP] Data: % x \n", udp.Data())
}
