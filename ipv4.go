package main

import (
	"log"
	"net"

	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
)

/*
https://www.ietf.org/rfc/rfc791.txt

    0                   1                   2                   3
    0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |Version|  IHL  |Type of Service|          Total Length         |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |         Identification        |Flags|      Fragment Offset    |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |  Time to Live |    Protocol   |         Header Checksum       |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |                       Source Address                          |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |                    Destination Address                        |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |                    Options                    |    Padding    |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

*/

type IPv4 []byte

//Version
func (f IPv4) Version() byte {
	return f[0] >> 4
}

//Internet Header Length
func (f IPv4) IHL() byte {
	/*
			Internet Header Length is the length of the internet header in 32
		    bit words, and thus points to the beginning of the data.  Note that
		    the minimum value for a correct header is 5.
	*/

	return f[0] & 0x0F
}

//Type of Service
func (f IPv4) Tos() byte {
	return f[1]
}

//Total Length
func (f IPv4) Tlen() []byte {
	return f[2:4]
}

//Identification
func (f IPv4) Id() []byte {
	return f[4:6]
}

//Flags
func (f IPv4) Flags() byte {
	return f[6]
}

//Offset
func (f IPv4) Offset() byte {
	return f[7]
}

//Time to Live
func (f IPv4) TTL() byte {
	return f[8]
}

//Protocol / ICMP = 0x1
func (f IPv4) Protocol() byte {
	return f[9]
}

//Checksum
func (f IPv4) Checksum() []byte {
	return f[10:12]
}

//Source
func (f IPv4) Source() net.IP {
	return net.IP(f[12:16])
}

//Source
func (f IPv4) Destination() net.IP {
	return net.IP(f[16:20])
}

//Options (if IHL > 5)
func (f IPv4) Options() []byte {
	//no need to handle this for now
	return f[20:23]
}

//Data
func (f IPv4) Data() []byte {
	return f[20:]
}

func handle_ipv4(ifce *water.Interface, frame ethernet.Frame) {
	var ipv4 IPv4 = frame.Payload()
	log.Println("ipv4 src :", ipv4.Source())
	log.Println("ipv4 Destination :", ipv4.Destination())
	log.Printf("ipv4 Payload: % x\n", ipv4.Data())

	switch ipv4.Protocol() {
	case 1: // ICMP
		go handle_icmp(ifce, &frame, &ipv4)
	case 17: //UDP
		go handle_udp(ifce, &frame, &ipv4)
	default:
		log.Println("[ipv4] UNKOWN Proto ", ipv4.Protocol())
	}
}
