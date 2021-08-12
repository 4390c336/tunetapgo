package main

import (
	"log"

	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
)

func main() {

	config := water.Config{
		DeviceType: water.TAP,
	}

	config.Name = "piw"

	ifce, err := water.New(config)

	ifce.IsTAP()

	if err != nil {
		log.Fatal("ERRRRRRRR:", err)
	}

	var frame ethernet.Frame

	for {
		frame.Resize(1500)

		n, err := ifce.Read([]byte(frame))
		if err != nil {
			log.Fatal(err)
		}

		//reset the size
		frame = frame[:n]

		switch frame.Ethertype() {
		case ethernet.ARP:
			go handle_arp(ifce, frame)
		case ethernet.IPv4:
			go handle_ipv4(ifce, frame)
		default:
			log.Println("[UNKOWN Ethertype]", frame.Ethertype())
		}
	}
}
