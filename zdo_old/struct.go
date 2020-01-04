package zdo_old

import (
	"errors"
	"fmt"
	"hemtjan.st/zcl"
)

type ResponseStatus zcl.Zu8

func (r ResponseStatus) MarshalZcl() ([]byte, error) { return zcl.Zu8(r).MarshalZcl() }
func (r *ResponseStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	if len(b) < 1 {
		return b, fmt.Errorf("responseStatus: expected 1 byte, got empty slice")
	}
	*r = ResponseStatus(b[0])
	return b[1:], nil
}

const (
	Success            ResponseStatus = 0x00 // The requested operation or transmission was completed successfully.
	InvalidRequestType ResponseStatus = 0x80 // The supplied request type was invalid.
	DeviceNotFound     ResponseStatus = 0x81 // The requested device did not exist on a device.
	InvalidEndpoint    ResponseStatus = 0x82 // The supplied endpoint was equal to = 0x00 or between 0xf1 and 0xff.
	NotActive          ResponseStatus = 0x83 // The requested endpoint is not described by a simple descriptor.
	NotSupported       ResponseStatus = 0x84 // The requested optional feature is not supported on the target device.
	Timeout            ResponseStatus = 0x85 // A timeout has occurred with the requested operation.
	NoMatch            ResponseStatus = 0x86 // The end device bind request was unsuccessful due to a failure to match any suitable clusters.
	NoEntry            ResponseStatus = 0x88 // The unbind request was unsuccessful due to the coordinator or source device not having an entry in its binding table to unbind.
	NoDescriptor       ResponseStatus = 0x89 // A child descriptor was not available following a discovery request to a parent.
	InsufficientSpace  ResponseStatus = 0x8a // The device does not have storage space to support the requested operation.
	NotPermitted       ResponseStatus = 0x8b // The device is not in the proper state to support the requested operation.
	TableFull          ResponseStatus = 0x8c // The device does not have table space to support the operation.
	NotAuthorized      ResponseStatus = 0x8d // The permissions configuration table on the target indicates that the request is not authorized from this device.
)

func (r ResponseStatus) String() string {
	switch r {
	case Success:
		return "Success"
	case InvalidRequestType:
		return "InvalidRequestType"
	case DeviceNotFound:
		return "DeviceNotFound"
	case InvalidEndpoint:
		return "InvalidEndpoint"
	case NotActive:
		return "NotActive"
	case NotSupported:
		return "NotSupported"
	case Timeout:
		return "Timeout"
	case NoMatch:
		return "NoMatch"
	case NoEntry:
		return "NoEntry"
	case NoDescriptor:
		return "NoDescriptor"
	case InsufficientSpace:
		return "InsufficientSpace"
	case NotPermitted:
		return "NotPermitted"
	case TableFull:
		return "TableFull"
	case NotAuthorized:
		return "NotAuthorized"
	default:
		return fmt.Sprintf("Unknown(%d)", uint8(r))
	}
}

type MultiAddress struct {
	Mode     zcl.Zu8
	NWK      zcl.Zu16
	IEEE     zcl.Zuid
	Endpoint zcl.Zu8
}

func (m MultiAddress) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = m.Mode.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if m.Mode != 3 {
		if tmp, err = m.NWK.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	} else {
		if tmp, err = m.IEEE.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	// 0x01 = Group (no endpoint)
	if m.Mode > 1 {
		if tmp, err = m.Endpoint.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (m *MultiAddress) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	if b, err = (&m.Mode).UnmarshalZcl(b); err != nil {
		return nil, err
	}

	if m.Mode != 3 {
		if b, err = (&m.NWK).UnmarshalZcl(b); err != nil {
			return nil, err
		}
	} else {
		if b, err = (&m.IEEE).UnmarshalZcl(b); err != nil {
			return nil, err
		}
	}
	if m.Mode > 1 {
		if b, err = (&m.Endpoint).UnmarshalZcl(b); err != nil {
			return nil, err
		}
	}

	return b, nil
}

func (m MultiAddress) String() string {
	switch m.Mode {
	case 1:
		return fmt.Sprintf("Group(%04X)", m.NWK)
	case 2:
		return fmt.Sprintf("%04X/%d", m.NWK, m.Endpoint)
	case 3:
		return fmt.Sprintf("%x/%d", m.IEEE, m.Endpoint)
	}
	return fmt.Sprintf("Addr[%d/%04x/%x/%d]", m.Mode, m.NWK, m.IEEE, m.Endpoint)
}

type Uint8List []zcl.Zu8

func (l Uint8List) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error
	for _, c := range l {
		if tmp, err = c.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (l *Uint8List) UnmarshalZcl(b []byte) ([]byte, error) {
	var n Uint8List
	var err error
	if len(b) < 1 {
		return nil, errors.New("too few bytes(LVuint8)")
	}
	for len(b) > 0 {
		c := new(zcl.Zu8)
		if b, err = (c).UnmarshalZcl(b); err != nil {
			return nil, err
		}
		n = append(n, *c)
	}
	*l = n
	return b, nil
}

type Uint16List []zcl.Zu16

func (l Uint16List) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error
	for _, c := range l {
		if tmp, err = c.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (l *Uint16List) UnmarshalZcl(b []byte) ([]byte, error) {
	var n Uint16List
	var err error
	if len(b) < 1 {
		return nil, errors.New("too few bytes(LVuint8)")
	}
	for len(b) > 0 {
		c := new(zcl.Zu16)
		if b, err = (c).UnmarshalZcl(b); err != nil {
			return nil, err
		}
		n = append(n, *c)
	}
	*l = n
	return b, nil
}

type LVuint8 []zcl.Zu8

func (l LVuint8) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error
	data = append(data, uint8(len(l)))
	for _, c := range l {
		if tmp, err = c.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (l *LVuint8) UnmarshalZcl(b []byte) ([]byte, error) {
	var n LVuint8
	var err error
	if len(b) < 1 {
		return nil, errors.New("too few bytes(LVuint8)")
	}
	inputLen := uint8(b[0])
	b = b[1:]
	for i := uint8(0); i < inputLen; i++ {
		c := new(zcl.Zu8)
		if b, err = (c).UnmarshalZcl(b); err != nil {
			return nil, err
		}
		n = append(n, *c)
	}
	*l = n
	return b, nil
}

type LVuint16 []zcl.Zu16

func (l LVuint16) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error
	data = append(data, uint8(len(l)))
	for _, c := range l {
		if tmp, err = c.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (l *LVuint16) UnmarshalZcl(b []byte) ([]byte, error) {
	var n LVuint16
	var err error
	if len(b) < 1 {
		return nil, errors.New("too few bytes(LVuint16)")
	}
	inputLen := uint8(b[0])
	b = b[1:]
	for i := uint8(0); i < inputLen; i++ {
		c := new(zcl.Zu16)
		if b, err = (c).UnmarshalZcl(b); err != nil {
			return nil, err
		}
		n = append(n, *c)
	}
	*l = n
	return b, nil
}

type SimpleDescriptor struct {
	Endpoint       zcl.Zu8
	Profile        zcl.Zu16
	DeviceType     zcl.Zu16
	DeviceVersion  zcl.Zu8
	InputClusters  LVuint16
	OutputClusters LVuint16
}

func (s SimpleDescriptor) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = s.Endpoint.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)
	if tmp, err = s.Profile.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)
	if tmp, err = s.DeviceType.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)
	if tmp, err = s.DeviceVersion.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)
	if tmp, err = s.InputClusters.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)
	if tmp, err = s.OutputClusters.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (s *SimpleDescriptor) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&s.Endpoint).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	if b, err = (&s.Profile).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	if b, err = (&s.DeviceType).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	if b, err = (&s.DeviceVersion).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	if b, err = (&s.InputClusters).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	if b, err = (&s.OutputClusters).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	return b, nil
}

type LVSimpleDescriptor SimpleDescriptor

func (s LVSimpleDescriptor) MarshalZcl() ([]byte, error) {
	b, err := SimpleDescriptor(s).MarshalZcl()
	if err != nil {
		return nil, err
	}
	return append([]byte{uint8(len(b))}, b...), nil
}

func (s *LVSimpleDescriptor) UnmarshalZcl(b []byte) ([]byte, error) {
	if len(b) < 1 {
		return nil, errors.New("simpleDescriptor: too few bytes")
	}
	ln := uint8(b[0])
	b = b[1:]
	if len(b) < int(ln) {
		return nil, fmt.Errorf("simpleDescriptor: too few bytes (expected %d, got %x)", ln, b)
	}
	b1 := b[:ln]
	b2 := b[ln:]
	desc := &SimpleDescriptor{}
	b3, err := desc.UnmarshalZcl(b1)
	*s = LVSimpleDescriptor(*desc)
	if err != nil {
		return nil, err
	}
	return append(b3, b2...), nil
}

type LVNeighbor []Neighbor

func (r LVNeighbor) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	for _, nb := range r {
		if tmp, err = nb.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return append([]byte{uint8(len(r))}, data...), nil
}

func (r *LVNeighbor) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	if len(b) < 1 {
		return nil, errors.New("too few bytes(LVRoute)")
	}
	n := LVNeighbor{}
	ln := uint8(b[0])
	for i := uint8(0); i < ln; i++ {
		nb := &Neighbor{}
		if b, err = nb.UnmarshalZcl(b); err != nil {
			return nil, err
		}
		n = append(n, *nb)
	}

	*r = n
	return b, nil
}

type Neighbor struct {
	PanID         zcl.Zuid
	IEEEAddr      zcl.Zuid
	NWKAddr       zcl.Zu16
	NeighborType  zcl.Zu8
	PermitJoining zcl.Zu8
	Depth         zcl.Zu8
	LQI           zcl.Zu8
}

func (s *Neighbor) Values() []zcl.Val {
	return []zcl.Val{
		&s.PanID,
		&s.IEEEAddr,
		&s.NWKAddr,
		&s.NeighborType,
		&s.PermitJoining,
		&s.Depth,
		&s.LQI,
	}
}

func (s Neighbor) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	for _, v := range (&s).Values() {
		if tmp, err = v.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)

	}
	return data, nil
}

func (s *Neighbor) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	for _, v := range s.Values() {
		if b, err = v.UnmarshalZcl(b); err != nil {
			return nil, err
		}
	}
	return b, nil
}

type LVRoute []Route

func (r LVRoute) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	for _, route := range r {
		if tmp, err = route.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return append([]byte{uint8(len(r))}, data...), nil
}

func (r *LVRoute) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	if len(b) < 1 {
		return nil, errors.New("too few bytes(LVRoute)")
	}
	n := LVRoute{}
	ln := uint8(b[0])
	for i := uint8(0); i < ln; i++ {
		route := &Route{}
		if b, err = route.UnmarshalZcl(b); err != nil {
			return nil, err
		}
		n = append(n, *route)
	}

	*r = n
	return b, nil
}

type Route struct {
	DstNWK      zcl.Zu16
	RouteStatus zcl.Zu8
	NextHop     zcl.Zu16
}

func (r Route) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	for _, v := range []zcl.Val{&r.DstNWK, &r.RouteStatus, &r.NextHop} {
		if tmp, err = v.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)

	}
	return data, nil
}

func (r *Route) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	for _, v := range []zcl.Val{&r.DstNWK, &r.RouteStatus, &r.NextHop} {
		if b, err = v.UnmarshalZcl(b); err != nil {
			return nil, err
		}
	}
	return b, nil
}
