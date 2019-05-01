// Attributes and commands.
package green_power

import (
	"neotor.se/zcl"
)

// GreenPower
// Attributes and commands.

func NewGreenPowerServer(profile zcl.ProfileID) *GreenPowerServer {
	return &GreenPowerServer{p: profile}
}
func NewGreenPowerClient(profile zcl.ProfileID) *GreenPowerClient {
	return &GreenPowerClient{p: profile}
}

const GreenPowerCluster zcl.ClusterID = 33

type GreenPowerServer struct {
	p zcl.ProfileID

	MaxSinkTableEntries   *MaxSinkTableEntries
	SinkTable             *SinkTable
	CommunicationMode     *CommunicationMode
	CommissioningExitMode *CommissioningExitMode
	CommissioningWindow   *CommissioningWindow
	SecurityLevel         *SecurityLevel
	Functionality         *Functionality
	ActiveFunctionality   *ActiveFunctionality
}

type GreenPowerClient struct {
	p zcl.ProfileID
}

/*
var GreenPowerServer = map[zcl.CommandID]func() zcl.Command{
}

var GreenPowerClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type MaxSinkTableEntries zcl.Zu8

func (a MaxSinkTableEntries) ID() zcl.AttrID               { return 0 }
func (a MaxSinkTableEntries) Cluster() zcl.ClusterID       { return GreenPowerCluster }
func (a *MaxSinkTableEntries) Value() *MaxSinkTableEntries { return a }
func (a MaxSinkTableEntries) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *MaxSinkTableEntries) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxSinkTableEntries(*nt)
	return br, err
}

func (a MaxSinkTableEntries) Readable() bool   { return true }
func (a MaxSinkTableEntries) Writable() bool   { return false }
func (a MaxSinkTableEntries) Reportable() bool { return false }
func (a MaxSinkTableEntries) SceneIndex() int  { return -1 }

func (a MaxSinkTableEntries) String() string {

	return zcl.Sprintf("0x%X", zcl.Zu8(a))
}

type SinkTable zcl.Zlostring

func (a SinkTable) ID() zcl.AttrID         { return 1 }
func (a SinkTable) Cluster() zcl.ClusterID { return GreenPowerCluster }
func (a *SinkTable) Value() *SinkTable     { return a }
func (a SinkTable) MarshalZcl() ([]byte, error) {
	return zcl.Zlostring(a).MarshalZcl()
}
func (a *SinkTable) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zlostring)
	br, err := nt.UnmarshalZcl(b)
	*a = SinkTable(*nt)
	return br, err
}

func (a SinkTable) Readable() bool   { return true }
func (a SinkTable) Writable() bool   { return false }
func (a SinkTable) Reportable() bool { return false }
func (a SinkTable) SceneIndex() int  { return -1 }

func (a SinkTable) String() string {

	return zcl.Sprintf("0x%X", zcl.Zlostring(a))
}

type CommunicationMode zcl.Zbmp8

func (a CommunicationMode) ID() zcl.AttrID             { return 2 }
func (a CommunicationMode) Cluster() zcl.ClusterID     { return GreenPowerCluster }
func (a *CommunicationMode) Value() *CommunicationMode { return a }
func (a CommunicationMode) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *CommunicationMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = CommunicationMode(*nt)
	return br, err
}

func (a CommunicationMode) Readable() bool   { return true }
func (a CommunicationMode) Writable() bool   { return true }
func (a CommunicationMode) Reportable() bool { return false }
func (a CommunicationMode) SceneIndex() int  { return -1 }

func (a CommunicationMode) String() string {

	return zcl.Sprintf("0x%X", zcl.Zbmp8(a))
}

type CommissioningExitMode zcl.Zbmp8

func (a CommissioningExitMode) ID() zcl.AttrID                 { return 3 }
func (a CommissioningExitMode) Cluster() zcl.ClusterID         { return GreenPowerCluster }
func (a *CommissioningExitMode) Value() *CommissioningExitMode { return a }
func (a CommissioningExitMode) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *CommissioningExitMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = CommissioningExitMode(*nt)
	return br, err
}

func (a CommissioningExitMode) Readable() bool   { return true }
func (a CommissioningExitMode) Writable() bool   { return true }
func (a CommissioningExitMode) Reportable() bool { return false }
func (a CommissioningExitMode) SceneIndex() int  { return -1 }

func (a CommissioningExitMode) String() string {

	return zcl.Sprintf("0x%X", zcl.Zbmp8(a))
}

type CommissioningWindow zcl.Zu16

func (a CommissioningWindow) ID() zcl.AttrID               { return 4 }
func (a CommissioningWindow) Cluster() zcl.ClusterID       { return GreenPowerCluster }
func (a *CommissioningWindow) Value() *CommissioningWindow { return a }
func (a CommissioningWindow) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *CommissioningWindow) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = CommissioningWindow(*nt)
	return br, err
}

func (a CommissioningWindow) Readable() bool   { return true }
func (a CommissioningWindow) Writable() bool   { return true }
func (a CommissioningWindow) Reportable() bool { return false }
func (a CommissioningWindow) SceneIndex() int  { return -1 }

func (a CommissioningWindow) String() string {

	return zcl.Sprintf("0x%X", zcl.Zu16(a))
}

type SecurityLevel zcl.Zbmp8

func (a SecurityLevel) ID() zcl.AttrID         { return 5 }
func (a SecurityLevel) Cluster() zcl.ClusterID { return GreenPowerCluster }
func (a *SecurityLevel) Value() *SecurityLevel { return a }
func (a SecurityLevel) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *SecurityLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = SecurityLevel(*nt)
	return br, err
}

func (a SecurityLevel) Readable() bool   { return true }
func (a SecurityLevel) Writable() bool   { return true }
func (a SecurityLevel) Reportable() bool { return false }
func (a SecurityLevel) SceneIndex() int  { return -1 }

func (a SecurityLevel) String() string {

	return zcl.Sprintf("0x%X", zcl.Zbmp8(a))
}

type Functionality zcl.Zbmp24

func (a Functionality) ID() zcl.AttrID         { return 6 }
func (a Functionality) Cluster() zcl.ClusterID { return GreenPowerCluster }
func (a *Functionality) Value() *Functionality { return a }
func (a Functionality) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp24(a).MarshalZcl()
}
func (a *Functionality) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp24)
	br, err := nt.UnmarshalZcl(b)
	*a = Functionality(*nt)
	return br, err
}

func (a Functionality) Readable() bool   { return true }
func (a Functionality) Writable() bool   { return false }
func (a Functionality) Reportable() bool { return false }
func (a Functionality) SceneIndex() int  { return -1 }

func (a Functionality) String() string {

	return zcl.Sprintf("0x%X", zcl.Zbmp24(a))
}

type ActiveFunctionality zcl.Zbmp24

func (a ActiveFunctionality) ID() zcl.AttrID               { return 7 }
func (a ActiveFunctionality) Cluster() zcl.ClusterID       { return GreenPowerCluster }
func (a *ActiveFunctionality) Value() *ActiveFunctionality { return a }
func (a ActiveFunctionality) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp24(a).MarshalZcl()
}
func (a *ActiveFunctionality) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp24)
	br, err := nt.UnmarshalZcl(b)
	*a = ActiveFunctionality(*nt)
	return br, err
}

func (a ActiveFunctionality) Readable() bool   { return true }
func (a ActiveFunctionality) Writable() bool   { return false }
func (a ActiveFunctionality) Reportable() bool { return false }
func (a ActiveFunctionality) SceneIndex() int  { return -1 }

func (a ActiveFunctionality) String() string {

	return zcl.Sprintf("0x%X", zcl.Zbmp24(a))
}
