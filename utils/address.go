package utils

import (
	"fmt"
	"hemtjan.st/zcl"
	"net"
)

type IEEEAddress []byte
type NWKAddress uint16
type GroupAddress uint16

func (a IEEEAddress) Mode() zcl.AddressMode   { return zcl.IEEEAddress }
func (a IEEEAddress) NWK() uint16             { return 0 }
func (a IEEEAddress) IEEE() net.HardwareAddr  { return net.HardwareAddr(a) }
func (a IEEEAddress) String() string          { return a.IEEE().String() }
func (a NWKAddress) Mode() zcl.AddressMode    { return zcl.NWKAddress }
func (a NWKAddress) NWK() uint16              { return uint16(a) }
func (a NWKAddress) IEEE() net.HardwareAddr   { return make([]byte, 8) }
func (a NWKAddress) String() string           { return fmt.Sprintf("NWK(%04X)", uint16(a)) }
func (a GroupAddress) Mode() zcl.AddressMode  { return zcl.GroupAddress }
func (a GroupAddress) NWK() uint16            { return uint16(a) }
func (a GroupAddress) IEEE() net.HardwareAddr { return make([]byte, 8) }
func (a GroupAddress) String() string         { return fmt.Sprintf("Group(%04X)", uint16(a)) }

const (
	BroadcastAll            = NWKAddress(zcl.BroadcastAll)
	BroadcastRxOnWhenIdle   = NWKAddress(zcl.BroadcastRxOnWhenIdle)
	BroadcastRoutersCoords  = NWKAddress(zcl.BroadcastRoutersCoords)
	BroadcastLowPowerRouter = NWKAddress(zcl.BroadcastLowPowerRouter)
)

type full struct {
	IEEEAddress
	NWKAddress
}

func (a full) Mode() zcl.AddressMode  { return zcl.FullAddress }
func (a full) NWK() uint16            { return a.NWKAddress.NWK() }
func (a full) IEEE() net.HardwareAddr { return a.IEEEAddress.IEEE() }
func (a full) String() string         { return a.IEEE().String() }

func FullAddress(ieee net.HardwareAddr, nwk uint16) zcl.Address {
	return full{IEEEAddress(ieee), NWKAddress(nwk)}
}
