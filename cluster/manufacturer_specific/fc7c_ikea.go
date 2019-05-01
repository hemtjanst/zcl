// IKEA control outlet cluster.
package manufacturer_specific

import (
	"neotor.se/zcl"
)

// Ikea
// Manufacturer code 0x117C
const IkeaID zcl.ClusterID = 64636

var IkeaCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		Unknown1Attr: func() zcl.Attr { return new(Unknown1) },
		Unknown2Attr: func() zcl.Attr { return new(Unknown2) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

const Unknown1Attr zcl.AttrID = 16

type Unknown1 zcl.Zu8

func (a Unknown1) ID() zcl.AttrID         { return Unknown1Attr }
func (a Unknown1) Cluster() zcl.ClusterID { return IkeaID }
func (a *Unknown1) Value() *Unknown1      { return a }
func (a Unknown1) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *Unknown1) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Unknown1(*nt)
	return br, err
}

func (a Unknown1) Readable() bool   { return true }
func (a Unknown1) Writable() bool   { return true }
func (a Unknown1) Reportable() bool { return false }
func (a Unknown1) SceneIndex() int  { return -1 }

func (a Unknown1) String() string {

	return zcl.Sprintf("0x%X", zcl.Zu8(a))
}

const Unknown2Attr zcl.AttrID = 65533

type Unknown2 zcl.Zu16

func (a Unknown2) ID() zcl.AttrID         { return Unknown2Attr }
func (a Unknown2) Cluster() zcl.ClusterID { return IkeaID }
func (a *Unknown2) Value() *Unknown2      { return a }
func (a Unknown2) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Unknown2) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Unknown2(*nt)
	return br, err
}

func (a Unknown2) Readable() bool   { return true }
func (a Unknown2) Writable() bool   { return true }
func (a Unknown2) Reportable() bool { return false }
func (a Unknown2) SceneIndex() int  { return -1 }

func (a Unknown2) String() string {

	return zcl.Sprintf("0x%X", zcl.Zu16(a))
}
