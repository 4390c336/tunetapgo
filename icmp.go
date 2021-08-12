package main

import (
	"log"

	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
)

type Icmp []byte

//Type
func (f Icmp) Type() byte {
	//8 for echo message
	//0 for echo reply message.
	return f[0]
}

//Code
func (f Icmp) Code() byte {
	//0 for ping/echo
	return f[1]
}

//Checksum
func (f Icmp) Checksum() []byte {
	//0 for ping/echo
	return f[2:4]
}

//Identifier
func (f Icmp) Identifier() []byte {
	return f[4:6]
}

//Sequence Number
func (f Icmp) SeqN() []byte {
	return f[6:8]
}

//Sequence Number
func (f Icmp) Data() []byte {
	return f[8:]
}

//checkSUM C
//codde from https://github.com/google/gopacket/blob/3aa782ce48d4a525acaebab344cedabfb561f870/layers/ip4.go#L158
func (f Icmp) CalcChecksum() uint16 {
	// Clear checksum bytes
	f[2] = 0
	f[3] = 0

	// Compute checksum
	var csum uint32
	for i := 0; i < len(f); i += 2 {
		csum += uint32(f[i]) << 8
		csum += uint32(f[i+1])
	}
	for {
		// Break when sum is less or equals to 0xFFFF
		if csum <= 65535 {
			break
		}
		// Add carry to the sum
		csum = (csum >> 16) + uint32(uint16(csum))
	}
	// Flip all the bits
	return ^uint16(csum)
}

func handle_icmp(ifce *water.Interface, frame *ethernet.Frame, ipv4 *IPv4) {
	var icmp Icmp = ipv4.Data()
	log.Println("[ICMP] Type:", icmp.Type())
	log.Println("[ICMP] Code:", icmp.Code())
	log.Printf("[ICMP] Checksum: %x", icmp.Checksum())
	log.Printf("[ICMP] Checksum: %x", icmp.CalcChecksum())
	log.Printf("[ICMP] Identifier: %x", icmp.Identifier())
	log.Printf("[ICMP] SeqN: %x", icmp.SeqN())
	log.Printf("[ICMP] Data: % x \n", icmp.Data())

	switch icmp.Type() {
	case 8:
		//ping
		//swap address for IP
		//change ICMP type to 0
		var MacSrc = frame.Source()
		var MacDst = frame.Destination()
		var response ethernet.Frame
		response.Prepare(MacSrc, MacDst, 0, ethernet.IPv4, len(*ipv4))

	default:
		log.Println("[ICMP] Uknown Code", icmp.Type())
	}
}
