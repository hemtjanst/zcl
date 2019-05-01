package zcl

import (
	"fmt"
)

type ClusterID Zu16
type ProfileID Zu16
type CommandID Zu8
type TypeID Zu8
type AttrID Zu16

func (i AttrID) MarshalZcl() ([]byte, error)    { return Zu16(i).MarshalZcl() }
func (i ClusterID) MarshalZcl() ([]byte, error) { return Zu16(i).MarshalZcl() }
func (i ProfileID) MarshalZcl() ([]byte, error) { return Zu16(i).MarshalZcl() }
func (i CommandID) MarshalZcl() ([]byte, error) { return Zu8(i).MarshalZcl() }
func (i TypeID) MarshalZcl() ([]byte, error)    { return Zu8(i).MarshalZcl() }

type AttrDef interface {
	ID() AttrID
	Value() Val
}

type Attribute struct {
	ID    AttrID
	Value Val
}

const (
	ReadAttr  CommandID = 0
	WriteAttr CommandID = 2

	ErrNotEnoughData errType = "not enough data"
	ErrTooMuchData   errType = "too much data"
	ErrNotImpl       errType = "not implemented"
)

type Frame struct {
	ExpectReply bool
	Profile     ProfileID
	Cluster     ClusterID
	MnfCode     []byte
	Type        uint8
	Seq         uint8
	CommandID   CommandID
	Payload     Val
}

type RequestFn func(fr *Frame) (*Frame, error)

func (f *Frame) Check() error {
	return nil
}

type Val interface {
	MarshalZcl() ([]byte, error)
	UnmarshalZcl([]byte) ([]byte, error)
}

type Command interface {
	Val
	ID() CommandID
	Cluster() ClusterID
	MnfCode() []byte
}

type Cluster interface {
}

type errType string

func (i *AttrID) UnmarshalZcl(b []byte) ([]byte, error) {
	v := new(Zu16)
	b, err := v.UnmarshalZcl(b)
	*i = AttrID(*v)
	return b, err
}

func (i *ClusterID) UnmarshalZcl(b []byte) ([]byte, error) {
	v := new(Zu16)
	b, err := v.UnmarshalZcl(b)
	*i = ClusterID(*v)
	return b, err
}

func (i *ProfileID) UnmarshalZcl(b []byte) ([]byte, error) {
	v := new(Zu16)
	b, err := v.UnmarshalZcl(b)
	*i = ProfileID(*v)
	return b, err
}

func (i *CommandID) UnmarshalZcl(b []byte) ([]byte, error) {
	v := new(Zu8)
	b, err := v.UnmarshalZcl(b)
	*i = CommandID(*v)
	return b, err
}

func (i *TypeID) UnmarshalZcl(b []byte) ([]byte, error) {
	v := new(Zu8)
	b, err := v.UnmarshalZcl(b)
	*i = TypeID(*v)
	return b, err
}

func (a Attribute) MarshalZcl() ([]byte, error) {
	b1, err := a.ID.MarshalZcl()
	if err != nil {
		return nil, err
	}
	b2, err := a.Value.MarshalZcl()
	if err != nil {
		return nil, err
	}

	return append(b1, b2...), nil
}
func (a *Attribute) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	b, err = (&a.ID).UnmarshalZcl(b)
	if err != nil {
		return b, err
	}
	b, err = a.Value.UnmarshalZcl(b)
	if err != nil {
		return b, err
	}
	return b, nil

}
func (a *Attribute) Values() []Val { return []Val{&a.ID, a.Value} }

func (e errType) Error() string { return string(e) }

func NewValue(dataType uint8) Val {
	switch dataType {
	case 8:
		return new(Zdat8)
	case 9:
		return new(Zdat16)
	case 10:
		return new(Zdat24)
	case 11:
		return new(Zdat32)
	case 12:
		return new(Zdat40)
	case 13:
		return new(Zdat48)
	case 14:
		return new(Zdat56)
	case 15:
		return new(Zdat64)
	case 16:
		return new(Zbool)
	case 24:
		return new(Zbmp8)
	case 25:
		return new(Zbmp16)
	case 26:
		return new(Zbmp24)
	case 27:
		return new(Zbmp32)
	case 28:
		return new(Zbmp40)
	case 29:
		return new(Zbmp48)
	case 30:
		return new(Zbmp56)
	case 31:
		return new(Zbmp64)
	case 32:
		return new(Zu8)
	case 33:
		return new(Zu16)
	case 34:
		return new(Zu24)
	case 35:
		return new(Zu32)
	case 36:
		return new(Zu40)
	case 37:
		return new(Zu48)
	case 38:
		return new(Zu56)
	case 39:
		return new(Zu64)
	case 40:
		return new(Zs8)
	case 41:
		return new(Zs16)
	case 42:
		return new(Zs24)
	case 43:
		return new(Zs32)
	case 44:
		return new(Zs40)
	case 45:
		return new(Zs48)
	case 46:
		return new(Zs56)
	case 47:
		return new(Zs64)
	case 48:
		return new(Zenum8)
	case 49:
		return new(Zenum16)
	case 56:
		return new(Zsemi)
	case 57:
		return new(Zfloat)
	case 58:
		return new(Zdouble)
	case 65:
		return new(Zostring)
	case 66:
		return new(Zcstring)
	case 67:
		return new(Zlostring)
	case 68:
		return new(Zlcstring)
	case 72:
		return new(Zarray)
	case 76:
		return new(Zstruct)
	case 80:
		return new(Zset)
	case 81:
		return new(Zbag)
	case 224:
		return new(Ztime)
	case 225:
		return new(Zdate)
	case 226:
		return new(Zutc)
	case 232:
		return new(Zcid)
	case 233:
		return new(Zaid)
	case 234:
		return new(Zoid)
	case 240:
		return new(Zuid)
	case 241:
		return new(Zseckey)
	}
	return nil
}

func (i TypeID) String() string {
	switch i {
	case 0:
		return "No data"
	case 8:
		return "8-bit data"
	case 9:
		return "16-bit data"
	case 10:
		return "24-bit data"
	case 11:
		return "32-bit data"
	case 12:
		return "40-bit data"
	case 13:
		return "48-bit data"
	case 14:
		return "56-bit data"
	case 15:
		return "64-bit data"
	case 16:
		return "Boolean"
	case 24:
		return "8-bit bitmap"
	case 25:
		return "16-bit bitmap"
	case 26:
		return "24-bit bitmap"
	case 27:
		return "32-bit bitmap"
	case 28:
		return "40-bit bitmap"
	case 29:
		return "48-bit bitmap"
	case 30:
		return "56-bit bitmap"
	case 31:
		return "64-bit bitmap"
	case 32:
		return "Unsigned 8-bit integer"
	case 33:
		return "Unsigned 16-bit integer"
	case 34:
		return "Unsigned 24-bit integer"
	case 35:
		return "Unsigned 32-bit integer"
	case 36:
		return "Unsigned 40-bit integer"
	case 37:
		return "Unsigned 48-bit integer"
	case 38:
		return "Unsigned 56-bit integer"
	case 39:
		return "Unsigned 64-bit integer"
	case 40:
		return "Signed 8-bit integer"
	case 41:
		return "Signed 16-bit integer"
	case 42:
		return "Signed 24-bit integer"
	case 43:
		return "Signed 32-bit integer"
	case 44:
		return "Signed 40-bit integer"
	case 45:
		return "Signed 48-bit integer"
	case 46:
		return "Signed 56-bit integer"
	case 47:
		return "Signed 64-bit integer"
	case 48:
		return "8-bit enumeration"
	case 49:
		return "16-bit enumeration"
	case 56:
		return "Semi-precicion"
	case 57:
		return "Single precicion"
	case 58:
		return "Double precicion"
	case 65:
		return "Octed string"
	case 66:
		return "Character string"
	case 67:
		return "Long octed string"
	case 68:
		return "Long character string"
	case 72:
		return "Array"
	case 76:
		return "Structure"
	case 80:
		return "Set"
	case 81:
		return "Bag"
	case 224:
		return "Time of day"
	case 225:
		return "Date"
	case 226:
		return "UTCTime"
	case 232:
		return "Cluster ID"
	case 233:
		return "Attribute ID"
	case 234:
		return "BACnet OID"
	case 240:
		return "IEEE address"
	case 241:
		return "128-bit security key"
	}
	return fmt.Sprintf("Unknown(0x%02x)", uint8(i))

}
