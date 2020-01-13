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

type IdentifyHandler interface {
	HandleIdentify(frame Frame, cmd *Identify) error
}

// IdentifyCommand is the Command ID of Identify
const IdentifyCommand CommandID = 0x0000

// Values returns all values of Identify
func (v *Identify) Values() []zcl.Val {
	return []zcl.Val{
		&v.IdentifyTime,
	}
}

// Arguments returns all values of Identify
func (v *Identify) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "IdentifyTime", Argument: &v.IdentifyTime},
	}
}

// Name of the command
func (Identify) Name() string { return `Identify` }

// Description of the command
func (Identify) Description() string { return `Start or stop the device identifying itself.` }

// ID of the command
func (Identify) ID() CommandID { return IdentifyCommand }

// Required
func (Identify) Required() bool { return true }

// Cluster ID of the command
func (Identify) Cluster() zcl.ClusterID { return IdentifyID }

// Direction of the command
func (Identify) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (Identify) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (Identify) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

func (v *Identify) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h IdentifyHandler
	if h, found = handler.(IdentifyHandler); found {
		err = h.HandleIdentify(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of Identify
func (v Identify) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

type IdentifyQueryHandler interface {
	HandleIdentifyQuery(frame Frame, cmd *IdentifyQuery) (*IdentifyQueryResponse, error)
}

// IdentifyQueryCommand is the Command ID of IdentifyQuery
const IdentifyQueryCommand CommandID = 0x0001

// Values returns all values of IdentifyQuery
func (v *IdentifyQuery) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of IdentifyQuery
func (v *IdentifyQuery) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (IdentifyQuery) Name() string { return `Identify Query` }

// Description of the command
func (IdentifyQuery) Description() string {
	return `Allows the sending device to request the target or targets to respond if they are currently identifying themselves.`
}

// ID of the command
func (IdentifyQuery) ID() CommandID { return IdentifyQueryCommand }

// Required
func (IdentifyQuery) Required() bool { return true }

// Cluster ID of the command
func (IdentifyQuery) Cluster() zcl.ClusterID { return IdentifyID }

// Direction of the command
func (IdentifyQuery) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (IdentifyQuery) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (IdentifyQuery) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

func (v *IdentifyQuery) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h IdentifyQueryHandler
	if h, found = handler.(IdentifyQueryHandler); found {
		rsp, err = h.HandleIdentifyQuery(frame, v)
	}
	return
}

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

type TriggerEffectHandler interface {
	HandleTriggerEffect(frame Frame, cmd *TriggerEffect) error
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

// Arguments returns all values of TriggerEffect
func (v *TriggerEffect) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "IdentifyEffect", Argument: &v.IdentifyEffect},
		{Name: "IdentifyEffectVariant", Argument: &v.IdentifyEffectVariant},
	}
}

// Name of the command
func (TriggerEffect) Name() string { return `Trigger Effect` }

// Description of the command
func (TriggerEffect) Description() string {
	return `The trigger effect command allows the support of feedback to the user, such as a certain light effect.`
}

// ID of the command
func (TriggerEffect) ID() CommandID { return TriggerEffectCommand }

// Required
func (TriggerEffect) Required() bool { return true }

// Cluster ID of the command
func (TriggerEffect) Cluster() zcl.ClusterID { return IdentifyID }

// Direction of the command
func (TriggerEffect) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (TriggerEffect) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (TriggerEffect) MarshalJSON() ([]byte, error) { return []byte("64"), nil }

func (v *TriggerEffect) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h TriggerEffectHandler
	if h, found = handler.(TriggerEffectHandler); found {
		err = h.HandleTriggerEffect(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of TriggerEffect
func (v TriggerEffect) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

type IdentifyQueryResponseHandler interface {
	HandleIdentifyQueryResponse(frame Frame, cmd *IdentifyQueryResponse) error
}

// IdentifyQueryResponseCommand is the Command ID of IdentifyQueryResponse
const IdentifyQueryResponseCommand CommandID = 0x0000

// Values returns all values of IdentifyQueryResponse
func (v *IdentifyQueryResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.IdentifyTimeout,
	}
}

// Arguments returns all values of IdentifyQueryResponse
func (v *IdentifyQueryResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "IdentifyTimeout", Argument: &v.IdentifyTimeout},
	}
}

// Name of the command
func (IdentifyQueryResponse) Name() string { return `Identify Query Response` }

// Description of the command
func (IdentifyQueryResponse) Description() string { return `Response of a identify query command.` }

// ID of the command
func (IdentifyQueryResponse) ID() CommandID { return IdentifyQueryResponseCommand }

// Required
func (IdentifyQueryResponse) Required() bool { return true }

// Cluster ID of the command
func (IdentifyQueryResponse) Cluster() zcl.ClusterID { return IdentifyID }

// Direction of the command
func (IdentifyQueryResponse) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (IdentifyQueryResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (IdentifyQueryResponse) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

func (v *IdentifyQueryResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h IdentifyQueryResponseHandler
	if h, found = handler.(IdentifyQueryResponseHandler); found {
		err = h.HandleIdentifyQueryResponse(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of IdentifyQueryResponse
func (v IdentifyQueryResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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
