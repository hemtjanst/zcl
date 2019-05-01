package zcl

import (
	"fmt"
	"time"
)

// UnixtimeOffset specifies the offset between zigbee time (epoch = 2000-01-01 UTC) and unixtime in seconds
const UnixtimeOffset = 946684800

// Zdate is Date (0xffffffff = invalid). A/D = A
type Zdate struct {
	Year    int
	Month   time.Month
	Day     int
	Weekday time.Weekday
}

// Ztime is Time of day (0xffffffff = invalid). A/D = A
type Ztime uint32

// Zutc is UTCTime (0xffffffff = invalid). A/D = A
type Zutc uint32

func (t *Ztime) UnmarshalZcl(buf []byte) ([]byte, error) {
	// (uint32) H(0) M(1) S(2) Centi(3)
	val, buf, err := uintLEUnmarshalZcl(4, buf)
	*t = Ztime(val)
	return buf, err
}
func (t Ztime) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(4, uint64(t)) }
func (t *Ztime) Values() []Val              { return []Val{t} }
func (t Ztime) ID() TypeID                  { return 224 }
func (t Ztime) String() string {
	h := (t & 0xFF000000) >> 24
	m := (t & 0x00FF0000) >> 16
	s := (t & 0x0000FF00) >> 8
	c := t & 0x000000FF
	return fmt.Sprintf("%02d:%02d:%02d.%03d", h, m, s, c)
}
func (t *Ztime) SetTime(hr, min, sec, hundreths int) {
	*t = Ztime(((hr & 0xFF) << 24) | ((min & 0xFF) << 16) | ((sec & 0xFF) << 8) | (hundreths & 0xFF))
}

func (t Ztime) Valid() bool { return t != Ztime(4294967295) }

func (d *Zdate) UnmarshalZcl(buf []byte) ([]byte, error) {
	// (uint32) Y=(b[0]+1900) M=b[1] D=b[2] DoW=b[3] (1-7 mon-sun)
	val, buf, err := uintLEUnmarshalZcl(4, buf)
	if err != nil {
		return buf, err
	}
	if val == 0xffffffff {
		return buf, err
	}
	d.Year = int((val&0xFF000000)>>24 + 1900)
	d.Month = time.Month((val & 0xFF0000) >> 16)
	d.Day = int((val & 0xFF00) >> 8)
	d.Weekday = time.Weekday(val & 0xFF)
	if d.Weekday == 7 {
		d.Weekday = time.Sunday
	}
	return buf, err
}
func (d Zdate) MarshalZcl() ([]byte, error) {
	var val uint64
	if d.Year == 0 {
		val = 0xffffffff
	} else {
		wd := int(d.Weekday)
		if wd == 0 {
			wd = 7
		}
		val = uint64(((d.Year-1900)&0xFF)<<24) |
			uint64((d.Month&0xFF)<<16) |
			uint64((d.Day&0xFF)<<8) |
			uint64(wd&0xFF)
	}
	return uintLEMarshalZcl(4, val)
}
func (d *Zdate) Values() []Val { return []Val{d} }
func (d Zdate) ID() TypeID     { return 225 }

func (d Zdate) Valid() bool { return d.Year > 0 }

func (u *Zutc) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(4, buf)
	*u = Zutc(val)
	return buf, err
}
func (u Zutc) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(4, uint64(u)) }
func (u *Zutc) Values() []Val              { return []Val{u} }
func (u Zutc) ID() TypeID                  { return 226 }
func (u Zutc) String() string              { return u.Time().String() }
func (u Zutc) Time() time.Time {
	// (uint32) Seconds since midnight 2000-01-01
	return time.Unix(int64(u)+UnixtimeOffset, 0)
}

func (u Zutc) Valid() bool { return u != Zutc(0xffffffff) }
