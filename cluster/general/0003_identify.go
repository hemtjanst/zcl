// Attributes and commands for putting a device into Identification mode (e.g. flashing a light)
package general

import (
	"neotor.se/zcl"
)

// Identify
const IdentifyID zcl.ClusterID = 3

var IdentifyCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		IdentifyCommand:      func() zcl.Command { return new(Identify) },
		IdentifyQueryCommand: func() zcl.Command { return new(IdentifyQuery) },
		TriggerEffectCommand: func() zcl.Command { return new(TriggerEffect) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{
		IdentifyQueryResponseCommand: func() zcl.Command { return new(IdentifyQueryResponse) },
	},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		IdentifyTimeAttr: func() zcl.Attr { return new(IdentifyTime) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

// Start or stop the device identifying itself.
type Identify struct {
	// The time in seconds for which a device will stay in identify mode.
	IdentifyTime zcl.Zu16
}

const IdentifyCommand zcl.CommandID = 0

func (v *Identify) Values() []zcl.Val {
	return []zcl.Val{
		&v.IdentifyTime,
	}
}

func (v Identify) ID() zcl.CommandID {
	return IdentifyCommand
}

func (v Identify) Cluster() zcl.ClusterID {
	return IdentifyID
}

func (v Identify) MnfCode() []byte {
	return []byte{}
}

func (v Identify) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.IdentifyTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *Identify) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.IdentifyTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Allows the sending device to request the target or targets to respond if they are currently identifying themselves.
type IdentifyQuery struct {
}

const IdentifyQueryCommand zcl.CommandID = 1

func (v *IdentifyQuery) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v IdentifyQuery) ID() zcl.CommandID {
	return IdentifyQueryCommand
}

func (v IdentifyQuery) Cluster() zcl.ClusterID {
	return IdentifyID
}

func (v IdentifyQuery) MnfCode() []byte {
	return []byte{}
}

func (v IdentifyQuery) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *IdentifyQuery) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// The trigger effect command allows the support of feedback to the user, such as a certain light effect.
type TriggerEffect struct {
	// The effect identifier field specifies the identify effect to use.
	EffectIdentifier zcl.Zenum8
	// The effect identifier field specifies the identify effect to use.
	EffectVariant zcl.Zenum8
}

const TriggerEffectCommand zcl.CommandID = 64

func (v *TriggerEffect) Values() []zcl.Val {
	return []zcl.Val{
		&v.EffectIdentifier,
		&v.EffectVariant,
	}
}

func (v TriggerEffect) ID() zcl.CommandID {
	return TriggerEffectCommand
}

func (v TriggerEffect) Cluster() zcl.ClusterID {
	return IdentifyID
}

func (v TriggerEffect) MnfCode() []byte {
	return []byte{}
}

func (v TriggerEffect) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.EffectIdentifier.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.EffectVariant.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *TriggerEffect) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.EffectIdentifier).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.EffectVariant).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Response of a identify query command.
type IdentifyQueryResponse struct {
	// The time in seconds for which a device will stay in identify mode.
	Timeout zcl.Zu16
}

const IdentifyQueryResponseCommand zcl.CommandID = 0

func (v *IdentifyQueryResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Timeout,
	}
}

func (v IdentifyQueryResponse) ID() zcl.CommandID {
	return IdentifyQueryResponseCommand
}

func (v IdentifyQueryResponse) Cluster() zcl.ClusterID {
	return IdentifyID
}

func (v IdentifyQueryResponse) MnfCode() []byte {
	return []byte{}
}

func (v IdentifyQueryResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Timeout.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *IdentifyQueryResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Timeout).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

const IdentifyTimeAttr zcl.AttrID = 0

type IdentifyTime zcl.Zu16

func (a IdentifyTime) ID() zcl.AttrID         { return IdentifyTimeAttr }
func (a IdentifyTime) Cluster() zcl.ClusterID { return IdentifyID }
func (a *IdentifyTime) Value() *IdentifyTime  { return a }
func (a IdentifyTime) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *IdentifyTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = IdentifyTime(*nt)
	return br, err
}

func (a IdentifyTime) Readable() bool   { return true }
func (a IdentifyTime) Writable() bool   { return true }
func (a IdentifyTime) Reportable() bool { return false }
func (a IdentifyTime) SceneIndex() int  { return -1 }

func (a IdentifyTime) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}