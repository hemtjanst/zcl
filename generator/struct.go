package codegen

import (
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var nativeTypes = map[string]string{
	"dat8":              "[]byte",
	"dat16":             "[]byte",
	"dat24":             "[]byte",
	"dat32":             "[]byte",
	"dat40":             "[]byte",
	"dat48":             "[]byte",
	"dat56":             "[]byte",
	"dat64":             "[]byte",
	"bool":              "uint8",
	"bmp8":              "[]byte",
	"bmp16":             "[]byte",
	"bmp24":             "[]byte",
	"bmp32":             "[]byte",
	"bmp40":             "[]byte",
	"bmp48":             "[]byte",
	"bmp56":             "[]byte",
	"bmp64":             "[]byte",
	"u8":                "uint8",
	"u16":               "uint16",
	"u24":               "uint32",
	"u32":               "uint32",
	"u40":               "uint64",
	"u48":               "uint64",
	"u56":               "uint64",
	"u64":               "uint64",
	"s8":                "int8",
	"s16":               "int16",
	"s24":               "int32",
	"s32":               "int32",
	"s40":               "int64",
	"s48":               "int64",
	"s56":               "int64",
	"s64":               "int64",
	"enum8":             "uint8",
	"enum16":            "uint16",
	"semi":              "float32",
	"float":             "float32",
	"double":            "float64",
	"ostring":           "[]byte",
	"cstring":           "string",
	"lostring":          "[]byte",
	"lcstring":          "string",
	"array":             "[]zcl.ZVal",
	"struct":            "[]zcl.ZVal",
	"set":               "[]zcl.ZVal",
	"bag":               "[]zcl.ZVal",
	"time":              "uint32",
	"date":              "zcl.Date",
	"utc":               "uint32",
	"cid":               "uint16",
	"aid":               "uint16",
	"oid":               "uint32",
	"uid":               "uint64",
	"seckey":            "[]byte",
	"SceneExtensionSet": "...zcl.SceneExtension",
}

type Root struct {
	// Package name
	Package string `xml:"-"`

	DataTypes []DataType    `xml:"datatypes>datatype"`
	Enum      []Enumeration `xml:"enumeration"`
	Domains   []Domain      `xml:"domain"`
	Profiles  []Profile     `xml:"profile"`
	Devices   []Device      `xml:"devices>device"`
}

type Command struct {
	ID          Hex    `xml:"id,attr" toml:"id,omitempty" yaml:"id,omitempty" json:"id,omitempty"`
	Name        Name   `xml:"name,attr" toml:"name,omitempty" yaml:"name,omitempty" json:"name,omitempty"`
	Desc        Desc   `xml:"description" toml:"description,omitempty" yaml:"description,omitempty" json:"description,omitempty"`
	Dir         string `xml:"dir,attr" toml:"dir,omitempty" yaml:"dir,omitempty" json:"dir,omitempty"`
	Required    string `xml:"required,attr" toml:"required,omitempty" yaml:"required,omitempty" json:"required,omitempty"`
	Response    Hex    `xml:"response,attr" toml:"response,omitempty" yaml:"response,omitempty" json:"response,omitempty"`
	ShowAs      string `xml:"showas,attr" toml:"showas,omitempty" yaml:"showas,omitempty" json:"showas,omitempty"`
	Vendor      string `xml:"vendor,attr" toml:"vendor,omitempty" yaml:"vendor,omitempty" json:"vendor,omitempty"`
	PayloadAttr []Attr `xml:"payload>attribute" toml:"payloadattr,omitempty" yaml:"payloadattr,omitempty" json:"payloadattr,omitempty"`
}

func (c Command) TypeCode(mnf MfCode) uint8 {
	var v uint8
	if mnf.Valid() {
		v = 4
	}

	return v
}

type Condition struct {
	Attr  Hex `toml:"attr,omitempty" yaml:"attr,omitempty" json:"attr,omitempty"`
	Value Hex `toml:"value,omitempty" yaml:"value,omitempty" json:"value,omitempty"`
	Mask  Hex `toml:"mask,omitempty" yaml:"mask,omitempty" json:"mask,omitempty"`
}

func (a Attr) Equal(o Attr) bool {
	if a.ID.Uint() != o.ID.Uint() || a.Name != o.Name || a.Type != o.Type {
		return false
	}
	if a.Access != o.Access || a.Default.Uint() != o.Default.Uint() || a.Required != o.Required {
		return false
	}
	if a.MfCode.Uint() != o.MfCode.Uint() || a.ShowAs != o.ShowAs || a.ListSize != o.ListSize {
		return false
	}
	if a.Enumeration != o.Enumeration || len(a.Values) != len(o.Values) {
		return false
	}
	if len(a.Values) > 0 {
		for i, v := range a.Values {
			ov, ok := o.Values[i]
			if !ok {
				return false
			}
			if ov != v {
				return false
			}
		}
	}
	return true
}

type Attr struct {
	ID          Hex             `xml:"id,attr" toml:"id,omitempty" yaml:"id,omitempty" json:"id,omitempty"`
	Name        Name            `xml:"name,attr" toml:"name,omitempty" yaml:"name,omitempty" json:"name,omitempty"`
	Type        Type            `xml:"type,attr" toml:"type,omitempty" yaml:"type,omitempty" json:"type,omitempty"`
	ArrayType   Type            `xml:"arrayType,attr" toml:"arraytype,omitempty" yaml:"arraytype,omitempty" json:"arraytype,omitempty"`
	Access      string          `xml:"access,attr" toml:"access,omitempty" yaml:"access,omitempty" json:"access,omitempty"`
	Default     Hex             `xml:"default,attr" toml:"default,omitempty" yaml:"default,omitempty" json:"default,omitempty"`
	Required    string          `xml:"required,attr" toml:"required,omitempty" yaml:"required,omitempty" json:"required,omitempty"`
	Report      bool            `xml:"-" toml:"report,omitempty" yaml:"report,omitempty" json:"report,omitempty"`
	Scene       Int             `xml:"-" toml:"scene,omitempty" yaml:"scene,omitempty" json:"scene,omitempty"`
	Range       string          `xml:"range,attr" toml:"range,omitempty" yaml:"range,omitempty" json:"range,omitempty"`
	MfCode      MfCode          `xml:"mfcode,attr" toml:"mnfcode,omitempty" yaml:"mnfcode,omitempty" json:"mnfcode,omitempty"`
	ShowAs      string          `xml:"showas,attr" toml:"showas,omitempty" yaml:"showas,omitempty" json:"showas,omitempty"`
	ListSize    string          `xml:"listSize,attr" toml:"listsize,omitempty" yaml:"listsize,omitempty" json:"listsize,omitempty"`
	Enumeration Name            `xml:"enumeration,attr" toml:"enumeration,omitempty" yaml:"enumeration,omitempty" json:"enumeration,omitempty"`
	Unit        string          `xml:"unit,attr" toml:"unit,omitempty" yaml:"unit,omitempty" json:"unit,omitempty"`
	Multiplier  Multiplier      `xml:"multiplier,attr" toml:"multiplier,omitempty" yaml:"multiplier,omitempty" json:"multiplier,omitempty"`
	Desc        Desc            `xml:"description" toml:"description,omitempty" yaml:"description,omitempty" json:"description,omitempty"`
	Values      map[string]Name `xml:"value" toml:"values,omitempty" yaml:"values,omitempty" json:"values,omitempty"`
	Bits        map[string]Name `xml:"bit" toml:"bits,omitempty" yaml:"bits,omitempty" json:"bits,omitempty"`
	Cond        []Condition     `xml:"payload>condition" toml:"cond,omitempty" yaml:"cond,omitempty" json:"cond,omitempty"`
}

func (a Attr) CanRead() bool {
	return a.Access == "rw" || a.Access == "r"
}
func (a Attr) CanWrite() bool {
	return a.Access == "rw" || a.Access == "w"
}

type AttrSet struct {
	ID     string `xml:"id,attr"` // Not hex since most ID:s are incomplete
	Desc   Desc   `xml:"description,attr"`
	MfCode MfCode `xml:"mfcode,attr"`
	Attr   []Attr `xml:"attribute"`
}

type CmdAttr struct {
	AttrSet []AttrSet `xml:"attribute-set"`
	Attr    []Attr    `xml:"attribute" toml:"attr,omitempty" yaml:"attr,omitempty" json:"attr,omitempty"`
	Command []Command `xml:"command" toml:"cmd,omitempty" yaml:"cmd,omitempty" json:"cmd,omitempty"`
}

func (c CmdAttr) SceneAttr() []Attr {
	var a []Attr

	for _, attr := range c.Attr {
		if attr.Scene.Valid() {
			a = append(a, attr)
		}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].Scene.Int() < a[j].Scene.Int()
	})

	return a
}

type Profile struct {
	ID      Hex         `xml:"id,attr"`
	Name    Name        `xml:"name,attr"`
	Desc    Desc        `xml:"description,attr"`
	Version string      `xml:"version,attr"`
	Rev     string      `xml:"rev,attr"`
	Icon    string      `xml:"icon,attr"`
	Domains []DomainRef `xml:"domain-ref"`
	Devices []Device    `xml:"device"`
}

type Device struct {
	ID   Hex    `xml:"id,attr"`
	Name Name   `xml:"name,attr"`
	Desc Desc   `xml:"description"`
	Icon string `xml:"icon,attr"`
}

type Cluster struct {
	ID     Hex     `xml:"id,attr" toml:"id,omitempty" yaml:"id,omitempty" json:"id,omitempty"`
	Name   Name    `xml:"name,attr" toml:"name,omitempty" yaml:"name,omitempty" json:"name,omitempty"`
	MfCode MfCode  `xml:"mfcode,attr" toml:"mnfcode,omitempty" yaml:"mnfcode,omitempty" json:"mnfcode,omitempty"`
	Desc   Desc    `xml:"description" toml:"description,omitempty" yaml:"description,omitempty" json:"description,omitempty"`
	Server CmdAttr `xml:"server" toml:"server,omitempty" yaml:"server,omitempty" json:"server,omitempty"`
	Client CmdAttr `xml:"client" toml:"client,omitempty" yaml:"client,omitempty" json:"client,omitempty"`
}

func (c Cluster) Command() []Command {
	return append(c.Server.Command, c.Client.Command...)
}

func (c Cluster) Attr() []Attr {
	attr := append(c.Server.Attr, c.Client.Attr...)
	for _, a := range append(c.Server.AttrSet, c.Client.AttrSet...) {
		attr = append(attr, a.Attr...)
	}
	return attr
}

type DomainRef struct {
	Name     Name `xml:"name,attr"`
	LowBound Hex  `xml:"low_bound,attr"`
}

func (d DomainRef) Resolve(r *Root) (cluster []Cluster) {
	for _, v := range r.Domains {
		if v.Name != d.Name {
			continue
		}
		for _, c := range v.Cluster {
			if c.ID.Uint() < d.LowBound.Uint() {
				continue
			}
			cluster = append(cluster, c)
		}
	}
	return
}

type Domain struct {
	Package   string    `xml:"package,attr"`
	Name      Name      `xml:"name,attr"`
	UseZCL    bool      `xml:"useZcl,attr"`
	Desc      Desc      `xml:"description,attr"`
	LowBound  Hex       `xml:"low_bound,attr"`
	HighBound Hex       `xml:"high_bound,attr"`
	Cluster   []Cluster `xml:"cluster"`
	Devices   []Device  `xml:"device"`
}

type Enumeration struct {
	ID     Hex         `xml:"id,attr"`
	Name   Name        `xml:"name,attr"`
	Values []EnumValue `xml:"value"`
}
type EnumValue struct {
	Value Hex  `xml:"value,attr"`
	Name  Name `xml:"name,attr"`
	Desc  Desc `xml:"description"`
}

type DataType struct {
	ID        Hex    `xml:"id,attr"`
	Name      Name   `xml:"name,attr"`
	ShortName Type   `xml:"shortname,attr"`
	Length    Length `xml:"length,attr"`
	AD        string `xml:"ad,attr"`
	Invalid   Hex    `xml:"inval,attr"`
}

func (d DataType) Native() string {
	if nt, ok := nativeTypes[string(d.ShortName)]; ok {
		if strings.HasPrefix(nt, "zcl.") {
			return nt[4:]
		}
		if strings.HasPrefix(nt, "[]zcl.") {
			return "[]" + nt[6:]
		}
		return nt
	}
	return ""
}

func (n Type) Native() string {
	if nt, ok := nativeTypes[string(n)]; ok {
		return nt
	}
	return "interface{}"
}

func (l Length) FixLen() bool {
	for _, v := range l {
		if v < '0' || v > '9' {
			return false
		}
	}
	return len(l) > 0
}

func (d DataType) marshalType() string {
	nt := d.Native()
	switch nt {
	case "[]ZVal":
		if d.ShortName == "struct" {
			return "struct"
		}
		if d.ShortName == "set" {
			return "arraySet"
		}
		return "array"
	case "[]byte":
		return "bytes"
	case "uint8", "uint16", "uint32", "uint64":
		return "uintLE"
	case "int8", "int16", "int32", "int64":
		return "intLE"
	case "string":
		return "string"
	case "float32":
		return "float32"
	case "float64":
		return "float64"
	}
	return "bytes"
}

func (d DataType) MarshalCast() string {
	nt := d.Native()
	switch nt {
	case "[]ZVal":
		return "[]ZVal"
	case "[]byte":
		return "[]byte"
	case "uint8", "uint16", "uint32", "uint64":
		return "uint64"
	case "int8", "int16", "int32", "int64":
		return "int64"
	case "string":
		return "string"
	case "float32":
		return "float32"
	case "float64":
		return "float64"
	}
	return "[]byte"
}

func (d DataType) MarshalFn() string {
	return d.marshalType() + "Marshal"
}

func (d DataType) UnmarshalFn() string {
	return d.marshalType() + "Unmarshal"
}

func (d DataType) Tag() string {
	return string(d.ShortName[0:1])
}

type Desc string
type Type string
type Name string
type Hex string
type MfCode string
type Length string
type Int string
type Multiplier string

func (i Int) Valid() bool {
	if i == "" {
		return false
	}
	_, err := strconv.Atoi(string(i))
	return err == nil
}

func (i Int) Int() int {
	v, _ := strconv.Atoi(string(i))
	return v
}

func (l Length) Uint() uint {
	v, _ := strconv.Atoi(string(l))
	return uint(v)
}

func (m MfCode) Trim() string  { return Hex(m).Trim() }
func (m MfCode) Valid() bool   { return Hex(m).Valid(16) }
func (m MfCode) Uint() uint64  { return Hex(m).Uint() }
func (m MfCode) Bytes() []byte { return Hex(m).Bytes(2) }
func (m MfCode) ByteStr() string {
	b := m.Bytes()
	if len(b) == 0 {
		return ""
	}
	return fmt.Sprintf("0x%02x, 0x%02x", b[0], b[1])
}
func (m MfCode) Name() string {
	switch m.Uint() {
	case 0x100b:
		return "Philips"
	case 0x1021:
		return "Legrand"
	case 0x10f2:
		return "Ubisys"
	case 0x117c:
		return "Ikea"
	case 0x115f:
		return "Xiaomi"
	case 0x1037:
		return "Eurotronic"
	default:
		return ""
	}
}
func (m MfCode) Stn(name Name) string {
	if !m.Valid() {
		return name.Fmt()
	}
	mf := m.Name()
	nn := name.Fmt()
	if mf == "" {
		return fmt.Sprintf("%s%04X", nn, m.Uint())
	}
	if !strings.HasPrefix(nn, mf) {
		nn = mf + nn
	}
	return nn
}

func (h Hex) Trim() string {
	v := strings.TrimSpace(string(h))
	if len(v) >= 2 && v[0:2] == "0x" {
		v = v[2:]
	}
	return v
}

func (h Hex) Valid(bits int) bool {
	v := h.Trim()
	if len(v) == 0 {
		return false
	}
	_, err := strconv.ParseUint(v, 16, bits)
	return err == nil
}

func (h Hex) Uint() uint64 {
	i, _ := strconv.ParseUint(h.Trim(), 16, 64)
	return i
}

func (h Hex) AsDecimal() uint64 {
	i, _ := strconv.ParseUint(h.Trim(), 10, 64)
	return i
}

func (h Hex) Hex() string {
	if !h.Valid(64) {
		return ""
	}
	return fmt.Sprintf("%x", h.Uint())
}

func (h Hex) Bytes(size int) []byte {
	v := h.Trim()
	if len(v) == 0 {
		return []byte{}
	}
	b, _ := hex.DecodeString(v)
	if len(b) > size {
		b = b[size-len(b):]
	}
	for len(b) < size {
		b = append([]byte{0x00}, b...)
	}
	return b
}

func (n Name) Pfx() string {
	o := strings.ToLower(string(n))
	uc := true
	var out []byte
	for i := 0; i < len(o); i++ {
		ch := o[i]
		if ch < 'a' || ch > 'z' {
			break
		}
		if uc {
			ch = ch & 0xDF
			uc = false
		}
		out = append(out, ch)
	}
	return string(out)
}

func (n Type) Type() string {
	if len(n) > 0 && (n[0]&0x60) == 0x40 {
		return string(n)
	}

	return fmt.Sprintf("Z%s", n)
	/*return fmt.Sprintf("%s%s",
		strings.ToUpper(string(n[0:1])),
		n[1:],
	)*/
}

func (n Name) FmtUs() string {
	o := strings.ToLower(string(n))
	uc := false
	var out []byte
	for i := 0; i < len(o); i++ {
		ch := o[i]
		if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
			if uc {
				out = append(out, '_')
				uc = false
			}
			out = append(out, ch)
			continue
		}
		uc = true
	}
	return string(out)
}

func (n Name) Fmt() string {
	o := strings.ToLower(string(n))
	uc := true
	var out []byte
	for i := 0; i < len(o); i++ {
		ch := o[i]
		if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
			if uc && (ch >= 'a' && ch <= 'z') {
				ch = ch & 0xDF
				uc = false
			}
			out = append(out, ch)
			continue
		}
		uc = true

	}
	return string(out)
}

func (d Desc) Trim() string {
	if len(d) == 0 {
		return ""
	}
	parts := strings.Split(string(d), "\n")
	for i, v := range parts {
		parts[i] = strings.TrimSpace(v)
	}
	return strings.Join(parts, " ")
}

func (d Desc) Comment() string {
	if len(d) == 0 {
		return ""
	}
	parts := strings.Split(string(d), "\n")
	var out []string
	for _, v := range parts {
		p := strings.TrimSpace(v)
		if p != "" {
			out = append(out, p)
		}
	}

	return fmt.Sprintf("// %s\n", strings.Join(out, "\n// "))

}
