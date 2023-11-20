package main

import (
	"net"
	"time"

	"github.com/bogushevich/godivert"
)

var cloudflareDNS = net.ParseIP("1.1.1.1")

func checkPacket(wd *godivert.WinDivertHandle, packetChan <-chan *godivert.Packet) {
	for packet := range packetChan {
		if !packet.DstIP().Equal(cloudflareDNS) {
			packet.Send(wd)
		}
	}
}

func main() {
	winDivert, err := godivert.OpenHandle("icmp", godivert.LayerNetwork, godivert.PriorityDefault, godivert.OpenFlagNone)
	if err != nil {
		panic(err)
	}
	defer winDivert.Close()

	packetChan, _, err := winDivert.Packets()
	if err != nil {
		panic(err)
	}

	go checkPacket(winDivert, packetChan)

	time.Sleep(1 * time.Minute)
}
