package zcl

import (
	"net"
)

type Transport interface {
	Send(Packet) error
	SetPermitJoin(time uint8) error
	NWKAddr() (uint16, error)
	IEEEAddr() (Zuid, error)
}

type InArg interface {
	MarshalZcl() ([]byte, error)
}
type OutArg interface {
	UnmarshalZcl([]byte) ([]byte, error)
}

type Val interface {
	InArg
	OutArg
}

type ScaledArg interface {
	Scaled() float64
}

type EnumArg interface {
	SingleOptions() []Option
}

type BitmapArg interface {
	MultiOptions() []Option
}

type Command interface {
	General
	Arguments() []ArgDesc
	Cluster() ClusterID
	Required() bool
	MnfCode() uint16
}

type General interface {
	Val
	Values() []Val
	ID() CommandID
	Name() string
	String() string
	Direction() Direction
	Handle(frame ReceivedZclFrame, handler interface{}) (response General, found bool, err error)
}

type ZdoCommand interface {
	Val
	Arguments() []ArgDesc
	Cluster() ClusterID
	Values() []Val
	ID() ZdoCmdID
	Name() string
	String() string
	Handle(frame ReceivedZdpFrame, handler interface{}) (response ZdoCommand, found bool, err error)
}

type ZclHandler interface {
	HandleZcl(ReceivedZclFrame) (response ZclFrame, err error)
}

type ZdpHandler interface {
	HandleZdp(ReceivedZdpFrame) (response ZdpFrame, err error)
}

type Handler interface {
	ZclHandler
	ZdpHandler
}

type receivedCommon interface {
	SrcAddr() Address
	LQI() uint8
	RSSI() int8
}

type zdpCommon interface {
	SeqNo() uint8
	Command() ZdoCmdID
	Payload() []byte
}

type zclCommon interface {
	SeqNo() uint8
	CommandType() CommandType
	Direction() Direction
	Manufacturer() uint16
	DisableDefaultResponse() bool
	Command() CommandID
	Payload() []byte
}

type Packet interface {
	SrcEP() uint8
	DstEP() uint8
	DstAddr() Address
	Profile() ProfileID
	Cluster() ClusterID
	Data() []byte
}

type ReceivedPacket interface {
	Packet
	receivedCommon
}

type ZdpFrame interface {
	Packet
	zdpCommon
}

type ReceivedZdpFrame interface {
	ZdpFrame
	receivedCommon
	Response(cmd ZdoCommand) (ZdpFrame, error)
}

type ZclFrame interface {
	Packet
	zclCommon
}

type ReceivedZclFrame interface {
	ZclFrame
	receivedCommon
	Response(srcEp uint8, cmd General) (ZclFrame, error)
}

type Address interface {
	Mode() AddressMode
	NWK() uint16
	IEEE() net.HardwareAddr
}
