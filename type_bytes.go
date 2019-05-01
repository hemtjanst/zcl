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
		return nil, nil, ErrNotEnoughData
	}
	return o[0:ln], o[ln:], nil
}

// Zdat8 is 8-bit data. A/D = D
type Zdat8 []byte

func (d *Zdat8) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(1, buf)
	*d = Zdat8(val)
	return buf, err
}
func (d Zdat8) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(1, []byte(d)) }
func (d *Zdat8) Values() []Val              { return []Val{d} }
func (d *Zdat8) ID() TypeID                 { return 8 }

//

// Zdat16 is 16-bit data. A/D = D
type Zdat16 []byte

func (d *Zdat16) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(2, buf)
	*d = Zdat16(val)
	return buf, err
}
func (d Zdat16) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(2, []byte(d)) }
func (d *Zdat16) Values() []Val              { return []Val{d} }
func (d *Zdat16) ID() TypeID                 { return 9 }

//

// Zdat24 is 24-bit data. A/D = D
type Zdat24 []byte

func (d *Zdat24) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(3, buf)
	*d = Zdat24(val)
	return buf, err
}
func (d Zdat24) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(3, []byte(d)) }
func (d *Zdat24) Values() []Val              { return []Val{d} }
func (d *Zdat24) ID() TypeID                 { return 10 }

//

// Zdat32 is 32-bit data. A/D = D
type Zdat32 []byte

func (d *Zdat32) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(4, buf)
	*d = Zdat32(val)
	return buf, err
}
func (d Zdat32) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(4, []byte(d)) }
func (d *Zdat32) Values() []Val              { return []Val{d} }
func (d *Zdat32) ID() TypeID                 { return 11 }

//

// Zdat40 is 40-bit data. A/D = D
type Zdat40 []byte

func (d *Zdat40) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(5, buf)
	*d = Zdat40(val)
	return buf, err
}
func (d Zdat40) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(5, []byte(d)) }
func (d *Zdat40) Values() []Val              { return []Val{d} }
func (d *Zdat40) ID() TypeID                 { return 12 }

//

// Zdat48 is 48-bit data. A/D = D
type Zdat48 []byte

func (d *Zdat48) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(6, buf)
	*d = Zdat48(val)
	return buf, err
}
func (d Zdat48) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(6, []byte(d)) }
func (d *Zdat48) Values() []Val              { return []Val{d} }
func (d *Zdat48) ID() TypeID                 { return 13 }

//

// Zdat56 is 56-bit data. A/D = D
type Zdat56 []byte

func (d *Zdat56) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(7, buf)
	*d = Zdat56(val)
	return buf, err
}
func (d Zdat56) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(7, []byte(d)) }
func (d *Zdat56) Values() []Val              { return []Val{d} }
func (d *Zdat56) ID() TypeID                 { return 14 }

//

// Zdat64 is 64-bit data. A/D = D
type Zdat64 []byte

func (d *Zdat64) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(8, buf)
	*d = Zdat64(val)
	return buf, err
}
func (d Zdat64) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(8, []byte(d)) }
func (d *Zdat64) Values() []Val              { return []Val{d} }
func (d *Zdat64) ID() TypeID                 { return 15 }

//
