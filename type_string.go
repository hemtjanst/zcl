package zcl

import "encoding/binary"

func stringMarshalZcl(ln, b string) ([]byte, error) {
	var buf []byte
	if ln == "o1" {
		bln := len(b)
		if bln > 255 {
			return nil, ErrTooMuchData
		}
		buf = append(buf, uint8(bln))
	} else if ln == "o2" {
		bln := len(b)
		if bln > 65535 {
			return nil, ErrTooMuchData
		}
		buf = make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, uint16(bln))
	} else {
		return nil, ErrNotImpl
	}
	return append(buf, []byte(b)...), nil
}

func stringUnmarshalZcl(ln string, b []byte) ([]byte, []byte, error) {
	var bln uint16
	if ln == "o1" {
		if len(b) < 1 {
			return nil, nil, ErrNotEnoughData
		}
		bln = uint16(b[0])
		b = b[1:]
	} else if ln == "o2" {
		if len(b) < 2 {
			return nil, nil, ErrNotEnoughData
		}
		bln = binary.LittleEndian.Uint16(b[0:2])
		b = b[2:]
	} else {
		return nil, nil, ErrNotImpl
	}
	if len(b) < int(bln) {
		return nil, nil, ErrNotEnoughData
	}

	return b[0:bln], b[bln:], nil
}

// Zostring is Octet string (0xff = invalid). A/D = D
type Zostring []byte

func (o *Zostring) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := stringUnmarshalZcl("o1", buf)
	*o = Zostring(val)
	return buf, err
}
func (o Zostring) MarshalZcl() ([]byte, error) { return stringMarshalZcl("o1", string(o)) }
func (o *Zostring) Values() []Val              { return []Val{o} }
func (o Zostring) ID() TypeID                  { return 65 }
func (o Zostring) String() string              { return string(o) }

//func (o Zostring) Valid() bool { return o != Zostring(255) }

// Zcstring is Character string (0xff = invalid). A/D = D
type Zcstring string

func (c *Zcstring) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := stringUnmarshalZcl("o1", buf)
	*c = Zcstring(val)
	return buf, err
}
func (c Zcstring) MarshalZcl() ([]byte, error) { return stringMarshalZcl("o1", string(c)) }
func (c *Zcstring) Values() []Val              { return []Val{c} }
func (c Zcstring) ID() TypeID                  { return 66 }
func (c *Zcstring) String() string             { return string(*c) }

//func (c Zcstring) Valid() bool { return c != Zcstring(255) }

// Zlostring is Long octet string (0xffff = invalid). A/D = D
type Zlostring string

func (l *Zlostring) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := stringUnmarshalZcl("o2", buf)
	*l = Zlostring(val)
	return buf, err
}
func (l Zlostring) MarshalZcl() ([]byte, error) { return stringMarshalZcl("o2", string(l)) }
func (l *Zlostring) Values() []Val              { return []Val{l} }
func (l Zlostring) ID() TypeID                  { return 67 }
func (l Zlostring) String() string              { return string(l) }

//func (l Zlostring) Valid() bool { return l != Zlostring(65535) }

// Zlcstring is Long character string (0xffff = invalid). A/D = D
type Zlcstring string

func (l *Zlcstring) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := stringUnmarshalZcl("o2", buf)
	*l = Zlcstring(val)
	return buf, err
}
func (l Zlcstring) MarshalZcl() ([]byte, error) { return stringMarshalZcl("o2", string(l)) }
func (l *Zlcstring) Values() []Val              { return []Val{l} }
func (l Zlcstring) ID() TypeID                  { return 68 }
func (l Zlcstring) String() string              { return string(l) }

//func (l Zlcstring) Valid() bool { return l != Zlcstring(65535) }
