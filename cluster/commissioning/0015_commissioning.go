// Attributes and commands for commissioning and managing a ZigBee device.
package commissioning

import (
	"neotor.se/zcl"
)

// Commissioning
const CommissioningID zcl.ClusterID = 21

var CommissioningCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		RestartDeviceCommand:            func() zcl.Command { return new(RestartDevice) },
		SaveStartupParametersCommand:    func() zcl.Command { return new(SaveStartupParameters) },
		RestoreStartupParametersCommand: func() zcl.Command { return new(RestoreStartupParameters) },
		ResetStartupParametersCommand:   func() zcl.Command { return new(ResetStartupParameters) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{
		RestartDeviceResponseCommand:            func() zcl.Command { return new(RestartDeviceResponse) },
		SaveStartupParametersResponseCommand:    func() zcl.Command { return new(SaveStartupParametersResponse) },
		RestoreStartupParametersResponseCommand: func() zcl.Command { return new(RestoreStartupParametersResponse) },
		ResetStartupParametersResponseCommand:   func() zcl.Command { return new(ResetStartupParametersResponse) },
	},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		ShortAddressAttr:              func() zcl.Attr { return new(ShortAddress) },
		ExtendedPanIdAttr:             func() zcl.Attr { return new(ExtendedPanId) },
		PanIdAttr:                     func() zcl.Attr { return new(PanId) },
		ChannelMaskAttr:               func() zcl.Attr { return new(ChannelMask) },
		ProtocolVersionAttr:           func() zcl.Attr { return new(ProtocolVersion) },
		StackProfileAttr:              func() zcl.Attr { return new(StackProfile) },
		StartupControlAttr:            func() zcl.Attr { return new(StartupControl) },
		TrustCenterAddressAttr:        func() zcl.Attr { return new(TrustCenterAddress) },
		TrustCenterMasterKeyAttr:      func() zcl.Attr { return new(TrustCenterMasterKey) },
		NetworkKeyAttr:                func() zcl.Attr { return new(NetworkKey) },
		UseInsecureJoinAttr:           func() zcl.Attr { return new(UseInsecureJoin) },
		PreconfiguredLinkKeyAttr:      func() zcl.Attr { return new(PreconfiguredLinkKey) },
		NetworkKeySeqNumAttr:          func() zcl.Attr { return new(NetworkKeySeqNum) },
		NetworkKeyTypeAttr:            func() zcl.Attr { return new(NetworkKeyType) },
		NetworkManagerAddressAttr:     func() zcl.Attr { return new(NetworkManagerAddress) },
		ScanAttemptsAttr:              func() zcl.Attr { return new(ScanAttempts) },
		TimeBetweenScansAttr:          func() zcl.Attr { return new(TimeBetweenScans) },
		RejoinIntervalAttr:            func() zcl.Attr { return new(RejoinInterval) },
		MaxRejoinIntervalAttr:         func() zcl.Attr { return new(MaxRejoinInterval) },
		IndirectPollRateAttr:          func() zcl.Attr { return new(IndirectPollRate) },
		ParentRetryThresholdAttr:      func() zcl.Attr { return new(ParentRetryThreshold) },
		ConcentratorFlagAttr:          func() zcl.Attr { return new(ConcentratorFlag) },
		ConcentratorRadiusAttr:        func() zcl.Attr { return new(ConcentratorRadius) },
		ConcentratorDiscoveryTimeAttr: func() zcl.Attr { return new(ConcentratorDiscoveryTime) },
		MacAddressAttr:                func() zcl.Attr { return new(MacAddress) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

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
	return CommissioningID
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
	return CommissioningID
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
	return CommissioningID
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
	return CommissioningID
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
	return CommissioningID
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
	return CommissioningID
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
	return CommissioningID
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
	return CommissioningID
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

// ShortAddress is an autogenerated attribute in the Commissioning cluster
type ShortAddress zcl.Zu16

const ShortAddressAttr zcl.AttrID = 0

func (a ShortAddress) ID() zcl.AttrID         { return ShortAddressAttr }
func (a ShortAddress) Cluster() zcl.ClusterID { return CommissioningID }
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

// ExtendedPanId is an autogenerated attribute in the Commissioning cluster
type ExtendedPanId zcl.Zuid

const ExtendedPanIdAttr zcl.AttrID = 1

func (a ExtendedPanId) ID() zcl.AttrID         { return ExtendedPanIdAttr }
func (a ExtendedPanId) Cluster() zcl.ClusterID { return CommissioningID }
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

// PanId is an autogenerated attribute in the Commissioning cluster
type PanId zcl.Zu16

const PanIdAttr zcl.AttrID = 2

func (a PanId) ID() zcl.AttrID         { return PanIdAttr }
func (a PanId) Cluster() zcl.ClusterID { return CommissioningID }
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

// ChannelMask is an autogenerated attribute in the Commissioning cluster
type ChannelMask zcl.Zbmp32

const ChannelMaskAttr zcl.AttrID = 3

func (a ChannelMask) ID() zcl.AttrID         { return ChannelMaskAttr }
func (a ChannelMask) Cluster() zcl.ClusterID { return CommissioningID }
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

// ProtocolVersion is an autogenerated attribute in the Commissioning cluster
type ProtocolVersion zcl.Zu8

const ProtocolVersionAttr zcl.AttrID = 4

func (a ProtocolVersion) ID() zcl.AttrID           { return ProtocolVersionAttr }
func (a ProtocolVersion) Cluster() zcl.ClusterID   { return CommissioningID }
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

// StackProfile is an autogenerated attribute in the Commissioning cluster
type StackProfile zcl.Zu8

const StackProfileAttr zcl.AttrID = 5

func (a StackProfile) ID() zcl.AttrID         { return StackProfileAttr }
func (a StackProfile) Cluster() zcl.ClusterID { return CommissioningID }
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

// StartupControl is an autogenerated attribute in the Commissioning cluster
type StartupControl zcl.Zenum8

const StartupControlAttr zcl.AttrID = 6

func (a StartupControl) ID() zcl.AttrID          { return StartupControlAttr }
func (a StartupControl) Cluster() zcl.ClusterID  { return CommissioningID }
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

// TrustCenterAddress is an autogenerated attribute in the Commissioning cluster
type TrustCenterAddress zcl.Zuid

const TrustCenterAddressAttr zcl.AttrID = 16

func (a TrustCenterAddress) ID() zcl.AttrID              { return TrustCenterAddressAttr }
func (a TrustCenterAddress) Cluster() zcl.ClusterID      { return CommissioningID }
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

// TrustCenterMasterKey is an autogenerated attribute in the Commissioning cluster
type TrustCenterMasterKey zcl.Zseckey

const TrustCenterMasterKeyAttr zcl.AttrID = 17

func (a TrustCenterMasterKey) ID() zcl.AttrID                { return TrustCenterMasterKeyAttr }
func (a TrustCenterMasterKey) Cluster() zcl.ClusterID        { return CommissioningID }
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

// NetworkKey is an autogenerated attribute in the Commissioning cluster
type NetworkKey zcl.Zseckey

const NetworkKeyAttr zcl.AttrID = 18

func (a NetworkKey) ID() zcl.AttrID         { return NetworkKeyAttr }
func (a NetworkKey) Cluster() zcl.ClusterID { return CommissioningID }
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

// UseInsecureJoin is an autogenerated attribute in the Commissioning cluster
type UseInsecureJoin zcl.Zbool

const UseInsecureJoinAttr zcl.AttrID = 19

func (a UseInsecureJoin) ID() zcl.AttrID           { return UseInsecureJoinAttr }
func (a UseInsecureJoin) Cluster() zcl.ClusterID   { return CommissioningID }
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

// PreconfiguredLinkKey is an autogenerated attribute in the Commissioning cluster
type PreconfiguredLinkKey zcl.Zseckey

const PreconfiguredLinkKeyAttr zcl.AttrID = 20

func (a PreconfiguredLinkKey) ID() zcl.AttrID                { return PreconfiguredLinkKeyAttr }
func (a PreconfiguredLinkKey) Cluster() zcl.ClusterID        { return CommissioningID }
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

// NetworkKeySeqNum is an autogenerated attribute in the Commissioning cluster
type NetworkKeySeqNum zcl.Zu8

const NetworkKeySeqNumAttr zcl.AttrID = 21

func (a NetworkKeySeqNum) ID() zcl.AttrID            { return NetworkKeySeqNumAttr }
func (a NetworkKeySeqNum) Cluster() zcl.ClusterID    { return CommissioningID }
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

// NetworkKeyType is an autogenerated attribute in the Commissioning cluster
type NetworkKeyType zcl.Zenum8

const NetworkKeyTypeAttr zcl.AttrID = 22

func (a NetworkKeyType) ID() zcl.AttrID          { return NetworkKeyTypeAttr }
func (a NetworkKeyType) Cluster() zcl.ClusterID  { return CommissioningID }
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

// NetworkManagerAddress is an autogenerated attribute in the Commissioning cluster
type NetworkManagerAddress zcl.Zu16

const NetworkManagerAddressAttr zcl.AttrID = 23

func (a NetworkManagerAddress) ID() zcl.AttrID                 { return NetworkManagerAddressAttr }
func (a NetworkManagerAddress) Cluster() zcl.ClusterID         { return CommissioningID }
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

// ScanAttempts is an autogenerated attribute in the Commissioning cluster
type ScanAttempts zcl.Zu8

const ScanAttemptsAttr zcl.AttrID = 32

func (a ScanAttempts) ID() zcl.AttrID         { return ScanAttemptsAttr }
func (a ScanAttempts) Cluster() zcl.ClusterID { return CommissioningID }
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

// TimeBetweenScans is an autogenerated attribute in the Commissioning cluster
type TimeBetweenScans zcl.Zu16

const TimeBetweenScansAttr zcl.AttrID = 33

func (a TimeBetweenScans) ID() zcl.AttrID            { return TimeBetweenScansAttr }
func (a TimeBetweenScans) Cluster() zcl.ClusterID    { return CommissioningID }
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

// RejoinInterval is an autogenerated attribute in the Commissioning cluster
type RejoinInterval zcl.Zu16

const RejoinIntervalAttr zcl.AttrID = 34

func (a RejoinInterval) ID() zcl.AttrID          { return RejoinIntervalAttr }
func (a RejoinInterval) Cluster() zcl.ClusterID  { return CommissioningID }
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

// MaxRejoinInterval is an autogenerated attribute in the Commissioning cluster
type MaxRejoinInterval zcl.Zu16

const MaxRejoinIntervalAttr zcl.AttrID = 35

func (a MaxRejoinInterval) ID() zcl.AttrID             { return MaxRejoinIntervalAttr }
func (a MaxRejoinInterval) Cluster() zcl.ClusterID     { return CommissioningID }
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

// IndirectPollRate is an autogenerated attribute in the Commissioning cluster
type IndirectPollRate zcl.Zu16

const IndirectPollRateAttr zcl.AttrID = 48

func (a IndirectPollRate) ID() zcl.AttrID            { return IndirectPollRateAttr }
func (a IndirectPollRate) Cluster() zcl.ClusterID    { return CommissioningID }
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

// ParentRetryThreshold is an autogenerated attribute in the Commissioning cluster
type ParentRetryThreshold zcl.Zu8

const ParentRetryThresholdAttr zcl.AttrID = 48

func (a ParentRetryThreshold) ID() zcl.AttrID                { return ParentRetryThresholdAttr }
func (a ParentRetryThreshold) Cluster() zcl.ClusterID        { return CommissioningID }
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

// ConcentratorFlag is an autogenerated attribute in the Commissioning cluster
type ConcentratorFlag zcl.Zbool

const ConcentratorFlagAttr zcl.AttrID = 64

func (a ConcentratorFlag) ID() zcl.AttrID            { return ConcentratorFlagAttr }
func (a ConcentratorFlag) Cluster() zcl.ClusterID    { return CommissioningID }
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

// ConcentratorRadius is an autogenerated attribute in the Commissioning cluster
type ConcentratorRadius zcl.Zu8

const ConcentratorRadiusAttr zcl.AttrID = 65

func (a ConcentratorRadius) ID() zcl.AttrID              { return ConcentratorRadiusAttr }
func (a ConcentratorRadius) Cluster() zcl.ClusterID      { return CommissioningID }
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

// ConcentratorDiscoveryTime is an autogenerated attribute in the Commissioning cluster
type ConcentratorDiscoveryTime zcl.Zu8

const ConcentratorDiscoveryTimeAttr zcl.AttrID = 66

func (a ConcentratorDiscoveryTime) ID() zcl.AttrID                     { return ConcentratorDiscoveryTimeAttr }
func (a ConcentratorDiscoveryTime) Cluster() zcl.ClusterID             { return CommissioningID }
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

// MacAddress is an autogenerated attribute in the Commissioning cluster
type MacAddress zcl.Zuid

const MacAddressAttr zcl.AttrID = 56833

func (a MacAddress) ID() zcl.AttrID         { return MacAddressAttr }
func (a MacAddress) Cluster() zcl.ClusterID { return CommissioningID }
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
