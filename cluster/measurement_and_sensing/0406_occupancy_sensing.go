package measurement_and_sensing

import (
	"neotor.se/zcl/cluster/zcl"
)

// OccupancySensing

func NewOccupancySensingServer(profile zcl.ProfileID) *OccupancySensingServer {
	return &OccupancySensingServer{p: profile}
}
func NewOccupancySensingClient(profile zcl.ProfileID) *OccupancySensingClient {
	return &OccupancySensingClient{p: profile}
}

const OccupancySensingCluster zcl.ClusterID = 1030

type OccupancySensingServer struct {
	p zcl.ProfileID

	Occupancy                               *Occupancy
	OccupancySensorType                     *OccupancySensorType
	PirOccupiedToUnoccupiedDelay            *PirOccupiedToUnoccupiedDelay
	PirUnoccupiedToOccupiedDelay            *PirUnoccupiedToOccupiedDelay
	PirUnoccupiedToOccupiedThreshold        *PirUnoccupiedToOccupiedThreshold
	UltrasonicOccupiedToUnoccupiedDelay     *UltrasonicOccupiedToUnoccupiedDelay
	UltrasonicUnoccupiedToOccupiedDelay     *UltrasonicUnoccupiedToOccupiedDelay
	UltrasonicUnoccupiedToOccupiedThreshold *UltrasonicUnoccupiedToOccupiedThreshold
	Sensitivity                             *Sensitivity
	SensitivityMax                          *SensitivityMax
}

type OccupancySensingClient struct {
	p zcl.ProfileID
}

/*
var OccupancySensingServer = map[zcl.CommandID]func() zcl.Command{
}

var OccupancySensingClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type Occupancy zcl.Zbmp8

func (a Occupancy) ID() zcl.AttrID         { return 0 }
func (a Occupancy) Cluster() zcl.ClusterID { return OccupancySensingCluster }
func (a *Occupancy) Value() *Occupancy     { return a }
func (a Occupancy) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *Occupancy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = Occupancy(*nt)
	return br, err
}

func (a Occupancy) Readable() bool   { return true }
func (a Occupancy) Writable() bool   { return false }
func (a Occupancy) Reportable() bool { return true }
func (a Occupancy) SceneIndex() int  { return -1 }

func (a Occupancy) String() string {

	return zcl.Sprintf("%s", zcl.Zbmp8(a))
}

type OccupancySensorType zcl.Zenum8

func (a OccupancySensorType) ID() zcl.AttrID               { return 1 }
func (a OccupancySensorType) Cluster() zcl.ClusterID       { return OccupancySensingCluster }
func (a *OccupancySensorType) Value() *OccupancySensorType { return a }
func (a OccupancySensorType) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *OccupancySensorType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = OccupancySensorType(*nt)
	return br, err
}

func (a OccupancySensorType) Readable() bool   { return true }
func (a OccupancySensorType) Writable() bool   { return false }
func (a OccupancySensorType) Reportable() bool { return false }
func (a OccupancySensorType) SceneIndex() int  { return -1 }

func (a OccupancySensorType) String() string {
	switch a {
	case 0x00:
		return "PIR"
	case 0x01:
		return "Ultrasonic"
	case 0x02:
		return "PIR and ultrasonic"
	}

	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsPir checks if OccupancySensorType equals the value for PIR (0x00)
func (a OccupancySensorType) IsPir() bool { return a == 0x00 }

// SetPir sets OccupancySensorType to PIR (0x00)
func (a *OccupancySensorType) SetPir() { *a = 0x00 }

// IsUltrasonic checks if OccupancySensorType equals the value for Ultrasonic (0x01)
func (a OccupancySensorType) IsUltrasonic() bool { return a == 0x01 }

// SetUltrasonic sets OccupancySensorType to Ultrasonic (0x01)
func (a *OccupancySensorType) SetUltrasonic() { *a = 0x01 }

// IsPirAndUltrasonic checks if OccupancySensorType equals the value for PIR and ultrasonic (0x02)
func (a OccupancySensorType) IsPirAndUltrasonic() bool { return a == 0x02 }

// SetPirAndUltrasonic sets OccupancySensorType to PIR and ultrasonic (0x02)
func (a *OccupancySensorType) SetPirAndUltrasonic() { *a = 0x02 }

type PirOccupiedToUnoccupiedDelay zcl.Zu16

func (a PirOccupiedToUnoccupiedDelay) ID() zcl.AttrID                        { return 16 }
func (a PirOccupiedToUnoccupiedDelay) Cluster() zcl.ClusterID                { return OccupancySensingCluster }
func (a *PirOccupiedToUnoccupiedDelay) Value() *PirOccupiedToUnoccupiedDelay { return a }
func (a PirOccupiedToUnoccupiedDelay) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *PirOccupiedToUnoccupiedDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PirOccupiedToUnoccupiedDelay(*nt)
	return br, err
}

func (a PirOccupiedToUnoccupiedDelay) Readable() bool   { return true }
func (a PirOccupiedToUnoccupiedDelay) Writable() bool   { return true }
func (a PirOccupiedToUnoccupiedDelay) Reportable() bool { return false }
func (a PirOccupiedToUnoccupiedDelay) SceneIndex() int  { return -1 }

func (a PirOccupiedToUnoccupiedDelay) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type PirUnoccupiedToOccupiedDelay zcl.Zu16

func (a PirUnoccupiedToOccupiedDelay) ID() zcl.AttrID                        { return 17 }
func (a PirUnoccupiedToOccupiedDelay) Cluster() zcl.ClusterID                { return OccupancySensingCluster }
func (a *PirUnoccupiedToOccupiedDelay) Value() *PirUnoccupiedToOccupiedDelay { return a }
func (a PirUnoccupiedToOccupiedDelay) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *PirUnoccupiedToOccupiedDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PirUnoccupiedToOccupiedDelay(*nt)
	return br, err
}

func (a PirUnoccupiedToOccupiedDelay) Readable() bool   { return true }
func (a PirUnoccupiedToOccupiedDelay) Writable() bool   { return true }
func (a PirUnoccupiedToOccupiedDelay) Reportable() bool { return false }
func (a PirUnoccupiedToOccupiedDelay) SceneIndex() int  { return -1 }

func (a PirUnoccupiedToOccupiedDelay) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type PirUnoccupiedToOccupiedThreshold zcl.Zu8

func (a PirUnoccupiedToOccupiedThreshold) ID() zcl.AttrID                            { return 18 }
func (a PirUnoccupiedToOccupiedThreshold) Cluster() zcl.ClusterID                    { return OccupancySensingCluster }
func (a *PirUnoccupiedToOccupiedThreshold) Value() *PirUnoccupiedToOccupiedThreshold { return a }
func (a PirUnoccupiedToOccupiedThreshold) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *PirUnoccupiedToOccupiedThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = PirUnoccupiedToOccupiedThreshold(*nt)
	return br, err
}

func (a PirUnoccupiedToOccupiedThreshold) Readable() bool   { return true }
func (a PirUnoccupiedToOccupiedThreshold) Writable() bool   { return true }
func (a PirUnoccupiedToOccupiedThreshold) Reportable() bool { return false }
func (a PirUnoccupiedToOccupiedThreshold) SceneIndex() int  { return -1 }

func (a PirUnoccupiedToOccupiedThreshold) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type UltrasonicOccupiedToUnoccupiedDelay zcl.Zu16

func (a UltrasonicOccupiedToUnoccupiedDelay) ID() zcl.AttrID                               { return 32 }
func (a UltrasonicOccupiedToUnoccupiedDelay) Cluster() zcl.ClusterID                       { return OccupancySensingCluster }
func (a *UltrasonicOccupiedToUnoccupiedDelay) Value() *UltrasonicOccupiedToUnoccupiedDelay { return a }
func (a UltrasonicOccupiedToUnoccupiedDelay) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *UltrasonicOccupiedToUnoccupiedDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = UltrasonicOccupiedToUnoccupiedDelay(*nt)
	return br, err
}

func (a UltrasonicOccupiedToUnoccupiedDelay) Readable() bool   { return true }
func (a UltrasonicOccupiedToUnoccupiedDelay) Writable() bool   { return true }
func (a UltrasonicOccupiedToUnoccupiedDelay) Reportable() bool { return false }
func (a UltrasonicOccupiedToUnoccupiedDelay) SceneIndex() int  { return -1 }

func (a UltrasonicOccupiedToUnoccupiedDelay) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type UltrasonicUnoccupiedToOccupiedDelay zcl.Zu16

func (a UltrasonicUnoccupiedToOccupiedDelay) ID() zcl.AttrID                               { return 33 }
func (a UltrasonicUnoccupiedToOccupiedDelay) Cluster() zcl.ClusterID                       { return OccupancySensingCluster }
func (a *UltrasonicUnoccupiedToOccupiedDelay) Value() *UltrasonicUnoccupiedToOccupiedDelay { return a }
func (a UltrasonicUnoccupiedToOccupiedDelay) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *UltrasonicUnoccupiedToOccupiedDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = UltrasonicUnoccupiedToOccupiedDelay(*nt)
	return br, err
}

func (a UltrasonicUnoccupiedToOccupiedDelay) Readable() bool   { return true }
func (a UltrasonicUnoccupiedToOccupiedDelay) Writable() bool   { return true }
func (a UltrasonicUnoccupiedToOccupiedDelay) Reportable() bool { return false }
func (a UltrasonicUnoccupiedToOccupiedDelay) SceneIndex() int  { return -1 }

func (a UltrasonicUnoccupiedToOccupiedDelay) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type UltrasonicUnoccupiedToOccupiedThreshold zcl.Zu8

func (a UltrasonicUnoccupiedToOccupiedThreshold) ID() zcl.AttrID { return 34 }
func (a UltrasonicUnoccupiedToOccupiedThreshold) Cluster() zcl.ClusterID {
	return OccupancySensingCluster
}
func (a *UltrasonicUnoccupiedToOccupiedThreshold) Value() *UltrasonicUnoccupiedToOccupiedThreshold {
	return a
}
func (a UltrasonicUnoccupiedToOccupiedThreshold) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *UltrasonicUnoccupiedToOccupiedThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = UltrasonicUnoccupiedToOccupiedThreshold(*nt)
	return br, err
}

func (a UltrasonicUnoccupiedToOccupiedThreshold) Readable() bool   { return true }
func (a UltrasonicUnoccupiedToOccupiedThreshold) Writable() bool   { return true }
func (a UltrasonicUnoccupiedToOccupiedThreshold) Reportable() bool { return false }
func (a UltrasonicUnoccupiedToOccupiedThreshold) SceneIndex() int  { return -1 }

func (a UltrasonicUnoccupiedToOccupiedThreshold) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type Sensitivity zcl.Zu8

func (a Sensitivity) ID() zcl.AttrID         { return 48 }
func (a Sensitivity) Cluster() zcl.ClusterID { return OccupancySensingCluster }
func (a *Sensitivity) Value() *Sensitivity   { return a }
func (a Sensitivity) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *Sensitivity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Sensitivity(*nt)
	return br, err
}

func (a Sensitivity) Readable() bool   { return true }
func (a Sensitivity) Writable() bool   { return true }
func (a Sensitivity) Reportable() bool { return false }
func (a Sensitivity) SceneIndex() int  { return -1 }

func (a Sensitivity) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type SensitivityMax zcl.Zu8

func (a SensitivityMax) ID() zcl.AttrID          { return 49 }
func (a SensitivityMax) Cluster() zcl.ClusterID  { return OccupancySensingCluster }
func (a *SensitivityMax) Value() *SensitivityMax { return a }
func (a SensitivityMax) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *SensitivityMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = SensitivityMax(*nt)
	return br, err
}

func (a SensitivityMax) Readable() bool   { return true }
func (a SensitivityMax) Writable() bool   { return true }
func (a SensitivityMax) Reportable() bool { return false }
func (a SensitivityMax) SceneIndex() int  { return -1 }

func (a SensitivityMax) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}
