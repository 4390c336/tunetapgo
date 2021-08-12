package main

import (
	"encoding/binary"
	"log"

	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
)

type Udp []byte

//Source port
func (f Udp) SrcPort() uint16 {
	return binary.BigEndian.Uint16(f[0:2])
}

//Destination port
func (f Udp) DstPort() uint16 {
	return binary.BigEndian.Uint16(f[2:4])
}

//Length
func (f Udp) Length() uint16 {
	return binary.BigEndian.Uint16(f[4:6])
}

//Checksum (Optional on IPv4)
func (f Udp) Checksum() []byte {
	return f[6:8]
}

//Data ( should be based on the LEN field ...)
func (f Udp) Data() []byte {
	return f[8:]
}

func handle_udp(ifce *water.Interface, frame *ethernet.Frame, ipv4 *IPv4) {
	var udp Udp = ipv4.Data()
	log.Printf("[UDP] data % x", udp.Data())
}
