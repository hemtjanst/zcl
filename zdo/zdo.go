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

func (v ComplexDescriptorAvailable) Scaled() float64 {
	return float64(v)
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
	case 0x02:
		return "End Device"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v LogicalType) IsCoordinator() bool { return v == 0x00 }
func (v LogicalType) IsRouter() bool      { return v == 0x01 }
func (v LogicalType) IsEndDevice() bool   { return v == 0x02 }
func (v *LogicalType) SetCoordinator()    { *v = 0x00 }
func (v *LogicalType) SetRouter()         { *v = 0x01 }
func (v *LogicalType) SetEndDevice()      { *v = 0x02 }

func (LogicalType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Coordinator"},
		{Value: 0x01, Name: "Router"},
		{Value: 0x02, Name: "End Device"},
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

func (v Lqi) Scaled() float64 {
	return float64(v) / 2.55
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
	case 0x1000:
		return "Cirronet"
	case 0x1001:
		return "Chipcon"
	case 0x1002:
		return "Ember"
	case 0x1003:
		return "National Tech"
	case 0x1004:
		return "Freescale"
	case 0x1005:
		return "IPCom"
	case 0x1006:
		return "San Juan Software"
	case 0x1007:
		return "TUV"
	case 0x1008:
		return "CompXs"
	case 0x1009:
		return "BM SpA"
	case 0x100a:
		return "AwarePoint"
	case 0x100b:
		return "Philips"
	case 0x100c:
		return "Luxoft"
	case 0x100d:
		return "Korvin"
	case 0x100e:
		return "One RF"
	case 0x100f:
		return "Software Technology Group"
	case 0x1010:
		return "Telegesis"
	case 0x1011:
		return "Visionic"
	case 0x1012:
		return "Insta"
	case 0x1013:
		return "Atalum"
	case 0x1014:
		return "Atmel"
	case 0x1015:
		return "Develco"
	case 0x1016:
		return "Honeywell"
	case 0x1017:
		return "RadioPulse"
	case 0x1018:
		return "Renesas"
	case 0x1019:
		return "Xanadu Wireless"
	case 0x101a:
		return "NEC Engineering"
	case 0x101b:
		return "Yamatake"
	case 0x101c:
		return "Tendril"
	case 0x101d:
		return "Assa Abloy"
	case 0x101e:
		return "Maxstream"
	case 0x101f:
		return "Neurocom"
	case 0x1020:
		return "Institute for Information Industry"
	case 0x1021:
		return "Vantage Controls"
	case 0x1022:
		return "iControl"
	case 0x1023:
		return "Raymarine"
	case 0x1024:
		return "LS Research"
	case 0x1025:
		return "Onity"
	case 0x1026:
		return "Mono Products"
	case 0x1027:
		return "RF Tech"
	case 0x1028:
		return "Itron"
	case 0x1029:
		return "Tritech"
	case 0x102a:
		return "Embedit"
	case 0x102b:
		return "S3C"
	case 0x102c:
		return "Siemens"
	case 0x102d:
		return "Mindtech"
	case 0x102e:
		return "LG Electronics"
	case 0x102f:
		return "Mitsubishi"
	case 0x1030:
		return "Johnson Controls"
	case 0x1031:
		return "PRI"
	case 0x1032:
		return "Knick"
	case 0x1033:
		return "Viconics"
	case 0x1034:
		return "Flexipanel"
	case 0x1035:
		return "Piasim Corporation"
	case 0x1036:
		return "Trane"
	case 0x1037:
		return "Jennic"
	case 0x1038:
		return "Living Independently"
	case 0x1039:
		return "AlertMe"
	case 0x103a:
		return "Daintree"
	case 0x103b:
		return "Aiji"
	case 0x103c:
		return "Telecom Italia"
	case 0x103d:
		return "Mikrokrets"
	case 0x103e:
		return "Oki Semi"
	case 0x103f:
		return "Newport Electronics"
	case 0x1040:
		return "Control4"
	case 0x1041:
		return "STMicro"
	case 0x1042:
		return "Ad-Sol Nissin"
	case 0x1043:
		return "DCSI"
	case 0x1044:
		return "France Telecom"
	case 0x1045:
		return "muNet"
	case 0x1046:
		return "Autani"
	case 0x1047:
		return "Colorado vNet"
	case 0x1048:
		return "Aerocomm"
	case 0x1049:
		return "Silicon Labs"
	case 0x104F:
		return "Crane"
	case 0x104a:
		return "Inncom"
	case 0x104b:
		return "Cannon"
	case 0x104c:
		return "Synapse"
	case 0x104d:
		return "Fisher Pierce/Sunrise"
	case 0x104e:
		return "CentraLite"
	case 0x1050:
		return "Mobilarm"
	case 0x1051:
		return "iMonitor"
	case 0x1052:
		return "Bartech"
	case 0x1053:
		return "Meshnetics"
	case 0x1054:
		return "LS Industrial"
	case 0x1055:
		return "Cason"
	case 0x1056:
		return "Wireless Glue"
	case 0x1057:
		return "Elster"
	case 0x1058:
		return "SMS Tec"
	case 0x1059:
		return "Onset Computer"
	case 0x105a:
		return "Riga Development"
	case 0x105b:
		return "Energate"
	case 0x105c:
		return "ConMed Linvatec"
	case 0x105d:
		return "PowerMand"
	case 0x105e:
		return "Schneider Electric"
	case 0x105f:
		return "Eaton"
	case 0x1060:
		return "Telular"
	case 0x1061:
		return "Delphi Medical"
	case 0x1062:
		return "EpiSensor"
	case 0x1063:
		return "Landis+Gyr"
	case 0x1064:
		return "Kaba Group"
	case 0x1065:
		return "Shure"
	case 0x1066:
		return "Comverge"
	case 0x1067:
		return "DBS Lodging"
	case 0x1068:
		return "Energy Aware"
	case 0x1069:
		return "Hidalgo"
	case 0x106a:
		return "Air2App"
	case 0x106b:
		return "AMX"
	case 0x106c:
		return "EDMI Pty"
	case 0x106d:
		return "Cyan Ltd"
	case 0x106e:
		return "System SPA"
	case 0x106f:
		return "Telit"
	case 0x1070:
		return "Kaga Electronics"
	case 0x1071:
		return "4-noks s.r.l."
	case 0x1072:
		return "Certicom"
	case 0x1073:
		return "Gridpoint"
	case 0x1074:
		return "Profile Systems"
	case 0x1075:
		return "Compacta International"
	case 0x1076:
		return "Freestyle Technology"
	case 0x1077:
		return "Alektrona"
	case 0x1078:
		return "Computime"
	case 0x1079:
		return "Remote Technologies"
	case 0x107a:
		return "Wavecom"
	case 0x107b:
		return "Energy Optimizers"
	case 0x107c:
		return "GE"
	case 0x107d:
		return "Jetlun"
	case 0x107e:
		return "Cipher Systems"
	case 0x107f:
		return "Corporate Systems Eng"
	case 0x1080:
		return "ecobee"
	case 0x1081:
		return "SMK"
	case 0x1082:
		return "Meshworks Wireless"
	case 0x1083:
		return "Ellips B.V."
	case 0x1084:
		return "Secure electrans"
	case 0x1085:
		return "CEDO"
	case 0x1086:
		return "Toshiba"
	case 0x1087:
		return "Digi International"
	case 0x1088:
		return "Ubilogix"
	case 0x1089:
		return "Echelon"
	case 0x1090:
		return "Green Energy Options"
	case 0x1091:
		return "Silver Spring Networks"
	case 0x1092:
		return "Black & Decker"
	case 0x1093:
		return "Aztech AssociatesInc."
	case 0x1094:
		return "A&D Co"
	case 0x1095:
		return "Rainforest Automation"
	case 0x1096:
		return "Carrier Electronics"
	case 0x1097:
		return "SyChip/Murata"
	case 0x1098:
		return "OpenPeak"
	case 0x1099:
		return "Passive Systems"
	case 0x109a:
		return "MMBResearch"
	case 0x109b:
		return "Leviton"
	case 0x109c:
		return "Korea Electric Power Data Network"
	case 0x109d:
		return "Comcast"
	case 0x109e:
		return "NEC Electronics"
	case 0x109f:
		return "Netvox"
	case 0x10a0:
		return "U-Control"
	case 0x10a1:
		return "Embedia Technologies"
	case 0x10a2:
		return "Sensus"
	case 0x10a3:
		return "SunriseTechnologies"
	case 0x10a4:
		return "MemtechCorp"
	case 0x10a5:
		return "Freebox"
	case 0x10a6:
		return "M2 Labs"
	case 0x10a7:
		return "BritishGas"
	case 0x10a8:
		return "Sentec"
	case 0x10a9:
		return "Navetas"
	case 0x10aa:
		return "Lightspeed Technologies"
	case 0x10ab:
		return "Oki Electric"
	case 0x10ac:
		return "Sistemas Inteligentes"
	case 0x10ad:
		return "Dometic"
	case 0x10ae:
		return "Alps"
	case 0x10af:
		return "EnergyHub"
	case 0x10b0:
		return "Kamstrup"
	case 0x10b1:
		return "EchoStar"
	case 0x10b2:
		return "EnerNOC"
	case 0x10b3:
		return "Eltav"
	case 0x10b4:
		return "Belkin"
	case 0x10b5:
		return "XStreamHD Wireless"
	case 0x10b6:
		return "Saturn South"
	case 0x10b7:
		return "GreenTrapOnline"
	case 0x10b8:
		return "SmartSynch"
	case 0x10b9:
		return "Nyce Control"
	case 0x10ba:
		return "ICM Controls"
	case 0x10bb:
		return "Millennium Electronics"
	case 0x10bc:
		return "Motorola"
	case 0x10bd:
		return "EmersonWhite-Rodgers"
	case 0x10be:
		return "Radio Thermostat"
	case 0x10bf:
		return "OMRONCorporation"
	case 0x10c0:
		return "GiiNii GlobalLimited"
	case 0x10c1:
		return "Fujitsu GeneralLimited"
	case 0x10c2:
		return "Peel Technologies"
	case 0x10c3:
		return "Accent"
	case 0x10c4:
		return "ByteSnap Design"
	case 0x10c5:
		return "NEC TOKIN Corporation"
	case 0x10c6:
		return "G4S JusticeServices"
	case 0x10c7:
		return "Trilliant Networks"
	case 0x10c8:
		return "Electrolux Italia"
	case 0x10c9:
		return "OnzoLtd"
	case 0x10ca:
		return "EnTekSystems"
	case 0x10cb:
		return "Philips 2"
	case 0x10cc:
		return "MainstreamEngineering"
	case 0x10cd:
		return "IndesitCompany"
	case 0x10ce:
		return "THINKECO"
	case 0x10cf:
		return "2D2C"
	case 0x10d0:
		return "GreenPeak"
	case 0x10d1:
		return "InterCEL"
	case 0x10d2:
		return "LG Electronics 2"
	case 0x10d3:
		return "Mitsumi Electric"
	case 0x10d4:
		return "Mitsumi Electric 2"
	case 0x10d5:
		return "Zentrum Mikroelektronik Dresden"
	case 0x10d6:
		return "Nest Labs"
	case 0x10d7:
		return "Exegin Technologies"
	case 0x10d8:
		return "Honeywell 2"
	case 0x10d9:
		return "Takahata Precision"
	case 0x10da:
		return "Sumitomo Electric Networks"
	case 0x10db:
		return "GE Energy"
	case 0x10dc:
		return "GE Appliances"
	case 0x10dd:
		return "Radiocrafts AS"
	case 0x10de:
		return "Ceiva"
	case 0x10df:
		return "TEC CO Co., Ltd"
	case 0x10e0:
		return "Chameleon Technology (UK) Ltd"
	case 0x10e1:
		return "Samsung"
	case 0x10e2:
		return "ruwido austria gmbh"
	case 0x10e3:
		return "Huawei Technologies Co., Ltd."
	case 0x10e4:
		return "Huawei Technologies Co., Ltd. 2"
	case 0x10e5:
		return "Greenwave Reality"
	case 0x10e6:
		return "BGlobal Metering Ltd"
	case 0x10e7:
		return "Mindteck"
	case 0x10e8:
		return "Ingersoll-Rand"
	case 0x10e9:
		return "Dius Computing Pty Ltd"
	case 0x10ea:
		return "Embedded Automation, Inc."
	case 0x10eb:
		return "ABB"
	case 0x10ec:
		return "Sony"
	case 0x10ed:
		return "Genus Power Infrastructures Limited"
	case 0x10ee:
		return "Universal Devices"
	case 0x10ef:
		return "Universal Devices 2"
	case 0x10f0:
		return "Metrum Technologies, LLC"
	case 0x10f1:
		return "Cisco"
	case 0x10f2:
		return "Ubisys technologies GmbH"
	case 0x10f3:
		return "Consert"
	case 0x10f4:
		return "Crestron Electronics"
	case 0x10f5:
		return "Enphase Energy"
	case 0x10f6:
		return "Invensys Controls"
	case 0x10f7:
		return "Mueller Systems, LLC"
	case 0x10f8:
		return "AAC Technologies Holding"
	case 0x10f9:
		return "U-NEXT Co., Ltd"
	case 0x10fa:
		return "Steelcase Inc."
	case 0x10fb:
		return "Telematics Wireless"
	case 0x10fc:
		return "Samil Power Co., Ltd"
	case 0x10fd:
		return "Pace Plc"
	case 0x10fe:
		return "Osborne Coinage Co."
	case 0x10ff:
		return "Powerwatch"
	case 0x1100:
		return "CANDELED GmbH"
	case 0x1101:
		return "FlexGrid S.R.L"
	case 0x1102:
		return "Humax"
	case 0x1103:
		return "Universal Electronics, Inc."
	case 0x1104:
		return "Advanced Energy"
	case 0x1105:
		return "BEGA Gantenbrink-Leuchten"
	case 0x1106:
		return "Brunel University"
	case 0x1107:
		return "Panasonic R&D Center Singapore"
	case 0x1108:
		return "eSystems Research"
	case 0x1109:
		return "Panamax"
	case 0x110a:
		return "Physical Graph Corporation"
	case 0x110b:
		return "EM-Lite Ltd."
	case 0x110c:
		return "Osram Sylvania"
	case 0x110d:
		return "2 Save Energy Ltd."
	case 0x110e:
		return "Planet Innovation Products Pty Ltd"
	case 0x110f:
		return "Ambient Devices, Inc."
	case 0x1110:
		return "Profalux"
	case 0x1111:
		return "Billion Electric Company (BEC)"
	case 0x1112:
		return "Embertec Pty Ltd"
	case 0x1113:
		return "IT Watchdogs"
	case 0x1114:
		return "Reloc"
	case 0x1115:
		return "Intel Corporation"
	case 0x1116:
		return "Trend Electronics Limited"
	case 0x1117:
		return "Moxa"
	case 0x1118:
		return "QEES"
	case 0x1119:
		return "SAYME Wireless Sensor Networks"
	case 0x111a:
		return "Pentair Aquatic Systems"
	case 0x111b:
		return "Orbit Irrigation"
	case 0x111c:
		return "California Eastern Laboratories"
	case 0x111d:
		return "Comcast 2"
	case 0x111e:
		return "IDT Technology Limited"
	case 0x111f:
		return "Pixela"
	case 0x1120:
		return "TiVo"
	case 0x1121:
		return "Fidure"
	case 0x1122:
		return "Marvell Semiconductor"
	case 0x1123:
		return "Wasion Group"
	case 0x1124:
		return "Jasco Products"
	case 0x1125:
		return "Shenzhen Kaifa Technology"
	case 0x1126:
		return "Netcomm Wireless"
	case 0x1127:
		return "Define Instruments"
	case 0x1128:
		return "In Home Displays"
	case 0x1129:
		return "Miele & Cie. KG"
	case 0x112a:
		return "Televes S.A."
	case 0x112b:
		return "Labelec"
	case 0x112c:
		return "China Electronics Standardization Institute"
	case 0x112d:
		return "Vectorform"
	case 0x112e:
		return "Busch-Jaeger Elektro"
	case 0x112f:
		return "Redpine Signals"
	case 0x1130:
		return "Bridges Electronic Technology"
	case 0x1131:
		return "Sercomm"
	case 0x1132:
		return "WSH GmbH wirsindheller"
	case 0x1133:
		return "Bosch Security Systems"
	case 0x1134:
		return "eZEX Corporation"
	case 0x1135:
		return "Dresden Elektronik Ingenieurtechnik GmbH"
	case 0x1136:
		return "MEAZON S.A."
	case 0x1137:
		return "Crow Electronic Engineering"
	case 0x1138:
		return "Harvard Engineering"
	case 0x1139:
		return "Andson(Beijing) Technology"
	case 0x113a:
		return "Adhoco AG"
	case 0x113b:
		return "Waxman Consumer Products Group"
	case 0x113c:
		return "Owon Technology"
	case 0x113d:
		return "Hitron Technologies"
	case 0x113e:
		return "Scemtec Steuerungstechnik GmbH"
	case 0x113f:
		return "Webee"
	case 0x1140:
		return "Grid2Home"
	case 0x1141:
		return "Telink Micro"
	case 0x1142:
		return "Jasmine Systems"
	case 0x1143:
		return "Bidgely"
	case 0x1144:
		return "Lutron"
	case 0x1145:
		return "IJENKO"
	case 0x1146:
		return "Starfield Electronic"
	case 0x1147:
		return "TCP"
	case 0x1148:
		return "Rogers Communications Partnership"
	case 0x1149:
		return "Cree"
	case 0x114a:
		return "Robert Bosch LLC"
	case 0x114b:
		return "Ibis Networks"
	case 0x114c:
		return "Quirky"
	case 0x114d:
		return "Efergy Technologies"
	case 0x114e:
		return "Smartlabs"
	case 0x114f:
		return "Everspring Industry"
	case 0x1150:
		return "Swann Communications"
	case 0x1151:
		return "Soneter"
	case 0x1152:
		return "Samsung SDS"
	case 0x1153:
		return "Uniband Electronic Corporation"
	case 0x1154:
		return "Accton Technology Corporation"
	case 0x1155:
		return "Bosch Thermotechnik GmbH"
	case 0x1156:
		return "Wincor Nixdorf Inc."
	case 0x1157:
		return "Ohsung Electronics"
	case 0x1158:
		return "Zen Within, Inc."
	case 0x1159:
		return "Tech4home, Lda."
	case 0x115A:
		return "Nanoleaf"
	case 0x115B:
		return "Keen Home, Inc."
	case 0x115C:
		return "Poly-Control APS"
	case 0x115D:
		return "Eastfield Lighting Co., Ltd Shenzhen"
	case 0x115E:
		return "IP Datatel, Inc."
	case 0x115F:
		return "Lumi United Techology, Ltd Shenzhen"
	case 0x1160:
		return "Sengled Optoelectronics Corp"
	case 0x1161:
		return "Remote Solution Co., Ltd."
	case 0x1162:
		return "ABB Genway Xiamen Electrical Equipment Co., Ltd."
	case 0x1163:
		return "Zhejiang Rexense Tech"
	case 0x1164:
		return "ForEE Technology"
	case 0x1165:
		return "Open Access Technology Intl."
	case 0x1166:
		return "INNR Lighting BV"
	case 0x1167:
		return "Techworld Industries"
	case 0x1168:
		return "Leedarson Lighting Co., Ltd."
	case 0x1169:
		return "Arzel Zoning"
	case 0x116A:
		return "Holley Technology"
	case 0x116B:
		return "Beldon Technologies"
	case 0x116C:
		return "Flextronics"
	case 0x116D:
		return "Shenzhen Meian"
	case 0x116E:
		return "Lowes"
	case 0x116F:
		return "Sigma Connectivity"
	case 0x1171:
		return "Wulian"
	case 0x1172:
		return "Plugwise B.V."
	case 0x1173:
		return "Titan Products"
	case 0x1174:
		return "Ecospectral"
	case 0x1175:
		return "D-Link"
	case 0x1176:
		return "Technicolor Home USA"
	case 0x1177:
		return "Opple Lighting"
	case 0x1178:
		return "Wistron NeWeb Corp."
	case 0x1179:
		return "QMotion Shades"
	case 0x117A:
		return "Insta Elektro GmbH"
	case 0x117B:
		return "Shanghai Vancount"
	case 0x117C:
		return "Ikea of Sweden"
	case 0x117D:
		return "RT-RK"
	case 0x117E:
		return "Shenzhen Feibit"
	case 0x117F:
		return "EuControls"
	case 0x1180:
		return "Telkonet"
	case 0x1181:
		return "Thermal Solution Resources"
	case 0x1182:
		return "PomCube"
	case 0x1183:
		return "Ei Electronics"
	case 0x1184:
		return "Optoga"
	case 0x1185:
		return "Stelpro"
	case 0x1186:
		return "Lynxus Technologies Corp."
	case 0x1187:
		return "Semiconductor Components"
	case 0x1188:
		return "TP-Link"
	case 0x1189:
		return "LEDVANCE LLC."
	case 0x118A:
		return "Nortek"
	case 0x118B:
		return "iRevo/Assa Abbloy Korea"
	case 0x118C:
		return "Midea"
	case 0x118D:
		return "ZF Friedrichshafen"
	case 0x118E:
		return "Checkit"
	case 0x118F:
		return "Aclara"
	case 0x1190:
		return "Nokia"
	case 0x1191:
		return "Goldcard High-tech Co., Ltd."
	case 0x1192:
		return "George Wilson Industries Ltd."
	case 0x1193:
		return "EASY SAVER CO.,INC"
	case 0x1194:
		return "ZTE Corporation"
	case 0x1195:
		return "ARRIS"
	case 0x1196:
		return "Reliance BIG TV"
	case 0x1197:
		return "Insight Energy Ventures/Powerley"
	case 0x1198:
		return "Thomas Research Products (Hubbell Lighting Inc.)"
	case 0x1199:
		return "Li Seng Technology"
	case 0x119A:
		return "System Level Solutions Inc."
	case 0x119B:
		return "Matrix Labs"
	case 0x119C:
		return "Sinope Technologies"
	case 0x119D:
		return "Jiuzhou Greeble"
	case 0x119E:
		return "Guangzhou Lanvee Tech. Co. Ltd."
	case 0x119F:
		return "Venstar"
	case 0x1200:
		return "SLV"
	case 0x1201:
		return "Halo Smart Labs"
	case 0x1202:
		return "Scout Security Inc."
	case 0x1203:
		return "Alibaba China Inc."
	case 0x1204:
		return "Resolution Products, Inc."
	case 0x1205:
		return "Smartlok Inc."
	case 0x1206:
		return "Lux Products Corp."
	case 0x1207:
		return "Vimar SpA"
	case 0x1208:
		return "Universal Lighting Technologies"
	case 0x1209:
		return "Robert Bosch, GmbH"
	case 0x120A:
		return "Accenture"
	case 0x120B:
		return "Heiman Technology Co., Ltd."
	case 0x120C:
		return "Shenzhen HOMA Technology Co., Ltd."
	case 0x120D:
		return "Vision-Electronics Technology"
	case 0x120E:
		return "Lenovo"
	case 0x120F:
		return "Presciense R&D"
	case 0x1210:
		return "Shenzhen Seastar Intelligence Co., Ltd."
	case 0x1211:
		return "Sensative AB"
	case 0x1212:
		return "SolarEdge"
	case 0x1213:
		return "Zipato"
	case 0x1214:
		return "China Fire & Security Sensing Manufacturing (iHorn)"
	case 0x1215:
		return "Quby BV"
	case 0x1216:
		return "Hangzhou Roombanker Technology Co., Ltd."
	case 0x1217:
		return "Amazon Lab126"
	case 0x1218:
		return "Paulmann Licht GmbH"
	case 0x1219:
		return "Shenzhen Orvibo Electronics Co. Ltd."
	case 0x121A:
		return "TCI Telecommunications"
	case 0x121B:
		return "Mueller-Licht International Inc."
	case 0x121C:
		return "Aurora Limited"
	case 0x121D:
		return "SmartDCC"
	case 0x121E:
		return "Shanghai UMEinfo Co. Ltd."
	case 0x121F:
		return "carbonTRACK"
	case 0x1220:
		return "Somfy"
	case 0x1221:
		return "Viessmann Elektronik GmbH"
	case 0x1222:
		return "Hildebrand Technology Ltd"
	case 0x1223:
		return "Onkyo Technology Corporation"
	case 0x1224:
		return "Shenzhen Sunricher Technology Ltd."
	case 0x1225:
		return "Xiu Xiu Technology Co., Ltd"
	case 0x1226:
		return "Zumtobel Group"
	case 0x1227:
		return "Shenzhen Kaadas Intelligent Technology Co. Ltd"
	case 0x1228:
		return "Shanghai Xiaoyan Technology Co. Ltd"
	case 0x1229:
		return "Cypress Semiconductor"
	case 0x122A:
		return "XAL GmbH"
	case 0x122B:
		return "Inergy Systems LLC"
	case 0x122C:
		return "Alfred Karcher GmbH & Co KG"
	case 0x122D:
		return "Adurolight Manufacturing"
	case 0x122E:
		return "Groupe Muller"
	case 0x122F:
		return "V-Mark Enterprises Inc."
	case 0x1230:
		return "Lead Energy AG"
	case 0x1231:
		return "UIOT Group"
	case 0x1232:
		return "Axxess Industries Inc."
	case 0x1233:
		return "Third Reality Inc."
	case 0x1234:
		return "DSR Corporation"
	case 0x1235:
		return "Guangzhou Vensi Intelligent Technology Co. Ltd."
	case 0x1236:
		return "Schlage Lock (Allegion)"
	case 0x1237:
		return "Net2Grid"
	case 0x1238:
		return "Airam Electric Oy Ab"
	case 0x1239:
		return "IMMAX WPB CZ"
	case 0x123A:
		return "ZIV Automation"
	case 0x123B:
		return "HangZhou iMagicTechnology Co., Ltd"
	case 0x123C:
		return "Xiamen Leelen Technology Co. Ltd."
	case 0x123D:
		return "Overkiz SAS"
	case 0x123E:
		return "Flonidan A/S"
	case 0x123F:
		return "HDL Automation Co., Ltd."
	case 0x1240:
		return "Ardomus Networks Corporation"
	case 0x1241:
		return "Samjin Co., Ltd."
	case 0x1242:
		return "Sprue Aegis PLC"
	case 0x1243:
		return "Indra Sistemas, S.A."
	case 0x1244:
		return "Shenzhen JBT Smart Lighting Co., Ltd."
	case 0x1245:
		return "GE Lighting & Current"
	case 0x1246:
		return "Danfoss A/S"
	case 0x1247:
		return "NIVISS PHP Sp. z o.o. Sp.k."
	case 0x1248:
		return "Shenzhen Fengliyuan Energy Conservating Technology Co. Ltd"
	case 0x1249:
		return "NEXELEC"
	case 0x124A:
		return "Sichuan Behome Prominent Technology Co., Ltd"
	case 0x124B:
		return "Fujian Star-net Communication Co., Ltd."
	case 0x124C:
		return "Toshiba Visual Solutions Corporation"
	case 0x124D:
		return "Latchable, Inc."
	case 0x124E:
		return "L&S Deutschland GmbH"
	case 0x124F:
		return "Gledopto Co., Ltd."
	case 0x1250:
		return "The Home Depot"
	case 0x1251:
		return "Neonlite International Ltd."
	case 0x1252:
		return "Arlo Technologies, Inc."
	case 0x1253:
		return "Xingluo Technology Co., Ltd."
	case 0x1254:
		return "Simon Electric (China) Co., Ltd."
	case 0x1255:
		return "Hangzhou Greatstar Industrial Co., Ltd."
	case 0x1256:
		return "Sequentric Energy Systems, LLC"
	case 0x1257:
		return "Solum Co., Ltd."
	case 0x1258:
		return "Eaglerise Electric & Electronic (China) Co., Ltd."
	case 0x1259:
		return "Fantem Technologies (Shenzhen) Co., Ltd."
	case 0x125A:
		return "Yunding Network Technology (Beijing) Co., Ltd."
	case 0x125B:
		return "Atlantic Group"
	case 0x125C:
		return "Xiamen Intretech, Inc."
	case 0x125D:
		return "Tuya Global Inc."
	case 0x125E:
		return "Xiamen Dnake Intelligent Technology Co., Ltd"
	case 0x125F:
		return "Niko nv"
	case 0x1260:
		return "Emporia Energy"
	case 0x1261:
		return "Sikom AS"
	case 0x1262:
		return "AXIS Labs, Inc."
	case 0x1263:
		return "Current Products Corporation"
	case 0x1264:
		return "MeteRSit SRL"
	case 0x1265:
		return "HORNBACH Baumarkt AG"
	case 0x1266:
		return "DiCEworld s.r.l. a socio unico"
	case 0x1267:
		return "ARC Technology Co., Ltd"
	case 0x1268:
		return "Hangzhou Konke Information Technology Co., Ltd."
	case 0x1269:
		return "SALTO Systems S.L."
	case 0x126A:
		return "Shenzhen Shyugj Technology Co., Ltd"
	case 0x126B:
		return "Brayden Automation Corporation"
	case 0x126C:
		return "Environexus Pty. Ltd."
	case 0x126D:
		return "Eltra nv/sa"
	case 0x126E:
		return "Xiaomi Communications Co., Ltd."
	case 0x126F:
		return "Shanghai Shuncom Electronic Technology Co., Ltd."
	case 0x1270:
		return "Voltalis S.A"
	case 0x1271:
		return "FEELUX Co., Ltd."
	case 0x1272:
		return "SmartPlus Inc."
	case 0x1273:
		return "Halemeier GmbH"
	case 0x1274:
		return "Trust International BBV"
	case 0x1275:
		return "Duke Energy Business Services LLC"
	case 0x1276:
		return "Calix, Inc."
	case 0x1277:
		return "ADEO"
	case 0x1278:
		return "Connected Response Limited"
	case 0x1279:
		return "StroyEnergoKom"
	case 0x127A:
		return "Lumitech Lighting Solution GmbH"
	case 0x127B:
		return "Verdant Environmental Technologies"
	case 0x127C:
		return "Alfred International"
	case 0x127D:
		return "Sansi LED Lighting"
	case 0x127E:
		return "Mindtree"
	case 0x127F:
		return "Nordic Semiconductor ASA"
	case 0x1280:
		return "Siterwell Electronics"
	case 0x1281:
		return "Briloner Leuchten GmbH"
	case 0x1282:
		return "Shenzhen SEI Technology"
	case 0x1283:
		return "Copper Labs"
	case 0x1284:
		return "Delta Dore"
	case 0x1285:
		return "Hager Group"
	case 0x1286:
		return "Shenzhen CoolKit Technology"
	case 0x1287:
		return "Hangzhou Sky-Lighting"
	case 0x1288:
		return "E.ON SE"
	case 0x1289:
		return "Lidl Stiftung"
	case 0x128A:
		return "Sichuan Changhong Network Technologies"
	case 0x128B:
		return "NodOn"
	case 0x128C:
		return "Jiangxi Innotech Technology"
	case 0x128D:
		return "Mercator Pty"
	case 0x128E:
		return "Beijing Ruying Tech"
	case 0x128F:
		return "EGLO Leuchten GmbH"
	case 0x1290:
		return "Pietro Fiorentini S.p.A"
	case 0x1291:
		return "Zehnder Group Vaux-Andigny"
	case 0x1292:
		return "BRK Brands"
	case 0x1293:
		return "Askey Computer"
	case 0x1294:
		return "PassiveBolt"
	case 0x1295:
		return "AVM Audiovisuelles"
	case 0x1296:
		return "Ningbo Suntech Lighting Tech"
	case 0x1297:
		return "Societe en Commandite Stello"
	case 0x1298:
		return "Vivint Smart Home"
	case 0x1299:
		return "Namron"
	case 0x129A:
		return "RADEMACHER Geraete Elektronik GmbH"
	case 0x129B:
		return "OMO Systems"
	case 0x129C:
		return "Siglis"
	case 0x129D:
		return "IMHOTEP CREATION"
	case 0x129E:
		return "icasa"
	case 0x129F:
		return "Level Home"
	case 0x1300:
		return "TIS Control"
	case 0x1301:
		return "Radisys India"
	case 0x1302:
		return "Veea"
	case 0x1303:
		return "FELL Technology"
	case 0x1304:
		return "Sowilo Design Services"
	case 0x1305:
		return "Lexi Devices"
	case 0x1306:
		return "Lifi Labs"
	case 0x1307:
		return "GRUNDFOS Holding"
	case 0x1308:
		return "SOURCING & CREATION"
	case 0x1309:
		return "Kraken Technologies"
	case 0x130A:
		return "EVE SYSTEMS"
	case 0x130B:
		return "LITE-ON TECHNOLOGY CORPORATION"
	case 0x130C:
		return "Focalcrest"
	case 0x130D:
		return "Bouffalo Lab (Nanjing)"
	case 0x130E:
		return "Wyze Labs"
	case 0x1337:
		return "Datek Wireless AS"
	case 0x1994:
		return "Gewiss S.p.A."
	case 0x2794:
		return "Climax Technology Cp., Ltd."
	}
	return zcl.Sprintf("%v", zcl.Zenum16(v))
}

func (v ManufacturerCode) IsCirronet() bool                                   { return v == 0x1000 }
func (v ManufacturerCode) IsChipcon() bool                                    { return v == 0x1001 }
func (v ManufacturerCode) IsEmber() bool                                      { return v == 0x1002 }
func (v ManufacturerCode) IsNationalTech() bool                               { return v == 0x1003 }
func (v ManufacturerCode) IsFreescale() bool                                  { return v == 0x1004 }
func (v ManufacturerCode) IsIpcom() bool                                      { return v == 0x1005 }
func (v ManufacturerCode) IsSanJuanSoftware() bool                            { return v == 0x1006 }
func (v ManufacturerCode) IsTuv() bool                                        { return v == 0x1007 }
func (v ManufacturerCode) IsCompxs() bool                                     { return v == 0x1008 }
func (v ManufacturerCode) IsBmSpa() bool                                      { return v == 0x1009 }
func (v ManufacturerCode) IsAwarepoint() bool                                 { return v == 0x100a }
func (v ManufacturerCode) IsPhilips() bool                                    { return v == 0x100b }
func (v ManufacturerCode) IsLuxoft() bool                                     { return v == 0x100c }
func (v ManufacturerCode) IsKorvin() bool                                     { return v == 0x100d }
func (v ManufacturerCode) IsOneRf() bool                                      { return v == 0x100e }
func (v ManufacturerCode) IsSoftwareTechnologyGroup() bool                    { return v == 0x100f }
func (v ManufacturerCode) IsTelegesis() bool                                  { return v == 0x1010 }
func (v ManufacturerCode) IsVisionic() bool                                   { return v == 0x1011 }
func (v ManufacturerCode) IsInsta() bool                                      { return v == 0x1012 }
func (v ManufacturerCode) IsAtalum() bool                                     { return v == 0x1013 }
func (v ManufacturerCode) IsAtmel() bool                                      { return v == 0x1014 }
func (v ManufacturerCode) IsDevelco() bool                                    { return v == 0x1015 }
func (v ManufacturerCode) IsHoneywell() bool                                  { return v == 0x1016 }
func (v ManufacturerCode) IsRadiopulse() bool                                 { return v == 0x1017 }
func (v ManufacturerCode) IsRenesas() bool                                    { return v == 0x1018 }
func (v ManufacturerCode) IsXanaduWireless() bool                             { return v == 0x1019 }
func (v ManufacturerCode) IsNecEngineering() bool                             { return v == 0x101a }
func (v ManufacturerCode) IsYamatake() bool                                   { return v == 0x101b }
func (v ManufacturerCode) IsTendril() bool                                    { return v == 0x101c }
func (v ManufacturerCode) IsAssaAbloy() bool                                  { return v == 0x101d }
func (v ManufacturerCode) IsMaxstream() bool                                  { return v == 0x101e }
func (v ManufacturerCode) IsNeurocom() bool                                   { return v == 0x101f }
func (v ManufacturerCode) IsInstituteForInformationIndustry() bool            { return v == 0x1020 }
func (v ManufacturerCode) IsVantageControls() bool                            { return v == 0x1021 }
func (v ManufacturerCode) IsIcontrol() bool                                   { return v == 0x1022 }
func (v ManufacturerCode) IsRaymarine() bool                                  { return v == 0x1023 }
func (v ManufacturerCode) IsLsResearch() bool                                 { return v == 0x1024 }
func (v ManufacturerCode) IsOnity() bool                                      { return v == 0x1025 }
func (v ManufacturerCode) IsMonoProducts() bool                               { return v == 0x1026 }
func (v ManufacturerCode) IsRfTech() bool                                     { return v == 0x1027 }
func (v ManufacturerCode) IsItron() bool                                      { return v == 0x1028 }
func (v ManufacturerCode) IsTritech() bool                                    { return v == 0x1029 }
func (v ManufacturerCode) IsEmbedit() bool                                    { return v == 0x102a }
func (v ManufacturerCode) IsS3c() bool                                        { return v == 0x102b }
func (v ManufacturerCode) IsSiemens() bool                                    { return v == 0x102c }
func (v ManufacturerCode) IsMindtech() bool                                   { return v == 0x102d }
func (v ManufacturerCode) IsLgElectronics() bool                              { return v == 0x102e }
func (v ManufacturerCode) IsMitsubishi() bool                                 { return v == 0x102f }
func (v ManufacturerCode) IsJohnsonControls() bool                            { return v == 0x1030 }
func (v ManufacturerCode) IsPri() bool                                        { return v == 0x1031 }
func (v ManufacturerCode) IsKnick() bool                                      { return v == 0x1032 }
func (v ManufacturerCode) IsViconics() bool                                   { return v == 0x1033 }
func (v ManufacturerCode) IsFlexipanel() bool                                 { return v == 0x1034 }
func (v ManufacturerCode) IsPiasimCorporation() bool                          { return v == 0x1035 }
func (v ManufacturerCode) IsTrane() bool                                      { return v == 0x1036 }
func (v ManufacturerCode) IsJennic() bool                                     { return v == 0x1037 }
func (v ManufacturerCode) IsLivingIndependently() bool                        { return v == 0x1038 }
func (v ManufacturerCode) IsAlertme() bool                                    { return v == 0x1039 }
func (v ManufacturerCode) IsDaintree() bool                                   { return v == 0x103a }
func (v ManufacturerCode) IsAiji() bool                                       { return v == 0x103b }
func (v ManufacturerCode) IsTelecomItalia() bool                              { return v == 0x103c }
func (v ManufacturerCode) IsMikrokrets() bool                                 { return v == 0x103d }
func (v ManufacturerCode) IsOkiSemi() bool                                    { return v == 0x103e }
func (v ManufacturerCode) IsNewportElectronics() bool                         { return v == 0x103f }
func (v ManufacturerCode) IsControl4() bool                                   { return v == 0x1040 }
func (v ManufacturerCode) IsStmicro() bool                                    { return v == 0x1041 }
func (v ManufacturerCode) IsAdSolNissin() bool                                { return v == 0x1042 }
func (v ManufacturerCode) IsDcsi() bool                                       { return v == 0x1043 }
func (v ManufacturerCode) IsFranceTelecom() bool                              { return v == 0x1044 }
func (v ManufacturerCode) IsMunet() bool                                      { return v == 0x1045 }
func (v ManufacturerCode) IsAutani() bool                                     { return v == 0x1046 }
func (v ManufacturerCode) IsColoradoVnet() bool                               { return v == 0x1047 }
func (v ManufacturerCode) IsAerocomm() bool                                   { return v == 0x1048 }
func (v ManufacturerCode) IsSiliconLabs() bool                                { return v == 0x1049 }
func (v ManufacturerCode) IsCrane() bool                                      { return v == 0x104F }
func (v ManufacturerCode) IsInncom() bool                                     { return v == 0x104a }
func (v ManufacturerCode) IsCannon() bool                                     { return v == 0x104b }
func (v ManufacturerCode) IsSynapse() bool                                    { return v == 0x104c }
func (v ManufacturerCode) IsFisherPierceSunrise() bool                        { return v == 0x104d }
func (v ManufacturerCode) IsCentralite() bool                                 { return v == 0x104e }
func (v ManufacturerCode) IsMobilarm() bool                                   { return v == 0x1050 }
func (v ManufacturerCode) IsImonitor() bool                                   { return v == 0x1051 }
func (v ManufacturerCode) IsBartech() bool                                    { return v == 0x1052 }
func (v ManufacturerCode) IsMeshnetics() bool                                 { return v == 0x1053 }
func (v ManufacturerCode) IsLsIndustrial() bool                               { return v == 0x1054 }
func (v ManufacturerCode) IsCason() bool                                      { return v == 0x1055 }
func (v ManufacturerCode) IsWirelessGlue() bool                               { return v == 0x1056 }
func (v ManufacturerCode) IsElster() bool                                     { return v == 0x1057 }
func (v ManufacturerCode) IsSmsTec() bool                                     { return v == 0x1058 }
func (v ManufacturerCode) IsOnsetComputer() bool                              { return v == 0x1059 }
func (v ManufacturerCode) IsRigaDevelopment() bool                            { return v == 0x105a }
func (v ManufacturerCode) IsEnergate() bool                                   { return v == 0x105b }
func (v ManufacturerCode) IsConmedLinvatec() bool                             { return v == 0x105c }
func (v ManufacturerCode) IsPowermand() bool                                  { return v == 0x105d }
func (v ManufacturerCode) IsSchneiderElectric() bool                          { return v == 0x105e }
func (v ManufacturerCode) IsEaton() bool                                      { return v == 0x105f }
func (v ManufacturerCode) IsTelular() bool                                    { return v == 0x1060 }
func (v ManufacturerCode) IsDelphiMedical() bool                              { return v == 0x1061 }
func (v ManufacturerCode) IsEpisensor() bool                                  { return v == 0x1062 }
func (v ManufacturerCode) IsLandisGyr() bool                                  { return v == 0x1063 }
func (v ManufacturerCode) IsKabaGroup() bool                                  { return v == 0x1064 }
func (v ManufacturerCode) IsShure() bool                                      { return v == 0x1065 }
func (v ManufacturerCode) IsComverge() bool                                   { return v == 0x1066 }
func (v ManufacturerCode) IsDbsLodging() bool                                 { return v == 0x1067 }
func (v ManufacturerCode) IsEnergyAware() bool                                { return v == 0x1068 }
func (v ManufacturerCode) IsHidalgo() bool                                    { return v == 0x1069 }
func (v ManufacturerCode) IsAir2app() bool                                    { return v == 0x106a }
func (v ManufacturerCode) IsAmx() bool                                        { return v == 0x106b }
func (v ManufacturerCode) IsEdmiPty() bool                                    { return v == 0x106c }
func (v ManufacturerCode) IsCyanLtd() bool                                    { return v == 0x106d }
func (v ManufacturerCode) IsSystemSpa() bool                                  { return v == 0x106e }
func (v ManufacturerCode) IsTelit() bool                                      { return v == 0x106f }
func (v ManufacturerCode) IsKagaElectronics() bool                            { return v == 0x1070 }
func (v ManufacturerCode) Is4NoksSRL() bool                                   { return v == 0x1071 }
func (v ManufacturerCode) IsCerticom() bool                                   { return v == 0x1072 }
func (v ManufacturerCode) IsGridpoint() bool                                  { return v == 0x1073 }
func (v ManufacturerCode) IsProfileSystems() bool                             { return v == 0x1074 }
func (v ManufacturerCode) IsCompactaInternational() bool                      { return v == 0x1075 }
func (v ManufacturerCode) IsFreestyleTechnology() bool                        { return v == 0x1076 }
func (v ManufacturerCode) IsAlektrona() bool                                  { return v == 0x1077 }
func (v ManufacturerCode) IsComputime() bool                                  { return v == 0x1078 }
func (v ManufacturerCode) IsRemoteTechnologies() bool                         { return v == 0x1079 }
func (v ManufacturerCode) IsWavecom() bool                                    { return v == 0x107a }
func (v ManufacturerCode) IsEnergyOptimizers() bool                           { return v == 0x107b }
func (v ManufacturerCode) IsGe() bool                                         { return v == 0x107c }
func (v ManufacturerCode) IsJetlun() bool                                     { return v == 0x107d }
func (v ManufacturerCode) IsCipherSystems() bool                              { return v == 0x107e }
func (v ManufacturerCode) IsCorporateSystemsEng() bool                        { return v == 0x107f }
func (v ManufacturerCode) IsEcobee() bool                                     { return v == 0x1080 }
func (v ManufacturerCode) IsSmk() bool                                        { return v == 0x1081 }
func (v ManufacturerCode) IsMeshworksWireless() bool                          { return v == 0x1082 }
func (v ManufacturerCode) IsEllipsBV() bool                                   { return v == 0x1083 }
func (v ManufacturerCode) IsSecureElectrans() bool                            { return v == 0x1084 }
func (v ManufacturerCode) IsCedo() bool                                       { return v == 0x1085 }
func (v ManufacturerCode) IsToshiba() bool                                    { return v == 0x1086 }
func (v ManufacturerCode) IsDigiInternational() bool                          { return v == 0x1087 }
func (v ManufacturerCode) IsUbilogix() bool                                   { return v == 0x1088 }
func (v ManufacturerCode) IsEchelon() bool                                    { return v == 0x1089 }
func (v ManufacturerCode) IsGreenEnergyOptions() bool                         { return v == 0x1090 }
func (v ManufacturerCode) IsSilverSpringNetworks() bool                       { return v == 0x1091 }
func (v ManufacturerCode) IsBlackDecker() bool                                { return v == 0x1092 }
func (v ManufacturerCode) IsAztechAssociatesinc() bool                        { return v == 0x1093 }
func (v ManufacturerCode) IsADCo() bool                                       { return v == 0x1094 }
func (v ManufacturerCode) IsRainforestAutomation() bool                       { return v == 0x1095 }
func (v ManufacturerCode) IsCarrierElectronics() bool                         { return v == 0x1096 }
func (v ManufacturerCode) IsSychipMurata() bool                               { return v == 0x1097 }
func (v ManufacturerCode) IsOpenpeak() bool                                   { return v == 0x1098 }
func (v ManufacturerCode) IsPassiveSystems() bool                             { return v == 0x1099 }
func (v ManufacturerCode) IsMmbresearch() bool                                { return v == 0x109a }
func (v ManufacturerCode) IsLeviton() bool                                    { return v == 0x109b }
func (v ManufacturerCode) IsKoreaElectricPowerDataNetwork() bool              { return v == 0x109c }
func (v ManufacturerCode) IsComcast() bool                                    { return v == 0x109d }
func (v ManufacturerCode) IsNecElectronics() bool                             { return v == 0x109e }
func (v ManufacturerCode) IsNetvox() bool                                     { return v == 0x109f }
func (v ManufacturerCode) IsUControl() bool                                   { return v == 0x10a0 }
func (v ManufacturerCode) IsEmbediaTechnologies() bool                        { return v == 0x10a1 }
func (v ManufacturerCode) IsSensus() bool                                     { return v == 0x10a2 }
func (v ManufacturerCode) IsSunrisetechnologies() bool                        { return v == 0x10a3 }
func (v ManufacturerCode) IsMemtechcorp() bool                                { return v == 0x10a4 }
func (v ManufacturerCode) IsFreebox() bool                                    { return v == 0x10a5 }
func (v ManufacturerCode) IsM2Labs() bool                                     { return v == 0x10a6 }
func (v ManufacturerCode) IsBritishgas() bool                                 { return v == 0x10a7 }
func (v ManufacturerCode) IsSentec() bool                                     { return v == 0x10a8 }
func (v ManufacturerCode) IsNavetas() bool                                    { return v == 0x10a9 }
func (v ManufacturerCode) IsLightspeedTechnologies() bool                     { return v == 0x10aa }
func (v ManufacturerCode) IsOkiElectric() bool                                { return v == 0x10ab }
func (v ManufacturerCode) IsSistemasInteligentes() bool                       { return v == 0x10ac }
func (v ManufacturerCode) IsDometic() bool                                    { return v == 0x10ad }
func (v ManufacturerCode) IsAlps() bool                                       { return v == 0x10ae }
func (v ManufacturerCode) IsEnergyhub() bool                                  { return v == 0x10af }
func (v ManufacturerCode) IsKamstrup() bool                                   { return v == 0x10b0 }
func (v ManufacturerCode) IsEchostar() bool                                   { return v == 0x10b1 }
func (v ManufacturerCode) IsEnernoc() bool                                    { return v == 0x10b2 }
func (v ManufacturerCode) IsEltav() bool                                      { return v == 0x10b3 }
func (v ManufacturerCode) IsBelkin() bool                                     { return v == 0x10b4 }
func (v ManufacturerCode) IsXstreamhdWireless() bool                          { return v == 0x10b5 }
func (v ManufacturerCode) IsSaturnSouth() bool                                { return v == 0x10b6 }
func (v ManufacturerCode) IsGreentraponline() bool                            { return v == 0x10b7 }
func (v ManufacturerCode) IsSmartsynch() bool                                 { return v == 0x10b8 }
func (v ManufacturerCode) IsNyceControl() bool                                { return v == 0x10b9 }
func (v ManufacturerCode) IsIcmControls() bool                                { return v == 0x10ba }
func (v ManufacturerCode) IsMillenniumElectronics() bool                      { return v == 0x10bb }
func (v ManufacturerCode) IsMotorola() bool                                   { return v == 0x10bc }
func (v ManufacturerCode) IsEmersonwhiteRodgers() bool                        { return v == 0x10bd }
func (v ManufacturerCode) IsRadioThermostat() bool                            { return v == 0x10be }
func (v ManufacturerCode) IsOmroncorporation() bool                           { return v == 0x10bf }
func (v ManufacturerCode) IsGiiniiGloballimited() bool                        { return v == 0x10c0 }
func (v ManufacturerCode) IsFujitsuGenerallimited() bool                      { return v == 0x10c1 }
func (v ManufacturerCode) IsPeelTechnologies() bool                           { return v == 0x10c2 }
func (v ManufacturerCode) IsAccent() bool                                     { return v == 0x10c3 }
func (v ManufacturerCode) IsBytesnapDesign() bool                             { return v == 0x10c4 }
func (v ManufacturerCode) IsNecTokinCorporation() bool                        { return v == 0x10c5 }
func (v ManufacturerCode) IsG4sJusticeservices() bool                         { return v == 0x10c6 }
func (v ManufacturerCode) IsTrilliantNetworks() bool                          { return v == 0x10c7 }
func (v ManufacturerCode) IsElectroluxItalia() bool                           { return v == 0x10c8 }
func (v ManufacturerCode) IsOnzoltd() bool                                    { return v == 0x10c9 }
func (v ManufacturerCode) IsEnteksystems() bool                               { return v == 0x10ca }
func (v ManufacturerCode) IsPhilips2() bool                                   { return v == 0x10cb }
func (v ManufacturerCode) IsMainstreamengineering() bool                      { return v == 0x10cc }
func (v ManufacturerCode) IsIndesitcompany() bool                             { return v == 0x10cd }
func (v ManufacturerCode) IsThinkeco() bool                                   { return v == 0x10ce }
func (v ManufacturerCode) Is2D2c() bool                                       { return v == 0x10cf }
func (v ManufacturerCode) IsGreenpeak() bool                                  { return v == 0x10d0 }
func (v ManufacturerCode) IsIntercel() bool                                   { return v == 0x10d1 }
func (v ManufacturerCode) IsLgElectronics2() bool                             { return v == 0x10d2 }
func (v ManufacturerCode) IsMitsumiElectric() bool                            { return v == 0x10d3 }
func (v ManufacturerCode) IsMitsumiElectric2() bool                           { return v == 0x10d4 }
func (v ManufacturerCode) IsZentrumMikroelektronikDresden() bool              { return v == 0x10d5 }
func (v ManufacturerCode) IsNestLabs() bool                                   { return v == 0x10d6 }
func (v ManufacturerCode) IsExeginTechnologies() bool                         { return v == 0x10d7 }
func (v ManufacturerCode) IsHoneywell2() bool                                 { return v == 0x10d8 }
func (v ManufacturerCode) IsTakahataPrecision() bool                          { return v == 0x10d9 }
func (v ManufacturerCode) IsSumitomoElectricNetworks() bool                   { return v == 0x10da }
func (v ManufacturerCode) IsGeEnergy() bool                                   { return v == 0x10db }
func (v ManufacturerCode) IsGeAppliances() bool                               { return v == 0x10dc }
func (v ManufacturerCode) IsRadiocraftsAs() bool                              { return v == 0x10dd }
func (v ManufacturerCode) IsCeiva() bool                                      { return v == 0x10de }
func (v ManufacturerCode) IsTecCoCoLtd() bool                                 { return v == 0x10df }
func (v ManufacturerCode) IsChameleonTechnologyUkLtd() bool                   { return v == 0x10e0 }
func (v ManufacturerCode) IsSamsung() bool                                    { return v == 0x10e1 }
func (v ManufacturerCode) IsRuwidoAustriaGmbh() bool                          { return v == 0x10e2 }
func (v ManufacturerCode) IsHuaweiTechnologiesCoLtd() bool                    { return v == 0x10e3 }
func (v ManufacturerCode) IsHuaweiTechnologiesCoLtd2() bool                   { return v == 0x10e4 }
func (v ManufacturerCode) IsGreenwaveReality() bool                           { return v == 0x10e5 }
func (v ManufacturerCode) IsBglobalMeteringLtd() bool                         { return v == 0x10e6 }
func (v ManufacturerCode) IsMindteck() bool                                   { return v == 0x10e7 }
func (v ManufacturerCode) IsIngersollRand() bool                              { return v == 0x10e8 }
func (v ManufacturerCode) IsDiusComputingPtyLtd() bool                        { return v == 0x10e9 }
func (v ManufacturerCode) IsEmbeddedAutomationInc() bool                      { return v == 0x10ea }
func (v ManufacturerCode) IsAbb() bool                                        { return v == 0x10eb }
func (v ManufacturerCode) IsSony() bool                                       { return v == 0x10ec }
func (v ManufacturerCode) IsGenusPowerInfrastructuresLimited() bool           { return v == 0x10ed }
func (v ManufacturerCode) IsUniversalDevices() bool                           { return v == 0x10ee }
func (v ManufacturerCode) IsUniversalDevices2() bool                          { return v == 0x10ef }
func (v ManufacturerCode) IsMetrumTechnologiesLlc() bool                      { return v == 0x10f0 }
func (v ManufacturerCode) IsCisco() bool                                      { return v == 0x10f1 }
func (v ManufacturerCode) IsUbisysTechnologiesGmbh() bool                     { return v == 0x10f2 }
func (v ManufacturerCode) IsConsert() bool                                    { return v == 0x10f3 }
func (v ManufacturerCode) IsCrestronElectronics() bool                        { return v == 0x10f4 }
func (v ManufacturerCode) IsEnphaseEnergy() bool                              { return v == 0x10f5 }
func (v ManufacturerCode) IsInvensysControls() bool                           { return v == 0x10f6 }
func (v ManufacturerCode) IsMuellerSystemsLlc() bool                          { return v == 0x10f7 }
func (v ManufacturerCode) IsAacTechnologiesHolding() bool                     { return v == 0x10f8 }
func (v ManufacturerCode) IsUNextCoLtd() bool                                 { return v == 0x10f9 }
func (v ManufacturerCode) IsSteelcaseInc() bool                               { return v == 0x10fa }
func (v ManufacturerCode) IsTelematicsWireless() bool                         { return v == 0x10fb }
func (v ManufacturerCode) IsSamilPowerCoLtd() bool                            { return v == 0x10fc }
func (v ManufacturerCode) IsPacePlc() bool                                    { return v == 0x10fd }
func (v ManufacturerCode) IsOsborneCoinageCo() bool                           { return v == 0x10fe }
func (v ManufacturerCode) IsPowerwatch() bool                                 { return v == 0x10ff }
func (v ManufacturerCode) IsCandeledGmbh() bool                               { return v == 0x1100 }
func (v ManufacturerCode) IsFlexgridSRL() bool                                { return v == 0x1101 }
func (v ManufacturerCode) IsHumax() bool                                      { return v == 0x1102 }
func (v ManufacturerCode) IsUniversalElectronicsInc() bool                    { return v == 0x1103 }
func (v ManufacturerCode) IsAdvancedEnergy() bool                             { return v == 0x1104 }
func (v ManufacturerCode) IsBegaGantenbrinkLeuchten() bool                    { return v == 0x1105 }
func (v ManufacturerCode) IsBrunelUniversity() bool                           { return v == 0x1106 }
func (v ManufacturerCode) IsPanasonicRDCenterSingapore() bool                 { return v == 0x1107 }
func (v ManufacturerCode) IsEsystemsResearch() bool                           { return v == 0x1108 }
func (v ManufacturerCode) IsPanamax() bool                                    { return v == 0x1109 }
func (v ManufacturerCode) IsPhysicalGraphCorporation() bool                   { return v == 0x110a }
func (v ManufacturerCode) IsEmLiteLtd() bool                                  { return v == 0x110b }
func (v ManufacturerCode) IsOsramSylvania() bool                              { return v == 0x110c }
func (v ManufacturerCode) Is2SaveEnergyLtd() bool                             { return v == 0x110d }
func (v ManufacturerCode) IsPlanetInnovationProductsPtyLtd() bool             { return v == 0x110e }
func (v ManufacturerCode) IsAmbientDevicesInc() bool                          { return v == 0x110f }
func (v ManufacturerCode) IsProfalux() bool                                   { return v == 0x1110 }
func (v ManufacturerCode) IsBillionElectricCompanyBec() bool                  { return v == 0x1111 }
func (v ManufacturerCode) IsEmbertecPtyLtd() bool                             { return v == 0x1112 }
func (v ManufacturerCode) IsItWatchdogs() bool                                { return v == 0x1113 }
func (v ManufacturerCode) IsReloc() bool                                      { return v == 0x1114 }
func (v ManufacturerCode) IsIntelCorporation() bool                           { return v == 0x1115 }
func (v ManufacturerCode) IsTrendElectronicsLimited() bool                    { return v == 0x1116 }
func (v ManufacturerCode) IsMoxa() bool                                       { return v == 0x1117 }
func (v ManufacturerCode) IsQees() bool                                       { return v == 0x1118 }
func (v ManufacturerCode) IsSaymeWirelessSensorNetworks() bool                { return v == 0x1119 }
func (v ManufacturerCode) IsPentairAquaticSystems() bool                      { return v == 0x111a }
func (v ManufacturerCode) IsOrbitIrrigation() bool                            { return v == 0x111b }
func (v ManufacturerCode) IsCaliforniaEasternLaboratories() bool              { return v == 0x111c }
func (v ManufacturerCode) IsComcast2() bool                                   { return v == 0x111d }
func (v ManufacturerCode) IsIdtTechnologyLimited() bool                       { return v == 0x111e }
func (v ManufacturerCode) IsPixela() bool                                     { return v == 0x111f }
func (v ManufacturerCode) IsTivo() bool                                       { return v == 0x1120 }
func (v ManufacturerCode) IsFidure() bool                                     { return v == 0x1121 }
func (v ManufacturerCode) IsMarvellSemiconductor() bool                       { return v == 0x1122 }
func (v ManufacturerCode) IsWasionGroup() bool                                { return v == 0x1123 }
func (v ManufacturerCode) IsJascoProducts() bool                              { return v == 0x1124 }
func (v ManufacturerCode) IsShenzhenKaifaTechnology() bool                    { return v == 0x1125 }
func (v ManufacturerCode) IsNetcommWireless() bool                            { return v == 0x1126 }
func (v ManufacturerCode) IsDefineInstruments() bool                          { return v == 0x1127 }
func (v ManufacturerCode) IsInHomeDisplays() bool                             { return v == 0x1128 }
func (v ManufacturerCode) IsMieleCieKg() bool                                 { return v == 0x1129 }
func (v ManufacturerCode) IsTelevesSA() bool                                  { return v == 0x112a }
func (v ManufacturerCode) IsLabelec() bool                                    { return v == 0x112b }
func (v ManufacturerCode) IsChinaElectronicsStandardizationInstitute() bool   { return v == 0x112c }
func (v ManufacturerCode) IsVectorform() bool                                 { return v == 0x112d }
func (v ManufacturerCode) IsBuschJaegerElektro() bool                         { return v == 0x112e }
func (v ManufacturerCode) IsRedpineSignals() bool                             { return v == 0x112f }
func (v ManufacturerCode) IsBridgesElectronicTechnology() bool                { return v == 0x1130 }
func (v ManufacturerCode) IsSercomm() bool                                    { return v == 0x1131 }
func (v ManufacturerCode) IsWshGmbhWirsindheller() bool                       { return v == 0x1132 }
func (v ManufacturerCode) IsBoschSecuritySystems() bool                       { return v == 0x1133 }
func (v ManufacturerCode) IsEzexCorporation() bool                            { return v == 0x1134 }
func (v ManufacturerCode) IsDresdenElektronikIngenieurtechnikGmbh() bool      { return v == 0x1135 }
func (v ManufacturerCode) IsMeazonSA() bool                                   { return v == 0x1136 }
func (v ManufacturerCode) IsCrowElectronicEngineering() bool                  { return v == 0x1137 }
func (v ManufacturerCode) IsHarvardEngineering() bool                         { return v == 0x1138 }
func (v ManufacturerCode) IsAndsonBeijingTechnology() bool                    { return v == 0x1139 }
func (v ManufacturerCode) IsAdhocoAg() bool                                   { return v == 0x113a }
func (v ManufacturerCode) IsWaxmanConsumerProductsGroup() bool                { return v == 0x113b }
func (v ManufacturerCode) IsOwonTechnology() bool                             { return v == 0x113c }
func (v ManufacturerCode) IsHitronTechnologies() bool                         { return v == 0x113d }
func (v ManufacturerCode) IsScemtecSteuerungstechnikGmbh() bool               { return v == 0x113e }
func (v ManufacturerCode) IsWebee() bool                                      { return v == 0x113f }
func (v ManufacturerCode) IsGrid2home() bool                                  { return v == 0x1140 }
func (v ManufacturerCode) IsTelinkMicro() bool                                { return v == 0x1141 }
func (v ManufacturerCode) IsJasmineSystems() bool                             { return v == 0x1142 }
func (v ManufacturerCode) IsBidgely() bool                                    { return v == 0x1143 }
func (v ManufacturerCode) IsLutron() bool                                     { return v == 0x1144 }
func (v ManufacturerCode) IsIjenko() bool                                     { return v == 0x1145 }
func (v ManufacturerCode) IsStarfieldElectronic() bool                        { return v == 0x1146 }
func (v ManufacturerCode) IsTcp() bool                                        { return v == 0x1147 }
func (v ManufacturerCode) IsRogersCommunicationsPartnership() bool            { return v == 0x1148 }
func (v ManufacturerCode) IsCree() bool                                       { return v == 0x1149 }
func (v ManufacturerCode) IsRobertBoschLlc() bool                             { return v == 0x114a }
func (v ManufacturerCode) IsIbisNetworks() bool                               { return v == 0x114b }
func (v ManufacturerCode) IsQuirky() bool                                     { return v == 0x114c }
func (v ManufacturerCode) IsEfergyTechnologies() bool                         { return v == 0x114d }
func (v ManufacturerCode) IsSmartlabs() bool                                  { return v == 0x114e }
func (v ManufacturerCode) IsEverspringIndustry() bool                         { return v == 0x114f }
func (v ManufacturerCode) IsSwannCommunications() bool                        { return v == 0x1150 }
func (v ManufacturerCode) IsSoneter() bool                                    { return v == 0x1151 }
func (v ManufacturerCode) IsSamsungSds() bool                                 { return v == 0x1152 }
func (v ManufacturerCode) IsUnibandElectronicCorporation() bool               { return v == 0x1153 }
func (v ManufacturerCode) IsAcctonTechnologyCorporation() bool                { return v == 0x1154 }
func (v ManufacturerCode) IsBoschThermotechnikGmbh() bool                     { return v == 0x1155 }
func (v ManufacturerCode) IsWincorNixdorfInc() bool                           { return v == 0x1156 }
func (v ManufacturerCode) IsOhsungElectronics() bool                          { return v == 0x1157 }
func (v ManufacturerCode) IsZenWithinInc() bool                               { return v == 0x1158 }
func (v ManufacturerCode) IsTech4homeLda() bool                               { return v == 0x1159 }
func (v ManufacturerCode) IsNanoleaf() bool                                   { return v == 0x115A }
func (v ManufacturerCode) IsKeenHomeInc() bool                                { return v == 0x115B }
func (v ManufacturerCode) IsPolyControlAps() bool                             { return v == 0x115C }
func (v ManufacturerCode) IsEastfieldLightingCoLtdShenzhen() bool             { return v == 0x115D }
func (v ManufacturerCode) IsIpDatatelInc() bool                               { return v == 0x115E }
func (v ManufacturerCode) IsLumiUnitedTechologyLtdShenzhen() bool             { return v == 0x115F }
func (v ManufacturerCode) IsSengledOptoelectronicsCorp() bool                 { return v == 0x1160 }
func (v ManufacturerCode) IsRemoteSolutionCoLtd() bool                        { return v == 0x1161 }
func (v ManufacturerCode) IsAbbGenwayXiamenElectricalEquipmentCoLtd() bool    { return v == 0x1162 }
func (v ManufacturerCode) IsZhejiangRexenseTech() bool                        { return v == 0x1163 }
func (v ManufacturerCode) IsForeeTechnology() bool                            { return v == 0x1164 }
func (v ManufacturerCode) IsOpenAccessTechnologyIntl() bool                   { return v == 0x1165 }
func (v ManufacturerCode) IsInnrLightingBv() bool                             { return v == 0x1166 }
func (v ManufacturerCode) IsTechworldIndustries() bool                        { return v == 0x1167 }
func (v ManufacturerCode) IsLeedarsonLightingCoLtd() bool                     { return v == 0x1168 }
func (v ManufacturerCode) IsArzelZoning() bool                                { return v == 0x1169 }
func (v ManufacturerCode) IsHolleyTechnology() bool                           { return v == 0x116A }
func (v ManufacturerCode) IsBeldonTechnologies() bool                         { return v == 0x116B }
func (v ManufacturerCode) IsFlextronics() bool                                { return v == 0x116C }
func (v ManufacturerCode) IsShenzhenMeian() bool                              { return v == 0x116D }
func (v ManufacturerCode) IsLowes() bool                                      { return v == 0x116E }
func (v ManufacturerCode) IsSigmaConnectivity() bool                          { return v == 0x116F }
func (v ManufacturerCode) IsWulian() bool                                     { return v == 0x1171 }
func (v ManufacturerCode) IsPlugwiseBV() bool                                 { return v == 0x1172 }
func (v ManufacturerCode) IsTitanProducts() bool                              { return v == 0x1173 }
func (v ManufacturerCode) IsEcospectral() bool                                { return v == 0x1174 }
func (v ManufacturerCode) IsDLink() bool                                      { return v == 0x1175 }
func (v ManufacturerCode) IsTechnicolorHomeUsa() bool                         { return v == 0x1176 }
func (v ManufacturerCode) IsOppleLighting() bool                              { return v == 0x1177 }
func (v ManufacturerCode) IsWistronNewebCorp() bool                           { return v == 0x1178 }
func (v ManufacturerCode) IsQmotionShades() bool                              { return v == 0x1179 }
func (v ManufacturerCode) IsInstaElektroGmbh() bool                           { return v == 0x117A }
func (v ManufacturerCode) IsShanghaiVancount() bool                           { return v == 0x117B }
func (v ManufacturerCode) IsIkeaOfSweden() bool                               { return v == 0x117C }
func (v ManufacturerCode) IsRtRk() bool                                       { return v == 0x117D }
func (v ManufacturerCode) IsShenzhenFeibit() bool                             { return v == 0x117E }
func (v ManufacturerCode) IsEucontrols() bool                                 { return v == 0x117F }
func (v ManufacturerCode) IsTelkonet() bool                                   { return v == 0x1180 }
func (v ManufacturerCode) IsThermalSolutionResources() bool                   { return v == 0x1181 }
func (v ManufacturerCode) IsPomcube() bool                                    { return v == 0x1182 }
func (v ManufacturerCode) IsEiElectronics() bool                              { return v == 0x1183 }
func (v ManufacturerCode) IsOptoga() bool                                     { return v == 0x1184 }
func (v ManufacturerCode) IsStelpro() bool                                    { return v == 0x1185 }
func (v ManufacturerCode) IsLynxusTechnologiesCorp() bool                     { return v == 0x1186 }
func (v ManufacturerCode) IsSemiconductorComponents() bool                    { return v == 0x1187 }
func (v ManufacturerCode) IsTpLink() bool                                     { return v == 0x1188 }
func (v ManufacturerCode) IsLedvanceLlc() bool                                { return v == 0x1189 }
func (v ManufacturerCode) IsNortek() bool                                     { return v == 0x118A }
func (v ManufacturerCode) IsIrevoAssaAbbloyKorea() bool                       { return v == 0x118B }
func (v ManufacturerCode) IsMidea() bool                                      { return v == 0x118C }
func (v ManufacturerCode) IsZfFriedrichshafen() bool                          { return v == 0x118D }
func (v ManufacturerCode) IsCheckit() bool                                    { return v == 0x118E }
func (v ManufacturerCode) IsAclara() bool                                     { return v == 0x118F }
func (v ManufacturerCode) IsNokia() bool                                      { return v == 0x1190 }
func (v ManufacturerCode) IsGoldcardHighTechCoLtd() bool                      { return v == 0x1191 }
func (v ManufacturerCode) IsGeorgeWilsonIndustriesLtd() bool                  { return v == 0x1192 }
func (v ManufacturerCode) IsEasySaverCoInc() bool                             { return v == 0x1193 }
func (v ManufacturerCode) IsZteCorporation() bool                             { return v == 0x1194 }
func (v ManufacturerCode) IsArris() bool                                      { return v == 0x1195 }
func (v ManufacturerCode) IsRelianceBigTv() bool                              { return v == 0x1196 }
func (v ManufacturerCode) IsInsightEnergyVenturesPowerley() bool              { return v == 0x1197 }
func (v ManufacturerCode) IsThomasResearchProductsHubbellLightingInc() bool   { return v == 0x1198 }
func (v ManufacturerCode) IsLiSengTechnology() bool                           { return v == 0x1199 }
func (v ManufacturerCode) IsSystemLevelSolutionsInc() bool                    { return v == 0x119A }
func (v ManufacturerCode) IsMatrixLabs() bool                                 { return v == 0x119B }
func (v ManufacturerCode) IsSinopeTechnologies() bool                         { return v == 0x119C }
func (v ManufacturerCode) IsJiuzhouGreeble() bool                             { return v == 0x119D }
func (v ManufacturerCode) IsGuangzhouLanveeTechCoLtd() bool                   { return v == 0x119E }
func (v ManufacturerCode) IsVenstar() bool                                    { return v == 0x119F }
func (v ManufacturerCode) IsSlv() bool                                        { return v == 0x1200 }
func (v ManufacturerCode) IsHaloSmartLabs() bool                              { return v == 0x1201 }
func (v ManufacturerCode) IsScoutSecurityInc() bool                           { return v == 0x1202 }
func (v ManufacturerCode) IsAlibabaChinaInc() bool                            { return v == 0x1203 }
func (v ManufacturerCode) IsResolutionProductsInc() bool                      { return v == 0x1204 }
func (v ManufacturerCode) IsSmartlokInc() bool                                { return v == 0x1205 }
func (v ManufacturerCode) IsLuxProductsCorp() bool                            { return v == 0x1206 }
func (v ManufacturerCode) IsVimarSpa() bool                                   { return v == 0x1207 }
func (v ManufacturerCode) IsUniversalLightingTechnologies() bool              { return v == 0x1208 }
func (v ManufacturerCode) IsRobertBoschGmbh() bool                            { return v == 0x1209 }
func (v ManufacturerCode) IsAccenture() bool                                  { return v == 0x120A }
func (v ManufacturerCode) IsHeimanTechnologyCoLtd() bool                      { return v == 0x120B }
func (v ManufacturerCode) IsShenzhenHomaTechnologyCoLtd() bool                { return v == 0x120C }
func (v ManufacturerCode) IsVisionElectronicsTechnology() bool                { return v == 0x120D }
func (v ManufacturerCode) IsLenovo() bool                                     { return v == 0x120E }
func (v ManufacturerCode) IsPrescienseRD() bool                               { return v == 0x120F }
func (v ManufacturerCode) IsShenzhenSeastarIntelligenceCoLtd() bool           { return v == 0x1210 }
func (v ManufacturerCode) IsSensativeAb() bool                                { return v == 0x1211 }
func (v ManufacturerCode) IsSolaredge() bool                                  { return v == 0x1212 }
func (v ManufacturerCode) IsZipato() bool                                     { return v == 0x1213 }
func (v ManufacturerCode) IsChinaFireSecuritySensingManufacturingIhorn() bool { return v == 0x1214 }
func (v ManufacturerCode) IsQubyBv() bool                                     { return v == 0x1215 }
func (v ManufacturerCode) IsHangzhouRoombankerTechnologyCoLtd() bool          { return v == 0x1216 }
func (v ManufacturerCode) IsAmazonLab126() bool                               { return v == 0x1217 }
func (v ManufacturerCode) IsPaulmannLichtGmbh() bool                          { return v == 0x1218 }
func (v ManufacturerCode) IsShenzhenOrviboElectronicsCoLtd() bool             { return v == 0x1219 }
func (v ManufacturerCode) IsTciTelecommunications() bool                      { return v == 0x121A }
func (v ManufacturerCode) IsMuellerLichtInternationalInc() bool               { return v == 0x121B }
func (v ManufacturerCode) IsAuroraLimited() bool                              { return v == 0x121C }
func (v ManufacturerCode) IsSmartdcc() bool                                   { return v == 0x121D }
func (v ManufacturerCode) IsShanghaiUmeinfoCoLtd() bool                       { return v == 0x121E }
func (v ManufacturerCode) IsCarbontrack() bool                                { return v == 0x121F }
func (v ManufacturerCode) IsSomfy() bool                                      { return v == 0x1220 }
func (v ManufacturerCode) IsViessmannElektronikGmbh() bool                    { return v == 0x1221 }
func (v ManufacturerCode) IsHildebrandTechnologyLtd() bool                    { return v == 0x1222 }
func (v ManufacturerCode) IsOnkyoTechnologyCorporation() bool                 { return v == 0x1223 }
func (v ManufacturerCode) IsShenzhenSunricherTechnologyLtd() bool             { return v == 0x1224 }
func (v ManufacturerCode) IsXiuXiuTechnologyCoLtd() bool                      { return v == 0x1225 }
func (v ManufacturerCode) IsZumtobelGroup() bool                              { return v == 0x1226 }
func (v ManufacturerCode) IsShenzhenKaadasIntelligentTechnologyCoLtd() bool   { return v == 0x1227 }
func (v ManufacturerCode) IsShanghaiXiaoyanTechnologyCoLtd() bool             { return v == 0x1228 }
func (v ManufacturerCode) IsCypressSemiconductor() bool                       { return v == 0x1229 }
func (v ManufacturerCode) IsXalGmbh() bool                                    { return v == 0x122A }
func (v ManufacturerCode) IsInergySystemsLlc() bool                           { return v == 0x122B }
func (v ManufacturerCode) IsAlfredKarcherGmbhCoKg() bool                      { return v == 0x122C }
func (v ManufacturerCode) IsAdurolightManufacturing() bool                    { return v == 0x122D }
func (v ManufacturerCode) IsGroupeMuller() bool                               { return v == 0x122E }
func (v ManufacturerCode) IsVMarkEnterprisesInc() bool                        { return v == 0x122F }
func (v ManufacturerCode) IsLeadEnergyAg() bool                               { return v == 0x1230 }
func (v ManufacturerCode) IsUiotGroup() bool                                  { return v == 0x1231 }
func (v ManufacturerCode) IsAxxessIndustriesInc() bool                        { return v == 0x1232 }
func (v ManufacturerCode) IsThirdRealityInc() bool                            { return v == 0x1233 }
func (v ManufacturerCode) IsDsrCorporation() bool                             { return v == 0x1234 }
func (v ManufacturerCode) IsGuangzhouVensiIntelligentTechnologyCoLtd() bool   { return v == 0x1235 }
func (v ManufacturerCode) IsSchlageLockAllegion() bool                        { return v == 0x1236 }
func (v ManufacturerCode) IsNet2grid() bool                                   { return v == 0x1237 }
func (v ManufacturerCode) IsAiramElectricOyAb() bool                          { return v == 0x1238 }
func (v ManufacturerCode) IsImmaxWpbCz() bool                                 { return v == 0x1239 }
func (v ManufacturerCode) IsZivAutomation() bool                              { return v == 0x123A }
func (v ManufacturerCode) IsHangzhouImagictechnologyCoLtd() bool              { return v == 0x123B }
func (v ManufacturerCode) IsXiamenLeelenTechnologyCoLtd() bool                { return v == 0x123C }
func (v ManufacturerCode) IsOverkizSas() bool                                 { return v == 0x123D }
func (v ManufacturerCode) IsFlonidanAS() bool                                 { return v == 0x123E }
func (v ManufacturerCode) IsHdlAutomationCoLtd() bool                         { return v == 0x123F }
func (v ManufacturerCode) IsArdomusNetworksCorporation() bool                 { return v == 0x1240 }
func (v ManufacturerCode) IsSamjinCoLtd() bool                                { return v == 0x1241 }
func (v ManufacturerCode) IsSprueAegisPlc() bool                              { return v == 0x1242 }
func (v ManufacturerCode) IsIndraSistemasSA() bool                            { return v == 0x1243 }
func (v ManufacturerCode) IsShenzhenJbtSmartLightingCoLtd() bool              { return v == 0x1244 }
func (v ManufacturerCode) IsGeLightingCurrent() bool                          { return v == 0x1245 }
func (v ManufacturerCode) IsDanfossAS() bool                                  { return v == 0x1246 }
func (v ManufacturerCode) IsNivissPhpSpZOOSpK() bool                          { return v == 0x1247 }
func (v ManufacturerCode) IsShenzhenFengliyuanEnergyConservatingTechnologyCoLtd() bool {
	return v == 0x1248
}
func (v ManufacturerCode) IsNexelec() bool                                          { return v == 0x1249 }
func (v ManufacturerCode) IsSichuanBehomeProminentTechnologyCoLtd() bool            { return v == 0x124A }
func (v ManufacturerCode) IsFujianStarNetCommunicationCoLtd() bool                  { return v == 0x124B }
func (v ManufacturerCode) IsToshibaVisualSolutionsCorporation() bool                { return v == 0x124C }
func (v ManufacturerCode) IsLatchableInc() bool                                     { return v == 0x124D }
func (v ManufacturerCode) IsLSDeutschlandGmbh() bool                                { return v == 0x124E }
func (v ManufacturerCode) IsGledoptoCoLtd() bool                                    { return v == 0x124F }
func (v ManufacturerCode) IsTheHomeDepot() bool                                     { return v == 0x1250 }
func (v ManufacturerCode) IsNeonliteInternationalLtd() bool                         { return v == 0x1251 }
func (v ManufacturerCode) IsArloTechnologiesInc() bool                              { return v == 0x1252 }
func (v ManufacturerCode) IsXingluoTechnologyCoLtd() bool                           { return v == 0x1253 }
func (v ManufacturerCode) IsSimonElectricChinaCoLtd() bool                          { return v == 0x1254 }
func (v ManufacturerCode) IsHangzhouGreatstarIndustrialCoLtd() bool                 { return v == 0x1255 }
func (v ManufacturerCode) IsSequentricEnergySystemsLlc() bool                       { return v == 0x1256 }
func (v ManufacturerCode) IsSolumCoLtd() bool                                       { return v == 0x1257 }
func (v ManufacturerCode) IsEagleriseElectricElectronicChinaCoLtd() bool            { return v == 0x1258 }
func (v ManufacturerCode) IsFantemTechnologiesShenzhenCoLtd() bool                  { return v == 0x1259 }
func (v ManufacturerCode) IsYundingNetworkTechnologyBeijingCoLtd() bool             { return v == 0x125A }
func (v ManufacturerCode) IsAtlanticGroup() bool                                    { return v == 0x125B }
func (v ManufacturerCode) IsXiamenIntretechInc() bool                               { return v == 0x125C }
func (v ManufacturerCode) IsTuyaGlobalInc() bool                                    { return v == 0x125D }
func (v ManufacturerCode) IsXiamenDnakeIntelligentTechnologyCoLtd() bool            { return v == 0x125E }
func (v ManufacturerCode) IsNikoNv() bool                                           { return v == 0x125F }
func (v ManufacturerCode) IsEmporiaEnergy() bool                                    { return v == 0x1260 }
func (v ManufacturerCode) IsSikomAs() bool                                          { return v == 0x1261 }
func (v ManufacturerCode) IsAxisLabsInc() bool                                      { return v == 0x1262 }
func (v ManufacturerCode) IsCurrentProductsCorporation() bool                       { return v == 0x1263 }
func (v ManufacturerCode) IsMetersitSrl() bool                                      { return v == 0x1264 }
func (v ManufacturerCode) IsHornbachBaumarktAg() bool                               { return v == 0x1265 }
func (v ManufacturerCode) IsDiceworldSRLASocioUnico() bool                          { return v == 0x1266 }
func (v ManufacturerCode) IsArcTechnologyCoLtd() bool                               { return v == 0x1267 }
func (v ManufacturerCode) IsHangzhouKonkeInformationTechnologyCoLtd() bool          { return v == 0x1268 }
func (v ManufacturerCode) IsSaltoSystemsSL() bool                                   { return v == 0x1269 }
func (v ManufacturerCode) IsShenzhenShyugjTechnologyCoLtd() bool                    { return v == 0x126A }
func (v ManufacturerCode) IsBraydenAutomationCorporation() bool                     { return v == 0x126B }
func (v ManufacturerCode) IsEnvironexusPtyLtd() bool                                { return v == 0x126C }
func (v ManufacturerCode) IsEltraNvSa() bool                                        { return v == 0x126D }
func (v ManufacturerCode) IsXiaomiCommunicationsCoLtd() bool                        { return v == 0x126E }
func (v ManufacturerCode) IsShanghaiShuncomElectronicTechnologyCoLtd() bool         { return v == 0x126F }
func (v ManufacturerCode) IsVoltalisSA() bool                                       { return v == 0x1270 }
func (v ManufacturerCode) IsFeeluxCoLtd() bool                                      { return v == 0x1271 }
func (v ManufacturerCode) IsSmartplusInc() bool                                     { return v == 0x1272 }
func (v ManufacturerCode) IsHalemeierGmbh() bool                                    { return v == 0x1273 }
func (v ManufacturerCode) IsTrustInternationalBbv() bool                            { return v == 0x1274 }
func (v ManufacturerCode) IsDukeEnergyBusinessServicesLlc() bool                    { return v == 0x1275 }
func (v ManufacturerCode) IsCalixInc() bool                                         { return v == 0x1276 }
func (v ManufacturerCode) IsAdeo() bool                                             { return v == 0x1277 }
func (v ManufacturerCode) IsConnectedResponseLimited() bool                         { return v == 0x1278 }
func (v ManufacturerCode) IsStroyenergokom() bool                                   { return v == 0x1279 }
func (v ManufacturerCode) IsLumitechLightingSolutionGmbh() bool                     { return v == 0x127A }
func (v ManufacturerCode) IsVerdantEnvironmentalTechnologies() bool                 { return v == 0x127B }
func (v ManufacturerCode) IsAlfredInternational() bool                              { return v == 0x127C }
func (v ManufacturerCode) IsSansiLedLighting() bool                                 { return v == 0x127D }
func (v ManufacturerCode) IsMindtree() bool                                         { return v == 0x127E }
func (v ManufacturerCode) IsNordicSemiconductorAsa() bool                           { return v == 0x127F }
func (v ManufacturerCode) IsSiterwellElectronics() bool                             { return v == 0x1280 }
func (v ManufacturerCode) IsBrilonerLeuchtenGmbh() bool                             { return v == 0x1281 }
func (v ManufacturerCode) IsShenzhenSeiTechnology() bool                            { return v == 0x1282 }
func (v ManufacturerCode) IsCopperLabs() bool                                       { return v == 0x1283 }
func (v ManufacturerCode) IsDeltaDore() bool                                        { return v == 0x1284 }
func (v ManufacturerCode) IsHagerGroup() bool                                       { return v == 0x1285 }
func (v ManufacturerCode) IsShenzhenCoolkitTechnology() bool                        { return v == 0x1286 }
func (v ManufacturerCode) IsHangzhouSkyLighting() bool                              { return v == 0x1287 }
func (v ManufacturerCode) IsEOnSe() bool                                            { return v == 0x1288 }
func (v ManufacturerCode) IsLidlStiftung() bool                                     { return v == 0x1289 }
func (v ManufacturerCode) IsSichuanChanghongNetworkTechnologies() bool              { return v == 0x128A }
func (v ManufacturerCode) IsNodon() bool                                            { return v == 0x128B }
func (v ManufacturerCode) IsJiangxiInnotechTechnology() bool                        { return v == 0x128C }
func (v ManufacturerCode) IsMercatorPty() bool                                      { return v == 0x128D }
func (v ManufacturerCode) IsBeijingRuyingTech() bool                                { return v == 0x128E }
func (v ManufacturerCode) IsEgloLeuchtenGmbh() bool                                 { return v == 0x128F }
func (v ManufacturerCode) IsPietroFiorentiniSPA() bool                              { return v == 0x1290 }
func (v ManufacturerCode) IsZehnderGroupVauxAndigny() bool                          { return v == 0x1291 }
func (v ManufacturerCode) IsBrkBrands() bool                                        { return v == 0x1292 }
func (v ManufacturerCode) IsAskeyComputer() bool                                    { return v == 0x1293 }
func (v ManufacturerCode) IsPassivebolt() bool                                      { return v == 0x1294 }
func (v ManufacturerCode) IsAvmAudiovisuelles() bool                                { return v == 0x1295 }
func (v ManufacturerCode) IsNingboSuntechLightingTech() bool                        { return v == 0x1296 }
func (v ManufacturerCode) IsSocieteEnCommanditeStello() bool                        { return v == 0x1297 }
func (v ManufacturerCode) IsVivintSmartHome() bool                                  { return v == 0x1298 }
func (v ManufacturerCode) IsNamron() bool                                           { return v == 0x1299 }
func (v ManufacturerCode) IsRademacherGeraeteElektronikGmbh() bool                  { return v == 0x129A }
func (v ManufacturerCode) IsOmoSystems() bool                                       { return v == 0x129B }
func (v ManufacturerCode) IsSiglis() bool                                           { return v == 0x129C }
func (v ManufacturerCode) IsImhotepCreation() bool                                  { return v == 0x129D }
func (v ManufacturerCode) IsIcasa() bool                                            { return v == 0x129E }
func (v ManufacturerCode) IsLevelHome() bool                                        { return v == 0x129F }
func (v ManufacturerCode) IsTisControl() bool                                       { return v == 0x1300 }
func (v ManufacturerCode) IsRadisysIndia() bool                                     { return v == 0x1301 }
func (v ManufacturerCode) IsVeea() bool                                             { return v == 0x1302 }
func (v ManufacturerCode) IsFellTechnology() bool                                   { return v == 0x1303 }
func (v ManufacturerCode) IsSowiloDesignServices() bool                             { return v == 0x1304 }
func (v ManufacturerCode) IsLexiDevices() bool                                      { return v == 0x1305 }
func (v ManufacturerCode) IsLifiLabs() bool                                         { return v == 0x1306 }
func (v ManufacturerCode) IsGrundfosHolding() bool                                  { return v == 0x1307 }
func (v ManufacturerCode) IsSourcingCreation() bool                                 { return v == 0x1308 }
func (v ManufacturerCode) IsKrakenTechnologies() bool                               { return v == 0x1309 }
func (v ManufacturerCode) IsEveSystems() bool                                       { return v == 0x130A }
func (v ManufacturerCode) IsLiteOnTechnologyCorporation() bool                      { return v == 0x130B }
func (v ManufacturerCode) IsFocalcrest() bool                                       { return v == 0x130C }
func (v ManufacturerCode) IsBouffaloLabNanjing() bool                               { return v == 0x130D }
func (v ManufacturerCode) IsWyzeLabs() bool                                         { return v == 0x130E }
func (v ManufacturerCode) IsDatekWirelessAs() bool                                  { return v == 0x1337 }
func (v ManufacturerCode) IsGewissSPA() bool                                        { return v == 0x1994 }
func (v ManufacturerCode) IsClimaxTechnologyCpLtd() bool                            { return v == 0x2794 }
func (v *ManufacturerCode) SetCirronet()                                            { *v = 0x1000 }
func (v *ManufacturerCode) SetChipcon()                                             { *v = 0x1001 }
func (v *ManufacturerCode) SetEmber()                                               { *v = 0x1002 }
func (v *ManufacturerCode) SetNationalTech()                                        { *v = 0x1003 }
func (v *ManufacturerCode) SetFreescale()                                           { *v = 0x1004 }
func (v *ManufacturerCode) SetIpcom()                                               { *v = 0x1005 }
func (v *ManufacturerCode) SetSanJuanSoftware()                                     { *v = 0x1006 }
func (v *ManufacturerCode) SetTuv()                                                 { *v = 0x1007 }
func (v *ManufacturerCode) SetCompxs()                                              { *v = 0x1008 }
func (v *ManufacturerCode) SetBmSpa()                                               { *v = 0x1009 }
func (v *ManufacturerCode) SetAwarepoint()                                          { *v = 0x100a }
func (v *ManufacturerCode) SetPhilips()                                             { *v = 0x100b }
func (v *ManufacturerCode) SetLuxoft()                                              { *v = 0x100c }
func (v *ManufacturerCode) SetKorvin()                                              { *v = 0x100d }
func (v *ManufacturerCode) SetOneRf()                                               { *v = 0x100e }
func (v *ManufacturerCode) SetSoftwareTechnologyGroup()                             { *v = 0x100f }
func (v *ManufacturerCode) SetTelegesis()                                           { *v = 0x1010 }
func (v *ManufacturerCode) SetVisionic()                                            { *v = 0x1011 }
func (v *ManufacturerCode) SetInsta()                                               { *v = 0x1012 }
func (v *ManufacturerCode) SetAtalum()                                              { *v = 0x1013 }
func (v *ManufacturerCode) SetAtmel()                                               { *v = 0x1014 }
func (v *ManufacturerCode) SetDevelco()                                             { *v = 0x1015 }
func (v *ManufacturerCode) SetHoneywell()                                           { *v = 0x1016 }
func (v *ManufacturerCode) SetRadiopulse()                                          { *v = 0x1017 }
func (v *ManufacturerCode) SetRenesas()                                             { *v = 0x1018 }
func (v *ManufacturerCode) SetXanaduWireless()                                      { *v = 0x1019 }
func (v *ManufacturerCode) SetNecEngineering()                                      { *v = 0x101a }
func (v *ManufacturerCode) SetYamatake()                                            { *v = 0x101b }
func (v *ManufacturerCode) SetTendril()                                             { *v = 0x101c }
func (v *ManufacturerCode) SetAssaAbloy()                                           { *v = 0x101d }
func (v *ManufacturerCode) SetMaxstream()                                           { *v = 0x101e }
func (v *ManufacturerCode) SetNeurocom()                                            { *v = 0x101f }
func (v *ManufacturerCode) SetInstituteForInformationIndustry()                     { *v = 0x1020 }
func (v *ManufacturerCode) SetVantageControls()                                     { *v = 0x1021 }
func (v *ManufacturerCode) SetIcontrol()                                            { *v = 0x1022 }
func (v *ManufacturerCode) SetRaymarine()                                           { *v = 0x1023 }
func (v *ManufacturerCode) SetLsResearch()                                          { *v = 0x1024 }
func (v *ManufacturerCode) SetOnity()                                               { *v = 0x1025 }
func (v *ManufacturerCode) SetMonoProducts()                                        { *v = 0x1026 }
func (v *ManufacturerCode) SetRfTech()                                              { *v = 0x1027 }
func (v *ManufacturerCode) SetItron()                                               { *v = 0x1028 }
func (v *ManufacturerCode) SetTritech()                                             { *v = 0x1029 }
func (v *ManufacturerCode) SetEmbedit()                                             { *v = 0x102a }
func (v *ManufacturerCode) SetS3c()                                                 { *v = 0x102b }
func (v *ManufacturerCode) SetSiemens()                                             { *v = 0x102c }
func (v *ManufacturerCode) SetMindtech()                                            { *v = 0x102d }
func (v *ManufacturerCode) SetLgElectronics()                                       { *v = 0x102e }
func (v *ManufacturerCode) SetMitsubishi()                                          { *v = 0x102f }
func (v *ManufacturerCode) SetJohnsonControls()                                     { *v = 0x1030 }
func (v *ManufacturerCode) SetPri()                                                 { *v = 0x1031 }
func (v *ManufacturerCode) SetKnick()                                               { *v = 0x1032 }
func (v *ManufacturerCode) SetViconics()                                            { *v = 0x1033 }
func (v *ManufacturerCode) SetFlexipanel()                                          { *v = 0x1034 }
func (v *ManufacturerCode) SetPiasimCorporation()                                   { *v = 0x1035 }
func (v *ManufacturerCode) SetTrane()                                               { *v = 0x1036 }
func (v *ManufacturerCode) SetJennic()                                              { *v = 0x1037 }
func (v *ManufacturerCode) SetLivingIndependently()                                 { *v = 0x1038 }
func (v *ManufacturerCode) SetAlertme()                                             { *v = 0x1039 }
func (v *ManufacturerCode) SetDaintree()                                            { *v = 0x103a }
func (v *ManufacturerCode) SetAiji()                                                { *v = 0x103b }
func (v *ManufacturerCode) SetTelecomItalia()                                       { *v = 0x103c }
func (v *ManufacturerCode) SetMikrokrets()                                          { *v = 0x103d }
func (v *ManufacturerCode) SetOkiSemi()                                             { *v = 0x103e }
func (v *ManufacturerCode) SetNewportElectronics()                                  { *v = 0x103f }
func (v *ManufacturerCode) SetControl4()                                            { *v = 0x1040 }
func (v *ManufacturerCode) SetStmicro()                                             { *v = 0x1041 }
func (v *ManufacturerCode) SetAdSolNissin()                                         { *v = 0x1042 }
func (v *ManufacturerCode) SetDcsi()                                                { *v = 0x1043 }
func (v *ManufacturerCode) SetFranceTelecom()                                       { *v = 0x1044 }
func (v *ManufacturerCode) SetMunet()                                               { *v = 0x1045 }
func (v *ManufacturerCode) SetAutani()                                              { *v = 0x1046 }
func (v *ManufacturerCode) SetColoradoVnet()                                        { *v = 0x1047 }
func (v *ManufacturerCode) SetAerocomm()                                            { *v = 0x1048 }
func (v *ManufacturerCode) SetSiliconLabs()                                         { *v = 0x1049 }
func (v *ManufacturerCode) SetCrane()                                               { *v = 0x104F }
func (v *ManufacturerCode) SetInncom()                                              { *v = 0x104a }
func (v *ManufacturerCode) SetCannon()                                              { *v = 0x104b }
func (v *ManufacturerCode) SetSynapse()                                             { *v = 0x104c }
func (v *ManufacturerCode) SetFisherPierceSunrise()                                 { *v = 0x104d }
func (v *ManufacturerCode) SetCentralite()                                          { *v = 0x104e }
func (v *ManufacturerCode) SetMobilarm()                                            { *v = 0x1050 }
func (v *ManufacturerCode) SetImonitor()                                            { *v = 0x1051 }
func (v *ManufacturerCode) SetBartech()                                             { *v = 0x1052 }
func (v *ManufacturerCode) SetMeshnetics()                                          { *v = 0x1053 }
func (v *ManufacturerCode) SetLsIndustrial()                                        { *v = 0x1054 }
func (v *ManufacturerCode) SetCason()                                               { *v = 0x1055 }
func (v *ManufacturerCode) SetWirelessGlue()                                        { *v = 0x1056 }
func (v *ManufacturerCode) SetElster()                                              { *v = 0x1057 }
func (v *ManufacturerCode) SetSmsTec()                                              { *v = 0x1058 }
func (v *ManufacturerCode) SetOnsetComputer()                                       { *v = 0x1059 }
func (v *ManufacturerCode) SetRigaDevelopment()                                     { *v = 0x105a }
func (v *ManufacturerCode) SetEnergate()                                            { *v = 0x105b }
func (v *ManufacturerCode) SetConmedLinvatec()                                      { *v = 0x105c }
func (v *ManufacturerCode) SetPowermand()                                           { *v = 0x105d }
func (v *ManufacturerCode) SetSchneiderElectric()                                   { *v = 0x105e }
func (v *ManufacturerCode) SetEaton()                                               { *v = 0x105f }
func (v *ManufacturerCode) SetTelular()                                             { *v = 0x1060 }
func (v *ManufacturerCode) SetDelphiMedical()                                       { *v = 0x1061 }
func (v *ManufacturerCode) SetEpisensor()                                           { *v = 0x1062 }
func (v *ManufacturerCode) SetLandisGyr()                                           { *v = 0x1063 }
func (v *ManufacturerCode) SetKabaGroup()                                           { *v = 0x1064 }
func (v *ManufacturerCode) SetShure()                                               { *v = 0x1065 }
func (v *ManufacturerCode) SetComverge()                                            { *v = 0x1066 }
func (v *ManufacturerCode) SetDbsLodging()                                          { *v = 0x1067 }
func (v *ManufacturerCode) SetEnergyAware()                                         { *v = 0x1068 }
func (v *ManufacturerCode) SetHidalgo()                                             { *v = 0x1069 }
func (v *ManufacturerCode) SetAir2app()                                             { *v = 0x106a }
func (v *ManufacturerCode) SetAmx()                                                 { *v = 0x106b }
func (v *ManufacturerCode) SetEdmiPty()                                             { *v = 0x106c }
func (v *ManufacturerCode) SetCyanLtd()                                             { *v = 0x106d }
func (v *ManufacturerCode) SetSystemSpa()                                           { *v = 0x106e }
func (v *ManufacturerCode) SetTelit()                                               { *v = 0x106f }
func (v *ManufacturerCode) SetKagaElectronics()                                     { *v = 0x1070 }
func (v *ManufacturerCode) Set4NoksSRL()                                            { *v = 0x1071 }
func (v *ManufacturerCode) SetCerticom()                                            { *v = 0x1072 }
func (v *ManufacturerCode) SetGridpoint()                                           { *v = 0x1073 }
func (v *ManufacturerCode) SetProfileSystems()                                      { *v = 0x1074 }
func (v *ManufacturerCode) SetCompactaInternational()                               { *v = 0x1075 }
func (v *ManufacturerCode) SetFreestyleTechnology()                                 { *v = 0x1076 }
func (v *ManufacturerCode) SetAlektrona()                                           { *v = 0x1077 }
func (v *ManufacturerCode) SetComputime()                                           { *v = 0x1078 }
func (v *ManufacturerCode) SetRemoteTechnologies()                                  { *v = 0x1079 }
func (v *ManufacturerCode) SetWavecom()                                             { *v = 0x107a }
func (v *ManufacturerCode) SetEnergyOptimizers()                                    { *v = 0x107b }
func (v *ManufacturerCode) SetGe()                                                  { *v = 0x107c }
func (v *ManufacturerCode) SetJetlun()                                              { *v = 0x107d }
func (v *ManufacturerCode) SetCipherSystems()                                       { *v = 0x107e }
func (v *ManufacturerCode) SetCorporateSystemsEng()                                 { *v = 0x107f }
func (v *ManufacturerCode) SetEcobee()                                              { *v = 0x1080 }
func (v *ManufacturerCode) SetSmk()                                                 { *v = 0x1081 }
func (v *ManufacturerCode) SetMeshworksWireless()                                   { *v = 0x1082 }
func (v *ManufacturerCode) SetEllipsBV()                                            { *v = 0x1083 }
func (v *ManufacturerCode) SetSecureElectrans()                                     { *v = 0x1084 }
func (v *ManufacturerCode) SetCedo()                                                { *v = 0x1085 }
func (v *ManufacturerCode) SetToshiba()                                             { *v = 0x1086 }
func (v *ManufacturerCode) SetDigiInternational()                                   { *v = 0x1087 }
func (v *ManufacturerCode) SetUbilogix()                                            { *v = 0x1088 }
func (v *ManufacturerCode) SetEchelon()                                             { *v = 0x1089 }
func (v *ManufacturerCode) SetGreenEnergyOptions()                                  { *v = 0x1090 }
func (v *ManufacturerCode) SetSilverSpringNetworks()                                { *v = 0x1091 }
func (v *ManufacturerCode) SetBlackDecker()                                         { *v = 0x1092 }
func (v *ManufacturerCode) SetAztechAssociatesinc()                                 { *v = 0x1093 }
func (v *ManufacturerCode) SetADCo()                                                { *v = 0x1094 }
func (v *ManufacturerCode) SetRainforestAutomation()                                { *v = 0x1095 }
func (v *ManufacturerCode) SetCarrierElectronics()                                  { *v = 0x1096 }
func (v *ManufacturerCode) SetSychipMurata()                                        { *v = 0x1097 }
func (v *ManufacturerCode) SetOpenpeak()                                            { *v = 0x1098 }
func (v *ManufacturerCode) SetPassiveSystems()                                      { *v = 0x1099 }
func (v *ManufacturerCode) SetMmbresearch()                                         { *v = 0x109a }
func (v *ManufacturerCode) SetLeviton()                                             { *v = 0x109b }
func (v *ManufacturerCode) SetKoreaElectricPowerDataNetwork()                       { *v = 0x109c }
func (v *ManufacturerCode) SetComcast()                                             { *v = 0x109d }
func (v *ManufacturerCode) SetNecElectronics()                                      { *v = 0x109e }
func (v *ManufacturerCode) SetNetvox()                                              { *v = 0x109f }
func (v *ManufacturerCode) SetUControl()                                            { *v = 0x10a0 }
func (v *ManufacturerCode) SetEmbediaTechnologies()                                 { *v = 0x10a1 }
func (v *ManufacturerCode) SetSensus()                                              { *v = 0x10a2 }
func (v *ManufacturerCode) SetSunrisetechnologies()                                 { *v = 0x10a3 }
func (v *ManufacturerCode) SetMemtechcorp()                                         { *v = 0x10a4 }
func (v *ManufacturerCode) SetFreebox()                                             { *v = 0x10a5 }
func (v *ManufacturerCode) SetM2Labs()                                              { *v = 0x10a6 }
func (v *ManufacturerCode) SetBritishgas()                                          { *v = 0x10a7 }
func (v *ManufacturerCode) SetSentec()                                              { *v = 0x10a8 }
func (v *ManufacturerCode) SetNavetas()                                             { *v = 0x10a9 }
func (v *ManufacturerCode) SetLightspeedTechnologies()                              { *v = 0x10aa }
func (v *ManufacturerCode) SetOkiElectric()                                         { *v = 0x10ab }
func (v *ManufacturerCode) SetSistemasInteligentes()                                { *v = 0x10ac }
func (v *ManufacturerCode) SetDometic()                                             { *v = 0x10ad }
func (v *ManufacturerCode) SetAlps()                                                { *v = 0x10ae }
func (v *ManufacturerCode) SetEnergyhub()                                           { *v = 0x10af }
func (v *ManufacturerCode) SetKamstrup()                                            { *v = 0x10b0 }
func (v *ManufacturerCode) SetEchostar()                                            { *v = 0x10b1 }
func (v *ManufacturerCode) SetEnernoc()                                             { *v = 0x10b2 }
func (v *ManufacturerCode) SetEltav()                                               { *v = 0x10b3 }
func (v *ManufacturerCode) SetBelkin()                                              { *v = 0x10b4 }
func (v *ManufacturerCode) SetXstreamhdWireless()                                   { *v = 0x10b5 }
func (v *ManufacturerCode) SetSaturnSouth()                                         { *v = 0x10b6 }
func (v *ManufacturerCode) SetGreentraponline()                                     { *v = 0x10b7 }
func (v *ManufacturerCode) SetSmartsynch()                                          { *v = 0x10b8 }
func (v *ManufacturerCode) SetNyceControl()                                         { *v = 0x10b9 }
func (v *ManufacturerCode) SetIcmControls()                                         { *v = 0x10ba }
func (v *ManufacturerCode) SetMillenniumElectronics()                               { *v = 0x10bb }
func (v *ManufacturerCode) SetMotorola()                                            { *v = 0x10bc }
func (v *ManufacturerCode) SetEmersonwhiteRodgers()                                 { *v = 0x10bd }
func (v *ManufacturerCode) SetRadioThermostat()                                     { *v = 0x10be }
func (v *ManufacturerCode) SetOmroncorporation()                                    { *v = 0x10bf }
func (v *ManufacturerCode) SetGiiniiGloballimited()                                 { *v = 0x10c0 }
func (v *ManufacturerCode) SetFujitsuGenerallimited()                               { *v = 0x10c1 }
func (v *ManufacturerCode) SetPeelTechnologies()                                    { *v = 0x10c2 }
func (v *ManufacturerCode) SetAccent()                                              { *v = 0x10c3 }
func (v *ManufacturerCode) SetBytesnapDesign()                                      { *v = 0x10c4 }
func (v *ManufacturerCode) SetNecTokinCorporation()                                 { *v = 0x10c5 }
func (v *ManufacturerCode) SetG4sJusticeservices()                                  { *v = 0x10c6 }
func (v *ManufacturerCode) SetTrilliantNetworks()                                   { *v = 0x10c7 }
func (v *ManufacturerCode) SetElectroluxItalia()                                    { *v = 0x10c8 }
func (v *ManufacturerCode) SetOnzoltd()                                             { *v = 0x10c9 }
func (v *ManufacturerCode) SetEnteksystems()                                        { *v = 0x10ca }
func (v *ManufacturerCode) SetPhilips2()                                            { *v = 0x10cb }
func (v *ManufacturerCode) SetMainstreamengineering()                               { *v = 0x10cc }
func (v *ManufacturerCode) SetIndesitcompany()                                      { *v = 0x10cd }
func (v *ManufacturerCode) SetThinkeco()                                            { *v = 0x10ce }
func (v *ManufacturerCode) Set2D2c()                                                { *v = 0x10cf }
func (v *ManufacturerCode) SetGreenpeak()                                           { *v = 0x10d0 }
func (v *ManufacturerCode) SetIntercel()                                            { *v = 0x10d1 }
func (v *ManufacturerCode) SetLgElectronics2()                                      { *v = 0x10d2 }
func (v *ManufacturerCode) SetMitsumiElectric()                                     { *v = 0x10d3 }
func (v *ManufacturerCode) SetMitsumiElectric2()                                    { *v = 0x10d4 }
func (v *ManufacturerCode) SetZentrumMikroelektronikDresden()                       { *v = 0x10d5 }
func (v *ManufacturerCode) SetNestLabs()                                            { *v = 0x10d6 }
func (v *ManufacturerCode) SetExeginTechnologies()                                  { *v = 0x10d7 }
func (v *ManufacturerCode) SetHoneywell2()                                          { *v = 0x10d8 }
func (v *ManufacturerCode) SetTakahataPrecision()                                   { *v = 0x10d9 }
func (v *ManufacturerCode) SetSumitomoElectricNetworks()                            { *v = 0x10da }
func (v *ManufacturerCode) SetGeEnergy()                                            { *v = 0x10db }
func (v *ManufacturerCode) SetGeAppliances()                                        { *v = 0x10dc }
func (v *ManufacturerCode) SetRadiocraftsAs()                                       { *v = 0x10dd }
func (v *ManufacturerCode) SetCeiva()                                               { *v = 0x10de }
func (v *ManufacturerCode) SetTecCoCoLtd()                                          { *v = 0x10df }
func (v *ManufacturerCode) SetChameleonTechnologyUkLtd()                            { *v = 0x10e0 }
func (v *ManufacturerCode) SetSamsung()                                             { *v = 0x10e1 }
func (v *ManufacturerCode) SetRuwidoAustriaGmbh()                                   { *v = 0x10e2 }
func (v *ManufacturerCode) SetHuaweiTechnologiesCoLtd()                             { *v = 0x10e3 }
func (v *ManufacturerCode) SetHuaweiTechnologiesCoLtd2()                            { *v = 0x10e4 }
func (v *ManufacturerCode) SetGreenwaveReality()                                    { *v = 0x10e5 }
func (v *ManufacturerCode) SetBglobalMeteringLtd()                                  { *v = 0x10e6 }
func (v *ManufacturerCode) SetMindteck()                                            { *v = 0x10e7 }
func (v *ManufacturerCode) SetIngersollRand()                                       { *v = 0x10e8 }
func (v *ManufacturerCode) SetDiusComputingPtyLtd()                                 { *v = 0x10e9 }
func (v *ManufacturerCode) SetEmbeddedAutomationInc()                               { *v = 0x10ea }
func (v *ManufacturerCode) SetAbb()                                                 { *v = 0x10eb }
func (v *ManufacturerCode) SetSony()                                                { *v = 0x10ec }
func (v *ManufacturerCode) SetGenusPowerInfrastructuresLimited()                    { *v = 0x10ed }
func (v *ManufacturerCode) SetUniversalDevices()                                    { *v = 0x10ee }
func (v *ManufacturerCode) SetUniversalDevices2()                                   { *v = 0x10ef }
func (v *ManufacturerCode) SetMetrumTechnologiesLlc()                               { *v = 0x10f0 }
func (v *ManufacturerCode) SetCisco()                                               { *v = 0x10f1 }
func (v *ManufacturerCode) SetUbisysTechnologiesGmbh()                              { *v = 0x10f2 }
func (v *ManufacturerCode) SetConsert()                                             { *v = 0x10f3 }
func (v *ManufacturerCode) SetCrestronElectronics()                                 { *v = 0x10f4 }
func (v *ManufacturerCode) SetEnphaseEnergy()                                       { *v = 0x10f5 }
func (v *ManufacturerCode) SetInvensysControls()                                    { *v = 0x10f6 }
func (v *ManufacturerCode) SetMuellerSystemsLlc()                                   { *v = 0x10f7 }
func (v *ManufacturerCode) SetAacTechnologiesHolding()                              { *v = 0x10f8 }
func (v *ManufacturerCode) SetUNextCoLtd()                                          { *v = 0x10f9 }
func (v *ManufacturerCode) SetSteelcaseInc()                                        { *v = 0x10fa }
func (v *ManufacturerCode) SetTelematicsWireless()                                  { *v = 0x10fb }
func (v *ManufacturerCode) SetSamilPowerCoLtd()                                     { *v = 0x10fc }
func (v *ManufacturerCode) SetPacePlc()                                             { *v = 0x10fd }
func (v *ManufacturerCode) SetOsborneCoinageCo()                                    { *v = 0x10fe }
func (v *ManufacturerCode) SetPowerwatch()                                          { *v = 0x10ff }
func (v *ManufacturerCode) SetCandeledGmbh()                                        { *v = 0x1100 }
func (v *ManufacturerCode) SetFlexgridSRL()                                         { *v = 0x1101 }
func (v *ManufacturerCode) SetHumax()                                               { *v = 0x1102 }
func (v *ManufacturerCode) SetUniversalElectronicsInc()                             { *v = 0x1103 }
func (v *ManufacturerCode) SetAdvancedEnergy()                                      { *v = 0x1104 }
func (v *ManufacturerCode) SetBegaGantenbrinkLeuchten()                             { *v = 0x1105 }
func (v *ManufacturerCode) SetBrunelUniversity()                                    { *v = 0x1106 }
func (v *ManufacturerCode) SetPanasonicRDCenterSingapore()                          { *v = 0x1107 }
func (v *ManufacturerCode) SetEsystemsResearch()                                    { *v = 0x1108 }
func (v *ManufacturerCode) SetPanamax()                                             { *v = 0x1109 }
func (v *ManufacturerCode) SetPhysicalGraphCorporation()                            { *v = 0x110a }
func (v *ManufacturerCode) SetEmLiteLtd()                                           { *v = 0x110b }
func (v *ManufacturerCode) SetOsramSylvania()                                       { *v = 0x110c }
func (v *ManufacturerCode) Set2SaveEnergyLtd()                                      { *v = 0x110d }
func (v *ManufacturerCode) SetPlanetInnovationProductsPtyLtd()                      { *v = 0x110e }
func (v *ManufacturerCode) SetAmbientDevicesInc()                                   { *v = 0x110f }
func (v *ManufacturerCode) SetProfalux()                                            { *v = 0x1110 }
func (v *ManufacturerCode) SetBillionElectricCompanyBec()                           { *v = 0x1111 }
func (v *ManufacturerCode) SetEmbertecPtyLtd()                                      { *v = 0x1112 }
func (v *ManufacturerCode) SetItWatchdogs()                                         { *v = 0x1113 }
func (v *ManufacturerCode) SetReloc()                                               { *v = 0x1114 }
func (v *ManufacturerCode) SetIntelCorporation()                                    { *v = 0x1115 }
func (v *ManufacturerCode) SetTrendElectronicsLimited()                             { *v = 0x1116 }
func (v *ManufacturerCode) SetMoxa()                                                { *v = 0x1117 }
func (v *ManufacturerCode) SetQees()                                                { *v = 0x1118 }
func (v *ManufacturerCode) SetSaymeWirelessSensorNetworks()                         { *v = 0x1119 }
func (v *ManufacturerCode) SetPentairAquaticSystems()                               { *v = 0x111a }
func (v *ManufacturerCode) SetOrbitIrrigation()                                     { *v = 0x111b }
func (v *ManufacturerCode) SetCaliforniaEasternLaboratories()                       { *v = 0x111c }
func (v *ManufacturerCode) SetComcast2()                                            { *v = 0x111d }
func (v *ManufacturerCode) SetIdtTechnologyLimited()                                { *v = 0x111e }
func (v *ManufacturerCode) SetPixela()                                              { *v = 0x111f }
func (v *ManufacturerCode) SetTivo()                                                { *v = 0x1120 }
func (v *ManufacturerCode) SetFidure()                                              { *v = 0x1121 }
func (v *ManufacturerCode) SetMarvellSemiconductor()                                { *v = 0x1122 }
func (v *ManufacturerCode) SetWasionGroup()                                         { *v = 0x1123 }
func (v *ManufacturerCode) SetJascoProducts()                                       { *v = 0x1124 }
func (v *ManufacturerCode) SetShenzhenKaifaTechnology()                             { *v = 0x1125 }
func (v *ManufacturerCode) SetNetcommWireless()                                     { *v = 0x1126 }
func (v *ManufacturerCode) SetDefineInstruments()                                   { *v = 0x1127 }
func (v *ManufacturerCode) SetInHomeDisplays()                                      { *v = 0x1128 }
func (v *ManufacturerCode) SetMieleCieKg()                                          { *v = 0x1129 }
func (v *ManufacturerCode) SetTelevesSA()                                           { *v = 0x112a }
func (v *ManufacturerCode) SetLabelec()                                             { *v = 0x112b }
func (v *ManufacturerCode) SetChinaElectronicsStandardizationInstitute()            { *v = 0x112c }
func (v *ManufacturerCode) SetVectorform()                                          { *v = 0x112d }
func (v *ManufacturerCode) SetBuschJaegerElektro()                                  { *v = 0x112e }
func (v *ManufacturerCode) SetRedpineSignals()                                      { *v = 0x112f }
func (v *ManufacturerCode) SetBridgesElectronicTechnology()                         { *v = 0x1130 }
func (v *ManufacturerCode) SetSercomm()                                             { *v = 0x1131 }
func (v *ManufacturerCode) SetWshGmbhWirsindheller()                                { *v = 0x1132 }
func (v *ManufacturerCode) SetBoschSecuritySystems()                                { *v = 0x1133 }
func (v *ManufacturerCode) SetEzexCorporation()                                     { *v = 0x1134 }
func (v *ManufacturerCode) SetDresdenElektronikIngenieurtechnikGmbh()               { *v = 0x1135 }
func (v *ManufacturerCode) SetMeazonSA()                                            { *v = 0x1136 }
func (v *ManufacturerCode) SetCrowElectronicEngineering()                           { *v = 0x1137 }
func (v *ManufacturerCode) SetHarvardEngineering()                                  { *v = 0x1138 }
func (v *ManufacturerCode) SetAndsonBeijingTechnology()                             { *v = 0x1139 }
func (v *ManufacturerCode) SetAdhocoAg()                                            { *v = 0x113a }
func (v *ManufacturerCode) SetWaxmanConsumerProductsGroup()                         { *v = 0x113b }
func (v *ManufacturerCode) SetOwonTechnology()                                      { *v = 0x113c }
func (v *ManufacturerCode) SetHitronTechnologies()                                  { *v = 0x113d }
func (v *ManufacturerCode) SetScemtecSteuerungstechnikGmbh()                        { *v = 0x113e }
func (v *ManufacturerCode) SetWebee()                                               { *v = 0x113f }
func (v *ManufacturerCode) SetGrid2home()                                           { *v = 0x1140 }
func (v *ManufacturerCode) SetTelinkMicro()                                         { *v = 0x1141 }
func (v *ManufacturerCode) SetJasmineSystems()                                      { *v = 0x1142 }
func (v *ManufacturerCode) SetBidgely()                                             { *v = 0x1143 }
func (v *ManufacturerCode) SetLutron()                                              { *v = 0x1144 }
func (v *ManufacturerCode) SetIjenko()                                              { *v = 0x1145 }
func (v *ManufacturerCode) SetStarfieldElectronic()                                 { *v = 0x1146 }
func (v *ManufacturerCode) SetTcp()                                                 { *v = 0x1147 }
func (v *ManufacturerCode) SetRogersCommunicationsPartnership()                     { *v = 0x1148 }
func (v *ManufacturerCode) SetCree()                                                { *v = 0x1149 }
func (v *ManufacturerCode) SetRobertBoschLlc()                                      { *v = 0x114a }
func (v *ManufacturerCode) SetIbisNetworks()                                        { *v = 0x114b }
func (v *ManufacturerCode) SetQuirky()                                              { *v = 0x114c }
func (v *ManufacturerCode) SetEfergyTechnologies()                                  { *v = 0x114d }
func (v *ManufacturerCode) SetSmartlabs()                                           { *v = 0x114e }
func (v *ManufacturerCode) SetEverspringIndustry()                                  { *v = 0x114f }
func (v *ManufacturerCode) SetSwannCommunications()                                 { *v = 0x1150 }
func (v *ManufacturerCode) SetSoneter()                                             { *v = 0x1151 }
func (v *ManufacturerCode) SetSamsungSds()                                          { *v = 0x1152 }
func (v *ManufacturerCode) SetUnibandElectronicCorporation()                        { *v = 0x1153 }
func (v *ManufacturerCode) SetAcctonTechnologyCorporation()                         { *v = 0x1154 }
func (v *ManufacturerCode) SetBoschThermotechnikGmbh()                              { *v = 0x1155 }
func (v *ManufacturerCode) SetWincorNixdorfInc()                                    { *v = 0x1156 }
func (v *ManufacturerCode) SetOhsungElectronics()                                   { *v = 0x1157 }
func (v *ManufacturerCode) SetZenWithinInc()                                        { *v = 0x1158 }
func (v *ManufacturerCode) SetTech4homeLda()                                        { *v = 0x1159 }
func (v *ManufacturerCode) SetNanoleaf()                                            { *v = 0x115A }
func (v *ManufacturerCode) SetKeenHomeInc()                                         { *v = 0x115B }
func (v *ManufacturerCode) SetPolyControlAps()                                      { *v = 0x115C }
func (v *ManufacturerCode) SetEastfieldLightingCoLtdShenzhen()                      { *v = 0x115D }
func (v *ManufacturerCode) SetIpDatatelInc()                                        { *v = 0x115E }
func (v *ManufacturerCode) SetLumiUnitedTechologyLtdShenzhen()                      { *v = 0x115F }
func (v *ManufacturerCode) SetSengledOptoelectronicsCorp()                          { *v = 0x1160 }
func (v *ManufacturerCode) SetRemoteSolutionCoLtd()                                 { *v = 0x1161 }
func (v *ManufacturerCode) SetAbbGenwayXiamenElectricalEquipmentCoLtd()             { *v = 0x1162 }
func (v *ManufacturerCode) SetZhejiangRexenseTech()                                 { *v = 0x1163 }
func (v *ManufacturerCode) SetForeeTechnology()                                     { *v = 0x1164 }
func (v *ManufacturerCode) SetOpenAccessTechnologyIntl()                            { *v = 0x1165 }
func (v *ManufacturerCode) SetInnrLightingBv()                                      { *v = 0x1166 }
func (v *ManufacturerCode) SetTechworldIndustries()                                 { *v = 0x1167 }
func (v *ManufacturerCode) SetLeedarsonLightingCoLtd()                              { *v = 0x1168 }
func (v *ManufacturerCode) SetArzelZoning()                                         { *v = 0x1169 }
func (v *ManufacturerCode) SetHolleyTechnology()                                    { *v = 0x116A }
func (v *ManufacturerCode) SetBeldonTechnologies()                                  { *v = 0x116B }
func (v *ManufacturerCode) SetFlextronics()                                         { *v = 0x116C }
func (v *ManufacturerCode) SetShenzhenMeian()                                       { *v = 0x116D }
func (v *ManufacturerCode) SetLowes()                                               { *v = 0x116E }
func (v *ManufacturerCode) SetSigmaConnectivity()                                   { *v = 0x116F }
func (v *ManufacturerCode) SetWulian()                                              { *v = 0x1171 }
func (v *ManufacturerCode) SetPlugwiseBV()                                          { *v = 0x1172 }
func (v *ManufacturerCode) SetTitanProducts()                                       { *v = 0x1173 }
func (v *ManufacturerCode) SetEcospectral()                                         { *v = 0x1174 }
func (v *ManufacturerCode) SetDLink()                                               { *v = 0x1175 }
func (v *ManufacturerCode) SetTechnicolorHomeUsa()                                  { *v = 0x1176 }
func (v *ManufacturerCode) SetOppleLighting()                                       { *v = 0x1177 }
func (v *ManufacturerCode) SetWistronNewebCorp()                                    { *v = 0x1178 }
func (v *ManufacturerCode) SetQmotionShades()                                       { *v = 0x1179 }
func (v *ManufacturerCode) SetInstaElektroGmbh()                                    { *v = 0x117A }
func (v *ManufacturerCode) SetShanghaiVancount()                                    { *v = 0x117B }
func (v *ManufacturerCode) SetIkeaOfSweden()                                        { *v = 0x117C }
func (v *ManufacturerCode) SetRtRk()                                                { *v = 0x117D }
func (v *ManufacturerCode) SetShenzhenFeibit()                                      { *v = 0x117E }
func (v *ManufacturerCode) SetEucontrols()                                          { *v = 0x117F }
func (v *ManufacturerCode) SetTelkonet()                                            { *v = 0x1180 }
func (v *ManufacturerCode) SetThermalSolutionResources()                            { *v = 0x1181 }
func (v *ManufacturerCode) SetPomcube()                                             { *v = 0x1182 }
func (v *ManufacturerCode) SetEiElectronics()                                       { *v = 0x1183 }
func (v *ManufacturerCode) SetOptoga()                                              { *v = 0x1184 }
func (v *ManufacturerCode) SetStelpro()                                             { *v = 0x1185 }
func (v *ManufacturerCode) SetLynxusTechnologiesCorp()                              { *v = 0x1186 }
func (v *ManufacturerCode) SetSemiconductorComponents()                             { *v = 0x1187 }
func (v *ManufacturerCode) SetTpLink()                                              { *v = 0x1188 }
func (v *ManufacturerCode) SetLedvanceLlc()                                         { *v = 0x1189 }
func (v *ManufacturerCode) SetNortek()                                              { *v = 0x118A }
func (v *ManufacturerCode) SetIrevoAssaAbbloyKorea()                                { *v = 0x118B }
func (v *ManufacturerCode) SetMidea()                                               { *v = 0x118C }
func (v *ManufacturerCode) SetZfFriedrichshafen()                                   { *v = 0x118D }
func (v *ManufacturerCode) SetCheckit()                                             { *v = 0x118E }
func (v *ManufacturerCode) SetAclara()                                              { *v = 0x118F }
func (v *ManufacturerCode) SetNokia()                                               { *v = 0x1190 }
func (v *ManufacturerCode) SetGoldcardHighTechCoLtd()                               { *v = 0x1191 }
func (v *ManufacturerCode) SetGeorgeWilsonIndustriesLtd()                           { *v = 0x1192 }
func (v *ManufacturerCode) SetEasySaverCoInc()                                      { *v = 0x1193 }
func (v *ManufacturerCode) SetZteCorporation()                                      { *v = 0x1194 }
func (v *ManufacturerCode) SetArris()                                               { *v = 0x1195 }
func (v *ManufacturerCode) SetRelianceBigTv()                                       { *v = 0x1196 }
func (v *ManufacturerCode) SetInsightEnergyVenturesPowerley()                       { *v = 0x1197 }
func (v *ManufacturerCode) SetThomasResearchProductsHubbellLightingInc()            { *v = 0x1198 }
func (v *ManufacturerCode) SetLiSengTechnology()                                    { *v = 0x1199 }
func (v *ManufacturerCode) SetSystemLevelSolutionsInc()                             { *v = 0x119A }
func (v *ManufacturerCode) SetMatrixLabs()                                          { *v = 0x119B }
func (v *ManufacturerCode) SetSinopeTechnologies()                                  { *v = 0x119C }
func (v *ManufacturerCode) SetJiuzhouGreeble()                                      { *v = 0x119D }
func (v *ManufacturerCode) SetGuangzhouLanveeTechCoLtd()                            { *v = 0x119E }
func (v *ManufacturerCode) SetVenstar()                                             { *v = 0x119F }
func (v *ManufacturerCode) SetSlv()                                                 { *v = 0x1200 }
func (v *ManufacturerCode) SetHaloSmartLabs()                                       { *v = 0x1201 }
func (v *ManufacturerCode) SetScoutSecurityInc()                                    { *v = 0x1202 }
func (v *ManufacturerCode) SetAlibabaChinaInc()                                     { *v = 0x1203 }
func (v *ManufacturerCode) SetResolutionProductsInc()                               { *v = 0x1204 }
func (v *ManufacturerCode) SetSmartlokInc()                                         { *v = 0x1205 }
func (v *ManufacturerCode) SetLuxProductsCorp()                                     { *v = 0x1206 }
func (v *ManufacturerCode) SetVimarSpa()                                            { *v = 0x1207 }
func (v *ManufacturerCode) SetUniversalLightingTechnologies()                       { *v = 0x1208 }
func (v *ManufacturerCode) SetRobertBoschGmbh()                                     { *v = 0x1209 }
func (v *ManufacturerCode) SetAccenture()                                           { *v = 0x120A }
func (v *ManufacturerCode) SetHeimanTechnologyCoLtd()                               { *v = 0x120B }
func (v *ManufacturerCode) SetShenzhenHomaTechnologyCoLtd()                         { *v = 0x120C }
func (v *ManufacturerCode) SetVisionElectronicsTechnology()                         { *v = 0x120D }
func (v *ManufacturerCode) SetLenovo()                                              { *v = 0x120E }
func (v *ManufacturerCode) SetPrescienseRD()                                        { *v = 0x120F }
func (v *ManufacturerCode) SetShenzhenSeastarIntelligenceCoLtd()                    { *v = 0x1210 }
func (v *ManufacturerCode) SetSensativeAb()                                         { *v = 0x1211 }
func (v *ManufacturerCode) SetSolaredge()                                           { *v = 0x1212 }
func (v *ManufacturerCode) SetZipato()                                              { *v = 0x1213 }
func (v *ManufacturerCode) SetChinaFireSecuritySensingManufacturingIhorn()          { *v = 0x1214 }
func (v *ManufacturerCode) SetQubyBv()                                              { *v = 0x1215 }
func (v *ManufacturerCode) SetHangzhouRoombankerTechnologyCoLtd()                   { *v = 0x1216 }
func (v *ManufacturerCode) SetAmazonLab126()                                        { *v = 0x1217 }
func (v *ManufacturerCode) SetPaulmannLichtGmbh()                                   { *v = 0x1218 }
func (v *ManufacturerCode) SetShenzhenOrviboElectronicsCoLtd()                      { *v = 0x1219 }
func (v *ManufacturerCode) SetTciTelecommunications()                               { *v = 0x121A }
func (v *ManufacturerCode) SetMuellerLichtInternationalInc()                        { *v = 0x121B }
func (v *ManufacturerCode) SetAuroraLimited()                                       { *v = 0x121C }
func (v *ManufacturerCode) SetSmartdcc()                                            { *v = 0x121D }
func (v *ManufacturerCode) SetShanghaiUmeinfoCoLtd()                                { *v = 0x121E }
func (v *ManufacturerCode) SetCarbontrack()                                         { *v = 0x121F }
func (v *ManufacturerCode) SetSomfy()                                               { *v = 0x1220 }
func (v *ManufacturerCode) SetViessmannElektronikGmbh()                             { *v = 0x1221 }
func (v *ManufacturerCode) SetHildebrandTechnologyLtd()                             { *v = 0x1222 }
func (v *ManufacturerCode) SetOnkyoTechnologyCorporation()                          { *v = 0x1223 }
func (v *ManufacturerCode) SetShenzhenSunricherTechnologyLtd()                      { *v = 0x1224 }
func (v *ManufacturerCode) SetXiuXiuTechnologyCoLtd()                               { *v = 0x1225 }
func (v *ManufacturerCode) SetZumtobelGroup()                                       { *v = 0x1226 }
func (v *ManufacturerCode) SetShenzhenKaadasIntelligentTechnologyCoLtd()            { *v = 0x1227 }
func (v *ManufacturerCode) SetShanghaiXiaoyanTechnologyCoLtd()                      { *v = 0x1228 }
func (v *ManufacturerCode) SetCypressSemiconductor()                                { *v = 0x1229 }
func (v *ManufacturerCode) SetXalGmbh()                                             { *v = 0x122A }
func (v *ManufacturerCode) SetInergySystemsLlc()                                    { *v = 0x122B }
func (v *ManufacturerCode) SetAlfredKarcherGmbhCoKg()                               { *v = 0x122C }
func (v *ManufacturerCode) SetAdurolightManufacturing()                             { *v = 0x122D }
func (v *ManufacturerCode) SetGroupeMuller()                                        { *v = 0x122E }
func (v *ManufacturerCode) SetVMarkEnterprisesInc()                                 { *v = 0x122F }
func (v *ManufacturerCode) SetLeadEnergyAg()                                        { *v = 0x1230 }
func (v *ManufacturerCode) SetUiotGroup()                                           { *v = 0x1231 }
func (v *ManufacturerCode) SetAxxessIndustriesInc()                                 { *v = 0x1232 }
func (v *ManufacturerCode) SetThirdRealityInc()                                     { *v = 0x1233 }
func (v *ManufacturerCode) SetDsrCorporation()                                      { *v = 0x1234 }
func (v *ManufacturerCode) SetGuangzhouVensiIntelligentTechnologyCoLtd()            { *v = 0x1235 }
func (v *ManufacturerCode) SetSchlageLockAllegion()                                 { *v = 0x1236 }
func (v *ManufacturerCode) SetNet2grid()                                            { *v = 0x1237 }
func (v *ManufacturerCode) SetAiramElectricOyAb()                                   { *v = 0x1238 }
func (v *ManufacturerCode) SetImmaxWpbCz()                                          { *v = 0x1239 }
func (v *ManufacturerCode) SetZivAutomation()                                       { *v = 0x123A }
func (v *ManufacturerCode) SetHangzhouImagictechnologyCoLtd()                       { *v = 0x123B }
func (v *ManufacturerCode) SetXiamenLeelenTechnologyCoLtd()                         { *v = 0x123C }
func (v *ManufacturerCode) SetOverkizSas()                                          { *v = 0x123D }
func (v *ManufacturerCode) SetFlonidanAS()                                          { *v = 0x123E }
func (v *ManufacturerCode) SetHdlAutomationCoLtd()                                  { *v = 0x123F }
func (v *ManufacturerCode) SetArdomusNetworksCorporation()                          { *v = 0x1240 }
func (v *ManufacturerCode) SetSamjinCoLtd()                                         { *v = 0x1241 }
func (v *ManufacturerCode) SetSprueAegisPlc()                                       { *v = 0x1242 }
func (v *ManufacturerCode) SetIndraSistemasSA()                                     { *v = 0x1243 }
func (v *ManufacturerCode) SetShenzhenJbtSmartLightingCoLtd()                       { *v = 0x1244 }
func (v *ManufacturerCode) SetGeLightingCurrent()                                   { *v = 0x1245 }
func (v *ManufacturerCode) SetDanfossAS()                                           { *v = 0x1246 }
func (v *ManufacturerCode) SetNivissPhpSpZOOSpK()                                   { *v = 0x1247 }
func (v *ManufacturerCode) SetShenzhenFengliyuanEnergyConservatingTechnologyCoLtd() { *v = 0x1248 }
func (v *ManufacturerCode) SetNexelec()                                             { *v = 0x1249 }
func (v *ManufacturerCode) SetSichuanBehomeProminentTechnologyCoLtd()               { *v = 0x124A }
func (v *ManufacturerCode) SetFujianStarNetCommunicationCoLtd()                     { *v = 0x124B }
func (v *ManufacturerCode) SetToshibaVisualSolutionsCorporation()                   { *v = 0x124C }
func (v *ManufacturerCode) SetLatchableInc()                                        { *v = 0x124D }
func (v *ManufacturerCode) SetLSDeutschlandGmbh()                                   { *v = 0x124E }
func (v *ManufacturerCode) SetGledoptoCoLtd()                                       { *v = 0x124F }
func (v *ManufacturerCode) SetTheHomeDepot()                                        { *v = 0x1250 }
func (v *ManufacturerCode) SetNeonliteInternationalLtd()                            { *v = 0x1251 }
func (v *ManufacturerCode) SetArloTechnologiesInc()                                 { *v = 0x1252 }
func (v *ManufacturerCode) SetXingluoTechnologyCoLtd()                              { *v = 0x1253 }
func (v *ManufacturerCode) SetSimonElectricChinaCoLtd()                             { *v = 0x1254 }
func (v *ManufacturerCode) SetHangzhouGreatstarIndustrialCoLtd()                    { *v = 0x1255 }
func (v *ManufacturerCode) SetSequentricEnergySystemsLlc()                          { *v = 0x1256 }
func (v *ManufacturerCode) SetSolumCoLtd()                                          { *v = 0x1257 }
func (v *ManufacturerCode) SetEagleriseElectricElectronicChinaCoLtd()               { *v = 0x1258 }
func (v *ManufacturerCode) SetFantemTechnologiesShenzhenCoLtd()                     { *v = 0x1259 }
func (v *ManufacturerCode) SetYundingNetworkTechnologyBeijingCoLtd()                { *v = 0x125A }
func (v *ManufacturerCode) SetAtlanticGroup()                                       { *v = 0x125B }
func (v *ManufacturerCode) SetXiamenIntretechInc()                                  { *v = 0x125C }
func (v *ManufacturerCode) SetTuyaGlobalInc()                                       { *v = 0x125D }
func (v *ManufacturerCode) SetXiamenDnakeIntelligentTechnologyCoLtd()               { *v = 0x125E }
func (v *ManufacturerCode) SetNikoNv()                                              { *v = 0x125F }
func (v *ManufacturerCode) SetEmporiaEnergy()                                       { *v = 0x1260 }
func (v *ManufacturerCode) SetSikomAs()                                             { *v = 0x1261 }
func (v *ManufacturerCode) SetAxisLabsInc()                                         { *v = 0x1262 }
func (v *ManufacturerCode) SetCurrentProductsCorporation()                          { *v = 0x1263 }
func (v *ManufacturerCode) SetMetersitSrl()                                         { *v = 0x1264 }
func (v *ManufacturerCode) SetHornbachBaumarktAg()                                  { *v = 0x1265 }
func (v *ManufacturerCode) SetDiceworldSRLASocioUnico()                             { *v = 0x1266 }
func (v *ManufacturerCode) SetArcTechnologyCoLtd()                                  { *v = 0x1267 }
func (v *ManufacturerCode) SetHangzhouKonkeInformationTechnologyCoLtd()             { *v = 0x1268 }
func (v *ManufacturerCode) SetSaltoSystemsSL()                                      { *v = 0x1269 }
func (v *ManufacturerCode) SetShenzhenShyugjTechnologyCoLtd()                       { *v = 0x126A }
func (v *ManufacturerCode) SetBraydenAutomationCorporation()                        { *v = 0x126B }
func (v *ManufacturerCode) SetEnvironexusPtyLtd()                                   { *v = 0x126C }
func (v *ManufacturerCode) SetEltraNvSa()                                           { *v = 0x126D }
func (v *ManufacturerCode) SetXiaomiCommunicationsCoLtd()                           { *v = 0x126E }
func (v *ManufacturerCode) SetShanghaiShuncomElectronicTechnologyCoLtd()            { *v = 0x126F }
func (v *ManufacturerCode) SetVoltalisSA()                                          { *v = 0x1270 }
func (v *ManufacturerCode) SetFeeluxCoLtd()                                         { *v = 0x1271 }
func (v *ManufacturerCode) SetSmartplusInc()                                        { *v = 0x1272 }
func (v *ManufacturerCode) SetHalemeierGmbh()                                       { *v = 0x1273 }
func (v *ManufacturerCode) SetTrustInternationalBbv()                               { *v = 0x1274 }
func (v *ManufacturerCode) SetDukeEnergyBusinessServicesLlc()                       { *v = 0x1275 }
func (v *ManufacturerCode) SetCalixInc()                                            { *v = 0x1276 }
func (v *ManufacturerCode) SetAdeo()                                                { *v = 0x1277 }
func (v *ManufacturerCode) SetConnectedResponseLimited()                            { *v = 0x1278 }
func (v *ManufacturerCode) SetStroyenergokom()                                      { *v = 0x1279 }
func (v *ManufacturerCode) SetLumitechLightingSolutionGmbh()                        { *v = 0x127A }
func (v *ManufacturerCode) SetVerdantEnvironmentalTechnologies()                    { *v = 0x127B }
func (v *ManufacturerCode) SetAlfredInternational()                                 { *v = 0x127C }
func (v *ManufacturerCode) SetSansiLedLighting()                                    { *v = 0x127D }
func (v *ManufacturerCode) SetMindtree()                                            { *v = 0x127E }
func (v *ManufacturerCode) SetNordicSemiconductorAsa()                              { *v = 0x127F }
func (v *ManufacturerCode) SetSiterwellElectronics()                                { *v = 0x1280 }
func (v *ManufacturerCode) SetBrilonerLeuchtenGmbh()                                { *v = 0x1281 }
func (v *ManufacturerCode) SetShenzhenSeiTechnology()                               { *v = 0x1282 }
func (v *ManufacturerCode) SetCopperLabs()                                          { *v = 0x1283 }
func (v *ManufacturerCode) SetDeltaDore()                                           { *v = 0x1284 }
func (v *ManufacturerCode) SetHagerGroup()                                          { *v = 0x1285 }
func (v *ManufacturerCode) SetShenzhenCoolkitTechnology()                           { *v = 0x1286 }
func (v *ManufacturerCode) SetHangzhouSkyLighting()                                 { *v = 0x1287 }
func (v *ManufacturerCode) SetEOnSe()                                               { *v = 0x1288 }
func (v *ManufacturerCode) SetLidlStiftung()                                        { *v = 0x1289 }
func (v *ManufacturerCode) SetSichuanChanghongNetworkTechnologies()                 { *v = 0x128A }
func (v *ManufacturerCode) SetNodon()                                               { *v = 0x128B }
func (v *ManufacturerCode) SetJiangxiInnotechTechnology()                           { *v = 0x128C }
func (v *ManufacturerCode) SetMercatorPty()                                         { *v = 0x128D }
func (v *ManufacturerCode) SetBeijingRuyingTech()                                   { *v = 0x128E }
func (v *ManufacturerCode) SetEgloLeuchtenGmbh()                                    { *v = 0x128F }
func (v *ManufacturerCode) SetPietroFiorentiniSPA()                                 { *v = 0x1290 }
func (v *ManufacturerCode) SetZehnderGroupVauxAndigny()                             { *v = 0x1291 }
func (v *ManufacturerCode) SetBrkBrands()                                           { *v = 0x1292 }
func (v *ManufacturerCode) SetAskeyComputer()                                       { *v = 0x1293 }
func (v *ManufacturerCode) SetPassivebolt()                                         { *v = 0x1294 }
func (v *ManufacturerCode) SetAvmAudiovisuelles()                                   { *v = 0x1295 }
func (v *ManufacturerCode) SetNingboSuntechLightingTech()                           { *v = 0x1296 }
func (v *ManufacturerCode) SetSocieteEnCommanditeStello()                           { *v = 0x1297 }
func (v *ManufacturerCode) SetVivintSmartHome()                                     { *v = 0x1298 }
func (v *ManufacturerCode) SetNamron()                                              { *v = 0x1299 }
func (v *ManufacturerCode) SetRademacherGeraeteElektronikGmbh()                     { *v = 0x129A }
func (v *ManufacturerCode) SetOmoSystems()                                          { *v = 0x129B }
func (v *ManufacturerCode) SetSiglis()                                              { *v = 0x129C }
func (v *ManufacturerCode) SetImhotepCreation()                                     { *v = 0x129D }
func (v *ManufacturerCode) SetIcasa()                                               { *v = 0x129E }
func (v *ManufacturerCode) SetLevelHome()                                           { *v = 0x129F }
func (v *ManufacturerCode) SetTisControl()                                          { *v = 0x1300 }
func (v *ManufacturerCode) SetRadisysIndia()                                        { *v = 0x1301 }
func (v *ManufacturerCode) SetVeea()                                                { *v = 0x1302 }
func (v *ManufacturerCode) SetFellTechnology()                                      { *v = 0x1303 }
func (v *ManufacturerCode) SetSowiloDesignServices()                                { *v = 0x1304 }
func (v *ManufacturerCode) SetLexiDevices()                                         { *v = 0x1305 }
func (v *ManufacturerCode) SetLifiLabs()                                            { *v = 0x1306 }
func (v *ManufacturerCode) SetGrundfosHolding()                                     { *v = 0x1307 }
func (v *ManufacturerCode) SetSourcingCreation()                                    { *v = 0x1308 }
func (v *ManufacturerCode) SetKrakenTechnologies()                                  { *v = 0x1309 }
func (v *ManufacturerCode) SetEveSystems()                                          { *v = 0x130A }
func (v *ManufacturerCode) SetLiteOnTechnologyCorporation()                         { *v = 0x130B }
func (v *ManufacturerCode) SetFocalcrest()                                          { *v = 0x130C }
func (v *ManufacturerCode) SetBouffaloLabNanjing()                                  { *v = 0x130D }
func (v *ManufacturerCode) SetWyzeLabs()                                            { *v = 0x130E }
func (v *ManufacturerCode) SetDatekWirelessAs()                                     { *v = 0x1337 }
func (v *ManufacturerCode) SetGewissSPA()                                           { *v = 0x1994 }
func (v *ManufacturerCode) SetClimaxTechnologyCpLtd()                               { *v = 0x2794 }

func (ManufacturerCode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x1000, Name: "Cirronet"},
		{Value: 0x1001, Name: "Chipcon"},
		{Value: 0x1002, Name: "Ember"},
		{Value: 0x1003, Name: "National Tech"},
		{Value: 0x1004, Name: "Freescale"},
		{Value: 0x1005, Name: "IPCom"},
		{Value: 0x1006, Name: "San Juan Software"},
		{Value: 0x1007, Name: "TUV"},
		{Value: 0x1008, Name: "CompXs"},
		{Value: 0x1009, Name: "BM SpA"},
		{Value: 0x100a, Name: "AwarePoint"},
		{Value: 0x100b, Name: "Philips"},
		{Value: 0x100c, Name: "Luxoft"},
		{Value: 0x100d, Name: "Korvin"},
		{Value: 0x100e, Name: "One RF"},
		{Value: 0x100f, Name: "Software Technology Group"},
		{Value: 0x1010, Name: "Telegesis"},
		{Value: 0x1011, Name: "Visionic"},
		{Value: 0x1012, Name: "Insta"},
		{Value: 0x1013, Name: "Atalum"},
		{Value: 0x1014, Name: "Atmel"},
		{Value: 0x1015, Name: "Develco"},
		{Value: 0x1016, Name: "Honeywell"},
		{Value: 0x1017, Name: "RadioPulse"},
		{Value: 0x1018, Name: "Renesas"},
		{Value: 0x1019, Name: "Xanadu Wireless"},
		{Value: 0x101a, Name: "NEC Engineering"},
		{Value: 0x101b, Name: "Yamatake"},
		{Value: 0x101c, Name: "Tendril"},
		{Value: 0x101d, Name: "Assa Abloy"},
		{Value: 0x101e, Name: "Maxstream"},
		{Value: 0x101f, Name: "Neurocom"},
		{Value: 0x1020, Name: "Institute for Information Industry"},
		{Value: 0x1021, Name: "Vantage Controls"},
		{Value: 0x1022, Name: "iControl"},
		{Value: 0x1023, Name: "Raymarine"},
		{Value: 0x1024, Name: "LS Research"},
		{Value: 0x1025, Name: "Onity"},
		{Value: 0x1026, Name: "Mono Products"},
		{Value: 0x1027, Name: "RF Tech"},
		{Value: 0x1028, Name: "Itron"},
		{Value: 0x1029, Name: "Tritech"},
		{Value: 0x102a, Name: "Embedit"},
		{Value: 0x102b, Name: "S3C"},
		{Value: 0x102c, Name: "Siemens"},
		{Value: 0x102d, Name: "Mindtech"},
		{Value: 0x102e, Name: "LG Electronics"},
		{Value: 0x102f, Name: "Mitsubishi"},
		{Value: 0x1030, Name: "Johnson Controls"},
		{Value: 0x1031, Name: "PRI"},
		{Value: 0x1032, Name: "Knick"},
		{Value: 0x1033, Name: "Viconics"},
		{Value: 0x1034, Name: "Flexipanel"},
		{Value: 0x1035, Name: "Piasim Corporation"},
		{Value: 0x1036, Name: "Trane"},
		{Value: 0x1037, Name: "Jennic"},
		{Value: 0x1038, Name: "Living Independently"},
		{Value: 0x1039, Name: "AlertMe"},
		{Value: 0x103a, Name: "Daintree"},
		{Value: 0x103b, Name: "Aiji"},
		{Value: 0x103c, Name: "Telecom Italia"},
		{Value: 0x103d, Name: "Mikrokrets"},
		{Value: 0x103e, Name: "Oki Semi"},
		{Value: 0x103f, Name: "Newport Electronics"},
		{Value: 0x1040, Name: "Control4"},
		{Value: 0x1041, Name: "STMicro"},
		{Value: 0x1042, Name: "Ad-Sol Nissin"},
		{Value: 0x1043, Name: "DCSI"},
		{Value: 0x1044, Name: "France Telecom"},
		{Value: 0x1045, Name: "muNet"},
		{Value: 0x1046, Name: "Autani"},
		{Value: 0x1047, Name: "Colorado vNet"},
		{Value: 0x1048, Name: "Aerocomm"},
		{Value: 0x1049, Name: "Silicon Labs"},
		{Value: 0x104F, Name: "Crane"},
		{Value: 0x104a, Name: "Inncom"},
		{Value: 0x104b, Name: "Cannon"},
		{Value: 0x104c, Name: "Synapse"},
		{Value: 0x104d, Name: "Fisher Pierce/Sunrise"},
		{Value: 0x104e, Name: "CentraLite"},
		{Value: 0x1050, Name: "Mobilarm"},
		{Value: 0x1051, Name: "iMonitor"},
		{Value: 0x1052, Name: "Bartech"},
		{Value: 0x1053, Name: "Meshnetics"},
		{Value: 0x1054, Name: "LS Industrial"},
		{Value: 0x1055, Name: "Cason"},
		{Value: 0x1056, Name: "Wireless Glue"},
		{Value: 0x1057, Name: "Elster"},
		{Value: 0x1058, Name: "SMS Tec"},
		{Value: 0x1059, Name: "Onset Computer"},
		{Value: 0x105a, Name: "Riga Development"},
		{Value: 0x105b, Name: "Energate"},
		{Value: 0x105c, Name: "ConMed Linvatec"},
		{Value: 0x105d, Name: "PowerMand"},
		{Value: 0x105e, Name: "Schneider Electric"},
		{Value: 0x105f, Name: "Eaton"},
		{Value: 0x1060, Name: "Telular"},
		{Value: 0x1061, Name: "Delphi Medical"},
		{Value: 0x1062, Name: "EpiSensor"},
		{Value: 0x1063, Name: "Landis+Gyr"},
		{Value: 0x1064, Name: "Kaba Group"},
		{Value: 0x1065, Name: "Shure"},
		{Value: 0x1066, Name: "Comverge"},
		{Value: 0x1067, Name: "DBS Lodging"},
		{Value: 0x1068, Name: "Energy Aware"},
		{Value: 0x1069, Name: "Hidalgo"},
		{Value: 0x106a, Name: "Air2App"},
		{Value: 0x106b, Name: "AMX"},
		{Value: 0x106c, Name: "EDMI Pty"},
		{Value: 0x106d, Name: "Cyan Ltd"},
		{Value: 0x106e, Name: "System SPA"},
		{Value: 0x106f, Name: "Telit"},
		{Value: 0x1070, Name: "Kaga Electronics"},
		{Value: 0x1071, Name: "4-noks s.r.l."},
		{Value: 0x1072, Name: "Certicom"},
		{Value: 0x1073, Name: "Gridpoint"},
		{Value: 0x1074, Name: "Profile Systems"},
		{Value: 0x1075, Name: "Compacta International"},
		{Value: 0x1076, Name: "Freestyle Technology"},
		{Value: 0x1077, Name: "Alektrona"},
		{Value: 0x1078, Name: "Computime"},
		{Value: 0x1079, Name: "Remote Technologies"},
		{Value: 0x107a, Name: "Wavecom"},
		{Value: 0x107b, Name: "Energy Optimizers"},
		{Value: 0x107c, Name: "GE"},
		{Value: 0x107d, Name: "Jetlun"},
		{Value: 0x107e, Name: "Cipher Systems"},
		{Value: 0x107f, Name: "Corporate Systems Eng"},
		{Value: 0x1080, Name: "ecobee"},
		{Value: 0x1081, Name: "SMK"},
		{Value: 0x1082, Name: "Meshworks Wireless"},
		{Value: 0x1083, Name: "Ellips B.V."},
		{Value: 0x1084, Name: "Secure electrans"},
		{Value: 0x1085, Name: "CEDO"},
		{Value: 0x1086, Name: "Toshiba"},
		{Value: 0x1087, Name: "Digi International"},
		{Value: 0x1088, Name: "Ubilogix"},
		{Value: 0x1089, Name: "Echelon"},
		{Value: 0x1090, Name: "Green Energy Options"},
		{Value: 0x1091, Name: "Silver Spring Networks"},
		{Value: 0x1092, Name: "Black & Decker"},
		{Value: 0x1093, Name: "Aztech AssociatesInc."},
		{Value: 0x1094, Name: "A&D Co"},
		{Value: 0x1095, Name: "Rainforest Automation"},
		{Value: 0x1096, Name: "Carrier Electronics"},
		{Value: 0x1097, Name: "SyChip/Murata"},
		{Value: 0x1098, Name: "OpenPeak"},
		{Value: 0x1099, Name: "Passive Systems"},
		{Value: 0x109a, Name: "MMBResearch"},
		{Value: 0x109b, Name: "Leviton"},
		{Value: 0x109c, Name: "Korea Electric Power Data Network"},
		{Value: 0x109d, Name: "Comcast"},
		{Value: 0x109e, Name: "NEC Electronics"},
		{Value: 0x109f, Name: "Netvox"},
		{Value: 0x10a0, Name: "U-Control"},
		{Value: 0x10a1, Name: "Embedia Technologies"},
		{Value: 0x10a2, Name: "Sensus"},
		{Value: 0x10a3, Name: "SunriseTechnologies"},
		{Value: 0x10a4, Name: "MemtechCorp"},
		{Value: 0x10a5, Name: "Freebox"},
		{Value: 0x10a6, Name: "M2 Labs"},
		{Value: 0x10a7, Name: "BritishGas"},
		{Value: 0x10a8, Name: "Sentec"},
		{Value: 0x10a9, Name: "Navetas"},
		{Value: 0x10aa, Name: "Lightspeed Technologies"},
		{Value: 0x10ab, Name: "Oki Electric"},
		{Value: 0x10ac, Name: "Sistemas Inteligentes"},
		{Value: 0x10ad, Name: "Dometic"},
		{Value: 0x10ae, Name: "Alps"},
		{Value: 0x10af, Name: "EnergyHub"},
		{Value: 0x10b0, Name: "Kamstrup"},
		{Value: 0x10b1, Name: "EchoStar"},
		{Value: 0x10b2, Name: "EnerNOC"},
		{Value: 0x10b3, Name: "Eltav"},
		{Value: 0x10b4, Name: "Belkin"},
		{Value: 0x10b5, Name: "XStreamHD Wireless"},
		{Value: 0x10b6, Name: "Saturn South"},
		{Value: 0x10b7, Name: "GreenTrapOnline"},
		{Value: 0x10b8, Name: "SmartSynch"},
		{Value: 0x10b9, Name: "Nyce Control"},
		{Value: 0x10ba, Name: "ICM Controls"},
		{Value: 0x10bb, Name: "Millennium Electronics"},
		{Value: 0x10bc, Name: "Motorola"},
		{Value: 0x10bd, Name: "EmersonWhite-Rodgers"},
		{Value: 0x10be, Name: "Radio Thermostat"},
		{Value: 0x10bf, Name: "OMRONCorporation"},
		{Value: 0x10c0, Name: "GiiNii GlobalLimited"},
		{Value: 0x10c1, Name: "Fujitsu GeneralLimited"},
		{Value: 0x10c2, Name: "Peel Technologies"},
		{Value: 0x10c3, Name: "Accent"},
		{Value: 0x10c4, Name: "ByteSnap Design"},
		{Value: 0x10c5, Name: "NEC TOKIN Corporation"},
		{Value: 0x10c6, Name: "G4S JusticeServices"},
		{Value: 0x10c7, Name: "Trilliant Networks"},
		{Value: 0x10c8, Name: "Electrolux Italia"},
		{Value: 0x10c9, Name: "OnzoLtd"},
		{Value: 0x10ca, Name: "EnTekSystems"},
		{Value: 0x10cb, Name: "Philips 2"},
		{Value: 0x10cc, Name: "MainstreamEngineering"},
		{Value: 0x10cd, Name: "IndesitCompany"},
		{Value: 0x10ce, Name: "THINKECO"},
		{Value: 0x10cf, Name: "2D2C"},
		{Value: 0x10d0, Name: "GreenPeak"},
		{Value: 0x10d1, Name: "InterCEL"},
		{Value: 0x10d2, Name: "LG Electronics 2"},
		{Value: 0x10d3, Name: "Mitsumi Electric"},
		{Value: 0x10d4, Name: "Mitsumi Electric 2"},
		{Value: 0x10d5, Name: "Zentrum Mikroelektronik Dresden"},
		{Value: 0x10d6, Name: "Nest Labs"},
		{Value: 0x10d7, Name: "Exegin Technologies"},
		{Value: 0x10d8, Name: "Honeywell 2"},
		{Value: 0x10d9, Name: "Takahata Precision"},
		{Value: 0x10da, Name: "Sumitomo Electric Networks"},
		{Value: 0x10db, Name: "GE Energy"},
		{Value: 0x10dc, Name: "GE Appliances"},
		{Value: 0x10dd, Name: "Radiocrafts AS"},
		{Value: 0x10de, Name: "Ceiva"},
		{Value: 0x10df, Name: "TEC CO Co., Ltd"},
		{Value: 0x10e0, Name: "Chameleon Technology (UK) Ltd"},
		{Value: 0x10e1, Name: "Samsung"},
		{Value: 0x10e2, Name: "ruwido austria gmbh"},
		{Value: 0x10e3, Name: "Huawei Technologies Co., Ltd."},
		{Value: 0x10e4, Name: "Huawei Technologies Co., Ltd. 2"},
		{Value: 0x10e5, Name: "Greenwave Reality"},
		{Value: 0x10e6, Name: "BGlobal Metering Ltd"},
		{Value: 0x10e7, Name: "Mindteck"},
		{Value: 0x10e8, Name: "Ingersoll-Rand"},
		{Value: 0x10e9, Name: "Dius Computing Pty Ltd"},
		{Value: 0x10ea, Name: "Embedded Automation, Inc."},
		{Value: 0x10eb, Name: "ABB"},
		{Value: 0x10ec, Name: "Sony"},
		{Value: 0x10ed, Name: "Genus Power Infrastructures Limited"},
		{Value: 0x10ee, Name: "Universal Devices"},
		{Value: 0x10ef, Name: "Universal Devices 2"},
		{Value: 0x10f0, Name: "Metrum Technologies, LLC"},
		{Value: 0x10f1, Name: "Cisco"},
		{Value: 0x10f2, Name: "Ubisys technologies GmbH"},
		{Value: 0x10f3, Name: "Consert"},
		{Value: 0x10f4, Name: "Crestron Electronics"},
		{Value: 0x10f5, Name: "Enphase Energy"},
		{Value: 0x10f6, Name: "Invensys Controls"},
		{Value: 0x10f7, Name: "Mueller Systems, LLC"},
		{Value: 0x10f8, Name: "AAC Technologies Holding"},
		{Value: 0x10f9, Name: "U-NEXT Co., Ltd"},
		{Value: 0x10fa, Name: "Steelcase Inc."},
		{Value: 0x10fb, Name: "Telematics Wireless"},
		{Value: 0x10fc, Name: "Samil Power Co., Ltd"},
		{Value: 0x10fd, Name: "Pace Plc"},
		{Value: 0x10fe, Name: "Osborne Coinage Co."},
		{Value: 0x10ff, Name: "Powerwatch"},
		{Value: 0x1100, Name: "CANDELED GmbH"},
		{Value: 0x1101, Name: "FlexGrid S.R.L"},
		{Value: 0x1102, Name: "Humax"},
		{Value: 0x1103, Name: "Universal Electronics, Inc."},
		{Value: 0x1104, Name: "Advanced Energy"},
		{Value: 0x1105, Name: "BEGA Gantenbrink-Leuchten"},
		{Value: 0x1106, Name: "Brunel University"},
		{Value: 0x1107, Name: "Panasonic R&D Center Singapore"},
		{Value: 0x1108, Name: "eSystems Research"},
		{Value: 0x1109, Name: "Panamax"},
		{Value: 0x110a, Name: "Physical Graph Corporation"},
		{Value: 0x110b, Name: "EM-Lite Ltd."},
		{Value: 0x110c, Name: "Osram Sylvania"},
		{Value: 0x110d, Name: "2 Save Energy Ltd."},
		{Value: 0x110e, Name: "Planet Innovation Products Pty Ltd"},
		{Value: 0x110f, Name: "Ambient Devices, Inc."},
		{Value: 0x1110, Name: "Profalux"},
		{Value: 0x1111, Name: "Billion Electric Company (BEC)"},
		{Value: 0x1112, Name: "Embertec Pty Ltd"},
		{Value: 0x1113, Name: "IT Watchdogs"},
		{Value: 0x1114, Name: "Reloc"},
		{Value: 0x1115, Name: "Intel Corporation"},
		{Value: 0x1116, Name: "Trend Electronics Limited"},
		{Value: 0x1117, Name: "Moxa"},
		{Value: 0x1118, Name: "QEES"},
		{Value: 0x1119, Name: "SAYME Wireless Sensor Networks"},
		{Value: 0x111a, Name: "Pentair Aquatic Systems"},
		{Value: 0x111b, Name: "Orbit Irrigation"},
		{Value: 0x111c, Name: "California Eastern Laboratories"},
		{Value: 0x111d, Name: "Comcast 2"},
		{Value: 0x111e, Name: "IDT Technology Limited"},
		{Value: 0x111f, Name: "Pixela"},
		{Value: 0x1120, Name: "TiVo"},
		{Value: 0x1121, Name: "Fidure"},
		{Value: 0x1122, Name: "Marvell Semiconductor"},
		{Value: 0x1123, Name: "Wasion Group"},
		{Value: 0x1124, Name: "Jasco Products"},
		{Value: 0x1125, Name: "Shenzhen Kaifa Technology"},
		{Value: 0x1126, Name: "Netcomm Wireless"},
		{Value: 0x1127, Name: "Define Instruments"},
		{Value: 0x1128, Name: "In Home Displays"},
		{Value: 0x1129, Name: "Miele & Cie. KG"},
		{Value: 0x112a, Name: "Televes S.A."},
		{Value: 0x112b, Name: "Labelec"},
		{Value: 0x112c, Name: "China Electronics Standardization Institute"},
		{Value: 0x112d, Name: "Vectorform"},
		{Value: 0x112e, Name: "Busch-Jaeger Elektro"},
		{Value: 0x112f, Name: "Redpine Signals"},
		{Value: 0x1130, Name: "Bridges Electronic Technology"},
		{Value: 0x1131, Name: "Sercomm"},
		{Value: 0x1132, Name: "WSH GmbH wirsindheller"},
		{Value: 0x1133, Name: "Bosch Security Systems"},
		{Value: 0x1134, Name: "eZEX Corporation"},
		{Value: 0x1135, Name: "Dresden Elektronik Ingenieurtechnik GmbH"},
		{Value: 0x1136, Name: "MEAZON S.A."},
		{Value: 0x1137, Name: "Crow Electronic Engineering"},
		{Value: 0x1138, Name: "Harvard Engineering"},
		{Value: 0x1139, Name: "Andson(Beijing) Technology"},
		{Value: 0x113a, Name: "Adhoco AG"},
		{Value: 0x113b, Name: "Waxman Consumer Products Group"},
		{Value: 0x113c, Name: "Owon Technology"},
		{Value: 0x113d, Name: "Hitron Technologies"},
		{Value: 0x113e, Name: "Scemtec Steuerungstechnik GmbH"},
		{Value: 0x113f, Name: "Webee"},
		{Value: 0x1140, Name: "Grid2Home"},
		{Value: 0x1141, Name: "Telink Micro"},
		{Value: 0x1142, Name: "Jasmine Systems"},
		{Value: 0x1143, Name: "Bidgely"},
		{Value: 0x1144, Name: "Lutron"},
		{Value: 0x1145, Name: "IJENKO"},
		{Value: 0x1146, Name: "Starfield Electronic"},
		{Value: 0x1147, Name: "TCP"},
		{Value: 0x1148, Name: "Rogers Communications Partnership"},
		{Value: 0x1149, Name: "Cree"},
		{Value: 0x114a, Name: "Robert Bosch LLC"},
		{Value: 0x114b, Name: "Ibis Networks"},
		{Value: 0x114c, Name: "Quirky"},
		{Value: 0x114d, Name: "Efergy Technologies"},
		{Value: 0x114e, Name: "Smartlabs"},
		{Value: 0x114f, Name: "Everspring Industry"},
		{Value: 0x1150, Name: "Swann Communications"},
		{Value: 0x1151, Name: "Soneter"},
		{Value: 0x1152, Name: "Samsung SDS"},
		{Value: 0x1153, Name: "Uniband Electronic Corporation"},
		{Value: 0x1154, Name: "Accton Technology Corporation"},
		{Value: 0x1155, Name: "Bosch Thermotechnik GmbH"},
		{Value: 0x1156, Name: "Wincor Nixdorf Inc."},
		{Value: 0x1157, Name: "Ohsung Electronics"},
		{Value: 0x1158, Name: "Zen Within, Inc."},
		{Value: 0x1159, Name: "Tech4home, Lda."},
		{Value: 0x115A, Name: "Nanoleaf"},
		{Value: 0x115B, Name: "Keen Home, Inc."},
		{Value: 0x115C, Name: "Poly-Control APS"},
		{Value: 0x115D, Name: "Eastfield Lighting Co., Ltd Shenzhen"},
		{Value: 0x115E, Name: "IP Datatel, Inc."},
		{Value: 0x115F, Name: "Lumi United Techology, Ltd Shenzhen"},
		{Value: 0x1160, Name: "Sengled Optoelectronics Corp"},
		{Value: 0x1161, Name: "Remote Solution Co., Ltd."},
		{Value: 0x1162, Name: "ABB Genway Xiamen Electrical Equipment Co., Ltd."},
		{Value: 0x1163, Name: "Zhejiang Rexense Tech"},
		{Value: 0x1164, Name: "ForEE Technology"},
		{Value: 0x1165, Name: "Open Access Technology Intl."},
		{Value: 0x1166, Name: "INNR Lighting BV"},
		{Value: 0x1167, Name: "Techworld Industries"},
		{Value: 0x1168, Name: "Leedarson Lighting Co., Ltd."},
		{Value: 0x1169, Name: "Arzel Zoning"},
		{Value: 0x116A, Name: "Holley Technology"},
		{Value: 0x116B, Name: "Beldon Technologies"},
		{Value: 0x116C, Name: "Flextronics"},
		{Value: 0x116D, Name: "Shenzhen Meian"},
		{Value: 0x116E, Name: "Lowes"},
		{Value: 0x116F, Name: "Sigma Connectivity"},
		{Value: 0x1171, Name: "Wulian"},
		{Value: 0x1172, Name: "Plugwise B.V."},
		{Value: 0x1173, Name: "Titan Products"},
		{Value: 0x1174, Name: "Ecospectral"},
		{Value: 0x1175, Name: "D-Link"},
		{Value: 0x1176, Name: "Technicolor Home USA"},
		{Value: 0x1177, Name: "Opple Lighting"},
		{Value: 0x1178, Name: "Wistron NeWeb Corp."},
		{Value: 0x1179, Name: "QMotion Shades"},
		{Value: 0x117A, Name: "Insta Elektro GmbH"},
		{Value: 0x117B, Name: "Shanghai Vancount"},
		{Value: 0x117C, Name: "Ikea of Sweden"},
		{Value: 0x117D, Name: "RT-RK"},
		{Value: 0x117E, Name: "Shenzhen Feibit"},
		{Value: 0x117F, Name: "EuControls"},
		{Value: 0x1180, Name: "Telkonet"},
		{Value: 0x1181, Name: "Thermal Solution Resources"},
		{Value: 0x1182, Name: "PomCube"},
		{Value: 0x1183, Name: "Ei Electronics"},
		{Value: 0x1184, Name: "Optoga"},
		{Value: 0x1185, Name: "Stelpro"},
		{Value: 0x1186, Name: "Lynxus Technologies Corp."},
		{Value: 0x1187, Name: "Semiconductor Components"},
		{Value: 0x1188, Name: "TP-Link"},
		{Value: 0x1189, Name: "LEDVANCE LLC."},
		{Value: 0x118A, Name: "Nortek"},
		{Value: 0x118B, Name: "iRevo/Assa Abbloy Korea"},
		{Value: 0x118C, Name: "Midea"},
		{Value: 0x118D, Name: "ZF Friedrichshafen"},
		{Value: 0x118E, Name: "Checkit"},
		{Value: 0x118F, Name: "Aclara"},
		{Value: 0x1190, Name: "Nokia"},
		{Value: 0x1191, Name: "Goldcard High-tech Co., Ltd."},
		{Value: 0x1192, Name: "George Wilson Industries Ltd."},
		{Value: 0x1193, Name: "EASY SAVER CO.,INC"},
		{Value: 0x1194, Name: "ZTE Corporation"},
		{Value: 0x1195, Name: "ARRIS"},
		{Value: 0x1196, Name: "Reliance BIG TV"},
		{Value: 0x1197, Name: "Insight Energy Ventures/Powerley"},
		{Value: 0x1198, Name: "Thomas Research Products (Hubbell Lighting Inc.)"},
		{Value: 0x1199, Name: "Li Seng Technology"},
		{Value: 0x119A, Name: "System Level Solutions Inc."},
		{Value: 0x119B, Name: "Matrix Labs"},
		{Value: 0x119C, Name: "Sinope Technologies"},
		{Value: 0x119D, Name: "Jiuzhou Greeble"},
		{Value: 0x119E, Name: "Guangzhou Lanvee Tech. Co. Ltd."},
		{Value: 0x119F, Name: "Venstar"},
		{Value: 0x1200, Name: "SLV"},
		{Value: 0x1201, Name: "Halo Smart Labs"},
		{Value: 0x1202, Name: "Scout Security Inc."},
		{Value: 0x1203, Name: "Alibaba China Inc."},
		{Value: 0x1204, Name: "Resolution Products, Inc."},
		{Value: 0x1205, Name: "Smartlok Inc."},
		{Value: 0x1206, Name: "Lux Products Corp."},
		{Value: 0x1207, Name: "Vimar SpA"},
		{Value: 0x1208, Name: "Universal Lighting Technologies"},
		{Value: 0x1209, Name: "Robert Bosch, GmbH"},
		{Value: 0x120A, Name: "Accenture"},
		{Value: 0x120B, Name: "Heiman Technology Co., Ltd."},
		{Value: 0x120C, Name: "Shenzhen HOMA Technology Co., Ltd."},
		{Value: 0x120D, Name: "Vision-Electronics Technology"},
		{Value: 0x120E, Name: "Lenovo"},
		{Value: 0x120F, Name: "Presciense R&D"},
		{Value: 0x1210, Name: "Shenzhen Seastar Intelligence Co., Ltd."},
		{Value: 0x1211, Name: "Sensative AB"},
		{Value: 0x1212, Name: "SolarEdge"},
		{Value: 0x1213, Name: "Zipato"},
		{Value: 0x1214, Name: "China Fire & Security Sensing Manufacturing (iHorn)"},
		{Value: 0x1215, Name: "Quby BV"},
		{Value: 0x1216, Name: "Hangzhou Roombanker Technology Co., Ltd."},
		{Value: 0x1217, Name: "Amazon Lab126"},
		{Value: 0x1218, Name: "Paulmann Licht GmbH"},
		{Value: 0x1219, Name: "Shenzhen Orvibo Electronics Co. Ltd."},
		{Value: 0x121A, Name: "TCI Telecommunications"},
		{Value: 0x121B, Name: "Mueller-Licht International Inc."},
		{Value: 0x121C, Name: "Aurora Limited"},
		{Value: 0x121D, Name: "SmartDCC"},
		{Value: 0x121E, Name: "Shanghai UMEinfo Co. Ltd."},
		{Value: 0x121F, Name: "carbonTRACK"},
		{Value: 0x1220, Name: "Somfy"},
		{Value: 0x1221, Name: "Viessmann Elektronik GmbH"},
		{Value: 0x1222, Name: "Hildebrand Technology Ltd"},
		{Value: 0x1223, Name: "Onkyo Technology Corporation"},
		{Value: 0x1224, Name: "Shenzhen Sunricher Technology Ltd."},
		{Value: 0x1225, Name: "Xiu Xiu Technology Co., Ltd"},
		{Value: 0x1226, Name: "Zumtobel Group"},
		{Value: 0x1227, Name: "Shenzhen Kaadas Intelligent Technology Co. Ltd"},
		{Value: 0x1228, Name: "Shanghai Xiaoyan Technology Co. Ltd"},
		{Value: 0x1229, Name: "Cypress Semiconductor"},
		{Value: 0x122A, Name: "XAL GmbH"},
		{Value: 0x122B, Name: "Inergy Systems LLC"},
		{Value: 0x122C, Name: "Alfred Karcher GmbH & Co KG"},
		{Value: 0x122D, Name: "Adurolight Manufacturing"},
		{Value: 0x122E, Name: "Groupe Muller"},
		{Value: 0x122F, Name: "V-Mark Enterprises Inc."},
		{Value: 0x1230, Name: "Lead Energy AG"},
		{Value: 0x1231, Name: "UIOT Group"},
		{Value: 0x1232, Name: "Axxess Industries Inc."},
		{Value: 0x1233, Name: "Third Reality Inc."},
		{Value: 0x1234, Name: "DSR Corporation"},
		{Value: 0x1235, Name: "Guangzhou Vensi Intelligent Technology Co. Ltd."},
		{Value: 0x1236, Name: "Schlage Lock (Allegion)"},
		{Value: 0x1237, Name: "Net2Grid"},
		{Value: 0x1238, Name: "Airam Electric Oy Ab"},
		{Value: 0x1239, Name: "IMMAX WPB CZ"},
		{Value: 0x123A, Name: "ZIV Automation"},
		{Value: 0x123B, Name: "HangZhou iMagicTechnology Co., Ltd"},
		{Value: 0x123C, Name: "Xiamen Leelen Technology Co. Ltd."},
		{Value: 0x123D, Name: "Overkiz SAS"},
		{Value: 0x123E, Name: "Flonidan A/S"},
		{Value: 0x123F, Name: "HDL Automation Co., Ltd."},
		{Value: 0x1240, Name: "Ardomus Networks Corporation"},
		{Value: 0x1241, Name: "Samjin Co., Ltd."},
		{Value: 0x1242, Name: "Sprue Aegis PLC"},
		{Value: 0x1243, Name: "Indra Sistemas, S.A."},
		{Value: 0x1244, Name: "Shenzhen JBT Smart Lighting Co., Ltd."},
		{Value: 0x1245, Name: "GE Lighting & Current"},
		{Value: 0x1246, Name: "Danfoss A/S"},
		{Value: 0x1247, Name: "NIVISS PHP Sp. z o.o. Sp.k."},
		{Value: 0x1248, Name: "Shenzhen Fengliyuan Energy Conservating Technology Co. Ltd"},
		{Value: 0x1249, Name: "NEXELEC"},
		{Value: 0x124A, Name: "Sichuan Behome Prominent Technology Co., Ltd"},
		{Value: 0x124B, Name: "Fujian Star-net Communication Co., Ltd."},
		{Value: 0x124C, Name: "Toshiba Visual Solutions Corporation"},
		{Value: 0x124D, Name: "Latchable, Inc."},
		{Value: 0x124E, Name: "L&S Deutschland GmbH"},
		{Value: 0x124F, Name: "Gledopto Co., Ltd."},
		{Value: 0x1250, Name: "The Home Depot"},
		{Value: 0x1251, Name: "Neonlite International Ltd."},
		{Value: 0x1252, Name: "Arlo Technologies, Inc."},
		{Value: 0x1253, Name: "Xingluo Technology Co., Ltd."},
		{Value: 0x1254, Name: "Simon Electric (China) Co., Ltd."},
		{Value: 0x1255, Name: "Hangzhou Greatstar Industrial Co., Ltd."},
		{Value: 0x1256, Name: "Sequentric Energy Systems, LLC"},
		{Value: 0x1257, Name: "Solum Co., Ltd."},
		{Value: 0x1258, Name: "Eaglerise Electric & Electronic (China) Co., Ltd."},
		{Value: 0x1259, Name: "Fantem Technologies (Shenzhen) Co., Ltd."},
		{Value: 0x125A, Name: "Yunding Network Technology (Beijing) Co., Ltd."},
		{Value: 0x125B, Name: "Atlantic Group"},
		{Value: 0x125C, Name: "Xiamen Intretech, Inc."},
		{Value: 0x125D, Name: "Tuya Global Inc."},
		{Value: 0x125E, Name: "Xiamen Dnake Intelligent Technology Co., Ltd"},
		{Value: 0x125F, Name: "Niko nv"},
		{Value: 0x1260, Name: "Emporia Energy"},
		{Value: 0x1261, Name: "Sikom AS"},
		{Value: 0x1262, Name: "AXIS Labs, Inc."},
		{Value: 0x1263, Name: "Current Products Corporation"},
		{Value: 0x1264, Name: "MeteRSit SRL"},
		{Value: 0x1265, Name: "HORNBACH Baumarkt AG"},
		{Value: 0x1266, Name: "DiCEworld s.r.l. a socio unico"},
		{Value: 0x1267, Name: "ARC Technology Co., Ltd"},
		{Value: 0x1268, Name: "Hangzhou Konke Information Technology Co., Ltd."},
		{Value: 0x1269, Name: "SALTO Systems S.L."},
		{Value: 0x126A, Name: "Shenzhen Shyugj Technology Co., Ltd"},
		{Value: 0x126B, Name: "Brayden Automation Corporation"},
		{Value: 0x126C, Name: "Environexus Pty. Ltd."},
		{Value: 0x126D, Name: "Eltra nv/sa"},
		{Value: 0x126E, Name: "Xiaomi Communications Co., Ltd."},
		{Value: 0x126F, Name: "Shanghai Shuncom Electronic Technology Co., Ltd."},
		{Value: 0x1270, Name: "Voltalis S.A"},
		{Value: 0x1271, Name: "FEELUX Co., Ltd."},
		{Value: 0x1272, Name: "SmartPlus Inc."},
		{Value: 0x1273, Name: "Halemeier GmbH"},
		{Value: 0x1274, Name: "Trust International BBV"},
		{Value: 0x1275, Name: "Duke Energy Business Services LLC"},
		{Value: 0x1276, Name: "Calix, Inc."},
		{Value: 0x1277, Name: "ADEO"},
		{Value: 0x1278, Name: "Connected Response Limited"},
		{Value: 0x1279, Name: "StroyEnergoKom"},
		{Value: 0x127A, Name: "Lumitech Lighting Solution GmbH"},
		{Value: 0x127B, Name: "Verdant Environmental Technologies"},
		{Value: 0x127C, Name: "Alfred International"},
		{Value: 0x127D, Name: "Sansi LED Lighting"},
		{Value: 0x127E, Name: "Mindtree"},
		{Value: 0x127F, Name: "Nordic Semiconductor ASA"},
		{Value: 0x1280, Name: "Siterwell Electronics"},
		{Value: 0x1281, Name: "Briloner Leuchten GmbH"},
		{Value: 0x1282, Name: "Shenzhen SEI Technology"},
		{Value: 0x1283, Name: "Copper Labs"},
		{Value: 0x1284, Name: "Delta Dore"},
		{Value: 0x1285, Name: "Hager Group"},
		{Value: 0x1286, Name: "Shenzhen CoolKit Technology"},
		{Value: 0x1287, Name: "Hangzhou Sky-Lighting"},
		{Value: 0x1288, Name: "E.ON SE"},
		{Value: 0x1289, Name: "Lidl Stiftung"},
		{Value: 0x128A, Name: "Sichuan Changhong Network Technologies"},
		{Value: 0x128B, Name: "NodOn"},
		{Value: 0x128C, Name: "Jiangxi Innotech Technology"},
		{Value: 0x128D, Name: "Mercator Pty"},
		{Value: 0x128E, Name: "Beijing Ruying Tech"},
		{Value: 0x128F, Name: "EGLO Leuchten GmbH"},
		{Value: 0x1290, Name: "Pietro Fiorentini S.p.A"},
		{Value: 0x1291, Name: "Zehnder Group Vaux-Andigny"},
		{Value: 0x1292, Name: "BRK Brands"},
		{Value: 0x1293, Name: "Askey Computer"},
		{Value: 0x1294, Name: "PassiveBolt"},
		{Value: 0x1295, Name: "AVM Audiovisuelles"},
		{Value: 0x1296, Name: "Ningbo Suntech Lighting Tech"},
		{Value: 0x1297, Name: "Societe en Commandite Stello"},
		{Value: 0x1298, Name: "Vivint Smart Home"},
		{Value: 0x1299, Name: "Namron"},
		{Value: 0x129A, Name: "RADEMACHER Geraete Elektronik GmbH"},
		{Value: 0x129B, Name: "OMO Systems"},
		{Value: 0x129C, Name: "Siglis"},
		{Value: 0x129D, Name: "IMHOTEP CREATION"},
		{Value: 0x129E, Name: "icasa"},
		{Value: 0x129F, Name: "Level Home"},
		{Value: 0x1300, Name: "TIS Control"},
		{Value: 0x1301, Name: "Radisys India"},
		{Value: 0x1302, Name: "Veea"},
		{Value: 0x1303, Name: "FELL Technology"},
		{Value: 0x1304, Name: "Sowilo Design Services"},
		{Value: 0x1305, Name: "Lexi Devices"},
		{Value: 0x1306, Name: "Lifi Labs"},
		{Value: 0x1307, Name: "GRUNDFOS Holding"},
		{Value: 0x1308, Name: "SOURCING & CREATION"},
		{Value: 0x1309, Name: "Kraken Technologies"},
		{Value: 0x130A, Name: "EVE SYSTEMS"},
		{Value: 0x130B, Name: "LITE-ON TECHNOLOGY CORPORATION"},
		{Value: 0x130C, Name: "Focalcrest"},
		{Value: 0x130D, Name: "Bouffalo Lab (Nanjing)"},
		{Value: 0x130E, Name: "Wyze Labs"},
		{Value: 0x1337, Name: "Datek Wireless AS"},
		{Value: 0x1994, Name: "Gewiss S.p.A."},
		{Value: 0x2794, Name: "Climax Technology Cp., Ltd."},
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
	case 0xA1E0:
		return "Green Power"
	case 0xC05E:
		return "Light Link"
	}
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (v ProfileId) IsHomeAutomation() bool { return v == 0x0104 }
func (v ProfileId) IsGreenPower() bool     { return v == 0xA1E0 }
func (v ProfileId) IsLightLink() bool      { return v == 0xC05E }
func (v *ProfileId) SetHomeAutomation()    { *v = 0x0104 }
func (v *ProfileId) SetGreenPower()        { *v = 0xA1E0 }
func (v *ProfileId) SetLightLink()         { *v = 0xC05E }

func (ProfileId) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x0104, Name: "Home Automation"},
		{Value: 0xA1E0, Name: "Green Power"},
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
func (v *ServerMask) SetBackupTrustCenter(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
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

func (v UserDescriptorAvailable) Scaled() float64 {
	return float64(v)
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
