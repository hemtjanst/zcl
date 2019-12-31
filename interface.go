package zcl

import (
	"fmt"
	"net"
)

type Device interface {
}

type Transport interface {
	Send(Packet) error
}

type Handler interface {
	Handle(ReceivedPacket)
}

type Packet interface {
	SrcEP() uint8
	DstEP() uint8
	DstAddr() Address
	Profile() uint16
	Cluster() uint16
	Data() []byte
}

type ReceivedPacket interface {
	Packet
	SrcAddr() Address
	LQI() uint8
	RSSI() int8
}

type Address interface {
	Mode() AddressMode
	NWK() uint16
	IEEE() net.HardwareAddr
}

type AddressMode uint8

const (
	GroupAddress AddressMode = 0x01
	NWKAddress   AddressMode = 0x02
	IEEEAddress  AddressMode = 0x03

	BroadcastAll            uint16 = 0xFFFF
	BroadcastRxOnWhenIdle   uint16 = 0xFFFD
	BroadcastRoutersCoords  uint16 = 0xFFFC
	BroadcastLowPowerRouter uint16 = 0xFFFB
)

func (a AddressMode) String() string {
	switch a {
	case GroupAddress:
		return "Group"
	case NWKAddress:
		return "NWK"
	case IEEEAddress:
		return "IEEE"
	}
	return fmt.Sprintf("AddrMode%d", int(a))
}
