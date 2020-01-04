package zcl

import (
	"errors"
	"fmt"
)

type ProfileID Zu16
type CommandID Zu8
type ZdoCmdID Zu16
type TypeID Zu8

func (i ProfileID) MarshalZcl() ([]byte, error) { return Zu16(i).MarshalZcl() }
func (i CommandID) MarshalZcl() ([]byte, error) { return Zu8(i).MarshalZcl() }
func (i ZdoCmdID) MarshalZcl() ([]byte, error)  { return Zu16(i).MarshalZcl() }
func (i TypeID) MarshalZcl() ([]byte, error)    { return Zu8(i).MarshalZcl() }
func (i TypeID) New(val ...interface{}) Val     { return NewValue(uint8(i), val...) }

const (
	ErrNotEnoughData      errType = "not enough data"
	ErrTooMuchData        errType = "too much data"
	ErrNotImpl            errType = "not implemented"
	ErrInvalidArrayLength errType = "invalid length for array"
	ErrInvalidType        errType = "invalid type"
	ErrTimeout            errType = "timeout"
)

type Frame struct {
	ExpectReply bool
	IsReply     bool
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

func (f *Frame) MarshalZcl() ([]byte, error) {
	data := []byte{f.Type & 0x03}
	if f.IsReply {
		data[0] = data[0] | 0x08
	}
	if len(f.MnfCode) == 2 {
		data[0] = data[0] | 0x04
		data = append(data, f.MnfCode[1], f.MnfCode[0])
	}
	data = append(data, uint8(f.Seq), uint8(f.CommandID))

	if f.Payload != nil {
		payload, err := f.Payload.MarshalZcl()
		if err != nil {
			return nil, err
		}
		data = append(data, payload...)
	}
	return data, nil
}

func (f *Frame) UnmarshalZcl(b []byte) ([]byte, error) {
	if len(b) < 3 {
		return nil, errors.New("too few bytes(frame:0)")
	}
	st := b[0]
	b = b[1:]
	f.Type = st & 0x03
	f.IsReply = st&0x08 == 0x08
	if st&0x04 == 0x04 {
		if len(b) < 4 {
			return nil, errors.New("too few bytes(frame:1)")
		}
		f.MnfCode = []byte{b[1], b[0]}
		b = b[2:]
	}
	f.Seq = b[0]
	f.CommandID = CommandID(b[1])
	return b[2:], nil
}

type Val interface {
	MarshalZcl() ([]byte, error)
	UnmarshalZcl([]byte) ([]byte, error)
}

type Option struct {
	Name  string
	Value int
}
type EnumAttr interface {
	Attr
	SingleOptions() []Option
}
type BitmapAttr interface {
	Attr
	MultiOptions() []Option
}

type Command interface {
	General
	Arguments() []Argument
	Cluster() ClusterID
	Required() bool
	MnfCode() []byte
}

type General interface {
	Val
	Values() []Val
	ID() CommandID
	Name() string
	String() string
}

type ZdoCommand interface {
	Cluster() uint16
	Values() []Val
}

/*type ZdoCommand interface {
	Val
	Values() []Val
	ID() ZdoCmdID
	Name() string
	String() string
}*/

type errType string

func (i *AttrID) UnmarshalZcl(b []byte) ([]byte, error) {
	v := new(Zu16)
	b, err := v.UnmarshalZcl(b)
	*i = AttrID(*v)
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

func (i *ZdoCmdID) UnmarshalZcl(b []byte) ([]byte, error) {
	v := new(Zu16)
	b, err := v.UnmarshalZcl(b)
	*i = ZdoCmdID(*v)
	return b, err
}

func (i *TypeID) UnmarshalZcl(b []byte) ([]byte, error) {
	v := new(Zu8)
	b, err := v.UnmarshalZcl(b)
	*i = TypeID(*v)
	return b, err
}

func (e errType) Error() string { return string(e) }

func NewValue(dataType uint8, val ...interface{}) Val {
	switch dataType {
	case 8:
		v := new(Zdat8)
		if len(val) == 1 {
			switch val[0].(type) {
			case string:
				vv := val[0].(string)
				if len(vv) > 0 {
					v[0] = vv[0]
				}
			}
		}
		return v
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

func (i TypeID) Analog() bool {
	switch i {

	case 32, 33, 34, 35, 36, 37, 38, 39:
		// unsigned ints
		return true

	case 40, 41, 42, 43, 44, 45, 46, 47:
		// signed ints
		return true

	case 56, 57, 58:
		// floats
		return true

	case 224, 225, 226:
		// time & date
		return true

	default:
		return false
	}
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
		return "Octet string"
	case 66:
		return "Character string"
	case 67:
		return "Long octet string"
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
