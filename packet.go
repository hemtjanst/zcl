package zcl

import "fmt"

func NewPacket(srcEp uint8, dstAddr Address, dstEp uint8, profile ProfileID, cluster ClusterID, data ...InArg) (p Packet, err error) {
	var rawData []byte
	for _, v := range data {
		if rawData, err = v.MarshalZcl(); err != nil {
			return
		}
	}
	return &packet{
		srcEp:   srcEp,
		dstAddr: dstAddr,
		dstEp:   dstEp,
		profile: profile,
		cluster: cluster,
		data:    rawData,
	}, nil
}

type packet struct {
	srcEp   uint8
	dstAddr Address
	dstEp   uint8
	profile ProfileID
	cluster ClusterID
	data    []byte
}

func (p packet) SrcEP() uint8       { return p.srcEp }
func (p packet) DstEP() uint8       { return p.dstEp }
func (p packet) DstAddr() Address   { return p.dstAddr }
func (p packet) Profile() ProfileID { return p.profile }
func (p packet) Cluster() ClusterID { return p.cluster }
func (p packet) Data() []byte       { return p.data }
func (p packet) String() string {
	return fmt.Sprintf("[*/%d->%s/%d] P[%04X] C[%04X] Sz[%d]",
		p.SrcEP(),
		p.DstAddr(),
		p.DstEP(),
		uint16(p.Profile()),
		uint16(p.Cluster()),
		len(p.data),
	)

}
