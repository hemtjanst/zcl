package zcl

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Zbool uint8

func (t *Zbool) UnmarshalZcl(buf []byte) ([]byte, error) {
	if len(buf) < 1 {
		return nil, ErrNotEnoughData
	}
	*t = Zbool(buf[0])
	return buf[1:], nil
}
func (t *Zbool) Values() []Val              { return []Val{t} }
func (t Zbool) MarshalZcl() ([]byte, error) { return []byte{byte(t)}, nil }
func (t Zbool) MarshalJSON() ([]byte, error) {
	if t == 0 {
		return []byte("false"), nil
	} else {
		return []byte("true"), nil
	}
}
func (t *Zbool) UnmarshalJSON(b []byte) error {
	if string(b) == "true" {
		*t = 1
	} else if string(b) == "false" {
		*t = 0
	} else {
		nv := new(uint8)
		if err := json.Unmarshal(b, nv); err != nil {
			return err
		}
		*t = Zbool(*nv)
	}
	return nil
}

func (t Zbool) TypeID() TypeID { return 16 }
func (t Zbool) Valid() bool    { return t < 2 }
func (t Zbool) String() string {
	if t == 1 {
		return "true"
	}
	if t == 0 {
		return "false"
	}
	return fmt.Sprintf("Bool(0x%02x)", uint8(t))
}

// Zcid is Cluster ID (0xffff = invalid). A/D = D
type Zcid uint16

func (c *Zcid) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(2, buf)
	*c = Zcid(val)
	return buf, err
}
func (c Zcid) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(2, uint64(c)) }
func (c *Zcid) Values() []Val              { return []Val{c} }
func (c Zcid) TypeID() TypeID              { return 232 }
func (c Zcid) String() string              { return fmt.Sprintf("0x%04x", uint16(c)) }
func (c Zcid) Valid() bool                 { return c != Zcid(0xFFFF) }

// Zaid is Attribute ID (0xffff = invalid). A/D = D
type Zaid uint16

func (a *Zaid) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(2, buf)
	*a = Zaid(val)
	return buf, err
}
func (a Zaid) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(2, uint64(a)) }
func (a *Zaid) Values() []Val              { return []Val{a} }
func (a Zaid) TypeID() TypeID              { return 233 }
func (a Zaid) String() string              { return fmt.Sprintf("0x%04x", uint16(a)) }
func (a Zaid) Valid() bool                 { return a != Zaid(0xFFFF) }

// Zoid is BACnet OID (0xffffffff = invalid). A/D = D
type Zoid uint32

func (o *Zoid) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(4, buf)
	*o = Zoid(val)
	return buf, err
}
func (o Zoid) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(4, uint64(o)) }
func (o *Zoid) Values() []Val              { return []Val{o} }
func (o Zoid) TypeID() TypeID              { return 234 }
func (o Zoid) Valid() bool                 { return o != Zoid(0xFFFFFFFF) }
func (o Zoid) String() string {
	ot := (o & 0xFFC00000) >> 22
	oi := o & 0x003FFFFF
	return fmt.Sprintf("%d:%d", ot, oi)
}

// Zuid is IEEE address (0xffffffffffffffff = invalid). A/D = D
type Zuid uint64

func (u *Zuid) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(8, buf)
	*u = Zuid(val)
	return buf, err
}
func (u Zuid) MarshalZcl() ([]byte, error)  { return uintLEMarshalZcl(8, uint64(u)) }
func (u Zuid) MarshalJSON() ([]byte, error) { return json.Marshal(u.String()) }
func (u *Zuid) UnmarshalJSON(b []byte) error {
	s := new(string)
	if err := json.Unmarshal(b, s); err != nil {
		return err
	}
	mac, err := net.ParseMAC(*s)
	if err != nil {
		return err
	}

	*u = Zuid(binary.BigEndian.Uint64(mac))
	return nil
}
func (u *Zuid) Values() []Val { return []Val{u} }
func (u Zuid) TypeID() TypeID { return 240 }
func (u Zuid) Valid() bool    { return u != Zuid(0xFFFFFFFFFFFFFFFF) }
func (u Zuid) String() string {
	return u.HWAddr().String()
}
func (u Zuid) HWAddr() net.HardwareAddr {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(u))
	return b
}

// Zseckey is 128-bit security key. A/D = D
type Zseckey []byte

func (s *Zseckey) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(16, buf)
	*s = Zseckey(val)
	return buf, err
}
func (s Zseckey) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(16, []byte(s)) }
func (s *Zseckey) Values() []Val              { return []Val{s} }
func (s Zseckey) TypeID() TypeID              { return 241 }
func (s Zseckey) String() string {
	return net.HardwareAddr(s).String()
}
