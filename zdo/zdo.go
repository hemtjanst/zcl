package zdo

import "hemtjan.st/zcl"

type CommandID = zcl.ZdoCmdID

var Commands = map[zcl.ZdoCmdID]func() zcl.ZdoCommand{
	NwkAddressRequestCommand:             func() zcl.ZdoCommand { return new(NwkAddressRequest) },
	NwkAddressResponseCommand:            func() zcl.ZdoCommand { return new(NwkAddressResponse) },
	IeeeAddressRequestCommand:            func() zcl.ZdoCommand { return new(IeeeAddressRequest) },
	IeeeAddressResponseCommand:           func() zcl.ZdoCommand { return new(IeeeAddressResponse) },
	NodeDescRequestCommand:               func() zcl.ZdoCommand { return new(NodeDescRequest) },
	NodeDescResponseCommand:              func() zcl.ZdoCommand { return new(NodeDescResponse) },
	PowerDescRequestCommand:              func() zcl.ZdoCommand { return new(PowerDescRequest) },
	PowerDescResponseCommand:             func() zcl.ZdoCommand { return new(PowerDescResponse) },
	SimpleDescRequestCommand:             func() zcl.ZdoCommand { return new(SimpleDescRequest) },
	SimpleDescResponseCommand:            func() zcl.ZdoCommand { return new(SimpleDescResponse) },
	ActiveEndpointRequestCommand:         func() zcl.ZdoCommand { return new(ActiveEndpointRequest) },
	ActiveEndpointResponseCommand:        func() zcl.ZdoCommand { return new(ActiveEndpointResponse) },
	MatchDescRequestCommand:              func() zcl.ZdoCommand { return new(MatchDescRequest) },
	MatchDescResponseCommand:             func() zcl.ZdoCommand { return new(MatchDescResponse) },
	ComplexDescRequestCommand:            func() zcl.ZdoCommand { return new(ComplexDescRequest) },
	ComplexDescResponseCommand:           func() zcl.ZdoCommand { return new(ComplexDescResponse) },
	UserDescRequestCommand:               func() zcl.ZdoCommand { return new(UserDescRequest) },
	DiscoveryCacheRequestCommand:         func() zcl.ZdoCommand { return new(DiscoveryCacheRequest) },
	DeviceAnnounceCommand:                func() zcl.ZdoCommand { return new(DeviceAnnounce) },
	UserDescSetRequestCommand:            func() zcl.ZdoCommand { return new(UserDescSetRequest) },
	SystemServerDiscoverRequestCommand:   func() zcl.ZdoCommand { return new(SystemServerDiscoverRequest) },
	DiscoveryStoreRequestCommand:         func() zcl.ZdoCommand { return new(DiscoveryStoreRequest) },
	NodeDescStoreRequestCommand:          func() zcl.ZdoCommand { return new(NodeDescStoreRequest) },
	PowerDescStoreRequestCommand:         func() zcl.ZdoCommand { return new(PowerDescStoreRequest) },
	ActiveEndpointStoreRequestCommand:    func() zcl.ZdoCommand { return new(ActiveEndpointStoreRequest) },
	SimpleDescStoreRequestCommand:        func() zcl.ZdoCommand { return new(SimpleDescStoreRequest) },
	RemoveNodeCacheRequestCommand:        func() zcl.ZdoCommand { return new(RemoveNodeCacheRequest) },
	FindNodeCacheRequestCommand:          func() zcl.ZdoCommand { return new(FindNodeCacheRequest) },
	ExtendedSimpleDescRequestCommand:     func() zcl.ZdoCommand { return new(ExtendedSimpleDescRequest) },
	ExtendedActiveEndpointRequestCommand: func() zcl.ZdoCommand { return new(ExtendedActiveEndpointRequest) },
	EndDeviceBindRequestCommand:          func() zcl.ZdoCommand { return new(EndDeviceBindRequest) },
	BindRequestCommand:                   func() zcl.ZdoCommand { return new(BindRequest) },
	BindResponseCommand:                  func() zcl.ZdoCommand { return new(BindResponse) },
	UnbindRequestCommand:                 func() zcl.ZdoCommand { return new(UnbindRequest) },
	UnbindResponseCommand:                func() zcl.ZdoCommand { return new(UnbindResponse) },
	MgmtLqiRequestCommand:                func() zcl.ZdoCommand { return new(MgmtLqiRequest) },
	MgmtLqiResponseCommand:               func() zcl.ZdoCommand { return new(MgmtLqiResponse) },
}

func (ApsFlags) Name() string { return "APS Flags" }

type ApsFlags zcl.Zu8

func (a *ApsFlags) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (a *ApsFlags) Value() zcl.Val     { return a }

func (a ApsFlags) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *ApsFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ApsFlags(*nt)
	return br, err
}

func (a ApsFlags) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(a))
}

func (a *ApsFlags) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = ApsFlags(*v)
	return nil
}

func (a *ApsFlags) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = ApsFlags(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ApsFlags) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

func (ActiveEndpointList) Name() string { return "Active Endpoint List" }

// ActiveEndpointList List of active endpoints
type ActiveEndpointList []*zcl.Zu8

func (a *ActiveEndpointList) TypeID() zcl.TypeID { return new(zcl.Zset).TypeID() }
func (a *ActiveEndpointList) Value() zcl.Val     { return a }

func (ActiveEndpointList) ArrayTypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }

func (a *ActiveEndpointList) ArrayValues() (o []uint8) {
	for _, v := range *a {
		o = append(o, uint8(*v))
	}
	return o
}

func (a *ActiveEndpointList) SetValues(val []uint8) error {
	*a = []*zcl.Zu8{}
	return a.AddValues(val...)
}

func (a *ActiveEndpointList) AddValues(val ...uint8) error {
	for _, v := range val {
		nv := zcl.Zu8(v)
		*a = append(*a, &nv)
	}
	return nil
}

func (a ActiveEndpointList) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, v := range a {
		vars = append(vars, v)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (a *ActiveEndpointList) UnmarshalZcl(b []byte) ([]byte, error) {
	*a = []*zcl.Zu8{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(zcl.Zu8)
		*a = append(*a, nv)
		return nv
	})
}

func (a *ActiveEndpointList) SetValue(v zcl.Val) error {
	if nv, ok := v.(*ActiveEndpointList); ok {
		*a = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ActiveEndpointList) String() string {
	var s []string
	for _, v := range a {
		s = append(s, zcl.Sprintf("%v", v))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

func (ActiveEndpointSize) Name() string { return "Active Endpoint Size" }

// ActiveEndpointSize Size in bytes of the Active Endpoints List
type ActiveEndpointSize zcl.Zu8

func (a *ActiveEndpointSize) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (a *ActiveEndpointSize) Value() zcl.Val     { return a }

func (a ActiveEndpointSize) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *ActiveEndpointSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ActiveEndpointSize(*nt)
	return br, err
}

func (a ActiveEndpointSize) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(a))
}

func (a *ActiveEndpointSize) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = ActiveEndpointSize(*v)
	return nil
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

func (AddressMode) Name() string { return "Address Mode" }

type AddressMode zcl.Zenum8

func (a *AddressMode) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (a *AddressMode) Value() zcl.Val     { return a }

func (a AddressMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *AddressMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = AddressMode(*nt)
	return br, err
}

func (a AddressMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(a))
}

func (a *AddressMode) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = AddressMode(*v)
	return nil
}

func (a *AddressMode) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = AddressMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a AddressMode) String() string {
	switch a {
	case 0x01:
		return "Group"
	case 0x02:
		return "NWK"
	case 0x03:
		return "IEEE"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a AddressMode) IsGroup() bool { return a == 0x01 }
func (a AddressMode) IsNwk() bool   { return a == 0x02 }
func (a AddressMode) IsIeee() bool  { return a == 0x03 }
func (a *AddressMode) SetGroup()    { *a = 0x01 }
func (a *AddressMode) SetNwk()      { *a = 0x02 }
func (a *AddressMode) SetIeee()     { *a = 0x03 }

func (AddressMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x01, Name: "Group"},
		{Value: 0x02, Name: "NWK"},
		{Value: 0x03, Name: "IEEE"},
	}
}

func (AssociatedDevices) Name() string { return "Associated Devices" }

type AssociatedDevices []*zcl.Zu16

func (a *AssociatedDevices) TypeID() zcl.TypeID { return new(zcl.Zlist).TypeID() }
func (a *AssociatedDevices) Value() zcl.Val     { return a }

func (AssociatedDevices) ArrayTypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }

func (a *AssociatedDevices) ArrayValues() (o []uint16) {
	for _, v := range *a {
		o = append(o, uint16(*v))
	}
	return o
}

func (a *AssociatedDevices) SetValues(val []uint16) error {
	*a = []*zcl.Zu16{}
	return a.AddValues(val...)
}

func (a *AssociatedDevices) AddValues(val ...uint16) error {
	for _, v := range val {
		nv := zcl.Zu16(v)
		*a = append(*a, &nv)
	}
	return nil
}

func (a AssociatedDevices) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, v := range a {
		vars = append(vars, v)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (a *AssociatedDevices) UnmarshalZcl(b []byte) ([]byte, error) {
	*a = []*zcl.Zu16{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(zcl.Zu16)
		*a = append(*a, nv)
		return nv
	})
}

func (a *AssociatedDevices) SetValue(v zcl.Val) error {
	if nv, ok := v.(*AssociatedDevices); ok {
		*a = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (a AssociatedDevices) String() string {
	var s []string
	for _, v := range a {
		s = append(s, zcl.Sprintf("%v", v))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

func (BindingTarget) Name() string { return "Binding Target" }

// BindingTarget NWK Address
type BindingTarget zcl.Zu16

func (a *BindingTarget) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *BindingTarget) Value() zcl.Val     { return a }

func (a BindingTarget) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *BindingTarget) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = BindingTarget(*nt)
	return br, err
}

func (a BindingTarget) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *BindingTarget) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = BindingTarget(*v)
	return nil
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

func (Capability) Name() string { return "Capability" }

// Capability specifies the device:s capabilities
type Capability zcl.Zbmp8

func (a *Capability) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (a *Capability) Value() zcl.Val     { return a }

func (a Capability) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *Capability) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = Capability(*nt)
	return br, err
}

func (a Capability) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(a))
}

func (a *Capability) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = Capability(*v)
	return nil
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

func (Capability) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Pan Coordinator"},
		{Value: 1, Name: "Full Function"},
		{Value: 2, Name: "Mains Power"},
		{Value: 3, Name: "RX on Idle"},
		{Value: 6, Name: "Security"},
		{Value: 7, Name: "Allocate Address"},
	}
}

func (ClusterId) Name() string { return "Cluster ID" }

type ClusterId zcl.Zu16

func (a *ClusterId) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *ClusterId) Value() zcl.Val     { return a }

func (a ClusterId) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ClusterId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ClusterId(*nt)
	return br, err
}

func (a ClusterId) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *ClusterId) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = ClusterId(*v)
	return nil
}

func (a *ClusterId) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ClusterId(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ClusterId) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

func (ComplexDescriptor) Name() string { return "Complex Descriptor" }

type ComplexDescriptor zcl.Zostring

func (a *ComplexDescriptor) TypeID() zcl.TypeID { return new(zcl.Zostring).TypeID() }
func (a *ComplexDescriptor) Value() zcl.Val     { return a }

func (a ComplexDescriptor) MarshalZcl() ([]byte, error) { return zcl.Zostring(a).MarshalZcl() }

func (a *ComplexDescriptor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*a = ComplexDescriptor(*nt)
	return br, err
}

func (a ComplexDescriptor) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zostring(a))
}

func (a *ComplexDescriptor) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zostring)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = ComplexDescriptor(*v)
	return nil
}

func (a *ComplexDescriptor) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zostring); ok {
		*a = ComplexDescriptor(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ComplexDescriptor) String() string {
	return zcl.Sprintf("%v", zcl.Zostring(a))
}

func (ComplexDescriptorAvailable) Name() string { return "Complex Descriptor Available" }

type ComplexDescriptorAvailable zcl.Zbool

func (a *ComplexDescriptorAvailable) TypeID() zcl.TypeID { return new(zcl.Zbool).TypeID() }
func (a *ComplexDescriptorAvailable) Value() zcl.Val     { return a }

func (a ComplexDescriptorAvailable) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *ComplexDescriptorAvailable) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = ComplexDescriptorAvailable(*nt)
	return br, err
}

func (a ComplexDescriptorAvailable) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbool(a))
}

func (a *ComplexDescriptorAvailable) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zbool)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = ComplexDescriptorAvailable(*v)
	return nil
}

func (a *ComplexDescriptorAvailable) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbool); ok {
		*a = ComplexDescriptorAvailable(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ComplexDescriptorAvailable) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

func (Depth) Name() string { return "Depth" }

// Depth of the neighbor device. A value of 0 indicates that the device is the coordinator for the network
type Depth zcl.Zu8

func (a *Depth) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (a *Depth) Value() zcl.Val     { return a }

func (a Depth) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Depth) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Depth(*nt)
	return br, err
}

func (a Depth) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(a))
}

func (a *Depth) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = Depth(*v)
	return nil
}

func (a *Depth) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = Depth(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Depth) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

func (DescriptorCapability) Name() string { return "Descriptor Capability" }

type DescriptorCapability zcl.Zbmp8

func (a *DescriptorCapability) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (a *DescriptorCapability) Value() zcl.Val     { return a }

func (a DescriptorCapability) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *DescriptorCapability) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = DescriptorCapability(*nt)
	return br, err
}

func (a DescriptorCapability) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(a))
}

func (a *DescriptorCapability) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = DescriptorCapability(*v)
	return nil
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

func (DescriptorCapability) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Extended Active Endpoint List Available"},
		{Value: 1, Name: "Extended Simple Descriptor List Available"},
	}
}

func (DeviceType) Name() string { return "Device Type" }

type DeviceType zcl.Zenum16

func (a *DeviceType) TypeID() zcl.TypeID { return new(zcl.Zenum16).TypeID() }
func (a *DeviceType) Value() zcl.Val     { return a }

func (a DeviceType) MarshalZcl() ([]byte, error) { return zcl.Zenum16(a).MarshalZcl() }

func (a *DeviceType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum16)
	br, err := nt.UnmarshalZcl(b)
	*a = DeviceType(*nt)
	return br, err
}

func (a DeviceType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum16(a))
}

func (a *DeviceType) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zenum16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = DeviceType(*v)
	return nil
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

func (DeviceType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x0000, Name: "On/Off Switch"},
		{Value: 0x0001, Name: "Level Control Switch"},
		{Value: 0x0002, Name: "On/Off Output"},
		{Value: 0x0003, Name: "Level Controllable Output"},
		{Value: 0x0004, Name: "Scene Selector"},
		{Value: 0x0005, Name: "Configuration Tool"},
		{Value: 0x0006, Name: "Remote Control"},
		{Value: 0x0007, Name: "Combined Interface"},
		{Value: 0x0008, Name: "Range Extender"},
		{Value: 0x0009, Name: "Mains Power Outlet"},
		{Value: 0x000a, Name: "Door lock"},
		{Value: 0x000c, Name: "Simple sensor"},
		{Value: 0x0010, Name: "On/off plug-in unit"},
		{Value: 0x0051, Name: "Smart plug"},
		{Value: 0x0060, Name: "GP Proxy"},
		{Value: 0x0061, Name: "GP Proxy Basic"},
		{Value: 0x0062, Name: "GP Target Plus"},
		{Value: 0x0063, Name: "GP Target"},
		{Value: 0x0064, Name: "GP Commissioning Tool"},
		{Value: 0x0065, Name: "GP Combo"},
		{Value: 0x0066, Name: "GP Combo Basic"},
		{Value: 0x0100, Name: "On/Off Light"},
		{Value: 0x0101, Name: "Dimmable Light"},
		{Value: 0x0102, Name: "Color Dimmable Light"},
		{Value: 0x0103, Name: "On/Off Light Switch"},
		{Value: 0x0104, Name: "Dimmer Switch"},
		{Value: 0x0105, Name: "Color Dimmer Switch"},
		{Value: 0x0106, Name: "Light Sensor"},
		{Value: 0x0107, Name: "Occupancy Sensor"},
		{Value: 0x0108, Name: "On/off ballast"},
		{Value: 0x0109, Name: "Dimmable ballast"},
		{Value: 0x010a, Name: "On/off plugin unit"},
		{Value: 0x010b, Name: "Dimmable plugin unit"},
		{Value: 0x010c, Name: "Color temperature light"},
		{Value: 0x010d, Name: "Extended color light"},
		{Value: 0x010e, Name: "Light level sensor"},
		{Value: 0x0110, Name: "Dimmable plug-in unit"},
		{Value: 0x0200, Name: "Shade"},
		{Value: 0x0201, Name: "Shade Controller"},
		{Value: 0x0202, Name: "Window Covering Device"},
		{Value: 0x0210, Name: "Extended color light 2"},
		{Value: 0x0220, Name: "Color temperature light 2"},
		{Value: 0x0300, Name: "Heating/Cooling Unit"},
		{Value: 0x0301, Name: "Thermostat"},
		{Value: 0x0302, Name: "Temperature Sensor"},
		{Value: 0x0303, Name: "Pump"},
		{Value: 0x0304, Name: "Pump Controller"},
		{Value: 0x0305, Name: "Pressure Sensor"},
		{Value: 0x0306, Name: "Flow Sensor"},
		{Value: 0x0400, Name: "IAS Control and Indicating Equipment"},
		{Value: 0x0401, Name: "IAS Ancillary Control Equipment"},
		{Value: 0x0402, Name: "IAS Zone"},
		{Value: 0x0403, Name: "IAS Warning Device"},
		{Value: 0x0500, Name: "Energy Service Portal"},
		{Value: 0x0501, Name: "Metering Device"},
		{Value: 0x0502, Name: "In-Premise Display"},
		{Value: 0x0503, Name: "Programmable Communicating Thermostat (PCT)"},
		{Value: 0x0504, Name: "Load Control Device"},
		{Value: 0x0505, Name: "Smart Appliance"},
		{Value: 0x0506, Name: "Prepayment Terminal"},
		{Value: 0x0507, Name: "Device Management"},
		{Value: 0x0800, Name: "Color controller"},
		{Value: 0x0810, Name: "Color scene controller"},
		{Value: 0x0820, Name: "Non color controller"},
		{Value: 0x0830, Name: "Non color scene controller"},
		{Value: 0x0840, Name: "Control bridge"},
		{Value: 0x0850, Name: "On/off sensor"},
	}
}

func (DeviceVersion) Name() string { return "Device Version" }

// DeviceVersion is dependant on profile
type DeviceVersion zcl.Zu8

func (a *DeviceVersion) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (a *DeviceVersion) Value() zcl.Val     { return a }

func (a DeviceVersion) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *DeviceVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = DeviceVersion(*nt)
	return br, err
}

func (a DeviceVersion) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(a))
}

func (a *DeviceVersion) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = DeviceVersion(*v)
	return nil
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

func (Endpoint) Name() string { return "Endpoint" }

type Endpoint zcl.Zu8

func (a *Endpoint) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (a *Endpoint) Value() zcl.Val     { return a }

func (a Endpoint) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Endpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Endpoint(*nt)
	return br, err
}

func (a Endpoint) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(a))
}

func (a *Endpoint) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = Endpoint(*v)
	return nil
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

func (EndpointList) Name() string { return "Endpoint List" }

type EndpointList []*zcl.Zu8

func (a *EndpointList) TypeID() zcl.TypeID { return new(zcl.Zset).TypeID() }
func (a *EndpointList) Value() zcl.Val     { return a }

func (EndpointList) ArrayTypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }

func (a *EndpointList) ArrayValues() (o []uint8) {
	for _, v := range *a {
		o = append(o, uint8(*v))
	}
	return o
}

func (a *EndpointList) SetValues(val []uint8) error {
	*a = []*zcl.Zu8{}
	return a.AddValues(val...)
}

func (a *EndpointList) AddValues(val ...uint8) error {
	for _, v := range val {
		nv := zcl.Zu8(v)
		*a = append(*a, &nv)
	}
	return nil
}

func (a EndpointList) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, v := range a {
		vars = append(vars, v)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (a *EndpointList) UnmarshalZcl(b []byte) ([]byte, error) {
	*a = []*zcl.Zu8{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(zcl.Zu8)
		*a = append(*a, nv)
		return nv
	})
}

func (a *EndpointList) SetValue(v zcl.Val) error {
	if nv, ok := v.(*EndpointList); ok {
		*a = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (a EndpointList) String() string {
	var s []string
	for _, v := range a {
		s = append(s, zcl.Sprintf("%v", v))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

func (FrequencyBands) Name() string { return "Frequency Bands" }

type FrequencyBands zcl.Zbmp8

func (a *FrequencyBands) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (a *FrequencyBands) Value() zcl.Val     { return a }

func (a FrequencyBands) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *FrequencyBands) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = FrequencyBands(*nt)
	return br, err
}

func (a FrequencyBands) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(a))
}

func (a *FrequencyBands) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = FrequencyBands(*v)
	return nil
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
			bstr = append(bstr, "868 MHz")
		case 2:
			bstr = append(bstr, "902-928 MHz")
		case 3:
			bstr = append(bstr, "2.4 GHz")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a FrequencyBands) Is868Mhz() bool       { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a FrequencyBands) Is902928Mhz() bool    { return zcl.BitmapTest([]byte(a[:]), 2) }
func (a FrequencyBands) Is24Ghz() bool        { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a *FrequencyBands) Set868Mhz(b bool)    { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b)) }
func (a *FrequencyBands) Set902928Mhz(b bool) { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b)) }
func (a *FrequencyBands) Set24Ghz(b bool)     { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b)) }

func (FrequencyBands) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "868 MHz"},
		{Value: 2, Name: "902-928 MHz"},
		{Value: 3, Name: "2.4 GHz"},
	}
}

func (IeeeAddress) Name() string { return "IEEE Address" }

// IeeeAddress is a 64-bit MAC address
type IeeeAddress zcl.Zuid

func (a *IeeeAddress) TypeID() zcl.TypeID { return new(zcl.Zuid).TypeID() }
func (a *IeeeAddress) Value() zcl.Val     { return a }

func (a IeeeAddress) MarshalZcl() ([]byte, error) { return zcl.Zuid(a).MarshalZcl() }

func (a *IeeeAddress) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*a = IeeeAddress(*nt)
	return br, err
}

func (a IeeeAddress) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zuid(a))
}

func (a *IeeeAddress) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zuid)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = IeeeAddress(*v)
	return nil
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

func (InClusterList) Name() string { return "In Cluster List" }

// InClusterList is a list of input clusters
type InClusterList []*zcl.Zu16

func (a *InClusterList) TypeID() zcl.TypeID { return new(zcl.Zset).TypeID() }
func (a *InClusterList) Value() zcl.Val     { return a }

func (InClusterList) ArrayTypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }

func (a *InClusterList) ArrayValues() (o []uint16) {
	for _, v := range *a {
		o = append(o, uint16(*v))
	}
	return o
}

func (a *InClusterList) SetValues(val []uint16) error {
	*a = []*zcl.Zu16{}
	return a.AddValues(val...)
}

func (a *InClusterList) AddValues(val ...uint16) error {
	for _, v := range val {
		nv := zcl.Zu16(v)
		*a = append(*a, &nv)
	}
	return nil
}

func (a InClusterList) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, v := range a {
		vars = append(vars, v)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (a *InClusterList) UnmarshalZcl(b []byte) ([]byte, error) {
	*a = []*zcl.Zu16{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(zcl.Zu16)
		*a = append(*a, nv)
		return nv
	})
}

func (a *InClusterList) SetValue(v zcl.Val) error {
	if nv, ok := v.(*InClusterList); ok {
		*a = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (a InClusterList) String() string {
	var s []string
	for _, v := range a {
		s = append(s, zcl.Sprintf("%v", v))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

func (LogicalType) Name() string { return "Logical Type" }

type LogicalType zcl.Zenum8

func (a *LogicalType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (a *LogicalType) Value() zcl.Val     { return a }

func (a LogicalType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *LogicalType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = LogicalType(*nt)
	return br, err
}

func (a LogicalType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(a))
}

func (a *LogicalType) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = LogicalType(*v)
	return nil
}

func (a *LogicalType) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = LogicalType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LogicalType) String() string {
	switch a {
	case 0x00:
		return "Coordinator"
	case 0x01:
		return "Router"
	case 0x10:
		return "End Device"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a LogicalType) IsCoordinator() bool { return a == 0x00 }
func (a LogicalType) IsRouter() bool      { return a == 0x01 }
func (a LogicalType) IsEndDevice() bool   { return a == 0x10 }
func (a *LogicalType) SetCoordinator()    { *a = 0x00 }
func (a *LogicalType) SetRouter()         { *a = 0x01 }
func (a *LogicalType) SetEndDevice()      { *a = 0x10 }

func (LogicalType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Coordinator"},
		{Value: 0x01, Name: "Router"},
		{Value: 0x10, Name: "End Device"},
	}
}

func (Lqi) Name() string { return "Lqi" }

// Lqi is the estimated link quality for RF transmissions from this device
type Lqi zcl.Zu8

func (a *Lqi) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (a *Lqi) Value() zcl.Val     { return a }

func (a Lqi) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Lqi) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Lqi(*nt)
	return br, err
}

func (a Lqi) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(a))
}

func (a *Lqi) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = Lqi(*v)
	return nil
}

func (a *Lqi) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = Lqi(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Lqi) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

func (MacCapabilityFlags) Name() string { return "MAC Capability Flags" }

type MacCapabilityFlags zcl.Zbmp8

func (a *MacCapabilityFlags) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (a *MacCapabilityFlags) Value() zcl.Val     { return a }

func (a MacCapabilityFlags) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *MacCapabilityFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = MacCapabilityFlags(*nt)
	return br, err
}

func (a MacCapabilityFlags) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(a))
}

func (a *MacCapabilityFlags) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = MacCapabilityFlags(*v)
	return nil
}

func (a *MacCapabilityFlags) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = MacCapabilityFlags(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MacCapabilityFlags) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Alternate PAN Coordinator")
		case 1:
			bstr = append(bstr, "Full function device")
		case 2:
			bstr = append(bstr, "Mains powered")
		case 3:
			bstr = append(bstr, "Receiver on when idle")
		case 6:
			bstr = append(bstr, "Supports secured frames")
		case 7:
			bstr = append(bstr, "Allocate Address")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a MacCapabilityFlags) IsAlternatePanCoordinator() bool { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a MacCapabilityFlags) IsFullFunctionDevice() bool      { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a MacCapabilityFlags) IsMainsPowered() bool            { return zcl.BitmapTest([]byte(a[:]), 2) }
func (a MacCapabilityFlags) IsReceiverOnWhenIdle() bool      { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a MacCapabilityFlags) IsSupportsSecuredFrames() bool   { return zcl.BitmapTest([]byte(a[:]), 6) }
func (a MacCapabilityFlags) IsAllocateAddress() bool         { return zcl.BitmapTest([]byte(a[:]), 7) }
func (a *MacCapabilityFlags) SetAlternatePanCoordinator(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *MacCapabilityFlags) SetFullFunctionDevice(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}
func (a *MacCapabilityFlags) SetMainsPowered(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b))
}
func (a *MacCapabilityFlags) SetReceiverOnWhenIdle(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b))
}
func (a *MacCapabilityFlags) SetSupportsSecuredFrames(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 6, b))
}
func (a *MacCapabilityFlags) SetAllocateAddress(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 7, b))
}

func (MacCapabilityFlags) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Alternate PAN Coordinator"},
		{Value: 1, Name: "Full function device"},
		{Value: 2, Name: "Mains powered"},
		{Value: 3, Name: "Receiver on when idle"},
		{Value: 6, Name: "Supports secured frames"},
		{Value: 7, Name: "Allocate Address"},
	}
}

func (ManufacturerCode) Name() string { return "Manufacturer Code" }

type ManufacturerCode zcl.Zenum16

func (a *ManufacturerCode) TypeID() zcl.TypeID { return new(zcl.Zenum16).TypeID() }
func (a *ManufacturerCode) Value() zcl.Val     { return a }

func (a ManufacturerCode) MarshalZcl() ([]byte, error) { return zcl.Zenum16(a).MarshalZcl() }

func (a *ManufacturerCode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum16)
	br, err := nt.UnmarshalZcl(b)
	*a = ManufacturerCode(*nt)
	return br, err
}

func (a ManufacturerCode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum16(a))
}

func (a *ManufacturerCode) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zenum16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = ManufacturerCode(*v)
	return nil
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
		return "NXP"
	case 0x10F2:
		return "Ubisys"
	case 0x115F:
		return "Xiaomi"
	case 0x117C:
		return "Ikea"
	}
	return zcl.Sprintf("%v", zcl.Zenum16(a))
}

func (a ManufacturerCode) IsPhilips() bool { return a == 0x100B }
func (a ManufacturerCode) IsLegrand() bool { return a == 0x1021 }
func (a ManufacturerCode) IsNxp() bool     { return a == 0x1037 }
func (a ManufacturerCode) IsUbisys() bool  { return a == 0x10F2 }
func (a ManufacturerCode) IsXiaomi() bool  { return a == 0x115F }
func (a ManufacturerCode) IsIkea() bool    { return a == 0x117C }
func (a *ManufacturerCode) SetPhilips()    { *a = 0x100B }
func (a *ManufacturerCode) SetLegrand()    { *a = 0x1021 }
func (a *ManufacturerCode) SetNxp()        { *a = 0x1037 }
func (a *ManufacturerCode) SetUbisys()     { *a = 0x10F2 }
func (a *ManufacturerCode) SetXiaomi()     { *a = 0x115F }
func (a *ManufacturerCode) SetIkea()       { *a = 0x117C }

func (ManufacturerCode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x100B, Name: "Philips"},
		{Value: 0x1021, Name: "Legrand"},
		{Value: 0x1037, Name: "NXP"},
		{Value: 0x10F2, Name: "Ubisys"},
		{Value: 0x115F, Name: "Xiaomi"},
		{Value: 0x117C, Name: "Ikea"},
	}
}

func (MaxBufferSize) Name() string { return "Max Buffer Size" }

// MaxBufferSize specifies the maximum size, in octets, of the network sub-layer data unit (NSDU) for this node.
// This is the maximum size of data or commands passed to or from the application by the application support sub-layer,
// before any fragmentation or re-assembly.
type MaxBufferSize zcl.Zu8

func (a *MaxBufferSize) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (a *MaxBufferSize) Value() zcl.Val     { return a }

func (a MaxBufferSize) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *MaxBufferSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxBufferSize(*nt)
	return br, err
}

func (a MaxBufferSize) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(a))
}

func (a *MaxBufferSize) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = MaxBufferSize(*v)
	return nil
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

func (MaxRxSize) Name() string { return "Max RX size" }

// MaxRxSize specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred to this node in one single message transfer.
type MaxRxSize zcl.Zu16

func (a *MaxRxSize) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *MaxRxSize) Value() zcl.Val     { return a }

func (a MaxRxSize) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *MaxRxSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxRxSize(*nt)
	return br, err
}

func (a MaxRxSize) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *MaxRxSize) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = MaxRxSize(*v)
	return nil
}

func (a *MaxRxSize) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = MaxRxSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MaxRxSize) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

func (MaxTxSize) Name() string { return "Max TX size" }

// MaxTxSize specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred from
// this node in one single message transfer.
type MaxTxSize zcl.Zu16

func (a *MaxTxSize) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *MaxTxSize) Value() zcl.Val     { return a }

func (a MaxTxSize) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *MaxTxSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxTxSize(*nt)
	return br, err
}

func (a MaxTxSize) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *MaxTxSize) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = MaxTxSize(*v)
	return nil
}

func (a *MaxTxSize) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = MaxTxSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MaxTxSize) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

func (NwkAddress) Name() string { return "NWK Address" }

// NwkAddress is a 16-bit Network address
type NwkAddress zcl.Zu16

func (a *NwkAddress) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *NwkAddress) Value() zcl.Val     { return a }

func (a NwkAddress) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *NwkAddress) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NwkAddress(*nt)
	return br, err
}

func (a NwkAddress) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *NwkAddress) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = NwkAddress(*v)
	return nil
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

func (NeighborTable) Name() string { return "Neighbor Table" }

type NeighborTable []*NeighborTableEntry

func (a *NeighborTable) TypeID() zcl.TypeID { return new(zcl.Zarray).TypeID() }
func (a *NeighborTable) Value() zcl.Val     { return a }

func (NeighborTable) ArrayTypeID() zcl.TypeID { return new(NeighborTableEntry).TypeID() }

func (a *NeighborTable) ArrayValues() (o []NeighborTableEntry) {
	for _, v := range *a {
		o = append(o, *v)
	}
	return o
}

func (a *NeighborTable) SetValues(val []NeighborTableEntry) error {
	*a = []*NeighborTableEntry{}
	return a.AddValues(val...)
}

func (a *NeighborTable) AddValues(val ...NeighborTableEntry) error {
	for _, v := range val {
		nv := v
		*a = append(*a, &nv)
	}
	return nil
}

func (a NeighborTable) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, v := range a {
		vars = append(vars, v)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (a *NeighborTable) UnmarshalZcl(b []byte) ([]byte, error) {
	*a = []*NeighborTableEntry{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(NeighborTableEntry)
		*a = append(*a, nv)
		return nv
	})
}

func (a *NeighborTable) SetValue(v zcl.Val) error {
	if nv, ok := v.(*NeighborTable); ok {
		*a = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NeighborTable) String() string {
	var s []string
	for _, v := range a {
		s = append(s, zcl.Sprintf("%v", v))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

func (NeighborTableEntry) Name() string { return "Neighbor Table Entry" }

type NeighborTableEntry struct {
	// PanId is a 64-bit MAC address
	PanId IeeeAddress
	// IeeeAddress is a 64-bit MAC address
	IeeeAddress IeeeAddress
	// NwkAddress is a 16-bit Network address
	NwkAddress    NwkAddress
	NeighborType  NeighborType
	RxOnWhenIdle  RxOnWhenIdle
	Relationship  Relationship
	PermitJoining PermitJoining
	// Depth of the neighbor device. A value of 0 indicates that the device is the coordinator for the network
	Depth Depth
	// Lqi is the estimated link quality for RF transmissions from this device
	Lqi Lqi
}

func (a *NeighborTableEntry) TypeID() zcl.TypeID { return zcl.TypeID(76) } // struct
func (a *NeighborTableEntry) Value() zcl.Val     { return a }
func (a NeighborTableEntry) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	var err error
	_ = tmp2
	_ = err

	{
		if tmp, err = a.PanId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = a.IeeeAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = a.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		tmp = []byte{}
		tmp2 = uint32(a.NeighborType&0x03) << 6

		data = append(data, tmp...)
	}
	{
		tmp2 |= uint32(a.RxOnWhenIdle&0x03) << 4

		data = append(data, tmp...)
	}
	{
		tmp2 |= uint32(a.Relationship&0x07) << 1
		tmp = []byte{uint8(tmp2)}

		data = append(data, tmp...)
	}
	{
		tmp2 = uint32(a.PermitJoining&0x03) << 6
		tmp = []byte{uint8(tmp2)}

		data = append(data, tmp...)
	}
	{
		if tmp, err = a.Depth.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = a.Lqi.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (a *NeighborTableEntry) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2
	_ = err

	if b, err = (&a.PanId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&a.IeeeAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&a.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if len(b) == 0 {
		return b, zcl.ErrNotEnoughData
	}
	a.NeighborType = NeighborType((b[0] >> 6) & 0x03)

	a.RxOnWhenIdle = RxOnWhenIdle((b[0] >> 4) & 0x03)

	a.Relationship = Relationship((b[0] >> 1) & 0x07)
	b = b[1:]

	if len(b) == 0 {
		return b, zcl.ErrNotEnoughData
	}
	a.PermitJoining = PermitJoining((b[0] >> 6) & 0x03)
	b = b[1:]

	if b, err = (&a.Depth).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&a.Lqi).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (a *NeighborTableEntry) SetValue(v zcl.Val) error {
	if nv, ok := v.(*NeighborTableEntry); ok {
		*a = *nv
		return nil
	}
	return zcl.ErrNotImpl
}

func (a *NeighborTableEntry) String() string {
	return zcl.Sprintf(
		"NeighborTableEntry{"+zcl.StrJoin([]string{
			"PanId(%v)",
			"IeeeAddress(%v)",
			"NwkAddress(%v)",
			"NeighborType(%v)",
			"RxOnWhenIdle(%v)",
			"Relationship(%v)",
			"PermitJoining(%v)",
			"Depth(%v)",
			"Lqi(%v)",
		}, " ")+"}",
		a.PanId,
		a.IeeeAddress,
		a.NwkAddress,
		a.NeighborType,
		a.RxOnWhenIdle,
		a.Relationship,
		a.PermitJoining,
		a.Depth,
		a.Lqi,
	)
}

func (NeighborType) Name() string { return "Neighbor Type" }

type NeighborType zcl.Zenum8

func (a *NeighborType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (a *NeighborType) Value() zcl.Val     { return a }

func (a NeighborType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *NeighborType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = NeighborType(*nt)
	return br, err
}

func (a NeighborType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(a))
}

func (a *NeighborType) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = NeighborType(*v)
	return nil
}

func (a *NeighborType) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = NeighborType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NeighborType) String() string {
	switch a {
	case 0x00:
		return "Coordinator"
	case 0x01:
		return "Router"
	case 0x03:
		return "End device"
	case 0x04:
		return "Unknown"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a NeighborType) IsCoordinator() bool { return a == 0x00 }
func (a NeighborType) IsRouter() bool      { return a == 0x01 }
func (a NeighborType) IsEndDevice() bool   { return a == 0x03 }
func (a NeighborType) IsUnknown() bool     { return a == 0x04 }
func (a *NeighborType) SetCoordinator()    { *a = 0x00 }
func (a *NeighborType) SetRouter()         { *a = 0x01 }
func (a *NeighborType) SetEndDevice()      { *a = 0x03 }
func (a *NeighborType) SetUnknown()        { *a = 0x04 }

func (NeighborType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Coordinator"},
		{Value: 0x01, Name: "Router"},
		{Value: 0x03, Name: "End device"},
		{Value: 0x04, Name: "Unknown"},
	}
}

func (NodeDescSize) Name() string { return "Node Desc Size" }

// NodeDescSize Size in bytes of the Node Descriptor
type NodeDescSize zcl.Zu8

func (a *NodeDescSize) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (a *NodeDescSize) Value() zcl.Val     { return a }

func (a NodeDescSize) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *NodeDescSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NodeDescSize(*nt)
	return br, err
}

func (a NodeDescSize) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(a))
}

func (a *NodeDescSize) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = NodeDescSize(*v)
	return nil
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

func (NodeDescriptor) Name() string { return "Node Descriptor" }

type NodeDescriptor struct {
	LogicalType                LogicalType
	ComplexDescriptorAvailable ComplexDescriptorAvailable
	UserDescriptorAvailable    UserDescriptorAvailable
	ApsFlags                   ApsFlags
	FrequencyBands             FrequencyBands
	MacCapabilityFlags         MacCapabilityFlags
	ManufacturerCode           ManufacturerCode
	// MaxBufferSize specifies the maximum size, in octets, of the network sub-layer data unit (NSDU) for this node.
	// This is the maximum size of data or commands passed to or from the application by the application support sub-layer,
	// before any fragmentation or re-assembly.
	MaxBufferSize MaxBufferSize
	// MaxRxSize specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred to this node in one single message transfer.
	MaxRxSize  MaxRxSize
	ServerMask ServerMask
	// MaxTxSize specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred from
	// this node in one single message transfer.
	MaxTxSize            MaxTxSize
	DescriptorCapability DescriptorCapability
}

func (a *NodeDescriptor) TypeID() zcl.TypeID { return zcl.TypeID(76) } // struct
func (a *NodeDescriptor) Value() zcl.Val     { return a }
func (a NodeDescriptor) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	var err error
	_ = tmp2
	_ = err

	{
		tmp = []byte{}
		tmp2 = uint32(a.LogicalType&0x07) << 5

		data = append(data, tmp...)
	}
	{
		tmp2 |= uint32(a.ComplexDescriptorAvailable&0x01) << 4

		data = append(data, tmp...)
	}
	{
		tmp2 |= uint32(a.UserDescriptorAvailable&0x01) << 3
		tmp = []byte{uint8(tmp2)}

		data = append(data, tmp...)
	}
	{
		tmp = []byte{}
		tmp2 = uint32(a.ApsFlags&0x07) << 5

		data = append(data, tmp...)
	}
	{
		tmp2 |= uint32(a.FrequencyBands[0] & 0x1F)
		tmp = []byte{uint8(tmp2)}

		data = append(data, tmp...)
	}
	{
		if tmp, err = a.MacCapabilityFlags.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = a.ManufacturerCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = a.MaxBufferSize.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = a.MaxRxSize.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = a.ServerMask.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = a.MaxTxSize.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = a.DescriptorCapability.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (a *NodeDescriptor) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2
	_ = err

	if len(b) == 0 {
		return b, zcl.ErrNotEnoughData
	}
	a.LogicalType = LogicalType((b[0] >> 5) & 0x07)

	a.ComplexDescriptorAvailable = ComplexDescriptorAvailable((b[0] >> 4) & 0x01)

	a.UserDescriptorAvailable = UserDescriptorAvailable((b[0] >> 3) & 0x01)
	b = b[1:]

	if len(b) == 0 {
		return b, zcl.ErrNotEnoughData
	}
	a.ApsFlags = ApsFlags((b[0] >> 5) & 0x07)

	a.FrequencyBands[0] = b[0] & 0x1F
	b = b[1:]

	if b, err = (&a.MacCapabilityFlags).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&a.ManufacturerCode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&a.MaxBufferSize).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&a.MaxRxSize).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&a.ServerMask).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&a.MaxTxSize).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&a.DescriptorCapability).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (a *NodeDescriptor) SetValue(v zcl.Val) error {
	if nv, ok := v.(*NodeDescriptor); ok {
		*a = *nv
		return nil
	}
	return zcl.ErrNotImpl
}

func (a *NodeDescriptor) String() string {
	return zcl.Sprintf(
		"NodeDescriptor{"+zcl.StrJoin([]string{
			"LogicalType(%v)",
			"ComplexDescriptorAvailable(%v)",
			"UserDescriptorAvailable(%v)",
			"ApsFlags(%v)",
			"FrequencyBands(%v)",
			"MacCapabilityFlags(%v)",
			"ManufacturerCode(%v)",
			"MaxBufferSize(%v)",
			"MaxRxSize(%v)",
			"ServerMask(%v)",
			"MaxTxSize(%v)",
			"DescriptorCapability(%v)",
		}, " ")+"}",
		a.LogicalType,
		a.ComplexDescriptorAvailable,
		a.UserDescriptorAvailable,
		a.ApsFlags,
		a.FrequencyBands,
		a.MacCapabilityFlags,
		a.ManufacturerCode,
		a.MaxBufferSize,
		a.MaxRxSize,
		a.ServerMask,
		a.MaxTxSize,
		a.DescriptorCapability,
	)
}

func (OutClusterList) Name() string { return "Out Cluster List" }

// OutClusterList is a list of output clusters
type OutClusterList []*zcl.Zu16

func (a *OutClusterList) TypeID() zcl.TypeID { return new(zcl.Zset).TypeID() }
func (a *OutClusterList) Value() zcl.Val     { return a }

func (OutClusterList) ArrayTypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }

func (a *OutClusterList) ArrayValues() (o []uint16) {
	for _, v := range *a {
		o = append(o, uint16(*v))
	}
	return o
}

func (a *OutClusterList) SetValues(val []uint16) error {
	*a = []*zcl.Zu16{}
	return a.AddValues(val...)
}

func (a *OutClusterList) AddValues(val ...uint16) error {
	for _, v := range val {
		nv := zcl.Zu16(v)
		*a = append(*a, &nv)
	}
	return nil
}

func (a OutClusterList) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, v := range a {
		vars = append(vars, v)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (a *OutClusterList) UnmarshalZcl(b []byte) ([]byte, error) {
	*a = []*zcl.Zu16{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(zcl.Zu16)
		*a = append(*a, nv)
		return nv
	})
}

func (a *OutClusterList) SetValue(v zcl.Val) error {
	if nv, ok := v.(*OutClusterList); ok {
		*a = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OutClusterList) String() string {
	var s []string
	for _, v := range a {
		s = append(s, zcl.Sprintf("%v", v))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

func (PermitJoining) Name() string { return "Permit Joining" }

type PermitJoining zcl.Zenum8

func (a *PermitJoining) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (a *PermitJoining) Value() zcl.Val     { return a }

func (a PermitJoining) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *PermitJoining) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = PermitJoining(*nt)
	return br, err
}

func (a PermitJoining) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(a))
}

func (a *PermitJoining) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = PermitJoining(*v)
	return nil
}

func (a *PermitJoining) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = PermitJoining(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PermitJoining) String() string {
	switch a {
	case 0x00:
		return "Not permitted"
	case 0x01:
		return "Permitted"
	case 0x02:
		return "Unknown"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a PermitJoining) IsNotPermitted() bool { return a == 0x00 }
func (a PermitJoining) IsPermitted() bool    { return a == 0x01 }
func (a PermitJoining) IsUnknown() bool      { return a == 0x02 }
func (a *PermitJoining) SetNotPermitted()    { *a = 0x00 }
func (a *PermitJoining) SetPermitted()       { *a = 0x01 }
func (a *PermitJoining) SetUnknown()         { *a = 0x02 }

func (PermitJoining) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Not permitted"},
		{Value: 0x01, Name: "Permitted"},
		{Value: 0x02, Name: "Unknown"},
	}
}

func (PowerDescSize) Name() string { return "Power Desc Size" }

// PowerDescSize Size in bytes of the Power Descriptor
type PowerDescSize zcl.Zu8

func (a *PowerDescSize) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (a *PowerDescSize) Value() zcl.Val     { return a }

func (a PowerDescSize) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *PowerDescSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerDescSize(*nt)
	return br, err
}

func (a PowerDescSize) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(a))
}

func (a *PowerDescSize) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = PowerDescSize(*v)
	return nil
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

func (PowerDescriptor) Name() string { return "Power Descriptor" }

type PowerDescriptor struct {
	PowerMode          PowerMode
	ActivePowerSource  PowerSource
	CurrentPowerSource PowerSource
	PowerLevel         PowerLevel
}

func (a *PowerDescriptor) TypeID() zcl.TypeID { return zcl.TypeID(76) } // struct
func (a *PowerDescriptor) Value() zcl.Val     { return a }
func (a PowerDescriptor) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	var err error
	_ = tmp2
	_ = err

	{
		tmp = []byte{}
		tmp2 = uint32(a.PowerMode[0]&0x0F) << 4

		data = append(data, tmp...)
	}
	{
		tmp2 |= uint32(a.ActivePowerSource[0] & 0x0F)
		tmp = []byte{uint8(tmp2)}

		data = append(data, tmp...)
	}
	{
		tmp = []byte{}
		tmp2 = uint32(a.CurrentPowerSource[0]&0x0F) << 4

		data = append(data, tmp...)
	}
	{
		tmp2 |= uint32(a.PowerLevel & 0x0F)
		tmp = []byte{uint8(tmp2)}

		data = append(data, tmp...)
	}
	return data, nil
}

func (a *PowerDescriptor) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2
	_ = err

	if len(b) == 0 {
		return b, zcl.ErrNotEnoughData
	}
	a.PowerMode[0] = (b[0] >> 4) & 0x0F

	a.ActivePowerSource[0] = b[0] & 0x0F
	b = b[1:]

	if len(b) == 0 {
		return b, zcl.ErrNotEnoughData
	}
	a.CurrentPowerSource[0] = (b[0] >> 4) & 0x0F

	a.PowerLevel = PowerLevel(b[0] & 0x0F)
	b = b[1:]

	return b, nil
}

func (a *PowerDescriptor) SetValue(v zcl.Val) error {
	if nv, ok := v.(*PowerDescriptor); ok {
		*a = *nv
		return nil
	}
	return zcl.ErrNotImpl
}

func (a *PowerDescriptor) String() string {
	return zcl.Sprintf(
		"PowerDescriptor{"+zcl.StrJoin([]string{
			"PowerMode(%v)",
			"ActivePowerSource(%v)",
			"CurrentPowerSource(%v)",
			"PowerLevel(%v)",
		}, " ")+"}",
		a.PowerMode,
		a.ActivePowerSource,
		a.CurrentPowerSource,
		a.PowerLevel,
	)
}

func (PowerLevel) Name() string { return "Power Level" }

type PowerLevel zcl.Zenum8

func (a *PowerLevel) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (a *PowerLevel) Value() zcl.Val     { return a }

func (a PowerLevel) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *PowerLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerLevel(*nt)
	return br, err
}

func (a PowerLevel) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(a))
}

func (a *PowerLevel) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = PowerLevel(*v)
	return nil
}

func (a *PowerLevel) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = PowerLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PowerLevel) String() string {
	switch a {
	case 0x00:
		return "Critical"
	case 0x04:
		return "33%"
	case 0x08:
		return "66%"
	case 0x0C:
		return "100%"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a PowerLevel) IsCritical() bool { return a == 0x00 }
func (a PowerLevel) Is33() bool       { return a == 0x04 }
func (a PowerLevel) Is66() bool       { return a == 0x08 }
func (a PowerLevel) Is100() bool      { return a == 0x0C }
func (a *PowerLevel) SetCritical()    { *a = 0x00 }
func (a *PowerLevel) Set33()          { *a = 0x04 }
func (a *PowerLevel) Set66()          { *a = 0x08 }
func (a *PowerLevel) Set100()         { *a = 0x0C }

func (PowerLevel) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Critical"},
		{Value: 0x04, Name: "33%"},
		{Value: 0x08, Name: "66%"},
		{Value: 0x0C, Name: "100%"},
	}
}

func (PowerMode) Name() string { return "Power Mode" }

type PowerMode zcl.Zbmp8

func (a *PowerMode) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (a *PowerMode) Value() zcl.Val     { return a }

func (a PowerMode) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *PowerMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerMode(*nt)
	return br, err
}

func (a PowerMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(a))
}

func (a *PowerMode) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = PowerMode(*v)
	return nil
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

func (PowerMode) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Constant Power Available"},
		{Value: 1, Name: "Rechargeable battery Available"},
		{Value: 2, Name: "Disposable Battery Available"},
		{Value: 4, Name: "Check In Periodically"},
		{Value: 5, Name: "Check In on Action"},
	}
}

func (PowerSource) Name() string { return "Power Source" }

type PowerSource zcl.Zbmp8

func (a *PowerSource) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (a *PowerSource) Value() zcl.Val     { return a }

func (a PowerSource) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *PowerSource) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerSource(*nt)
	return br, err
}

func (a PowerSource) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(a))
}

func (a *PowerSource) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = PowerSource(*v)
	return nil
}

func (a *PowerSource) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = PowerSource(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PowerSource) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Mains power")
		case 1:
			bstr = append(bstr, "Rechargeable battery")
		case 2:
			bstr = append(bstr, "Disposable battery")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a PowerSource) IsMainsPower() bool          { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a PowerSource) IsRechargeableBattery() bool { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a PowerSource) IsDisposableBattery() bool   { return zcl.BitmapTest([]byte(a[:]), 2) }
func (a *PowerSource) SetMainsPower(b bool)       { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b)) }
func (a *PowerSource) SetRechargeableBattery(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}
func (a *PowerSource) SetDisposableBattery(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b))
}

func (PowerSource) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Mains power"},
		{Value: 1, Name: "Rechargeable battery"},
		{Value: 2, Name: "Disposable battery"},
	}
}

func (ProfileId) Name() string { return "Profile ID" }

type ProfileId zcl.Zu16

func (a *ProfileId) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *ProfileId) Value() zcl.Val     { return a }

func (a ProfileId) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ProfileId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ProfileId(*nt)
	return br, err
}

func (a ProfileId) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *ProfileId) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = ProfileId(*v)
	return nil
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

func (Relationship) Name() string { return "Relationship" }

type Relationship zcl.Zenum8

func (a *Relationship) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (a *Relationship) Value() zcl.Val     { return a }

func (a Relationship) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *Relationship) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = Relationship(*nt)
	return br, err
}

func (a Relationship) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(a))
}

func (a *Relationship) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = Relationship(*v)
	return nil
}

func (a *Relationship) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = Relationship(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Relationship) String() string {
	switch a {
	case 0x00:
		return "Parent"
	case 0x01:
		return "Child"
	case 0x02:
		return "Sibling"
	case 0x03:
		return "None"
	case 0x04:
		return "Previous Child"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a Relationship) IsParent() bool        { return a == 0x00 }
func (a Relationship) IsChild() bool         { return a == 0x01 }
func (a Relationship) IsSibling() bool       { return a == 0x02 }
func (a Relationship) IsNone() bool          { return a == 0x03 }
func (a Relationship) IsPreviousChild() bool { return a == 0x04 }
func (a *Relationship) SetParent()           { *a = 0x00 }
func (a *Relationship) SetChild()            { *a = 0x01 }
func (a *Relationship) SetSibling()          { *a = 0x02 }
func (a *Relationship) SetNone()             { *a = 0x03 }
func (a *Relationship) SetPreviousChild()    { *a = 0x04 }

func (Relationship) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Parent"},
		{Value: 0x01, Name: "Child"},
		{Value: 0x02, Name: "Sibling"},
		{Value: 0x03, Name: "None"},
		{Value: 0x04, Name: "Previous Child"},
	}
}

func (RequestType) Name() string { return "Request Type" }

// RequestType should be set to 1 if extended response is requested, 0 otherwise
type RequestType zcl.Zenum8

func (a *RequestType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (a *RequestType) Value() zcl.Val     { return a }

func (a RequestType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *RequestType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = RequestType(*nt)
	return br, err
}

func (a RequestType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(a))
}

func (a *RequestType) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = RequestType(*v)
	return nil
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

func (RequestType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Single Device Response"},
		{Value: 0x01, Name: "Extended Response"},
	}
}

func (RxOnWhenIdle) Name() string { return "Rx On When Idle" }

type RxOnWhenIdle zcl.Zenum8

func (a *RxOnWhenIdle) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (a *RxOnWhenIdle) Value() zcl.Val     { return a }

func (a RxOnWhenIdle) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *RxOnWhenIdle) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = RxOnWhenIdle(*nt)
	return br, err
}

func (a RxOnWhenIdle) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(a))
}

func (a *RxOnWhenIdle) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = RxOnWhenIdle(*v)
	return nil
}

func (a *RxOnWhenIdle) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = RxOnWhenIdle(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a RxOnWhenIdle) String() string {
	switch a {
	case 0x00:
		return "Receiver is off"
	case 0x01:
		return "Receiver is on"
	case 0x02:
		return "Unknown"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a RxOnWhenIdle) IsReceiverIsOff() bool { return a == 0x00 }
func (a RxOnWhenIdle) IsReceiverIsOn() bool  { return a == 0x01 }
func (a RxOnWhenIdle) IsUnknown() bool       { return a == 0x02 }
func (a *RxOnWhenIdle) SetReceiverIsOff()    { *a = 0x00 }
func (a *RxOnWhenIdle) SetReceiverIsOn()     { *a = 0x01 }
func (a *RxOnWhenIdle) SetUnknown()          { *a = 0x02 }

func (RxOnWhenIdle) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Receiver is off"},
		{Value: 0x01, Name: "Receiver is on"},
		{Value: 0x02, Name: "Unknown"},
	}
}

func (ServerMask) Name() string { return "Server Mask" }

type ServerMask zcl.Zbmp16

func (a *ServerMask) TypeID() zcl.TypeID { return new(zcl.Zbmp16).TypeID() }
func (a *ServerMask) Value() zcl.Val     { return a }

func (a ServerMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp16(a).MarshalZcl() }

func (a *ServerMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*a = ServerMask(*nt)
	return br, err
}

func (a ServerMask) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp16(a))
}

func (a *ServerMask) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zbmp16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = ServerMask(*v)
	return nil
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

func (ServerMask) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Primary Trust Center"},
		{Value: 1, Name: "Backup Trust Center"},
		{Value: 2, Name: "Primary Binding Table Cache"},
		{Value: 3, Name: "Backup Binding Table Cache"},
		{Value: 4, Name: "Primary Discovery Cache"},
		{Value: 5, Name: "Backup Discovery Cache"},
		{Value: 6, Name: "Network Manager"},
	}
}

func (SimpleDescSizeList) Name() string { return "Simple Desc Size List" }

// SimpleDescSizeList List of sizes for the different Simple Descriptors
type SimpleDescSizeList []*zcl.Zu8

func (a *SimpleDescSizeList) TypeID() zcl.TypeID { return new(zcl.Zset).TypeID() }
func (a *SimpleDescSizeList) Value() zcl.Val     { return a }

func (SimpleDescSizeList) ArrayTypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }

func (a *SimpleDescSizeList) ArrayValues() (o []uint8) {
	for _, v := range *a {
		o = append(o, uint8(*v))
	}
	return o
}

func (a *SimpleDescSizeList) SetValues(val []uint8) error {
	*a = []*zcl.Zu8{}
	return a.AddValues(val...)
}

func (a *SimpleDescSizeList) AddValues(val ...uint8) error {
	for _, v := range val {
		nv := zcl.Zu8(v)
		*a = append(*a, &nv)
	}
	return nil
}

func (a SimpleDescSizeList) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, v := range a {
		vars = append(vars, v)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (a *SimpleDescSizeList) UnmarshalZcl(b []byte) ([]byte, error) {
	*a = []*zcl.Zu8{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(zcl.Zu8)
		*a = append(*a, nv)
		return nv
	})
}

func (a *SimpleDescSizeList) SetValue(v zcl.Val) error {
	if nv, ok := v.(*SimpleDescSizeList); ok {
		*a = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SimpleDescSizeList) String() string {
	var s []string
	for _, v := range a {
		s = append(s, zcl.Sprintf("%v", v))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

func (SimpleDescriptor) Name() string { return "Simple Descriptor" }

type SimpleDescriptor struct {
	Endpoint   Endpoint
	ProfileId  ProfileId
	DeviceType DeviceType
	// DeviceVersion is dependant on profile
	DeviceVersion DeviceVersion
	// InClusterList is a list of input clusters
	InClusterList InClusterList
	// OutClusterList is a list of output clusters
	OutClusterList OutClusterList
}

func (a *SimpleDescriptor) TypeID() zcl.TypeID { return zcl.TypeID(76) } // struct
func (a *SimpleDescriptor) Value() zcl.Val     { return a }
func (a SimpleDescriptor) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	var err error
	_ = tmp2
	_ = err

	{
		if tmp, err = a.Endpoint.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = a.ProfileId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = a.DeviceType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = a.DeviceVersion.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = a.InClusterList.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = a.OutClusterList.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (a *SimpleDescriptor) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2
	_ = err

	if b, err = (&a.Endpoint).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&a.ProfileId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&a.DeviceType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&a.DeviceVersion).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&a.InClusterList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&a.OutClusterList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (a *SimpleDescriptor) SetValue(v zcl.Val) error {
	if nv, ok := v.(*SimpleDescriptor); ok {
		*a = *nv
		return nil
	}
	return zcl.ErrNotImpl
}

func (a *SimpleDescriptor) String() string {
	return zcl.Sprintf(
		"SimpleDescriptor{"+zcl.StrJoin([]string{
			"Endpoint(%v)",
			"ProfileId(%v)",
			"DeviceType(%v)",
			"DeviceVersion(%v)",
			"InClusterList(%v)",
			"OutClusterList(%v)",
		}, " ")+"}",
		a.Endpoint,
		a.ProfileId,
		a.DeviceType,
		a.DeviceVersion,
		a.InClusterList,
		a.OutClusterList,
	)
}

func (SourceAddress) Name() string { return "Source Address" }

// SourceAddress of device generating the request
type SourceAddress zcl.Zuid

func (a *SourceAddress) TypeID() zcl.TypeID { return new(zcl.Zuid).TypeID() }
func (a *SourceAddress) Value() zcl.Val     { return a }

func (a SourceAddress) MarshalZcl() ([]byte, error) { return zcl.Zuid(a).MarshalZcl() }

func (a *SourceAddress) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*a = SourceAddress(*nt)
	return br, err
}

func (a SourceAddress) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zuid(a))
}

func (a *SourceAddress) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zuid)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = SourceAddress(*v)
	return nil
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

func (SourceEndpoint) Name() string { return "Source Endpoint" }

// SourceEndpoint of device generating the request
type SourceEndpoint zcl.Zu8

func (a *SourceEndpoint) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (a *SourceEndpoint) Value() zcl.Val     { return a }

func (a SourceEndpoint) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *SourceEndpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = SourceEndpoint(*nt)
	return br, err
}

func (a SourceEndpoint) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(a))
}

func (a *SourceEndpoint) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = SourceEndpoint(*v)
	return nil
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

func (StartIndex) Name() string { return "Start Index" }

// StartIndex provides the starting index for the requested elements of the associated list.
type StartIndex zcl.Zu8

func (a *StartIndex) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (a *StartIndex) Value() zcl.Val     { return a }

func (a StartIndex) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *StartIndex) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = StartIndex(*nt)
	return br, err
}

func (a StartIndex) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(a))
}

func (a *StartIndex) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = StartIndex(*v)
	return nil
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

func (Status) Name() string { return "Status" }

// Status Code, command is normally empty unless status is `Success`
type Status zcl.Zenum8

func (a *Status) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (a *Status) Value() zcl.Val     { return a }

func (a Status) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *Status) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = Status(*nt)
	return br, err
}

func (a Status) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(a))
}

func (a *Status) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = Status(*v)
	return nil
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

func (Status) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Success"},
		{Value: 0x80, Name: "Invalid Request Type"},
		{Value: 0x81, Name: "Device Not Found"},
		{Value: 0x82, Name: "Invalid Endpoint"},
		{Value: 0x83, Name: "Not Active"},
		{Value: 0x84, Name: "Not Supported"},
		{Value: 0x85, Name: "Timeout"},
		{Value: 0x86, Name: "No Match"},
		{Value: 0x88, Name: "No Entry"},
		{Value: 0x89, Name: "No Descriptor"},
		{Value: 0x8A, Name: "Insufficient Space"},
		{Value: 0x8B, Name: "Not Permitted"},
		{Value: 0x8C, Name: "Table Full"},
		{Value: 0x8D, Name: "Not Authorized"},
	}
}

func (TotalEntries) Name() string { return "Total Entries" }

// TotalEntries is the total number of entries that can be queried for
type TotalEntries zcl.Zu8

func (a *TotalEntries) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (a *TotalEntries) Value() zcl.Val     { return a }

func (a TotalEntries) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *TotalEntries) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = TotalEntries(*nt)
	return br, err
}

func (a TotalEntries) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(a))
}

func (a *TotalEntries) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = TotalEntries(*v)
	return nil
}

func (a *TotalEntries) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = TotalEntries(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a TotalEntries) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

func (UnknownU8) Name() string { return "Unknown u8" }

type UnknownU8 zcl.Zu8

func (a *UnknownU8) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (a *UnknownU8) Value() zcl.Val     { return a }

func (a UnknownU8) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *UnknownU8) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = UnknownU8(*nt)
	return br, err
}

func (a UnknownU8) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(a))
}

func (a *UnknownU8) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = UnknownU8(*v)
	return nil
}

func (a *UnknownU8) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = UnknownU8(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a UnknownU8) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

func (UserDesc) Name() string { return "User Desc" }

// UserDesc is a user provided ASCII string of 16 characters set on a remote device.
// If the string is shorter than 16 characters it should be padded with space characters (0x20).
// Control characters (0x00-0x1f) are not allowed.
type UserDesc zcl.Zcstring

func (a *UserDesc) TypeID() zcl.TypeID { return new(zcl.Zcstring).TypeID() }
func (a *UserDesc) Value() zcl.Val     { return a }

func (a UserDesc) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *UserDesc) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = UserDesc(*nt)
	return br, err
}

func (a UserDesc) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zcstring(a))
}

func (a *UserDesc) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zcstring)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = UserDesc(*v)
	return nil
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

func (UserDescriptorAvailable) Name() string { return "User Descriptor Available" }

type UserDescriptorAvailable zcl.Zbool

func (a *UserDescriptorAvailable) TypeID() zcl.TypeID { return new(zcl.Zbool).TypeID() }
func (a *UserDescriptorAvailable) Value() zcl.Val     { return a }

func (a UserDescriptorAvailable) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *UserDescriptorAvailable) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = UserDescriptorAvailable(*nt)
	return br, err
}

func (a UserDescriptorAvailable) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbool(a))
}

func (a *UserDescriptorAvailable) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zbool)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = UserDescriptorAvailable(*v)
	return nil
}

func (a *UserDescriptorAvailable) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbool); ok {
		*a = UserDescriptorAvailable(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a UserDescriptorAvailable) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(a))
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

// Arguments returns all values of NwkAddressRequest
func (v *NwkAddressRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "IeeeAddress", Argument: &v.IeeeAddress},
		{Name: "RequestType", Argument: &v.RequestType},
		{Name: "StartIndex", Argument: &v.StartIndex},
	}
}

// Name of the command
func (NwkAddressRequest) Name() string { return "NWK Address Request" }

// ID of the command
func (NwkAddressRequest) ID() CommandID { return NwkAddressRequestCommand }

// Required
func (NwkAddressRequest) Required() bool { return true }

// Cluster ID of the command
func (NwkAddressRequest) Cluster() zcl.ClusterID { return 0x0000 }

// MnfCode returns the manufacturer code (if any) of the command
func (NwkAddressRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (NwkAddressRequest) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

// MarshalZcl returns the wire format representation of NwkAddressRequest
func (v NwkAddressRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

// Arguments returns all values of NwkAddressResponse
func (v *NwkAddressResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "IeeeAddress", Argument: &v.IeeeAddress},
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "StartIndex", Argument: &v.StartIndex},
		{Name: "AssociatedDevices", Argument: &v.AssociatedDevices},
	}
}

// Name of the command
func (NwkAddressResponse) Name() string { return "NWK Address Response" }

// ID of the command
func (NwkAddressResponse) ID() CommandID { return NwkAddressResponseCommand }

// Required
func (NwkAddressResponse) Required() bool { return true }

// Cluster ID of the command
func (NwkAddressResponse) Cluster() zcl.ClusterID { return 0x8000 }

// MnfCode returns the manufacturer code (if any) of the command
func (NwkAddressResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (NwkAddressResponse) MarshalJSON() ([]byte, error) { return []byte("32768"), nil }

// MarshalZcl returns the wire format representation of NwkAddressResponse
func (v NwkAddressResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
		tmp = []byte{uint8(2 * len(v.AssociatedDevices))}

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
	tmp2 := uint32(0)
	_ = tmp2

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
		if len(b) == 0 {
			return b, nil
		}
		nv := new(zcl.Zu8)
		b, _ = nv.UnmarshalZcl(b)
		tmp2 = uint32(*nv)

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

// Arguments returns all values of IeeeAddressRequest
func (v *IeeeAddressRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "RequestType", Argument: &v.RequestType},
		{Name: "StartIndex", Argument: &v.StartIndex},
	}
}

// Name of the command
func (IeeeAddressRequest) Name() string { return "IEEE Address Request" }

// ID of the command
func (IeeeAddressRequest) ID() CommandID { return IeeeAddressRequestCommand }

// Required
func (IeeeAddressRequest) Required() bool { return true }

// Cluster ID of the command
func (IeeeAddressRequest) Cluster() zcl.ClusterID { return 0x0001 }

// MnfCode returns the manufacturer code (if any) of the command
func (IeeeAddressRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (IeeeAddressRequest) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

// MarshalZcl returns the wire format representation of IeeeAddressRequest
func (v IeeeAddressRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

type IeeeAddressResponse struct {
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

// IeeeAddressResponseCommand is the Command ID of IeeeAddressResponse
const IeeeAddressResponseCommand CommandID = 0x8001

// Values returns all values of IeeeAddressResponse
func (v *IeeeAddressResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.IeeeAddress,
		&v.NwkAddress,
		&v.StartIndex,
		&v.AssociatedDevices,
	}
}

// Arguments returns all values of IeeeAddressResponse
func (v *IeeeAddressResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "IeeeAddress", Argument: &v.IeeeAddress},
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "StartIndex", Argument: &v.StartIndex},
		{Name: "AssociatedDevices", Argument: &v.AssociatedDevices},
	}
}

// Name of the command
func (IeeeAddressResponse) Name() string { return "IEEE Address Response" }

// ID of the command
func (IeeeAddressResponse) ID() CommandID { return IeeeAddressResponseCommand }

// Required
func (IeeeAddressResponse) Required() bool { return true }

// Cluster ID of the command
func (IeeeAddressResponse) Cluster() zcl.ClusterID { return 0x8001 }

// MnfCode returns the manufacturer code (if any) of the command
func (IeeeAddressResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (IeeeAddressResponse) MarshalJSON() ([]byte, error) { return []byte("32769"), nil }

// MarshalZcl returns the wire format representation of IeeeAddressResponse
func (v IeeeAddressResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
		tmp = []byte{uint8(2 * len(v.AssociatedDevices))}

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
		// Ignore length, it was added earlier
		tmp = tmp[1:]

		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the IeeeAddressResponse struct
func (v *IeeeAddressResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

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
		if len(b) == 0 {
			return b, nil
		}
		nv := new(zcl.Zu8)
		b, _ = nv.UnmarshalZcl(b)
		tmp2 = uint32(*nv)

	}

	// StartIndex is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.StartIndex).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// AssociatedDevices is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.AssociatedDevices).UnmarshalZcl(append([]byte{uint8(tmp2)}, b...)); err != nil {
			return b, err
		}

	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v IeeeAddressResponse) String() string {
	return zcl.Sprintf(
		"IeeeAddressResponse{"+zcl.StrJoin([]string{
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

// Arguments returns all values of NodeDescRequest
func (v *NodeDescRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
	}
}

// Name of the command
func (NodeDescRequest) Name() string { return "Node Desc Request" }

// ID of the command
func (NodeDescRequest) ID() CommandID { return NodeDescRequestCommand }

// Required
func (NodeDescRequest) Required() bool { return true }

// Cluster ID of the command
func (NodeDescRequest) Cluster() zcl.ClusterID { return 0x0002 }

// MnfCode returns the manufacturer code (if any) of the command
func (NodeDescRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (NodeDescRequest) MarshalJSON() ([]byte, error) { return []byte("2"), nil }

// MarshalZcl returns the wire format representation of NodeDescRequest
func (v NodeDescRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

type NodeDescResponse struct {
	// Status Code, command is normally empty unless status is `Success`
	Status Status
	// NwkAddress is a 16-bit Network address
	NwkAddress     NwkAddress
	NodeDescriptor NodeDescriptor
}

// NodeDescResponseCommand is the Command ID of NodeDescResponse
const NodeDescResponseCommand CommandID = 0x8002

// Values returns all values of NodeDescResponse
func (v *NodeDescResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.NwkAddress,
		&v.NodeDescriptor,
	}
}

// Arguments returns all values of NodeDescResponse
func (v *NodeDescResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "NodeDescriptor", Argument: &v.NodeDescriptor},
	}
}

// Name of the command
func (NodeDescResponse) Name() string { return "Node Desc Response" }

// ID of the command
func (NodeDescResponse) ID() CommandID { return NodeDescResponseCommand }

// Required
func (NodeDescResponse) Required() bool { return true }

// Cluster ID of the command
func (NodeDescResponse) Cluster() zcl.ClusterID { return 0x8002 }

// MnfCode returns the manufacturer code (if any) of the command
func (NodeDescResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (NodeDescResponse) MarshalJSON() ([]byte, error) { return []byte("32770"), nil }

// MarshalZcl returns the wire format representation of NodeDescResponse
func (v NodeDescResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
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
	{
		if tmp, err = v.NodeDescriptor.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the NodeDescResponse struct
func (v *NodeDescResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NodeDescriptor).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v NodeDescResponse) String() string {
	return zcl.Sprintf(
		"NodeDescResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"NwkAddress(%v)",
			"NodeDescriptor(%v)",
		}, " ")+"}",
		v.Status,
		v.NwkAddress,
		v.NodeDescriptor,
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

// Arguments returns all values of PowerDescRequest
func (v *PowerDescRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
	}
}

// Name of the command
func (PowerDescRequest) Name() string { return "Power Desc Request" }

// ID of the command
func (PowerDescRequest) ID() CommandID { return PowerDescRequestCommand }

// Required
func (PowerDescRequest) Required() bool { return true }

// Cluster ID of the command
func (PowerDescRequest) Cluster() zcl.ClusterID { return 0x0003 }

// MnfCode returns the manufacturer code (if any) of the command
func (PowerDescRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (PowerDescRequest) MarshalJSON() ([]byte, error) { return []byte("3"), nil }

// MarshalZcl returns the wire format representation of PowerDescRequest
func (v PowerDescRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

type PowerDescResponse struct {
}

// PowerDescResponseCommand is the Command ID of PowerDescResponse
const PowerDescResponseCommand CommandID = 0x8003

// Values returns all values of PowerDescResponse
func (v *PowerDescResponse) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of PowerDescResponse
func (v *PowerDescResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (PowerDescResponse) Name() string { return "Power Desc Response" }

// ID of the command
func (PowerDescResponse) ID() CommandID { return PowerDescResponseCommand }

// Required
func (PowerDescResponse) Required() bool { return true }

// Cluster ID of the command
func (PowerDescResponse) Cluster() zcl.ClusterID { return 0x8003 }

// MnfCode returns the manufacturer code (if any) of the command
func (PowerDescResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (PowerDescResponse) MarshalJSON() ([]byte, error) { return []byte("32771"), nil }

// MarshalZcl returns the wire format representation of PowerDescResponse
func (v PowerDescResponse) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the PowerDescResponse struct
func (v *PowerDescResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v PowerDescResponse) String() string {
	return zcl.Sprintf(
		"PowerDescResponse{" + zcl.StrJoin([]string{}, " ") + "}",
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

// Arguments returns all values of SimpleDescRequest
func (v *SimpleDescRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "Endpoint", Argument: &v.Endpoint},
	}
}

// Name of the command
func (SimpleDescRequest) Name() string { return "Simple Desc Request" }

// ID of the command
func (SimpleDescRequest) ID() CommandID { return SimpleDescRequestCommand }

// Required
func (SimpleDescRequest) Required() bool { return true }

// Cluster ID of the command
func (SimpleDescRequest) Cluster() zcl.ClusterID { return 0x0004 }

// MnfCode returns the manufacturer code (if any) of the command
func (SimpleDescRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (SimpleDescRequest) MarshalJSON() ([]byte, error) { return []byte("4"), nil }

// MarshalZcl returns the wire format representation of SimpleDescRequest
func (v SimpleDescRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

type SimpleDescResponse struct {
	// Status Code, command is normally empty unless status is `Success`
	Status Status
	// NwkAddress is a 16-bit Network address
	NwkAddress       NwkAddress
	SimpleDescriptor SimpleDescriptor
}

// SimpleDescResponseCommand is the Command ID of SimpleDescResponse
const SimpleDescResponseCommand CommandID = 0x8004

// Values returns all values of SimpleDescResponse
func (v *SimpleDescResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.NwkAddress,
		&v.SimpleDescriptor,
	}
}

// Arguments returns all values of SimpleDescResponse
func (v *SimpleDescResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "SimpleDescriptor", Argument: &v.SimpleDescriptor},
	}
}

// Name of the command
func (SimpleDescResponse) Name() string { return "Simple Desc Response" }

// ID of the command
func (SimpleDescResponse) ID() CommandID { return SimpleDescResponseCommand }

// Required
func (SimpleDescResponse) Required() bool { return false }

// Cluster ID of the command
func (SimpleDescResponse) Cluster() zcl.ClusterID { return 0x8004 }

// MnfCode returns the manufacturer code (if any) of the command
func (SimpleDescResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (SimpleDescResponse) MarshalJSON() ([]byte, error) { return []byte("32772"), nil }

// MarshalZcl returns the wire format representation of SimpleDescResponse
func (v SimpleDescResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
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
		sd, _ := v.SimpleDescriptor.MarshalZcl()
		tmp = []byte{uint8(len(sd))}

		data = append(data, tmp...)
	}
	// SimpleDescriptor is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.SimpleDescriptor.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the SimpleDescResponse struct
func (v *SimpleDescResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	// is only included if successful
	if v.Status == 0x00 {
		// Ignore descriptor length
		if b, err = new(zcl.Zu8).UnmarshalZcl(b); err != nil {
			return b, err
		}

	}

	// SimpleDescriptor is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.SimpleDescriptor).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v SimpleDescResponse) String() string {
	return zcl.Sprintf(
		"SimpleDescResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"NwkAddress(%v)",
			"SimpleDescriptor(%v)",
		}, " ")+"}",
		v.Status,
		v.NwkAddress,
		v.SimpleDescriptor,
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

// Arguments returns all values of ActiveEndpointRequest
func (v *ActiveEndpointRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
	}
}

// Name of the command
func (ActiveEndpointRequest) Name() string { return "Active Endpoint Request" }

// ID of the command
func (ActiveEndpointRequest) ID() CommandID { return ActiveEndpointRequestCommand }

// Required
func (ActiveEndpointRequest) Required() bool { return true }

// Cluster ID of the command
func (ActiveEndpointRequest) Cluster() zcl.ClusterID { return 0x0005 }

// MnfCode returns the manufacturer code (if any) of the command
func (ActiveEndpointRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ActiveEndpointRequest) MarshalJSON() ([]byte, error) { return []byte("5"), nil }

// MarshalZcl returns the wire format representation of ActiveEndpointRequest
func (v ActiveEndpointRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

type ActiveEndpointResponse struct {
	// Status Code, command is normally empty unless status is `Success`
	Status Status
	// NwkAddress is a 16-bit Network address
	NwkAddress   NwkAddress
	EndpointList EndpointList
}

// ActiveEndpointResponseCommand is the Command ID of ActiveEndpointResponse
const ActiveEndpointResponseCommand CommandID = 0x8005

// Values returns all values of ActiveEndpointResponse
func (v *ActiveEndpointResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.NwkAddress,
		&v.EndpointList,
	}
}

// Arguments returns all values of ActiveEndpointResponse
func (v *ActiveEndpointResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "EndpointList", Argument: &v.EndpointList},
	}
}

// Name of the command
func (ActiveEndpointResponse) Name() string { return "Active Endpoint Response" }

// ID of the command
func (ActiveEndpointResponse) ID() CommandID { return ActiveEndpointResponseCommand }

// Required
func (ActiveEndpointResponse) Required() bool { return true }

// Cluster ID of the command
func (ActiveEndpointResponse) Cluster() zcl.ClusterID { return 0x8005 }

// MnfCode returns the manufacturer code (if any) of the command
func (ActiveEndpointResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ActiveEndpointResponse) MarshalJSON() ([]byte, error) { return []byte("32773"), nil }

// MarshalZcl returns the wire format representation of ActiveEndpointResponse
func (v ActiveEndpointResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// NwkAddress is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// EndpointList is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.EndpointList.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ActiveEndpointResponse struct
func (v *ActiveEndpointResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	// NwkAddress is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// EndpointList is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.EndpointList).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ActiveEndpointResponse) String() string {
	return zcl.Sprintf(
		"ActiveEndpointResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"NwkAddress(%v)",
			"EndpointList(%v)",
		}, " ")+"}",
		v.Status,
		v.NwkAddress,
		v.EndpointList,
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

// Arguments returns all values of MatchDescRequest
func (v *MatchDescRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "ProfileId", Argument: &v.ProfileId},
		{Name: "InClusterList", Argument: &v.InClusterList},
		{Name: "OutClusterList", Argument: &v.OutClusterList},
	}
}

// Name of the command
func (MatchDescRequest) Name() string { return "Match Desc Request" }

// ID of the command
func (MatchDescRequest) ID() CommandID { return MatchDescRequestCommand }

// Required
func (MatchDescRequest) Required() bool { return true }

// Cluster ID of the command
func (MatchDescRequest) Cluster() zcl.ClusterID { return 0x0006 }

// MnfCode returns the manufacturer code (if any) of the command
func (MatchDescRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MatchDescRequest) MarshalJSON() ([]byte, error) { return []byte("6"), nil }

// MarshalZcl returns the wire format representation of MatchDescRequest
func (v MatchDescRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

type MatchDescResponse struct {
	// Status Code, command is normally empty unless status is `Success`
	Status       Status
	EndpointList EndpointList
}

// MatchDescResponseCommand is the Command ID of MatchDescResponse
const MatchDescResponseCommand CommandID = 0x8006

// Values returns all values of MatchDescResponse
func (v *MatchDescResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.EndpointList,
	}
}

// Arguments returns all values of MatchDescResponse
func (v *MatchDescResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "EndpointList", Argument: &v.EndpointList},
	}
}

// Name of the command
func (MatchDescResponse) Name() string { return "Match Desc Response" }

// ID of the command
func (MatchDescResponse) ID() CommandID { return MatchDescResponseCommand }

// Required
func (MatchDescResponse) Required() bool { return false }

// Cluster ID of the command
func (MatchDescResponse) Cluster() zcl.ClusterID { return 0x8006 }

// MnfCode returns the manufacturer code (if any) of the command
func (MatchDescResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MatchDescResponse) MarshalJSON() ([]byte, error) { return []byte("32774"), nil }

// MarshalZcl returns the wire format representation of MatchDescResponse
func (v MatchDescResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// EndpointList is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.EndpointList.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the MatchDescResponse struct
func (v *MatchDescResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	// EndpointList is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.EndpointList).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MatchDescResponse) String() string {
	return zcl.Sprintf(
		"MatchDescResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"EndpointList(%v)",
		}, " ")+"}",
		v.Status,
		v.EndpointList,
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

// Arguments returns all values of ComplexDescRequest
func (v *ComplexDescRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
	}
}

// Name of the command
func (ComplexDescRequest) Name() string { return "Complex Desc Request" }

// ID of the command
func (ComplexDescRequest) ID() CommandID { return ComplexDescRequestCommand }

// Required
func (ComplexDescRequest) Required() bool { return false }

// Cluster ID of the command
func (ComplexDescRequest) Cluster() zcl.ClusterID { return 0x0010 }

// MnfCode returns the manufacturer code (if any) of the command
func (ComplexDescRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ComplexDescRequest) MarshalJSON() ([]byte, error) { return []byte("16"), nil }

// MarshalZcl returns the wire format representation of ComplexDescRequest
func (v ComplexDescRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

type ComplexDescResponse struct {
	// Status Code, command is normally empty unless status is `Success`
	Status Status
	// NwkAddress is a 16-bit Network address
	NwkAddress        NwkAddress
	ComplexDescriptor ComplexDescriptor
}

// ComplexDescResponseCommand is the Command ID of ComplexDescResponse
const ComplexDescResponseCommand CommandID = 0x8010

// Values returns all values of ComplexDescResponse
func (v *ComplexDescResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.NwkAddress,
		&v.ComplexDescriptor,
	}
}

// Arguments returns all values of ComplexDescResponse
func (v *ComplexDescResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "ComplexDescriptor", Argument: &v.ComplexDescriptor},
	}
}

// Name of the command
func (ComplexDescResponse) Name() string { return "Complex Desc Response" }

// ID of the command
func (ComplexDescResponse) ID() CommandID { return ComplexDescResponseCommand }

// Required
func (ComplexDescResponse) Required() bool { return false }

// Cluster ID of the command
func (ComplexDescResponse) Cluster() zcl.ClusterID { return 0x8010 }

// MnfCode returns the manufacturer code (if any) of the command
func (ComplexDescResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ComplexDescResponse) MarshalJSON() ([]byte, error) { return []byte("32784"), nil }

// MarshalZcl returns the wire format representation of ComplexDescResponse
func (v ComplexDescResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// NwkAddress is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// ComplexDescriptor is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.ComplexDescriptor.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ComplexDescResponse struct
func (v *ComplexDescResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	// NwkAddress is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// ComplexDescriptor is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.ComplexDescriptor).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ComplexDescResponse) String() string {
	return zcl.Sprintf(
		"ComplexDescResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"NwkAddress(%v)",
			"ComplexDescriptor(%v)",
		}, " ")+"}",
		v.Status,
		v.NwkAddress,
		v.ComplexDescriptor,
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

// Arguments returns all values of UserDescRequest
func (v *UserDescRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
	}
}

// Name of the command
func (UserDescRequest) Name() string { return "User Desc Request" }

// ID of the command
func (UserDescRequest) ID() CommandID { return UserDescRequestCommand }

// Required
func (UserDescRequest) Required() bool { return false }

// Cluster ID of the command
func (UserDescRequest) Cluster() zcl.ClusterID { return 0x0011 }

// MnfCode returns the manufacturer code (if any) of the command
func (UserDescRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (UserDescRequest) MarshalJSON() ([]byte, error) { return []byte("17"), nil }

// MarshalZcl returns the wire format representation of UserDescRequest
func (v UserDescRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

// Arguments returns all values of DiscoveryCacheRequest
func (v *DiscoveryCacheRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "IeeeAddress", Argument: &v.IeeeAddress},
	}
}

// Name of the command
func (DiscoveryCacheRequest) Name() string { return "Discovery Cache Request" }

// ID of the command
func (DiscoveryCacheRequest) ID() CommandID { return DiscoveryCacheRequestCommand }

// Required
func (DiscoveryCacheRequest) Required() bool { return false }

// Cluster ID of the command
func (DiscoveryCacheRequest) Cluster() zcl.ClusterID { return 0x0012 }

// MnfCode returns the manufacturer code (if any) of the command
func (DiscoveryCacheRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (DiscoveryCacheRequest) MarshalJSON() ([]byte, error) { return []byte("18"), nil }

// MarshalZcl returns the wire format representation of DiscoveryCacheRequest
func (v DiscoveryCacheRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

// Arguments returns all values of DeviceAnnounce
func (v *DeviceAnnounce) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "IeeeAddress", Argument: &v.IeeeAddress},
		{Name: "Capability", Argument: &v.Capability},
	}
}

// Name of the command
func (DeviceAnnounce) Name() string { return "Device Announce" }

// ID of the command
func (DeviceAnnounce) ID() CommandID { return DeviceAnnounceCommand }

// Required
func (DeviceAnnounce) Required() bool { return true }

// Cluster ID of the command
func (DeviceAnnounce) Cluster() zcl.ClusterID { return 0x0013 }

// MnfCode returns the manufacturer code (if any) of the command
func (DeviceAnnounce) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (DeviceAnnounce) MarshalJSON() ([]byte, error) { return []byte("19"), nil }

// MarshalZcl returns the wire format representation of DeviceAnnounce
func (v DeviceAnnounce) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

// Arguments returns all values of UserDescSetRequest
func (v *UserDescSetRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "UserDesc", Argument: &v.UserDesc},
	}
}

// Name of the command
func (UserDescSetRequest) Name() string { return "User Desc Set Request" }

// ID of the command
func (UserDescSetRequest) ID() CommandID { return UserDescSetRequestCommand }

// Required
func (UserDescSetRequest) Required() bool { return false }

// Cluster ID of the command
func (UserDescSetRequest) Cluster() zcl.ClusterID { return 0x0014 }

// MnfCode returns the manufacturer code (if any) of the command
func (UserDescSetRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (UserDescSetRequest) MarshalJSON() ([]byte, error) { return []byte("20"), nil }

// MarshalZcl returns the wire format representation of UserDescSetRequest
func (v UserDescSetRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

// Arguments returns all values of SystemServerDiscoverRequest
func (v *SystemServerDiscoverRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "ServerMask", Argument: &v.ServerMask},
	}
}

// Name of the command
func (SystemServerDiscoverRequest) Name() string { return "System Server Discover Request" }

// ID of the command
func (SystemServerDiscoverRequest) ID() CommandID { return SystemServerDiscoverRequestCommand }

// Required
func (SystemServerDiscoverRequest) Required() bool { return false }

// Cluster ID of the command
func (SystemServerDiscoverRequest) Cluster() zcl.ClusterID { return 0x0015 }

// MnfCode returns the manufacturer code (if any) of the command
func (SystemServerDiscoverRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (SystemServerDiscoverRequest) MarshalJSON() ([]byte, error) { return []byte("21"), nil }

// MarshalZcl returns the wire format representation of SystemServerDiscoverRequest
func (v SystemServerDiscoverRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

// Arguments returns all values of DiscoveryStoreRequest
func (v *DiscoveryStoreRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "IeeeAddress", Argument: &v.IeeeAddress},
		{Name: "NodeDescSize", Argument: &v.NodeDescSize},
		{Name: "PowerDescSize", Argument: &v.PowerDescSize},
		{Name: "ActiveEndpointSize", Argument: &v.ActiveEndpointSize},
		{Name: "SimpleDescSizeList", Argument: &v.SimpleDescSizeList},
	}
}

// Name of the command
func (DiscoveryStoreRequest) Name() string { return "Discovery Store Request" }

// ID of the command
func (DiscoveryStoreRequest) ID() CommandID { return DiscoveryStoreRequestCommand }

// Required
func (DiscoveryStoreRequest) Required() bool { return false }

// Cluster ID of the command
func (DiscoveryStoreRequest) Cluster() zcl.ClusterID { return 0x0016 }

// MnfCode returns the manufacturer code (if any) of the command
func (DiscoveryStoreRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (DiscoveryStoreRequest) MarshalJSON() ([]byte, error) { return []byte("22"), nil }

// MarshalZcl returns the wire format representation of DiscoveryStoreRequest
func (v DiscoveryStoreRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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
	IeeeAddress IeeeAddress
	// Capability specifies the device:s capabilities
	Capability       Capability
	ManufacturerCode ManufacturerCode
	// MaxBufferSize specifies the maximum size, in octets, of the network sub-layer data unit (NSDU) for this node.
	// This is the maximum size of data or commands passed to or from the application by the application support sub-layer,
	// before any fragmentation or re-assembly.
	MaxBufferSize MaxBufferSize
	// MaxRxSize specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred to this node in one single message transfer.
	MaxRxSize  MaxRxSize
	ServerMask ServerMask
	// MaxTxSize specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred from
	// this node in one single message transfer.
	MaxTxSize MaxTxSize
}

// NodeDescStoreRequestCommand is the Command ID of NodeDescStoreRequest
const NodeDescStoreRequestCommand CommandID = 0x0017

// Values returns all values of NodeDescStoreRequest
func (v *NodeDescStoreRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.IeeeAddress,
		&v.Capability,
		&v.ManufacturerCode,
		&v.MaxBufferSize,
		&v.MaxRxSize,
		&v.ServerMask,
		&v.MaxTxSize,
	}
}

// Arguments returns all values of NodeDescStoreRequest
func (v *NodeDescStoreRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "IeeeAddress", Argument: &v.IeeeAddress},
		{Name: "Capability", Argument: &v.Capability},
		{Name: "ManufacturerCode", Argument: &v.ManufacturerCode},
		{Name: "MaxBufferSize", Argument: &v.MaxBufferSize},
		{Name: "MaxRxSize", Argument: &v.MaxRxSize},
		{Name: "ServerMask", Argument: &v.ServerMask},
		{Name: "MaxTxSize", Argument: &v.MaxTxSize},
	}
}

// Name of the command
func (NodeDescStoreRequest) Name() string { return "Node Desc Store Request" }

// ID of the command
func (NodeDescStoreRequest) ID() CommandID { return NodeDescStoreRequestCommand }

// Required
func (NodeDescStoreRequest) Required() bool { return false }

// Cluster ID of the command
func (NodeDescStoreRequest) Cluster() zcl.ClusterID { return 0x0017 }

// MnfCode returns the manufacturer code (if any) of the command
func (NodeDescStoreRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (NodeDescStoreRequest) MarshalJSON() ([]byte, error) { return []byte("23"), nil }

// MarshalZcl returns the wire format representation of NodeDescStoreRequest
func (v NodeDescStoreRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
		if tmp, err = v.MaxRxSize.MarshalZcl(); err != nil {
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
		if tmp, err = v.MaxTxSize.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the NodeDescStoreRequest struct
func (v *NodeDescStoreRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.IeeeAddress).UnmarshalZcl(b); err != nil {
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

	if b, err = (&v.MaxRxSize).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ServerMask).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.MaxTxSize).UnmarshalZcl(b); err != nil {
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
			"Capability(%v)",
			"ManufacturerCode(%v)",
			"MaxBufferSize(%v)",
			"MaxRxSize(%v)",
			"ServerMask(%v)",
			"MaxTxSize(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.IeeeAddress,
		v.Capability,
		v.ManufacturerCode,
		v.MaxBufferSize,
		v.MaxRxSize,
		v.ServerMask,
		v.MaxTxSize,
	)
}

// PowerDescStoreRequest sent to a Discovery Cache Node to store the provided Power Descriptor.
// A DiscoveryStoreRequest should be sent and acknowledged with successful status before requesting storage.
// Should be sent to the unicast address of a discovery cache device.
type PowerDescStoreRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	// IeeeAddress is a 64-bit MAC address
	IeeeAddress IeeeAddress
	PowerMode   PowerMode
	PowerSource PowerSource
}

// PowerDescStoreRequestCommand is the Command ID of PowerDescStoreRequest
const PowerDescStoreRequestCommand CommandID = 0x0018

// Values returns all values of PowerDescStoreRequest
func (v *PowerDescStoreRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.IeeeAddress,
		&v.PowerMode,
		&v.PowerSource,
	}
}

// Arguments returns all values of PowerDescStoreRequest
func (v *PowerDescStoreRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "IeeeAddress", Argument: &v.IeeeAddress},
		{Name: "PowerMode", Argument: &v.PowerMode},
		{Name: "PowerSource", Argument: &v.PowerSource},
	}
}

// Name of the command
func (PowerDescStoreRequest) Name() string { return "Power Desc Store Request" }

// ID of the command
func (PowerDescStoreRequest) ID() CommandID { return PowerDescStoreRequestCommand }

// Required
func (PowerDescStoreRequest) Required() bool { return false }

// Cluster ID of the command
func (PowerDescStoreRequest) Cluster() zcl.ClusterID { return 0x0018 }

// MnfCode returns the manufacturer code (if any) of the command
func (PowerDescStoreRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (PowerDescStoreRequest) MarshalJSON() ([]byte, error) { return []byte("24"), nil }

// MarshalZcl returns the wire format representation of PowerDescStoreRequest
func (v PowerDescStoreRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
		if tmp, err = v.PowerSource.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the PowerDescStoreRequest struct
func (v *PowerDescStoreRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.IeeeAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.PowerMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.PowerSource).UnmarshalZcl(b); err != nil {
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
			"PowerSource(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.IeeeAddress,
		v.PowerMode,
		v.PowerSource,
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

// Arguments returns all values of ActiveEndpointStoreRequest
func (v *ActiveEndpointStoreRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "IeeeAddress", Argument: &v.IeeeAddress},
		{Name: "EndpointList", Argument: &v.EndpointList},
	}
}

// Name of the command
func (ActiveEndpointStoreRequest) Name() string { return "Active Endpoint Store Request" }

// ID of the command
func (ActiveEndpointStoreRequest) ID() CommandID { return ActiveEndpointStoreRequestCommand }

// Required
func (ActiveEndpointStoreRequest) Required() bool { return false }

// Cluster ID of the command
func (ActiveEndpointStoreRequest) Cluster() zcl.ClusterID { return 0x0019 }

// MnfCode returns the manufacturer code (if any) of the command
func (ActiveEndpointStoreRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ActiveEndpointStoreRequest) MarshalJSON() ([]byte, error) { return []byte("25"), nil }

// MarshalZcl returns the wire format representation of ActiveEndpointStoreRequest
func (v ActiveEndpointStoreRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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
	IeeeAddress      IeeeAddress
	SimpleDescriptor SimpleDescriptor
}

// SimpleDescStoreRequestCommand is the Command ID of SimpleDescStoreRequest
const SimpleDescStoreRequestCommand CommandID = 0x001A

// Values returns all values of SimpleDescStoreRequest
func (v *SimpleDescStoreRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.IeeeAddress,
		&v.SimpleDescriptor,
	}
}

// Arguments returns all values of SimpleDescStoreRequest
func (v *SimpleDescStoreRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "IeeeAddress", Argument: &v.IeeeAddress},
		{Name: "SimpleDescriptor", Argument: &v.SimpleDescriptor},
	}
}

// Name of the command
func (SimpleDescStoreRequest) Name() string { return "Simple Desc Store Request" }

// ID of the command
func (SimpleDescStoreRequest) ID() CommandID { return SimpleDescStoreRequestCommand }

// Required
func (SimpleDescStoreRequest) Required() bool { return false }

// Cluster ID of the command
func (SimpleDescStoreRequest) Cluster() zcl.ClusterID { return 0x001A }

// MnfCode returns the manufacturer code (if any) of the command
func (SimpleDescStoreRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (SimpleDescStoreRequest) MarshalJSON() ([]byte, error) { return []byte("26"), nil }

// MarshalZcl returns the wire format representation of SimpleDescStoreRequest
func (v SimpleDescStoreRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
		if tmp, err = v.SimpleDescriptor.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, uint8(len(tmp)))

		data = append(data, tmp...)
	}
	{
		// Marshalled above
		tmp = []byte{}

		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the SimpleDescStoreRequest struct
func (v *SimpleDescStoreRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.IeeeAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = new(zcl.Zu8).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SimpleDescriptor).UnmarshalZcl(b); err != nil {
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
			"SimpleDescriptor(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.IeeeAddress,
		v.SimpleDescriptor,
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

// Arguments returns all values of RemoveNodeCacheRequest
func (v *RemoveNodeCacheRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "IeeeAddress", Argument: &v.IeeeAddress},
	}
}

// Name of the command
func (RemoveNodeCacheRequest) Name() string { return "Remove Node Cache Request" }

// ID of the command
func (RemoveNodeCacheRequest) ID() CommandID { return RemoveNodeCacheRequestCommand }

// Required
func (RemoveNodeCacheRequest) Required() bool { return false }

// Cluster ID of the command
func (RemoveNodeCacheRequest) Cluster() zcl.ClusterID { return 0x001B }

// MnfCode returns the manufacturer code (if any) of the command
func (RemoveNodeCacheRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (RemoveNodeCacheRequest) MarshalJSON() ([]byte, error) { return []byte("27"), nil }

// MarshalZcl returns the wire format representation of RemoveNodeCacheRequest
func (v RemoveNodeCacheRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

// Arguments returns all values of FindNodeCacheRequest
func (v *FindNodeCacheRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "IeeeAddress", Argument: &v.IeeeAddress},
	}
}

// Name of the command
func (FindNodeCacheRequest) Name() string { return "Find Node Cache Request" }

// ID of the command
func (FindNodeCacheRequest) ID() CommandID { return FindNodeCacheRequestCommand }

// Required
func (FindNodeCacheRequest) Required() bool { return false }

// Cluster ID of the command
func (FindNodeCacheRequest) Cluster() zcl.ClusterID { return 0x001C }

// MnfCode returns the manufacturer code (if any) of the command
func (FindNodeCacheRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (FindNodeCacheRequest) MarshalJSON() ([]byte, error) { return []byte("28"), nil }

// MarshalZcl returns the wire format representation of FindNodeCacheRequest
func (v FindNodeCacheRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

// Arguments returns all values of ExtendedSimpleDescRequest
func (v *ExtendedSimpleDescRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "Endpoint", Argument: &v.Endpoint},
		{Name: "StartIndex", Argument: &v.StartIndex},
	}
}

// Name of the command
func (ExtendedSimpleDescRequest) Name() string { return "Extended Simple Desc Request" }

// ID of the command
func (ExtendedSimpleDescRequest) ID() CommandID { return ExtendedSimpleDescRequestCommand }

// Required
func (ExtendedSimpleDescRequest) Required() bool { return false }

// Cluster ID of the command
func (ExtendedSimpleDescRequest) Cluster() zcl.ClusterID { return 0x001D }

// MnfCode returns the manufacturer code (if any) of the command
func (ExtendedSimpleDescRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ExtendedSimpleDescRequest) MarshalJSON() ([]byte, error) { return []byte("29"), nil }

// MarshalZcl returns the wire format representation of ExtendedSimpleDescRequest
func (v ExtendedSimpleDescRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

// Arguments returns all values of ExtendedActiveEndpointRequest
func (v *ExtendedActiveEndpointRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "StartIndex", Argument: &v.StartIndex},
	}
}

// Name of the command
func (ExtendedActiveEndpointRequest) Name() string { return "Extended Active Endpoint Request" }

// ID of the command
func (ExtendedActiveEndpointRequest) ID() CommandID { return ExtendedActiveEndpointRequestCommand }

// Required
func (ExtendedActiveEndpointRequest) Required() bool { return false }

// Cluster ID of the command
func (ExtendedActiveEndpointRequest) Cluster() zcl.ClusterID { return 0x001E }

// MnfCode returns the manufacturer code (if any) of the command
func (ExtendedActiveEndpointRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ExtendedActiveEndpointRequest) MarshalJSON() ([]byte, error) { return []byte("30"), nil }

// MarshalZcl returns the wire format representation of ExtendedActiveEndpointRequest
func (v ExtendedActiveEndpointRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

// Arguments returns all values of EndDeviceBindRequest
func (v *EndDeviceBindRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "BindingTarget", Argument: &v.BindingTarget},
		{Name: "SourceAddress", Argument: &v.SourceAddress},
		{Name: "SourceEndpoint", Argument: &v.SourceEndpoint},
		{Name: "ProfileId", Argument: &v.ProfileId},
		{Name: "InClusterList", Argument: &v.InClusterList},
		{Name: "OutClusterList", Argument: &v.OutClusterList},
	}
}

// Name of the command
func (EndDeviceBindRequest) Name() string { return "End Device Bind Request" }

// ID of the command
func (EndDeviceBindRequest) ID() CommandID { return EndDeviceBindRequestCommand }

// Required
func (EndDeviceBindRequest) Required() bool { return false }

// Cluster ID of the command
func (EndDeviceBindRequest) Cluster() zcl.ClusterID { return 0x0020 }

// MnfCode returns the manufacturer code (if any) of the command
func (EndDeviceBindRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (EndDeviceBindRequest) MarshalJSON() ([]byte, error) { return []byte("32"), nil }

// MarshalZcl returns the wire format representation of EndDeviceBindRequest
func (v EndDeviceBindRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
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
	tmp2 := uint32(0)
	_ = tmp2

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

type BindRequest struct {
	// SourceAddress of device generating the request
	SourceAddress SourceAddress
	// SourceEndpoint of device generating the request
	SourceEndpoint SourceEndpoint
	ClusterId      ClusterId
	AddressMode    AddressMode
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	// IeeeAddress is a 64-bit MAC address
	IeeeAddress IeeeAddress
	Endpoint    Endpoint
}

// BindRequestCommand is the Command ID of BindRequest
const BindRequestCommand CommandID = 0x0021

// Values returns all values of BindRequest
func (v *BindRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.SourceAddress,
		&v.SourceEndpoint,
		&v.ClusterId,
		&v.AddressMode,
		&v.NwkAddress,
		&v.IeeeAddress,
		&v.Endpoint,
	}
}

// Arguments returns all values of BindRequest
func (v *BindRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "SourceAddress", Argument: &v.SourceAddress},
		{Name: "SourceEndpoint", Argument: &v.SourceEndpoint},
		{Name: "ClusterId", Argument: &v.ClusterId},
		{Name: "AddressMode", Argument: &v.AddressMode},
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "IeeeAddress", Argument: &v.IeeeAddress},
		{Name: "Endpoint", Argument: &v.Endpoint},
	}
}

// Name of the command
func (BindRequest) Name() string { return "Bind Request" }

// ID of the command
func (BindRequest) ID() CommandID { return BindRequestCommand }

// Required
func (BindRequest) Required() bool { return false }

// Cluster ID of the command
func (BindRequest) Cluster() zcl.ClusterID { return 0x0021 }

// MnfCode returns the manufacturer code (if any) of the command
func (BindRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (BindRequest) MarshalJSON() ([]byte, error) { return []byte("33"), nil }

// MarshalZcl returns the wire format representation of BindRequest
func (v BindRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

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
		if tmp, err = v.ClusterId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.AddressMode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// NwkAddress is only included if short address mode
	if v.AddressMode != 0x03 {
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// IeeeAddress is only included if long address mode
	if v.AddressMode == 0x03 {
		if tmp, err = v.IeeeAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// Endpoint is only included if is not group address
	if v.AddressMode != 0x01 {
		if tmp, err = v.Endpoint.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the BindRequest struct
func (v *BindRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.SourceAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SourceEndpoint).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ClusterId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.AddressMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	// NwkAddress is only included if short address mode
	if v.AddressMode != 0x03 {
		if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// IeeeAddress is only included if long address mode
	if v.AddressMode == 0x03 {
		if b, err = (&v.IeeeAddress).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// Endpoint is only included if is not group address
	if v.AddressMode != 0x01 {
		if b, err = (&v.Endpoint).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v BindRequest) String() string {
	return zcl.Sprintf(
		"BindRequest{"+zcl.StrJoin([]string{
			"SourceAddress(%v)",
			"SourceEndpoint(%v)",
			"ClusterId(%v)",
			"AddressMode(%v)",
			"NwkAddress(%v)",
			"IeeeAddress(%v)",
			"Endpoint(%v)",
		}, " ")+"}",
		v.SourceAddress,
		v.SourceEndpoint,
		v.ClusterId,
		v.AddressMode,
		v.NwkAddress,
		v.IeeeAddress,
		v.Endpoint,
	)
}

type BindResponse struct {
	// Status Code, command is normally empty unless status is `Success`
	Status Status
}

// BindResponseCommand is the Command ID of BindResponse
const BindResponseCommand CommandID = 0x8021

// Values returns all values of BindResponse
func (v *BindResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
	}
}

// Arguments returns all values of BindResponse
func (v *BindResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
	}
}

// Name of the command
func (BindResponse) Name() string { return "Bind Response" }

// ID of the command
func (BindResponse) ID() CommandID { return BindResponseCommand }

// Required
func (BindResponse) Required() bool { return false }

// Cluster ID of the command
func (BindResponse) Cluster() zcl.ClusterID { return 0x8021 }

// MnfCode returns the manufacturer code (if any) of the command
func (BindResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (BindResponse) MarshalJSON() ([]byte, error) { return []byte("32801"), nil }

// MarshalZcl returns the wire format representation of BindResponse
func (v BindResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the BindResponse struct
func (v *BindResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v BindResponse) String() string {
	return zcl.Sprintf(
		"BindResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
		}, " ")+"}",
		v.Status,
	)
}

type UnbindRequest struct {
	// SourceAddress of device generating the request
	SourceAddress SourceAddress
	// SourceEndpoint of device generating the request
	SourceEndpoint SourceEndpoint
	ClusterId      ClusterId
	AddressMode    AddressMode
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	// IeeeAddress is a 64-bit MAC address
	IeeeAddress IeeeAddress
	Endpoint    Endpoint
}

// UnbindRequestCommand is the Command ID of UnbindRequest
const UnbindRequestCommand CommandID = 0x0022

// Values returns all values of UnbindRequest
func (v *UnbindRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.SourceAddress,
		&v.SourceEndpoint,
		&v.ClusterId,
		&v.AddressMode,
		&v.NwkAddress,
		&v.IeeeAddress,
		&v.Endpoint,
	}
}

// Arguments returns all values of UnbindRequest
func (v *UnbindRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "SourceAddress", Argument: &v.SourceAddress},
		{Name: "SourceEndpoint", Argument: &v.SourceEndpoint},
		{Name: "ClusterId", Argument: &v.ClusterId},
		{Name: "AddressMode", Argument: &v.AddressMode},
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "IeeeAddress", Argument: &v.IeeeAddress},
		{Name: "Endpoint", Argument: &v.Endpoint},
	}
}

// Name of the command
func (UnbindRequest) Name() string { return "Unbind Request" }

// ID of the command
func (UnbindRequest) ID() CommandID { return UnbindRequestCommand }

// Required
func (UnbindRequest) Required() bool { return false }

// Cluster ID of the command
func (UnbindRequest) Cluster() zcl.ClusterID { return 0x0022 }

// MnfCode returns the manufacturer code (if any) of the command
func (UnbindRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (UnbindRequest) MarshalJSON() ([]byte, error) { return []byte("34"), nil }

// MarshalZcl returns the wire format representation of UnbindRequest
func (v UnbindRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

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
		if tmp, err = v.ClusterId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.AddressMode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// NwkAddress is only included if short address mode
	if v.AddressMode != 0x03 {
		if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// IeeeAddress is only included if long address mode
	if v.AddressMode == 0x03 {
		if tmp, err = v.IeeeAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// Endpoint is only included if is not group address
	if v.AddressMode != 0x01 {
		if tmp, err = v.Endpoint.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the UnbindRequest struct
func (v *UnbindRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.SourceAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SourceEndpoint).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ClusterId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.AddressMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	// NwkAddress is only included if short address mode
	if v.AddressMode != 0x03 {
		if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// IeeeAddress is only included if long address mode
	if v.AddressMode == 0x03 {
		if b, err = (&v.IeeeAddress).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// Endpoint is only included if is not group address
	if v.AddressMode != 0x01 {
		if b, err = (&v.Endpoint).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v UnbindRequest) String() string {
	return zcl.Sprintf(
		"UnbindRequest{"+zcl.StrJoin([]string{
			"SourceAddress(%v)",
			"SourceEndpoint(%v)",
			"ClusterId(%v)",
			"AddressMode(%v)",
			"NwkAddress(%v)",
			"IeeeAddress(%v)",
			"Endpoint(%v)",
		}, " ")+"}",
		v.SourceAddress,
		v.SourceEndpoint,
		v.ClusterId,
		v.AddressMode,
		v.NwkAddress,
		v.IeeeAddress,
		v.Endpoint,
	)
}

type UnbindResponse struct {
	// Status Code, command is normally empty unless status is `Success`
	Status Status
}

// UnbindResponseCommand is the Command ID of UnbindResponse
const UnbindResponseCommand CommandID = 0x8022

// Values returns all values of UnbindResponse
func (v *UnbindResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
	}
}

// Arguments returns all values of UnbindResponse
func (v *UnbindResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
	}
}

// Name of the command
func (UnbindResponse) Name() string { return "Unbind Response" }

// ID of the command
func (UnbindResponse) ID() CommandID { return UnbindResponseCommand }

// Required
func (UnbindResponse) Required() bool { return false }

// Cluster ID of the command
func (UnbindResponse) Cluster() zcl.ClusterID { return 0x8022 }

// MnfCode returns the manufacturer code (if any) of the command
func (UnbindResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (UnbindResponse) MarshalJSON() ([]byte, error) { return []byte("32802"), nil }

// MarshalZcl returns the wire format representation of UnbindResponse
func (v UnbindResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the UnbindResponse struct
func (v *UnbindResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v UnbindResponse) String() string {
	return zcl.Sprintf(
		"UnbindResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
		}, " ")+"}",
		v.Status,
	)
}

type MgmtLqiRequest struct {
	// StartIndex provides the starting index for the requested elements of the associated list.
	StartIndex StartIndex
}

// MgmtLqiRequestCommand is the Command ID of MgmtLqiRequest
const MgmtLqiRequestCommand CommandID = 0x0031

// Values returns all values of MgmtLqiRequest
func (v *MgmtLqiRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.StartIndex,
	}
}

// Arguments returns all values of MgmtLqiRequest
func (v *MgmtLqiRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "StartIndex", Argument: &v.StartIndex},
	}
}

// Name of the command
func (MgmtLqiRequest) Name() string { return "Mgmt Lqi Request" }

// ID of the command
func (MgmtLqiRequest) ID() CommandID { return MgmtLqiRequestCommand }

// Required
func (MgmtLqiRequest) Required() bool { return false }

// Cluster ID of the command
func (MgmtLqiRequest) Cluster() zcl.ClusterID { return 0x0031 }

// MnfCode returns the manufacturer code (if any) of the command
func (MgmtLqiRequest) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MgmtLqiRequest) MarshalJSON() ([]byte, error) { return []byte("49"), nil }

// MarshalZcl returns the wire format representation of MgmtLqiRequest
func (v MgmtLqiRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.StartIndex.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the MgmtLqiRequest struct
func (v *MgmtLqiRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.StartIndex).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MgmtLqiRequest) String() string {
	return zcl.Sprintf(
		"MgmtLqiRequest{"+zcl.StrJoin([]string{
			"StartIndex(%v)",
		}, " ")+"}",
		v.StartIndex,
	)
}

type MgmtLqiResponse struct {
	// Status Code, command is normally empty unless status is `Success`
	Status Status
	// TotalEntries is the total number of entries that can be queried for
	TotalEntries TotalEntries
	// StartIndex provides the starting index for the requested elements of the associated list.
	StartIndex    StartIndex
	NeighborTable NeighborTable
}

// MgmtLqiResponseCommand is the Command ID of MgmtLqiResponse
const MgmtLqiResponseCommand CommandID = 0x8031

// Values returns all values of MgmtLqiResponse
func (v *MgmtLqiResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.TotalEntries,
		&v.StartIndex,
		&v.NeighborTable,
	}
}

// Arguments returns all values of MgmtLqiResponse
func (v *MgmtLqiResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "TotalEntries", Argument: &v.TotalEntries},
		{Name: "StartIndex", Argument: &v.StartIndex},
		{Name: "NeighborTable", Argument: &v.NeighborTable},
	}
}

// Name of the command
func (MgmtLqiResponse) Name() string { return "Mgmt Lqi Response" }

// ID of the command
func (MgmtLqiResponse) ID() CommandID { return MgmtLqiResponseCommand }

// Required
func (MgmtLqiResponse) Required() bool { return false }

// Cluster ID of the command
func (MgmtLqiResponse) Cluster() zcl.ClusterID { return 0x8031 }

// MnfCode returns the manufacturer code (if any) of the command
func (MgmtLqiResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MgmtLqiResponse) MarshalJSON() ([]byte, error) { return []byte("32817"), nil }

// MarshalZcl returns the wire format representation of MgmtLqiResponse
func (v MgmtLqiResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.TotalEntries.MarshalZcl(); err != nil {
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
	{
		if tmp, err = v.NeighborTable.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the MgmtLqiResponse struct
func (v *MgmtLqiResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TotalEntries).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StartIndex).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NeighborTable).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MgmtLqiResponse) String() string {
	return zcl.Sprintf(
		"MgmtLqiResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"TotalEntries(%v)",
			"StartIndex(%v)",
			"NeighborTable(%v)",
		}, " ")+"}",
		v.Status,
		v.TotalEntries,
		v.StartIndex,
		v.NeighborTable,
	)
}
