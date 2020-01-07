package zcl

import "fmt"

func ArrayMarshalZcl(ln string, t TypeID, v []Val) ([]byte, error) {
	var tmp []byte
	var err error

	if tmp, err = t.MarshalZcl(); err != nil {
		return nil, err
	}

	data, err := ArrayNoTypeMarshalZcl(ln, v)
	return append(tmp, data...), err
}

func ArrayNoTypeMarshalZcl(ln string, v []Val) ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	switch ln {
	case "0":
	case "sloc", "1":
		lv := Zu8(len(v))
		if tmp, err = lv.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	case "2+sloc", "2":
		lv := Zu16(len(v))
		if tmp, err = lv.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	default:
		return nil, ErrInvalidArrayLength
	}

	for _, val := range v {
		if tmp, err = val.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil

}

func UnmarshalList(ln string, b []byte, newVal func() Val) ([]byte, error) {
	var err error
	var length int

	switch ln {
	case "0":
		length = -1
	case "sloc", "1":
		lv := new(Zu8)
		if b, err = lv.UnmarshalZcl(b); err != nil {
			return nil, err
		}
		length = int(*lv)
	case "2+sloc", "2":
		lv := new(Zu16)
		if b, err = lv.UnmarshalZcl(b); err != nil {
			return nil, err
		}
		length = int(*lv)
	default:
		return nil, ErrInvalidArrayLength
	}

	var val []Val

	for i := 0; i < length; i++ {
		if length == -1 && len(b) == 0 {
			break
		}
		v := newVal()
		if v == nil {
			return nil, ErrInvalidType
		}
		if b, err = v.UnmarshalZcl(b); err != nil {
			return nil, err
		}
		val = append(val, v)
	}

	return b, nil
}

func ArrayUnmarshalZcl(ln string, b []byte) (TypeID, []Val, []byte, error) {
	var err error
	var t TypeID
	var val []Val

	if b, err = (&t).UnmarshalZcl(b); err != nil {
		return 0, nil, nil, err
	}

	val, b, err = ArrayNoTypeUnmarshalZcl(ln, b, t)

	return t, val, b, err
}

func ArrayNoTypeUnmarshalZcl(ln string, b []byte, t TypeID) ([]Val, []byte, error) {
	var val []Val
	ret, err := UnmarshalList(ln, b, func() Val {
		nv := NewValue(uint8(t))
		val = append(val, nv)
		return nv
	})
	return val, ret, err
}

func StructMarshalZcl(ln string, v []StructField) ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error
	switch ln {
	case "sloc", "1":
		lv := Zu8(len(v))
		if tmp, err = lv.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	case "2+sloc", "2":
		lv := Zu16(len(v))
		if tmp, err = lv.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	default:
		return nil, ErrInvalidArrayLength
	}

	for _, val := range v {
		if tmp, err = val.Type.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)

		if tmp, err = val.Content.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	return data, nil
}
func StructUnmarshalZcl(ln string, b []byte) ([]StructField, []byte, error) {
	var err error

	var length int
	switch ln {
	case "sloc", "1":
		lv := new(Zu8)
		if b, err = lv.UnmarshalZcl(b); err != nil {
			return nil, nil, err
		}
		length = int(*lv)
	case "2+sloc", "2":
		lv := new(Zu16)
		if b, err = lv.UnmarshalZcl(b); err != nil {
			return nil, nil, err
		}
		length = int(*lv)
	default:
		return nil, nil, ErrInvalidArrayLength
	}

	var val []StructField

	for i := 0; i < length; i++ {
		var t TypeID
		if b, err = (&t).UnmarshalZcl(b); err != nil {
			return nil, nil, err
		}
		v := NewValue(uint8(t))
		if v == nil {
			return nil, nil, ErrInvalidType
		}
		if b, err = v.UnmarshalZcl(b); err != nil {
			return nil, nil, err
		}
		val = append(val, StructField{Type: t, Content: v})
	}

	return val, b, nil
}

func arraySetMarshalZcl(ln string, t TypeID, v []Val) ([]byte, error) {
	return ArrayMarshalZcl(ln, t, v)
}
func arraySetUnmarshalZcl(ln string, b []byte) (TypeID, []Val, []byte, error) {
	return ArrayUnmarshalZcl(ln, b)
}

type StructField struct {
	Type    TypeID
	Content Val
}

func (s StructField) String() string {
	return fmt.Sprintf("%v", s.Content)
}

// Zarray is Array (0xffff = invalid). A/D = D
type Zarray struct {
	Type    TypeID
	Content []Val
}

func (a Zarray) String() string {
	var str []string
	for _, v := range a.Content {
		str = append(str, fmt.Sprintf("%v", v))
	}
	return fmt.Sprintf("[%s]", StrJoin(str, ", "))
}

func (a *Zarray) UnmarshalZcl(buf []byte) ([]byte, error) {
	var err error
	a.Type, a.Content, buf, err = ArrayUnmarshalZcl("2+sloc", buf)
	return buf, err
}
func (a Zarray) MarshalZcl() ([]byte, error) { return ArrayMarshalZcl("2+sloc", a.Type, a.Content) }
func (a *Zarray) Values() []Val              { return a.Content }
func (a Zarray) TypeID() TypeID              { return 72 }

//func (a Zarray) Valid() bool { return a != Zarray(65535) }

// Zstruct is Structure (0xffff = invalid). A/D = D
type Zstruct []StructField

func (s Zstruct) String() string {
	var str []string
	for _, v := range s {
		str = append(str, fmt.Sprintf("%v", v))
	}
	return fmt.Sprintf("[%s]", StrJoin(str, ", "))
}

func (s *Zstruct) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := StructUnmarshalZcl("2+sloc", buf)
	*s = Zstruct(val)
	return buf, err
}
func (s Zstruct) MarshalZcl() ([]byte, error) { return StructMarshalZcl("2+sloc", []StructField(s)) }
func (s *Zstruct) Values() []Val              { return []Val{s} }
func (s Zstruct) TypeID() TypeID              { return 76 }

//func (s Zstruct) Valid() bool { return s != Zstruct(65535) }

// Zset is Set (0xffff = invalid). A/D = D
type Zset struct {
	Type    TypeID
	Content []Val
}

func (s Zset) String() string {
	var str []string
	for _, v := range s.Content {
		str = append(str, fmt.Sprintf("%v", v))
	}
	return fmt.Sprintf("[%s]", StrJoin(str, ", "))
}

func (s *Zset) UnmarshalZcl(buf []byte) ([]byte, error) {
	var err error
	s.Type, s.Content, buf, err = arraySetUnmarshalZcl("sloc", buf)
	return buf, err
}
func (s Zset) MarshalZcl() ([]byte, error) { return arraySetMarshalZcl("sloc", s.Type, s.Content) }

func (s *Zset) Values() []Val { return s.Content }
func (s Zset) TypeID() TypeID { return 80 }

//func (s Zset) Valid() bool { return s != Zset(65535) }

// Zbag is Bag (0xffff = invalid). A/D = D
type Zbag struct {
	Type    TypeID
	Content []Val
}

func (b Zbag) String() string {
	var str []string
	for _, v := range b.Content {
		str = append(str, fmt.Sprintf("%v", v))
	}
	return fmt.Sprintf("[%s]", StrJoin(str, ", "))
}

func (b *Zbag) UnmarshalZcl(buf []byte) ([]byte, error) {
	var err error
	b.Type, b.Content, buf, err = ArrayUnmarshalZcl("sloc", buf)
	return buf, err
}
func (b Zbag) MarshalZcl() ([]byte, error) { return ArrayMarshalZcl("sloc", b.Type, b.Content) }
func (b *Zbag) Values() []Val              { return b.Content }
func (b Zbag) TypeID() TypeID              { return 81 }

//func (b Zbag) Valid() bool { return b != Zbag(65535) }

type Zlist struct {
	Type    TypeID
	Content []Val
}

func (b Zlist) String() string {
	var str []string
	for _, v := range b.Content {
		str = append(str, fmt.Sprintf("%v", v))
	}
	return fmt.Sprintf("[%s]", StrJoin(str, ", "))
}

func (b *Zlist) UnmarshalZcl(buf []byte) ([]byte, error) {
	var err error
	b.Type, b.Content, buf, err = ArrayUnmarshalZcl("0", buf)
	return buf, err
}

func (b Zlist) MarshalZcl() ([]byte, error) { return ArrayMarshalZcl("0", b.Type, b.Content) }
func (b *Zlist) Values() []Val              { return b.Content }
func (b Zlist) TypeID() TypeID              { return 81 }
