package zdo

import "hemtjan.st/zcl"

type CommandID = zcl.ZdoCmdID

// ActiveEndpointSize Size in bytes of the Active Endpoints List
type ActiveEndpointSize zcl.Zu8

func (a *ActiveEndpointSize) Value() *ActiveEndpointSize { return a }
func (a ActiveEndpointSize) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *ActiveEndpointSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ActiveEndpointSize(*nt)
	return br, err
}

func (a *ActiveEndpointSize) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = ActiveEndpointSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ActiveEndpointSize) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type AssociatedDevices zcl.Zlist

func (a *AssociatedDevices) Value() *AssociatedDevices { return a }
func (a AssociatedDevices) MarshalZcl() ([]byte, error) {
	return nil, zcl.Errorf("not implemented")
}

func (a *AssociatedDevices) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zlist)
	br, err := nt.UnmarshalZcl(b)
	*a = AssociatedDevices(*nt)
	return br, err
}

func (a *AssociatedDevices) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zlist); ok {
		*a = AssociatedDevices(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a AssociatedDevices) String() string {
	return zcl.Sprintf("%v", zcl.Zlist(a))
}

// BindingTarget NWK Address
type BindingTarget zcl.Zu16

func (a *BindingTarget) Value() *BindingTarget      { return a }
func (a BindingTarget) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *BindingTarget) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = BindingTarget(*nt)
	return br, err
}

func (a *BindingTarget) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = BindingTarget(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BindingTarget) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// Capability specifies the device:s capabilities
type Capability zcl.Zbmp8

func (a *Capability) Value() *Capability         { return a }
func (a Capability) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *Capability) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = Capability(*nt)
	return br, err
}

func (a *Capability) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = Capability(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Capability) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Pan Coordinator")
		case 1:
			bstr = append(bstr, "Full Function")
		case 2:
			bstr = append(bstr, "Mains Power")
		case 3:
			bstr = append(bstr, "RX on Idle")
		case 6:
			bstr = append(bstr, "Security")
		case 7:
			bstr = append(bstr, "Allocate Address")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a Capability) IsPanCoordinator() bool     { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a Capability) IsFullFunction() bool       { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a Capability) IsMainsPower() bool         { return zcl.BitmapTest([]byte(a[:]), 2) }
func (a Capability) IsRxOnIdle() bool           { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a Capability) IsSecurity() bool           { return zcl.BitmapTest([]byte(a[:]), 6) }
func (a Capability) IsAllocateAddress() bool    { return zcl.BitmapTest([]byte(a[:]), 7) }
func (a *Capability) SetPanCoordinator(b bool)  { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b)) }
func (a *Capability) SetFullFunction(b bool)    { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b)) }
func (a *Capability) SetMainsPower(b bool)      { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b)) }
func (a *Capability) SetRxOnIdle(b bool)        { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b)) }
func (a *Capability) SetSecurity(b bool)        { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 6, b)) }
func (a *Capability) SetAllocateAddress(b bool) { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 7, b)) }

type CurrentPowerSource zcl.Zenum8

func (a *CurrentPowerSource) Value() *CurrentPowerSource { return a }
func (a CurrentPowerSource) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *CurrentPowerSource) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentPowerSource(*nt)
	return br, err
}

func (a *CurrentPowerSource) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = CurrentPowerSource(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CurrentPowerSource) String() string {
	switch a {
	case 0x00:
		return "Constant Power (Critical)"
	case 0x04:
		return "Constant Power (33%)"
	case 0x08:
		return "Constant Power (66%)"
	case 0x0B:
		return "Constant Power (100%)"
	case 0x10:
		return "Rech. battery (Critical)"
	case 0x14:
		return "Rech. battery (33%)"
	case 0x18:
		return "Rech. battery (66%)"
	case 0x1B:
		return "Rech. battery (100%)"
	case 0x20:
		return "Battery (Critical)"
	case 0x24:
		return "Battery (33%)"
	case 0x28:
		return "Battery (66%)"
	case 0x2B:
		return "Battery (100%)"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a CurrentPowerSource) IsConstantPowerCritical() bool { return a == 0x00 }
func (a CurrentPowerSource) IsConstantPower33() bool       { return a == 0x04 }
func (a CurrentPowerSource) IsConstantPower66() bool       { return a == 0x08 }
func (a CurrentPowerSource) IsConstantPower100() bool      { return a == 0x0B }
func (a CurrentPowerSource) IsRechBatteryCritical() bool   { return a == 0x10 }
func (a CurrentPowerSource) IsRechBattery33() bool         { return a == 0x14 }
func (a CurrentPowerSource) IsRechBattery66() bool         { return a == 0x18 }
func (a CurrentPowerSource) IsRechBattery100() bool        { return a == 0x1B }
func (a CurrentPowerSource) IsBatteryCritical() bool       { return a == 0x20 }
func (a CurrentPowerSource) IsBattery33() bool             { return a == 0x24 }
func (a CurrentPowerSource) IsBattery66() bool             { return a == 0x28 }
func (a CurrentPowerSource) IsBattery100() bool            { return a == 0x2B }
func (a *CurrentPowerSource) SetConstantPowerCritical()    { *a = 0x00 }
func (a *CurrentPowerSource) SetConstantPower33()          { *a = 0x04 }
func (a *CurrentPowerSource) SetConstantPower66()          { *a = 0x08 }
func (a *CurrentPowerSource) SetConstantPower100()         { *a = 0x0B }
func (a *CurrentPowerSource) SetRechBatteryCritical()      { *a = 0x10 }
func (a *CurrentPowerSource) SetRechBattery33()            { *a = 0x14 }
func (a *CurrentPowerSource) SetRechBattery66()            { *a = 0x18 }
func (a *CurrentPowerSource) SetRechBattery100()           { *a = 0x1B }
func (a *CurrentPowerSource) SetBatteryCritical()          { *a = 0x20 }
func (a *CurrentPowerSource) SetBattery33()                { *a = 0x24 }
func (a *CurrentPowerSource) SetBattery66()                { *a = 0x28 }
func (a *CurrentPowerSource) SetBattery100()               { *a = 0x2B }

type DescriptorCapability zcl.Zbmp8

func (a *DescriptorCapability) Value() *DescriptorCapability { return a }
func (a DescriptorCapability) MarshalZcl() ([]byte, error)   { return zcl.Zbmp8(a).MarshalZcl() }

func (a *DescriptorCapability) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = DescriptorCapability(*nt)
	return br, err
}

func (a *DescriptorCapability) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = DescriptorCapability(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DescriptorCapability) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Extended Active Endpoint List Available")
		case 1:
			bstr = append(bstr, "Extended Simple Descriptor List Available")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a DescriptorCapability) IsExtendedActiveEndpointListAvailable() bool {
	return zcl.BitmapTest([]byte(a[:]), 0)
}
func (a DescriptorCapability) IsExtendedSimpleDescriptorListAvailable() bool {
	return zcl.BitmapTest([]byte(a[:]), 1)
}
func (a *DescriptorCapability) SetExtendedActiveEndpointListAvailable(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *DescriptorCapability) SetExtendedSimpleDescriptorListAvailable(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}

type DeviceType zcl.Zenum16

func (a *DeviceType) Value() *DeviceType         { return a }
func (a DeviceType) MarshalZcl() ([]byte, error) { return zcl.Zenum16(a).MarshalZcl() }

func (a *DeviceType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum16)
	br, err := nt.UnmarshalZcl(b)
	*a = DeviceType(*nt)
	return br, err
}

func (a *DeviceType) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum16); ok {
		*a = DeviceType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DeviceType) String() string {
	switch a {
	case 0x0000:
		return "On/Off Switch"
	case 0x0001:
		return "Level Control Switch"
	case 0x0002:
		return "On/Off Output"
	case 0x0003:
		return "Level Controllable Output"
	case 0x0004:
		return "Scene Selector"
	case 0x0005:
		return "Configuration Tool"
	case 0x0006:
		return "Remote Control"
	case 0x0007:
		return "Combined Interface"
	case 0x0008:
		return "Range Extender"
	case 0x0009:
		return "Mains Power Outlet"
	case 0x000a:
		return "Door lock"
	case 0x000c:
		return "Simple sensor"
	case 0x0010:
		return "On/off plug-in unit"
	case 0x0051:
		return "Smart plug"
	case 0x0060:
		return "GP Proxy"
	case 0x0061:
		return "GP Proxy Basic"
	case 0x0062:
		return "GP Target Plus"
	case 0x0063:
		return "GP Target"
	case 0x0064:
		return "GP Commissioning Tool"
	case 0x0065:
		return "GP Combo"
	case 0x0066:
		return "GP Combo Basic"
	case 0x0100:
		return "On/Off Light"
	case 0x0101:
		return "Dimmable Light"
	case 0x0102:
		return "Color Dimmable Light"
	case 0x0103:
		return "On/Off Light Switch"
	case 0x0104:
		return "Dimmer Switch"
	case 0x0105:
		return "Color Dimmer Switch"
	case 0x0106:
		return "Light Sensor"
	case 0x0107:
		return "Occupancy Sensor"
	case 0x0108:
		return "On/off ballast"
	case 0x0109:
		return "Dimmable ballast"
	case 0x010a:
		return "On/off plugin unit"
	case 0x010b:
		return "Dimmable plugin unit"
	case 0x010c:
		return "Color temperature light"
	case 0x010d:
		return "Extended color light"
	case 0x010e:
		return "Light level sensor"
	case 0x0110:
		return "Dimmable plug-in unit"
	case 0x0200:
		return "Shade"
	case 0x0201:
		return "Shade Controller"
	case 0x0202:
		return "Window Covering Device"
	case 0x0210:
		return "Extended color light 2"
	case 0x0220:
		return "Color temperature light 2"
	case 0x0300:
		return "Heating/Cooling Unit"
	case 0x0301:
		return "Thermostat"
	case 0x0302:
		return "Temperature Sensor"
	case 0x0303:
		return "Pump"
	case 0x0304:
		return "Pump Controller"
	case 0x0305:
		return "Pressure Sensor"
	case 0x0306:
		return "Flow Sensor"
	case 0x0400:
		return "IAS Control and Indicating Equipment"
	case 0x0401:
		return "IAS Ancillary Control Equipment"
	case 0x0402:
		return "IAS Zone"
	case 0x0403:
		return "IAS Warning Device"
	case 0x0500:
		return "Energy Service Portal"
	case 0x0501:
		return "Metering Device"
	case 0x0502:
		return "In-Premise Display"
	case 0x0503:
		return "Programmable Communicating Thermostat (PCT)"
	case 0x0504:
		return "Load Control Device"
	case 0x0505:
		return "Smart Appliance"
	case 0x0506:
		return "Prepayment Terminal"
	case 0x0507:
		return "Device Management"
	case 0x0800:
		return "Color controller"
	case 0x0810:
		return "Color scene controller"
	case 0x0820:
		return "Non color controller"
	case 0x0830:
		return "Non color scene controller"
	case 0x0840:
		return "Control bridge"
	case 0x0850:
		return "On/off sensor"
	}
	return zcl.Sprintf("%v", zcl.Zenum16(a))
}

func (a DeviceType) IsOnOffSwitch() bool                            { return a == 0x0000 }
func (a DeviceType) IsLevelControlSwitch() bool                     { return a == 0x0001 }
func (a DeviceType) IsOnOffOutput() bool                            { return a == 0x0002 }
func (a DeviceType) IsLevelControllableOutput() bool                { return a == 0x0003 }
func (a DeviceType) IsSceneSelector() bool                          { return a == 0x0004 }
func (a DeviceType) IsConfigurationTool() bool                      { return a == 0x0005 }
func (a DeviceType) IsRemoteControl() bool                          { return a == 0x0006 }
func (a DeviceType) IsCombinedInterface() bool                      { return a == 0x0007 }
func (a DeviceType) IsRangeExtender() bool                          { return a == 0x0008 }
func (a DeviceType) IsMainsPowerOutlet() bool                       { return a == 0x0009 }
func (a DeviceType) IsDoorLock() bool                               { return a == 0x000a }
func (a DeviceType) IsSimpleSensor() bool                           { return a == 0x000c }
func (a DeviceType) IsOnOffPlugInUnit() bool                        { return a == 0x0010 }
func (a DeviceType) IsSmartPlug() bool                              { return a == 0x0051 }
func (a DeviceType) IsGpProxy() bool                                { return a == 0x0060 }
func (a DeviceType) IsGpProxyBasic() bool                           { return a == 0x0061 }
func (a DeviceType) IsGpTargetPlus() bool                           { return a == 0x0062 }
func (a DeviceType) IsGpTarget() bool                               { return a == 0x0063 }
func (a DeviceType) IsGpCommissioningTool() bool                    { return a == 0x0064 }
func (a DeviceType) IsGpCombo() bool                                { return a == 0x0065 }
func (a DeviceType) IsGpComboBasic() bool                           { return a == 0x0066 }
func (a DeviceType) IsOnOffLight() bool                             { return a == 0x0100 }
func (a DeviceType) IsDimmableLight() bool                          { return a == 0x0101 }
func (a DeviceType) IsColorDimmableLight() bool                     { return a == 0x0102 }
func (a DeviceType) IsOnOffLightSwitch() bool                       { return a == 0x0103 }
func (a DeviceType) IsDimmerSwitch() bool                           { return a == 0x0104 }
func (a DeviceType) IsColorDimmerSwitch() bool                      { return a == 0x0105 }
func (a DeviceType) IsLightSensor() bool                            { return a == 0x0106 }
func (a DeviceType) IsOccupancySensor() bool                        { return a == 0x0107 }
func (a DeviceType) IsOnOffBallast() bool                           { return a == 0x0108 }
func (a DeviceType) IsDimmableBallast() bool                        { return a == 0x0109 }
func (a DeviceType) IsOnOffPluginUnit() bool                        { return a == 0x010a }
func (a DeviceType) IsDimmablePluginUnit() bool                     { return a == 0x010b }
func (a DeviceType) IsColorTemperatureLight() bool                  { return a == 0x010c }
func (a DeviceType) IsExtendedColorLight() bool                     { return a == 0x010d }
func (a DeviceType) IsLightLevelSensor() bool                       { return a == 0x010e }
func (a DeviceType) IsDimmablePlugInUnit() bool                     { return a == 0x0110 }
func (a DeviceType) IsShade() bool                                  { return a == 0x0200 }
func (a DeviceType) IsShadeController() bool                        { return a == 0x0201 }
func (a DeviceType) IsWindowCoveringDevice() bool                   { return a == 0x0202 }
func (a DeviceType) IsExtendedColorLight2() bool                    { return a == 0x0210 }
func (a DeviceType) IsColorTemperatureLight2() bool                 { return a == 0x0220 }
func (a DeviceType) IsHeatingCoolingUnit() bool                     { return a == 0x0300 }
func (a DeviceType) IsThermostat() bool                             { return a == 0x0301 }
func (a DeviceType) IsTemperatureSensor() bool                      { return a == 0x0302 }
func (a DeviceType) IsPump() bool                                   { return a == 0x0303 }
func (a DeviceType) IsPumpController() bool                         { return a == 0x0304 }
func (a DeviceType) IsPressureSensor() bool                         { return a == 0x0305 }
func (a DeviceType) IsFlowSensor() bool                             { return a == 0x0306 }
func (a DeviceType) IsIasControlAndIndicatingEquipment() bool       { return a == 0x0400 }
func (a DeviceType) IsIasAncillaryControlEquipment() bool           { return a == 0x0401 }
func (a DeviceType) IsIasZone() bool                                { return a == 0x0402 }
func (a DeviceType) IsIasWarningDevice() bool                       { return a == 0x0403 }
func (a DeviceType) IsEnergyServicePortal() bool                    { return a == 0x0500 }
func (a DeviceType) IsMeteringDevice() bool                         { return a == 0x0501 }
func (a DeviceType) IsInPremiseDisplay() bool                       { return a == 0x0502 }
func (a DeviceType) IsProgrammableCommunicatingThermostatPct() bool { return a == 0x0503 }
func (a DeviceType) IsLoadControlDevice() bool                      { return a == 0x0504 }
func (a DeviceType) IsSmartAppliance() bool                         { return a == 0x0505 }
func (a DeviceType) IsPrepaymentTerminal() bool                     { return a == 0x0506 }
func (a DeviceType) IsDeviceManagement() bool                       { return a == 0x0507 }
func (a DeviceType) IsColorController() bool                        { return a == 0x0800 }
func (a DeviceType) IsColorSceneController() bool                   { return a == 0x0810 }
func (a DeviceType) IsNonColorController() bool                     { return a == 0x0820 }
func (a DeviceType) IsNonColorSceneController() bool                { return a == 0x0830 }
func (a DeviceType) IsControlBridge() bool                          { return a == 0x0840 }
func (a DeviceType) IsOnOffSensor() bool                            { return a == 0x0850 }
func (a *DeviceType) SetOnOffSwitch()                               { *a = 0x0000 }
func (a *DeviceType) SetLevelControlSwitch()                        { *a = 0x0001 }
func (a *DeviceType) SetOnOffOutput()                               { *a = 0x0002 }
func (a *DeviceType) SetLevelControllableOutput()                   { *a = 0x0003 }
func (a *DeviceType) SetSceneSelector()                             { *a = 0x0004 }
func (a *DeviceType) SetConfigurationTool()                         { *a = 0x0005 }
func (a *DeviceType) SetRemoteControl()                             { *a = 0x0006 }
func (a *DeviceType) SetCombinedInterface()                         { *a = 0x0007 }
func (a *DeviceType) SetRangeExtender()                             { *a = 0x0008 }
func (a *DeviceType) SetMainsPowerOutlet()                          { *a = 0x0009 }
func (a *DeviceType) SetDoorLock()                                  { *a = 0x000a }
func (a *DeviceType) SetSimpleSensor()                              { *a = 0x000c }
func (a *DeviceType) SetOnOffPlugInUnit()                           { *a = 0x0010 }
func (a *DeviceType) SetSmartPlug()                                 { *a = 0x0051 }
func (a *DeviceType) SetGpProxy()                                   { *a = 0x0060 }
func (a *DeviceType) SetGpProxyBasic()                              { *a = 0x0061 }
func (a *DeviceType) SetGpTargetPlus()                              { *a = 0x0062 }
func (a *DeviceType) SetGpTarget()                                  { *a = 0x0063 }
func (a *DeviceType) SetGpCommissioningTool()                       { *a = 0x0064 }
func (a *DeviceType) SetGpCombo()                                   { *a = 0x0065 }
func (a *DeviceType) SetGpComboBasic()                              { *a = 0x0066 }
func (a *DeviceType) SetOnOffLight()                                { *a = 0x0100 }
func (a *DeviceType) SetDimmableLight()                             { *a = 0x0101 }
func (a *DeviceType) SetColorDimmableLight()                        { *a = 0x0102 }
func (a *DeviceType) SetOnOffLightSwitch()                          { *a = 0x0103 }
func (a *DeviceType) SetDimmerSwitch()                              { *a = 0x0104 }
func (a *DeviceType) SetColorDimmerSwitch()                         { *a = 0x0105 }
func (a *DeviceType) SetLightSensor()                               { *a = 0x0106 }
func (a *DeviceType) SetOccupancySensor()                           { *a = 0x0107 }
func (a *DeviceType) SetOnOffBallast()                              { *a = 0x0108 }
func (a *DeviceType) SetDimmableBallast()                           { *a = 0x0109 }
func (a *DeviceType) SetOnOffPluginUnit()                           { *a = 0x010a }
func (a *DeviceType) SetDimmablePluginUnit()                        { *a = 0x010b }
func (a *DeviceType) SetColorTemperatureLight()                     { *a = 0x010c }
func (a *DeviceType) SetExtendedColorLight()                        { *a = 0x010d }
func (a *DeviceType) SetLightLevelSensor()                          { *a = 0x010e }
func (a *DeviceType) SetDimmablePlugInUnit()                        { *a = 0x0110 }
func (a *DeviceType) SetShade()                                     { *a = 0x0200 }
func (a *DeviceType) SetShadeController()                           { *a = 0x0201 }
func (a *DeviceType) SetWindowCoveringDevice()                      { *a = 0x0202 }
func (a *DeviceType) SetExtendedColorLight2()                       { *a = 0x0210 }
func (a *DeviceType) SetColorTemperatureLight2()                    { *a = 0x0220 }
func (a *DeviceType) SetHeatingCoolingUnit()                        { *a = 0x0300 }
func (a *DeviceType) SetThermostat()                                { *a = 0x0301 }
func (a *DeviceType) SetTemperatureSensor()                         { *a = 0x0302 }
func (a *DeviceType) SetPump()                                      { *a = 0x0303 }
func (a *DeviceType) SetPumpController()                            { *a = 0x0304 }
func (a *DeviceType) SetPressureSensor()                            { *a = 0x0305 }
func (a *DeviceType) SetFlowSensor()                                { *a = 0x0306 }
func (a *DeviceType) SetIasControlAndIndicatingEquipment()          { *a = 0x0400 }
func (a *DeviceType) SetIasAncillaryControlEquipment()              { *a = 0x0401 }
func (a *DeviceType) SetIasZone()                                   { *a = 0x0402 }
func (a *DeviceType) SetIasWarningDevice()                          { *a = 0x0403 }
func (a *DeviceType) SetEnergyServicePortal()                       { *a = 0x0500 }
func (a *DeviceType) SetMeteringDevice()                            { *a = 0x0501 }
func (a *DeviceType) SetInPremiseDisplay()                          { *a = 0x0502 }
func (a *DeviceType) SetProgrammableCommunicatingThermostatPct()    { *a = 0x0503 }
func (a *DeviceType) SetLoadControlDevice()                         { *a = 0x0504 }
func (a *DeviceType) SetSmartAppliance()                            { *a = 0x0505 }
func (a *DeviceType) SetPrepaymentTerminal()                        { *a = 0x0506 }
func (a *DeviceType) SetDeviceManagement()                          { *a = 0x0507 }
func (a *DeviceType) SetColorController()                           { *a = 0x0800 }
func (a *DeviceType) SetColorSceneController()                      { *a = 0x0810 }
func (a *DeviceType) SetNonColorController()                        { *a = 0x0820 }
func (a *DeviceType) SetNonColorSceneController()                   { *a = 0x0830 }
func (a *DeviceType) SetControlBridge()                             { *a = 0x0840 }
func (a *DeviceType) SetOnOffSensor()                               { *a = 0x0850 }

// DeviceVersion is dependant on profile
type DeviceVersion zcl.Zu8

func (a *DeviceVersion) Value() *DeviceVersion      { return a }
func (a DeviceVersion) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *DeviceVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = DeviceVersion(*nt)
	return br, err
}

func (a *DeviceVersion) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = DeviceVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DeviceVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type Endpoint zcl.Zu8

func (a *Endpoint) Value() *Endpoint           { return a }
func (a Endpoint) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Endpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Endpoint(*nt)
	return br, err
}

func (a *Endpoint) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = Endpoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Endpoint) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type EndpointList zcl.Zset

func (a *EndpointList) Value() *EndpointList { return a }
func (a EndpointList) MarshalZcl() ([]byte, error) {
	return nil, zcl.Errorf("not implemented")
}

func (a *EndpointList) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zset)
	br, err := nt.UnmarshalZcl(b)
	*a = EndpointList(*nt)
	return br, err
}

func (a *EndpointList) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zset); ok {
		*a = EndpointList(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a EndpointList) String() string {
	return zcl.Sprintf("%v", zcl.Zset(a))
}

type FrequencyBands zcl.Zbmp8

func (a *FrequencyBands) Value() *FrequencyBands     { return a }
func (a FrequencyBands) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *FrequencyBands) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = FrequencyBands(*nt)
	return br, err
}

func (a *FrequencyBands) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = FrequencyBands(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a FrequencyBands) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "868.0-868.6 MHz")
		case 2:
			bstr = append(bstr, "902.0-928 MHz")
		case 3:
			bstr = append(bstr, "2400.0-2483.5 MHz")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a FrequencyBands) Is86808686Mhz() bool      { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a FrequencyBands) Is9020928Mhz() bool       { return zcl.BitmapTest([]byte(a[:]), 2) }
func (a FrequencyBands) Is2400024835Mhz() bool    { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a *FrequencyBands) Set86808686Mhz(b bool)   { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b)) }
func (a *FrequencyBands) Set9020928Mhz(b bool)    { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b)) }
func (a *FrequencyBands) Set2400024835Mhz(b bool) { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b)) }

// IeeeAddress is a 64-bit MAC address
type IeeeAddress zcl.Zuid

func (a *IeeeAddress) Value() *IeeeAddress        { return a }
func (a IeeeAddress) MarshalZcl() ([]byte, error) { return zcl.Zuid(a).MarshalZcl() }

func (a *IeeeAddress) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*a = IeeeAddress(*nt)
	return br, err
}

func (a *IeeeAddress) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zuid); ok {
		*a = IeeeAddress(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a IeeeAddress) String() string {
	return zcl.Sprintf("%v", zcl.Zuid(a))
}

// InClusterList is a list of input clusters
type InClusterList zcl.Zset

func (a *InClusterList) Value() *InClusterList { return a }
func (a InClusterList) MarshalZcl() ([]byte, error) {
	return nil, zcl.Errorf("not implemented")
}

func (a *InClusterList) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zset)
	br, err := nt.UnmarshalZcl(b)
	*a = InClusterList(*nt)
	return br, err
}

func (a *InClusterList) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zset); ok {
		*a = InClusterList(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a InClusterList) String() string {
	return zcl.Sprintf("%v", zcl.Zset(a))
}

type ManufacturerCode zcl.Zenum16

func (a *ManufacturerCode) Value() *ManufacturerCode   { return a }
func (a ManufacturerCode) MarshalZcl() ([]byte, error) { return zcl.Zenum16(a).MarshalZcl() }

func (a *ManufacturerCode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum16)
	br, err := nt.UnmarshalZcl(b)
	*a = ManufacturerCode(*nt)
	return br, err
}

func (a *ManufacturerCode) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum16); ok {
		*a = ManufacturerCode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ManufacturerCode) String() string {
	switch a {
	case 0x100B:
		return "Philips"
	case 0x1021:
		return "Legrand"
	case 0x1037:
		return "Eurotronic"
	case 0x10F2:
		return "Ubisys"
	case 0x115F:
		return "Xiaomi"
	case 0x117C:
		return "Ikea"
	}
	return zcl.Sprintf("%v", zcl.Zenum16(a))
}

func (a ManufacturerCode) IsPhilips() bool    { return a == 0x100B }
func (a ManufacturerCode) IsLegrand() bool    { return a == 0x1021 }
func (a ManufacturerCode) IsEurotronic() bool { return a == 0x1037 }
func (a ManufacturerCode) IsUbisys() bool     { return a == 0x10F2 }
func (a ManufacturerCode) IsXiaomi() bool     { return a == 0x115F }
func (a ManufacturerCode) IsIkea() bool       { return a == 0x117C }
func (a *ManufacturerCode) SetPhilips()       { *a = 0x100B }
func (a *ManufacturerCode) SetLegrand()       { *a = 0x1021 }
func (a *ManufacturerCode) SetEurotronic()    { *a = 0x1037 }
func (a *ManufacturerCode) SetUbisys()        { *a = 0x10F2 }
func (a *ManufacturerCode) SetXiaomi()        { *a = 0x115F }
func (a *ManufacturerCode) SetIkea()          { *a = 0x117C }

// MaxBufferSize specifies the maximum size, in octets, of the network sub-layer data unit (NSDU) for this node.
// This is the maximum size of data or commands passed to or from the application by the application support sub-layer,
// before any fragmentation or re-assembly.
type MaxBufferSize zcl.Zu8

func (a *MaxBufferSize) Value() *MaxBufferSize      { return a }
func (a MaxBufferSize) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *MaxBufferSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxBufferSize(*nt)
	return br, err
}

func (a *MaxBufferSize) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = MaxBufferSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MaxBufferSize) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// MaxIncomingTransferSize specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred to this node in one single message transfer.
type MaxIncomingTransferSize zcl.Zu16

func (a *MaxIncomingTransferSize) Value() *MaxIncomingTransferSize { return a }
func (a MaxIncomingTransferSize) MarshalZcl() ([]byte, error)      { return zcl.Zu16(a).MarshalZcl() }

func (a *MaxIncomingTransferSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxIncomingTransferSize(*nt)
	return br, err
}

func (a *MaxIncomingTransferSize) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = MaxIncomingTransferSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MaxIncomingTransferSize) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// MaxOutgoingTransferSize specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred from
// this node in one single message transfer.
type MaxOutgoingTransferSize zcl.Zu16

func (a *MaxOutgoingTransferSize) Value() *MaxOutgoingTransferSize { return a }
func (a MaxOutgoingTransferSize) MarshalZcl() ([]byte, error)      { return zcl.Zu16(a).MarshalZcl() }

func (a *MaxOutgoingTransferSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxOutgoingTransferSize(*nt)
	return br, err
}

func (a *MaxOutgoingTransferSize) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = MaxOutgoingTransferSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MaxOutgoingTransferSize) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// NwkAddress is a 16-bit Network address
type NwkAddress zcl.Zu16

func (a *NwkAddress) Value() *NwkAddress         { return a }
func (a NwkAddress) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *NwkAddress) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NwkAddress(*nt)
	return br, err
}

func (a *NwkAddress) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = NwkAddress(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NwkAddress) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// NodeDescSize Size in bytes of the Node Descriptor
type NodeDescSize zcl.Zu8

func (a *NodeDescSize) Value() *NodeDescSize       { return a }
func (a NodeDescSize) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *NodeDescSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NodeDescSize(*nt)
	return br, err
}

func (a *NodeDescSize) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = NodeDescSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NodeDescSize) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// OutClusterList is a list of output clusters
type OutClusterList zcl.Zset

func (a *OutClusterList) Value() *OutClusterList { return a }
func (a OutClusterList) MarshalZcl() ([]byte, error) {
	return nil, zcl.Errorf("not implemented")
}

func (a *OutClusterList) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zset)
	br, err := nt.UnmarshalZcl(b)
	*a = OutClusterList(*nt)
	return br, err
}

func (a *OutClusterList) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zset); ok {
		*a = OutClusterList(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OutClusterList) String() string {
	return zcl.Sprintf("%v", zcl.Zset(a))
}

// PowerDescSize Size in bytes of the Power Descriptor
type PowerDescSize zcl.Zu8

func (a *PowerDescSize) Value() *PowerDescSize      { return a }
func (a PowerDescSize) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *PowerDescSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerDescSize(*nt)
	return br, err
}

func (a *PowerDescSize) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = PowerDescSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PowerDescSize) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type PowerMode zcl.Zbmp8

func (a *PowerMode) Value() *PowerMode          { return a }
func (a PowerMode) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *PowerMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerMode(*nt)
	return br, err
}

func (a *PowerMode) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = PowerMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PowerMode) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Constant Power Available")
		case 1:
			bstr = append(bstr, "Rechargeable battery Available")
		case 2:
			bstr = append(bstr, "Disposable Battery Available")
		case 4:
			bstr = append(bstr, "Check In Periodically")
		case 5:
			bstr = append(bstr, "Check In on Action")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a PowerMode) IsConstantPowerAvailable() bool       { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a PowerMode) IsRechargeableBatteryAvailable() bool { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a PowerMode) IsDisposableBatteryAvailable() bool   { return zcl.BitmapTest([]byte(a[:]), 2) }
func (a PowerMode) IsCheckInPeriodically() bool          { return zcl.BitmapTest([]byte(a[:]), 4) }
func (a PowerMode) IsCheckInOnAction() bool              { return zcl.BitmapTest([]byte(a[:]), 5) }
func (a *PowerMode) SetConstantPowerAvailable(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *PowerMode) SetRechargeableBatteryAvailable(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}
func (a *PowerMode) SetDisposableBatteryAvailable(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b))
}
func (a *PowerMode) SetCheckInPeriodically(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 4, b))
}
func (a *PowerMode) SetCheckInOnAction(b bool) { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 5, b)) }

type ProfileId zcl.Zu16

func (a *ProfileId) Value() *ProfileId          { return a }
func (a ProfileId) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ProfileId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ProfileId(*nt)
	return br, err
}

func (a *ProfileId) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ProfileId(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ProfileId) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// RequestType should be set to 1 if extended response is requested, 0 otherwise
type RequestType zcl.Zenum8

func (a *RequestType) Value() *RequestType        { return a }
func (a RequestType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *RequestType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = RequestType(*nt)
	return br, err
}

func (a *RequestType) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = RequestType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a RequestType) String() string {
	switch a {
	case 0x00:
		return "Single Device Response"
	case 0x01:
		return "Extended Response"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a RequestType) IsSingleDeviceResponse() bool { return a == 0x00 }
func (a RequestType) IsExtendedResponse() bool     { return a == 0x01 }
func (a *RequestType) SetSingleDeviceResponse()    { *a = 0x00 }
func (a *RequestType) SetExtendedResponse()        { *a = 0x01 }

type ServerMask zcl.Zbmp16

func (a *ServerMask) Value() *ServerMask         { return a }
func (a ServerMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp16(a).MarshalZcl() }

func (a *ServerMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*a = ServerMask(*nt)
	return br, err
}

func (a *ServerMask) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp16); ok {
		*a = ServerMask(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ServerMask) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Primary Trust Center")
		case 1:
			bstr = append(bstr, "Backup Trust Center")
		case 2:
			bstr = append(bstr, "Primary Binding Table Cache")
		case 3:
			bstr = append(bstr, "Backup Binding Table Cache")
		case 4:
			bstr = append(bstr, "Primary Discovery Cache")
		case 5:
			bstr = append(bstr, "Backup Discovery Cache")
		case 6:
			bstr = append(bstr, "Network Manager")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a ServerMask) IsPrimaryTrustCenter() bool       { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a ServerMask) IsBackupTrustCenter() bool        { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a ServerMask) IsPrimaryBindingTableCache() bool { return zcl.BitmapTest([]byte(a[:]), 2) }
func (a ServerMask) IsBackupBindingTableCache() bool  { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a ServerMask) IsPrimaryDiscoveryCache() bool    { return zcl.BitmapTest([]byte(a[:]), 4) }
func (a ServerMask) IsBackupDiscoveryCache() bool     { return zcl.BitmapTest([]byte(a[:]), 5) }
func (a ServerMask) IsNetworkManager() bool           { return zcl.BitmapTest([]byte(a[:]), 6) }
func (a *ServerMask) SetPrimaryTrustCenter(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *ServerMask) SetBackupTrustCenter(b bool) { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b)) }
func (a *ServerMask) SetPrimaryBindingTableCache(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b))
}
func (a *ServerMask) SetBackupBindingTableCache(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b))
}
func (a *ServerMask) SetPrimaryDiscoveryCache(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 4, b))
}
func (a *ServerMask) SetBackupDiscoveryCache(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 5, b))
}
func (a *ServerMask) SetNetworkManager(b bool) { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 6, b)) }

// SimpleDescSizeList List of sizes for the different Simple Descriptors
type SimpleDescSizeList zcl.Zset

func (a *SimpleDescSizeList) Value() *SimpleDescSizeList { return a }
func (a SimpleDescSizeList) MarshalZcl() ([]byte, error) {
	return nil, zcl.Errorf("not implemented")
}

func (a *SimpleDescSizeList) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zset)
	br, err := nt.UnmarshalZcl(b)
	*a = SimpleDescSizeList(*nt)
	return br, err
}

func (a *SimpleDescSizeList) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zset); ok {
		*a = SimpleDescSizeList(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SimpleDescSizeList) String() string {
	return zcl.Sprintf("%v", zcl.Zset(a))
}

// SourceAddress of device generating the request
type SourceAddress zcl.Zuid

func (a *SourceAddress) Value() *SourceAddress      { return a }
func (a SourceAddress) MarshalZcl() ([]byte, error) { return zcl.Zuid(a).MarshalZcl() }

func (a *SourceAddress) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*a = SourceAddress(*nt)
	return br, err
}

func (a *SourceAddress) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zuid); ok {
		*a = SourceAddress(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SourceAddress) String() string {
	return zcl.Sprintf("%v", zcl.Zuid(a))
}

// SourceEndpoint of device generating the request
type SourceEndpoint zcl.Zu8

func (a *SourceEndpoint) Value() *SourceEndpoint     { return a }
func (a SourceEndpoint) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *SourceEndpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = SourceEndpoint(*nt)
	return br, err
}

func (a *SourceEndpoint) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = SourceEndpoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SourceEndpoint) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// StartIndex provides the starting index for the requested elements of the associated list.
type StartIndex zcl.Zu8

func (a *StartIndex) Value() *StartIndex         { return a }
func (a StartIndex) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *StartIndex) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = StartIndex(*nt)
	return br, err
}

func (a *StartIndex) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = StartIndex(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a StartIndex) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// Status Code, command is normally empty unless status is `Success`
type Status zcl.Zenum8

func (a *Status) Value() *Status             { return a }
func (a Status) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *Status) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = Status(*nt)
	return br, err
}

func (a *Status) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = Status(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Status) String() string {
	switch a {
	case 0x00:
		return "Success"
	case 0x80:
		return "Invalid Request Type"
	case 0x81:
		return "Device Not Found"
	case 0x82:
		return "Invalid Endpoint"
	case 0x83:
		return "Not Active"
	case 0x84:
		return "Not Supported"
	case 0x85:
		return "Timeout"
	case 0x86:
		return "No Match"
	case 0x88:
		return "No Entry"
	case 0x89:
		return "No Descriptor"
	case 0x8A:
		return "Insufficient Space"
	case 0x8B:
		return "Not Permitted"
	case 0x8C:
		return "Table Full"
	case 0x8D:
		return "Not Authorized"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a Status) IsSuccess() bool            { return a == 0x00 }
func (a Status) IsInvalidRequestType() bool { return a == 0x80 }
func (a Status) IsDeviceNotFound() bool     { return a == 0x81 }
func (a Status) IsInvalidEndpoint() bool    { return a == 0x82 }
func (a Status) IsNotActive() bool          { return a == 0x83 }
func (a Status) IsNotSupported() bool       { return a == 0x84 }
func (a Status) IsTimeout() bool            { return a == 0x85 }
func (a Status) IsNoMatch() bool            { return a == 0x86 }
func (a Status) IsNoEntry() bool            { return a == 0x88 }
func (a Status) IsNoDescriptor() bool       { return a == 0x89 }
func (a Status) IsInsufficientSpace() bool  { return a == 0x8A }
func (a Status) IsNotPermitted() bool       { return a == 0x8B }
func (a Status) IsTableFull() bool          { return a == 0x8C }
func (a Status) IsNotAuthorized() bool      { return a == 0x8D }
func (a *Status) SetSuccess()               { *a = 0x00 }
func (a *Status) SetInvalidRequestType()    { *a = 0x80 }
func (a *Status) SetDeviceNotFound()        { *a = 0x81 }
func (a *Status) SetInvalidEndpoint()       { *a = 0x82 }
func (a *Status) SetNotActive()             { *a = 0x83 }
func (a *Status) SetNotSupported()          { *a = 0x84 }
func (a *Status) SetTimeout()               { *a = 0x85 }
func (a *Status) SetNoMatch()               { *a = 0x86 }
func (a *Status) SetNoEntry()               { *a = 0x88 }
func (a *Status) SetNoDescriptor()          { *a = 0x89 }
func (a *Status) SetInsufficientSpace()     { *a = 0x8A }
func (a *Status) SetNotPermitted()          { *a = 0x8B }
func (a *Status) SetTableFull()             { *a = 0x8C }
func (a *Status) SetNotAuthorized()         { *a = 0x8D }

type Type zcl.Zbmp8

func (a *Type) Value() *Type               { return a }
func (a Type) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *Type) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = Type(*nt)
	return br, err
}

func (a *Type) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = Type(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Type) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 3:
			bstr = append(bstr, "User Desc Available")
		case 4:
			bstr = append(bstr, "Complex Desc Available")
		case 5:
			bstr = append(bstr, "Router")
		case 6:
			bstr = append(bstr, "End Device")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a Type) IsUserDescAvailable() bool       { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a Type) IsComplexDescAvailable() bool    { return zcl.BitmapTest([]byte(a[:]), 4) }
func (a Type) IsRouter() bool                  { return zcl.BitmapTest([]byte(a[:]), 5) }
func (a Type) IsEndDevice() bool               { return zcl.BitmapTest([]byte(a[:]), 6) }
func (a *Type) SetUserDescAvailable(b bool)    { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b)) }
func (a *Type) SetComplexDescAvailable(b bool) { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 4, b)) }
func (a *Type) SetRouter(b bool)               { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 5, b)) }
func (a *Type) SetEndDevice(b bool)            { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 6, b)) }

// UserDesc is a user provided ASCII string of 16 characters set on a remote device.
// If the string is shorter than 16 characters it should be padded with space characters (0x20).
// Control characters (0x00-0x1f) are not allowed.
type UserDesc zcl.Zcstring

func (a *UserDesc) Value() *UserDesc           { return a }
func (a UserDesc) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *UserDesc) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = UserDesc(*nt)
	return br, err
}

func (a *UserDesc) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = UserDesc(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a UserDesc) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

// NwkAddressRequest queries the 16-bit address of the Remote Device based on its known IEEE address.
// The destination addressing on this command shall be unicast or broadcast to all
// devices supporting RX on when Idle (0xFFFD)
type NwkAddressRequest struct {
	// IeeeAddress is a 64-bit MAC address
	IeeeAddress IeeeAddress
	// RequestType should be set to 1 if extended response is requested, 0 otherwise
	RequestType RequestType
	// StartIndex provides the starting index for the requested elements of the associated list.
	StartIndex StartIndex
}

// NwkAddressRequestCommand is the Command ID of NwkAddressRequest
const NwkAddressRequestCommand CommandID = 0x0000

// Values returns all values of NwkAddressRequest
func (v *NwkAddressRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.IeeeAddress,
		&v.RequestType,
		&v.StartIndex,
	}
}

// Name of the command (needed to fulfill interface)
func (NwkAddressRequest) Name() string { return "NWK Address Request" }

// ID of the command (needed to fulfill interface)
func (NwkAddressRequest) ID() CommandID { return NwkAddressRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (NwkAddressRequest) Cluster() zcl.ClusterID { return 0x0000 }

// MnfCode returns the manufacturer code (if any) of the command
func (NwkAddressRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of NwkAddressRequest
func (v NwkAddressRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.IeeeAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.RequestType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.StartIndex.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the NwkAddressRequest struct
func (v *NwkAddressRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.IeeeAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.RequestType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StartIndex).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v NwkAddressRequest) String() string {
	return zcl.Sprintf(
		"NwkAddressRequest{"+zcl.StrJoin([]string{
			"IeeeAddress(%v)",
			"RequestType(%v)",
			"StartIndex(%v)",
		}, " ")+"}",
		v.IeeeAddress,
		v.RequestType,
		v.StartIndex,
	)
}

type NwkAddressResponse struct {
	// Status Code, command is normally empty unless status is `Success`
	Status Status
	// IeeeAddress is a 64-bit MAC address
	IeeeAddress IeeeAddress
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	// StartIndex provides the starting index for the requested elements of the associated list.
	StartIndex        StartIndex
	AssociatedDevices AssociatedDevices
}

// NwkAddressResponseCommand is the Command ID of NwkAddressResponse
const NwkAddressResponseCommand CommandID = 0x8000

// Values returns all values of NwkAddressResponse
func (v *NwkAddressResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.IeeeAddress,
		&v.NwkAddress,
		&v.StartIndex,
		&v.AssociatedDevices,
	}
}

// Name of the command (needed to fulfill interface)
func (NwkAddressResponse) Name() string { return "NWK Address Response" }

// ID of the command (needed to fulfill interface)
func (NwkAddressResponse) ID() CommandID { return NwkAddressResponseCommand }

// Cluster ID of the command (needed to fulfill interface)
func (NwkAddressResponse) Cluster() zcl.ClusterID { return 0x8000 }

// MnfCode returns the manufacturer code (if any) of the command
func (NwkAddressResponse) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of NwkAddressResponse
func (v NwkAddressResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.IeeeAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// is only included if successful
	if v.Status == 0x00 {
		x := zcl.Zu8(2 * len(v.AssociatedDevices.Content))
		if tmp, err = x.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// StartIndex is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.StartIndex.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// AssociatedDevices is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.AssociatedDevices.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the NwkAddressResponse struct
func (v *NwkAddressResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.IeeeAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	// is only included if successful
	if v.Status == 0x00 {
		if b, err = new(zcl.Zu8).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// StartIndex is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.StartIndex).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// AssociatedDevices is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.AssociatedDevices).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v NwkAddressResponse) String() string {
	return zcl.Sprintf(
		"NwkAddressResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"IeeeAddress(%v)",
			"NwkAddress(%v)",
			"StartIndex(%v)",
			"AssociatedDevices(%v)",
		}, " ")+"}",
		v.Status,
		v.IeeeAddress,
		v.NwkAddress,
		v.StartIndex,
		v.AssociatedDevices,
	)
}

// IeeeAddressRequest queries the 64-bit IEEE address of the Remote Device based on its known 16-bit NWK address.
// The command should always be sent by unicast.
type IeeeAddressRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	// RequestType should be set to 1 if extended response is requested, 0 otherwise
	RequestType RequestType
	// StartIndex provides the starting index for the requested elements of the associated list.
	StartIndex StartIndex
}

// IeeeAddressRequestCommand is the Command ID of IeeeAddressRequest
const IeeeAddressRequestCommand CommandID = 0x0001

// Values returns all values of IeeeAddressRequest
func (v *IeeeAddressRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.RequestType,
		&v.StartIndex,
	}
}

// Name of the command (needed to fulfill interface)
func (IeeeAddressRequest) Name() string { return "IEEE Address Request" }

// ID of the command (needed to fulfill interface)
func (IeeeAddressRequest) ID() CommandID { return IeeeAddressRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (IeeeAddressRequest) Cluster() zcl.ClusterID { return 0x0001 }

// MnfCode returns the manufacturer code (if any) of the command
func (IeeeAddressRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of IeeeAddressRequest
func (v IeeeAddressRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.RequestType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.StartIndex.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the IeeeAddressRequest struct
func (v *IeeeAddressRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.RequestType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StartIndex).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v IeeeAddressRequest) String() string {
	return zcl.Sprintf(
		"IeeeAddressRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
			"RequestType(%v)",
			"StartIndex(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.RequestType,
		v.StartIndex,
	)
}

// NodeDescRequest queries the node descriptor of a remote device. Should be unicast to the remote device directly,
// or to a device that contains the discovery information of the remote device.
type NodeDescRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
}

// NodeDescRequestCommand is the Command ID of NodeDescRequest
const NodeDescRequestCommand CommandID = 0x0002

// Values returns all values of NodeDescRequest
func (v *NodeDescRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
	}
}

// Name of the command (needed to fulfill interface)
func (NodeDescRequest) Name() string { return "Node Desc Request" }

// ID of the command (needed to fulfill interface)
func (NodeDescRequest) ID() CommandID { return NodeDescRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (NodeDescRequest) Cluster() zcl.ClusterID { return 0x0002 }

// MnfCode returns the manufacturer code (if any) of the command
func (NodeDescRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of NodeDescRequest
func (v NodeDescRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the NodeDescRequest struct
func (v *NodeDescRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v NodeDescRequest) String() string {
	return zcl.Sprintf(
		"NodeDescRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
		}, " ")+"}",
		v.NwkAddress,
	)
}

// PowerDescRequest queries the power descriptor of a remote device. Should be unicast to the remote device directly,
// or to a device that contains the discovery information of the remote device.
type PowerDescRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
}

// PowerDescRequestCommand is the Command ID of PowerDescRequest
const PowerDescRequestCommand CommandID = 0x0003

// Values returns all values of PowerDescRequest
func (v *PowerDescRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
	}
}

// Name of the command (needed to fulfill interface)
func (PowerDescRequest) Name() string { return "Power Desc Request" }

// ID of the command (needed to fulfill interface)
func (PowerDescRequest) ID() CommandID { return PowerDescRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (PowerDescRequest) Cluster() zcl.ClusterID { return 0x0003 }

// MnfCode returns the manufacturer code (if any) of the command
func (PowerDescRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of PowerDescRequest
func (v PowerDescRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the PowerDescRequest struct
func (v *PowerDescRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v PowerDescRequest) String() string {
	return zcl.Sprintf(
		"PowerDescRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
		}, " ")+"}",
		v.NwkAddress,
	)
}

// SimpleDescRequest queries the Simple Descriptor of a remote device on a specific endpoint.
// Simple Descriptor contains information about which clusters the device supports on the given endpoint.
// Should be unicast to the remote device directly, or to a device that contains the discovery information of the remote device.
type SimpleDescRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	Endpoint   Endpoint
}

// SimpleDescRequestCommand is the Command ID of SimpleDescRequest
const SimpleDescRequestCommand CommandID = 0x0004

// Values returns all values of SimpleDescRequest
func (v *SimpleDescRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.Endpoint,
	}
}

// Name of the command (needed to fulfill interface)
func (SimpleDescRequest) Name() string { return "Simple Desc Request" }

// ID of the command (needed to fulfill interface)
func (SimpleDescRequest) ID() CommandID { return SimpleDescRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (SimpleDescRequest) Cluster() zcl.ClusterID { return 0x0004 }

// MnfCode returns the manufacturer code (if any) of the command
func (SimpleDescRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of SimpleDescRequest
func (v SimpleDescRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Endpoint.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the SimpleDescRequest struct
func (v *SimpleDescRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Endpoint).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v SimpleDescRequest) String() string {
	return zcl.Sprintf(
		"SimpleDescRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
			"Endpoint(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.Endpoint,
	)
}

// ActiveEndpointRequest queries the remote device for a list of active endpoints. Should be unicast to the remote device directly,
// or to a device that contains the discovery information of the remote device.
type ActiveEndpointRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
}

// ActiveEndpointRequestCommand is the Command ID of ActiveEndpointRequest
const ActiveEndpointRequestCommand CommandID = 0x0005

// Values returns all values of ActiveEndpointRequest
func (v *ActiveEndpointRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
	}
}

// Name of the command (needed to fulfill interface)
func (ActiveEndpointRequest) Name() string { return "Active Endpoint Request" }

// ID of the command (needed to fulfill interface)
func (ActiveEndpointRequest) ID() CommandID { return ActiveEndpointRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (ActiveEndpointRequest) Cluster() zcl.ClusterID { return 0x0005 }

// MnfCode returns the manufacturer code (if any) of the command
func (ActiveEndpointRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of ActiveEndpointRequest
func (v ActiveEndpointRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ActiveEndpointRequest struct
func (v *ActiveEndpointRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ActiveEndpointRequest) String() string {
	return zcl.Sprintf(
		"ActiveEndpointRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
		}, " ")+"}",
		v.NwkAddress,
	)
}

// MatchDescRequest is used to find remote devices supporting a specific simple descriptor match criterion. Normally broadcast
// to all devices supporting RX on when Idle (0xFFFD)
type MatchDescRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	ProfileId  ProfileId
	// InClusterList is a list of input clusters
	InClusterList InClusterList
	// OutClusterList is a list of output clusters
	OutClusterList OutClusterList
}

// MatchDescRequestCommand is the Command ID of MatchDescRequest
const MatchDescRequestCommand CommandID = 0x0006

// Values returns all values of MatchDescRequest
func (v *MatchDescRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.ProfileId,
		&v.InClusterList,
		&v.OutClusterList,
	}
}

// Name of the command (needed to fulfill interface)
func (MatchDescRequest) Name() string { return "Match Desc Request" }

// ID of the command (needed to fulfill interface)
func (MatchDescRequest) ID() CommandID { return MatchDescRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (MatchDescRequest) Cluster() zcl.ClusterID { return 0x0006 }

// MnfCode returns the manufacturer code (if any) of the command
func (MatchDescRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of MatchDescRequest
func (v MatchDescRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ProfileId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.InClusterList.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.OutClusterList.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the MatchDescRequest struct
func (v *MatchDescRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ProfileId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.InClusterList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.OutClusterList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MatchDescRequest) String() string {
	return zcl.Sprintf(
		"MatchDescRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
			"ProfileId(%v)",
			"InClusterList(%v)",
			"OutClusterList(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.ProfileId,
		v.InClusterList,
		v.OutClusterList,
	)
}

// ComplexDescRequest queries the Complex Descriptor of a remote device. Should be unicast to the remote device directly,
// or to a device that contains the discovery information of the remote device.
type ComplexDescRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
}

// ComplexDescRequestCommand is the Command ID of ComplexDescRequest
const ComplexDescRequestCommand CommandID = 0x0010

// Values returns all values of ComplexDescRequest
func (v *ComplexDescRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
	}
}

// Name of the command (needed to fulfill interface)
func (ComplexDescRequest) Name() string { return "Complex Desc Request" }

// ID of the command (needed to fulfill interface)
func (ComplexDescRequest) ID() CommandID { return ComplexDescRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (ComplexDescRequest) Cluster() zcl.ClusterID { return 0x0010 }

// MnfCode returns the manufacturer code (if any) of the command
func (ComplexDescRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of ComplexDescRequest
func (v ComplexDescRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ComplexDescRequest struct
func (v *ComplexDescRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ComplexDescRequest) String() string {
	return zcl.Sprintf(
		"ComplexDescRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
		}, " ")+"}",
		v.NwkAddress,
	)
}

// UserDescRequest queries the User Descriptor of a remote device. Should be unicast to the remote device directly,
// or to a device that contains the discovery information of the remote device.
type UserDescRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
}

// UserDescRequestCommand is the Command ID of UserDescRequest
const UserDescRequestCommand CommandID = 0x0011

// Values returns all values of UserDescRequest
func (v *UserDescRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
	}
}

// Name of the command (needed to fulfill interface)
func (UserDescRequest) Name() string { return "User Desc Request" }

// ID of the command (needed to fulfill interface)
func (UserDescRequest) ID() CommandID { return UserDescRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (UserDescRequest) Cluster() zcl.ClusterID { return 0x0011 }

// MnfCode returns the manufacturer code (if any) of the command
func (UserDescRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of UserDescRequest
func (v UserDescRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the UserDescRequest struct
func (v *UserDescRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v UserDescRequest) String() string {
	return zcl.Sprintf(
		"UserDescRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
		}, " ")+"}",
		v.NwkAddress,
	)
}

// DiscoveryCacheRequest locates a Primary Discovery Cache on the network. Should be addressed to broadcast RXOnWhenIdle (0xFFFD)
type DiscoveryCacheRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	// IeeeAddress is a 64-bit MAC address
	IeeeAddress IeeeAddress
}

// DiscoveryCacheRequestCommand is the Command ID of DiscoveryCacheRequest
const DiscoveryCacheRequestCommand CommandID = 0x0012

// Values returns all values of DiscoveryCacheRequest
func (v *DiscoveryCacheRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.IeeeAddress,
	}
}

// Name of the command (needed to fulfill interface)
func (DiscoveryCacheRequest) Name() string { return "Discovery Cache Request" }

// ID of the command (needed to fulfill interface)
func (DiscoveryCacheRequest) ID() CommandID { return DiscoveryCacheRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (DiscoveryCacheRequest) Cluster() zcl.ClusterID { return 0x0012 }

// MnfCode returns the manufacturer code (if any) of the command
func (DiscoveryCacheRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of DiscoveryCacheRequest
func (v DiscoveryCacheRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.IeeeAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the DiscoveryCacheRequest struct
func (v *DiscoveryCacheRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.IeeeAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v DiscoveryCacheRequest) String() string {
	return zcl.Sprintf(
		"DiscoveryCacheRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
			"IeeeAddress(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.IeeeAddress,
	)
}

// DeviceAnnounce is sent by a joining or returning device, identifying it's NWK address, IEEE address and capabilities.
// Normally sent to broadcast RXOnWhenIdle (0xFFFD)
type DeviceAnnounce struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	// IeeeAddress is a 64-bit MAC address
	IeeeAddress IeeeAddress
	// Capability specifies the device:s capabilities
	Capability Capability
}

// DeviceAnnounceCommand is the Command ID of DeviceAnnounce
const DeviceAnnounceCommand CommandID = 0x0013

// Values returns all values of DeviceAnnounce
func (v *DeviceAnnounce) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.IeeeAddress,
		&v.Capability,
	}
}

// Name of the command (needed to fulfill interface)
func (DeviceAnnounce) Name() string { return "Device Announce" }

// ID of the command (needed to fulfill interface)
func (DeviceAnnounce) ID() CommandID { return DeviceAnnounceCommand }

// Cluster ID of the command (needed to fulfill interface)
func (DeviceAnnounce) Cluster() zcl.ClusterID { return 0x0013 }

// MnfCode returns the manufacturer code (if any) of the command
func (DeviceAnnounce) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of DeviceAnnounce
func (v DeviceAnnounce) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.IeeeAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Capability.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the DeviceAnnounce struct
func (v *DeviceAnnounce) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.IeeeAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Capability).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v DeviceAnnounce) String() string {
	return zcl.Sprintf(
		"DeviceAnnounce{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
			"IeeeAddress(%v)",
			"Capability(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.IeeeAddress,
		v.Capability,
	)
}

// UserDescSetRequest requests the remote device to update it's user descriptor.
type UserDescSetRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	// UserDesc is a user provided ASCII string of 16 characters set on a remote device.
	// If the string is shorter than 16 characters it should be padded with space characters (0x20).
	// Control characters (0x00-0x1f) are not allowed.
	UserDesc UserDesc
}

// UserDescSetRequestCommand is the Command ID of UserDescSetRequest
const UserDescSetRequestCommand CommandID = 0x0014

// Values returns all values of UserDescSetRequest
func (v *UserDescSetRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.UserDesc,
	}
}

// Name of the command (needed to fulfill interface)
func (UserDescSetRequest) Name() string { return "User Desc Set Request" }

// ID of the command (needed to fulfill interface)
func (UserDescSetRequest) ID() CommandID { return UserDescSetRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (UserDescSetRequest) Cluster() zcl.ClusterID { return 0x0014 }

// MnfCode returns the manufacturer code (if any) of the command
func (UserDescSetRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of UserDescSetRequest
func (v UserDescSetRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.UserDesc.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the UserDescSetRequest struct
func (v *UserDescSetRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.UserDesc).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v UserDescSetRequest) String() string {
	return zcl.Sprintf(
		"UserDescSetRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
			"UserDesc(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.UserDesc,
	)
}

// SystemServerDiscoverRequest discovers the location of a particular system server or servers as indicated by the Server Mask. Should be sent to broadcast RXOnWhenIdle (0xFFFD)
type SystemServerDiscoverRequest struct {
	ServerMask ServerMask
}

// SystemServerDiscoverRequestCommand is the Command ID of SystemServerDiscoverRequest
const SystemServerDiscoverRequestCommand CommandID = 0x0015

// Values returns all values of SystemServerDiscoverRequest
func (v *SystemServerDiscoverRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.ServerMask,
	}
}

// Name of the command (needed to fulfill interface)
func (SystemServerDiscoverRequest) Name() string { return "System Server Discover Request" }

// ID of the command (needed to fulfill interface)
func (SystemServerDiscoverRequest) ID() CommandID { return SystemServerDiscoverRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (SystemServerDiscoverRequest) Cluster() zcl.ClusterID { return 0x0015 }

// MnfCode returns the manufacturer code (if any) of the command
func (SystemServerDiscoverRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of SystemServerDiscoverRequest
func (v SystemServerDiscoverRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.ServerMask.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the SystemServerDiscoverRequest struct
func (v *SystemServerDiscoverRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.ServerMask).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v SystemServerDiscoverRequest) String() string {
	return zcl.Sprintf(
		"SystemServerDiscoverRequest{"+zcl.StrJoin([]string{
			"ServerMask(%v)",
		}, " ")+"}",
		v.ServerMask,
	)
}

// DiscoveryStoreRequest sent to a Discovery Cache Node to allocate memory of the provided sizes for cache storage.
// If the node already exists in cache, it will be removed to allow storage of updated values.
// Should be sent to the unicast address of a discovery cache device.
type DiscoveryStoreRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	// IeeeAddress is a 64-bit MAC address
	IeeeAddress IeeeAddress
	// NodeDescSize Size in bytes of the Node Descriptor
	NodeDescSize NodeDescSize
	// PowerDescSize Size in bytes of the Power Descriptor
	PowerDescSize PowerDescSize
	// ActiveEndpointSize Size in bytes of the Active Endpoints List
	ActiveEndpointSize ActiveEndpointSize
	// SimpleDescSizeList List of sizes for the different Simple Descriptors
	SimpleDescSizeList SimpleDescSizeList
}

// DiscoveryStoreRequestCommand is the Command ID of DiscoveryStoreRequest
const DiscoveryStoreRequestCommand CommandID = 0x0016

// Values returns all values of DiscoveryStoreRequest
func (v *DiscoveryStoreRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.IeeeAddress,
		&v.NodeDescSize,
		&v.PowerDescSize,
		&v.ActiveEndpointSize,
		&v.SimpleDescSizeList,
	}
}

// Name of the command (needed to fulfill interface)
func (DiscoveryStoreRequest) Name() string { return "Discovery Store Request" }

// ID of the command (needed to fulfill interface)
func (DiscoveryStoreRequest) ID() CommandID { return DiscoveryStoreRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (DiscoveryStoreRequest) Cluster() zcl.ClusterID { return 0x0016 }

// MnfCode returns the manufacturer code (if any) of the command
func (DiscoveryStoreRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of DiscoveryStoreRequest
func (v DiscoveryStoreRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.IeeeAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.NodeDescSize.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.PowerDescSize.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ActiveEndpointSize.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SimpleDescSizeList.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the DiscoveryStoreRequest struct
func (v *DiscoveryStoreRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.IeeeAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NodeDescSize).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.PowerDescSize).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ActiveEndpointSize).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SimpleDescSizeList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v DiscoveryStoreRequest) String() string {
	return zcl.Sprintf(
		"DiscoveryStoreRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
			"IeeeAddress(%v)",
			"NodeDescSize(%v)",
			"PowerDescSize(%v)",
			"ActiveEndpointSize(%v)",
			"SimpleDescSizeList(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.IeeeAddress,
		v.NodeDescSize,
		v.PowerDescSize,
		v.ActiveEndpointSize,
		v.SimpleDescSizeList,
	)
}

// NodeDescStoreRequest sent to a Discovery Cache Node to store the provided Node Descriptor.
// A DiscoveryStoreRequest should be sent and acknowledged with successful status before requesting storage.
// Should be sent to the unicast address of a discovery cache device.
type NodeDescStoreRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	// IeeeAddress is a 64-bit MAC address
	IeeeAddress    IeeeAddress
	Type           Type
	FrequencyBands FrequencyBands
	// Capability specifies the device:s capabilities
	Capability       Capability
	ManufacturerCode ManufacturerCode
	// MaxBufferSize specifies the maximum size, in octets, of the network sub-layer data unit (NSDU) for this node.
	// This is the maximum size of data or commands passed to or from the application by the application support sub-layer,
	// before any fragmentation or re-assembly.
	MaxBufferSize MaxBufferSize
	// MaxIncomingTransferSize specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred to this node in one single message transfer.
	MaxIncomingTransferSize MaxIncomingTransferSize
	ServerMask              ServerMask
	// MaxOutgoingTransferSize specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred from
	// this node in one single message transfer.
	MaxOutgoingTransferSize MaxOutgoingTransferSize
	DescriptorCapability    DescriptorCapability
}

// NodeDescStoreRequestCommand is the Command ID of NodeDescStoreRequest
const NodeDescStoreRequestCommand CommandID = 0x0017

// Values returns all values of NodeDescStoreRequest
func (v *NodeDescStoreRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.IeeeAddress,
		&v.Type,
		&v.FrequencyBands,
		&v.Capability,
		&v.ManufacturerCode,
		&v.MaxBufferSize,
		&v.MaxIncomingTransferSize,
		&v.ServerMask,
		&v.MaxOutgoingTransferSize,
		&v.DescriptorCapability,
	}
}

// Name of the command (needed to fulfill interface)
func (NodeDescStoreRequest) Name() string { return "Node Desc Store Request" }

// ID of the command (needed to fulfill interface)
func (NodeDescStoreRequest) ID() CommandID { return NodeDescStoreRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (NodeDescStoreRequest) Cluster() zcl.ClusterID { return 0x0017 }

// MnfCode returns the manufacturer code (if any) of the command
func (NodeDescStoreRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of NodeDescStoreRequest
func (v NodeDescStoreRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.IeeeAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Type.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.FrequencyBands.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Capability.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ManufacturerCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.MaxBufferSize.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.MaxIncomingTransferSize.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ServerMask.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.MaxOutgoingTransferSize.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.DescriptorCapability.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the NodeDescStoreRequest struct
func (v *NodeDescStoreRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.IeeeAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Type).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.FrequencyBands).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Capability).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ManufacturerCode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.MaxBufferSize).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.MaxIncomingTransferSize).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ServerMask).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.MaxOutgoingTransferSize).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.DescriptorCapability).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v NodeDescStoreRequest) String() string {
	return zcl.Sprintf(
		"NodeDescStoreRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
			"IeeeAddress(%v)",
			"Type(%v)",
			"FrequencyBands(%v)",
			"Capability(%v)",
			"ManufacturerCode(%v)",
			"MaxBufferSize(%v)",
			"MaxIncomingTransferSize(%v)",
			"ServerMask(%v)",
			"MaxOutgoingTransferSize(%v)",
			"DescriptorCapability(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.IeeeAddress,
		v.Type,
		v.FrequencyBands,
		v.Capability,
		v.ManufacturerCode,
		v.MaxBufferSize,
		v.MaxIncomingTransferSize,
		v.ServerMask,
		v.MaxOutgoingTransferSize,
		v.DescriptorCapability,
	)
}

// PowerDescStoreRequest sent to a Discovery Cache Node to store the provided Power Descriptor.
// A DiscoveryStoreRequest should be sent and acknowledged with successful status before requesting storage.
// Should be sent to the unicast address of a discovery cache device.
type PowerDescStoreRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	// IeeeAddress is a 64-bit MAC address
	IeeeAddress        IeeeAddress
	PowerMode          PowerMode
	CurrentPowerSource CurrentPowerSource
}

// PowerDescStoreRequestCommand is the Command ID of PowerDescStoreRequest
const PowerDescStoreRequestCommand CommandID = 0x0018

// Values returns all values of PowerDescStoreRequest
func (v *PowerDescStoreRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.IeeeAddress,
		&v.PowerMode,
		&v.CurrentPowerSource,
	}
}

// Name of the command (needed to fulfill interface)
func (PowerDescStoreRequest) Name() string { return "Power Desc Store Request" }

// ID of the command (needed to fulfill interface)
func (PowerDescStoreRequest) ID() CommandID { return PowerDescStoreRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (PowerDescStoreRequest) Cluster() zcl.ClusterID { return 0x0018 }

// MnfCode returns the manufacturer code (if any) of the command
func (PowerDescStoreRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of PowerDescStoreRequest
func (v PowerDescStoreRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.IeeeAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.PowerMode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.CurrentPowerSource.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the PowerDescStoreRequest struct
func (v *PowerDescStoreRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.IeeeAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.PowerMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.CurrentPowerSource).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v PowerDescStoreRequest) String() string {
	return zcl.Sprintf(
		"PowerDescStoreRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
			"IeeeAddress(%v)",
			"PowerMode(%v)",
			"CurrentPowerSource(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.IeeeAddress,
		v.PowerMode,
		v.CurrentPowerSource,
	)
}

// ActiveEndpointStoreRequest sent to a Discovery Cache Node to store the provided Active Endpoint list.
// A DiscoveryStoreRequest should be sent and acknowledged with successful status before requesting storage.
// Should be sent to the unicast address of a discovery cache device.
type ActiveEndpointStoreRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	// IeeeAddress is a 64-bit MAC address
	IeeeAddress  IeeeAddress
	EndpointList EndpointList
}

// ActiveEndpointStoreRequestCommand is the Command ID of ActiveEndpointStoreRequest
const ActiveEndpointStoreRequestCommand CommandID = 0x0019

// Values returns all values of ActiveEndpointStoreRequest
func (v *ActiveEndpointStoreRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.IeeeAddress,
		&v.EndpointList,
	}
}

// Name of the command (needed to fulfill interface)
func (ActiveEndpointStoreRequest) Name() string { return "Active Endpoint Store Request" }

// ID of the command (needed to fulfill interface)
func (ActiveEndpointStoreRequest) ID() CommandID { return ActiveEndpointStoreRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (ActiveEndpointStoreRequest) Cluster() zcl.ClusterID { return 0x0019 }

// MnfCode returns the manufacturer code (if any) of the command
func (ActiveEndpointStoreRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of ActiveEndpointStoreRequest
func (v ActiveEndpointStoreRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.IeeeAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.EndpointList.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ActiveEndpointStoreRequest struct
func (v *ActiveEndpointStoreRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.IeeeAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.EndpointList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ActiveEndpointStoreRequest) String() string {
	return zcl.Sprintf(
		"ActiveEndpointStoreRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
			"IeeeAddress(%v)",
			"EndpointList(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.IeeeAddress,
		v.EndpointList,
	)
}

// SimpleDescStoreRequest sent to a Discovery Cache Node to store the provided Simple Descriptor.
// A DiscoveryStoreRequest should be sent and acknowledged with successful status before requesting storage.
// Should be sent to the unicast address of a discovery cache device.
type SimpleDescStoreRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	// IeeeAddress is a 64-bit MAC address
	IeeeAddress IeeeAddress
	Endpoint    Endpoint
	ProfileId   ProfileId
	DeviceType  DeviceType
	// DeviceVersion is dependant on profile
	DeviceVersion DeviceVersion
	// InClusterList is a list of input clusters
	InClusterList InClusterList
	// OutClusterList is a list of output clusters
	OutClusterList OutClusterList
}

// SimpleDescStoreRequestCommand is the Command ID of SimpleDescStoreRequest
const SimpleDescStoreRequestCommand CommandID = 0x001A

// Values returns all values of SimpleDescStoreRequest
func (v *SimpleDescStoreRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.IeeeAddress,
		&v.Endpoint,
		&v.ProfileId,
		&v.DeviceType,
		&v.DeviceVersion,
		&v.InClusterList,
		&v.OutClusterList,
	}
}

// Name of the command (needed to fulfill interface)
func (SimpleDescStoreRequest) Name() string { return "Simple Desc Store Request" }

// ID of the command (needed to fulfill interface)
func (SimpleDescStoreRequest) ID() CommandID { return SimpleDescStoreRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (SimpleDescStoreRequest) Cluster() zcl.ClusterID { return 0x001A }

// MnfCode returns the manufacturer code (if any) of the command
func (SimpleDescStoreRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of SimpleDescStoreRequest
func (v SimpleDescStoreRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.IeeeAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		x := zcl.Zu8(8 + 2*len(v.InClusterList.Content) + 2*len(v.OutClusterList.Content))
		if tmp, err = x.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Endpoint.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ProfileId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.DeviceType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.DeviceVersion.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.InClusterList.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.OutClusterList.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the SimpleDescStoreRequest struct
func (v *SimpleDescStoreRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.IeeeAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = new(zcl.Zu8).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Endpoint).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ProfileId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.DeviceType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.DeviceVersion).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.InClusterList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.OutClusterList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v SimpleDescStoreRequest) String() string {
	return zcl.Sprintf(
		"SimpleDescStoreRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
			"IeeeAddress(%v)",
			"Endpoint(%v)",
			"ProfileId(%v)",
			"DeviceType(%v)",
			"DeviceVersion(%v)",
			"InClusterList(%v)",
			"OutClusterList(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.IeeeAddress,
		v.Endpoint,
		v.ProfileId,
		v.DeviceType,
		v.DeviceVersion,
		v.InClusterList,
		v.OutClusterList,
	)
}

// RemoveNodeCacheRequest sent to a Discovery Cache Node will request it to remove cached values for the provided node. Should be sent to the unicast address of a discovery cache device.
type RemoveNodeCacheRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	// IeeeAddress is a 64-bit MAC address
	IeeeAddress IeeeAddress
}

// RemoveNodeCacheRequestCommand is the Command ID of RemoveNodeCacheRequest
const RemoveNodeCacheRequestCommand CommandID = 0x001B

// Values returns all values of RemoveNodeCacheRequest
func (v *RemoveNodeCacheRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.IeeeAddress,
	}
}

// Name of the command (needed to fulfill interface)
func (RemoveNodeCacheRequest) Name() string { return "Remove Node Cache Request" }

// ID of the command (needed to fulfill interface)
func (RemoveNodeCacheRequest) ID() CommandID { return RemoveNodeCacheRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (RemoveNodeCacheRequest) Cluster() zcl.ClusterID { return 0x001B }

// MnfCode returns the manufacturer code (if any) of the command
func (RemoveNodeCacheRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of RemoveNodeCacheRequest
func (v RemoveNodeCacheRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.IeeeAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the RemoveNodeCacheRequest struct
func (v *RemoveNodeCacheRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.IeeeAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v RemoveNodeCacheRequest) String() string {
	return zcl.Sprintf(
		"RemoveNodeCacheRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
			"IeeeAddress(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.IeeeAddress,
	)
}

// FindNodeCacheRequest broadcast to the network will generate a response from the Primary Discovery Cache holding information
// for the device of interest
type FindNodeCacheRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	// IeeeAddress is a 64-bit MAC address
	IeeeAddress IeeeAddress
}

// FindNodeCacheRequestCommand is the Command ID of FindNodeCacheRequest
const FindNodeCacheRequestCommand CommandID = 0x001C

// Values returns all values of FindNodeCacheRequest
func (v *FindNodeCacheRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.IeeeAddress,
	}
}

// Name of the command (needed to fulfill interface)
func (FindNodeCacheRequest) Name() string { return "Find Node Cache Request" }

// ID of the command (needed to fulfill interface)
func (FindNodeCacheRequest) ID() CommandID { return FindNodeCacheRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (FindNodeCacheRequest) Cluster() zcl.ClusterID { return 0x001C }

// MnfCode returns the manufacturer code (if any) of the command
func (FindNodeCacheRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of FindNodeCacheRequest
func (v FindNodeCacheRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.IeeeAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the FindNodeCacheRequest struct
func (v *FindNodeCacheRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.IeeeAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v FindNodeCacheRequest) String() string {
	return zcl.Sprintf(
		"FindNodeCacheRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
			"IeeeAddress(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.IeeeAddress,
	)
}

// ExtendedSimpleDescRequest should be unicast to the remote device or a discovery cache node. It is used to request information from
// nodes that implements a larger number of clusters than can be described by a SimpleDescRequest
type ExtendedSimpleDescRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	Endpoint   Endpoint
	// StartIndex provides the starting index for the requested elements of the associated list.
	StartIndex StartIndex
}

// ExtendedSimpleDescRequestCommand is the Command ID of ExtendedSimpleDescRequest
const ExtendedSimpleDescRequestCommand CommandID = 0x001D

// Values returns all values of ExtendedSimpleDescRequest
func (v *ExtendedSimpleDescRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.Endpoint,
		&v.StartIndex,
	}
}

// Name of the command (needed to fulfill interface)
func (ExtendedSimpleDescRequest) Name() string { return "Extended Simple Desc Request" }

// ID of the command (needed to fulfill interface)
func (ExtendedSimpleDescRequest) ID() CommandID { return ExtendedSimpleDescRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (ExtendedSimpleDescRequest) Cluster() zcl.ClusterID { return 0x001D }

// MnfCode returns the manufacturer code (if any) of the command
func (ExtendedSimpleDescRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of ExtendedSimpleDescRequest
func (v ExtendedSimpleDescRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Endpoint.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.StartIndex.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ExtendedSimpleDescRequest struct
func (v *ExtendedSimpleDescRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Endpoint).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StartIndex).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ExtendedSimpleDescRequest) String() string {
	return zcl.Sprintf(
		"ExtendedSimpleDescRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
			"Endpoint(%v)",
			"StartIndex(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.Endpoint,
		v.StartIndex,
	)
}

// ExtendedActiveEndpointRequest should be unicast to the remote device or a discovery cache node. It is used to request information from
// nodes that implements a larger number of endpoints than can be described by a ActiveEndpointRequest
type ExtendedActiveEndpointRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	// StartIndex provides the starting index for the requested elements of the associated list.
	StartIndex StartIndex
}

// ExtendedActiveEndpointRequestCommand is the Command ID of ExtendedActiveEndpointRequest
const ExtendedActiveEndpointRequestCommand CommandID = 0x001E

// Values returns all values of ExtendedActiveEndpointRequest
func (v *ExtendedActiveEndpointRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.StartIndex,
	}
}

// Name of the command (needed to fulfill interface)
func (ExtendedActiveEndpointRequest) Name() string { return "Extended Active Endpoint Request" }

// ID of the command (needed to fulfill interface)
func (ExtendedActiveEndpointRequest) ID() CommandID { return ExtendedActiveEndpointRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (ExtendedActiveEndpointRequest) Cluster() zcl.ClusterID { return 0x001E }

// MnfCode returns the manufacturer code (if any) of the command
func (ExtendedActiveEndpointRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of ExtendedActiveEndpointRequest
func (v ExtendedActiveEndpointRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.StartIndex.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ExtendedActiveEndpointRequest struct
func (v *ExtendedActiveEndpointRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StartIndex).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ExtendedActiveEndpointRequest) String() string {
	return zcl.Sprintf(
		"ExtendedActiveEndpointRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
			"StartIndex(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.StartIndex,
	)
}

// EndDeviceBindRequest is sent to the ZigBee coordinator from two different devices simultaneously.
// * If the supplied endpoint is outside the specified range - a reply is sent with status `InvalidEndpoint`
// * If only one device sends the request within a pre-configure timeout - a reply is sent with status `Timeout`
// * If the ProfileID doesn't match or none of the In/OutClusterList elements match - a reply is sent with status `NoMatch`
// * Otherwise, a reply is sent with status `Success` to each device
// The Coordinator then needs to either send `BindRequest` or `UnbindRequest` for each matched `ClusterID`.
// This is done by first issuing a `UnbindRequest` with any of the matched `ClusterID`:
// * If the reply status is `NoEntry` - `BindRequest` will instead be sent for each matched `ClusterID`
// * If the reply status is `Success` - additional unbind requests are sent the rest of the matched cluster
// This enables the request to toggle the binding.
type EndDeviceBindRequest struct {
	// BindingTarget NWK Address
	BindingTarget BindingTarget
	// SourceAddress of device generating the request
	SourceAddress SourceAddress
	// SourceEndpoint of device generating the request
	SourceEndpoint SourceEndpoint
	ProfileId      ProfileId
	// InClusterList is a list of input clusters
	InClusterList InClusterList
	// OutClusterList is a list of output clusters
	OutClusterList OutClusterList
}

// EndDeviceBindRequestCommand is the Command ID of EndDeviceBindRequest
const EndDeviceBindRequestCommand CommandID = 0x0020

// Values returns all values of EndDeviceBindRequest
func (v *EndDeviceBindRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.BindingTarget,
		&v.SourceAddress,
		&v.SourceEndpoint,
		&v.ProfileId,
		&v.InClusterList,
		&v.OutClusterList,
	}
}

// Name of the command (needed to fulfill interface)
func (EndDeviceBindRequest) Name() string { return "End Device Bind Request" }

// ID of the command (needed to fulfill interface)
func (EndDeviceBindRequest) ID() CommandID { return EndDeviceBindRequestCommand }

// Cluster ID of the command (needed to fulfill interface)
func (EndDeviceBindRequest) Cluster() zcl.ClusterID { return 0x0020 }

// MnfCode returns the manufacturer code (if any) of the command
func (EndDeviceBindRequest) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of EndDeviceBindRequest
func (v EndDeviceBindRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.BindingTarget.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SourceAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SourceEndpoint.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ProfileId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.InClusterList.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.OutClusterList.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the EndDeviceBindRequest struct
func (v *EndDeviceBindRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.BindingTarget).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SourceAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SourceEndpoint).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ProfileId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.InClusterList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.OutClusterList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v EndDeviceBindRequest) String() string {
	return zcl.Sprintf(
		"EndDeviceBindRequest{"+zcl.StrJoin([]string{
			"BindingTarget(%v)",
			"SourceAddress(%v)",
			"SourceEndpoint(%v)",
			"ProfileId(%v)",
			"InClusterList(%v)",
			"OutClusterList(%v)",
		}, " ")+"}",
		v.BindingTarget,
		v.SourceAddress,
		v.SourceEndpoint,
		v.ProfileId,
		v.InClusterList,
		v.OutClusterList,
	)
}
