// Attributes for determining basic information about a device, setting user device information
// such as description of location, and enabling a device.
//
package general

import (
	"neotor.se/zcl"
)

// Basic
const BasicID zcl.ClusterID = 0

var BasicCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		ResetToFactoryDefaultsCommand: func() zcl.Command { return new(ResetToFactoryDefaults) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		ClusterRevisionAttr:     func() zcl.Attr { return new(ClusterRevision) },
		ZclVersionAttr:          func() zcl.Attr { return new(ZclVersion) },
		ApplicationVersionAttr:  func() zcl.Attr { return new(ApplicationVersion) },
		StackVersionAttr:        func() zcl.Attr { return new(StackVersion) },
		HwVersionAttr:           func() zcl.Attr { return new(HwVersion) },
		ManufacturerNameAttr:    func() zcl.Attr { return new(ManufacturerName) },
		ModelIdentifierAttr:     func() zcl.Attr { return new(ModelIdentifier) },
		DateCodeAttr:            func() zcl.Attr { return new(DateCode) },
		PowerSourceAttr:         func() zcl.Attr { return new(PowerSource) },
		LocationDescriptionAttr: func() zcl.Attr { return new(LocationDescription) },
		PhysicalEnvironmentAttr: func() zcl.Attr { return new(PhysicalEnvironment) },
		DeviceEnabledAttr:       func() zcl.Attr { return new(DeviceEnabled) },
		AlarmMaskAttr:           func() zcl.Attr { return new(AlarmMask) },
		DisableLocalConfigAttr:  func() zcl.Attr { return new(DisableLocalConfig) },
		SwBuildIdAttr:           func() zcl.Attr { return new(SwBuildId) },
		Unknown1Attr:            func() zcl.Attr { return new(Unknown1) },
		Unknown2Attr:            func() zcl.Attr { return new(Unknown2) },
		ProductCodeAttr:         func() zcl.Attr { return new(ProductCode) },
		XiaomiSensitivityAttr:   func() zcl.Attr { return new(XiaomiSensitivity) },
		XiaomiDisconnect1Attr:   func() zcl.Attr { return new(XiaomiDisconnect1) },
		XiaomiDisconnect2Attr:   func() zcl.Attr { return new(XiaomiDisconnect2) },
		SensitivityAttr:         func() zcl.Attr { return new(Sensitivity) },
		ConfigurationAttr:       func() zcl.Attr { return new(Configuration) },
		UsertestAttr:            func() zcl.Attr { return new(Usertest) },
		LedIndicationAttr:       func() zcl.Attr { return new(LedIndication) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

type ResetToFactoryDefaults struct {
}

const ResetToFactoryDefaultsCommand zcl.CommandID = 0

func (v *ResetToFactoryDefaults) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v ResetToFactoryDefaults) ID() zcl.CommandID {
	return ResetToFactoryDefaultsCommand
}

func (v ResetToFactoryDefaults) Cluster() zcl.ClusterID {
	return BasicID
}

func (v ResetToFactoryDefaults) MnfCode() []byte {
	return []byte{}
}

func (v ResetToFactoryDefaults) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *ResetToFactoryDefaults) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

const ClusterRevisionAttr zcl.AttrID = 65533

type ClusterRevision zcl.Zu16

func (a ClusterRevision) ID() zcl.AttrID           { return ClusterRevisionAttr }
func (a ClusterRevision) Cluster() zcl.ClusterID   { return BasicID }
func (a *ClusterRevision) Value() *ClusterRevision { return a }
func (a ClusterRevision) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ClusterRevision) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ClusterRevision(*nt)
	return br, err
}

func (a ClusterRevision) Readable() bool   { return true }
func (a ClusterRevision) Writable() bool   { return true }
func (a ClusterRevision) Reportable() bool { return false }
func (a ClusterRevision) SceneIndex() int  { return -1 }

func (a ClusterRevision) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ZclVersionAttr zcl.AttrID = 0

type ZclVersion zcl.Zu8

func (a ZclVersion) ID() zcl.AttrID         { return ZclVersionAttr }
func (a ZclVersion) Cluster() zcl.ClusterID { return BasicID }
func (a *ZclVersion) Value() *ZclVersion    { return a }
func (a ZclVersion) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *ZclVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ZclVersion(*nt)
	return br, err
}

func (a ZclVersion) Readable() bool   { return true }
func (a ZclVersion) Writable() bool   { return false }
func (a ZclVersion) Reportable() bool { return false }
func (a ZclVersion) SceneIndex() int  { return -1 }

func (a ZclVersion) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const ApplicationVersionAttr zcl.AttrID = 1

type ApplicationVersion zcl.Zu8

func (a ApplicationVersion) ID() zcl.AttrID              { return ApplicationVersionAttr }
func (a ApplicationVersion) Cluster() zcl.ClusterID      { return BasicID }
func (a *ApplicationVersion) Value() *ApplicationVersion { return a }
func (a ApplicationVersion) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *ApplicationVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ApplicationVersion(*nt)
	return br, err
}

func (a ApplicationVersion) Readable() bool   { return true }
func (a ApplicationVersion) Writable() bool   { return false }
func (a ApplicationVersion) Reportable() bool { return false }
func (a ApplicationVersion) SceneIndex() int  { return -1 }

func (a ApplicationVersion) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const StackVersionAttr zcl.AttrID = 2

type StackVersion zcl.Zu8

func (a StackVersion) ID() zcl.AttrID         { return StackVersionAttr }
func (a StackVersion) Cluster() zcl.ClusterID { return BasicID }
func (a *StackVersion) Value() *StackVersion  { return a }
func (a StackVersion) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *StackVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = StackVersion(*nt)
	return br, err
}

func (a StackVersion) Readable() bool   { return true }
func (a StackVersion) Writable() bool   { return false }
func (a StackVersion) Reportable() bool { return false }
func (a StackVersion) SceneIndex() int  { return -1 }

func (a StackVersion) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const HwVersionAttr zcl.AttrID = 3

type HwVersion zcl.Zu8

func (a HwVersion) ID() zcl.AttrID         { return HwVersionAttr }
func (a HwVersion) Cluster() zcl.ClusterID { return BasicID }
func (a *HwVersion) Value() *HwVersion     { return a }
func (a HwVersion) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *HwVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = HwVersion(*nt)
	return br, err
}

func (a HwVersion) Readable() bool   { return true }
func (a HwVersion) Writable() bool   { return false }
func (a HwVersion) Reportable() bool { return false }
func (a HwVersion) SceneIndex() int  { return -1 }

func (a HwVersion) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const ManufacturerNameAttr zcl.AttrID = 4

type ManufacturerName zcl.Zcstring

func (a ManufacturerName) ID() zcl.AttrID            { return ManufacturerNameAttr }
func (a ManufacturerName) Cluster() zcl.ClusterID    { return BasicID }
func (a *ManufacturerName) Value() *ManufacturerName { return a }
func (a ManufacturerName) MarshalZcl() ([]byte, error) {
	return zcl.Zcstring(a).MarshalZcl()
}
func (a *ManufacturerName) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = ManufacturerName(*nt)
	return br, err
}

func (a ManufacturerName) Readable() bool   { return true }
func (a ManufacturerName) Writable() bool   { return false }
func (a ManufacturerName) Reportable() bool { return false }
func (a ManufacturerName) SceneIndex() int  { return -1 }

func (a ManufacturerName) String() string {

	return zcl.Sprintf("%s", zcl.Zcstring(a))
}

const ModelIdentifierAttr zcl.AttrID = 5

type ModelIdentifier zcl.Zcstring

func (a ModelIdentifier) ID() zcl.AttrID           { return ModelIdentifierAttr }
func (a ModelIdentifier) Cluster() zcl.ClusterID   { return BasicID }
func (a *ModelIdentifier) Value() *ModelIdentifier { return a }
func (a ModelIdentifier) MarshalZcl() ([]byte, error) {
	return zcl.Zcstring(a).MarshalZcl()
}
func (a *ModelIdentifier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = ModelIdentifier(*nt)
	return br, err
}

func (a ModelIdentifier) Readable() bool   { return true }
func (a ModelIdentifier) Writable() bool   { return false }
func (a ModelIdentifier) Reportable() bool { return false }
func (a ModelIdentifier) SceneIndex() int  { return -1 }

func (a ModelIdentifier) String() string {

	return zcl.Sprintf("%s", zcl.Zcstring(a))
}

const DateCodeAttr zcl.AttrID = 6

type DateCode zcl.Zcstring

func (a DateCode) ID() zcl.AttrID         { return DateCodeAttr }
func (a DateCode) Cluster() zcl.ClusterID { return BasicID }
func (a *DateCode) Value() *DateCode      { return a }
func (a DateCode) MarshalZcl() ([]byte, error) {
	return zcl.Zcstring(a).MarshalZcl()
}
func (a *DateCode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = DateCode(*nt)
	return br, err
}

func (a DateCode) Readable() bool   { return true }
func (a DateCode) Writable() bool   { return false }
func (a DateCode) Reportable() bool { return false }
func (a DateCode) SceneIndex() int  { return -1 }

func (a DateCode) String() string {

	return zcl.Sprintf("%s", zcl.Zcstring(a))
}

const PowerSourceAttr zcl.AttrID = 7

type PowerSource zcl.Zenum8

func (a PowerSource) ID() zcl.AttrID         { return PowerSourceAttr }
func (a PowerSource) Cluster() zcl.ClusterID { return BasicID }
func (a *PowerSource) Value() *PowerSource   { return a }
func (a PowerSource) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *PowerSource) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerSource(*nt)
	return br, err
}

func (a PowerSource) Readable() bool   { return true }
func (a PowerSource) Writable() bool   { return false }
func (a PowerSource) Reportable() bool { return false }
func (a PowerSource) SceneIndex() int  { return -1 }

func (a PowerSource) String() string {
	switch a {
	case 0x00:
		return "Unknown"
	case 0x01:
		return "Mains (single phase)"
	case 0x02:
		return "Mains (3 phase)"
	case 0x03:
		return "Battery"
	case 0x04:
		return "DC Source"
	case 0x05:
		return "Emergency mains constantly powered"
	case 0x06:
		return "Emergency mains and transfer switch"
	case 0x80:
		return "Unknown with battery backup"
	case 0x81:
		return "Mains (single phase) with battery backup"
	case 0x82:
		return "Mains (3 phase) with battery backup"
	case 0x83:
		return "Battery with battery backup"
	case 0x84:
		return "DC Source with battery backup"
	case 0x85:
		return "Emergency mains constantly powered with battery backup"
	case 0x86:
		return "Emergency mains and transfer switch with battery backup"
	}

	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsUnknown checks if PowerSource equals the value for Unknown (0x00)
func (a PowerSource) IsUnknown() bool { return a == 0x00 }

// SetUnknown sets PowerSource to Unknown (0x00)
func (a *PowerSource) SetUnknown() { *a = 0x00 }

// IsMainsSinglePhase checks if PowerSource equals the value for Mains (single phase) (0x01)
func (a PowerSource) IsMainsSinglePhase() bool { return a == 0x01 }

// SetMainsSinglePhase sets PowerSource to Mains (single phase) (0x01)
func (a *PowerSource) SetMainsSinglePhase() { *a = 0x01 }

// IsMains3Phase checks if PowerSource equals the value for Mains (3 phase) (0x02)
func (a PowerSource) IsMains3Phase() bool { return a == 0x02 }

// SetMains3Phase sets PowerSource to Mains (3 phase) (0x02)
func (a *PowerSource) SetMains3Phase() { *a = 0x02 }

// IsBattery checks if PowerSource equals the value for Battery (0x03)
func (a PowerSource) IsBattery() bool { return a == 0x03 }

// SetBattery sets PowerSource to Battery (0x03)
func (a *PowerSource) SetBattery() { *a = 0x03 }

// IsDcSource checks if PowerSource equals the value for DC Source (0x04)
func (a PowerSource) IsDcSource() bool { return a == 0x04 }

// SetDcSource sets PowerSource to DC Source (0x04)
func (a *PowerSource) SetDcSource() { *a = 0x04 }

// IsEmergencyMainsConstantlyPowered checks if PowerSource equals the value for Emergency mains constantly powered (0x05)
func (a PowerSource) IsEmergencyMainsConstantlyPowered() bool { return a == 0x05 }

// SetEmergencyMainsConstantlyPowered sets PowerSource to Emergency mains constantly powered (0x05)
func (a *PowerSource) SetEmergencyMainsConstantlyPowered() { *a = 0x05 }

// IsEmergencyMainsAndTransferSwitch checks if PowerSource equals the value for Emergency mains and transfer switch (0x06)
func (a PowerSource) IsEmergencyMainsAndTransferSwitch() bool { return a == 0x06 }

// SetEmergencyMainsAndTransferSwitch sets PowerSource to Emergency mains and transfer switch (0x06)
func (a *PowerSource) SetEmergencyMainsAndTransferSwitch() { *a = 0x06 }

// IsUnknownWithBatteryBackup checks if PowerSource equals the value for Unknown with battery backup (0x80)
func (a PowerSource) IsUnknownWithBatteryBackup() bool { return a == 0x80 }

// SetUnknownWithBatteryBackup sets PowerSource to Unknown with battery backup (0x80)
func (a *PowerSource) SetUnknownWithBatteryBackup() { *a = 0x80 }

// IsMainsSinglePhaseWithBatteryBackup checks if PowerSource equals the value for Mains (single phase) with battery backup (0x81)
func (a PowerSource) IsMainsSinglePhaseWithBatteryBackup() bool { return a == 0x81 }

// SetMainsSinglePhaseWithBatteryBackup sets PowerSource to Mains (single phase) with battery backup (0x81)
func (a *PowerSource) SetMainsSinglePhaseWithBatteryBackup() { *a = 0x81 }

// IsMains3PhaseWithBatteryBackup checks if PowerSource equals the value for Mains (3 phase) with battery backup (0x82)
func (a PowerSource) IsMains3PhaseWithBatteryBackup() bool { return a == 0x82 }

// SetMains3PhaseWithBatteryBackup sets PowerSource to Mains (3 phase) with battery backup (0x82)
func (a *PowerSource) SetMains3PhaseWithBatteryBackup() { *a = 0x82 }

// IsBatteryWithBatteryBackup checks if PowerSource equals the value for Battery with battery backup (0x83)
func (a PowerSource) IsBatteryWithBatteryBackup() bool { return a == 0x83 }

// SetBatteryWithBatteryBackup sets PowerSource to Battery with battery backup (0x83)
func (a *PowerSource) SetBatteryWithBatteryBackup() { *a = 0x83 }

// IsDcSourceWithBatteryBackup checks if PowerSource equals the value for DC Source with battery backup (0x84)
func (a PowerSource) IsDcSourceWithBatteryBackup() bool { return a == 0x84 }

// SetDcSourceWithBatteryBackup sets PowerSource to DC Source with battery backup (0x84)
func (a *PowerSource) SetDcSourceWithBatteryBackup() { *a = 0x84 }

// IsEmergencyMainsConstantlyPoweredWithBatteryBackup checks if PowerSource equals the value for Emergency mains constantly powered with battery backup (0x85)
func (a PowerSource) IsEmergencyMainsConstantlyPoweredWithBatteryBackup() bool { return a == 0x85 }

// SetEmergencyMainsConstantlyPoweredWithBatteryBackup sets PowerSource to Emergency mains constantly powered with battery backup (0x85)
func (a *PowerSource) SetEmergencyMainsConstantlyPoweredWithBatteryBackup() { *a = 0x85 }

// IsEmergencyMainsAndTransferSwitchWithBatteryBackup checks if PowerSource equals the value for Emergency mains and transfer switch with battery backup (0x86)
func (a PowerSource) IsEmergencyMainsAndTransferSwitchWithBatteryBackup() bool { return a == 0x86 }

// SetEmergencyMainsAndTransferSwitchWithBatteryBackup sets PowerSource to Emergency mains and transfer switch with battery backup (0x86)
func (a *PowerSource) SetEmergencyMainsAndTransferSwitchWithBatteryBackup() { *a = 0x86 }

const LocationDescriptionAttr zcl.AttrID = 16

type LocationDescription zcl.Zcstring

func (a LocationDescription) ID() zcl.AttrID               { return LocationDescriptionAttr }
func (a LocationDescription) Cluster() zcl.ClusterID       { return BasicID }
func (a *LocationDescription) Value() *LocationDescription { return a }
func (a LocationDescription) MarshalZcl() ([]byte, error) {
	return zcl.Zcstring(a).MarshalZcl()
}
func (a *LocationDescription) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = LocationDescription(*nt)
	return br, err
}

func (a LocationDescription) Readable() bool   { return true }
func (a LocationDescription) Writable() bool   { return true }
func (a LocationDescription) Reportable() bool { return false }
func (a LocationDescription) SceneIndex() int  { return -1 }

func (a LocationDescription) String() string {

	return zcl.Sprintf("%s", zcl.Zcstring(a))
}

const PhysicalEnvironmentAttr zcl.AttrID = 17

type PhysicalEnvironment zcl.Zenum8

func (a PhysicalEnvironment) ID() zcl.AttrID               { return PhysicalEnvironmentAttr }
func (a PhysicalEnvironment) Cluster() zcl.ClusterID       { return BasicID }
func (a *PhysicalEnvironment) Value() *PhysicalEnvironment { return a }
func (a PhysicalEnvironment) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *PhysicalEnvironment) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = PhysicalEnvironment(*nt)
	return br, err
}

func (a PhysicalEnvironment) Readable() bool   { return true }
func (a PhysicalEnvironment) Writable() bool   { return true }
func (a PhysicalEnvironment) Reportable() bool { return false }
func (a PhysicalEnvironment) SceneIndex() int  { return -1 }

func (a PhysicalEnvironment) String() string {
	switch a {
	case 0x00:
		return "Unspecified Environment"
	case 0x01:
		return "Atrium"
	case 0x02:
		return "Bar"
	case 0x03:
		return "Courtyard"
	case 0x04:
		return "Bathroom"
	case 0x05:
		return "Bedroom"
	case 0x06:
		return "Billiard Room"
	case 0x07:
		return "Utility Room"
	case 0x08:
		return "Cellar"
	case 0x09:
		return "Storage Closet"
	case 0x0a:
		return "Theater"
	case 0x0b:
		return "Office"
	case 0x0c:
		return "Deck"
	case 0x0d:
		return "Den"
	case 0x0e:
		return "Dining Room"
	case 0x0f:
		return "Electrical Room"
	case 0x10:
		return "Elevator"
	case 0x11:
		return "Entry"
	case 0x12:
		return "Family Room"
	case 0x13:
		return "Main Floor"
	case 0x14:
		return "Upstairs"
	case 0x15:
		return "Downstairs"
	case 0x16:
		return "Basement/Lower Level"
	case 0x17:
		return "Gallery"
	case 0x18:
		return "Game Room"
	case 0x19:
		return "Garage"
	case 0x1a:
		return "Gym"
	case 0x1b:
		return "Hallway"
	case 0x1c:
		return "House"
	case 0x1d:
		return "Kitchen"
	case 0x1e:
		return "Laundry Room"
	case 0x1f:
		return "Library"
	case 0x20:
		return "Master Bedroom"
	case 0x21:
		return "Mud Room (small room for coats and boots)"
	case 0x22:
		return "Nursery"
	case 0x23:
		return "Pantry"
	case 0x24:
		return "Office 2"
	case 0x25:
		return "Outside"
	case 0x26:
		return "Pool"
	case 0x27:
		return "Porch"
	case 0x28:
		return "Sewing Room"
	case 0x29:
		return "Sitting Room"
	case 0x2a:
		return "Stairway"
	case 0x2b:
		return "Yard"
	case 0x2c:
		return "Attic"
	case 0x2d:
		return "Hot Tub"
	case 0x2e:
		return "Living Room"
	case 0x2f:
		return "Sauna"
	case 0x30:
		return "Shop/Workshop"
	case 0x31:
		return "Guest Bedroom"
	case 0x32:
		return "Guest Bath"
	case 0x33:
		return "Powder Room (1/2 bath)"
	case 0x34:
		return "Back Yard"
	case 0x35:
		return "Front Yard"
	case 0x36:
		return "Patio"
	case 0x37:
		return "Driveway"
	case 0x38:
		return "Sun Room"
	case 0x39:
		return "Living Room 2"
	case 0x3a:
		return "Spa"
	case 0x3b:
		return "Whirlpool"
	case 0x3c:
		return "Shed"
	case 0x3d:
		return "Equipment Storage"
	case 0x3e:
		return "Hobby/Craft Room"
	case 0x3f:
		return "Fountain"
	case 0x40:
		return "Pond"
	case 0x41:
		return "Reception Room"
	case 0x42:
		return "Breakfast Room"
	case 0x43:
		return "Nook"
	case 0x44:
		return "Garden"
	case 0x45:
		return "Balcony"
	case 0x46:
		return "Panic Room"
	case 0x47:
		return "Terrace"
	case 0x48:
		return "Roof"
	case 0x49:
		return "Toilet"
	case 0x4a:
		return "Toilet Main"
	case 0x4b:
		return "Outside Toilet"
	case 0x4c:
		return "Shower room"
	case 0x4d:
		return "Study"
	case 0x4e:
		return "Front Garden"
	case 0x4f:
		return "Back Garden"
	case 0x50:
		return "Kettle"
	case 0x51:
		return "Television"
	case 0x52:
		return "Stove"
	case 0x53:
		return "Microwave"
	case 0x54:
		return "Toaster"
	case 0x55:
		return "Vacuum"
	case 0x56:
		return "Appliance"
	case 0x57:
		return "Front Door"
	case 0x58:
		return "Back Door"
	case 0x59:
		return "Fridge Door"
	case 0x60:
		return "Medication Cabinet Door"
	case 0x61:
		return "Wardrobe Door"
	case 0x62:
		return "Front Cupboard Door"
	case 0x63:
		return "Other Door"
	case 0x64:
		return "Waiting Room"
	case 0x65:
		return "Triage Room"
	case 0x66:
		return "Doctor’s Office"
	case 0x67:
		return "Patient’s Private Room"
	case 0x68:
		return "Consultation Room"
	case 0x69:
		return "Nurse Station"
	case 0x6a:
		return "Ward"
	case 0x6b:
		return "Corridor"
	case 0x6c:
		return "Operating Theatre"
	case 0x6d:
		return "Dental Surgery Room"
	case 0x6e:
		return "Medical Imaging Room"
	case 0x6f:
		return "Decontamination Room"
	case 0xFF:
		return "Unknown Environment"
	}

	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsUnspecifiedEnvironment checks if PhysicalEnvironment equals the value for Unspecified Environment (0x00)
func (a PhysicalEnvironment) IsUnspecifiedEnvironment() bool { return a == 0x00 }

// SetUnspecifiedEnvironment sets PhysicalEnvironment to Unspecified Environment (0x00)
func (a *PhysicalEnvironment) SetUnspecifiedEnvironment() { *a = 0x00 }

// IsAtrium checks if PhysicalEnvironment equals the value for Atrium (0x01)
func (a PhysicalEnvironment) IsAtrium() bool { return a == 0x01 }

// SetAtrium sets PhysicalEnvironment to Atrium (0x01)
func (a *PhysicalEnvironment) SetAtrium() { *a = 0x01 }

// IsBar checks if PhysicalEnvironment equals the value for Bar (0x02)
func (a PhysicalEnvironment) IsBar() bool { return a == 0x02 }

// SetBar sets PhysicalEnvironment to Bar (0x02)
func (a *PhysicalEnvironment) SetBar() { *a = 0x02 }

// IsCourtyard checks if PhysicalEnvironment equals the value for Courtyard (0x03)
func (a PhysicalEnvironment) IsCourtyard() bool { return a == 0x03 }

// SetCourtyard sets PhysicalEnvironment to Courtyard (0x03)
func (a *PhysicalEnvironment) SetCourtyard() { *a = 0x03 }

// IsBathroom checks if PhysicalEnvironment equals the value for Bathroom (0x04)
func (a PhysicalEnvironment) IsBathroom() bool { return a == 0x04 }

// SetBathroom sets PhysicalEnvironment to Bathroom (0x04)
func (a *PhysicalEnvironment) SetBathroom() { *a = 0x04 }

// IsBedroom checks if PhysicalEnvironment equals the value for Bedroom (0x05)
func (a PhysicalEnvironment) IsBedroom() bool { return a == 0x05 }

// SetBedroom sets PhysicalEnvironment to Bedroom (0x05)
func (a *PhysicalEnvironment) SetBedroom() { *a = 0x05 }

// IsBilliardRoom checks if PhysicalEnvironment equals the value for Billiard Room (0x06)
func (a PhysicalEnvironment) IsBilliardRoom() bool { return a == 0x06 }

// SetBilliardRoom sets PhysicalEnvironment to Billiard Room (0x06)
func (a *PhysicalEnvironment) SetBilliardRoom() { *a = 0x06 }

// IsUtilityRoom checks if PhysicalEnvironment equals the value for Utility Room (0x07)
func (a PhysicalEnvironment) IsUtilityRoom() bool { return a == 0x07 }

// SetUtilityRoom sets PhysicalEnvironment to Utility Room (0x07)
func (a *PhysicalEnvironment) SetUtilityRoom() { *a = 0x07 }

// IsCellar checks if PhysicalEnvironment equals the value for Cellar (0x08)
func (a PhysicalEnvironment) IsCellar() bool { return a == 0x08 }

// SetCellar sets PhysicalEnvironment to Cellar (0x08)
func (a *PhysicalEnvironment) SetCellar() { *a = 0x08 }

// IsStorageCloset checks if PhysicalEnvironment equals the value for Storage Closet (0x09)
func (a PhysicalEnvironment) IsStorageCloset() bool { return a == 0x09 }

// SetStorageCloset sets PhysicalEnvironment to Storage Closet (0x09)
func (a *PhysicalEnvironment) SetStorageCloset() { *a = 0x09 }

// IsTheater checks if PhysicalEnvironment equals the value for Theater (0x0a)
func (a PhysicalEnvironment) IsTheater() bool { return a == 0x0a }

// SetTheater sets PhysicalEnvironment to Theater (0x0a)
func (a *PhysicalEnvironment) SetTheater() { *a = 0x0a }

// IsOffice checks if PhysicalEnvironment equals the value for Office (0x0b)
func (a PhysicalEnvironment) IsOffice() bool { return a == 0x0b }

// SetOffice sets PhysicalEnvironment to Office (0x0b)
func (a *PhysicalEnvironment) SetOffice() { *a = 0x0b }

// IsDeck checks if PhysicalEnvironment equals the value for Deck (0x0c)
func (a PhysicalEnvironment) IsDeck() bool { return a == 0x0c }

// SetDeck sets PhysicalEnvironment to Deck (0x0c)
func (a *PhysicalEnvironment) SetDeck() { *a = 0x0c }

// IsDen checks if PhysicalEnvironment equals the value for Den (0x0d)
func (a PhysicalEnvironment) IsDen() bool { return a == 0x0d }

// SetDen sets PhysicalEnvironment to Den (0x0d)
func (a *PhysicalEnvironment) SetDen() { *a = 0x0d }

// IsDiningRoom checks if PhysicalEnvironment equals the value for Dining Room (0x0e)
func (a PhysicalEnvironment) IsDiningRoom() bool { return a == 0x0e }

// SetDiningRoom sets PhysicalEnvironment to Dining Room (0x0e)
func (a *PhysicalEnvironment) SetDiningRoom() { *a = 0x0e }

// IsElectricalRoom checks if PhysicalEnvironment equals the value for Electrical Room (0x0f)
func (a PhysicalEnvironment) IsElectricalRoom() bool { return a == 0x0f }

// SetElectricalRoom sets PhysicalEnvironment to Electrical Room (0x0f)
func (a *PhysicalEnvironment) SetElectricalRoom() { *a = 0x0f }

// IsElevator checks if PhysicalEnvironment equals the value for Elevator (0x10)
func (a PhysicalEnvironment) IsElevator() bool { return a == 0x10 }

// SetElevator sets PhysicalEnvironment to Elevator (0x10)
func (a *PhysicalEnvironment) SetElevator() { *a = 0x10 }

// IsEntry checks if PhysicalEnvironment equals the value for Entry (0x11)
func (a PhysicalEnvironment) IsEntry() bool { return a == 0x11 }

// SetEntry sets PhysicalEnvironment to Entry (0x11)
func (a *PhysicalEnvironment) SetEntry() { *a = 0x11 }

// IsFamilyRoom checks if PhysicalEnvironment equals the value for Family Room (0x12)
func (a PhysicalEnvironment) IsFamilyRoom() bool { return a == 0x12 }

// SetFamilyRoom sets PhysicalEnvironment to Family Room (0x12)
func (a *PhysicalEnvironment) SetFamilyRoom() { *a = 0x12 }

// IsMainFloor checks if PhysicalEnvironment equals the value for Main Floor (0x13)
func (a PhysicalEnvironment) IsMainFloor() bool { return a == 0x13 }

// SetMainFloor sets PhysicalEnvironment to Main Floor (0x13)
func (a *PhysicalEnvironment) SetMainFloor() { *a = 0x13 }

// IsUpstairs checks if PhysicalEnvironment equals the value for Upstairs (0x14)
func (a PhysicalEnvironment) IsUpstairs() bool { return a == 0x14 }

// SetUpstairs sets PhysicalEnvironment to Upstairs (0x14)
func (a *PhysicalEnvironment) SetUpstairs() { *a = 0x14 }

// IsDownstairs checks if PhysicalEnvironment equals the value for Downstairs (0x15)
func (a PhysicalEnvironment) IsDownstairs() bool { return a == 0x15 }

// SetDownstairs sets PhysicalEnvironment to Downstairs (0x15)
func (a *PhysicalEnvironment) SetDownstairs() { *a = 0x15 }

// IsBasementLowerLevel checks if PhysicalEnvironment equals the value for Basement/Lower Level (0x16)
func (a PhysicalEnvironment) IsBasementLowerLevel() bool { return a == 0x16 }

// SetBasementLowerLevel sets PhysicalEnvironment to Basement/Lower Level (0x16)
func (a *PhysicalEnvironment) SetBasementLowerLevel() { *a = 0x16 }

// IsGallery checks if PhysicalEnvironment equals the value for Gallery (0x17)
func (a PhysicalEnvironment) IsGallery() bool { return a == 0x17 }

// SetGallery sets PhysicalEnvironment to Gallery (0x17)
func (a *PhysicalEnvironment) SetGallery() { *a = 0x17 }

// IsGameRoom checks if PhysicalEnvironment equals the value for Game Room (0x18)
func (a PhysicalEnvironment) IsGameRoom() bool { return a == 0x18 }

// SetGameRoom sets PhysicalEnvironment to Game Room (0x18)
func (a *PhysicalEnvironment) SetGameRoom() { *a = 0x18 }

// IsGarage checks if PhysicalEnvironment equals the value for Garage (0x19)
func (a PhysicalEnvironment) IsGarage() bool { return a == 0x19 }

// SetGarage sets PhysicalEnvironment to Garage (0x19)
func (a *PhysicalEnvironment) SetGarage() { *a = 0x19 }

// IsGym checks if PhysicalEnvironment equals the value for Gym (0x1a)
func (a PhysicalEnvironment) IsGym() bool { return a == 0x1a }

// SetGym sets PhysicalEnvironment to Gym (0x1a)
func (a *PhysicalEnvironment) SetGym() { *a = 0x1a }

// IsHallway checks if PhysicalEnvironment equals the value for Hallway (0x1b)
func (a PhysicalEnvironment) IsHallway() bool { return a == 0x1b }

// SetHallway sets PhysicalEnvironment to Hallway (0x1b)
func (a *PhysicalEnvironment) SetHallway() { *a = 0x1b }

// IsHouse checks if PhysicalEnvironment equals the value for House (0x1c)
func (a PhysicalEnvironment) IsHouse() bool { return a == 0x1c }

// SetHouse sets PhysicalEnvironment to House (0x1c)
func (a *PhysicalEnvironment) SetHouse() { *a = 0x1c }

// IsKitchen checks if PhysicalEnvironment equals the value for Kitchen (0x1d)
func (a PhysicalEnvironment) IsKitchen() bool { return a == 0x1d }

// SetKitchen sets PhysicalEnvironment to Kitchen (0x1d)
func (a *PhysicalEnvironment) SetKitchen() { *a = 0x1d }

// IsLaundryRoom checks if PhysicalEnvironment equals the value for Laundry Room (0x1e)
func (a PhysicalEnvironment) IsLaundryRoom() bool { return a == 0x1e }

// SetLaundryRoom sets PhysicalEnvironment to Laundry Room (0x1e)
func (a *PhysicalEnvironment) SetLaundryRoom() { *a = 0x1e }

// IsLibrary checks if PhysicalEnvironment equals the value for Library (0x1f)
func (a PhysicalEnvironment) IsLibrary() bool { return a == 0x1f }

// SetLibrary sets PhysicalEnvironment to Library (0x1f)
func (a *PhysicalEnvironment) SetLibrary() { *a = 0x1f }

// IsMasterBedroom checks if PhysicalEnvironment equals the value for Master Bedroom (0x20)
func (a PhysicalEnvironment) IsMasterBedroom() bool { return a == 0x20 }

// SetMasterBedroom sets PhysicalEnvironment to Master Bedroom (0x20)
func (a *PhysicalEnvironment) SetMasterBedroom() { *a = 0x20 }

// IsMudRoomSmallRoomForCoatsAndBoots checks if PhysicalEnvironment equals the value for Mud Room (small room for coats and boots) (0x21)
func (a PhysicalEnvironment) IsMudRoomSmallRoomForCoatsAndBoots() bool { return a == 0x21 }

// SetMudRoomSmallRoomForCoatsAndBoots sets PhysicalEnvironment to Mud Room (small room for coats and boots) (0x21)
func (a *PhysicalEnvironment) SetMudRoomSmallRoomForCoatsAndBoots() { *a = 0x21 }

// IsNursery checks if PhysicalEnvironment equals the value for Nursery (0x22)
func (a PhysicalEnvironment) IsNursery() bool { return a == 0x22 }

// SetNursery sets PhysicalEnvironment to Nursery (0x22)
func (a *PhysicalEnvironment) SetNursery() { *a = 0x22 }

// IsPantry checks if PhysicalEnvironment equals the value for Pantry (0x23)
func (a PhysicalEnvironment) IsPantry() bool { return a == 0x23 }

// SetPantry sets PhysicalEnvironment to Pantry (0x23)
func (a *PhysicalEnvironment) SetPantry() { *a = 0x23 }

// IsOffice2 checks if PhysicalEnvironment equals the value for Office 2 (0x24)
func (a PhysicalEnvironment) IsOffice2() bool { return a == 0x24 }

// SetOffice2 sets PhysicalEnvironment to Office 2 (0x24)
func (a *PhysicalEnvironment) SetOffice2() { *a = 0x24 }

// IsOutside checks if PhysicalEnvironment equals the value for Outside (0x25)
func (a PhysicalEnvironment) IsOutside() bool { return a == 0x25 }

// SetOutside sets PhysicalEnvironment to Outside (0x25)
func (a *PhysicalEnvironment) SetOutside() { *a = 0x25 }

// IsPool checks if PhysicalEnvironment equals the value for Pool (0x26)
func (a PhysicalEnvironment) IsPool() bool { return a == 0x26 }

// SetPool sets PhysicalEnvironment to Pool (0x26)
func (a *PhysicalEnvironment) SetPool() { *a = 0x26 }

// IsPorch checks if PhysicalEnvironment equals the value for Porch (0x27)
func (a PhysicalEnvironment) IsPorch() bool { return a == 0x27 }

// SetPorch sets PhysicalEnvironment to Porch (0x27)
func (a *PhysicalEnvironment) SetPorch() { *a = 0x27 }

// IsSewingRoom checks if PhysicalEnvironment equals the value for Sewing Room (0x28)
func (a PhysicalEnvironment) IsSewingRoom() bool { return a == 0x28 }

// SetSewingRoom sets PhysicalEnvironment to Sewing Room (0x28)
func (a *PhysicalEnvironment) SetSewingRoom() { *a = 0x28 }

// IsSittingRoom checks if PhysicalEnvironment equals the value for Sitting Room (0x29)
func (a PhysicalEnvironment) IsSittingRoom() bool { return a == 0x29 }

// SetSittingRoom sets PhysicalEnvironment to Sitting Room (0x29)
func (a *PhysicalEnvironment) SetSittingRoom() { *a = 0x29 }

// IsStairway checks if PhysicalEnvironment equals the value for Stairway (0x2a)
func (a PhysicalEnvironment) IsStairway() bool { return a == 0x2a }

// SetStairway sets PhysicalEnvironment to Stairway (0x2a)
func (a *PhysicalEnvironment) SetStairway() { *a = 0x2a }

// IsYard checks if PhysicalEnvironment equals the value for Yard (0x2b)
func (a PhysicalEnvironment) IsYard() bool { return a == 0x2b }

// SetYard sets PhysicalEnvironment to Yard (0x2b)
func (a *PhysicalEnvironment) SetYard() { *a = 0x2b }

// IsAttic checks if PhysicalEnvironment equals the value for Attic (0x2c)
func (a PhysicalEnvironment) IsAttic() bool { return a == 0x2c }

// SetAttic sets PhysicalEnvironment to Attic (0x2c)
func (a *PhysicalEnvironment) SetAttic() { *a = 0x2c }

// IsHotTub checks if PhysicalEnvironment equals the value for Hot Tub (0x2d)
func (a PhysicalEnvironment) IsHotTub() bool { return a == 0x2d }

// SetHotTub sets PhysicalEnvironment to Hot Tub (0x2d)
func (a *PhysicalEnvironment) SetHotTub() { *a = 0x2d }

// IsLivingRoom checks if PhysicalEnvironment equals the value for Living Room (0x2e)
func (a PhysicalEnvironment) IsLivingRoom() bool { return a == 0x2e }

// SetLivingRoom sets PhysicalEnvironment to Living Room (0x2e)
func (a *PhysicalEnvironment) SetLivingRoom() { *a = 0x2e }

// IsSauna checks if PhysicalEnvironment equals the value for Sauna (0x2f)
func (a PhysicalEnvironment) IsSauna() bool { return a == 0x2f }

// SetSauna sets PhysicalEnvironment to Sauna (0x2f)
func (a *PhysicalEnvironment) SetSauna() { *a = 0x2f }

// IsShopWorkshop checks if PhysicalEnvironment equals the value for Shop/Workshop (0x30)
func (a PhysicalEnvironment) IsShopWorkshop() bool { return a == 0x30 }

// SetShopWorkshop sets PhysicalEnvironment to Shop/Workshop (0x30)
func (a *PhysicalEnvironment) SetShopWorkshop() { *a = 0x30 }

// IsGuestBedroom checks if PhysicalEnvironment equals the value for Guest Bedroom (0x31)
func (a PhysicalEnvironment) IsGuestBedroom() bool { return a == 0x31 }

// SetGuestBedroom sets PhysicalEnvironment to Guest Bedroom (0x31)
func (a *PhysicalEnvironment) SetGuestBedroom() { *a = 0x31 }

// IsGuestBath checks if PhysicalEnvironment equals the value for Guest Bath (0x32)
func (a PhysicalEnvironment) IsGuestBath() bool { return a == 0x32 }

// SetGuestBath sets PhysicalEnvironment to Guest Bath (0x32)
func (a *PhysicalEnvironment) SetGuestBath() { *a = 0x32 }

// IsPowderRoom12Bath checks if PhysicalEnvironment equals the value for Powder Room (1/2 bath) (0x33)
func (a PhysicalEnvironment) IsPowderRoom12Bath() bool { return a == 0x33 }

// SetPowderRoom12Bath sets PhysicalEnvironment to Powder Room (1/2 bath) (0x33)
func (a *PhysicalEnvironment) SetPowderRoom12Bath() { *a = 0x33 }

// IsBackYard checks if PhysicalEnvironment equals the value for Back Yard (0x34)
func (a PhysicalEnvironment) IsBackYard() bool { return a == 0x34 }

// SetBackYard sets PhysicalEnvironment to Back Yard (0x34)
func (a *PhysicalEnvironment) SetBackYard() { *a = 0x34 }

// IsFrontYard checks if PhysicalEnvironment equals the value for Front Yard (0x35)
func (a PhysicalEnvironment) IsFrontYard() bool { return a == 0x35 }

// SetFrontYard sets PhysicalEnvironment to Front Yard (0x35)
func (a *PhysicalEnvironment) SetFrontYard() { *a = 0x35 }

// IsPatio checks if PhysicalEnvironment equals the value for Patio (0x36)
func (a PhysicalEnvironment) IsPatio() bool { return a == 0x36 }

// SetPatio sets PhysicalEnvironment to Patio (0x36)
func (a *PhysicalEnvironment) SetPatio() { *a = 0x36 }

// IsDriveway checks if PhysicalEnvironment equals the value for Driveway (0x37)
func (a PhysicalEnvironment) IsDriveway() bool { return a == 0x37 }

// SetDriveway sets PhysicalEnvironment to Driveway (0x37)
func (a *PhysicalEnvironment) SetDriveway() { *a = 0x37 }

// IsSunRoom checks if PhysicalEnvironment equals the value for Sun Room (0x38)
func (a PhysicalEnvironment) IsSunRoom() bool { return a == 0x38 }

// SetSunRoom sets PhysicalEnvironment to Sun Room (0x38)
func (a *PhysicalEnvironment) SetSunRoom() { *a = 0x38 }

// IsLivingRoom2 checks if PhysicalEnvironment equals the value for Living Room 2 (0x39)
func (a PhysicalEnvironment) IsLivingRoom2() bool { return a == 0x39 }

// SetLivingRoom2 sets PhysicalEnvironment to Living Room 2 (0x39)
func (a *PhysicalEnvironment) SetLivingRoom2() { *a = 0x39 }

// IsSpa checks if PhysicalEnvironment equals the value for Spa (0x3a)
func (a PhysicalEnvironment) IsSpa() bool { return a == 0x3a }

// SetSpa sets PhysicalEnvironment to Spa (0x3a)
func (a *PhysicalEnvironment) SetSpa() { *a = 0x3a }

// IsWhirlpool checks if PhysicalEnvironment equals the value for Whirlpool (0x3b)
func (a PhysicalEnvironment) IsWhirlpool() bool { return a == 0x3b }

// SetWhirlpool sets PhysicalEnvironment to Whirlpool (0x3b)
func (a *PhysicalEnvironment) SetWhirlpool() { *a = 0x3b }

// IsShed checks if PhysicalEnvironment equals the value for Shed (0x3c)
func (a PhysicalEnvironment) IsShed() bool { return a == 0x3c }

// SetShed sets PhysicalEnvironment to Shed (0x3c)
func (a *PhysicalEnvironment) SetShed() { *a = 0x3c }

// IsEquipmentStorage checks if PhysicalEnvironment equals the value for Equipment Storage (0x3d)
func (a PhysicalEnvironment) IsEquipmentStorage() bool { return a == 0x3d }

// SetEquipmentStorage sets PhysicalEnvironment to Equipment Storage (0x3d)
func (a *PhysicalEnvironment) SetEquipmentStorage() { *a = 0x3d }

// IsHobbyCraftRoom checks if PhysicalEnvironment equals the value for Hobby/Craft Room (0x3e)
func (a PhysicalEnvironment) IsHobbyCraftRoom() bool { return a == 0x3e }

// SetHobbyCraftRoom sets PhysicalEnvironment to Hobby/Craft Room (0x3e)
func (a *PhysicalEnvironment) SetHobbyCraftRoom() { *a = 0x3e }

// IsFountain checks if PhysicalEnvironment equals the value for Fountain (0x3f)
func (a PhysicalEnvironment) IsFountain() bool { return a == 0x3f }

// SetFountain sets PhysicalEnvironment to Fountain (0x3f)
func (a *PhysicalEnvironment) SetFountain() { *a = 0x3f }

// IsPond checks if PhysicalEnvironment equals the value for Pond (0x40)
func (a PhysicalEnvironment) IsPond() bool { return a == 0x40 }

// SetPond sets PhysicalEnvironment to Pond (0x40)
func (a *PhysicalEnvironment) SetPond() { *a = 0x40 }

// IsReceptionRoom checks if PhysicalEnvironment equals the value for Reception Room (0x41)
func (a PhysicalEnvironment) IsReceptionRoom() bool { return a == 0x41 }

// SetReceptionRoom sets PhysicalEnvironment to Reception Room (0x41)
func (a *PhysicalEnvironment) SetReceptionRoom() { *a = 0x41 }

// IsBreakfastRoom checks if PhysicalEnvironment equals the value for Breakfast Room (0x42)
func (a PhysicalEnvironment) IsBreakfastRoom() bool { return a == 0x42 }

// SetBreakfastRoom sets PhysicalEnvironment to Breakfast Room (0x42)
func (a *PhysicalEnvironment) SetBreakfastRoom() { *a = 0x42 }

// IsNook checks if PhysicalEnvironment equals the value for Nook (0x43)
func (a PhysicalEnvironment) IsNook() bool { return a == 0x43 }

// SetNook sets PhysicalEnvironment to Nook (0x43)
func (a *PhysicalEnvironment) SetNook() { *a = 0x43 }

// IsGarden checks if PhysicalEnvironment equals the value for Garden (0x44)
func (a PhysicalEnvironment) IsGarden() bool { return a == 0x44 }

// SetGarden sets PhysicalEnvironment to Garden (0x44)
func (a *PhysicalEnvironment) SetGarden() { *a = 0x44 }

// IsBalcony checks if PhysicalEnvironment equals the value for Balcony (0x45)
func (a PhysicalEnvironment) IsBalcony() bool { return a == 0x45 }

// SetBalcony sets PhysicalEnvironment to Balcony (0x45)
func (a *PhysicalEnvironment) SetBalcony() { *a = 0x45 }

// IsPanicRoom checks if PhysicalEnvironment equals the value for Panic Room (0x46)
func (a PhysicalEnvironment) IsPanicRoom() bool { return a == 0x46 }

// SetPanicRoom sets PhysicalEnvironment to Panic Room (0x46)
func (a *PhysicalEnvironment) SetPanicRoom() { *a = 0x46 }

// IsTerrace checks if PhysicalEnvironment equals the value for Terrace (0x47)
func (a PhysicalEnvironment) IsTerrace() bool { return a == 0x47 }

// SetTerrace sets PhysicalEnvironment to Terrace (0x47)
func (a *PhysicalEnvironment) SetTerrace() { *a = 0x47 }

// IsRoof checks if PhysicalEnvironment equals the value for Roof (0x48)
func (a PhysicalEnvironment) IsRoof() bool { return a == 0x48 }

// SetRoof sets PhysicalEnvironment to Roof (0x48)
func (a *PhysicalEnvironment) SetRoof() { *a = 0x48 }

// IsToilet checks if PhysicalEnvironment equals the value for Toilet (0x49)
func (a PhysicalEnvironment) IsToilet() bool { return a == 0x49 }

// SetToilet sets PhysicalEnvironment to Toilet (0x49)
func (a *PhysicalEnvironment) SetToilet() { *a = 0x49 }

// IsToiletMain checks if PhysicalEnvironment equals the value for Toilet Main (0x4a)
func (a PhysicalEnvironment) IsToiletMain() bool { return a == 0x4a }

// SetToiletMain sets PhysicalEnvironment to Toilet Main (0x4a)
func (a *PhysicalEnvironment) SetToiletMain() { *a = 0x4a }

// IsOutsideToilet checks if PhysicalEnvironment equals the value for Outside Toilet (0x4b)
func (a PhysicalEnvironment) IsOutsideToilet() bool { return a == 0x4b }

// SetOutsideToilet sets PhysicalEnvironment to Outside Toilet (0x4b)
func (a *PhysicalEnvironment) SetOutsideToilet() { *a = 0x4b }

// IsShowerRoom checks if PhysicalEnvironment equals the value for Shower room (0x4c)
func (a PhysicalEnvironment) IsShowerRoom() bool { return a == 0x4c }

// SetShowerRoom sets PhysicalEnvironment to Shower room (0x4c)
func (a *PhysicalEnvironment) SetShowerRoom() { *a = 0x4c }

// IsStudy checks if PhysicalEnvironment equals the value for Study (0x4d)
func (a PhysicalEnvironment) IsStudy() bool { return a == 0x4d }

// SetStudy sets PhysicalEnvironment to Study (0x4d)
func (a *PhysicalEnvironment) SetStudy() { *a = 0x4d }

// IsFrontGarden checks if PhysicalEnvironment equals the value for Front Garden (0x4e)
func (a PhysicalEnvironment) IsFrontGarden() bool { return a == 0x4e }

// SetFrontGarden sets PhysicalEnvironment to Front Garden (0x4e)
func (a *PhysicalEnvironment) SetFrontGarden() { *a = 0x4e }

// IsBackGarden checks if PhysicalEnvironment equals the value for Back Garden (0x4f)
func (a PhysicalEnvironment) IsBackGarden() bool { return a == 0x4f }

// SetBackGarden sets PhysicalEnvironment to Back Garden (0x4f)
func (a *PhysicalEnvironment) SetBackGarden() { *a = 0x4f }

// IsKettle checks if PhysicalEnvironment equals the value for Kettle (0x50)
func (a PhysicalEnvironment) IsKettle() bool { return a == 0x50 }

// SetKettle sets PhysicalEnvironment to Kettle (0x50)
func (a *PhysicalEnvironment) SetKettle() { *a = 0x50 }

// IsTelevision checks if PhysicalEnvironment equals the value for Television (0x51)
func (a PhysicalEnvironment) IsTelevision() bool { return a == 0x51 }

// SetTelevision sets PhysicalEnvironment to Television (0x51)
func (a *PhysicalEnvironment) SetTelevision() { *a = 0x51 }

// IsStove checks if PhysicalEnvironment equals the value for Stove (0x52)
func (a PhysicalEnvironment) IsStove() bool { return a == 0x52 }

// SetStove sets PhysicalEnvironment to Stove (0x52)
func (a *PhysicalEnvironment) SetStove() { *a = 0x52 }

// IsMicrowave checks if PhysicalEnvironment equals the value for Microwave (0x53)
func (a PhysicalEnvironment) IsMicrowave() bool { return a == 0x53 }

// SetMicrowave sets PhysicalEnvironment to Microwave (0x53)
func (a *PhysicalEnvironment) SetMicrowave() { *a = 0x53 }

// IsToaster checks if PhysicalEnvironment equals the value for Toaster (0x54)
func (a PhysicalEnvironment) IsToaster() bool { return a == 0x54 }

// SetToaster sets PhysicalEnvironment to Toaster (0x54)
func (a *PhysicalEnvironment) SetToaster() { *a = 0x54 }

// IsVacuum checks if PhysicalEnvironment equals the value for Vacuum (0x55)
func (a PhysicalEnvironment) IsVacuum() bool { return a == 0x55 }

// SetVacuum sets PhysicalEnvironment to Vacuum (0x55)
func (a *PhysicalEnvironment) SetVacuum() { *a = 0x55 }

// IsAppliance checks if PhysicalEnvironment equals the value for Appliance (0x56)
func (a PhysicalEnvironment) IsAppliance() bool { return a == 0x56 }

// SetAppliance sets PhysicalEnvironment to Appliance (0x56)
func (a *PhysicalEnvironment) SetAppliance() { *a = 0x56 }

// IsFrontDoor checks if PhysicalEnvironment equals the value for Front Door (0x57)
func (a PhysicalEnvironment) IsFrontDoor() bool { return a == 0x57 }

// SetFrontDoor sets PhysicalEnvironment to Front Door (0x57)
func (a *PhysicalEnvironment) SetFrontDoor() { *a = 0x57 }

// IsBackDoor checks if PhysicalEnvironment equals the value for Back Door (0x58)
func (a PhysicalEnvironment) IsBackDoor() bool { return a == 0x58 }

// SetBackDoor sets PhysicalEnvironment to Back Door (0x58)
func (a *PhysicalEnvironment) SetBackDoor() { *a = 0x58 }

// IsFridgeDoor checks if PhysicalEnvironment equals the value for Fridge Door (0x59)
func (a PhysicalEnvironment) IsFridgeDoor() bool { return a == 0x59 }

// SetFridgeDoor sets PhysicalEnvironment to Fridge Door (0x59)
func (a *PhysicalEnvironment) SetFridgeDoor() { *a = 0x59 }

// IsMedicationCabinetDoor checks if PhysicalEnvironment equals the value for Medication Cabinet Door (0x60)
func (a PhysicalEnvironment) IsMedicationCabinetDoor() bool { return a == 0x60 }

// SetMedicationCabinetDoor sets PhysicalEnvironment to Medication Cabinet Door (0x60)
func (a *PhysicalEnvironment) SetMedicationCabinetDoor() { *a = 0x60 }

// IsWardrobeDoor checks if PhysicalEnvironment equals the value for Wardrobe Door (0x61)
func (a PhysicalEnvironment) IsWardrobeDoor() bool { return a == 0x61 }

// SetWardrobeDoor sets PhysicalEnvironment to Wardrobe Door (0x61)
func (a *PhysicalEnvironment) SetWardrobeDoor() { *a = 0x61 }

// IsFrontCupboardDoor checks if PhysicalEnvironment equals the value for Front Cupboard Door (0x62)
func (a PhysicalEnvironment) IsFrontCupboardDoor() bool { return a == 0x62 }

// SetFrontCupboardDoor sets PhysicalEnvironment to Front Cupboard Door (0x62)
func (a *PhysicalEnvironment) SetFrontCupboardDoor() { *a = 0x62 }

// IsOtherDoor checks if PhysicalEnvironment equals the value for Other Door (0x63)
func (a PhysicalEnvironment) IsOtherDoor() bool { return a == 0x63 }

// SetOtherDoor sets PhysicalEnvironment to Other Door (0x63)
func (a *PhysicalEnvironment) SetOtherDoor() { *a = 0x63 }

// IsWaitingRoom checks if PhysicalEnvironment equals the value for Waiting Room (0x64)
func (a PhysicalEnvironment) IsWaitingRoom() bool { return a == 0x64 }

// SetWaitingRoom sets PhysicalEnvironment to Waiting Room (0x64)
func (a *PhysicalEnvironment) SetWaitingRoom() { *a = 0x64 }

// IsTriageRoom checks if PhysicalEnvironment equals the value for Triage Room (0x65)
func (a PhysicalEnvironment) IsTriageRoom() bool { return a == 0x65 }

// SetTriageRoom sets PhysicalEnvironment to Triage Room (0x65)
func (a *PhysicalEnvironment) SetTriageRoom() { *a = 0x65 }

// IsDoctorSOffice checks if PhysicalEnvironment equals the value for Doctor’s Office (0x66)
func (a PhysicalEnvironment) IsDoctorSOffice() bool { return a == 0x66 }

// SetDoctorSOffice sets PhysicalEnvironment to Doctor’s Office (0x66)
func (a *PhysicalEnvironment) SetDoctorSOffice() { *a = 0x66 }

// IsPatientSPrivateRoom checks if PhysicalEnvironment equals the value for Patient’s Private Room (0x67)
func (a PhysicalEnvironment) IsPatientSPrivateRoom() bool { return a == 0x67 }

// SetPatientSPrivateRoom sets PhysicalEnvironment to Patient’s Private Room (0x67)
func (a *PhysicalEnvironment) SetPatientSPrivateRoom() { *a = 0x67 }

// IsConsultationRoom checks if PhysicalEnvironment equals the value for Consultation Room (0x68)
func (a PhysicalEnvironment) IsConsultationRoom() bool { return a == 0x68 }

// SetConsultationRoom sets PhysicalEnvironment to Consultation Room (0x68)
func (a *PhysicalEnvironment) SetConsultationRoom() { *a = 0x68 }

// IsNurseStation checks if PhysicalEnvironment equals the value for Nurse Station (0x69)
func (a PhysicalEnvironment) IsNurseStation() bool { return a == 0x69 }

// SetNurseStation sets PhysicalEnvironment to Nurse Station (0x69)
func (a *PhysicalEnvironment) SetNurseStation() { *a = 0x69 }

// IsWard checks if PhysicalEnvironment equals the value for Ward (0x6a)
func (a PhysicalEnvironment) IsWard() bool { return a == 0x6a }

// SetWard sets PhysicalEnvironment to Ward (0x6a)
func (a *PhysicalEnvironment) SetWard() { *a = 0x6a }

// IsCorridor checks if PhysicalEnvironment equals the value for Corridor (0x6b)
func (a PhysicalEnvironment) IsCorridor() bool { return a == 0x6b }

// SetCorridor sets PhysicalEnvironment to Corridor (0x6b)
func (a *PhysicalEnvironment) SetCorridor() { *a = 0x6b }

// IsOperatingTheatre checks if PhysicalEnvironment equals the value for Operating Theatre (0x6c)
func (a PhysicalEnvironment) IsOperatingTheatre() bool { return a == 0x6c }

// SetOperatingTheatre sets PhysicalEnvironment to Operating Theatre (0x6c)
func (a *PhysicalEnvironment) SetOperatingTheatre() { *a = 0x6c }

// IsDentalSurgeryRoom checks if PhysicalEnvironment equals the value for Dental Surgery Room (0x6d)
func (a PhysicalEnvironment) IsDentalSurgeryRoom() bool { return a == 0x6d }

// SetDentalSurgeryRoom sets PhysicalEnvironment to Dental Surgery Room (0x6d)
func (a *PhysicalEnvironment) SetDentalSurgeryRoom() { *a = 0x6d }

// IsMedicalImagingRoom checks if PhysicalEnvironment equals the value for Medical Imaging Room (0x6e)
func (a PhysicalEnvironment) IsMedicalImagingRoom() bool { return a == 0x6e }

// SetMedicalImagingRoom sets PhysicalEnvironment to Medical Imaging Room (0x6e)
func (a *PhysicalEnvironment) SetMedicalImagingRoom() { *a = 0x6e }

// IsDecontaminationRoom checks if PhysicalEnvironment equals the value for Decontamination Room (0x6f)
func (a PhysicalEnvironment) IsDecontaminationRoom() bool { return a == 0x6f }

// SetDecontaminationRoom sets PhysicalEnvironment to Decontamination Room (0x6f)
func (a *PhysicalEnvironment) SetDecontaminationRoom() { *a = 0x6f }

// IsUnknownEnvironment checks if PhysicalEnvironment equals the value for Unknown Environment (0xFF)
func (a PhysicalEnvironment) IsUnknownEnvironment() bool { return a == 0xFF }

// SetUnknownEnvironment sets PhysicalEnvironment to Unknown Environment (0xFF)
func (a *PhysicalEnvironment) SetUnknownEnvironment() { *a = 0xFF }

const DeviceEnabledAttr zcl.AttrID = 18

type DeviceEnabled zcl.Zbool

func (a DeviceEnabled) ID() zcl.AttrID         { return DeviceEnabledAttr }
func (a DeviceEnabled) Cluster() zcl.ClusterID { return BasicID }
func (a *DeviceEnabled) Value() *DeviceEnabled { return a }
func (a DeviceEnabled) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *DeviceEnabled) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = DeviceEnabled(*nt)
	return br, err
}

func (a DeviceEnabled) Readable() bool   { return true }
func (a DeviceEnabled) Writable() bool   { return true }
func (a DeviceEnabled) Reportable() bool { return false }
func (a DeviceEnabled) SceneIndex() int  { return -1 }

func (a DeviceEnabled) String() string {

	return zcl.Sprintf("%s", zcl.Zbool(a))
}

const AlarmMaskAttr zcl.AttrID = 19

type AlarmMask zcl.Zbmp8

func (a AlarmMask) ID() zcl.AttrID         { return AlarmMaskAttr }
func (a AlarmMask) Cluster() zcl.ClusterID { return BasicID }
func (a *AlarmMask) Value() *AlarmMask     { return a }
func (a AlarmMask) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *AlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = AlarmMask(*nt)
	return br, err
}

func (a AlarmMask) Readable() bool   { return true }
func (a AlarmMask) Writable() bool   { return true }
func (a AlarmMask) Reportable() bool { return false }
func (a AlarmMask) SceneIndex() int  { return -1 }

func (a AlarmMask) String() string {

	var bstr []string
	if a.IsGeneralHardwareFault() {
		bstr = append(bstr, "General Hardware Fault")
	}
	if a.IsGeneralSoftwareFault() {
		bstr = append(bstr, "General Software Fault")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a AlarmMask) IsGeneralHardwareFault() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *AlarmMask) SetGeneralHardwareFault(b bool) {
	*a = AlarmMask(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a AlarmMask) IsGeneralSoftwareFault() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *AlarmMask) SetGeneralSoftwareFault(b bool) {
	*a = AlarmMask(zcl.BitmapSet([]byte(*a), 1, b))
}

const DisableLocalConfigAttr zcl.AttrID = 20

type DisableLocalConfig zcl.Zbmp8

func (a DisableLocalConfig) ID() zcl.AttrID              { return DisableLocalConfigAttr }
func (a DisableLocalConfig) Cluster() zcl.ClusterID      { return BasicID }
func (a *DisableLocalConfig) Value() *DisableLocalConfig { return a }
func (a DisableLocalConfig) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *DisableLocalConfig) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = DisableLocalConfig(*nt)
	return br, err
}

func (a DisableLocalConfig) Readable() bool   { return true }
func (a DisableLocalConfig) Writable() bool   { return true }
func (a DisableLocalConfig) Reportable() bool { return false }
func (a DisableLocalConfig) SceneIndex() int  { return -1 }

func (a DisableLocalConfig) String() string {

	var bstr []string
	if a.IsDisableFactoryReset() {
		bstr = append(bstr, "Disable Factory Reset")
	}
	if a.IsDisableDeviceConfiguration() {
		bstr = append(bstr, "Disable Device Configuration")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a DisableLocalConfig) IsDisableFactoryReset() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *DisableLocalConfig) SetDisableFactoryReset(b bool) {
	*a = DisableLocalConfig(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a DisableLocalConfig) IsDisableDeviceConfiguration() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *DisableLocalConfig) SetDisableDeviceConfiguration(b bool) {
	*a = DisableLocalConfig(zcl.BitmapSet([]byte(*a), 1, b))
}

const SwBuildIdAttr zcl.AttrID = 16384

type SwBuildId zcl.Zcstring

func (a SwBuildId) ID() zcl.AttrID         { return SwBuildIdAttr }
func (a SwBuildId) Cluster() zcl.ClusterID { return BasicID }
func (a *SwBuildId) Value() *SwBuildId     { return a }
func (a SwBuildId) MarshalZcl() ([]byte, error) {
	return zcl.Zcstring(a).MarshalZcl()
}
func (a *SwBuildId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = SwBuildId(*nt)
	return br, err
}

func (a SwBuildId) Readable() bool   { return true }
func (a SwBuildId) Writable() bool   { return false }
func (a SwBuildId) Reportable() bool { return false }
func (a SwBuildId) SceneIndex() int  { return -1 }

func (a SwBuildId) String() string {

	return zcl.Sprintf("%s", zcl.Zcstring(a))
}

const Unknown1Attr zcl.AttrID = 8

type Unknown1 zcl.Zenum8

func (a Unknown1) ID() zcl.AttrID         { return Unknown1Attr }
func (a Unknown1) Cluster() zcl.ClusterID { return BasicID }
func (a *Unknown1) Value() *Unknown1      { return a }
func (a Unknown1) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *Unknown1) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = Unknown1(*nt)
	return br, err
}

func (a Unknown1) Readable() bool   { return true }
func (a Unknown1) Writable() bool   { return false }
func (a Unknown1) Reportable() bool { return false }
func (a Unknown1) SceneIndex() int  { return -1 }

func (a Unknown1) String() string {

	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

const Unknown2Attr zcl.AttrID = 9

type Unknown2 zcl.Zenum8

func (a Unknown2) ID() zcl.AttrID         { return Unknown2Attr }
func (a Unknown2) Cluster() zcl.ClusterID { return BasicID }
func (a *Unknown2) Value() *Unknown2      { return a }
func (a Unknown2) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *Unknown2) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = Unknown2(*nt)
	return br, err
}

func (a Unknown2) Readable() bool   { return true }
func (a Unknown2) Writable() bool   { return false }
func (a Unknown2) Reportable() bool { return false }
func (a Unknown2) SceneIndex() int  { return -1 }

func (a Unknown2) String() string {

	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

const ProductCodeAttr zcl.AttrID = 10

type ProductCode zcl.Zostring

func (a ProductCode) ID() zcl.AttrID         { return ProductCodeAttr }
func (a ProductCode) Cluster() zcl.ClusterID { return BasicID }
func (a *ProductCode) Value() *ProductCode   { return a }
func (a ProductCode) MarshalZcl() ([]byte, error) {
	return zcl.Zostring(a).MarshalZcl()
}
func (a *ProductCode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*a = ProductCode(*nt)
	return br, err
}

func (a ProductCode) Readable() bool   { return true }
func (a ProductCode) Writable() bool   { return false }
func (a ProductCode) Reportable() bool { return false }
func (a ProductCode) SceneIndex() int  { return -1 }

func (a ProductCode) String() string {

	return zcl.Sprintf("%s", zcl.Zostring(a))
}

const XiaomiSensitivityAttr zcl.AttrID = 65293

type XiaomiSensitivity zcl.Zu8

func (a XiaomiSensitivity) ID() zcl.AttrID             { return XiaomiSensitivityAttr }
func (a XiaomiSensitivity) Cluster() zcl.ClusterID     { return BasicID }
func (a *XiaomiSensitivity) Value() *XiaomiSensitivity { return a }
func (a XiaomiSensitivity) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *XiaomiSensitivity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = XiaomiSensitivity(*nt)
	return br, err
}

func (a XiaomiSensitivity) Readable() bool   { return true }
func (a XiaomiSensitivity) Writable() bool   { return true }
func (a XiaomiSensitivity) Reportable() bool { return false }
func (a XiaomiSensitivity) SceneIndex() int  { return -1 }

func (a XiaomiSensitivity) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const XiaomiDisconnect1Attr zcl.AttrID = 65314

type XiaomiDisconnect1 zcl.Zu8

func (a XiaomiDisconnect1) ID() zcl.AttrID             { return XiaomiDisconnect1Attr }
func (a XiaomiDisconnect1) Cluster() zcl.ClusterID     { return BasicID }
func (a *XiaomiDisconnect1) Value() *XiaomiDisconnect1 { return a }
func (a XiaomiDisconnect1) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *XiaomiDisconnect1) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = XiaomiDisconnect1(*nt)
	return br, err
}

func (a XiaomiDisconnect1) Readable() bool   { return true }
func (a XiaomiDisconnect1) Writable() bool   { return true }
func (a XiaomiDisconnect1) Reportable() bool { return false }
func (a XiaomiDisconnect1) SceneIndex() int  { return -1 }

func (a XiaomiDisconnect1) String() string {

	return zcl.Sprintf("0x%X", zcl.Zu8(a))
}

const XiaomiDisconnect2Attr zcl.AttrID = 65315

type XiaomiDisconnect2 zcl.Zu8

func (a XiaomiDisconnect2) ID() zcl.AttrID             { return XiaomiDisconnect2Attr }
func (a XiaomiDisconnect2) Cluster() zcl.ClusterID     { return BasicID }
func (a *XiaomiDisconnect2) Value() *XiaomiDisconnect2 { return a }
func (a XiaomiDisconnect2) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *XiaomiDisconnect2) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = XiaomiDisconnect2(*nt)
	return br, err
}

func (a XiaomiDisconnect2) Readable() bool   { return true }
func (a XiaomiDisconnect2) Writable() bool   { return true }
func (a XiaomiDisconnect2) Reportable() bool { return false }
func (a XiaomiDisconnect2) SceneIndex() int  { return -1 }

func (a XiaomiDisconnect2) String() string {

	return zcl.Sprintf("0x%X", zcl.Zu8(a))
}

const SensitivityAttr zcl.AttrID = 48

type Sensitivity zcl.Zenum8

func (a Sensitivity) ID() zcl.AttrID         { return SensitivityAttr }
func (a Sensitivity) Cluster() zcl.ClusterID { return BasicID }
func (a *Sensitivity) Value() *Sensitivity   { return a }
func (a Sensitivity) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *Sensitivity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = Sensitivity(*nt)
	return br, err
}

func (a Sensitivity) Readable() bool   { return true }
func (a Sensitivity) Writable() bool   { return true }
func (a Sensitivity) Reportable() bool { return false }
func (a Sensitivity) SceneIndex() int  { return -1 }

func (a Sensitivity) String() string {
	switch a {
	case 0x00:
		return "default"
	case 0x01:
		return "high"
	case 0x02:
		return "max"
	}

	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsDefault checks if Sensitivity equals the value for default (0x00)
func (a Sensitivity) IsDefault() bool { return a == 0x00 }

// SetDefault sets Sensitivity to default (0x00)
func (a *Sensitivity) SetDefault() { *a = 0x00 }

// IsHigh checks if Sensitivity equals the value for high (0x01)
func (a Sensitivity) IsHigh() bool { return a == 0x01 }

// SetHigh sets Sensitivity to high (0x01)
func (a *Sensitivity) SetHigh() { *a = 0x01 }

// IsMax checks if Sensitivity equals the value for max (0x02)
func (a Sensitivity) IsMax() bool { return a == 0x02 }

// SetMax sets Sensitivity to max (0x02)
func (a *Sensitivity) SetMax() { *a = 0x02 }

const ConfigurationAttr zcl.AttrID = 49

type Configuration zcl.Zbmp16

func (a Configuration) ID() zcl.AttrID         { return ConfigurationAttr }
func (a Configuration) Cluster() zcl.ClusterID { return BasicID }
func (a *Configuration) Value() *Configuration { return a }
func (a Configuration) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp16(a).MarshalZcl()
}
func (a *Configuration) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*a = Configuration(*nt)
	return br, err
}

func (a Configuration) Readable() bool   { return true }
func (a Configuration) Writable() bool   { return true }
func (a Configuration) Reportable() bool { return false }
func (a Configuration) SceneIndex() int  { return -1 }

func (a Configuration) String() string {

	var bstr []string
	if a.IsTouchlinkEnabled0() {
		bstr = append(bstr, "Touchlink enabled 0")
	}
	if a.IsTouchlinkEnabled1() {
		bstr = append(bstr, "Touchlink enabled 1")
	}
	if a.IsTouchlinkEnabled2() {
		bstr = append(bstr, "Touchlink enabled 2")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a Configuration) IsTouchlinkEnabled0() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *Configuration) SetTouchlinkEnabled0(b bool) {
	*a = Configuration(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a Configuration) IsTouchlinkEnabled1() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *Configuration) SetTouchlinkEnabled1(b bool) {
	*a = Configuration(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a Configuration) IsTouchlinkEnabled2() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *Configuration) SetTouchlinkEnabled2(b bool) {
	*a = Configuration(zcl.BitmapSet([]byte(*a), 3, b))
}

const UsertestAttr zcl.AttrID = 50

type Usertest zcl.Zbool

func (a Usertest) ID() zcl.AttrID         { return UsertestAttr }
func (a Usertest) Cluster() zcl.ClusterID { return BasicID }
func (a *Usertest) Value() *Usertest      { return a }
func (a Usertest) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *Usertest) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = Usertest(*nt)
	return br, err
}

func (a Usertest) Readable() bool   { return true }
func (a Usertest) Writable() bool   { return true }
func (a Usertest) Reportable() bool { return false }
func (a Usertest) SceneIndex() int  { return -1 }

func (a Usertest) String() string {

	return zcl.Sprintf("%s", zcl.Zbool(a))
}

const LedIndicationAttr zcl.AttrID = 51

type LedIndication zcl.Zbool

func (a LedIndication) ID() zcl.AttrID         { return LedIndicationAttr }
func (a LedIndication) Cluster() zcl.ClusterID { return BasicID }
func (a *LedIndication) Value() *LedIndication { return a }
func (a LedIndication) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *LedIndication) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = LedIndication(*nt)
	return br, err
}

func (a LedIndication) Readable() bool   { return true }
func (a LedIndication) Writable() bool   { return true }
func (a LedIndication) Reportable() bool { return false }
func (a LedIndication) SceneIndex() int  { return -1 }

func (a LedIndication) String() string {

	return zcl.Sprintf("%s", zcl.Zbool(a))
}
