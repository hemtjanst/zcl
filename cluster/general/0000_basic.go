package general

import "hemtjan.st/zcl"

// Basic
const BasicID zcl.ClusterID = 0

var BasicCluster = zcl.Cluster{
	Name: "Basic",
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
		GenericDeviceClassAttr:  func() zcl.Attr { return new(GenericDeviceClass) },
		GenericDeviceTypeAttr:   func() zcl.Attr { return new(GenericDeviceType) },
		LocationDescriptionAttr: func() zcl.Attr { return new(LocationDescription) },
		PhysicalEnvironmentAttr: func() zcl.Attr { return new(PhysicalEnvironment) },
		DeviceEnabledAttr:       func() zcl.Attr { return new(DeviceEnabled) },
		AlarmMaskAttr:           func() zcl.Attr { return new(AlarmMask) },
		DisableLocalConfigAttr:  func() zcl.Attr { return new(DisableLocalConfig) },
		SwBuildIdAttr:           func() zcl.Attr { return new(SwBuildId) },
		ProductCodeAttr:         func() zcl.Attr { return new(ProductCode) },
		ProductUrlAttr:          func() zcl.Attr { return new(ProductUrl) },
		SensitivityAttr:         func() zcl.Attr { return new(Sensitivity) },
		ConfigurationAttr:       func() zcl.Attr { return new(Configuration) },
		UserTestAttr:            func() zcl.Attr { return new(UserTest) },
		LedIndicationAttr:       func() zcl.Attr { return new(LedIndication) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

type ResetToFactoryDefaults struct {
}

// ResetToFactoryDefaultsCommand is the Command ID of ResetToFactoryDefaults
const ResetToFactoryDefaultsCommand CommandID = 0x0000

// Values returns all values of ResetToFactoryDefaults
func (v *ResetToFactoryDefaults) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of ResetToFactoryDefaults
func (v *ResetToFactoryDefaults) Arguments() []zcl.Argument {
	return []zcl.Argument{}
}

// Name of the command
func (ResetToFactoryDefaults) Name() string { return "Reset to Factory Defaults" }

// ID of the command
func (ResetToFactoryDefaults) ID() CommandID { return ResetToFactoryDefaultsCommand }

// Required
func (ResetToFactoryDefaults) Required() bool { return false }

// Cluster ID of the command
func (ResetToFactoryDefaults) Cluster() zcl.ClusterID { return BasicID }

// MnfCode returns the manufacturer code (if any) of the command
func (ResetToFactoryDefaults) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (ResetToFactoryDefaults) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

// MarshalZcl returns the wire format representation of ResetToFactoryDefaults
func (v ResetToFactoryDefaults) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the ResetToFactoryDefaults struct
func (v *ResetToFactoryDefaults) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ResetToFactoryDefaults) String() string {
	return zcl.Sprintf(
		"ResetToFactoryDefaults{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}
