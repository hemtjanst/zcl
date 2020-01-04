package zcl

// Zenum8 is 8-bit enumeration (0xff = invalid). A/D = D
type Zenum8 uint8

func (e Zenum8) String() string { return Sprintf("0x%02X", uint8(e)) }

func (e *Zenum8) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(1, buf)
	*e = Zenum8(val)
	return buf, err
}
func (e Zenum8) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(1, uint64(e)) }
func (e *Zenum8) Values() []Val              { return []Val{e} }
func (e Zenum8) ID() TypeID                  { return 48 }

func (e Zenum8) Valid() bool { return e != Zenum8(0xFF) }

// Zenum16 is 16-bit enumeration (0xffff = invalid). A/D = D
type Zenum16 uint16

func (e Zenum16) String() string { return Sprintf("0x%04X", uint16(e)) }

func (e *Zenum16) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(2, buf)
	*e = Zenum16(val)
	return buf, err
}
func (e Zenum16) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(2, uint64(e)) }
func (e *Zenum16) Values() []Val              { return []Val{e} }
func (e Zenum16) ID() TypeID                  { return 49 }

func (e Zenum16) Valid() bool { return e != Zenum16(0xFFFF) }
