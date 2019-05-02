// Attributes and commands.
package manufacturer_specific

import (
	"neotor.se/zcl"
)

// UbisysDeviceSetup
// Manufacturer code 0x10F2
const UbisysDeviceSetupID zcl.ClusterID = 64512

var UbisysDeviceSetupCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		InputConfigurationsAttr: func() zcl.Attr { return new(InputConfigurations) },
		InputActionsAttr:        func() zcl.Attr { return new(InputActions) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

const InputConfigurationsAttr zcl.AttrID = 0

type InputConfigurations zcl.Zarray

func (a InputConfigurations) ID() zcl.AttrID               { return InputConfigurationsAttr }
func (a InputConfigurations) Cluster() zcl.ClusterID       { return UbisysDeviceSetupID }
func (a *InputConfigurations) Value() *InputConfigurations { return a }
func (a InputConfigurations) MarshalZcl() ([]byte, error) {
	return zcl.Zarray(a).MarshalZcl()
}
func (a *InputConfigurations) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zarray)
	br, err := nt.UnmarshalZcl(b)
	*a = InputConfigurations(*nt)
	return br, err
}

func (a InputConfigurations) Readable() bool   { return true }
func (a InputConfigurations) Writable() bool   { return false }
func (a InputConfigurations) Reportable() bool { return false }
func (a InputConfigurations) SceneIndex() int  { return -1 }

func (a InputConfigurations) String() string {
	return zcl.Sprintf("%s", zcl.Zarray(a))
}

const InputActionsAttr zcl.AttrID = 1

type InputActions zcl.Zarray

func (a InputActions) ID() zcl.AttrID         { return InputActionsAttr }
func (a InputActions) Cluster() zcl.ClusterID { return UbisysDeviceSetupID }
func (a *InputActions) Value() *InputActions  { return a }
func (a InputActions) MarshalZcl() ([]byte, error) {
	return zcl.Zarray(a).MarshalZcl()
}
func (a *InputActions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zarray)
	br, err := nt.UnmarshalZcl(b)
	*a = InputActions(*nt)
	return br, err
}

func (a InputActions) Readable() bool   { return true }
func (a InputActions) Writable() bool   { return false }
func (a InputActions) Reportable() bool { return false }
func (a InputActions) SceneIndex() int  { return -1 }

func (a InputActions) String() string {
	return zcl.Sprintf("%s", zcl.Zarray(a))
}
