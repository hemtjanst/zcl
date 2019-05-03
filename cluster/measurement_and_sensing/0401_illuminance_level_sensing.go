package measurement_and_sensing

import (
	"neotor.se/zcl"
)

// IlluminanceLevelSensing
const IlluminanceLevelSensingID zcl.ClusterID = 1025

var IlluminanceLevelSensingCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		LevelStatusAttr:            func() zcl.Attr { return new(LevelStatus) },
		IlluminanceSensorTypeAttr:  func() zcl.Attr { return new(IlluminanceSensorType) },
		IlluminanceTargetLevelAttr: func() zcl.Attr { return new(IlluminanceTargetLevel) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

// LevelStatus is an autogenerated attribute in the IlluminanceLevelSensing cluster
type LevelStatus zcl.Zenum8

const LevelStatusAttr zcl.AttrID = 0

func (a LevelStatus) ID() zcl.AttrID         { return LevelStatusAttr }
func (a LevelStatus) Cluster() zcl.ClusterID { return IlluminanceLevelSensingID }
func (a *LevelStatus) Value() *LevelStatus   { return a }
func (a LevelStatus) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *LevelStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = LevelStatus(*nt)
	return br, err
}
func (LevelStatus) Name() string     { return "Level Status" }
func (LevelStatus) Readable() bool   { return true }
func (LevelStatus) Writable() bool   { return false }
func (LevelStatus) Reportable() bool { return true }
func (LevelStatus) SceneIndex() int  { return -1 }

func (a LevelStatus) String() string {
	switch a {
	case 0x00:
		return "Illuminance on target"
	case 0x01:
		return "Illuminance below target"
	case 0x02:
		return "Illuminance above target"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

// IsIlluminanceOnTarget checks if LevelStatus equals the value for Illuminance on target (0x00)
func (a LevelStatus) IsIlluminanceOnTarget() bool { return a == 0x00 }

// SetIlluminanceOnTarget sets LevelStatus to Illuminance on target (0x00)
func (a *LevelStatus) SetIlluminanceOnTarget() { *a = 0x00 }

// IsIlluminanceBelowTarget checks if LevelStatus equals the value for Illuminance below target (0x01)
func (a LevelStatus) IsIlluminanceBelowTarget() bool { return a == 0x01 }

// SetIlluminanceBelowTarget sets LevelStatus to Illuminance below target (0x01)
func (a *LevelStatus) SetIlluminanceBelowTarget() { *a = 0x01 }

// IsIlluminanceAboveTarget checks if LevelStatus equals the value for Illuminance above target (0x02)
func (a LevelStatus) IsIlluminanceAboveTarget() bool { return a == 0x02 }

// SetIlluminanceAboveTarget sets LevelStatus to Illuminance above target (0x02)
func (a *LevelStatus) SetIlluminanceAboveTarget() { *a = 0x02 }

// IlluminanceSensorType is an autogenerated attribute in the IlluminanceLevelSensing cluster
type IlluminanceSensorType zcl.Zenum8

const IlluminanceSensorTypeAttr zcl.AttrID = 1

func (a IlluminanceSensorType) ID() zcl.AttrID                 { return IlluminanceSensorTypeAttr }
func (a IlluminanceSensorType) Cluster() zcl.ClusterID         { return IlluminanceLevelSensingID }
func (a *IlluminanceSensorType) Value() *IlluminanceSensorType { return a }
func (a IlluminanceSensorType) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *IlluminanceSensorType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = IlluminanceSensorType(*nt)
	return br, err
}
func (IlluminanceSensorType) Name() string     { return "Illuminance Sensor Type" }
func (IlluminanceSensorType) Readable() bool   { return true }
func (IlluminanceSensorType) Writable() bool   { return false }
func (IlluminanceSensorType) Reportable() bool { return false }
func (IlluminanceSensorType) SceneIndex() int  { return -1 }

func (a IlluminanceSensorType) String() string {
	switch a {
	case 0x00:
		return "Photodiode"
	case 0x01:
		return "CMOS"
	case 0xFF:
		return "Unknown"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

// IsPhotodiode checks if IlluminanceSensorType equals the value for Photodiode (0x00)
func (a IlluminanceSensorType) IsPhotodiode() bool { return a == 0x00 }

// SetPhotodiode sets IlluminanceSensorType to Photodiode (0x00)
func (a *IlluminanceSensorType) SetPhotodiode() { *a = 0x00 }

// IsCmos checks if IlluminanceSensorType equals the value for CMOS (0x01)
func (a IlluminanceSensorType) IsCmos() bool { return a == 0x01 }

// SetCmos sets IlluminanceSensorType to CMOS (0x01)
func (a *IlluminanceSensorType) SetCmos() { *a = 0x01 }

// IsUnknown checks if IlluminanceSensorType equals the value for Unknown (0xFF)
func (a IlluminanceSensorType) IsUnknown() bool { return a == 0xFF }

// SetUnknown sets IlluminanceSensorType to Unknown (0xFF)
func (a *IlluminanceSensorType) SetUnknown() { *a = 0xFF }

// IlluminanceTargetLevel is an autogenerated attribute in the IlluminanceLevelSensing cluster
type IlluminanceTargetLevel zcl.Zu16

const IlluminanceTargetLevelAttr zcl.AttrID = 16

func (a IlluminanceTargetLevel) ID() zcl.AttrID                  { return IlluminanceTargetLevelAttr }
func (a IlluminanceTargetLevel) Cluster() zcl.ClusterID          { return IlluminanceLevelSensingID }
func (a *IlluminanceTargetLevel) Value() *IlluminanceTargetLevel { return a }
func (a IlluminanceTargetLevel) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *IlluminanceTargetLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = IlluminanceTargetLevel(*nt)
	return br, err
}
func (IlluminanceTargetLevel) Name() string     { return "Illuminance Target Level" }
func (IlluminanceTargetLevel) Readable() bool   { return true }
func (IlluminanceTargetLevel) Writable() bool   { return true }
func (IlluminanceTargetLevel) Reportable() bool { return false }
func (IlluminanceTargetLevel) SceneIndex() int  { return -1 }

func (a IlluminanceTargetLevel) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}
