// provides general/common attributes and operations for most zigbee devices
package general

import "hemtjan.st/zcl"

type CommandID = zcl.CommandID

type ApsDecryptFailures zcl.Zu16

const ApsDecryptFailuresAttr zcl.AttrID = 278

func (ApsDecryptFailures) ID() zcl.AttrID   { return ApsDecryptFailuresAttr }
func (ApsDecryptFailures) Readable() bool   { return true }
func (ApsDecryptFailures) Writable() bool   { return false }
func (ApsDecryptFailures) Reportable() bool { return false }
func (ApsDecryptFailures) SceneIndex() int  { return -1 }

func (ApsDecryptFailures) Name() string                  { return "APS Decrypt Failures" }
func (a *ApsDecryptFailures) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ApsDecryptFailures) Value() zcl.Val             { return a }
func (a ApsDecryptFailures) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApsDecryptFailures) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsDecryptFailures(*nt)
	return br, err
}

func (a *ApsDecryptFailures) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ApsDecryptFailures(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ApsDecryptFailures) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type ApsFcFailure zcl.Zu16

const ApsFcFailureAttr zcl.AttrID = 275

func (ApsFcFailure) ID() zcl.AttrID   { return ApsFcFailureAttr }
func (ApsFcFailure) Readable() bool   { return true }
func (ApsFcFailure) Writable() bool   { return false }
func (ApsFcFailure) Reportable() bool { return false }
func (ApsFcFailure) SceneIndex() int  { return -1 }

func (ApsFcFailure) Name() string                  { return "APS FC Failure" }
func (a *ApsFcFailure) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ApsFcFailure) Value() zcl.Val             { return a }
func (a ApsFcFailure) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApsFcFailure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsFcFailure(*nt)
	return br, err
}

func (a *ApsFcFailure) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ApsFcFailure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ApsFcFailure) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type ApsRxBcast zcl.Zu16

const ApsRxBcastAttr zcl.AttrID = 262

func (ApsRxBcast) ID() zcl.AttrID   { return ApsRxBcastAttr }
func (ApsRxBcast) Readable() bool   { return true }
func (ApsRxBcast) Writable() bool   { return false }
func (ApsRxBcast) Reportable() bool { return false }
func (ApsRxBcast) SceneIndex() int  { return -1 }

func (ApsRxBcast) Name() string                  { return "APS Rx Bcast" }
func (a *ApsRxBcast) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ApsRxBcast) Value() zcl.Val             { return a }
func (a ApsRxBcast) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApsRxBcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsRxBcast(*nt)
	return br, err
}

func (a *ApsRxBcast) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ApsRxBcast(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ApsRxBcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type ApsRxUcast zcl.Zu16

const ApsRxUcastAttr zcl.AttrID = 264

func (ApsRxUcast) ID() zcl.AttrID   { return ApsRxUcastAttr }
func (ApsRxUcast) Readable() bool   { return true }
func (ApsRxUcast) Writable() bool   { return false }
func (ApsRxUcast) Reportable() bool { return false }
func (ApsRxUcast) SceneIndex() int  { return -1 }

func (ApsRxUcast) Name() string                  { return "APS Rx Ucast" }
func (a *ApsRxUcast) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ApsRxUcast) Value() zcl.Val             { return a }
func (a ApsRxUcast) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApsRxUcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsRxUcast(*nt)
	return br, err
}

func (a *ApsRxUcast) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ApsRxUcast(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ApsRxUcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type ApsTxBcast zcl.Zu16

const ApsTxBcastAttr zcl.AttrID = 263

func (ApsTxBcast) ID() zcl.AttrID   { return ApsTxBcastAttr }
func (ApsTxBcast) Readable() bool   { return true }
func (ApsTxBcast) Writable() bool   { return false }
func (ApsTxBcast) Reportable() bool { return false }
func (ApsTxBcast) SceneIndex() int  { return -1 }

func (ApsTxBcast) Name() string                  { return "APS Tx Bcast" }
func (a *ApsTxBcast) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ApsTxBcast) Value() zcl.Val             { return a }
func (a ApsTxBcast) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApsTxBcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsTxBcast(*nt)
	return br, err
}

func (a *ApsTxBcast) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ApsTxBcast(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ApsTxBcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type ApsTxUcastFail zcl.Zu16

const ApsTxUcastFailAttr zcl.AttrID = 267

func (ApsTxUcastFail) ID() zcl.AttrID   { return ApsTxUcastFailAttr }
func (ApsTxUcastFail) Readable() bool   { return true }
func (ApsTxUcastFail) Writable() bool   { return false }
func (ApsTxUcastFail) Reportable() bool { return false }
func (ApsTxUcastFail) SceneIndex() int  { return -1 }

func (ApsTxUcastFail) Name() string                  { return "APS Tx Ucast Fail" }
func (a *ApsTxUcastFail) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ApsTxUcastFail) Value() zcl.Val             { return a }
func (a ApsTxUcastFail) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApsTxUcastFail) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsTxUcastFail(*nt)
	return br, err
}

func (a *ApsTxUcastFail) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ApsTxUcastFail(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ApsTxUcastFail) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type ApsTxUcastRetry zcl.Zu16

const ApsTxUcastRetryAttr zcl.AttrID = 266

func (ApsTxUcastRetry) ID() zcl.AttrID   { return ApsTxUcastRetryAttr }
func (ApsTxUcastRetry) Readable() bool   { return true }
func (ApsTxUcastRetry) Writable() bool   { return false }
func (ApsTxUcastRetry) Reportable() bool { return false }
func (ApsTxUcastRetry) SceneIndex() int  { return -1 }

func (ApsTxUcastRetry) Name() string                  { return "APS Tx Ucast Retry" }
func (a *ApsTxUcastRetry) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ApsTxUcastRetry) Value() zcl.Val             { return a }
func (a ApsTxUcastRetry) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApsTxUcastRetry) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsTxUcastRetry(*nt)
	return br, err
}

func (a *ApsTxUcastRetry) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ApsTxUcastRetry(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ApsTxUcastRetry) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type ApsTxUcastSuccess zcl.Zu16

const ApsTxUcastSuccessAttr zcl.AttrID = 265

func (ApsTxUcastSuccess) ID() zcl.AttrID   { return ApsTxUcastSuccessAttr }
func (ApsTxUcastSuccess) Readable() bool   { return true }
func (ApsTxUcastSuccess) Writable() bool   { return false }
func (ApsTxUcastSuccess) Reportable() bool { return false }
func (ApsTxUcastSuccess) SceneIndex() int  { return -1 }

func (ApsTxUcastSuccess) Name() string                  { return "APS Tx Ucast Success" }
func (a *ApsTxUcastSuccess) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ApsTxUcastSuccess) Value() zcl.Val             { return a }
func (a ApsTxUcastSuccess) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApsTxUcastSuccess) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsTxUcastSuccess(*nt)
	return br, err
}

func (a *ApsTxUcastSuccess) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ApsTxUcastSuccess(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ApsTxUcastSuccess) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type ApsUnauthorizedKey zcl.Zu16

const ApsUnauthorizedKeyAttr zcl.AttrID = 276

func (ApsUnauthorizedKey) ID() zcl.AttrID   { return ApsUnauthorizedKeyAttr }
func (ApsUnauthorizedKey) Readable() bool   { return true }
func (ApsUnauthorizedKey) Writable() bool   { return false }
func (ApsUnauthorizedKey) Reportable() bool { return false }
func (ApsUnauthorizedKey) SceneIndex() int  { return -1 }

func (ApsUnauthorizedKey) Name() string                  { return "APS Unauthorized Key" }
func (a *ApsUnauthorizedKey) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ApsUnauthorizedKey) Value() zcl.Val             { return a }
func (a ApsUnauthorizedKey) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApsUnauthorizedKey) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsUnauthorizedKey(*nt)
	return br, err
}

func (a *ApsUnauthorizedKey) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ApsUnauthorizedKey(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ApsUnauthorizedKey) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// AlarmCount Number of alarms currently defined
type AlarmCount zcl.Zu16

const AlarmCountAttr zcl.AttrID = 0

func (AlarmCount) ID() zcl.AttrID   { return AlarmCountAttr }
func (AlarmCount) Readable() bool   { return true }
func (AlarmCount) Writable() bool   { return true }
func (AlarmCount) Reportable() bool { return false }
func (AlarmCount) SceneIndex() int  { return -1 }

func (AlarmCount) Name() string                  { return "Alarm Count" }
func (a *AlarmCount) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *AlarmCount) Value() zcl.Val             { return a }
func (a AlarmCount) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *AlarmCount) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AlarmCount(*nt)
	return br, err
}

func (a *AlarmCount) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = AlarmCount(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a AlarmCount) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type AlarmMask zcl.Zbmp8

const AlarmMaskAttr zcl.AttrID = 19

func (AlarmMask) ID() zcl.AttrID   { return AlarmMaskAttr }
func (AlarmMask) Readable() bool   { return true }
func (AlarmMask) Writable() bool   { return true }
func (AlarmMask) Reportable() bool { return false }
func (AlarmMask) SceneIndex() int  { return -1 }

func (AlarmMask) Name() string                  { return "Alarm Mask" }
func (a *AlarmMask) TypeID() zcl.TypeID         { return zcl.Zbmp8(*a).ID() }
func (a *AlarmMask) Value() zcl.Val             { return a }
func (a AlarmMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *AlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = AlarmMask(*nt)
	return br, err
}

func (a *AlarmMask) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = AlarmMask(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a AlarmMask) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "General Hardware Fault")
		case 1:
			bstr = append(bstr, "General Software Fault")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a AlarmMask) IsGeneralHardwareFault() bool { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a AlarmMask) IsGeneralSoftwareFault() bool { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a *AlarmMask) SetGeneralHardwareFault(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *AlarmMask) SetGeneralSoftwareFault(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}

func (AlarmMask) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "General Hardware Fault"},
		{Value: 1, Name: "General Software Fault"},
	}
}

type AlarmCode zcl.Zenum8

func (AlarmCode) Name() string                  { return "Alarm code" }
func (a *AlarmCode) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *AlarmCode) Value() zcl.Val             { return a }
func (a AlarmCode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *AlarmCode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = AlarmCode(*nt)
	return br, err
}

func (a *AlarmCode) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = AlarmCode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a AlarmCode) String() string {
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

type AnalogMaxPresentValue zcl.Zfloat

const AnalogMaxPresentValueAttr zcl.AttrID = 65

func (AnalogMaxPresentValue) ID() zcl.AttrID   { return AnalogMaxPresentValueAttr }
func (AnalogMaxPresentValue) Readable() bool   { return true }
func (AnalogMaxPresentValue) Writable() bool   { return true }
func (AnalogMaxPresentValue) Reportable() bool { return false }
func (AnalogMaxPresentValue) SceneIndex() int  { return -1 }

func (AnalogMaxPresentValue) Name() string                  { return "Analog Max Present Value" }
func (a *AnalogMaxPresentValue) TypeID() zcl.TypeID         { return zcl.Zfloat(*a).ID() }
func (a *AnalogMaxPresentValue) Value() zcl.Val             { return a }
func (a AnalogMaxPresentValue) MarshalZcl() ([]byte, error) { return zcl.Zfloat(a).MarshalZcl() }

func (a *AnalogMaxPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogMaxPresentValue(*nt)
	return br, err
}

func (a *AnalogMaxPresentValue) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zfloat); ok {
		*a = AnalogMaxPresentValue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a AnalogMaxPresentValue) String() string {
	return zcl.Sprintf("%v", zcl.Zfloat(a))
}

type AnalogMinPresentValue zcl.Zfloat

const AnalogMinPresentValueAttr zcl.AttrID = 69

func (AnalogMinPresentValue) ID() zcl.AttrID   { return AnalogMinPresentValueAttr }
func (AnalogMinPresentValue) Readable() bool   { return true }
func (AnalogMinPresentValue) Writable() bool   { return true }
func (AnalogMinPresentValue) Reportable() bool { return false }
func (AnalogMinPresentValue) SceneIndex() int  { return -1 }

func (AnalogMinPresentValue) Name() string                  { return "Analog Min Present Value" }
func (a *AnalogMinPresentValue) TypeID() zcl.TypeID         { return zcl.Zfloat(*a).ID() }
func (a *AnalogMinPresentValue) Value() zcl.Val             { return a }
func (a AnalogMinPresentValue) MarshalZcl() ([]byte, error) { return zcl.Zfloat(a).MarshalZcl() }

func (a *AnalogMinPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogMinPresentValue(*nt)
	return br, err
}

func (a *AnalogMinPresentValue) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zfloat); ok {
		*a = AnalogMinPresentValue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a AnalogMinPresentValue) String() string {
	return zcl.Sprintf("%v", zcl.Zfloat(a))
}

type AnalogPresentValue zcl.Zfloat

const AnalogPresentValueAttr zcl.AttrID = 85

func (AnalogPresentValue) ID() zcl.AttrID   { return AnalogPresentValueAttr }
func (AnalogPresentValue) Readable() bool   { return true }
func (AnalogPresentValue) Writable() bool   { return true }
func (AnalogPresentValue) Reportable() bool { return true }
func (AnalogPresentValue) SceneIndex() int  { return -1 }

func (AnalogPresentValue) Name() string                  { return "Analog Present value" }
func (a *AnalogPresentValue) TypeID() zcl.TypeID         { return zcl.Zfloat(*a).ID() }
func (a *AnalogPresentValue) Value() zcl.Val             { return a }
func (a AnalogPresentValue) MarshalZcl() ([]byte, error) { return zcl.Zfloat(a).MarshalZcl() }

func (a *AnalogPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogPresentValue(*nt)
	return br, err
}

func (a *AnalogPresentValue) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zfloat); ok {
		*a = AnalogPresentValue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a AnalogPresentValue) String() string {
	return zcl.Sprintf("%v", zcl.Zfloat(a))
}

type AnalogPriorityArray zcl.Zarray

const AnalogPriorityArrayAttr zcl.AttrID = 87

func (AnalogPriorityArray) ID() zcl.AttrID   { return AnalogPriorityArrayAttr }
func (AnalogPriorityArray) Readable() bool   { return true }
func (AnalogPriorityArray) Writable() bool   { return true }
func (AnalogPriorityArray) Reportable() bool { return false }
func (AnalogPriorityArray) SceneIndex() int  { return -1 }

func (AnalogPriorityArray) Name() string          { return "Analog Priority Array" }
func (a *AnalogPriorityArray) TypeID() zcl.TypeID { return zcl.Zarray(*a).ID() }
func (a *AnalogPriorityArray) Value() zcl.Val     { return a }
func (a AnalogPriorityArray) MarshalZcl() ([]byte, error) {
	return nil, zcl.Errorf("not implemented")
}

func (a *AnalogPriorityArray) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zarray)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogPriorityArray(*nt)
	return br, err
}

func (a *AnalogPriorityArray) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zarray); ok {
		*a = AnalogPriorityArray(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a AnalogPriorityArray) String() string {
	return zcl.Sprintf("%v", zcl.Zarray(a))
}

type AnalogRelinquishDefault zcl.Zfloat

const AnalogRelinquishDefaultAttr zcl.AttrID = 104

func (AnalogRelinquishDefault) ID() zcl.AttrID   { return AnalogRelinquishDefaultAttr }
func (AnalogRelinquishDefault) Readable() bool   { return true }
func (AnalogRelinquishDefault) Writable() bool   { return true }
func (AnalogRelinquishDefault) Reportable() bool { return false }
func (AnalogRelinquishDefault) SceneIndex() int  { return -1 }

func (AnalogRelinquishDefault) Name() string                  { return "Analog Relinquish Default" }
func (a *AnalogRelinquishDefault) TypeID() zcl.TypeID         { return zcl.Zfloat(*a).ID() }
func (a *AnalogRelinquishDefault) Value() zcl.Val             { return a }
func (a AnalogRelinquishDefault) MarshalZcl() ([]byte, error) { return zcl.Zfloat(a).MarshalZcl() }

func (a *AnalogRelinquishDefault) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogRelinquishDefault(*nt)
	return br, err
}

func (a *AnalogRelinquishDefault) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zfloat); ok {
		*a = AnalogRelinquishDefault(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a AnalogRelinquishDefault) String() string {
	return zcl.Sprintf("%v", zcl.Zfloat(a))
}

type AnalogResolution zcl.Zfloat

const AnalogResolutionAttr zcl.AttrID = 106

func (AnalogResolution) ID() zcl.AttrID   { return AnalogResolutionAttr }
func (AnalogResolution) Readable() bool   { return true }
func (AnalogResolution) Writable() bool   { return true }
func (AnalogResolution) Reportable() bool { return false }
func (AnalogResolution) SceneIndex() int  { return -1 }

func (AnalogResolution) Name() string                  { return "Analog Resolution" }
func (a *AnalogResolution) TypeID() zcl.TypeID         { return zcl.Zfloat(*a).ID() }
func (a *AnalogResolution) Value() zcl.Val             { return a }
func (a AnalogResolution) MarshalZcl() ([]byte, error) { return zcl.Zfloat(a).MarshalZcl() }

func (a *AnalogResolution) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogResolution(*nt)
	return br, err
}

func (a *AnalogResolution) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zfloat); ok {
		*a = AnalogResolution(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a AnalogResolution) String() string {
	return zcl.Sprintf("%v", zcl.Zfloat(a))
}

type ApplicationVersion zcl.Zu8

const ApplicationVersionAttr zcl.AttrID = 1

func (ApplicationVersion) ID() zcl.AttrID   { return ApplicationVersionAttr }
func (ApplicationVersion) Readable() bool   { return true }
func (ApplicationVersion) Writable() bool   { return false }
func (ApplicationVersion) Reportable() bool { return false }
func (ApplicationVersion) SceneIndex() int  { return -1 }

func (ApplicationVersion) Name() string                  { return "Application Version" }
func (a *ApplicationVersion) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *ApplicationVersion) Value() zcl.Val             { return a }
func (a ApplicationVersion) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *ApplicationVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ApplicationVersion(*nt)
	return br, err
}

func (a *ApplicationVersion) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = ApplicationVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ApplicationVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type AvgMacRetryPerApsMsgSent zcl.Zu16

const AvgMacRetryPerApsMsgSentAttr zcl.AttrID = 283

func (AvgMacRetryPerApsMsgSent) ID() zcl.AttrID   { return AvgMacRetryPerApsMsgSentAttr }
func (AvgMacRetryPerApsMsgSent) Readable() bool   { return true }
func (AvgMacRetryPerApsMsgSent) Writable() bool   { return false }
func (AvgMacRetryPerApsMsgSent) Reportable() bool { return false }
func (AvgMacRetryPerApsMsgSent) SceneIndex() int  { return -1 }

func (AvgMacRetryPerApsMsgSent) Name() string                  { return "Avg MAC Retry per APS Msg Sent" }
func (a *AvgMacRetryPerApsMsgSent) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *AvgMacRetryPerApsMsgSent) Value() zcl.Val             { return a }
func (a AvgMacRetryPerApsMsgSent) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *AvgMacRetryPerApsMsgSent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AvgMacRetryPerApsMsgSent(*nt)
	return br, err
}

func (a *AvgMacRetryPerApsMsgSent) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = AvgMacRetryPerApsMsgSent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a AvgMacRetryPerApsMsgSent) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type BatteryAlarmMask zcl.Zbmp8

const BatteryAlarmMaskAttr zcl.AttrID = 53

func (BatteryAlarmMask) ID() zcl.AttrID   { return BatteryAlarmMaskAttr }
func (BatteryAlarmMask) Readable() bool   { return true }
func (BatteryAlarmMask) Writable() bool   { return true }
func (BatteryAlarmMask) Reportable() bool { return false }
func (BatteryAlarmMask) SceneIndex() int  { return -1 }

func (BatteryAlarmMask) Name() string                  { return "Battery Alarm Mask" }
func (a *BatteryAlarmMask) TypeID() zcl.TypeID         { return zcl.Zbmp8(*a).ID() }
func (a *BatteryAlarmMask) Value() zcl.Val             { return a }
func (a BatteryAlarmMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *BatteryAlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryAlarmMask(*nt)
	return br, err
}

func (a *BatteryAlarmMask) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = BatteryAlarmMask(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryAlarmMask) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Battery Voltage too low")
		case 1:
			bstr = append(bstr, "Battery Alarm 1")
		case 2:
			bstr = append(bstr, "Battery Alarm 2")
		case 3:
			bstr = append(bstr, "Battery Alarm 3")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a BatteryAlarmMask) IsBatteryVoltageTooLow() bool { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a BatteryAlarmMask) IsBatteryAlarm1() bool        { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a BatteryAlarmMask) IsBatteryAlarm2() bool        { return zcl.BitmapTest([]byte(a[:]), 2) }
func (a BatteryAlarmMask) IsBatteryAlarm3() bool        { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a *BatteryAlarmMask) SetBatteryVoltageTooLow(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *BatteryAlarmMask) SetBatteryAlarm1(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}
func (a *BatteryAlarmMask) SetBatteryAlarm2(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b))
}
func (a *BatteryAlarmMask) SetBatteryAlarm3(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b))
}

func (BatteryAlarmMask) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Battery Voltage too low"},
		{Value: 1, Name: "Battery Alarm 1"},
		{Value: 2, Name: "Battery Alarm 2"},
		{Value: 3, Name: "Battery Alarm 3"},
	}
}

type BatteryAlarmState zcl.Zbmp32

const BatteryAlarmStateAttr zcl.AttrID = 62

func (BatteryAlarmState) ID() zcl.AttrID   { return BatteryAlarmStateAttr }
func (BatteryAlarmState) Readable() bool   { return true }
func (BatteryAlarmState) Writable() bool   { return true }
func (BatteryAlarmState) Reportable() bool { return false }
func (BatteryAlarmState) SceneIndex() int  { return -1 }

func (BatteryAlarmState) Name() string                  { return "Battery Alarm State" }
func (a *BatteryAlarmState) TypeID() zcl.TypeID         { return zcl.Zbmp32(*a).ID() }
func (a *BatteryAlarmState) Value() zcl.Val             { return a }
func (a BatteryAlarmState) MarshalZcl() ([]byte, error) { return zcl.Zbmp32(a).MarshalZcl() }

func (a *BatteryAlarmState) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp32)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryAlarmState(*nt)
	return br, err
}

func (a *BatteryAlarmState) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp32); ok {
		*a = BatteryAlarmState(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryAlarmState) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Minimum Threshold 1 for Voltage or Percentage reached for Battery Source 1")
		case 1:
			bstr = append(bstr, "Threshold 1 for Voltage or Percentage reached for Battery Source 1")
		case 10:
			bstr = append(bstr, "Minimum Threshold 1 for Voltage or Percentage reached for Battery Source 2")
		case 11:
			bstr = append(bstr, "Threshold 1 for Voltage or Percentage reached for Battery Source 2")
		case 12:
			bstr = append(bstr, "Threshold 2 for Voltage or Percentage reached for Battery Source 2")
		case 13:
			bstr = append(bstr, "Threshold 3 for Voltage or Percentage reached for Battery Source 2")
		case 2:
			bstr = append(bstr, "Threshold 2 for Voltage or Percentage reached for Battery Source 1")
		case 20:
			bstr = append(bstr, "Minimum Threshold 1 for Voltage or Percentage reached for Battery Source 3")
		case 21:
			bstr = append(bstr, "Threshold 1 for Voltage or Percentage reached for Battery Source 3")
		case 22:
			bstr = append(bstr, "Threshold 2 for Voltage or Percentage reached for Battery Source 3")
		case 23:
			bstr = append(bstr, "Threshold 3 for Voltage or Percentage reached for Battery Source 3")
		case 3:
			bstr = append(bstr, "Threshold 3 for Voltage or Percentage reached for Battery Source 1")
		case 30:
			bstr = append(bstr, "Mains power lost/unavailable")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a BatteryAlarmState) IsMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource1() bool {
	return zcl.BitmapTest([]byte(a[:]), 0)
}
func (a BatteryAlarmState) IsThreshold1ForVoltageOrPercentageReachedForBatterySource1() bool {
	return zcl.BitmapTest([]byte(a[:]), 1)
}
func (a BatteryAlarmState) IsMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource2() bool {
	return zcl.BitmapTest([]byte(a[:]), 10)
}
func (a BatteryAlarmState) IsThreshold1ForVoltageOrPercentageReachedForBatterySource2() bool {
	return zcl.BitmapTest([]byte(a[:]), 11)
}
func (a BatteryAlarmState) IsThreshold2ForVoltageOrPercentageReachedForBatterySource2() bool {
	return zcl.BitmapTest([]byte(a[:]), 12)
}
func (a BatteryAlarmState) IsThreshold3ForVoltageOrPercentageReachedForBatterySource2() bool {
	return zcl.BitmapTest([]byte(a[:]), 13)
}
func (a BatteryAlarmState) IsThreshold2ForVoltageOrPercentageReachedForBatterySource1() bool {
	return zcl.BitmapTest([]byte(a[:]), 2)
}
func (a BatteryAlarmState) IsMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource3() bool {
	return zcl.BitmapTest([]byte(a[:]), 20)
}
func (a BatteryAlarmState) IsThreshold1ForVoltageOrPercentageReachedForBatterySource3() bool {
	return zcl.BitmapTest([]byte(a[:]), 21)
}
func (a BatteryAlarmState) IsThreshold2ForVoltageOrPercentageReachedForBatterySource3() bool {
	return zcl.BitmapTest([]byte(a[:]), 22)
}
func (a BatteryAlarmState) IsThreshold3ForVoltageOrPercentageReachedForBatterySource3() bool {
	return zcl.BitmapTest([]byte(a[:]), 23)
}
func (a BatteryAlarmState) IsThreshold3ForVoltageOrPercentageReachedForBatterySource1() bool {
	return zcl.BitmapTest([]byte(a[:]), 3)
}
func (a BatteryAlarmState) IsMainsPowerLostUnavailable() bool { return zcl.BitmapTest([]byte(a[:]), 30) }
func (a *BatteryAlarmState) SetMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource1(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *BatteryAlarmState) SetThreshold1ForVoltageOrPercentageReachedForBatterySource1(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}
func (a *BatteryAlarmState) SetMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource2(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 10, b))
}
func (a *BatteryAlarmState) SetThreshold1ForVoltageOrPercentageReachedForBatterySource2(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 11, b))
}
func (a *BatteryAlarmState) SetThreshold2ForVoltageOrPercentageReachedForBatterySource2(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 12, b))
}
func (a *BatteryAlarmState) SetThreshold3ForVoltageOrPercentageReachedForBatterySource2(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 13, b))
}
func (a *BatteryAlarmState) SetThreshold2ForVoltageOrPercentageReachedForBatterySource1(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b))
}
func (a *BatteryAlarmState) SetMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource3(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 20, b))
}
func (a *BatteryAlarmState) SetThreshold1ForVoltageOrPercentageReachedForBatterySource3(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 21, b))
}
func (a *BatteryAlarmState) SetThreshold2ForVoltageOrPercentageReachedForBatterySource3(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 22, b))
}
func (a *BatteryAlarmState) SetThreshold3ForVoltageOrPercentageReachedForBatterySource3(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 23, b))
}
func (a *BatteryAlarmState) SetThreshold3ForVoltageOrPercentageReachedForBatterySource1(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b))
}
func (a *BatteryAlarmState) SetMainsPowerLostUnavailable(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 30, b))
}

func (BatteryAlarmState) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Minimum Threshold 1 for Voltage or Percentage reached for Battery Source 1"},
		{Value: 1, Name: "Threshold 1 for Voltage or Percentage reached for Battery Source 1"},
		{Value: 10, Name: "Minimum Threshold 1 for Voltage or Percentage reached for Battery Source 2"},
		{Value: 11, Name: "Threshold 1 for Voltage or Percentage reached for Battery Source 2"},
		{Value: 12, Name: "Threshold 2 for Voltage or Percentage reached for Battery Source 2"},
		{Value: 13, Name: "Threshold 3 for Voltage or Percentage reached for Battery Source 2"},
		{Value: 2, Name: "Threshold 2 for Voltage or Percentage reached for Battery Source 1"},
		{Value: 20, Name: "Minimum Threshold 1 for Voltage or Percentage reached for Battery Source 3"},
		{Value: 21, Name: "Threshold 1 for Voltage or Percentage reached for Battery Source 3"},
		{Value: 22, Name: "Threshold 2 for Voltage or Percentage reached for Battery Source 3"},
		{Value: 23, Name: "Threshold 3 for Voltage or Percentage reached for Battery Source 3"},
		{Value: 3, Name: "Threshold 3 for Voltage or Percentage reached for Battery Source 1"},
		{Value: 30, Name: "Mains power lost/unavailable"},
	}
}

type BatteryManufacturer zcl.Zcstring

const BatteryManufacturerAttr zcl.AttrID = 48

func (BatteryManufacturer) ID() zcl.AttrID   { return BatteryManufacturerAttr }
func (BatteryManufacturer) Readable() bool   { return true }
func (BatteryManufacturer) Writable() bool   { return true }
func (BatteryManufacturer) Reportable() bool { return false }
func (BatteryManufacturer) SceneIndex() int  { return -1 }

func (BatteryManufacturer) Name() string                  { return "Battery Manufacturer" }
func (a *BatteryManufacturer) TypeID() zcl.TypeID         { return zcl.Zcstring(*a).ID() }
func (a *BatteryManufacturer) Value() zcl.Val             { return a }
func (a BatteryManufacturer) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *BatteryManufacturer) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryManufacturer(*nt)
	return br, err
}

func (a *BatteryManufacturer) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = BatteryManufacturer(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryManufacturer) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

type BatteryPercentageMinThreshold zcl.Zu8

const BatteryPercentageMinThresholdAttr zcl.AttrID = 58

func (BatteryPercentageMinThreshold) ID() zcl.AttrID   { return BatteryPercentageMinThresholdAttr }
func (BatteryPercentageMinThreshold) Readable() bool   { return true }
func (BatteryPercentageMinThreshold) Writable() bool   { return true }
func (BatteryPercentageMinThreshold) Reportable() bool { return false }
func (BatteryPercentageMinThreshold) SceneIndex() int  { return -1 }

func (BatteryPercentageMinThreshold) Name() string                  { return "Battery Percentage Min Threshold" }
func (a *BatteryPercentageMinThreshold) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *BatteryPercentageMinThreshold) Value() zcl.Val             { return a }
func (a BatteryPercentageMinThreshold) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryPercentageMinThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryPercentageMinThreshold(*nt)
	return br, err
}

func (a *BatteryPercentageMinThreshold) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryPercentageMinThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryPercentageMinThreshold) String() string {
	return zcl.Percent.Format(float64(a) / 2)
}

type BatteryPercentageThreshold1 zcl.Zu8

const BatteryPercentageThreshold1Attr zcl.AttrID = 59

func (BatteryPercentageThreshold1) ID() zcl.AttrID   { return BatteryPercentageThreshold1Attr }
func (BatteryPercentageThreshold1) Readable() bool   { return true }
func (BatteryPercentageThreshold1) Writable() bool   { return true }
func (BatteryPercentageThreshold1) Reportable() bool { return false }
func (BatteryPercentageThreshold1) SceneIndex() int  { return -1 }

func (BatteryPercentageThreshold1) Name() string                  { return "Battery Percentage Threshold 1" }
func (a *BatteryPercentageThreshold1) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *BatteryPercentageThreshold1) Value() zcl.Val             { return a }
func (a BatteryPercentageThreshold1) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryPercentageThreshold1) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryPercentageThreshold1(*nt)
	return br, err
}

func (a *BatteryPercentageThreshold1) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryPercentageThreshold1(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryPercentageThreshold1) String() string {
	return zcl.Percent.Format(float64(a) / 2)
}

type BatteryPercentageThreshold2 zcl.Zu8

const BatteryPercentageThreshold2Attr zcl.AttrID = 60

func (BatteryPercentageThreshold2) ID() zcl.AttrID   { return BatteryPercentageThreshold2Attr }
func (BatteryPercentageThreshold2) Readable() bool   { return true }
func (BatteryPercentageThreshold2) Writable() bool   { return true }
func (BatteryPercentageThreshold2) Reportable() bool { return false }
func (BatteryPercentageThreshold2) SceneIndex() int  { return -1 }

func (BatteryPercentageThreshold2) Name() string                  { return "Battery Percentage Threshold 2" }
func (a *BatteryPercentageThreshold2) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *BatteryPercentageThreshold2) Value() zcl.Val             { return a }
func (a BatteryPercentageThreshold2) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryPercentageThreshold2) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryPercentageThreshold2(*nt)
	return br, err
}

func (a *BatteryPercentageThreshold2) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryPercentageThreshold2(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryPercentageThreshold2) String() string {
	return zcl.Percent.Format(float64(a) / 2)
}

type BatteryPercentageThreshold3 zcl.Zu8

const BatteryPercentageThreshold3Attr zcl.AttrID = 61

func (BatteryPercentageThreshold3) ID() zcl.AttrID   { return BatteryPercentageThreshold3Attr }
func (BatteryPercentageThreshold3) Readable() bool   { return true }
func (BatteryPercentageThreshold3) Writable() bool   { return true }
func (BatteryPercentageThreshold3) Reportable() bool { return false }
func (BatteryPercentageThreshold3) SceneIndex() int  { return -1 }

func (BatteryPercentageThreshold3) Name() string                  { return "Battery Percentage Threshold 3" }
func (a *BatteryPercentageThreshold3) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *BatteryPercentageThreshold3) Value() zcl.Val             { return a }
func (a BatteryPercentageThreshold3) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryPercentageThreshold3) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryPercentageThreshold3(*nt)
	return br, err
}

func (a *BatteryPercentageThreshold3) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryPercentageThreshold3(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryPercentageThreshold3) String() string {
	return zcl.Percent.Format(float64(a) / 2)
}

type BatteryQuantity zcl.Zu8

const BatteryQuantityAttr zcl.AttrID = 51

func (BatteryQuantity) ID() zcl.AttrID   { return BatteryQuantityAttr }
func (BatteryQuantity) Readable() bool   { return true }
func (BatteryQuantity) Writable() bool   { return true }
func (BatteryQuantity) Reportable() bool { return false }
func (BatteryQuantity) SceneIndex() int  { return -1 }

func (BatteryQuantity) Name() string                  { return "Battery Quantity" }
func (a *BatteryQuantity) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *BatteryQuantity) Value() zcl.Val             { return a }
func (a BatteryQuantity) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryQuantity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryQuantity(*nt)
	return br, err
}

func (a *BatteryQuantity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryQuantity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryQuantity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type BatteryRatedVoltage zcl.Zu8

const BatteryRatedVoltageAttr zcl.AttrID = 52

func (BatteryRatedVoltage) ID() zcl.AttrID   { return BatteryRatedVoltageAttr }
func (BatteryRatedVoltage) Readable() bool   { return true }
func (BatteryRatedVoltage) Writable() bool   { return true }
func (BatteryRatedVoltage) Reportable() bool { return false }
func (BatteryRatedVoltage) SceneIndex() int  { return -1 }

func (BatteryRatedVoltage) Name() string                  { return "Battery Rated Voltage" }
func (a *BatteryRatedVoltage) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *BatteryRatedVoltage) Value() zcl.Val             { return a }
func (a BatteryRatedVoltage) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryRatedVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryRatedVoltage(*nt)
	return br, err
}

func (a *BatteryRatedVoltage) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryRatedVoltage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryRatedVoltage) String() string {
	return zcl.Volts.Format(float64(a) / 10)
}

type BatteryRemaining zcl.Zu8

const BatteryRemainingAttr zcl.AttrID = 33

func (BatteryRemaining) ID() zcl.AttrID   { return BatteryRemainingAttr }
func (BatteryRemaining) Readable() bool   { return true }
func (BatteryRemaining) Writable() bool   { return false }
func (BatteryRemaining) Reportable() bool { return true }
func (BatteryRemaining) SceneIndex() int  { return -1 }

func (BatteryRemaining) Name() string                  { return "Battery Remaining" }
func (a *BatteryRemaining) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *BatteryRemaining) Value() zcl.Val             { return a }
func (a BatteryRemaining) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryRemaining) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryRemaining(*nt)
	return br, err
}

func (a *BatteryRemaining) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryRemaining(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryRemaining) String() string {
	return zcl.Percent.Format(float64(a) / 2)
}

type BatterySize zcl.Zenum8

const BatterySizeAttr zcl.AttrID = 49

func (BatterySize) ID() zcl.AttrID   { return BatterySizeAttr }
func (BatterySize) Readable() bool   { return true }
func (BatterySize) Writable() bool   { return true }
func (BatterySize) Reportable() bool { return false }
func (BatterySize) SceneIndex() int  { return -1 }

func (BatterySize) Name() string                  { return "Battery Size" }
func (a *BatterySize) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *BatterySize) Value() zcl.Val             { return a }
func (a BatterySize) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *BatterySize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatterySize(*nt)
	return br, err
}

func (a *BatterySize) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = BatterySize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatterySize) String() string {
	switch a {
	case 0x00:
		return "No Battery"
	case 0x01:
		return "Built in"
	case 0x02:
		return "Other"
	case 0x03:
		return "AA"
	case 0x04:
		return "AAA"
	case 0x05:
		return "C"
	case 0x06:
		return "D"
	case 0x07:
		return "CR2"
	case 0x08:
		return "CR123A"
	case 0xFF:
		return "Unknown"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a BatterySize) IsNoBattery() bool { return a == 0x00 }
func (a BatterySize) IsBuiltIn() bool   { return a == 0x01 }
func (a BatterySize) IsOther() bool     { return a == 0x02 }
func (a BatterySize) IsAa() bool        { return a == 0x03 }
func (a BatterySize) IsAaa() bool       { return a == 0x04 }
func (a BatterySize) IsC() bool         { return a == 0x05 }
func (a BatterySize) IsD() bool         { return a == 0x06 }
func (a BatterySize) IsCr2() bool       { return a == 0x07 }
func (a BatterySize) IsCr123a() bool    { return a == 0x08 }
func (a BatterySize) IsUnknown() bool   { return a == 0xFF }
func (a *BatterySize) SetNoBattery()    { *a = 0x00 }
func (a *BatterySize) SetBuiltIn()      { *a = 0x01 }
func (a *BatterySize) SetOther()        { *a = 0x02 }
func (a *BatterySize) SetAa()           { *a = 0x03 }
func (a *BatterySize) SetAaa()          { *a = 0x04 }
func (a *BatterySize) SetC()            { *a = 0x05 }
func (a *BatterySize) SetD()            { *a = 0x06 }
func (a *BatterySize) SetCr2()          { *a = 0x07 }
func (a *BatterySize) SetCr123a()       { *a = 0x08 }
func (a *BatterySize) SetUnknown()      { *a = 0xFF }

func (BatterySize) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "No Battery"},
		{Value: 0x01, Name: "Built in"},
		{Value: 0x02, Name: "Other"},
		{Value: 0x03, Name: "AA"},
		{Value: 0x04, Name: "AAA"},
		{Value: 0x05, Name: "C"},
		{Value: 0x06, Name: "D"},
		{Value: 0x07, Name: "CR2"},
		{Value: 0x08, Name: "CR123A"},
		{Value: 0xFF, Name: "Unknown"},
	}
}

type BatteryVoltage zcl.Zu8

const BatteryVoltageAttr zcl.AttrID = 32

func (BatteryVoltage) ID() zcl.AttrID   { return BatteryVoltageAttr }
func (BatteryVoltage) Readable() bool   { return true }
func (BatteryVoltage) Writable() bool   { return false }
func (BatteryVoltage) Reportable() bool { return false }
func (BatteryVoltage) SceneIndex() int  { return -1 }

func (BatteryVoltage) Name() string                  { return "Battery Voltage" }
func (a *BatteryVoltage) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *BatteryVoltage) Value() zcl.Val             { return a }
func (a BatteryVoltage) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryVoltage(*nt)
	return br, err
}

func (a *BatteryVoltage) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryVoltage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryVoltage) String() string {
	return zcl.Volts.Format(float64(a) / 10)
}

type BatteryVoltageMinThreshold zcl.Zu8

const BatteryVoltageMinThresholdAttr zcl.AttrID = 54

func (BatteryVoltageMinThreshold) ID() zcl.AttrID   { return BatteryVoltageMinThresholdAttr }
func (BatteryVoltageMinThreshold) Readable() bool   { return true }
func (BatteryVoltageMinThreshold) Writable() bool   { return true }
func (BatteryVoltageMinThreshold) Reportable() bool { return false }
func (BatteryVoltageMinThreshold) SceneIndex() int  { return -1 }

func (BatteryVoltageMinThreshold) Name() string                  { return "Battery Voltage Min Threshold" }
func (a *BatteryVoltageMinThreshold) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *BatteryVoltageMinThreshold) Value() zcl.Val             { return a }
func (a BatteryVoltageMinThreshold) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryVoltageMinThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryVoltageMinThreshold(*nt)
	return br, err
}

func (a *BatteryVoltageMinThreshold) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryVoltageMinThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryVoltageMinThreshold) String() string {
	return zcl.Volts.Format(float64(a) / 10)
}

type BatteryVoltageThreshold1 zcl.Zu8

const BatteryVoltageThreshold1Attr zcl.AttrID = 55

func (BatteryVoltageThreshold1) ID() zcl.AttrID   { return BatteryVoltageThreshold1Attr }
func (BatteryVoltageThreshold1) Readable() bool   { return true }
func (BatteryVoltageThreshold1) Writable() bool   { return true }
func (BatteryVoltageThreshold1) Reportable() bool { return false }
func (BatteryVoltageThreshold1) SceneIndex() int  { return -1 }

func (BatteryVoltageThreshold1) Name() string                  { return "Battery Voltage Threshold 1" }
func (a *BatteryVoltageThreshold1) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *BatteryVoltageThreshold1) Value() zcl.Val             { return a }
func (a BatteryVoltageThreshold1) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryVoltageThreshold1) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryVoltageThreshold1(*nt)
	return br, err
}

func (a *BatteryVoltageThreshold1) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryVoltageThreshold1(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryVoltageThreshold1) String() string {
	return zcl.Volts.Format(float64(a) / 10)
}

type BatteryVoltageThreshold2 zcl.Zu8

const BatteryVoltageThreshold2Attr zcl.AttrID = 56

func (BatteryVoltageThreshold2) ID() zcl.AttrID   { return BatteryVoltageThreshold2Attr }
func (BatteryVoltageThreshold2) Readable() bool   { return true }
func (BatteryVoltageThreshold2) Writable() bool   { return true }
func (BatteryVoltageThreshold2) Reportable() bool { return false }
func (BatteryVoltageThreshold2) SceneIndex() int  { return -1 }

func (BatteryVoltageThreshold2) Name() string                  { return "Battery Voltage Threshold 2" }
func (a *BatteryVoltageThreshold2) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *BatteryVoltageThreshold2) Value() zcl.Val             { return a }
func (a BatteryVoltageThreshold2) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryVoltageThreshold2) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryVoltageThreshold2(*nt)
	return br, err
}

func (a *BatteryVoltageThreshold2) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryVoltageThreshold2(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryVoltageThreshold2) String() string {
	return zcl.Volts.Format(float64(a) / 10)
}

type BatteryVoltageThreshold3 zcl.Zu8

const BatteryVoltageThreshold3Attr zcl.AttrID = 57

func (BatteryVoltageThreshold3) ID() zcl.AttrID   { return BatteryVoltageThreshold3Attr }
func (BatteryVoltageThreshold3) Readable() bool   { return true }
func (BatteryVoltageThreshold3) Writable() bool   { return true }
func (BatteryVoltageThreshold3) Reportable() bool { return false }
func (BatteryVoltageThreshold3) SceneIndex() int  { return -1 }

func (BatteryVoltageThreshold3) Name() string                  { return "Battery Voltage Threshold 3" }
func (a *BatteryVoltageThreshold3) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *BatteryVoltageThreshold3) Value() zcl.Val             { return a }
func (a BatteryVoltageThreshold3) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryVoltageThreshold3) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryVoltageThreshold3(*nt)
	return br, err
}

func (a *BatteryVoltageThreshold3) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryVoltageThreshold3(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryVoltageThreshold3) String() string {
	return zcl.Volts.Format(float64(a) / 10)
}

type BatteryCapacity zcl.Zu16

const BatteryCapacityAttr zcl.AttrID = 50

func (BatteryCapacity) ID() zcl.AttrID   { return BatteryCapacityAttr }
func (BatteryCapacity) Readable() bool   { return true }
func (BatteryCapacity) Writable() bool   { return true }
func (BatteryCapacity) Reportable() bool { return false }
func (BatteryCapacity) SceneIndex() int  { return -1 }

func (BatteryCapacity) Name() string                  { return "Battery capacity" }
func (a *BatteryCapacity) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *BatteryCapacity) Value() zcl.Val             { return a }
func (a BatteryCapacity) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *BatteryCapacity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryCapacity(*nt)
	return br, err
}

func (a *BatteryCapacity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = BatteryCapacity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryCapacity) String() string {
	return zcl.MilliAmpereHours.Format(float64(a) / 0.1)
}

type BinaryActiveText zcl.Zcstring

const BinaryActiveTextAttr zcl.AttrID = 4

func (BinaryActiveText) ID() zcl.AttrID   { return BinaryActiveTextAttr }
func (BinaryActiveText) Readable() bool   { return true }
func (BinaryActiveText) Writable() bool   { return true }
func (BinaryActiveText) Reportable() bool { return false }
func (BinaryActiveText) SceneIndex() int  { return -1 }

func (BinaryActiveText) Name() string                  { return "Binary Active Text" }
func (a *BinaryActiveText) TypeID() zcl.TypeID         { return zcl.Zcstring(*a).ID() }
func (a *BinaryActiveText) Value() zcl.Val             { return a }
func (a BinaryActiveText) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *BinaryActiveText) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = BinaryActiveText(*nt)
	return br, err
}

func (a *BinaryActiveText) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = BinaryActiveText(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BinaryActiveText) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

type BinaryInactiveText zcl.Zcstring

const BinaryInactiveTextAttr zcl.AttrID = 46

func (BinaryInactiveText) ID() zcl.AttrID   { return BinaryInactiveTextAttr }
func (BinaryInactiveText) Readable() bool   { return true }
func (BinaryInactiveText) Writable() bool   { return true }
func (BinaryInactiveText) Reportable() bool { return false }
func (BinaryInactiveText) SceneIndex() int  { return -1 }

func (BinaryInactiveText) Name() string                  { return "Binary Inactive Text" }
func (a *BinaryInactiveText) TypeID() zcl.TypeID         { return zcl.Zcstring(*a).ID() }
func (a *BinaryInactiveText) Value() zcl.Val             { return a }
func (a BinaryInactiveText) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *BinaryInactiveText) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = BinaryInactiveText(*nt)
	return br, err
}

func (a *BinaryInactiveText) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = BinaryInactiveText(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BinaryInactiveText) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

type BinaryMaxOffTime zcl.Zu32

const BinaryMaxOffTimeAttr zcl.AttrID = 67

func (BinaryMaxOffTime) ID() zcl.AttrID   { return BinaryMaxOffTimeAttr }
func (BinaryMaxOffTime) Readable() bool   { return true }
func (BinaryMaxOffTime) Writable() bool   { return true }
func (BinaryMaxOffTime) Reportable() bool { return false }
func (BinaryMaxOffTime) SceneIndex() int  { return -1 }

func (BinaryMaxOffTime) Name() string                  { return "Binary Max Off-time" }
func (a *BinaryMaxOffTime) TypeID() zcl.TypeID         { return zcl.Zu32(*a).ID() }
func (a *BinaryMaxOffTime) Value() zcl.Val             { return a }
func (a BinaryMaxOffTime) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *BinaryMaxOffTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = BinaryMaxOffTime(*nt)
	return br, err
}

func (a *BinaryMaxOffTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu32); ok {
		*a = BinaryMaxOffTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BinaryMaxOffTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

type BinaryMinOffTime zcl.Zu32

const BinaryMinOffTimeAttr zcl.AttrID = 66

func (BinaryMinOffTime) ID() zcl.AttrID   { return BinaryMinOffTimeAttr }
func (BinaryMinOffTime) Readable() bool   { return true }
func (BinaryMinOffTime) Writable() bool   { return true }
func (BinaryMinOffTime) Reportable() bool { return false }
func (BinaryMinOffTime) SceneIndex() int  { return -1 }

func (BinaryMinOffTime) Name() string                  { return "Binary Min Off-time" }
func (a *BinaryMinOffTime) TypeID() zcl.TypeID         { return zcl.Zu32(*a).ID() }
func (a *BinaryMinOffTime) Value() zcl.Val             { return a }
func (a BinaryMinOffTime) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *BinaryMinOffTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = BinaryMinOffTime(*nt)
	return br, err
}

func (a *BinaryMinOffTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu32); ok {
		*a = BinaryMinOffTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BinaryMinOffTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

type BinaryPolarity zcl.Zenum8

const BinaryPolarityAttr zcl.AttrID = 84

func (BinaryPolarity) ID() zcl.AttrID   { return BinaryPolarityAttr }
func (BinaryPolarity) Readable() bool   { return true }
func (BinaryPolarity) Writable() bool   { return false }
func (BinaryPolarity) Reportable() bool { return false }
func (BinaryPolarity) SceneIndex() int  { return -1 }

func (BinaryPolarity) Name() string                  { return "Binary Polarity" }
func (a *BinaryPolarity) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *BinaryPolarity) Value() zcl.Val             { return a }
func (a BinaryPolarity) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *BinaryPolarity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = BinaryPolarity(*nt)
	return br, err
}

func (a *BinaryPolarity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = BinaryPolarity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BinaryPolarity) String() string {
	switch a {
	case 0x00:
		return "Normal"
	case 0x01:
		return "Reverse"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a BinaryPolarity) IsNormal() bool  { return a == 0x00 }
func (a BinaryPolarity) IsReverse() bool { return a == 0x01 }
func (a *BinaryPolarity) SetNormal()     { *a = 0x00 }
func (a *BinaryPolarity) SetReverse()    { *a = 0x01 }

func (BinaryPolarity) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Normal"},
		{Value: 0x01, Name: "Reverse"},
	}
}

type BinaryPresentValue zcl.Zbool

const BinaryPresentValueAttr zcl.AttrID = 85

func (BinaryPresentValue) ID() zcl.AttrID   { return BinaryPresentValueAttr }
func (BinaryPresentValue) Readable() bool   { return true }
func (BinaryPresentValue) Writable() bool   { return true }
func (BinaryPresentValue) Reportable() bool { return true }
func (BinaryPresentValue) SceneIndex() int  { return -1 }

func (BinaryPresentValue) Name() string                  { return "Binary Present Value" }
func (a *BinaryPresentValue) TypeID() zcl.TypeID         { return zcl.Zbool(*a).ID() }
func (a *BinaryPresentValue) Value() zcl.Val             { return a }
func (a BinaryPresentValue) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *BinaryPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = BinaryPresentValue(*nt)
	return br, err
}

func (a *BinaryPresentValue) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbool); ok {
		*a = BinaryPresentValue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BinaryPresentValue) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

type BinaryPriorityArray zcl.Zarray

const BinaryPriorityArrayAttr zcl.AttrID = 87

func (BinaryPriorityArray) ID() zcl.AttrID   { return BinaryPriorityArrayAttr }
func (BinaryPriorityArray) Readable() bool   { return true }
func (BinaryPriorityArray) Writable() bool   { return true }
func (BinaryPriorityArray) Reportable() bool { return false }
func (BinaryPriorityArray) SceneIndex() int  { return -1 }

func (BinaryPriorityArray) Name() string          { return "Binary Priority Array" }
func (a *BinaryPriorityArray) TypeID() zcl.TypeID { return zcl.Zarray(*a).ID() }
func (a *BinaryPriorityArray) Value() zcl.Val     { return a }
func (a BinaryPriorityArray) MarshalZcl() ([]byte, error) {
	return nil, zcl.Errorf("not implemented")
}

func (a *BinaryPriorityArray) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zarray)
	br, err := nt.UnmarshalZcl(b)
	*a = BinaryPriorityArray(*nt)
	return br, err
}

func (a *BinaryPriorityArray) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zarray); ok {
		*a = BinaryPriorityArray(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BinaryPriorityArray) String() string {
	return zcl.Sprintf("%v", zcl.Zarray(a))
}

type BinaryRelinquishDefault zcl.Zbool

const BinaryRelinquishDefaultAttr zcl.AttrID = 104

func (BinaryRelinquishDefault) ID() zcl.AttrID   { return BinaryRelinquishDefaultAttr }
func (BinaryRelinquishDefault) Readable() bool   { return true }
func (BinaryRelinquishDefault) Writable() bool   { return true }
func (BinaryRelinquishDefault) Reportable() bool { return false }
func (BinaryRelinquishDefault) SceneIndex() int  { return -1 }

func (BinaryRelinquishDefault) Name() string                  { return "Binary Relinquish Default" }
func (a *BinaryRelinquishDefault) TypeID() zcl.TypeID         { return zcl.Zbool(*a).ID() }
func (a *BinaryRelinquishDefault) Value() zcl.Val             { return a }
func (a BinaryRelinquishDefault) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *BinaryRelinquishDefault) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = BinaryRelinquishDefault(*nt)
	return br, err
}

func (a *BinaryRelinquishDefault) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbool); ok {
		*a = BinaryRelinquishDefault(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BinaryRelinquishDefault) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

type CalculationPeriod zcl.Zu16

const CalculationPeriodAttr zcl.AttrID = 22

func (CalculationPeriod) ID() zcl.AttrID   { return CalculationPeriodAttr }
func (CalculationPeriod) Readable() bool   { return true }
func (CalculationPeriod) Writable() bool   { return true }
func (CalculationPeriod) Reportable() bool { return false }
func (CalculationPeriod) SceneIndex() int  { return -1 }

func (CalculationPeriod) Name() string                  { return "Calculation Period" }
func (a *CalculationPeriod) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *CalculationPeriod) Value() zcl.Val             { return a }
func (a CalculationPeriod) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *CalculationPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = CalculationPeriod(*nt)
	return br, err
}

func (a *CalculationPeriod) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = CalculationPeriod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CalculationPeriod) String() string {
	return zcl.Seconds.Format(float64(a))
}

type CheckInInterval zcl.Zu32

const CheckInIntervalAttr zcl.AttrID = 0

func (CheckInInterval) ID() zcl.AttrID   { return CheckInIntervalAttr }
func (CheckInInterval) Readable() bool   { return true }
func (CheckInInterval) Writable() bool   { return true }
func (CheckInInterval) Reportable() bool { return false }
func (CheckInInterval) SceneIndex() int  { return -1 }

func (CheckInInterval) Name() string                  { return "Check-in Interval" }
func (a *CheckInInterval) TypeID() zcl.TypeID         { return zcl.Zu32(*a).ID() }
func (a *CheckInInterval) Value() zcl.Val             { return a }
func (a CheckInInterval) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *CheckInInterval) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = CheckInInterval(*nt)
	return br, err
}

func (a *CheckInInterval) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu32); ok {
		*a = CheckInInterval(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CheckInInterval) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

type CheckInIntervalMin zcl.Zu32

const CheckInIntervalMinAttr zcl.AttrID = 4

func (CheckInIntervalMin) ID() zcl.AttrID   { return CheckInIntervalMinAttr }
func (CheckInIntervalMin) Readable() bool   { return true }
func (CheckInIntervalMin) Writable() bool   { return false }
func (CheckInIntervalMin) Reportable() bool { return false }
func (CheckInIntervalMin) SceneIndex() int  { return -1 }

func (CheckInIntervalMin) Name() string                  { return "Check-in Interval Min" }
func (a *CheckInIntervalMin) TypeID() zcl.TypeID         { return zcl.Zu32(*a).ID() }
func (a *CheckInIntervalMin) Value() zcl.Val             { return a }
func (a CheckInIntervalMin) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *CheckInIntervalMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = CheckInIntervalMin(*nt)
	return br, err
}

func (a *CheckInIntervalMin) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu32); ok {
		*a = CheckInIntervalMin(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CheckInIntervalMin) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

type ChildMoved zcl.Zu16

const ChildMovedAttr zcl.AttrID = 273

func (ChildMoved) ID() zcl.AttrID   { return ChildMovedAttr }
func (ChildMoved) Readable() bool   { return true }
func (ChildMoved) Writable() bool   { return false }
func (ChildMoved) Reportable() bool { return false }
func (ChildMoved) SceneIndex() int  { return -1 }

func (ChildMoved) Name() string                  { return "Child Moved" }
func (a *ChildMoved) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ChildMoved) Value() zcl.Val             { return a }
func (a ChildMoved) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ChildMoved) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ChildMoved(*nt)
	return br, err
}

func (a *ChildMoved) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ChildMoved(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ChildMoved) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type ClusterId zcl.Zu16

func (ClusterId) Name() string                  { return "Cluster Id" }
func (a *ClusterId) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ClusterId) Value() zcl.Val             { return a }
func (a ClusterId) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ClusterId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ClusterId(*nt)
	return br, err
}

func (a *ClusterId) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ClusterId(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ClusterId) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type ClusterRevision zcl.Zu16

const ClusterRevisionAttr zcl.AttrID = 65533

func (ClusterRevision) ID() zcl.AttrID   { return ClusterRevisionAttr }
func (ClusterRevision) Readable() bool   { return true }
func (ClusterRevision) Writable() bool   { return true }
func (ClusterRevision) Reportable() bool { return false }
func (ClusterRevision) SceneIndex() int  { return -1 }

func (ClusterRevision) Name() string                  { return "Cluster Revision" }
func (a *ClusterRevision) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ClusterRevision) Value() zcl.Val             { return a }
func (a ClusterRevision) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ClusterRevision) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ClusterRevision(*nt)
	return br, err
}

func (a *ClusterRevision) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ClusterRevision(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ClusterRevision) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type Configuration zcl.Zbmp16

const ConfigurationAttr zcl.AttrID = 49

func (Configuration) ID() zcl.AttrID   { return ConfigurationAttr }
func (Configuration) Readable() bool   { return true }
func (Configuration) Writable() bool   { return true }
func (Configuration) Reportable() bool { return false }
func (Configuration) SceneIndex() int  { return -1 }

func (Configuration) Name() string                  { return "Configuration" }
func (a *Configuration) TypeID() zcl.TypeID         { return zcl.Zbmp16(*a).ID() }
func (a *Configuration) Value() zcl.Val             { return a }
func (a Configuration) MarshalZcl() ([]byte, error) { return zcl.Zbmp16(a).MarshalZcl() }

func (a *Configuration) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*a = Configuration(*nt)
	return br, err
}

func (a *Configuration) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp16); ok {
		*a = Configuration(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Configuration) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Touchlink enabled 0")
		case 1:
			bstr = append(bstr, "Touchlink enabled 1")
		case 3:
			bstr = append(bstr, "Touchlink enabled 2")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a Configuration) IsTouchlinkEnabled0() bool { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a Configuration) IsTouchlinkEnabled1() bool { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a Configuration) IsTouchlinkEnabled2() bool { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a *Configuration) SetTouchlinkEnabled0(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *Configuration) SetTouchlinkEnabled1(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}
func (a *Configuration) SetTouchlinkEnabled2(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b))
}

func (Configuration) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Touchlink enabled 0"},
		{Value: 1, Name: "Touchlink enabled 1"},
		{Value: 3, Name: "Touchlink enabled 2"},
	}
}

type CurrentGroup zcl.Zu16

const CurrentGroupAttr zcl.AttrID = 2

func (CurrentGroup) ID() zcl.AttrID   { return CurrentGroupAttr }
func (CurrentGroup) Readable() bool   { return true }
func (CurrentGroup) Writable() bool   { return false }
func (CurrentGroup) Reportable() bool { return false }
func (CurrentGroup) SceneIndex() int  { return -1 }

func (CurrentGroup) Name() string                  { return "Current Group" }
func (a *CurrentGroup) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *CurrentGroup) Value() zcl.Val             { return a }
func (a CurrentGroup) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *CurrentGroup) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentGroup(*nt)
	return br, err
}

func (a *CurrentGroup) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = CurrentGroup(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CurrentGroup) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// CurrentLevel The CurrentLevel attribute represents the current level of this device. meaning of 'level' is device dependent.
type CurrentLevel zcl.Zu8

const CurrentLevelAttr zcl.AttrID = 0

func (CurrentLevel) ID() zcl.AttrID   { return CurrentLevelAttr }
func (CurrentLevel) Readable() bool   { return true }
func (CurrentLevel) Writable() bool   { return false }
func (CurrentLevel) Reportable() bool { return true }
func (CurrentLevel) SceneIndex() int  { return 1 }

func (CurrentLevel) Name() string                  { return "Current Level" }
func (a *CurrentLevel) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *CurrentLevel) Value() zcl.Val             { return a }
func (a CurrentLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *CurrentLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentLevel(*nt)
	return br, err
}

func (a *CurrentLevel) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = CurrentLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CurrentLevel) String() string {
	return zcl.Percent.Format(float64(a) / 2.54)
}

type CurrentScene zcl.Zu8

const CurrentSceneAttr zcl.AttrID = 1

func (CurrentScene) ID() zcl.AttrID   { return CurrentSceneAttr }
func (CurrentScene) Readable() bool   { return true }
func (CurrentScene) Writable() bool   { return false }
func (CurrentScene) Reportable() bool { return false }
func (CurrentScene) SceneIndex() int  { return -1 }

func (CurrentScene) Name() string                  { return "Current Scene" }
func (a *CurrentScene) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *CurrentScene) Value() zcl.Val             { return a }
func (a CurrentScene) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *CurrentScene) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentScene(*nt)
	return br, err
}

func (a *CurrentScene) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = CurrentScene(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CurrentScene) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type CurrentTemperature zcl.Zs16

const CurrentTemperatureAttr zcl.AttrID = 0

func (CurrentTemperature) ID() zcl.AttrID   { return CurrentTemperatureAttr }
func (CurrentTemperature) Readable() bool   { return true }
func (CurrentTemperature) Writable() bool   { return false }
func (CurrentTemperature) Reportable() bool { return false }
func (CurrentTemperature) SceneIndex() int  { return -1 }

func (CurrentTemperature) Name() string                  { return "Current Temperature" }
func (a *CurrentTemperature) TypeID() zcl.TypeID         { return zcl.Zs16(*a).ID() }
func (a *CurrentTemperature) Value() zcl.Val             { return a }
func (a CurrentTemperature) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *CurrentTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentTemperature(*nt)
	return br, err
}

func (a *CurrentTemperature) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = CurrentTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CurrentTemperature) String() string {
	return zcl.DegreesCelsius.Format(float64(a))
}

type DateCode zcl.Zcstring

const DateCodeAttr zcl.AttrID = 6

func (DateCode) ID() zcl.AttrID   { return DateCodeAttr }
func (DateCode) Readable() bool   { return true }
func (DateCode) Writable() bool   { return false }
func (DateCode) Reportable() bool { return false }
func (DateCode) SceneIndex() int  { return -1 }

func (DateCode) Name() string                  { return "Date Code" }
func (a *DateCode) TypeID() zcl.TypeID         { return zcl.Zcstring(*a).ID() }
func (a *DateCode) Value() zcl.Val             { return a }
func (a DateCode) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *DateCode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = DateCode(*nt)
	return br, err
}

func (a *DateCode) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = DateCode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DateCode) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

type DefaultMoveRate zcl.Zu8

const DefaultMoveRateAttr zcl.AttrID = 20

func (DefaultMoveRate) ID() zcl.AttrID   { return DefaultMoveRateAttr }
func (DefaultMoveRate) Readable() bool   { return true }
func (DefaultMoveRate) Writable() bool   { return true }
func (DefaultMoveRate) Reportable() bool { return false }
func (DefaultMoveRate) SceneIndex() int  { return -1 }

func (DefaultMoveRate) Name() string                  { return "Default Move Rate" }
func (a *DefaultMoveRate) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *DefaultMoveRate) Value() zcl.Val             { return a }
func (a DefaultMoveRate) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *DefaultMoveRate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = DefaultMoveRate(*nt)
	return br, err
}

func (a *DefaultMoveRate) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = DefaultMoveRate(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DefaultMoveRate) String() string {
	return zcl.PercentPerSecond.Format(float64(a) / 2.54)
}

type Device zcl.Zuid

func (Device) Name() string                  { return "Device" }
func (a *Device) TypeID() zcl.TypeID         { return zcl.Zuid(*a).ID() }
func (a *Device) Value() zcl.Val             { return a }
func (a Device) MarshalZcl() ([]byte, error) { return zcl.Zuid(a).MarshalZcl() }

func (a *Device) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*a = Device(*nt)
	return br, err
}

func (a *Device) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zuid); ok {
		*a = Device(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Device) String() string {
	return zcl.Sprintf("%v", zcl.Zuid(a))
}

type DeviceEnabled zcl.Zbool

const DeviceEnabledAttr zcl.AttrID = 18

func (DeviceEnabled) ID() zcl.AttrID   { return DeviceEnabledAttr }
func (DeviceEnabled) Readable() bool   { return true }
func (DeviceEnabled) Writable() bool   { return true }
func (DeviceEnabled) Reportable() bool { return false }
func (DeviceEnabled) SceneIndex() int  { return -1 }

func (DeviceEnabled) Name() string                  { return "Device Enabled" }
func (a *DeviceEnabled) TypeID() zcl.TypeID         { return zcl.Zbool(*a).ID() }
func (a *DeviceEnabled) Value() zcl.Val             { return a }
func (a DeviceEnabled) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *DeviceEnabled) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = DeviceEnabled(*nt)
	return br, err
}

func (a *DeviceEnabled) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbool); ok {
		*a = DeviceEnabled(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DeviceEnabled) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

type DeviceTempAlarmMask zcl.Zbmp8

const DeviceTempAlarmMaskAttr zcl.AttrID = 16

func (DeviceTempAlarmMask) ID() zcl.AttrID   { return DeviceTempAlarmMaskAttr }
func (DeviceTempAlarmMask) Readable() bool   { return true }
func (DeviceTempAlarmMask) Writable() bool   { return true }
func (DeviceTempAlarmMask) Reportable() bool { return false }
func (DeviceTempAlarmMask) SceneIndex() int  { return -1 }

func (DeviceTempAlarmMask) Name() string                  { return "Device Temp Alarm Mask" }
func (a *DeviceTempAlarmMask) TypeID() zcl.TypeID         { return zcl.Zbmp8(*a).ID() }
func (a *DeviceTempAlarmMask) Value() zcl.Val             { return a }
func (a DeviceTempAlarmMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *DeviceTempAlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = DeviceTempAlarmMask(*nt)
	return br, err
}

func (a *DeviceTempAlarmMask) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = DeviceTempAlarmMask(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DeviceTempAlarmMask) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Temperature too low")
		case 1:
			bstr = append(bstr, "Temperature too high")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a DeviceTempAlarmMask) IsTemperatureTooLow() bool  { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a DeviceTempAlarmMask) IsTemperatureTooHigh() bool { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a *DeviceTempAlarmMask) SetTemperatureTooLow(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *DeviceTempAlarmMask) SetTemperatureTooHigh(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}

func (DeviceTempAlarmMask) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Temperature too low"},
		{Value: 1, Name: "Temperature too high"},
	}
}

type DisableLocalConfig zcl.Zbmp8

const DisableLocalConfigAttr zcl.AttrID = 20

func (DisableLocalConfig) ID() zcl.AttrID   { return DisableLocalConfigAttr }
func (DisableLocalConfig) Readable() bool   { return true }
func (DisableLocalConfig) Writable() bool   { return true }
func (DisableLocalConfig) Reportable() bool { return false }
func (DisableLocalConfig) SceneIndex() int  { return -1 }

func (DisableLocalConfig) Name() string                  { return "Disable Local Config" }
func (a *DisableLocalConfig) TypeID() zcl.TypeID         { return zcl.Zbmp8(*a).ID() }
func (a *DisableLocalConfig) Value() zcl.Val             { return a }
func (a DisableLocalConfig) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *DisableLocalConfig) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = DisableLocalConfig(*nt)
	return br, err
}

func (a *DisableLocalConfig) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = DisableLocalConfig(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DisableLocalConfig) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Disable Factory Reset")
		case 1:
			bstr = append(bstr, "Disable Device Configuration")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a DisableLocalConfig) IsDisableFactoryReset() bool { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a DisableLocalConfig) IsDisableDeviceConfiguration() bool {
	return zcl.BitmapTest([]byte(a[:]), 1)
}
func (a *DisableLocalConfig) SetDisableFactoryReset(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *DisableLocalConfig) SetDisableDeviceConfiguration(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}

func (DisableLocalConfig) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Disable Factory Reset"},
		{Value: 1, Name: "Disable Device Configuration"},
	}
}

type Distance zcl.Zu16

func (Distance) Name() string                  { return "Distance" }
func (a *Distance) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *Distance) Value() zcl.Val             { return a }
func (a Distance) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *Distance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Distance(*nt)
	return br, err
}

func (a *Distance) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = Distance(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Distance) String() string {
	return zcl.Meters.Format(float64(a) / 10)
}

type DstEnd zcl.Zutc

const DstEndAttr zcl.AttrID = 4

func (DstEnd) ID() zcl.AttrID   { return DstEndAttr }
func (DstEnd) Readable() bool   { return true }
func (DstEnd) Writable() bool   { return true }
func (DstEnd) Reportable() bool { return false }
func (DstEnd) SceneIndex() int  { return -1 }

func (DstEnd) Name() string                  { return "Dst End" }
func (a *DstEnd) TypeID() zcl.TypeID         { return zcl.Zutc(*a).ID() }
func (a *DstEnd) Value() zcl.Val             { return a }
func (a DstEnd) MarshalZcl() ([]byte, error) { return zcl.Zutc(a).MarshalZcl() }

func (a *DstEnd) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*a = DstEnd(*nt)
	return br, err
}

func (a *DstEnd) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zutc); ok {
		*a = DstEnd(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DstEnd) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(a))
}

type DstShift zcl.Zs32

const DstShiftAttr zcl.AttrID = 5

func (DstShift) ID() zcl.AttrID   { return DstShiftAttr }
func (DstShift) Readable() bool   { return true }
func (DstShift) Writable() bool   { return true }
func (DstShift) Reportable() bool { return false }
func (DstShift) SceneIndex() int  { return -1 }

func (DstShift) Name() string                  { return "Dst Shift" }
func (a *DstShift) TypeID() zcl.TypeID         { return zcl.Zs32(*a).ID() }
func (a *DstShift) Value() zcl.Val             { return a }
func (a DstShift) MarshalZcl() ([]byte, error) { return zcl.Zs32(a).MarshalZcl() }

func (a *DstShift) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs32)
	br, err := nt.UnmarshalZcl(b)
	*a = DstShift(*nt)
	return br, err
}

func (a *DstShift) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs32); ok {
		*a = DstShift(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DstShift) String() string {
	return zcl.Seconds.Format(float64(a))
}

type DstStart zcl.Zutc

const DstStartAttr zcl.AttrID = 3

func (DstStart) ID() zcl.AttrID   { return DstStartAttr }
func (DstStart) Readable() bool   { return true }
func (DstStart) Writable() bool   { return true }
func (DstStart) Reportable() bool { return false }
func (DstStart) SceneIndex() int  { return -1 }

func (DstStart) Name() string                  { return "Dst Start" }
func (a *DstStart) TypeID() zcl.TypeID         { return zcl.Zutc(*a).ID() }
func (a *DstStart) Value() zcl.Val             { return a }
func (a DstStart) MarshalZcl() ([]byte, error) { return zcl.Zutc(a).MarshalZcl() }

func (a *DstStart) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*a = DstStart(*nt)
	return br, err
}

func (a *DstStart) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zutc); ok {
		*a = DstStart(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DstStart) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(a))
}

// EffectIdentifier when turning lights off
type EffectIdentifier zcl.Zenum8

func (EffectIdentifier) Name() string                  { return "Effect Identifier" }
func (a *EffectIdentifier) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *EffectIdentifier) Value() zcl.Val             { return a }
func (a EffectIdentifier) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *EffectIdentifier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = EffectIdentifier(*nt)
	return br, err
}

func (a *EffectIdentifier) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = EffectIdentifier(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a EffectIdentifier) String() string {
	switch a {
	case 0x00:
		return "Delayed all off"
	case 0x01:
		return "Dying light"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a EffectIdentifier) IsDelayedAllOff() bool { return a == 0x00 }
func (a EffectIdentifier) IsDyingLight() bool    { return a == 0x01 }
func (a *EffectIdentifier) SetDelayedAllOff()    { *a = 0x00 }
func (a *EffectIdentifier) SetDyingLight()       { *a = 0x01 }

func (EffectIdentifier) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Delayed all off"},
		{Value: 0x01, Name: "Dying light"},
	}
}

type EffectVariant zcl.Zenum8

func (EffectVariant) Name() string                  { return "Effect Variant" }
func (a *EffectVariant) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *EffectVariant) Value() zcl.Val             { return a }
func (a EffectVariant) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *EffectVariant) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = EffectVariant(*nt)
	return br, err
}

func (a *EffectVariant) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = EffectVariant(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a EffectVariant) String() string {
	switch a {
	case 0x00:
		return "Fade to off in 0.8 seconds / 20% dim up in 0.5 then fade to off in 1 second"
	case 0x01:
		return "No Fade"
	case 0x02:
		return "50% dim down in 0.8s then fade off in 12s"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a EffectVariant) IsFadeToOffIn08Seconds20DimUpIn05ThenFadeToOffIn1Second() bool {
	return a == 0x00
}
func (a EffectVariant) IsNoFade() bool                                             { return a == 0x01 }
func (a EffectVariant) Is50DimDownIn08SThenFadeOffIn12S() bool                     { return a == 0x02 }
func (a *EffectVariant) SetFadeToOffIn08Seconds20DimUpIn05ThenFadeToOffIn1Second() { *a = 0x00 }
func (a *EffectVariant) SetNoFade()                                                { *a = 0x01 }
func (a *EffectVariant) Set50DimDownIn08SThenFadeOffIn12S()                        { *a = 0x02 }

func (EffectVariant) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Fade to off in 0.8 seconds / 20% dim up in 0.5 then fade to off in 1 second"},
		{Value: 0x01, Name: "No Fade"},
		{Value: 0x02, Name: "50% dim down in 0.8s then fade off in 12s"},
	}
}

type FastPollTimeout zcl.Zu16

const FastPollTimeoutAttr zcl.AttrID = 3

func (FastPollTimeout) ID() zcl.AttrID   { return FastPollTimeoutAttr }
func (FastPollTimeout) Readable() bool   { return true }
func (FastPollTimeout) Writable() bool   { return true }
func (FastPollTimeout) Reportable() bool { return false }
func (FastPollTimeout) SceneIndex() int  { return -1 }

func (FastPollTimeout) Name() string                  { return "Fast Poll Timeout" }
func (a *FastPollTimeout) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *FastPollTimeout) Value() zcl.Val             { return a }
func (a FastPollTimeout) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *FastPollTimeout) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = FastPollTimeout(*nt)
	return br, err
}

func (a *FastPollTimeout) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = FastPollTimeout(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a FastPollTimeout) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type FastPollTimeoutMax zcl.Zu16

const FastPollTimeoutMaxAttr zcl.AttrID = 6

func (FastPollTimeoutMax) ID() zcl.AttrID   { return FastPollTimeoutMaxAttr }
func (FastPollTimeoutMax) Readable() bool   { return true }
func (FastPollTimeoutMax) Writable() bool   { return false }
func (FastPollTimeoutMax) Reportable() bool { return false }
func (FastPollTimeoutMax) SceneIndex() int  { return -1 }

func (FastPollTimeoutMax) Name() string                  { return "Fast Poll Timeout Max" }
func (a *FastPollTimeoutMax) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *FastPollTimeoutMax) Value() zcl.Val             { return a }
func (a FastPollTimeoutMax) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *FastPollTimeoutMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = FastPollTimeoutMax(*nt)
	return br, err
}

func (a *FastPollTimeoutMax) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = FastPollTimeoutMax(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a FastPollTimeoutMax) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// GenericDeviceClass defines the field of application of the GenericDeviceType attribute
type GenericDeviceClass zcl.Zenum8

const GenericDeviceClassAttr zcl.AttrID = 8

func (GenericDeviceClass) ID() zcl.AttrID   { return GenericDeviceClassAttr }
func (GenericDeviceClass) Readable() bool   { return true }
func (GenericDeviceClass) Writable() bool   { return false }
func (GenericDeviceClass) Reportable() bool { return false }
func (GenericDeviceClass) SceneIndex() int  { return -1 }

func (GenericDeviceClass) Name() string                  { return "Generic Device Class" }
func (a *GenericDeviceClass) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *GenericDeviceClass) Value() zcl.Val             { return a }
func (a GenericDeviceClass) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *GenericDeviceClass) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = GenericDeviceClass(*nt)
	return br, err
}

func (a *GenericDeviceClass) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = GenericDeviceClass(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a GenericDeviceClass) String() string {
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

// GenericDeviceType allows an application to show an icon on a rich user interface (e.g. smartphone app)
type GenericDeviceType zcl.Zenum8

const GenericDeviceTypeAttr zcl.AttrID = 9

func (GenericDeviceType) ID() zcl.AttrID   { return GenericDeviceTypeAttr }
func (GenericDeviceType) Readable() bool   { return true }
func (GenericDeviceType) Writable() bool   { return false }
func (GenericDeviceType) Reportable() bool { return false }
func (GenericDeviceType) SceneIndex() int  { return -1 }

func (GenericDeviceType) Name() string                  { return "Generic Device Type" }
func (a *GenericDeviceType) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *GenericDeviceType) Value() zcl.Val             { return a }
func (a GenericDeviceType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *GenericDeviceType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = GenericDeviceType(*nt)
	return br, err
}

func (a *GenericDeviceType) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = GenericDeviceType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a GenericDeviceType) String() string {
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

type GlobalSceneControl zcl.Zbool

const GlobalSceneControlAttr zcl.AttrID = 16384

func (GlobalSceneControl) ID() zcl.AttrID   { return GlobalSceneControlAttr }
func (GlobalSceneControl) Readable() bool   { return true }
func (GlobalSceneControl) Writable() bool   { return false }
func (GlobalSceneControl) Reportable() bool { return false }
func (GlobalSceneControl) SceneIndex() int  { return -1 }

func (GlobalSceneControl) Name() string                  { return "Global Scene Control" }
func (a *GlobalSceneControl) TypeID() zcl.TypeID         { return zcl.Zbool(*a).ID() }
func (a *GlobalSceneControl) Value() zcl.Val             { return a }
func (a GlobalSceneControl) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *GlobalSceneControl) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = GlobalSceneControl(*nt)
	return br, err
}

func (a *GlobalSceneControl) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbool); ok {
		*a = GlobalSceneControl(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a GlobalSceneControl) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

type GroupId zcl.Zu16

func (GroupId) Name() string                  { return "Group ID" }
func (a *GroupId) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *GroupId) Value() zcl.Val             { return a }
func (a GroupId) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *GroupId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = GroupId(*nt)
	return br, err
}

func (a *GroupId) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = GroupId(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a GroupId) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type GroupNameSupport zcl.Zbmp8

const GroupNameSupportAttr zcl.AttrID = 0

func (GroupNameSupport) ID() zcl.AttrID   { return GroupNameSupportAttr }
func (GroupNameSupport) Readable() bool   { return true }
func (GroupNameSupport) Writable() bool   { return false }
func (GroupNameSupport) Reportable() bool { return false }
func (GroupNameSupport) SceneIndex() int  { return -1 }

func (GroupNameSupport) Name() string                  { return "Group Name Support" }
func (a *GroupNameSupport) TypeID() zcl.TypeID         { return zcl.Zbmp8(*a).ID() }
func (a *GroupNameSupport) Value() zcl.Val             { return a }
func (a GroupNameSupport) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *GroupNameSupport) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = GroupNameSupport(*nt)
	return br, err
}

func (a *GroupNameSupport) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = GroupNameSupport(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a GroupNameSupport) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 7:
			bstr = append(bstr, "Names Supported")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a GroupNameSupport) IsNamesSupported() bool { return zcl.BitmapTest([]byte(a[:]), 7) }
func (a *GroupNameSupport) SetNamesSupported(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 7, b))
}

func (GroupNameSupport) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 7, Name: "Names Supported"},
	}
}

// GroupCapacity specifies remaining number of groups that can be added.
// If set to 0xFE, at least one more group can be added (exact number unknown)
// If set to 0xFF, it's unknown if any more groups can be added
type GroupCapacity zcl.Zu8

func (GroupCapacity) Name() string                  { return "Group capacity" }
func (a *GroupCapacity) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *GroupCapacity) Value() zcl.Val             { return a }
func (a GroupCapacity) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *GroupCapacity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = GroupCapacity(*nt)
	return br, err
}

func (a *GroupCapacity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = GroupCapacity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a GroupCapacity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type GroupList zcl.Zset

func (GroupList) Name() string          { return "Group list" }
func (a *GroupList) TypeID() zcl.TypeID { return zcl.Zset(*a).ID() }
func (a *GroupList) Value() zcl.Val     { return a }
func (a GroupList) MarshalZcl() ([]byte, error) {
	return nil, zcl.Errorf("not implemented")
}

func (a *GroupList) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zset)
	br, err := nt.UnmarshalZcl(b)
	*a = GroupList(*nt)
	return br, err
}

func (a *GroupList) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zset); ok {
		*a = GroupList(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a GroupList) String() string {
	return zcl.Sprintf("%v", zcl.Zset(a))
}

type GroupName zcl.Zcstring

func (GroupName) Name() string                  { return "Group name" }
func (a *GroupName) TypeID() zcl.TypeID         { return zcl.Zcstring(*a).ID() }
func (a *GroupName) Value() zcl.Val             { return a }
func (a GroupName) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *GroupName) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = GroupName(*nt)
	return br, err
}

func (a *GroupName) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = GroupName(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a GroupName) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

type HwVersion zcl.Zu8

const HwVersionAttr zcl.AttrID = 3

func (HwVersion) ID() zcl.AttrID   { return HwVersionAttr }
func (HwVersion) Readable() bool   { return true }
func (HwVersion) Writable() bool   { return false }
func (HwVersion) Reportable() bool { return false }
func (HwVersion) SceneIndex() int  { return -1 }

func (HwVersion) Name() string                  { return "HW Version" }
func (a *HwVersion) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *HwVersion) Value() zcl.Val             { return a }
func (a HwVersion) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *HwVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = HwVersion(*nt)
	return br, err
}

func (a *HwVersion) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = HwVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a HwVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type HighTempDwellTripPoint zcl.Zu24

const HighTempDwellTripPointAttr zcl.AttrID = 20

func (HighTempDwellTripPoint) ID() zcl.AttrID   { return HighTempDwellTripPointAttr }
func (HighTempDwellTripPoint) Readable() bool   { return false }
func (HighTempDwellTripPoint) Writable() bool   { return false }
func (HighTempDwellTripPoint) Reportable() bool { return false }
func (HighTempDwellTripPoint) SceneIndex() int  { return -1 }

func (HighTempDwellTripPoint) Name() string                  { return "High Temp Dwell Trip Point" }
func (a *HighTempDwellTripPoint) TypeID() zcl.TypeID         { return zcl.Zu24(*a).ID() }
func (a *HighTempDwellTripPoint) Value() zcl.Val             { return a }
func (a HighTempDwellTripPoint) MarshalZcl() ([]byte, error) { return zcl.Zu24(a).MarshalZcl() }

func (a *HighTempDwellTripPoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*a = HighTempDwellTripPoint(*nt)
	return br, err
}

func (a *HighTempDwellTripPoint) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu24); ok {
		*a = HighTempDwellTripPoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a HighTempDwellTripPoint) String() string {
	return zcl.Seconds.Format(float64(a))
}

// HighTempThreshold If the current temperature goes above the threshold for longer
// than the time specified by high temp dwell, an alarm will be triggered
type HighTempThreshold zcl.Zs16

const HighTempThresholdAttr zcl.AttrID = 18

func (HighTempThreshold) ID() zcl.AttrID   { return HighTempThresholdAttr }
func (HighTempThreshold) Readable() bool   { return true }
func (HighTempThreshold) Writable() bool   { return true }
func (HighTempThreshold) Reportable() bool { return false }
func (HighTempThreshold) SceneIndex() int  { return -1 }

func (HighTempThreshold) Name() string                  { return "High Temp Threshold" }
func (a *HighTempThreshold) TypeID() zcl.TypeID         { return zcl.Zs16(*a).ID() }
func (a *HighTempThreshold) Value() zcl.Val             { return a }
func (a HighTempThreshold) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *HighTempThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = HighTempThreshold(*nt)
	return br, err
}

func (a *HighTempThreshold) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = HighTempThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a HighTempThreshold) String() string {
	return zcl.DegreesCelsius.Format(float64(a))
}

type IOApplicationType zcl.Zu32

const IOApplicationTypeAttr zcl.AttrID = 256

func (IOApplicationType) ID() zcl.AttrID   { return IOApplicationTypeAttr }
func (IOApplicationType) Readable() bool   { return true }
func (IOApplicationType) Writable() bool   { return false }
func (IOApplicationType) Reportable() bool { return false }
func (IOApplicationType) SceneIndex() int  { return -1 }

func (IOApplicationType) Name() string                  { return "I/O Application Type" }
func (a *IOApplicationType) TypeID() zcl.TypeID         { return zcl.Zu32(*a).ID() }
func (a *IOApplicationType) Value() zcl.Val             { return a }
func (a IOApplicationType) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *IOApplicationType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = IOApplicationType(*nt)
	return br, err
}

func (a *IOApplicationType) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu32); ok {
		*a = IOApplicationType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a IOApplicationType) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

type IODescription zcl.Zcstring

const IODescriptionAttr zcl.AttrID = 28

func (IODescription) ID() zcl.AttrID   { return IODescriptionAttr }
func (IODescription) Readable() bool   { return true }
func (IODescription) Writable() bool   { return true }
func (IODescription) Reportable() bool { return false }
func (IODescription) SceneIndex() int  { return -1 }

func (IODescription) Name() string                  { return "I/O Description" }
func (a *IODescription) TypeID() zcl.TypeID         { return zcl.Zcstring(*a).ID() }
func (a *IODescription) Value() zcl.Val             { return a }
func (a IODescription) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *IODescription) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = IODescription(*nt)
	return br, err
}

func (a *IODescription) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = IODescription(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a IODescription) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

type IOOutOfService zcl.Zbool

const IOOutOfServiceAttr zcl.AttrID = 81

func (IOOutOfService) ID() zcl.AttrID   { return IOOutOfServiceAttr }
func (IOOutOfService) Readable() bool   { return true }
func (IOOutOfService) Writable() bool   { return true }
func (IOOutOfService) Reportable() bool { return false }
func (IOOutOfService) SceneIndex() int  { return -1 }

func (IOOutOfService) Name() string                  { return "I/O Out of service" }
func (a *IOOutOfService) TypeID() zcl.TypeID         { return zcl.Zbool(*a).ID() }
func (a *IOOutOfService) Value() zcl.Val             { return a }
func (a IOOutOfService) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *IOOutOfService) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = IOOutOfService(*nt)
	return br, err
}

func (a *IOOutOfService) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbool); ok {
		*a = IOOutOfService(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a IOOutOfService) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

type IOReliability zcl.Zenum8

const IOReliabilityAttr zcl.AttrID = 103

func (IOReliability) ID() zcl.AttrID   { return IOReliabilityAttr }
func (IOReliability) Readable() bool   { return true }
func (IOReliability) Writable() bool   { return true }
func (IOReliability) Reportable() bool { return false }
func (IOReliability) SceneIndex() int  { return -1 }

func (IOReliability) Name() string                  { return "I/O Reliability" }
func (a *IOReliability) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *IOReliability) Value() zcl.Val             { return a }
func (a IOReliability) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *IOReliability) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = IOReliability(*nt)
	return br, err
}

func (a *IOReliability) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = IOReliability(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a IOReliability) String() string {
	switch a {
	case 0x00:
		return "No fault detected"
	case 0x01:
		return "No Sensor"
	case 0x02:
		return "Over Range"
	case 0x03:
		return "Under Range"
	case 0x04:
		return "Open Loop"
	case 0x05:
		return "Shorted Loop"
	case 0x06:
		return "No Output"
	case 0x07:
		return "Unreliable (other)"
	case 0x08:
		return "Process Error"
	case 0x09:
		return "Multi state fault"
	case 0x0A:
		return "Configuration Error"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a IOReliability) IsNoFaultDetected() bool    { return a == 0x00 }
func (a IOReliability) IsNoSensor() bool           { return a == 0x01 }
func (a IOReliability) IsOverRange() bool          { return a == 0x02 }
func (a IOReliability) IsUnderRange() bool         { return a == 0x03 }
func (a IOReliability) IsOpenLoop() bool           { return a == 0x04 }
func (a IOReliability) IsShortedLoop() bool        { return a == 0x05 }
func (a IOReliability) IsNoOutput() bool           { return a == 0x06 }
func (a IOReliability) IsUnreliableOther() bool    { return a == 0x07 }
func (a IOReliability) IsProcessError() bool       { return a == 0x08 }
func (a IOReliability) IsMultiStateFault() bool    { return a == 0x09 }
func (a IOReliability) IsConfigurationError() bool { return a == 0x0A }
func (a *IOReliability) SetNoFaultDetected()       { *a = 0x00 }
func (a *IOReliability) SetNoSensor()              { *a = 0x01 }
func (a *IOReliability) SetOverRange()             { *a = 0x02 }
func (a *IOReliability) SetUnderRange()            { *a = 0x03 }
func (a *IOReliability) SetOpenLoop()              { *a = 0x04 }
func (a *IOReliability) SetShortedLoop()           { *a = 0x05 }
func (a *IOReliability) SetNoOutput()              { *a = 0x06 }
func (a *IOReliability) SetUnreliableOther()       { *a = 0x07 }
func (a *IOReliability) SetProcessError()          { *a = 0x08 }
func (a *IOReliability) SetMultiStateFault()       { *a = 0x09 }
func (a *IOReliability) SetConfigurationError()    { *a = 0x0A }

func (IOReliability) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "No fault detected"},
		{Value: 0x01, Name: "No Sensor"},
		{Value: 0x02, Name: "Over Range"},
		{Value: 0x03, Name: "Under Range"},
		{Value: 0x04, Name: "Open Loop"},
		{Value: 0x05, Name: "Shorted Loop"},
		{Value: 0x06, Name: "No Output"},
		{Value: 0x07, Name: "Unreliable (other)"},
		{Value: 0x08, Name: "Process Error"},
		{Value: 0x09, Name: "Multi state fault"},
		{Value: 0x0A, Name: "Configuration Error"},
	}
}

type IOStatusFlags zcl.Zbmp8

const IOStatusFlagsAttr zcl.AttrID = 111

func (IOStatusFlags) ID() zcl.AttrID   { return IOStatusFlagsAttr }
func (IOStatusFlags) Readable() bool   { return true }
func (IOStatusFlags) Writable() bool   { return false }
func (IOStatusFlags) Reportable() bool { return true }
func (IOStatusFlags) SceneIndex() int  { return -1 }

func (IOStatusFlags) Name() string                  { return "I/O Status flags" }
func (a *IOStatusFlags) TypeID() zcl.TypeID         { return zcl.Zbmp8(*a).ID() }
func (a *IOStatusFlags) Value() zcl.Val             { return a }
func (a IOStatusFlags) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *IOStatusFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = IOStatusFlags(*nt)
	return br, err
}

func (a *IOStatusFlags) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = IOStatusFlags(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a IOStatusFlags) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "In Alarm")
		case 1:
			bstr = append(bstr, "Fault")
		case 2:
			bstr = append(bstr, "Overidden")
		case 3:
			bstr = append(bstr, "Out of Service")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a IOStatusFlags) IsInAlarm() bool         { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a IOStatusFlags) IsFault() bool           { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a IOStatusFlags) IsOveridden() bool       { return zcl.BitmapTest([]byte(a[:]), 2) }
func (a IOStatusFlags) IsOutOfService() bool    { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a *IOStatusFlags) SetInAlarm(b bool)      { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b)) }
func (a *IOStatusFlags) SetFault(b bool)        { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b)) }
func (a *IOStatusFlags) SetOveridden(b bool)    { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b)) }
func (a *IOStatusFlags) SetOutOfService(b bool) { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b)) }

func (IOStatusFlags) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "In Alarm"},
		{Value: 1, Name: "Fault"},
		{Value: 2, Name: "Overidden"},
		{Value: 3, Name: "Out of Service"},
	}
}

type IOUnitType zcl.EngineeringUnit

const IOUnitTypeAttr zcl.AttrID = 117

func (IOUnitType) ID() zcl.AttrID   { return IOUnitTypeAttr }
func (IOUnitType) Readable() bool   { return true }
func (IOUnitType) Writable() bool   { return true }
func (IOUnitType) Reportable() bool { return false }
func (IOUnitType) SceneIndex() int  { return -1 }

func (IOUnitType) Name() string                  { return "I/O Unit Type" }
func (a *IOUnitType) TypeID() zcl.TypeID         { return zcl.EngineeringUnit(*a).ID() }
func (a *IOUnitType) Value() zcl.Val             { return a }
func (a IOUnitType) MarshalZcl() ([]byte, error) { return zcl.EngineeringUnit(a).MarshalZcl() }

func (a *IOUnitType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.EngineeringUnit)
	br, err := nt.UnmarshalZcl(b)
	*a = IOUnitType(*nt)
	return br, err
}

func (a *IOUnitType) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.EngineeringUnit); ok {
		*a = IOUnitType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a IOUnitType) String() string {
	return zcl.Sprintf("%v", zcl.EngineeringUnit(a))
}

// IdentifyEffect The effect identifier field specifies the identify effect to use.
type IdentifyEffect zcl.Zenum8

func (IdentifyEffect) Name() string                  { return "Identify Effect" }
func (a *IdentifyEffect) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *IdentifyEffect) Value() zcl.Val             { return a }
func (a IdentifyEffect) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *IdentifyEffect) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = IdentifyEffect(*nt)
	return br, err
}

func (a *IdentifyEffect) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = IdentifyEffect(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a IdentifyEffect) String() string {
	switch a {
	case 0x00:
		return "Blink"
	case 0x01:
		return "Breathe"
	case 0x02:
		return "Okay"
	case 0x0B:
		return "Channel change"
	case 0xFE:
		return "Finish"
	case 0xFF:
		return "Stop"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a IdentifyEffect) IsBlink() bool         { return a == 0x00 }
func (a IdentifyEffect) IsBreathe() bool       { return a == 0x01 }
func (a IdentifyEffect) IsOkay() bool          { return a == 0x02 }
func (a IdentifyEffect) IsChannelChange() bool { return a == 0x0B }
func (a IdentifyEffect) IsFinish() bool        { return a == 0xFE }
func (a IdentifyEffect) IsStop() bool          { return a == 0xFF }
func (a *IdentifyEffect) SetBlink()            { *a = 0x00 }
func (a *IdentifyEffect) SetBreathe()          { *a = 0x01 }
func (a *IdentifyEffect) SetOkay()             { *a = 0x02 }
func (a *IdentifyEffect) SetChannelChange()    { *a = 0x0B }
func (a *IdentifyEffect) SetFinish()           { *a = 0xFE }
func (a *IdentifyEffect) SetStop()             { *a = 0xFF }

func (IdentifyEffect) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Blink"},
		{Value: 0x01, Name: "Breathe"},
		{Value: 0x02, Name: "Okay"},
		{Value: 0x0B, Name: "Channel change"},
		{Value: 0xFE, Name: "Finish"},
		{Value: 0xFF, Name: "Stop"},
	}
}

// IdentifyEffectVariant The effect identifier field specifies the identify effect to use.
type IdentifyEffectVariant zcl.Zenum8

func (IdentifyEffectVariant) Name() string                  { return "Identify Effect variant" }
func (a *IdentifyEffectVariant) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *IdentifyEffectVariant) Value() zcl.Val             { return a }
func (a IdentifyEffectVariant) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *IdentifyEffectVariant) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = IdentifyEffectVariant(*nt)
	return br, err
}

func (a *IdentifyEffectVariant) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = IdentifyEffectVariant(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a IdentifyEffectVariant) String() string {
	switch a {
	case 0x00:
		return "Default"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a IdentifyEffectVariant) IsDefault() bool { return a == 0x00 }
func (a *IdentifyEffectVariant) SetDefault()    { *a = 0x00 }

func (IdentifyEffectVariant) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Default"},
	}
}

// IdentifyTime The time in seconds for which a device will stay in identify mode.
type IdentifyTime zcl.Zu16

const IdentifyTimeAttr zcl.AttrID = 0

func (IdentifyTime) ID() zcl.AttrID   { return IdentifyTimeAttr }
func (IdentifyTime) Readable() bool   { return true }
func (IdentifyTime) Writable() bool   { return true }
func (IdentifyTime) Reportable() bool { return false }
func (IdentifyTime) SceneIndex() int  { return -1 }

func (IdentifyTime) Name() string                  { return "Identify Time" }
func (a *IdentifyTime) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *IdentifyTime) Value() zcl.Val             { return a }
func (a IdentifyTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *IdentifyTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = IdentifyTime(*nt)
	return br, err
}

func (a *IdentifyTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = IdentifyTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a IdentifyTime) String() string {
	return zcl.Seconds.Format(float64(a))
}

// IdentifyTimeout The time in seconds for which a device will stay in identify mode.
type IdentifyTimeout zcl.Zu16

func (IdentifyTimeout) Name() string                  { return "Identify Timeout" }
func (a *IdentifyTimeout) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *IdentifyTimeout) Value() zcl.Val             { return a }
func (a IdentifyTimeout) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *IdentifyTimeout) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = IdentifyTimeout(*nt)
	return br, err
}

func (a *IdentifyTimeout) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = IdentifyTimeout(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a IdentifyTimeout) String() string {
	return zcl.Seconds.Format(float64(a))
}

type IkeaRemoteDirection zcl.Zenum8

func (IkeaRemoteDirection) Name() string                  { return "Ikea Remote Direction" }
func (a *IkeaRemoteDirection) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *IkeaRemoteDirection) Value() zcl.Val             { return a }
func (a IkeaRemoteDirection) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *IkeaRemoteDirection) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = IkeaRemoteDirection(*nt)
	return br, err
}

func (a *IkeaRemoteDirection) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = IkeaRemoteDirection(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a IkeaRemoteDirection) String() string {
	switch a {
	case 0x00:
		return "Right"
	case 0x01:
		return "Left"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a IkeaRemoteDirection) IsRight() bool { return a == 0x00 }
func (a IkeaRemoteDirection) IsLeft() bool  { return a == 0x01 }
func (a *IkeaRemoteDirection) SetRight()    { *a = 0x00 }
func (a *IkeaRemoteDirection) SetLeft()     { *a = 0x01 }

func (IkeaRemoteDirection) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Right"},
		{Value: 0x01, Name: "Left"},
	}
}

type JoinIndication zcl.Zu16

const JoinIndicationAttr zcl.AttrID = 272

func (JoinIndication) ID() zcl.AttrID   { return JoinIndicationAttr }
func (JoinIndication) Readable() bool   { return true }
func (JoinIndication) Writable() bool   { return false }
func (JoinIndication) Reportable() bool { return false }
func (JoinIndication) SceneIndex() int  { return -1 }

func (JoinIndication) Name() string                  { return "Join Indication" }
func (a *JoinIndication) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *JoinIndication) Value() zcl.Val             { return a }
func (a JoinIndication) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *JoinIndication) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = JoinIndication(*nt)
	return br, err
}

func (a *JoinIndication) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = JoinIndication(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a JoinIndication) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type LedIndication zcl.Zbool

const LedIndicationAttr zcl.AttrID = 51

func (LedIndication) ID() zcl.AttrID   { return LedIndicationAttr }
func (LedIndication) Readable() bool   { return true }
func (LedIndication) Writable() bool   { return true }
func (LedIndication) Reportable() bool { return false }
func (LedIndication) SceneIndex() int  { return -1 }

func (LedIndication) Name() string                  { return "LED Indication" }
func (a *LedIndication) TypeID() zcl.TypeID         { return zcl.Zbool(*a).ID() }
func (a *LedIndication) Value() zcl.Val             { return a }
func (a LedIndication) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *LedIndication) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = LedIndication(*nt)
	return br, err
}

func (a *LedIndication) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbool); ok {
		*a = LedIndication(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LedIndication) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

type LastMessageLqi zcl.Zu8

const LastMessageLqiAttr zcl.AttrID = 284

func (LastMessageLqi) ID() zcl.AttrID   { return LastMessageLqiAttr }
func (LastMessageLqi) Readable() bool   { return true }
func (LastMessageLqi) Writable() bool   { return false }
func (LastMessageLqi) Reportable() bool { return false }
func (LastMessageLqi) SceneIndex() int  { return -1 }

func (LastMessageLqi) Name() string                  { return "Last Message LQI" }
func (a *LastMessageLqi) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *LastMessageLqi) Value() zcl.Val             { return a }
func (a LastMessageLqi) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *LastMessageLqi) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = LastMessageLqi(*nt)
	return br, err
}

func (a *LastMessageLqi) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = LastMessageLqi(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LastMessageLqi) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type LastMessageRssi zcl.Zs8

const LastMessageRssiAttr zcl.AttrID = 285

func (LastMessageRssi) ID() zcl.AttrID   { return LastMessageRssiAttr }
func (LastMessageRssi) Readable() bool   { return true }
func (LastMessageRssi) Writable() bool   { return false }
func (LastMessageRssi) Reportable() bool { return false }
func (LastMessageRssi) SceneIndex() int  { return -1 }

func (LastMessageRssi) Name() string                  { return "Last Message RSSI" }
func (a *LastMessageRssi) TypeID() zcl.TypeID         { return zcl.Zs8(*a).ID() }
func (a *LastMessageRssi) Value() zcl.Val             { return a }
func (a LastMessageRssi) MarshalZcl() ([]byte, error) { return zcl.Zs8(a).MarshalZcl() }

func (a *LastMessageRssi) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = LastMessageRssi(*nt)
	return br, err
}

func (a *LastMessageRssi) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs8); ok {
		*a = LastMessageRssi(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LastMessageRssi) String() string {
	return zcl.Sprintf("%v", zcl.Zs8(a))
}

type LastSetTime zcl.Zutc

const LastSetTimeAttr zcl.AttrID = 8

func (LastSetTime) ID() zcl.AttrID   { return LastSetTimeAttr }
func (LastSetTime) Readable() bool   { return true }
func (LastSetTime) Writable() bool   { return false }
func (LastSetTime) Reportable() bool { return false }
func (LastSetTime) SceneIndex() int  { return -1 }

func (LastSetTime) Name() string                  { return "Last Set Time" }
func (a *LastSetTime) TypeID() zcl.TypeID         { return zcl.Zutc(*a).ID() }
func (a *LastSetTime) Value() zcl.Val             { return a }
func (a LastSetTime) MarshalZcl() ([]byte, error) { return zcl.Zutc(a).MarshalZcl() }

func (a *LastSetTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*a = LastSetTime(*nt)
	return br, err
}

func (a *LastSetTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zutc); ok {
		*a = LastSetTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LastSetTime) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(a))
}

type Level zcl.Zu8

func (Level) Name() string                  { return "Level" }
func (a *Level) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *Level) Value() zcl.Val             { return a }
func (a Level) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Level) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Level(*nt)
	return br, err
}

func (a *Level) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = Level(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Level) String() string {
	return zcl.Percent.Format(float64(a) / 2.54)
}

// LevelControlOptions is a bitmap that determines the default behavior of some cluster commands
type LevelControlOptions zcl.Zbmp8

const LevelControlOptionsAttr zcl.AttrID = 15

func (LevelControlOptions) ID() zcl.AttrID   { return LevelControlOptionsAttr }
func (LevelControlOptions) Readable() bool   { return true }
func (LevelControlOptions) Writable() bool   { return true }
func (LevelControlOptions) Reportable() bool { return false }
func (LevelControlOptions) SceneIndex() int  { return -1 }

func (LevelControlOptions) Name() string                  { return "Level Control Options" }
func (a *LevelControlOptions) TypeID() zcl.TypeID         { return zcl.Zbmp8(*a).ID() }
func (a *LevelControlOptions) Value() zcl.Val             { return a }
func (a LevelControlOptions) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *LevelControlOptions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = LevelControlOptions(*nt)
	return br, err
}

func (a *LevelControlOptions) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = LevelControlOptions(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LevelControlOptions) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0x00:
			bstr = append(bstr, "Execute if off")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a LevelControlOptions) IsExecuteIfOff() bool { return zcl.BitmapTest([]byte(a[:]), 0x00) }
func (a *LevelControlOptions) SetExecuteIfOff(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0x00, b))
}

func (LevelControlOptions) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Execute if off"},
	}
}

type LevelDirection zcl.Zenum8

func (LevelDirection) Name() string                  { return "Level direction" }
func (a *LevelDirection) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *LevelDirection) Value() zcl.Val             { return a }
func (a LevelDirection) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *LevelDirection) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = LevelDirection(*nt)
	return br, err
}

func (a *LevelDirection) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = LevelDirection(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LevelDirection) String() string {
	switch a {
	case 0x00:
		return "Up"
	case 0x01:
		return "Down"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a LevelDirection) IsUp() bool   { return a == 0x00 }
func (a LevelDirection) IsDown() bool { return a == 0x01 }
func (a *LevelDirection) SetUp()      { *a = 0x00 }
func (a *LevelDirection) SetDown()    { *a = 0x01 }

func (LevelDirection) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Up"},
		{Value: 0x01, Name: "Down"},
	}
}

// LocalTime Local time
type LocalTime zcl.Zu32

const LocalTimeAttr zcl.AttrID = 7

func (LocalTime) ID() zcl.AttrID   { return LocalTimeAttr }
func (LocalTime) Readable() bool   { return true }
func (LocalTime) Writable() bool   { return false }
func (LocalTime) Reportable() bool { return false }
func (LocalTime) SceneIndex() int  { return -1 }

func (LocalTime) Name() string                  { return "Local Time" }
func (a *LocalTime) TypeID() zcl.TypeID         { return zcl.Zu32(*a).ID() }
func (a *LocalTime) Value() zcl.Val             { return a }
func (a LocalTime) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *LocalTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = LocalTime(*nt)
	return br, err
}

func (a *LocalTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu32); ok {
		*a = LocalTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LocalTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

type LocationAge zcl.Zu16

const LocationAgeAttr zcl.AttrID = 2

func (LocationAge) ID() zcl.AttrID   { return LocationAgeAttr }
func (LocationAge) Readable() bool   { return true }
func (LocationAge) Writable() bool   { return false }
func (LocationAge) Reportable() bool { return false }
func (LocationAge) SceneIndex() int  { return -1 }

func (LocationAge) Name() string                  { return "Location Age" }
func (a *LocationAge) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *LocationAge) Value() zcl.Val             { return a }
func (a LocationAge) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *LocationAge) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LocationAge(*nt)
	return br, err
}

func (a *LocationAge) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = LocationAge(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LocationAge) String() string {
	return zcl.Seconds.Format(float64(a))
}

type LocationDescription zcl.Zcstring

const LocationDescriptionAttr zcl.AttrID = 16

func (LocationDescription) ID() zcl.AttrID   { return LocationDescriptionAttr }
func (LocationDescription) Readable() bool   { return true }
func (LocationDescription) Writable() bool   { return true }
func (LocationDescription) Reportable() bool { return false }
func (LocationDescription) SceneIndex() int  { return -1 }

func (LocationDescription) Name() string                  { return "Location Description" }
func (a *LocationDescription) TypeID() zcl.TypeID         { return zcl.Zcstring(*a).ID() }
func (a *LocationDescription) Value() zcl.Val             { return a }
func (a LocationDescription) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *LocationDescription) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = LocationDescription(*nt)
	return br, err
}

func (a *LocationDescription) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = LocationDescription(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LocationDescription) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

type LocationMethod zcl.Zenum8

const LocationMethodAttr zcl.AttrID = 1

func (LocationMethod) ID() zcl.AttrID   { return LocationMethodAttr }
func (LocationMethod) Readable() bool   { return true }
func (LocationMethod) Writable() bool   { return true }
func (LocationMethod) Reportable() bool { return false }
func (LocationMethod) SceneIndex() int  { return -1 }

func (LocationMethod) Name() string                  { return "Location Method" }
func (a *LocationMethod) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *LocationMethod) Value() zcl.Val             { return a }
func (a LocationMethod) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *LocationMethod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = LocationMethod(*nt)
	return br, err
}

func (a *LocationMethod) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = LocationMethod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LocationMethod) String() string {
	switch a {
	case 0x00:
		return "Lateration"
	case 0x01:
		return "Signposting"
	case 0x02:
		return "RF fingerprinting"
	case 0x03:
		return "Out of band"
	case 0x04:
		return "Centralized"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a LocationMethod) IsLateration() bool       { return a == 0x00 }
func (a LocationMethod) IsSignposting() bool      { return a == 0x01 }
func (a LocationMethod) IsRfFingerprinting() bool { return a == 0x02 }
func (a LocationMethod) IsOutOfBand() bool        { return a == 0x03 }
func (a LocationMethod) IsCentralized() bool      { return a == 0x04 }
func (a *LocationMethod) SetLateration()          { *a = 0x00 }
func (a *LocationMethod) SetSignposting()         { *a = 0x01 }
func (a *LocationMethod) SetRfFingerprinting()    { *a = 0x02 }
func (a *LocationMethod) SetOutOfBand()           { *a = 0x03 }
func (a *LocationMethod) SetCentralized()         { *a = 0x04 }

func (LocationMethod) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Lateration"},
		{Value: 0x01, Name: "Signposting"},
		{Value: 0x02, Name: "RF fingerprinting"},
		{Value: 0x03, Name: "Out of band"},
		{Value: 0x04, Name: "Centralized"},
	}
}

type LocationType zcl.Zenum8

const LocationTypeAttr zcl.AttrID = 0

func (LocationType) ID() zcl.AttrID   { return LocationTypeAttr }
func (LocationType) Readable() bool   { return true }
func (LocationType) Writable() bool   { return true }
func (LocationType) Reportable() bool { return false }
func (LocationType) SceneIndex() int  { return -1 }

func (LocationType) Name() string                  { return "Location Type" }
func (a *LocationType) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *LocationType) Value() zcl.Val             { return a }
func (a LocationType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *LocationType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = LocationType(*nt)
	return br, err
}

func (a *LocationType) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = LocationType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LocationType) String() string {
	switch a {
	case 0x00:
		return "3D Location"
	case 0x01:
		return "Absolute 3D Location"
	case 0x02:
		return "2D Location"
	case 0x03:
		return "Absolute 2D Location"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a LocationType) Is3DLocation() bool         { return a == 0x00 }
func (a LocationType) IsAbsolute3DLocation() bool { return a == 0x01 }
func (a LocationType) Is2DLocation() bool         { return a == 0x02 }
func (a LocationType) IsAbsolute2DLocation() bool { return a == 0x03 }
func (a *LocationType) Set3DLocation()            { *a = 0x00 }
func (a *LocationType) SetAbsolute3DLocation()    { *a = 0x01 }
func (a *LocationType) Set2DLocation()            { *a = 0x02 }
func (a *LocationType) SetAbsolute2DLocation()    { *a = 0x03 }

func (LocationType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "3D Location"},
		{Value: 0x01, Name: "Absolute 3D Location"},
		{Value: 0x02, Name: "2D Location"},
		{Value: 0x03, Name: "Absolute 2D Location"},
	}
}

type LocationFlags zcl.Zbmp8

func (LocationFlags) Name() string                  { return "Location flags" }
func (a *LocationFlags) TypeID() zcl.TypeID         { return zcl.Zbmp8(*a).ID() }
func (a *LocationFlags) Value() zcl.Val             { return a }
func (a LocationFlags) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *LocationFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = LocationFlags(*nt)
	return br, err
}

func (a *LocationFlags) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = LocationFlags(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LocationFlags) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Absolute Only")
		case 1:
			bstr = append(bstr, "Recalculate")
		case 2:
			bstr = append(bstr, "Broadcast Indicator")
		case 3:
			bstr = append(bstr, "Broadcast Response")
		case 4:
			bstr = append(bstr, "Compact Response")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a LocationFlags) IsAbsoluteOnly() bool       { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a LocationFlags) IsRecalculate() bool        { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a LocationFlags) IsBroadcastIndicator() bool { return zcl.BitmapTest([]byte(a[:]), 2) }
func (a LocationFlags) IsBroadcastResponse() bool  { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a LocationFlags) IsCompactResponse() bool    { return zcl.BitmapTest([]byte(a[:]), 4) }
func (a *LocationFlags) SetAbsoluteOnly(b bool)    { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b)) }
func (a *LocationFlags) SetRecalculate(b bool)     { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b)) }
func (a *LocationFlags) SetBroadcastIndicator(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b))
}
func (a *LocationFlags) SetBroadcastResponse(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b))
}
func (a *LocationFlags) SetCompactResponse(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 4, b))
}

func (LocationFlags) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Absolute Only"},
		{Value: 1, Name: "Recalculate"},
		{Value: 2, Name: "Broadcast Indicator"},
		{Value: 3, Name: "Broadcast Response"},
		{Value: 4, Name: "Compact Response"},
	}
}

type LongPollInterval zcl.Zu32

const LongPollIntervalAttr zcl.AttrID = 1

func (LongPollInterval) ID() zcl.AttrID   { return LongPollIntervalAttr }
func (LongPollInterval) Readable() bool   { return true }
func (LongPollInterval) Writable() bool   { return false }
func (LongPollInterval) Reportable() bool { return false }
func (LongPollInterval) SceneIndex() int  { return -1 }

func (LongPollInterval) Name() string                  { return "Long Poll Interval" }
func (a *LongPollInterval) TypeID() zcl.TypeID         { return zcl.Zu32(*a).ID() }
func (a *LongPollInterval) Value() zcl.Val             { return a }
func (a LongPollInterval) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *LongPollInterval) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = LongPollInterval(*nt)
	return br, err
}

func (a *LongPollInterval) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu32); ok {
		*a = LongPollInterval(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LongPollInterval) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

type LongPollIntervalMin zcl.Zu32

const LongPollIntervalMinAttr zcl.AttrID = 5

func (LongPollIntervalMin) ID() zcl.AttrID   { return LongPollIntervalMinAttr }
func (LongPollIntervalMin) Readable() bool   { return true }
func (LongPollIntervalMin) Writable() bool   { return false }
func (LongPollIntervalMin) Reportable() bool { return false }
func (LongPollIntervalMin) SceneIndex() int  { return -1 }

func (LongPollIntervalMin) Name() string                  { return "Long Poll Interval Min" }
func (a *LongPollIntervalMin) TypeID() zcl.TypeID         { return zcl.Zu32(*a).ID() }
func (a *LongPollIntervalMin) Value() zcl.Val             { return a }
func (a LongPollIntervalMin) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *LongPollIntervalMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = LongPollIntervalMin(*nt)
	return br, err
}

func (a *LongPollIntervalMin) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu32); ok {
		*a = LongPollIntervalMin(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LongPollIntervalMin) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

type LowTempDwellTripPoint zcl.Zu24

const LowTempDwellTripPointAttr zcl.AttrID = 19

func (LowTempDwellTripPoint) ID() zcl.AttrID   { return LowTempDwellTripPointAttr }
func (LowTempDwellTripPoint) Readable() bool   { return false }
func (LowTempDwellTripPoint) Writable() bool   { return false }
func (LowTempDwellTripPoint) Reportable() bool { return false }
func (LowTempDwellTripPoint) SceneIndex() int  { return -1 }

func (LowTempDwellTripPoint) Name() string                  { return "Low Temp Dwell Trip Point" }
func (a *LowTempDwellTripPoint) TypeID() zcl.TypeID         { return zcl.Zu24(*a).ID() }
func (a *LowTempDwellTripPoint) Value() zcl.Val             { return a }
func (a LowTempDwellTripPoint) MarshalZcl() ([]byte, error) { return zcl.Zu24(a).MarshalZcl() }

func (a *LowTempDwellTripPoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*a = LowTempDwellTripPoint(*nt)
	return br, err
}

func (a *LowTempDwellTripPoint) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu24); ok {
		*a = LowTempDwellTripPoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LowTempDwellTripPoint) String() string {
	return zcl.Seconds.Format(float64(a))
}

// LowTempThreshold If the current temperature drops below the threshold for longer
// than the time specified by low temp dwell, an alarm will be triggered
type LowTempThreshold zcl.Zs16

const LowTempThresholdAttr zcl.AttrID = 17

func (LowTempThreshold) ID() zcl.AttrID   { return LowTempThresholdAttr }
func (LowTempThreshold) Readable() bool   { return true }
func (LowTempThreshold) Writable() bool   { return true }
func (LowTempThreshold) Reportable() bool { return false }
func (LowTempThreshold) SceneIndex() int  { return -1 }

func (LowTempThreshold) Name() string                  { return "Low Temp Threshold" }
func (a *LowTempThreshold) TypeID() zcl.TypeID         { return zcl.Zs16(*a).ID() }
func (a *LowTempThreshold) Value() zcl.Val             { return a }
func (a LowTempThreshold) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *LowTempThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = LowTempThreshold(*nt)
	return br, err
}

func (a *LowTempThreshold) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = LowTempThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LowTempThreshold) String() string {
	return zcl.DegreesCelsius.Format(float64(a))
}

type MacRxBcast zcl.Zu32

const MacRxBcastAttr zcl.AttrID = 256

func (MacRxBcast) ID() zcl.AttrID   { return MacRxBcastAttr }
func (MacRxBcast) Readable() bool   { return true }
func (MacRxBcast) Writable() bool   { return false }
func (MacRxBcast) Reportable() bool { return false }
func (MacRxBcast) SceneIndex() int  { return -1 }

func (MacRxBcast) Name() string                  { return "Mac Rx Bcast" }
func (a *MacRxBcast) TypeID() zcl.TypeID         { return zcl.Zu32(*a).ID() }
func (a *MacRxBcast) Value() zcl.Val             { return a }
func (a MacRxBcast) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *MacRxBcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = MacRxBcast(*nt)
	return br, err
}

func (a *MacRxBcast) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu32); ok {
		*a = MacRxBcast(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MacRxBcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

type MacRxUcast zcl.Zu32

const MacRxUcastAttr zcl.AttrID = 258

func (MacRxUcast) ID() zcl.AttrID   { return MacRxUcastAttr }
func (MacRxUcast) Readable() bool   { return true }
func (MacRxUcast) Writable() bool   { return false }
func (MacRxUcast) Reportable() bool { return false }
func (MacRxUcast) SceneIndex() int  { return -1 }

func (MacRxUcast) Name() string                  { return "Mac Rx Ucast" }
func (a *MacRxUcast) TypeID() zcl.TypeID         { return zcl.Zu32(*a).ID() }
func (a *MacRxUcast) Value() zcl.Val             { return a }
func (a MacRxUcast) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *MacRxUcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = MacRxUcast(*nt)
	return br, err
}

func (a *MacRxUcast) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu32); ok {
		*a = MacRxUcast(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MacRxUcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

type MacTxBcast zcl.Zu32

const MacTxBcastAttr zcl.AttrID = 257

func (MacTxBcast) ID() zcl.AttrID   { return MacTxBcastAttr }
func (MacTxBcast) Readable() bool   { return true }
func (MacTxBcast) Writable() bool   { return false }
func (MacTxBcast) Reportable() bool { return false }
func (MacTxBcast) SceneIndex() int  { return -1 }

func (MacTxBcast) Name() string                  { return "Mac Tx Bcast" }
func (a *MacTxBcast) TypeID() zcl.TypeID         { return zcl.Zu32(*a).ID() }
func (a *MacTxBcast) Value() zcl.Val             { return a }
func (a MacTxBcast) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *MacTxBcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = MacTxBcast(*nt)
	return br, err
}

func (a *MacTxBcast) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu32); ok {
		*a = MacTxBcast(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MacTxBcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

type MacTxUcast zcl.Zu32

const MacTxUcastAttr zcl.AttrID = 259

func (MacTxUcast) ID() zcl.AttrID   { return MacTxUcastAttr }
func (MacTxUcast) Readable() bool   { return true }
func (MacTxUcast) Writable() bool   { return false }
func (MacTxUcast) Reportable() bool { return false }
func (MacTxUcast) SceneIndex() int  { return -1 }

func (MacTxUcast) Name() string                  { return "Mac Tx Ucast" }
func (a *MacTxUcast) TypeID() zcl.TypeID         { return zcl.Zu32(*a).ID() }
func (a *MacTxUcast) Value() zcl.Val             { return a }
func (a MacTxUcast) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *MacTxUcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = MacTxUcast(*nt)
	return br, err
}

func (a *MacTxUcast) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu32); ok {
		*a = MacTxUcast(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MacTxUcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

type MacTxUcastFail zcl.Zu16

const MacTxUcastFailAttr zcl.AttrID = 261

func (MacTxUcastFail) ID() zcl.AttrID   { return MacTxUcastFailAttr }
func (MacTxUcastFail) Readable() bool   { return true }
func (MacTxUcastFail) Writable() bool   { return false }
func (MacTxUcastFail) Reportable() bool { return false }
func (MacTxUcastFail) SceneIndex() int  { return -1 }

func (MacTxUcastFail) Name() string                  { return "Mac Tx Ucast Fail" }
func (a *MacTxUcastFail) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *MacTxUcastFail) Value() zcl.Val             { return a }
func (a MacTxUcastFail) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *MacTxUcastFail) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MacTxUcastFail(*nt)
	return br, err
}

func (a *MacTxUcastFail) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = MacTxUcastFail(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MacTxUcastFail) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type MacTxUcastRetry zcl.Zu16

const MacTxUcastRetryAttr zcl.AttrID = 260

func (MacTxUcastRetry) ID() zcl.AttrID   { return MacTxUcastRetryAttr }
func (MacTxUcastRetry) Readable() bool   { return true }
func (MacTxUcastRetry) Writable() bool   { return false }
func (MacTxUcastRetry) Reportable() bool { return false }
func (MacTxUcastRetry) SceneIndex() int  { return -1 }

func (MacTxUcastRetry) Name() string                  { return "Mac Tx Ucast Retry" }
func (a *MacTxUcastRetry) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *MacTxUcastRetry) Value() zcl.Val             { return a }
func (a MacTxUcastRetry) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *MacTxUcastRetry) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MacTxUcastRetry(*nt)
	return br, err
}

func (a *MacTxUcastRetry) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = MacTxUcastRetry(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MacTxUcastRetry) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type MainsAlarmMask zcl.Zbmp8

const MainsAlarmMaskAttr zcl.AttrID = 16

func (MainsAlarmMask) ID() zcl.AttrID   { return MainsAlarmMaskAttr }
func (MainsAlarmMask) Readable() bool   { return true }
func (MainsAlarmMask) Writable() bool   { return true }
func (MainsAlarmMask) Reportable() bool { return false }
func (MainsAlarmMask) SceneIndex() int  { return -1 }

func (MainsAlarmMask) Name() string                  { return "Mains Alarm Mask" }
func (a *MainsAlarmMask) TypeID() zcl.TypeID         { return zcl.Zbmp8(*a).ID() }
func (a *MainsAlarmMask) Value() zcl.Val             { return a }
func (a MainsAlarmMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *MainsAlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = MainsAlarmMask(*nt)
	return br, err
}

func (a *MainsAlarmMask) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = MainsAlarmMask(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MainsAlarmMask) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Mains Voltage too low")
		case 1:
			bstr = append(bstr, "Mains Voltage too high")
		case 2:
			bstr = append(bstr, "Mains power supply lost/unavailable")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a MainsAlarmMask) IsMainsVoltageTooLow() bool  { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a MainsAlarmMask) IsMainsVoltageTooHigh() bool { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a MainsAlarmMask) IsMainsPowerSupplyLostUnavailable() bool {
	return zcl.BitmapTest([]byte(a[:]), 2)
}
func (a *MainsAlarmMask) SetMainsVoltageTooLow(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *MainsAlarmMask) SetMainsVoltageTooHigh(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}
func (a *MainsAlarmMask) SetMainsPowerSupplyLostUnavailable(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b))
}

func (MainsAlarmMask) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Mains Voltage too low"},
		{Value: 1, Name: "Mains Voltage too high"},
		{Value: 2, Name: "Mains power supply lost/unavailable"},
	}
}

// MainsFrequency Resolution of 2Hz
// Special values:
// * 0x00 frequency too low to measure
// * 0xfe frequency too high to measure
// * 0xff unable to measure
type MainsFrequency zcl.Zu8

const MainsFrequencyAttr zcl.AttrID = 1

func (MainsFrequency) ID() zcl.AttrID   { return MainsFrequencyAttr }
func (MainsFrequency) Readable() bool   { return true }
func (MainsFrequency) Writable() bool   { return false }
func (MainsFrequency) Reportable() bool { return false }
func (MainsFrequency) SceneIndex() int  { return -1 }

func (MainsFrequency) Name() string                  { return "Mains Frequency" }
func (a *MainsFrequency) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *MainsFrequency) Value() zcl.Val             { return a }
func (a MainsFrequency) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *MainsFrequency) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = MainsFrequency(*nt)
	return br, err
}

func (a *MainsFrequency) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = MainsFrequency(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MainsFrequency) String() string {
	return zcl.Hertz.Format(float64(a) / 2)
}

type MainsVoltage zcl.Zu16

const MainsVoltageAttr zcl.AttrID = 0

func (MainsVoltage) ID() zcl.AttrID   { return MainsVoltageAttr }
func (MainsVoltage) Readable() bool   { return true }
func (MainsVoltage) Writable() bool   { return false }
func (MainsVoltage) Reportable() bool { return false }
func (MainsVoltage) SceneIndex() int  { return -1 }

func (MainsVoltage) Name() string                  { return "Mains Voltage" }
func (a *MainsVoltage) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *MainsVoltage) Value() zcl.Val             { return a }
func (a MainsVoltage) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *MainsVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MainsVoltage(*nt)
	return br, err
}

func (a *MainsVoltage) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = MainsVoltage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MainsVoltage) String() string {
	return zcl.Volts.Format(float64(a) / 10)
}

// MainsVoltageDwellTripPoint Length of time that the value of MainsVoltage MAY exist beyond either
// of its thresholds before an alarm is generated
type MainsVoltageDwellTripPoint zcl.Zu16

const MainsVoltageDwellTripPointAttr zcl.AttrID = 19

func (MainsVoltageDwellTripPoint) ID() zcl.AttrID   { return MainsVoltageDwellTripPointAttr }
func (MainsVoltageDwellTripPoint) Readable() bool   { return true }
func (MainsVoltageDwellTripPoint) Writable() bool   { return true }
func (MainsVoltageDwellTripPoint) Reportable() bool { return false }
func (MainsVoltageDwellTripPoint) SceneIndex() int  { return -1 }

func (MainsVoltageDwellTripPoint) Name() string                  { return "Mains Voltage Dwell Trip Point" }
func (a *MainsVoltageDwellTripPoint) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *MainsVoltageDwellTripPoint) Value() zcl.Val             { return a }
func (a MainsVoltageDwellTripPoint) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *MainsVoltageDwellTripPoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MainsVoltageDwellTripPoint(*nt)
	return br, err
}

func (a *MainsVoltageDwellTripPoint) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = MainsVoltageDwellTripPoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MainsVoltageDwellTripPoint) String() string {
	return zcl.Seconds.Format(float64(a))
}

type MainsVoltageMaxThreshold zcl.Zu16

const MainsVoltageMaxThresholdAttr zcl.AttrID = 18

func (MainsVoltageMaxThreshold) ID() zcl.AttrID   { return MainsVoltageMaxThresholdAttr }
func (MainsVoltageMaxThreshold) Readable() bool   { return true }
func (MainsVoltageMaxThreshold) Writable() bool   { return true }
func (MainsVoltageMaxThreshold) Reportable() bool { return false }
func (MainsVoltageMaxThreshold) SceneIndex() int  { return -1 }

func (MainsVoltageMaxThreshold) Name() string                  { return "Mains Voltage Max Threshold" }
func (a *MainsVoltageMaxThreshold) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *MainsVoltageMaxThreshold) Value() zcl.Val             { return a }
func (a MainsVoltageMaxThreshold) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *MainsVoltageMaxThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MainsVoltageMaxThreshold(*nt)
	return br, err
}

func (a *MainsVoltageMaxThreshold) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = MainsVoltageMaxThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MainsVoltageMaxThreshold) String() string {
	return zcl.Volts.Format(float64(a) / 10)
}

type MainsVoltageMinThreshold zcl.Zu16

const MainsVoltageMinThresholdAttr zcl.AttrID = 17

func (MainsVoltageMinThreshold) ID() zcl.AttrID   { return MainsVoltageMinThresholdAttr }
func (MainsVoltageMinThreshold) Readable() bool   { return true }
func (MainsVoltageMinThreshold) Writable() bool   { return true }
func (MainsVoltageMinThreshold) Reportable() bool { return false }
func (MainsVoltageMinThreshold) SceneIndex() int  { return -1 }

func (MainsVoltageMinThreshold) Name() string                  { return "Mains Voltage Min Threshold" }
func (a *MainsVoltageMinThreshold) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *MainsVoltageMinThreshold) Value() zcl.Val             { return a }
func (a MainsVoltageMinThreshold) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *MainsVoltageMinThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MainsVoltageMinThreshold(*nt)
	return br, err
}

func (a *MainsVoltageMinThreshold) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = MainsVoltageMinThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MainsVoltageMinThreshold) String() string {
	return zcl.Volts.Format(float64(a) / 10)
}

type ManufacturerName zcl.Zcstring

const ManufacturerNameAttr zcl.AttrID = 4

func (ManufacturerName) ID() zcl.AttrID   { return ManufacturerNameAttr }
func (ManufacturerName) Readable() bool   { return true }
func (ManufacturerName) Writable() bool   { return false }
func (ManufacturerName) Reportable() bool { return false }
func (ManufacturerName) SceneIndex() int  { return -1 }

func (ManufacturerName) Name() string                  { return "Manufacturer Name" }
func (a *ManufacturerName) TypeID() zcl.TypeID         { return zcl.Zcstring(*a).ID() }
func (a *ManufacturerName) Value() zcl.Val             { return a }
func (a ManufacturerName) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *ManufacturerName) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = ManufacturerName(*nt)
	return br, err
}

func (a *ManufacturerName) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = ManufacturerName(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ManufacturerName) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

type MaxTempExperienced zcl.Zs16

const MaxTempExperiencedAttr zcl.AttrID = 2

func (MaxTempExperienced) ID() zcl.AttrID   { return MaxTempExperiencedAttr }
func (MaxTempExperienced) Readable() bool   { return true }
func (MaxTempExperienced) Writable() bool   { return false }
func (MaxTempExperienced) Reportable() bool { return false }
func (MaxTempExperienced) SceneIndex() int  { return -1 }

func (MaxTempExperienced) Name() string                  { return "Max Temp Experienced" }
func (a *MaxTempExperienced) TypeID() zcl.TypeID         { return zcl.Zs16(*a).ID() }
func (a *MaxTempExperienced) Value() zcl.Val             { return a }
func (a MaxTempExperienced) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *MaxTempExperienced) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxTempExperienced(*nt)
	return br, err
}

func (a *MaxTempExperienced) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = MaxTempExperienced(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MaxTempExperienced) String() string {
	return zcl.DegreesCelsius.Format(float64(a))
}

type MinTempExperienced zcl.Zs16

const MinTempExperiencedAttr zcl.AttrID = 1

func (MinTempExperienced) ID() zcl.AttrID   { return MinTempExperiencedAttr }
func (MinTempExperienced) Readable() bool   { return true }
func (MinTempExperienced) Writable() bool   { return false }
func (MinTempExperienced) Reportable() bool { return false }
func (MinTempExperienced) SceneIndex() int  { return -1 }

func (MinTempExperienced) Name() string                  { return "Min Temp Experienced" }
func (a *MinTempExperienced) TypeID() zcl.TypeID         { return zcl.Zs16(*a).ID() }
func (a *MinTempExperienced) Value() zcl.Val             { return a }
func (a MinTempExperienced) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *MinTempExperienced) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MinTempExperienced(*nt)
	return br, err
}

func (a *MinTempExperienced) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = MinTempExperienced(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MinTempExperienced) String() string {
	return zcl.DegreesCelsius.Format(float64(a))
}

type ModelIdentifier zcl.Zcstring

const ModelIdentifierAttr zcl.AttrID = 5

func (ModelIdentifier) ID() zcl.AttrID   { return ModelIdentifierAttr }
func (ModelIdentifier) Readable() bool   { return true }
func (ModelIdentifier) Writable() bool   { return false }
func (ModelIdentifier) Reportable() bool { return false }
func (ModelIdentifier) SceneIndex() int  { return -1 }

func (ModelIdentifier) Name() string                  { return "Model Identifier" }
func (a *ModelIdentifier) TypeID() zcl.TypeID         { return zcl.Zcstring(*a).ID() }
func (a *ModelIdentifier) Value() zcl.Val             { return a }
func (a ModelIdentifier) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *ModelIdentifier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = ModelIdentifier(*nt)
	return br, err
}

func (a *ModelIdentifier) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = ModelIdentifier(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ModelIdentifier) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

type MultistateNumberOfStates zcl.Zu16

const MultistateNumberOfStatesAttr zcl.AttrID = 74

func (MultistateNumberOfStates) ID() zcl.AttrID   { return MultistateNumberOfStatesAttr }
func (MultistateNumberOfStates) Readable() bool   { return true }
func (MultistateNumberOfStates) Writable() bool   { return true }
func (MultistateNumberOfStates) Reportable() bool { return false }
func (MultistateNumberOfStates) SceneIndex() int  { return -1 }

func (MultistateNumberOfStates) Name() string                  { return "Multistate Number of States" }
func (a *MultistateNumberOfStates) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *MultistateNumberOfStates) Value() zcl.Val             { return a }
func (a MultistateNumberOfStates) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *MultistateNumberOfStates) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MultistateNumberOfStates(*nt)
	return br, err
}

func (a *MultistateNumberOfStates) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = MultistateNumberOfStates(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MultistateNumberOfStates) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type MultistatePresentValue zcl.Zbmp16

const MultistatePresentValueAttr zcl.AttrID = 85

func (MultistatePresentValue) ID() zcl.AttrID   { return MultistatePresentValueAttr }
func (MultistatePresentValue) Readable() bool   { return true }
func (MultistatePresentValue) Writable() bool   { return true }
func (MultistatePresentValue) Reportable() bool { return true }
func (MultistatePresentValue) SceneIndex() int  { return -1 }

func (MultistatePresentValue) Name() string                  { return "Multistate Present value" }
func (a *MultistatePresentValue) TypeID() zcl.TypeID         { return zcl.Zbmp16(*a).ID() }
func (a *MultistatePresentValue) Value() zcl.Val             { return a }
func (a MultistatePresentValue) MarshalZcl() ([]byte, error) { return zcl.Zbmp16(a).MarshalZcl() }

func (a *MultistatePresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*a = MultistatePresentValue(*nt)
	return br, err
}

func (a *MultistatePresentValue) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp16); ok {
		*a = MultistatePresentValue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MultistatePresentValue) String() string {
	return zcl.Sprintf("%v", zcl.Zbmp16(a))
}

type MultistatePriorityArray zcl.Zarray

const MultistatePriorityArrayAttr zcl.AttrID = 87

func (MultistatePriorityArray) ID() zcl.AttrID   { return MultistatePriorityArrayAttr }
func (MultistatePriorityArray) Readable() bool   { return true }
func (MultistatePriorityArray) Writable() bool   { return true }
func (MultistatePriorityArray) Reportable() bool { return false }
func (MultistatePriorityArray) SceneIndex() int  { return -1 }

func (MultistatePriorityArray) Name() string          { return "Multistate Priority Array" }
func (a *MultistatePriorityArray) TypeID() zcl.TypeID { return zcl.Zarray(*a).ID() }
func (a *MultistatePriorityArray) Value() zcl.Val     { return a }
func (a MultistatePriorityArray) MarshalZcl() ([]byte, error) {
	return nil, zcl.Errorf("not implemented")
}

func (a *MultistatePriorityArray) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zarray)
	br, err := nt.UnmarshalZcl(b)
	*a = MultistatePriorityArray(*nt)
	return br, err
}

func (a *MultistatePriorityArray) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zarray); ok {
		*a = MultistatePriorityArray(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MultistatePriorityArray) String() string {
	return zcl.Sprintf("%v", zcl.Zarray(a))
}

type MultistateRelinquishDefault zcl.Zbmp16

const MultistateRelinquishDefaultAttr zcl.AttrID = 104

func (MultistateRelinquishDefault) ID() zcl.AttrID   { return MultistateRelinquishDefaultAttr }
func (MultistateRelinquishDefault) Readable() bool   { return true }
func (MultistateRelinquishDefault) Writable() bool   { return true }
func (MultistateRelinquishDefault) Reportable() bool { return false }
func (MultistateRelinquishDefault) SceneIndex() int  { return -1 }

func (MultistateRelinquishDefault) Name() string                  { return "Multistate Relinquish Default" }
func (a *MultistateRelinquishDefault) TypeID() zcl.TypeID         { return zcl.Zbmp16(*a).ID() }
func (a *MultistateRelinquishDefault) Value() zcl.Val             { return a }
func (a MultistateRelinquishDefault) MarshalZcl() ([]byte, error) { return zcl.Zbmp16(a).MarshalZcl() }

func (a *MultistateRelinquishDefault) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*a = MultistateRelinquishDefault(*nt)
	return br, err
}

func (a *MultistateRelinquishDefault) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp16); ok {
		*a = MultistateRelinquishDefault(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MultistateRelinquishDefault) String() string {
	return zcl.Sprintf("%v", zcl.Zbmp16(a))
}

type MultistateText zcl.Zcstring

const MultistateTextAttr zcl.AttrID = 14

func (MultistateText) ID() zcl.AttrID   { return MultistateTextAttr }
func (MultistateText) Readable() bool   { return true }
func (MultistateText) Writable() bool   { return true }
func (MultistateText) Reportable() bool { return false }
func (MultistateText) SceneIndex() int  { return -1 }

func (MultistateText) Name() string                  { return "Multistate Text" }
func (a *MultistateText) TypeID() zcl.TypeID         { return zcl.Zcstring(*a).ID() }
func (a *MultistateText) Value() zcl.Val             { return a }
func (a MultistateText) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *MultistateText) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = MultistateText(*nt)
	return br, err
}

func (a *MultistateText) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = MultistateText(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MultistateText) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

type NwkDecryptFailures zcl.Zu16

const NwkDecryptFailuresAttr zcl.AttrID = 277

func (NwkDecryptFailures) ID() zcl.AttrID   { return NwkDecryptFailuresAttr }
func (NwkDecryptFailures) Readable() bool   { return true }
func (NwkDecryptFailures) Writable() bool   { return false }
func (NwkDecryptFailures) Reportable() bool { return false }
func (NwkDecryptFailures) SceneIndex() int  { return -1 }

func (NwkDecryptFailures) Name() string                  { return "NWK Decrypt Failures" }
func (a *NwkDecryptFailures) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *NwkDecryptFailures) Value() zcl.Val             { return a }
func (a NwkDecryptFailures) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *NwkDecryptFailures) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NwkDecryptFailures(*nt)
	return br, err
}

func (a *NwkDecryptFailures) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = NwkDecryptFailures(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NwkDecryptFailures) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type NwkFcFailure zcl.Zu16

const NwkFcFailureAttr zcl.AttrID = 274

func (NwkFcFailure) ID() zcl.AttrID   { return NwkFcFailureAttr }
func (NwkFcFailure) Readable() bool   { return true }
func (NwkFcFailure) Writable() bool   { return false }
func (NwkFcFailure) Reportable() bool { return false }
func (NwkFcFailure) SceneIndex() int  { return -1 }

func (NwkFcFailure) Name() string                  { return "NWK FC Failure" }
func (a *NwkFcFailure) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *NwkFcFailure) Value() zcl.Val             { return a }
func (a NwkFcFailure) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *NwkFcFailure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NwkFcFailure(*nt)
	return br, err
}

func (a *NwkFcFailure) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = NwkFcFailure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NwkFcFailure) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type NeighborAdded zcl.Zu16

const NeighborAddedAttr zcl.AttrID = 269

func (NeighborAdded) ID() zcl.AttrID   { return NeighborAddedAttr }
func (NeighborAdded) Readable() bool   { return true }
func (NeighborAdded) Writable() bool   { return false }
func (NeighborAdded) Reportable() bool { return false }
func (NeighborAdded) SceneIndex() int  { return -1 }

func (NeighborAdded) Name() string                  { return "Neighbor Added" }
func (a *NeighborAdded) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *NeighborAdded) Value() zcl.Val             { return a }
func (a NeighborAdded) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *NeighborAdded) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NeighborAdded(*nt)
	return br, err
}

func (a *NeighborAdded) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = NeighborAdded(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NeighborAdded) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type NeighborRemoved zcl.Zu16

const NeighborRemovedAttr zcl.AttrID = 270

func (NeighborRemoved) ID() zcl.AttrID   { return NeighborRemovedAttr }
func (NeighborRemoved) Readable() bool   { return true }
func (NeighborRemoved) Writable() bool   { return false }
func (NeighborRemoved) Reportable() bool { return false }
func (NeighborRemoved) SceneIndex() int  { return -1 }

func (NeighborRemoved) Name() string                  { return "Neighbor Removed" }
func (a *NeighborRemoved) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *NeighborRemoved) Value() zcl.Val             { return a }
func (a NeighborRemoved) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *NeighborRemoved) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NeighborRemoved(*nt)
	return br, err
}

func (a *NeighborRemoved) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = NeighborRemoved(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NeighborRemoved) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type NeighborStale zcl.Zu16

const NeighborStaleAttr zcl.AttrID = 271

func (NeighborStale) ID() zcl.AttrID   { return NeighborStaleAttr }
func (NeighborStale) Readable() bool   { return true }
func (NeighborStale) Writable() bool   { return false }
func (NeighborStale) Reportable() bool { return false }
func (NeighborStale) SceneIndex() int  { return -1 }

func (NeighborStale) Name() string                  { return "Neighbor Stale" }
func (a *NeighborStale) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *NeighborStale) Value() zcl.Val             { return a }
func (a NeighborStale) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *NeighborStale) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NeighborStale(*nt)
	return br, err
}

func (a *NeighborStale) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = NeighborStale(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NeighborStale) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type NeighborsInfo zcl.Zset

func (NeighborsInfo) Name() string          { return "Neighbors Info" }
func (a *NeighborsInfo) TypeID() zcl.TypeID { return zcl.Zset(*a).ID() }
func (a *NeighborsInfo) Value() zcl.Val     { return a }
func (a NeighborsInfo) MarshalZcl() ([]byte, error) {
	return nil, zcl.Errorf("not implemented")
}

func (a *NeighborsInfo) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zset)
	br, err := nt.UnmarshalZcl(b)
	*a = NeighborsInfo(*nt)
	return br, err
}

func (a *NeighborsInfo) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zset); ok {
		*a = NeighborsInfo(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NeighborsInfo) String() string {
	return zcl.Sprintf("%v", zcl.Zset(a))
}

// NumberRssiMeasurements is the number of measurements to use to generate one location estimate
type NumberRssiMeasurements zcl.Zu8

const NumberRssiMeasurementsAttr zcl.AttrID = 23

func (NumberRssiMeasurements) ID() zcl.AttrID   { return NumberRssiMeasurementsAttr }
func (NumberRssiMeasurements) Readable() bool   { return true }
func (NumberRssiMeasurements) Writable() bool   { return true }
func (NumberRssiMeasurements) Reportable() bool { return false }
func (NumberRssiMeasurements) SceneIndex() int  { return -1 }

func (NumberRssiMeasurements) Name() string                  { return "Number RSSI Measurements" }
func (a *NumberRssiMeasurements) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *NumberRssiMeasurements) Value() zcl.Val             { return a }
func (a NumberRssiMeasurements) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *NumberRssiMeasurements) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberRssiMeasurements(*nt)
	return br, err
}

func (a *NumberRssiMeasurements) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = NumberRssiMeasurements(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NumberRssiMeasurements) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type NumberResponses zcl.Zu8

func (NumberResponses) Name() string                  { return "Number Responses" }
func (a *NumberResponses) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *NumberResponses) Value() zcl.Val             { return a }
func (a NumberResponses) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *NumberResponses) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberResponses(*nt)
	return br, err
}

func (a *NumberResponses) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = NumberResponses(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NumberResponses) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type NumberOfDevices zcl.Zu8

const NumberOfDevicesAttr zcl.AttrID = 4

func (NumberOfDevices) ID() zcl.AttrID   { return NumberOfDevicesAttr }
func (NumberOfDevices) Readable() bool   { return true }
func (NumberOfDevices) Writable() bool   { return false }
func (NumberOfDevices) Reportable() bool { return false }
func (NumberOfDevices) SceneIndex() int  { return -1 }

func (NumberOfDevices) Name() string                  { return "Number of Devices" }
func (a *NumberOfDevices) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *NumberOfDevices) Value() zcl.Val             { return a }
func (a NumberOfDevices) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *NumberOfDevices) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfDevices(*nt)
	return br, err
}

func (a *NumberOfDevices) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = NumberOfDevices(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NumberOfDevices) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type NumberOfResets zcl.Zu16

const NumberOfResetsAttr zcl.AttrID = 0

func (NumberOfResets) ID() zcl.AttrID   { return NumberOfResetsAttr }
func (NumberOfResets) Readable() bool   { return true }
func (NumberOfResets) Writable() bool   { return false }
func (NumberOfResets) Reportable() bool { return false }
func (NumberOfResets) SceneIndex() int  { return -1 }

func (NumberOfResets) Name() string                  { return "Number of Resets" }
func (a *NumberOfResets) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *NumberOfResets) Value() zcl.Val             { return a }
func (a NumberOfResets) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *NumberOfResets) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfResets(*nt)
	return br, err
}

func (a *NumberOfResets) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = NumberOfResets(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NumberOfResets) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type OffTransitionTime zcl.Zu16

const OffTransitionTimeAttr zcl.AttrID = 19

func (OffTransitionTime) ID() zcl.AttrID   { return OffTransitionTimeAttr }
func (OffTransitionTime) Readable() bool   { return true }
func (OffTransitionTime) Writable() bool   { return true }
func (OffTransitionTime) Reportable() bool { return false }
func (OffTransitionTime) SceneIndex() int  { return -1 }

func (OffTransitionTime) Name() string                  { return "Off Transition Time" }
func (a *OffTransitionTime) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *OffTransitionTime) Value() zcl.Val             { return a }
func (a OffTransitionTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *OffTransitionTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = OffTransitionTime(*nt)
	return br, err
}

func (a *OffTransitionTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = OffTransitionTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OffTransitionTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

type OffWaitTime zcl.Zu16

const OffWaitTimeAttr zcl.AttrID = 16386

func (OffWaitTime) ID() zcl.AttrID   { return OffWaitTimeAttr }
func (OffWaitTime) Readable() bool   { return true }
func (OffWaitTime) Writable() bool   { return false }
func (OffWaitTime) Reportable() bool { return false }
func (OffWaitTime) SceneIndex() int  { return -1 }

func (OffWaitTime) Name() string                  { return "Off Wait Time" }
func (a *OffWaitTime) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *OffWaitTime) Value() zcl.Val             { return a }
func (a OffWaitTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *OffWaitTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = OffWaitTime(*nt)
	return br, err
}

func (a *OffWaitTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = OffWaitTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OffWaitTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

// OnLevel determines the value that the CurrentLevel attribute is set to when the OnOff attribute of an On/Off cluster on the
// same endpoint is set to On. If the OnLevel attribute is not implemented, or is set to 0xff, it has no effect.
type OnLevel zcl.Zu8

const OnLevelAttr zcl.AttrID = 17

func (OnLevel) ID() zcl.AttrID   { return OnLevelAttr }
func (OnLevel) Readable() bool   { return true }
func (OnLevel) Writable() bool   { return true }
func (OnLevel) Reportable() bool { return false }
func (OnLevel) SceneIndex() int  { return -1 }

func (OnLevel) Name() string                  { return "On Level" }
func (a *OnLevel) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *OnLevel) Value() zcl.Val             { return a }
func (a OnLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *OnLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = OnLevel(*nt)
	return br, err
}

func (a *OnLevel) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = OnLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OnLevel) String() string {
	return zcl.Percent.Format(float64(a) / 2.54)
}

type OnOff zcl.Zbool

const OnOffAttr zcl.AttrID = 0

func (OnOff) ID() zcl.AttrID   { return OnOffAttr }
func (OnOff) Readable() bool   { return true }
func (OnOff) Writable() bool   { return false }
func (OnOff) Reportable() bool { return true }
func (OnOff) SceneIndex() int  { return 1 }

func (OnOff) Name() string                  { return "On Off" }
func (a *OnOff) TypeID() zcl.TypeID         { return zcl.Zbool(*a).ID() }
func (a *OnOff) Value() zcl.Val             { return a }
func (a OnOff) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *OnOff) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = OnOff(*nt)
	return br, err
}

func (a *OnOff) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbool); ok {
		*a = OnOff(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OnOff) String() string {
	switch a {
	case 0x00:
		return "Off"
	case 0x01:
		return "On"
	}
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

func (a OnOff) IsOff() bool { return a == 0x00 }
func (a OnOff) IsOn() bool  { return a == 0x01 }
func (a *OnOff) SetOff()    { *a = 0x00 }
func (a *OnOff) SetOn()     { *a = 0x01 }

func (OnOff) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Off"},
		{Value: 0x01, Name: "On"},
	}
}

type OnTime zcl.Zu16

const OnTimeAttr zcl.AttrID = 16385

func (OnTime) ID() zcl.AttrID   { return OnTimeAttr }
func (OnTime) Readable() bool   { return true }
func (OnTime) Writable() bool   { return false }
func (OnTime) Reportable() bool { return false }
func (OnTime) SceneIndex() int  { return -1 }

func (OnTime) Name() string                  { return "On Time" }
func (a *OnTime) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *OnTime) Value() zcl.Val             { return a }
func (a OnTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *OnTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = OnTime(*nt)
	return br, err
}

func (a *OnTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = OnTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OnTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

type OnTransitionTime zcl.Zu16

const OnTransitionTimeAttr zcl.AttrID = 18

func (OnTransitionTime) ID() zcl.AttrID   { return OnTransitionTimeAttr }
func (OnTransitionTime) Readable() bool   { return true }
func (OnTransitionTime) Writable() bool   { return true }
func (OnTransitionTime) Reportable() bool { return false }
func (OnTransitionTime) SceneIndex() int  { return -1 }

func (OnTransitionTime) Name() string                  { return "On Transition Time" }
func (a *OnTransitionTime) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *OnTransitionTime) Value() zcl.Val             { return a }
func (a OnTransitionTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *OnTransitionTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = OnTransitionTime(*nt)
	return br, err
}

func (a *OnTransitionTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = OnTransitionTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OnTransitionTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

// OnOffTransistionTime represents the time taken to move to or from the target level when On of Off commands are received
// by an On/Off cluster on the same endpoint.
// The actual time taken should be as close to OnOffTransitionTime as the device is able.
type OnOffTransistionTime zcl.Zu16

const OnOffTransistionTimeAttr zcl.AttrID = 16

func (OnOffTransistionTime) ID() zcl.AttrID   { return OnOffTransistionTimeAttr }
func (OnOffTransistionTime) Readable() bool   { return true }
func (OnOffTransistionTime) Writable() bool   { return true }
func (OnOffTransistionTime) Reportable() bool { return false }
func (OnOffTransistionTime) SceneIndex() int  { return -1 }

func (OnOffTransistionTime) Name() string                  { return "On/Off Transistion Time" }
func (a *OnOffTransistionTime) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *OnOffTransistionTime) Value() zcl.Val             { return a }
func (a OnOffTransistionTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *OnOffTransistionTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = OnOffTransistionTime(*nt)
	return br, err
}

func (a *OnOffTransistionTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = OnOffTransistionTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OnOffTransistionTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

type OnOffControl zcl.Zbmp8

func (OnOffControl) Name() string                  { return "On/off control" }
func (a *OnOffControl) TypeID() zcl.TypeID         { return zcl.Zbmp8(*a).ID() }
func (a *OnOffControl) Value() zcl.Val             { return a }
func (a OnOffControl) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *OnOffControl) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = OnOffControl(*nt)
	return br, err
}

func (a *OnOffControl) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = OnOffControl(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OnOffControl) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Accept only when on")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a OnOffControl) IsAcceptOnlyWhenOn() bool { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a *OnOffControl) SetAcceptOnlyWhenOn(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}

func (OnOffControl) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Accept only when on"},
	}
}

// OverTempTotalDwell Total time the device has spent above the tmperature specified by High Temp Threshold
type OverTempTotalDwell zcl.Zu16

const OverTempTotalDwellAttr zcl.AttrID = 3

func (OverTempTotalDwell) ID() zcl.AttrID   { return OverTempTotalDwellAttr }
func (OverTempTotalDwell) Readable() bool   { return false }
func (OverTempTotalDwell) Writable() bool   { return false }
func (OverTempTotalDwell) Reportable() bool { return false }
func (OverTempTotalDwell) SceneIndex() int  { return -1 }

func (OverTempTotalDwell) Name() string                  { return "Over Temp Total Dwell" }
func (a *OverTempTotalDwell) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *OverTempTotalDwell) Value() zcl.Val             { return a }
func (a OverTempTotalDwell) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *OverTempTotalDwell) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = OverTempTotalDwell(*nt)
	return br, err
}

func (a *OverTempTotalDwell) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = OverTempTotalDwell(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OverTempTotalDwell) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type PacketBufferAllocFailures zcl.Zu16

const PacketBufferAllocFailuresAttr zcl.AttrID = 279

func (PacketBufferAllocFailures) ID() zcl.AttrID   { return PacketBufferAllocFailuresAttr }
func (PacketBufferAllocFailures) Readable() bool   { return true }
func (PacketBufferAllocFailures) Writable() bool   { return false }
func (PacketBufferAllocFailures) Reportable() bool { return false }
func (PacketBufferAllocFailures) SceneIndex() int  { return -1 }

func (PacketBufferAllocFailures) Name() string                  { return "Packet Buffer Alloc Failures" }
func (a *PacketBufferAllocFailures) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *PacketBufferAllocFailures) Value() zcl.Val             { return a }
func (a PacketBufferAllocFailures) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *PacketBufferAllocFailures) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PacketBufferAllocFailures(*nt)
	return br, err
}

func (a *PacketBufferAllocFailures) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = PacketBufferAllocFailures(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PacketBufferAllocFailures) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type PacketValidateDropcount zcl.Zu16

const PacketValidateDropcountAttr zcl.AttrID = 282

func (PacketValidateDropcount) ID() zcl.AttrID   { return PacketValidateDropcountAttr }
func (PacketValidateDropcount) Readable() bool   { return true }
func (PacketValidateDropcount) Writable() bool   { return false }
func (PacketValidateDropcount) Reportable() bool { return false }
func (PacketValidateDropcount) SceneIndex() int  { return -1 }

func (PacketValidateDropcount) Name() string                  { return "Packet Validate Dropcount" }
func (a *PacketValidateDropcount) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *PacketValidateDropcount) Value() zcl.Val             { return a }
func (a PacketValidateDropcount) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *PacketValidateDropcount) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PacketValidateDropcount(*nt)
	return br, err
}

func (a *PacketValidateDropcount) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = PacketValidateDropcount(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PacketValidateDropcount) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// PathLossExponent is the rate at which the signal power decays with increasing distance
type PathLossExponent zcl.Zu16

const PathLossExponentAttr zcl.AttrID = 20

func (PathLossExponent) ID() zcl.AttrID   { return PathLossExponentAttr }
func (PathLossExponent) Readable() bool   { return true }
func (PathLossExponent) Writable() bool   { return true }
func (PathLossExponent) Reportable() bool { return false }
func (PathLossExponent) SceneIndex() int  { return -1 }

func (PathLossExponent) Name() string                  { return "Path loss Exponent" }
func (a *PathLossExponent) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *PathLossExponent) Value() zcl.Val             { return a }
func (a PathLossExponent) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *PathLossExponent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PathLossExponent(*nt)
	return br, err
}

func (a *PathLossExponent) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = PathLossExponent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PathLossExponent) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type PersistensMemoryWrites zcl.Zu16

const PersistensMemoryWritesAttr zcl.AttrID = 1

func (PersistensMemoryWrites) ID() zcl.AttrID   { return PersistensMemoryWritesAttr }
func (PersistensMemoryWrites) Readable() bool   { return true }
func (PersistensMemoryWrites) Writable() bool   { return false }
func (PersistensMemoryWrites) Reportable() bool { return false }
func (PersistensMemoryWrites) SceneIndex() int  { return -1 }

func (PersistensMemoryWrites) Name() string                  { return "Persistens Memory Writes" }
func (a *PersistensMemoryWrites) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *PersistensMemoryWrites) Value() zcl.Val             { return a }
func (a PersistensMemoryWrites) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *PersistensMemoryWrites) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PersistensMemoryWrites(*nt)
	return br, err
}

func (a *PersistensMemoryWrites) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = PersistensMemoryWrites(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PersistensMemoryWrites) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type PhyToMacQueueLimitReached zcl.Zu16

const PhyToMacQueueLimitReachedAttr zcl.AttrID = 281

func (PhyToMacQueueLimitReached) ID() zcl.AttrID   { return PhyToMacQueueLimitReachedAttr }
func (PhyToMacQueueLimitReached) Readable() bool   { return true }
func (PhyToMacQueueLimitReached) Writable() bool   { return false }
func (PhyToMacQueueLimitReached) Reportable() bool { return false }
func (PhyToMacQueueLimitReached) SceneIndex() int  { return -1 }

func (PhyToMacQueueLimitReached) Name() string                  { return "Phy to MAC queue limit reached" }
func (a *PhyToMacQueueLimitReached) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *PhyToMacQueueLimitReached) Value() zcl.Val             { return a }
func (a PhyToMacQueueLimitReached) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *PhyToMacQueueLimitReached) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PhyToMacQueueLimitReached(*nt)
	return br, err
}

func (a *PhyToMacQueueLimitReached) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = PhyToMacQueueLimitReached(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PhyToMacQueueLimitReached) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type PhysicalEnvironment zcl.Zenum8

const PhysicalEnvironmentAttr zcl.AttrID = 17

func (PhysicalEnvironment) ID() zcl.AttrID   { return PhysicalEnvironmentAttr }
func (PhysicalEnvironment) Readable() bool   { return true }
func (PhysicalEnvironment) Writable() bool   { return true }
func (PhysicalEnvironment) Reportable() bool { return false }
func (PhysicalEnvironment) SceneIndex() int  { return -1 }

func (PhysicalEnvironment) Name() string                  { return "Physical Environment" }
func (a *PhysicalEnvironment) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *PhysicalEnvironment) Value() zcl.Val             { return a }
func (a PhysicalEnvironment) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *PhysicalEnvironment) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = PhysicalEnvironment(*nt)
	return br, err
}

func (a *PhysicalEnvironment) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = PhysicalEnvironment(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PhysicalEnvironment) String() string {
	switch a {
	case 0x00:
		return "Unspecified Environment"
	case 0x01:
		return "Atrium"
	case 0x02:
		return "Bar"
	case 0x03:
		return "Courtyard"
	case 0x04:
		return "Bathroom"
	case 0x05:
		return "Bedroom"
	case 0x06:
		return "Billiard Room"
	case 0x07:
		return "Utility Room"
	case 0x08:
		return "Cellar"
	case 0x09:
		return "Storage Closet"
	case 0x0a:
		return "Theater"
	case 0x0b:
		return "Office"
	case 0x0c:
		return "Deck"
	case 0x0d:
		return "Den"
	case 0x0e:
		return "Dining Room"
	case 0x0f:
		return "Electrical Room"
	case 0x10:
		return "Elevator"
	case 0x11:
		return "Entry"
	case 0x12:
		return "Family Room"
	case 0x13:
		return "Main Floor"
	case 0x14:
		return "Upstairs"
	case 0x15:
		return "Downstairs"
	case 0x16:
		return "Basement/Lower Level"
	case 0x17:
		return "Gallery"
	case 0x18:
		return "Game Room"
	case 0x19:
		return "Garage"
	case 0x1a:
		return "Gym"
	case 0x1b:
		return "Hallway"
	case 0x1c:
		return "House"
	case 0x1d:
		return "Kitchen"
	case 0x1e:
		return "Laundry Room"
	case 0x1f:
		return "Library"
	case 0x20:
		return "Master Bedroom"
	case 0x21:
		return "Mud Room (small room for coats and boots)"
	case 0x22:
		return "Nursery"
	case 0x23:
		return "Pantry"
	case 0x24:
		return "Office 2"
	case 0x25:
		return "Outside"
	case 0x26:
		return "Pool"
	case 0x27:
		return "Porch"
	case 0x28:
		return "Sewing Room"
	case 0x29:
		return "Sitting Room"
	case 0x2a:
		return "Stairway"
	case 0x2b:
		return "Yard"
	case 0x2c:
		return "Attic"
	case 0x2d:
		return "Hot Tub"
	case 0x2e:
		return "Living Room"
	case 0x2f:
		return "Sauna"
	case 0x30:
		return "Shop/Workshop"
	case 0x31:
		return "Guest Bedroom"
	case 0x32:
		return "Guest Bath"
	case 0x33:
		return "Powder Room (1/2 bath)"
	case 0x34:
		return "Back Yard"
	case 0x35:
		return "Front Yard"
	case 0x36:
		return "Patio"
	case 0x37:
		return "Driveway"
	case 0x38:
		return "Sun Room"
	case 0x39:
		return "Living Room 2"
	case 0x3a:
		return "Spa"
	case 0x3b:
		return "Whirlpool"
	case 0x3c:
		return "Shed"
	case 0x3d:
		return "Equipment Storage"
	case 0x3e:
		return "Hobby/Craft Room"
	case 0x3f:
		return "Fountain"
	case 0x40:
		return "Pond"
	case 0x41:
		return "Reception Room"
	case 0x42:
		return "Breakfast Room"
	case 0x43:
		return "Nook"
	case 0x44:
		return "Garden"
	case 0x45:
		return "Balcony"
	case 0x46:
		return "Panic Room"
	case 0x47:
		return "Terrace"
	case 0x48:
		return "Roof"
	case 0x49:
		return "Toilet"
	case 0x4a:
		return "Toilet Main"
	case 0x4b:
		return "Outside Toilet"
	case 0x4c:
		return "Shower room"
	case 0x4d:
		return "Study"
	case 0x4e:
		return "Front Garden"
	case 0x4f:
		return "Back Garden"
	case 0x50:
		return "Kettle"
	case 0x51:
		return "Television"
	case 0x52:
		return "Stove"
	case 0x53:
		return "Microwave"
	case 0x54:
		return "Toaster"
	case 0x55:
		return "Vacuum"
	case 0x56:
		return "Appliance"
	case 0x57:
		return "Front Door"
	case 0x58:
		return "Back Door"
	case 0x59:
		return "Fridge Door"
	case 0x60:
		return "Medication Cabinet Door"
	case 0x61:
		return "Wardrobe Door"
	case 0x62:
		return "Front Cupboard Door"
	case 0x63:
		return "Other Door"
	case 0x64:
		return "Waiting Room"
	case 0x65:
		return "Triage Room"
	case 0x66:
		return "Doctor’s Office"
	case 0x67:
		return "Patient’s Private Room"
	case 0x68:
		return "Consultation Room"
	case 0x69:
		return "Nurse Station"
	case 0x6a:
		return "Ward"
	case 0x6b:
		return "Corridor"
	case 0x6c:
		return "Operating Theatre"
	case 0x6d:
		return "Dental Surgery Room"
	case 0x6e:
		return "Medical Imaging Room"
	case 0x6f:
		return "Decontamination Room"
	case 0xFF:
		return "Unknown Environment"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a PhysicalEnvironment) IsUnspecifiedEnvironment() bool           { return a == 0x00 }
func (a PhysicalEnvironment) IsAtrium() bool                           { return a == 0x01 }
func (a PhysicalEnvironment) IsBar() bool                              { return a == 0x02 }
func (a PhysicalEnvironment) IsCourtyard() bool                        { return a == 0x03 }
func (a PhysicalEnvironment) IsBathroom() bool                         { return a == 0x04 }
func (a PhysicalEnvironment) IsBedroom() bool                          { return a == 0x05 }
func (a PhysicalEnvironment) IsBilliardRoom() bool                     { return a == 0x06 }
func (a PhysicalEnvironment) IsUtilityRoom() bool                      { return a == 0x07 }
func (a PhysicalEnvironment) IsCellar() bool                           { return a == 0x08 }
func (a PhysicalEnvironment) IsStorageCloset() bool                    { return a == 0x09 }
func (a PhysicalEnvironment) IsTheater() bool                          { return a == 0x0a }
func (a PhysicalEnvironment) IsOffice() bool                           { return a == 0x0b }
func (a PhysicalEnvironment) IsDeck() bool                             { return a == 0x0c }
func (a PhysicalEnvironment) IsDen() bool                              { return a == 0x0d }
func (a PhysicalEnvironment) IsDiningRoom() bool                       { return a == 0x0e }
func (a PhysicalEnvironment) IsElectricalRoom() bool                   { return a == 0x0f }
func (a PhysicalEnvironment) IsElevator() bool                         { return a == 0x10 }
func (a PhysicalEnvironment) IsEntry() bool                            { return a == 0x11 }
func (a PhysicalEnvironment) IsFamilyRoom() bool                       { return a == 0x12 }
func (a PhysicalEnvironment) IsMainFloor() bool                        { return a == 0x13 }
func (a PhysicalEnvironment) IsUpstairs() bool                         { return a == 0x14 }
func (a PhysicalEnvironment) IsDownstairs() bool                       { return a == 0x15 }
func (a PhysicalEnvironment) IsBasementLowerLevel() bool               { return a == 0x16 }
func (a PhysicalEnvironment) IsGallery() bool                          { return a == 0x17 }
func (a PhysicalEnvironment) IsGameRoom() bool                         { return a == 0x18 }
func (a PhysicalEnvironment) IsGarage() bool                           { return a == 0x19 }
func (a PhysicalEnvironment) IsGym() bool                              { return a == 0x1a }
func (a PhysicalEnvironment) IsHallway() bool                          { return a == 0x1b }
func (a PhysicalEnvironment) IsHouse() bool                            { return a == 0x1c }
func (a PhysicalEnvironment) IsKitchen() bool                          { return a == 0x1d }
func (a PhysicalEnvironment) IsLaundryRoom() bool                      { return a == 0x1e }
func (a PhysicalEnvironment) IsLibrary() bool                          { return a == 0x1f }
func (a PhysicalEnvironment) IsMasterBedroom() bool                    { return a == 0x20 }
func (a PhysicalEnvironment) IsMudRoomSmallRoomForCoatsAndBoots() bool { return a == 0x21 }
func (a PhysicalEnvironment) IsNursery() bool                          { return a == 0x22 }
func (a PhysicalEnvironment) IsPantry() bool                           { return a == 0x23 }
func (a PhysicalEnvironment) IsOffice2() bool                          { return a == 0x24 }
func (a PhysicalEnvironment) IsOutside() bool                          { return a == 0x25 }
func (a PhysicalEnvironment) IsPool() bool                             { return a == 0x26 }
func (a PhysicalEnvironment) IsPorch() bool                            { return a == 0x27 }
func (a PhysicalEnvironment) IsSewingRoom() bool                       { return a == 0x28 }
func (a PhysicalEnvironment) IsSittingRoom() bool                      { return a == 0x29 }
func (a PhysicalEnvironment) IsStairway() bool                         { return a == 0x2a }
func (a PhysicalEnvironment) IsYard() bool                             { return a == 0x2b }
func (a PhysicalEnvironment) IsAttic() bool                            { return a == 0x2c }
func (a PhysicalEnvironment) IsHotTub() bool                           { return a == 0x2d }
func (a PhysicalEnvironment) IsLivingRoom() bool                       { return a == 0x2e }
func (a PhysicalEnvironment) IsSauna() bool                            { return a == 0x2f }
func (a PhysicalEnvironment) IsShopWorkshop() bool                     { return a == 0x30 }
func (a PhysicalEnvironment) IsGuestBedroom() bool                     { return a == 0x31 }
func (a PhysicalEnvironment) IsGuestBath() bool                        { return a == 0x32 }
func (a PhysicalEnvironment) IsPowderRoom12Bath() bool                 { return a == 0x33 }
func (a PhysicalEnvironment) IsBackYard() bool                         { return a == 0x34 }
func (a PhysicalEnvironment) IsFrontYard() bool                        { return a == 0x35 }
func (a PhysicalEnvironment) IsPatio() bool                            { return a == 0x36 }
func (a PhysicalEnvironment) IsDriveway() bool                         { return a == 0x37 }
func (a PhysicalEnvironment) IsSunRoom() bool                          { return a == 0x38 }
func (a PhysicalEnvironment) IsLivingRoom2() bool                      { return a == 0x39 }
func (a PhysicalEnvironment) IsSpa() bool                              { return a == 0x3a }
func (a PhysicalEnvironment) IsWhirlpool() bool                        { return a == 0x3b }
func (a PhysicalEnvironment) IsShed() bool                             { return a == 0x3c }
func (a PhysicalEnvironment) IsEquipmentStorage() bool                 { return a == 0x3d }
func (a PhysicalEnvironment) IsHobbyCraftRoom() bool                   { return a == 0x3e }
func (a PhysicalEnvironment) IsFountain() bool                         { return a == 0x3f }
func (a PhysicalEnvironment) IsPond() bool                             { return a == 0x40 }
func (a PhysicalEnvironment) IsReceptionRoom() bool                    { return a == 0x41 }
func (a PhysicalEnvironment) IsBreakfastRoom() bool                    { return a == 0x42 }
func (a PhysicalEnvironment) IsNook() bool                             { return a == 0x43 }
func (a PhysicalEnvironment) IsGarden() bool                           { return a == 0x44 }
func (a PhysicalEnvironment) IsBalcony() bool                          { return a == 0x45 }
func (a PhysicalEnvironment) IsPanicRoom() bool                        { return a == 0x46 }
func (a PhysicalEnvironment) IsTerrace() bool                          { return a == 0x47 }
func (a PhysicalEnvironment) IsRoof() bool                             { return a == 0x48 }
func (a PhysicalEnvironment) IsToilet() bool                           { return a == 0x49 }
func (a PhysicalEnvironment) IsToiletMain() bool                       { return a == 0x4a }
func (a PhysicalEnvironment) IsOutsideToilet() bool                    { return a == 0x4b }
func (a PhysicalEnvironment) IsShowerRoom() bool                       { return a == 0x4c }
func (a PhysicalEnvironment) IsStudy() bool                            { return a == 0x4d }
func (a PhysicalEnvironment) IsFrontGarden() bool                      { return a == 0x4e }
func (a PhysicalEnvironment) IsBackGarden() bool                       { return a == 0x4f }
func (a PhysicalEnvironment) IsKettle() bool                           { return a == 0x50 }
func (a PhysicalEnvironment) IsTelevision() bool                       { return a == 0x51 }
func (a PhysicalEnvironment) IsStove() bool                            { return a == 0x52 }
func (a PhysicalEnvironment) IsMicrowave() bool                        { return a == 0x53 }
func (a PhysicalEnvironment) IsToaster() bool                          { return a == 0x54 }
func (a PhysicalEnvironment) IsVacuum() bool                           { return a == 0x55 }
func (a PhysicalEnvironment) IsAppliance() bool                        { return a == 0x56 }
func (a PhysicalEnvironment) IsFrontDoor() bool                        { return a == 0x57 }
func (a PhysicalEnvironment) IsBackDoor() bool                         { return a == 0x58 }
func (a PhysicalEnvironment) IsFridgeDoor() bool                       { return a == 0x59 }
func (a PhysicalEnvironment) IsMedicationCabinetDoor() bool            { return a == 0x60 }
func (a PhysicalEnvironment) IsWardrobeDoor() bool                     { return a == 0x61 }
func (a PhysicalEnvironment) IsFrontCupboardDoor() bool                { return a == 0x62 }
func (a PhysicalEnvironment) IsOtherDoor() bool                        { return a == 0x63 }
func (a PhysicalEnvironment) IsWaitingRoom() bool                      { return a == 0x64 }
func (a PhysicalEnvironment) IsTriageRoom() bool                       { return a == 0x65 }
func (a PhysicalEnvironment) IsDoctorSOffice() bool                    { return a == 0x66 }
func (a PhysicalEnvironment) IsPatientSPrivateRoom() bool              { return a == 0x67 }
func (a PhysicalEnvironment) IsConsultationRoom() bool                 { return a == 0x68 }
func (a PhysicalEnvironment) IsNurseStation() bool                     { return a == 0x69 }
func (a PhysicalEnvironment) IsWard() bool                             { return a == 0x6a }
func (a PhysicalEnvironment) IsCorridor() bool                         { return a == 0x6b }
func (a PhysicalEnvironment) IsOperatingTheatre() bool                 { return a == 0x6c }
func (a PhysicalEnvironment) IsDentalSurgeryRoom() bool                { return a == 0x6d }
func (a PhysicalEnvironment) IsMedicalImagingRoom() bool               { return a == 0x6e }
func (a PhysicalEnvironment) IsDecontaminationRoom() bool              { return a == 0x6f }
func (a PhysicalEnvironment) IsUnknownEnvironment() bool               { return a == 0xFF }
func (a *PhysicalEnvironment) SetUnspecifiedEnvironment()              { *a = 0x00 }
func (a *PhysicalEnvironment) SetAtrium()                              { *a = 0x01 }
func (a *PhysicalEnvironment) SetBar()                                 { *a = 0x02 }
func (a *PhysicalEnvironment) SetCourtyard()                           { *a = 0x03 }
func (a *PhysicalEnvironment) SetBathroom()                            { *a = 0x04 }
func (a *PhysicalEnvironment) SetBedroom()                             { *a = 0x05 }
func (a *PhysicalEnvironment) SetBilliardRoom()                        { *a = 0x06 }
func (a *PhysicalEnvironment) SetUtilityRoom()                         { *a = 0x07 }
func (a *PhysicalEnvironment) SetCellar()                              { *a = 0x08 }
func (a *PhysicalEnvironment) SetStorageCloset()                       { *a = 0x09 }
func (a *PhysicalEnvironment) SetTheater()                             { *a = 0x0a }
func (a *PhysicalEnvironment) SetOffice()                              { *a = 0x0b }
func (a *PhysicalEnvironment) SetDeck()                                { *a = 0x0c }
func (a *PhysicalEnvironment) SetDen()                                 { *a = 0x0d }
func (a *PhysicalEnvironment) SetDiningRoom()                          { *a = 0x0e }
func (a *PhysicalEnvironment) SetElectricalRoom()                      { *a = 0x0f }
func (a *PhysicalEnvironment) SetElevator()                            { *a = 0x10 }
func (a *PhysicalEnvironment) SetEntry()                               { *a = 0x11 }
func (a *PhysicalEnvironment) SetFamilyRoom()                          { *a = 0x12 }
func (a *PhysicalEnvironment) SetMainFloor()                           { *a = 0x13 }
func (a *PhysicalEnvironment) SetUpstairs()                            { *a = 0x14 }
func (a *PhysicalEnvironment) SetDownstairs()                          { *a = 0x15 }
func (a *PhysicalEnvironment) SetBasementLowerLevel()                  { *a = 0x16 }
func (a *PhysicalEnvironment) SetGallery()                             { *a = 0x17 }
func (a *PhysicalEnvironment) SetGameRoom()                            { *a = 0x18 }
func (a *PhysicalEnvironment) SetGarage()                              { *a = 0x19 }
func (a *PhysicalEnvironment) SetGym()                                 { *a = 0x1a }
func (a *PhysicalEnvironment) SetHallway()                             { *a = 0x1b }
func (a *PhysicalEnvironment) SetHouse()                               { *a = 0x1c }
func (a *PhysicalEnvironment) SetKitchen()                             { *a = 0x1d }
func (a *PhysicalEnvironment) SetLaundryRoom()                         { *a = 0x1e }
func (a *PhysicalEnvironment) SetLibrary()                             { *a = 0x1f }
func (a *PhysicalEnvironment) SetMasterBedroom()                       { *a = 0x20 }
func (a *PhysicalEnvironment) SetMudRoomSmallRoomForCoatsAndBoots()    { *a = 0x21 }
func (a *PhysicalEnvironment) SetNursery()                             { *a = 0x22 }
func (a *PhysicalEnvironment) SetPantry()                              { *a = 0x23 }
func (a *PhysicalEnvironment) SetOffice2()                             { *a = 0x24 }
func (a *PhysicalEnvironment) SetOutside()                             { *a = 0x25 }
func (a *PhysicalEnvironment) SetPool()                                { *a = 0x26 }
func (a *PhysicalEnvironment) SetPorch()                               { *a = 0x27 }
func (a *PhysicalEnvironment) SetSewingRoom()                          { *a = 0x28 }
func (a *PhysicalEnvironment) SetSittingRoom()                         { *a = 0x29 }
func (a *PhysicalEnvironment) SetStairway()                            { *a = 0x2a }
func (a *PhysicalEnvironment) SetYard()                                { *a = 0x2b }
func (a *PhysicalEnvironment) SetAttic()                               { *a = 0x2c }
func (a *PhysicalEnvironment) SetHotTub()                              { *a = 0x2d }
func (a *PhysicalEnvironment) SetLivingRoom()                          { *a = 0x2e }
func (a *PhysicalEnvironment) SetSauna()                               { *a = 0x2f }
func (a *PhysicalEnvironment) SetShopWorkshop()                        { *a = 0x30 }
func (a *PhysicalEnvironment) SetGuestBedroom()                        { *a = 0x31 }
func (a *PhysicalEnvironment) SetGuestBath()                           { *a = 0x32 }
func (a *PhysicalEnvironment) SetPowderRoom12Bath()                    { *a = 0x33 }
func (a *PhysicalEnvironment) SetBackYard()                            { *a = 0x34 }
func (a *PhysicalEnvironment) SetFrontYard()                           { *a = 0x35 }
func (a *PhysicalEnvironment) SetPatio()                               { *a = 0x36 }
func (a *PhysicalEnvironment) SetDriveway()                            { *a = 0x37 }
func (a *PhysicalEnvironment) SetSunRoom()                             { *a = 0x38 }
func (a *PhysicalEnvironment) SetLivingRoom2()                         { *a = 0x39 }
func (a *PhysicalEnvironment) SetSpa()                                 { *a = 0x3a }
func (a *PhysicalEnvironment) SetWhirlpool()                           { *a = 0x3b }
func (a *PhysicalEnvironment) SetShed()                                { *a = 0x3c }
func (a *PhysicalEnvironment) SetEquipmentStorage()                    { *a = 0x3d }
func (a *PhysicalEnvironment) SetHobbyCraftRoom()                      { *a = 0x3e }
func (a *PhysicalEnvironment) SetFountain()                            { *a = 0x3f }
func (a *PhysicalEnvironment) SetPond()                                { *a = 0x40 }
func (a *PhysicalEnvironment) SetReceptionRoom()                       { *a = 0x41 }
func (a *PhysicalEnvironment) SetBreakfastRoom()                       { *a = 0x42 }
func (a *PhysicalEnvironment) SetNook()                                { *a = 0x43 }
func (a *PhysicalEnvironment) SetGarden()                              { *a = 0x44 }
func (a *PhysicalEnvironment) SetBalcony()                             { *a = 0x45 }
func (a *PhysicalEnvironment) SetPanicRoom()                           { *a = 0x46 }
func (a *PhysicalEnvironment) SetTerrace()                             { *a = 0x47 }
func (a *PhysicalEnvironment) SetRoof()                                { *a = 0x48 }
func (a *PhysicalEnvironment) SetToilet()                              { *a = 0x49 }
func (a *PhysicalEnvironment) SetToiletMain()                          { *a = 0x4a }
func (a *PhysicalEnvironment) SetOutsideToilet()                       { *a = 0x4b }
func (a *PhysicalEnvironment) SetShowerRoom()                          { *a = 0x4c }
func (a *PhysicalEnvironment) SetStudy()                               { *a = 0x4d }
func (a *PhysicalEnvironment) SetFrontGarden()                         { *a = 0x4e }
func (a *PhysicalEnvironment) SetBackGarden()                          { *a = 0x4f }
func (a *PhysicalEnvironment) SetKettle()                              { *a = 0x50 }
func (a *PhysicalEnvironment) SetTelevision()                          { *a = 0x51 }
func (a *PhysicalEnvironment) SetStove()                               { *a = 0x52 }
func (a *PhysicalEnvironment) SetMicrowave()                           { *a = 0x53 }
func (a *PhysicalEnvironment) SetToaster()                             { *a = 0x54 }
func (a *PhysicalEnvironment) SetVacuum()                              { *a = 0x55 }
func (a *PhysicalEnvironment) SetAppliance()                           { *a = 0x56 }
func (a *PhysicalEnvironment) SetFrontDoor()                           { *a = 0x57 }
func (a *PhysicalEnvironment) SetBackDoor()                            { *a = 0x58 }
func (a *PhysicalEnvironment) SetFridgeDoor()                          { *a = 0x59 }
func (a *PhysicalEnvironment) SetMedicationCabinetDoor()               { *a = 0x60 }
func (a *PhysicalEnvironment) SetWardrobeDoor()                        { *a = 0x61 }
func (a *PhysicalEnvironment) SetFrontCupboardDoor()                   { *a = 0x62 }
func (a *PhysicalEnvironment) SetOtherDoor()                           { *a = 0x63 }
func (a *PhysicalEnvironment) SetWaitingRoom()                         { *a = 0x64 }
func (a *PhysicalEnvironment) SetTriageRoom()                          { *a = 0x65 }
func (a *PhysicalEnvironment) SetDoctorSOffice()                       { *a = 0x66 }
func (a *PhysicalEnvironment) SetPatientSPrivateRoom()                 { *a = 0x67 }
func (a *PhysicalEnvironment) SetConsultationRoom()                    { *a = 0x68 }
func (a *PhysicalEnvironment) SetNurseStation()                        { *a = 0x69 }
func (a *PhysicalEnvironment) SetWard()                                { *a = 0x6a }
func (a *PhysicalEnvironment) SetCorridor()                            { *a = 0x6b }
func (a *PhysicalEnvironment) SetOperatingTheatre()                    { *a = 0x6c }
func (a *PhysicalEnvironment) SetDentalSurgeryRoom()                   { *a = 0x6d }
func (a *PhysicalEnvironment) SetMedicalImagingRoom()                  { *a = 0x6e }
func (a *PhysicalEnvironment) SetDecontaminationRoom()                 { *a = 0x6f }
func (a *PhysicalEnvironment) SetUnknownEnvironment()                  { *a = 0xFF }

func (PhysicalEnvironment) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Unspecified Environment"},
		{Value: 0x01, Name: "Atrium"},
		{Value: 0x02, Name: "Bar"},
		{Value: 0x03, Name: "Courtyard"},
		{Value: 0x04, Name: "Bathroom"},
		{Value: 0x05, Name: "Bedroom"},
		{Value: 0x06, Name: "Billiard Room"},
		{Value: 0x07, Name: "Utility Room"},
		{Value: 0x08, Name: "Cellar"},
		{Value: 0x09, Name: "Storage Closet"},
		{Value: 0x0a, Name: "Theater"},
		{Value: 0x0b, Name: "Office"},
		{Value: 0x0c, Name: "Deck"},
		{Value: 0x0d, Name: "Den"},
		{Value: 0x0e, Name: "Dining Room"},
		{Value: 0x0f, Name: "Electrical Room"},
		{Value: 0x10, Name: "Elevator"},
		{Value: 0x11, Name: "Entry"},
		{Value: 0x12, Name: "Family Room"},
		{Value: 0x13, Name: "Main Floor"},
		{Value: 0x14, Name: "Upstairs"},
		{Value: 0x15, Name: "Downstairs"},
		{Value: 0x16, Name: "Basement/Lower Level"},
		{Value: 0x17, Name: "Gallery"},
		{Value: 0x18, Name: "Game Room"},
		{Value: 0x19, Name: "Garage"},
		{Value: 0x1a, Name: "Gym"},
		{Value: 0x1b, Name: "Hallway"},
		{Value: 0x1c, Name: "House"},
		{Value: 0x1d, Name: "Kitchen"},
		{Value: 0x1e, Name: "Laundry Room"},
		{Value: 0x1f, Name: "Library"},
		{Value: 0x20, Name: "Master Bedroom"},
		{Value: 0x21, Name: "Mud Room (small room for coats and boots)"},
		{Value: 0x22, Name: "Nursery"},
		{Value: 0x23, Name: "Pantry"},
		{Value: 0x24, Name: "Office 2"},
		{Value: 0x25, Name: "Outside"},
		{Value: 0x26, Name: "Pool"},
		{Value: 0x27, Name: "Porch"},
		{Value: 0x28, Name: "Sewing Room"},
		{Value: 0x29, Name: "Sitting Room"},
		{Value: 0x2a, Name: "Stairway"},
		{Value: 0x2b, Name: "Yard"},
		{Value: 0x2c, Name: "Attic"},
		{Value: 0x2d, Name: "Hot Tub"},
		{Value: 0x2e, Name: "Living Room"},
		{Value: 0x2f, Name: "Sauna"},
		{Value: 0x30, Name: "Shop/Workshop"},
		{Value: 0x31, Name: "Guest Bedroom"},
		{Value: 0x32, Name: "Guest Bath"},
		{Value: 0x33, Name: "Powder Room (1/2 bath)"},
		{Value: 0x34, Name: "Back Yard"},
		{Value: 0x35, Name: "Front Yard"},
		{Value: 0x36, Name: "Patio"},
		{Value: 0x37, Name: "Driveway"},
		{Value: 0x38, Name: "Sun Room"},
		{Value: 0x39, Name: "Living Room 2"},
		{Value: 0x3a, Name: "Spa"},
		{Value: 0x3b, Name: "Whirlpool"},
		{Value: 0x3c, Name: "Shed"},
		{Value: 0x3d, Name: "Equipment Storage"},
		{Value: 0x3e, Name: "Hobby/Craft Room"},
		{Value: 0x3f, Name: "Fountain"},
		{Value: 0x40, Name: "Pond"},
		{Value: 0x41, Name: "Reception Room"},
		{Value: 0x42, Name: "Breakfast Room"},
		{Value: 0x43, Name: "Nook"},
		{Value: 0x44, Name: "Garden"},
		{Value: 0x45, Name: "Balcony"},
		{Value: 0x46, Name: "Panic Room"},
		{Value: 0x47, Name: "Terrace"},
		{Value: 0x48, Name: "Roof"},
		{Value: 0x49, Name: "Toilet"},
		{Value: 0x4a, Name: "Toilet Main"},
		{Value: 0x4b, Name: "Outside Toilet"},
		{Value: 0x4c, Name: "Shower room"},
		{Value: 0x4d, Name: "Study"},
		{Value: 0x4e, Name: "Front Garden"},
		{Value: 0x4f, Name: "Back Garden"},
		{Value: 0x50, Name: "Kettle"},
		{Value: 0x51, Name: "Television"},
		{Value: 0x52, Name: "Stove"},
		{Value: 0x53, Name: "Microwave"},
		{Value: 0x54, Name: "Toaster"},
		{Value: 0x55, Name: "Vacuum"},
		{Value: 0x56, Name: "Appliance"},
		{Value: 0x57, Name: "Front Door"},
		{Value: 0x58, Name: "Back Door"},
		{Value: 0x59, Name: "Fridge Door"},
		{Value: 0x60, Name: "Medication Cabinet Door"},
		{Value: 0x61, Name: "Wardrobe Door"},
		{Value: 0x62, Name: "Front Cupboard Door"},
		{Value: 0x63, Name: "Other Door"},
		{Value: 0x64, Name: "Waiting Room"},
		{Value: 0x65, Name: "Triage Room"},
		{Value: 0x66, Name: "Doctor’s Office"},
		{Value: 0x67, Name: "Patient’s Private Room"},
		{Value: 0x68, Name: "Consultation Room"},
		{Value: 0x69, Name: "Nurse Station"},
		{Value: 0x6a, Name: "Ward"},
		{Value: 0x6b, Name: "Corridor"},
		{Value: 0x6c, Name: "Operating Theatre"},
		{Value: 0x6d, Name: "Dental Surgery Room"},
		{Value: 0x6e, Name: "Medical Imaging Room"},
		{Value: 0x6f, Name: "Decontamination Room"},
		{Value: 0xFF, Name: "Unknown Environment"},
	}
}

type Power zcl.Zs16

const PowerAttr zcl.AttrID = 19

func (Power) ID() zcl.AttrID   { return PowerAttr }
func (Power) Readable() bool   { return true }
func (Power) Writable() bool   { return true }
func (Power) Reportable() bool { return false }
func (Power) SceneIndex() int  { return -1 }

func (Power) Name() string                  { return "Power" }
func (a *Power) TypeID() zcl.TypeID         { return zcl.Zs16(*a).ID() }
func (a *Power) Value() zcl.Val             { return a }
func (a Power) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *Power) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Power(*nt)
	return br, err
}

func (a *Power) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = Power(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Power) String() string {
	return zcl.DecibelMilliWatts.Format(float64(a) / 100)
}

type PowerOnLevel zcl.Zu8

const PowerOnLevelAttr zcl.AttrID = 16384

func (PowerOnLevel) ID() zcl.AttrID   { return PowerOnLevelAttr }
func (PowerOnLevel) Readable() bool   { return true }
func (PowerOnLevel) Writable() bool   { return true }
func (PowerOnLevel) Reportable() bool { return false }
func (PowerOnLevel) SceneIndex() int  { return -1 }

func (PowerOnLevel) Name() string                  { return "Power On level" }
func (a *PowerOnLevel) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *PowerOnLevel) Value() zcl.Val             { return a }
func (a PowerOnLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *PowerOnLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerOnLevel(*nt)
	return br, err
}

func (a *PowerOnLevel) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = PowerOnLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PowerOnLevel) String() string {
	return zcl.Percent.Format(float64(a) / 2.54)
}

type PowerSource zcl.Zenum8

const PowerSourceAttr zcl.AttrID = 7

func (PowerSource) ID() zcl.AttrID   { return PowerSourceAttr }
func (PowerSource) Readable() bool   { return true }
func (PowerSource) Writable() bool   { return false }
func (PowerSource) Reportable() bool { return false }
func (PowerSource) SceneIndex() int  { return -1 }

func (PowerSource) Name() string                  { return "Power Source" }
func (a *PowerSource) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *PowerSource) Value() zcl.Val             { return a }
func (a PowerSource) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *PowerSource) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerSource(*nt)
	return br, err
}

func (a *PowerSource) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = PowerSource(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PowerSource) String() string {
	switch a {
	case 0x00:
		return "Unknown"
	case 0x01:
		return "Mains (single phase)"
	case 0x02:
		return "Mains (3 phase)"
	case 0x03:
		return "Battery"
	case 0x04:
		return "DC Source"
	case 0x05:
		return "Emergency mains constantly powered"
	case 0x06:
		return "Emergency mains and transfer switch"
	case 0x80:
		return "Unknown with battery backup"
	case 0x81:
		return "Mains (single phase) with battery backup"
	case 0x82:
		return "Mains (3 phase) with battery backup"
	case 0x83:
		return "Battery with battery backup"
	case 0x84:
		return "DC Source with battery backup"
	case 0x85:
		return "Emergency mains constantly powered with battery backup"
	case 0x86:
		return "Emergency mains and transfer switch with battery backup"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a PowerSource) IsUnknown() bool                                          { return a == 0x00 }
func (a PowerSource) IsMainsSinglePhase() bool                                 { return a == 0x01 }
func (a PowerSource) IsMains3Phase() bool                                      { return a == 0x02 }
func (a PowerSource) IsBattery() bool                                          { return a == 0x03 }
func (a PowerSource) IsDcSource() bool                                         { return a == 0x04 }
func (a PowerSource) IsEmergencyMainsConstantlyPowered() bool                  { return a == 0x05 }
func (a PowerSource) IsEmergencyMainsAndTransferSwitch() bool                  { return a == 0x06 }
func (a PowerSource) IsUnknownWithBatteryBackup() bool                         { return a == 0x80 }
func (a PowerSource) IsMainsSinglePhaseWithBatteryBackup() bool                { return a == 0x81 }
func (a PowerSource) IsMains3PhaseWithBatteryBackup() bool                     { return a == 0x82 }
func (a PowerSource) IsBatteryWithBatteryBackup() bool                         { return a == 0x83 }
func (a PowerSource) IsDcSourceWithBatteryBackup() bool                        { return a == 0x84 }
func (a PowerSource) IsEmergencyMainsConstantlyPoweredWithBatteryBackup() bool { return a == 0x85 }
func (a PowerSource) IsEmergencyMainsAndTransferSwitchWithBatteryBackup() bool { return a == 0x86 }
func (a *PowerSource) SetUnknown()                                             { *a = 0x00 }
func (a *PowerSource) SetMainsSinglePhase()                                    { *a = 0x01 }
func (a *PowerSource) SetMains3Phase()                                         { *a = 0x02 }
func (a *PowerSource) SetBattery()                                             { *a = 0x03 }
func (a *PowerSource) SetDcSource()                                            { *a = 0x04 }
func (a *PowerSource) SetEmergencyMainsConstantlyPowered()                     { *a = 0x05 }
func (a *PowerSource) SetEmergencyMainsAndTransferSwitch()                     { *a = 0x06 }
func (a *PowerSource) SetUnknownWithBatteryBackup()                            { *a = 0x80 }
func (a *PowerSource) SetMainsSinglePhaseWithBatteryBackup()                   { *a = 0x81 }
func (a *PowerSource) SetMains3PhaseWithBatteryBackup()                        { *a = 0x82 }
func (a *PowerSource) SetBatteryWithBatteryBackup()                            { *a = 0x83 }
func (a *PowerSource) SetDcSourceWithBatteryBackup()                           { *a = 0x84 }
func (a *PowerSource) SetEmergencyMainsConstantlyPoweredWithBatteryBackup()    { *a = 0x85 }
func (a *PowerSource) SetEmergencyMainsAndTransferSwitchWithBatteryBackup()    { *a = 0x86 }

func (PowerSource) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Unknown"},
		{Value: 0x01, Name: "Mains (single phase)"},
		{Value: 0x02, Name: "Mains (3 phase)"},
		{Value: 0x03, Name: "Battery"},
		{Value: 0x04, Name: "DC Source"},
		{Value: 0x05, Name: "Emergency mains constantly powered"},
		{Value: 0x06, Name: "Emergency mains and transfer switch"},
		{Value: 0x80, Name: "Unknown with battery backup"},
		{Value: 0x81, Name: "Mains (single phase) with battery backup"},
		{Value: 0x82, Name: "Mains (3 phase) with battery backup"},
		{Value: 0x83, Name: "Battery with battery backup"},
		{Value: 0x84, Name: "DC Source with battery backup"},
		{Value: 0x85, Name: "Emergency mains constantly powered with battery backup"},
		{Value: 0x86, Name: "Emergency mains and transfer switch with battery backup"},
	}
}

type PoweronOnOff zcl.Zenum8

const PoweronOnOffAttr zcl.AttrID = 16387

func (PoweronOnOff) ID() zcl.AttrID   { return PoweronOnOffAttr }
func (PoweronOnOff) Readable() bool   { return true }
func (PoweronOnOff) Writable() bool   { return true }
func (PoweronOnOff) Reportable() bool { return false }
func (PoweronOnOff) SceneIndex() int  { return -1 }

func (PoweronOnOff) Name() string                  { return "PowerOn On/Off" }
func (a *PoweronOnOff) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *PoweronOnOff) Value() zcl.Val             { return a }
func (a PoweronOnOff) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *PoweronOnOff) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = PoweronOnOff(*nt)
	return br, err
}

func (a *PoweronOnOff) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = PoweronOnOff(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PoweronOnOff) String() string {
	switch a {
	case 0x00:
		return "Off"
	case 0x01:
		return "On"
	case 0xFF:
		return "Previous"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a PoweronOnOff) IsOff() bool      { return a == 0x00 }
func (a PoweronOnOff) IsOn() bool       { return a == 0x01 }
func (a PoweronOnOff) IsPrevious() bool { return a == 0xFF }
func (a *PoweronOnOff) SetOff()         { *a = 0x00 }
func (a *PoweronOnOff) SetOn()          { *a = 0x01 }
func (a *PoweronOnOff) SetPrevious()    { *a = 0xFF }

func (PoweronOnOff) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Off"},
		{Value: 0x01, Name: "On"},
		{Value: 0xFF, Name: "Previous"},
	}
}

// ProductUrl specifies a link to a web page containing specific product information.
type ProductUrl zcl.Zcstring

const ProductUrlAttr zcl.AttrID = 11

func (ProductUrl) ID() zcl.AttrID   { return ProductUrlAttr }
func (ProductUrl) Readable() bool   { return true }
func (ProductUrl) Writable() bool   { return false }
func (ProductUrl) Reportable() bool { return false }
func (ProductUrl) SceneIndex() int  { return -1 }

func (ProductUrl) Name() string                  { return "Product URL" }
func (a *ProductUrl) TypeID() zcl.TypeID         { return zcl.Zcstring(*a).ID() }
func (a *ProductUrl) Value() zcl.Val             { return a }
func (a ProductUrl) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *ProductUrl) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = ProductUrl(*nt)
	return br, err
}

func (a *ProductUrl) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = ProductUrl(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ProductUrl) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

// ProductCode As printed on the product.
type ProductCode zcl.Zostring

const ProductCodeAttr zcl.AttrID = 10

func (ProductCode) ID() zcl.AttrID   { return ProductCodeAttr }
func (ProductCode) Readable() bool   { return true }
func (ProductCode) Writable() bool   { return false }
func (ProductCode) Reportable() bool { return false }
func (ProductCode) SceneIndex() int  { return -1 }

func (ProductCode) Name() string                  { return "Product code" }
func (a *ProductCode) TypeID() zcl.TypeID         { return zcl.Zostring(*a).ID() }
func (a *ProductCode) Value() zcl.Val             { return a }
func (a ProductCode) MarshalZcl() ([]byte, error) { return zcl.Zostring(a).MarshalZcl() }

func (a *ProductCode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*a = ProductCode(*nt)
	return br, err
}

func (a *ProductCode) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zostring); ok {
		*a = ProductCode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ProductCode) String() string {
	return zcl.Sprintf("%v", zcl.Zostring(a))
}

type QualityMeasure zcl.Zu8

const QualityMeasureAttr zcl.AttrID = 3

func (QualityMeasure) ID() zcl.AttrID   { return QualityMeasureAttr }
func (QualityMeasure) Readable() bool   { return true }
func (QualityMeasure) Writable() bool   { return false }
func (QualityMeasure) Reportable() bool { return false }
func (QualityMeasure) SceneIndex() int  { return -1 }

func (QualityMeasure) Name() string                  { return "Quality Measure" }
func (a *QualityMeasure) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *QualityMeasure) Value() zcl.Val             { return a }
func (a QualityMeasure) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *QualityMeasure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = QualityMeasure(*nt)
	return br, err
}

func (a *QualityMeasure) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = QualityMeasure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a QualityMeasure) String() string {
	return zcl.Percent.Format(float64(a))
}

type QualityIndex zcl.Zu16

func (QualityIndex) Name() string                  { return "Quality index" }
func (a *QualityIndex) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *QualityIndex) Value() zcl.Val             { return a }
func (a QualityIndex) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *QualityIndex) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = QualityIndex(*nt)
	return br, err
}

func (a *QualityIndex) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = QualityIndex(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a QualityIndex) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type Rssi zcl.Zs8

func (Rssi) Name() string                  { return "RSSI" }
func (a *Rssi) TypeID() zcl.TypeID         { return zcl.Zs8(*a).ID() }
func (a *Rssi) Value() zcl.Val             { return a }
func (a Rssi) MarshalZcl() ([]byte, error) { return zcl.Zs8(a).MarshalZcl() }

func (a *Rssi) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = Rssi(*nt)
	return br, err
}

func (a *Rssi) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs8); ok {
		*a = Rssi(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Rssi) String() string {
	return zcl.DecibelMilliWatts.Format(float64(a))
}

type Rate zcl.Zu8

func (Rate) Name() string                  { return "Rate" }
func (a *Rate) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *Rate) Value() zcl.Val             { return a }
func (a Rate) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Rate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Rate(*nt)
	return br, err
}

func (a *Rate) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = Rate(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Rate) String() string {
	return zcl.PercentPerSecond.Format(float64(a) / 2.54)
}

type RelayedUcast zcl.Zu16

const RelayedUcastAttr zcl.AttrID = 280

func (RelayedUcast) ID() zcl.AttrID   { return RelayedUcastAttr }
func (RelayedUcast) Readable() bool   { return true }
func (RelayedUcast) Writable() bool   { return false }
func (RelayedUcast) Reportable() bool { return false }
func (RelayedUcast) SceneIndex() int  { return -1 }

func (RelayedUcast) Name() string                  { return "Relayed Ucast" }
func (a *RelayedUcast) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *RelayedUcast) Value() zcl.Val             { return a }
func (a RelayedUcast) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *RelayedUcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RelayedUcast(*nt)
	return br, err
}

func (a *RelayedUcast) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = RelayedUcast(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a RelayedUcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// RemainingTime represents the time remaining until the current command is complete - it is specified in 1/10ths of a second.
type RemainingTime zcl.Zu16

const RemainingTimeAttr zcl.AttrID = 1

func (RemainingTime) ID() zcl.AttrID   { return RemainingTimeAttr }
func (RemainingTime) Readable() bool   { return true }
func (RemainingTime) Writable() bool   { return false }
func (RemainingTime) Reportable() bool { return false }
func (RemainingTime) SceneIndex() int  { return -1 }

func (RemainingTime) Name() string                  { return "Remaining Time" }
func (a *RemainingTime) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *RemainingTime) Value() zcl.Val             { return a }
func (a RemainingTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *RemainingTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RemainingTime(*nt)
	return br, err
}

func (a *RemainingTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = RemainingTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a RemainingTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

type ReportingPeriod zcl.Zu16

const ReportingPeriodAttr zcl.AttrID = 21

func (ReportingPeriod) ID() zcl.AttrID   { return ReportingPeriodAttr }
func (ReportingPeriod) Readable() bool   { return true }
func (ReportingPeriod) Writable() bool   { return true }
func (ReportingPeriod) Reportable() bool { return false }
func (ReportingPeriod) SceneIndex() int  { return -1 }

func (ReportingPeriod) Name() string                  { return "Reporting Period" }
func (a *ReportingPeriod) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ReportingPeriod) Value() zcl.Val             { return a }
func (a ReportingPeriod) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ReportingPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReportingPeriod(*nt)
	return br, err
}

func (a *ReportingPeriod) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ReportingPeriod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ReportingPeriod) String() string {
	return zcl.Seconds.Format(float64(a))
}

type Resolution zcl.Zenum8

func (Resolution) Name() string                  { return "Resolution" }
func (a *Resolution) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *Resolution) Value() zcl.Val             { return a }
func (a Resolution) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *Resolution) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = Resolution(*nt)
	return br, err
}

func (a *Resolution) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = Resolution(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Resolution) String() string {
	switch a {
	case 0x00:
		return "High"
	case 0x01:
		return "Mid"
	case 0x02:
		return "Low"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a Resolution) IsHigh() bool { return a == 0x00 }
func (a Resolution) IsMid() bool  { return a == 0x01 }
func (a Resolution) IsLow() bool  { return a == 0x02 }
func (a *Resolution) SetHigh()    { *a = 0x00 }
func (a *Resolution) SetMid()     { *a = 0x01 }
func (a *Resolution) SetLow()     { *a = 0x02 }

func (Resolution) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "High"},
		{Value: 0x01, Name: "Mid"},
		{Value: 0x02, Name: "Low"},
	}
}

type RouteDiscInitiated zcl.Zu16

const RouteDiscInitiatedAttr zcl.AttrID = 268

func (RouteDiscInitiated) ID() zcl.AttrID   { return RouteDiscInitiatedAttr }
func (RouteDiscInitiated) Readable() bool   { return true }
func (RouteDiscInitiated) Writable() bool   { return false }
func (RouteDiscInitiated) Reportable() bool { return false }
func (RouteDiscInitiated) SceneIndex() int  { return -1 }

func (RouteDiscInitiated) Name() string                  { return "Route Disc Initiated" }
func (a *RouteDiscInitiated) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *RouteDiscInitiated) Value() zcl.Val             { return a }
func (a RouteDiscInitiated) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *RouteDiscInitiated) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RouteDiscInitiated(*nt)
	return br, err
}

func (a *RouteDiscInitiated) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = RouteDiscInitiated(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a RouteDiscInitiated) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type SwBuildId zcl.Zcstring

const SwBuildIdAttr zcl.AttrID = 16384

func (SwBuildId) ID() zcl.AttrID   { return SwBuildIdAttr }
func (SwBuildId) Readable() bool   { return true }
func (SwBuildId) Writable() bool   { return false }
func (SwBuildId) Reportable() bool { return false }
func (SwBuildId) SceneIndex() int  { return -1 }

func (SwBuildId) Name() string                  { return "SW Build ID" }
func (a *SwBuildId) TypeID() zcl.TypeID         { return zcl.Zcstring(*a).ID() }
func (a *SwBuildId) Value() zcl.Val             { return a }
func (a SwBuildId) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *SwBuildId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = SwBuildId(*nt)
	return br, err
}

func (a *SwBuildId) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = SwBuildId(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SwBuildId) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

// SceneCapacity specifies remaining number of scenes that can be added
type SceneCapacity zcl.Zu8

func (SceneCapacity) Name() string                  { return "Scene Capacity" }
func (a *SceneCapacity) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *SceneCapacity) Value() zcl.Val             { return a }
func (a SceneCapacity) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *SceneCapacity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneCapacity(*nt)
	return br, err
}

func (a *SceneCapacity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = SceneCapacity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneCapacity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type SceneCopyMode zcl.Zbmp8

func (SceneCopyMode) Name() string                  { return "Scene Copy Mode" }
func (a *SceneCopyMode) TypeID() zcl.TypeID         { return zcl.Zbmp8(*a).ID() }
func (a *SceneCopyMode) Value() zcl.Val             { return a }
func (a SceneCopyMode) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *SceneCopyMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneCopyMode(*nt)
	return br, err
}

func (a *SceneCopyMode) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = SceneCopyMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneCopyMode) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Copy All Scenes")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a SceneCopyMode) IsCopyAllScenes() bool    { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a *SceneCopyMode) SetCopyAllScenes(b bool) { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b)) }

func (SceneCopyMode) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Copy All Scenes"},
	}
}

type SceneCount zcl.Zu8

const SceneCountAttr zcl.AttrID = 0

func (SceneCount) ID() zcl.AttrID   { return SceneCountAttr }
func (SceneCount) Readable() bool   { return true }
func (SceneCount) Writable() bool   { return false }
func (SceneCount) Reportable() bool { return false }
func (SceneCount) SceneIndex() int  { return -1 }

func (SceneCount) Name() string                  { return "Scene Count" }
func (a *SceneCount) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *SceneCount) Value() zcl.Val             { return a }
func (a SceneCount) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *SceneCount) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneCount(*nt)
	return br, err
}

func (a *SceneCount) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = SceneCount(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneCount) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// SceneExtensions The format of each extension field set is a 16 bit field carrying the cluster ID,
// followed by an 8 bit length field and the set of scene extension fields specified
// in  the  relevant  cluster. The length field holds the length in octets of that
// extension field set. Extension field set format:
// {{clusterId1, length 1, {extension field set 1}}, {clusterId2, length 2, {extension field set 2}} ...}
// I.e. the field would be a repeating struct with [ClusterID uint16] [OctetLength uint8] [AttrID ...uint16]
type SceneExtensions zcl.SceneExtensionSet

func (SceneExtensions) Name() string                  { return "Scene Extensions" }
func (a *SceneExtensions) TypeID() zcl.TypeID         { return zcl.SceneExtensionSet(*a).ID() }
func (a *SceneExtensions) Value() zcl.Val             { return a }
func (a SceneExtensions) MarshalZcl() ([]byte, error) { return zcl.SceneExtensionSet(a).MarshalZcl() }

func (a *SceneExtensions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.SceneExtensionSet)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneExtensions(*nt)
	return br, err
}

func (a *SceneExtensions) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.SceneExtensionSet); ok {
		*a = SceneExtensions(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneExtensions) String() string {
	return zcl.Sprintf("%v", zcl.SceneExtensionSet(a))
}

type SceneId zcl.Zu8

func (SceneId) Name() string                  { return "Scene ID" }
func (a *SceneId) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *SceneId) Value() zcl.Val             { return a }
func (a SceneId) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *SceneId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneId(*nt)
	return br, err
}

func (a *SceneId) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = SceneId(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneId) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type SceneLastConfiguredBy zcl.Zuid

const SceneLastConfiguredByAttr zcl.AttrID = 5

func (SceneLastConfiguredBy) ID() zcl.AttrID   { return SceneLastConfiguredByAttr }
func (SceneLastConfiguredBy) Readable() bool   { return true }
func (SceneLastConfiguredBy) Writable() bool   { return false }
func (SceneLastConfiguredBy) Reportable() bool { return false }
func (SceneLastConfiguredBy) SceneIndex() int  { return -1 }

func (SceneLastConfiguredBy) Name() string                  { return "Scene Last Configured By" }
func (a *SceneLastConfiguredBy) TypeID() zcl.TypeID         { return zcl.Zuid(*a).ID() }
func (a *SceneLastConfiguredBy) Value() zcl.Val             { return a }
func (a SceneLastConfiguredBy) MarshalZcl() ([]byte, error) { return zcl.Zuid(a).MarshalZcl() }

func (a *SceneLastConfiguredBy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneLastConfiguredBy(*nt)
	return br, err
}

func (a *SceneLastConfiguredBy) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zuid); ok {
		*a = SceneLastConfiguredBy(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneLastConfiguredBy) String() string {
	return zcl.Sprintf("%v", zcl.Zuid(a))
}

type SceneName zcl.Zcstring

func (SceneName) Name() string                  { return "Scene Name" }
func (a *SceneName) TypeID() zcl.TypeID         { return zcl.Zcstring(*a).ID() }
func (a *SceneName) Value() zcl.Val             { return a }
func (a SceneName) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *SceneName) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneName(*nt)
	return br, err
}

func (a *SceneName) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = SceneName(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneName) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

type SceneNameSupport zcl.Zbmp8

const SceneNameSupportAttr zcl.AttrID = 4

func (SceneNameSupport) ID() zcl.AttrID   { return SceneNameSupportAttr }
func (SceneNameSupport) Readable() bool   { return true }
func (SceneNameSupport) Writable() bool   { return false }
func (SceneNameSupport) Reportable() bool { return false }
func (SceneNameSupport) SceneIndex() int  { return -1 }

func (SceneNameSupport) Name() string                  { return "Scene Name Support" }
func (a *SceneNameSupport) TypeID() zcl.TypeID         { return zcl.Zbmp8(*a).ID() }
func (a *SceneNameSupport) Value() zcl.Val             { return a }
func (a SceneNameSupport) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *SceneNameSupport) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneNameSupport(*nt)
	return br, err
}

func (a *SceneNameSupport) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = SceneNameSupport(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneNameSupport) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 7:
			bstr = append(bstr, "Names Supported")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a SceneNameSupport) IsNamesSupported() bool { return zcl.BitmapTest([]byte(a[:]), 7) }
func (a *SceneNameSupport) SetNamesSupported(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 7, b))
}

func (SceneNameSupport) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 7, Name: "Names Supported"},
	}
}

type SceneValid zcl.Zbool

const SceneValidAttr zcl.AttrID = 3

func (SceneValid) ID() zcl.AttrID   { return SceneValidAttr }
func (SceneValid) Readable() bool   { return true }
func (SceneValid) Writable() bool   { return false }
func (SceneValid) Reportable() bool { return false }
func (SceneValid) SceneIndex() int  { return -1 }

func (SceneValid) Name() string                  { return "Scene Valid" }
func (a *SceneValid) TypeID() zcl.TypeID         { return zcl.Zbool(*a).ID() }
func (a *SceneValid) Value() zcl.Val             { return a }
func (a SceneValid) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *SceneValid) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneValid(*nt)
	return br, err
}

func (a *SceneValid) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbool); ok {
		*a = SceneValid(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneValid) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

type SceneList zcl.Zset

func (SceneList) Name() string          { return "Scene list" }
func (a *SceneList) TypeID() zcl.TypeID { return zcl.Zset(*a).ID() }
func (a *SceneList) Value() zcl.Val     { return a }
func (a SceneList) MarshalZcl() ([]byte, error) {
	return nil, zcl.Errorf("not implemented")
}

func (a *SceneList) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zset)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneList(*nt)
	return br, err
}

func (a *SceneList) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zset); ok {
		*a = SceneList(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneList) String() string {
	return zcl.Sprintf("%v", zcl.Zset(a))
}

type Sensitivity zcl.Zenum8

const SensitivityAttr zcl.AttrID = 48

func (Sensitivity) ID() zcl.AttrID   { return SensitivityAttr }
func (Sensitivity) Readable() bool   { return true }
func (Sensitivity) Writable() bool   { return true }
func (Sensitivity) Reportable() bool { return false }
func (Sensitivity) SceneIndex() int  { return -1 }

func (Sensitivity) Name() string                  { return "Sensitivity" }
func (a *Sensitivity) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *Sensitivity) Value() zcl.Val             { return a }
func (a Sensitivity) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *Sensitivity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = Sensitivity(*nt)
	return br, err
}

func (a *Sensitivity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = Sensitivity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Sensitivity) String() string {
	switch a {
	case 0x00:
		return "default"
	case 0x01:
		return "high"
	case 0x02:
		return "max"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a Sensitivity) IsDefault() bool { return a == 0x00 }
func (a Sensitivity) IsHigh() bool    { return a == 0x01 }
func (a Sensitivity) IsMax() bool     { return a == 0x02 }
func (a *Sensitivity) SetDefault()    { *a = 0x00 }
func (a *Sensitivity) SetHigh()       { *a = 0x01 }
func (a *Sensitivity) SetMax()        { *a = 0x02 }

func (Sensitivity) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "default"},
		{Value: 0x01, Name: "high"},
		{Value: 0x02, Name: "max"},
	}
}

type ShortPollInterval zcl.Zu16

const ShortPollIntervalAttr zcl.AttrID = 2

func (ShortPollInterval) ID() zcl.AttrID   { return ShortPollIntervalAttr }
func (ShortPollInterval) Readable() bool   { return true }
func (ShortPollInterval) Writable() bool   { return false }
func (ShortPollInterval) Reportable() bool { return false }
func (ShortPollInterval) SceneIndex() int  { return -1 }

func (ShortPollInterval) Name() string                  { return "Short Poll Interval" }
func (a *ShortPollInterval) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ShortPollInterval) Value() zcl.Val             { return a }
func (a ShortPollInterval) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ShortPollInterval) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ShortPollInterval(*nt)
	return br, err
}

func (a *ShortPollInterval) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ShortPollInterval(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ShortPollInterval) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type StackVersion zcl.Zu8

const StackVersionAttr zcl.AttrID = 2

func (StackVersion) ID() zcl.AttrID   { return StackVersionAttr }
func (StackVersion) Readable() bool   { return true }
func (StackVersion) Writable() bool   { return false }
func (StackVersion) Reportable() bool { return false }
func (StackVersion) SceneIndex() int  { return -1 }

func (StackVersion) Name() string                  { return "Stack Version" }
func (a *StackVersion) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *StackVersion) Value() zcl.Val             { return a }
func (a StackVersion) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *StackVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = StackVersion(*nt)
	return br, err
}

func (a *StackVersion) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = StackVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a StackVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// StandardTime Local time (without DST offset)
type StandardTime zcl.Zu32

const StandardTimeAttr zcl.AttrID = 6

func (StandardTime) ID() zcl.AttrID   { return StandardTimeAttr }
func (StandardTime) Readable() bool   { return true }
func (StandardTime) Writable() bool   { return false }
func (StandardTime) Reportable() bool { return false }
func (StandardTime) SceneIndex() int  { return -1 }

func (StandardTime) Name() string                  { return "Standard Time" }
func (a *StandardTime) TypeID() zcl.TypeID         { return zcl.Zu32(*a).ID() }
func (a *StandardTime) Value() zcl.Val             { return a }
func (a StandardTime) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *StandardTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = StandardTime(*nt)
	return br, err
}

func (a *StandardTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu32); ok {
		*a = StandardTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a StandardTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

type Status zcl.Status

func (Status) Name() string                  { return "Status" }
func (a *Status) TypeID() zcl.TypeID         { return zcl.Status(*a).ID() }
func (a *Status) Value() zcl.Val             { return a }
func (a Status) MarshalZcl() ([]byte, error) { return zcl.Status(a).MarshalZcl() }

func (a *Status) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Status)
	br, err := nt.UnmarshalZcl(b)
	*a = Status(*nt)
	return br, err
}

func (a *Status) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Status); ok {
		*a = Status(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Status) String() string {
	return zcl.Sprintf("%v", zcl.Status(a))
}

type StepSize zcl.Zu8

func (StepSize) Name() string                  { return "Step size" }
func (a *StepSize) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *StepSize) Value() zcl.Val             { return a }
func (a StepSize) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *StepSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = StepSize(*nt)
	return br, err
}

func (a *StepSize) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = StepSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a StepSize) String() string {
	return zcl.Percent.Format(float64(a) / 2.54)
}

// SwitchActions specifies the commands of the On/Off cluster to be generated when the switch moves between its two states.
type SwitchActions zcl.Zenum8

const SwitchActionsAttr zcl.AttrID = 16

func (SwitchActions) ID() zcl.AttrID   { return SwitchActionsAttr }
func (SwitchActions) Readable() bool   { return true }
func (SwitchActions) Writable() bool   { return true }
func (SwitchActions) Reportable() bool { return false }
func (SwitchActions) SceneIndex() int  { return -1 }

func (SwitchActions) Name() string                  { return "Switch actions" }
func (a *SwitchActions) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *SwitchActions) Value() zcl.Val             { return a }
func (a SwitchActions) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *SwitchActions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = SwitchActions(*nt)
	return br, err
}

func (a *SwitchActions) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = SwitchActions(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SwitchActions) String() string {
	switch a {
	case 0x00:
		return "On-Off"
	case 0x01:
		return "Off-On"
	case 0x02:
		return "Toggle"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a SwitchActions) IsOnOff() bool  { return a == 0x00 }
func (a SwitchActions) IsOffOn() bool  { return a == 0x01 }
func (a SwitchActions) IsToggle() bool { return a == 0x02 }
func (a *SwitchActions) SetOnOff()     { *a = 0x00 }
func (a *SwitchActions) SetOffOn()     { *a = 0x01 }
func (a *SwitchActions) SetToggle()    { *a = 0x02 }

func (SwitchActions) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "On-Off"},
		{Value: 0x01, Name: "Off-On"},
		{Value: 0x02, Name: "Toggle"},
	}
}

// SwitchType specifies the basic functionality of the On/Off switching device.
type SwitchType zcl.Zenum8

const SwitchTypeAttr zcl.AttrID = 0

func (SwitchType) ID() zcl.AttrID   { return SwitchTypeAttr }
func (SwitchType) Readable() bool   { return true }
func (SwitchType) Writable() bool   { return false }
func (SwitchType) Reportable() bool { return false }
func (SwitchType) SceneIndex() int  { return -1 }

func (SwitchType) Name() string                  { return "Switch type" }
func (a *SwitchType) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *SwitchType) Value() zcl.Val             { return a }
func (a SwitchType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *SwitchType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = SwitchType(*nt)
	return br, err
}

func (a *SwitchType) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = SwitchType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SwitchType) String() string {
	switch a {
	case 0x00:
		return "Toggle"
	case 0x01:
		return "Momentary"
	case 0x02:
		return "Multifunction"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a SwitchType) IsToggle() bool        { return a == 0x00 }
func (a SwitchType) IsMomentary() bool     { return a == 0x01 }
func (a SwitchType) IsMultifunction() bool { return a == 0x02 }
func (a *SwitchType) SetToggle()           { *a = 0x00 }
func (a *SwitchType) SetMomentary()        { *a = 0x01 }
func (a *SwitchType) SetMultifunction()    { *a = 0x02 }

func (SwitchType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Toggle"},
		{Value: 0x01, Name: "Momentary"},
		{Value: 0x02, Name: "Multifunction"},
	}
}

type Time zcl.Zutc

const TimeAttr zcl.AttrID = 0

func (Time) ID() zcl.AttrID   { return TimeAttr }
func (Time) Readable() bool   { return true }
func (Time) Writable() bool   { return true }
func (Time) Reportable() bool { return false }
func (Time) SceneIndex() int  { return -1 }

func (Time) Name() string                  { return "Time" }
func (a *Time) TypeID() zcl.TypeID         { return zcl.Zutc(*a).ID() }
func (a *Time) Value() zcl.Val             { return a }
func (a Time) MarshalZcl() ([]byte, error) { return zcl.Zutc(a).MarshalZcl() }

func (a *Time) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*a = Time(*nt)
	return br, err
}

func (a *Time) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zutc); ok {
		*a = Time(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Time) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(a))
}

type TimeStatus zcl.Zbmp8

const TimeStatusAttr zcl.AttrID = 1

func (TimeStatus) ID() zcl.AttrID   { return TimeStatusAttr }
func (TimeStatus) Readable() bool   { return true }
func (TimeStatus) Writable() bool   { return true }
func (TimeStatus) Reportable() bool { return false }
func (TimeStatus) SceneIndex() int  { return -1 }

func (TimeStatus) Name() string                  { return "Time Status" }
func (a *TimeStatus) TypeID() zcl.TypeID         { return zcl.Zbmp8(*a).ID() }
func (a *TimeStatus) Value() zcl.Val             { return a }
func (a TimeStatus) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *TimeStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = TimeStatus(*nt)
	return br, err
}

func (a *TimeStatus) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = TimeStatus(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a TimeStatus) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Master Clock")
		case 1:
			bstr = append(bstr, "Synchronized")
		case 2:
			bstr = append(bstr, "Master for Timezone and Dst")
		case 3:
			bstr = append(bstr, "Superseding")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a TimeStatus) IsMasterClock() bool             { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a TimeStatus) IsSynchronized() bool            { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a TimeStatus) IsMasterForTimezoneAndDst() bool { return zcl.BitmapTest([]byte(a[:]), 2) }
func (a TimeStatus) IsSuperseding() bool             { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a *TimeStatus) SetMasterClock(b bool)          { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b)) }
func (a *TimeStatus) SetSynchronized(b bool)         { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b)) }
func (a *TimeStatus) SetMasterForTimezoneAndDst(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b))
}
func (a *TimeStatus) SetSuperseding(b bool) { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b)) }

func (TimeStatus) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Master Clock"},
		{Value: 1, Name: "Synchronized"},
		{Value: 2, Name: "Master for Timezone and Dst"},
		{Value: 3, Name: "Superseding"},
	}
}

// TimeZone Offset during normal time from UTC in seconds
type TimeZone zcl.Zs32

const TimeZoneAttr zcl.AttrID = 2

func (TimeZone) ID() zcl.AttrID   { return TimeZoneAttr }
func (TimeZone) Readable() bool   { return true }
func (TimeZone) Writable() bool   { return true }
func (TimeZone) Reportable() bool { return false }
func (TimeZone) SceneIndex() int  { return -1 }

func (TimeZone) Name() string                  { return "Time Zone" }
func (a *TimeZone) TypeID() zcl.TypeID         { return zcl.Zs32(*a).ID() }
func (a *TimeZone) Value() zcl.Val             { return a }
func (a TimeZone) MarshalZcl() ([]byte, error) { return zcl.Zs32(a).MarshalZcl() }

func (a *TimeZone) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs32)
	br, err := nt.UnmarshalZcl(b)
	*a = TimeZone(*nt)
	return br, err
}

func (a *TimeZone) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs32); ok {
		*a = TimeZone(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a TimeZone) String() string {
	return zcl.Seconds.Format(float64(a))
}

type TransitionTime zcl.Zu16

func (TransitionTime) Name() string                  { return "Transition Time" }
func (a *TransitionTime) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *TransitionTime) Value() zcl.Val             { return a }
func (a TransitionTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *TransitionTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = TransitionTime(*nt)
	return br, err
}

func (a *TransitionTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = TransitionTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a TransitionTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

type TransitionTimeSec zcl.Zu16

func (TransitionTimeSec) Name() string                  { return "Transition time (Sec)" }
func (a *TransitionTimeSec) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *TransitionTimeSec) Value() zcl.Val             { return a }
func (a TransitionTimeSec) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *TransitionTimeSec) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = TransitionTimeSec(*nt)
	return br, err
}

func (a *TransitionTimeSec) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = TransitionTimeSec(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a TransitionTimeSec) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

type UserTest zcl.Zbool

const UserTestAttr zcl.AttrID = 50

func (UserTest) ID() zcl.AttrID   { return UserTestAttr }
func (UserTest) Readable() bool   { return true }
func (UserTest) Writable() bool   { return true }
func (UserTest) Reportable() bool { return false }
func (UserTest) SceneIndex() int  { return -1 }

func (UserTest) Name() string                  { return "User test" }
func (a *UserTest) TypeID() zcl.TypeID         { return zcl.Zbool(*a).ID() }
func (a *UserTest) Value() zcl.Val             { return a }
func (a UserTest) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *UserTest) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = UserTest(*nt)
	return br, err
}

func (a *UserTest) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbool); ok {
		*a = UserTest(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a UserTest) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

type ValidUntilTime zcl.Zutc

const ValidUntilTimeAttr zcl.AttrID = 9

func (ValidUntilTime) ID() zcl.AttrID   { return ValidUntilTimeAttr }
func (ValidUntilTime) Readable() bool   { return true }
func (ValidUntilTime) Writable() bool   { return true }
func (ValidUntilTime) Reportable() bool { return false }
func (ValidUntilTime) SceneIndex() int  { return -1 }

func (ValidUntilTime) Name() string                  { return "Valid Until Time" }
func (a *ValidUntilTime) TypeID() zcl.TypeID         { return zcl.Zutc(*a).ID() }
func (a *ValidUntilTime) Value() zcl.Val             { return a }
func (a ValidUntilTime) MarshalZcl() ([]byte, error) { return zcl.Zutc(a).MarshalZcl() }

func (a *ValidUntilTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*a = ValidUntilTime(*nt)
	return br, err
}

func (a *ValidUntilTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zutc); ok {
		*a = ValidUntilTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ValidUntilTime) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(a))
}

type XCoordinate zcl.Zs16

const XCoordinateAttr zcl.AttrID = 16

func (XCoordinate) ID() zcl.AttrID   { return XCoordinateAttr }
func (XCoordinate) Readable() bool   { return true }
func (XCoordinate) Writable() bool   { return true }
func (XCoordinate) Reportable() bool { return false }
func (XCoordinate) SceneIndex() int  { return -1 }

func (XCoordinate) Name() string                  { return "X Coordinate" }
func (a *XCoordinate) TypeID() zcl.TypeID         { return zcl.Zs16(*a).ID() }
func (a *XCoordinate) Value() zcl.Val             { return a }
func (a XCoordinate) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *XCoordinate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = XCoordinate(*nt)
	return br, err
}

func (a *XCoordinate) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = XCoordinate(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a XCoordinate) String() string {
	return zcl.Meters.Format(float64(a) / 10)
}

type YCoordinate zcl.Zs16

const YCoordinateAttr zcl.AttrID = 17

func (YCoordinate) ID() zcl.AttrID   { return YCoordinateAttr }
func (YCoordinate) Readable() bool   { return true }
func (YCoordinate) Writable() bool   { return true }
func (YCoordinate) Reportable() bool { return false }
func (YCoordinate) SceneIndex() int  { return -1 }

func (YCoordinate) Name() string                  { return "Y Coordinate" }
func (a *YCoordinate) TypeID() zcl.TypeID         { return zcl.Zs16(*a).ID() }
func (a *YCoordinate) Value() zcl.Val             { return a }
func (a YCoordinate) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *YCoordinate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = YCoordinate(*nt)
	return br, err
}

func (a *YCoordinate) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = YCoordinate(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a YCoordinate) String() string {
	return zcl.Meters.Format(float64(a) / 10)
}

type ZCoordinate zcl.Zs16

const ZCoordinateAttr zcl.AttrID = 18

func (ZCoordinate) ID() zcl.AttrID   { return ZCoordinateAttr }
func (ZCoordinate) Readable() bool   { return true }
func (ZCoordinate) Writable() bool   { return true }
func (ZCoordinate) Reportable() bool { return false }
func (ZCoordinate) SceneIndex() int  { return -1 }

func (ZCoordinate) Name() string                  { return "Z Coordinate" }
func (a *ZCoordinate) TypeID() zcl.TypeID         { return zcl.Zs16(*a).ID() }
func (a *ZCoordinate) Value() zcl.Val             { return a }
func (a ZCoordinate) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *ZCoordinate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ZCoordinate(*nt)
	return br, err
}

func (a *ZCoordinate) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = ZCoordinate(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ZCoordinate) String() string {
	return zcl.Meters.Format(float64(a) / 10)
}

type ZclVersion zcl.Zu8

const ZclVersionAttr zcl.AttrID = 0

func (ZclVersion) ID() zcl.AttrID   { return ZclVersionAttr }
func (ZclVersion) Readable() bool   { return true }
func (ZclVersion) Writable() bool   { return false }
func (ZclVersion) Reportable() bool { return false }
func (ZclVersion) SceneIndex() int  { return -1 }

func (ZclVersion) Name() string                  { return "ZCL Version" }
func (a *ZclVersion) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *ZclVersion) Value() zcl.Val             { return a }
func (a ZclVersion) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *ZclVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ZclVersion(*nt)
	return br, err
}

func (a *ZclVersion) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = ZclVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ZclVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}
