// Legrand Specific clusters.
package manufacturer_specific

import (
	"neotor.se/zcl"
)

// LegrandSpecificClusters
// Manufacturer code 0x1021
const LegrandSpecificClustersID zcl.ClusterID = 64513

var LegrandSpecificClustersCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		DimmerAttr: func() zcl.Attr { return new(Dimmer) },
		LedAttr:    func() zcl.Attr { return new(Led) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

const DimmerAttr zcl.AttrID = 0

type Dimmer zcl.Zdat16

func (a Dimmer) ID() zcl.AttrID         { return DimmerAttr }
func (a Dimmer) Cluster() zcl.ClusterID { return LegrandSpecificClustersID }
func (a *Dimmer) Value() *Dimmer        { return a }
func (a Dimmer) MarshalZcl() ([]byte, error) {
	return zcl.Zdat16(a).MarshalZcl()
}
func (a *Dimmer) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zdat16)
	br, err := nt.UnmarshalZcl(b)
	*a = Dimmer(*nt)
	return br, err
}

func (a Dimmer) Readable() bool   { return true }
func (a Dimmer) Writable() bool   { return true }
func (a Dimmer) Reportable() bool { return false }
func (a Dimmer) SceneIndex() int  { return -1 }

func (a Dimmer) String() string {
	return zcl.Sprintf("0x%X", zcl.Zdat16(a))
}

const LedAttr zcl.AttrID = 1

type Led zcl.Zbool

func (a Led) ID() zcl.AttrID         { return LedAttr }
func (a Led) Cluster() zcl.ClusterID { return LegrandSpecificClustersID }
func (a *Led) Value() *Led           { return a }
func (a Led) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *Led) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = Led(*nt)
	return br, err
}

func (a Led) Readable() bool   { return true }
func (a Led) Writable() bool   { return true }
func (a Led) Reportable() bool { return false }
func (a Led) SceneIndex() int  { return -1 }

func (a Led) String() string {
	return zcl.Sprintf("%s", zcl.Zbool(a))
}
