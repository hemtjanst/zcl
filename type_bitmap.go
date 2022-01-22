package zcl

import (
	"encoding/binary"
	"encoding/json"
	"strconv"
)

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

func BitmapList(b []byte) []int {
	var bits []int
	i := 0
	for ln := len(b) - 1; ln >= 0; ln-- {
		for bit := 0x01; bit < 0xff; bit = bit << 1 { // 1 2 4 8 16 32 64 128
			if b[ln]&uint8(bit) == uint8(bit) {
				bits = append(bits, i)
			}
			i++
		}
	}
	return bits
}

func bitmapMarshalJson(b []byte) ([]byte, error) {
	var s []int
	for i := 0; i < len(b)*8; i++ {
		if BitmapTest(b, uint(i)) {
			s = append(s, i)
		}
	}
	return json.Marshal(s)
}

func bitmapUnmarshalJson(ln int, b []byte) ([]byte, error) {
	o := make([]byte, ln)
	v := new([]uint8)
	if err := json.Unmarshal(b, v); err != nil {
		v = new([]uint8)
		// Try with strings instead
		vv := new([]string)
		if err := json.Unmarshal(b, vv); err != nil {
			return o, err
		}
		for _, s := range *vv {
			if n, err := strconv.ParseUint(s, 10, 8); err != nil {
				return o, err
			} else {
				*v = append(*v, uint8(n))
			}
		}

	}
	for _, i := range *v {
		o = BitmapSet(o, uint(i), true)
	}
	return o, nil
}

func BitmapStringer(b []byte) string {
	bits := BitmapList(b)
	return Sprintf("Bits%v", bits)
}

type Zbmp8 [1]byte
type Zbmp16 [2]byte
type Zbmp24 [3]byte
type Zbmp32 [4]byte
type Zbmp40 [5]byte
type Zbmp48 [6]byte
type Zbmp56 [7]byte
type Zbmp64 [8]byte

func (b Zbmp8) String() string                { return BitmapStringer(b[:]) }
func (b Zbmp16) String() string               { return BitmapStringer(b[:]) }
func (b Zbmp24) String() string               { return BitmapStringer(b[:]) }
func (b Zbmp32) String() string               { return BitmapStringer(b[:]) }
func (b Zbmp40) String() string               { return BitmapStringer(b[:]) }
func (b Zbmp48) String() string               { return BitmapStringer(b[:]) }
func (b Zbmp56) String() string               { return BitmapStringer(b[:]) }
func (b Zbmp64) String() string               { return BitmapStringer(b[:]) }
func (b Zbmp8) MarshalJSON() ([]byte, error)  { return bitmapMarshalJson(b[:]) }
func (b Zbmp16) MarshalJSON() ([]byte, error) { return bitmapMarshalJson(b[:]) }
func (b Zbmp24) MarshalJSON() ([]byte, error) { return bitmapMarshalJson(b[:]) }
func (b Zbmp32) MarshalJSON() ([]byte, error) { return bitmapMarshalJson(b[:]) }
func (b Zbmp40) MarshalJSON() ([]byte, error) { return bitmapMarshalJson(b[:]) }
func (b Zbmp48) MarshalJSON() ([]byte, error) { return bitmapMarshalJson(b[:]) }
func (b Zbmp56) MarshalJSON() ([]byte, error) { return bitmapMarshalJson(b[:]) }
func (b Zbmp64) MarshalJSON() ([]byte, error) { return bitmapMarshalJson(b[:]) }
func (b Zbmp8) MarshalZcl() ([]byte, error)   { return bytesMarshalZcl(1, b[:]) }
func (b Zbmp16) MarshalZcl() ([]byte, error)  { return bytesMarshalZcl(2, reverse(b[:])) }
func (b Zbmp24) MarshalZcl() ([]byte, error)  { return bytesMarshalZcl(3, reverse(b[:])) }
func (b Zbmp32) MarshalZcl() ([]byte, error)  { return bytesMarshalZcl(4, reverse(b[:])) }
func (b Zbmp40) MarshalZcl() ([]byte, error)  { return bytesMarshalZcl(5, reverse(b[:])) }
func (b Zbmp48) MarshalZcl() ([]byte, error)  { return bytesMarshalZcl(6, reverse(b[:])) }
func (b Zbmp56) MarshalZcl() ([]byte, error)  { return bytesMarshalZcl(7, reverse(b[:])) }
func (b Zbmp64) MarshalZcl() ([]byte, error)  { return bytesMarshalZcl(8, reverse(b[:])) }
func (b *Zbmp8) Values() []Val                { return []Val{b} }
func (b *Zbmp16) Values() []Val               { return []Val{b} }
func (b *Zbmp24) Values() []Val               { return []Val{b} }
func (b *Zbmp32) Values() []Val               { return []Val{b} }
func (b *Zbmp40) Values() []Val               { return []Val{b} }
func (b *Zbmp48) Values() []Val               { return []Val{b} }
func (b *Zbmp56) Values() []Val               { return []Val{b} }
func (b *Zbmp64) Values() []Val               { return []Val{b} }
func (b Zbmp8) TypeID() TypeID                { return 24 }
func (b Zbmp16) TypeID() TypeID               { return 25 }
func (b Zbmp24) TypeID() TypeID               { return 26 }
func (b Zbmp32) TypeID() TypeID               { return 27 }
func (b Zbmp40) TypeID() TypeID               { return 28 }
func (b Zbmp48) TypeID() TypeID               { return 29 }
func (b Zbmp56) TypeID() TypeID               { return 30 }
func (b Zbmp64) TypeID() TypeID               { return 31 }

func (b *Zbmp8) UnmarshalJSON(buf []byte) error {
	val, err := bitmapUnmarshalJson(1, buf)
	copy((*b)[:], val)
	return err
}
func (b *Zbmp16) UnmarshalJSON(buf []byte) error {
	val, err := bitmapUnmarshalJson(2, buf)
	copy((*b)[:], val)
	return err
}
func (b *Zbmp24) UnmarshalJSON(buf []byte) error {
	val, err := bitmapUnmarshalJson(3, buf)
	copy((*b)[:], val)
	return err
}
func (b *Zbmp32) UnmarshalJSON(buf []byte) error {
	val, err := bitmapUnmarshalJson(4, buf)
	copy((*b)[:], val)
	return err
}
func (b *Zbmp40) UnmarshalJSON(buf []byte) error {
	val, err := bitmapUnmarshalJson(5, buf)
	copy((*b)[:], val)
	return err
}
func (b *Zbmp48) UnmarshalJSON(buf []byte) error {
	val, err := bitmapUnmarshalJson(6, buf)
	copy((*b)[:], val)
	return err
}
func (b *Zbmp56) UnmarshalJSON(buf []byte) error {
	val, err := bitmapUnmarshalJson(7, buf)
	copy((*b)[:], val)
	return err
}
func (b *Zbmp64) UnmarshalJSON(buf []byte) error {
	val, err := bitmapUnmarshalJson(8, buf)
	copy((*b)[:], val)
	return err
}

func (b *Zbmp8) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(1, buf)
	copy((*b)[:], val)
	return buf, err
}

func (b *Zbmp16) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(2, buf)
	copy((*b)[:], reverse(val))
	return buf, err
}
func (b *Zbmp24) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(3, buf)
	copy((*b)[:], reverse(val))
	return buf, err
}

func (b *Zbmp32) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(4, buf)
	copy((*b)[:], reverse(val))
	return buf, err
}

func (b *Zbmp40) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(5, buf)
	copy((*b)[:], reverse(val))
	return buf, err
}

func (b *Zbmp48) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(6, buf)
	copy((*b)[:], reverse(val))
	return buf, err
}

func (b *Zbmp56) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(7, buf)
	copy((*b)[:], reverse(val))
	return buf, err
}

func (b *Zbmp64) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := bytesUnmarshalZcl(8, buf)
	copy((*b)[:], reverse(val))
	return buf, err
}
