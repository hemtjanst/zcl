package general

import "hemtjan.st/zcl"

// PollControl
const PollControlID zcl.ClusterID = 32

var PollControlCluster = zcl.Cluster{
	Name: "Poll Control",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		CheckInCommand: func() zcl.Command { return new(CheckIn) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		CheckInIntervalAttr:     func() zcl.Attr { return new(CheckInInterval) },
		LongPollIntervalAttr:    func() zcl.Attr { return new(LongPollInterval) },
		ShortPollIntervalAttr:   func() zcl.Attr { return new(ShortPollInterval) },
		FastPollTimeoutAttr:     func() zcl.Attr { return new(FastPollTimeout) },
		CheckInIntervalMinAttr:  func() zcl.Attr { return new(CheckInIntervalMin) },
		LongPollIntervalMinAttr: func() zcl.Attr { return new(LongPollIntervalMin) },
		FastPollTimeoutMaxAttr:  func() zcl.Attr { return new(FastPollTimeoutMax) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

type CheckIn struct {
}

type CheckInHandler interface {
	HandleCheckIn(frame Frame, cmd *CheckIn) error
}

// CheckInCommand is the Command ID of CheckIn
const CheckInCommand CommandID = 0x0000

// Values returns all values of CheckIn
func (v *CheckIn) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of CheckIn
func (v *CheckIn) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (CheckIn) Name() string { return `Check-in` }

// Description of the command
func (CheckIn) Description() string { return `` }

// ID of the command
func (CheckIn) ID() CommandID { return CheckInCommand }

// Required
func (CheckIn) Required() bool { return true }

// Cluster ID of the command
func (CheckIn) Cluster() zcl.ClusterID { return PollControlID }

// Direction of the command
func (CheckIn) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (CheckIn) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (CheckIn) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

func (v *CheckIn) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h CheckInHandler
	if h, found = handler.(CheckInHandler); found {
		err = h.HandleCheckIn(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of CheckIn
func (v CheckIn) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the CheckIn struct
func (v *CheckIn) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v CheckIn) String() string {
	return zcl.Sprintf(
		"CheckIn{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}
