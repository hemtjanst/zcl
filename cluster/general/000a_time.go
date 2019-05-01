// This cluster provides a basic interface to a real-time clock.
package general

import (
	"neotor.se/zcl"
)

// Time
// This cluster provides a basic interface to a real-time clock.

func NewTimeServer(profile zcl.ProfileID) *TimeServer { return &TimeServer{p: profile} }
func NewTimeClient(profile zcl.ProfileID) *TimeClient { return &TimeClient{p: profile} }

const TimeCluster zcl.ClusterID = 10

type TimeServer struct {
	p zcl.ProfileID

	Time           *Time
	TimeStatus     *TimeStatus
	TimeZone       *TimeZone
	DstStart       *DstStart
	DstEnd         *DstEnd
	DstShift       *DstShift
	StandardTime   *StandardTime
	LocalTime      *LocalTime
	LastSetTime    *LastSetTime
	ValidUntilTime *ValidUntilTime
}

type TimeClient struct {
	p zcl.ProfileID
}

/*
var TimeServer = map[zcl.CommandID]func() zcl.Command{
}

var TimeClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type Time zcl.Zutc

func (a Time) ID() zcl.AttrID         { return 0 }
func (a Time) Cluster() zcl.ClusterID { return TimeCluster }
func (a *Time) Value() *Time          { return a }
func (a Time) MarshalZcl() ([]byte, error) {
	return zcl.Zutc(a).MarshalZcl()
}
func (a *Time) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*a = Time(*nt)
	return br, err
}

func (a Time) Readable() bool   { return true }
func (a Time) Writable() bool   { return true }
func (a Time) Reportable() bool { return false }
func (a Time) SceneIndex() int  { return -1 }

func (a Time) String() string {

	return zcl.Sprintf("%s", zcl.Zutc(a))
}

type TimeStatus zcl.Zbmp8

func (a TimeStatus) ID() zcl.AttrID         { return 1 }
func (a TimeStatus) Cluster() zcl.ClusterID { return TimeCluster }
func (a *TimeStatus) Value() *TimeStatus    { return a }
func (a TimeStatus) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *TimeStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = TimeStatus(*nt)
	return br, err
}

func (a TimeStatus) Readable() bool   { return true }
func (a TimeStatus) Writable() bool   { return true }
func (a TimeStatus) Reportable() bool { return false }
func (a TimeStatus) SceneIndex() int  { return -1 }

func (a TimeStatus) String() string {

	var bstr []string
	if a.IsMasterClock() {
		bstr = append(bstr, "Master Clock")
	}
	if a.IsSynchronized() {
		bstr = append(bstr, "Synchronized")
	}
	if a.IsMasterForTimezoneAndDst() {
		bstr = append(bstr, "Master for Timezone and Dst")
	}
	if a.IsSuperseding() {
		bstr = append(bstr, "Superseding")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a TimeStatus) IsMasterClock() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *TimeStatus) SetMasterClock(b bool) {
	*a = TimeStatus(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a TimeStatus) IsSynchronized() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *TimeStatus) SetSynchronized(b bool) {
	*a = TimeStatus(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a TimeStatus) IsMasterForTimezoneAndDst() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *TimeStatus) SetMasterForTimezoneAndDst(b bool) {
	*a = TimeStatus(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a TimeStatus) IsSuperseding() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *TimeStatus) SetSuperseding(b bool) {
	*a = TimeStatus(zcl.BitmapSet([]byte(*a), 3, b))
}

type TimeZone zcl.Zs32

func (a TimeZone) ID() zcl.AttrID         { return 2 }
func (a TimeZone) Cluster() zcl.ClusterID { return TimeCluster }
func (a *TimeZone) Value() *TimeZone      { return a }
func (a TimeZone) MarshalZcl() ([]byte, error) {
	return zcl.Zs32(a).MarshalZcl()
}
func (a *TimeZone) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs32)
	br, err := nt.UnmarshalZcl(b)
	*a = TimeZone(*nt)
	return br, err
}

func (a TimeZone) Readable() bool   { return true }
func (a TimeZone) Writable() bool   { return true }
func (a TimeZone) Reportable() bool { return false }
func (a TimeZone) SceneIndex() int  { return -1 }

func (a TimeZone) String() string {

	return zcl.Sprintf("%s", zcl.Zs32(a))
}

type DstStart zcl.Zutc

func (a DstStart) ID() zcl.AttrID         { return 3 }
func (a DstStart) Cluster() zcl.ClusterID { return TimeCluster }
func (a *DstStart) Value() *DstStart      { return a }
func (a DstStart) MarshalZcl() ([]byte, error) {
	return zcl.Zutc(a).MarshalZcl()
}
func (a *DstStart) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*a = DstStart(*nt)
	return br, err
}

func (a DstStart) Readable() bool   { return true }
func (a DstStart) Writable() bool   { return true }
func (a DstStart) Reportable() bool { return false }
func (a DstStart) SceneIndex() int  { return -1 }

func (a DstStart) String() string {

	return zcl.Sprintf("%s", zcl.Zutc(a))
}

type DstEnd zcl.Zutc

func (a DstEnd) ID() zcl.AttrID         { return 4 }
func (a DstEnd) Cluster() zcl.ClusterID { return TimeCluster }
func (a *DstEnd) Value() *DstEnd        { return a }
func (a DstEnd) MarshalZcl() ([]byte, error) {
	return zcl.Zutc(a).MarshalZcl()
}
func (a *DstEnd) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*a = DstEnd(*nt)
	return br, err
}

func (a DstEnd) Readable() bool   { return true }
func (a DstEnd) Writable() bool   { return true }
func (a DstEnd) Reportable() bool { return false }
func (a DstEnd) SceneIndex() int  { return -1 }

func (a DstEnd) String() string {

	return zcl.Sprintf("%s", zcl.Zutc(a))
}

type DstShift zcl.Zs32

func (a DstShift) ID() zcl.AttrID         { return 5 }
func (a DstShift) Cluster() zcl.ClusterID { return TimeCluster }
func (a *DstShift) Value() *DstShift      { return a }
func (a DstShift) MarshalZcl() ([]byte, error) {
	return zcl.Zs32(a).MarshalZcl()
}
func (a *DstShift) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs32)
	br, err := nt.UnmarshalZcl(b)
	*a = DstShift(*nt)
	return br, err
}

func (a DstShift) Readable() bool   { return true }
func (a DstShift) Writable() bool   { return true }
func (a DstShift) Reportable() bool { return false }
func (a DstShift) SceneIndex() int  { return -1 }

func (a DstShift) String() string {

	return zcl.Sprintf("%s", zcl.Zs32(a))
}

type StandardTime zcl.Zu32

func (a StandardTime) ID() zcl.AttrID         { return 6 }
func (a StandardTime) Cluster() zcl.ClusterID { return TimeCluster }
func (a *StandardTime) Value() *StandardTime  { return a }
func (a StandardTime) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *StandardTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = StandardTime(*nt)
	return br, err
}

func (a StandardTime) Readable() bool   { return true }
func (a StandardTime) Writable() bool   { return false }
func (a StandardTime) Reportable() bool { return false }
func (a StandardTime) SceneIndex() int  { return -1 }

func (a StandardTime) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}

type LocalTime zcl.Zu32

func (a LocalTime) ID() zcl.AttrID         { return 7 }
func (a LocalTime) Cluster() zcl.ClusterID { return TimeCluster }
func (a *LocalTime) Value() *LocalTime     { return a }
func (a LocalTime) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *LocalTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = LocalTime(*nt)
	return br, err
}

func (a LocalTime) Readable() bool   { return true }
func (a LocalTime) Writable() bool   { return false }
func (a LocalTime) Reportable() bool { return false }
func (a LocalTime) SceneIndex() int  { return -1 }

func (a LocalTime) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}

type LastSetTime zcl.Zutc

func (a LastSetTime) ID() zcl.AttrID         { return 8 }
func (a LastSetTime) Cluster() zcl.ClusterID { return TimeCluster }
func (a *LastSetTime) Value() *LastSetTime   { return a }
func (a LastSetTime) MarshalZcl() ([]byte, error) {
	return zcl.Zutc(a).MarshalZcl()
}
func (a *LastSetTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*a = LastSetTime(*nt)
	return br, err
}

func (a LastSetTime) Readable() bool   { return true }
func (a LastSetTime) Writable() bool   { return false }
func (a LastSetTime) Reportable() bool { return false }
func (a LastSetTime) SceneIndex() int  { return -1 }

func (a LastSetTime) String() string {

	return zcl.Sprintf("%s", zcl.Zutc(a))
}

type ValidUntilTime zcl.Zutc

func (a ValidUntilTime) ID() zcl.AttrID          { return 9 }
func (a ValidUntilTime) Cluster() zcl.ClusterID  { return TimeCluster }
func (a *ValidUntilTime) Value() *ValidUntilTime { return a }
func (a ValidUntilTime) MarshalZcl() ([]byte, error) {
	return zcl.Zutc(a).MarshalZcl()
}
func (a *ValidUntilTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*a = ValidUntilTime(*nt)
	return br, err
}

func (a ValidUntilTime) Readable() bool   { return true }
func (a ValidUntilTime) Writable() bool   { return true }
func (a ValidUntilTime) Reportable() bool { return false }
func (a ValidUntilTime) SceneIndex() int  { return -1 }

func (a ValidUntilTime) String() string {

	return zcl.Sprintf("%s", zcl.Zutc(a))
}
