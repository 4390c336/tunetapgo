package main

import "encoding/binary"

//
type Dhcp []byte

//OP
func (f Dhcp) OP() byte {
	return f[0]
}

//HTYPE
func (f Dhcp) HTYPE() byte {
	//HTYPE = 1 => Ethernet
	return f[1]
}

//HLEN
func (f Dhcp) HLEN() byte {
	//HLEN = 6 => ipv4
	return f[2]
}

//HOPS
func (f Dhcp) HOPS() byte {
	//Client sets to zero, optionally used by relay agents
	return f[3]
}

//XID
func (f Dhcp) XID() uint32 {
	return binary.BigEndian.Uint32(f[4:8])
}

//SECS
func (f Dhcp) SECS() []byte {
	return f[8:10]
}

//FLAGS
func (f Dhcp) FLAGS() []byte {
	return f[10:12]
}
