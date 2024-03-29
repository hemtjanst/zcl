package foundation

import (
	"encoding/json"
	"fmt"
	"hemtjan.st/zcl"
)

type AttributeList []zcl.AttrID

func (v *AttributeList) Values() []zcl.Val {
	var val []zcl.Val
	for _, a := range *v {
		val = append(val, &a)
	}
	return val
}

func (v AttributeList) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	for _, a := range v {
		if tmp, err = a.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (v *AttributeList) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	var list []zcl.AttrID

	for len(b) > 0 {
		a := new(zcl.AttrID)
		if b, err = a.UnmarshalZcl(b); err != nil {
			return b, err
		}
		list = append(list, *a)
	}
	*v = list
	return b, nil
}

func (v *AttributeList) TypeID() zcl.TypeID { return new(zcl.Zlist).TypeID() }
func (v *AttributeList) Name() string       { return "Attribute list" }
func (v *AttributeList) String() string     { return fmt.Sprintf("%v", []zcl.AttrID(*v)) }

type ReadAttributes struct {
	// AttributeList to read
	AttributeList AttributeList
}

func (v *ReadAttributes) Values() []zcl.Val {
	return []zcl.Val{&v.AttributeList}
}

func (v ReadAttributes) ID() zcl.CommandID {
	return ReadAttributesCommand
}

func (v ReadAttributes) MarshalZcl() ([]byte, error) {
	return v.AttributeList.MarshalZcl()
}
func (v *ReadAttributes) UnmarshalZcl(b []byte) ([]byte, error) {
	return v.AttributeList.UnmarshalZcl(b)
}

func (v *ReadAttributes) AttributeListString() string {
	var str []string
	for _, a := range v.AttributeList {
		str = append(str, zcl.Sprintf("%04x", a))
	}
	return zcl.StrJoin(str, ", ")
}

func (v *ReadAttributes) String() string {
	var str []string
	str = append(str, "AttributeList["+v.AttributeListString()+"]")
	return "ReadAttributes{" + zcl.StrJoin(str, " ") + "}"
}

func (ReadAttributes) Name() string { return "Read Attributes" }

type ReadAttributeStatusRecord struct {
	AttributeID zcl.AttrID
	Status      zcl.Status
	DataType    zcl.TypeID
	Value       zcl.Val
}

func (v *ReadAttributeStatusRecord) StatusCode() zcl.Status { return v.Status }
func (v *ReadAttributeStatusRecord) AttrID() zcl.AttrID     { return v.AttributeID }
func (v *ReadAttributeStatusRecord) AttrType() zcl.TypeID   { return v.DataType }
func (v *ReadAttributeStatusRecord) AttrValue() zcl.Val     { return v.Value }

func (v ReadAttributeStatusRecord) String() string {
	return zcl.Sprintf(
		"ID[0x%04X] Status[%s] Type[%s] Value[%s]",
		v.AttributeID,
		v.Status,
		v.DataType,
		v.Value,
	)
}

func (v ReadAttributeStatusRecord) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.AttributeID.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if v.Status != zcl.Success {
		// Exit early if status is not success
		return data, nil
	}

	if tmp, err = v.DataType.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if v.Value == nil {
		return nil, zcl.Errorf("value is nil")
	}
	switch v.Value.(type) {
	case *zcl.Zarray:
		arr := v.Value.(*zcl.Zarray)
		if tmp, err = zcl.ArrayMarshalZcl("2", arr.Type, arr.Content); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	case *zcl.Zbag:
		arr := v.Value.(*zcl.Zbag)
		if tmp, err = zcl.ArrayMarshalZcl("2", arr.Type, arr.Content); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	case *zcl.Zset:
		arr := v.Value.(*zcl.Zset)
		if tmp, err = zcl.ArrayMarshalZcl("2", arr.Type, arr.Content); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	case *zcl.Zstruct:
		str := v.Value.(*zcl.Zstruct)
		if tmp, err = zcl.StructMarshalZcl("2", []zcl.StructField(*str)); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	default:
		if tmp, err = v.Value.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	return data, nil
}

func (v *ReadAttributeStatusRecord) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.AttributeID).UnmarshalZcl(b); err != nil {
		return b, err
	}
	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}
	if v.Status != zcl.Success {
		// Exit early if status is not success
		return b, nil
	}

	if b, err = (&v.DataType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	val := zcl.NewValue(uint8(v.DataType))
	if val == nil {
		return nil, zcl.Errorf("unknown datatype %d", v.DataType)
	}

	switch val.(type) {
	case *zcl.Zarray:
		arr := val.(*zcl.Zarray)
		if arr.Type, arr.Content, b, err = zcl.ArrayUnmarshalZcl("2", b); err != nil {
			return b, err
		}
	case *zcl.Zbag:
		arr := val.(*zcl.Zbag)
		if arr.Type, arr.Content, b, err = zcl.ArrayUnmarshalZcl("2", b); err != nil {
			return b, err
		}
	case *zcl.Zset:
		arr := val.(*zcl.Zset)
		if arr.Type, arr.Content, b, err = zcl.ArrayUnmarshalZcl("2", b); err != nil {
			return b, err
		}
	case *zcl.Zstruct:
		str := val.(*zcl.Zstruct)
		var fields []zcl.StructField
		if fields, b, err = zcl.StructUnmarshalZcl("2", b); err != nil {
			return b, err
		}
		*str = zcl.Zstruct(fields)
	default:
		if b, err = val.UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	v.Value = val

	if v.Value == nil {
		return nil, zcl.Errorf("value is nil")
	}

	return b, nil
}

type ReadAttributesResponse struct {
	// AttributeList to read
	Attributes []ReadAttributeStatusRecord
	hasError   bool
}

func (v ReadAttributesResponse) HasError() bool {
	return v.hasError
}

func (v *ReadAttributesResponse) Values() []zcl.Val {
	var val []zcl.Val
	for _, a := range v.Attributes {
		val = append(val, &a)
	}
	return val
}

func (v ReadAttributesResponse) ID() zcl.CommandID {
	return ReadAttributesResponseCommand
}

func (v ReadAttributesResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	for _, a := range v.Attributes {
		if tmp, err = a.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (v *ReadAttributesResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	v.Attributes = []ReadAttributeStatusRecord{}

	for len(b) > 0 {
		a := new(ReadAttributeStatusRecord)
		nb := b
		if b, err = a.UnmarshalZcl(b); err != nil {
			if len(v.Attributes) > 0 {
				v.hasError = true
				// We read at least one record, handle as success
				return nb, nil
			}
			return b, err
		}
		v.Attributes = append(v.Attributes, *a)
	}
	return b, nil
}

func (v ReadAttributesResponse) AttributesString() string {
	var str []string
	for _, a := range v.Attributes {
		str = append(str, zcl.Sprintf("%s", a))
	}
	return zcl.StrJoin(str, ", ")
}

func (v ReadAttributesResponse) String() string {
	var str []string
	str = append(str, "Attributes["+v.AttributesString()+"]")
	return "ReadAttributes{" + zcl.StrJoin(str, " ") + "}"
}

func (ReadAttributesResponse) Name() string { return "Read Attributes Response" }

type DiscoverAttributes struct {
	StartIndex zcl.Zu16
	MaxEntries zcl.Zu8
}

func (v *DiscoverAttributes) Values() []zcl.Val { return []zcl.Val{&v.StartIndex, &v.MaxEntries} }

func (v DiscoverAttributes) ID() zcl.CommandID {
	return DiscoverAttributesCommand
}

func (v DiscoverAttributes) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.StartIndex.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)
	if tmp, err = v.MaxEntries.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *DiscoverAttributes) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.StartIndex).UnmarshalZcl(b); err != nil {
		return b, err
	}
	if b, err = (&v.MaxEntries).UnmarshalZcl(b); err != nil {
		return b, err
	}
	return b, nil
}

func (v DiscoverAttributes) String() string {
	return zcl.Sprintf("StartIndex[%d] MaxEntries[%d]", v.StartIndex, v.MaxEntries)
}

func (DiscoverAttributes) Name() string { return "Discover Attributes" }

type DiscoveredAttribute struct {
	ID   zcl.AttrID
	Type zcl.TypeID
}

func (v DiscoveredAttribute) AttrID() zcl.AttrID   { return v.ID }
func (v DiscoveredAttribute) AttrType() zcl.TypeID { return v.Type }

func (v DiscoveredAttribute) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.ID.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)
	if tmp, err = v.Type.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)
	return data, nil
}

func (v *DiscoveredAttribute) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	if b, err = v.ID.UnmarshalZcl(b); err != nil {
		return b, err
	}
	if b, err = v.Type.UnmarshalZcl(b); err != nil {
		return b, err
	}
	return b, nil
}

func (v DiscoveredAttribute) String() string {
	return zcl.Sprintf("0x%04X(%s)", v.ID, v.Type.String())
}

type DiscoverAttributesResponse struct {
	DiscoveryComplete zcl.Zbool
	Attributes        []DiscoveredAttribute
}

func (v *DiscoverAttributesResponse) Values() []zcl.Val {
	val := []zcl.Val{&v.DiscoveryComplete}
	for _, a := range v.Attributes {
		val = append(val, &a)
	}
	return val
}

func (v DiscoverAttributesResponse) ID() zcl.CommandID {
	return DiscoverAttributesResponseCommand
}

func (v DiscoverAttributesResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.DiscoveryComplete.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)
	for _, a := range v.Attributes {
		if tmp, err = a.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	return data, nil
}

func (v *DiscoverAttributesResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.DiscoveryComplete).UnmarshalZcl(b); err != nil {
		return b, err
	}
	v.Attributes = []DiscoveredAttribute{}
	for len(b) > 0 {
		attr := new(DiscoveredAttribute)
		if b, err = attr.UnmarshalZcl(b); err != nil {
			return b, err
		}
		v.Attributes = append(v.Attributes, *attr)
	}
	return b, nil
}

func (v DiscoverAttributesResponse) String() string {
	var str []string
	for _, a := range v.Attributes {
		str = append(str, fmt.Sprintf("0x%04X(%s)", a.ID, a.Type.String()))
	}
	return zcl.Sprintf("DiscoveryComplete[%v] Attributes[%s]", v.DiscoveryComplete == 1, zcl.StrJoin(str, ", "))
}

func (DiscoverAttributesResponse) Name() string { return "Discover Attributes Response" }

type WriteAttributeRecord struct {
	AttributeID zcl.AttrID
	DataType    zcl.TypeID
	Value       zcl.Val
}

func (v *WriteAttributeRecord) AttrID() zcl.AttrID   { return v.AttributeID }
func (v *WriteAttributeRecord) AttrType() zcl.TypeID { return v.DataType }
func (v *WriteAttributeRecord) AttrValue() zcl.Val   { return v.Value }

func (v WriteAttributeRecord) String() string {
	return zcl.Sprintf(
		"ID[0x%04X] Type[%s] Value[%s]",
		v.AttributeID,
		v.DataType,
		v.Value,
	)
}

func (v WriteAttributeRecord) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.AttributeID.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.DataType.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if v.Value == nil {
		return nil, zcl.Errorf("value is nil")
	}
	switch v.Value.(type) {
	case *zcl.Zarray:
		arr := v.Value.(*zcl.Zarray)
		if tmp, err = zcl.ArrayMarshalZcl("2", arr.Type, arr.Content); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	case *zcl.Zbag:
		arr := v.Value.(*zcl.Zbag)
		if tmp, err = zcl.ArrayMarshalZcl("2", arr.Type, arr.Content); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	case *zcl.Zset:
		arr := v.Value.(*zcl.Zset)
		if tmp, err = zcl.ArrayMarshalZcl("2", arr.Type, arr.Content); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	case *zcl.Zstruct:
		str := v.Value.(*zcl.Zstruct)
		if tmp, err = zcl.StructMarshalZcl("2", []zcl.StructField(*str)); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	default:
		if tmp, err = v.Value.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	return data, nil
}

func (v *WriteAttributeRecord) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.AttributeID).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.DataType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	val := zcl.NewValue(uint8(v.DataType))
	if val == nil {
		return nil, zcl.Errorf("unknown datatype %d", v.DataType)
	}

	switch val.(type) {
	case *zcl.Zarray:
		arr := val.(*zcl.Zarray)
		if arr.Type, arr.Content, b, err = zcl.ArrayUnmarshalZcl("2", b); err != nil {
			return b, err
		}
	case *zcl.Zbag:
		arr := val.(*zcl.Zbag)
		if arr.Type, arr.Content, b, err = zcl.ArrayUnmarshalZcl("2", b); err != nil {
			return b, err
		}
	case *zcl.Zset:
		arr := val.(*zcl.Zset)
		if arr.Type, arr.Content, b, err = zcl.ArrayUnmarshalZcl("2", b); err != nil {
			return b, err
		}
	case *zcl.Zstruct:
		str := val.(*zcl.Zstruct)
		var fields []zcl.StructField
		if fields, b, err = zcl.StructUnmarshalZcl("2", b); err != nil {
			return b, err
		}
		*str = zcl.Zstruct(fields)
	default:
		if b, err = val.UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	v.Value = val

	return b, nil
}

func (v *WriteAttributeRecord) UnmarshalJSON(b []byte) error {
	a := &struct {
		AttributeID zcl.AttrID
		DataType    zcl.TypeID
		Value       string
	}{}
	if err := json.Unmarshal(b, a); err != nil {
		return err
	}

	v.AttributeID = a.AttributeID
	v.DataType = a.DataType
	v.Value = zcl.FromString(uint8(v.DataType), a.Value)
	return nil
}

type WriteAttributes struct {
	// AttributeList to write
	AttributeList []WriteAttributeRecord
}

func (v *WriteAttributes) Values() (o []zcl.Val) {
	for _, a := range v.AttributeList {
		o = append(o, &a)
	}
	return
}

func (v WriteAttributes) ID() zcl.CommandID {
	return WriteAttributesCommand
}

func (v WriteAttributes) MarshalZcl() ([]byte, error) {
	var b []byte
	var err error
	for _, e := range v.AttributeList {
		if b, err = e.MarshalZcl(); err != nil {
			return b, err
		}
	}
	return b, err
}
func (v *WriteAttributes) UnmarshalZcl(b []byte) (o []byte, err error) {
	o = b
	v.AttributeList = []WriteAttributeRecord{}
	for len(o) > 0 {
		nv := new(WriteAttributeRecord)
		if o, err = nv.UnmarshalZcl(o); err != nil {
			return
		}
		v.AttributeList = append(v.AttributeList, *nv)
	}
	return
}

func (v *WriteAttributes) AttributeListString() string {
	var str []string
	for _, a := range v.AttributeList {
		str = append(str, zcl.Sprintf("%04x", a))
	}
	return zcl.StrJoin(str, ", ")
}

func (v *WriteAttributes) String() string {
	var str []string
	str = append(str, "AttributeList["+v.AttributeListString()+"]")
	return "WriteAttributes{" + zcl.StrJoin(str, " ") + "}"
}

func (WriteAttributes) Name() string { return "Write Attributes" }

type WriteAttributeStatusRecord struct {
	Status      zcl.Status
	AttributeID zcl.AttrID
}

func (v *WriteAttributeStatusRecord) StatusCode() zcl.Status { return v.Status }
func (v *WriteAttributeStatusRecord) AttrID() zcl.AttrID     { return v.AttributeID }

func (v WriteAttributeStatusRecord) String() string {
	return zcl.Sprintf(
		"ID[0x%04X] Status[%s]",
		v.AttributeID,
		v.Status,
	)
}

func (v WriteAttributeStatusRecord) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.AttributeID.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *WriteAttributeStatusRecord) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}
	if len(b) == 0 {
		return nil, nil
	}
	if b, err = (&v.AttributeID).UnmarshalZcl(b); err != nil {
		return b, err
	}
	return b, nil
}

type WriteAttributesResponse struct {
	// AttributeList to read
	Attributes []WriteAttributeStatusRecord
	hasError   bool
}

func (v WriteAttributesResponse) HasError() bool {
	return v.hasError
}

func (v *WriteAttributesResponse) Values() []zcl.Val {
	var val []zcl.Val
	for _, a := range v.Attributes {
		val = append(val, &a)
	}
	return val
}

func (v WriteAttributesResponse) ID() zcl.CommandID {
	return WriteAttributesResponseCommand
}

func (v WriteAttributesResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	for _, a := range v.Attributes {
		if tmp, err = a.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (v *WriteAttributesResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	v.Attributes = []WriteAttributeStatusRecord{}

	for len(b) > 0 {
		a := new(WriteAttributeStatusRecord)
		nb := b
		if b, err = a.UnmarshalZcl(b); err != nil {
			if len(v.Attributes) > 0 {
				v.hasError = true
				// We read at least one record, handle as success
				return nb, nil
			}
			return b, err
		}
		v.Attributes = append(v.Attributes, *a)
	}
	return b, nil
}

func (v WriteAttributesResponse) AttributesString() string {
	var str []string
	for _, a := range v.Attributes {
		str = append(str, zcl.Sprintf("%s", a))
	}
	return zcl.StrJoin(str, ", ")
}

func (v WriteAttributesResponse) String() string {
	var str []string
	str = append(str, "Attributes["+v.AttributesString()+"]")
	return "WriteAttributesResponse{" + zcl.StrJoin(str, " ") + "}"
}

func (WriteAttributesResponse) Name() string { return "Write Attributes Response" }
