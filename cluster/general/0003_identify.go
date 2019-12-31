package general

import "hemtjan.st/zcl"

// Identify
const IdentifyID zcl.ClusterID = 3

var IdentifyCluster = zcl.Cluster{
	Name: "Identify",
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

// Identify Start or stop the device identifying itself.
type Identify struct {
	// IdentifyTime The time in seconds for which a device will stay in identify mode.
	IdentifyTime IdentifyTime
}

// IdentifyCommand is the Command ID of Identify
const IdentifyCommand CommandID = 0x0000

// Values returns all values of Identify
func (v *Identify) Values() []zcl.Val {
	return []zcl.Val{
		&v.IdentifyTime,
	}
}

// Name of the command (needed to fulfill interface)
func (Identify) Name() string { return "Identify" }

// ID of the command (needed to fulfill interface)
func (Identify) ID() CommandID { return IdentifyCommand }

// Cluster ID of the command (needed to fulfill interface)
func (Identify) Cluster() zcl.ClusterID { return IdentifyID }

// MnfCode returns the manufacturer code (if any) of the command
func (Identify) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of Identify
func (v Identify) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.IdentifyTime.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the Identify struct
func (v *Identify) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.IdentifyTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v Identify) String() string {
	return zcl.Sprintf(
		"Identify{"+zcl.StrJoin([]string{
			"IdentifyTime(%v)",
		}, " ")+"}",
		v.IdentifyTime,
	)
}

// IdentifyQuery Allows the sending device to request the target or targets to respond if they are currently identifying themselves.
type IdentifyQuery struct {
}

// IdentifyQueryCommand is the Command ID of IdentifyQuery
const IdentifyQueryCommand CommandID = 0x0001

// Values returns all values of IdentifyQuery
func (v *IdentifyQuery) Values() []zcl.Val {
	return []zcl.Val{}
}

// Name of the command (needed to fulfill interface)
func (IdentifyQuery) Name() string { return "Identify Query" }

// ID of the command (needed to fulfill interface)
func (IdentifyQuery) ID() CommandID { return IdentifyQueryCommand }

// Cluster ID of the command (needed to fulfill interface)
func (IdentifyQuery) Cluster() zcl.ClusterID { return IdentifyID }

// MnfCode returns the manufacturer code (if any) of the command
func (IdentifyQuery) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of IdentifyQuery
func (v IdentifyQuery) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the IdentifyQuery struct
func (v *IdentifyQuery) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v IdentifyQuery) String() string {
	return zcl.Sprintf(
		"IdentifyQuery{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

// TriggerEffect The trigger effect command allows the support of feedback to the user, such as a certain light effect.
type TriggerEffect struct {
	// IdentifyEffect The effect identifier field specifies the identify effect to use.
	IdentifyEffect IdentifyEffect
	// IdentifyEffectVariant The effect identifier field specifies the identify effect to use.
	IdentifyEffectVariant IdentifyEffectVariant
}

// TriggerEffectCommand is the Command ID of TriggerEffect
const TriggerEffectCommand CommandID = 0x0040

// Values returns all values of TriggerEffect
func (v *TriggerEffect) Values() []zcl.Val {
	return []zcl.Val{
		&v.IdentifyEffect,
		&v.IdentifyEffectVariant,
	}
}

// Name of the command (needed to fulfill interface)
func (TriggerEffect) Name() string { return "Trigger Effect" }

// ID of the command (needed to fulfill interface)
func (TriggerEffect) ID() CommandID { return TriggerEffectCommand }

// Cluster ID of the command (needed to fulfill interface)
func (TriggerEffect) Cluster() zcl.ClusterID { return IdentifyID }

// MnfCode returns the manufacturer code (if any) of the command
func (TriggerEffect) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of TriggerEffect
func (v TriggerEffect) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.IdentifyEffect.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.IdentifyEffectVariant.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the TriggerEffect struct
func (v *TriggerEffect) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.IdentifyEffect).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.IdentifyEffectVariant).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v TriggerEffect) String() string {
	return zcl.Sprintf(
		"TriggerEffect{"+zcl.StrJoin([]string{
			"IdentifyEffect(%v)",
			"IdentifyEffectVariant(%v)",
		}, " ")+"}",
		v.IdentifyEffect,
		v.IdentifyEffectVariant,
	)
}

// IdentifyQueryResponse Response of a identify query command.
type IdentifyQueryResponse struct {
	// IdentifyTimeout The time in seconds for which a device will stay in identify mode.
	IdentifyTimeout IdentifyTimeout
}

// IdentifyQueryResponseCommand is the Command ID of IdentifyQueryResponse
const IdentifyQueryResponseCommand CommandID = 0x0000

// Values returns all values of IdentifyQueryResponse
func (v *IdentifyQueryResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.IdentifyTimeout,
	}
}

// Name of the command (needed to fulfill interface)
func (IdentifyQueryResponse) Name() string { return "Identify Query Response" }

// ID of the command (needed to fulfill interface)
func (IdentifyQueryResponse) ID() CommandID { return IdentifyQueryResponseCommand }

// Cluster ID of the command (needed to fulfill interface)
func (IdentifyQueryResponse) Cluster() zcl.ClusterID { return IdentifyID }

// MnfCode returns the manufacturer code (if any) of the command
func (IdentifyQueryResponse) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of IdentifyQueryResponse
func (v IdentifyQueryResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.IdentifyTimeout.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the IdentifyQueryResponse struct
func (v *IdentifyQueryResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.IdentifyTimeout).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v IdentifyQueryResponse) String() string {
	return zcl.Sprintf(
		"IdentifyQueryResponse{"+zcl.StrJoin([]string{
			"IdentifyTimeout(%v)",
		}, " ")+"}",
		v.IdentifyTimeout,
	)
}
