// The Simple Metering Cluster provides a mechanism to retrieve usage information from Electric, Gas, Water, and potentially Thermal metering devices. These devices can operate on either battery or mains power, and can have a wide variety of sophistication.
package smart_energy

import (
	"neotor.se/zcl"
)

// SimpleMetering
const SimpleMeteringID zcl.ClusterID = 1794

var SimpleMeteringCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		GetProfileCommand:            func() zcl.Command { return new(GetProfile) },
		RequestMirrorResponseCommand: func() zcl.Command { return new(RequestMirrorResponse) },
		MirrorRemovedCommand:         func() zcl.Command { return new(MirrorRemoved) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		CurrentSummationDeliveredAttr: func() zcl.Attr { return new(CurrentSummationDelivered) },
		StatusAttr:                    func() zcl.Attr { return new(Status) },
		UnitOfMeasureAttr:             func() zcl.Attr { return new(UnitOfMeasure) },
		MultiplierAttr:                func() zcl.Attr { return new(Multiplier) },
		DivisorAttr:                   func() zcl.Attr { return new(Divisor) },
		SummationFormattingAttr:       func() zcl.Attr { return new(SummationFormatting) },
		MeteringDeviceTypeAttr:        func() zcl.Attr { return new(MeteringDeviceType) },
		InstantaneousDemandAttr:       func() zcl.Attr { return new(InstantaneousDemand) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

type GetProfile struct {
	IntervalChannel zcl.Zu8
	EndTime         zcl.Zutc
	NumberOfPeriods zcl.Zu8
}

const GetProfileCommand zcl.CommandID = 0

func (v *GetProfile) Values() []zcl.Val {
	return []zcl.Val{
		&v.IntervalChannel,
		&v.EndTime,
		&v.NumberOfPeriods,
	}
}

func (v GetProfile) ID() zcl.CommandID {
	return GetProfileCommand
}

func (v GetProfile) Cluster() zcl.ClusterID {
	return SimpleMeteringID
}

func (v GetProfile) MnfCode() []byte {
	return []byte{}
}

func (v GetProfile) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.IntervalChannel.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.EndTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.NumberOfPeriods.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetProfile) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.IntervalChannel).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.EndTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NumberOfPeriods).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v *GetProfile) IntervalChannelString() string {
	switch v.IntervalChannel {
	case 0x00:
		return "Consumption Delivered"
	case 0x01:
		return "Consumption Received"
	}
	return zcl.Sprintf("%v", zcl.Zu8(v.IntervalChannel))
}
func (v *GetProfile) EndTimeString() string {
	return zcl.Sprintf("%v", zcl.Zutc(v.EndTime))
}
func (v *GetProfile) NumberOfPeriodsString() string {
	return zcl.Sprintf("%v", zcl.Zu8(v.NumberOfPeriods))
}

func (v *GetProfile) String() string {
	var str []string
	str = append(str, "IntervalChannel["+v.IntervalChannelString()+"]")
	str = append(str, "EndTime["+v.EndTimeString()+"]")
	str = append(str, "NumberOfPeriods["+v.NumberOfPeriodsString()+"]")
	return "GetProfile{" + zcl.StrJoin(str, " ") + "}"
}

type RequestMirrorResponse struct {
	EndpointId zcl.Zu16
}

const RequestMirrorResponseCommand zcl.CommandID = 1

func (v *RequestMirrorResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.EndpointId,
	}
}

func (v RequestMirrorResponse) ID() zcl.CommandID {
	return RequestMirrorResponseCommand
}

func (v RequestMirrorResponse) Cluster() zcl.ClusterID {
	return SimpleMeteringID
}

func (v RequestMirrorResponse) MnfCode() []byte {
	return []byte{}
}

func (v RequestMirrorResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.EndpointId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *RequestMirrorResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.EndpointId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v *RequestMirrorResponse) EndpointIdString() string {
	return zcl.Sprintf("%v", zcl.Zu16(v.EndpointId))
}

func (v *RequestMirrorResponse) String() string {
	var str []string
	str = append(str, "EndpointId["+v.EndpointIdString()+"]")
	return "RequestMirrorResponse{" + zcl.StrJoin(str, " ") + "}"
}

type MirrorRemoved struct {
	RemovedEndpointId zcl.Zu16
}

const MirrorRemovedCommand zcl.CommandID = 2

func (v *MirrorRemoved) Values() []zcl.Val {
	return []zcl.Val{
		&v.RemovedEndpointId,
	}
}

func (v MirrorRemoved) ID() zcl.CommandID {
	return MirrorRemovedCommand
}

func (v MirrorRemoved) Cluster() zcl.ClusterID {
	return SimpleMeteringID
}

func (v MirrorRemoved) MnfCode() []byte {
	return []byte{}
}

func (v MirrorRemoved) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.RemovedEndpointId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *MirrorRemoved) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.RemovedEndpointId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v *MirrorRemoved) RemovedEndpointIdString() string {
	return zcl.Sprintf("%v", zcl.Zu16(v.RemovedEndpointId))
}

func (v *MirrorRemoved) String() string {
	var str []string
	str = append(str, "RemovedEndpointId["+v.RemovedEndpointIdString()+"]")
	return "MirrorRemoved{" + zcl.StrJoin(str, " ") + "}"
}

// CurrentSummationDelivered is an autogenerated attribute in the SimpleMetering cluster
type CurrentSummationDelivered zcl.Zu48

const CurrentSummationDeliveredAttr zcl.AttrID = 0

func (a CurrentSummationDelivered) ID() zcl.AttrID                     { return CurrentSummationDeliveredAttr }
func (a CurrentSummationDelivered) Cluster() zcl.ClusterID             { return SimpleMeteringID }
func (a *CurrentSummationDelivered) Value() *CurrentSummationDelivered { return a }
func (a CurrentSummationDelivered) MarshalZcl() ([]byte, error) {
	return zcl.Zu48(a).MarshalZcl()
}
func (a *CurrentSummationDelivered) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu48)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentSummationDelivered(*nt)
	return br, err
}

func (a CurrentSummationDelivered) Readable() bool   { return true }
func (a CurrentSummationDelivered) Writable() bool   { return false }
func (a CurrentSummationDelivered) Reportable() bool { return false }
func (a CurrentSummationDelivered) SceneIndex() int  { return -1 }

func (a CurrentSummationDelivered) String() string {
	return zcl.Sprintf("%v", zcl.Zu48(a))
}

// Status is an autogenerated attribute in the SimpleMetering cluster
type Status zcl.Zbmp8

const StatusAttr zcl.AttrID = 512

func (a Status) ID() zcl.AttrID         { return StatusAttr }
func (a Status) Cluster() zcl.ClusterID { return SimpleMeteringID }
func (a *Status) Value() *Status        { return a }
func (a Status) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *Status) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = Status(*nt)
	return br, err
}

func (a Status) Readable() bool   { return true }
func (a Status) Writable() bool   { return false }
func (a Status) Reportable() bool { return false }
func (a Status) SceneIndex() int  { return -1 }

func (a Status) String() string {
	var bstr []string
	if a.IsCheckMeter() {
		bstr = append(bstr, "Check Meter")
	}
	if a.IsLowBattery() {
		bstr = append(bstr, "Low Battery")
	}
	if a.IsTamperDetect() {
		bstr = append(bstr, "Tamper Detect")
	}
	if a.IsPowerFailure() {
		bstr = append(bstr, "Power Failure")
	}
	if a.IsPowerQuality() {
		bstr = append(bstr, "Power Quality")
	}
	if a.IsLeakDetect() {
		bstr = append(bstr, "Leak Detect")
	}
	if a.IsServiceDisconnectOpen() {
		bstr = append(bstr, "Service Disconnect Open")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a Status) IsCheckMeter() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *Status) SetCheckMeter(b bool) {
	*a = Status(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a Status) IsLowBattery() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *Status) SetLowBattery(b bool) {
	*a = Status(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a Status) IsTamperDetect() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *Status) SetTamperDetect(b bool) {
	*a = Status(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a Status) IsPowerFailure() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *Status) SetPowerFailure(b bool) {
	*a = Status(zcl.BitmapSet([]byte(*a), 3, b))
}

func (a Status) IsPowerQuality() bool {
	return zcl.BitmapTest([]byte(a), 4)
}
func (a *Status) SetPowerQuality(b bool) {
	*a = Status(zcl.BitmapSet([]byte(*a), 4, b))
}

func (a Status) IsLeakDetect() bool {
	return zcl.BitmapTest([]byte(a), 5)
}
func (a *Status) SetLeakDetect(b bool) {
	*a = Status(zcl.BitmapSet([]byte(*a), 5, b))
}

func (a Status) IsServiceDisconnectOpen() bool {
	return zcl.BitmapTest([]byte(a), 6)
}
func (a *Status) SetServiceDisconnectOpen(b bool) {
	*a = Status(zcl.BitmapSet([]byte(*a), 6, b))
}

// UnitOfMeasure is an autogenerated attribute in the SimpleMetering cluster
type UnitOfMeasure zcl.Zenum8

const UnitOfMeasureAttr zcl.AttrID = 768

func (a UnitOfMeasure) ID() zcl.AttrID         { return UnitOfMeasureAttr }
func (a UnitOfMeasure) Cluster() zcl.ClusterID { return SimpleMeteringID }
func (a *UnitOfMeasure) Value() *UnitOfMeasure { return a }
func (a UnitOfMeasure) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *UnitOfMeasure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = UnitOfMeasure(*nt)
	return br, err
}

func (a UnitOfMeasure) Readable() bool   { return true }
func (a UnitOfMeasure) Writable() bool   { return false }
func (a UnitOfMeasure) Reportable() bool { return false }
func (a UnitOfMeasure) SceneIndex() int  { return -1 }

func (a UnitOfMeasure) String() string {
	switch a {
	case 0x00:
		return "kW & kWh binary"
	case 0x01:
		return "m³ & m³/h binary"
	case 0x02:
		return "ft³ & ft³/h binary"
	case 0x03:
		return "ccf & ccf/h binary"
	case 0x04:
		return "US gl & Us gl/h binary"
	case 0x05:
		return "IMP gl & IMP gl/h binary"
	case 0x06:
		return "BTUs & BTU/h binary"
	case 0x07:
		return "Liters & l/h binary"
	case 0x08:
		return "kPA(gauge) binary"
	case 0x09:
		return "kPA(absolute) binary"
	case 0x80:
		return "kW & kWh BCD"
	case 0x81:
		return "m³ & m³/h BCD"
	case 0x82:
		return "ft³ & ft³/h BCD"
	case 0x83:
		return "ccf & ccf/h BCD"
	case 0x84:
		return "US gl & Us gl/h BCD"
	case 0x85:
		return "IMP gl & IMP gl/h BCD"
	case 0x86:
		return "BTUs & BTU/h BCD"
	case 0x87:
		return "Liters & l/h BCD"
	case 0x88:
		return "kPA(gauge) BCD"
	case 0x89:
		return "kPA(absolute) BCD"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

// IsKwKwhBinary checks if UnitOfMeasure equals the value for kW & kWh binary (0x00)
func (a UnitOfMeasure) IsKwKwhBinary() bool { return a == 0x00 }

// SetKwKwhBinary sets UnitOfMeasure to kW & kWh binary (0x00)
func (a *UnitOfMeasure) SetKwKwhBinary() { *a = 0x00 }

// IsMMHBinary checks if UnitOfMeasure equals the value for m³ & m³/h binary (0x01)
func (a UnitOfMeasure) IsMMHBinary() bool { return a == 0x01 }

// SetMMHBinary sets UnitOfMeasure to m³ & m³/h binary (0x01)
func (a *UnitOfMeasure) SetMMHBinary() { *a = 0x01 }

// IsFtFtHBinary checks if UnitOfMeasure equals the value for ft³ & ft³/h binary (0x02)
func (a UnitOfMeasure) IsFtFtHBinary() bool { return a == 0x02 }

// SetFtFtHBinary sets UnitOfMeasure to ft³ & ft³/h binary (0x02)
func (a *UnitOfMeasure) SetFtFtHBinary() { *a = 0x02 }

// IsCcfCcfHBinary checks if UnitOfMeasure equals the value for ccf & ccf/h binary (0x03)
func (a UnitOfMeasure) IsCcfCcfHBinary() bool { return a == 0x03 }

// SetCcfCcfHBinary sets UnitOfMeasure to ccf & ccf/h binary (0x03)
func (a *UnitOfMeasure) SetCcfCcfHBinary() { *a = 0x03 }

// IsUsGlUsGlHBinary checks if UnitOfMeasure equals the value for US gl & Us gl/h binary (0x04)
func (a UnitOfMeasure) IsUsGlUsGlHBinary() bool { return a == 0x04 }

// SetUsGlUsGlHBinary sets UnitOfMeasure to US gl & Us gl/h binary (0x04)
func (a *UnitOfMeasure) SetUsGlUsGlHBinary() { *a = 0x04 }

// IsImpGlImpGlHBinary checks if UnitOfMeasure equals the value for IMP gl & IMP gl/h binary (0x05)
func (a UnitOfMeasure) IsImpGlImpGlHBinary() bool { return a == 0x05 }

// SetImpGlImpGlHBinary sets UnitOfMeasure to IMP gl & IMP gl/h binary (0x05)
func (a *UnitOfMeasure) SetImpGlImpGlHBinary() { *a = 0x05 }

// IsBtusBtuHBinary checks if UnitOfMeasure equals the value for BTUs & BTU/h binary (0x06)
func (a UnitOfMeasure) IsBtusBtuHBinary() bool { return a == 0x06 }

// SetBtusBtuHBinary sets UnitOfMeasure to BTUs & BTU/h binary (0x06)
func (a *UnitOfMeasure) SetBtusBtuHBinary() { *a = 0x06 }

// IsLitersLHBinary checks if UnitOfMeasure equals the value for Liters & l/h binary (0x07)
func (a UnitOfMeasure) IsLitersLHBinary() bool { return a == 0x07 }

// SetLitersLHBinary sets UnitOfMeasure to Liters & l/h binary (0x07)
func (a *UnitOfMeasure) SetLitersLHBinary() { *a = 0x07 }

// IsKpaGaugeBinary checks if UnitOfMeasure equals the value for kPA(gauge) binary (0x08)
func (a UnitOfMeasure) IsKpaGaugeBinary() bool { return a == 0x08 }

// SetKpaGaugeBinary sets UnitOfMeasure to kPA(gauge) binary (0x08)
func (a *UnitOfMeasure) SetKpaGaugeBinary() { *a = 0x08 }

// IsKpaAbsoluteBinary checks if UnitOfMeasure equals the value for kPA(absolute) binary (0x09)
func (a UnitOfMeasure) IsKpaAbsoluteBinary() bool { return a == 0x09 }

// SetKpaAbsoluteBinary sets UnitOfMeasure to kPA(absolute) binary (0x09)
func (a *UnitOfMeasure) SetKpaAbsoluteBinary() { *a = 0x09 }

// IsKwKwhBcd checks if UnitOfMeasure equals the value for kW & kWh BCD (0x80)
func (a UnitOfMeasure) IsKwKwhBcd() bool { return a == 0x80 }

// SetKwKwhBcd sets UnitOfMeasure to kW & kWh BCD (0x80)
func (a *UnitOfMeasure) SetKwKwhBcd() { *a = 0x80 }

// IsMMHBcd checks if UnitOfMeasure equals the value for m³ & m³/h BCD (0x81)
func (a UnitOfMeasure) IsMMHBcd() bool { return a == 0x81 }

// SetMMHBcd sets UnitOfMeasure to m³ & m³/h BCD (0x81)
func (a *UnitOfMeasure) SetMMHBcd() { *a = 0x81 }

// IsFtFtHBcd checks if UnitOfMeasure equals the value for ft³ & ft³/h BCD (0x82)
func (a UnitOfMeasure) IsFtFtHBcd() bool { return a == 0x82 }

// SetFtFtHBcd sets UnitOfMeasure to ft³ & ft³/h BCD (0x82)
func (a *UnitOfMeasure) SetFtFtHBcd() { *a = 0x82 }

// IsCcfCcfHBcd checks if UnitOfMeasure equals the value for ccf & ccf/h BCD (0x83)
func (a UnitOfMeasure) IsCcfCcfHBcd() bool { return a == 0x83 }

// SetCcfCcfHBcd sets UnitOfMeasure to ccf & ccf/h BCD (0x83)
func (a *UnitOfMeasure) SetCcfCcfHBcd() { *a = 0x83 }

// IsUsGlUsGlHBcd checks if UnitOfMeasure equals the value for US gl & Us gl/h BCD (0x84)
func (a UnitOfMeasure) IsUsGlUsGlHBcd() bool { return a == 0x84 }

// SetUsGlUsGlHBcd sets UnitOfMeasure to US gl & Us gl/h BCD (0x84)
func (a *UnitOfMeasure) SetUsGlUsGlHBcd() { *a = 0x84 }

// IsImpGlImpGlHBcd checks if UnitOfMeasure equals the value for IMP gl & IMP gl/h BCD (0x85)
func (a UnitOfMeasure) IsImpGlImpGlHBcd() bool { return a == 0x85 }

// SetImpGlImpGlHBcd sets UnitOfMeasure to IMP gl & IMP gl/h BCD (0x85)
func (a *UnitOfMeasure) SetImpGlImpGlHBcd() { *a = 0x85 }

// IsBtusBtuHBcd checks if UnitOfMeasure equals the value for BTUs & BTU/h BCD (0x86)
func (a UnitOfMeasure) IsBtusBtuHBcd() bool { return a == 0x86 }

// SetBtusBtuHBcd sets UnitOfMeasure to BTUs & BTU/h BCD (0x86)
func (a *UnitOfMeasure) SetBtusBtuHBcd() { *a = 0x86 }

// IsLitersLHBcd checks if UnitOfMeasure equals the value for Liters & l/h BCD (0x87)
func (a UnitOfMeasure) IsLitersLHBcd() bool { return a == 0x87 }

// SetLitersLHBcd sets UnitOfMeasure to Liters & l/h BCD (0x87)
func (a *UnitOfMeasure) SetLitersLHBcd() { *a = 0x87 }

// IsKpaGaugeBcd checks if UnitOfMeasure equals the value for kPA(gauge) BCD (0x88)
func (a UnitOfMeasure) IsKpaGaugeBcd() bool { return a == 0x88 }

// SetKpaGaugeBcd sets UnitOfMeasure to kPA(gauge) BCD (0x88)
func (a *UnitOfMeasure) SetKpaGaugeBcd() { *a = 0x88 }

// IsKpaAbsoluteBcd checks if UnitOfMeasure equals the value for kPA(absolute) BCD (0x89)
func (a UnitOfMeasure) IsKpaAbsoluteBcd() bool { return a == 0x89 }

// SetKpaAbsoluteBcd sets UnitOfMeasure to kPA(absolute) BCD (0x89)
func (a *UnitOfMeasure) SetKpaAbsoluteBcd() { *a = 0x89 }

// Multiplier is an autogenerated attribute in the SimpleMetering cluster
type Multiplier zcl.Zu24

const MultiplierAttr zcl.AttrID = 769

func (a Multiplier) ID() zcl.AttrID         { return MultiplierAttr }
func (a Multiplier) Cluster() zcl.ClusterID { return SimpleMeteringID }
func (a *Multiplier) Value() *Multiplier    { return a }
func (a Multiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zu24(a).MarshalZcl()
}
func (a *Multiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*a = Multiplier(*nt)
	return br, err
}

func (a Multiplier) Readable() bool   { return true }
func (a Multiplier) Writable() bool   { return false }
func (a Multiplier) Reportable() bool { return false }
func (a Multiplier) SceneIndex() int  { return -1 }

func (a Multiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zu24(a))
}

// Divisor is an autogenerated attribute in the SimpleMetering cluster
type Divisor zcl.Zu24

const DivisorAttr zcl.AttrID = 770

func (a Divisor) ID() zcl.AttrID         { return DivisorAttr }
func (a Divisor) Cluster() zcl.ClusterID { return SimpleMeteringID }
func (a *Divisor) Value() *Divisor       { return a }
func (a Divisor) MarshalZcl() ([]byte, error) {
	return zcl.Zu24(a).MarshalZcl()
}
func (a *Divisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*a = Divisor(*nt)
	return br, err
}

func (a Divisor) Readable() bool   { return true }
func (a Divisor) Writable() bool   { return false }
func (a Divisor) Reportable() bool { return false }
func (a Divisor) SceneIndex() int  { return -1 }

func (a Divisor) String() string {
	return zcl.Sprintf("%v", zcl.Zu24(a))
}

// SummationFormatting is an autogenerated attribute in the SimpleMetering cluster
type SummationFormatting zcl.Zbmp8

const SummationFormattingAttr zcl.AttrID = 771

func (a SummationFormatting) ID() zcl.AttrID               { return SummationFormattingAttr }
func (a SummationFormatting) Cluster() zcl.ClusterID       { return SimpleMeteringID }
func (a *SummationFormatting) Value() *SummationFormatting { return a }
func (a SummationFormatting) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *SummationFormatting) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = SummationFormatting(*nt)
	return br, err
}

func (a SummationFormatting) Readable() bool   { return true }
func (a SummationFormatting) Writable() bool   { return false }
func (a SummationFormatting) Reportable() bool { return false }
func (a SummationFormatting) SceneIndex() int  { return -1 }

func (a SummationFormatting) String() string {
	var bstr []string
	if a.IsNumberDigitsLeft() {
		bstr = append(bstr, "Number Digits Left")
	}
	if a.IsSurpressLeadingZeros() {
		bstr = append(bstr, "Surpress Leading Zeros")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a SummationFormatting) IsNumberDigitsLeft() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *SummationFormatting) SetNumberDigitsLeft(b bool) {
	*a = SummationFormatting(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a SummationFormatting) IsSurpressLeadingZeros() bool {
	return zcl.BitmapTest([]byte(a), 7)
}
func (a *SummationFormatting) SetSurpressLeadingZeros(b bool) {
	*a = SummationFormatting(zcl.BitmapSet([]byte(*a), 7, b))
}

// MeteringDeviceType is an autogenerated attribute in the SimpleMetering cluster
type MeteringDeviceType zcl.Zenum8

const MeteringDeviceTypeAttr zcl.AttrID = 774

func (a MeteringDeviceType) ID() zcl.AttrID              { return MeteringDeviceTypeAttr }
func (a MeteringDeviceType) Cluster() zcl.ClusterID      { return SimpleMeteringID }
func (a *MeteringDeviceType) Value() *MeteringDeviceType { return a }
func (a MeteringDeviceType) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *MeteringDeviceType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = MeteringDeviceType(*nt)
	return br, err
}

func (a MeteringDeviceType) Readable() bool   { return true }
func (a MeteringDeviceType) Writable() bool   { return false }
func (a MeteringDeviceType) Reportable() bool { return false }
func (a MeteringDeviceType) SceneIndex() int  { return -1 }

func (a MeteringDeviceType) String() string {
	switch a {
	case 0x00:
		return "Electric Metering"
	case 0x01:
		return "Gas Metering"
	case 0x02:
		return "Water Metering"
	case 0x03:
		return "Thermal Metering"
	case 0x04:
		return "Pressure Metering"
	case 0x05:
		return "Heat Metering"
	case 0x06:
		return "Cooling Metering"
	case 0x80:
		return "Mirrored Gas Metering"
	case 0x81:
		return "Mirrored Water Metering"
	case 0x82:
		return "Mirrored Thermal Metering"
	case 0x83:
		return "Mirrored Pressure Metering"
	case 0x84:
		return "Mirrored Heat Metering"
	case 0x85:
		return "Mirrored Cooling Metering"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

// IsElectricMetering checks if MeteringDeviceType equals the value for Electric Metering (0x00)
func (a MeteringDeviceType) IsElectricMetering() bool { return a == 0x00 }

// SetElectricMetering sets MeteringDeviceType to Electric Metering (0x00)
func (a *MeteringDeviceType) SetElectricMetering() { *a = 0x00 }

// IsGasMetering checks if MeteringDeviceType equals the value for Gas Metering (0x01)
func (a MeteringDeviceType) IsGasMetering() bool { return a == 0x01 }

// SetGasMetering sets MeteringDeviceType to Gas Metering (0x01)
func (a *MeteringDeviceType) SetGasMetering() { *a = 0x01 }

// IsWaterMetering checks if MeteringDeviceType equals the value for Water Metering (0x02)
func (a MeteringDeviceType) IsWaterMetering() bool { return a == 0x02 }

// SetWaterMetering sets MeteringDeviceType to Water Metering (0x02)
func (a *MeteringDeviceType) SetWaterMetering() { *a = 0x02 }

// IsThermalMetering checks if MeteringDeviceType equals the value for Thermal Metering (0x03)
func (a MeteringDeviceType) IsThermalMetering() bool { return a == 0x03 }

// SetThermalMetering sets MeteringDeviceType to Thermal Metering (0x03)
func (a *MeteringDeviceType) SetThermalMetering() { *a = 0x03 }

// IsPressureMetering checks if MeteringDeviceType equals the value for Pressure Metering (0x04)
func (a MeteringDeviceType) IsPressureMetering() bool { return a == 0x04 }

// SetPressureMetering sets MeteringDeviceType to Pressure Metering (0x04)
func (a *MeteringDeviceType) SetPressureMetering() { *a = 0x04 }

// IsHeatMetering checks if MeteringDeviceType equals the value for Heat Metering (0x05)
func (a MeteringDeviceType) IsHeatMetering() bool { return a == 0x05 }

// SetHeatMetering sets MeteringDeviceType to Heat Metering (0x05)
func (a *MeteringDeviceType) SetHeatMetering() { *a = 0x05 }

// IsCoolingMetering checks if MeteringDeviceType equals the value for Cooling Metering (0x06)
func (a MeteringDeviceType) IsCoolingMetering() bool { return a == 0x06 }

// SetCoolingMetering sets MeteringDeviceType to Cooling Metering (0x06)
func (a *MeteringDeviceType) SetCoolingMetering() { *a = 0x06 }

// IsMirroredGasMetering checks if MeteringDeviceType equals the value for Mirrored Gas Metering (0x80)
func (a MeteringDeviceType) IsMirroredGasMetering() bool { return a == 0x80 }

// SetMirroredGasMetering sets MeteringDeviceType to Mirrored Gas Metering (0x80)
func (a *MeteringDeviceType) SetMirroredGasMetering() { *a = 0x80 }

// IsMirroredWaterMetering checks if MeteringDeviceType equals the value for Mirrored Water Metering (0x81)
func (a MeteringDeviceType) IsMirroredWaterMetering() bool { return a == 0x81 }

// SetMirroredWaterMetering sets MeteringDeviceType to Mirrored Water Metering (0x81)
func (a *MeteringDeviceType) SetMirroredWaterMetering() { *a = 0x81 }

// IsMirroredThermalMetering checks if MeteringDeviceType equals the value for Mirrored Thermal Metering (0x82)
func (a MeteringDeviceType) IsMirroredThermalMetering() bool { return a == 0x82 }

// SetMirroredThermalMetering sets MeteringDeviceType to Mirrored Thermal Metering (0x82)
func (a *MeteringDeviceType) SetMirroredThermalMetering() { *a = 0x82 }

// IsMirroredPressureMetering checks if MeteringDeviceType equals the value for Mirrored Pressure Metering (0x83)
func (a MeteringDeviceType) IsMirroredPressureMetering() bool { return a == 0x83 }

// SetMirroredPressureMetering sets MeteringDeviceType to Mirrored Pressure Metering (0x83)
func (a *MeteringDeviceType) SetMirroredPressureMetering() { *a = 0x83 }

// IsMirroredHeatMetering checks if MeteringDeviceType equals the value for Mirrored Heat Metering (0x84)
func (a MeteringDeviceType) IsMirroredHeatMetering() bool { return a == 0x84 }

// SetMirroredHeatMetering sets MeteringDeviceType to Mirrored Heat Metering (0x84)
func (a *MeteringDeviceType) SetMirroredHeatMetering() { *a = 0x84 }

// IsMirroredCoolingMetering checks if MeteringDeviceType equals the value for Mirrored Cooling Metering (0x85)
func (a MeteringDeviceType) IsMirroredCoolingMetering() bool { return a == 0x85 }

// SetMirroredCoolingMetering sets MeteringDeviceType to Mirrored Cooling Metering (0x85)
func (a *MeteringDeviceType) SetMirroredCoolingMetering() { *a = 0x85 }

// InstantaneousDemand is an autogenerated attribute in the SimpleMetering cluster
type InstantaneousDemand zcl.Zs24

const InstantaneousDemandAttr zcl.AttrID = 1024

func (a InstantaneousDemand) ID() zcl.AttrID               { return InstantaneousDemandAttr }
func (a InstantaneousDemand) Cluster() zcl.ClusterID       { return SimpleMeteringID }
func (a *InstantaneousDemand) Value() *InstantaneousDemand { return a }
func (a InstantaneousDemand) MarshalZcl() ([]byte, error) {
	return zcl.Zs24(a).MarshalZcl()
}
func (a *InstantaneousDemand) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs24)
	br, err := nt.UnmarshalZcl(b)
	*a = InstantaneousDemand(*nt)
	return br, err
}

func (a InstantaneousDemand) Readable() bool   { return true }
func (a InstantaneousDemand) Writable() bool   { return false }
func (a InstantaneousDemand) Reportable() bool { return false }
func (a InstantaneousDemand) SceneIndex() int  { return -1 }

func (a InstantaneousDemand) String() string {
	return zcl.Sprintf("%v", zcl.Zs24(a))
}
