// Attributes and commands to configure a ballast.
package lighting

import (
	"neotor.se/zcl"
)

// BallastConfiguration
// Attributes and commands to configure a ballast.

func NewBallastConfigurationServer(profile zcl.ProfileID) *BallastConfigurationServer {
	return &BallastConfigurationServer{p: profile}
}
func NewBallastConfigurationClient(profile zcl.ProfileID) *BallastConfigurationClient {
	return &BallastConfigurationClient{p: profile}
}

const BallastConfigurationCluster zcl.ClusterID = 769

type BallastConfigurationServer struct {
	p zcl.ProfileID

	PhysicalMinLevel        *PhysicalMinLevel
	PhysicalMaxLevel        *PhysicalMaxLevel
	BallastStatus           *BallastStatus
	MinLevel                *MinLevel
	MaxLevel                *MaxLevel
	PowerOnLevel            *PowerOnLevel
	PowerOnFadeTime         *PowerOnFadeTime
	IntrinsicBallastFactor  *IntrinsicBallastFactor
	BallastFactorAdjustment *BallastFactorAdjustment
	LampQuantity            *LampQuantity
	LampType                *LampType
	LampManufacturer        *LampManufacturer
	LampRatedHours          *LampRatedHours
	LampBurnHours           *LampBurnHours
	LampAlarmMode           *LampAlarmMode
	LampBurnHoursTripPoint  *LampBurnHoursTripPoint
}

type BallastConfigurationClient struct {
	p zcl.ProfileID
}

/*
var BallastConfigurationServer = map[zcl.CommandID]func() zcl.Command{
}

var BallastConfigurationClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type PhysicalMinLevel zcl.Zu8

func (a PhysicalMinLevel) ID() zcl.AttrID            { return 0 }
func (a PhysicalMinLevel) Cluster() zcl.ClusterID    { return BallastConfigurationCluster }
func (a *PhysicalMinLevel) Value() *PhysicalMinLevel { return a }
func (a PhysicalMinLevel) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *PhysicalMinLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = PhysicalMinLevel(*nt)
	return br, err
}

func (a PhysicalMinLevel) Readable() bool   { return true }
func (a PhysicalMinLevel) Writable() bool   { return false }
func (a PhysicalMinLevel) Reportable() bool { return false }
func (a PhysicalMinLevel) SceneIndex() int  { return -1 }

func (a PhysicalMinLevel) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type PhysicalMaxLevel zcl.Zu8

func (a PhysicalMaxLevel) ID() zcl.AttrID            { return 1 }
func (a PhysicalMaxLevel) Cluster() zcl.ClusterID    { return BallastConfigurationCluster }
func (a *PhysicalMaxLevel) Value() *PhysicalMaxLevel { return a }
func (a PhysicalMaxLevel) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *PhysicalMaxLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = PhysicalMaxLevel(*nt)
	return br, err
}

func (a PhysicalMaxLevel) Readable() bool   { return true }
func (a PhysicalMaxLevel) Writable() bool   { return false }
func (a PhysicalMaxLevel) Reportable() bool { return false }
func (a PhysicalMaxLevel) SceneIndex() int  { return -1 }

func (a PhysicalMaxLevel) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type BallastStatus zcl.Zbmp8

func (a BallastStatus) ID() zcl.AttrID         { return 2 }
func (a BallastStatus) Cluster() zcl.ClusterID { return BallastConfigurationCluster }
func (a *BallastStatus) Value() *BallastStatus { return a }
func (a BallastStatus) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *BallastStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = BallastStatus(*nt)
	return br, err
}

func (a BallastStatus) Readable() bool   { return true }
func (a BallastStatus) Writable() bool   { return false }
func (a BallastStatus) Reportable() bool { return false }
func (a BallastStatus) SceneIndex() int  { return -1 }

func (a BallastStatus) String() string {

	var bstr []string
	if a.IsNonOperational() {
		bstr = append(bstr, "Non-operational")
	}
	if a.IsLampNotInSocket() {
		bstr = append(bstr, "Lamp not in socket")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a BallastStatus) IsNonOperational() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *BallastStatus) SetNonOperational(b bool) {
	*a = BallastStatus(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a BallastStatus) IsLampNotInSocket() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *BallastStatus) SetLampNotInSocket(b bool) {
	*a = BallastStatus(zcl.BitmapSet([]byte(*a), 1, b))
}

type MinLevel zcl.Zu8

func (a MinLevel) ID() zcl.AttrID         { return 16 }
func (a MinLevel) Cluster() zcl.ClusterID { return BallastConfigurationCluster }
func (a *MinLevel) Value() *MinLevel      { return a }
func (a MinLevel) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *MinLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = MinLevel(*nt)
	return br, err
}

func (a MinLevel) Readable() bool   { return true }
func (a MinLevel) Writable() bool   { return true }
func (a MinLevel) Reportable() bool { return false }
func (a MinLevel) SceneIndex() int  { return -1 }

func (a MinLevel) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type MaxLevel zcl.Zu8

func (a MaxLevel) ID() zcl.AttrID         { return 17 }
func (a MaxLevel) Cluster() zcl.ClusterID { return BallastConfigurationCluster }
func (a *MaxLevel) Value() *MaxLevel      { return a }
func (a MaxLevel) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *MaxLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxLevel(*nt)
	return br, err
}

func (a MaxLevel) Readable() bool   { return true }
func (a MaxLevel) Writable() bool   { return true }
func (a MaxLevel) Reportable() bool { return false }
func (a MaxLevel) SceneIndex() int  { return -1 }

func (a MaxLevel) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type PowerOnLevel zcl.Zu8

func (a PowerOnLevel) ID() zcl.AttrID         { return 18 }
func (a PowerOnLevel) Cluster() zcl.ClusterID { return BallastConfigurationCluster }
func (a *PowerOnLevel) Value() *PowerOnLevel  { return a }
func (a PowerOnLevel) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *PowerOnLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerOnLevel(*nt)
	return br, err
}

func (a PowerOnLevel) Readable() bool   { return true }
func (a PowerOnLevel) Writable() bool   { return true }
func (a PowerOnLevel) Reportable() bool { return false }
func (a PowerOnLevel) SceneIndex() int  { return -1 }

func (a PowerOnLevel) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type PowerOnFadeTime zcl.Zu16

func (a PowerOnFadeTime) ID() zcl.AttrID           { return 19 }
func (a PowerOnFadeTime) Cluster() zcl.ClusterID   { return BallastConfigurationCluster }
func (a *PowerOnFadeTime) Value() *PowerOnFadeTime { return a }
func (a PowerOnFadeTime) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *PowerOnFadeTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerOnFadeTime(*nt)
	return br, err
}

func (a PowerOnFadeTime) Readable() bool   { return true }
func (a PowerOnFadeTime) Writable() bool   { return true }
func (a PowerOnFadeTime) Reportable() bool { return false }
func (a PowerOnFadeTime) SceneIndex() int  { return -1 }

func (a PowerOnFadeTime) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type IntrinsicBallastFactor zcl.Zu8

func (a IntrinsicBallastFactor) ID() zcl.AttrID                  { return 20 }
func (a IntrinsicBallastFactor) Cluster() zcl.ClusterID          { return BallastConfigurationCluster }
func (a *IntrinsicBallastFactor) Value() *IntrinsicBallastFactor { return a }
func (a IntrinsicBallastFactor) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *IntrinsicBallastFactor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = IntrinsicBallastFactor(*nt)
	return br, err
}

func (a IntrinsicBallastFactor) Readable() bool   { return true }
func (a IntrinsicBallastFactor) Writable() bool   { return true }
func (a IntrinsicBallastFactor) Reportable() bool { return false }
func (a IntrinsicBallastFactor) SceneIndex() int  { return -1 }

func (a IntrinsicBallastFactor) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type BallastFactorAdjustment zcl.Zu8

func (a BallastFactorAdjustment) ID() zcl.AttrID                   { return 21 }
func (a BallastFactorAdjustment) Cluster() zcl.ClusterID           { return BallastConfigurationCluster }
func (a *BallastFactorAdjustment) Value() *BallastFactorAdjustment { return a }
func (a BallastFactorAdjustment) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *BallastFactorAdjustment) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BallastFactorAdjustment(*nt)
	return br, err
}

func (a BallastFactorAdjustment) Readable() bool   { return true }
func (a BallastFactorAdjustment) Writable() bool   { return true }
func (a BallastFactorAdjustment) Reportable() bool { return false }
func (a BallastFactorAdjustment) SceneIndex() int  { return -1 }

func (a BallastFactorAdjustment) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type LampQuantity zcl.Zu8

func (a LampQuantity) ID() zcl.AttrID         { return 32 }
func (a LampQuantity) Cluster() zcl.ClusterID { return BallastConfigurationCluster }
func (a *LampQuantity) Value() *LampQuantity  { return a }
func (a LampQuantity) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *LampQuantity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = LampQuantity(*nt)
	return br, err
}

func (a LampQuantity) Readable() bool   { return true }
func (a LampQuantity) Writable() bool   { return false }
func (a LampQuantity) Reportable() bool { return false }
func (a LampQuantity) SceneIndex() int  { return -1 }

func (a LampQuantity) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type LampType zcl.Zcstring

func (a LampType) ID() zcl.AttrID         { return 48 }
func (a LampType) Cluster() zcl.ClusterID { return BallastConfigurationCluster }
func (a *LampType) Value() *LampType      { return a }
func (a LampType) MarshalZcl() ([]byte, error) {
	return zcl.Zcstring(a).MarshalZcl()
}
func (a *LampType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = LampType(*nt)
	return br, err
}

func (a LampType) Readable() bool   { return true }
func (a LampType) Writable() bool   { return true }
func (a LampType) Reportable() bool { return false }
func (a LampType) SceneIndex() int  { return -1 }

func (a LampType) String() string {

	return zcl.Sprintf("%s", zcl.Zcstring(a))
}

type LampManufacturer zcl.Zcstring

func (a LampManufacturer) ID() zcl.AttrID            { return 49 }
func (a LampManufacturer) Cluster() zcl.ClusterID    { return BallastConfigurationCluster }
func (a *LampManufacturer) Value() *LampManufacturer { return a }
func (a LampManufacturer) MarshalZcl() ([]byte, error) {
	return zcl.Zcstring(a).MarshalZcl()
}
func (a *LampManufacturer) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = LampManufacturer(*nt)
	return br, err
}

func (a LampManufacturer) Readable() bool   { return true }
func (a LampManufacturer) Writable() bool   { return true }
func (a LampManufacturer) Reportable() bool { return false }
func (a LampManufacturer) SceneIndex() int  { return -1 }

func (a LampManufacturer) String() string {

	return zcl.Sprintf("%s", zcl.Zcstring(a))
}

type LampRatedHours zcl.Zu24

func (a LampRatedHours) ID() zcl.AttrID          { return 50 }
func (a LampRatedHours) Cluster() zcl.ClusterID  { return BallastConfigurationCluster }
func (a *LampRatedHours) Value() *LampRatedHours { return a }
func (a LampRatedHours) MarshalZcl() ([]byte, error) {
	return zcl.Zu24(a).MarshalZcl()
}
func (a *LampRatedHours) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*a = LampRatedHours(*nt)
	return br, err
}

func (a LampRatedHours) Readable() bool   { return true }
func (a LampRatedHours) Writable() bool   { return true }
func (a LampRatedHours) Reportable() bool { return false }
func (a LampRatedHours) SceneIndex() int  { return -1 }

func (a LampRatedHours) String() string {

	return zcl.Sprintf("%s", zcl.Zu24(a))
}

type LampBurnHours zcl.Zu24

func (a LampBurnHours) ID() zcl.AttrID         { return 51 }
func (a LampBurnHours) Cluster() zcl.ClusterID { return BallastConfigurationCluster }
func (a *LampBurnHours) Value() *LampBurnHours { return a }
func (a LampBurnHours) MarshalZcl() ([]byte, error) {
	return zcl.Zu24(a).MarshalZcl()
}
func (a *LampBurnHours) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*a = LampBurnHours(*nt)
	return br, err
}

func (a LampBurnHours) Readable() bool   { return true }
func (a LampBurnHours) Writable() bool   { return true }
func (a LampBurnHours) Reportable() bool { return false }
func (a LampBurnHours) SceneIndex() int  { return -1 }

func (a LampBurnHours) String() string {

	return zcl.Sprintf("%s", zcl.Zu24(a))
}

type LampAlarmMode zcl.Zbmp8

func (a LampAlarmMode) ID() zcl.AttrID         { return 52 }
func (a LampAlarmMode) Cluster() zcl.ClusterID { return BallastConfigurationCluster }
func (a *LampAlarmMode) Value() *LampAlarmMode { return a }
func (a LampAlarmMode) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *LampAlarmMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = LampAlarmMode(*nt)
	return br, err
}

func (a LampAlarmMode) Readable() bool   { return true }
func (a LampAlarmMode) Writable() bool   { return true }
func (a LampAlarmMode) Reportable() bool { return false }
func (a LampAlarmMode) SceneIndex() int  { return -1 }

func (a LampAlarmMode) String() string {

	var bstr []string
	if a.IsLampBurnHours() {
		bstr = append(bstr, "Lamp Burn Hours")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a LampAlarmMode) IsLampBurnHours() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *LampAlarmMode) SetLampBurnHours(b bool) {
	*a = LampAlarmMode(zcl.BitmapSet([]byte(*a), 0, b))
}

type LampBurnHoursTripPoint zcl.Zu24

func (a LampBurnHoursTripPoint) ID() zcl.AttrID                  { return 53 }
func (a LampBurnHoursTripPoint) Cluster() zcl.ClusterID          { return BallastConfigurationCluster }
func (a *LampBurnHoursTripPoint) Value() *LampBurnHoursTripPoint { return a }
func (a LampBurnHoursTripPoint) MarshalZcl() ([]byte, error) {
	return zcl.Zu24(a).MarshalZcl()
}
func (a *LampBurnHoursTripPoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*a = LampBurnHoursTripPoint(*nt)
	return br, err
}

func (a LampBurnHoursTripPoint) Readable() bool   { return true }
func (a LampBurnHoursTripPoint) Writable() bool   { return true }
func (a LampBurnHoursTripPoint) Reportable() bool { return false }
func (a LampBurnHoursTripPoint) SceneIndex() int  { return -1 }

func (a LampBurnHoursTripPoint) String() string {

	return zcl.Sprintf("%s", zcl.Zu24(a))
}
