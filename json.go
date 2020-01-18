package zcl

import (
	"encoding/json"
	"strconv"
)

func intUnmarshalJSON(b []byte) (int64, error) {
	i := int64(0)
	if err := json.Unmarshal(b, &i); err == nil {
		return i, nil
	}
	s := ""
	if err := json.Unmarshal(b, &s); err == nil {
		return strconv.ParseInt(s, 10, 64)
	} else {
		return 0, err
	}
}

func floatUnmarshalJSON(b []byte) (float64, error) {
	i := float64(0)
	if err := json.Unmarshal(b, &i); err == nil {
		return i, nil
	}
	s := ""
	if err := json.Unmarshal(b, &s); err == nil {
		return strconv.ParseFloat(s, 64)
	} else {
		return 0, err
	}
}

func (u *Zenum8) UnmarshalJSON(b []byte) error {
	if v, e := intUnmarshalJSON(b); e != nil {
		return e
	} else {
		*u = Zenum8(v)
	}
	return nil
}
func (u *Zenum16) UnmarshalJSON(b []byte) error {
	if v, e := intUnmarshalJSON(b); e != nil {
		return e
	} else {
		*u = Zenum16(v)
	}
	return nil
}
func (u *Zu8) UnmarshalJSON(b []byte) error {
	if v, e := intUnmarshalJSON(b); e != nil {
		return e
	} else {
		*u = Zu8(v)
	}
	return nil
}
func (u *Zu16) UnmarshalJSON(b []byte) error {
	if v, e := intUnmarshalJSON(b); e != nil {
		return e
	} else {
		*u = Zu16(v)
	}
	return nil
}
func (u *Zu24) UnmarshalJSON(b []byte) error {
	if v, e := intUnmarshalJSON(b); e != nil {
		return e
	} else {
		*u = Zu24(v)
	}
	return nil
}
func (u *Zu32) UnmarshalJSON(b []byte) error {
	if v, e := intUnmarshalJSON(b); e != nil {
		return e
	} else {
		*u = Zu32(v)
	}
	return nil
}
func (u *Zu40) UnmarshalJSON(b []byte) error {
	if v, e := intUnmarshalJSON(b); e != nil {
		return e
	} else {
		*u = Zu40(v)
	}
	return nil
}
func (u *Zu48) UnmarshalJSON(b []byte) error {
	if v, e := intUnmarshalJSON(b); e != nil {
		return e
	} else {
		*u = Zu48(v)
	}
	return nil
}
func (u *Zu56) UnmarshalJSON(b []byte) error {
	if v, e := intUnmarshalJSON(b); e != nil {
		return e
	} else {
		*u = Zu56(v)
	}
	return nil
}
func (u *Zu64) UnmarshalJSON(b []byte) error {
	if v, e := intUnmarshalJSON(b); e != nil {
		return e
	} else {
		*u = Zu64(v)
	}
	return nil
}
func (s *Zs8) UnmarshalJSON(b []byte) error {
	if v, e := intUnmarshalJSON(b); e != nil {
		return e
	} else {
		*s = Zs8(v)
	}
	return nil
}
func (s *Zs16) UnmarshalJSON(b []byte) error {
	if v, e := intUnmarshalJSON(b); e != nil {
		return e
	} else {
		*s = Zs16(v)
	}
	return nil
}
func (s *Zs24) UnmarshalJSON(b []byte) error {
	if v, e := intUnmarshalJSON(b); e != nil {
		return e
	} else {
		*s = Zs24(v)
	}
	return nil
}
func (s *Zs32) UnmarshalJSON(b []byte) error {
	if v, e := intUnmarshalJSON(b); e != nil {
		return e
	} else {
		*s = Zs32(v)
	}
	return nil
}
func (s *Zs40) UnmarshalJSON(b []byte) error {
	if v, e := intUnmarshalJSON(b); e != nil {
		return e
	} else {
		*s = Zs40(v)
	}
	return nil
}
func (s *Zs48) UnmarshalJSON(b []byte) error {
	if v, e := intUnmarshalJSON(b); e != nil {
		return e
	} else {
		*s = Zs48(v)
	}
	return nil
}
func (s *Zs56) UnmarshalJSON(b []byte) error {
	if v, e := intUnmarshalJSON(b); e != nil {
		return e
	} else {
		*s = Zs56(v)
	}
	return nil
}
func (s *Zs64) UnmarshalJSON(b []byte) error {
	if v, e := intUnmarshalJSON(b); e != nil {
		return e
	} else {
		*s = Zs64(v)
	}
	return nil
}
func (s *Zsemi) UnmarshalJSON(b []byte) error {
	if v, e := floatUnmarshalJSON(b); e != nil {
		return e
	} else {
		*s = Zsemi(v)
	}
	return nil
}
func (f *Zfloat) UnmarshalJSON(b []byte) error {
	if v, e := floatUnmarshalJSON(b); e != nil {
		return e
	} else {
		*f = Zfloat(v)
	}
	return nil
}
func (d *Zdouble) UnmarshalJSON(b []byte) error {
	if v, e := floatUnmarshalJSON(b); e != nil {
		return e
	} else {
		*d = Zdouble(v)
	}
	return nil
}
