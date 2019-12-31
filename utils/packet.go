package utils

import "hemtjan.st/zcl"

func Packet(srcEp uint8, dstAddr zcl.Address, dstEp uint8, profile uint16, cluster uint16, data []byte) zcl.Packet {
	return packet{
		srcEp,
		dstAddr,
		dstEp,
		profile,
		cluster,
		data,
	}
}

type packet struct {
	srcEp   uint8
	dstAddr zcl.Address
	dstEp   uint8
	profile uint16
	cluster uint16
	data    []byte
}

func (p packet) SrcEP() uint8         { return p.srcEp }
func (p packet) DstEP() uint8         { return p.dstEp }
func (p packet) DstAddr() zcl.Address { return p.dstAddr }
func (p packet) Profile() uint16      { return p.profile }
func (p packet) Cluster() uint16      { return p.cluster }
func (p packet) Data() []byte         { return p.data }
