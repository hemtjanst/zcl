// Attributes and commands.
package manufacturer_specific

import (
	"neotor.se/zcl"
)

// UbisysDeviceSetup
// Manufacturer code 0x10F2
// Attributes and commands.

func NewUbisysDeviceSetupServer(profile zcl.ProfileID) *UbisysDeviceSetupServer {
	return &UbisysDeviceSetupServer{p: profile}
}
func NewUbisysDeviceSetupClient(profile zcl.ProfileID) *UbisysDeviceSetupClient {
	return &UbisysDeviceSetupClient{p: profile}
}

const UbisysDeviceSetupCluster zcl.ClusterID = 64512

type UbisysDeviceSetupServer struct {
	p zcl.ProfileID

	InputConfigurations *InputConfigurations
	InputActions        *InputActions
}

type UbisysDeviceSetupClient struct {
	p zcl.ProfileID
}

/*
var UbisysDeviceSetupServer = map[zcl.CommandID]func() zcl.Command{
}

var UbisysDeviceSetupClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type InputConfigurations zcl.Zarray

func (a InputConfigurations) ID() zcl.AttrID               { return 0 }
func (a InputConfigurations) Cluster() zcl.ClusterID       { return UbisysDeviceSetupCluster }
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

type InputActions zcl.Zarray

func (a InputActions) ID() zcl.AttrID         { return 1 }
func (a InputActions) Cluster() zcl.ClusterID { return UbisysDeviceSetupCluster }
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
