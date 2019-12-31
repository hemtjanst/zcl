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
