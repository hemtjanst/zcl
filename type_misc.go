package zcl

import (
	"fmt"
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
func (t Zbool) ID() TypeID                  { return 16 }

// Zcid is Cluster ID (0xffff = invalid). A/D = D
type Zcid uint16

func (c *Zcid) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(2, buf)
	*c = Zcid(val)
	return buf, err
}
func (c Zcid) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(2, uint64(c)) }
func (c *Zcid) Values() []Val              { return []Val{c} }
func (c Zcid) ID() TypeID                  { return 232 }
func (c Zcid) String() string              { return fmt.Sprintf("0x%04x", uint16(c)) }

//func (c Zcid) Valid() bool { return c != Zcid(65535) }

// Zaid is Attribute ID (0xffff = invalid). A/D = D
type Zaid uint16

func (a *Zaid) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(2, buf)
	*a = Zaid(val)
	return buf, err
}
func (a Zaid) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(2, uint64(a)) }
func (a *Zaid) Values() []Val              { return []Val{a} }
func (a *Zaid) ID() TypeID                 { return 233 }
func (c Zaid) String() string              { return fmt.Sprintf("0x%04x", uint16(c)) }

//func (a Zaid) Valid() bool { return a != Zaid(65535) }

// Zoid is BACnet OID (0xffffffff = invalid). A/D = D
type Zoid uint32

func (o *Zoid) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(4, buf)
	*o = Zoid(val)
	return buf, err
}
func (o Zoid) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(4, uint64(o)) }
func (o *Zoid) Values() []Val              { return []Val{o} }
func (o *Zoid) ID() TypeID                 { return 234 }

//func (o Zoid) Valid() bool { return o != Zoid(4294967295) }

// Zuid is IEEE address (0xffffffffffffffff = invalid). A/D = D
type Zuid uint64

func (u *Zuid) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(8, buf)
	*u = Zuid(val)
	return buf, err
}
func (u Zuid) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(8, uint64(u)) }
func (u *Zuid) Values() []Val              { return []Val{u} }
func (u *Zuid) ID() TypeID                 { return 240 }

//func (u Zuid) Valid() bool { return u != Zuid(18446744073709551615) }

// Zseckey is 128-bit security key. A/D = D
type Zseckey []byte

func (s *Zseckey) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(16, buf)
	*s = Zseckey(val)
	return buf, err
}
func (s Zseckey) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(16, []byte(s)) }
func (s *Zseckey) Values() []Val              { return []Val{s} }
func (s *Zseckey) ID() TypeID                 { return 241 }
