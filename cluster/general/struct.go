// provides general/common attributes and operations for most zigbee devices
package general

import "hemtjan.st/zcl"

type CommandID = zcl.CommandID
type Frame = zcl.ReceivedZclFrame

const ApsDecryptFailuresAttr zcl.AttrID = 278

func (ApsDecryptFailures) ID() zcl.AttrID   { return ApsDecryptFailuresAttr }
func (ApsDecryptFailures) Readable() bool   { return true }
func (ApsDecryptFailures) Writable() bool   { return false }
func (ApsDecryptFailures) Reportable() bool { return false }
func (ApsDecryptFailures) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ApsDecryptFailures) AttrID() zcl.AttrID   { return v.ID() }
func (v ApsDecryptFailures) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ApsDecryptFailures) AttrValue() zcl.Val  { return v.Value() }

func (ApsDecryptFailures) Name() string        { return `APS Decrypt Failures` }
func (ApsDecryptFailures) Description() string { return `` }

type ApsDecryptFailures zcl.Zu16

func (v *ApsDecryptFailures) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ApsDecryptFailures) Value() zcl.Val     { return v }

func (v ApsDecryptFailures) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ApsDecryptFailures) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ApsDecryptFailures(*nt)
	return br, err
}

func (v ApsDecryptFailures) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ApsDecryptFailures) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ApsDecryptFailures(*a)
	return nil
}

func (v *ApsDecryptFailures) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ApsDecryptFailures(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ApsDecryptFailures) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ApsFcFailureAttr zcl.AttrID = 275

func (ApsFcFailure) ID() zcl.AttrID   { return ApsFcFailureAttr }
func (ApsFcFailure) Readable() bool   { return true }
func (ApsFcFailure) Writable() bool   { return false }
func (ApsFcFailure) Reportable() bool { return false }
func (ApsFcFailure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ApsFcFailure) AttrID() zcl.AttrID   { return v.ID() }
func (v ApsFcFailure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ApsFcFailure) AttrValue() zcl.Val  { return v.Value() }

func (ApsFcFailure) Name() string        { return `APS FC Failure` }
func (ApsFcFailure) Description() string { return `` }

type ApsFcFailure zcl.Zu16

func (v *ApsFcFailure) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ApsFcFailure) Value() zcl.Val     { return v }

func (v ApsFcFailure) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ApsFcFailure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ApsFcFailure(*nt)
	return br, err
}

func (v ApsFcFailure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ApsFcFailure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ApsFcFailure(*a)
	return nil
}

func (v *ApsFcFailure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ApsFcFailure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ApsFcFailure) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ApsRxBcastAttr zcl.AttrID = 262

func (ApsRxBcast) ID() zcl.AttrID   { return ApsRxBcastAttr }
func (ApsRxBcast) Readable() bool   { return true }
func (ApsRxBcast) Writable() bool   { return false }
func (ApsRxBcast) Reportable() bool { return false }
func (ApsRxBcast) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ApsRxBcast) AttrID() zcl.AttrID   { return v.ID() }
func (v ApsRxBcast) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ApsRxBcast) AttrValue() zcl.Val  { return v.Value() }

func (ApsRxBcast) Name() string        { return `APS Rx Bcast` }
func (ApsRxBcast) Description() string { return `` }

type ApsRxBcast zcl.Zu16

func (v *ApsRxBcast) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ApsRxBcast) Value() zcl.Val     { return v }

func (v ApsRxBcast) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ApsRxBcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ApsRxBcast(*nt)
	return br, err
}

func (v ApsRxBcast) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ApsRxBcast) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ApsRxBcast(*a)
	return nil
}

func (v *ApsRxBcast) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ApsRxBcast(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ApsRxBcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ApsRxUcastAttr zcl.AttrID = 264

func (ApsRxUcast) ID() zcl.AttrID   { return ApsRxUcastAttr }
func (ApsRxUcast) Readable() bool   { return true }
func (ApsRxUcast) Writable() bool   { return false }
func (ApsRxUcast) Reportable() bool { return false }
func (ApsRxUcast) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ApsRxUcast) AttrID() zcl.AttrID   { return v.ID() }
func (v ApsRxUcast) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ApsRxUcast) AttrValue() zcl.Val  { return v.Value() }

func (ApsRxUcast) Name() string        { return `APS Rx Ucast` }
func (ApsRxUcast) Description() string { return `` }

type ApsRxUcast zcl.Zu16

func (v *ApsRxUcast) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ApsRxUcast) Value() zcl.Val     { return v }

func (v ApsRxUcast) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ApsRxUcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ApsRxUcast(*nt)
	return br, err
}

func (v ApsRxUcast) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ApsRxUcast) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ApsRxUcast(*a)
	return nil
}

func (v *ApsRxUcast) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ApsRxUcast(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ApsRxUcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ApsTxBcastAttr zcl.AttrID = 263

func (ApsTxBcast) ID() zcl.AttrID   { return ApsTxBcastAttr }
func (ApsTxBcast) Readable() bool   { return true }
func (ApsTxBcast) Writable() bool   { return false }
func (ApsTxBcast) Reportable() bool { return false }
func (ApsTxBcast) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ApsTxBcast) AttrID() zcl.AttrID   { return v.ID() }
func (v ApsTxBcast) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ApsTxBcast) AttrValue() zcl.Val  { return v.Value() }

func (ApsTxBcast) Name() string        { return `APS Tx Bcast` }
func (ApsTxBcast) Description() string { return `` }

type ApsTxBcast zcl.Zu16

func (v *ApsTxBcast) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ApsTxBcast) Value() zcl.Val     { return v }

func (v ApsTxBcast) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ApsTxBcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ApsTxBcast(*nt)
	return br, err
}

func (v ApsTxBcast) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ApsTxBcast) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ApsTxBcast(*a)
	return nil
}

func (v *ApsTxBcast) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ApsTxBcast(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ApsTxBcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ApsTxUcastFailAttr zcl.AttrID = 267

func (ApsTxUcastFail) ID() zcl.AttrID   { return ApsTxUcastFailAttr }
func (ApsTxUcastFail) Readable() bool   { return true }
func (ApsTxUcastFail) Writable() bool   { return false }
func (ApsTxUcastFail) Reportable() bool { return false }
func (ApsTxUcastFail) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ApsTxUcastFail) AttrID() zcl.AttrID   { return v.ID() }
func (v ApsTxUcastFail) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ApsTxUcastFail) AttrValue() zcl.Val  { return v.Value() }

func (ApsTxUcastFail) Name() string        { return `APS Tx Ucast Fail` }
func (ApsTxUcastFail) Description() string { return `` }

type ApsTxUcastFail zcl.Zu16

func (v *ApsTxUcastFail) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ApsTxUcastFail) Value() zcl.Val     { return v }

func (v ApsTxUcastFail) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ApsTxUcastFail) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ApsTxUcastFail(*nt)
	return br, err
}

func (v ApsTxUcastFail) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ApsTxUcastFail) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ApsTxUcastFail(*a)
	return nil
}

func (v *ApsTxUcastFail) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ApsTxUcastFail(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ApsTxUcastFail) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ApsTxUcastRetryAttr zcl.AttrID = 266

func (ApsTxUcastRetry) ID() zcl.AttrID   { return ApsTxUcastRetryAttr }
func (ApsTxUcastRetry) Readable() bool   { return true }
func (ApsTxUcastRetry) Writable() bool   { return false }
func (ApsTxUcastRetry) Reportable() bool { return false }
func (ApsTxUcastRetry) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ApsTxUcastRetry) AttrID() zcl.AttrID   { return v.ID() }
func (v ApsTxUcastRetry) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ApsTxUcastRetry) AttrValue() zcl.Val  { return v.Value() }

func (ApsTxUcastRetry) Name() string        { return `APS Tx Ucast Retry` }
func (ApsTxUcastRetry) Description() string { return `` }

type ApsTxUcastRetry zcl.Zu16

func (v *ApsTxUcastRetry) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ApsTxUcastRetry) Value() zcl.Val     { return v }

func (v ApsTxUcastRetry) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ApsTxUcastRetry) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ApsTxUcastRetry(*nt)
	return br, err
}

func (v ApsTxUcastRetry) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ApsTxUcastRetry) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ApsTxUcastRetry(*a)
	return nil
}

func (v *ApsTxUcastRetry) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ApsTxUcastRetry(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ApsTxUcastRetry) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ApsTxUcastSuccessAttr zcl.AttrID = 265

func (ApsTxUcastSuccess) ID() zcl.AttrID   { return ApsTxUcastSuccessAttr }
func (ApsTxUcastSuccess) Readable() bool   { return true }
func (ApsTxUcastSuccess) Writable() bool   { return false }
func (ApsTxUcastSuccess) Reportable() bool { return false }
func (ApsTxUcastSuccess) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ApsTxUcastSuccess) AttrID() zcl.AttrID   { return v.ID() }
func (v ApsTxUcastSuccess) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ApsTxUcastSuccess) AttrValue() zcl.Val  { return v.Value() }

func (ApsTxUcastSuccess) Name() string        { return `APS Tx Ucast Success` }
func (ApsTxUcastSuccess) Description() string { return `` }

type ApsTxUcastSuccess zcl.Zu16

func (v *ApsTxUcastSuccess) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ApsTxUcastSuccess) Value() zcl.Val     { return v }

func (v ApsTxUcastSuccess) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ApsTxUcastSuccess) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ApsTxUcastSuccess(*nt)
	return br, err
}

func (v ApsTxUcastSuccess) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ApsTxUcastSuccess) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ApsTxUcastSuccess(*a)
	return nil
}

func (v *ApsTxUcastSuccess) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ApsTxUcastSuccess(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ApsTxUcastSuccess) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ApsUnauthorizedKeyAttr zcl.AttrID = 276

func (ApsUnauthorizedKey) ID() zcl.AttrID   { return ApsUnauthorizedKeyAttr }
func (ApsUnauthorizedKey) Readable() bool   { return true }
func (ApsUnauthorizedKey) Writable() bool   { return false }
func (ApsUnauthorizedKey) Reportable() bool { return false }
func (ApsUnauthorizedKey) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ApsUnauthorizedKey) AttrID() zcl.AttrID   { return v.ID() }
func (v ApsUnauthorizedKey) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ApsUnauthorizedKey) AttrValue() zcl.Val  { return v.Value() }

func (ApsUnauthorizedKey) Name() string        { return `APS Unauthorized Key` }
func (ApsUnauthorizedKey) Description() string { return `` }

type ApsUnauthorizedKey zcl.Zu16

func (v *ApsUnauthorizedKey) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ApsUnauthorizedKey) Value() zcl.Val     { return v }

func (v ApsUnauthorizedKey) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ApsUnauthorizedKey) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ApsUnauthorizedKey(*nt)
	return br, err
}

func (v ApsUnauthorizedKey) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ApsUnauthorizedKey) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ApsUnauthorizedKey(*a)
	return nil
}

func (v *ApsUnauthorizedKey) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ApsUnauthorizedKey(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ApsUnauthorizedKey) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AlarmCountAttr zcl.AttrID = 0

func (AlarmCount) ID() zcl.AttrID   { return AlarmCountAttr }
func (AlarmCount) Readable() bool   { return true }
func (AlarmCount) Writable() bool   { return true }
func (AlarmCount) Reportable() bool { return false }
func (AlarmCount) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AlarmCount) AttrID() zcl.AttrID   { return v.ID() }
func (v AlarmCount) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AlarmCount) AttrValue() zcl.Val  { return v.Value() }

func (AlarmCount) Name() string        { return `Alarm Count` }
func (AlarmCount) Description() string { return `Number of alarms currently defined` }

// AlarmCount Number of alarms currently defined
type AlarmCount zcl.Zu16

func (v *AlarmCount) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AlarmCount) Value() zcl.Val     { return v }

func (v AlarmCount) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *AlarmCount) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AlarmCount(*nt)
	return br, err
}

func (v AlarmCount) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AlarmCount) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AlarmCount(*a)
	return nil
}

func (v *AlarmCount) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AlarmCount(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AlarmCount) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AlarmMaskAttr zcl.AttrID = 19

func (AlarmMask) ID() zcl.AttrID   { return AlarmMaskAttr }
func (AlarmMask) Readable() bool   { return true }
func (AlarmMask) Writable() bool   { return true }
func (AlarmMask) Reportable() bool { return false }
func (AlarmMask) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AlarmMask) AttrID() zcl.AttrID   { return v.ID() }
func (v AlarmMask) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AlarmMask) AttrValue() zcl.Val  { return v.Value() }

func (AlarmMask) Name() string        { return `Alarm Mask` }
func (AlarmMask) Description() string { return `` }

type AlarmMask zcl.Zbmp8

func (v *AlarmMask) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *AlarmMask) Value() zcl.Val     { return v }

func (v AlarmMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *AlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = AlarmMask(*nt)
	return br, err
}

func (v AlarmMask) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *AlarmMask) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AlarmMask(*a)
	return nil
}

func (v *AlarmMask) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = AlarmMask(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AlarmMask) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v AlarmMask) IsGeneralHardwareFault() bool { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v AlarmMask) IsGeneralSoftwareFault() bool { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v *AlarmMask) SetGeneralHardwareFault(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *AlarmMask) SetGeneralSoftwareFault(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}

func (AlarmMask) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "General Hardware Fault"},
		{Value: 1, Name: "General Software Fault"},
	}
}

func (AlarmCode) Name() string        { return `Alarm code` }
func (AlarmCode) Description() string { return `` }

type AlarmCode zcl.Zenum8

func (v *AlarmCode) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *AlarmCode) Value() zcl.Val     { return v }

func (v AlarmCode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *AlarmCode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = AlarmCode(*nt)
	return br, err
}

func (v AlarmCode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *AlarmCode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AlarmCode(*a)
	return nil
}

func (v *AlarmCode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = AlarmCode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AlarmCode) String() string {
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

const AnalogMaxPresentValueAttr zcl.AttrID = 65

func (AnalogMaxPresentValue) ID() zcl.AttrID   { return AnalogMaxPresentValueAttr }
func (AnalogMaxPresentValue) Readable() bool   { return true }
func (AnalogMaxPresentValue) Writable() bool   { return true }
func (AnalogMaxPresentValue) Reportable() bool { return false }
func (AnalogMaxPresentValue) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AnalogMaxPresentValue) AttrID() zcl.AttrID   { return v.ID() }
func (v AnalogMaxPresentValue) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AnalogMaxPresentValue) AttrValue() zcl.Val  { return v.Value() }

func (AnalogMaxPresentValue) Name() string        { return `Analog Max Present Value` }
func (AnalogMaxPresentValue) Description() string { return `` }

type AnalogMaxPresentValue zcl.Zfloat

func (v *AnalogMaxPresentValue) TypeID() zcl.TypeID { return new(zcl.Zfloat).TypeID() }
func (v *AnalogMaxPresentValue) Value() zcl.Val     { return v }

func (v AnalogMaxPresentValue) MarshalZcl() ([]byte, error) { return zcl.Zfloat(v).MarshalZcl() }

func (v *AnalogMaxPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*v = AnalogMaxPresentValue(*nt)
	return br, err
}

func (v AnalogMaxPresentValue) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zfloat(v))
}

func (v *AnalogMaxPresentValue) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zfloat)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AnalogMaxPresentValue(*a)
	return nil
}

func (v *AnalogMaxPresentValue) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zfloat); ok {
		*v = AnalogMaxPresentValue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AnalogMaxPresentValue) String() string {
	return zcl.Sprintf("%v", zcl.Zfloat(v))
}

const AnalogMinPresentValueAttr zcl.AttrID = 69

func (AnalogMinPresentValue) ID() zcl.AttrID   { return AnalogMinPresentValueAttr }
func (AnalogMinPresentValue) Readable() bool   { return true }
func (AnalogMinPresentValue) Writable() bool   { return true }
func (AnalogMinPresentValue) Reportable() bool { return false }
func (AnalogMinPresentValue) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AnalogMinPresentValue) AttrID() zcl.AttrID   { return v.ID() }
func (v AnalogMinPresentValue) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AnalogMinPresentValue) AttrValue() zcl.Val  { return v.Value() }

func (AnalogMinPresentValue) Name() string        { return `Analog Min Present Value` }
func (AnalogMinPresentValue) Description() string { return `` }

type AnalogMinPresentValue zcl.Zfloat

func (v *AnalogMinPresentValue) TypeID() zcl.TypeID { return new(zcl.Zfloat).TypeID() }
func (v *AnalogMinPresentValue) Value() zcl.Val     { return v }

func (v AnalogMinPresentValue) MarshalZcl() ([]byte, error) { return zcl.Zfloat(v).MarshalZcl() }

func (v *AnalogMinPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*v = AnalogMinPresentValue(*nt)
	return br, err
}

func (v AnalogMinPresentValue) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zfloat(v))
}

func (v *AnalogMinPresentValue) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zfloat)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AnalogMinPresentValue(*a)
	return nil
}

func (v *AnalogMinPresentValue) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zfloat); ok {
		*v = AnalogMinPresentValue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AnalogMinPresentValue) String() string {
	return zcl.Sprintf("%v", zcl.Zfloat(v))
}

const AnalogPresentValueAttr zcl.AttrID = 85

func (AnalogPresentValue) ID() zcl.AttrID   { return AnalogPresentValueAttr }
func (AnalogPresentValue) Readable() bool   { return true }
func (AnalogPresentValue) Writable() bool   { return true }
func (AnalogPresentValue) Reportable() bool { return true }
func (AnalogPresentValue) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AnalogPresentValue) AttrID() zcl.AttrID   { return v.ID() }
func (v AnalogPresentValue) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AnalogPresentValue) AttrValue() zcl.Val  { return v.Value() }

func (AnalogPresentValue) Name() string        { return `Analog Present value` }
func (AnalogPresentValue) Description() string { return `` }

type AnalogPresentValue zcl.Zfloat

func (v *AnalogPresentValue) TypeID() zcl.TypeID { return new(zcl.Zfloat).TypeID() }
func (v *AnalogPresentValue) Value() zcl.Val     { return v }

func (v AnalogPresentValue) MarshalZcl() ([]byte, error) { return zcl.Zfloat(v).MarshalZcl() }

func (v *AnalogPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*v = AnalogPresentValue(*nt)
	return br, err
}

func (v AnalogPresentValue) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zfloat(v))
}

func (v *AnalogPresentValue) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zfloat)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AnalogPresentValue(*a)
	return nil
}

func (v *AnalogPresentValue) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zfloat); ok {
		*v = AnalogPresentValue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AnalogPresentValue) String() string {
	return zcl.Sprintf("%v", zcl.Zfloat(v))
}

func (AnalogPriority) Name() string        { return `Analog Priority` }
func (AnalogPriority) Description() string { return `` }

type AnalogPriority struct {
	Ispriority    zcl.Zbool
	Priorityvalue zcl.Zfloat
}

func (v *AnalogPriority) TypeID() zcl.TypeID { return zcl.TypeID(76) } // struct
func (v *AnalogPriority) Value() zcl.Val     { return v }
func (v AnalogPriority) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	var err error
	_ = tmp2
	_ = err

	{
		if tmp, err = v.Ispriority.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Priorityvalue.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (v *AnalogPriority) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2
	_ = err

	if b, err = new(zcl.Zbool).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = new(zcl.Zfloat).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v *AnalogPriority) SetValue(a zcl.Val) error {
	if nv, ok := a.(*AnalogPriority); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrNotImpl
}

func (v *AnalogPriority) String() string {
	return zcl.Sprintf(
		"AnalogPriority{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

const AnalogPriorityArrayAttr zcl.AttrID = 87

func (AnalogPriorityArray) ID() zcl.AttrID   { return AnalogPriorityArrayAttr }
func (AnalogPriorityArray) Readable() bool   { return true }
func (AnalogPriorityArray) Writable() bool   { return true }
func (AnalogPriorityArray) Reportable() bool { return false }
func (AnalogPriorityArray) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AnalogPriorityArray) AttrID() zcl.AttrID   { return v.ID() }
func (v AnalogPriorityArray) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AnalogPriorityArray) AttrValue() zcl.Val  { return v.Value() }

func (AnalogPriorityArray) Name() string        { return `Analog Priority Array` }
func (AnalogPriorityArray) Description() string { return `` }

type AnalogPriorityArray []*AnalogPriority

func (v *AnalogPriorityArray) TypeID() zcl.TypeID { return new(zcl.Zarray).TypeID() }
func (v *AnalogPriorityArray) Value() zcl.Val     { return v }

func (AnalogPriorityArray) ArrayTypeID() zcl.TypeID { return new(AnalogPriority).TypeID() }

func (v *AnalogPriorityArray) ArrayValues() (o []AnalogPriority) {
	for _, a := range *v {
		o = append(o, *a)
	}
	return o
}

func (v *AnalogPriorityArray) SetValues(val []AnalogPriority) error {
	*v = []*AnalogPriority{}
	return v.AddValues(val...)
}

func (v *AnalogPriorityArray) AddValues(val ...AnalogPriority) error {
	for _, a := range val {
		nv := a
		*v = append(*v, &nv)
	}
	return nil
}

func (v AnalogPriorityArray) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, a := range v {
		vars = append(vars, a)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (v *AnalogPriorityArray) UnmarshalZcl(b []byte) ([]byte, error) {
	*v = []*AnalogPriority{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(AnalogPriority)
		*v = append(*v, nv)
		return nv
	})
}

func (v *AnalogPriorityArray) SetValue(a zcl.Val) error {
	if nv, ok := a.(*AnalogPriorityArray); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AnalogPriorityArray) String() string {
	var s []string
	for _, a := range v {
		s = append(s, zcl.Sprintf("%v", a))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

const AnalogRelinquishDefaultAttr zcl.AttrID = 104

func (AnalogRelinquishDefault) ID() zcl.AttrID   { return AnalogRelinquishDefaultAttr }
func (AnalogRelinquishDefault) Readable() bool   { return true }
func (AnalogRelinquishDefault) Writable() bool   { return true }
func (AnalogRelinquishDefault) Reportable() bool { return false }
func (AnalogRelinquishDefault) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AnalogRelinquishDefault) AttrID() zcl.AttrID   { return v.ID() }
func (v AnalogRelinquishDefault) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AnalogRelinquishDefault) AttrValue() zcl.Val  { return v.Value() }

func (AnalogRelinquishDefault) Name() string        { return `Analog Relinquish Default` }
func (AnalogRelinquishDefault) Description() string { return `` }

type AnalogRelinquishDefault zcl.Zfloat

func (v *AnalogRelinquishDefault) TypeID() zcl.TypeID { return new(zcl.Zfloat).TypeID() }
func (v *AnalogRelinquishDefault) Value() zcl.Val     { return v }

func (v AnalogRelinquishDefault) MarshalZcl() ([]byte, error) { return zcl.Zfloat(v).MarshalZcl() }

func (v *AnalogRelinquishDefault) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*v = AnalogRelinquishDefault(*nt)
	return br, err
}

func (v AnalogRelinquishDefault) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zfloat(v))
}

func (v *AnalogRelinquishDefault) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zfloat)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AnalogRelinquishDefault(*a)
	return nil
}

func (v *AnalogRelinquishDefault) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zfloat); ok {
		*v = AnalogRelinquishDefault(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AnalogRelinquishDefault) String() string {
	return zcl.Sprintf("%v", zcl.Zfloat(v))
}

const AnalogResolutionAttr zcl.AttrID = 106

func (AnalogResolution) ID() zcl.AttrID   { return AnalogResolutionAttr }
func (AnalogResolution) Readable() bool   { return true }
func (AnalogResolution) Writable() bool   { return true }
func (AnalogResolution) Reportable() bool { return false }
func (AnalogResolution) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AnalogResolution) AttrID() zcl.AttrID   { return v.ID() }
func (v AnalogResolution) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AnalogResolution) AttrValue() zcl.Val  { return v.Value() }

func (AnalogResolution) Name() string        { return `Analog Resolution` }
func (AnalogResolution) Description() string { return `` }

type AnalogResolution zcl.Zfloat

func (v *AnalogResolution) TypeID() zcl.TypeID { return new(zcl.Zfloat).TypeID() }
func (v *AnalogResolution) Value() zcl.Val     { return v }

func (v AnalogResolution) MarshalZcl() ([]byte, error) { return zcl.Zfloat(v).MarshalZcl() }

func (v *AnalogResolution) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*v = AnalogResolution(*nt)
	return br, err
}

func (v AnalogResolution) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zfloat(v))
}

func (v *AnalogResolution) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zfloat)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AnalogResolution(*a)
	return nil
}

func (v *AnalogResolution) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zfloat); ok {
		*v = AnalogResolution(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AnalogResolution) String() string {
	return zcl.Sprintf("%v", zcl.Zfloat(v))
}

const ApplicationVersionAttr zcl.AttrID = 1

func (ApplicationVersion) ID() zcl.AttrID   { return ApplicationVersionAttr }
func (ApplicationVersion) Readable() bool   { return true }
func (ApplicationVersion) Writable() bool   { return false }
func (ApplicationVersion) Reportable() bool { return false }
func (ApplicationVersion) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ApplicationVersion) AttrID() zcl.AttrID   { return v.ID() }
func (v ApplicationVersion) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ApplicationVersion) AttrValue() zcl.Val  { return v.Value() }

func (ApplicationVersion) Name() string        { return `Application Version` }
func (ApplicationVersion) Description() string { return `` }

type ApplicationVersion zcl.Zu8

func (v *ApplicationVersion) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *ApplicationVersion) Value() zcl.Val     { return v }

func (v ApplicationVersion) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *ApplicationVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = ApplicationVersion(*nt)
	return br, err
}

func (v ApplicationVersion) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *ApplicationVersion) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ApplicationVersion(*a)
	return nil
}

func (v *ApplicationVersion) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = ApplicationVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ApplicationVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const AvgMacRetryPerApsMsgSentAttr zcl.AttrID = 283

func (AvgMacRetryPerApsMsgSent) ID() zcl.AttrID   { return AvgMacRetryPerApsMsgSentAttr }
func (AvgMacRetryPerApsMsgSent) Readable() bool   { return true }
func (AvgMacRetryPerApsMsgSent) Writable() bool   { return false }
func (AvgMacRetryPerApsMsgSent) Reportable() bool { return false }
func (AvgMacRetryPerApsMsgSent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AvgMacRetryPerApsMsgSent) AttrID() zcl.AttrID   { return v.ID() }
func (v AvgMacRetryPerApsMsgSent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AvgMacRetryPerApsMsgSent) AttrValue() zcl.Val  { return v.Value() }

func (AvgMacRetryPerApsMsgSent) Name() string        { return `Avg MAC Retry per APS Msg Sent` }
func (AvgMacRetryPerApsMsgSent) Description() string { return `` }

type AvgMacRetryPerApsMsgSent zcl.Zu16

func (v *AvgMacRetryPerApsMsgSent) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AvgMacRetryPerApsMsgSent) Value() zcl.Val     { return v }

func (v AvgMacRetryPerApsMsgSent) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *AvgMacRetryPerApsMsgSent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AvgMacRetryPerApsMsgSent(*nt)
	return br, err
}

func (v AvgMacRetryPerApsMsgSent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AvgMacRetryPerApsMsgSent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AvgMacRetryPerApsMsgSent(*a)
	return nil
}

func (v *AvgMacRetryPerApsMsgSent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AvgMacRetryPerApsMsgSent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AvgMacRetryPerApsMsgSent) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const BatteryAlarmMaskAttr zcl.AttrID = 53

func (BatteryAlarmMask) ID() zcl.AttrID   { return BatteryAlarmMaskAttr }
func (BatteryAlarmMask) Readable() bool   { return true }
func (BatteryAlarmMask) Writable() bool   { return true }
func (BatteryAlarmMask) Reportable() bool { return false }
func (BatteryAlarmMask) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BatteryAlarmMask) AttrID() zcl.AttrID   { return v.ID() }
func (v BatteryAlarmMask) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BatteryAlarmMask) AttrValue() zcl.Val  { return v.Value() }

func (BatteryAlarmMask) Name() string        { return `Battery Alarm Mask` }
func (BatteryAlarmMask) Description() string { return `` }

type BatteryAlarmMask zcl.Zbmp8

func (v *BatteryAlarmMask) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *BatteryAlarmMask) Value() zcl.Val     { return v }

func (v BatteryAlarmMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *BatteryAlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = BatteryAlarmMask(*nt)
	return br, err
}

func (v BatteryAlarmMask) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *BatteryAlarmMask) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BatteryAlarmMask(*a)
	return nil
}

func (v *BatteryAlarmMask) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = BatteryAlarmMask(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BatteryAlarmMask) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v BatteryAlarmMask) IsBatteryVoltageTooLow() bool { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v BatteryAlarmMask) IsBatteryAlarm1() bool        { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v BatteryAlarmMask) IsBatteryAlarm2() bool        { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v BatteryAlarmMask) IsBatteryAlarm3() bool        { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v *BatteryAlarmMask) SetBatteryVoltageTooLow(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *BatteryAlarmMask) SetBatteryAlarm1(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *BatteryAlarmMask) SetBatteryAlarm2(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}
func (v *BatteryAlarmMask) SetBatteryAlarm3(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b))
}

func (BatteryAlarmMask) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Battery Voltage too low"},
		{Value: 1, Name: "Battery Alarm 1"},
		{Value: 2, Name: "Battery Alarm 2"},
		{Value: 3, Name: "Battery Alarm 3"},
	}
}

const BatteryAlarmStateAttr zcl.AttrID = 62

func (BatteryAlarmState) ID() zcl.AttrID   { return BatteryAlarmStateAttr }
func (BatteryAlarmState) Readable() bool   { return true }
func (BatteryAlarmState) Writable() bool   { return true }
func (BatteryAlarmState) Reportable() bool { return false }
func (BatteryAlarmState) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BatteryAlarmState) AttrID() zcl.AttrID   { return v.ID() }
func (v BatteryAlarmState) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BatteryAlarmState) AttrValue() zcl.Val  { return v.Value() }

func (BatteryAlarmState) Name() string        { return `Battery Alarm State` }
func (BatteryAlarmState) Description() string { return `` }

type BatteryAlarmState zcl.Zbmp32

func (v *BatteryAlarmState) TypeID() zcl.TypeID { return new(zcl.Zbmp32).TypeID() }
func (v *BatteryAlarmState) Value() zcl.Val     { return v }

func (v BatteryAlarmState) MarshalZcl() ([]byte, error) { return zcl.Zbmp32(v).MarshalZcl() }

func (v *BatteryAlarmState) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp32)
	br, err := nt.UnmarshalZcl(b)
	*v = BatteryAlarmState(*nt)
	return br, err
}

func (v BatteryAlarmState) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp32(v))
}

func (v *BatteryAlarmState) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BatteryAlarmState(*a)
	return nil
}

func (v *BatteryAlarmState) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp32); ok {
		*v = BatteryAlarmState(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BatteryAlarmState) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v BatteryAlarmState) IsMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource1() bool {
	return zcl.BitmapTest([]byte(v[:]), 0)
}
func (v BatteryAlarmState) IsThreshold1ForVoltageOrPercentageReachedForBatterySource1() bool {
	return zcl.BitmapTest([]byte(v[:]), 1)
}
func (v BatteryAlarmState) IsMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource2() bool {
	return zcl.BitmapTest([]byte(v[:]), 10)
}
func (v BatteryAlarmState) IsThreshold1ForVoltageOrPercentageReachedForBatterySource2() bool {
	return zcl.BitmapTest([]byte(v[:]), 11)
}
func (v BatteryAlarmState) IsThreshold2ForVoltageOrPercentageReachedForBatterySource2() bool {
	return zcl.BitmapTest([]byte(v[:]), 12)
}
func (v BatteryAlarmState) IsThreshold3ForVoltageOrPercentageReachedForBatterySource2() bool {
	return zcl.BitmapTest([]byte(v[:]), 13)
}
func (v BatteryAlarmState) IsThreshold2ForVoltageOrPercentageReachedForBatterySource1() bool {
	return zcl.BitmapTest([]byte(v[:]), 2)
}
func (v BatteryAlarmState) IsMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource3() bool {
	return zcl.BitmapTest([]byte(v[:]), 20)
}
func (v BatteryAlarmState) IsThreshold1ForVoltageOrPercentageReachedForBatterySource3() bool {
	return zcl.BitmapTest([]byte(v[:]), 21)
}
func (v BatteryAlarmState) IsThreshold2ForVoltageOrPercentageReachedForBatterySource3() bool {
	return zcl.BitmapTest([]byte(v[:]), 22)
}
func (v BatteryAlarmState) IsThreshold3ForVoltageOrPercentageReachedForBatterySource3() bool {
	return zcl.BitmapTest([]byte(v[:]), 23)
}
func (v BatteryAlarmState) IsThreshold3ForVoltageOrPercentageReachedForBatterySource1() bool {
	return zcl.BitmapTest([]byte(v[:]), 3)
}
func (v BatteryAlarmState) IsMainsPowerLostUnavailable() bool { return zcl.BitmapTest([]byte(v[:]), 30) }
func (v *BatteryAlarmState) SetMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource1(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *BatteryAlarmState) SetThreshold1ForVoltageOrPercentageReachedForBatterySource1(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *BatteryAlarmState) SetMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource2(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 10, b))
}
func (v *BatteryAlarmState) SetThreshold1ForVoltageOrPercentageReachedForBatterySource2(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 11, b))
}
func (v *BatteryAlarmState) SetThreshold2ForVoltageOrPercentageReachedForBatterySource2(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 12, b))
}
func (v *BatteryAlarmState) SetThreshold3ForVoltageOrPercentageReachedForBatterySource2(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 13, b))
}
func (v *BatteryAlarmState) SetThreshold2ForVoltageOrPercentageReachedForBatterySource1(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}
func (v *BatteryAlarmState) SetMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource3(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 20, b))
}
func (v *BatteryAlarmState) SetThreshold1ForVoltageOrPercentageReachedForBatterySource3(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 21, b))
}
func (v *BatteryAlarmState) SetThreshold2ForVoltageOrPercentageReachedForBatterySource3(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 22, b))
}
func (v *BatteryAlarmState) SetThreshold3ForVoltageOrPercentageReachedForBatterySource3(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 23, b))
}
func (v *BatteryAlarmState) SetThreshold3ForVoltageOrPercentageReachedForBatterySource1(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b))
}
func (v *BatteryAlarmState) SetMainsPowerLostUnavailable(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 30, b))
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

const BatteryManufacturerAttr zcl.AttrID = 48

func (BatteryManufacturer) ID() zcl.AttrID   { return BatteryManufacturerAttr }
func (BatteryManufacturer) Readable() bool   { return true }
func (BatteryManufacturer) Writable() bool   { return true }
func (BatteryManufacturer) Reportable() bool { return false }
func (BatteryManufacturer) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BatteryManufacturer) AttrID() zcl.AttrID   { return v.ID() }
func (v BatteryManufacturer) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BatteryManufacturer) AttrValue() zcl.Val  { return v.Value() }

func (BatteryManufacturer) Name() string        { return `Battery Manufacturer` }
func (BatteryManufacturer) Description() string { return `` }

type BatteryManufacturer zcl.Zcstring

func (v *BatteryManufacturer) TypeID() zcl.TypeID { return new(zcl.Zcstring).TypeID() }
func (v *BatteryManufacturer) Value() zcl.Val     { return v }

func (v BatteryManufacturer) MarshalZcl() ([]byte, error) { return zcl.Zcstring(v).MarshalZcl() }

func (v *BatteryManufacturer) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*v = BatteryManufacturer(*nt)
	return br, err
}

func (v BatteryManufacturer) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zcstring(v))
}

func (v *BatteryManufacturer) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zcstring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BatteryManufacturer(*a)
	return nil
}

func (v *BatteryManufacturer) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zcstring); ok {
		*v = BatteryManufacturer(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BatteryManufacturer) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v))
}

const BatteryPercentageMinThresholdAttr zcl.AttrID = 58

func (BatteryPercentageMinThreshold) ID() zcl.AttrID   { return BatteryPercentageMinThresholdAttr }
func (BatteryPercentageMinThreshold) Readable() bool   { return true }
func (BatteryPercentageMinThreshold) Writable() bool   { return true }
func (BatteryPercentageMinThreshold) Reportable() bool { return false }
func (BatteryPercentageMinThreshold) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BatteryPercentageMinThreshold) AttrID() zcl.AttrID   { return v.ID() }
func (v BatteryPercentageMinThreshold) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BatteryPercentageMinThreshold) AttrValue() zcl.Val  { return v.Value() }

func (BatteryPercentageMinThreshold) Name() string        { return `Battery Percentage Min Threshold` }
func (BatteryPercentageMinThreshold) Description() string { return `` }

type BatteryPercentageMinThreshold zcl.Zu8

func (v *BatteryPercentageMinThreshold) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *BatteryPercentageMinThreshold) Value() zcl.Val     { return v }

func (v BatteryPercentageMinThreshold) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *BatteryPercentageMinThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = BatteryPercentageMinThreshold(*nt)
	return br, err
}

func (v BatteryPercentageMinThreshold) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *BatteryPercentageMinThreshold) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BatteryPercentageMinThreshold(*a)
	return nil
}

func (v *BatteryPercentageMinThreshold) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = BatteryPercentageMinThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BatteryPercentageMinThreshold) String() string {
	return zcl.Percent.Format(float64(v) / 2)
}

const BatteryPercentageThreshold1Attr zcl.AttrID = 59

func (BatteryPercentageThreshold1) ID() zcl.AttrID   { return BatteryPercentageThreshold1Attr }
func (BatteryPercentageThreshold1) Readable() bool   { return true }
func (BatteryPercentageThreshold1) Writable() bool   { return true }
func (BatteryPercentageThreshold1) Reportable() bool { return false }
func (BatteryPercentageThreshold1) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BatteryPercentageThreshold1) AttrID() zcl.AttrID   { return v.ID() }
func (v BatteryPercentageThreshold1) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BatteryPercentageThreshold1) AttrValue() zcl.Val  { return v.Value() }

func (BatteryPercentageThreshold1) Name() string        { return `Battery Percentage Threshold 1` }
func (BatteryPercentageThreshold1) Description() string { return `` }

type BatteryPercentageThreshold1 zcl.Zu8

func (v *BatteryPercentageThreshold1) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *BatteryPercentageThreshold1) Value() zcl.Val     { return v }

func (v BatteryPercentageThreshold1) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *BatteryPercentageThreshold1) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = BatteryPercentageThreshold1(*nt)
	return br, err
}

func (v BatteryPercentageThreshold1) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *BatteryPercentageThreshold1) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BatteryPercentageThreshold1(*a)
	return nil
}

func (v *BatteryPercentageThreshold1) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = BatteryPercentageThreshold1(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BatteryPercentageThreshold1) String() string {
	return zcl.Percent.Format(float64(v) / 2)
}

const BatteryPercentageThreshold2Attr zcl.AttrID = 60

func (BatteryPercentageThreshold2) ID() zcl.AttrID   { return BatteryPercentageThreshold2Attr }
func (BatteryPercentageThreshold2) Readable() bool   { return true }
func (BatteryPercentageThreshold2) Writable() bool   { return true }
func (BatteryPercentageThreshold2) Reportable() bool { return false }
func (BatteryPercentageThreshold2) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BatteryPercentageThreshold2) AttrID() zcl.AttrID   { return v.ID() }
func (v BatteryPercentageThreshold2) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BatteryPercentageThreshold2) AttrValue() zcl.Val  { return v.Value() }

func (BatteryPercentageThreshold2) Name() string        { return `Battery Percentage Threshold 2` }
func (BatteryPercentageThreshold2) Description() string { return `` }

type BatteryPercentageThreshold2 zcl.Zu8

func (v *BatteryPercentageThreshold2) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *BatteryPercentageThreshold2) Value() zcl.Val     { return v }

func (v BatteryPercentageThreshold2) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *BatteryPercentageThreshold2) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = BatteryPercentageThreshold2(*nt)
	return br, err
}

func (v BatteryPercentageThreshold2) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *BatteryPercentageThreshold2) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BatteryPercentageThreshold2(*a)
	return nil
}

func (v *BatteryPercentageThreshold2) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = BatteryPercentageThreshold2(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BatteryPercentageThreshold2) String() string {
	return zcl.Percent.Format(float64(v) / 2)
}

const BatteryPercentageThreshold3Attr zcl.AttrID = 61

func (BatteryPercentageThreshold3) ID() zcl.AttrID   { return BatteryPercentageThreshold3Attr }
func (BatteryPercentageThreshold3) Readable() bool   { return true }
func (BatteryPercentageThreshold3) Writable() bool   { return true }
func (BatteryPercentageThreshold3) Reportable() bool { return false }
func (BatteryPercentageThreshold3) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BatteryPercentageThreshold3) AttrID() zcl.AttrID   { return v.ID() }
func (v BatteryPercentageThreshold3) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BatteryPercentageThreshold3) AttrValue() zcl.Val  { return v.Value() }

func (BatteryPercentageThreshold3) Name() string        { return `Battery Percentage Threshold 3` }
func (BatteryPercentageThreshold3) Description() string { return `` }

type BatteryPercentageThreshold3 zcl.Zu8

func (v *BatteryPercentageThreshold3) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *BatteryPercentageThreshold3) Value() zcl.Val     { return v }

func (v BatteryPercentageThreshold3) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *BatteryPercentageThreshold3) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = BatteryPercentageThreshold3(*nt)
	return br, err
}

func (v BatteryPercentageThreshold3) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *BatteryPercentageThreshold3) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BatteryPercentageThreshold3(*a)
	return nil
}

func (v *BatteryPercentageThreshold3) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = BatteryPercentageThreshold3(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BatteryPercentageThreshold3) String() string {
	return zcl.Percent.Format(float64(v) / 2)
}

const BatteryQuantityAttr zcl.AttrID = 51

func (BatteryQuantity) ID() zcl.AttrID   { return BatteryQuantityAttr }
func (BatteryQuantity) Readable() bool   { return true }
func (BatteryQuantity) Writable() bool   { return true }
func (BatteryQuantity) Reportable() bool { return false }
func (BatteryQuantity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BatteryQuantity) AttrID() zcl.AttrID   { return v.ID() }
func (v BatteryQuantity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BatteryQuantity) AttrValue() zcl.Val  { return v.Value() }

func (BatteryQuantity) Name() string        { return `Battery Quantity` }
func (BatteryQuantity) Description() string { return `` }

type BatteryQuantity zcl.Zu8

func (v *BatteryQuantity) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *BatteryQuantity) Value() zcl.Val     { return v }

func (v BatteryQuantity) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *BatteryQuantity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = BatteryQuantity(*nt)
	return br, err
}

func (v BatteryQuantity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *BatteryQuantity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BatteryQuantity(*a)
	return nil
}

func (v *BatteryQuantity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = BatteryQuantity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BatteryQuantity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const BatteryRatedVoltageAttr zcl.AttrID = 52

func (BatteryRatedVoltage) ID() zcl.AttrID   { return BatteryRatedVoltageAttr }
func (BatteryRatedVoltage) Readable() bool   { return true }
func (BatteryRatedVoltage) Writable() bool   { return true }
func (BatteryRatedVoltage) Reportable() bool { return false }
func (BatteryRatedVoltage) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BatteryRatedVoltage) AttrID() zcl.AttrID   { return v.ID() }
func (v BatteryRatedVoltage) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BatteryRatedVoltage) AttrValue() zcl.Val  { return v.Value() }

func (BatteryRatedVoltage) Name() string        { return `Battery Rated Voltage` }
func (BatteryRatedVoltage) Description() string { return `` }

type BatteryRatedVoltage zcl.Zu8

func (v *BatteryRatedVoltage) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *BatteryRatedVoltage) Value() zcl.Val     { return v }

func (v BatteryRatedVoltage) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *BatteryRatedVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = BatteryRatedVoltage(*nt)
	return br, err
}

func (v BatteryRatedVoltage) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *BatteryRatedVoltage) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BatteryRatedVoltage(*a)
	return nil
}

func (v *BatteryRatedVoltage) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = BatteryRatedVoltage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BatteryRatedVoltage) String() string {
	return zcl.Volts.Format(float64(v) / 10)
}

const BatteryRemainingAttr zcl.AttrID = 33

func (BatteryRemaining) ID() zcl.AttrID   { return BatteryRemainingAttr }
func (BatteryRemaining) Readable() bool   { return true }
func (BatteryRemaining) Writable() bool   { return false }
func (BatteryRemaining) Reportable() bool { return true }
func (BatteryRemaining) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BatteryRemaining) AttrID() zcl.AttrID   { return v.ID() }
func (v BatteryRemaining) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BatteryRemaining) AttrValue() zcl.Val  { return v.Value() }

func (BatteryRemaining) Name() string        { return `Battery Remaining` }
func (BatteryRemaining) Description() string { return `` }

type BatteryRemaining zcl.Zu8

func (v *BatteryRemaining) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *BatteryRemaining) Value() zcl.Val     { return v }

func (v BatteryRemaining) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *BatteryRemaining) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = BatteryRemaining(*nt)
	return br, err
}

func (v BatteryRemaining) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *BatteryRemaining) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BatteryRemaining(*a)
	return nil
}

func (v *BatteryRemaining) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = BatteryRemaining(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BatteryRemaining) String() string {
	return zcl.Percent.Format(float64(v) / 2)
}

const BatterySizeAttr zcl.AttrID = 49

func (BatterySize) ID() zcl.AttrID   { return BatterySizeAttr }
func (BatterySize) Readable() bool   { return true }
func (BatterySize) Writable() bool   { return true }
func (BatterySize) Reportable() bool { return false }
func (BatterySize) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BatterySize) AttrID() zcl.AttrID   { return v.ID() }
func (v BatterySize) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BatterySize) AttrValue() zcl.Val  { return v.Value() }

func (BatterySize) Name() string        { return `Battery Size` }
func (BatterySize) Description() string { return `` }

type BatterySize zcl.Zenum8

func (v *BatterySize) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *BatterySize) Value() zcl.Val     { return v }

func (v BatterySize) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *BatterySize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = BatterySize(*nt)
	return br, err
}

func (v BatterySize) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *BatterySize) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BatterySize(*a)
	return nil
}

func (v *BatterySize) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = BatterySize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BatterySize) String() string {
	switch v {
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
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v BatterySize) IsNoBattery() bool { return v == 0x00 }
func (v BatterySize) IsBuiltIn() bool   { return v == 0x01 }
func (v BatterySize) IsOther() bool     { return v == 0x02 }
func (v BatterySize) IsAa() bool        { return v == 0x03 }
func (v BatterySize) IsAaa() bool       { return v == 0x04 }
func (v BatterySize) IsC() bool         { return v == 0x05 }
func (v BatterySize) IsD() bool         { return v == 0x06 }
func (v BatterySize) IsCr2() bool       { return v == 0x07 }
func (v BatterySize) IsCr123a() bool    { return v == 0x08 }
func (v BatterySize) IsUnknown() bool   { return v == 0xFF }
func (v *BatterySize) SetNoBattery()    { *v = 0x00 }
func (v *BatterySize) SetBuiltIn()      { *v = 0x01 }
func (v *BatterySize) SetOther()        { *v = 0x02 }
func (v *BatterySize) SetAa()           { *v = 0x03 }
func (v *BatterySize) SetAaa()          { *v = 0x04 }
func (v *BatterySize) SetC()            { *v = 0x05 }
func (v *BatterySize) SetD()            { *v = 0x06 }
func (v *BatterySize) SetCr2()          { *v = 0x07 }
func (v *BatterySize) SetCr123a()       { *v = 0x08 }
func (v *BatterySize) SetUnknown()      { *v = 0xFF }

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

const BatteryVoltageAttr zcl.AttrID = 32

func (BatteryVoltage) ID() zcl.AttrID   { return BatteryVoltageAttr }
func (BatteryVoltage) Readable() bool   { return true }
func (BatteryVoltage) Writable() bool   { return false }
func (BatteryVoltage) Reportable() bool { return false }
func (BatteryVoltage) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BatteryVoltage) AttrID() zcl.AttrID   { return v.ID() }
func (v BatteryVoltage) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BatteryVoltage) AttrValue() zcl.Val  { return v.Value() }

func (BatteryVoltage) Name() string        { return `Battery Voltage` }
func (BatteryVoltage) Description() string { return `` }

type BatteryVoltage zcl.Zu8

func (v *BatteryVoltage) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *BatteryVoltage) Value() zcl.Val     { return v }

func (v BatteryVoltage) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *BatteryVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = BatteryVoltage(*nt)
	return br, err
}

func (v BatteryVoltage) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *BatteryVoltage) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BatteryVoltage(*a)
	return nil
}

func (v *BatteryVoltage) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = BatteryVoltage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BatteryVoltage) String() string {
	return zcl.Volts.Format(float64(v) / 10)
}

const BatteryVoltageMinThresholdAttr zcl.AttrID = 54

func (BatteryVoltageMinThreshold) ID() zcl.AttrID   { return BatteryVoltageMinThresholdAttr }
func (BatteryVoltageMinThreshold) Readable() bool   { return true }
func (BatteryVoltageMinThreshold) Writable() bool   { return true }
func (BatteryVoltageMinThreshold) Reportable() bool { return false }
func (BatteryVoltageMinThreshold) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BatteryVoltageMinThreshold) AttrID() zcl.AttrID   { return v.ID() }
func (v BatteryVoltageMinThreshold) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BatteryVoltageMinThreshold) AttrValue() zcl.Val  { return v.Value() }

func (BatteryVoltageMinThreshold) Name() string        { return `Battery Voltage Min Threshold` }
func (BatteryVoltageMinThreshold) Description() string { return `` }

type BatteryVoltageMinThreshold zcl.Zu8

func (v *BatteryVoltageMinThreshold) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *BatteryVoltageMinThreshold) Value() zcl.Val     { return v }

func (v BatteryVoltageMinThreshold) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *BatteryVoltageMinThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = BatteryVoltageMinThreshold(*nt)
	return br, err
}

func (v BatteryVoltageMinThreshold) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *BatteryVoltageMinThreshold) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BatteryVoltageMinThreshold(*a)
	return nil
}

func (v *BatteryVoltageMinThreshold) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = BatteryVoltageMinThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BatteryVoltageMinThreshold) String() string {
	return zcl.Volts.Format(float64(v) / 10)
}

const BatteryVoltageThreshold1Attr zcl.AttrID = 55

func (BatteryVoltageThreshold1) ID() zcl.AttrID   { return BatteryVoltageThreshold1Attr }
func (BatteryVoltageThreshold1) Readable() bool   { return true }
func (BatteryVoltageThreshold1) Writable() bool   { return true }
func (BatteryVoltageThreshold1) Reportable() bool { return false }
func (BatteryVoltageThreshold1) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BatteryVoltageThreshold1) AttrID() zcl.AttrID   { return v.ID() }
func (v BatteryVoltageThreshold1) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BatteryVoltageThreshold1) AttrValue() zcl.Val  { return v.Value() }

func (BatteryVoltageThreshold1) Name() string        { return `Battery Voltage Threshold 1` }
func (BatteryVoltageThreshold1) Description() string { return `` }

type BatteryVoltageThreshold1 zcl.Zu8

func (v *BatteryVoltageThreshold1) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *BatteryVoltageThreshold1) Value() zcl.Val     { return v }

func (v BatteryVoltageThreshold1) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *BatteryVoltageThreshold1) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = BatteryVoltageThreshold1(*nt)
	return br, err
}

func (v BatteryVoltageThreshold1) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *BatteryVoltageThreshold1) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BatteryVoltageThreshold1(*a)
	return nil
}

func (v *BatteryVoltageThreshold1) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = BatteryVoltageThreshold1(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BatteryVoltageThreshold1) String() string {
	return zcl.Volts.Format(float64(v) / 10)
}

const BatteryVoltageThreshold2Attr zcl.AttrID = 56

func (BatteryVoltageThreshold2) ID() zcl.AttrID   { return BatteryVoltageThreshold2Attr }
func (BatteryVoltageThreshold2) Readable() bool   { return true }
func (BatteryVoltageThreshold2) Writable() bool   { return true }
func (BatteryVoltageThreshold2) Reportable() bool { return false }
func (BatteryVoltageThreshold2) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BatteryVoltageThreshold2) AttrID() zcl.AttrID   { return v.ID() }
func (v BatteryVoltageThreshold2) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BatteryVoltageThreshold2) AttrValue() zcl.Val  { return v.Value() }

func (BatteryVoltageThreshold2) Name() string        { return `Battery Voltage Threshold 2` }
func (BatteryVoltageThreshold2) Description() string { return `` }

type BatteryVoltageThreshold2 zcl.Zu8

func (v *BatteryVoltageThreshold2) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *BatteryVoltageThreshold2) Value() zcl.Val     { return v }

func (v BatteryVoltageThreshold2) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *BatteryVoltageThreshold2) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = BatteryVoltageThreshold2(*nt)
	return br, err
}

func (v BatteryVoltageThreshold2) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *BatteryVoltageThreshold2) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BatteryVoltageThreshold2(*a)
	return nil
}

func (v *BatteryVoltageThreshold2) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = BatteryVoltageThreshold2(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BatteryVoltageThreshold2) String() string {
	return zcl.Volts.Format(float64(v) / 10)
}

const BatteryVoltageThreshold3Attr zcl.AttrID = 57

func (BatteryVoltageThreshold3) ID() zcl.AttrID   { return BatteryVoltageThreshold3Attr }
func (BatteryVoltageThreshold3) Readable() bool   { return true }
func (BatteryVoltageThreshold3) Writable() bool   { return true }
func (BatteryVoltageThreshold3) Reportable() bool { return false }
func (BatteryVoltageThreshold3) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BatteryVoltageThreshold3) AttrID() zcl.AttrID   { return v.ID() }
func (v BatteryVoltageThreshold3) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BatteryVoltageThreshold3) AttrValue() zcl.Val  { return v.Value() }

func (BatteryVoltageThreshold3) Name() string        { return `Battery Voltage Threshold 3` }
func (BatteryVoltageThreshold3) Description() string { return `` }

type BatteryVoltageThreshold3 zcl.Zu8

func (v *BatteryVoltageThreshold3) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *BatteryVoltageThreshold3) Value() zcl.Val     { return v }

func (v BatteryVoltageThreshold3) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *BatteryVoltageThreshold3) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = BatteryVoltageThreshold3(*nt)
	return br, err
}

func (v BatteryVoltageThreshold3) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *BatteryVoltageThreshold3) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BatteryVoltageThreshold3(*a)
	return nil
}

func (v *BatteryVoltageThreshold3) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = BatteryVoltageThreshold3(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BatteryVoltageThreshold3) String() string {
	return zcl.Volts.Format(float64(v) / 10)
}

const BatteryCapacityAttr zcl.AttrID = 50

func (BatteryCapacity) ID() zcl.AttrID   { return BatteryCapacityAttr }
func (BatteryCapacity) Readable() bool   { return true }
func (BatteryCapacity) Writable() bool   { return true }
func (BatteryCapacity) Reportable() bool { return false }
func (BatteryCapacity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BatteryCapacity) AttrID() zcl.AttrID   { return v.ID() }
func (v BatteryCapacity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BatteryCapacity) AttrValue() zcl.Val  { return v.Value() }

func (BatteryCapacity) Name() string        { return `Battery capacity` }
func (BatteryCapacity) Description() string { return `` }

type BatteryCapacity zcl.Zu16

func (v *BatteryCapacity) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *BatteryCapacity) Value() zcl.Val     { return v }

func (v BatteryCapacity) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *BatteryCapacity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = BatteryCapacity(*nt)
	return br, err
}

func (v BatteryCapacity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *BatteryCapacity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BatteryCapacity(*a)
	return nil
}

func (v *BatteryCapacity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = BatteryCapacity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BatteryCapacity) String() string {
	return zcl.MilliAmpereHours.Format(float64(v) / 0.1)
}

const BinaryActiveTextAttr zcl.AttrID = 4

func (BinaryActiveText) ID() zcl.AttrID   { return BinaryActiveTextAttr }
func (BinaryActiveText) Readable() bool   { return true }
func (BinaryActiveText) Writable() bool   { return true }
func (BinaryActiveText) Reportable() bool { return false }
func (BinaryActiveText) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BinaryActiveText) AttrID() zcl.AttrID   { return v.ID() }
func (v BinaryActiveText) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BinaryActiveText) AttrValue() zcl.Val  { return v.Value() }

func (BinaryActiveText) Name() string        { return `Binary Active Text` }
func (BinaryActiveText) Description() string { return `` }

type BinaryActiveText zcl.Zcstring

func (v *BinaryActiveText) TypeID() zcl.TypeID { return new(zcl.Zcstring).TypeID() }
func (v *BinaryActiveText) Value() zcl.Val     { return v }

func (v BinaryActiveText) MarshalZcl() ([]byte, error) { return zcl.Zcstring(v).MarshalZcl() }

func (v *BinaryActiveText) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*v = BinaryActiveText(*nt)
	return br, err
}

func (v BinaryActiveText) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zcstring(v))
}

func (v *BinaryActiveText) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zcstring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BinaryActiveText(*a)
	return nil
}

func (v *BinaryActiveText) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zcstring); ok {
		*v = BinaryActiveText(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BinaryActiveText) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v))
}

const BinaryInactiveTextAttr zcl.AttrID = 46

func (BinaryInactiveText) ID() zcl.AttrID   { return BinaryInactiveTextAttr }
func (BinaryInactiveText) Readable() bool   { return true }
func (BinaryInactiveText) Writable() bool   { return true }
func (BinaryInactiveText) Reportable() bool { return false }
func (BinaryInactiveText) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BinaryInactiveText) AttrID() zcl.AttrID   { return v.ID() }
func (v BinaryInactiveText) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BinaryInactiveText) AttrValue() zcl.Val  { return v.Value() }

func (BinaryInactiveText) Name() string        { return `Binary Inactive Text` }
func (BinaryInactiveText) Description() string { return `` }

type BinaryInactiveText zcl.Zcstring

func (v *BinaryInactiveText) TypeID() zcl.TypeID { return new(zcl.Zcstring).TypeID() }
func (v *BinaryInactiveText) Value() zcl.Val     { return v }

func (v BinaryInactiveText) MarshalZcl() ([]byte, error) { return zcl.Zcstring(v).MarshalZcl() }

func (v *BinaryInactiveText) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*v = BinaryInactiveText(*nt)
	return br, err
}

func (v BinaryInactiveText) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zcstring(v))
}

func (v *BinaryInactiveText) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zcstring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BinaryInactiveText(*a)
	return nil
}

func (v *BinaryInactiveText) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zcstring); ok {
		*v = BinaryInactiveText(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BinaryInactiveText) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v))
}

const BinaryMaxOffTimeAttr zcl.AttrID = 67

func (BinaryMaxOffTime) ID() zcl.AttrID   { return BinaryMaxOffTimeAttr }
func (BinaryMaxOffTime) Readable() bool   { return true }
func (BinaryMaxOffTime) Writable() bool   { return true }
func (BinaryMaxOffTime) Reportable() bool { return false }
func (BinaryMaxOffTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BinaryMaxOffTime) AttrID() zcl.AttrID   { return v.ID() }
func (v BinaryMaxOffTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BinaryMaxOffTime) AttrValue() zcl.Val  { return v.Value() }

func (BinaryMaxOffTime) Name() string        { return `Binary Max Off-time` }
func (BinaryMaxOffTime) Description() string { return `` }

type BinaryMaxOffTime zcl.Zu32

func (v *BinaryMaxOffTime) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *BinaryMaxOffTime) Value() zcl.Val     { return v }

func (v BinaryMaxOffTime) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *BinaryMaxOffTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = BinaryMaxOffTime(*nt)
	return br, err
}

func (v BinaryMaxOffTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *BinaryMaxOffTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BinaryMaxOffTime(*a)
	return nil
}

func (v *BinaryMaxOffTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = BinaryMaxOffTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BinaryMaxOffTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

const BinaryMinOffTimeAttr zcl.AttrID = 66

func (BinaryMinOffTime) ID() zcl.AttrID   { return BinaryMinOffTimeAttr }
func (BinaryMinOffTime) Readable() bool   { return true }
func (BinaryMinOffTime) Writable() bool   { return true }
func (BinaryMinOffTime) Reportable() bool { return false }
func (BinaryMinOffTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BinaryMinOffTime) AttrID() zcl.AttrID   { return v.ID() }
func (v BinaryMinOffTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BinaryMinOffTime) AttrValue() zcl.Val  { return v.Value() }

func (BinaryMinOffTime) Name() string        { return `Binary Min Off-time` }
func (BinaryMinOffTime) Description() string { return `` }

type BinaryMinOffTime zcl.Zu32

func (v *BinaryMinOffTime) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *BinaryMinOffTime) Value() zcl.Val     { return v }

func (v BinaryMinOffTime) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *BinaryMinOffTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = BinaryMinOffTime(*nt)
	return br, err
}

func (v BinaryMinOffTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *BinaryMinOffTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BinaryMinOffTime(*a)
	return nil
}

func (v *BinaryMinOffTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = BinaryMinOffTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BinaryMinOffTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

const BinaryPolarityAttr zcl.AttrID = 84

func (BinaryPolarity) ID() zcl.AttrID   { return BinaryPolarityAttr }
func (BinaryPolarity) Readable() bool   { return true }
func (BinaryPolarity) Writable() bool   { return false }
func (BinaryPolarity) Reportable() bool { return false }
func (BinaryPolarity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BinaryPolarity) AttrID() zcl.AttrID   { return v.ID() }
func (v BinaryPolarity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BinaryPolarity) AttrValue() zcl.Val  { return v.Value() }

func (BinaryPolarity) Name() string        { return `Binary Polarity` }
func (BinaryPolarity) Description() string { return `` }

type BinaryPolarity zcl.Zenum8

func (v *BinaryPolarity) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *BinaryPolarity) Value() zcl.Val     { return v }

func (v BinaryPolarity) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *BinaryPolarity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = BinaryPolarity(*nt)
	return br, err
}

func (v BinaryPolarity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *BinaryPolarity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BinaryPolarity(*a)
	return nil
}

func (v *BinaryPolarity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = BinaryPolarity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BinaryPolarity) String() string {
	switch v {
	case 0x00:
		return "Normal"
	case 0x01:
		return "Reverse"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v BinaryPolarity) IsNormal() bool  { return v == 0x00 }
func (v BinaryPolarity) IsReverse() bool { return v == 0x01 }
func (v *BinaryPolarity) SetNormal()     { *v = 0x00 }
func (v *BinaryPolarity) SetReverse()    { *v = 0x01 }

func (BinaryPolarity) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Normal"},
		{Value: 0x01, Name: "Reverse"},
	}
}

const BinaryPresentValueAttr zcl.AttrID = 85

func (BinaryPresentValue) ID() zcl.AttrID   { return BinaryPresentValueAttr }
func (BinaryPresentValue) Readable() bool   { return true }
func (BinaryPresentValue) Writable() bool   { return true }
func (BinaryPresentValue) Reportable() bool { return true }
func (BinaryPresentValue) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BinaryPresentValue) AttrID() zcl.AttrID   { return v.ID() }
func (v BinaryPresentValue) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BinaryPresentValue) AttrValue() zcl.Val  { return v.Value() }

func (BinaryPresentValue) Name() string        { return `Binary Present Value` }
func (BinaryPresentValue) Description() string { return `` }

type BinaryPresentValue zcl.Zbool

func (v *BinaryPresentValue) TypeID() zcl.TypeID { return new(zcl.Zbool).TypeID() }
func (v *BinaryPresentValue) Value() zcl.Val     { return v }

func (v BinaryPresentValue) MarshalZcl() ([]byte, error) { return zcl.Zbool(v).MarshalZcl() }

func (v *BinaryPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*v = BinaryPresentValue(*nt)
	return br, err
}

func (v BinaryPresentValue) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbool(v))
}

func (v *BinaryPresentValue) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbool)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BinaryPresentValue(*a)
	return nil
}

func (v *BinaryPresentValue) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbool); ok {
		*v = BinaryPresentValue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BinaryPresentValue) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(v))
}

func (BinaryPriority) Name() string        { return `Binary Priority` }
func (BinaryPriority) Description() string { return `` }

type BinaryPriority struct {
	Ispriority    zcl.Zbool
	Priorityvalue zcl.Zbool
}

func (v *BinaryPriority) TypeID() zcl.TypeID { return zcl.TypeID(76) } // struct
func (v *BinaryPriority) Value() zcl.Val     { return v }
func (v BinaryPriority) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	var err error
	_ = tmp2
	_ = err

	{
		if tmp, err = v.Ispriority.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Priorityvalue.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (v *BinaryPriority) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2
	_ = err

	if b, err = new(zcl.Zbool).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = new(zcl.Zbool).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v *BinaryPriority) SetValue(a zcl.Val) error {
	if nv, ok := a.(*BinaryPriority); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrNotImpl
}

func (v *BinaryPriority) String() string {
	return zcl.Sprintf(
		"BinaryPriority{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

const BinaryPriorityArrayAttr zcl.AttrID = 87

func (BinaryPriorityArray) ID() zcl.AttrID   { return BinaryPriorityArrayAttr }
func (BinaryPriorityArray) Readable() bool   { return true }
func (BinaryPriorityArray) Writable() bool   { return true }
func (BinaryPriorityArray) Reportable() bool { return false }
func (BinaryPriorityArray) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BinaryPriorityArray) AttrID() zcl.AttrID   { return v.ID() }
func (v BinaryPriorityArray) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BinaryPriorityArray) AttrValue() zcl.Val  { return v.Value() }

func (BinaryPriorityArray) Name() string        { return `Binary Priority Array` }
func (BinaryPriorityArray) Description() string { return `` }

type BinaryPriorityArray []*BinaryPriority

func (v *BinaryPriorityArray) TypeID() zcl.TypeID { return new(zcl.Zarray).TypeID() }
func (v *BinaryPriorityArray) Value() zcl.Val     { return v }

func (BinaryPriorityArray) ArrayTypeID() zcl.TypeID { return new(BinaryPriority).TypeID() }

func (v *BinaryPriorityArray) ArrayValues() (o []BinaryPriority) {
	for _, a := range *v {
		o = append(o, *a)
	}
	return o
}

func (v *BinaryPriorityArray) SetValues(val []BinaryPriority) error {
	*v = []*BinaryPriority{}
	return v.AddValues(val...)
}

func (v *BinaryPriorityArray) AddValues(val ...BinaryPriority) error {
	for _, a := range val {
		nv := a
		*v = append(*v, &nv)
	}
	return nil
}

func (v BinaryPriorityArray) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, a := range v {
		vars = append(vars, a)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (v *BinaryPriorityArray) UnmarshalZcl(b []byte) ([]byte, error) {
	*v = []*BinaryPriority{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(BinaryPriority)
		*v = append(*v, nv)
		return nv
	})
}

func (v *BinaryPriorityArray) SetValue(a zcl.Val) error {
	if nv, ok := a.(*BinaryPriorityArray); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BinaryPriorityArray) String() string {
	var s []string
	for _, a := range v {
		s = append(s, zcl.Sprintf("%v", a))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

const BinaryRelinquishDefaultAttr zcl.AttrID = 104

func (BinaryRelinquishDefault) ID() zcl.AttrID   { return BinaryRelinquishDefaultAttr }
func (BinaryRelinquishDefault) Readable() bool   { return true }
func (BinaryRelinquishDefault) Writable() bool   { return true }
func (BinaryRelinquishDefault) Reportable() bool { return false }
func (BinaryRelinquishDefault) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BinaryRelinquishDefault) AttrID() zcl.AttrID   { return v.ID() }
func (v BinaryRelinquishDefault) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BinaryRelinquishDefault) AttrValue() zcl.Val  { return v.Value() }

func (BinaryRelinquishDefault) Name() string        { return `Binary Relinquish Default` }
func (BinaryRelinquishDefault) Description() string { return `` }

type BinaryRelinquishDefault zcl.Zbool

func (v *BinaryRelinquishDefault) TypeID() zcl.TypeID { return new(zcl.Zbool).TypeID() }
func (v *BinaryRelinquishDefault) Value() zcl.Val     { return v }

func (v BinaryRelinquishDefault) MarshalZcl() ([]byte, error) { return zcl.Zbool(v).MarshalZcl() }

func (v *BinaryRelinquishDefault) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*v = BinaryRelinquishDefault(*nt)
	return br, err
}

func (v BinaryRelinquishDefault) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbool(v))
}

func (v *BinaryRelinquishDefault) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbool)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BinaryRelinquishDefault(*a)
	return nil
}

func (v *BinaryRelinquishDefault) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbool); ok {
		*v = BinaryRelinquishDefault(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BinaryRelinquishDefault) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(v))
}

const CalculationPeriodAttr zcl.AttrID = 22

func (CalculationPeriod) ID() zcl.AttrID   { return CalculationPeriodAttr }
func (CalculationPeriod) Readable() bool   { return true }
func (CalculationPeriod) Writable() bool   { return true }
func (CalculationPeriod) Reportable() bool { return false }
func (CalculationPeriod) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CalculationPeriod) AttrID() zcl.AttrID   { return v.ID() }
func (v CalculationPeriod) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CalculationPeriod) AttrValue() zcl.Val  { return v.Value() }

func (CalculationPeriod) Name() string        { return `Calculation Period` }
func (CalculationPeriod) Description() string { return `` }

type CalculationPeriod zcl.Zu16

func (v *CalculationPeriod) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *CalculationPeriod) Value() zcl.Val     { return v }

func (v CalculationPeriod) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *CalculationPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = CalculationPeriod(*nt)
	return br, err
}

func (v CalculationPeriod) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *CalculationPeriod) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CalculationPeriod(*a)
	return nil
}

func (v *CalculationPeriod) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = CalculationPeriod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CalculationPeriod) String() string {
	return zcl.Seconds.Format(float64(v))
}

const CheckInIntervalAttr zcl.AttrID = 0

func (CheckInInterval) ID() zcl.AttrID   { return CheckInIntervalAttr }
func (CheckInInterval) Readable() bool   { return true }
func (CheckInInterval) Writable() bool   { return true }
func (CheckInInterval) Reportable() bool { return false }
func (CheckInInterval) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CheckInInterval) AttrID() zcl.AttrID   { return v.ID() }
func (v CheckInInterval) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CheckInInterval) AttrValue() zcl.Val  { return v.Value() }

func (CheckInInterval) Name() string        { return `Check-in Interval` }
func (CheckInInterval) Description() string { return `` }

type CheckInInterval zcl.Zu32

func (v *CheckInInterval) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *CheckInInterval) Value() zcl.Val     { return v }

func (v CheckInInterval) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *CheckInInterval) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = CheckInInterval(*nt)
	return br, err
}

func (v CheckInInterval) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *CheckInInterval) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CheckInInterval(*a)
	return nil
}

func (v *CheckInInterval) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = CheckInInterval(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CheckInInterval) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

const CheckInIntervalMinAttr zcl.AttrID = 4

func (CheckInIntervalMin) ID() zcl.AttrID   { return CheckInIntervalMinAttr }
func (CheckInIntervalMin) Readable() bool   { return true }
func (CheckInIntervalMin) Writable() bool   { return false }
func (CheckInIntervalMin) Reportable() bool { return false }
func (CheckInIntervalMin) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CheckInIntervalMin) AttrID() zcl.AttrID   { return v.ID() }
func (v CheckInIntervalMin) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CheckInIntervalMin) AttrValue() zcl.Val  { return v.Value() }

func (CheckInIntervalMin) Name() string        { return `Check-in Interval Min` }
func (CheckInIntervalMin) Description() string { return `` }

type CheckInIntervalMin zcl.Zu32

func (v *CheckInIntervalMin) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *CheckInIntervalMin) Value() zcl.Val     { return v }

func (v CheckInIntervalMin) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *CheckInIntervalMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = CheckInIntervalMin(*nt)
	return br, err
}

func (v CheckInIntervalMin) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *CheckInIntervalMin) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CheckInIntervalMin(*a)
	return nil
}

func (v *CheckInIntervalMin) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = CheckInIntervalMin(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CheckInIntervalMin) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

const ChildMovedAttr zcl.AttrID = 273

func (ChildMoved) ID() zcl.AttrID   { return ChildMovedAttr }
func (ChildMoved) Readable() bool   { return true }
func (ChildMoved) Writable() bool   { return false }
func (ChildMoved) Reportable() bool { return false }
func (ChildMoved) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ChildMoved) AttrID() zcl.AttrID   { return v.ID() }
func (v ChildMoved) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ChildMoved) AttrValue() zcl.Val  { return v.Value() }

func (ChildMoved) Name() string        { return `Child Moved` }
func (ChildMoved) Description() string { return `` }

type ChildMoved zcl.Zu16

func (v *ChildMoved) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ChildMoved) Value() zcl.Val     { return v }

func (v ChildMoved) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ChildMoved) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ChildMoved(*nt)
	return br, err
}

func (v ChildMoved) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ChildMoved) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ChildMoved(*a)
	return nil
}

func (v *ChildMoved) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ChildMoved(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ChildMoved) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (ClusterId) Name() string        { return `Cluster Id` }
func (ClusterId) Description() string { return `` }

type ClusterId zcl.Zu16

func (v *ClusterId) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ClusterId) Value() zcl.Val     { return v }

func (v ClusterId) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ClusterId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ClusterId(*nt)
	return br, err
}

func (v ClusterId) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ClusterId) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ClusterId(*a)
	return nil
}

func (v *ClusterId) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ClusterId(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ClusterId) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ClusterRevisionAttr zcl.AttrID = 65533

func (ClusterRevision) ID() zcl.AttrID   { return ClusterRevisionAttr }
func (ClusterRevision) Readable() bool   { return true }
func (ClusterRevision) Writable() bool   { return true }
func (ClusterRevision) Reportable() bool { return false }
func (ClusterRevision) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ClusterRevision) AttrID() zcl.AttrID   { return v.ID() }
func (v ClusterRevision) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ClusterRevision) AttrValue() zcl.Val  { return v.Value() }

func (ClusterRevision) Name() string        { return `Cluster Revision` }
func (ClusterRevision) Description() string { return `` }

type ClusterRevision zcl.Zu16

func (v *ClusterRevision) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ClusterRevision) Value() zcl.Val     { return v }

func (v ClusterRevision) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ClusterRevision) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ClusterRevision(*nt)
	return br, err
}

func (v ClusterRevision) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ClusterRevision) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ClusterRevision(*a)
	return nil
}

func (v *ClusterRevision) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ClusterRevision(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ClusterRevision) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ConfigurationAttr zcl.AttrID = 49

func (Configuration) ID() zcl.AttrID   { return ConfigurationAttr }
func (Configuration) Readable() bool   { return true }
func (Configuration) Writable() bool   { return true }
func (Configuration) Reportable() bool { return false }
func (Configuration) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Configuration) AttrID() zcl.AttrID   { return v.ID() }
func (v Configuration) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Configuration) AttrValue() zcl.Val  { return v.Value() }

func (Configuration) Name() string        { return `Configuration` }
func (Configuration) Description() string { return `` }

type Configuration zcl.Zbmp16

func (v *Configuration) TypeID() zcl.TypeID { return new(zcl.Zbmp16).TypeID() }
func (v *Configuration) Value() zcl.Val     { return v }

func (v Configuration) MarshalZcl() ([]byte, error) { return zcl.Zbmp16(v).MarshalZcl() }

func (v *Configuration) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*v = Configuration(*nt)
	return br, err
}

func (v Configuration) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp16(v))
}

func (v *Configuration) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Configuration(*a)
	return nil
}

func (v *Configuration) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp16); ok {
		*v = Configuration(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Configuration) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v Configuration) IsTouchlinkEnabled0() bool { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v Configuration) IsTouchlinkEnabled1() bool { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v Configuration) IsTouchlinkEnabled2() bool { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v *Configuration) SetTouchlinkEnabled0(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *Configuration) SetTouchlinkEnabled1(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *Configuration) SetTouchlinkEnabled2(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b))
}

func (Configuration) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Touchlink enabled 0"},
		{Value: 1, Name: "Touchlink enabled 1"},
		{Value: 3, Name: "Touchlink enabled 2"},
	}
}

const CurrentGroupAttr zcl.AttrID = 2

func (CurrentGroup) ID() zcl.AttrID   { return CurrentGroupAttr }
func (CurrentGroup) Readable() bool   { return true }
func (CurrentGroup) Writable() bool   { return false }
func (CurrentGroup) Reportable() bool { return false }
func (CurrentGroup) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentGroup) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentGroup) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentGroup) AttrValue() zcl.Val  { return v.Value() }

func (CurrentGroup) Name() string        { return `Current Group` }
func (CurrentGroup) Description() string { return `` }

type CurrentGroup zcl.Zu16

func (v *CurrentGroup) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *CurrentGroup) Value() zcl.Val     { return v }

func (v CurrentGroup) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *CurrentGroup) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentGroup(*nt)
	return br, err
}

func (v CurrentGroup) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *CurrentGroup) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentGroup(*a)
	return nil
}

func (v *CurrentGroup) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = CurrentGroup(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentGroup) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const CurrentLevelAttr zcl.AttrID = 0

func (CurrentLevel) ID() zcl.AttrID   { return CurrentLevelAttr }
func (CurrentLevel) Readable() bool   { return true }
func (CurrentLevel) Writable() bool   { return false }
func (CurrentLevel) Reportable() bool { return true }
func (CurrentLevel) SceneIndex() int  { return 1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentLevel) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentLevel) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentLevel) AttrValue() zcl.Val  { return v.Value() }

func (CurrentLevel) Name() string { return `Current Level` }
func (CurrentLevel) Description() string {
	return `represents the current level of this device. Meaning of 'level' is device dependent.`
}

// CurrentLevel represents the current level of this device. Meaning of 'level' is device dependent.
type CurrentLevel zcl.Zu8

func (v *CurrentLevel) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *CurrentLevel) Value() zcl.Val     { return v }

func (v CurrentLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *CurrentLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentLevel(*nt)
	return br, err
}

func (v CurrentLevel) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *CurrentLevel) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentLevel(*a)
	return nil
}

func (v *CurrentLevel) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = CurrentLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentLevel) String() string {
	return zcl.Percent.Format(float64(v) / 2.54)
}

const CurrentSceneAttr zcl.AttrID = 1

func (CurrentScene) ID() zcl.AttrID   { return CurrentSceneAttr }
func (CurrentScene) Readable() bool   { return true }
func (CurrentScene) Writable() bool   { return false }
func (CurrentScene) Reportable() bool { return false }
func (CurrentScene) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentScene) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentScene) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentScene) AttrValue() zcl.Val  { return v.Value() }

func (CurrentScene) Name() string        { return `Current Scene` }
func (CurrentScene) Description() string { return `` }

type CurrentScene zcl.Zu8

func (v *CurrentScene) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *CurrentScene) Value() zcl.Val     { return v }

func (v CurrentScene) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *CurrentScene) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentScene(*nt)
	return br, err
}

func (v CurrentScene) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *CurrentScene) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentScene(*a)
	return nil
}

func (v *CurrentScene) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = CurrentScene(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentScene) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const CurrentTemperatureAttr zcl.AttrID = 0

func (CurrentTemperature) ID() zcl.AttrID   { return CurrentTemperatureAttr }
func (CurrentTemperature) Readable() bool   { return true }
func (CurrentTemperature) Writable() bool   { return false }
func (CurrentTemperature) Reportable() bool { return false }
func (CurrentTemperature) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentTemperature) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentTemperature) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentTemperature) AttrValue() zcl.Val  { return v.Value() }

func (CurrentTemperature) Name() string        { return `Current Temperature` }
func (CurrentTemperature) Description() string { return `` }

type CurrentTemperature zcl.Zs16

func (v *CurrentTemperature) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *CurrentTemperature) Value() zcl.Val     { return v }

func (v CurrentTemperature) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *CurrentTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentTemperature(*nt)
	return br, err
}

func (v CurrentTemperature) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *CurrentTemperature) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentTemperature(*a)
	return nil
}

func (v *CurrentTemperature) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = CurrentTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentTemperature) String() string {
	return zcl.DegreesCelsius.Format(float64(v))
}

const DateCodeAttr zcl.AttrID = 6

func (DateCode) ID() zcl.AttrID   { return DateCodeAttr }
func (DateCode) Readable() bool   { return true }
func (DateCode) Writable() bool   { return false }
func (DateCode) Reportable() bool { return false }
func (DateCode) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DateCode) AttrID() zcl.AttrID   { return v.ID() }
func (v DateCode) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DateCode) AttrValue() zcl.Val  { return v.Value() }

func (DateCode) Name() string        { return `Date Code` }
func (DateCode) Description() string { return `` }

type DateCode zcl.Zcstring

func (v *DateCode) TypeID() zcl.TypeID { return new(zcl.Zcstring).TypeID() }
func (v *DateCode) Value() zcl.Val     { return v }

func (v DateCode) MarshalZcl() ([]byte, error) { return zcl.Zcstring(v).MarshalZcl() }

func (v *DateCode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*v = DateCode(*nt)
	return br, err
}

func (v DateCode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zcstring(v))
}

func (v *DateCode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zcstring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DateCode(*a)
	return nil
}

func (v *DateCode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zcstring); ok {
		*v = DateCode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DateCode) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v))
}

const DefaultMoveRateAttr zcl.AttrID = 20

func (DefaultMoveRate) ID() zcl.AttrID   { return DefaultMoveRateAttr }
func (DefaultMoveRate) Readable() bool   { return true }
func (DefaultMoveRate) Writable() bool   { return true }
func (DefaultMoveRate) Reportable() bool { return false }
func (DefaultMoveRate) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DefaultMoveRate) AttrID() zcl.AttrID   { return v.ID() }
func (v DefaultMoveRate) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DefaultMoveRate) AttrValue() zcl.Val  { return v.Value() }

func (DefaultMoveRate) Name() string        { return `Default Move Rate` }
func (DefaultMoveRate) Description() string { return `` }

type DefaultMoveRate zcl.Zu8

func (v *DefaultMoveRate) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *DefaultMoveRate) Value() zcl.Val     { return v }

func (v DefaultMoveRate) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *DefaultMoveRate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = DefaultMoveRate(*nt)
	return br, err
}

func (v DefaultMoveRate) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *DefaultMoveRate) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DefaultMoveRate(*a)
	return nil
}

func (v *DefaultMoveRate) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = DefaultMoveRate(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DefaultMoveRate) String() string {
	return zcl.PercentPerSecond.Format(float64(v) / 2.54)
}

func (Device) Name() string        { return `Device` }
func (Device) Description() string { return `` }

type Device zcl.Zuid

func (v *Device) TypeID() zcl.TypeID { return new(zcl.Zuid).TypeID() }
func (v *Device) Value() zcl.Val     { return v }

func (v Device) MarshalZcl() ([]byte, error) { return zcl.Zuid(v).MarshalZcl() }

func (v *Device) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*v = Device(*nt)
	return br, err
}

func (v Device) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zuid(v))
}

func (v *Device) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zuid)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Device(*a)
	return nil
}

func (v *Device) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zuid); ok {
		*v = Device(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Device) String() string {
	return zcl.Sprintf("%v", zcl.Zuid(v))
}

const DeviceEnabledAttr zcl.AttrID = 18

func (DeviceEnabled) ID() zcl.AttrID   { return DeviceEnabledAttr }
func (DeviceEnabled) Readable() bool   { return true }
func (DeviceEnabled) Writable() bool   { return true }
func (DeviceEnabled) Reportable() bool { return false }
func (DeviceEnabled) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DeviceEnabled) AttrID() zcl.AttrID   { return v.ID() }
func (v DeviceEnabled) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DeviceEnabled) AttrValue() zcl.Val  { return v.Value() }

func (DeviceEnabled) Name() string        { return `Device Enabled` }
func (DeviceEnabled) Description() string { return `` }

type DeviceEnabled zcl.Zbool

func (v *DeviceEnabled) TypeID() zcl.TypeID { return new(zcl.Zbool).TypeID() }
func (v *DeviceEnabled) Value() zcl.Val     { return v }

func (v DeviceEnabled) MarshalZcl() ([]byte, error) { return zcl.Zbool(v).MarshalZcl() }

func (v *DeviceEnabled) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*v = DeviceEnabled(*nt)
	return br, err
}

func (v DeviceEnabled) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbool(v))
}

func (v *DeviceEnabled) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbool)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DeviceEnabled(*a)
	return nil
}

func (v *DeviceEnabled) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbool); ok {
		*v = DeviceEnabled(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DeviceEnabled) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(v))
}

const DeviceTempAlarmMaskAttr zcl.AttrID = 16

func (DeviceTempAlarmMask) ID() zcl.AttrID   { return DeviceTempAlarmMaskAttr }
func (DeviceTempAlarmMask) Readable() bool   { return true }
func (DeviceTempAlarmMask) Writable() bool   { return true }
func (DeviceTempAlarmMask) Reportable() bool { return false }
func (DeviceTempAlarmMask) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DeviceTempAlarmMask) AttrID() zcl.AttrID   { return v.ID() }
func (v DeviceTempAlarmMask) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DeviceTempAlarmMask) AttrValue() zcl.Val  { return v.Value() }

func (DeviceTempAlarmMask) Name() string        { return `Device Temp Alarm Mask` }
func (DeviceTempAlarmMask) Description() string { return `` }

type DeviceTempAlarmMask zcl.Zbmp8

func (v *DeviceTempAlarmMask) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *DeviceTempAlarmMask) Value() zcl.Val     { return v }

func (v DeviceTempAlarmMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *DeviceTempAlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = DeviceTempAlarmMask(*nt)
	return br, err
}

func (v DeviceTempAlarmMask) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *DeviceTempAlarmMask) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DeviceTempAlarmMask(*a)
	return nil
}

func (v *DeviceTempAlarmMask) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = DeviceTempAlarmMask(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DeviceTempAlarmMask) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v DeviceTempAlarmMask) IsTemperatureTooLow() bool  { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v DeviceTempAlarmMask) IsTemperatureTooHigh() bool { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v *DeviceTempAlarmMask) SetTemperatureTooLow(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *DeviceTempAlarmMask) SetTemperatureTooHigh(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}

func (DeviceTempAlarmMask) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Temperature too low"},
		{Value: 1, Name: "Temperature too high"},
	}
}

const DisableLocalConfigAttr zcl.AttrID = 20

func (DisableLocalConfig) ID() zcl.AttrID   { return DisableLocalConfigAttr }
func (DisableLocalConfig) Readable() bool   { return true }
func (DisableLocalConfig) Writable() bool   { return true }
func (DisableLocalConfig) Reportable() bool { return false }
func (DisableLocalConfig) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DisableLocalConfig) AttrID() zcl.AttrID   { return v.ID() }
func (v DisableLocalConfig) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DisableLocalConfig) AttrValue() zcl.Val  { return v.Value() }

func (DisableLocalConfig) Name() string        { return `Disable Local Config` }
func (DisableLocalConfig) Description() string { return `` }

type DisableLocalConfig zcl.Zbmp8

func (v *DisableLocalConfig) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *DisableLocalConfig) Value() zcl.Val     { return v }

func (v DisableLocalConfig) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *DisableLocalConfig) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = DisableLocalConfig(*nt)
	return br, err
}

func (v DisableLocalConfig) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *DisableLocalConfig) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DisableLocalConfig(*a)
	return nil
}

func (v *DisableLocalConfig) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = DisableLocalConfig(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DisableLocalConfig) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v DisableLocalConfig) IsDisableFactoryReset() bool { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v DisableLocalConfig) IsDisableDeviceConfiguration() bool {
	return zcl.BitmapTest([]byte(v[:]), 1)
}
func (v *DisableLocalConfig) SetDisableFactoryReset(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *DisableLocalConfig) SetDisableDeviceConfiguration(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}

func (DisableLocalConfig) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Disable Factory Reset"},
		{Value: 1, Name: "Disable Device Configuration"},
	}
}

func (Distance) Name() string        { return `Distance` }
func (Distance) Description() string { return `` }

type Distance zcl.Zu16

func (v *Distance) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Distance) Value() zcl.Val     { return v }

func (v Distance) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Distance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Distance(*nt)
	return br, err
}

func (v Distance) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Distance) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Distance(*a)
	return nil
}

func (v *Distance) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Distance(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Distance) String() string {
	return zcl.Meters.Format(float64(v) / 10)
}

const DstEndAttr zcl.AttrID = 4

func (DstEnd) ID() zcl.AttrID   { return DstEndAttr }
func (DstEnd) Readable() bool   { return true }
func (DstEnd) Writable() bool   { return true }
func (DstEnd) Reportable() bool { return false }
func (DstEnd) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DstEnd) AttrID() zcl.AttrID   { return v.ID() }
func (v DstEnd) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DstEnd) AttrValue() zcl.Val  { return v.Value() }

func (DstEnd) Name() string        { return `Dst End` }
func (DstEnd) Description() string { return `` }

type DstEnd zcl.Zutc

func (v *DstEnd) TypeID() zcl.TypeID { return new(zcl.Zutc).TypeID() }
func (v *DstEnd) Value() zcl.Val     { return v }

func (v DstEnd) MarshalZcl() ([]byte, error) { return zcl.Zutc(v).MarshalZcl() }

func (v *DstEnd) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*v = DstEnd(*nt)
	return br, err
}

func (v DstEnd) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zutc(v))
}

func (v *DstEnd) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zutc)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DstEnd(*a)
	return nil
}

func (v *DstEnd) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zutc); ok {
		*v = DstEnd(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DstEnd) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(v))
}

const DstShiftAttr zcl.AttrID = 5

func (DstShift) ID() zcl.AttrID   { return DstShiftAttr }
func (DstShift) Readable() bool   { return true }
func (DstShift) Writable() bool   { return true }
func (DstShift) Reportable() bool { return false }
func (DstShift) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DstShift) AttrID() zcl.AttrID   { return v.ID() }
func (v DstShift) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DstShift) AttrValue() zcl.Val  { return v.Value() }

func (DstShift) Name() string        { return `Dst Shift` }
func (DstShift) Description() string { return `` }

type DstShift zcl.Zs32

func (v *DstShift) TypeID() zcl.TypeID { return new(zcl.Zs32).TypeID() }
func (v *DstShift) Value() zcl.Val     { return v }

func (v DstShift) MarshalZcl() ([]byte, error) { return zcl.Zs32(v).MarshalZcl() }

func (v *DstShift) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs32)
	br, err := nt.UnmarshalZcl(b)
	*v = DstShift(*nt)
	return br, err
}

func (v DstShift) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs32(v))
}

func (v *DstShift) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DstShift(*a)
	return nil
}

func (v *DstShift) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs32); ok {
		*v = DstShift(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DstShift) String() string {
	return zcl.Seconds.Format(float64(v))
}

const DstStartAttr zcl.AttrID = 3

func (DstStart) ID() zcl.AttrID   { return DstStartAttr }
func (DstStart) Readable() bool   { return true }
func (DstStart) Writable() bool   { return true }
func (DstStart) Reportable() bool { return false }
func (DstStart) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DstStart) AttrID() zcl.AttrID   { return v.ID() }
func (v DstStart) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DstStart) AttrValue() zcl.Val  { return v.Value() }

func (DstStart) Name() string        { return `Dst Start` }
func (DstStart) Description() string { return `` }

type DstStart zcl.Zutc

func (v *DstStart) TypeID() zcl.TypeID { return new(zcl.Zutc).TypeID() }
func (v *DstStart) Value() zcl.Val     { return v }

func (v DstStart) MarshalZcl() ([]byte, error) { return zcl.Zutc(v).MarshalZcl() }

func (v *DstStart) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*v = DstStart(*nt)
	return br, err
}

func (v DstStart) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zutc(v))
}

func (v *DstStart) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zutc)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DstStart(*a)
	return nil
}

func (v *DstStart) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zutc); ok {
		*v = DstStart(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DstStart) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(v))
}

func (EffectIdentifier) Name() string        { return `Effect Identifier` }
func (EffectIdentifier) Description() string { return `when turning lights off` }

// EffectIdentifier when turning lights off
type EffectIdentifier zcl.Zenum8

func (v *EffectIdentifier) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *EffectIdentifier) Value() zcl.Val     { return v }

func (v EffectIdentifier) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *EffectIdentifier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = EffectIdentifier(*nt)
	return br, err
}

func (v EffectIdentifier) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *EffectIdentifier) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = EffectIdentifier(*a)
	return nil
}

func (v *EffectIdentifier) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = EffectIdentifier(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v EffectIdentifier) String() string {
	switch v {
	case 0x00:
		return "Delayed all off"
	case 0x01:
		return "Dying light"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v EffectIdentifier) IsDelayedAllOff() bool { return v == 0x00 }
func (v EffectIdentifier) IsDyingLight() bool    { return v == 0x01 }
func (v *EffectIdentifier) SetDelayedAllOff()    { *v = 0x00 }
func (v *EffectIdentifier) SetDyingLight()       { *v = 0x01 }

func (EffectIdentifier) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Delayed all off"},
		{Value: 0x01, Name: "Dying light"},
	}
}

func (EffectVariant) Name() string        { return `Effect Variant` }
func (EffectVariant) Description() string { return `` }

type EffectVariant zcl.Zenum8

func (v *EffectVariant) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *EffectVariant) Value() zcl.Val     { return v }

func (v EffectVariant) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *EffectVariant) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = EffectVariant(*nt)
	return br, err
}

func (v EffectVariant) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *EffectVariant) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = EffectVariant(*a)
	return nil
}

func (v *EffectVariant) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = EffectVariant(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v EffectVariant) String() string {
	switch v {
	case 0x00:
		return "Fade to off in 0.8 seconds / 20% dim up in 0.5 then fade to off in 1 second"
	case 0x01:
		return "No Fade"
	case 0x02:
		return "50% dim down in 0.8s then fade off in 12s"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v EffectVariant) IsFadeToOffIn08Seconds20DimUpIn05ThenFadeToOffIn1Second() bool {
	return v == 0x00
}
func (v EffectVariant) IsNoFade() bool                                             { return v == 0x01 }
func (v EffectVariant) Is50DimDownIn08SThenFadeOffIn12S() bool                     { return v == 0x02 }
func (v *EffectVariant) SetFadeToOffIn08Seconds20DimUpIn05ThenFadeToOffIn1Second() { *v = 0x00 }
func (v *EffectVariant) SetNoFade()                                                { *v = 0x01 }
func (v *EffectVariant) Set50DimDownIn08SThenFadeOffIn12S()                        { *v = 0x02 }

func (EffectVariant) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Fade to off in 0.8 seconds / 20% dim up in 0.5 then fade to off in 1 second"},
		{Value: 0x01, Name: "No Fade"},
		{Value: 0x02, Name: "50% dim down in 0.8s then fade off in 12s"},
	}
}

const FastPollTimeoutAttr zcl.AttrID = 3

func (FastPollTimeout) ID() zcl.AttrID   { return FastPollTimeoutAttr }
func (FastPollTimeout) Readable() bool   { return true }
func (FastPollTimeout) Writable() bool   { return true }
func (FastPollTimeout) Reportable() bool { return false }
func (FastPollTimeout) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v FastPollTimeout) AttrID() zcl.AttrID   { return v.ID() }
func (v FastPollTimeout) AttrType() zcl.TypeID { return v.TypeID() }
func (v *FastPollTimeout) AttrValue() zcl.Val  { return v.Value() }

func (FastPollTimeout) Name() string        { return `Fast Poll Timeout` }
func (FastPollTimeout) Description() string { return `` }

type FastPollTimeout zcl.Zu16

func (v *FastPollTimeout) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *FastPollTimeout) Value() zcl.Val     { return v }

func (v FastPollTimeout) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *FastPollTimeout) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = FastPollTimeout(*nt)
	return br, err
}

func (v FastPollTimeout) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *FastPollTimeout) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = FastPollTimeout(*a)
	return nil
}

func (v *FastPollTimeout) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = FastPollTimeout(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v FastPollTimeout) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const FastPollTimeoutMaxAttr zcl.AttrID = 6

func (FastPollTimeoutMax) ID() zcl.AttrID   { return FastPollTimeoutMaxAttr }
func (FastPollTimeoutMax) Readable() bool   { return true }
func (FastPollTimeoutMax) Writable() bool   { return false }
func (FastPollTimeoutMax) Reportable() bool { return false }
func (FastPollTimeoutMax) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v FastPollTimeoutMax) AttrID() zcl.AttrID   { return v.ID() }
func (v FastPollTimeoutMax) AttrType() zcl.TypeID { return v.TypeID() }
func (v *FastPollTimeoutMax) AttrValue() zcl.Val  { return v.Value() }

func (FastPollTimeoutMax) Name() string        { return `Fast Poll Timeout Max` }
func (FastPollTimeoutMax) Description() string { return `` }

type FastPollTimeoutMax zcl.Zu16

func (v *FastPollTimeoutMax) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *FastPollTimeoutMax) Value() zcl.Val     { return v }

func (v FastPollTimeoutMax) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *FastPollTimeoutMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = FastPollTimeoutMax(*nt)
	return br, err
}

func (v FastPollTimeoutMax) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *FastPollTimeoutMax) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = FastPollTimeoutMax(*a)
	return nil
}

func (v *FastPollTimeoutMax) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = FastPollTimeoutMax(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v FastPollTimeoutMax) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const GenericDeviceClassAttr zcl.AttrID = 8

func (GenericDeviceClass) ID() zcl.AttrID   { return GenericDeviceClassAttr }
func (GenericDeviceClass) Readable() bool   { return true }
func (GenericDeviceClass) Writable() bool   { return false }
func (GenericDeviceClass) Reportable() bool { return false }
func (GenericDeviceClass) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v GenericDeviceClass) AttrID() zcl.AttrID   { return v.ID() }
func (v GenericDeviceClass) AttrType() zcl.TypeID { return v.TypeID() }
func (v *GenericDeviceClass) AttrValue() zcl.Val  { return v.Value() }

func (GenericDeviceClass) Name() string { return `Generic Device Class` }
func (GenericDeviceClass) Description() string {
	return `defines the field of application of the GenericDeviceType attribute`
}

// GenericDeviceClass defines the field of application of the GenericDeviceType attribute
type GenericDeviceClass zcl.Zenum8

func (v *GenericDeviceClass) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *GenericDeviceClass) Value() zcl.Val     { return v }

func (v GenericDeviceClass) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *GenericDeviceClass) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = GenericDeviceClass(*nt)
	return br, err
}

func (v GenericDeviceClass) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *GenericDeviceClass) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = GenericDeviceClass(*a)
	return nil
}

func (v *GenericDeviceClass) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = GenericDeviceClass(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v GenericDeviceClass) String() string {
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

const GenericDeviceTypeAttr zcl.AttrID = 9

func (GenericDeviceType) ID() zcl.AttrID   { return GenericDeviceTypeAttr }
func (GenericDeviceType) Readable() bool   { return true }
func (GenericDeviceType) Writable() bool   { return false }
func (GenericDeviceType) Reportable() bool { return false }
func (GenericDeviceType) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v GenericDeviceType) AttrID() zcl.AttrID   { return v.ID() }
func (v GenericDeviceType) AttrType() zcl.TypeID { return v.TypeID() }
func (v *GenericDeviceType) AttrValue() zcl.Val  { return v.Value() }

func (GenericDeviceType) Name() string { return `Generic Device Type` }
func (GenericDeviceType) Description() string {
	return `allows an application to show an icon on a rich user interface (e.g. smartphone app)`
}

// GenericDeviceType allows an application to show an icon on a rich user interface (e.g. smartphone app)
type GenericDeviceType zcl.Zenum8

func (v *GenericDeviceType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *GenericDeviceType) Value() zcl.Val     { return v }

func (v GenericDeviceType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *GenericDeviceType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = GenericDeviceType(*nt)
	return br, err
}

func (v GenericDeviceType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *GenericDeviceType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = GenericDeviceType(*a)
	return nil
}

func (v *GenericDeviceType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = GenericDeviceType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v GenericDeviceType) String() string {
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

const GlobalSceneControlAttr zcl.AttrID = 16384

func (GlobalSceneControl) ID() zcl.AttrID   { return GlobalSceneControlAttr }
func (GlobalSceneControl) Readable() bool   { return true }
func (GlobalSceneControl) Writable() bool   { return false }
func (GlobalSceneControl) Reportable() bool { return false }
func (GlobalSceneControl) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v GlobalSceneControl) AttrID() zcl.AttrID   { return v.ID() }
func (v GlobalSceneControl) AttrType() zcl.TypeID { return v.TypeID() }
func (v *GlobalSceneControl) AttrValue() zcl.Val  { return v.Value() }

func (GlobalSceneControl) Name() string        { return `Global Scene Control` }
func (GlobalSceneControl) Description() string { return `` }

type GlobalSceneControl zcl.Zbool

func (v *GlobalSceneControl) TypeID() zcl.TypeID { return new(zcl.Zbool).TypeID() }
func (v *GlobalSceneControl) Value() zcl.Val     { return v }

func (v GlobalSceneControl) MarshalZcl() ([]byte, error) { return zcl.Zbool(v).MarshalZcl() }

func (v *GlobalSceneControl) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*v = GlobalSceneControl(*nt)
	return br, err
}

func (v GlobalSceneControl) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbool(v))
}

func (v *GlobalSceneControl) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbool)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = GlobalSceneControl(*a)
	return nil
}

func (v *GlobalSceneControl) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbool); ok {
		*v = GlobalSceneControl(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v GlobalSceneControl) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(v))
}

func (GroupId) Name() string        { return `Group ID` }
func (GroupId) Description() string { return `` }

type GroupId zcl.Zu16

func (v *GroupId) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *GroupId) Value() zcl.Val     { return v }

func (v GroupId) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *GroupId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = GroupId(*nt)
	return br, err
}

func (v GroupId) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *GroupId) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = GroupId(*a)
	return nil
}

func (v *GroupId) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = GroupId(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v GroupId) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const GroupNameSupportAttr zcl.AttrID = 0

func (GroupNameSupport) ID() zcl.AttrID   { return GroupNameSupportAttr }
func (GroupNameSupport) Readable() bool   { return true }
func (GroupNameSupport) Writable() bool   { return false }
func (GroupNameSupport) Reportable() bool { return false }
func (GroupNameSupport) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v GroupNameSupport) AttrID() zcl.AttrID   { return v.ID() }
func (v GroupNameSupport) AttrType() zcl.TypeID { return v.TypeID() }
func (v *GroupNameSupport) AttrValue() zcl.Val  { return v.Value() }

func (GroupNameSupport) Name() string        { return `Group Name Support` }
func (GroupNameSupport) Description() string { return `` }

type GroupNameSupport zcl.Zbmp8

func (v *GroupNameSupport) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *GroupNameSupport) Value() zcl.Val     { return v }

func (v GroupNameSupport) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *GroupNameSupport) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = GroupNameSupport(*nt)
	return br, err
}

func (v GroupNameSupport) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *GroupNameSupport) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = GroupNameSupport(*a)
	return nil
}

func (v *GroupNameSupport) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = GroupNameSupport(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v GroupNameSupport) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v GroupNameSupport) IsNamesSupported() bool { return zcl.BitmapTest([]byte(v[:]), 7) }
func (v *GroupNameSupport) SetNamesSupported(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 7, b))
}

func (GroupNameSupport) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 7, Name: "Names Supported"},
	}
}

func (GroupCapacity) Name() string { return `Group capacity` }
func (GroupCapacity) Description() string {
	return `specifies remaining number of groups that can be added.
If set to 0xFE, at least one more group can be added (exact number unknown)
If set to 0xFF, it's unknown if any more groups can be added`
}

// GroupCapacity specifies remaining number of groups that can be added.
// If set to 0xFE, at least one more group can be added (exact number unknown)
// If set to 0xFF, it's unknown if any more groups can be added
type GroupCapacity zcl.Zu8

func (v *GroupCapacity) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *GroupCapacity) Value() zcl.Val     { return v }

func (v GroupCapacity) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *GroupCapacity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = GroupCapacity(*nt)
	return br, err
}

func (v GroupCapacity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *GroupCapacity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = GroupCapacity(*a)
	return nil
}

func (v *GroupCapacity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = GroupCapacity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v GroupCapacity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (GroupList) Name() string        { return `Group list` }
func (GroupList) Description() string { return `` }

type GroupList []*zcl.Zu16

func (v *GroupList) TypeID() zcl.TypeID { return new(zcl.Zset).TypeID() }
func (v *GroupList) Value() zcl.Val     { return v }

func (GroupList) ArrayTypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }

func (v *GroupList) ArrayValues() (o []uint16) {
	for _, a := range *v {
		o = append(o, uint16(*a))
	}
	return o
}

func (v *GroupList) SetValues(val []uint16) error {
	*v = []*zcl.Zu16{}
	return v.AddValues(val...)
}

func (v *GroupList) AddValues(val ...uint16) error {
	for _, a := range val {
		nv := zcl.Zu16(a)
		*v = append(*v, &nv)
	}
	return nil
}

func (v GroupList) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, a := range v {
		vars = append(vars, a)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (v *GroupList) UnmarshalZcl(b []byte) ([]byte, error) {
	*v = []*zcl.Zu16{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(zcl.Zu16)
		*v = append(*v, nv)
		return nv
	})
}

func (v *GroupList) SetValue(a zcl.Val) error {
	if nv, ok := a.(*GroupList); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (v GroupList) String() string {
	var s []string
	for _, a := range v {
		s = append(s, zcl.Sprintf("%v", a))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

func (GroupName) Name() string        { return `Group name` }
func (GroupName) Description() string { return `` }

type GroupName zcl.Zcstring

func (v *GroupName) TypeID() zcl.TypeID { return new(zcl.Zcstring).TypeID() }
func (v *GroupName) Value() zcl.Val     { return v }

func (v GroupName) MarshalZcl() ([]byte, error) { return zcl.Zcstring(v).MarshalZcl() }

func (v *GroupName) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*v = GroupName(*nt)
	return br, err
}

func (v GroupName) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zcstring(v))
}

func (v *GroupName) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zcstring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = GroupName(*a)
	return nil
}

func (v *GroupName) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zcstring); ok {
		*v = GroupName(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v GroupName) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v))
}

const HwVersionAttr zcl.AttrID = 3

func (HwVersion) ID() zcl.AttrID   { return HwVersionAttr }
func (HwVersion) Readable() bool   { return true }
func (HwVersion) Writable() bool   { return false }
func (HwVersion) Reportable() bool { return false }
func (HwVersion) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v HwVersion) AttrID() zcl.AttrID   { return v.ID() }
func (v HwVersion) AttrType() zcl.TypeID { return v.TypeID() }
func (v *HwVersion) AttrValue() zcl.Val  { return v.Value() }

func (HwVersion) Name() string        { return `HW Version` }
func (HwVersion) Description() string { return `` }

type HwVersion zcl.Zu8

func (v *HwVersion) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *HwVersion) Value() zcl.Val     { return v }

func (v HwVersion) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *HwVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = HwVersion(*nt)
	return br, err
}

func (v HwVersion) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *HwVersion) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = HwVersion(*a)
	return nil
}

func (v *HwVersion) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = HwVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v HwVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const HighTempDwellTripPointAttr zcl.AttrID = 20

func (HighTempDwellTripPoint) ID() zcl.AttrID   { return HighTempDwellTripPointAttr }
func (HighTempDwellTripPoint) Readable() bool   { return false }
func (HighTempDwellTripPoint) Writable() bool   { return false }
func (HighTempDwellTripPoint) Reportable() bool { return false }
func (HighTempDwellTripPoint) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v HighTempDwellTripPoint) AttrID() zcl.AttrID   { return v.ID() }
func (v HighTempDwellTripPoint) AttrType() zcl.TypeID { return v.TypeID() }
func (v *HighTempDwellTripPoint) AttrValue() zcl.Val  { return v.Value() }

func (HighTempDwellTripPoint) Name() string        { return `High Temp Dwell Trip Point` }
func (HighTempDwellTripPoint) Description() string { return `` }

type HighTempDwellTripPoint zcl.Zu24

func (v *HighTempDwellTripPoint) TypeID() zcl.TypeID { return new(zcl.Zu24).TypeID() }
func (v *HighTempDwellTripPoint) Value() zcl.Val     { return v }

func (v HighTempDwellTripPoint) MarshalZcl() ([]byte, error) { return zcl.Zu24(v).MarshalZcl() }

func (v *HighTempDwellTripPoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*v = HighTempDwellTripPoint(*nt)
	return br, err
}

func (v HighTempDwellTripPoint) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu24(v))
}

func (v *HighTempDwellTripPoint) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu24)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = HighTempDwellTripPoint(*a)
	return nil
}

func (v *HighTempDwellTripPoint) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu24); ok {
		*v = HighTempDwellTripPoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v HighTempDwellTripPoint) String() string {
	return zcl.Seconds.Format(float64(v))
}

const HighTempThresholdAttr zcl.AttrID = 18

func (HighTempThreshold) ID() zcl.AttrID   { return HighTempThresholdAttr }
func (HighTempThreshold) Readable() bool   { return true }
func (HighTempThreshold) Writable() bool   { return true }
func (HighTempThreshold) Reportable() bool { return false }
func (HighTempThreshold) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v HighTempThreshold) AttrID() zcl.AttrID   { return v.ID() }
func (v HighTempThreshold) AttrType() zcl.TypeID { return v.TypeID() }
func (v *HighTempThreshold) AttrValue() zcl.Val  { return v.Value() }

func (HighTempThreshold) Name() string { return `High Temp Threshold` }
func (HighTempThreshold) Description() string {
	return `If the current temperature goes above the threshold for longer
than the time specified by high temp dwell, an alarm will be triggered`
}

// HighTempThreshold If the current temperature goes above the threshold for longer
// than the time specified by high temp dwell, an alarm will be triggered
type HighTempThreshold zcl.Zs16

func (v *HighTempThreshold) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *HighTempThreshold) Value() zcl.Val     { return v }

func (v HighTempThreshold) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *HighTempThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = HighTempThreshold(*nt)
	return br, err
}

func (v HighTempThreshold) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *HighTempThreshold) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = HighTempThreshold(*a)
	return nil
}

func (v *HighTempThreshold) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = HighTempThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v HighTempThreshold) String() string {
	return zcl.DegreesCelsius.Format(float64(v))
}

const IOApplicationTypeAttr zcl.AttrID = 256

func (IOApplicationType) ID() zcl.AttrID   { return IOApplicationTypeAttr }
func (IOApplicationType) Readable() bool   { return true }
func (IOApplicationType) Writable() bool   { return false }
func (IOApplicationType) Reportable() bool { return false }
func (IOApplicationType) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v IOApplicationType) AttrID() zcl.AttrID   { return v.ID() }
func (v IOApplicationType) AttrType() zcl.TypeID { return v.TypeID() }
func (v *IOApplicationType) AttrValue() zcl.Val  { return v.Value() }

func (IOApplicationType) Name() string        { return `I/O Application Type` }
func (IOApplicationType) Description() string { return `` }

type IOApplicationType zcl.Zu32

func (v *IOApplicationType) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *IOApplicationType) Value() zcl.Val     { return v }

func (v IOApplicationType) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *IOApplicationType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = IOApplicationType(*nt)
	return br, err
}

func (v IOApplicationType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *IOApplicationType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IOApplicationType(*a)
	return nil
}

func (v *IOApplicationType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = IOApplicationType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IOApplicationType) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

const IODescriptionAttr zcl.AttrID = 28

func (IODescription) ID() zcl.AttrID   { return IODescriptionAttr }
func (IODescription) Readable() bool   { return true }
func (IODescription) Writable() bool   { return true }
func (IODescription) Reportable() bool { return false }
func (IODescription) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v IODescription) AttrID() zcl.AttrID   { return v.ID() }
func (v IODescription) AttrType() zcl.TypeID { return v.TypeID() }
func (v *IODescription) AttrValue() zcl.Val  { return v.Value() }

func (IODescription) Name() string        { return `I/O Description` }
func (IODescription) Description() string { return `` }

type IODescription zcl.Zcstring

func (v *IODescription) TypeID() zcl.TypeID { return new(zcl.Zcstring).TypeID() }
func (v *IODescription) Value() zcl.Val     { return v }

func (v IODescription) MarshalZcl() ([]byte, error) { return zcl.Zcstring(v).MarshalZcl() }

func (v *IODescription) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*v = IODescription(*nt)
	return br, err
}

func (v IODescription) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zcstring(v))
}

func (v *IODescription) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zcstring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IODescription(*a)
	return nil
}

func (v *IODescription) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zcstring); ok {
		*v = IODescription(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IODescription) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v))
}

const IOOutOfServiceAttr zcl.AttrID = 81

func (IOOutOfService) ID() zcl.AttrID   { return IOOutOfServiceAttr }
func (IOOutOfService) Readable() bool   { return true }
func (IOOutOfService) Writable() bool   { return true }
func (IOOutOfService) Reportable() bool { return false }
func (IOOutOfService) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v IOOutOfService) AttrID() zcl.AttrID   { return v.ID() }
func (v IOOutOfService) AttrType() zcl.TypeID { return v.TypeID() }
func (v *IOOutOfService) AttrValue() zcl.Val  { return v.Value() }

func (IOOutOfService) Name() string        { return `I/O Out of service` }
func (IOOutOfService) Description() string { return `` }

type IOOutOfService zcl.Zbool

func (v *IOOutOfService) TypeID() zcl.TypeID { return new(zcl.Zbool).TypeID() }
func (v *IOOutOfService) Value() zcl.Val     { return v }

func (v IOOutOfService) MarshalZcl() ([]byte, error) { return zcl.Zbool(v).MarshalZcl() }

func (v *IOOutOfService) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*v = IOOutOfService(*nt)
	return br, err
}

func (v IOOutOfService) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbool(v))
}

func (v *IOOutOfService) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbool)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IOOutOfService(*a)
	return nil
}

func (v *IOOutOfService) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbool); ok {
		*v = IOOutOfService(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IOOutOfService) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(v))
}

const IOReliabilityAttr zcl.AttrID = 103

func (IOReliability) ID() zcl.AttrID   { return IOReliabilityAttr }
func (IOReliability) Readable() bool   { return true }
func (IOReliability) Writable() bool   { return true }
func (IOReliability) Reportable() bool { return false }
func (IOReliability) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v IOReliability) AttrID() zcl.AttrID   { return v.ID() }
func (v IOReliability) AttrType() zcl.TypeID { return v.TypeID() }
func (v *IOReliability) AttrValue() zcl.Val  { return v.Value() }

func (IOReliability) Name() string        { return `I/O Reliability` }
func (IOReliability) Description() string { return `` }

type IOReliability zcl.Zenum8

func (v *IOReliability) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *IOReliability) Value() zcl.Val     { return v }

func (v IOReliability) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *IOReliability) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = IOReliability(*nt)
	return br, err
}

func (v IOReliability) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *IOReliability) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IOReliability(*a)
	return nil
}

func (v *IOReliability) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = IOReliability(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IOReliability) String() string {
	switch v {
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
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v IOReliability) IsNoFaultDetected() bool    { return v == 0x00 }
func (v IOReliability) IsNoSensor() bool           { return v == 0x01 }
func (v IOReliability) IsOverRange() bool          { return v == 0x02 }
func (v IOReliability) IsUnderRange() bool         { return v == 0x03 }
func (v IOReliability) IsOpenLoop() bool           { return v == 0x04 }
func (v IOReliability) IsShortedLoop() bool        { return v == 0x05 }
func (v IOReliability) IsNoOutput() bool           { return v == 0x06 }
func (v IOReliability) IsUnreliableOther() bool    { return v == 0x07 }
func (v IOReliability) IsProcessError() bool       { return v == 0x08 }
func (v IOReliability) IsMultiStateFault() bool    { return v == 0x09 }
func (v IOReliability) IsConfigurationError() bool { return v == 0x0A }
func (v *IOReliability) SetNoFaultDetected()       { *v = 0x00 }
func (v *IOReliability) SetNoSensor()              { *v = 0x01 }
func (v *IOReliability) SetOverRange()             { *v = 0x02 }
func (v *IOReliability) SetUnderRange()            { *v = 0x03 }
func (v *IOReliability) SetOpenLoop()              { *v = 0x04 }
func (v *IOReliability) SetShortedLoop()           { *v = 0x05 }
func (v *IOReliability) SetNoOutput()              { *v = 0x06 }
func (v *IOReliability) SetUnreliableOther()       { *v = 0x07 }
func (v *IOReliability) SetProcessError()          { *v = 0x08 }
func (v *IOReliability) SetMultiStateFault()       { *v = 0x09 }
func (v *IOReliability) SetConfigurationError()    { *v = 0x0A }

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

const IOStatusFlagsAttr zcl.AttrID = 111

func (IOStatusFlags) ID() zcl.AttrID   { return IOStatusFlagsAttr }
func (IOStatusFlags) Readable() bool   { return true }
func (IOStatusFlags) Writable() bool   { return false }
func (IOStatusFlags) Reportable() bool { return true }
func (IOStatusFlags) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v IOStatusFlags) AttrID() zcl.AttrID   { return v.ID() }
func (v IOStatusFlags) AttrType() zcl.TypeID { return v.TypeID() }
func (v *IOStatusFlags) AttrValue() zcl.Val  { return v.Value() }

func (IOStatusFlags) Name() string        { return `I/O Status flags` }
func (IOStatusFlags) Description() string { return `` }

type IOStatusFlags zcl.Zbmp8

func (v *IOStatusFlags) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *IOStatusFlags) Value() zcl.Val     { return v }

func (v IOStatusFlags) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *IOStatusFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = IOStatusFlags(*nt)
	return br, err
}

func (v IOStatusFlags) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *IOStatusFlags) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IOStatusFlags(*a)
	return nil
}

func (v *IOStatusFlags) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = IOStatusFlags(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IOStatusFlags) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v IOStatusFlags) IsInAlarm() bool         { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v IOStatusFlags) IsFault() bool           { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v IOStatusFlags) IsOveridden() bool       { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v IOStatusFlags) IsOutOfService() bool    { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v *IOStatusFlags) SetInAlarm(b bool)      { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b)) }
func (v *IOStatusFlags) SetFault(b bool)        { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b)) }
func (v *IOStatusFlags) SetOveridden(b bool)    { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b)) }
func (v *IOStatusFlags) SetOutOfService(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b)) }

func (IOStatusFlags) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "In Alarm"},
		{Value: 1, Name: "Fault"},
		{Value: 2, Name: "Overidden"},
		{Value: 3, Name: "Out of Service"},
	}
}

const IOUnitTypeAttr zcl.AttrID = 117

func (IOUnitType) ID() zcl.AttrID   { return IOUnitTypeAttr }
func (IOUnitType) Readable() bool   { return true }
func (IOUnitType) Writable() bool   { return true }
func (IOUnitType) Reportable() bool { return false }
func (IOUnitType) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v IOUnitType) AttrID() zcl.AttrID   { return v.ID() }
func (v IOUnitType) AttrType() zcl.TypeID { return v.TypeID() }
func (v *IOUnitType) AttrValue() zcl.Val  { return v.Value() }

func (IOUnitType) Name() string        { return `I/O Unit Type` }
func (IOUnitType) Description() string { return `` }

type IOUnitType zcl.EngineeringUnit

func (v *IOUnitType) TypeID() zcl.TypeID { return new(zcl.EngineeringUnit).TypeID() }
func (v *IOUnitType) Value() zcl.Val     { return v }

func (v IOUnitType) MarshalZcl() ([]byte, error) { return zcl.EngineeringUnit(v).MarshalZcl() }

func (v *IOUnitType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.EngineeringUnit)
	br, err := nt.UnmarshalZcl(b)
	*v = IOUnitType(*nt)
	return br, err
}

func (v IOUnitType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.EngineeringUnit(v))
}

func (v *IOUnitType) UnmarshalJSON(b []byte) error {
	a := new(zcl.EngineeringUnit)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IOUnitType(*a)
	return nil
}

func (v *IOUnitType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.EngineeringUnit); ok {
		*v = IOUnitType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IOUnitType) String() string {
	return zcl.Sprintf("%v", zcl.EngineeringUnit(v))
}

func (IdentifyEffect) Name() string { return `Identify Effect` }
func (IdentifyEffect) Description() string {
	return `The effect identifier field specifies the identify effect to use.`
}

// IdentifyEffect The effect identifier field specifies the identify effect to use.
type IdentifyEffect zcl.Zenum8

func (v *IdentifyEffect) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *IdentifyEffect) Value() zcl.Val     { return v }

func (v IdentifyEffect) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *IdentifyEffect) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = IdentifyEffect(*nt)
	return br, err
}

func (v IdentifyEffect) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *IdentifyEffect) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IdentifyEffect(*a)
	return nil
}

func (v *IdentifyEffect) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = IdentifyEffect(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IdentifyEffect) String() string {
	switch v {
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
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v IdentifyEffect) IsBlink() bool         { return v == 0x00 }
func (v IdentifyEffect) IsBreathe() bool       { return v == 0x01 }
func (v IdentifyEffect) IsOkay() bool          { return v == 0x02 }
func (v IdentifyEffect) IsChannelChange() bool { return v == 0x0B }
func (v IdentifyEffect) IsFinish() bool        { return v == 0xFE }
func (v IdentifyEffect) IsStop() bool          { return v == 0xFF }
func (v *IdentifyEffect) SetBlink()            { *v = 0x00 }
func (v *IdentifyEffect) SetBreathe()          { *v = 0x01 }
func (v *IdentifyEffect) SetOkay()             { *v = 0x02 }
func (v *IdentifyEffect) SetChannelChange()    { *v = 0x0B }
func (v *IdentifyEffect) SetFinish()           { *v = 0xFE }
func (v *IdentifyEffect) SetStop()             { *v = 0xFF }

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

func (IdentifyEffectVariant) Name() string { return `Identify Effect variant` }
func (IdentifyEffectVariant) Description() string {
	return `The effect identifier field specifies the identify effect to use.`
}

// IdentifyEffectVariant The effect identifier field specifies the identify effect to use.
type IdentifyEffectVariant zcl.Zenum8

func (v *IdentifyEffectVariant) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *IdentifyEffectVariant) Value() zcl.Val     { return v }

func (v IdentifyEffectVariant) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *IdentifyEffectVariant) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = IdentifyEffectVariant(*nt)
	return br, err
}

func (v IdentifyEffectVariant) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *IdentifyEffectVariant) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IdentifyEffectVariant(*a)
	return nil
}

func (v *IdentifyEffectVariant) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = IdentifyEffectVariant(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IdentifyEffectVariant) String() string {
	switch v {
	case 0x00:
		return "Default"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v IdentifyEffectVariant) IsDefault() bool { return v == 0x00 }
func (v *IdentifyEffectVariant) SetDefault()    { *v = 0x00 }

func (IdentifyEffectVariant) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Default"},
	}
}

const IdentifyTimeAttr zcl.AttrID = 0

func (IdentifyTime) ID() zcl.AttrID   { return IdentifyTimeAttr }
func (IdentifyTime) Readable() bool   { return true }
func (IdentifyTime) Writable() bool   { return true }
func (IdentifyTime) Reportable() bool { return false }
func (IdentifyTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v IdentifyTime) AttrID() zcl.AttrID   { return v.ID() }
func (v IdentifyTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *IdentifyTime) AttrValue() zcl.Val  { return v.Value() }

func (IdentifyTime) Name() string { return `Identify Time` }
func (IdentifyTime) Description() string {
	return `The time in seconds for which a device will stay in identify mode.`
}

// IdentifyTime The time in seconds for which a device will stay in identify mode.
type IdentifyTime zcl.Zu16

func (v *IdentifyTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *IdentifyTime) Value() zcl.Val     { return v }

func (v IdentifyTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *IdentifyTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = IdentifyTime(*nt)
	return br, err
}

func (v IdentifyTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *IdentifyTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IdentifyTime(*a)
	return nil
}

func (v *IdentifyTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = IdentifyTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IdentifyTime) String() string {
	return zcl.Seconds.Format(float64(v))
}

func (IdentifyTimeout) Name() string { return `Identify Timeout` }
func (IdentifyTimeout) Description() string {
	return `The time in seconds for which a device will stay in identify mode.`
}

// IdentifyTimeout The time in seconds for which a device will stay in identify mode.
type IdentifyTimeout zcl.Zu16

func (v *IdentifyTimeout) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *IdentifyTimeout) Value() zcl.Val     { return v }

func (v IdentifyTimeout) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *IdentifyTimeout) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = IdentifyTimeout(*nt)
	return br, err
}

func (v IdentifyTimeout) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *IdentifyTimeout) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IdentifyTimeout(*a)
	return nil
}

func (v *IdentifyTimeout) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = IdentifyTimeout(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IdentifyTimeout) String() string {
	return zcl.Seconds.Format(float64(v))
}

func (IkeaRemoteDirection) Name() string        { return `Ikea Remote Direction` }
func (IkeaRemoteDirection) Description() string { return `` }

type IkeaRemoteDirection zcl.Zenum8

func (v *IkeaRemoteDirection) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *IkeaRemoteDirection) Value() zcl.Val     { return v }

func (v IkeaRemoteDirection) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *IkeaRemoteDirection) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = IkeaRemoteDirection(*nt)
	return br, err
}

func (v IkeaRemoteDirection) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *IkeaRemoteDirection) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IkeaRemoteDirection(*a)
	return nil
}

func (v *IkeaRemoteDirection) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = IkeaRemoteDirection(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IkeaRemoteDirection) String() string {
	switch v {
	case 0x00:
		return "Right"
	case 0x01:
		return "Left"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v IkeaRemoteDirection) IsRight() bool { return v == 0x00 }
func (v IkeaRemoteDirection) IsLeft() bool  { return v == 0x01 }
func (v *IkeaRemoteDirection) SetRight()    { *v = 0x00 }
func (v *IkeaRemoteDirection) SetLeft()     { *v = 0x01 }

func (IkeaRemoteDirection) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Right"},
		{Value: 0x01, Name: "Left"},
	}
}

const JoinIndicationAttr zcl.AttrID = 272

func (JoinIndication) ID() zcl.AttrID   { return JoinIndicationAttr }
func (JoinIndication) Readable() bool   { return true }
func (JoinIndication) Writable() bool   { return false }
func (JoinIndication) Reportable() bool { return false }
func (JoinIndication) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v JoinIndication) AttrID() zcl.AttrID   { return v.ID() }
func (v JoinIndication) AttrType() zcl.TypeID { return v.TypeID() }
func (v *JoinIndication) AttrValue() zcl.Val  { return v.Value() }

func (JoinIndication) Name() string        { return `Join Indication` }
func (JoinIndication) Description() string { return `` }

type JoinIndication zcl.Zu16

func (v *JoinIndication) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *JoinIndication) Value() zcl.Val     { return v }

func (v JoinIndication) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *JoinIndication) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = JoinIndication(*nt)
	return br, err
}

func (v JoinIndication) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *JoinIndication) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = JoinIndication(*a)
	return nil
}

func (v *JoinIndication) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = JoinIndication(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v JoinIndication) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const LedIndicationAttr zcl.AttrID = 51

func (LedIndication) ID() zcl.AttrID   { return LedIndicationAttr }
func (LedIndication) Readable() bool   { return true }
func (LedIndication) Writable() bool   { return true }
func (LedIndication) Reportable() bool { return false }
func (LedIndication) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LedIndication) AttrID() zcl.AttrID   { return v.ID() }
func (v LedIndication) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LedIndication) AttrValue() zcl.Val  { return v.Value() }

func (LedIndication) Name() string        { return `LED Indication` }
func (LedIndication) Description() string { return `` }

type LedIndication zcl.Zbool

func (v *LedIndication) TypeID() zcl.TypeID { return new(zcl.Zbool).TypeID() }
func (v *LedIndication) Value() zcl.Val     { return v }

func (v LedIndication) MarshalZcl() ([]byte, error) { return zcl.Zbool(v).MarshalZcl() }

func (v *LedIndication) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*v = LedIndication(*nt)
	return br, err
}

func (v LedIndication) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbool(v))
}

func (v *LedIndication) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbool)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LedIndication(*a)
	return nil
}

func (v *LedIndication) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbool); ok {
		*v = LedIndication(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LedIndication) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(v))
}

const LastMessageLqiAttr zcl.AttrID = 284

func (LastMessageLqi) ID() zcl.AttrID   { return LastMessageLqiAttr }
func (LastMessageLqi) Readable() bool   { return true }
func (LastMessageLqi) Writable() bool   { return false }
func (LastMessageLqi) Reportable() bool { return false }
func (LastMessageLqi) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LastMessageLqi) AttrID() zcl.AttrID   { return v.ID() }
func (v LastMessageLqi) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LastMessageLqi) AttrValue() zcl.Val  { return v.Value() }

func (LastMessageLqi) Name() string        { return `Last Message LQI` }
func (LastMessageLqi) Description() string { return `` }

type LastMessageLqi zcl.Zu8

func (v *LastMessageLqi) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *LastMessageLqi) Value() zcl.Val     { return v }

func (v LastMessageLqi) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *LastMessageLqi) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = LastMessageLqi(*nt)
	return br, err
}

func (v LastMessageLqi) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *LastMessageLqi) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LastMessageLqi(*a)
	return nil
}

func (v *LastMessageLqi) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = LastMessageLqi(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LastMessageLqi) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const LastMessageRssiAttr zcl.AttrID = 285

func (LastMessageRssi) ID() zcl.AttrID   { return LastMessageRssiAttr }
func (LastMessageRssi) Readable() bool   { return true }
func (LastMessageRssi) Writable() bool   { return false }
func (LastMessageRssi) Reportable() bool { return false }
func (LastMessageRssi) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LastMessageRssi) AttrID() zcl.AttrID   { return v.ID() }
func (v LastMessageRssi) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LastMessageRssi) AttrValue() zcl.Val  { return v.Value() }

func (LastMessageRssi) Name() string        { return `Last Message RSSI` }
func (LastMessageRssi) Description() string { return `` }

type LastMessageRssi zcl.Zs8

func (v *LastMessageRssi) TypeID() zcl.TypeID { return new(zcl.Zs8).TypeID() }
func (v *LastMessageRssi) Value() zcl.Val     { return v }

func (v LastMessageRssi) MarshalZcl() ([]byte, error) { return zcl.Zs8(v).MarshalZcl() }

func (v *LastMessageRssi) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*v = LastMessageRssi(*nt)
	return br, err
}

func (v LastMessageRssi) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs8(v))
}

func (v *LastMessageRssi) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LastMessageRssi(*a)
	return nil
}

func (v *LastMessageRssi) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs8); ok {
		*v = LastMessageRssi(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LastMessageRssi) String() string {
	return zcl.Sprintf("%v", zcl.Zs8(v))
}

const LastSetTimeAttr zcl.AttrID = 8

func (LastSetTime) ID() zcl.AttrID   { return LastSetTimeAttr }
func (LastSetTime) Readable() bool   { return true }
func (LastSetTime) Writable() bool   { return false }
func (LastSetTime) Reportable() bool { return false }
func (LastSetTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LastSetTime) AttrID() zcl.AttrID   { return v.ID() }
func (v LastSetTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LastSetTime) AttrValue() zcl.Val  { return v.Value() }

func (LastSetTime) Name() string        { return `Last Set Time` }
func (LastSetTime) Description() string { return `` }

type LastSetTime zcl.Zutc

func (v *LastSetTime) TypeID() zcl.TypeID { return new(zcl.Zutc).TypeID() }
func (v *LastSetTime) Value() zcl.Val     { return v }

func (v LastSetTime) MarshalZcl() ([]byte, error) { return zcl.Zutc(v).MarshalZcl() }

func (v *LastSetTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*v = LastSetTime(*nt)
	return br, err
}

func (v LastSetTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zutc(v))
}

func (v *LastSetTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zutc)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LastSetTime(*a)
	return nil
}

func (v *LastSetTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zutc); ok {
		*v = LastSetTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LastSetTime) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(v))
}

func (Level) Name() string        { return `Level` }
func (Level) Description() string { return `` }

type Level zcl.Zu8

func (v *Level) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *Level) Value() zcl.Val     { return v }

func (v Level) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *Level) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = Level(*nt)
	return br, err
}

func (v Level) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *Level) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Level(*a)
	return nil
}

func (v *Level) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = Level(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Level) String() string {
	return zcl.Percent.Format(float64(v) / 2.54)
}

const LevelControlOptionsAttr zcl.AttrID = 15

func (LevelControlOptions) ID() zcl.AttrID   { return LevelControlOptionsAttr }
func (LevelControlOptions) Readable() bool   { return true }
func (LevelControlOptions) Writable() bool   { return true }
func (LevelControlOptions) Reportable() bool { return false }
func (LevelControlOptions) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LevelControlOptions) AttrID() zcl.AttrID   { return v.ID() }
func (v LevelControlOptions) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LevelControlOptions) AttrValue() zcl.Val  { return v.Value() }

func (LevelControlOptions) Name() string { return `Level Control Options` }
func (LevelControlOptions) Description() string {
	return `is a bitmap that determines the default behavior of some cluster commands`
}

// LevelControlOptions is a bitmap that determines the default behavior of some cluster commands
type LevelControlOptions zcl.Zbmp8

func (v *LevelControlOptions) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *LevelControlOptions) Value() zcl.Val     { return v }

func (v LevelControlOptions) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *LevelControlOptions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = LevelControlOptions(*nt)
	return br, err
}

func (v LevelControlOptions) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *LevelControlOptions) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LevelControlOptions(*a)
	return nil
}

func (v *LevelControlOptions) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = LevelControlOptions(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LevelControlOptions) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v LevelControlOptions) IsExecuteIfOff() bool { return zcl.BitmapTest([]byte(v[:]), 0x00) }
func (v *LevelControlOptions) SetExecuteIfOff(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0x00, b))
}

func (LevelControlOptions) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Execute if off"},
	}
}

func (LevelDirection) Name() string        { return `Level direction` }
func (LevelDirection) Description() string { return `` }

type LevelDirection zcl.Zenum8

func (v *LevelDirection) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *LevelDirection) Value() zcl.Val     { return v }

func (v LevelDirection) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *LevelDirection) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = LevelDirection(*nt)
	return br, err
}

func (v LevelDirection) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *LevelDirection) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LevelDirection(*a)
	return nil
}

func (v *LevelDirection) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = LevelDirection(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LevelDirection) String() string {
	switch v {
	case 0x00:
		return "Up"
	case 0x01:
		return "Down"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v LevelDirection) IsUp() bool   { return v == 0x00 }
func (v LevelDirection) IsDown() bool { return v == 0x01 }
func (v *LevelDirection) SetUp()      { *v = 0x00 }
func (v *LevelDirection) SetDown()    { *v = 0x01 }

func (LevelDirection) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Up"},
		{Value: 0x01, Name: "Down"},
	}
}

const LocalTimeAttr zcl.AttrID = 7

func (LocalTime) ID() zcl.AttrID   { return LocalTimeAttr }
func (LocalTime) Readable() bool   { return true }
func (LocalTime) Writable() bool   { return false }
func (LocalTime) Reportable() bool { return false }
func (LocalTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LocalTime) AttrID() zcl.AttrID   { return v.ID() }
func (v LocalTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LocalTime) AttrValue() zcl.Val  { return v.Value() }

func (LocalTime) Name() string        { return `Local Time` }
func (LocalTime) Description() string { return `Local time` }

// LocalTime Local time
type LocalTime zcl.Zu32

func (v *LocalTime) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *LocalTime) Value() zcl.Val     { return v }

func (v LocalTime) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *LocalTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = LocalTime(*nt)
	return br, err
}

func (v LocalTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *LocalTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LocalTime(*a)
	return nil
}

func (v *LocalTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = LocalTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LocalTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

const LocationAgeAttr zcl.AttrID = 2

func (LocationAge) ID() zcl.AttrID   { return LocationAgeAttr }
func (LocationAge) Readable() bool   { return true }
func (LocationAge) Writable() bool   { return false }
func (LocationAge) Reportable() bool { return false }
func (LocationAge) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LocationAge) AttrID() zcl.AttrID   { return v.ID() }
func (v LocationAge) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LocationAge) AttrValue() zcl.Val  { return v.Value() }

func (LocationAge) Name() string        { return `Location Age` }
func (LocationAge) Description() string { return `` }

type LocationAge zcl.Zu16

func (v *LocationAge) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *LocationAge) Value() zcl.Val     { return v }

func (v LocationAge) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *LocationAge) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = LocationAge(*nt)
	return br, err
}

func (v LocationAge) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *LocationAge) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LocationAge(*a)
	return nil
}

func (v *LocationAge) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = LocationAge(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LocationAge) String() string {
	return zcl.Seconds.Format(float64(v))
}

const LocationDescriptionAttr zcl.AttrID = 16

func (LocationDescription) ID() zcl.AttrID   { return LocationDescriptionAttr }
func (LocationDescription) Readable() bool   { return true }
func (LocationDescription) Writable() bool   { return true }
func (LocationDescription) Reportable() bool { return false }
func (LocationDescription) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LocationDescription) AttrID() zcl.AttrID   { return v.ID() }
func (v LocationDescription) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LocationDescription) AttrValue() zcl.Val  { return v.Value() }

func (LocationDescription) Name() string        { return `Location Description` }
func (LocationDescription) Description() string { return `` }

type LocationDescription zcl.Zcstring

func (v *LocationDescription) TypeID() zcl.TypeID { return new(zcl.Zcstring).TypeID() }
func (v *LocationDescription) Value() zcl.Val     { return v }

func (v LocationDescription) MarshalZcl() ([]byte, error) { return zcl.Zcstring(v).MarshalZcl() }

func (v *LocationDescription) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*v = LocationDescription(*nt)
	return br, err
}

func (v LocationDescription) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zcstring(v))
}

func (v *LocationDescription) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zcstring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LocationDescription(*a)
	return nil
}

func (v *LocationDescription) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zcstring); ok {
		*v = LocationDescription(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LocationDescription) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v))
}

const LocationMethodAttr zcl.AttrID = 1

func (LocationMethod) ID() zcl.AttrID   { return LocationMethodAttr }
func (LocationMethod) Readable() bool   { return true }
func (LocationMethod) Writable() bool   { return true }
func (LocationMethod) Reportable() bool { return false }
func (LocationMethod) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LocationMethod) AttrID() zcl.AttrID   { return v.ID() }
func (v LocationMethod) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LocationMethod) AttrValue() zcl.Val  { return v.Value() }

func (LocationMethod) Name() string        { return `Location Method` }
func (LocationMethod) Description() string { return `` }

type LocationMethod zcl.Zenum8

func (v *LocationMethod) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *LocationMethod) Value() zcl.Val     { return v }

func (v LocationMethod) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *LocationMethod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = LocationMethod(*nt)
	return br, err
}

func (v LocationMethod) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *LocationMethod) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LocationMethod(*a)
	return nil
}

func (v *LocationMethod) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = LocationMethod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LocationMethod) String() string {
	switch v {
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
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v LocationMethod) IsLateration() bool       { return v == 0x00 }
func (v LocationMethod) IsSignposting() bool      { return v == 0x01 }
func (v LocationMethod) IsRfFingerprinting() bool { return v == 0x02 }
func (v LocationMethod) IsOutOfBand() bool        { return v == 0x03 }
func (v LocationMethod) IsCentralized() bool      { return v == 0x04 }
func (v *LocationMethod) SetLateration()          { *v = 0x00 }
func (v *LocationMethod) SetSignposting()         { *v = 0x01 }
func (v *LocationMethod) SetRfFingerprinting()    { *v = 0x02 }
func (v *LocationMethod) SetOutOfBand()           { *v = 0x03 }
func (v *LocationMethod) SetCentralized()         { *v = 0x04 }

func (LocationMethod) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Lateration"},
		{Value: 0x01, Name: "Signposting"},
		{Value: 0x02, Name: "RF fingerprinting"},
		{Value: 0x03, Name: "Out of band"},
		{Value: 0x04, Name: "Centralized"},
	}
}

const LocationTypeAttr zcl.AttrID = 0

func (LocationType) ID() zcl.AttrID   { return LocationTypeAttr }
func (LocationType) Readable() bool   { return true }
func (LocationType) Writable() bool   { return true }
func (LocationType) Reportable() bool { return false }
func (LocationType) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LocationType) AttrID() zcl.AttrID   { return v.ID() }
func (v LocationType) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LocationType) AttrValue() zcl.Val  { return v.Value() }

func (LocationType) Name() string        { return `Location Type` }
func (LocationType) Description() string { return `` }

type LocationType zcl.Zenum8

func (v *LocationType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *LocationType) Value() zcl.Val     { return v }

func (v LocationType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *LocationType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = LocationType(*nt)
	return br, err
}

func (v LocationType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *LocationType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LocationType(*a)
	return nil
}

func (v *LocationType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = LocationType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LocationType) String() string {
	switch v {
	case 0x00:
		return "3D Location"
	case 0x01:
		return "Absolute 3D Location"
	case 0x02:
		return "2D Location"
	case 0x03:
		return "Absolute 2D Location"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v LocationType) Is3DLocation() bool         { return v == 0x00 }
func (v LocationType) IsAbsolute3DLocation() bool { return v == 0x01 }
func (v LocationType) Is2DLocation() bool         { return v == 0x02 }
func (v LocationType) IsAbsolute2DLocation() bool { return v == 0x03 }
func (v *LocationType) Set3DLocation()            { *v = 0x00 }
func (v *LocationType) SetAbsolute3DLocation()    { *v = 0x01 }
func (v *LocationType) Set2DLocation()            { *v = 0x02 }
func (v *LocationType) SetAbsolute2DLocation()    { *v = 0x03 }

func (LocationType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "3D Location"},
		{Value: 0x01, Name: "Absolute 3D Location"},
		{Value: 0x02, Name: "2D Location"},
		{Value: 0x03, Name: "Absolute 2D Location"},
	}
}

func (LocationFlags) Name() string        { return `Location flags` }
func (LocationFlags) Description() string { return `` }

type LocationFlags zcl.Zbmp8

func (v *LocationFlags) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *LocationFlags) Value() zcl.Val     { return v }

func (v LocationFlags) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *LocationFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = LocationFlags(*nt)
	return br, err
}

func (v LocationFlags) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *LocationFlags) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LocationFlags(*a)
	return nil
}

func (v *LocationFlags) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = LocationFlags(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LocationFlags) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v LocationFlags) IsAbsoluteOnly() bool       { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v LocationFlags) IsRecalculate() bool        { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v LocationFlags) IsBroadcastIndicator() bool { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v LocationFlags) IsBroadcastResponse() bool  { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v LocationFlags) IsCompactResponse() bool    { return zcl.BitmapTest([]byte(v[:]), 4) }
func (v *LocationFlags) SetAbsoluteOnly(b bool)    { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b)) }
func (v *LocationFlags) SetRecalculate(b bool)     { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b)) }
func (v *LocationFlags) SetBroadcastIndicator(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}
func (v *LocationFlags) SetBroadcastResponse(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b))
}
func (v *LocationFlags) SetCompactResponse(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 4, b))
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

const LongPollIntervalAttr zcl.AttrID = 1

func (LongPollInterval) ID() zcl.AttrID   { return LongPollIntervalAttr }
func (LongPollInterval) Readable() bool   { return true }
func (LongPollInterval) Writable() bool   { return false }
func (LongPollInterval) Reportable() bool { return false }
func (LongPollInterval) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LongPollInterval) AttrID() zcl.AttrID   { return v.ID() }
func (v LongPollInterval) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LongPollInterval) AttrValue() zcl.Val  { return v.Value() }

func (LongPollInterval) Name() string        { return `Long Poll Interval` }
func (LongPollInterval) Description() string { return `` }

type LongPollInterval zcl.Zu32

func (v *LongPollInterval) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *LongPollInterval) Value() zcl.Val     { return v }

func (v LongPollInterval) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *LongPollInterval) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = LongPollInterval(*nt)
	return br, err
}

func (v LongPollInterval) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *LongPollInterval) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LongPollInterval(*a)
	return nil
}

func (v *LongPollInterval) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = LongPollInterval(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LongPollInterval) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

const LongPollIntervalMinAttr zcl.AttrID = 5

func (LongPollIntervalMin) ID() zcl.AttrID   { return LongPollIntervalMinAttr }
func (LongPollIntervalMin) Readable() bool   { return true }
func (LongPollIntervalMin) Writable() bool   { return false }
func (LongPollIntervalMin) Reportable() bool { return false }
func (LongPollIntervalMin) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LongPollIntervalMin) AttrID() zcl.AttrID   { return v.ID() }
func (v LongPollIntervalMin) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LongPollIntervalMin) AttrValue() zcl.Val  { return v.Value() }

func (LongPollIntervalMin) Name() string        { return `Long Poll Interval Min` }
func (LongPollIntervalMin) Description() string { return `` }

type LongPollIntervalMin zcl.Zu32

func (v *LongPollIntervalMin) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *LongPollIntervalMin) Value() zcl.Val     { return v }

func (v LongPollIntervalMin) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *LongPollIntervalMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = LongPollIntervalMin(*nt)
	return br, err
}

func (v LongPollIntervalMin) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *LongPollIntervalMin) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LongPollIntervalMin(*a)
	return nil
}

func (v *LongPollIntervalMin) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = LongPollIntervalMin(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LongPollIntervalMin) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

const LowTempDwellTripPointAttr zcl.AttrID = 19

func (LowTempDwellTripPoint) ID() zcl.AttrID   { return LowTempDwellTripPointAttr }
func (LowTempDwellTripPoint) Readable() bool   { return false }
func (LowTempDwellTripPoint) Writable() bool   { return false }
func (LowTempDwellTripPoint) Reportable() bool { return false }
func (LowTempDwellTripPoint) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LowTempDwellTripPoint) AttrID() zcl.AttrID   { return v.ID() }
func (v LowTempDwellTripPoint) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LowTempDwellTripPoint) AttrValue() zcl.Val  { return v.Value() }

func (LowTempDwellTripPoint) Name() string        { return `Low Temp Dwell Trip Point` }
func (LowTempDwellTripPoint) Description() string { return `` }

type LowTempDwellTripPoint zcl.Zu24

func (v *LowTempDwellTripPoint) TypeID() zcl.TypeID { return new(zcl.Zu24).TypeID() }
func (v *LowTempDwellTripPoint) Value() zcl.Val     { return v }

func (v LowTempDwellTripPoint) MarshalZcl() ([]byte, error) { return zcl.Zu24(v).MarshalZcl() }

func (v *LowTempDwellTripPoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*v = LowTempDwellTripPoint(*nt)
	return br, err
}

func (v LowTempDwellTripPoint) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu24(v))
}

func (v *LowTempDwellTripPoint) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu24)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LowTempDwellTripPoint(*a)
	return nil
}

func (v *LowTempDwellTripPoint) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu24); ok {
		*v = LowTempDwellTripPoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LowTempDwellTripPoint) String() string {
	return zcl.Seconds.Format(float64(v))
}

const LowTempThresholdAttr zcl.AttrID = 17

func (LowTempThreshold) ID() zcl.AttrID   { return LowTempThresholdAttr }
func (LowTempThreshold) Readable() bool   { return true }
func (LowTempThreshold) Writable() bool   { return true }
func (LowTempThreshold) Reportable() bool { return false }
func (LowTempThreshold) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LowTempThreshold) AttrID() zcl.AttrID   { return v.ID() }
func (v LowTempThreshold) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LowTempThreshold) AttrValue() zcl.Val  { return v.Value() }

func (LowTempThreshold) Name() string { return `Low Temp Threshold` }
func (LowTempThreshold) Description() string {
	return `If the current temperature drops below the threshold for longer
than the time specified by low temp dwell, an alarm will be triggered`
}

// LowTempThreshold If the current temperature drops below the threshold for longer
// than the time specified by low temp dwell, an alarm will be triggered
type LowTempThreshold zcl.Zs16

func (v *LowTempThreshold) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *LowTempThreshold) Value() zcl.Val     { return v }

func (v LowTempThreshold) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *LowTempThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = LowTempThreshold(*nt)
	return br, err
}

func (v LowTempThreshold) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *LowTempThreshold) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LowTempThreshold(*a)
	return nil
}

func (v *LowTempThreshold) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = LowTempThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LowTempThreshold) String() string {
	return zcl.DegreesCelsius.Format(float64(v))
}

const MacRxBcastAttr zcl.AttrID = 256

func (MacRxBcast) ID() zcl.AttrID   { return MacRxBcastAttr }
func (MacRxBcast) Readable() bool   { return true }
func (MacRxBcast) Writable() bool   { return false }
func (MacRxBcast) Reportable() bool { return false }
func (MacRxBcast) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MacRxBcast) AttrID() zcl.AttrID   { return v.ID() }
func (v MacRxBcast) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MacRxBcast) AttrValue() zcl.Val  { return v.Value() }

func (MacRxBcast) Name() string        { return `Mac Rx Bcast` }
func (MacRxBcast) Description() string { return `` }

type MacRxBcast zcl.Zu32

func (v *MacRxBcast) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *MacRxBcast) Value() zcl.Val     { return v }

func (v MacRxBcast) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *MacRxBcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = MacRxBcast(*nt)
	return br, err
}

func (v MacRxBcast) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *MacRxBcast) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MacRxBcast(*a)
	return nil
}

func (v *MacRxBcast) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = MacRxBcast(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MacRxBcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

const MacRxUcastAttr zcl.AttrID = 258

func (MacRxUcast) ID() zcl.AttrID   { return MacRxUcastAttr }
func (MacRxUcast) Readable() bool   { return true }
func (MacRxUcast) Writable() bool   { return false }
func (MacRxUcast) Reportable() bool { return false }
func (MacRxUcast) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MacRxUcast) AttrID() zcl.AttrID   { return v.ID() }
func (v MacRxUcast) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MacRxUcast) AttrValue() zcl.Val  { return v.Value() }

func (MacRxUcast) Name() string        { return `Mac Rx Ucast` }
func (MacRxUcast) Description() string { return `` }

type MacRxUcast zcl.Zu32

func (v *MacRxUcast) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *MacRxUcast) Value() zcl.Val     { return v }

func (v MacRxUcast) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *MacRxUcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = MacRxUcast(*nt)
	return br, err
}

func (v MacRxUcast) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *MacRxUcast) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MacRxUcast(*a)
	return nil
}

func (v *MacRxUcast) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = MacRxUcast(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MacRxUcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

const MacTxBcastAttr zcl.AttrID = 257

func (MacTxBcast) ID() zcl.AttrID   { return MacTxBcastAttr }
func (MacTxBcast) Readable() bool   { return true }
func (MacTxBcast) Writable() bool   { return false }
func (MacTxBcast) Reportable() bool { return false }
func (MacTxBcast) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MacTxBcast) AttrID() zcl.AttrID   { return v.ID() }
func (v MacTxBcast) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MacTxBcast) AttrValue() zcl.Val  { return v.Value() }

func (MacTxBcast) Name() string        { return `Mac Tx Bcast` }
func (MacTxBcast) Description() string { return `` }

type MacTxBcast zcl.Zu32

func (v *MacTxBcast) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *MacTxBcast) Value() zcl.Val     { return v }

func (v MacTxBcast) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *MacTxBcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = MacTxBcast(*nt)
	return br, err
}

func (v MacTxBcast) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *MacTxBcast) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MacTxBcast(*a)
	return nil
}

func (v *MacTxBcast) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = MacTxBcast(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MacTxBcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

const MacTxUcastAttr zcl.AttrID = 259

func (MacTxUcast) ID() zcl.AttrID   { return MacTxUcastAttr }
func (MacTxUcast) Readable() bool   { return true }
func (MacTxUcast) Writable() bool   { return false }
func (MacTxUcast) Reportable() bool { return false }
func (MacTxUcast) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MacTxUcast) AttrID() zcl.AttrID   { return v.ID() }
func (v MacTxUcast) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MacTxUcast) AttrValue() zcl.Val  { return v.Value() }

func (MacTxUcast) Name() string        { return `Mac Tx Ucast` }
func (MacTxUcast) Description() string { return `` }

type MacTxUcast zcl.Zu32

func (v *MacTxUcast) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *MacTxUcast) Value() zcl.Val     { return v }

func (v MacTxUcast) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *MacTxUcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = MacTxUcast(*nt)
	return br, err
}

func (v MacTxUcast) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *MacTxUcast) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MacTxUcast(*a)
	return nil
}

func (v *MacTxUcast) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = MacTxUcast(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MacTxUcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

const MacTxUcastFailAttr zcl.AttrID = 261

func (MacTxUcastFail) ID() zcl.AttrID   { return MacTxUcastFailAttr }
func (MacTxUcastFail) Readable() bool   { return true }
func (MacTxUcastFail) Writable() bool   { return false }
func (MacTxUcastFail) Reportable() bool { return false }
func (MacTxUcastFail) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MacTxUcastFail) AttrID() zcl.AttrID   { return v.ID() }
func (v MacTxUcastFail) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MacTxUcastFail) AttrValue() zcl.Val  { return v.Value() }

func (MacTxUcastFail) Name() string        { return `Mac Tx Ucast Fail` }
func (MacTxUcastFail) Description() string { return `` }

type MacTxUcastFail zcl.Zu16

func (v *MacTxUcastFail) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MacTxUcastFail) Value() zcl.Val     { return v }

func (v MacTxUcastFail) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MacTxUcastFail) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MacTxUcastFail(*nt)
	return br, err
}

func (v MacTxUcastFail) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MacTxUcastFail) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MacTxUcastFail(*a)
	return nil
}

func (v *MacTxUcastFail) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MacTxUcastFail(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MacTxUcastFail) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const MacTxUcastRetryAttr zcl.AttrID = 260

func (MacTxUcastRetry) ID() zcl.AttrID   { return MacTxUcastRetryAttr }
func (MacTxUcastRetry) Readable() bool   { return true }
func (MacTxUcastRetry) Writable() bool   { return false }
func (MacTxUcastRetry) Reportable() bool { return false }
func (MacTxUcastRetry) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MacTxUcastRetry) AttrID() zcl.AttrID   { return v.ID() }
func (v MacTxUcastRetry) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MacTxUcastRetry) AttrValue() zcl.Val  { return v.Value() }

func (MacTxUcastRetry) Name() string        { return `Mac Tx Ucast Retry` }
func (MacTxUcastRetry) Description() string { return `` }

type MacTxUcastRetry zcl.Zu16

func (v *MacTxUcastRetry) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MacTxUcastRetry) Value() zcl.Val     { return v }

func (v MacTxUcastRetry) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MacTxUcastRetry) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MacTxUcastRetry(*nt)
	return br, err
}

func (v MacTxUcastRetry) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MacTxUcastRetry) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MacTxUcastRetry(*a)
	return nil
}

func (v *MacTxUcastRetry) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MacTxUcastRetry(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MacTxUcastRetry) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const MainsAlarmMaskAttr zcl.AttrID = 16

func (MainsAlarmMask) ID() zcl.AttrID   { return MainsAlarmMaskAttr }
func (MainsAlarmMask) Readable() bool   { return true }
func (MainsAlarmMask) Writable() bool   { return true }
func (MainsAlarmMask) Reportable() bool { return false }
func (MainsAlarmMask) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MainsAlarmMask) AttrID() zcl.AttrID   { return v.ID() }
func (v MainsAlarmMask) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MainsAlarmMask) AttrValue() zcl.Val  { return v.Value() }

func (MainsAlarmMask) Name() string        { return `Mains Alarm Mask` }
func (MainsAlarmMask) Description() string { return `` }

type MainsAlarmMask zcl.Zbmp8

func (v *MainsAlarmMask) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *MainsAlarmMask) Value() zcl.Val     { return v }

func (v MainsAlarmMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *MainsAlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = MainsAlarmMask(*nt)
	return br, err
}

func (v MainsAlarmMask) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *MainsAlarmMask) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MainsAlarmMask(*a)
	return nil
}

func (v *MainsAlarmMask) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = MainsAlarmMask(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MainsAlarmMask) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v MainsAlarmMask) IsMainsVoltageTooLow() bool  { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v MainsAlarmMask) IsMainsVoltageTooHigh() bool { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v MainsAlarmMask) IsMainsPowerSupplyLostUnavailable() bool {
	return zcl.BitmapTest([]byte(v[:]), 2)
}
func (v *MainsAlarmMask) SetMainsVoltageTooLow(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *MainsAlarmMask) SetMainsVoltageTooHigh(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *MainsAlarmMask) SetMainsPowerSupplyLostUnavailable(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}

func (MainsAlarmMask) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Mains Voltage too low"},
		{Value: 1, Name: "Mains Voltage too high"},
		{Value: 2, Name: "Mains power supply lost/unavailable"},
	}
}

const MainsFrequencyAttr zcl.AttrID = 1

func (MainsFrequency) ID() zcl.AttrID   { return MainsFrequencyAttr }
func (MainsFrequency) Readable() bool   { return true }
func (MainsFrequency) Writable() bool   { return false }
func (MainsFrequency) Reportable() bool { return false }
func (MainsFrequency) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MainsFrequency) AttrID() zcl.AttrID   { return v.ID() }
func (v MainsFrequency) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MainsFrequency) AttrValue() zcl.Val  { return v.Value() }

func (MainsFrequency) Name() string { return `Mains Frequency` }
func (MainsFrequency) Description() string {
	return `Resolution of 2Hz
Special values:
* 0x00 frequency too low to measure
* 0xfe frequency too high to measure
* 0xff unable to measure`
}

// MainsFrequency Resolution of 2Hz
// Special values:
// * 0x00 frequency too low to measure
// * 0xfe frequency too high to measure
// * 0xff unable to measure
type MainsFrequency zcl.Zu8

func (v *MainsFrequency) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *MainsFrequency) Value() zcl.Val     { return v }

func (v MainsFrequency) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *MainsFrequency) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = MainsFrequency(*nt)
	return br, err
}

func (v MainsFrequency) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *MainsFrequency) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MainsFrequency(*a)
	return nil
}

func (v *MainsFrequency) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = MainsFrequency(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MainsFrequency) String() string {
	return zcl.Hertz.Format(float64(v) / 2)
}

const MainsVoltageAttr zcl.AttrID = 0

func (MainsVoltage) ID() zcl.AttrID   { return MainsVoltageAttr }
func (MainsVoltage) Readable() bool   { return true }
func (MainsVoltage) Writable() bool   { return false }
func (MainsVoltage) Reportable() bool { return false }
func (MainsVoltage) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MainsVoltage) AttrID() zcl.AttrID   { return v.ID() }
func (v MainsVoltage) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MainsVoltage) AttrValue() zcl.Val  { return v.Value() }

func (MainsVoltage) Name() string        { return `Mains Voltage` }
func (MainsVoltage) Description() string { return `` }

type MainsVoltage zcl.Zu16

func (v *MainsVoltage) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MainsVoltage) Value() zcl.Val     { return v }

func (v MainsVoltage) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MainsVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MainsVoltage(*nt)
	return br, err
}

func (v MainsVoltage) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MainsVoltage) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MainsVoltage(*a)
	return nil
}

func (v *MainsVoltage) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MainsVoltage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MainsVoltage) String() string {
	return zcl.Volts.Format(float64(v) / 10)
}

const MainsVoltageDwellTripPointAttr zcl.AttrID = 19

func (MainsVoltageDwellTripPoint) ID() zcl.AttrID   { return MainsVoltageDwellTripPointAttr }
func (MainsVoltageDwellTripPoint) Readable() bool   { return true }
func (MainsVoltageDwellTripPoint) Writable() bool   { return true }
func (MainsVoltageDwellTripPoint) Reportable() bool { return false }
func (MainsVoltageDwellTripPoint) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MainsVoltageDwellTripPoint) AttrID() zcl.AttrID   { return v.ID() }
func (v MainsVoltageDwellTripPoint) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MainsVoltageDwellTripPoint) AttrValue() zcl.Val  { return v.Value() }

func (MainsVoltageDwellTripPoint) Name() string { return `Mains Voltage Dwell Trip Point` }
func (MainsVoltageDwellTripPoint) Description() string {
	return `Length of time that the value of MainsVoltage MAY exist beyond either
of its thresholds before an alarm is generated`
}

// MainsVoltageDwellTripPoint Length of time that the value of MainsVoltage MAY exist beyond either
// of its thresholds before an alarm is generated
type MainsVoltageDwellTripPoint zcl.Zu16

func (v *MainsVoltageDwellTripPoint) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MainsVoltageDwellTripPoint) Value() zcl.Val     { return v }

func (v MainsVoltageDwellTripPoint) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MainsVoltageDwellTripPoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MainsVoltageDwellTripPoint(*nt)
	return br, err
}

func (v MainsVoltageDwellTripPoint) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MainsVoltageDwellTripPoint) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MainsVoltageDwellTripPoint(*a)
	return nil
}

func (v *MainsVoltageDwellTripPoint) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MainsVoltageDwellTripPoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MainsVoltageDwellTripPoint) String() string {
	return zcl.Seconds.Format(float64(v))
}

const MainsVoltageMaxThresholdAttr zcl.AttrID = 18

func (MainsVoltageMaxThreshold) ID() zcl.AttrID   { return MainsVoltageMaxThresholdAttr }
func (MainsVoltageMaxThreshold) Readable() bool   { return true }
func (MainsVoltageMaxThreshold) Writable() bool   { return true }
func (MainsVoltageMaxThreshold) Reportable() bool { return false }
func (MainsVoltageMaxThreshold) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MainsVoltageMaxThreshold) AttrID() zcl.AttrID   { return v.ID() }
func (v MainsVoltageMaxThreshold) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MainsVoltageMaxThreshold) AttrValue() zcl.Val  { return v.Value() }

func (MainsVoltageMaxThreshold) Name() string        { return `Mains Voltage Max Threshold` }
func (MainsVoltageMaxThreshold) Description() string { return `` }

type MainsVoltageMaxThreshold zcl.Zu16

func (v *MainsVoltageMaxThreshold) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MainsVoltageMaxThreshold) Value() zcl.Val     { return v }

func (v MainsVoltageMaxThreshold) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MainsVoltageMaxThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MainsVoltageMaxThreshold(*nt)
	return br, err
}

func (v MainsVoltageMaxThreshold) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MainsVoltageMaxThreshold) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MainsVoltageMaxThreshold(*a)
	return nil
}

func (v *MainsVoltageMaxThreshold) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MainsVoltageMaxThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MainsVoltageMaxThreshold) String() string {
	return zcl.Volts.Format(float64(v) / 10)
}

const MainsVoltageMinThresholdAttr zcl.AttrID = 17

func (MainsVoltageMinThreshold) ID() zcl.AttrID   { return MainsVoltageMinThresholdAttr }
func (MainsVoltageMinThreshold) Readable() bool   { return true }
func (MainsVoltageMinThreshold) Writable() bool   { return true }
func (MainsVoltageMinThreshold) Reportable() bool { return false }
func (MainsVoltageMinThreshold) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MainsVoltageMinThreshold) AttrID() zcl.AttrID   { return v.ID() }
func (v MainsVoltageMinThreshold) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MainsVoltageMinThreshold) AttrValue() zcl.Val  { return v.Value() }

func (MainsVoltageMinThreshold) Name() string        { return `Mains Voltage Min Threshold` }
func (MainsVoltageMinThreshold) Description() string { return `` }

type MainsVoltageMinThreshold zcl.Zu16

func (v *MainsVoltageMinThreshold) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MainsVoltageMinThreshold) Value() zcl.Val     { return v }

func (v MainsVoltageMinThreshold) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MainsVoltageMinThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MainsVoltageMinThreshold(*nt)
	return br, err
}

func (v MainsVoltageMinThreshold) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MainsVoltageMinThreshold) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MainsVoltageMinThreshold(*a)
	return nil
}

func (v *MainsVoltageMinThreshold) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MainsVoltageMinThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MainsVoltageMinThreshold) String() string {
	return zcl.Volts.Format(float64(v) / 10)
}

const ManufacturerNameAttr zcl.AttrID = 4

func (ManufacturerName) ID() zcl.AttrID   { return ManufacturerNameAttr }
func (ManufacturerName) Readable() bool   { return true }
func (ManufacturerName) Writable() bool   { return false }
func (ManufacturerName) Reportable() bool { return false }
func (ManufacturerName) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ManufacturerName) AttrID() zcl.AttrID   { return v.ID() }
func (v ManufacturerName) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ManufacturerName) AttrValue() zcl.Val  { return v.Value() }

func (ManufacturerName) Name() string        { return `Manufacturer Name` }
func (ManufacturerName) Description() string { return `` }

type ManufacturerName zcl.Zcstring

func (v *ManufacturerName) TypeID() zcl.TypeID { return new(zcl.Zcstring).TypeID() }
func (v *ManufacturerName) Value() zcl.Val     { return v }

func (v ManufacturerName) MarshalZcl() ([]byte, error) { return zcl.Zcstring(v).MarshalZcl() }

func (v *ManufacturerName) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*v = ManufacturerName(*nt)
	return br, err
}

func (v ManufacturerName) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zcstring(v))
}

func (v *ManufacturerName) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zcstring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ManufacturerName(*a)
	return nil
}

func (v *ManufacturerName) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zcstring); ok {
		*v = ManufacturerName(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ManufacturerName) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v))
}

const MaxTempExperiencedAttr zcl.AttrID = 2

func (MaxTempExperienced) ID() zcl.AttrID   { return MaxTempExperiencedAttr }
func (MaxTempExperienced) Readable() bool   { return true }
func (MaxTempExperienced) Writable() bool   { return false }
func (MaxTempExperienced) Reportable() bool { return false }
func (MaxTempExperienced) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxTempExperienced) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxTempExperienced) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxTempExperienced) AttrValue() zcl.Val  { return v.Value() }

func (MaxTempExperienced) Name() string        { return `Max Temp Experienced` }
func (MaxTempExperienced) Description() string { return `` }

type MaxTempExperienced zcl.Zs16

func (v *MaxTempExperienced) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MaxTempExperienced) Value() zcl.Val     { return v }

func (v MaxTempExperienced) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MaxTempExperienced) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxTempExperienced(*nt)
	return br, err
}

func (v MaxTempExperienced) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MaxTempExperienced) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxTempExperienced(*a)
	return nil
}

func (v *MaxTempExperienced) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MaxTempExperienced(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxTempExperienced) String() string {
	return zcl.DegreesCelsius.Format(float64(v))
}

const MinTempExperiencedAttr zcl.AttrID = 1

func (MinTempExperienced) ID() zcl.AttrID   { return MinTempExperiencedAttr }
func (MinTempExperienced) Readable() bool   { return true }
func (MinTempExperienced) Writable() bool   { return false }
func (MinTempExperienced) Reportable() bool { return false }
func (MinTempExperienced) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinTempExperienced) AttrID() zcl.AttrID   { return v.ID() }
func (v MinTempExperienced) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinTempExperienced) AttrValue() zcl.Val  { return v.Value() }

func (MinTempExperienced) Name() string        { return `Min Temp Experienced` }
func (MinTempExperienced) Description() string { return `` }

type MinTempExperienced zcl.Zs16

func (v *MinTempExperienced) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MinTempExperienced) Value() zcl.Val     { return v }

func (v MinTempExperienced) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MinTempExperienced) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinTempExperienced(*nt)
	return br, err
}

func (v MinTempExperienced) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MinTempExperienced) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinTempExperienced(*a)
	return nil
}

func (v *MinTempExperienced) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MinTempExperienced(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinTempExperienced) String() string {
	return zcl.DegreesCelsius.Format(float64(v))
}

const ModelIdentifierAttr zcl.AttrID = 5

func (ModelIdentifier) ID() zcl.AttrID   { return ModelIdentifierAttr }
func (ModelIdentifier) Readable() bool   { return true }
func (ModelIdentifier) Writable() bool   { return false }
func (ModelIdentifier) Reportable() bool { return false }
func (ModelIdentifier) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ModelIdentifier) AttrID() zcl.AttrID   { return v.ID() }
func (v ModelIdentifier) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ModelIdentifier) AttrValue() zcl.Val  { return v.Value() }

func (ModelIdentifier) Name() string        { return `Model Identifier` }
func (ModelIdentifier) Description() string { return `` }

type ModelIdentifier zcl.Zcstring

func (v *ModelIdentifier) TypeID() zcl.TypeID { return new(zcl.Zcstring).TypeID() }
func (v *ModelIdentifier) Value() zcl.Val     { return v }

func (v ModelIdentifier) MarshalZcl() ([]byte, error) { return zcl.Zcstring(v).MarshalZcl() }

func (v *ModelIdentifier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*v = ModelIdentifier(*nt)
	return br, err
}

func (v ModelIdentifier) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zcstring(v))
}

func (v *ModelIdentifier) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zcstring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ModelIdentifier(*a)
	return nil
}

func (v *ModelIdentifier) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zcstring); ok {
		*v = ModelIdentifier(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ModelIdentifier) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v))
}

const MultistateNumberOfStatesAttr zcl.AttrID = 74

func (MultistateNumberOfStates) ID() zcl.AttrID   { return MultistateNumberOfStatesAttr }
func (MultistateNumberOfStates) Readable() bool   { return true }
func (MultistateNumberOfStates) Writable() bool   { return true }
func (MultistateNumberOfStates) Reportable() bool { return false }
func (MultistateNumberOfStates) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MultistateNumberOfStates) AttrID() zcl.AttrID   { return v.ID() }
func (v MultistateNumberOfStates) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MultistateNumberOfStates) AttrValue() zcl.Val  { return v.Value() }

func (MultistateNumberOfStates) Name() string        { return `Multistate Number of States` }
func (MultistateNumberOfStates) Description() string { return `` }

type MultistateNumberOfStates zcl.Zu16

func (v *MultistateNumberOfStates) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MultistateNumberOfStates) Value() zcl.Val     { return v }

func (v MultistateNumberOfStates) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MultistateNumberOfStates) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MultistateNumberOfStates(*nt)
	return br, err
}

func (v MultistateNumberOfStates) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MultistateNumberOfStates) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MultistateNumberOfStates(*a)
	return nil
}

func (v *MultistateNumberOfStates) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MultistateNumberOfStates(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MultistateNumberOfStates) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const MultistatePresentValueAttr zcl.AttrID = 85

func (MultistatePresentValue) ID() zcl.AttrID   { return MultistatePresentValueAttr }
func (MultistatePresentValue) Readable() bool   { return true }
func (MultistatePresentValue) Writable() bool   { return true }
func (MultistatePresentValue) Reportable() bool { return true }
func (MultistatePresentValue) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MultistatePresentValue) AttrID() zcl.AttrID   { return v.ID() }
func (v MultistatePresentValue) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MultistatePresentValue) AttrValue() zcl.Val  { return v.Value() }

func (MultistatePresentValue) Name() string        { return `Multistate Present value` }
func (MultistatePresentValue) Description() string { return `` }

type MultistatePresentValue zcl.Zu16

func (v *MultistatePresentValue) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MultistatePresentValue) Value() zcl.Val     { return v }

func (v MultistatePresentValue) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MultistatePresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MultistatePresentValue(*nt)
	return br, err
}

func (v MultistatePresentValue) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MultistatePresentValue) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MultistatePresentValue(*a)
	return nil
}

func (v *MultistatePresentValue) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MultistatePresentValue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MultistatePresentValue) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (MultistatePriority) Name() string        { return `Multistate Priority` }
func (MultistatePriority) Description() string { return `` }

type MultistatePriority struct {
	Ispriority    zcl.Zbool
	Priorityvalue zcl.Zu16
}

func (v *MultistatePriority) TypeID() zcl.TypeID { return zcl.TypeID(76) } // struct
func (v *MultistatePriority) Value() zcl.Val     { return v }
func (v MultistatePriority) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	var err error
	_ = tmp2
	_ = err

	{
		if tmp, err = v.Ispriority.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Priorityvalue.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (v *MultistatePriority) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2
	_ = err

	if b, err = new(zcl.Zbool).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = new(zcl.Zu16).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v *MultistatePriority) SetValue(a zcl.Val) error {
	if nv, ok := a.(*MultistatePriority); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrNotImpl
}

func (v *MultistatePriority) String() string {
	return zcl.Sprintf(
		"MultistatePriority{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

const MultistatePriorityArrayAttr zcl.AttrID = 87

func (MultistatePriorityArray) ID() zcl.AttrID   { return MultistatePriorityArrayAttr }
func (MultistatePriorityArray) Readable() bool   { return true }
func (MultistatePriorityArray) Writable() bool   { return true }
func (MultistatePriorityArray) Reportable() bool { return false }
func (MultistatePriorityArray) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MultistatePriorityArray) AttrID() zcl.AttrID   { return v.ID() }
func (v MultistatePriorityArray) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MultistatePriorityArray) AttrValue() zcl.Val  { return v.Value() }

func (MultistatePriorityArray) Name() string        { return `Multistate Priority Array` }
func (MultistatePriorityArray) Description() string { return `` }

type MultistatePriorityArray []*MultistatePriority

func (v *MultistatePriorityArray) TypeID() zcl.TypeID { return new(zcl.Zarray).TypeID() }
func (v *MultistatePriorityArray) Value() zcl.Val     { return v }

func (MultistatePriorityArray) ArrayTypeID() zcl.TypeID { return new(MultistatePriority).TypeID() }

func (v *MultistatePriorityArray) ArrayValues() (o []MultistatePriority) {
	for _, a := range *v {
		o = append(o, *a)
	}
	return o
}

func (v *MultistatePriorityArray) SetValues(val []MultistatePriority) error {
	*v = []*MultistatePriority{}
	return v.AddValues(val...)
}

func (v *MultistatePriorityArray) AddValues(val ...MultistatePriority) error {
	for _, a := range val {
		nv := a
		*v = append(*v, &nv)
	}
	return nil
}

func (v MultistatePriorityArray) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, a := range v {
		vars = append(vars, a)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (v *MultistatePriorityArray) UnmarshalZcl(b []byte) ([]byte, error) {
	*v = []*MultistatePriority{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(MultistatePriority)
		*v = append(*v, nv)
		return nv
	})
}

func (v *MultistatePriorityArray) SetValue(a zcl.Val) error {
	if nv, ok := a.(*MultistatePriorityArray); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MultistatePriorityArray) String() string {
	var s []string
	for _, a := range v {
		s = append(s, zcl.Sprintf("%v", a))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

const MultistateRelinquishDefaultAttr zcl.AttrID = 104

func (MultistateRelinquishDefault) ID() zcl.AttrID   { return MultistateRelinquishDefaultAttr }
func (MultistateRelinquishDefault) Readable() bool   { return true }
func (MultistateRelinquishDefault) Writable() bool   { return true }
func (MultistateRelinquishDefault) Reportable() bool { return false }
func (MultistateRelinquishDefault) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MultistateRelinquishDefault) AttrID() zcl.AttrID   { return v.ID() }
func (v MultistateRelinquishDefault) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MultistateRelinquishDefault) AttrValue() zcl.Val  { return v.Value() }

func (MultistateRelinquishDefault) Name() string        { return `Multistate Relinquish Default` }
func (MultistateRelinquishDefault) Description() string { return `` }

type MultistateRelinquishDefault zcl.Zu16

func (v *MultistateRelinquishDefault) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MultistateRelinquishDefault) Value() zcl.Val     { return v }

func (v MultistateRelinquishDefault) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MultistateRelinquishDefault) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MultistateRelinquishDefault(*nt)
	return br, err
}

func (v MultistateRelinquishDefault) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MultistateRelinquishDefault) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MultistateRelinquishDefault(*a)
	return nil
}

func (v *MultistateRelinquishDefault) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MultistateRelinquishDefault(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MultistateRelinquishDefault) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const MultistateTextAttr zcl.AttrID = 14

func (MultistateText) ID() zcl.AttrID   { return MultistateTextAttr }
func (MultistateText) Readable() bool   { return true }
func (MultistateText) Writable() bool   { return true }
func (MultistateText) Reportable() bool { return false }
func (MultistateText) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MultistateText) AttrID() zcl.AttrID   { return v.ID() }
func (v MultistateText) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MultistateText) AttrValue() zcl.Val  { return v.Value() }

func (MultistateText) Name() string        { return `Multistate Text` }
func (MultistateText) Description() string { return `` }

type MultistateText zcl.Zcstring

func (v *MultistateText) TypeID() zcl.TypeID { return new(zcl.Zcstring).TypeID() }
func (v *MultistateText) Value() zcl.Val     { return v }

func (v MultistateText) MarshalZcl() ([]byte, error) { return zcl.Zcstring(v).MarshalZcl() }

func (v *MultistateText) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*v = MultistateText(*nt)
	return br, err
}

func (v MultistateText) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zcstring(v))
}

func (v *MultistateText) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zcstring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MultistateText(*a)
	return nil
}

func (v *MultistateText) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zcstring); ok {
		*v = MultistateText(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MultistateText) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v))
}

const NwkDecryptFailuresAttr zcl.AttrID = 277

func (NwkDecryptFailures) ID() zcl.AttrID   { return NwkDecryptFailuresAttr }
func (NwkDecryptFailures) Readable() bool   { return true }
func (NwkDecryptFailures) Writable() bool   { return false }
func (NwkDecryptFailures) Reportable() bool { return false }
func (NwkDecryptFailures) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v NwkDecryptFailures) AttrID() zcl.AttrID   { return v.ID() }
func (v NwkDecryptFailures) AttrType() zcl.TypeID { return v.TypeID() }
func (v *NwkDecryptFailures) AttrValue() zcl.Val  { return v.Value() }

func (NwkDecryptFailures) Name() string        { return `NWK Decrypt Failures` }
func (NwkDecryptFailures) Description() string { return `` }

type NwkDecryptFailures zcl.Zu16

func (v *NwkDecryptFailures) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *NwkDecryptFailures) Value() zcl.Val     { return v }

func (v NwkDecryptFailures) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *NwkDecryptFailures) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = NwkDecryptFailures(*nt)
	return br, err
}

func (v NwkDecryptFailures) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *NwkDecryptFailures) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NwkDecryptFailures(*a)
	return nil
}

func (v *NwkDecryptFailures) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = NwkDecryptFailures(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NwkDecryptFailures) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const NwkFcFailureAttr zcl.AttrID = 274

func (NwkFcFailure) ID() zcl.AttrID   { return NwkFcFailureAttr }
func (NwkFcFailure) Readable() bool   { return true }
func (NwkFcFailure) Writable() bool   { return false }
func (NwkFcFailure) Reportable() bool { return false }
func (NwkFcFailure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v NwkFcFailure) AttrID() zcl.AttrID   { return v.ID() }
func (v NwkFcFailure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *NwkFcFailure) AttrValue() zcl.Val  { return v.Value() }

func (NwkFcFailure) Name() string        { return `NWK FC Failure` }
func (NwkFcFailure) Description() string { return `` }

type NwkFcFailure zcl.Zu16

func (v *NwkFcFailure) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *NwkFcFailure) Value() zcl.Val     { return v }

func (v NwkFcFailure) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *NwkFcFailure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = NwkFcFailure(*nt)
	return br, err
}

func (v NwkFcFailure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *NwkFcFailure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NwkFcFailure(*a)
	return nil
}

func (v *NwkFcFailure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = NwkFcFailure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NwkFcFailure) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const NeighborAddedAttr zcl.AttrID = 269

func (NeighborAdded) ID() zcl.AttrID   { return NeighborAddedAttr }
func (NeighborAdded) Readable() bool   { return true }
func (NeighborAdded) Writable() bool   { return false }
func (NeighborAdded) Reportable() bool { return false }
func (NeighborAdded) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v NeighborAdded) AttrID() zcl.AttrID   { return v.ID() }
func (v NeighborAdded) AttrType() zcl.TypeID { return v.TypeID() }
func (v *NeighborAdded) AttrValue() zcl.Val  { return v.Value() }

func (NeighborAdded) Name() string        { return `Neighbor Added` }
func (NeighborAdded) Description() string { return `` }

type NeighborAdded zcl.Zu16

func (v *NeighborAdded) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *NeighborAdded) Value() zcl.Val     { return v }

func (v NeighborAdded) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *NeighborAdded) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = NeighborAdded(*nt)
	return br, err
}

func (v NeighborAdded) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *NeighborAdded) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NeighborAdded(*a)
	return nil
}

func (v *NeighborAdded) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = NeighborAdded(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NeighborAdded) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (NeighborInfo) Name() string        { return `Neighbor Info` }
func (NeighborInfo) Description() string { return `` }

type NeighborInfo struct {
	Device      Device
	XCoordinate XCoordinate
	YCoordinate YCoordinate
	ZCoordinate ZCoordinate
	Rssi        Rssi
	// NumberRssiMeasurements is the number of measurements to use to generate one location estimate
	NumberRssiMeasurements NumberRssiMeasurements
}

func (v *NeighborInfo) TypeID() zcl.TypeID { return zcl.TypeID(76) } // struct
func (v *NeighborInfo) Value() zcl.Val     { return v }
func (v NeighborInfo) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	var err error
	_ = tmp2
	_ = err

	{
		if tmp, err = v.Device.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.XCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.YCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ZCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Rssi.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.NumberRssiMeasurements.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (v *NeighborInfo) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2
	_ = err

	if b, err = (&v.Device).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.XCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.YCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ZCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Rssi).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NumberRssiMeasurements).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v *NeighborInfo) SetValue(a zcl.Val) error {
	if nv, ok := a.(*NeighborInfo); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrNotImpl
}

func (v *NeighborInfo) String() string {
	return zcl.Sprintf(
		"NeighborInfo{"+zcl.StrJoin([]string{
			"Device(%v)",
			"XCoordinate(%v)",
			"YCoordinate(%v)",
			"ZCoordinate(%v)",
			"Rssi(%v)",
			"NumberRssiMeasurements(%v)",
		}, " ")+"}",
		v.Device,
		v.XCoordinate,
		v.YCoordinate,
		v.ZCoordinate,
		v.Rssi,
		v.NumberRssiMeasurements,
	)
}

const NeighborRemovedAttr zcl.AttrID = 270

func (NeighborRemoved) ID() zcl.AttrID   { return NeighborRemovedAttr }
func (NeighborRemoved) Readable() bool   { return true }
func (NeighborRemoved) Writable() bool   { return false }
func (NeighborRemoved) Reportable() bool { return false }
func (NeighborRemoved) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v NeighborRemoved) AttrID() zcl.AttrID   { return v.ID() }
func (v NeighborRemoved) AttrType() zcl.TypeID { return v.TypeID() }
func (v *NeighborRemoved) AttrValue() zcl.Val  { return v.Value() }

func (NeighborRemoved) Name() string        { return `Neighbor Removed` }
func (NeighborRemoved) Description() string { return `` }

type NeighborRemoved zcl.Zu16

func (v *NeighborRemoved) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *NeighborRemoved) Value() zcl.Val     { return v }

func (v NeighborRemoved) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *NeighborRemoved) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = NeighborRemoved(*nt)
	return br, err
}

func (v NeighborRemoved) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *NeighborRemoved) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NeighborRemoved(*a)
	return nil
}

func (v *NeighborRemoved) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = NeighborRemoved(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NeighborRemoved) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const NeighborStaleAttr zcl.AttrID = 271

func (NeighborStale) ID() zcl.AttrID   { return NeighborStaleAttr }
func (NeighborStale) Readable() bool   { return true }
func (NeighborStale) Writable() bool   { return false }
func (NeighborStale) Reportable() bool { return false }
func (NeighborStale) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v NeighborStale) AttrID() zcl.AttrID   { return v.ID() }
func (v NeighborStale) AttrType() zcl.TypeID { return v.TypeID() }
func (v *NeighborStale) AttrValue() zcl.Val  { return v.Value() }

func (NeighborStale) Name() string        { return `Neighbor Stale` }
func (NeighborStale) Description() string { return `` }

type NeighborStale zcl.Zu16

func (v *NeighborStale) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *NeighborStale) Value() zcl.Val     { return v }

func (v NeighborStale) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *NeighborStale) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = NeighborStale(*nt)
	return br, err
}

func (v NeighborStale) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *NeighborStale) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NeighborStale(*a)
	return nil
}

func (v *NeighborStale) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = NeighborStale(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NeighborStale) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (NeighborsInfoList) Name() string        { return `Neighbors Info List` }
func (NeighborsInfoList) Description() string { return `` }

type NeighborsInfoList []*NeighborInfo

func (v *NeighborsInfoList) TypeID() zcl.TypeID { return new(zcl.Zset).TypeID() }
func (v *NeighborsInfoList) Value() zcl.Val     { return v }

func (NeighborsInfoList) ArrayTypeID() zcl.TypeID { return new(NeighborInfo).TypeID() }

func (v *NeighborsInfoList) ArrayValues() (o []NeighborInfo) {
	for _, a := range *v {
		o = append(o, *a)
	}
	return o
}

func (v *NeighborsInfoList) SetValues(val []NeighborInfo) error {
	*v = []*NeighborInfo{}
	return v.AddValues(val...)
}

func (v *NeighborsInfoList) AddValues(val ...NeighborInfo) error {
	for _, a := range val {
		nv := a
		*v = append(*v, &nv)
	}
	return nil
}

func (v NeighborsInfoList) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, a := range v {
		vars = append(vars, a)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (v *NeighborsInfoList) UnmarshalZcl(b []byte) ([]byte, error) {
	*v = []*NeighborInfo{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(NeighborInfo)
		*v = append(*v, nv)
		return nv
	})
}

func (v *NeighborsInfoList) SetValue(a zcl.Val) error {
	if nv, ok := a.(*NeighborsInfoList); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NeighborsInfoList) String() string {
	var s []string
	for _, a := range v {
		s = append(s, zcl.Sprintf("%v", a))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

const NumberRssiMeasurementsAttr zcl.AttrID = 23

func (NumberRssiMeasurements) ID() zcl.AttrID   { return NumberRssiMeasurementsAttr }
func (NumberRssiMeasurements) Readable() bool   { return true }
func (NumberRssiMeasurements) Writable() bool   { return true }
func (NumberRssiMeasurements) Reportable() bool { return false }
func (NumberRssiMeasurements) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v NumberRssiMeasurements) AttrID() zcl.AttrID   { return v.ID() }
func (v NumberRssiMeasurements) AttrType() zcl.TypeID { return v.TypeID() }
func (v *NumberRssiMeasurements) AttrValue() zcl.Val  { return v.Value() }

func (NumberRssiMeasurements) Name() string { return `Number RSSI Measurements` }
func (NumberRssiMeasurements) Description() string {
	return `is the number of measurements to use to generate one location estimate`
}

// NumberRssiMeasurements is the number of measurements to use to generate one location estimate
type NumberRssiMeasurements zcl.Zu8

func (v *NumberRssiMeasurements) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *NumberRssiMeasurements) Value() zcl.Val     { return v }

func (v NumberRssiMeasurements) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *NumberRssiMeasurements) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = NumberRssiMeasurements(*nt)
	return br, err
}

func (v NumberRssiMeasurements) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *NumberRssiMeasurements) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NumberRssiMeasurements(*a)
	return nil
}

func (v *NumberRssiMeasurements) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = NumberRssiMeasurements(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NumberRssiMeasurements) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (NumberResponses) Name() string        { return `Number Responses` }
func (NumberResponses) Description() string { return `` }

type NumberResponses zcl.Zu8

func (v *NumberResponses) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *NumberResponses) Value() zcl.Val     { return v }

func (v NumberResponses) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *NumberResponses) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = NumberResponses(*nt)
	return br, err
}

func (v NumberResponses) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *NumberResponses) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NumberResponses(*a)
	return nil
}

func (v *NumberResponses) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = NumberResponses(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NumberResponses) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const NumberOfDevicesAttr zcl.AttrID = 4

func (NumberOfDevices) ID() zcl.AttrID   { return NumberOfDevicesAttr }
func (NumberOfDevices) Readable() bool   { return true }
func (NumberOfDevices) Writable() bool   { return false }
func (NumberOfDevices) Reportable() bool { return false }
func (NumberOfDevices) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v NumberOfDevices) AttrID() zcl.AttrID   { return v.ID() }
func (v NumberOfDevices) AttrType() zcl.TypeID { return v.TypeID() }
func (v *NumberOfDevices) AttrValue() zcl.Val  { return v.Value() }

func (NumberOfDevices) Name() string        { return `Number of Devices` }
func (NumberOfDevices) Description() string { return `` }

type NumberOfDevices zcl.Zu8

func (v *NumberOfDevices) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *NumberOfDevices) Value() zcl.Val     { return v }

func (v NumberOfDevices) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *NumberOfDevices) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = NumberOfDevices(*nt)
	return br, err
}

func (v NumberOfDevices) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *NumberOfDevices) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NumberOfDevices(*a)
	return nil
}

func (v *NumberOfDevices) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = NumberOfDevices(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NumberOfDevices) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const NumberOfResetsAttr zcl.AttrID = 0

func (NumberOfResets) ID() zcl.AttrID   { return NumberOfResetsAttr }
func (NumberOfResets) Readable() bool   { return true }
func (NumberOfResets) Writable() bool   { return false }
func (NumberOfResets) Reportable() bool { return false }
func (NumberOfResets) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v NumberOfResets) AttrID() zcl.AttrID   { return v.ID() }
func (v NumberOfResets) AttrType() zcl.TypeID { return v.TypeID() }
func (v *NumberOfResets) AttrValue() zcl.Val  { return v.Value() }

func (NumberOfResets) Name() string        { return `Number of Resets` }
func (NumberOfResets) Description() string { return `` }

type NumberOfResets zcl.Zu16

func (v *NumberOfResets) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *NumberOfResets) Value() zcl.Val     { return v }

func (v NumberOfResets) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *NumberOfResets) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = NumberOfResets(*nt)
	return br, err
}

func (v NumberOfResets) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *NumberOfResets) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NumberOfResets(*a)
	return nil
}

func (v *NumberOfResets) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = NumberOfResets(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NumberOfResets) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const OffTransitionTimeAttr zcl.AttrID = 19

func (OffTransitionTime) ID() zcl.AttrID   { return OffTransitionTimeAttr }
func (OffTransitionTime) Readable() bool   { return true }
func (OffTransitionTime) Writable() bool   { return true }
func (OffTransitionTime) Reportable() bool { return false }
func (OffTransitionTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v OffTransitionTime) AttrID() zcl.AttrID   { return v.ID() }
func (v OffTransitionTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *OffTransitionTime) AttrValue() zcl.Val  { return v.Value() }

func (OffTransitionTime) Name() string        { return `Off Transition Time` }
func (OffTransitionTime) Description() string { return `` }

type OffTransitionTime zcl.Zu16

func (v *OffTransitionTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *OffTransitionTime) Value() zcl.Val     { return v }

func (v OffTransitionTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *OffTransitionTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = OffTransitionTime(*nt)
	return br, err
}

func (v OffTransitionTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *OffTransitionTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OffTransitionTime(*a)
	return nil
}

func (v *OffTransitionTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = OffTransitionTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OffTransitionTime) String() string {
	return zcl.Seconds.Format(float64(v) / 10)
}

const OffWaitTimeAttr zcl.AttrID = 16386

func (OffWaitTime) ID() zcl.AttrID   { return OffWaitTimeAttr }
func (OffWaitTime) Readable() bool   { return true }
func (OffWaitTime) Writable() bool   { return false }
func (OffWaitTime) Reportable() bool { return false }
func (OffWaitTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v OffWaitTime) AttrID() zcl.AttrID   { return v.ID() }
func (v OffWaitTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *OffWaitTime) AttrValue() zcl.Val  { return v.Value() }

func (OffWaitTime) Name() string        { return `Off Wait Time` }
func (OffWaitTime) Description() string { return `` }

type OffWaitTime zcl.Zu16

func (v *OffWaitTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *OffWaitTime) Value() zcl.Val     { return v }

func (v OffWaitTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *OffWaitTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = OffWaitTime(*nt)
	return br, err
}

func (v OffWaitTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *OffWaitTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OffWaitTime(*a)
	return nil
}

func (v *OffWaitTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = OffWaitTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OffWaitTime) String() string {
	return zcl.Seconds.Format(float64(v) / 10)
}

const OnLevelAttr zcl.AttrID = 17

func (OnLevel) ID() zcl.AttrID   { return OnLevelAttr }
func (OnLevel) Readable() bool   { return true }
func (OnLevel) Writable() bool   { return true }
func (OnLevel) Reportable() bool { return false }
func (OnLevel) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v OnLevel) AttrID() zcl.AttrID   { return v.ID() }
func (v OnLevel) AttrType() zcl.TypeID { return v.TypeID() }
func (v *OnLevel) AttrValue() zcl.Val  { return v.Value() }

func (OnLevel) Name() string { return `On Level` }
func (OnLevel) Description() string {
	return `determines the value that the CurrentLevel attribute is set to when the OnOff attribute of an On/Off cluster
on the same endpoint is set to On. If the OnLevel attribute is not implemented, or is set to 0xff, it has no
effect.`
}

// OnLevel determines the value that the CurrentLevel attribute is set to when the OnOff attribute of an On/Off cluster
// on the same endpoint is set to On. If the OnLevel attribute is not implemented, or is set to 0xff, it has no
// effect.
type OnLevel zcl.Zu8

func (v *OnLevel) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *OnLevel) Value() zcl.Val     { return v }

func (v OnLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *OnLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = OnLevel(*nt)
	return br, err
}

func (v OnLevel) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *OnLevel) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OnLevel(*a)
	return nil
}

func (v *OnLevel) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = OnLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OnLevel) String() string {
	switch v {
	case 0xFF:
		return "Previous"
	}
	return zcl.Percent.Format(float64(v) / 2.54)
}

func (v OnLevel) IsPrevious() bool { return v == 0xFF }
func (v *OnLevel) SetPrevious()    { *v = 0xFF }

func (OnLevel) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0xFF, Name: "Previous"},
	}
}

const OnOffAttr zcl.AttrID = 0

func (OnOff) ID() zcl.AttrID   { return OnOffAttr }
func (OnOff) Readable() bool   { return true }
func (OnOff) Writable() bool   { return false }
func (OnOff) Reportable() bool { return true }
func (OnOff) SceneIndex() int  { return 1 }

// Implements AttrDef/AttrValue interfaces
func (v OnOff) AttrID() zcl.AttrID   { return v.ID() }
func (v OnOff) AttrType() zcl.TypeID { return v.TypeID() }
func (v *OnOff) AttrValue() zcl.Val  { return v.Value() }

func (OnOff) Name() string        { return `On Off` }
func (OnOff) Description() string { return `` }

type OnOff zcl.Zbool

func (v *OnOff) TypeID() zcl.TypeID { return new(zcl.Zbool).TypeID() }
func (v *OnOff) Value() zcl.Val     { return v }

func (v OnOff) MarshalZcl() ([]byte, error) { return zcl.Zbool(v).MarshalZcl() }

func (v *OnOff) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*v = OnOff(*nt)
	return br, err
}

func (v OnOff) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbool(v))
}

func (v *OnOff) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbool)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OnOff(*a)
	return nil
}

func (v *OnOff) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbool); ok {
		*v = OnOff(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OnOff) String() string {
	switch v {
	case 0x00:
		return "Off"
	case 0x01:
		return "On"
	}
	return zcl.Sprintf("%v", zcl.Zbool(v))
}

func (v OnOff) IsOff() bool { return v == 0x00 }
func (v OnOff) IsOn() bool  { return v == 0x01 }
func (v *OnOff) SetOff()    { *v = 0x00 }
func (v *OnOff) SetOn()     { *v = 0x01 }

func (OnOff) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Off"},
		{Value: 0x01, Name: "On"},
	}
}

const OnTimeAttr zcl.AttrID = 16385

func (OnTime) ID() zcl.AttrID   { return OnTimeAttr }
func (OnTime) Readable() bool   { return true }
func (OnTime) Writable() bool   { return false }
func (OnTime) Reportable() bool { return false }
func (OnTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v OnTime) AttrID() zcl.AttrID   { return v.ID() }
func (v OnTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *OnTime) AttrValue() zcl.Val  { return v.Value() }

func (OnTime) Name() string        { return `On Time` }
func (OnTime) Description() string { return `` }

type OnTime zcl.Zu16

func (v *OnTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *OnTime) Value() zcl.Val     { return v }

func (v OnTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *OnTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = OnTime(*nt)
	return br, err
}

func (v OnTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *OnTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OnTime(*a)
	return nil
}

func (v *OnTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = OnTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OnTime) String() string {
	return zcl.Seconds.Format(float64(v) / 10)
}

const OnTransitionTimeAttr zcl.AttrID = 18

func (OnTransitionTime) ID() zcl.AttrID   { return OnTransitionTimeAttr }
func (OnTransitionTime) Readable() bool   { return true }
func (OnTransitionTime) Writable() bool   { return true }
func (OnTransitionTime) Reportable() bool { return false }
func (OnTransitionTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v OnTransitionTime) AttrID() zcl.AttrID   { return v.ID() }
func (v OnTransitionTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *OnTransitionTime) AttrValue() zcl.Val  { return v.Value() }

func (OnTransitionTime) Name() string        { return `On Transition Time` }
func (OnTransitionTime) Description() string { return `` }

type OnTransitionTime zcl.Zu16

func (v *OnTransitionTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *OnTransitionTime) Value() zcl.Val     { return v }

func (v OnTransitionTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *OnTransitionTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = OnTransitionTime(*nt)
	return br, err
}

func (v OnTransitionTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *OnTransitionTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OnTransitionTime(*a)
	return nil
}

func (v *OnTransitionTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = OnTransitionTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OnTransitionTime) String() string {
	return zcl.Seconds.Format(float64(v) / 10)
}

const OnOffTransistionTimeAttr zcl.AttrID = 16

func (OnOffTransistionTime) ID() zcl.AttrID   { return OnOffTransistionTimeAttr }
func (OnOffTransistionTime) Readable() bool   { return true }
func (OnOffTransistionTime) Writable() bool   { return true }
func (OnOffTransistionTime) Reportable() bool { return false }
func (OnOffTransistionTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v OnOffTransistionTime) AttrID() zcl.AttrID   { return v.ID() }
func (v OnOffTransistionTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *OnOffTransistionTime) AttrValue() zcl.Val  { return v.Value() }

func (OnOffTransistionTime) Name() string { return `On/Off Transistion Time` }
func (OnOffTransistionTime) Description() string {
	return `represents the time taken to move to or from the target level when On of Off commands are received
by an On/Off cluster on the same endpoint.
The actual time taken should be as close to OnOffTransitionTime as the device is able.`
}

// OnOffTransistionTime represents the time taken to move to or from the target level when On of Off commands are received
// by an On/Off cluster on the same endpoint.
// The actual time taken should be as close to OnOffTransitionTime as the device is able.
type OnOffTransistionTime zcl.Zu16

func (v *OnOffTransistionTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *OnOffTransistionTime) Value() zcl.Val     { return v }

func (v OnOffTransistionTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *OnOffTransistionTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = OnOffTransistionTime(*nt)
	return br, err
}

func (v OnOffTransistionTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *OnOffTransistionTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OnOffTransistionTime(*a)
	return nil
}

func (v *OnOffTransistionTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = OnOffTransistionTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OnOffTransistionTime) String() string {
	return zcl.Seconds.Format(float64(v) / 10)
}

func (OnOffControl) Name() string        { return `On/off control` }
func (OnOffControl) Description() string { return `` }

type OnOffControl zcl.Zbmp8

func (v *OnOffControl) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *OnOffControl) Value() zcl.Val     { return v }

func (v OnOffControl) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *OnOffControl) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = OnOffControl(*nt)
	return br, err
}

func (v OnOffControl) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *OnOffControl) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OnOffControl(*a)
	return nil
}

func (v *OnOffControl) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = OnOffControl(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OnOffControl) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v OnOffControl) IsAcceptOnlyWhenOn() bool { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v *OnOffControl) SetAcceptOnlyWhenOn(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}

func (OnOffControl) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Accept only when on"},
	}
}

const OverTempTotalDwellAttr zcl.AttrID = 3

func (OverTempTotalDwell) ID() zcl.AttrID   { return OverTempTotalDwellAttr }
func (OverTempTotalDwell) Readable() bool   { return false }
func (OverTempTotalDwell) Writable() bool   { return false }
func (OverTempTotalDwell) Reportable() bool { return false }
func (OverTempTotalDwell) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v OverTempTotalDwell) AttrID() zcl.AttrID   { return v.ID() }
func (v OverTempTotalDwell) AttrType() zcl.TypeID { return v.TypeID() }
func (v *OverTempTotalDwell) AttrValue() zcl.Val  { return v.Value() }

func (OverTempTotalDwell) Name() string { return `Over Temp Total Dwell` }
func (OverTempTotalDwell) Description() string {
	return `Total time the device has spent above the tmperature specified by High Temp Threshold`
}

// OverTempTotalDwell Total time the device has spent above the tmperature specified by High Temp Threshold
type OverTempTotalDwell zcl.Zu16

func (v *OverTempTotalDwell) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *OverTempTotalDwell) Value() zcl.Val     { return v }

func (v OverTempTotalDwell) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *OverTempTotalDwell) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = OverTempTotalDwell(*nt)
	return br, err
}

func (v OverTempTotalDwell) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *OverTempTotalDwell) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OverTempTotalDwell(*a)
	return nil
}

func (v *OverTempTotalDwell) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = OverTempTotalDwell(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OverTempTotalDwell) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const PacketBufferAllocFailuresAttr zcl.AttrID = 279

func (PacketBufferAllocFailures) ID() zcl.AttrID   { return PacketBufferAllocFailuresAttr }
func (PacketBufferAllocFailures) Readable() bool   { return true }
func (PacketBufferAllocFailures) Writable() bool   { return false }
func (PacketBufferAllocFailures) Reportable() bool { return false }
func (PacketBufferAllocFailures) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PacketBufferAllocFailures) AttrID() zcl.AttrID   { return v.ID() }
func (v PacketBufferAllocFailures) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PacketBufferAllocFailures) AttrValue() zcl.Val  { return v.Value() }

func (PacketBufferAllocFailures) Name() string        { return `Packet Buffer Alloc Failures` }
func (PacketBufferAllocFailures) Description() string { return `` }

type PacketBufferAllocFailures zcl.Zu16

func (v *PacketBufferAllocFailures) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *PacketBufferAllocFailures) Value() zcl.Val     { return v }

func (v PacketBufferAllocFailures) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *PacketBufferAllocFailures) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = PacketBufferAllocFailures(*nt)
	return br, err
}

func (v PacketBufferAllocFailures) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *PacketBufferAllocFailures) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PacketBufferAllocFailures(*a)
	return nil
}

func (v *PacketBufferAllocFailures) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = PacketBufferAllocFailures(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PacketBufferAllocFailures) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const PacketValidateDropcountAttr zcl.AttrID = 282

func (PacketValidateDropcount) ID() zcl.AttrID   { return PacketValidateDropcountAttr }
func (PacketValidateDropcount) Readable() bool   { return true }
func (PacketValidateDropcount) Writable() bool   { return false }
func (PacketValidateDropcount) Reportable() bool { return false }
func (PacketValidateDropcount) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PacketValidateDropcount) AttrID() zcl.AttrID   { return v.ID() }
func (v PacketValidateDropcount) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PacketValidateDropcount) AttrValue() zcl.Val  { return v.Value() }

func (PacketValidateDropcount) Name() string        { return `Packet Validate Dropcount` }
func (PacketValidateDropcount) Description() string { return `` }

type PacketValidateDropcount zcl.Zu16

func (v *PacketValidateDropcount) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *PacketValidateDropcount) Value() zcl.Val     { return v }

func (v PacketValidateDropcount) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *PacketValidateDropcount) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = PacketValidateDropcount(*nt)
	return br, err
}

func (v PacketValidateDropcount) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *PacketValidateDropcount) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PacketValidateDropcount(*a)
	return nil
}

func (v *PacketValidateDropcount) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = PacketValidateDropcount(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PacketValidateDropcount) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const PathLossExponentAttr zcl.AttrID = 20

func (PathLossExponent) ID() zcl.AttrID   { return PathLossExponentAttr }
func (PathLossExponent) Readable() bool   { return true }
func (PathLossExponent) Writable() bool   { return true }
func (PathLossExponent) Reportable() bool { return false }
func (PathLossExponent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PathLossExponent) AttrID() zcl.AttrID   { return v.ID() }
func (v PathLossExponent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PathLossExponent) AttrValue() zcl.Val  { return v.Value() }

func (PathLossExponent) Name() string { return `Path loss Exponent` }
func (PathLossExponent) Description() string {
	return `is the rate at which the signal power decays with increasing distance`
}

// PathLossExponent is the rate at which the signal power decays with increasing distance
type PathLossExponent zcl.Zu16

func (v *PathLossExponent) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *PathLossExponent) Value() zcl.Val     { return v }

func (v PathLossExponent) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *PathLossExponent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = PathLossExponent(*nt)
	return br, err
}

func (v PathLossExponent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *PathLossExponent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PathLossExponent(*a)
	return nil
}

func (v *PathLossExponent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = PathLossExponent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PathLossExponent) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const PersistensMemoryWritesAttr zcl.AttrID = 1

func (PersistensMemoryWrites) ID() zcl.AttrID   { return PersistensMemoryWritesAttr }
func (PersistensMemoryWrites) Readable() bool   { return true }
func (PersistensMemoryWrites) Writable() bool   { return false }
func (PersistensMemoryWrites) Reportable() bool { return false }
func (PersistensMemoryWrites) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PersistensMemoryWrites) AttrID() zcl.AttrID   { return v.ID() }
func (v PersistensMemoryWrites) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PersistensMemoryWrites) AttrValue() zcl.Val  { return v.Value() }

func (PersistensMemoryWrites) Name() string        { return `Persistens Memory Writes` }
func (PersistensMemoryWrites) Description() string { return `` }

type PersistensMemoryWrites zcl.Zu16

func (v *PersistensMemoryWrites) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *PersistensMemoryWrites) Value() zcl.Val     { return v }

func (v PersistensMemoryWrites) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *PersistensMemoryWrites) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = PersistensMemoryWrites(*nt)
	return br, err
}

func (v PersistensMemoryWrites) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *PersistensMemoryWrites) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PersistensMemoryWrites(*a)
	return nil
}

func (v *PersistensMemoryWrites) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = PersistensMemoryWrites(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PersistensMemoryWrites) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const PhyToMacQueueLimitReachedAttr zcl.AttrID = 281

func (PhyToMacQueueLimitReached) ID() zcl.AttrID   { return PhyToMacQueueLimitReachedAttr }
func (PhyToMacQueueLimitReached) Readable() bool   { return true }
func (PhyToMacQueueLimitReached) Writable() bool   { return false }
func (PhyToMacQueueLimitReached) Reportable() bool { return false }
func (PhyToMacQueueLimitReached) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PhyToMacQueueLimitReached) AttrID() zcl.AttrID   { return v.ID() }
func (v PhyToMacQueueLimitReached) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PhyToMacQueueLimitReached) AttrValue() zcl.Val  { return v.Value() }

func (PhyToMacQueueLimitReached) Name() string        { return `Phy to MAC queue limit reached` }
func (PhyToMacQueueLimitReached) Description() string { return `` }

type PhyToMacQueueLimitReached zcl.Zu16

func (v *PhyToMacQueueLimitReached) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *PhyToMacQueueLimitReached) Value() zcl.Val     { return v }

func (v PhyToMacQueueLimitReached) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *PhyToMacQueueLimitReached) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = PhyToMacQueueLimitReached(*nt)
	return br, err
}

func (v PhyToMacQueueLimitReached) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *PhyToMacQueueLimitReached) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PhyToMacQueueLimitReached(*a)
	return nil
}

func (v *PhyToMacQueueLimitReached) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = PhyToMacQueueLimitReached(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PhyToMacQueueLimitReached) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const PhysicalEnvironmentAttr zcl.AttrID = 17

func (PhysicalEnvironment) ID() zcl.AttrID   { return PhysicalEnvironmentAttr }
func (PhysicalEnvironment) Readable() bool   { return true }
func (PhysicalEnvironment) Writable() bool   { return true }
func (PhysicalEnvironment) Reportable() bool { return false }
func (PhysicalEnvironment) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PhysicalEnvironment) AttrID() zcl.AttrID   { return v.ID() }
func (v PhysicalEnvironment) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PhysicalEnvironment) AttrValue() zcl.Val  { return v.Value() }

func (PhysicalEnvironment) Name() string        { return `Physical Environment` }
func (PhysicalEnvironment) Description() string { return `` }

type PhysicalEnvironment zcl.Zenum8

func (v *PhysicalEnvironment) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *PhysicalEnvironment) Value() zcl.Val     { return v }

func (v PhysicalEnvironment) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *PhysicalEnvironment) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = PhysicalEnvironment(*nt)
	return br, err
}

func (v PhysicalEnvironment) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *PhysicalEnvironment) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PhysicalEnvironment(*a)
	return nil
}

func (v *PhysicalEnvironment) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = PhysicalEnvironment(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PhysicalEnvironment) String() string {
	switch v {
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
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v PhysicalEnvironment) IsUnspecifiedEnvironment() bool           { return v == 0x00 }
func (v PhysicalEnvironment) IsAtrium() bool                           { return v == 0x01 }
func (v PhysicalEnvironment) IsBar() bool                              { return v == 0x02 }
func (v PhysicalEnvironment) IsCourtyard() bool                        { return v == 0x03 }
func (v PhysicalEnvironment) IsBathroom() bool                         { return v == 0x04 }
func (v PhysicalEnvironment) IsBedroom() bool                          { return v == 0x05 }
func (v PhysicalEnvironment) IsBilliardRoom() bool                     { return v == 0x06 }
func (v PhysicalEnvironment) IsUtilityRoom() bool                      { return v == 0x07 }
func (v PhysicalEnvironment) IsCellar() bool                           { return v == 0x08 }
func (v PhysicalEnvironment) IsStorageCloset() bool                    { return v == 0x09 }
func (v PhysicalEnvironment) IsTheater() bool                          { return v == 0x0a }
func (v PhysicalEnvironment) IsOffice() bool                           { return v == 0x0b }
func (v PhysicalEnvironment) IsDeck() bool                             { return v == 0x0c }
func (v PhysicalEnvironment) IsDen() bool                              { return v == 0x0d }
func (v PhysicalEnvironment) IsDiningRoom() bool                       { return v == 0x0e }
func (v PhysicalEnvironment) IsElectricalRoom() bool                   { return v == 0x0f }
func (v PhysicalEnvironment) IsElevator() bool                         { return v == 0x10 }
func (v PhysicalEnvironment) IsEntry() bool                            { return v == 0x11 }
func (v PhysicalEnvironment) IsFamilyRoom() bool                       { return v == 0x12 }
func (v PhysicalEnvironment) IsMainFloor() bool                        { return v == 0x13 }
func (v PhysicalEnvironment) IsUpstairs() bool                         { return v == 0x14 }
func (v PhysicalEnvironment) IsDownstairs() bool                       { return v == 0x15 }
func (v PhysicalEnvironment) IsBasementLowerLevel() bool               { return v == 0x16 }
func (v PhysicalEnvironment) IsGallery() bool                          { return v == 0x17 }
func (v PhysicalEnvironment) IsGameRoom() bool                         { return v == 0x18 }
func (v PhysicalEnvironment) IsGarage() bool                           { return v == 0x19 }
func (v PhysicalEnvironment) IsGym() bool                              { return v == 0x1a }
func (v PhysicalEnvironment) IsHallway() bool                          { return v == 0x1b }
func (v PhysicalEnvironment) IsHouse() bool                            { return v == 0x1c }
func (v PhysicalEnvironment) IsKitchen() bool                          { return v == 0x1d }
func (v PhysicalEnvironment) IsLaundryRoom() bool                      { return v == 0x1e }
func (v PhysicalEnvironment) IsLibrary() bool                          { return v == 0x1f }
func (v PhysicalEnvironment) IsMasterBedroom() bool                    { return v == 0x20 }
func (v PhysicalEnvironment) IsMudRoomSmallRoomForCoatsAndBoots() bool { return v == 0x21 }
func (v PhysicalEnvironment) IsNursery() bool                          { return v == 0x22 }
func (v PhysicalEnvironment) IsPantry() bool                           { return v == 0x23 }
func (v PhysicalEnvironment) IsOffice2() bool                          { return v == 0x24 }
func (v PhysicalEnvironment) IsOutside() bool                          { return v == 0x25 }
func (v PhysicalEnvironment) IsPool() bool                             { return v == 0x26 }
func (v PhysicalEnvironment) IsPorch() bool                            { return v == 0x27 }
func (v PhysicalEnvironment) IsSewingRoom() bool                       { return v == 0x28 }
func (v PhysicalEnvironment) IsSittingRoom() bool                      { return v == 0x29 }
func (v PhysicalEnvironment) IsStairway() bool                         { return v == 0x2a }
func (v PhysicalEnvironment) IsYard() bool                             { return v == 0x2b }
func (v PhysicalEnvironment) IsAttic() bool                            { return v == 0x2c }
func (v PhysicalEnvironment) IsHotTub() bool                           { return v == 0x2d }
func (v PhysicalEnvironment) IsLivingRoom() bool                       { return v == 0x2e }
func (v PhysicalEnvironment) IsSauna() bool                            { return v == 0x2f }
func (v PhysicalEnvironment) IsShopWorkshop() bool                     { return v == 0x30 }
func (v PhysicalEnvironment) IsGuestBedroom() bool                     { return v == 0x31 }
func (v PhysicalEnvironment) IsGuestBath() bool                        { return v == 0x32 }
func (v PhysicalEnvironment) IsPowderRoom12Bath() bool                 { return v == 0x33 }
func (v PhysicalEnvironment) IsBackYard() bool                         { return v == 0x34 }
func (v PhysicalEnvironment) IsFrontYard() bool                        { return v == 0x35 }
func (v PhysicalEnvironment) IsPatio() bool                            { return v == 0x36 }
func (v PhysicalEnvironment) IsDriveway() bool                         { return v == 0x37 }
func (v PhysicalEnvironment) IsSunRoom() bool                          { return v == 0x38 }
func (v PhysicalEnvironment) IsLivingRoom2() bool                      { return v == 0x39 }
func (v PhysicalEnvironment) IsSpa() bool                              { return v == 0x3a }
func (v PhysicalEnvironment) IsWhirlpool() bool                        { return v == 0x3b }
func (v PhysicalEnvironment) IsShed() bool                             { return v == 0x3c }
func (v PhysicalEnvironment) IsEquipmentStorage() bool                 { return v == 0x3d }
func (v PhysicalEnvironment) IsHobbyCraftRoom() bool                   { return v == 0x3e }
func (v PhysicalEnvironment) IsFountain() bool                         { return v == 0x3f }
func (v PhysicalEnvironment) IsPond() bool                             { return v == 0x40 }
func (v PhysicalEnvironment) IsReceptionRoom() bool                    { return v == 0x41 }
func (v PhysicalEnvironment) IsBreakfastRoom() bool                    { return v == 0x42 }
func (v PhysicalEnvironment) IsNook() bool                             { return v == 0x43 }
func (v PhysicalEnvironment) IsGarden() bool                           { return v == 0x44 }
func (v PhysicalEnvironment) IsBalcony() bool                          { return v == 0x45 }
func (v PhysicalEnvironment) IsPanicRoom() bool                        { return v == 0x46 }
func (v PhysicalEnvironment) IsTerrace() bool                          { return v == 0x47 }
func (v PhysicalEnvironment) IsRoof() bool                             { return v == 0x48 }
func (v PhysicalEnvironment) IsToilet() bool                           { return v == 0x49 }
func (v PhysicalEnvironment) IsToiletMain() bool                       { return v == 0x4a }
func (v PhysicalEnvironment) IsOutsideToilet() bool                    { return v == 0x4b }
func (v PhysicalEnvironment) IsShowerRoom() bool                       { return v == 0x4c }
func (v PhysicalEnvironment) IsStudy() bool                            { return v == 0x4d }
func (v PhysicalEnvironment) IsFrontGarden() bool                      { return v == 0x4e }
func (v PhysicalEnvironment) IsBackGarden() bool                       { return v == 0x4f }
func (v PhysicalEnvironment) IsKettle() bool                           { return v == 0x50 }
func (v PhysicalEnvironment) IsTelevision() bool                       { return v == 0x51 }
func (v PhysicalEnvironment) IsStove() bool                            { return v == 0x52 }
func (v PhysicalEnvironment) IsMicrowave() bool                        { return v == 0x53 }
func (v PhysicalEnvironment) IsToaster() bool                          { return v == 0x54 }
func (v PhysicalEnvironment) IsVacuum() bool                           { return v == 0x55 }
func (v PhysicalEnvironment) IsAppliance() bool                        { return v == 0x56 }
func (v PhysicalEnvironment) IsFrontDoor() bool                        { return v == 0x57 }
func (v PhysicalEnvironment) IsBackDoor() bool                         { return v == 0x58 }
func (v PhysicalEnvironment) IsFridgeDoor() bool                       { return v == 0x59 }
func (v PhysicalEnvironment) IsMedicationCabinetDoor() bool            { return v == 0x60 }
func (v PhysicalEnvironment) IsWardrobeDoor() bool                     { return v == 0x61 }
func (v PhysicalEnvironment) IsFrontCupboardDoor() bool                { return v == 0x62 }
func (v PhysicalEnvironment) IsOtherDoor() bool                        { return v == 0x63 }
func (v PhysicalEnvironment) IsWaitingRoom() bool                      { return v == 0x64 }
func (v PhysicalEnvironment) IsTriageRoom() bool                       { return v == 0x65 }
func (v PhysicalEnvironment) IsDoctorSOffice() bool                    { return v == 0x66 }
func (v PhysicalEnvironment) IsPatientSPrivateRoom() bool              { return v == 0x67 }
func (v PhysicalEnvironment) IsConsultationRoom() bool                 { return v == 0x68 }
func (v PhysicalEnvironment) IsNurseStation() bool                     { return v == 0x69 }
func (v PhysicalEnvironment) IsWard() bool                             { return v == 0x6a }
func (v PhysicalEnvironment) IsCorridor() bool                         { return v == 0x6b }
func (v PhysicalEnvironment) IsOperatingTheatre() bool                 { return v == 0x6c }
func (v PhysicalEnvironment) IsDentalSurgeryRoom() bool                { return v == 0x6d }
func (v PhysicalEnvironment) IsMedicalImagingRoom() bool               { return v == 0x6e }
func (v PhysicalEnvironment) IsDecontaminationRoom() bool              { return v == 0x6f }
func (v PhysicalEnvironment) IsUnknownEnvironment() bool               { return v == 0xFF }
func (v *PhysicalEnvironment) SetUnspecifiedEnvironment()              { *v = 0x00 }
func (v *PhysicalEnvironment) SetAtrium()                              { *v = 0x01 }
func (v *PhysicalEnvironment) SetBar()                                 { *v = 0x02 }
func (v *PhysicalEnvironment) SetCourtyard()                           { *v = 0x03 }
func (v *PhysicalEnvironment) SetBathroom()                            { *v = 0x04 }
func (v *PhysicalEnvironment) SetBedroom()                             { *v = 0x05 }
func (v *PhysicalEnvironment) SetBilliardRoom()                        { *v = 0x06 }
func (v *PhysicalEnvironment) SetUtilityRoom()                         { *v = 0x07 }
func (v *PhysicalEnvironment) SetCellar()                              { *v = 0x08 }
func (v *PhysicalEnvironment) SetStorageCloset()                       { *v = 0x09 }
func (v *PhysicalEnvironment) SetTheater()                             { *v = 0x0a }
func (v *PhysicalEnvironment) SetOffice()                              { *v = 0x0b }
func (v *PhysicalEnvironment) SetDeck()                                { *v = 0x0c }
func (v *PhysicalEnvironment) SetDen()                                 { *v = 0x0d }
func (v *PhysicalEnvironment) SetDiningRoom()                          { *v = 0x0e }
func (v *PhysicalEnvironment) SetElectricalRoom()                      { *v = 0x0f }
func (v *PhysicalEnvironment) SetElevator()                            { *v = 0x10 }
func (v *PhysicalEnvironment) SetEntry()                               { *v = 0x11 }
func (v *PhysicalEnvironment) SetFamilyRoom()                          { *v = 0x12 }
func (v *PhysicalEnvironment) SetMainFloor()                           { *v = 0x13 }
func (v *PhysicalEnvironment) SetUpstairs()                            { *v = 0x14 }
func (v *PhysicalEnvironment) SetDownstairs()                          { *v = 0x15 }
func (v *PhysicalEnvironment) SetBasementLowerLevel()                  { *v = 0x16 }
func (v *PhysicalEnvironment) SetGallery()                             { *v = 0x17 }
func (v *PhysicalEnvironment) SetGameRoom()                            { *v = 0x18 }
func (v *PhysicalEnvironment) SetGarage()                              { *v = 0x19 }
func (v *PhysicalEnvironment) SetGym()                                 { *v = 0x1a }
func (v *PhysicalEnvironment) SetHallway()                             { *v = 0x1b }
func (v *PhysicalEnvironment) SetHouse()                               { *v = 0x1c }
func (v *PhysicalEnvironment) SetKitchen()                             { *v = 0x1d }
func (v *PhysicalEnvironment) SetLaundryRoom()                         { *v = 0x1e }
func (v *PhysicalEnvironment) SetLibrary()                             { *v = 0x1f }
func (v *PhysicalEnvironment) SetMasterBedroom()                       { *v = 0x20 }
func (v *PhysicalEnvironment) SetMudRoomSmallRoomForCoatsAndBoots()    { *v = 0x21 }
func (v *PhysicalEnvironment) SetNursery()                             { *v = 0x22 }
func (v *PhysicalEnvironment) SetPantry()                              { *v = 0x23 }
func (v *PhysicalEnvironment) SetOffice2()                             { *v = 0x24 }
func (v *PhysicalEnvironment) SetOutside()                             { *v = 0x25 }
func (v *PhysicalEnvironment) SetPool()                                { *v = 0x26 }
func (v *PhysicalEnvironment) SetPorch()                               { *v = 0x27 }
func (v *PhysicalEnvironment) SetSewingRoom()                          { *v = 0x28 }
func (v *PhysicalEnvironment) SetSittingRoom()                         { *v = 0x29 }
func (v *PhysicalEnvironment) SetStairway()                            { *v = 0x2a }
func (v *PhysicalEnvironment) SetYard()                                { *v = 0x2b }
func (v *PhysicalEnvironment) SetAttic()                               { *v = 0x2c }
func (v *PhysicalEnvironment) SetHotTub()                              { *v = 0x2d }
func (v *PhysicalEnvironment) SetLivingRoom()                          { *v = 0x2e }
func (v *PhysicalEnvironment) SetSauna()                               { *v = 0x2f }
func (v *PhysicalEnvironment) SetShopWorkshop()                        { *v = 0x30 }
func (v *PhysicalEnvironment) SetGuestBedroom()                        { *v = 0x31 }
func (v *PhysicalEnvironment) SetGuestBath()                           { *v = 0x32 }
func (v *PhysicalEnvironment) SetPowderRoom12Bath()                    { *v = 0x33 }
func (v *PhysicalEnvironment) SetBackYard()                            { *v = 0x34 }
func (v *PhysicalEnvironment) SetFrontYard()                           { *v = 0x35 }
func (v *PhysicalEnvironment) SetPatio()                               { *v = 0x36 }
func (v *PhysicalEnvironment) SetDriveway()                            { *v = 0x37 }
func (v *PhysicalEnvironment) SetSunRoom()                             { *v = 0x38 }
func (v *PhysicalEnvironment) SetLivingRoom2()                         { *v = 0x39 }
func (v *PhysicalEnvironment) SetSpa()                                 { *v = 0x3a }
func (v *PhysicalEnvironment) SetWhirlpool()                           { *v = 0x3b }
func (v *PhysicalEnvironment) SetShed()                                { *v = 0x3c }
func (v *PhysicalEnvironment) SetEquipmentStorage()                    { *v = 0x3d }
func (v *PhysicalEnvironment) SetHobbyCraftRoom()                      { *v = 0x3e }
func (v *PhysicalEnvironment) SetFountain()                            { *v = 0x3f }
func (v *PhysicalEnvironment) SetPond()                                { *v = 0x40 }
func (v *PhysicalEnvironment) SetReceptionRoom()                       { *v = 0x41 }
func (v *PhysicalEnvironment) SetBreakfastRoom()                       { *v = 0x42 }
func (v *PhysicalEnvironment) SetNook()                                { *v = 0x43 }
func (v *PhysicalEnvironment) SetGarden()                              { *v = 0x44 }
func (v *PhysicalEnvironment) SetBalcony()                             { *v = 0x45 }
func (v *PhysicalEnvironment) SetPanicRoom()                           { *v = 0x46 }
func (v *PhysicalEnvironment) SetTerrace()                             { *v = 0x47 }
func (v *PhysicalEnvironment) SetRoof()                                { *v = 0x48 }
func (v *PhysicalEnvironment) SetToilet()                              { *v = 0x49 }
func (v *PhysicalEnvironment) SetToiletMain()                          { *v = 0x4a }
func (v *PhysicalEnvironment) SetOutsideToilet()                       { *v = 0x4b }
func (v *PhysicalEnvironment) SetShowerRoom()                          { *v = 0x4c }
func (v *PhysicalEnvironment) SetStudy()                               { *v = 0x4d }
func (v *PhysicalEnvironment) SetFrontGarden()                         { *v = 0x4e }
func (v *PhysicalEnvironment) SetBackGarden()                          { *v = 0x4f }
func (v *PhysicalEnvironment) SetKettle()                              { *v = 0x50 }
func (v *PhysicalEnvironment) SetTelevision()                          { *v = 0x51 }
func (v *PhysicalEnvironment) SetStove()                               { *v = 0x52 }
func (v *PhysicalEnvironment) SetMicrowave()                           { *v = 0x53 }
func (v *PhysicalEnvironment) SetToaster()                             { *v = 0x54 }
func (v *PhysicalEnvironment) SetVacuum()                              { *v = 0x55 }
func (v *PhysicalEnvironment) SetAppliance()                           { *v = 0x56 }
func (v *PhysicalEnvironment) SetFrontDoor()                           { *v = 0x57 }
func (v *PhysicalEnvironment) SetBackDoor()                            { *v = 0x58 }
func (v *PhysicalEnvironment) SetFridgeDoor()                          { *v = 0x59 }
func (v *PhysicalEnvironment) SetMedicationCabinetDoor()               { *v = 0x60 }
func (v *PhysicalEnvironment) SetWardrobeDoor()                        { *v = 0x61 }
func (v *PhysicalEnvironment) SetFrontCupboardDoor()                   { *v = 0x62 }
func (v *PhysicalEnvironment) SetOtherDoor()                           { *v = 0x63 }
func (v *PhysicalEnvironment) SetWaitingRoom()                         { *v = 0x64 }
func (v *PhysicalEnvironment) SetTriageRoom()                          { *v = 0x65 }
func (v *PhysicalEnvironment) SetDoctorSOffice()                       { *v = 0x66 }
func (v *PhysicalEnvironment) SetPatientSPrivateRoom()                 { *v = 0x67 }
func (v *PhysicalEnvironment) SetConsultationRoom()                    { *v = 0x68 }
func (v *PhysicalEnvironment) SetNurseStation()                        { *v = 0x69 }
func (v *PhysicalEnvironment) SetWard()                                { *v = 0x6a }
func (v *PhysicalEnvironment) SetCorridor()                            { *v = 0x6b }
func (v *PhysicalEnvironment) SetOperatingTheatre()                    { *v = 0x6c }
func (v *PhysicalEnvironment) SetDentalSurgeryRoom()                   { *v = 0x6d }
func (v *PhysicalEnvironment) SetMedicalImagingRoom()                  { *v = 0x6e }
func (v *PhysicalEnvironment) SetDecontaminationRoom()                 { *v = 0x6f }
func (v *PhysicalEnvironment) SetUnknownEnvironment()                  { *v = 0xFF }

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

const PowerAttr zcl.AttrID = 19

func (Power) ID() zcl.AttrID   { return PowerAttr }
func (Power) Readable() bool   { return true }
func (Power) Writable() bool   { return true }
func (Power) Reportable() bool { return false }
func (Power) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Power) AttrID() zcl.AttrID   { return v.ID() }
func (v Power) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Power) AttrValue() zcl.Val  { return v.Value() }

func (Power) Name() string        { return `Power` }
func (Power) Description() string { return `` }

type Power zcl.Zs16

func (v *Power) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *Power) Value() zcl.Val     { return v }

func (v Power) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *Power) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = Power(*nt)
	return br, err
}

func (v Power) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *Power) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Power(*a)
	return nil
}

func (v *Power) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = Power(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Power) String() string {
	return zcl.DecibelMilliWatts.Format(float64(v) / 100)
}

const PowerOnLevelAttr zcl.AttrID = 16384

func (PowerOnLevel) ID() zcl.AttrID   { return PowerOnLevelAttr }
func (PowerOnLevel) Readable() bool   { return true }
func (PowerOnLevel) Writable() bool   { return true }
func (PowerOnLevel) Reportable() bool { return false }
func (PowerOnLevel) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PowerOnLevel) AttrID() zcl.AttrID   { return v.ID() }
func (v PowerOnLevel) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PowerOnLevel) AttrValue() zcl.Val  { return v.Value() }

func (PowerOnLevel) Name() string        { return `Power On level` }
func (PowerOnLevel) Description() string { return `` }

type PowerOnLevel zcl.Zu8

func (v *PowerOnLevel) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *PowerOnLevel) Value() zcl.Val     { return v }

func (v PowerOnLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *PowerOnLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = PowerOnLevel(*nt)
	return br, err
}

func (v PowerOnLevel) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *PowerOnLevel) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PowerOnLevel(*a)
	return nil
}

func (v *PowerOnLevel) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = PowerOnLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PowerOnLevel) String() string {
	switch v {
	case 0xFF:
		return "Previous"
	}
	return zcl.Percent.Format(float64(v) / 2.54)
}

func (v PowerOnLevel) IsPrevious() bool { return v == 0xFF }
func (v *PowerOnLevel) SetPrevious()    { *v = 0xFF }

func (PowerOnLevel) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0xFF, Name: "Previous"},
	}
}

const PowerSourceAttr zcl.AttrID = 7

func (PowerSource) ID() zcl.AttrID   { return PowerSourceAttr }
func (PowerSource) Readable() bool   { return true }
func (PowerSource) Writable() bool   { return false }
func (PowerSource) Reportable() bool { return false }
func (PowerSource) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PowerSource) AttrID() zcl.AttrID   { return v.ID() }
func (v PowerSource) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PowerSource) AttrValue() zcl.Val  { return v.Value() }

func (PowerSource) Name() string        { return `Power Source` }
func (PowerSource) Description() string { return `` }

type PowerSource zcl.Zenum8

func (v *PowerSource) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *PowerSource) Value() zcl.Val     { return v }

func (v PowerSource) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *PowerSource) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = PowerSource(*nt)
	return br, err
}

func (v PowerSource) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *PowerSource) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PowerSource(*a)
	return nil
}

func (v *PowerSource) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = PowerSource(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PowerSource) String() string {
	switch v {
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
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v PowerSource) IsUnknown() bool                                          { return v == 0x00 }
func (v PowerSource) IsMainsSinglePhase() bool                                 { return v == 0x01 }
func (v PowerSource) IsMains3Phase() bool                                      { return v == 0x02 }
func (v PowerSource) IsBattery() bool                                          { return v == 0x03 }
func (v PowerSource) IsDcSource() bool                                         { return v == 0x04 }
func (v PowerSource) IsEmergencyMainsConstantlyPowered() bool                  { return v == 0x05 }
func (v PowerSource) IsEmergencyMainsAndTransferSwitch() bool                  { return v == 0x06 }
func (v PowerSource) IsUnknownWithBatteryBackup() bool                         { return v == 0x80 }
func (v PowerSource) IsMainsSinglePhaseWithBatteryBackup() bool                { return v == 0x81 }
func (v PowerSource) IsMains3PhaseWithBatteryBackup() bool                     { return v == 0x82 }
func (v PowerSource) IsBatteryWithBatteryBackup() bool                         { return v == 0x83 }
func (v PowerSource) IsDcSourceWithBatteryBackup() bool                        { return v == 0x84 }
func (v PowerSource) IsEmergencyMainsConstantlyPoweredWithBatteryBackup() bool { return v == 0x85 }
func (v PowerSource) IsEmergencyMainsAndTransferSwitchWithBatteryBackup() bool { return v == 0x86 }
func (v *PowerSource) SetUnknown()                                             { *v = 0x00 }
func (v *PowerSource) SetMainsSinglePhase()                                    { *v = 0x01 }
func (v *PowerSource) SetMains3Phase()                                         { *v = 0x02 }
func (v *PowerSource) SetBattery()                                             { *v = 0x03 }
func (v *PowerSource) SetDcSource()                                            { *v = 0x04 }
func (v *PowerSource) SetEmergencyMainsConstantlyPowered()                     { *v = 0x05 }
func (v *PowerSource) SetEmergencyMainsAndTransferSwitch()                     { *v = 0x06 }
func (v *PowerSource) SetUnknownWithBatteryBackup()                            { *v = 0x80 }
func (v *PowerSource) SetMainsSinglePhaseWithBatteryBackup()                   { *v = 0x81 }
func (v *PowerSource) SetMains3PhaseWithBatteryBackup()                        { *v = 0x82 }
func (v *PowerSource) SetBatteryWithBatteryBackup()                            { *v = 0x83 }
func (v *PowerSource) SetDcSourceWithBatteryBackup()                           { *v = 0x84 }
func (v *PowerSource) SetEmergencyMainsConstantlyPoweredWithBatteryBackup()    { *v = 0x85 }
func (v *PowerSource) SetEmergencyMainsAndTransferSwitchWithBatteryBackup()    { *v = 0x86 }

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

const PoweronOnOffAttr zcl.AttrID = 16387

func (PoweronOnOff) ID() zcl.AttrID   { return PoweronOnOffAttr }
func (PoweronOnOff) Readable() bool   { return true }
func (PoweronOnOff) Writable() bool   { return true }
func (PoweronOnOff) Reportable() bool { return false }
func (PoweronOnOff) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PoweronOnOff) AttrID() zcl.AttrID   { return v.ID() }
func (v PoweronOnOff) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PoweronOnOff) AttrValue() zcl.Val  { return v.Value() }

func (PoweronOnOff) Name() string        { return `PowerOn On/Off` }
func (PoweronOnOff) Description() string { return `` }

type PoweronOnOff zcl.Zenum8

func (v *PoweronOnOff) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *PoweronOnOff) Value() zcl.Val     { return v }

func (v PoweronOnOff) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *PoweronOnOff) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = PoweronOnOff(*nt)
	return br, err
}

func (v PoweronOnOff) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *PoweronOnOff) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PoweronOnOff(*a)
	return nil
}

func (v *PoweronOnOff) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = PoweronOnOff(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PoweronOnOff) String() string {
	switch v {
	case 0x00:
		return "Off"
	case 0x01:
		return "On"
	case 0xFF:
		return "Previous"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v PoweronOnOff) IsOff() bool      { return v == 0x00 }
func (v PoweronOnOff) IsOn() bool       { return v == 0x01 }
func (v PoweronOnOff) IsPrevious() bool { return v == 0xFF }
func (v *PoweronOnOff) SetOff()         { *v = 0x00 }
func (v *PoweronOnOff) SetOn()          { *v = 0x01 }
func (v *PoweronOnOff) SetPrevious()    { *v = 0xFF }

func (PoweronOnOff) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Off"},
		{Value: 0x01, Name: "On"},
		{Value: 0xFF, Name: "Previous"},
	}
}

const ProductUrlAttr zcl.AttrID = 11

func (ProductUrl) ID() zcl.AttrID   { return ProductUrlAttr }
func (ProductUrl) Readable() bool   { return true }
func (ProductUrl) Writable() bool   { return false }
func (ProductUrl) Reportable() bool { return false }
func (ProductUrl) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ProductUrl) AttrID() zcl.AttrID   { return v.ID() }
func (v ProductUrl) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ProductUrl) AttrValue() zcl.Val  { return v.Value() }

func (ProductUrl) Name() string { return `Product URL` }
func (ProductUrl) Description() string {
	return `specifies a link to a web page containing specific product information.`
}

// ProductUrl specifies a link to a web page containing specific product information.
type ProductUrl zcl.Zcstring

func (v *ProductUrl) TypeID() zcl.TypeID { return new(zcl.Zcstring).TypeID() }
func (v *ProductUrl) Value() zcl.Val     { return v }

func (v ProductUrl) MarshalZcl() ([]byte, error) { return zcl.Zcstring(v).MarshalZcl() }

func (v *ProductUrl) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*v = ProductUrl(*nt)
	return br, err
}

func (v ProductUrl) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zcstring(v))
}

func (v *ProductUrl) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zcstring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ProductUrl(*a)
	return nil
}

func (v *ProductUrl) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zcstring); ok {
		*v = ProductUrl(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ProductUrl) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v))
}

const ProductCodeAttr zcl.AttrID = 10

func (ProductCode) ID() zcl.AttrID   { return ProductCodeAttr }
func (ProductCode) Readable() bool   { return true }
func (ProductCode) Writable() bool   { return false }
func (ProductCode) Reportable() bool { return false }
func (ProductCode) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ProductCode) AttrID() zcl.AttrID   { return v.ID() }
func (v ProductCode) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ProductCode) AttrValue() zcl.Val  { return v.Value() }

func (ProductCode) Name() string        { return `Product code` }
func (ProductCode) Description() string { return `As printed on the product.` }

// ProductCode As printed on the product.
type ProductCode zcl.Zostring

func (v *ProductCode) TypeID() zcl.TypeID { return new(zcl.Zostring).TypeID() }
func (v *ProductCode) Value() zcl.Val     { return v }

func (v ProductCode) MarshalZcl() ([]byte, error) { return zcl.Zostring(v).MarshalZcl() }

func (v *ProductCode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*v = ProductCode(*nt)
	return br, err
}

func (v ProductCode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zostring(v))
}

func (v *ProductCode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zostring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ProductCode(*a)
	return nil
}

func (v *ProductCode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zostring); ok {
		*v = ProductCode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ProductCode) String() string {
	return zcl.Sprintf("%v", zcl.Zostring(v))
}

const QualityMeasureAttr zcl.AttrID = 3

func (QualityMeasure) ID() zcl.AttrID   { return QualityMeasureAttr }
func (QualityMeasure) Readable() bool   { return true }
func (QualityMeasure) Writable() bool   { return false }
func (QualityMeasure) Reportable() bool { return false }
func (QualityMeasure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v QualityMeasure) AttrID() zcl.AttrID   { return v.ID() }
func (v QualityMeasure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *QualityMeasure) AttrValue() zcl.Val  { return v.Value() }

func (QualityMeasure) Name() string        { return `Quality Measure` }
func (QualityMeasure) Description() string { return `` }

type QualityMeasure zcl.Zu8

func (v *QualityMeasure) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *QualityMeasure) Value() zcl.Val     { return v }

func (v QualityMeasure) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *QualityMeasure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = QualityMeasure(*nt)
	return br, err
}

func (v QualityMeasure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *QualityMeasure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = QualityMeasure(*a)
	return nil
}

func (v *QualityMeasure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = QualityMeasure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v QualityMeasure) String() string {
	return zcl.Percent.Format(float64(v))
}

func (QualityIndex) Name() string        { return `Quality index` }
func (QualityIndex) Description() string { return `` }

type QualityIndex zcl.Zu16

func (v *QualityIndex) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *QualityIndex) Value() zcl.Val     { return v }

func (v QualityIndex) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *QualityIndex) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = QualityIndex(*nt)
	return br, err
}

func (v QualityIndex) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *QualityIndex) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = QualityIndex(*a)
	return nil
}

func (v *QualityIndex) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = QualityIndex(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v QualityIndex) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (Rssi) Name() string        { return `RSSI` }
func (Rssi) Description() string { return `` }

type Rssi zcl.Zs8

func (v *Rssi) TypeID() zcl.TypeID { return new(zcl.Zs8).TypeID() }
func (v *Rssi) Value() zcl.Val     { return v }

func (v Rssi) MarshalZcl() ([]byte, error) { return zcl.Zs8(v).MarshalZcl() }

func (v *Rssi) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*v = Rssi(*nt)
	return br, err
}

func (v Rssi) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs8(v))
}

func (v *Rssi) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Rssi(*a)
	return nil
}

func (v *Rssi) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs8); ok {
		*v = Rssi(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Rssi) String() string {
	return zcl.DecibelMilliWatts.Format(float64(v))
}

func (Rate) Name() string        { return `Rate` }
func (Rate) Description() string { return `` }

type Rate zcl.Zu8

func (v *Rate) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *Rate) Value() zcl.Val     { return v }

func (v Rate) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *Rate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = Rate(*nt)
	return br, err
}

func (v Rate) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *Rate) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Rate(*a)
	return nil
}

func (v *Rate) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = Rate(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Rate) String() string {
	return zcl.PercentPerSecond.Format(float64(v) / 2.54)
}

const RelayedUcastAttr zcl.AttrID = 280

func (RelayedUcast) ID() zcl.AttrID   { return RelayedUcastAttr }
func (RelayedUcast) Readable() bool   { return true }
func (RelayedUcast) Writable() bool   { return false }
func (RelayedUcast) Reportable() bool { return false }
func (RelayedUcast) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RelayedUcast) AttrID() zcl.AttrID   { return v.ID() }
func (v RelayedUcast) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RelayedUcast) AttrValue() zcl.Val  { return v.Value() }

func (RelayedUcast) Name() string        { return `Relayed Ucast` }
func (RelayedUcast) Description() string { return `` }

type RelayedUcast zcl.Zu16

func (v *RelayedUcast) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RelayedUcast) Value() zcl.Val     { return v }

func (v RelayedUcast) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RelayedUcast) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RelayedUcast(*nt)
	return br, err
}

func (v RelayedUcast) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RelayedUcast) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RelayedUcast(*a)
	return nil
}

func (v *RelayedUcast) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RelayedUcast(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RelayedUcast) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RemainingTimeAttr zcl.AttrID = 1

func (RemainingTime) ID() zcl.AttrID   { return RemainingTimeAttr }
func (RemainingTime) Readable() bool   { return true }
func (RemainingTime) Writable() bool   { return false }
func (RemainingTime) Reportable() bool { return false }
func (RemainingTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RemainingTime) AttrID() zcl.AttrID   { return v.ID() }
func (v RemainingTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RemainingTime) AttrValue() zcl.Val  { return v.Value() }

func (RemainingTime) Name() string { return `Remaining Time` }
func (RemainingTime) Description() string {
	return `represents the time remaining until the current command is complete. It is specified in 1/10ths of a second.`
}

// RemainingTime represents the time remaining until the current command is complete. It is specified in 1/10ths of a second.
type RemainingTime zcl.Zu16

func (v *RemainingTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RemainingTime) Value() zcl.Val     { return v }

func (v RemainingTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RemainingTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RemainingTime(*nt)
	return br, err
}

func (v RemainingTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RemainingTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RemainingTime(*a)
	return nil
}

func (v *RemainingTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RemainingTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RemainingTime) String() string {
	return zcl.Seconds.Format(float64(v) / 10)
}

const ReportingPeriodAttr zcl.AttrID = 21

func (ReportingPeriod) ID() zcl.AttrID   { return ReportingPeriodAttr }
func (ReportingPeriod) Readable() bool   { return true }
func (ReportingPeriod) Writable() bool   { return true }
func (ReportingPeriod) Reportable() bool { return false }
func (ReportingPeriod) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ReportingPeriod) AttrID() zcl.AttrID   { return v.ID() }
func (v ReportingPeriod) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ReportingPeriod) AttrValue() zcl.Val  { return v.Value() }

func (ReportingPeriod) Name() string        { return `Reporting Period` }
func (ReportingPeriod) Description() string { return `` }

type ReportingPeriod zcl.Zu16

func (v *ReportingPeriod) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ReportingPeriod) Value() zcl.Val     { return v }

func (v ReportingPeriod) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ReportingPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ReportingPeriod(*nt)
	return br, err
}

func (v ReportingPeriod) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ReportingPeriod) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ReportingPeriod(*a)
	return nil
}

func (v *ReportingPeriod) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ReportingPeriod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ReportingPeriod) String() string {
	return zcl.Seconds.Format(float64(v))
}

func (Resolution) Name() string        { return `Resolution` }
func (Resolution) Description() string { return `` }

type Resolution zcl.Zenum8

func (v *Resolution) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *Resolution) Value() zcl.Val     { return v }

func (v Resolution) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *Resolution) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = Resolution(*nt)
	return br, err
}

func (v Resolution) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *Resolution) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Resolution(*a)
	return nil
}

func (v *Resolution) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = Resolution(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Resolution) String() string {
	switch v {
	case 0x00:
		return "High"
	case 0x01:
		return "Mid"
	case 0x02:
		return "Low"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v Resolution) IsHigh() bool { return v == 0x00 }
func (v Resolution) IsMid() bool  { return v == 0x01 }
func (v Resolution) IsLow() bool  { return v == 0x02 }
func (v *Resolution) SetHigh()    { *v = 0x00 }
func (v *Resolution) SetMid()     { *v = 0x01 }
func (v *Resolution) SetLow()     { *v = 0x02 }

func (Resolution) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "High"},
		{Value: 0x01, Name: "Mid"},
		{Value: 0x02, Name: "Low"},
	}
}

const RouteDiscInitiatedAttr zcl.AttrID = 268

func (RouteDiscInitiated) ID() zcl.AttrID   { return RouteDiscInitiatedAttr }
func (RouteDiscInitiated) Readable() bool   { return true }
func (RouteDiscInitiated) Writable() bool   { return false }
func (RouteDiscInitiated) Reportable() bool { return false }
func (RouteDiscInitiated) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RouteDiscInitiated) AttrID() zcl.AttrID   { return v.ID() }
func (v RouteDiscInitiated) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RouteDiscInitiated) AttrValue() zcl.Val  { return v.Value() }

func (RouteDiscInitiated) Name() string        { return `Route Disc Initiated` }
func (RouteDiscInitiated) Description() string { return `` }

type RouteDiscInitiated zcl.Zu16

func (v *RouteDiscInitiated) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RouteDiscInitiated) Value() zcl.Val     { return v }

func (v RouteDiscInitiated) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RouteDiscInitiated) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RouteDiscInitiated(*nt)
	return br, err
}

func (v RouteDiscInitiated) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RouteDiscInitiated) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RouteDiscInitiated(*a)
	return nil
}

func (v *RouteDiscInitiated) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RouteDiscInitiated(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RouteDiscInitiated) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const SwBuildIdAttr zcl.AttrID = 16384

func (SwBuildId) ID() zcl.AttrID   { return SwBuildIdAttr }
func (SwBuildId) Readable() bool   { return true }
func (SwBuildId) Writable() bool   { return false }
func (SwBuildId) Reportable() bool { return false }
func (SwBuildId) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v SwBuildId) AttrID() zcl.AttrID   { return v.ID() }
func (v SwBuildId) AttrType() zcl.TypeID { return v.TypeID() }
func (v *SwBuildId) AttrValue() zcl.Val  { return v.Value() }

func (SwBuildId) Name() string        { return `SW Build ID` }
func (SwBuildId) Description() string { return `` }

type SwBuildId zcl.Zcstring

func (v *SwBuildId) TypeID() zcl.TypeID { return new(zcl.Zcstring).TypeID() }
func (v *SwBuildId) Value() zcl.Val     { return v }

func (v SwBuildId) MarshalZcl() ([]byte, error) { return zcl.Zcstring(v).MarshalZcl() }

func (v *SwBuildId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*v = SwBuildId(*nt)
	return br, err
}

func (v SwBuildId) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zcstring(v))
}

func (v *SwBuildId) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zcstring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SwBuildId(*a)
	return nil
}

func (v *SwBuildId) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zcstring); ok {
		*v = SwBuildId(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SwBuildId) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v))
}

func (SceneCapacity) Name() string { return `Scene Capacity` }
func (SceneCapacity) Description() string {
	return `specifies remaining number of scenes that can be added`
}

// SceneCapacity specifies remaining number of scenes that can be added
type SceneCapacity zcl.Zu8

func (v *SceneCapacity) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *SceneCapacity) Value() zcl.Val     { return v }

func (v SceneCapacity) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *SceneCapacity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = SceneCapacity(*nt)
	return br, err
}

func (v SceneCapacity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *SceneCapacity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SceneCapacity(*a)
	return nil
}

func (v *SceneCapacity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = SceneCapacity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SceneCapacity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (SceneCopyMode) Name() string        { return `Scene Copy Mode` }
func (SceneCopyMode) Description() string { return `` }

type SceneCopyMode zcl.Zbmp8

func (v *SceneCopyMode) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *SceneCopyMode) Value() zcl.Val     { return v }

func (v SceneCopyMode) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *SceneCopyMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = SceneCopyMode(*nt)
	return br, err
}

func (v SceneCopyMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *SceneCopyMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SceneCopyMode(*a)
	return nil
}

func (v *SceneCopyMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = SceneCopyMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SceneCopyMode) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v SceneCopyMode) IsCopyAllScenes() bool    { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v *SceneCopyMode) SetCopyAllScenes(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b)) }

func (SceneCopyMode) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Copy All Scenes"},
	}
}

const SceneCountAttr zcl.AttrID = 0

func (SceneCount) ID() zcl.AttrID   { return SceneCountAttr }
func (SceneCount) Readable() bool   { return true }
func (SceneCount) Writable() bool   { return false }
func (SceneCount) Reportable() bool { return false }
func (SceneCount) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v SceneCount) AttrID() zcl.AttrID   { return v.ID() }
func (v SceneCount) AttrType() zcl.TypeID { return v.TypeID() }
func (v *SceneCount) AttrValue() zcl.Val  { return v.Value() }

func (SceneCount) Name() string        { return `Scene Count` }
func (SceneCount) Description() string { return `` }

type SceneCount zcl.Zu8

func (v *SceneCount) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *SceneCount) Value() zcl.Val     { return v }

func (v SceneCount) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *SceneCount) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = SceneCount(*nt)
	return br, err
}

func (v SceneCount) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *SceneCount) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SceneCount(*a)
	return nil
}

func (v *SceneCount) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = SceneCount(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SceneCount) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (SceneExtensions) Name() string { return `Scene Extensions` }
func (SceneExtensions) Description() string {
	return `The format of each extension field set is a 16 bit field carrying the cluster ID,
followed by an 8 bit length field and the set of scene extension fields specified
in  the  relevant  cluster. The length field holds the length in octets of that
extension field set. Extension field set format:
{{clusterId1, length 1, {extension field set 1}}, {clusterId2, length 2, {extension field set 2}} ...}
I.e. the field would be a repeating struct with [ClusterID uint16] [OctetLength uint8] [AttrID ...uint16]`
}

// SceneExtensions The format of each extension field set is a 16 bit field carrying the cluster ID,
// followed by an 8 bit length field and the set of scene extension fields specified
// in  the  relevant  cluster. The length field holds the length in octets of that
// extension field set. Extension field set format:
// {{clusterId1, length 1, {extension field set 1}}, {clusterId2, length 2, {extension field set 2}} ...}
// I.e. the field would be a repeating struct with [ClusterID uint16] [OctetLength uint8] [AttrID ...uint16]
type SceneExtensions zcl.SceneExtensionSet

func (v *SceneExtensions) TypeID() zcl.TypeID { return new(zcl.SceneExtensionSet).TypeID() }
func (v *SceneExtensions) Value() zcl.Val     { return v }

func (v SceneExtensions) MarshalZcl() ([]byte, error) { return zcl.SceneExtensionSet(v).MarshalZcl() }

func (v *SceneExtensions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.SceneExtensionSet)
	br, err := nt.UnmarshalZcl(b)
	*v = SceneExtensions(*nt)
	return br, err
}

func (v SceneExtensions) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.SceneExtensionSet(v))
}

func (v *SceneExtensions) UnmarshalJSON(b []byte) error {
	a := new(zcl.SceneExtensionSet)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SceneExtensions(*a)
	return nil
}

func (v *SceneExtensions) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.SceneExtensionSet); ok {
		*v = SceneExtensions(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SceneExtensions) String() string {
	return zcl.Sprintf("%v", zcl.SceneExtensionSet(v))
}

func (SceneId) Name() string        { return `Scene ID` }
func (SceneId) Description() string { return `` }

type SceneId zcl.Zu8

func (v *SceneId) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *SceneId) Value() zcl.Val     { return v }

func (v SceneId) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *SceneId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = SceneId(*nt)
	return br, err
}

func (v SceneId) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *SceneId) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SceneId(*a)
	return nil
}

func (v *SceneId) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = SceneId(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SceneId) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const SceneLastConfiguredByAttr zcl.AttrID = 5

func (SceneLastConfiguredBy) ID() zcl.AttrID   { return SceneLastConfiguredByAttr }
func (SceneLastConfiguredBy) Readable() bool   { return true }
func (SceneLastConfiguredBy) Writable() bool   { return false }
func (SceneLastConfiguredBy) Reportable() bool { return false }
func (SceneLastConfiguredBy) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v SceneLastConfiguredBy) AttrID() zcl.AttrID   { return v.ID() }
func (v SceneLastConfiguredBy) AttrType() zcl.TypeID { return v.TypeID() }
func (v *SceneLastConfiguredBy) AttrValue() zcl.Val  { return v.Value() }

func (SceneLastConfiguredBy) Name() string        { return `Scene Last Configured By` }
func (SceneLastConfiguredBy) Description() string { return `` }

type SceneLastConfiguredBy zcl.Zuid

func (v *SceneLastConfiguredBy) TypeID() zcl.TypeID { return new(zcl.Zuid).TypeID() }
func (v *SceneLastConfiguredBy) Value() zcl.Val     { return v }

func (v SceneLastConfiguredBy) MarshalZcl() ([]byte, error) { return zcl.Zuid(v).MarshalZcl() }

func (v *SceneLastConfiguredBy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*v = SceneLastConfiguredBy(*nt)
	return br, err
}

func (v SceneLastConfiguredBy) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zuid(v))
}

func (v *SceneLastConfiguredBy) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zuid)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SceneLastConfiguredBy(*a)
	return nil
}

func (v *SceneLastConfiguredBy) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zuid); ok {
		*v = SceneLastConfiguredBy(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SceneLastConfiguredBy) String() string {
	return zcl.Sprintf("%v", zcl.Zuid(v))
}

func (SceneName) Name() string        { return `Scene Name` }
func (SceneName) Description() string { return `` }

type SceneName zcl.Zcstring

func (v *SceneName) TypeID() zcl.TypeID { return new(zcl.Zcstring).TypeID() }
func (v *SceneName) Value() zcl.Val     { return v }

func (v SceneName) MarshalZcl() ([]byte, error) { return zcl.Zcstring(v).MarshalZcl() }

func (v *SceneName) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*v = SceneName(*nt)
	return br, err
}

func (v SceneName) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zcstring(v))
}

func (v *SceneName) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zcstring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SceneName(*a)
	return nil
}

func (v *SceneName) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zcstring); ok {
		*v = SceneName(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SceneName) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v))
}

const SceneNameSupportAttr zcl.AttrID = 4

func (SceneNameSupport) ID() zcl.AttrID   { return SceneNameSupportAttr }
func (SceneNameSupport) Readable() bool   { return true }
func (SceneNameSupport) Writable() bool   { return false }
func (SceneNameSupport) Reportable() bool { return false }
func (SceneNameSupport) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v SceneNameSupport) AttrID() zcl.AttrID   { return v.ID() }
func (v SceneNameSupport) AttrType() zcl.TypeID { return v.TypeID() }
func (v *SceneNameSupport) AttrValue() zcl.Val  { return v.Value() }

func (SceneNameSupport) Name() string        { return `Scene Name Support` }
func (SceneNameSupport) Description() string { return `` }

type SceneNameSupport zcl.Zbmp8

func (v *SceneNameSupport) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *SceneNameSupport) Value() zcl.Val     { return v }

func (v SceneNameSupport) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *SceneNameSupport) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = SceneNameSupport(*nt)
	return br, err
}

func (v SceneNameSupport) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *SceneNameSupport) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SceneNameSupport(*a)
	return nil
}

func (v *SceneNameSupport) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = SceneNameSupport(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SceneNameSupport) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v SceneNameSupport) IsNamesSupported() bool { return zcl.BitmapTest([]byte(v[:]), 7) }
func (v *SceneNameSupport) SetNamesSupported(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 7, b))
}

func (SceneNameSupport) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 7, Name: "Names Supported"},
	}
}

const SceneValidAttr zcl.AttrID = 3

func (SceneValid) ID() zcl.AttrID   { return SceneValidAttr }
func (SceneValid) Readable() bool   { return true }
func (SceneValid) Writable() bool   { return false }
func (SceneValid) Reportable() bool { return false }
func (SceneValid) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v SceneValid) AttrID() zcl.AttrID   { return v.ID() }
func (v SceneValid) AttrType() zcl.TypeID { return v.TypeID() }
func (v *SceneValid) AttrValue() zcl.Val  { return v.Value() }

func (SceneValid) Name() string        { return `Scene Valid` }
func (SceneValid) Description() string { return `` }

type SceneValid zcl.Zbool

func (v *SceneValid) TypeID() zcl.TypeID { return new(zcl.Zbool).TypeID() }
func (v *SceneValid) Value() zcl.Val     { return v }

func (v SceneValid) MarshalZcl() ([]byte, error) { return zcl.Zbool(v).MarshalZcl() }

func (v *SceneValid) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*v = SceneValid(*nt)
	return br, err
}

func (v SceneValid) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbool(v))
}

func (v *SceneValid) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbool)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SceneValid(*a)
	return nil
}

func (v *SceneValid) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbool); ok {
		*v = SceneValid(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SceneValid) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(v))
}

func (SceneList) Name() string        { return `Scene list` }
func (SceneList) Description() string { return `` }

type SceneList []*zcl.Zu8

func (v *SceneList) TypeID() zcl.TypeID { return new(zcl.Zset).TypeID() }
func (v *SceneList) Value() zcl.Val     { return v }

func (SceneList) ArrayTypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }

func (v *SceneList) ArrayValues() (o []uint8) {
	for _, a := range *v {
		o = append(o, uint8(*a))
	}
	return o
}

func (v *SceneList) SetValues(val []uint8) error {
	*v = []*zcl.Zu8{}
	return v.AddValues(val...)
}

func (v *SceneList) AddValues(val ...uint8) error {
	for _, a := range val {
		nv := zcl.Zu8(a)
		*v = append(*v, &nv)
	}
	return nil
}

func (v SceneList) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, a := range v {
		vars = append(vars, a)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (v *SceneList) UnmarshalZcl(b []byte) ([]byte, error) {
	*v = []*zcl.Zu8{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(zcl.Zu8)
		*v = append(*v, nv)
		return nv
	})
}

func (v *SceneList) SetValue(a zcl.Val) error {
	if nv, ok := a.(*SceneList); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SceneList) String() string {
	var s []string
	for _, a := range v {
		s = append(s, zcl.Sprintf("%v", a))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

const SensitivityAttr zcl.AttrID = 48

func (Sensitivity) ID() zcl.AttrID   { return SensitivityAttr }
func (Sensitivity) Readable() bool   { return true }
func (Sensitivity) Writable() bool   { return true }
func (Sensitivity) Reportable() bool { return false }
func (Sensitivity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Sensitivity) AttrID() zcl.AttrID   { return v.ID() }
func (v Sensitivity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Sensitivity) AttrValue() zcl.Val  { return v.Value() }

func (Sensitivity) Name() string        { return `Sensitivity` }
func (Sensitivity) Description() string { return `` }

type Sensitivity zcl.Zenum8

func (v *Sensitivity) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *Sensitivity) Value() zcl.Val     { return v }

func (v Sensitivity) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *Sensitivity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = Sensitivity(*nt)
	return br, err
}

func (v Sensitivity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *Sensitivity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Sensitivity(*a)
	return nil
}

func (v *Sensitivity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = Sensitivity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Sensitivity) String() string {
	switch v {
	case 0x00:
		return "default"
	case 0x01:
		return "high"
	case 0x02:
		return "max"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v Sensitivity) IsDefault() bool { return v == 0x00 }
func (v Sensitivity) IsHigh() bool    { return v == 0x01 }
func (v Sensitivity) IsMax() bool     { return v == 0x02 }
func (v *Sensitivity) SetDefault()    { *v = 0x00 }
func (v *Sensitivity) SetHigh()       { *v = 0x01 }
func (v *Sensitivity) SetMax()        { *v = 0x02 }

func (Sensitivity) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "default"},
		{Value: 0x01, Name: "high"},
		{Value: 0x02, Name: "max"},
	}
}

const ShortPollIntervalAttr zcl.AttrID = 2

func (ShortPollInterval) ID() zcl.AttrID   { return ShortPollIntervalAttr }
func (ShortPollInterval) Readable() bool   { return true }
func (ShortPollInterval) Writable() bool   { return false }
func (ShortPollInterval) Reportable() bool { return false }
func (ShortPollInterval) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ShortPollInterval) AttrID() zcl.AttrID   { return v.ID() }
func (v ShortPollInterval) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ShortPollInterval) AttrValue() zcl.Val  { return v.Value() }

func (ShortPollInterval) Name() string        { return `Short Poll Interval` }
func (ShortPollInterval) Description() string { return `` }

type ShortPollInterval zcl.Zu16

func (v *ShortPollInterval) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ShortPollInterval) Value() zcl.Val     { return v }

func (v ShortPollInterval) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ShortPollInterval) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ShortPollInterval(*nt)
	return br, err
}

func (v ShortPollInterval) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ShortPollInterval) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ShortPollInterval(*a)
	return nil
}

func (v *ShortPollInterval) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ShortPollInterval(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ShortPollInterval) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const StackVersionAttr zcl.AttrID = 2

func (StackVersion) ID() zcl.AttrID   { return StackVersionAttr }
func (StackVersion) Readable() bool   { return true }
func (StackVersion) Writable() bool   { return false }
func (StackVersion) Reportable() bool { return false }
func (StackVersion) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v StackVersion) AttrID() zcl.AttrID   { return v.ID() }
func (v StackVersion) AttrType() zcl.TypeID { return v.TypeID() }
func (v *StackVersion) AttrValue() zcl.Val  { return v.Value() }

func (StackVersion) Name() string        { return `Stack Version` }
func (StackVersion) Description() string { return `` }

type StackVersion zcl.Zu8

func (v *StackVersion) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *StackVersion) Value() zcl.Val     { return v }

func (v StackVersion) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *StackVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = StackVersion(*nt)
	return br, err
}

func (v StackVersion) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *StackVersion) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = StackVersion(*a)
	return nil
}

func (v *StackVersion) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = StackVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v StackVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const StandardTimeAttr zcl.AttrID = 6

func (StandardTime) ID() zcl.AttrID   { return StandardTimeAttr }
func (StandardTime) Readable() bool   { return true }
func (StandardTime) Writable() bool   { return false }
func (StandardTime) Reportable() bool { return false }
func (StandardTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v StandardTime) AttrID() zcl.AttrID   { return v.ID() }
func (v StandardTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *StandardTime) AttrValue() zcl.Val  { return v.Value() }

func (StandardTime) Name() string        { return `Standard Time` }
func (StandardTime) Description() string { return `Local time (without DST offset)` }

// StandardTime Local time (without DST offset)
type StandardTime zcl.Zu32

func (v *StandardTime) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *StandardTime) Value() zcl.Val     { return v }

func (v StandardTime) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *StandardTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = StandardTime(*nt)
	return br, err
}

func (v StandardTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *StandardTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = StandardTime(*a)
	return nil
}

func (v *StandardTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = StandardTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v StandardTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

func (Status) Name() string        { return `Status` }
func (Status) Description() string { return `` }

type Status zcl.Status

func (v *Status) TypeID() zcl.TypeID { return new(zcl.Status).TypeID() }
func (v *Status) Value() zcl.Val     { return v }

func (v Status) MarshalZcl() ([]byte, error) { return zcl.Status(v).MarshalZcl() }

func (v *Status) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Status)
	br, err := nt.UnmarshalZcl(b)
	*v = Status(*nt)
	return br, err
}

func (v Status) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Status(v))
}

func (v *Status) UnmarshalJSON(b []byte) error {
	a := new(zcl.Status)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Status(*a)
	return nil
}

func (v *Status) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Status); ok {
		*v = Status(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Status) String() string {
	return zcl.Sprintf("%v", zcl.Status(v))
}

func (StepSize) Name() string        { return `Step size` }
func (StepSize) Description() string { return `` }

type StepSize zcl.Zu8

func (v *StepSize) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *StepSize) Value() zcl.Val     { return v }

func (v StepSize) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *StepSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = StepSize(*nt)
	return br, err
}

func (v StepSize) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *StepSize) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = StepSize(*a)
	return nil
}

func (v *StepSize) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = StepSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v StepSize) String() string {
	return zcl.Percent.Format(float64(v) / 2.54)
}

const SwitchActionsAttr zcl.AttrID = 16

func (SwitchActions) ID() zcl.AttrID   { return SwitchActionsAttr }
func (SwitchActions) Readable() bool   { return true }
func (SwitchActions) Writable() bool   { return true }
func (SwitchActions) Reportable() bool { return false }
func (SwitchActions) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v SwitchActions) AttrID() zcl.AttrID   { return v.ID() }
func (v SwitchActions) AttrType() zcl.TypeID { return v.TypeID() }
func (v *SwitchActions) AttrValue() zcl.Val  { return v.Value() }

func (SwitchActions) Name() string { return `Switch actions` }
func (SwitchActions) Description() string {
	return `specifies the commands of the On/Off cluster to be generated when the switch moves between its two states.`
}

// SwitchActions specifies the commands of the On/Off cluster to be generated when the switch moves between its two states.
type SwitchActions zcl.Zenum8

func (v *SwitchActions) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *SwitchActions) Value() zcl.Val     { return v }

func (v SwitchActions) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *SwitchActions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = SwitchActions(*nt)
	return br, err
}

func (v SwitchActions) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *SwitchActions) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SwitchActions(*a)
	return nil
}

func (v *SwitchActions) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = SwitchActions(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SwitchActions) String() string {
	switch v {
	case 0x00:
		return "On-Off"
	case 0x01:
		return "Off-On"
	case 0x02:
		return "Toggle"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v SwitchActions) IsOnOff() bool  { return v == 0x00 }
func (v SwitchActions) IsOffOn() bool  { return v == 0x01 }
func (v SwitchActions) IsToggle() bool { return v == 0x02 }
func (v *SwitchActions) SetOnOff()     { *v = 0x00 }
func (v *SwitchActions) SetOffOn()     { *v = 0x01 }
func (v *SwitchActions) SetToggle()    { *v = 0x02 }

func (SwitchActions) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "On-Off"},
		{Value: 0x01, Name: "Off-On"},
		{Value: 0x02, Name: "Toggle"},
	}
}

const SwitchTypeAttr zcl.AttrID = 0

func (SwitchType) ID() zcl.AttrID   { return SwitchTypeAttr }
func (SwitchType) Readable() bool   { return true }
func (SwitchType) Writable() bool   { return false }
func (SwitchType) Reportable() bool { return false }
func (SwitchType) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v SwitchType) AttrID() zcl.AttrID   { return v.ID() }
func (v SwitchType) AttrType() zcl.TypeID { return v.TypeID() }
func (v *SwitchType) AttrValue() zcl.Val  { return v.Value() }

func (SwitchType) Name() string { return `Switch type` }
func (SwitchType) Description() string {
	return `specifies the basic functionality of the On/Off switching device.`
}

// SwitchType specifies the basic functionality of the On/Off switching device.
type SwitchType zcl.Zenum8

func (v *SwitchType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *SwitchType) Value() zcl.Val     { return v }

func (v SwitchType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *SwitchType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = SwitchType(*nt)
	return br, err
}

func (v SwitchType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *SwitchType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SwitchType(*a)
	return nil
}

func (v *SwitchType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = SwitchType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SwitchType) String() string {
	switch v {
	case 0x00:
		return "Toggle"
	case 0x01:
		return "Momentary"
	case 0x02:
		return "Multifunction"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v SwitchType) IsToggle() bool        { return v == 0x00 }
func (v SwitchType) IsMomentary() bool     { return v == 0x01 }
func (v SwitchType) IsMultifunction() bool { return v == 0x02 }
func (v *SwitchType) SetToggle()           { *v = 0x00 }
func (v *SwitchType) SetMomentary()        { *v = 0x01 }
func (v *SwitchType) SetMultifunction()    { *v = 0x02 }

func (SwitchType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Toggle"},
		{Value: 0x01, Name: "Momentary"},
		{Value: 0x02, Name: "Multifunction"},
	}
}

const TimeAttr zcl.AttrID = 0

func (Time) ID() zcl.AttrID   { return TimeAttr }
func (Time) Readable() bool   { return true }
func (Time) Writable() bool   { return true }
func (Time) Reportable() bool { return false }
func (Time) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Time) AttrID() zcl.AttrID   { return v.ID() }
func (v Time) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Time) AttrValue() zcl.Val  { return v.Value() }

func (Time) Name() string        { return `Time` }
func (Time) Description() string { return `` }

type Time zcl.Zutc

func (v *Time) TypeID() zcl.TypeID { return new(zcl.Zutc).TypeID() }
func (v *Time) Value() zcl.Val     { return v }

func (v Time) MarshalZcl() ([]byte, error) { return zcl.Zutc(v).MarshalZcl() }

func (v *Time) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*v = Time(*nt)
	return br, err
}

func (v Time) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zutc(v))
}

func (v *Time) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zutc)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Time(*a)
	return nil
}

func (v *Time) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zutc); ok {
		*v = Time(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Time) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(v))
}

const TimeStatusAttr zcl.AttrID = 1

func (TimeStatus) ID() zcl.AttrID   { return TimeStatusAttr }
func (TimeStatus) Readable() bool   { return true }
func (TimeStatus) Writable() bool   { return true }
func (TimeStatus) Reportable() bool { return false }
func (TimeStatus) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v TimeStatus) AttrID() zcl.AttrID   { return v.ID() }
func (v TimeStatus) AttrType() zcl.TypeID { return v.TypeID() }
func (v *TimeStatus) AttrValue() zcl.Val  { return v.Value() }

func (TimeStatus) Name() string        { return `Time Status` }
func (TimeStatus) Description() string { return `` }

type TimeStatus zcl.Zbmp8

func (v *TimeStatus) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *TimeStatus) Value() zcl.Val     { return v }

func (v TimeStatus) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *TimeStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = TimeStatus(*nt)
	return br, err
}

func (v TimeStatus) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *TimeStatus) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TimeStatus(*a)
	return nil
}

func (v *TimeStatus) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = TimeStatus(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TimeStatus) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v TimeStatus) IsMasterClock() bool             { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v TimeStatus) IsSynchronized() bool            { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v TimeStatus) IsMasterForTimezoneAndDst() bool { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v TimeStatus) IsSuperseding() bool             { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v *TimeStatus) SetMasterClock(b bool)          { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b)) }
func (v *TimeStatus) SetSynchronized(b bool)         { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b)) }
func (v *TimeStatus) SetMasterForTimezoneAndDst(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}
func (v *TimeStatus) SetSuperseding(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b)) }

func (TimeStatus) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Master Clock"},
		{Value: 1, Name: "Synchronized"},
		{Value: 2, Name: "Master for Timezone and Dst"},
		{Value: 3, Name: "Superseding"},
	}
}

const TimeZoneAttr zcl.AttrID = 2

func (TimeZone) ID() zcl.AttrID   { return TimeZoneAttr }
func (TimeZone) Readable() bool   { return true }
func (TimeZone) Writable() bool   { return true }
func (TimeZone) Reportable() bool { return false }
func (TimeZone) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v TimeZone) AttrID() zcl.AttrID   { return v.ID() }
func (v TimeZone) AttrType() zcl.TypeID { return v.TypeID() }
func (v *TimeZone) AttrValue() zcl.Val  { return v.Value() }

func (TimeZone) Name() string        { return `Time Zone` }
func (TimeZone) Description() string { return `Offset during normal time from UTC in seconds` }

// TimeZone Offset during normal time from UTC in seconds
type TimeZone zcl.Zs32

func (v *TimeZone) TypeID() zcl.TypeID { return new(zcl.Zs32).TypeID() }
func (v *TimeZone) Value() zcl.Val     { return v }

func (v TimeZone) MarshalZcl() ([]byte, error) { return zcl.Zs32(v).MarshalZcl() }

func (v *TimeZone) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs32)
	br, err := nt.UnmarshalZcl(b)
	*v = TimeZone(*nt)
	return br, err
}

func (v TimeZone) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs32(v))
}

func (v *TimeZone) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TimeZone(*a)
	return nil
}

func (v *TimeZone) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs32); ok {
		*v = TimeZone(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TimeZone) String() string {
	return zcl.Seconds.Format(float64(v))
}

func (TransitionTime) Name() string        { return `Transition Time` }
func (TransitionTime) Description() string { return `` }

type TransitionTime zcl.Zu16

func (v *TransitionTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *TransitionTime) Value() zcl.Val     { return v }

func (v TransitionTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *TransitionTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = TransitionTime(*nt)
	return br, err
}

func (v TransitionTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *TransitionTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TransitionTime(*a)
	return nil
}

func (v *TransitionTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = TransitionTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TransitionTime) String() string {
	return zcl.Seconds.Format(float64(v) / 10)
}

func (TransitionTimeSec) Name() string        { return `Transition time (Sec)` }
func (TransitionTimeSec) Description() string { return `` }

type TransitionTimeSec zcl.Zu16

func (v *TransitionTimeSec) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *TransitionTimeSec) Value() zcl.Val     { return v }

func (v TransitionTimeSec) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *TransitionTimeSec) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = TransitionTimeSec(*nt)
	return br, err
}

func (v TransitionTimeSec) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *TransitionTimeSec) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TransitionTimeSec(*a)
	return nil
}

func (v *TransitionTimeSec) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = TransitionTimeSec(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TransitionTimeSec) String() string {
	return zcl.Seconds.Format(float64(v) / 10)
}

const UserTestAttr zcl.AttrID = 50

func (UserTest) ID() zcl.AttrID   { return UserTestAttr }
func (UserTest) Readable() bool   { return true }
func (UserTest) Writable() bool   { return true }
func (UserTest) Reportable() bool { return false }
func (UserTest) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v UserTest) AttrID() zcl.AttrID   { return v.ID() }
func (v UserTest) AttrType() zcl.TypeID { return v.TypeID() }
func (v *UserTest) AttrValue() zcl.Val  { return v.Value() }

func (UserTest) Name() string        { return `User test` }
func (UserTest) Description() string { return `` }

type UserTest zcl.Zbool

func (v *UserTest) TypeID() zcl.TypeID { return new(zcl.Zbool).TypeID() }
func (v *UserTest) Value() zcl.Val     { return v }

func (v UserTest) MarshalZcl() ([]byte, error) { return zcl.Zbool(v).MarshalZcl() }

func (v *UserTest) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*v = UserTest(*nt)
	return br, err
}

func (v UserTest) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbool(v))
}

func (v *UserTest) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbool)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = UserTest(*a)
	return nil
}

func (v *UserTest) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbool); ok {
		*v = UserTest(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v UserTest) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(v))
}

const ValidUntilTimeAttr zcl.AttrID = 9

func (ValidUntilTime) ID() zcl.AttrID   { return ValidUntilTimeAttr }
func (ValidUntilTime) Readable() bool   { return true }
func (ValidUntilTime) Writable() bool   { return true }
func (ValidUntilTime) Reportable() bool { return false }
func (ValidUntilTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ValidUntilTime) AttrID() zcl.AttrID   { return v.ID() }
func (v ValidUntilTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ValidUntilTime) AttrValue() zcl.Val  { return v.Value() }

func (ValidUntilTime) Name() string        { return `Valid Until Time` }
func (ValidUntilTime) Description() string { return `` }

type ValidUntilTime zcl.Zutc

func (v *ValidUntilTime) TypeID() zcl.TypeID { return new(zcl.Zutc).TypeID() }
func (v *ValidUntilTime) Value() zcl.Val     { return v }

func (v ValidUntilTime) MarshalZcl() ([]byte, error) { return zcl.Zutc(v).MarshalZcl() }

func (v *ValidUntilTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*v = ValidUntilTime(*nt)
	return br, err
}

func (v ValidUntilTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zutc(v))
}

func (v *ValidUntilTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zutc)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ValidUntilTime(*a)
	return nil
}

func (v *ValidUntilTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zutc); ok {
		*v = ValidUntilTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ValidUntilTime) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(v))
}

const XCoordinateAttr zcl.AttrID = 16

func (XCoordinate) ID() zcl.AttrID   { return XCoordinateAttr }
func (XCoordinate) Readable() bool   { return true }
func (XCoordinate) Writable() bool   { return true }
func (XCoordinate) Reportable() bool { return false }
func (XCoordinate) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v XCoordinate) AttrID() zcl.AttrID   { return v.ID() }
func (v XCoordinate) AttrType() zcl.TypeID { return v.TypeID() }
func (v *XCoordinate) AttrValue() zcl.Val  { return v.Value() }

func (XCoordinate) Name() string        { return `X Coordinate` }
func (XCoordinate) Description() string { return `` }

type XCoordinate zcl.Zs16

func (v *XCoordinate) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *XCoordinate) Value() zcl.Val     { return v }

func (v XCoordinate) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *XCoordinate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = XCoordinate(*nt)
	return br, err
}

func (v XCoordinate) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *XCoordinate) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = XCoordinate(*a)
	return nil
}

func (v *XCoordinate) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = XCoordinate(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v XCoordinate) String() string {
	return zcl.Meters.Format(float64(v) / 10)
}

const YCoordinateAttr zcl.AttrID = 17

func (YCoordinate) ID() zcl.AttrID   { return YCoordinateAttr }
func (YCoordinate) Readable() bool   { return true }
func (YCoordinate) Writable() bool   { return true }
func (YCoordinate) Reportable() bool { return false }
func (YCoordinate) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v YCoordinate) AttrID() zcl.AttrID   { return v.ID() }
func (v YCoordinate) AttrType() zcl.TypeID { return v.TypeID() }
func (v *YCoordinate) AttrValue() zcl.Val  { return v.Value() }

func (YCoordinate) Name() string        { return `Y Coordinate` }
func (YCoordinate) Description() string { return `` }

type YCoordinate zcl.Zs16

func (v *YCoordinate) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *YCoordinate) Value() zcl.Val     { return v }

func (v YCoordinate) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *YCoordinate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = YCoordinate(*nt)
	return br, err
}

func (v YCoordinate) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *YCoordinate) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = YCoordinate(*a)
	return nil
}

func (v *YCoordinate) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = YCoordinate(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v YCoordinate) String() string {
	return zcl.Meters.Format(float64(v) / 10)
}

const ZCoordinateAttr zcl.AttrID = 18

func (ZCoordinate) ID() zcl.AttrID   { return ZCoordinateAttr }
func (ZCoordinate) Readable() bool   { return true }
func (ZCoordinate) Writable() bool   { return true }
func (ZCoordinate) Reportable() bool { return false }
func (ZCoordinate) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ZCoordinate) AttrID() zcl.AttrID   { return v.ID() }
func (v ZCoordinate) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ZCoordinate) AttrValue() zcl.Val  { return v.Value() }

func (ZCoordinate) Name() string        { return `Z Coordinate` }
func (ZCoordinate) Description() string { return `` }

type ZCoordinate zcl.Zs16

func (v *ZCoordinate) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ZCoordinate) Value() zcl.Val     { return v }

func (v ZCoordinate) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ZCoordinate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ZCoordinate(*nt)
	return br, err
}

func (v ZCoordinate) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ZCoordinate) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ZCoordinate(*a)
	return nil
}

func (v *ZCoordinate) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ZCoordinate(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ZCoordinate) String() string {
	return zcl.Meters.Format(float64(v) / 10)
}

const ZclVersionAttr zcl.AttrID = 0

func (ZclVersion) ID() zcl.AttrID   { return ZclVersionAttr }
func (ZclVersion) Readable() bool   { return true }
func (ZclVersion) Writable() bool   { return false }
func (ZclVersion) Reportable() bool { return false }
func (ZclVersion) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ZclVersion) AttrID() zcl.AttrID   { return v.ID() }
func (v ZclVersion) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ZclVersion) AttrValue() zcl.Val  { return v.Value() }

func (ZclVersion) Name() string        { return `ZCL Version` }
func (ZclVersion) Description() string { return `` }

type ZclVersion zcl.Zu8

func (v *ZclVersion) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *ZclVersion) Value() zcl.Val     { return v }

func (v ZclVersion) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *ZclVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = ZclVersion(*nt)
	return br, err
}

func (v ZclVersion) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *ZclVersion) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ZclVersion(*a)
	return nil
}

func (v *ZclVersion) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = ZclVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ZclVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}
