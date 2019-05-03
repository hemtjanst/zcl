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

// Dimmer is an autogenerated attribute in the LegrandSpecificClusters cluster
// 0100 = Dimmer Off, 0101 = Dimmer On
type Dimmer zcl.Zdat16

const DimmerAttr zcl.AttrID = 0

func (Dimmer) ID() zcl.AttrID                { return DimmerAttr }
func (Dimmer) Cluster() zcl.ClusterID        { return LegrandSpecificClustersID }
func (Dimmer) Name() string                  { return "Dimmer" }
func (Dimmer) Readable() bool                { return true }
func (Dimmer) Writable() bool                { return true }
func (Dimmer) Reportable() bool              { return false }
func (Dimmer) SceneIndex() int               { return -1 }
func (a *Dimmer) Value() *Dimmer             { return a }
func (a Dimmer) MarshalZcl() ([]byte, error) { return zcl.Zdat16(a).MarshalZcl() }

func (a *Dimmer) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zdat16)
	br, err := nt.UnmarshalZcl(b)
	*a = Dimmer(*nt)
	return br, err
}

func (a Dimmer) String() string {
	return zcl.Sprintf("%v", zcl.Zdat16(a))
}

// Led is an autogenerated attribute in the LegrandSpecificClusters cluster
// Enable LED in dark
type Led zcl.Zbool

const LedAttr zcl.AttrID = 1

func (Led) ID() zcl.AttrID                { return LedAttr }
func (Led) Cluster() zcl.ClusterID        { return LegrandSpecificClustersID }
func (Led) Name() string                  { return "LED" }
func (Led) Readable() bool                { return true }
func (Led) Writable() bool                { return true }
func (Led) Reportable() bool              { return false }
func (Led) SceneIndex() int               { return -1 }
func (a *Led) Value() *Led                { return a }
func (a Led) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *Led) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = Led(*nt)
	return br, err
}

func (a Led) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(a))
}
