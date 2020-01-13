package utils

import "hemtjan.st/zcl"

func Packet(srcEp uint8, dstAddr zcl.Address, dstEp uint8, profile zcl.ProfileID, cluster zcl.ClusterID, data []byte) zcl.Packet {
	p, _ := zcl.NewPacket(srcEp, dstAddr, dstEp, profile, cluster, zcl.Zbytes(data))
	return p
}

type packet struct {
	srcEp   uint8
	dstAddr zcl.Address
	dstEp   uint8
	profile zcl.ProfileID
	cluster zcl.ClusterID
	data    []byte
}

func (p packet) SrcEP() uint8           { return p.srcEp }
func (p packet) DstEP() uint8           { return p.dstEp }
func (p packet) DstAddr() zcl.Address   { return p.dstAddr }
func (p packet) Profile() zcl.ProfileID { return p.profile }
func (p packet) Cluster() zcl.ClusterID { return p.cluster }
func (p packet) Data() []byte           { return p.data }
