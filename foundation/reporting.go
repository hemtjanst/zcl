package foundation

import (
	"encoding/json"
	"hemtjan.st/zcl"
)

type AttributeReportingConfig struct {
	Direction        zcl.Zenum8
	AttrID           zcl.AttrID
	Type             zcl.TypeID
	MinInterval      zcl.Zu16
	MaxInterval      zcl.Zu16
	ReportableChange zcl.Val
	TimeoutPeriod    zcl.Zu16
}

func (v *AttributeReportingConfig) UnmarshalJSON(b []byte) error {
	a := &struct {
		Direction        zcl.Zenum8
		AttrID           zcl.AttrID
		Type             zcl.TypeID
		MinInterval      zcl.Zu16
		MaxInterval      zcl.Zu16
		TimeoutPeriod    zcl.Zu16
		ReportableChange *string
	}{}
	if err := json.Unmarshal(b, a); err != nil {
		return err
	}

	v.Direction = a.Direction
	v.AttrID = a.AttrID
	v.Type = a.Type
	v.MinInterval = a.MinInterval
	v.MaxInterval = a.MaxInterval
	v.TimeoutPeriod = a.TimeoutPeriod

	if a.ReportableChange != nil {
		v.ReportableChange = zcl.FromString(uint8(v.Type), *a.ReportableChange)
	}
	return nil
}

func (v AttributeReportingConfig) String() string {
	if v.Direction == 1 {
		timeoutStr := zcl.Seconds.Format(float64(v.TimeoutPeriod))
		if v.TimeoutPeriod == 0 {
			timeoutStr = "never"
		}
		return zcl.Sprintf("Recv[0x%04X Timeout(%s)]", v.AttrID, timeoutStr)
	}
	maxIntStr := zcl.Seconds.Format(float64(v.MaxInterval))
	if v.MaxInterval == 0xFFFF {
		maxIntStr = "never"
	}

	reportChStr := ""
	if v.Type.Analog() && v.ReportableChange != nil {
		reportChStr = zcl.Sprintf(" MinChange(%v)", v.ReportableChange)
	}

	return zcl.Sprintf(
		"Send[0x%04X Type(%s) MinInt(%s) MaxInt(%s)%s]",
		v.AttrID,
		v.Type.String(),
		zcl.Seconds.Format(float64(v.MinInterval)),
		maxIntStr,
		reportChStr,
	)
}
func (v AttributeReportingConfig) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Direction.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.AttrID.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if v.Direction == 1 {
		if tmp, err = v.TimeoutPeriod.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
		return data, nil
	}

	if tmp, err = v.Type.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.MinInterval.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.MaxInterval.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if v.Type.Analog() {
		if tmp, err = v.ReportableChange.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	return data, nil
}
func (v *AttributeReportingConfig) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	if b, err = (&v.Direction).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	if b, err = (&v.AttrID).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	if v.Direction == 1 {
		if b, err = (&v.TimeoutPeriod).UnmarshalZcl(b); err != nil {
			return nil, err
		}
		return b, nil
	}

	if b, err = (&v.Type).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	if b, err = (&v.MinInterval).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	if b, err = (&v.MaxInterval).UnmarshalZcl(b); err != nil {
		return nil, err
	}

	if v.Type.Analog() {
		v.ReportableChange = zcl.NewValue(uint8(v.Type))
		if v.ReportableChange == nil {
			return nil, zcl.ErrInvalidType
		}

		if b, err = v.ReportableChange.UnmarshalZcl(b); err != nil {
			return nil, err
		}
	}
	return b, nil
}

type ConfigureReporting []AttributeReportingConfig

func (v ConfigureReporting) String() string {
	return zcl.Sprintf("ConfigureReporting{%v}", []AttributeReportingConfig(v))
}
func (v *ConfigureReporting) Values() (r []zcl.Val) {
	for _, c := range *v {
		r = append(r, &c)
	}
	return
}

func (ConfigureReporting) ID() zcl.CommandID { return ConfigureReportingCommand }
func (ConfigureReporting) Name() string      { return "ConfigureReporting" }

func (v ConfigureReporting) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct{ ReportingConfig []AttributeReportingConfig }{v})
}
func (v *ConfigureReporting) UnmarshalJSON(b []byte) error {
	rc := &struct{ ReportingConfig []AttributeReportingConfig }{}
	if err := json.Unmarshal(b, rc); err != nil {
		return err
	}
	*v = rc.ReportingConfig
	return nil
}

func (v ConfigureReporting) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	for _, c := range v {
		if tmp, err = c.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (v *ConfigureReporting) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	var list ConfigureReporting

	for len(b) > 0 {
		c := new(AttributeReportingConfig)
		if b, err = c.UnmarshalZcl(b); err != nil {
			return nil, err
		}
		list = append(list, *c)
	}

	*v = list
	return b, nil
}

type AttributeReportingConfigStatus struct {
	Status    zcl.Status
	Direction zcl.Zenum8
	AttrID    zcl.AttrID
}

func (v AttributeReportingConfigStatus) String() string {
	dirStr := "Send"
	if v.Direction == 1 {
		dirStr = "Recv"
	}
	return zcl.Sprintf(
		"%s[0x%04X %s]",
		dirStr,
		v.AttrID,
		v.Status.String(),
	)
}
func (v AttributeReportingConfigStatus) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Direction.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.AttrID.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}
func (v *AttributeReportingConfigStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	if len(b) == 0 {
		return nil, nil
	}
	if b, err = (&v.Direction).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	if b, err = (&v.AttrID).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	return b, nil
}

type ConfigureReportingResponse []AttributeReportingConfigStatus

func (v ConfigureReportingResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	for _, c := range v {
		if tmp, err = c.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (v *ConfigureReportingResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	var list ConfigureReportingResponse

	for len(b) > 0 {
		c := new(AttributeReportingConfigStatus)
		if b, err = c.UnmarshalZcl(b); err != nil {
			return nil, err
		}
		list = append(list, *c)
	}

	*v = list
	return b, nil
}

func (v *ConfigureReportingResponse) Values() (r []zcl.Val) {
	for _, c := range *v {
		r = append(r, &c)
	}
	return
}

func (ConfigureReportingResponse) ID() zcl.CommandID { return ConfigureReportingResponseCommand }
func (ConfigureReportingResponse) Name() string      { return "ConfigureReportingResponse" }
func (v ConfigureReportingResponse) String() string {
	return zcl.Sprintf("ConfigureReportingResponse{%v}", []AttributeReportingConfigStatus(v))
}

type ReadReportingConfigEntry struct {
	Direction zcl.Zenum8
	AttrID    zcl.AttrID
}

func (v ReadReportingConfigEntry) String() string {
	dirStr := "Send"
	if v.Direction == 1 {
		dirStr = "Recv"
	}
	return zcl.Sprintf(
		"%s[0x%04X]",
		dirStr,
		v.AttrID,
	)
}
func (v ReadReportingConfigEntry) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Direction.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.AttrID.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}
func (v *ReadReportingConfigEntry) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	if b, err = (&v.Direction).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	if b, err = (&v.AttrID).UnmarshalZcl(b); err != nil {
		return nil, err
	}
	return b, nil
}

type ReadReportingConfiguration []ReadReportingConfigEntry

func (v ReadReportingConfiguration) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	for _, c := range v {
		if tmp, err = c.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (v *ReadReportingConfiguration) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	var list ReadReportingConfiguration

	for len(b) > 0 {
		c := new(ReadReportingConfigEntry)
		if b, err = c.UnmarshalZcl(b); err != nil {
			return nil, err
		}
		list = append(list, *c)
	}

	*v = list
	return b, nil
}

func (v *ReadReportingConfiguration) Values() (r []zcl.Val) {
	for _, c := range *v {
		r = append(r, &c)
	}
	return
}

func (v ReadReportingConfiguration) String() string {
	return zcl.Sprintf("ReadReportingConfiguration{%v}", []ReadReportingConfigEntry(v))
}

func (ReadReportingConfiguration) ID() zcl.CommandID { return ReadReportingConfigurationCommand }
func (ReadReportingConfiguration) Name() string      { return "ReadReportingConfiguration" }

type ReadReportingConfigStatus struct {
	Status zcl.Status
	AttributeReportingConfig
}

func (v ReadReportingConfigStatus) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if v.Status == zcl.Success {
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)

	}
	nv := &AttributeReportingConfigStatus{
		Status:    v.Status,
		Direction: v.Direction,
		AttrID:    v.AttrID,
	}
	return nv.MarshalZcl()
}

func (v *ReadReportingConfigStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	if len(b) < 1 {
		return nil, zcl.ErrNotEnoughData
	}
	v.Status = zcl.Status(b[0])
	if v.Status == zcl.Success {
		return (&v.AttributeReportingConfig).UnmarshalZcl(b[1:])
	}
	nv := &AttributeReportingConfigStatus{}
	b, err = nv.UnmarshalZcl(b)
	v.Direction = nv.Direction
	v.AttrID = nv.AttrID
	return b, err
}

func (v ReadReportingConfigStatus) String() string {
	if v.Status == zcl.Success {
		return v.AttributeReportingConfig.String()
	}
	return AttributeReportingConfigStatus{
		Status:    v.Status,
		Direction: v.Direction,
		AttrID:    v.AttrID,
	}.String()
}

type ReadReportingConfigurationResponse []ReadReportingConfigStatus

func (v ReadReportingConfigurationResponse) String() string {
	return zcl.Sprintf("ReadReportingConfigurationResponse{%v}", []ReadReportingConfigStatus(v))
}
func (v *ReadReportingConfigurationResponse) Values() (r []zcl.Val) {
	for _, c := range *v {
		r = append(r, &c)
	}
	return
}

func (ReadReportingConfigurationResponse) ID() zcl.CommandID {
	return ReadReportingConfigurationResponseCommand
}
func (ReadReportingConfigurationResponse) Name() string { return "ReadReportingConfigurationResponse" }

func (v ReadReportingConfigurationResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	for _, c := range v {
		if tmp, err = c.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

func (v *ReadReportingConfigurationResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	var list ReadReportingConfigurationResponse

	for len(b) > 0 {
		c := new(ReadReportingConfigStatus)
		if b, err = c.UnmarshalZcl(b); err != nil {
			return nil, err
		}
		list = append(list, *c)
	}

	*v = list
	return b, nil
}

type ReportAttributeRecord struct {
	AttributeID zcl.AttrID
	DataType    zcl.TypeID
	Value       zcl.Val
}

func (v *ReportAttributeRecord) AttrID() zcl.AttrID   { return v.AttributeID }
func (v *ReportAttributeRecord) AttrType() zcl.TypeID { return v.DataType }
func (v *ReportAttributeRecord) AttrValue() zcl.Val   { return v.Value }

func (v ReportAttributeRecord) String() string {
	return zcl.Sprintf(
		"Attr[0x%04X = %v(%s)]",
		v.AttributeID,
		v.Value,
		v.DataType,
	)
}

func (v ReportAttributeRecord) MarshalZcl() ([]byte, error) {
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

func (v *ReportAttributeRecord) UnmarshalZcl(b []byte) ([]byte, error) {
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

type ReportAttributes struct {
	// AttributeList to read
	Attributes []ReportAttributeRecord
}

func (v *ReportAttributes) Values() []zcl.Val {
	var val []zcl.Val
	for _, a := range v.Attributes {
		val = append(val, &a)
	}
	return val
}

func (v ReportAttributes) ID() zcl.CommandID {
	return ReportAttributesCommand
}

func (v ReportAttributes) MarshalZcl() ([]byte, error) {
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

func (v *ReportAttributes) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	var list []ReportAttributeRecord

	for len(b) > 0 {
		a := new(ReportAttributeRecord)
		if b, err = a.UnmarshalZcl(b); err != nil {
			return b, err
		}
		list = append(list, *a)
	}
	v.Attributes = list
	return b, nil
}

func (v ReportAttributes) AttributesString() string {
	var str []string
	for _, a := range v.Attributes {
		str = append(str, zcl.Sprintf("%v", a))
	}
	return zcl.StrJoin(str, ", ")
}

func (v ReportAttributes) String() string {
	return "ReportAttributes{" + v.AttributesString() + "}"
}

func (ReportAttributes) Name() string { return "Read Attributes Response" }
