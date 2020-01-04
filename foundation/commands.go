package foundation

import (
	"hemtjan.st/zcl"
)

type DiscoverCommandsReceived struct {
	StartIndex zcl.Zu8
	MaxEntries zcl.Zu8
}

func (v DiscoverCommandsReceived) MarshalZcl() ([]byte, error) {
	return []byte{
		uint8(v.StartIndex),
		uint8(v.MaxEntries),
	}, nil
}

func (v *DiscoverCommandsReceived) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	if b, err = (&v.StartIndex).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	if b, err = (&v.MaxEntries).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	return b, nil
}

func (v DiscoverCommandsReceived) Values() []zcl.Val {
	return []zcl.Val{&v.StartIndex, &v.MaxEntries}
}

func (DiscoverCommandsReceived) ID() zcl.CommandID {
	return DiscoverCommandsReceivedCommand
}

func (DiscoverCommandsReceived) Name() string {
	return "DiscoverCommandsReceived"
}

func (v DiscoverCommandsReceived) String() string {
	return zcl.Sprintf("DiscoverCommandsReceived{Start[0x%02X] Max[%d]}", v.StartIndex, v.MaxEntries)
}

type DiscoverCommandsGenerated struct {
	StartIndex zcl.Zu8
	MaxEntries zcl.Zu8
}

func (v DiscoverCommandsGenerated) MarshalZcl() ([]byte, error) {
	return []byte{
		uint8(v.StartIndex),
		uint8(v.MaxEntries),
	}, nil
}

func (v *DiscoverCommandsGenerated) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	if b, err = (&v.StartIndex).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	if b, err = (&v.MaxEntries).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	return b, nil
}

func (v DiscoverCommandsGenerated) Values() []zcl.Val {
	return []zcl.Val{&v.StartIndex, &v.MaxEntries}
}

func (DiscoverCommandsGenerated) ID() zcl.CommandID {
	return DiscoverCommandsGeneratedCommand
}

func (DiscoverCommandsGenerated) Name() string {
	return "DiscoverCommandsGenerated"
}

func (v DiscoverCommandsGenerated) String() string {
	return zcl.Sprintf("DiscoverCommandsGenerated{Start[0x%02X] Max[%d]}", v.StartIndex, v.MaxEntries)
}

type DiscoverCommandsReceivedResponse struct {
	DiscoveryComplete zcl.Zbool
	Commands          []zcl.CommandID
}

func (v DiscoverCommandsReceivedResponse) MarshalZcl() ([]byte, error) {
	var tmp []byte
	var data []byte
	var err error
	if tmp, err = v.DiscoveryComplete.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)
	for _, c := range v.Commands {
		if tmp, err = c.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (v *DiscoverCommandsReceivedResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	if b, err = (&v.DiscoveryComplete).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	v.Commands = []zcl.CommandID{}
	for len(b) > 0 {
		c := new(zcl.CommandID)
		if b, err = c.UnmarshalZcl(b); err != nil {
			return nil, err
		}
		v.Commands = append(v.Commands, *c)
	}
	return b, nil
}

func (v DiscoverCommandsReceivedResponse) Values() (r []zcl.Val) {
	r = []zcl.Val{&v.DiscoveryComplete}
	for _, c := range v.Commands {
		r = append(r, &c)
	}
	return
}

func (DiscoverCommandsReceivedResponse) ID() zcl.CommandID {
	return DiscoverCommandsReceivedResponseCommand
}

func (DiscoverCommandsReceivedResponse) Name() string {
	return "DiscoverCommandsReceivedResponse"
}

func (v DiscoverCommandsReceivedResponse) String() string {
	return zcl.Sprintf("DiscoverCommandsReceivedResponse{Complete(%v) Commands(%v)}", v.DiscoveryComplete, v.Commands)
}

type DiscoverCommandsGeneratedResponse struct {
	DiscoveryComplete zcl.Zbool
	Commands          []zcl.CommandID
}

func (v DiscoverCommandsGeneratedResponse) MarshalZcl() ([]byte, error) {
	var tmp []byte
	var data []byte
	var err error
	if tmp, err = v.DiscoveryComplete.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)
	for _, c := range v.Commands {
		if tmp, err = c.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (v *DiscoverCommandsGeneratedResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	if b, err = (&v.DiscoveryComplete).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	v.Commands = []zcl.CommandID{}
	for len(b) > 0 {
		c := new(zcl.CommandID)
		if b, err = c.UnmarshalZcl(b); err != nil {
			return nil, err
		}
		v.Commands = append(v.Commands, *c)
	}
	return b, nil
}

func (v DiscoverCommandsGeneratedResponse) Values() (r []zcl.Val) {
	r = []zcl.Val{&v.DiscoveryComplete}
	for _, c := range v.Commands {
		r = append(r, &c)
	}
	return
}

func (DiscoverCommandsGeneratedResponse) ID() zcl.CommandID {
	return DiscoverCommandsGeneratedResponseCommand
}

func (DiscoverCommandsGeneratedResponse) Name() string {
	return "DiscoverCommandsGeneratedResponse"
}

func (v DiscoverCommandsGeneratedResponse) String() string {
	return zcl.Sprintf("DiscoverCommandsGeneratedResponse{Complete(%v) Commands(%v)}", v.DiscoveryComplete, v.Commands)
}
