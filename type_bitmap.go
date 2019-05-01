package zcl

import "encoding/binary"

func BitmapTest(a []byte, b uint) bool {
	// Pad to 64 bits
	aa := append(make([]byte, 8-len(a)), a...)
	ui := binary.BigEndian.Uint64(aa)
	return (ui & (1 << b)) == (1 << b)
}

func BitmapSet(a []byte, b uint, v bool) []byte {
	for i := len(a) - 1; i >= 0; i-- {
		if b >= 8 {
			b -= 8
			continue
		}
		if v {
			a[i] = a[i] | (1 << b)
		} else {
			a[i] = a[i] & ^(1 << b)
		}
		break
	}
	return a
}

// Zbmp8 is 8-bit bitmap. A/D = D
type Zbmp8 []byte

func (b *Zbmp8) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(1, buf)
	*b = Zbmp8(val)
	return buf, err
}
func (b Zbmp8) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(1, []byte(b)) }
func (b *Zbmp8) Values() []Val              { return []Val{b} }
func (b *Zbmp8) ID() TypeID                 { return 24 }

//

// Zbmp16 is 16-bit bitmap. A/D = D
type Zbmp16 []byte

func (b *Zbmp16) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(2, buf)
	*b = Zbmp16(val)
	return buf, err
}
func (b Zbmp16) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(2, []byte(b)) }
func (b *Zbmp16) Values() []Val              { return []Val{b} }
func (b *Zbmp16) ID() TypeID                 { return 25 }

//

// Zbmp24 is 24-bit bitmap. A/D = D
type Zbmp24 []byte

func (b *Zbmp24) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(3, buf)
	*b = Zbmp24(val)
	return buf, err
}
func (b Zbmp24) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(3, []byte(b)) }
func (b *Zbmp24) Values() []Val              { return []Val{b} }
func (b *Zbmp24) ID() TypeID                 { return 26 }

//

// Zbmp32 is 32-bit bitmap. A/D = D
type Zbmp32 []byte

func (b *Zbmp32) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(4, buf)
	*b = Zbmp32(val)
	return buf, err
}
func (b Zbmp32) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(4, []byte(b)) }
func (b *Zbmp32) Values() []Val              { return []Val{b} }
func (b *Zbmp32) ID() TypeID                 { return 27 }

//

// Zbmp40 is 40-bit bitmap. A/D = D
type Zbmp40 []byte

func (b *Zbmp40) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(5, buf)
	*b = Zbmp40(val)
	return buf, err
}
func (b Zbmp40) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(5, []byte(b)) }
func (b *Zbmp40) Values() []Val              { return []Val{b} }
func (b *Zbmp40) ID() TypeID                 { return 28 }

//

// Zbmp48 is 48-bit bitmap. A/D = D
type Zbmp48 []byte

func (b *Zbmp48) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(6, buf)
	*b = Zbmp48(val)
	return buf, err
}
func (b Zbmp48) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(6, []byte(b)) }
func (b *Zbmp48) Values() []Val              { return []Val{b} }
func (b *Zbmp48) ID() TypeID                 { return 29 }

//

// Zbmp56 is 56-bit bitmap. A/D = D
type Zbmp56 []byte

func (b *Zbmp56) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(7, buf)
	*b = Zbmp56(val)
	return buf, err
}
func (b Zbmp56) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(7, []byte(b)) }
func (b *Zbmp56) Values() []Val              { return []Val{b} }
func (b *Zbmp56) ID() TypeID                 { return 30 }

//

// Zbmp64 is 64-bit bitmap. A/D = D
type Zbmp64 []byte

func (b *Zbmp64) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(8, buf)
	*b = Zbmp64(val)
	return buf, err
}
func (b Zbmp64) MarshalZcl() ([]byte, error) { return bytesMarshalZcl(8, []byte(b)) }
func (b *Zbmp64) Values() []Val              { return []Val{b} }
func (b *Zbmp64) ID() TypeID                 { return 31 }
