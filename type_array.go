package zcl

func arrayMarshalZcl(ln string, v []Val) ([]byte, error)           { return nil, ErrNotImpl }
func arrayUnmarshalZcl(ln string, b []byte) ([]Val, []byte, error) { return nil, nil, ErrNotImpl }

func structMarshalZcl(ln string, v []Val) ([]byte, error)           { return nil, ErrNotImpl }
func structUnmarshalZcl(ln string, b []byte) ([]Val, []byte, error) { return nil, nil, ErrNotImpl }

func arraySetMarshalZcl(ln string, v []Val) ([]byte, error)           { return nil, ErrNotImpl }
func arraySetUnmarshalZcl(ln string, b []byte) ([]Val, []byte, error) { return nil, nil, ErrNotImpl }

// Zarray is Array (0xffff = invalid). A/D = D
type Zarray []Val

func (a *Zarray) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := arrayUnmarshalZcl("2+sloc", buf)
	*a = Zarray(val)
	return buf, err
}
func (a Zarray) MarshalZcl() ([]byte, error) { return arrayMarshalZcl("2+sloc", []Val(a)) }
func (a *Zarray) Values() []Val              { return []Val{a} }
func (a *Zarray) ID() TypeID                 { return 72 }

//func (a Zarray) Valid() bool { return a != Zarray(65535) }

// Zstruct is Structure (0xffff = invalid). A/D = D
type Zstruct []Val

func (s *Zstruct) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := structUnmarshalZcl("2+sloc", buf)
	*s = Zstruct(val)
	return buf, err
}
func (s Zstruct) MarshalZcl() ([]byte, error) { return structMarshalZcl("2+sloc", []Val(s)) }
func (s *Zstruct) Values() []Val              { return []Val{s} }
func (s *Zstruct) ID() TypeID                 { return 76 }

//func (s Zstruct) Valid() bool { return s != Zstruct(65535) }

// Zset is Set (0xffff = invalid). A/D = D
type Zset []Val

func (s *Zset) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := arraySetUnmarshalZcl("sloc", buf)
	*s = Zset(val)
	return buf, err
}
func (s Zset) MarshalZcl() ([]byte, error) { return arraySetMarshalZcl("sloc", []Val(s)) }
func (s *Zset) Values() []Val              { return []Val{s} }
func (s *Zset) ID() TypeID                 { return 80 }

//func (s Zset) Valid() bool { return s != Zset(65535) }

// Zbag is Bag (0xffff = invalid). A/D = D
type Zbag []Val

func (b *Zbag) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := arrayUnmarshalZcl("sloc", buf)
	*b = Zbag(val)
	return buf, err
}
func (b Zbag) MarshalZcl() ([]byte, error) { return arrayMarshalZcl("sloc", []Val(b)) }
func (b *Zbag) Values() []Val              { return []Val{b} }
func (b *Zbag) ID() TypeID                 { return 81 }

//func (b Zbag) Valid() bool { return b != Zbag(65535) }
