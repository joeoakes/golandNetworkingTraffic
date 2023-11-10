package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
)

func main() {

	devices, err := pcap.FindAllDevs()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, device := range devices {
		fmt.Println("Name:", device.Name)
		fmt.Println("Description:", device.Description)
		fmt.Println()
	}
	// Define the network interface you want to capture traffic from (e.g., "eth0")
	device := "\\Device\\NPF_Loopback "

	// Open the network device for packet capture
	handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set a BPF (Berkeley Packet Filter) filter to capture specific traffic (optional)
	filter := "tcp and port 80" // Example: Capture HTTP traffic
	if err := handle.SetBPFFilter(filter); err != nil {
		log.Fatal(err)
	}

	// Start capturing and processing packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process the captured packet
		fmt.Println(packet.String())
	}
}
