package main

import (
	"log"
	"net"

	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
)

//arp struct :
type Arp []byte

//Hardware Type
func (f Arp) HT() []byte {
	return f[0:2]
}

//Protocol Type
func (f Arp) PT() []byte {
	return f[2:4]
}

//Hardware Address Length
func (f Arp) HAL() byte {
	return f[4]
}

// Protocol Address Length
func (f Arp) PAL() byte {
	return f[5]
}

//Operation Code : 1 => Request, 2 => Reply , 3 RARP request , 4 RARP reply
func (f Arp) OP() []byte {
	return f[6:8]
}

func (f Arp) SetOP(b []byte) {
	copy(f[6:8], b)
}

//Sender Hardware Address
func (f Arp) S_HA() net.HardwareAddr {
	return net.HardwareAddr(f[8:14])
}
func (f Arp) SetS_HA(b []byte) {
	copy(f[8:14], b)
}

//Sender L32  (* same as Sender IPv4 address for ARP)
func (f Arp) S_L32() net.IP {
	return net.IP(f[14:18])
}
func (f Arp) SetS_L32(b []byte) {
	copy(f[14:18], b)
}

//Target Hardware Address
func (f Arp) T_HA() net.HardwareAddr {
	return net.HardwareAddr(f[18:24])
}
func (f Arp) SetsT_HA(b []byte) {
	copy(f[18:24], b)
}

//Protocol address of target
func (f Arp) T_L32() net.IP {
	return net.IP(f[24:28])
}

func (f Arp) SetT_L32(b []byte) {
	copy(f[24:28], b)
}

func handle_arp(ifce *water.Interface, frame ethernet.Frame) {
	//leave this here for now
	var tapnet, _ = net.InterfaceByName("O_O")

	var arp Arp = frame.Payload()

	//is the request for us ? ( stupid but enough for now)
	//if arp.T_L32().Equal(netip) {
	//	log.Println("ARP TO US")
	//}

	//cache the information , so we know that a speicifc IP/HA exists ?
	log.Println("ARP src :", arp.S_L32(), arp.S_HA())
	log.Println("ARP dst :", arp.T_L32(), arp.T_HA())

	var response ethernet.Frame

	//28 is the size of ARP in the case of IPv4/eth <= we need to change this later
	response.Prepare(arp.S_HA(), tapnet.HardwareAddr, 0, ethernet.ARP, 28)

	// make the arp packet a reply
	arp.SetOP([]byte{0, 2})

	//flip the proto address
	{
		var s_ha = arp.S_HA()
		arp.SetS_HA(tapnet.HardwareAddr)
		arp.SetsT_HA(s_ha)
	}

	//flip mac address
	{
		var s_l = arp.S_L32()
		var t_l = arp.T_L32()
		arp.SetS_L32(t_l)
		arp.SetT_L32(s_l)
	}

	//set the payload
	copy(response[14:], arp)
	ifce.Write(response)
}
