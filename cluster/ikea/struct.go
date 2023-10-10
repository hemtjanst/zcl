// IKEA Specific clusters
package ikea

import "hemtjan.st/zcl"

type CommandID = zcl.CommandID
type Frame = zcl.ReceivedZclFrame

const VocIndexAttr zcl.AttrID = 0

func (VocIndex) ID() zcl.AttrID   { return VocIndexAttr }
func (VocIndex) Readable() bool   { return true }
func (VocIndex) Writable() bool   { return false }
func (VocIndex) Reportable() bool { return true }
func (VocIndex) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v VocIndex) AttrID() zcl.AttrID   { return v.ID() }
func (v VocIndex) AttrType() zcl.TypeID { return v.TypeID() }
func (v *VocIndex) AttrValue() zcl.Val  { return v.Value() }

func (VocIndex) Name() string { return `VOC Index` }
func (VocIndex) Description() string {
	return `represents the Sensirion AG measurement of VOC status relative to the sensors recent history.
The VOC Index uses a moving average over the past 24 hours (called the "learning time") as offset.
The VOC Index mimics the human nose’s perception of odors with a relative intensity compared to recent history.
A VOC Index above 100 means that there are more VOCs compared to the average (e.g., induced by a VOC event from
cooking, cleaning, breathing, etc.) while a VOC Index below 100 means that there are fewer VOCs compared to the
average (e.g., induced by fresh air from an open window, using an air purifier, etc.).
The VOC Index ranges from 1 to 500.`
}

// VocIndex represents the Sensirion AG measurement of VOC status relative to the sensors recent history.
// The VOC Index uses a moving average over the past 24 hours (called the "learning time") as offset.
// The VOC Index mimics the human nose’s perception of odors with a relative intensity compared to recent history.
// A VOC Index above 100 means that there are more VOCs compared to the average (e.g., induced by a VOC event from
// cooking, cleaning, breathing, etc.) while a VOC Index below 100 means that there are fewer VOCs compared to the
// average (e.g., induced by fresh air from an open window, using an air purifier, etc.).
// The VOC Index ranges from 1 to 500.
type VocIndex zcl.Zfloat

func (v *VocIndex) TypeID() zcl.TypeID { return new(zcl.Zfloat).TypeID() }
func (v *VocIndex) Value() zcl.Val     { return v }

func (v VocIndex) MarshalZcl() ([]byte, error) { return zcl.Zfloat(v).MarshalZcl() }

func (v *VocIndex) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*v = VocIndex(*nt)
	return br, err
}

func (v VocIndex) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zfloat(v))
}

func (v *VocIndex) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zfloat)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = VocIndex(*a)
	return nil
}

func (v *VocIndex) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zfloat); ok {
		*v = VocIndex(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v VocIndex) String() string {
	return zcl.Sprintf("%v", zcl.Zfloat(v))
}
