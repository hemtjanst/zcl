// Provides a mechanism for the management of an end device’s MAC Data Request rate.
package general

import (
	"neotor.se/zcl"
)

// PollControl
const PollControlID zcl.ClusterID = 32

var PollControlCluster = zcl.Cluster{
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

const CheckInCommand zcl.CommandID = 0

func (v *CheckIn) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v CheckIn) ID() zcl.CommandID {
	return CheckInCommand
}

func (v CheckIn) Cluster() zcl.ClusterID {
	return PollControlID
}

func (v CheckIn) MnfCode() []byte {
	return []byte{}
}

func (v CheckIn) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *CheckIn) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

const CheckInIntervalAttr zcl.AttrID = 0

type CheckInInterval zcl.Zu32

func (a CheckInInterval) ID() zcl.AttrID           { return CheckInIntervalAttr }
func (a CheckInInterval) Cluster() zcl.ClusterID   { return PollControlID }
func (a *CheckInInterval) Value() *CheckInInterval { return a }
func (a CheckInInterval) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *CheckInInterval) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = CheckInInterval(*nt)
	return br, err
}

func (a CheckInInterval) Readable() bool   { return true }
func (a CheckInInterval) Writable() bool   { return true }
func (a CheckInInterval) Reportable() bool { return false }
func (a CheckInInterval) SceneIndex() int  { return -1 }

func (a CheckInInterval) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}

const LongPollIntervalAttr zcl.AttrID = 1

type LongPollInterval zcl.Zu32

func (a LongPollInterval) ID() zcl.AttrID            { return LongPollIntervalAttr }
func (a LongPollInterval) Cluster() zcl.ClusterID    { return PollControlID }
func (a *LongPollInterval) Value() *LongPollInterval { return a }
func (a LongPollInterval) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *LongPollInterval) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = LongPollInterval(*nt)
	return br, err
}

func (a LongPollInterval) Readable() bool   { return true }
func (a LongPollInterval) Writable() bool   { return false }
func (a LongPollInterval) Reportable() bool { return false }
func (a LongPollInterval) SceneIndex() int  { return -1 }

func (a LongPollInterval) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}

const ShortPollIntervalAttr zcl.AttrID = 2

type ShortPollInterval zcl.Zu16

func (a ShortPollInterval) ID() zcl.AttrID             { return ShortPollIntervalAttr }
func (a ShortPollInterval) Cluster() zcl.ClusterID     { return PollControlID }
func (a *ShortPollInterval) Value() *ShortPollInterval { return a }
func (a ShortPollInterval) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ShortPollInterval) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ShortPollInterval(*nt)
	return br, err
}

func (a ShortPollInterval) Readable() bool   { return true }
func (a ShortPollInterval) Writable() bool   { return false }
func (a ShortPollInterval) Reportable() bool { return false }
func (a ShortPollInterval) SceneIndex() int  { return -1 }

func (a ShortPollInterval) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const FastPollTimeoutAttr zcl.AttrID = 3

type FastPollTimeout zcl.Zu16

func (a FastPollTimeout) ID() zcl.AttrID           { return FastPollTimeoutAttr }
func (a FastPollTimeout) Cluster() zcl.ClusterID   { return PollControlID }
func (a *FastPollTimeout) Value() *FastPollTimeout { return a }
func (a FastPollTimeout) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *FastPollTimeout) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = FastPollTimeout(*nt)
	return br, err
}

func (a FastPollTimeout) Readable() bool   { return true }
func (a FastPollTimeout) Writable() bool   { return true }
func (a FastPollTimeout) Reportable() bool { return false }
func (a FastPollTimeout) SceneIndex() int  { return -1 }

func (a FastPollTimeout) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const CheckInIntervalMinAttr zcl.AttrID = 4

type CheckInIntervalMin zcl.Zu32

func (a CheckInIntervalMin) ID() zcl.AttrID              { return CheckInIntervalMinAttr }
func (a CheckInIntervalMin) Cluster() zcl.ClusterID      { return PollControlID }
func (a *CheckInIntervalMin) Value() *CheckInIntervalMin { return a }
func (a CheckInIntervalMin) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *CheckInIntervalMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = CheckInIntervalMin(*nt)
	return br, err
}

func (a CheckInIntervalMin) Readable() bool   { return true }
func (a CheckInIntervalMin) Writable() bool   { return false }
func (a CheckInIntervalMin) Reportable() bool { return false }
func (a CheckInIntervalMin) SceneIndex() int  { return -1 }

func (a CheckInIntervalMin) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}

const LongPollIntervalMinAttr zcl.AttrID = 5

type LongPollIntervalMin zcl.Zu32

func (a LongPollIntervalMin) ID() zcl.AttrID               { return LongPollIntervalMinAttr }
func (a LongPollIntervalMin) Cluster() zcl.ClusterID       { return PollControlID }
func (a *LongPollIntervalMin) Value() *LongPollIntervalMin { return a }
func (a LongPollIntervalMin) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *LongPollIntervalMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = LongPollIntervalMin(*nt)
	return br, err
}

func (a LongPollIntervalMin) Readable() bool   { return true }
func (a LongPollIntervalMin) Writable() bool   { return false }
func (a LongPollIntervalMin) Reportable() bool { return false }
func (a LongPollIntervalMin) SceneIndex() int  { return -1 }

func (a LongPollIntervalMin) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}

const FastPollTimeoutMaxAttr zcl.AttrID = 6

type FastPollTimeoutMax zcl.Zu16

func (a FastPollTimeoutMax) ID() zcl.AttrID              { return FastPollTimeoutMaxAttr }
func (a FastPollTimeoutMax) Cluster() zcl.ClusterID      { return PollControlID }
func (a *FastPollTimeoutMax) Value() *FastPollTimeoutMax { return a }
func (a FastPollTimeoutMax) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *FastPollTimeoutMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = FastPollTimeoutMax(*nt)
	return br, err
}

func (a FastPollTimeoutMax) Readable() bool   { return true }
func (a FastPollTimeoutMax) Writable() bool   { return false }
func (a FastPollTimeoutMax) Reportable() bool { return false }
func (a FastPollTimeoutMax) SceneIndex() int  { return -1 }

func (a FastPollTimeoutMax) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}
