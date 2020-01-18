package zdo

import "hemtjan.st/zcl"

type CommandID = zcl.ZdoCmdID
type Frame = zcl.ReceivedZdpFrame

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
	UserDescResponseCommand:              func() zcl.ZdoCommand { return new(UserDescResponse) },
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
	MgmtBindRequestCommand:               func() zcl.ZdoCommand { return new(MgmtBindRequest) },
	MgmtBindResponseCommand:              func() zcl.ZdoCommand { return new(MgmtBindResponse) },
	UnbindRequestCommand:                 func() zcl.ZdoCommand { return new(UnbindRequest) },
	UnbindResponseCommand:                func() zcl.ZdoCommand { return new(UnbindResponse) },
	MgmtLqiRequestCommand:                func() zcl.ZdoCommand { return new(MgmtLqiRequest) },
	MgmtLqiResponseCommand:               func() zcl.ZdoCommand { return new(MgmtLqiResponse) },
	MgmtNwkUpdateCommand:                 func() zcl.ZdoCommand { return new(MgmtNwkUpdate) },
}

func (ApsFlags) Name() string        { return `APS Flags` }
func (ApsFlags) Description() string { return `` }

type ApsFlags zcl.Zu8

func (v *ApsFlags) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *ApsFlags) Value() zcl.Val     { return v }

func (v ApsFlags) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *ApsFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = ApsFlags(*nt)
	return br, err
}

func (v ApsFlags) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *ApsFlags) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ApsFlags(*a)
	return nil
}

func (v *ApsFlags) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = ApsFlags(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ApsFlags) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (ActiveEndpointList) Name() string        { return `Active Endpoint List` }
func (ActiveEndpointList) Description() string { return `List of active endpoints` }

// ActiveEndpointList List of active endpoints
type ActiveEndpointList []*zcl.Zu8

func (v *ActiveEndpointList) TypeID() zcl.TypeID { return new(zcl.Zset).TypeID() }
func (v *ActiveEndpointList) Value() zcl.Val     { return v }

func (ActiveEndpointList) ArrayTypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }

func (v *ActiveEndpointList) ArrayValues() (o []uint8) {
	for _, a := range *v {
		o = append(o, uint8(*a))
	}
	return o
}

func (v *ActiveEndpointList) SetValues(val []uint8) error {
	*v = []*zcl.Zu8{}
	return v.AddValues(val...)
}

func (v *ActiveEndpointList) AddValues(val ...uint8) error {
	for _, a := range val {
		nv := zcl.Zu8(a)
		*v = append(*v, &nv)
	}
	return nil
}

func (v ActiveEndpointList) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, a := range v {
		vars = append(vars, a)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (v *ActiveEndpointList) UnmarshalZcl(b []byte) ([]byte, error) {
	*v = []*zcl.Zu8{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(zcl.Zu8)
		*v = append(*v, nv)
		return nv
	})
}

func (v *ActiveEndpointList) SetValue(a zcl.Val) error {
	if nv, ok := a.(*ActiveEndpointList); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ActiveEndpointList) String() string {
	var s []string
	for _, a := range v {
		s = append(s, zcl.Sprintf("%v", a))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

func (ActiveEndpointSize) Name() string        { return `Active Endpoint Size` }
func (ActiveEndpointSize) Description() string { return `Size in bytes of the Active Endpoints List` }

// ActiveEndpointSize Size in bytes of the Active Endpoints List
type ActiveEndpointSize zcl.Zu8

func (v *ActiveEndpointSize) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *ActiveEndpointSize) Value() zcl.Val     { return v }

func (v ActiveEndpointSize) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *ActiveEndpointSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = ActiveEndpointSize(*nt)
	return br, err
}

func (v ActiveEndpointSize) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *ActiveEndpointSize) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ActiveEndpointSize(*a)
	return nil
}

func (v *ActiveEndpointSize) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = ActiveEndpointSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ActiveEndpointSize) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (AddressMode) Name() string        { return `Address Mode` }
func (AddressMode) Description() string { return `` }

type AddressMode zcl.Zenum8

func (v *AddressMode) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *AddressMode) Value() zcl.Val     { return v }

func (v AddressMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *AddressMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = AddressMode(*nt)
	return br, err
}

func (v AddressMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *AddressMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AddressMode(*a)
	return nil
}

func (v *AddressMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = AddressMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AddressMode) String() string {
	switch v {
	case 0x01:
		return "Group"
	case 0x02:
		return "NWK"
	case 0x03:
		return "IEEE"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v AddressMode) IsGroup() bool { return v == 0x01 }
func (v AddressMode) IsNwk() bool   { return v == 0x02 }
func (v AddressMode) IsIeee() bool  { return v == 0x03 }
func (v *AddressMode) SetGroup()    { *v = 0x01 }
func (v *AddressMode) SetNwk()      { *v = 0x02 }
func (v *AddressMode) SetIeee()     { *v = 0x03 }

func (AddressMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x01, Name: "Group"},
		{Value: 0x02, Name: "NWK"},
		{Value: 0x03, Name: "IEEE"},
	}
}

func (AssociatedDevices) Name() string        { return `Associated Devices` }
func (AssociatedDevices) Description() string { return `` }

type AssociatedDevices []*zcl.Zu16

func (v *AssociatedDevices) TypeID() zcl.TypeID { return new(zcl.Zlist).TypeID() }
func (v *AssociatedDevices) Value() zcl.Val     { return v }

func (AssociatedDevices) ArrayTypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }

func (v *AssociatedDevices) ArrayValues() (o []uint16) {
	for _, a := range *v {
		o = append(o, uint16(*a))
	}
	return o
}

func (v *AssociatedDevices) SetValues(val []uint16) error {
	*v = []*zcl.Zu16{}
	return v.AddValues(val...)
}

func (v *AssociatedDevices) AddValues(val ...uint16) error {
	for _, a := range val {
		nv := zcl.Zu16(a)
		*v = append(*v, &nv)
	}
	return nil
}

func (v AssociatedDevices) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, a := range v {
		vars = append(vars, a)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (v *AssociatedDevices) UnmarshalZcl(b []byte) ([]byte, error) {
	*v = []*zcl.Zu16{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(zcl.Zu16)
		*v = append(*v, nv)
		return nv
	})
}

func (v *AssociatedDevices) SetValue(a zcl.Val) error {
	if nv, ok := a.(*AssociatedDevices); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AssociatedDevices) String() string {
	var s []string
	for _, a := range v {
		s = append(s, zcl.Sprintf("%v", a))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

func (BindingTable) Name() string        { return `Binding Table` }
func (BindingTable) Description() string { return `` }

type BindingTable []*BindingTableEntry

func (v *BindingTable) TypeID() zcl.TypeID { return new(zcl.Zarray).TypeID() }
func (v *BindingTable) Value() zcl.Val     { return v }

func (BindingTable) ArrayTypeID() zcl.TypeID { return new(BindingTableEntry).TypeID() }

func (v *BindingTable) ArrayValues() (o []BindingTableEntry) {
	for _, a := range *v {
		o = append(o, *a)
	}
	return o
}

func (v *BindingTable) SetValues(val []BindingTableEntry) error {
	*v = []*BindingTableEntry{}
	return v.AddValues(val...)
}

func (v *BindingTable) AddValues(val ...BindingTableEntry) error {
	for _, a := range val {
		nv := a
		*v = append(*v, &nv)
	}
	return nil
}

func (v BindingTable) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, a := range v {
		vars = append(vars, a)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (v *BindingTable) UnmarshalZcl(b []byte) ([]byte, error) {
	*v = []*BindingTableEntry{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(BindingTableEntry)
		*v = append(*v, nv)
		return nv
	})
}

func (v *BindingTable) SetValue(a zcl.Val) error {
	if nv, ok := a.(*BindingTable); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BindingTable) String() string {
	var s []string
	for _, a := range v {
		s = append(s, zcl.Sprintf("%v", a))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

func (BindingTableEntry) Name() string        { return `Binding Table Entry` }
func (BindingTableEntry) Description() string { return `` }

type BindingTableEntry struct {
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

func (v *BindingTableEntry) TypeID() zcl.TypeID { return zcl.TypeID(76) } // struct
func (v *BindingTableEntry) Value() zcl.Val     { return v }
func (v BindingTableEntry) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	var err error
	_ = tmp2
	_ = err

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

func (v *BindingTableEntry) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2
	_ = err

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

func (v *BindingTableEntry) SetValue(a zcl.Val) error {
	if nv, ok := a.(*BindingTableEntry); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrNotImpl
}

func (v *BindingTableEntry) String() string {
	return zcl.Sprintf(
		"BindingTableEntry{"+zcl.StrJoin([]string{
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

func (BindingTarget) Name() string        { return `Binding Target` }
func (BindingTarget) Description() string { return `NWK Address` }

// BindingTarget NWK Address
type BindingTarget zcl.Zu16

func (v *BindingTarget) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *BindingTarget) Value() zcl.Val     { return v }

func (v BindingTarget) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *BindingTarget) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = BindingTarget(*nt)
	return br, err
}

func (v BindingTarget) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *BindingTarget) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BindingTarget(*a)
	return nil
}

func (v *BindingTarget) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = BindingTarget(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BindingTarget) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (Capability) Name() string        { return `Capability` }
func (Capability) Description() string { return `specifies the device:s capabilities` }

// Capability specifies the device:s capabilities
type Capability zcl.Zbmp8

func (v *Capability) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *Capability) Value() zcl.Val     { return v }

func (v Capability) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *Capability) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = Capability(*nt)
	return br, err
}

func (v Capability) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *Capability) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Capability(*a)
	return nil
}

func (v *Capability) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = Capability(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Capability) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v Capability) IsPanCoordinator() bool     { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v Capability) IsFullFunction() bool       { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v Capability) IsMainsPower() bool         { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v Capability) IsRxOnIdle() bool           { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v Capability) IsSecurity() bool           { return zcl.BitmapTest([]byte(v[:]), 6) }
func (v Capability) IsAllocateAddress() bool    { return zcl.BitmapTest([]byte(v[:]), 7) }
func (v *Capability) SetPanCoordinator(b bool)  { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b)) }
func (v *Capability) SetFullFunction(b bool)    { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b)) }
func (v *Capability) SetMainsPower(b bool)      { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b)) }
func (v *Capability) SetRxOnIdle(b bool)        { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b)) }
func (v *Capability) SetSecurity(b bool)        { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 6, b)) }
func (v *Capability) SetAllocateAddress(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 7, b)) }

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

func (ClusterId) Name() string        { return `Cluster ID` }
func (ClusterId) Description() string { return `` }

type ClusterId zcl.Zu16

func (v *ClusterId) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ClusterId) Value() zcl.Val     { return v }

func (v ClusterId) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ClusterId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ClusterId(*nt)
	return br, err
}

func (v ClusterId) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ClusterId) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ClusterId(*a)
	return nil
}

func (v *ClusterId) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ClusterId(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ClusterId) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (ComplexDescriptor) Name() string        { return `Complex Descriptor` }
func (ComplexDescriptor) Description() string { return `` }

type ComplexDescriptor zcl.Zostring

func (v *ComplexDescriptor) TypeID() zcl.TypeID { return new(zcl.Zostring).TypeID() }
func (v *ComplexDescriptor) Value() zcl.Val     { return v }

func (v ComplexDescriptor) MarshalZcl() ([]byte, error) { return zcl.Zostring(v).MarshalZcl() }

func (v *ComplexDescriptor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*v = ComplexDescriptor(*nt)
	return br, err
}

func (v ComplexDescriptor) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zostring(v))
}

func (v *ComplexDescriptor) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zostring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ComplexDescriptor(*a)
	return nil
}

func (v *ComplexDescriptor) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zostring); ok {
		*v = ComplexDescriptor(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ComplexDescriptor) String() string {
	return zcl.Sprintf("%v", zcl.Zostring(v))
}

func (ComplexDescriptorAvailable) Name() string        { return `Complex Descriptor Available` }
func (ComplexDescriptorAvailable) Description() string { return `` }

type ComplexDescriptorAvailable zcl.Zbool

func (v *ComplexDescriptorAvailable) TypeID() zcl.TypeID { return new(zcl.Zbool).TypeID() }
func (v *ComplexDescriptorAvailable) Value() zcl.Val     { return v }

func (v ComplexDescriptorAvailable) MarshalZcl() ([]byte, error) { return zcl.Zbool(v).MarshalZcl() }

func (v *ComplexDescriptorAvailable) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*v = ComplexDescriptorAvailable(*nt)
	return br, err
}

func (v ComplexDescriptorAvailable) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbool(v))
}

func (v *ComplexDescriptorAvailable) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbool)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ComplexDescriptorAvailable(*a)
	return nil
}

func (v *ComplexDescriptorAvailable) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbool); ok {
		*v = ComplexDescriptorAvailable(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ComplexDescriptorAvailable) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(v))
}

func (Depth) Name() string { return `Depth` }
func (Depth) Description() string {
	return `of the neighbor device. A value of 0 indicates that the device is the coordinator for the network`
}

// Depth of the neighbor device. A value of 0 indicates that the device is the coordinator for the network
type Depth zcl.Zu8

func (v *Depth) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *Depth) Value() zcl.Val     { return v }

func (v Depth) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *Depth) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = Depth(*nt)
	return br, err
}

func (v Depth) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *Depth) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Depth(*a)
	return nil
}

func (v *Depth) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = Depth(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Depth) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (DescriptorCapability) Name() string        { return `Descriptor Capability` }
func (DescriptorCapability) Description() string { return `` }

type DescriptorCapability zcl.Zbmp8

func (v *DescriptorCapability) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *DescriptorCapability) Value() zcl.Val     { return v }

func (v DescriptorCapability) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *DescriptorCapability) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = DescriptorCapability(*nt)
	return br, err
}

func (v DescriptorCapability) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *DescriptorCapability) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DescriptorCapability(*a)
	return nil
}

func (v *DescriptorCapability) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = DescriptorCapability(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DescriptorCapability) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v DescriptorCapability) IsExtendedActiveEndpointListAvailable() bool {
	return zcl.BitmapTest([]byte(v[:]), 0)
}
func (v DescriptorCapability) IsExtendedSimpleDescriptorListAvailable() bool {
	return zcl.BitmapTest([]byte(v[:]), 1)
}
func (v *DescriptorCapability) SetExtendedActiveEndpointListAvailable(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *DescriptorCapability) SetExtendedSimpleDescriptorListAvailable(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}

func (DescriptorCapability) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Extended Active Endpoint List Available"},
		{Value: 1, Name: "Extended Simple Descriptor List Available"},
	}
}

func (DeviceType) Name() string        { return `Device Type` }
func (DeviceType) Description() string { return `` }

type DeviceType zcl.Zenum16

func (v *DeviceType) TypeID() zcl.TypeID { return new(zcl.Zenum16).TypeID() }
func (v *DeviceType) Value() zcl.Val     { return v }

func (v DeviceType) MarshalZcl() ([]byte, error) { return zcl.Zenum16(v).MarshalZcl() }

func (v *DeviceType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum16)
	br, err := nt.UnmarshalZcl(b)
	*v = DeviceType(*nt)
	return br, err
}

func (v DeviceType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum16(v))
}

func (v *DeviceType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DeviceType(*a)
	return nil
}

func (v *DeviceType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum16); ok {
		*v = DeviceType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DeviceType) String() string {
	switch v {
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
	return zcl.Sprintf("%v", zcl.Zenum16(v))
}

func (v DeviceType) IsOnOffSwitch() bool                            { return v == 0x0000 }
func (v DeviceType) IsLevelControlSwitch() bool                     { return v == 0x0001 }
func (v DeviceType) IsOnOffOutput() bool                            { return v == 0x0002 }
func (v DeviceType) IsLevelControllableOutput() bool                { return v == 0x0003 }
func (v DeviceType) IsSceneSelector() bool                          { return v == 0x0004 }
func (v DeviceType) IsConfigurationTool() bool                      { return v == 0x0005 }
func (v DeviceType) IsRemoteControl() bool                          { return v == 0x0006 }
func (v DeviceType) IsCombinedInterface() bool                      { return v == 0x0007 }
func (v DeviceType) IsRangeExtender() bool                          { return v == 0x0008 }
func (v DeviceType) IsMainsPowerOutlet() bool                       { return v == 0x0009 }
func (v DeviceType) IsDoorLock() bool                               { return v == 0x000a }
func (v DeviceType) IsSimpleSensor() bool                           { return v == 0x000c }
func (v DeviceType) IsOnOffPlugInUnit() bool                        { return v == 0x0010 }
func (v DeviceType) IsSmartPlug() bool                              { return v == 0x0051 }
func (v DeviceType) IsGpProxy() bool                                { return v == 0x0060 }
func (v DeviceType) IsGpProxyBasic() bool                           { return v == 0x0061 }
func (v DeviceType) IsGpTargetPlus() bool                           { return v == 0x0062 }
func (v DeviceType) IsGpTarget() bool                               { return v == 0x0063 }
func (v DeviceType) IsGpCommissioningTool() bool                    { return v == 0x0064 }
func (v DeviceType) IsGpCombo() bool                                { return v == 0x0065 }
func (v DeviceType) IsGpComboBasic() bool                           { return v == 0x0066 }
func (v DeviceType) IsOnOffLight() bool                             { return v == 0x0100 }
func (v DeviceType) IsDimmableLight() bool                          { return v == 0x0101 }
func (v DeviceType) IsColorDimmableLight() bool                     { return v == 0x0102 }
func (v DeviceType) IsOnOffLightSwitch() bool                       { return v == 0x0103 }
func (v DeviceType) IsDimmerSwitch() bool                           { return v == 0x0104 }
func (v DeviceType) IsColorDimmerSwitch() bool                      { return v == 0x0105 }
func (v DeviceType) IsLightSensor() bool                            { return v == 0x0106 }
func (v DeviceType) IsOccupancySensor() bool                        { return v == 0x0107 }
func (v DeviceType) IsOnOffBallast() bool                           { return v == 0x0108 }
func (v DeviceType) IsDimmableBallast() bool                        { return v == 0x0109 }
func (v DeviceType) IsOnOffPluginUnit() bool                        { return v == 0x010a }
func (v DeviceType) IsDimmablePluginUnit() bool                     { return v == 0x010b }
func (v DeviceType) IsColorTemperatureLight() bool                  { return v == 0x010c }
func (v DeviceType) IsExtendedColorLight() bool                     { return v == 0x010d }
func (v DeviceType) IsLightLevelSensor() bool                       { return v == 0x010e }
func (v DeviceType) IsDimmablePlugInUnit() bool                     { return v == 0x0110 }
func (v DeviceType) IsShade() bool                                  { return v == 0x0200 }
func (v DeviceType) IsShadeController() bool                        { return v == 0x0201 }
func (v DeviceType) IsWindowCoveringDevice() bool                   { return v == 0x0202 }
func (v DeviceType) IsExtendedColorLight2() bool                    { return v == 0x0210 }
func (v DeviceType) IsColorTemperatureLight2() bool                 { return v == 0x0220 }
func (v DeviceType) IsHeatingCoolingUnit() bool                     { return v == 0x0300 }
func (v DeviceType) IsThermostat() bool                             { return v == 0x0301 }
func (v DeviceType) IsTemperatureSensor() bool                      { return v == 0x0302 }
func (v DeviceType) IsPump() bool                                   { return v == 0x0303 }
func (v DeviceType) IsPumpController() bool                         { return v == 0x0304 }
func (v DeviceType) IsPressureSensor() bool                         { return v == 0x0305 }
func (v DeviceType) IsFlowSensor() bool                             { return v == 0x0306 }
func (v DeviceType) IsIasControlAndIndicatingEquipment() bool       { return v == 0x0400 }
func (v DeviceType) IsIasAncillaryControlEquipment() bool           { return v == 0x0401 }
func (v DeviceType) IsIasZone() bool                                { return v == 0x0402 }
func (v DeviceType) IsIasWarningDevice() bool                       { return v == 0x0403 }
func (v DeviceType) IsEnergyServicePortal() bool                    { return v == 0x0500 }
func (v DeviceType) IsMeteringDevice() bool                         { return v == 0x0501 }
func (v DeviceType) IsInPremiseDisplay() bool                       { return v == 0x0502 }
func (v DeviceType) IsProgrammableCommunicatingThermostatPct() bool { return v == 0x0503 }
func (v DeviceType) IsLoadControlDevice() bool                      { return v == 0x0504 }
func (v DeviceType) IsSmartAppliance() bool                         { return v == 0x0505 }
func (v DeviceType) IsPrepaymentTerminal() bool                     { return v == 0x0506 }
func (v DeviceType) IsDeviceManagement() bool                       { return v == 0x0507 }
func (v DeviceType) IsColorController() bool                        { return v == 0x0800 }
func (v DeviceType) IsColorSceneController() bool                   { return v == 0x0810 }
func (v DeviceType) IsNonColorController() bool                     { return v == 0x0820 }
func (v DeviceType) IsNonColorSceneController() bool                { return v == 0x0830 }
func (v DeviceType) IsControlBridge() bool                          { return v == 0x0840 }
func (v DeviceType) IsOnOffSensor() bool                            { return v == 0x0850 }
func (v *DeviceType) SetOnOffSwitch()                               { *v = 0x0000 }
func (v *DeviceType) SetLevelControlSwitch()                        { *v = 0x0001 }
func (v *DeviceType) SetOnOffOutput()                               { *v = 0x0002 }
func (v *DeviceType) SetLevelControllableOutput()                   { *v = 0x0003 }
func (v *DeviceType) SetSceneSelector()                             { *v = 0x0004 }
func (v *DeviceType) SetConfigurationTool()                         { *v = 0x0005 }
func (v *DeviceType) SetRemoteControl()                             { *v = 0x0006 }
func (v *DeviceType) SetCombinedInterface()                         { *v = 0x0007 }
func (v *DeviceType) SetRangeExtender()                             { *v = 0x0008 }
func (v *DeviceType) SetMainsPowerOutlet()                          { *v = 0x0009 }
func (v *DeviceType) SetDoorLock()                                  { *v = 0x000a }
func (v *DeviceType) SetSimpleSensor()                              { *v = 0x000c }
func (v *DeviceType) SetOnOffPlugInUnit()                           { *v = 0x0010 }
func (v *DeviceType) SetSmartPlug()                                 { *v = 0x0051 }
func (v *DeviceType) SetGpProxy()                                   { *v = 0x0060 }
func (v *DeviceType) SetGpProxyBasic()                              { *v = 0x0061 }
func (v *DeviceType) SetGpTargetPlus()                              { *v = 0x0062 }
func (v *DeviceType) SetGpTarget()                                  { *v = 0x0063 }
func (v *DeviceType) SetGpCommissioningTool()                       { *v = 0x0064 }
func (v *DeviceType) SetGpCombo()                                   { *v = 0x0065 }
func (v *DeviceType) SetGpComboBasic()                              { *v = 0x0066 }
func (v *DeviceType) SetOnOffLight()                                { *v = 0x0100 }
func (v *DeviceType) SetDimmableLight()                             { *v = 0x0101 }
func (v *DeviceType) SetColorDimmableLight()                        { *v = 0x0102 }
func (v *DeviceType) SetOnOffLightSwitch()                          { *v = 0x0103 }
func (v *DeviceType) SetDimmerSwitch()                              { *v = 0x0104 }
func (v *DeviceType) SetColorDimmerSwitch()                         { *v = 0x0105 }
func (v *DeviceType) SetLightSensor()                               { *v = 0x0106 }
func (v *DeviceType) SetOccupancySensor()                           { *v = 0x0107 }
func (v *DeviceType) SetOnOffBallast()                              { *v = 0x0108 }
func (v *DeviceType) SetDimmableBallast()                           { *v = 0x0109 }
func (v *DeviceType) SetOnOffPluginUnit()                           { *v = 0x010a }
func (v *DeviceType) SetDimmablePluginUnit()                        { *v = 0x010b }
func (v *DeviceType) SetColorTemperatureLight()                     { *v = 0x010c }
func (v *DeviceType) SetExtendedColorLight()                        { *v = 0x010d }
func (v *DeviceType) SetLightLevelSensor()                          { *v = 0x010e }
func (v *DeviceType) SetDimmablePlugInUnit()                        { *v = 0x0110 }
func (v *DeviceType) SetShade()                                     { *v = 0x0200 }
func (v *DeviceType) SetShadeController()                           { *v = 0x0201 }
func (v *DeviceType) SetWindowCoveringDevice()                      { *v = 0x0202 }
func (v *DeviceType) SetExtendedColorLight2()                       { *v = 0x0210 }
func (v *DeviceType) SetColorTemperatureLight2()                    { *v = 0x0220 }
func (v *DeviceType) SetHeatingCoolingUnit()                        { *v = 0x0300 }
func (v *DeviceType) SetThermostat()                                { *v = 0x0301 }
func (v *DeviceType) SetTemperatureSensor()                         { *v = 0x0302 }
func (v *DeviceType) SetPump()                                      { *v = 0x0303 }
func (v *DeviceType) SetPumpController()                            { *v = 0x0304 }
func (v *DeviceType) SetPressureSensor()                            { *v = 0x0305 }
func (v *DeviceType) SetFlowSensor()                                { *v = 0x0306 }
func (v *DeviceType) SetIasControlAndIndicatingEquipment()          { *v = 0x0400 }
func (v *DeviceType) SetIasAncillaryControlEquipment()              { *v = 0x0401 }
func (v *DeviceType) SetIasZone()                                   { *v = 0x0402 }
func (v *DeviceType) SetIasWarningDevice()                          { *v = 0x0403 }
func (v *DeviceType) SetEnergyServicePortal()                       { *v = 0x0500 }
func (v *DeviceType) SetMeteringDevice()                            { *v = 0x0501 }
func (v *DeviceType) SetInPremiseDisplay()                          { *v = 0x0502 }
func (v *DeviceType) SetProgrammableCommunicatingThermostatPct()    { *v = 0x0503 }
func (v *DeviceType) SetLoadControlDevice()                         { *v = 0x0504 }
func (v *DeviceType) SetSmartAppliance()                            { *v = 0x0505 }
func (v *DeviceType) SetPrepaymentTerminal()                        { *v = 0x0506 }
func (v *DeviceType) SetDeviceManagement()                          { *v = 0x0507 }
func (v *DeviceType) SetColorController()                           { *v = 0x0800 }
func (v *DeviceType) SetColorSceneController()                      { *v = 0x0810 }
func (v *DeviceType) SetNonColorController()                        { *v = 0x0820 }
func (v *DeviceType) SetNonColorSceneController()                   { *v = 0x0830 }
func (v *DeviceType) SetControlBridge()                             { *v = 0x0840 }
func (v *DeviceType) SetOnOffSensor()                               { *v = 0x0850 }

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

func (DeviceVersion) Name() string        { return `Device Version` }
func (DeviceVersion) Description() string { return `is dependant on profile` }

// DeviceVersion is dependant on profile
type DeviceVersion zcl.Zu8

func (v *DeviceVersion) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *DeviceVersion) Value() zcl.Val     { return v }

func (v DeviceVersion) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *DeviceVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = DeviceVersion(*nt)
	return br, err
}

func (v DeviceVersion) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *DeviceVersion) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DeviceVersion(*a)
	return nil
}

func (v *DeviceVersion) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = DeviceVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DeviceVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (Endpoint) Name() string        { return `Endpoint` }
func (Endpoint) Description() string { return `` }

type Endpoint zcl.Zu8

func (v *Endpoint) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *Endpoint) Value() zcl.Val     { return v }

func (v Endpoint) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *Endpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = Endpoint(*nt)
	return br, err
}

func (v Endpoint) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *Endpoint) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Endpoint(*a)
	return nil
}

func (v *Endpoint) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = Endpoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Endpoint) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (EndpointList) Name() string        { return `Endpoint List` }
func (EndpointList) Description() string { return `` }

type EndpointList []*zcl.Zu8

func (v *EndpointList) TypeID() zcl.TypeID { return new(zcl.Zset).TypeID() }
func (v *EndpointList) Value() zcl.Val     { return v }

func (EndpointList) ArrayTypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }

func (v *EndpointList) ArrayValues() (o []uint8) {
	for _, a := range *v {
		o = append(o, uint8(*a))
	}
	return o
}

func (v *EndpointList) SetValues(val []uint8) error {
	*v = []*zcl.Zu8{}
	return v.AddValues(val...)
}

func (v *EndpointList) AddValues(val ...uint8) error {
	for _, a := range val {
		nv := zcl.Zu8(a)
		*v = append(*v, &nv)
	}
	return nil
}

func (v EndpointList) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, a := range v {
		vars = append(vars, a)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (v *EndpointList) UnmarshalZcl(b []byte) ([]byte, error) {
	*v = []*zcl.Zu8{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(zcl.Zu8)
		*v = append(*v, nv)
		return nv
	})
}

func (v *EndpointList) SetValue(a zcl.Val) error {
	if nv, ok := a.(*EndpointList); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (v EndpointList) String() string {
	var s []string
	for _, a := range v {
		s = append(s, zcl.Sprintf("%v", a))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

func (EnergyValues) Name() string        { return `Energy Values` }
func (EnergyValues) Description() string { return `` }

type EnergyValues []*zcl.Zu8

func (v *EnergyValues) TypeID() zcl.TypeID { return new(zcl.Zarray).TypeID() }
func (v *EnergyValues) Value() zcl.Val     { return v }

func (EnergyValues) ArrayTypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }

func (v *EnergyValues) ArrayValues() (o []uint8) {
	for _, a := range *v {
		o = append(o, uint8(*a))
	}
	return o
}

func (v *EnergyValues) SetValues(val []uint8) error {
	*v = []*zcl.Zu8{}
	return v.AddValues(val...)
}

func (v *EnergyValues) AddValues(val ...uint8) error {
	for _, a := range val {
		nv := zcl.Zu8(a)
		*v = append(*v, &nv)
	}
	return nil
}

func (v EnergyValues) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, a := range v {
		vars = append(vars, a)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (v *EnergyValues) UnmarshalZcl(b []byte) ([]byte, error) {
	*v = []*zcl.Zu8{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(zcl.Zu8)
		*v = append(*v, nv)
		return nv
	})
}

func (v *EnergyValues) SetValue(a zcl.Val) error {
	if nv, ok := a.(*EnergyValues); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (v EnergyValues) String() string {
	var s []string
	for _, a := range v {
		s = append(s, zcl.Sprintf("%v", a))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

func (FrequencyBands) Name() string        { return `Frequency Bands` }
func (FrequencyBands) Description() string { return `` }

type FrequencyBands zcl.Zbmp8

func (v *FrequencyBands) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *FrequencyBands) Value() zcl.Val     { return v }

func (v FrequencyBands) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *FrequencyBands) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = FrequencyBands(*nt)
	return br, err
}

func (v FrequencyBands) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *FrequencyBands) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = FrequencyBands(*a)
	return nil
}

func (v *FrequencyBands) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = FrequencyBands(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v FrequencyBands) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "868 MHz")
		case 2:
			bstr = append(bstr, "902-928 MHz")
		case 3:
			bstr = append(bstr, "2.4 GHz")
		case 4:
			bstr = append(bstr, "EU Sub-GHz")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v FrequencyBands) Is868Mhz() bool       { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v FrequencyBands) Is902928Mhz() bool    { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v FrequencyBands) Is24Ghz() bool        { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v FrequencyBands) IsEuSubGhz() bool     { return zcl.BitmapTest([]byte(v[:]), 4) }
func (v *FrequencyBands) Set868Mhz(b bool)    { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b)) }
func (v *FrequencyBands) Set902928Mhz(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b)) }
func (v *FrequencyBands) Set24Ghz(b bool)     { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b)) }
func (v *FrequencyBands) SetEuSubGhz(b bool)  { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 4, b)) }

func (FrequencyBands) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "868 MHz"},
		{Value: 2, Name: "902-928 MHz"},
		{Value: 3, Name: "2.4 GHz"},
		{Value: 4, Name: "EU Sub-GHz"},
	}
}

func (IeeeAddress) Name() string        { return `IEEE Address` }
func (IeeeAddress) Description() string { return `is a 64-bit MAC address` }

// IeeeAddress is a 64-bit MAC address
type IeeeAddress zcl.Zuid

func (v *IeeeAddress) TypeID() zcl.TypeID { return new(zcl.Zuid).TypeID() }
func (v *IeeeAddress) Value() zcl.Val     { return v }

func (v IeeeAddress) MarshalZcl() ([]byte, error) { return zcl.Zuid(v).MarshalZcl() }

func (v *IeeeAddress) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*v = IeeeAddress(*nt)
	return br, err
}

func (v IeeeAddress) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zuid(v))
}

func (v *IeeeAddress) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zuid)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IeeeAddress(*a)
	return nil
}

func (v *IeeeAddress) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zuid); ok {
		*v = IeeeAddress(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IeeeAddress) String() string {
	return zcl.Sprintf("%v", zcl.Zuid(v))
}

func (InClusterList) Name() string        { return `In Cluster List` }
func (InClusterList) Description() string { return `is a list of input clusters` }

// InClusterList is a list of input clusters
type InClusterList []*zcl.Zu16

func (v *InClusterList) TypeID() zcl.TypeID { return new(zcl.Zset).TypeID() }
func (v *InClusterList) Value() zcl.Val     { return v }

func (InClusterList) ArrayTypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }

func (v *InClusterList) ArrayValues() (o []uint16) {
	for _, a := range *v {
		o = append(o, uint16(*a))
	}
	return o
}

func (v *InClusterList) SetValues(val []uint16) error {
	*v = []*zcl.Zu16{}
	return v.AddValues(val...)
}

func (v *InClusterList) AddValues(val ...uint16) error {
	for _, a := range val {
		nv := zcl.Zu16(a)
		*v = append(*v, &nv)
	}
	return nil
}

func (v InClusterList) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, a := range v {
		vars = append(vars, a)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (v *InClusterList) UnmarshalZcl(b []byte) ([]byte, error) {
	*v = []*zcl.Zu16{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(zcl.Zu16)
		*v = append(*v, nv)
		return nv
	})
}

func (v *InClusterList) SetValue(a zcl.Val) error {
	if nv, ok := a.(*InClusterList); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (v InClusterList) String() string {
	var s []string
	for _, a := range v {
		s = append(s, zcl.Sprintf("%v", a))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

func (LogicalType) Name() string        { return `Logical Type` }
func (LogicalType) Description() string { return `` }

type LogicalType zcl.Zenum8

func (v *LogicalType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *LogicalType) Value() zcl.Val     { return v }

func (v LogicalType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *LogicalType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = LogicalType(*nt)
	return br, err
}

func (v LogicalType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *LogicalType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LogicalType(*a)
	return nil
}

func (v *LogicalType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = LogicalType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LogicalType) String() string {
	switch v {
	case 0x00:
		return "Coordinator"
	case 0x01:
		return "Router"
	case 0x10:
		return "End Device"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v LogicalType) IsCoordinator() bool { return v == 0x00 }
func (v LogicalType) IsRouter() bool      { return v == 0x01 }
func (v LogicalType) IsEndDevice() bool   { return v == 0x10 }
func (v *LogicalType) SetCoordinator()    { *v = 0x00 }
func (v *LogicalType) SetRouter()         { *v = 0x01 }
func (v *LogicalType) SetEndDevice()      { *v = 0x10 }

func (LogicalType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Coordinator"},
		{Value: 0x01, Name: "Router"},
		{Value: 0x10, Name: "End Device"},
	}
}

func (Lqi) Name() string { return `Lqi` }
func (Lqi) Description() string {
	return `is the estimated link quality for RF transmissions from this device`
}

// Lqi is the estimated link quality for RF transmissions from this device
type Lqi zcl.Zu8

func (v *Lqi) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *Lqi) Value() zcl.Val     { return v }

func (v Lqi) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *Lqi) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = Lqi(*nt)
	return br, err
}

func (v Lqi) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *Lqi) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Lqi(*a)
	return nil
}

func (v *Lqi) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = Lqi(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Lqi) String() string {
	return zcl.Percent.Format(float64(v) / 2.55)
}

func (MacCapabilityFlags) Name() string        { return `MAC Capability Flags` }
func (MacCapabilityFlags) Description() string { return `` }

type MacCapabilityFlags zcl.Zbmp8

func (v *MacCapabilityFlags) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *MacCapabilityFlags) Value() zcl.Val     { return v }

func (v MacCapabilityFlags) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *MacCapabilityFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = MacCapabilityFlags(*nt)
	return br, err
}

func (v MacCapabilityFlags) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *MacCapabilityFlags) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MacCapabilityFlags(*a)
	return nil
}

func (v *MacCapabilityFlags) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = MacCapabilityFlags(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MacCapabilityFlags) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v MacCapabilityFlags) IsAlternatePanCoordinator() bool { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v MacCapabilityFlags) IsFullFunctionDevice() bool      { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v MacCapabilityFlags) IsMainsPowered() bool            { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v MacCapabilityFlags) IsReceiverOnWhenIdle() bool      { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v MacCapabilityFlags) IsSupportsSecuredFrames() bool   { return zcl.BitmapTest([]byte(v[:]), 6) }
func (v MacCapabilityFlags) IsAllocateAddress() bool         { return zcl.BitmapTest([]byte(v[:]), 7) }
func (v *MacCapabilityFlags) SetAlternatePanCoordinator(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *MacCapabilityFlags) SetFullFunctionDevice(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *MacCapabilityFlags) SetMainsPowered(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}
func (v *MacCapabilityFlags) SetReceiverOnWhenIdle(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b))
}
func (v *MacCapabilityFlags) SetSupportsSecuredFrames(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 6, b))
}
func (v *MacCapabilityFlags) SetAllocateAddress(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 7, b))
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

func (ManufacturerCode) Name() string        { return `Manufacturer Code` }
func (ManufacturerCode) Description() string { return `` }

type ManufacturerCode zcl.Zenum16

func (v *ManufacturerCode) TypeID() zcl.TypeID { return new(zcl.Zenum16).TypeID() }
func (v *ManufacturerCode) Value() zcl.Val     { return v }

func (v ManufacturerCode) MarshalZcl() ([]byte, error) { return zcl.Zenum16(v).MarshalZcl() }

func (v *ManufacturerCode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum16)
	br, err := nt.UnmarshalZcl(b)
	*v = ManufacturerCode(*nt)
	return br, err
}

func (v ManufacturerCode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum16(v))
}

func (v *ManufacturerCode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ManufacturerCode(*a)
	return nil
}

func (v *ManufacturerCode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum16); ok {
		*v = ManufacturerCode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ManufacturerCode) String() string {
	switch v {
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
	case 0x1166:
		return "innr"
	case 0x117C:
		return "Ikea"
	}
	return zcl.Sprintf("%v", zcl.Zenum16(v))
}

func (v ManufacturerCode) IsPhilips() bool { return v == 0x100B }
func (v ManufacturerCode) IsLegrand() bool { return v == 0x1021 }
func (v ManufacturerCode) IsNxp() bool     { return v == 0x1037 }
func (v ManufacturerCode) IsUbisys() bool  { return v == 0x10F2 }
func (v ManufacturerCode) IsXiaomi() bool  { return v == 0x115F }
func (v ManufacturerCode) IsInnr() bool    { return v == 0x1166 }
func (v ManufacturerCode) IsIkea() bool    { return v == 0x117C }
func (v *ManufacturerCode) SetPhilips()    { *v = 0x100B }
func (v *ManufacturerCode) SetLegrand()    { *v = 0x1021 }
func (v *ManufacturerCode) SetNxp()        { *v = 0x1037 }
func (v *ManufacturerCode) SetUbisys()     { *v = 0x10F2 }
func (v *ManufacturerCode) SetXiaomi()     { *v = 0x115F }
func (v *ManufacturerCode) SetInnr()       { *v = 0x1166 }
func (v *ManufacturerCode) SetIkea()       { *v = 0x117C }

func (ManufacturerCode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x100B, Name: "Philips"},
		{Value: 0x1021, Name: "Legrand"},
		{Value: 0x1037, Name: "NXP"},
		{Value: 0x10F2, Name: "Ubisys"},
		{Value: 0x115F, Name: "Xiaomi"},
		{Value: 0x1166, Name: "innr"},
		{Value: 0x117C, Name: "Ikea"},
	}
}

func (MaxBufferSize) Name() string { return `Max Buffer Size` }
func (MaxBufferSize) Description() string {
	return `specifies the maximum size, in octets, of the network sub-layer data unit (NSDU) for this node.
This is the maximum size of data or commands passed to or from the application by the application support
sub-layer, before any fragmentation or re-assembly.`
}

// MaxBufferSize specifies the maximum size, in octets, of the network sub-layer data unit (NSDU) for this node.
// This is the maximum size of data or commands passed to or from the application by the application support
// sub-layer, before any fragmentation or re-assembly.
type MaxBufferSize zcl.Zu8

func (v *MaxBufferSize) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *MaxBufferSize) Value() zcl.Val     { return v }

func (v MaxBufferSize) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *MaxBufferSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxBufferSize(*nt)
	return br, err
}

func (v MaxBufferSize) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *MaxBufferSize) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxBufferSize(*a)
	return nil
}

func (v *MaxBufferSize) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = MaxBufferSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxBufferSize) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (MaxRxSize) Name() string { return `Max RX size` }
func (MaxRxSize) Description() string {
	return `specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred to this node in one single message transfer.`
}

// MaxRxSize specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred to this node in one single message transfer.
type MaxRxSize zcl.Zu16

func (v *MaxRxSize) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MaxRxSize) Value() zcl.Val     { return v }

func (v MaxRxSize) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MaxRxSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxRxSize(*nt)
	return br, err
}

func (v MaxRxSize) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MaxRxSize) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxRxSize(*a)
	return nil
}

func (v *MaxRxSize) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MaxRxSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxRxSize) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (MaxTxSize) Name() string { return `Max TX size` }
func (MaxTxSize) Description() string {
	return `specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred from
this node in one single message transfer.`
}

// MaxTxSize specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred from
// this node in one single message transfer.
type MaxTxSize zcl.Zu16

func (v *MaxTxSize) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MaxTxSize) Value() zcl.Val     { return v }

func (v MaxTxSize) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MaxTxSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxTxSize(*nt)
	return br, err
}

func (v MaxTxSize) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MaxTxSize) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxTxSize(*a)
	return nil
}

func (v *MaxTxSize) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MaxTxSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxTxSize) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (NwkAddress) Name() string        { return `NWK Address` }
func (NwkAddress) Description() string { return `is a 16-bit Network address` }

// NwkAddress is a 16-bit Network address
type NwkAddress zcl.Zu16

func (v *NwkAddress) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *NwkAddress) Value() zcl.Val     { return v }

func (v NwkAddress) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *NwkAddress) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = NwkAddress(*nt)
	return br, err
}

func (v NwkAddress) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *NwkAddress) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NwkAddress(*a)
	return nil
}

func (v *NwkAddress) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = NwkAddress(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NwkAddress) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (NeighborTable) Name() string        { return `Neighbor Table` }
func (NeighborTable) Description() string { return `` }

type NeighborTable []*NeighborTableEntry

func (v *NeighborTable) TypeID() zcl.TypeID { return new(zcl.Zarray).TypeID() }
func (v *NeighborTable) Value() zcl.Val     { return v }

func (NeighborTable) ArrayTypeID() zcl.TypeID { return new(NeighborTableEntry).TypeID() }

func (v *NeighborTable) ArrayValues() (o []NeighborTableEntry) {
	for _, a := range *v {
		o = append(o, *a)
	}
	return o
}

func (v *NeighborTable) SetValues(val []NeighborTableEntry) error {
	*v = []*NeighborTableEntry{}
	return v.AddValues(val...)
}

func (v *NeighborTable) AddValues(val ...NeighborTableEntry) error {
	for _, a := range val {
		nv := a
		*v = append(*v, &nv)
	}
	return nil
}

func (v NeighborTable) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, a := range v {
		vars = append(vars, a)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (v *NeighborTable) UnmarshalZcl(b []byte) ([]byte, error) {
	*v = []*NeighborTableEntry{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(NeighborTableEntry)
		*v = append(*v, nv)
		return nv
	})
}

func (v *NeighborTable) SetValue(a zcl.Val) error {
	if nv, ok := a.(*NeighborTable); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NeighborTable) String() string {
	var s []string
	for _, a := range v {
		s = append(s, zcl.Sprintf("%v", a))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

func (NeighborTableEntry) Name() string        { return `Neighbor Table Entry` }
func (NeighborTableEntry) Description() string { return `` }

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

func (v *NeighborTableEntry) TypeID() zcl.TypeID { return zcl.TypeID(76) } // struct
func (v *NeighborTableEntry) Value() zcl.Val     { return v }
func (v NeighborTableEntry) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	var err error
	_ = tmp2
	_ = err

	{
		if tmp, err = v.PanId.MarshalZcl(); err != nil {
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
	{
		tmp = []byte{}
		tmp2 = uint32(v.NeighborType&0x03) << 6

		data = append(data, tmp...)
	}
	{
		tmp2 |= uint32(v.RxOnWhenIdle&0x03) << 4

		data = append(data, tmp...)
	}
	{
		tmp2 |= uint32(v.Relationship&0x07) << 1
		tmp = []byte{uint8(tmp2)}

		data = append(data, tmp...)
	}
	{
		tmp2 = uint32(v.PermitJoining&0x03) << 6
		tmp = []byte{uint8(tmp2)}

		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Depth.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Lqi.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (v *NeighborTableEntry) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2
	_ = err

	if b, err = (&v.PanId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.IeeeAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if len(b) == 0 {
		return b, zcl.ErrNotEnoughData
	}
	v.NeighborType = NeighborType((b[0] >> 6) & 0x03)

	v.RxOnWhenIdle = RxOnWhenIdle((b[0] >> 4) & 0x03)

	v.Relationship = Relationship((b[0] >> 1) & 0x07)
	b = b[1:]

	if len(b) == 0 {
		return b, zcl.ErrNotEnoughData
	}
	v.PermitJoining = PermitJoining((b[0] >> 6) & 0x03)
	b = b[1:]

	if b, err = (&v.Depth).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Lqi).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v *NeighborTableEntry) SetValue(a zcl.Val) error {
	if nv, ok := a.(*NeighborTableEntry); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrNotImpl
}

func (v *NeighborTableEntry) String() string {
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
		v.PanId,
		v.IeeeAddress,
		v.NwkAddress,
		v.NeighborType,
		v.RxOnWhenIdle,
		v.Relationship,
		v.PermitJoining,
		v.Depth,
		v.Lqi,
	)
}

func (NeighborType) Name() string        { return `Neighbor Type` }
func (NeighborType) Description() string { return `` }

type NeighborType zcl.Zenum8

func (v *NeighborType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *NeighborType) Value() zcl.Val     { return v }

func (v NeighborType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *NeighborType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = NeighborType(*nt)
	return br, err
}

func (v NeighborType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *NeighborType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NeighborType(*a)
	return nil
}

func (v *NeighborType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = NeighborType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NeighborType) String() string {
	switch v {
	case 0x00:
		return "Coordinator"
	case 0x01:
		return "Router"
	case 0x03:
		return "End device"
	case 0x04:
		return "Unknown"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v NeighborType) IsCoordinator() bool { return v == 0x00 }
func (v NeighborType) IsRouter() bool      { return v == 0x01 }
func (v NeighborType) IsEndDevice() bool   { return v == 0x03 }
func (v NeighborType) IsUnknown() bool     { return v == 0x04 }
func (v *NeighborType) SetCoordinator()    { *v = 0x00 }
func (v *NeighborType) SetRouter()         { *v = 0x01 }
func (v *NeighborType) SetEndDevice()      { *v = 0x03 }
func (v *NeighborType) SetUnknown()        { *v = 0x04 }

func (NeighborType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Coordinator"},
		{Value: 0x01, Name: "Router"},
		{Value: 0x03, Name: "End device"},
		{Value: 0x04, Name: "Unknown"},
	}
}

func (NodeDescSize) Name() string        { return `Node Desc Size` }
func (NodeDescSize) Description() string { return `Size in bytes of the Node Descriptor` }

// NodeDescSize Size in bytes of the Node Descriptor
type NodeDescSize zcl.Zu8

func (v *NodeDescSize) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *NodeDescSize) Value() zcl.Val     { return v }

func (v NodeDescSize) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *NodeDescSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = NodeDescSize(*nt)
	return br, err
}

func (v NodeDescSize) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *NodeDescSize) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NodeDescSize(*a)
	return nil
}

func (v *NodeDescSize) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = NodeDescSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NodeDescSize) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (NodeDescriptor) Name() string        { return `Node Descriptor` }
func (NodeDescriptor) Description() string { return `` }

type NodeDescriptor struct {
	LogicalType                LogicalType
	ComplexDescriptorAvailable ComplexDescriptorAvailable
	UserDescriptorAvailable    UserDescriptorAvailable
	ApsFlags                   ApsFlags
	FrequencyBands             FrequencyBands
	MacCapabilityFlags         MacCapabilityFlags
	ManufacturerCode           ManufacturerCode
	// MaxBufferSize specifies the maximum size, in octets, of the network sub-layer data unit (NSDU) for this node.
	// This is the maximum size of data or commands passed to or from the application by the application support
	// sub-layer, before any fragmentation or re-assembly.
	MaxBufferSize MaxBufferSize
	// MaxRxSize specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred to this node in one single message transfer.
	MaxRxSize  MaxRxSize
	ServerMask ServerMask
	// MaxTxSize specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred from
	// this node in one single message transfer.
	MaxTxSize            MaxTxSize
	DescriptorCapability DescriptorCapability
}

func (v *NodeDescriptor) TypeID() zcl.TypeID { return zcl.TypeID(76) } // struct
func (v *NodeDescriptor) Value() zcl.Val     { return v }
func (v NodeDescriptor) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	var err error
	_ = tmp2
	_ = err

	{
		tmp = []byte{}
		tmp2 = uint32(v.LogicalType & 0x07)

		data = append(data, tmp...)
	}
	{
		tmp2 |= uint32(v.ComplexDescriptorAvailable&0x01) << 3

		data = append(data, tmp...)
	}
	{
		tmp2 |= uint32(v.UserDescriptorAvailable&0x01) << 4
		tmp = []byte{uint8(tmp2)}

		data = append(data, tmp...)
	}
	{
		tmp = []byte{}
		tmp2 = uint32(v.ApsFlags & 0x07)

		data = append(data, tmp...)
	}
	{
		tmp2 |= uint32(v.FrequencyBands[0]&0x1F) << 3
		tmp = []byte{uint8(tmp2)}

		data = append(data, tmp...)
	}
	{
		if tmp, err = v.MacCapabilityFlags.MarshalZcl(); err != nil {
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
	{
		if tmp, err = v.DescriptorCapability.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (v *NodeDescriptor) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2
	_ = err

	if len(b) == 0 {
		return b, zcl.ErrNotEnoughData
	}
	v.LogicalType = LogicalType(b[0] & 0x07)

	v.ComplexDescriptorAvailable = ComplexDescriptorAvailable((b[0] >> 3) & 0x01)

	v.UserDescriptorAvailable = UserDescriptorAvailable((b[0] >> 4) & 0x01)
	b = b[1:]

	if len(b) == 0 {
		return b, zcl.ErrNotEnoughData
	}
	v.ApsFlags = ApsFlags(b[0] & 0x07)

	v.FrequencyBands[0] = (b[0] >> 3) & 0x1F
	b = b[1:]

	if b, err = (&v.MacCapabilityFlags).UnmarshalZcl(b); err != nil {
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

	if b, err = (&v.DescriptorCapability).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v *NodeDescriptor) SetValue(a zcl.Val) error {
	if nv, ok := a.(*NodeDescriptor); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrNotImpl
}

func (v *NodeDescriptor) String() string {
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
		v.LogicalType,
		v.ComplexDescriptorAvailable,
		v.UserDescriptorAvailable,
		v.ApsFlags,
		v.FrequencyBands,
		v.MacCapabilityFlags,
		v.ManufacturerCode,
		v.MaxBufferSize,
		v.MaxRxSize,
		v.ServerMask,
		v.MaxTxSize,
		v.DescriptorCapability,
	)
}

func (OutClusterList) Name() string        { return `Out Cluster List` }
func (OutClusterList) Description() string { return `is a list of output clusters` }

// OutClusterList is a list of output clusters
type OutClusterList []*zcl.Zu16

func (v *OutClusterList) TypeID() zcl.TypeID { return new(zcl.Zset).TypeID() }
func (v *OutClusterList) Value() zcl.Val     { return v }

func (OutClusterList) ArrayTypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }

func (v *OutClusterList) ArrayValues() (o []uint16) {
	for _, a := range *v {
		o = append(o, uint16(*a))
	}
	return o
}

func (v *OutClusterList) SetValues(val []uint16) error {
	*v = []*zcl.Zu16{}
	return v.AddValues(val...)
}

func (v *OutClusterList) AddValues(val ...uint16) error {
	for _, a := range val {
		nv := zcl.Zu16(a)
		*v = append(*v, &nv)
	}
	return nil
}

func (v OutClusterList) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, a := range v {
		vars = append(vars, a)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (v *OutClusterList) UnmarshalZcl(b []byte) ([]byte, error) {
	*v = []*zcl.Zu16{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(zcl.Zu16)
		*v = append(*v, nv)
		return nv
	})
}

func (v *OutClusterList) SetValue(a zcl.Val) error {
	if nv, ok := a.(*OutClusterList); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OutClusterList) String() string {
	var s []string
	for _, a := range v {
		s = append(s, zcl.Sprintf("%v", a))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

func (PermitJoining) Name() string        { return `Permit Joining` }
func (PermitJoining) Description() string { return `` }

type PermitJoining zcl.Zenum8

func (v *PermitJoining) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *PermitJoining) Value() zcl.Val     { return v }

func (v PermitJoining) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *PermitJoining) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = PermitJoining(*nt)
	return br, err
}

func (v PermitJoining) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *PermitJoining) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PermitJoining(*a)
	return nil
}

func (v *PermitJoining) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = PermitJoining(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PermitJoining) String() string {
	switch v {
	case 0x00:
		return "Not permitted"
	case 0x01:
		return "Permitted"
	case 0x02:
		return "Unknown"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v PermitJoining) IsNotPermitted() bool { return v == 0x00 }
func (v PermitJoining) IsPermitted() bool    { return v == 0x01 }
func (v PermitJoining) IsUnknown() bool      { return v == 0x02 }
func (v *PermitJoining) SetNotPermitted()    { *v = 0x00 }
func (v *PermitJoining) SetPermitted()       { *v = 0x01 }
func (v *PermitJoining) SetUnknown()         { *v = 0x02 }

func (PermitJoining) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Not permitted"},
		{Value: 0x01, Name: "Permitted"},
		{Value: 0x02, Name: "Unknown"},
	}
}

func (PowerDescSize) Name() string        { return `Power Desc Size` }
func (PowerDescSize) Description() string { return `Size in bytes of the Power Descriptor` }

// PowerDescSize Size in bytes of the Power Descriptor
type PowerDescSize zcl.Zu8

func (v *PowerDescSize) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *PowerDescSize) Value() zcl.Val     { return v }

func (v PowerDescSize) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *PowerDescSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = PowerDescSize(*nt)
	return br, err
}

func (v PowerDescSize) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *PowerDescSize) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PowerDescSize(*a)
	return nil
}

func (v *PowerDescSize) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = PowerDescSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PowerDescSize) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (PowerDescriptor) Name() string        { return `Power Descriptor` }
func (PowerDescriptor) Description() string { return `` }

type PowerDescriptor struct {
	PowerMode          PowerMode
	ActivePowerSource  PowerSource
	CurrentPowerSource PowerSource
	PowerLevel         PowerLevel
}

func (v *PowerDescriptor) TypeID() zcl.TypeID { return zcl.TypeID(76) } // struct
func (v *PowerDescriptor) Value() zcl.Val     { return v }
func (v PowerDescriptor) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	var err error
	_ = tmp2
	_ = err

	{
		tmp = []byte{}
		tmp2 = uint32(v.PowerMode[0] & 0x0F)

		data = append(data, tmp...)
	}
	{
		tmp2 |= uint32(v.ActivePowerSource[0]&0x0F) << 4
		tmp = []byte{uint8(tmp2)}

		data = append(data, tmp...)
	}
	{
		tmp = []byte{}
		tmp2 = uint32(v.CurrentPowerSource[0] & 0x0F)

		data = append(data, tmp...)
	}
	{
		tmp2 |= uint32(v.PowerLevel&0x0F) << 4
		tmp = []byte{uint8(tmp2)}

		data = append(data, tmp...)
	}
	return data, nil
}

func (v *PowerDescriptor) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2
	_ = err

	if len(b) == 0 {
		return b, zcl.ErrNotEnoughData
	}
	v.PowerMode[0] = b[0] & 0x0F

	v.ActivePowerSource[0] = (b[0] >> 4) & 0x0F
	b = b[1:]

	if len(b) == 0 {
		return b, zcl.ErrNotEnoughData
	}
	v.CurrentPowerSource[0] = b[0] & 0x0F

	v.PowerLevel = PowerLevel(b[0]>>4) & 0x0F
	b = b[1:]

	return b, nil
}

func (v *PowerDescriptor) SetValue(a zcl.Val) error {
	if nv, ok := a.(*PowerDescriptor); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrNotImpl
}

func (v *PowerDescriptor) String() string {
	return zcl.Sprintf(
		"PowerDescriptor{"+zcl.StrJoin([]string{
			"PowerMode(%v)",
			"ActivePowerSource(%v)",
			"CurrentPowerSource(%v)",
			"PowerLevel(%v)",
		}, " ")+"}",
		v.PowerMode,
		v.ActivePowerSource,
		v.CurrentPowerSource,
		v.PowerLevel,
	)
}

func (PowerLevel) Name() string        { return `Power Level` }
func (PowerLevel) Description() string { return `` }

type PowerLevel zcl.Zenum8

func (v *PowerLevel) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *PowerLevel) Value() zcl.Val     { return v }

func (v PowerLevel) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *PowerLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = PowerLevel(*nt)
	return br, err
}

func (v PowerLevel) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *PowerLevel) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PowerLevel(*a)
	return nil
}

func (v *PowerLevel) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = PowerLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PowerLevel) String() string {
	switch v {
	case 0x00:
		return "Critical"
	case 0x04:
		return "33%"
	case 0x08:
		return "66%"
	case 0x0C:
		return "100%"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v PowerLevel) IsCritical() bool { return v == 0x00 }
func (v PowerLevel) Is33() bool       { return v == 0x04 }
func (v PowerLevel) Is66() bool       { return v == 0x08 }
func (v PowerLevel) Is100() bool      { return v == 0x0C }
func (v *PowerLevel) SetCritical()    { *v = 0x00 }
func (v *PowerLevel) Set33()          { *v = 0x04 }
func (v *PowerLevel) Set66()          { *v = 0x08 }
func (v *PowerLevel) Set100()         { *v = 0x0C }

func (PowerLevel) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Critical"},
		{Value: 0x04, Name: "33%"},
		{Value: 0x08, Name: "66%"},
		{Value: 0x0C, Name: "100%"},
	}
}

func (PowerMode) Name() string        { return `Power Mode` }
func (PowerMode) Description() string { return `` }

type PowerMode zcl.Zbmp8

func (v *PowerMode) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *PowerMode) Value() zcl.Val     { return v }

func (v PowerMode) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *PowerMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = PowerMode(*nt)
	return br, err
}

func (v PowerMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *PowerMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PowerMode(*a)
	return nil
}

func (v *PowerMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = PowerMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PowerMode) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v PowerMode) IsConstantPowerAvailable() bool       { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v PowerMode) IsRechargeableBatteryAvailable() bool { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v PowerMode) IsDisposableBatteryAvailable() bool   { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v PowerMode) IsCheckInPeriodically() bool          { return zcl.BitmapTest([]byte(v[:]), 4) }
func (v PowerMode) IsCheckInOnAction() bool              { return zcl.BitmapTest([]byte(v[:]), 5) }
func (v *PowerMode) SetConstantPowerAvailable(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *PowerMode) SetRechargeableBatteryAvailable(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *PowerMode) SetDisposableBatteryAvailable(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}
func (v *PowerMode) SetCheckInPeriodically(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 4, b))
}
func (v *PowerMode) SetCheckInOnAction(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 5, b)) }

func (PowerMode) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Constant Power Available"},
		{Value: 1, Name: "Rechargeable battery Available"},
		{Value: 2, Name: "Disposable Battery Available"},
		{Value: 4, Name: "Check In Periodically"},
		{Value: 5, Name: "Check In on Action"},
	}
}

func (PowerSource) Name() string        { return `Power Source` }
func (PowerSource) Description() string { return `` }

type PowerSource zcl.Zbmp8

func (v *PowerSource) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *PowerSource) Value() zcl.Val     { return v }

func (v PowerSource) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *PowerSource) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = PowerSource(*nt)
	return br, err
}

func (v PowerSource) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *PowerSource) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PowerSource(*a)
	return nil
}

func (v *PowerSource) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = PowerSource(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PowerSource) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v PowerSource) IsMainsPower() bool          { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v PowerSource) IsRechargeableBattery() bool { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v PowerSource) IsDisposableBattery() bool   { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v *PowerSource) SetMainsPower(b bool)       { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b)) }
func (v *PowerSource) SetRechargeableBattery(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *PowerSource) SetDisposableBattery(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}

func (PowerSource) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Mains power"},
		{Value: 1, Name: "Rechargeable battery"},
		{Value: 2, Name: "Disposable battery"},
	}
}

func (ProfileId) Name() string        { return `Profile ID` }
func (ProfileId) Description() string { return `` }

type ProfileId zcl.Zu16

func (v *ProfileId) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ProfileId) Value() zcl.Val     { return v }

func (v ProfileId) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ProfileId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ProfileId(*nt)
	return br, err
}

func (v ProfileId) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ProfileId) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ProfileId(*a)
	return nil
}

func (v *ProfileId) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ProfileId(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ProfileId) String() string {
	switch v {
	case 0x0104:
		return "Home Automation"
	case 0xC05E:
		return "Light Link"
	}
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (v ProfileId) IsHomeAutomation() bool { return v == 0x0104 }
func (v ProfileId) IsLightLink() bool      { return v == 0xC05E }
func (v *ProfileId) SetHomeAutomation()    { *v = 0x0104 }
func (v *ProfileId) SetLightLink()         { *v = 0xC05E }

func (ProfileId) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x0104, Name: "Home Automation"},
		{Value: 0xC05E, Name: "Light Link"},
	}
}

func (Relationship) Name() string        { return `Relationship` }
func (Relationship) Description() string { return `` }

type Relationship zcl.Zenum8

func (v *Relationship) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *Relationship) Value() zcl.Val     { return v }

func (v Relationship) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *Relationship) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = Relationship(*nt)
	return br, err
}

func (v Relationship) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *Relationship) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Relationship(*a)
	return nil
}

func (v *Relationship) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = Relationship(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Relationship) String() string {
	switch v {
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
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v Relationship) IsParent() bool        { return v == 0x00 }
func (v Relationship) IsChild() bool         { return v == 0x01 }
func (v Relationship) IsSibling() bool       { return v == 0x02 }
func (v Relationship) IsNone() bool          { return v == 0x03 }
func (v Relationship) IsPreviousChild() bool { return v == 0x04 }
func (v *Relationship) SetParent()           { *v = 0x00 }
func (v *Relationship) SetChild()            { *v = 0x01 }
func (v *Relationship) SetSibling()          { *v = 0x02 }
func (v *Relationship) SetNone()             { *v = 0x03 }
func (v *Relationship) SetPreviousChild()    { *v = 0x04 }

func (Relationship) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Parent"},
		{Value: 0x01, Name: "Child"},
		{Value: 0x02, Name: "Sibling"},
		{Value: 0x03, Name: "None"},
		{Value: 0x04, Name: "Previous Child"},
	}
}

func (RequestType) Name() string { return `Request Type` }
func (RequestType) Description() string {
	return `should be set to 1 if extended response is requested, 0 otherwise`
}

// RequestType should be set to 1 if extended response is requested, 0 otherwise
type RequestType zcl.Zenum8

func (v *RequestType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *RequestType) Value() zcl.Val     { return v }

func (v RequestType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *RequestType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = RequestType(*nt)
	return br, err
}

func (v RequestType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *RequestType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RequestType(*a)
	return nil
}

func (v *RequestType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = RequestType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RequestType) String() string {
	switch v {
	case 0x00:
		return "Single Device Response"
	case 0x01:
		return "Extended Response"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v RequestType) IsSingleDeviceResponse() bool { return v == 0x00 }
func (v RequestType) IsExtendedResponse() bool     { return v == 0x01 }
func (v *RequestType) SetSingleDeviceResponse()    { *v = 0x00 }
func (v *RequestType) SetExtendedResponse()        { *v = 0x01 }

func (RequestType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Single Device Response"},
		{Value: 0x01, Name: "Extended Response"},
	}
}

func (RxOnWhenIdle) Name() string        { return `Rx On When Idle` }
func (RxOnWhenIdle) Description() string { return `` }

type RxOnWhenIdle zcl.Zenum8

func (v *RxOnWhenIdle) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *RxOnWhenIdle) Value() zcl.Val     { return v }

func (v RxOnWhenIdle) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *RxOnWhenIdle) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = RxOnWhenIdle(*nt)
	return br, err
}

func (v RxOnWhenIdle) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *RxOnWhenIdle) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RxOnWhenIdle(*a)
	return nil
}

func (v *RxOnWhenIdle) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = RxOnWhenIdle(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RxOnWhenIdle) String() string {
	switch v {
	case 0x00:
		return "Receiver is off"
	case 0x01:
		return "Receiver is on"
	case 0x02:
		return "Unknown"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v RxOnWhenIdle) IsReceiverIsOff() bool { return v == 0x00 }
func (v RxOnWhenIdle) IsReceiverIsOn() bool  { return v == 0x01 }
func (v RxOnWhenIdle) IsUnknown() bool       { return v == 0x02 }
func (v *RxOnWhenIdle) SetReceiverIsOff()    { *v = 0x00 }
func (v *RxOnWhenIdle) SetReceiverIsOn()     { *v = 0x01 }
func (v *RxOnWhenIdle) SetUnknown()          { *v = 0x02 }

func (RxOnWhenIdle) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Receiver is off"},
		{Value: 0x01, Name: "Receiver is on"},
		{Value: 0x02, Name: "Unknown"},
	}
}

func (ScannedChannels) Name() string        { return `Scanned Channels` }
func (ScannedChannels) Description() string { return `` }

type ScannedChannels zcl.Zbmp32

func (v *ScannedChannels) TypeID() zcl.TypeID { return new(zcl.Zbmp32).TypeID() }
func (v *ScannedChannels) Value() zcl.Val     { return v }

func (v ScannedChannels) MarshalZcl() ([]byte, error) { return zcl.Zbmp32(v).MarshalZcl() }

func (v *ScannedChannels) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp32)
	br, err := nt.UnmarshalZcl(b)
	*v = ScannedChannels(*nt)
	return br, err
}

func (v ScannedChannels) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp32(v))
}

func (v *ScannedChannels) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ScannedChannels(*a)
	return nil
}

func (v *ScannedChannels) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp32); ok {
		*v = ScannedChannels(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ScannedChannels) String() string {
	return zcl.Sprintf("%v", zcl.Zbmp32(v))
}

func (ServerMask) Name() string        { return `Server Mask` }
func (ServerMask) Description() string { return `` }

type ServerMask zcl.Zbmp16

func (v *ServerMask) TypeID() zcl.TypeID { return new(zcl.Zbmp16).TypeID() }
func (v *ServerMask) Value() zcl.Val     { return v }

func (v ServerMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp16(v).MarshalZcl() }

func (v *ServerMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*v = ServerMask(*nt)
	return br, err
}

func (v ServerMask) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp16(v))
}

func (v *ServerMask) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ServerMask(*a)
	return nil
}

func (v *ServerMask) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp16); ok {
		*v = ServerMask(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ServerMask) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v ServerMask) IsPrimaryTrustCenter() bool       { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v ServerMask) IsBackupTrustCenter() bool        { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v ServerMask) IsPrimaryBindingTableCache() bool { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v ServerMask) IsBackupBindingTableCache() bool  { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v ServerMask) IsPrimaryDiscoveryCache() bool    { return zcl.BitmapTest([]byte(v[:]), 4) }
func (v ServerMask) IsBackupDiscoveryCache() bool     { return zcl.BitmapTest([]byte(v[:]), 5) }
func (v ServerMask) IsNetworkManager() bool           { return zcl.BitmapTest([]byte(v[:]), 6) }
func (v *ServerMask) SetPrimaryTrustCenter(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *ServerMask) SetBackupTrustCenter(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b)) }
func (v *ServerMask) SetPrimaryBindingTableCache(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}
func (v *ServerMask) SetBackupBindingTableCache(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b))
}
func (v *ServerMask) SetPrimaryDiscoveryCache(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 4, b))
}
func (v *ServerMask) SetBackupDiscoveryCache(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 5, b))
}
func (v *ServerMask) SetNetworkManager(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 6, b)) }

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

func (SimpleDescSizeList) Name() string { return `Simple Desc Size List` }
func (SimpleDescSizeList) Description() string {
	return `List of sizes for the different Simple Descriptors`
}

// SimpleDescSizeList List of sizes for the different Simple Descriptors
type SimpleDescSizeList []*zcl.Zu8

func (v *SimpleDescSizeList) TypeID() zcl.TypeID { return new(zcl.Zset).TypeID() }
func (v *SimpleDescSizeList) Value() zcl.Val     { return v }

func (SimpleDescSizeList) ArrayTypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }

func (v *SimpleDescSizeList) ArrayValues() (o []uint8) {
	for _, a := range *v {
		o = append(o, uint8(*a))
	}
	return o
}

func (v *SimpleDescSizeList) SetValues(val []uint8) error {
	*v = []*zcl.Zu8{}
	return v.AddValues(val...)
}

func (v *SimpleDescSizeList) AddValues(val ...uint8) error {
	for _, a := range val {
		nv := zcl.Zu8(a)
		*v = append(*v, &nv)
	}
	return nil
}

func (v SimpleDescSizeList) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, a := range v {
		vars = append(vars, a)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (v *SimpleDescSizeList) UnmarshalZcl(b []byte) ([]byte, error) {
	*v = []*zcl.Zu8{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(zcl.Zu8)
		*v = append(*v, nv)
		return nv
	})
}

func (v *SimpleDescSizeList) SetValue(a zcl.Val) error {
	if nv, ok := a.(*SimpleDescSizeList); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SimpleDescSizeList) String() string {
	var s []string
	for _, a := range v {
		s = append(s, zcl.Sprintf("%v", a))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

func (SimpleDescriptor) Name() string        { return `Simple Descriptor` }
func (SimpleDescriptor) Description() string { return `` }

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

func (v *SimpleDescriptor) TypeID() zcl.TypeID { return zcl.TypeID(76) } // struct
func (v *SimpleDescriptor) Value() zcl.Val     { return v }
func (v SimpleDescriptor) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	var err error
	_ = tmp2
	_ = err

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

func (v *SimpleDescriptor) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2
	_ = err

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

func (v *SimpleDescriptor) SetValue(a zcl.Val) error {
	if nv, ok := a.(*SimpleDescriptor); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrNotImpl
}

func (v *SimpleDescriptor) String() string {
	return zcl.Sprintf(
		"SimpleDescriptor{"+zcl.StrJoin([]string{
			"Endpoint(%v)",
			"ProfileId(%v)",
			"DeviceType(%v)",
			"DeviceVersion(%v)",
			"InClusterList(%v)",
			"OutClusterList(%v)",
		}, " ")+"}",
		v.Endpoint,
		v.ProfileId,
		v.DeviceType,
		v.DeviceVersion,
		v.InClusterList,
		v.OutClusterList,
	)
}

func (SourceAddress) Name() string        { return `Source Address` }
func (SourceAddress) Description() string { return `of device generating the request` }

// SourceAddress of device generating the request
type SourceAddress zcl.Zuid

func (v *SourceAddress) TypeID() zcl.TypeID { return new(zcl.Zuid).TypeID() }
func (v *SourceAddress) Value() zcl.Val     { return v }

func (v SourceAddress) MarshalZcl() ([]byte, error) { return zcl.Zuid(v).MarshalZcl() }

func (v *SourceAddress) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*v = SourceAddress(*nt)
	return br, err
}

func (v SourceAddress) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zuid(v))
}

func (v *SourceAddress) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zuid)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SourceAddress(*a)
	return nil
}

func (v *SourceAddress) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zuid); ok {
		*v = SourceAddress(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SourceAddress) String() string {
	return zcl.Sprintf("%v", zcl.Zuid(v))
}

func (SourceEndpoint) Name() string        { return `Source Endpoint` }
func (SourceEndpoint) Description() string { return `of device generating the request` }

// SourceEndpoint of device generating the request
type SourceEndpoint zcl.Zu8

func (v *SourceEndpoint) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *SourceEndpoint) Value() zcl.Val     { return v }

func (v SourceEndpoint) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *SourceEndpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = SourceEndpoint(*nt)
	return br, err
}

func (v SourceEndpoint) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *SourceEndpoint) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SourceEndpoint(*a)
	return nil
}

func (v *SourceEndpoint) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = SourceEndpoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SourceEndpoint) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (StartIndex) Name() string { return `Start Index` }
func (StartIndex) Description() string {
	return `provides the starting index for the requested elements of the associated list.`
}

// StartIndex provides the starting index for the requested elements of the associated list.
type StartIndex zcl.Zu8

func (v *StartIndex) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *StartIndex) Value() zcl.Val     { return v }

func (v StartIndex) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *StartIndex) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = StartIndex(*nt)
	return br, err
}

func (v StartIndex) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *StartIndex) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = StartIndex(*a)
	return nil
}

func (v *StartIndex) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = StartIndex(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v StartIndex) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (Status) Name() string { return `Status` }
func (Status) Description() string {
	return "Code, command is normally empty unless status is `Success`"
}

// Status Code, command is normally empty unless status is `Success`
type Status zcl.Zenum8

func (v *Status) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *Status) Value() zcl.Val     { return v }

func (v Status) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *Status) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = Status(*nt)
	return br, err
}

func (v Status) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *Status) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Status(*a)
	return nil
}

func (v *Status) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = Status(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Status) String() string {
	switch v {
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
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v Status) IsSuccess() bool            { return v == 0x00 }
func (v Status) IsInvalidRequestType() bool { return v == 0x80 }
func (v Status) IsDeviceNotFound() bool     { return v == 0x81 }
func (v Status) IsInvalidEndpoint() bool    { return v == 0x82 }
func (v Status) IsNotActive() bool          { return v == 0x83 }
func (v Status) IsNotSupported() bool       { return v == 0x84 }
func (v Status) IsTimeout() bool            { return v == 0x85 }
func (v Status) IsNoMatch() bool            { return v == 0x86 }
func (v Status) IsNoEntry() bool            { return v == 0x88 }
func (v Status) IsNoDescriptor() bool       { return v == 0x89 }
func (v Status) IsInsufficientSpace() bool  { return v == 0x8A }
func (v Status) IsNotPermitted() bool       { return v == 0x8B }
func (v Status) IsTableFull() bool          { return v == 0x8C }
func (v Status) IsNotAuthorized() bool      { return v == 0x8D }
func (v *Status) SetSuccess()               { *v = 0x00 }
func (v *Status) SetInvalidRequestType()    { *v = 0x80 }
func (v *Status) SetDeviceNotFound()        { *v = 0x81 }
func (v *Status) SetInvalidEndpoint()       { *v = 0x82 }
func (v *Status) SetNotActive()             { *v = 0x83 }
func (v *Status) SetNotSupported()          { *v = 0x84 }
func (v *Status) SetTimeout()               { *v = 0x85 }
func (v *Status) SetNoMatch()               { *v = 0x86 }
func (v *Status) SetNoEntry()               { *v = 0x88 }
func (v *Status) SetNoDescriptor()          { *v = 0x89 }
func (v *Status) SetInsufficientSpace()     { *v = 0x8A }
func (v *Status) SetNotPermitted()          { *v = 0x8B }
func (v *Status) SetTableFull()             { *v = 0x8C }
func (v *Status) SetNotAuthorized()         { *v = 0x8D }

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

func (TotalEntries) Name() string { return `Total Entries` }
func (TotalEntries) Description() string {
	return `is the total number of entries that can be queried for`
}

// TotalEntries is the total number of entries that can be queried for
type TotalEntries zcl.Zu8

func (v *TotalEntries) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *TotalEntries) Value() zcl.Val     { return v }

func (v TotalEntries) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *TotalEntries) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = TotalEntries(*nt)
	return br, err
}

func (v TotalEntries) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *TotalEntries) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TotalEntries(*a)
	return nil
}

func (v *TotalEntries) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = TotalEntries(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TotalEntries) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (TotalTransmissions) Name() string        { return `Total Transmissions` }
func (TotalTransmissions) Description() string { return `` }

type TotalTransmissions zcl.Zu16

func (v *TotalTransmissions) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *TotalTransmissions) Value() zcl.Val     { return v }

func (v TotalTransmissions) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *TotalTransmissions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = TotalTransmissions(*nt)
	return br, err
}

func (v TotalTransmissions) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *TotalTransmissions) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TotalTransmissions(*a)
	return nil
}

func (v *TotalTransmissions) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = TotalTransmissions(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TotalTransmissions) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (TransmissionFailures) Name() string        { return `Transmission Failures` }
func (TransmissionFailures) Description() string { return `` }

type TransmissionFailures zcl.Zu16

func (v *TransmissionFailures) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *TransmissionFailures) Value() zcl.Val     { return v }

func (v TransmissionFailures) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *TransmissionFailures) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = TransmissionFailures(*nt)
	return br, err
}

func (v TransmissionFailures) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *TransmissionFailures) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TransmissionFailures(*a)
	return nil
}

func (v *TransmissionFailures) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = TransmissionFailures(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TransmissionFailures) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (UnknownU8) Name() string        { return `Unknown u8` }
func (UnknownU8) Description() string { return `` }

type UnknownU8 zcl.Zu8

func (v *UnknownU8) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *UnknownU8) Value() zcl.Val     { return v }

func (v UnknownU8) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *UnknownU8) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = UnknownU8(*nt)
	return br, err
}

func (v UnknownU8) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *UnknownU8) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = UnknownU8(*a)
	return nil
}

func (v *UnknownU8) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = UnknownU8(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v UnknownU8) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (UserDescriptor) Name() string { return `User Descriptor` }
func (UserDescriptor) Description() string {
	return `is a user provided ASCII string of 16 characters set on a remote device.
If the string is shorter than 16 characters it should be padded with space characters (0x20).
Control characters (0x00-0x1f) are not allowed.`
}

// UserDescriptor is a user provided ASCII string of 16 characters set on a remote device.
// If the string is shorter than 16 characters it should be padded with space characters (0x20).
// Control characters (0x00-0x1f) are not allowed.
type UserDescriptor zcl.Zcstring

func (v *UserDescriptor) TypeID() zcl.TypeID { return new(zcl.Zcstring).TypeID() }
func (v *UserDescriptor) Value() zcl.Val     { return v }

func (v UserDescriptor) MarshalZcl() ([]byte, error) { return zcl.Zcstring(v).MarshalZcl() }

func (v *UserDescriptor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*v = UserDescriptor(*nt)
	return br, err
}

func (v UserDescriptor) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zcstring(v))
}

func (v *UserDescriptor) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zcstring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = UserDescriptor(*a)
	return nil
}

func (v *UserDescriptor) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zcstring); ok {
		*v = UserDescriptor(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v UserDescriptor) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v))
}

func (UserDescriptorAvailable) Name() string        { return `User Descriptor Available` }
func (UserDescriptorAvailable) Description() string { return `` }

type UserDescriptorAvailable zcl.Zbool

func (v *UserDescriptorAvailable) TypeID() zcl.TypeID { return new(zcl.Zbool).TypeID() }
func (v *UserDescriptorAvailable) Value() zcl.Val     { return v }

func (v UserDescriptorAvailable) MarshalZcl() ([]byte, error) { return zcl.Zbool(v).MarshalZcl() }

func (v *UserDescriptorAvailable) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*v = UserDescriptorAvailable(*nt)
	return br, err
}

func (v UserDescriptorAvailable) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbool(v))
}

func (v *UserDescriptorAvailable) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbool)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = UserDescriptorAvailable(*a)
	return nil
}

func (v *UserDescriptorAvailable) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbool); ok {
		*v = UserDescriptorAvailable(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v UserDescriptorAvailable) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(v))
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

type NwkAddressRequestHandler interface {
	HandleNwkAddressRequest(frame Frame, cmd *NwkAddressRequest) (*NwkAddressResponse, error)
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
func (NwkAddressRequest) Name() string { return `NWK Address Request` }

// Description of the command
func (NwkAddressRequest) Description() string {
	return `queries the 16-bit address of the Remote Device based on its known IEEE address.
The destination addressing on this command shall be unicast or broadcast to all
devices supporting RX on when Idle (0xFFFD)`
}

// ID of the command
func (NwkAddressRequest) ID() CommandID { return NwkAddressRequestCommand }

// Required
func (NwkAddressRequest) Required() bool { return true }

// Cluster ID of the command
func (NwkAddressRequest) Cluster() zcl.ClusterID { return 0x0000 }

// MnfCode returns the manufacturer code (if any) of the command
func (NwkAddressRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (NwkAddressRequest) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

func (v *NwkAddressRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h NwkAddressRequestHandler
	if h, found = handler.(NwkAddressRequestHandler); found {
		rsp, err = h.HandleNwkAddressRequest(frame, v)
	} else {
		rsp = &NwkAddressResponse{}
		rsp.(*NwkAddressResponse).Status.SetNotSupported()
	}
	return
}

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

type NwkAddressResponseHandler interface {
	HandleNwkAddressResponse(frame Frame, cmd *NwkAddressResponse) error
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
func (NwkAddressResponse) Name() string { return `NWK Address Response` }

// Description of the command
func (NwkAddressResponse) Description() string { return `` }

// ID of the command
func (NwkAddressResponse) ID() CommandID { return NwkAddressResponseCommand }

// Required
func (NwkAddressResponse) Required() bool { return true }

// Cluster ID of the command
func (NwkAddressResponse) Cluster() zcl.ClusterID { return 0x8000 }

// MnfCode returns the manufacturer code (if any) of the command
func (NwkAddressResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (NwkAddressResponse) MarshalJSON() ([]byte, error) { return []byte("32768"), nil }

func (v *NwkAddressResponse) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h NwkAddressResponseHandler
	if h, found = handler.(NwkAddressResponseHandler); found {
		err = h.HandleNwkAddressResponse(frame, v)
	}
	return
}

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

type IeeeAddressRequestHandler interface {
	HandleIeeeAddressRequest(frame Frame, cmd *IeeeAddressRequest) (*IeeeAddressResponse, error)
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
func (IeeeAddressRequest) Name() string { return `IEEE Address Request` }

// Description of the command
func (IeeeAddressRequest) Description() string {
	return `queries the 64-bit IEEE address of the Remote Device based on its known 16-bit NWK address.
The command should always be sent by unicast.`
}

// ID of the command
func (IeeeAddressRequest) ID() CommandID { return IeeeAddressRequestCommand }

// Required
func (IeeeAddressRequest) Required() bool { return true }

// Cluster ID of the command
func (IeeeAddressRequest) Cluster() zcl.ClusterID { return 0x0001 }

// MnfCode returns the manufacturer code (if any) of the command
func (IeeeAddressRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (IeeeAddressRequest) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

func (v *IeeeAddressRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h IeeeAddressRequestHandler
	if h, found = handler.(IeeeAddressRequestHandler); found {
		rsp, err = h.HandleIeeeAddressRequest(frame, v)
	} else {
		rsp = &IeeeAddressResponse{}
		rsp.(*IeeeAddressResponse).Status.SetNotSupported()
	}
	return
}

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

type IeeeAddressResponseHandler interface {
	HandleIeeeAddressResponse(frame Frame, cmd *IeeeAddressResponse) error
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
func (IeeeAddressResponse) Name() string { return `IEEE Address Response` }

// Description of the command
func (IeeeAddressResponse) Description() string { return `` }

// ID of the command
func (IeeeAddressResponse) ID() CommandID { return IeeeAddressResponseCommand }

// Required
func (IeeeAddressResponse) Required() bool { return true }

// Cluster ID of the command
func (IeeeAddressResponse) Cluster() zcl.ClusterID { return 0x8001 }

// MnfCode returns the manufacturer code (if any) of the command
func (IeeeAddressResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (IeeeAddressResponse) MarshalJSON() ([]byte, error) { return []byte("32769"), nil }

func (v *IeeeAddressResponse) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h IeeeAddressResponseHandler
	if h, found = handler.(IeeeAddressResponseHandler); found {
		err = h.HandleIeeeAddressResponse(frame, v)
	}
	return
}

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

type NodeDescRequestHandler interface {
	HandleNodeDescRequest(frame Frame, cmd *NodeDescRequest) (*NodeDescResponse, error)
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
func (NodeDescRequest) Name() string { return `Node Desc Request` }

// Description of the command
func (NodeDescRequest) Description() string {
	return `queries the node descriptor of a remote device. Should be unicast to the remote device directly,
or to a device that contains the discovery information of the remote device.`
}

// ID of the command
func (NodeDescRequest) ID() CommandID { return NodeDescRequestCommand }

// Required
func (NodeDescRequest) Required() bool { return true }

// Cluster ID of the command
func (NodeDescRequest) Cluster() zcl.ClusterID { return 0x0002 }

// MnfCode returns the manufacturer code (if any) of the command
func (NodeDescRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (NodeDescRequest) MarshalJSON() ([]byte, error) { return []byte("2"), nil }

func (v *NodeDescRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h NodeDescRequestHandler
	if h, found = handler.(NodeDescRequestHandler); found {
		rsp, err = h.HandleNodeDescRequest(frame, v)
	} else {
		rsp = &NodeDescResponse{}
		rsp.(*NodeDescResponse).Status.SetNotSupported()
	}
	return
}

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

type NodeDescResponseHandler interface {
	HandleNodeDescResponse(frame Frame, cmd *NodeDescResponse) error
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
func (NodeDescResponse) Name() string { return `Node Desc Response` }

// Description of the command
func (NodeDescResponse) Description() string { return `` }

// ID of the command
func (NodeDescResponse) ID() CommandID { return NodeDescResponseCommand }

// Required
func (NodeDescResponse) Required() bool { return true }

// Cluster ID of the command
func (NodeDescResponse) Cluster() zcl.ClusterID { return 0x8002 }

// MnfCode returns the manufacturer code (if any) of the command
func (NodeDescResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (NodeDescResponse) MarshalJSON() ([]byte, error) { return []byte("32770"), nil }

func (v *NodeDescResponse) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h NodeDescResponseHandler
	if h, found = handler.(NodeDescResponseHandler); found {
		err = h.HandleNodeDescResponse(frame, v)
	}
	return
}

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

type PowerDescRequestHandler interface {
	HandlePowerDescRequest(frame Frame, cmd *PowerDescRequest) (*PowerDescResponse, error)
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
func (PowerDescRequest) Name() string { return `Power Desc Request` }

// Description of the command
func (PowerDescRequest) Description() string {
	return `queries the power descriptor of a remote device. Should be unicast to the remote device directly,
or to a device that contains the discovery information of the remote device.`
}

// ID of the command
func (PowerDescRequest) ID() CommandID { return PowerDescRequestCommand }

// Required
func (PowerDescRequest) Required() bool { return true }

// Cluster ID of the command
func (PowerDescRequest) Cluster() zcl.ClusterID { return 0x0003 }

// MnfCode returns the manufacturer code (if any) of the command
func (PowerDescRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (PowerDescRequest) MarshalJSON() ([]byte, error) { return []byte("3"), nil }

func (v *PowerDescRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h PowerDescRequestHandler
	if h, found = handler.(PowerDescRequestHandler); found {
		rsp, err = h.HandlePowerDescRequest(frame, v)
	} else {
		rsp = &PowerDescResponse{}
		rsp.(*PowerDescResponse).Status.SetNotSupported()
	}
	return
}

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
	// Status Code, command is normally empty unless status is `Success`
	Status Status
	// NwkAddress is a 16-bit Network address
	NwkAddress      NwkAddress
	PowerDescriptor PowerDescriptor
}

type PowerDescResponseHandler interface {
	HandlePowerDescResponse(frame Frame, cmd *PowerDescResponse) error
}

// PowerDescResponseCommand is the Command ID of PowerDescResponse
const PowerDescResponseCommand CommandID = 0x8003

// Values returns all values of PowerDescResponse
func (v *PowerDescResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.NwkAddress,
		&v.PowerDescriptor,
	}
}

// Arguments returns all values of PowerDescResponse
func (v *PowerDescResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "PowerDescriptor", Argument: &v.PowerDescriptor},
	}
}

// Name of the command
func (PowerDescResponse) Name() string { return `Power Desc Response` }

// Description of the command
func (PowerDescResponse) Description() string { return `` }

// ID of the command
func (PowerDescResponse) ID() CommandID { return PowerDescResponseCommand }

// Required
func (PowerDescResponse) Required() bool { return true }

// Cluster ID of the command
func (PowerDescResponse) Cluster() zcl.ClusterID { return 0x8003 }

// MnfCode returns the manufacturer code (if any) of the command
func (PowerDescResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (PowerDescResponse) MarshalJSON() ([]byte, error) { return []byte("32771"), nil }

func (v *PowerDescResponse) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h PowerDescResponseHandler
	if h, found = handler.(PowerDescResponseHandler); found {
		err = h.HandlePowerDescResponse(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of PowerDescResponse
func (v PowerDescResponse) MarshalZcl() ([]byte, error) {
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
		if tmp, err = v.PowerDescriptor.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the PowerDescResponse struct
func (v *PowerDescResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.PowerDescriptor).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v PowerDescResponse) String() string {
	return zcl.Sprintf(
		"PowerDescResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"NwkAddress(%v)",
			"PowerDescriptor(%v)",
		}, " ")+"}",
		v.Status,
		v.NwkAddress,
		v.PowerDescriptor,
	)
}

// SimpleDescRequest queries the Simple Descriptor of a remote device on a specific endpoint.
// Simple Descriptor contains information about which clusters the device supports on the given endpoint.
// Should be unicast to the remote device directly, or to a device that contains the discovery information
// of the remote device.
type SimpleDescRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	Endpoint   Endpoint
}

type SimpleDescRequestHandler interface {
	HandleSimpleDescRequest(frame Frame, cmd *SimpleDescRequest) (*SimpleDescResponse, error)
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
func (SimpleDescRequest) Name() string { return `Simple Desc Request` }

// Description of the command
func (SimpleDescRequest) Description() string {
	return `queries the Simple Descriptor of a remote device on a specific endpoint.
Simple Descriptor contains information about which clusters the device supports on the given endpoint.
Should be unicast to the remote device directly, or to a device that contains the discovery information
of the remote device.`
}

// ID of the command
func (SimpleDescRequest) ID() CommandID { return SimpleDescRequestCommand }

// Required
func (SimpleDescRequest) Required() bool { return true }

// Cluster ID of the command
func (SimpleDescRequest) Cluster() zcl.ClusterID { return 0x0004 }

// MnfCode returns the manufacturer code (if any) of the command
func (SimpleDescRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (SimpleDescRequest) MarshalJSON() ([]byte, error) { return []byte("4"), nil }

func (v *SimpleDescRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h SimpleDescRequestHandler
	if h, found = handler.(SimpleDescRequestHandler); found {
		rsp, err = h.HandleSimpleDescRequest(frame, v)
	} else {
		rsp = &SimpleDescResponse{}
		rsp.(*SimpleDescResponse).Status.SetNotSupported()
	}
	return
}

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

type SimpleDescResponseHandler interface {
	HandleSimpleDescResponse(frame Frame, cmd *SimpleDescResponse) error
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
func (SimpleDescResponse) Name() string { return `Simple Desc Response` }

// Description of the command
func (SimpleDescResponse) Description() string { return `` }

// ID of the command
func (SimpleDescResponse) ID() CommandID { return SimpleDescResponseCommand }

// Required
func (SimpleDescResponse) Required() bool { return false }

// Cluster ID of the command
func (SimpleDescResponse) Cluster() zcl.ClusterID { return 0x8004 }

// MnfCode returns the manufacturer code (if any) of the command
func (SimpleDescResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (SimpleDescResponse) MarshalJSON() ([]byte, error) { return []byte("32772"), nil }

func (v *SimpleDescResponse) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h SimpleDescResponseHandler
	if h, found = handler.(SimpleDescResponseHandler); found {
		err = h.HandleSimpleDescResponse(frame, v)
	}
	return
}

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

type ActiveEndpointRequestHandler interface {
	HandleActiveEndpointRequest(frame Frame, cmd *ActiveEndpointRequest) (*ActiveEndpointResponse, error)
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
func (ActiveEndpointRequest) Name() string { return `Active Endpoint Request` }

// Description of the command
func (ActiveEndpointRequest) Description() string {
	return `queries the remote device for a list of active endpoints. Should be unicast to the remote device directly,
or to a device that contains the discovery information of the remote device.`
}

// ID of the command
func (ActiveEndpointRequest) ID() CommandID { return ActiveEndpointRequestCommand }

// Required
func (ActiveEndpointRequest) Required() bool { return true }

// Cluster ID of the command
func (ActiveEndpointRequest) Cluster() zcl.ClusterID { return 0x0005 }

// MnfCode returns the manufacturer code (if any) of the command
func (ActiveEndpointRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ActiveEndpointRequest) MarshalJSON() ([]byte, error) { return []byte("5"), nil }

func (v *ActiveEndpointRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h ActiveEndpointRequestHandler
	if h, found = handler.(ActiveEndpointRequestHandler); found {
		rsp, err = h.HandleActiveEndpointRequest(frame, v)
	} else {
		rsp = &ActiveEndpointResponse{}
		rsp.(*ActiveEndpointResponse).Status.SetNotSupported()
	}
	return
}

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

type ActiveEndpointResponseHandler interface {
	HandleActiveEndpointResponse(frame Frame, cmd *ActiveEndpointResponse) error
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
func (ActiveEndpointResponse) Name() string { return `Active Endpoint Response` }

// Description of the command
func (ActiveEndpointResponse) Description() string { return `` }

// ID of the command
func (ActiveEndpointResponse) ID() CommandID { return ActiveEndpointResponseCommand }

// Required
func (ActiveEndpointResponse) Required() bool { return true }

// Cluster ID of the command
func (ActiveEndpointResponse) Cluster() zcl.ClusterID { return 0x8005 }

// MnfCode returns the manufacturer code (if any) of the command
func (ActiveEndpointResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ActiveEndpointResponse) MarshalJSON() ([]byte, error) { return []byte("32773"), nil }

func (v *ActiveEndpointResponse) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h ActiveEndpointResponseHandler
	if h, found = handler.(ActiveEndpointResponseHandler); found {
		err = h.HandleActiveEndpointResponse(frame, v)
	}
	return
}

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

type MatchDescRequestHandler interface {
	HandleMatchDescRequest(frame Frame, cmd *MatchDescRequest) error
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
func (MatchDescRequest) Name() string { return `Match Desc Request` }

// Description of the command
func (MatchDescRequest) Description() string {
	return `is used to find remote devices supporting a specific simple descriptor match criterion. Normally broadcast
to all devices supporting RX on when Idle (0xFFFD)`
}

// ID of the command
func (MatchDescRequest) ID() CommandID { return MatchDescRequestCommand }

// Required
func (MatchDescRequest) Required() bool { return true }

// Cluster ID of the command
func (MatchDescRequest) Cluster() zcl.ClusterID { return 0x0006 }

// MnfCode returns the manufacturer code (if any) of the command
func (MatchDescRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MatchDescRequest) MarshalJSON() ([]byte, error) { return []byte("6"), nil }

func (v *MatchDescRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h MatchDescRequestHandler
	if h, found = handler.(MatchDescRequestHandler); found {
		err = h.HandleMatchDescRequest(frame, v)
	}
	return
}

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
	Status Status
	// NwkAddress is a 16-bit Network address
	NwkAddress   NwkAddress
	EndpointList EndpointList
}

type MatchDescResponseHandler interface {
	HandleMatchDescResponse(frame Frame, cmd *MatchDescResponse) error
}

// MatchDescResponseCommand is the Command ID of MatchDescResponse
const MatchDescResponseCommand CommandID = 0x8006

// Values returns all values of MatchDescResponse
func (v *MatchDescResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.NwkAddress,
		&v.EndpointList,
	}
}

// Arguments returns all values of MatchDescResponse
func (v *MatchDescResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "EndpointList", Argument: &v.EndpointList},
	}
}

// Name of the command
func (MatchDescResponse) Name() string { return `Match Desc Response` }

// Description of the command
func (MatchDescResponse) Description() string { return `` }

// ID of the command
func (MatchDescResponse) ID() CommandID { return MatchDescResponseCommand }

// Required
func (MatchDescResponse) Required() bool { return false }

// Cluster ID of the command
func (MatchDescResponse) Cluster() zcl.ClusterID { return 0x8006 }

// MnfCode returns the manufacturer code (if any) of the command
func (MatchDescResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MatchDescResponse) MarshalJSON() ([]byte, error) { return []byte("32774"), nil }

func (v *MatchDescResponse) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h MatchDescResponseHandler
	if h, found = handler.(MatchDescResponseHandler); found {
		err = h.HandleMatchDescResponse(frame, v)
	}
	return
}

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

// UnmarshalZcl parses the wire format representation into the MatchDescResponse struct
func (v *MatchDescResponse) UnmarshalZcl(b []byte) ([]byte, error) {
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
func (v MatchDescResponse) String() string {
	return zcl.Sprintf(
		"MatchDescResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"NwkAddress(%v)",
			"EndpointList(%v)",
		}, " ")+"}",
		v.Status,
		v.NwkAddress,
		v.EndpointList,
	)
}

// ComplexDescRequest queries the Complex Descriptor of a remote device. Should be unicast to the remote device directly,
// or to a device that contains the discovery information of the remote device.
type ComplexDescRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
}

type ComplexDescRequestHandler interface {
	HandleComplexDescRequest(frame Frame, cmd *ComplexDescRequest) error
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
func (ComplexDescRequest) Name() string { return `Complex Desc Request` }

// Description of the command
func (ComplexDescRequest) Description() string {
	return `queries the Complex Descriptor of a remote device. Should be unicast to the remote device directly,
or to a device that contains the discovery information of the remote device.`
}

// ID of the command
func (ComplexDescRequest) ID() CommandID { return ComplexDescRequestCommand }

// Required
func (ComplexDescRequest) Required() bool { return false }

// Cluster ID of the command
func (ComplexDescRequest) Cluster() zcl.ClusterID { return 0x0010 }

// MnfCode returns the manufacturer code (if any) of the command
func (ComplexDescRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ComplexDescRequest) MarshalJSON() ([]byte, error) { return []byte("16"), nil }

func (v *ComplexDescRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h ComplexDescRequestHandler
	if h, found = handler.(ComplexDescRequestHandler); found {
		err = h.HandleComplexDescRequest(frame, v)
	}
	return
}

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

type ComplexDescResponseHandler interface {
	HandleComplexDescResponse(frame Frame, cmd *ComplexDescResponse) error
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
func (ComplexDescResponse) Name() string { return `Complex Desc Response` }

// Description of the command
func (ComplexDescResponse) Description() string { return `` }

// ID of the command
func (ComplexDescResponse) ID() CommandID { return ComplexDescResponseCommand }

// Required
func (ComplexDescResponse) Required() bool { return false }

// Cluster ID of the command
func (ComplexDescResponse) Cluster() zcl.ClusterID { return 0x8010 }

// MnfCode returns the manufacturer code (if any) of the command
func (ComplexDescResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ComplexDescResponse) MarshalJSON() ([]byte, error) { return []byte("32784"), nil }

func (v *ComplexDescResponse) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h ComplexDescResponseHandler
	if h, found = handler.(ComplexDescResponseHandler); found {
		err = h.HandleComplexDescResponse(frame, v)
	}
	return
}

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

type UserDescRequestHandler interface {
	HandleUserDescRequest(frame Frame, cmd *UserDescRequest) error
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
func (UserDescRequest) Name() string { return `User Desc Request` }

// Description of the command
func (UserDescRequest) Description() string {
	return `queries the User Descriptor of a remote device. Should be unicast to the remote device directly,
or to a device that contains the discovery information of the remote device.`
}

// ID of the command
func (UserDescRequest) ID() CommandID { return UserDescRequestCommand }

// Required
func (UserDescRequest) Required() bool { return false }

// Cluster ID of the command
func (UserDescRequest) Cluster() zcl.ClusterID { return 0x0011 }

// MnfCode returns the manufacturer code (if any) of the command
func (UserDescRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (UserDescRequest) MarshalJSON() ([]byte, error) { return []byte("17"), nil }

func (v *UserDescRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h UserDescRequestHandler
	if h, found = handler.(UserDescRequestHandler); found {
		err = h.HandleUserDescRequest(frame, v)
	}
	return
}

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

type UserDescResponse struct {
	// Status Code, command is normally empty unless status is `Success`
	Status Status
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	// UserDescriptor is a user provided ASCII string of 16 characters set on a remote device.
	// If the string is shorter than 16 characters it should be padded with space characters (0x20).
	// Control characters (0x00-0x1f) are not allowed.
	UserDescriptor UserDescriptor
}

type UserDescResponseHandler interface {
	HandleUserDescResponse(frame Frame, cmd *UserDescResponse) error
}

// UserDescResponseCommand is the Command ID of UserDescResponse
const UserDescResponseCommand CommandID = 0x8011

// Values returns all values of UserDescResponse
func (v *UserDescResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.NwkAddress,
		&v.UserDescriptor,
	}
}

// Arguments returns all values of UserDescResponse
func (v *UserDescResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "UserDescriptor", Argument: &v.UserDescriptor},
	}
}

// Name of the command
func (UserDescResponse) Name() string { return `User Desc Response` }

// Description of the command
func (UserDescResponse) Description() string { return `` }

// ID of the command
func (UserDescResponse) ID() CommandID { return UserDescResponseCommand }

// Required
func (UserDescResponse) Required() bool { return false }

// Cluster ID of the command
func (UserDescResponse) Cluster() zcl.ClusterID { return 0x8011 }

// MnfCode returns the manufacturer code (if any) of the command
func (UserDescResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (UserDescResponse) MarshalJSON() ([]byte, error) { return []byte("32785"), nil }

func (v *UserDescResponse) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h UserDescResponseHandler
	if h, found = handler.(UserDescResponseHandler); found {
		err = h.HandleUserDescResponse(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of UserDescResponse
func (v UserDescResponse) MarshalZcl() ([]byte, error) {
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
	// UserDescriptor is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.UserDescriptor.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the UserDescResponse struct
func (v *UserDescResponse) UnmarshalZcl(b []byte) ([]byte, error) {
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

	// UserDescriptor is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.UserDescriptor).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v UserDescResponse) String() string {
	return zcl.Sprintf(
		"UserDescResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"NwkAddress(%v)",
			"UserDescriptor(%v)",
		}, " ")+"}",
		v.Status,
		v.NwkAddress,
		v.UserDescriptor,
	)
}

// DiscoveryCacheRequest locates a Primary Discovery Cache on the network. Should be addressed to broadcast RXOnWhenIdle (0xFFFD)
type DiscoveryCacheRequest struct {
	// NwkAddress is a 16-bit Network address
	NwkAddress NwkAddress
	// IeeeAddress is a 64-bit MAC address
	IeeeAddress IeeeAddress
}

type DiscoveryCacheRequestHandler interface {
	HandleDiscoveryCacheRequest(frame Frame, cmd *DiscoveryCacheRequest) error
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
func (DiscoveryCacheRequest) Name() string { return `Discovery Cache Request` }

// Description of the command
func (DiscoveryCacheRequest) Description() string {
	return `locates a Primary Discovery Cache on the network. Should be addressed to broadcast RXOnWhenIdle (0xFFFD)`
}

// ID of the command
func (DiscoveryCacheRequest) ID() CommandID { return DiscoveryCacheRequestCommand }

// Required
func (DiscoveryCacheRequest) Required() bool { return false }

// Cluster ID of the command
func (DiscoveryCacheRequest) Cluster() zcl.ClusterID { return 0x0012 }

// MnfCode returns the manufacturer code (if any) of the command
func (DiscoveryCacheRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (DiscoveryCacheRequest) MarshalJSON() ([]byte, error) { return []byte("18"), nil }

func (v *DiscoveryCacheRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h DiscoveryCacheRequestHandler
	if h, found = handler.(DiscoveryCacheRequestHandler); found {
		err = h.HandleDiscoveryCacheRequest(frame, v)
	}
	return
}

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

type DeviceAnnounceHandler interface {
	HandleDeviceAnnounce(frame Frame, cmd *DeviceAnnounce) error
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
func (DeviceAnnounce) Name() string { return `Device Announce` }

// Description of the command
func (DeviceAnnounce) Description() string {
	return `is sent by a joining or returning device, identifying it's NWK address, IEEE address and capabilities.
Normally sent to broadcast RXOnWhenIdle (0xFFFD)`
}

// ID of the command
func (DeviceAnnounce) ID() CommandID { return DeviceAnnounceCommand }

// Required
func (DeviceAnnounce) Required() bool { return true }

// Cluster ID of the command
func (DeviceAnnounce) Cluster() zcl.ClusterID { return 0x0013 }

// MnfCode returns the manufacturer code (if any) of the command
func (DeviceAnnounce) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (DeviceAnnounce) MarshalJSON() ([]byte, error) { return []byte("19"), nil }

func (v *DeviceAnnounce) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h DeviceAnnounceHandler
	if h, found = handler.(DeviceAnnounceHandler); found {
		err = h.HandleDeviceAnnounce(frame, v)
	}
	return
}

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
	// UserDescriptor is a user provided ASCII string of 16 characters set on a remote device.
	// If the string is shorter than 16 characters it should be padded with space characters (0x20).
	// Control characters (0x00-0x1f) are not allowed.
	UserDescriptor UserDescriptor
}

type UserDescSetRequestHandler interface {
	HandleUserDescSetRequest(frame Frame, cmd *UserDescSetRequest) error
}

// UserDescSetRequestCommand is the Command ID of UserDescSetRequest
const UserDescSetRequestCommand CommandID = 0x0014

// Values returns all values of UserDescSetRequest
func (v *UserDescSetRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.NwkAddress,
		&v.UserDescriptor,
	}
}

// Arguments returns all values of UserDescSetRequest
func (v *UserDescSetRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "NwkAddress", Argument: &v.NwkAddress},
		{Name: "UserDescriptor", Argument: &v.UserDescriptor},
	}
}

// Name of the command
func (UserDescSetRequest) Name() string { return `User Desc Set Request` }

// Description of the command
func (UserDescSetRequest) Description() string {
	return `requests the remote device to update it's user descriptor.`
}

// ID of the command
func (UserDescSetRequest) ID() CommandID { return UserDescSetRequestCommand }

// Required
func (UserDescSetRequest) Required() bool { return false }

// Cluster ID of the command
func (UserDescSetRequest) Cluster() zcl.ClusterID { return 0x0014 }

// MnfCode returns the manufacturer code (if any) of the command
func (UserDescSetRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (UserDescSetRequest) MarshalJSON() ([]byte, error) { return []byte("20"), nil }

func (v *UserDescSetRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h UserDescSetRequestHandler
	if h, found = handler.(UserDescSetRequestHandler); found {
		err = h.HandleUserDescSetRequest(frame, v)
	}
	return
}

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
		if tmp, err = v.UserDescriptor.MarshalZcl(); err != nil {
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

	if b, err = (&v.UserDescriptor).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v UserDescSetRequest) String() string {
	return zcl.Sprintf(
		"UserDescSetRequest{"+zcl.StrJoin([]string{
			"NwkAddress(%v)",
			"UserDescriptor(%v)",
		}, " ")+"}",
		v.NwkAddress,
		v.UserDescriptor,
	)
}

// SystemServerDiscoverRequest discovers the location of a particular system server or servers as indicated by the Server Mask. Should be sent to broadcast RXOnWhenIdle (0xFFFD)
type SystemServerDiscoverRequest struct {
	ServerMask ServerMask
}

type SystemServerDiscoverRequestHandler interface {
	HandleSystemServerDiscoverRequest(frame Frame, cmd *SystemServerDiscoverRequest) error
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
func (SystemServerDiscoverRequest) Name() string { return `System Server Discover Request` }

// Description of the command
func (SystemServerDiscoverRequest) Description() string {
	return `discovers the location of a particular system server or servers as indicated by the Server Mask. Should be sent to broadcast RXOnWhenIdle (0xFFFD)`
}

// ID of the command
func (SystemServerDiscoverRequest) ID() CommandID { return SystemServerDiscoverRequestCommand }

// Required
func (SystemServerDiscoverRequest) Required() bool { return false }

// Cluster ID of the command
func (SystemServerDiscoverRequest) Cluster() zcl.ClusterID { return 0x0015 }

// MnfCode returns the manufacturer code (if any) of the command
func (SystemServerDiscoverRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (SystemServerDiscoverRequest) MarshalJSON() ([]byte, error) { return []byte("21"), nil }

func (v *SystemServerDiscoverRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h SystemServerDiscoverRequestHandler
	if h, found = handler.(SystemServerDiscoverRequestHandler); found {
		err = h.HandleSystemServerDiscoverRequest(frame, v)
	}
	return
}

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

type DiscoveryStoreRequestHandler interface {
	HandleDiscoveryStoreRequest(frame Frame, cmd *DiscoveryStoreRequest) error
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
func (DiscoveryStoreRequest) Name() string { return `Discovery Store Request` }

// Description of the command
func (DiscoveryStoreRequest) Description() string {
	return `sent to a Discovery Cache Node to allocate memory of the provided sizes for cache storage.
If the node already exists in cache, it will be removed to allow storage of updated values.
Should be sent to the unicast address of a discovery cache device.`
}

// ID of the command
func (DiscoveryStoreRequest) ID() CommandID { return DiscoveryStoreRequestCommand }

// Required
func (DiscoveryStoreRequest) Required() bool { return false }

// Cluster ID of the command
func (DiscoveryStoreRequest) Cluster() zcl.ClusterID { return 0x0016 }

// MnfCode returns the manufacturer code (if any) of the command
func (DiscoveryStoreRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (DiscoveryStoreRequest) MarshalJSON() ([]byte, error) { return []byte("22"), nil }

func (v *DiscoveryStoreRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h DiscoveryStoreRequestHandler
	if h, found = handler.(DiscoveryStoreRequestHandler); found {
		err = h.HandleDiscoveryStoreRequest(frame, v)
	}
	return
}

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
	// This is the maximum size of data or commands passed to or from the application by the application support
	// sub-layer, before any fragmentation or re-assembly.
	MaxBufferSize MaxBufferSize
	// MaxRxSize specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred to this node in one single message transfer.
	MaxRxSize  MaxRxSize
	ServerMask ServerMask
	// MaxTxSize specifies the maximum size, in octets, of the application sub-layer data unit (ASDU) that can be transferred from
	// this node in one single message transfer.
	MaxTxSize MaxTxSize
}

type NodeDescStoreRequestHandler interface {
	HandleNodeDescStoreRequest(frame Frame, cmd *NodeDescStoreRequest) error
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
func (NodeDescStoreRequest) Name() string { return `Node Desc Store Request` }

// Description of the command
func (NodeDescStoreRequest) Description() string {
	return `sent to a Discovery Cache Node to store the provided Node Descriptor.
A DiscoveryStoreRequest should be sent and acknowledged with successful status before requesting storage.
Should be sent to the unicast address of a discovery cache device.`
}

// ID of the command
func (NodeDescStoreRequest) ID() CommandID { return NodeDescStoreRequestCommand }

// Required
func (NodeDescStoreRequest) Required() bool { return false }

// Cluster ID of the command
func (NodeDescStoreRequest) Cluster() zcl.ClusterID { return 0x0017 }

// MnfCode returns the manufacturer code (if any) of the command
func (NodeDescStoreRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (NodeDescStoreRequest) MarshalJSON() ([]byte, error) { return []byte("23"), nil }

func (v *NodeDescStoreRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h NodeDescStoreRequestHandler
	if h, found = handler.(NodeDescStoreRequestHandler); found {
		err = h.HandleNodeDescStoreRequest(frame, v)
	}
	return
}

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

type PowerDescStoreRequestHandler interface {
	HandlePowerDescStoreRequest(frame Frame, cmd *PowerDescStoreRequest) error
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
func (PowerDescStoreRequest) Name() string { return `Power Desc Store Request` }

// Description of the command
func (PowerDescStoreRequest) Description() string {
	return `sent to a Discovery Cache Node to store the provided Power Descriptor.
A DiscoveryStoreRequest should be sent and acknowledged with successful status before requesting storage.
Should be sent to the unicast address of a discovery cache device.`
}

// ID of the command
func (PowerDescStoreRequest) ID() CommandID { return PowerDescStoreRequestCommand }

// Required
func (PowerDescStoreRequest) Required() bool { return false }

// Cluster ID of the command
func (PowerDescStoreRequest) Cluster() zcl.ClusterID { return 0x0018 }

// MnfCode returns the manufacturer code (if any) of the command
func (PowerDescStoreRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (PowerDescStoreRequest) MarshalJSON() ([]byte, error) { return []byte("24"), nil }

func (v *PowerDescStoreRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h PowerDescStoreRequestHandler
	if h, found = handler.(PowerDescStoreRequestHandler); found {
		err = h.HandlePowerDescStoreRequest(frame, v)
	}
	return
}

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

type ActiveEndpointStoreRequestHandler interface {
	HandleActiveEndpointStoreRequest(frame Frame, cmd *ActiveEndpointStoreRequest) error
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
func (ActiveEndpointStoreRequest) Name() string { return `Active Endpoint Store Request` }

// Description of the command
func (ActiveEndpointStoreRequest) Description() string {
	return `sent to a Discovery Cache Node to store the provided Active Endpoint list.
A DiscoveryStoreRequest should be sent and acknowledged with successful status before requesting storage.
Should be sent to the unicast address of a discovery cache device.`
}

// ID of the command
func (ActiveEndpointStoreRequest) ID() CommandID { return ActiveEndpointStoreRequestCommand }

// Required
func (ActiveEndpointStoreRequest) Required() bool { return false }

// Cluster ID of the command
func (ActiveEndpointStoreRequest) Cluster() zcl.ClusterID { return 0x0019 }

// MnfCode returns the manufacturer code (if any) of the command
func (ActiveEndpointStoreRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ActiveEndpointStoreRequest) MarshalJSON() ([]byte, error) { return []byte("25"), nil }

func (v *ActiveEndpointStoreRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h ActiveEndpointStoreRequestHandler
	if h, found = handler.(ActiveEndpointStoreRequestHandler); found {
		err = h.HandleActiveEndpointStoreRequest(frame, v)
	}
	return
}

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

type SimpleDescStoreRequestHandler interface {
	HandleSimpleDescStoreRequest(frame Frame, cmd *SimpleDescStoreRequest) error
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
func (SimpleDescStoreRequest) Name() string { return `Simple Desc Store Request` }

// Description of the command
func (SimpleDescStoreRequest) Description() string {
	return `sent to a Discovery Cache Node to store the provided Simple Descriptor.
A DiscoveryStoreRequest should be sent and acknowledged with successful status before requesting storage.
Should be sent to the unicast address of a discovery cache device.`
}

// ID of the command
func (SimpleDescStoreRequest) ID() CommandID { return SimpleDescStoreRequestCommand }

// Required
func (SimpleDescStoreRequest) Required() bool { return false }

// Cluster ID of the command
func (SimpleDescStoreRequest) Cluster() zcl.ClusterID { return 0x001A }

// MnfCode returns the manufacturer code (if any) of the command
func (SimpleDescStoreRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (SimpleDescStoreRequest) MarshalJSON() ([]byte, error) { return []byte("26"), nil }

func (v *SimpleDescStoreRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h SimpleDescStoreRequestHandler
	if h, found = handler.(SimpleDescStoreRequestHandler); found {
		err = h.HandleSimpleDescStoreRequest(frame, v)
	}
	return
}

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

type RemoveNodeCacheRequestHandler interface {
	HandleRemoveNodeCacheRequest(frame Frame, cmd *RemoveNodeCacheRequest) error
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
func (RemoveNodeCacheRequest) Name() string { return `Remove Node Cache Request` }

// Description of the command
func (RemoveNodeCacheRequest) Description() string {
	return `sent to a Discovery Cache Node will request it to remove cached values for the provided node. Should be sent to the unicast address of a discovery cache device.`
}

// ID of the command
func (RemoveNodeCacheRequest) ID() CommandID { return RemoveNodeCacheRequestCommand }

// Required
func (RemoveNodeCacheRequest) Required() bool { return false }

// Cluster ID of the command
func (RemoveNodeCacheRequest) Cluster() zcl.ClusterID { return 0x001B }

// MnfCode returns the manufacturer code (if any) of the command
func (RemoveNodeCacheRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (RemoveNodeCacheRequest) MarshalJSON() ([]byte, error) { return []byte("27"), nil }

func (v *RemoveNodeCacheRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h RemoveNodeCacheRequestHandler
	if h, found = handler.(RemoveNodeCacheRequestHandler); found {
		err = h.HandleRemoveNodeCacheRequest(frame, v)
	}
	return
}

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

type FindNodeCacheRequestHandler interface {
	HandleFindNodeCacheRequest(frame Frame, cmd *FindNodeCacheRequest) error
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
func (FindNodeCacheRequest) Name() string { return `Find Node Cache Request` }

// Description of the command
func (FindNodeCacheRequest) Description() string {
	return `broadcast to the network will generate a response from the Primary Discovery Cache holding information
for the device of interest`
}

// ID of the command
func (FindNodeCacheRequest) ID() CommandID { return FindNodeCacheRequestCommand }

// Required
func (FindNodeCacheRequest) Required() bool { return false }

// Cluster ID of the command
func (FindNodeCacheRequest) Cluster() zcl.ClusterID { return 0x001C }

// MnfCode returns the manufacturer code (if any) of the command
func (FindNodeCacheRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (FindNodeCacheRequest) MarshalJSON() ([]byte, error) { return []byte("28"), nil }

func (v *FindNodeCacheRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h FindNodeCacheRequestHandler
	if h, found = handler.(FindNodeCacheRequestHandler); found {
		err = h.HandleFindNodeCacheRequest(frame, v)
	}
	return
}

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

type ExtendedSimpleDescRequestHandler interface {
	HandleExtendedSimpleDescRequest(frame Frame, cmd *ExtendedSimpleDescRequest) error
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
func (ExtendedSimpleDescRequest) Name() string { return `Extended Simple Desc Request` }

// Description of the command
func (ExtendedSimpleDescRequest) Description() string {
	return `should be unicast to the remote device or a discovery cache node. It is used to request information from
nodes that implements a larger number of clusters than can be described by a SimpleDescRequest`
}

// ID of the command
func (ExtendedSimpleDescRequest) ID() CommandID { return ExtendedSimpleDescRequestCommand }

// Required
func (ExtendedSimpleDescRequest) Required() bool { return false }

// Cluster ID of the command
func (ExtendedSimpleDescRequest) Cluster() zcl.ClusterID { return 0x001D }

// MnfCode returns the manufacturer code (if any) of the command
func (ExtendedSimpleDescRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ExtendedSimpleDescRequest) MarshalJSON() ([]byte, error) { return []byte("29"), nil }

func (v *ExtendedSimpleDescRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h ExtendedSimpleDescRequestHandler
	if h, found = handler.(ExtendedSimpleDescRequestHandler); found {
		err = h.HandleExtendedSimpleDescRequest(frame, v)
	}
	return
}

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

type ExtendedActiveEndpointRequestHandler interface {
	HandleExtendedActiveEndpointRequest(frame Frame, cmd *ExtendedActiveEndpointRequest) error
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
func (ExtendedActiveEndpointRequest) Name() string { return `Extended Active Endpoint Request` }

// Description of the command
func (ExtendedActiveEndpointRequest) Description() string {
	return `should be unicast to the remote device or a discovery cache node. It is used to request information from
nodes that implements a larger number of endpoints than can be described by a ActiveEndpointRequest`
}

// ID of the command
func (ExtendedActiveEndpointRequest) ID() CommandID { return ExtendedActiveEndpointRequestCommand }

// Required
func (ExtendedActiveEndpointRequest) Required() bool { return false }

// Cluster ID of the command
func (ExtendedActiveEndpointRequest) Cluster() zcl.ClusterID { return 0x001E }

// MnfCode returns the manufacturer code (if any) of the command
func (ExtendedActiveEndpointRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ExtendedActiveEndpointRequest) MarshalJSON() ([]byte, error) { return []byte("30"), nil }

func (v *ExtendedActiveEndpointRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h ExtendedActiveEndpointRequestHandler
	if h, found = handler.(ExtendedActiveEndpointRequestHandler); found {
		err = h.HandleExtendedActiveEndpointRequest(frame, v)
	}
	return
}

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
// * If the ProfileID doesn't match or none of the In/OutClusterList elements match - a reply is sent with status
// `NoMatch`
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

type EndDeviceBindRequestHandler interface {
	HandleEndDeviceBindRequest(frame Frame, cmd *EndDeviceBindRequest) error
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
func (EndDeviceBindRequest) Name() string { return `End Device Bind Request` }

// Description of the command
func (EndDeviceBindRequest) Description() string {
	return "is sent to the ZigBee coordinator from two different devices simultaneously." + "\n" +
		"* If the supplied endpoint is outside the specified range - a reply is sent with status `InvalidEndpoint`" + "\n" +
		"* If only one device sends the request within a pre-configure timeout - a reply is sent with status `Timeout`" + "\n" +
		"* If the ProfileID doesn't match or none of the In/OutClusterList elements match - a reply is sent with status" + "\n" +
		"  `NoMatch`" + "\n" +
		"* Otherwise, a reply is sent with status `Success` to each device" + "\n" +
		"" + "\n" +
		"The Coordinator then needs to either send `BindRequest` or `UnbindRequest` for each matched `ClusterID`." + "\n" +
		"This is done by first issuing a `UnbindRequest` with any of the matched `ClusterID`:" + "\n" +
		"* If the reply status is `NoEntry` - `BindRequest` will instead be sent for each matched `ClusterID`" + "\n" +
		"* If the reply status is `Success` - additional unbind requests are sent the rest of the matched cluster" + "\n" +
		"" + "\n" +
		"This enables the request to toggle the binding."
}

// ID of the command
func (EndDeviceBindRequest) ID() CommandID { return EndDeviceBindRequestCommand }

// Required
func (EndDeviceBindRequest) Required() bool { return false }

// Cluster ID of the command
func (EndDeviceBindRequest) Cluster() zcl.ClusterID { return 0x0020 }

// MnfCode returns the manufacturer code (if any) of the command
func (EndDeviceBindRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (EndDeviceBindRequest) MarshalJSON() ([]byte, error) { return []byte("32"), nil }

func (v *EndDeviceBindRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h EndDeviceBindRequestHandler
	if h, found = handler.(EndDeviceBindRequestHandler); found {
		err = h.HandleEndDeviceBindRequest(frame, v)
	}
	return
}

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
	BindingTableEntry BindingTableEntry
}

type BindRequestHandler interface {
	HandleBindRequest(frame Frame, cmd *BindRequest) (*MgmtBindResponse, error)
}

// BindRequestCommand is the Command ID of BindRequest
const BindRequestCommand CommandID = 0x0021

// Values returns all values of BindRequest
func (v *BindRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.BindingTableEntry,
	}
}

// Arguments returns all values of BindRequest
func (v *BindRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "BindingTableEntry", Argument: &v.BindingTableEntry},
	}
}

// Name of the command
func (BindRequest) Name() string { return `Bind Request` }

// Description of the command
func (BindRequest) Description() string { return `` }

// ID of the command
func (BindRequest) ID() CommandID { return BindRequestCommand }

// Required
func (BindRequest) Required() bool { return false }

// Cluster ID of the command
func (BindRequest) Cluster() zcl.ClusterID { return 0x0021 }

// MnfCode returns the manufacturer code (if any) of the command
func (BindRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (BindRequest) MarshalJSON() ([]byte, error) { return []byte("33"), nil }

func (v *BindRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h BindRequestHandler
	if h, found = handler.(BindRequestHandler); found {
		rsp, err = h.HandleBindRequest(frame, v)
	} else {
		rsp = &MgmtBindResponse{}
		rsp.(*MgmtBindResponse).Status.SetNotSupported()
	}
	return
}

// MarshalZcl returns the wire format representation of BindRequest
func (v BindRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.BindingTableEntry.MarshalZcl(); err != nil {
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

	if b, err = (&v.BindingTableEntry).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v BindRequest) String() string {
	return zcl.Sprintf(
		"BindRequest{"+zcl.StrJoin([]string{
			"BindingTableEntry(%v)",
		}, " ")+"}",
		v.BindingTableEntry,
	)
}

type BindResponse struct {
	// Status Code, command is normally empty unless status is `Success`
	Status Status
}

type BindResponseHandler interface {
	HandleBindResponse(frame Frame, cmd *BindResponse) error
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
func (BindResponse) Name() string { return `Bind Response` }

// Description of the command
func (BindResponse) Description() string { return `` }

// ID of the command
func (BindResponse) ID() CommandID { return BindResponseCommand }

// Required
func (BindResponse) Required() bool { return false }

// Cluster ID of the command
func (BindResponse) Cluster() zcl.ClusterID { return 0x8021 }

// MnfCode returns the manufacturer code (if any) of the command
func (BindResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (BindResponse) MarshalJSON() ([]byte, error) { return []byte("32801"), nil }

func (v *BindResponse) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h BindResponseHandler
	if h, found = handler.(BindResponseHandler); found {
		err = h.HandleBindResponse(frame, v)
	}
	return
}

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

type MgmtBindRequest struct {
	// StartIndex provides the starting index for the requested elements of the associated list.
	StartIndex StartIndex
}

type MgmtBindRequestHandler interface {
	HandleMgmtBindRequest(frame Frame, cmd *MgmtBindRequest) (*MgmtBindResponse, error)
}

// MgmtBindRequestCommand is the Command ID of MgmtBindRequest
const MgmtBindRequestCommand CommandID = 0x0033

// Values returns all values of MgmtBindRequest
func (v *MgmtBindRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.StartIndex,
	}
}

// Arguments returns all values of MgmtBindRequest
func (v *MgmtBindRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "StartIndex", Argument: &v.StartIndex},
	}
}

// Name of the command
func (MgmtBindRequest) Name() string { return `Mgmt Bind Request` }

// Description of the command
func (MgmtBindRequest) Description() string { return `` }

// ID of the command
func (MgmtBindRequest) ID() CommandID { return MgmtBindRequestCommand }

// Required
func (MgmtBindRequest) Required() bool { return false }

// Cluster ID of the command
func (MgmtBindRequest) Cluster() zcl.ClusterID { return 0x0033 }

// MnfCode returns the manufacturer code (if any) of the command
func (MgmtBindRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MgmtBindRequest) MarshalJSON() ([]byte, error) { return []byte("51"), nil }

func (v *MgmtBindRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h MgmtBindRequestHandler
	if h, found = handler.(MgmtBindRequestHandler); found {
		rsp, err = h.HandleMgmtBindRequest(frame, v)
	} else {
		rsp = &MgmtBindResponse{}
		rsp.(*MgmtBindResponse).Status.SetNotSupported()
	}
	return
}

// MarshalZcl returns the wire format representation of MgmtBindRequest
func (v MgmtBindRequest) MarshalZcl() ([]byte, error) {
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

// UnmarshalZcl parses the wire format representation into the MgmtBindRequest struct
func (v *MgmtBindRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.StartIndex).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MgmtBindRequest) String() string {
	return zcl.Sprintf(
		"MgmtBindRequest{"+zcl.StrJoin([]string{
			"StartIndex(%v)",
		}, " ")+"}",
		v.StartIndex,
	)
}

type MgmtBindResponse struct {
	// Status Code, command is normally empty unless status is `Success`
	Status Status
	// TotalEntries is the total number of entries that can be queried for
	TotalEntries TotalEntries
	// StartIndex provides the starting index for the requested elements of the associated list.
	StartIndex   StartIndex
	BindingTable BindingTable
}

type MgmtBindResponseHandler interface {
	HandleMgmtBindResponse(frame Frame, cmd *MgmtBindResponse) error
}

// MgmtBindResponseCommand is the Command ID of MgmtBindResponse
const MgmtBindResponseCommand CommandID = 0x8033

// Values returns all values of MgmtBindResponse
func (v *MgmtBindResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.TotalEntries,
		&v.StartIndex,
		&v.BindingTable,
	}
}

// Arguments returns all values of MgmtBindResponse
func (v *MgmtBindResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "TotalEntries", Argument: &v.TotalEntries},
		{Name: "StartIndex", Argument: &v.StartIndex},
		{Name: "BindingTable", Argument: &v.BindingTable},
	}
}

// Name of the command
func (MgmtBindResponse) Name() string { return `Mgmt Bind Response` }

// Description of the command
func (MgmtBindResponse) Description() string { return `` }

// ID of the command
func (MgmtBindResponse) ID() CommandID { return MgmtBindResponseCommand }

// Required
func (MgmtBindResponse) Required() bool { return false }

// Cluster ID of the command
func (MgmtBindResponse) Cluster() zcl.ClusterID { return 0x8033 }

// MnfCode returns the manufacturer code (if any) of the command
func (MgmtBindResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MgmtBindResponse) MarshalJSON() ([]byte, error) { return []byte("32819"), nil }

func (v *MgmtBindResponse) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h MgmtBindResponseHandler
	if h, found = handler.(MgmtBindResponseHandler); found {
		err = h.HandleMgmtBindResponse(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of MgmtBindResponse
func (v MgmtBindResponse) MarshalZcl() ([]byte, error) {
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
		if tmp, err = v.BindingTable.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the MgmtBindResponse struct
func (v *MgmtBindResponse) UnmarshalZcl(b []byte) ([]byte, error) {
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

	if b, err = (&v.BindingTable).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MgmtBindResponse) String() string {
	return zcl.Sprintf(
		"MgmtBindResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"TotalEntries(%v)",
			"StartIndex(%v)",
			"BindingTable(%v)",
		}, " ")+"}",
		v.Status,
		v.TotalEntries,
		v.StartIndex,
		v.BindingTable,
	)
}

type UnbindRequest struct {
	BindingTableEntry BindingTableEntry
}

type UnbindRequestHandler interface {
	HandleUnbindRequest(frame Frame, cmd *UnbindRequest) error
}

// UnbindRequestCommand is the Command ID of UnbindRequest
const UnbindRequestCommand CommandID = 0x0022

// Values returns all values of UnbindRequest
func (v *UnbindRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.BindingTableEntry,
	}
}

// Arguments returns all values of UnbindRequest
func (v *UnbindRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "BindingTableEntry", Argument: &v.BindingTableEntry},
	}
}

// Name of the command
func (UnbindRequest) Name() string { return `Unbind Request` }

// Description of the command
func (UnbindRequest) Description() string { return `` }

// ID of the command
func (UnbindRequest) ID() CommandID { return UnbindRequestCommand }

// Required
func (UnbindRequest) Required() bool { return false }

// Cluster ID of the command
func (UnbindRequest) Cluster() zcl.ClusterID { return 0x0022 }

// MnfCode returns the manufacturer code (if any) of the command
func (UnbindRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (UnbindRequest) MarshalJSON() ([]byte, error) { return []byte("34"), nil }

func (v *UnbindRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h UnbindRequestHandler
	if h, found = handler.(UnbindRequestHandler); found {
		err = h.HandleUnbindRequest(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of UnbindRequest
func (v UnbindRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.BindingTableEntry.MarshalZcl(); err != nil {
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

	if b, err = (&v.BindingTableEntry).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v UnbindRequest) String() string {
	return zcl.Sprintf(
		"UnbindRequest{"+zcl.StrJoin([]string{
			"BindingTableEntry(%v)",
		}, " ")+"}",
		v.BindingTableEntry,
	)
}

type UnbindResponse struct {
	// Status Code, command is normally empty unless status is `Success`
	Status Status
}

type UnbindResponseHandler interface {
	HandleUnbindResponse(frame Frame, cmd *UnbindResponse) error
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
func (UnbindResponse) Name() string { return `Unbind Response` }

// Description of the command
func (UnbindResponse) Description() string { return `` }

// ID of the command
func (UnbindResponse) ID() CommandID { return UnbindResponseCommand }

// Required
func (UnbindResponse) Required() bool { return false }

// Cluster ID of the command
func (UnbindResponse) Cluster() zcl.ClusterID { return 0x8022 }

// MnfCode returns the manufacturer code (if any) of the command
func (UnbindResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (UnbindResponse) MarshalJSON() ([]byte, error) { return []byte("32802"), nil }

func (v *UnbindResponse) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h UnbindResponseHandler
	if h, found = handler.(UnbindResponseHandler); found {
		err = h.HandleUnbindResponse(frame, v)
	}
	return
}

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

type MgmtLqiRequestHandler interface {
	HandleMgmtLqiRequest(frame Frame, cmd *MgmtLqiRequest) error
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
func (MgmtLqiRequest) Name() string { return `Mgmt Lqi Request` }

// Description of the command
func (MgmtLqiRequest) Description() string { return `` }

// ID of the command
func (MgmtLqiRequest) ID() CommandID { return MgmtLqiRequestCommand }

// Required
func (MgmtLqiRequest) Required() bool { return false }

// Cluster ID of the command
func (MgmtLqiRequest) Cluster() zcl.ClusterID { return 0x0031 }

// MnfCode returns the manufacturer code (if any) of the command
func (MgmtLqiRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MgmtLqiRequest) MarshalJSON() ([]byte, error) { return []byte("49"), nil }

func (v *MgmtLqiRequest) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h MgmtLqiRequestHandler
	if h, found = handler.(MgmtLqiRequestHandler); found {
		err = h.HandleMgmtLqiRequest(frame, v)
	}
	return
}

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

type MgmtLqiResponseHandler interface {
	HandleMgmtLqiResponse(frame Frame, cmd *MgmtLqiResponse) error
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
func (MgmtLqiResponse) Name() string { return `Mgmt Lqi Response` }

// Description of the command
func (MgmtLqiResponse) Description() string { return `` }

// ID of the command
func (MgmtLqiResponse) ID() CommandID { return MgmtLqiResponseCommand }

// Required
func (MgmtLqiResponse) Required() bool { return false }

// Cluster ID of the command
func (MgmtLqiResponse) Cluster() zcl.ClusterID { return 0x8031 }

// MnfCode returns the manufacturer code (if any) of the command
func (MgmtLqiResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MgmtLqiResponse) MarshalJSON() ([]byte, error) { return []byte("32817"), nil }

func (v *MgmtLqiResponse) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h MgmtLqiResponseHandler
	if h, found = handler.(MgmtLqiResponseHandler); found {
		err = h.HandleMgmtLqiResponse(frame, v)
	}
	return
}

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

type MgmtNwkUpdate struct {
	// Status Code, command is normally empty unless status is `Success`
	Status               Status
	ScannedChannels      ScannedChannels
	TotalTransmissions   TotalTransmissions
	TransmissionFailures TransmissionFailures
	EnergyValues         EnergyValues
}

type MgmtNwkUpdateHandler interface {
	HandleMgmtNwkUpdate(frame Frame, cmd *MgmtNwkUpdate) error
}

// MgmtNwkUpdateCommand is the Command ID of MgmtNwkUpdate
const MgmtNwkUpdateCommand CommandID = 0x8038

// Values returns all values of MgmtNwkUpdate
func (v *MgmtNwkUpdate) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.ScannedChannels,
		&v.TotalTransmissions,
		&v.TransmissionFailures,
		&v.EnergyValues,
	}
}

// Arguments returns all values of MgmtNwkUpdate
func (v *MgmtNwkUpdate) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "ScannedChannels", Argument: &v.ScannedChannels},
		{Name: "TotalTransmissions", Argument: &v.TotalTransmissions},
		{Name: "TransmissionFailures", Argument: &v.TransmissionFailures},
		{Name: "EnergyValues", Argument: &v.EnergyValues},
	}
}

// Name of the command
func (MgmtNwkUpdate) Name() string { return `Mgmt Nwk Update` }

// Description of the command
func (MgmtNwkUpdate) Description() string { return `` }

// ID of the command
func (MgmtNwkUpdate) ID() CommandID { return MgmtNwkUpdateCommand }

// Required
func (MgmtNwkUpdate) Required() bool { return false }

// Cluster ID of the command
func (MgmtNwkUpdate) Cluster() zcl.ClusterID { return 0x8038 }

// MnfCode returns the manufacturer code (if any) of the command
func (MgmtNwkUpdate) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MgmtNwkUpdate) MarshalJSON() ([]byte, error) { return []byte("32824"), nil }

func (v *MgmtNwkUpdate) Handle(frame Frame, handler interface{}) (rsp zcl.ZdoCommand, found bool, err error) {
	var h MgmtNwkUpdateHandler
	if h, found = handler.(MgmtNwkUpdateHandler); found {
		err = h.HandleMgmtNwkUpdate(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of MgmtNwkUpdate
func (v MgmtNwkUpdate) MarshalZcl() ([]byte, error) {
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
		if tmp, err = v.ScannedChannels.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.TotalTransmissions.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.TransmissionFailures.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.EnergyValues.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the MgmtNwkUpdate struct
func (v *MgmtNwkUpdate) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ScannedChannels).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TotalTransmissions).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransmissionFailures).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.EnergyValues).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MgmtNwkUpdate) String() string {
	return zcl.Sprintf(
		"MgmtNwkUpdate{"+zcl.StrJoin([]string{
			"Status(%v)",
			"ScannedChannels(%v)",
			"TotalTransmissions(%v)",
			"TransmissionFailures(%v)",
			"EnergyValues(%v)",
		}, " ")+"}",
		v.Status,
		v.ScannedChannels,
		v.TotalTransmissions,
		v.TransmissionFailures,
		v.EnergyValues,
	)
}
