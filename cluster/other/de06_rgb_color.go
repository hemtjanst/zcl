// Attributes and commands for setting devices light color. The color is specified in the RGB range from 0 - 255.
package other

import (
	"neotor.se/zcl/cluster/zcl"
)

// RgbColor
// Attributes and commands for setting devices light color. The color is specified in the RGB range from 0 - 255.

func NewRgbColorServer(profile zcl.ProfileID) *RgbColorServer { return &RgbColorServer{p: profile} }
func NewRgbColorClient(profile zcl.ProfileID) *RgbColorClient { return &RgbColorClient{p: profile} }

const RgbColorCluster zcl.ClusterID = 56838

type RgbColorServer struct {
	p zcl.ProfileID

	Currentcolorset *Currentcolorset
	Colorsetcount   *Colorsetcount
}

func (s *RgbColorServer) SetColor() *SetColor { return new(SetColor) }

type RgbColorClient struct {
	p zcl.ProfileID
}

/*
var RgbColorServer = map[zcl.CommandID]func() zcl.Command{
    SetColorID: func() zcl.Command { return new(SetColor) },
}

var RgbColorClient = map[zcl.CommandID]func() zcl.Command{
}
*/

// On receipt of this command, the color of the light shall be changed and the current index updatet.
type SetColor struct {
	Red      zcl.Zu8
	Green    zcl.Zu8
	Blue     zcl.Zu8
	SetIndex zcl.Zu8
	Options  zcl.Zbmp8
}

const SetColorCommand zcl.CommandID = 0

func (v *SetColor) Values() []zcl.Val {
	return []zcl.Val{
		&v.Red,
		&v.Green,
		&v.Blue,
		&v.SetIndex,
		&v.Options,
	}
}

func (v SetColor) ID() zcl.CommandID {
	return SetColorCommand
}

func (v SetColor) Cluster() zcl.ClusterID {
	return RgbColorCluster
}

func (v SetColor) MnfCode() []byte {
	return []byte{}
}

func (v SetColor) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Red.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Green.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Blue.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.SetIndex.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Options.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *SetColor) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Red).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Green).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Blue).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SetIndex).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Options).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type Currentcolorset zcl.Zu32

func (a Currentcolorset) ID() zcl.AttrID           { return 0 }
func (a Currentcolorset) Cluster() zcl.ClusterID   { return RgbColorCluster }
func (a *Currentcolorset) Value() *Currentcolorset { return a }
func (a Currentcolorset) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *Currentcolorset) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = Currentcolorset(*nt)
	return br, err
}

func (a Currentcolorset) Readable() bool   { return true }
func (a Currentcolorset) Writable() bool   { return false }
func (a Currentcolorset) Reportable() bool { return false }
func (a Currentcolorset) SceneIndex() int  { return -1 }

func (a Currentcolorset) String() string {

	return zcl.Sprintf("0x%X", zcl.Zu32(a))
}

type Colorsetcount zcl.Zu8

func (a Colorsetcount) ID() zcl.AttrID         { return 1 }
func (a Colorsetcount) Cluster() zcl.ClusterID { return RgbColorCluster }
func (a *Colorsetcount) Value() *Colorsetcount { return a }
func (a Colorsetcount) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *Colorsetcount) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Colorsetcount(*nt)
	return br, err
}

func (a Colorsetcount) Readable() bool   { return true }
func (a Colorsetcount) Writable() bool   { return false }
func (a Colorsetcount) Reportable() bool { return false }
func (a Colorsetcount) SceneIndex() int  { return -1 }

func (a Colorsetcount) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}
