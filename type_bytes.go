package zcl

func bytesMarshalZcl(ln int, o []byte) ([]byte, error) {
	if len(o) < ln {
		return nil, ErrNotEnoughData
	}
	if len(o) > ln {
		return nil, ErrTooMuchData
	}
	return o, nil
}

func bytesUnmarshalZcl(ln int, o []byte) ([]byte, []byte, error) {
	if len(o) < ln {
		return make([]byte, ln), nil, ErrNotEnoughData
	}
	return o[0:ln], o[ln:], nil
}

func BytesStringer(b []byte) string {
	var str []string
	for _, v := range b {
		str = append(str, Sprintf("0x%02X", v))
	}
	return StrJoin(str, " ")
}

type Zdat8 [1]byte
type Zdat16 [2]byte
type Zdat24 [3]byte
type Zdat32 [4]byte
type Zdat40 [5]byte
type Zdat48 [6]byte
type Zdat56 [7]byte
type Zdat64 [8]byte
type Zbytes []byte

func (d Zdat8) String() string               { return BytesStringer([]byte(d[:])) }
func (d Zdat16) String() string              { return BytesStringer([]byte(d[:])) }
func (d Zdat24) String() string              { return BytesStringer([]byte(d[:])) }
func (d Zdat32) String() string              { return BytesStringer([]byte(d[:])) }
func (d Zdat40) String() string              { return BytesStringer([]byte(d[:])) }
func (d Zdat48) String() string              { return BytesStringer([]byte(d[:])) }
func (d Zdat56) String() string              { return BytesStringer([]byte(d[:])) }
func (d Zdat64) String() string              { return BytesStringer([]byte(d[:])) }
func (d Zbytes) String() string              { return BytesStringer([]byte(d)) }
func (d Zdat8) MarshalZcl() ([]byte, error)  { return bytesMarshalZcl(1, []byte(d[:])) }
func (d Zdat16) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(2, []byte(d[:])) }
func (d Zdat24) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(3, []byte(d[:])) }
func (d Zdat32) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(4, []byte(d[:])) }
func (d Zdat40) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(5, []byte(d[:])) }
func (d Zdat48) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(6, []byte(d[:])) }
func (d Zdat56) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(7, []byte(d[:])) }
func (d Zdat64) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(8, []byte(d[:])) }
func (d Zbytes) MarshalZcl() ([]byte, error) { return []byte(d), nil }
func (d *Zdat8) Values() []Val               { return []Val{d} }
func (d *Zdat16) Values() []Val              { return []Val{d} }
func (d *Zdat24) Values() []Val              { return []Val{d} }
func (d *Zdat32) Values() []Val              { return []Val{d} }
func (d *Zdat40) Values() []Val              { return []Val{d} }
func (d *Zdat48) Values() []Val              { return []Val{d} }
func (d *Zdat56) Values() []Val              { return []Val{d} }
func (d *Zdat64) Values() []Val              { return []Val{d} }
func (d *Zbytes) Values() []Val              { return []Val{d} }
func (d Zdat8) TypeID() TypeID               { return 8 }
func (d Zdat16) TypeID() TypeID              { return 9 }
func (d Zdat24) TypeID() TypeID              { return 10 }
func (d Zdat32) TypeID() TypeID              { return 11 }
func (d Zdat40) TypeID() TypeID              { return 12 }
func (d Zdat48) TypeID() TypeID              { return 13 }
func (d Zdat56) TypeID() TypeID              { return 14 }
func (d Zdat64) TypeID() TypeID              { return 15 }

func (d *Zdat8) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(1, buf)
	copy((*d)[:], val)
	return buf, err
}

func (d *Zdat16) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(2, buf)
	copy((*d)[:], val)
	return buf, err
}

func (d *Zdat24) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(3, buf)
	copy((*d)[:], val)
	return buf, err
}

func (d *Zdat32) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(4, buf)
	copy((*d)[:], val)
	return buf, err
}

func (d *Zdat40) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(5, buf)
	copy((*d)[:], val)
	return buf, err
}

func (d *Zdat48) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(6, buf)
	copy((*d)[:], val)
	return buf, err
}

func (d *Zdat56) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(7, buf)
	copy((*d)[:], val)
	return buf, err
}

func (d *Zdat64) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(8, buf)
	copy((*d)[:], val)
	return buf, err
}

func (d *Zbytes) UnmarshalZcl(buf []byte) ([]byte, error) {
	*d = Zbytes(buf)
	return nil, nil
}
