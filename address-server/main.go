package main

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
)

const (
	// This is the default port for client and server; if running locally, change the game to use `hostPort=2301`
	AddressServerPort = 2300
	UdpBufferSize     = 1200
)

func main() {
	log.StandardLogger().Formatter = &log.TextFormatter{
		FullTimestamp: true,
	}
	log.WithField("port", AddressServerPort).Infoln("Starting address server")

	udpAddress, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", "0.0.0.0", AddressServerPort))
	if err != nil {
		log.WithError(err).Panic("Failed to resolve udp address")
	}

	udpConn, err := net.ListenUDP("udp", udpAddress)
	if err != nil {
		log.WithError(err).Panic("Failed to listen for UDP packets.")
	}

	defer udpConn.Close()

	udpBuffer := make([]byte, UdpBufferSize)

	for {
		bytes, udpClientAddress, err := udpConn.ReadFromUDP(udpBuffer)
		if err != nil {
			log.WithError(err).Warn("Failed to read UDP packet.")
			continue
		}

		if bytes == 0 {
			log.Warn("Failed to read any data from the UDP packet.")
			continue
		}

		switch udpBuffer[0] {
		case 9:
			reply := make([]byte, 16)
			reply[0] = 9                                    // reply with the same first byte
			copy(reply[1:], []byte{0x02, 0x00, 0xd9, 0x49}) // these 4 are unknown, but it seems to work
			copy(reply[5:], udpClientAddress.IP)
			// from [9:12], all 0x00s
			copy(reply[13:], udpBuffer[1:4])

			log.WithField("reply", reply).WithField("ip", udpClientAddress.IP).Info("Replying to request.")
			udpConn.WriteToUDP(reply, udpClientAddress)
		case 10:
			// TODO: reply back with the same message
		}

		log.WithField("fist_byte", udpBuffer[0]).WithField("ip", udpClientAddress.IP).Info("Got first byte")
	}
}
