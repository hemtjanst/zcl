package zcl

import (
	"encoding/binary"
	"hemtjan.st/zcl/utils/half"
	"math"
	"strconv"
)

func uintLEMarshalZcl(ln int, o uint64) ([]byte, error) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, o)
	return b[0:ln], nil
}
func uintLEUnmarshalZcl(ln int, b []byte) (uint64, []byte, error) {
	if len(b) < ln {
		return 0, nil, ErrNotEnoughData
	}
	buf := make([]byte, 8)
	for i := 0; i < ln; i++ {
		buf[i] = b[i]
	}
	return binary.LittleEndian.Uint64(buf), b[ln:], nil
}

func intLEMarshalZcl(ln int, o int64) ([]byte, error) {
	b := make([]byte, 8)
	ui := uint64(o)
	if ln < 8 {
		sign64 := uint64(0x80000000)
		sign := uint64((1 << uint(ln*8)) - 1)
		if (ui & sign64) == sign64 {
			// Move the sign to the correct place
			ui = (ui ^ sign64) | sign
		}
	}
	binary.LittleEndian.PutUint64(b, ui)
	return b[0:ln], nil
}
func intLEUnmarshalZcl(ln int, b []byte) (int64, []byte, error) {
	if len(b) < ln {
		return 0, nil, ErrNotEnoughData
	}
	buf := make([]byte, 8)
	for i := 0; i < ln; i++ {
		buf[i] = b[i]
	}
	ui := binary.LittleEndian.Uint64(buf)
	if ln < 8 {
		sign64 := uint64(0x80000000)
		sign := uint64((1 << uint(ln*8)) - 1)
		if (ui & sign) == sign {
			// Move the sign to the correct place
			ui = (ui ^ sign) | sign64
		}
	}
	return int64(ui), b[ln:], nil
}

func float32MarshalZcl(ln int, b float32) ([]byte, error) {
	var buf []byte
	if ln == 2 {
		buf = make([]byte, 2)
		hf := half.NewFloat16(b)
		binary.LittleEndian.PutUint16(buf, uint16(hf))
	} else {
		buf = make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, math.Float32bits(b))
	}
	return buf, nil
}

func float32UnmarshalZcl(ln int, b []byte) (float32, []byte, error) {
	if len(b) < ln {
		return 0, nil, ErrNotEnoughData
	}
	var f float32
	if ln == 2 {
		ui := binary.LittleEndian.Uint16(b[0:2])
		f = half.Float16(ui).Float32()
	} else {
		ui := binary.LittleEndian.Uint32(b[0:4])
		f = math.Float32frombits(ui)
	}
	return f, b[ln:], nil
}

func float64MarshalZcl(ln int, b float64) ([]byte, error) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, math.Float64bits(b))
	return buf, nil
}

func float64UnmarshalZcl(ln int, b []byte) (float64, []byte, error) {
	if len(b) < 8 {
		return 0, nil, ErrNotEnoughData
	}
	ui := binary.LittleEndian.Uint64(b[0:8])
	return math.Float64frombits(ui), b[8:], nil
}

func (u Zu8) String() string     { return strconv.FormatUint(uint64(u), 10) }
func (u Zu16) String() string    { return strconv.FormatUint(uint64(u), 10) }
func (u Zu24) String() string    { return strconv.FormatUint(uint64(u), 10) }
func (u Zu32) String() string    { return strconv.FormatUint(uint64(u), 10) }
func (u Zu40) String() string    { return strconv.FormatUint(uint64(u), 10) }
func (u Zu48) String() string    { return strconv.FormatUint(uint64(u), 10) }
func (u Zu56) String() string    { return strconv.FormatUint(uint64(u), 10) }
func (u Zu64) String() string    { return strconv.FormatUint(uint64(u), 10) }
func (s Zs8) String() string     { return strconv.FormatInt(int64(s), 10) }
func (s Zs16) String() string    { return strconv.FormatInt(int64(s), 10) }
func (s Zs24) String() string    { return strconv.FormatInt(int64(s), 10) }
func (s Zs32) String() string    { return strconv.FormatInt(int64(s), 10) }
func (s Zs40) String() string    { return strconv.FormatInt(int64(s), 10) }
func (s Zs48) String() string    { return strconv.FormatInt(int64(s), 10) }
func (s Zs56) String() string    { return strconv.FormatInt(int64(s), 10) }
func (s Zs64) String() string    { return strconv.FormatInt(int64(s), 10) }
func (s Zsemi) String() string   { return strconv.FormatFloat(float64(s), 'f', -1, 32) }
func (f Zfloat) String() string  { return strconv.FormatFloat(float64(f), 'f', -1, 32) }
func (d Zdouble) String() string { return strconv.FormatFloat(float64(d), 'f', -1, 64) }

// Zu8 is Unsigned 8-bit integer (0xff = invalid). A/D = A
type Zu8 uint8

func (u *Zu8) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(1, buf)
	*u = Zu8(val)
	return buf, err
}
func (u Zu8) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(1, uint64(u)) }
func (u *Zu8) Values() []Val              { return []Val{u} }
func (u Zu8) TypeID() TypeID              { return 32 }
func (u Zu8) Valid() bool                 { return u < 0xFF }

// Zu16 is Unsigned 16-bit integer (0xffff = invalid). A/D = A
type Zu16 uint16

func (u *Zu16) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(2, buf)
	*u = Zu16(val)
	return buf, err
}
func (u Zu16) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(2, uint64(u)) }
func (u *Zu16) Values() []Val              { return []Val{u} }
func (u Zu16) TypeID() TypeID              { return 33 }
func (u Zu16) Valid() bool                 { return u < 0xFFFF }

// Zu24 is Unsigned 24-bit integer (0xffffff = invalid). A/D = A
type Zu24 uint32

func (u *Zu24) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(3, buf)
	*u = Zu24(val)
	return buf, err
}
func (u Zu24) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(3, uint64(u)) }
func (u *Zu24) Values() []Val              { return []Val{u} }
func (u Zu24) TypeID() TypeID              { return 34 }
func (u Zu24) Valid() bool                 { return u < 0xFFFFFF }

// Zu32 is Unsigned 32-bit integer (0xffffffff = invalid). A/D = A
type Zu32 uint32

func (u *Zu32) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(4, buf)
	*u = Zu32(val)
	return buf, err
}
func (u Zu32) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(4, uint64(u)) }
func (u *Zu32) Values() []Val              { return []Val{u} }
func (u Zu32) TypeID() TypeID              { return 35 }
func (u Zu32) Valid() bool                 { return u < 0xFFFFFFFF }

// Zu40 is Unsigned 40-bit integer (0xffffffffff = invalid). A/D = A
type Zu40 uint64

func (u *Zu40) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(5, buf)
	*u = Zu40(val)
	return buf, err
}
func (u Zu40) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(5, uint64(u)) }
func (u *Zu40) Values() []Val              { return []Val{u} }
func (u Zu40) TypeID() TypeID              { return 36 }
func (u Zu40) Valid() bool                 { return u < 0xFFFFFFFFFF }

// Zu48 is Unsigned 48-bit integer (0xffffffffffff = invalid). A/D = A
type Zu48 uint64

func (u *Zu48) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(6, buf)
	*u = Zu48(val)
	return buf, err
}
func (u Zu48) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(6, uint64(u)) }
func (u *Zu48) Values() []Val              { return []Val{u} }
func (u Zu48) TypeID() TypeID              { return 37 }

func (u Zu48) Valid() bool { return u < 0xFFFFFFFFFFFF }

// Zu56 is Unsigned 56-bit integer (0xffffffffffffff = invalid). A/D = A
type Zu56 uint64

func (u *Zu56) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(7, buf)
	*u = Zu56(val)
	return buf, err
}
func (u Zu56) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(7, uint64(u)) }
func (u *Zu56) Values() []Val              { return []Val{u} }
func (u Zu56) TypeID() TypeID              { return 38 }

func (u Zu56) Valid() bool { return u < 0xFFFFFFFFFFFFFF }

// Zu64 is Unsigned 64-bit integer (0xffffffffffffffff = invalid). A/D = A
type Zu64 uint64

func (u *Zu64) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(8, buf)
	*u = Zu64(val)
	return buf, err
}
func (u Zu64) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(8, uint64(u)) }
func (u *Zu64) Values() []Val              { return []Val{u} }
func (u Zu64) TypeID() TypeID              { return 39 }

func (u Zu64) Valid() bool { return u < 0xFFFFFFFFFFFFFFFF }

// Zs8 is Signed 8-bit integer (0x80 = invalid). A/D = A
type Zs8 int8

func (s *Zs8) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := intLEUnmarshalZcl(1, buf)
	*s = Zs8(val)
	return buf, err
}
func (s Zs8) MarshalZcl() ([]byte, error) { return intLEMarshalZcl(1, int64(s)) }
func (s *Zs8) Values() []Val              { return []Val{s} }
func (s Zs8) TypeID() TypeID              { return 40 }

//func (s Zs8) Valid() bool { return s != Zs8(128) }

// Zs16 is Signed 16-bit integer (0x8000 = invalid). A/D = A
type Zs16 int16

func (s *Zs16) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := intLEUnmarshalZcl(2, buf)
	*s = Zs16(val)
	return buf, err
}
func (s Zs16) MarshalZcl() ([]byte, error) { return intLEMarshalZcl(2, int64(s)) }
func (s *Zs16) Values() []Val              { return []Val{s} }
func (s Zs16) TypeID() TypeID              { return 41 }

//func (s Zs16) Valid() bool { return s != Zs16(32768) }

// Zs24 is Signed 24-bit integer (0x800000 = invalid). A/D = A
type Zs24 int32

func (s *Zs24) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := intLEUnmarshalZcl(3, buf)
	*s = Zs24(val)
	return buf, err
}
func (s Zs24) MarshalZcl() ([]byte, error) { return intLEMarshalZcl(3, int64(s)) }
func (s *Zs24) Values() []Val              { return []Val{s} }
func (s Zs24) TypeID() TypeID              { return 42 }

//func (s Zs24) Valid() bool { return s != Zs24(8388608) }

// Zs32 is Signed 32-bit integer (0x80000000 = invalid). A/D = A
type Zs32 int32

func (s *Zs32) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := intLEUnmarshalZcl(4, buf)
	*s = Zs32(val)
	return buf, err
}
func (s Zs32) MarshalZcl() ([]byte, error) { return intLEMarshalZcl(4, int64(s)) }
func (s *Zs32) Values() []Val              { return []Val{s} }
func (s Zs32) TypeID() TypeID              { return 43 }

//func (s Zs32) Valid() bool { return s != Zs32(2147483648) }

// Zs40 is Signed 40-bit integer (0x8000000000 = invalid). A/D = A
type Zs40 int64

func (s *Zs40) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := intLEUnmarshalZcl(5, buf)
	*s = Zs40(val)
	return buf, err
}
func (s Zs40) MarshalZcl() ([]byte, error) { return intLEMarshalZcl(5, int64(s)) }
func (s *Zs40) Values() []Val              { return []Val{s} }
func (s Zs40) TypeID() TypeID              { return 44 }

//func (s Zs40) Valid() bool { return s != Zs40(549755813888) }

// Zs48 is Signed 48-bit integer (0x800000000000 = invalid). A/D = A
type Zs48 int64

func (s *Zs48) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := intLEUnmarshalZcl(6, buf)
	*s = Zs48(val)
	return buf, err
}
func (s Zs48) MarshalZcl() ([]byte, error) { return intLEMarshalZcl(6, int64(s)) }
func (s *Zs48) Values() []Val              { return []Val{s} }
func (s Zs48) TypeID() TypeID              { return 45 }

//func (s Zs48) Valid() bool { return s != Zs48(140737488355328) }

// Zs56 is Signed 56-bit integer (0x80000000000000 = invalid). A/D = A
type Zs56 int64

func (s *Zs56) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := intLEUnmarshalZcl(7, buf)
	*s = Zs56(val)
	return buf, err
}
func (s Zs56) MarshalZcl() ([]byte, error) { return intLEMarshalZcl(7, int64(s)) }
func (s *Zs56) Values() []Val              { return []Val{s} }
func (s Zs56) TypeID() TypeID              { return 46 }

//func (s Zs56) Valid() bool { return s != Zs56(36028797018963968) }

// Zs64 is Signed 64-bit integer (0x8000000000000000 = invalid). A/D = A
type Zs64 int64

func (s *Zs64) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := intLEUnmarshalZcl(8, buf)
	*s = Zs64(val)
	return buf, err
}
func (s Zs64) MarshalZcl() ([]byte, error) { return intLEMarshalZcl(8, int64(s)) }
func (s *Zs64) Values() []Val              { return []Val{s} }
func (s Zs64) TypeID() TypeID              { return 47 }

//func (s Zs64) Valid() bool { return s != Zs64(9223372036854775808) }

// Zsemi is Semi-precicion (0x = invalid). A/D = A
type Zsemi float32

func (s *Zsemi) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := float32UnmarshalZcl(2, buf)
	*s = Zsemi(val)
	return buf, err
}
func (s Zsemi) MarshalZcl() ([]byte, error) { return float32MarshalZcl(2, float32(s)) }
func (s *Zsemi) Values() []Val              { return []Val{s} }
func (s Zsemi) TypeID() TypeID              { return 56 }

//func (s Zsemi) Valid() bool { return s != Zsemi(0) }

// Zfloat is Single precicion (0x = invalid). A/D = A
type Zfloat float32

func (f *Zfloat) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := float32UnmarshalZcl(4, buf)
	*f = Zfloat(val)
	return buf, err
}
func (f Zfloat) MarshalZcl() ([]byte, error) { return float32MarshalZcl(4, float32(f)) }
func (f *Zfloat) Values() []Val              { return []Val{f} }
func (f Zfloat) TypeID() TypeID              { return 57 }

//func (f Zfloat) Valid() bool { return f != Zfloat(0) }

// Zdouble is Double precicion (0x = invalid). A/D = A
type Zdouble float64

func (d *Zdouble) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := float64UnmarshalZcl(8, buf)
	*d = Zdouble(val)
	return buf, err
}
func (d Zdouble) MarshalZcl() ([]byte, error) { return float64MarshalZcl(8, float64(d)) }
func (d *Zdouble) Values() []Val              { return []Val{d} }
func (d Zdouble) TypeID() TypeID              { return 58 }

//func (d Zdouble) Valid() bool { return d != Zdouble(0) }
