// Attributes and commands for commissioning and managing a ZigBee device.
package commissioning

import (
	"neotor.se/zcl"
)

// Commissioning
// Attributes and commands for commissioning and managing a ZigBee device.

func NewCommissioningServer(profile zcl.ProfileID) *CommissioningServer {
	return &CommissioningServer{p: profile}
}
func NewCommissioningClient(profile zcl.ProfileID) *CommissioningClient {
	return &CommissioningClient{p: profile}
}

const CommissioningCluster zcl.ClusterID = 21

type CommissioningServer struct {
	p zcl.ProfileID

	ShortAddress              *ShortAddress
	ExtendedPanId             *ExtendedPanId
	PanId                     *PanId
	ChannelMask               *ChannelMask
	ProtocolVersion           *ProtocolVersion
	StackProfile              *StackProfile
	StartupControl            *StartupControl
	TrustCenterAddress        *TrustCenterAddress
	TrustCenterMasterKey      *TrustCenterMasterKey
	NetworkKey                *NetworkKey
	UseInsecureJoin           *UseInsecureJoin
	PreconfiguredLinkKey      *PreconfiguredLinkKey
	NetworkKeySeqNum          *NetworkKeySeqNum
	NetworkKeyType            *NetworkKeyType
	NetworkManagerAddress     *NetworkManagerAddress
	ScanAttempts              *ScanAttempts
	TimeBetweenScans          *TimeBetweenScans
	RejoinInterval            *RejoinInterval
	MaxRejoinInterval         *MaxRejoinInterval
	IndirectPollRate          *IndirectPollRate
	ParentRetryThreshold      *ParentRetryThreshold
	ConcentratorFlag          *ConcentratorFlag
	ConcentratorRadius        *ConcentratorRadius
	ConcentratorDiscoveryTime *ConcentratorDiscoveryTime
	MacAddress                *MacAddress
}

func (s *CommissioningServer) RestartDevice() *RestartDevice { return new(RestartDevice) }
func (s *CommissioningServer) SaveStartupParameters() *SaveStartupParameters {
	return new(SaveStartupParameters)
}
func (s *CommissioningServer) RestoreStartupParameters() *RestoreStartupParameters {
	return new(RestoreStartupParameters)
}
func (s *CommissioningServer) ResetStartupParameters() *ResetStartupParameters {
	return new(ResetStartupParameters)
}

type CommissioningClient struct {
	p zcl.ProfileID
}

func (s *CommissioningClient) RestartDeviceResponse() *RestartDeviceResponse {
	return new(RestartDeviceResponse)
}
func (s *CommissioningClient) SaveStartupParametersResponse() *SaveStartupParametersResponse {
	return new(SaveStartupParametersResponse)
}
func (s *CommissioningClient) RestoreStartupParametersResponse() *RestoreStartupParametersResponse {
	return new(RestoreStartupParametersResponse)
}
func (s *CommissioningClient) ResetStartupParametersResponse() *ResetStartupParametersResponse {
	return new(ResetStartupParametersResponse)
}

/*
var CommissioningServer = map[zcl.CommandID]func() zcl.Command{
    RestartDeviceID: func() zcl.Command { return new(RestartDevice) },
    SaveStartupParametersID: func() zcl.Command { return new(SaveStartupParameters) },
    RestoreStartupParametersID: func() zcl.Command { return new(RestoreStartupParameters) },
    ResetStartupParametersID: func() zcl.Command { return new(ResetStartupParameters) },
}

var CommissioningClient = map[zcl.CommandID]func() zcl.Command{
    RestartDeviceResponseID: func() zcl.Command { return new(RestartDeviceResponse) },
    SaveStartupParametersResponseID: func() zcl.Command { return new(SaveStartupParametersResponse) },
    RestoreStartupParametersResponseID: func() zcl.Command { return new(RestoreStartupParametersResponse) },
    ResetStartupParametersResponseID: func() zcl.Command { return new(ResetStartupParametersResponse) },
}
*/

// The Restart Device command is used to optionally install a set of startup parameters in a device and run the startup procedure so as to put the new values into effect. The new values may take effect immediately or after an optional delay with optional jitter. The server will send a Restart Device Response command back to the client device before executing the procedure or starting the countdown timer required to time the delay.
type RestartDevice struct {
	Options zcl.Zbmp8
	Delay   zcl.Zu8
	Jitter  zcl.Zu8
}

const RestartDeviceCommand zcl.CommandID = 0

func (v *RestartDevice) Values() []zcl.Val {
	return []zcl.Val{
		&v.Options,
		&v.Delay,
		&v.Jitter,
	}
}

func (v RestartDevice) ID() zcl.CommandID {
	return RestartDeviceCommand
}

func (v RestartDevice) Cluster() zcl.ClusterID {
	return CommissioningCluster
}

func (v RestartDevice) MnfCode() []byte {
	return []byte{}
}

func (v RestartDevice) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Options.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Delay.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Jitter.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *RestartDevice) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Options).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Delay).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Jitter).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// The Save Startup Parameters Request command allows for the current attribute set to be stored under a given index.
type SaveStartupParameters struct {
	Options zcl.Zbmp8
	Index   zcl.Zu8
}

const SaveStartupParametersCommand zcl.CommandID = 1

func (v *SaveStartupParameters) Values() []zcl.Val {
	return []zcl.Val{
		&v.Options,
		&v.Index,
	}
}

func (v SaveStartupParameters) ID() zcl.CommandID {
	return SaveStartupParametersCommand
}

func (v SaveStartupParameters) Cluster() zcl.ClusterID {
	return CommissioningCluster
}

func (v SaveStartupParameters) MnfCode() []byte {
	return []byte{}
}

func (v SaveStartupParameters) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Options.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Index.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *SaveStartupParameters) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Options).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Index).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// This command allows a saved startup parameters attribute set to be restored to current status overwriting whatever was there previously.
type RestoreStartupParameters struct {
	Options zcl.Zbmp8
	Index   zcl.Zu8
}

const RestoreStartupParametersCommand zcl.CommandID = 2

func (v *RestoreStartupParameters) Values() []zcl.Val {
	return []zcl.Val{
		&v.Options,
		&v.Index,
	}
}

func (v RestoreStartupParameters) ID() zcl.CommandID {
	return RestoreStartupParametersCommand
}

func (v RestoreStartupParameters) Cluster() zcl.ClusterID {
	return CommissioningCluster
}

func (v RestoreStartupParameters) MnfCode() []byte {
	return []byte{}
}

func (v RestoreStartupParameters) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Options.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Index.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *RestoreStartupParameters) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Options).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Index).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// This command allows current startup parameters attribute set and one or all of the saved attribute sets to be set to default values. There is also an option for erasing the index under which an attribute set is saved thereby freeing up storage capacity.
type ResetStartupParameters struct {
	Options zcl.Zbmp8
	Index   zcl.Zu8
}

const ResetStartupParametersCommand zcl.CommandID = 3

func (v *ResetStartupParameters) Values() []zcl.Val {
	return []zcl.Val{
		&v.Options,
		&v.Index,
	}
}

func (v ResetStartupParameters) ID() zcl.CommandID {
	return ResetStartupParametersCommand
}

func (v ResetStartupParameters) Cluster() zcl.ClusterID {
	return CommissioningCluster
}

func (v ResetStartupParameters) MnfCode() []byte {
	return []byte{}
}

func (v ResetStartupParameters) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Options.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Index.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *ResetStartupParameters) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Options).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Index).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// On receipt of this command the client is made aware that the server has received the corresponding request and is informed of the status of the request.
type RestartDeviceResponse struct {
	Status zcl.Status
}

const RestartDeviceResponseCommand zcl.CommandID = 0

func (v *RestartDeviceResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
	}
}

func (v RestartDeviceResponse) ID() zcl.CommandID {
	return RestartDeviceResponseCommand
}

func (v RestartDeviceResponse) Cluster() zcl.ClusterID {
	return CommissioningCluster
}

func (v RestartDeviceResponse) MnfCode() []byte {
	return []byte{}
}

func (v RestartDeviceResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *RestartDeviceResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// On receipt of this command the client is made aware that the server has received the corresponding request and is informed of the status of the request.
type SaveStartupParametersResponse struct {
	Status zcl.Status
}

const SaveStartupParametersResponseCommand zcl.CommandID = 1

func (v *SaveStartupParametersResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
	}
}

func (v SaveStartupParametersResponse) ID() zcl.CommandID {
	return SaveStartupParametersResponseCommand
}

func (v SaveStartupParametersResponse) Cluster() zcl.ClusterID {
	return CommissioningCluster
}

func (v SaveStartupParametersResponse) MnfCode() []byte {
	return []byte{}
}

func (v SaveStartupParametersResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *SaveStartupParametersResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// On receipt of this command the client is made aware that the server has received the corresponding request and is informed of the status of the request.
type RestoreStartupParametersResponse struct {
	Status zcl.Status
}

const RestoreStartupParametersResponseCommand zcl.CommandID = 2

func (v *RestoreStartupParametersResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
	}
}

func (v RestoreStartupParametersResponse) ID() zcl.CommandID {
	return RestoreStartupParametersResponseCommand
}

func (v RestoreStartupParametersResponse) Cluster() zcl.ClusterID {
	return CommissioningCluster
}

func (v RestoreStartupParametersResponse) MnfCode() []byte {
	return []byte{}
}

func (v RestoreStartupParametersResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *RestoreStartupParametersResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// On receipt of this command the client is made aware that the server has received the corresponding request and is informed of the status of the request.
type ResetStartupParametersResponse struct {
	Status zcl.Status
}

const ResetStartupParametersResponseCommand zcl.CommandID = 3

func (v *ResetStartupParametersResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
	}
}

func (v ResetStartupParametersResponse) ID() zcl.CommandID {
	return ResetStartupParametersResponseCommand
}

func (v ResetStartupParametersResponse) Cluster() zcl.ClusterID {
	return CommissioningCluster
}

func (v ResetStartupParametersResponse) MnfCode() []byte {
	return []byte{}
}

func (v ResetStartupParametersResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *ResetStartupParametersResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type ShortAddress zcl.Zu16

func (a ShortAddress) ID() zcl.AttrID         { return 0 }
func (a ShortAddress) Cluster() zcl.ClusterID { return CommissioningCluster }
func (a *ShortAddress) Value() *ShortAddress  { return a }
func (a ShortAddress) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ShortAddress) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ShortAddress(*nt)
	return br, err
}

func (a ShortAddress) Readable() bool   { return true }
func (a ShortAddress) Writable() bool   { return true }
func (a ShortAddress) Reportable() bool { return false }
func (a ShortAddress) SceneIndex() int  { return -1 }

func (a ShortAddress) String() string {

	return zcl.Sprintf("0x%X", zcl.Zu16(a))
}

type ExtendedPanId zcl.Zuid

func (a ExtendedPanId) ID() zcl.AttrID         { return 1 }
func (a ExtendedPanId) Cluster() zcl.ClusterID { return CommissioningCluster }
func (a *ExtendedPanId) Value() *ExtendedPanId { return a }
func (a ExtendedPanId) MarshalZcl() ([]byte, error) {
	return zcl.Zuid(a).MarshalZcl()
}
func (a *ExtendedPanId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*a = ExtendedPanId(*nt)
	return br, err
}

func (a ExtendedPanId) Readable() bool   { return true }
func (a ExtendedPanId) Writable() bool   { return true }
func (a ExtendedPanId) Reportable() bool { return false }
func (a ExtendedPanId) SceneIndex() int  { return -1 }

func (a ExtendedPanId) String() string {

	return zcl.Sprintf("%s", zcl.Zuid(a))
}

type PanId zcl.Zu16

func (a PanId) ID() zcl.AttrID         { return 2 }
func (a PanId) Cluster() zcl.ClusterID { return CommissioningCluster }
func (a *PanId) Value() *PanId         { return a }
func (a PanId) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *PanId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PanId(*nt)
	return br, err
}

func (a PanId) Readable() bool   { return true }
func (a PanId) Writable() bool   { return true }
func (a PanId) Reportable() bool { return false }
func (a PanId) SceneIndex() int  { return -1 }

func (a PanId) String() string {

	return zcl.Sprintf("0x%X", zcl.Zu16(a))
}

type ChannelMask zcl.Zbmp32

func (a ChannelMask) ID() zcl.AttrID         { return 3 }
func (a ChannelMask) Cluster() zcl.ClusterID { return CommissioningCluster }
func (a *ChannelMask) Value() *ChannelMask   { return a }
func (a ChannelMask) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp32(a).MarshalZcl()
}
func (a *ChannelMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp32)
	br, err := nt.UnmarshalZcl(b)
	*a = ChannelMask(*nt)
	return br, err
}

func (a ChannelMask) Readable() bool   { return true }
func (a ChannelMask) Writable() bool   { return true }
func (a ChannelMask) Reportable() bool { return false }
func (a ChannelMask) SceneIndex() int  { return -1 }

func (a ChannelMask) String() string {

	var bstr []string
	if a.IsCh11() {
		bstr = append(bstr, "CH 11")
	}
	if a.IsCh12() {
		bstr = append(bstr, "CH 12")
	}
	if a.IsCh13() {
		bstr = append(bstr, "CH 13")
	}
	if a.IsCh14() {
		bstr = append(bstr, "CH 14")
	}
	if a.IsCh15() {
		bstr = append(bstr, "CH 15")
	}
	if a.IsCh16() {
		bstr = append(bstr, "CH 16")
	}
	if a.IsCh17() {
		bstr = append(bstr, "CH 17")
	}
	if a.IsCh18() {
		bstr = append(bstr, "CH 18")
	}
	if a.IsCh19() {
		bstr = append(bstr, "CH 19")
	}
	if a.IsCh20() {
		bstr = append(bstr, "CH 20")
	}
	if a.IsCh21() {
		bstr = append(bstr, "CH 21")
	}
	if a.IsCh22() {
		bstr = append(bstr, "CH 22")
	}
	if a.IsCh23() {
		bstr = append(bstr, "CH 23")
	}
	if a.IsCh24() {
		bstr = append(bstr, "CH 24")
	}
	if a.IsCh25() {
		bstr = append(bstr, "CH 25")
	}
	if a.IsCh26() {
		bstr = append(bstr, "CH 26")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a ChannelMask) IsCh11() bool {
	return zcl.BitmapTest([]byte(a), 11)
}
func (a *ChannelMask) SetCh11(b bool) {
	*a = ChannelMask(zcl.BitmapSet([]byte(*a), 11, b))
}

func (a ChannelMask) IsCh12() bool {
	return zcl.BitmapTest([]byte(a), 12)
}
func (a *ChannelMask) SetCh12(b bool) {
	*a = ChannelMask(zcl.BitmapSet([]byte(*a), 12, b))
}

func (a ChannelMask) IsCh13() bool {
	return zcl.BitmapTest([]byte(a), 13)
}
func (a *ChannelMask) SetCh13(b bool) {
	*a = ChannelMask(zcl.BitmapSet([]byte(*a), 13, b))
}

func (a ChannelMask) IsCh14() bool {
	return zcl.BitmapTest([]byte(a), 14)
}
func (a *ChannelMask) SetCh14(b bool) {
	*a = ChannelMask(zcl.BitmapSet([]byte(*a), 14, b))
}

func (a ChannelMask) IsCh15() bool {
	return zcl.BitmapTest([]byte(a), 15)
}
func (a *ChannelMask) SetCh15(b bool) {
	*a = ChannelMask(zcl.BitmapSet([]byte(*a), 15, b))
}

func (a ChannelMask) IsCh16() bool {
	return zcl.BitmapTest([]byte(a), 16)
}
func (a *ChannelMask) SetCh16(b bool) {
	*a = ChannelMask(zcl.BitmapSet([]byte(*a), 16, b))
}

func (a ChannelMask) IsCh17() bool {
	return zcl.BitmapTest([]byte(a), 17)
}
func (a *ChannelMask) SetCh17(b bool) {
	*a = ChannelMask(zcl.BitmapSet([]byte(*a), 17, b))
}

func (a ChannelMask) IsCh18() bool {
	return zcl.BitmapTest([]byte(a), 18)
}
func (a *ChannelMask) SetCh18(b bool) {
	*a = ChannelMask(zcl.BitmapSet([]byte(*a), 18, b))
}

func (a ChannelMask) IsCh19() bool {
	return zcl.BitmapTest([]byte(a), 19)
}
func (a *ChannelMask) SetCh19(b bool) {
	*a = ChannelMask(zcl.BitmapSet([]byte(*a), 19, b))
}

func (a ChannelMask) IsCh20() bool {
	return zcl.BitmapTest([]byte(a), 20)
}
func (a *ChannelMask) SetCh20(b bool) {
	*a = ChannelMask(zcl.BitmapSet([]byte(*a), 20, b))
}

func (a ChannelMask) IsCh21() bool {
	return zcl.BitmapTest([]byte(a), 21)
}
func (a *ChannelMask) SetCh21(b bool) {
	*a = ChannelMask(zcl.BitmapSet([]byte(*a), 21, b))
}

func (a ChannelMask) IsCh22() bool {
	return zcl.BitmapTest([]byte(a), 22)
}
func (a *ChannelMask) SetCh22(b bool) {
	*a = ChannelMask(zcl.BitmapSet([]byte(*a), 22, b))
}

func (a ChannelMask) IsCh23() bool {
	return zcl.BitmapTest([]byte(a), 23)
}
func (a *ChannelMask) SetCh23(b bool) {
	*a = ChannelMask(zcl.BitmapSet([]byte(*a), 23, b))
}

func (a ChannelMask) IsCh24() bool {
	return zcl.BitmapTest([]byte(a), 24)
}
func (a *ChannelMask) SetCh24(b bool) {
	*a = ChannelMask(zcl.BitmapSet([]byte(*a), 24, b))
}

func (a ChannelMask) IsCh25() bool {
	return zcl.BitmapTest([]byte(a), 25)
}
func (a *ChannelMask) SetCh25(b bool) {
	*a = ChannelMask(zcl.BitmapSet([]byte(*a), 25, b))
}

func (a ChannelMask) IsCh26() bool {
	return zcl.BitmapTest([]byte(a), 26)
}
func (a *ChannelMask) SetCh26(b bool) {
	*a = ChannelMask(zcl.BitmapSet([]byte(*a), 26, b))
}

type ProtocolVersion zcl.Zu8

func (a ProtocolVersion) ID() zcl.AttrID           { return 4 }
func (a ProtocolVersion) Cluster() zcl.ClusterID   { return CommissioningCluster }
func (a *ProtocolVersion) Value() *ProtocolVersion { return a }
func (a ProtocolVersion) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *ProtocolVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ProtocolVersion(*nt)
	return br, err
}

func (a ProtocolVersion) Readable() bool   { return true }
func (a ProtocolVersion) Writable() bool   { return true }
func (a ProtocolVersion) Reportable() bool { return false }
func (a ProtocolVersion) SceneIndex() int  { return -1 }

func (a ProtocolVersion) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type StackProfile zcl.Zu8

func (a StackProfile) ID() zcl.AttrID         { return 5 }
func (a StackProfile) Cluster() zcl.ClusterID { return CommissioningCluster }
func (a *StackProfile) Value() *StackProfile  { return a }
func (a StackProfile) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *StackProfile) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = StackProfile(*nt)
	return br, err
}

func (a StackProfile) Readable() bool   { return true }
func (a StackProfile) Writable() bool   { return true }
func (a StackProfile) Reportable() bool { return false }
func (a StackProfile) SceneIndex() int  { return -1 }

func (a StackProfile) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type StartupControl zcl.Zenum8

func (a StartupControl) ID() zcl.AttrID          { return 6 }
func (a StartupControl) Cluster() zcl.ClusterID  { return CommissioningCluster }
func (a *StartupControl) Value() *StartupControl { return a }
func (a StartupControl) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *StartupControl) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = StartupControl(*nt)
	return br, err
}

func (a StartupControl) Readable() bool   { return true }
func (a StartupControl) Writable() bool   { return true }
func (a StartupControl) Reportable() bool { return false }
func (a StartupControl) SceneIndex() int  { return -1 }

func (a StartupControl) String() string {
	switch a {
	case 0x00:
		return "Part of the network"
	case 0x01:
		return "Form a network"
	case 0x02:
		return "Rejoin the network"
	case 0x03:
		return "Start from scratch"
	}

	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsPartOfTheNetwork checks if StartupControl equals the value for Part of the network (0x00)
func (a StartupControl) IsPartOfTheNetwork() bool { return a == 0x00 }

// SetPartOfTheNetwork sets StartupControl to Part of the network (0x00)
func (a *StartupControl) SetPartOfTheNetwork() { *a = 0x00 }

// IsFormANetwork checks if StartupControl equals the value for Form a network (0x01)
func (a StartupControl) IsFormANetwork() bool { return a == 0x01 }

// SetFormANetwork sets StartupControl to Form a network (0x01)
func (a *StartupControl) SetFormANetwork() { *a = 0x01 }

// IsRejoinTheNetwork checks if StartupControl equals the value for Rejoin the network (0x02)
func (a StartupControl) IsRejoinTheNetwork() bool { return a == 0x02 }

// SetRejoinTheNetwork sets StartupControl to Rejoin the network (0x02)
func (a *StartupControl) SetRejoinTheNetwork() { *a = 0x02 }

// IsStartFromScratch checks if StartupControl equals the value for Start from scratch (0x03)
func (a StartupControl) IsStartFromScratch() bool { return a == 0x03 }

// SetStartFromScratch sets StartupControl to Start from scratch (0x03)
func (a *StartupControl) SetStartFromScratch() { *a = 0x03 }

type TrustCenterAddress zcl.Zuid

func (a TrustCenterAddress) ID() zcl.AttrID              { return 16 }
func (a TrustCenterAddress) Cluster() zcl.ClusterID      { return CommissioningCluster }
func (a *TrustCenterAddress) Value() *TrustCenterAddress { return a }
func (a TrustCenterAddress) MarshalZcl() ([]byte, error) {
	return zcl.Zuid(a).MarshalZcl()
}
func (a *TrustCenterAddress) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*a = TrustCenterAddress(*nt)
	return br, err
}

func (a TrustCenterAddress) Readable() bool   { return true }
func (a TrustCenterAddress) Writable() bool   { return true }
func (a TrustCenterAddress) Reportable() bool { return false }
func (a TrustCenterAddress) SceneIndex() int  { return -1 }

func (a TrustCenterAddress) String() string {

	return zcl.Sprintf("%s", zcl.Zuid(a))
}

type TrustCenterMasterKey zcl.Zseckey

func (a TrustCenterMasterKey) ID() zcl.AttrID                { return 17 }
func (a TrustCenterMasterKey) Cluster() zcl.ClusterID        { return CommissioningCluster }
func (a *TrustCenterMasterKey) Value() *TrustCenterMasterKey { return a }
func (a TrustCenterMasterKey) MarshalZcl() ([]byte, error) {
	return zcl.Zseckey(a).MarshalZcl()
}
func (a *TrustCenterMasterKey) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zseckey)
	br, err := nt.UnmarshalZcl(b)
	*a = TrustCenterMasterKey(*nt)
	return br, err
}

func (a TrustCenterMasterKey) Readable() bool   { return true }
func (a TrustCenterMasterKey) Writable() bool   { return true }
func (a TrustCenterMasterKey) Reportable() bool { return false }
func (a TrustCenterMasterKey) SceneIndex() int  { return -1 }

func (a TrustCenterMasterKey) String() string {

	return zcl.Sprintf("%s", zcl.Zseckey(a))
}

type NetworkKey zcl.Zseckey

func (a NetworkKey) ID() zcl.AttrID         { return 18 }
func (a NetworkKey) Cluster() zcl.ClusterID { return CommissioningCluster }
func (a *NetworkKey) Value() *NetworkKey    { return a }
func (a NetworkKey) MarshalZcl() ([]byte, error) {
	return zcl.Zseckey(a).MarshalZcl()
}
func (a *NetworkKey) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zseckey)
	br, err := nt.UnmarshalZcl(b)
	*a = NetworkKey(*nt)
	return br, err
}

func (a NetworkKey) Readable() bool   { return true }
func (a NetworkKey) Writable() bool   { return true }
func (a NetworkKey) Reportable() bool { return false }
func (a NetworkKey) SceneIndex() int  { return -1 }

func (a NetworkKey) String() string {

	return zcl.Sprintf("%s", zcl.Zseckey(a))
}

type UseInsecureJoin zcl.Zbool

func (a UseInsecureJoin) ID() zcl.AttrID           { return 19 }
func (a UseInsecureJoin) Cluster() zcl.ClusterID   { return CommissioningCluster }
func (a *UseInsecureJoin) Value() *UseInsecureJoin { return a }
func (a UseInsecureJoin) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *UseInsecureJoin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = UseInsecureJoin(*nt)
	return br, err
}

func (a UseInsecureJoin) Readable() bool   { return true }
func (a UseInsecureJoin) Writable() bool   { return true }
func (a UseInsecureJoin) Reportable() bool { return false }
func (a UseInsecureJoin) SceneIndex() int  { return -1 }

func (a UseInsecureJoin) String() string {

	return zcl.Sprintf("%s", zcl.Zbool(a))
}

type PreconfiguredLinkKey zcl.Zseckey

func (a PreconfiguredLinkKey) ID() zcl.AttrID                { return 20 }
func (a PreconfiguredLinkKey) Cluster() zcl.ClusterID        { return CommissioningCluster }
func (a *PreconfiguredLinkKey) Value() *PreconfiguredLinkKey { return a }
func (a PreconfiguredLinkKey) MarshalZcl() ([]byte, error) {
	return zcl.Zseckey(a).MarshalZcl()
}
func (a *PreconfiguredLinkKey) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zseckey)
	br, err := nt.UnmarshalZcl(b)
	*a = PreconfiguredLinkKey(*nt)
	return br, err
}

func (a PreconfiguredLinkKey) Readable() bool   { return true }
func (a PreconfiguredLinkKey) Writable() bool   { return true }
func (a PreconfiguredLinkKey) Reportable() bool { return false }
func (a PreconfiguredLinkKey) SceneIndex() int  { return -1 }

func (a PreconfiguredLinkKey) String() string {

	return zcl.Sprintf("%s", zcl.Zseckey(a))
}

type NetworkKeySeqNum zcl.Zu8

func (a NetworkKeySeqNum) ID() zcl.AttrID            { return 21 }
func (a NetworkKeySeqNum) Cluster() zcl.ClusterID    { return CommissioningCluster }
func (a *NetworkKeySeqNum) Value() *NetworkKeySeqNum { return a }
func (a NetworkKeySeqNum) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *NetworkKeySeqNum) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NetworkKeySeqNum(*nt)
	return br, err
}

func (a NetworkKeySeqNum) Readable() bool   { return true }
func (a NetworkKeySeqNum) Writable() bool   { return true }
func (a NetworkKeySeqNum) Reportable() bool { return false }
func (a NetworkKeySeqNum) SceneIndex() int  { return -1 }

func (a NetworkKeySeqNum) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type NetworkKeyType zcl.Zenum8

func (a NetworkKeyType) ID() zcl.AttrID          { return 22 }
func (a NetworkKeyType) Cluster() zcl.ClusterID  { return CommissioningCluster }
func (a *NetworkKeyType) Value() *NetworkKeyType { return a }
func (a NetworkKeyType) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *NetworkKeyType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = NetworkKeyType(*nt)
	return br, err
}

func (a NetworkKeyType) Readable() bool   { return true }
func (a NetworkKeyType) Writable() bool   { return true }
func (a NetworkKeyType) Reportable() bool { return false }
func (a NetworkKeyType) SceneIndex() int  { return -1 }

func (a NetworkKeyType) String() string {
	switch a {
	case 0x01:
		return "Standard"
	case 0x05:
		return "High Security"
	}

	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsStandard checks if NetworkKeyType equals the value for Standard (0x01)
func (a NetworkKeyType) IsStandard() bool { return a == 0x01 }

// SetStandard sets NetworkKeyType to Standard (0x01)
func (a *NetworkKeyType) SetStandard() { *a = 0x01 }

// IsHighSecurity checks if NetworkKeyType equals the value for High Security (0x05)
func (a NetworkKeyType) IsHighSecurity() bool { return a == 0x05 }

// SetHighSecurity sets NetworkKeyType to High Security (0x05)
func (a *NetworkKeyType) SetHighSecurity() { *a = 0x05 }

type NetworkManagerAddress zcl.Zu16

func (a NetworkManagerAddress) ID() zcl.AttrID                 { return 23 }
func (a NetworkManagerAddress) Cluster() zcl.ClusterID         { return CommissioningCluster }
func (a *NetworkManagerAddress) Value() *NetworkManagerAddress { return a }
func (a NetworkManagerAddress) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *NetworkManagerAddress) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NetworkManagerAddress(*nt)
	return br, err
}

func (a NetworkManagerAddress) Readable() bool   { return true }
func (a NetworkManagerAddress) Writable() bool   { return true }
func (a NetworkManagerAddress) Reportable() bool { return false }
func (a NetworkManagerAddress) SceneIndex() int  { return -1 }

func (a NetworkManagerAddress) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type ScanAttempts zcl.Zu8

func (a ScanAttempts) ID() zcl.AttrID         { return 32 }
func (a ScanAttempts) Cluster() zcl.ClusterID { return CommissioningCluster }
func (a *ScanAttempts) Value() *ScanAttempts  { return a }
func (a ScanAttempts) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *ScanAttempts) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ScanAttempts(*nt)
	return br, err
}

func (a ScanAttempts) Readable() bool   { return true }
func (a ScanAttempts) Writable() bool   { return true }
func (a ScanAttempts) Reportable() bool { return false }
func (a ScanAttempts) SceneIndex() int  { return -1 }

func (a ScanAttempts) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type TimeBetweenScans zcl.Zu16

func (a TimeBetweenScans) ID() zcl.AttrID            { return 33 }
func (a TimeBetweenScans) Cluster() zcl.ClusterID    { return CommissioningCluster }
func (a *TimeBetweenScans) Value() *TimeBetweenScans { return a }
func (a TimeBetweenScans) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *TimeBetweenScans) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = TimeBetweenScans(*nt)
	return br, err
}

func (a TimeBetweenScans) Readable() bool   { return true }
func (a TimeBetweenScans) Writable() bool   { return true }
func (a TimeBetweenScans) Reportable() bool { return false }
func (a TimeBetweenScans) SceneIndex() int  { return -1 }

func (a TimeBetweenScans) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RejoinInterval zcl.Zu16

func (a RejoinInterval) ID() zcl.AttrID          { return 34 }
func (a RejoinInterval) Cluster() zcl.ClusterID  { return CommissioningCluster }
func (a *RejoinInterval) Value() *RejoinInterval { return a }
func (a RejoinInterval) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RejoinInterval) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RejoinInterval(*nt)
	return br, err
}

func (a RejoinInterval) Readable() bool   { return true }
func (a RejoinInterval) Writable() bool   { return true }
func (a RejoinInterval) Reportable() bool { return false }
func (a RejoinInterval) SceneIndex() int  { return -1 }

func (a RejoinInterval) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type MaxRejoinInterval zcl.Zu16

func (a MaxRejoinInterval) ID() zcl.AttrID             { return 35 }
func (a MaxRejoinInterval) Cluster() zcl.ClusterID     { return CommissioningCluster }
func (a *MaxRejoinInterval) Value() *MaxRejoinInterval { return a }
func (a MaxRejoinInterval) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MaxRejoinInterval) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxRejoinInterval(*nt)
	return br, err
}

func (a MaxRejoinInterval) Readable() bool   { return true }
func (a MaxRejoinInterval) Writable() bool   { return true }
func (a MaxRejoinInterval) Reportable() bool { return false }
func (a MaxRejoinInterval) SceneIndex() int  { return -1 }

func (a MaxRejoinInterval) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type IndirectPollRate zcl.Zu16

func (a IndirectPollRate) ID() zcl.AttrID            { return 48 }
func (a IndirectPollRate) Cluster() zcl.ClusterID    { return CommissioningCluster }
func (a *IndirectPollRate) Value() *IndirectPollRate { return a }
func (a IndirectPollRate) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *IndirectPollRate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = IndirectPollRate(*nt)
	return br, err
}

func (a IndirectPollRate) Readable() bool   { return true }
func (a IndirectPollRate) Writable() bool   { return true }
func (a IndirectPollRate) Reportable() bool { return false }
func (a IndirectPollRate) SceneIndex() int  { return -1 }

func (a IndirectPollRate) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type ParentRetryThreshold zcl.Zu8

func (a ParentRetryThreshold) ID() zcl.AttrID                { return 48 }
func (a ParentRetryThreshold) Cluster() zcl.ClusterID        { return CommissioningCluster }
func (a *ParentRetryThreshold) Value() *ParentRetryThreshold { return a }
func (a ParentRetryThreshold) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *ParentRetryThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ParentRetryThreshold(*nt)
	return br, err
}

func (a ParentRetryThreshold) Readable() bool   { return true }
func (a ParentRetryThreshold) Writable() bool   { return true }
func (a ParentRetryThreshold) Reportable() bool { return false }
func (a ParentRetryThreshold) SceneIndex() int  { return -1 }

func (a ParentRetryThreshold) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type ConcentratorFlag zcl.Zbool

func (a ConcentratorFlag) ID() zcl.AttrID            { return 64 }
func (a ConcentratorFlag) Cluster() zcl.ClusterID    { return CommissioningCluster }
func (a *ConcentratorFlag) Value() *ConcentratorFlag { return a }
func (a ConcentratorFlag) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *ConcentratorFlag) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = ConcentratorFlag(*nt)
	return br, err
}

func (a ConcentratorFlag) Readable() bool   { return true }
func (a ConcentratorFlag) Writable() bool   { return true }
func (a ConcentratorFlag) Reportable() bool { return false }
func (a ConcentratorFlag) SceneIndex() int  { return -1 }

func (a ConcentratorFlag) String() string {

	return zcl.Sprintf("%s", zcl.Zbool(a))
}

type ConcentratorRadius zcl.Zu8

func (a ConcentratorRadius) ID() zcl.AttrID              { return 65 }
func (a ConcentratorRadius) Cluster() zcl.ClusterID      { return CommissioningCluster }
func (a *ConcentratorRadius) Value() *ConcentratorRadius { return a }
func (a ConcentratorRadius) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *ConcentratorRadius) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ConcentratorRadius(*nt)
	return br, err
}

func (a ConcentratorRadius) Readable() bool   { return true }
func (a ConcentratorRadius) Writable() bool   { return true }
func (a ConcentratorRadius) Reportable() bool { return false }
func (a ConcentratorRadius) SceneIndex() int  { return -1 }

func (a ConcentratorRadius) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type ConcentratorDiscoveryTime zcl.Zu8

func (a ConcentratorDiscoveryTime) ID() zcl.AttrID                     { return 66 }
func (a ConcentratorDiscoveryTime) Cluster() zcl.ClusterID             { return CommissioningCluster }
func (a *ConcentratorDiscoveryTime) Value() *ConcentratorDiscoveryTime { return a }
func (a ConcentratorDiscoveryTime) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *ConcentratorDiscoveryTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ConcentratorDiscoveryTime(*nt)
	return br, err
}

func (a ConcentratorDiscoveryTime) Readable() bool   { return true }
func (a ConcentratorDiscoveryTime) Writable() bool   { return true }
func (a ConcentratorDiscoveryTime) Reportable() bool { return false }
func (a ConcentratorDiscoveryTime) SceneIndex() int  { return -1 }

func (a ConcentratorDiscoveryTime) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type MacAddress zcl.Zuid

func (a MacAddress) ID() zcl.AttrID         { return 56833 }
func (a MacAddress) Cluster() zcl.ClusterID { return CommissioningCluster }
func (a *MacAddress) Value() *MacAddress    { return a }
func (a MacAddress) MarshalZcl() ([]byte, error) {
	return zcl.Zuid(a).MarshalZcl()
}
func (a *MacAddress) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*a = MacAddress(*nt)
	return br, err
}

func (a MacAddress) Readable() bool   { return true }
func (a MacAddress) Writable() bool   { return true }
func (a MacAddress) Reportable() bool { return false }
func (a MacAddress) SceneIndex() int  { return -1 }

func (a MacAddress) String() string {

	return zcl.Sprintf("%s", zcl.Zuid(a))
}
