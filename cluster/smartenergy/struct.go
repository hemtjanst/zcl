package smartenergy

import "hemtjan.st/zcl"

type CommandID = zcl.CommandID
type Frame = zcl.ReceivedZclFrame

const ControlTemperatureAttr zcl.AttrID = 25

func (ControlTemperature) ID() zcl.AttrID   { return ControlTemperatureAttr }
func (ControlTemperature) Readable() bool   { return false }
func (ControlTemperature) Writable() bool   { return false }
func (ControlTemperature) Reportable() bool { return false }
func (ControlTemperature) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ControlTemperature) AttrID() zcl.AttrID   { return v.ID() }
func (v ControlTemperature) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ControlTemperature) AttrValue() zcl.Val  { return v.Value() }

func (ControlTemperature) Name() string        { return `Control Temperature` }
func (ControlTemperature) Description() string { return `` }

type ControlTemperature zcl.Zs24

func (v *ControlTemperature) TypeID() zcl.TypeID { return new(zcl.Zs24).TypeID() }
func (v *ControlTemperature) Value() zcl.Val     { return v }

func (v ControlTemperature) MarshalZcl() ([]byte, error) { return zcl.Zs24(v).MarshalZcl() }

func (v *ControlTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs24)
	br, err := nt.UnmarshalZcl(b)
	*v = ControlTemperature(*nt)
	return br, err
}

func (v ControlTemperature) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs24(v))
}

func (v *ControlTemperature) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs24)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ControlTemperature(*a)
	return nil
}

func (v *ControlTemperature) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs24); ok {
		*v = ControlTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ControlTemperature) String() string {
	return zcl.Sprintf("%v", zcl.Zs24(v))
}

const CurrentBlockAttr zcl.AttrID = 14

func (CurrentBlock) ID() zcl.AttrID   { return CurrentBlockAttr }
func (CurrentBlock) Readable() bool   { return false }
func (CurrentBlock) Writable() bool   { return false }
func (CurrentBlock) Reportable() bool { return false }
func (CurrentBlock) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentBlock) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentBlock) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentBlock) AttrValue() zcl.Val  { return v.Value() }

func (CurrentBlock) Name() string { return `Current Block` }
func (CurrentBlock) Description() string {
	return `is an 8-bit Enumeration which indicates the currently active block. If blocks are active then the current
active block is based on the CurrentBlockPeriodConsumptionDelivered and the block thresholds.
Block 1 is active when the value of CurrentBlockPeriodConsumptionDelivered is less than Block1Threshold value;
Block 2 is active when CurrentBlockPeriodConsumptionDelivered is greater than Block1Threshold value and less than
Block2Threshold value, and so on. Block 16 is active when the value of CurrentBlockPeriodConsumptionDelivered is
greater than Block15Threshold value.`
}

// CurrentBlock is an 8-bit Enumeration which indicates the currently active block. If blocks are active then the current
// active block is based on the CurrentBlockPeriodConsumptionDelivered and the block thresholds.
// Block 1 is active when the value of CurrentBlockPeriodConsumptionDelivered is less than Block1Threshold value;
// Block 2 is active when CurrentBlockPeriodConsumptionDelivered is greater than Block1Threshold value and less than
// Block2Threshold value, and so on. Block 16 is active when the value of CurrentBlockPeriodConsumptionDelivered is
// greater than Block15Threshold value.
type CurrentBlock zcl.Zenum8

func (v *CurrentBlock) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *CurrentBlock) Value() zcl.Val     { return v }

func (v CurrentBlock) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *CurrentBlock) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentBlock(*nt)
	return br, err
}

func (v CurrentBlock) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *CurrentBlock) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentBlock(*a)
	return nil
}

func (v *CurrentBlock) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = CurrentBlock(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentBlock) String() string {
	switch v {
	case 0x00:
		return "No blocks in use"
	case 0x01:
		return "Block1"
	case 0x02:
		return "Block2"
	case 0x03:
		return "Block3"
	case 0x04:
		return "Block4"
	case 0x05:
		return "Block5"
	case 0x06:
		return "Block6"
	case 0x07:
		return "Block7"
	case 0x08:
		return "Block8"
	case 0x09:
		return "Block9"
	case 0x0A:
		return "Block10"
	case 0x0B:
		return "Block11"
	case 0x0C:
		return "Block12"
	case 0x0D:
		return "Block13"
	case 0x0E:
		return "Block14"
	case 0x0F:
		return "Block15"
	case 0x10:
		return "Block16"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v CurrentBlock) IsNoBlocksInUse() bool { return v == 0x00 }
func (v CurrentBlock) IsBlock1() bool        { return v == 0x01 }
func (v CurrentBlock) IsBlock2() bool        { return v == 0x02 }
func (v CurrentBlock) IsBlock3() bool        { return v == 0x03 }
func (v CurrentBlock) IsBlock4() bool        { return v == 0x04 }
func (v CurrentBlock) IsBlock5() bool        { return v == 0x05 }
func (v CurrentBlock) IsBlock6() bool        { return v == 0x06 }
func (v CurrentBlock) IsBlock7() bool        { return v == 0x07 }
func (v CurrentBlock) IsBlock8() bool        { return v == 0x08 }
func (v CurrentBlock) IsBlock9() bool        { return v == 0x09 }
func (v CurrentBlock) IsBlock10() bool       { return v == 0x0A }
func (v CurrentBlock) IsBlock11() bool       { return v == 0x0B }
func (v CurrentBlock) IsBlock12() bool       { return v == 0x0C }
func (v CurrentBlock) IsBlock13() bool       { return v == 0x0D }
func (v CurrentBlock) IsBlock14() bool       { return v == 0x0E }
func (v CurrentBlock) IsBlock15() bool       { return v == 0x0F }
func (v CurrentBlock) IsBlock16() bool       { return v == 0x10 }
func (v *CurrentBlock) SetNoBlocksInUse()    { *v = 0x00 }
func (v *CurrentBlock) SetBlock1()           { *v = 0x01 }
func (v *CurrentBlock) SetBlock2()           { *v = 0x02 }
func (v *CurrentBlock) SetBlock3()           { *v = 0x03 }
func (v *CurrentBlock) SetBlock4()           { *v = 0x04 }
func (v *CurrentBlock) SetBlock5()           { *v = 0x05 }
func (v *CurrentBlock) SetBlock6()           { *v = 0x06 }
func (v *CurrentBlock) SetBlock7()           { *v = 0x07 }
func (v *CurrentBlock) SetBlock8()           { *v = 0x08 }
func (v *CurrentBlock) SetBlock9()           { *v = 0x09 }
func (v *CurrentBlock) SetBlock10()          { *v = 0x0A }
func (v *CurrentBlock) SetBlock11()          { *v = 0x0B }
func (v *CurrentBlock) SetBlock12()          { *v = 0x0C }
func (v *CurrentBlock) SetBlock13()          { *v = 0x0D }
func (v *CurrentBlock) SetBlock14()          { *v = 0x0E }
func (v *CurrentBlock) SetBlock15()          { *v = 0x0F }
func (v *CurrentBlock) SetBlock16()          { *v = 0x10 }

func (CurrentBlock) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "No blocks in use"},
		{Value: 0x01, Name: "Block1"},
		{Value: 0x02, Name: "Block2"},
		{Value: 0x03, Name: "Block3"},
		{Value: 0x04, Name: "Block4"},
		{Value: 0x05, Name: "Block5"},
		{Value: 0x06, Name: "Block6"},
		{Value: 0x07, Name: "Block7"},
		{Value: 0x08, Name: "Block8"},
		{Value: 0x09, Name: "Block9"},
		{Value: 0x0A, Name: "Block10"},
		{Value: 0x0B, Name: "Block11"},
		{Value: 0x0C, Name: "Block12"},
		{Value: 0x0D, Name: "Block13"},
		{Value: 0x0E, Name: "Block14"},
		{Value: 0x0F, Name: "Block15"},
		{Value: 0x10, Name: "Block16"},
	}
}

const CurrentBlockPeriodConsumptionDeliveredAttr zcl.AttrID = 12

func (CurrentBlockPeriodConsumptionDelivered) ID() zcl.AttrID {
	return CurrentBlockPeriodConsumptionDeliveredAttr
}
func (CurrentBlockPeriodConsumptionDelivered) Readable() bool   { return false }
func (CurrentBlockPeriodConsumptionDelivered) Writable() bool   { return false }
func (CurrentBlockPeriodConsumptionDelivered) Reportable() bool { return false }
func (CurrentBlockPeriodConsumptionDelivered) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentBlockPeriodConsumptionDelivered) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentBlockPeriodConsumptionDelivered) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentBlockPeriodConsumptionDelivered) AttrValue() zcl.Val  { return v.Value() }

func (CurrentBlockPeriodConsumptionDelivered) Name() string {
	return `Current Block Period Consumption Delivered`
}
func (CurrentBlockPeriodConsumptionDelivered) Description() string {
	return `represents the most recent summed value of Energy, Gas or Water delivered and consumed in the premises during the
Block Tariff Period. The CurrentBlockPeriodConsumptionDelivered is reset at the start of each Block Tariff Period.`
}

// CurrentBlockPeriodConsumptionDelivered represents the most recent summed value of Energy, Gas or Water delivered and consumed in the premises during the
// Block Tariff Period. The CurrentBlockPeriodConsumptionDelivered is reset at the start of each Block Tariff Period.
type CurrentBlockPeriodConsumptionDelivered zcl.Zu48

func (v *CurrentBlockPeriodConsumptionDelivered) TypeID() zcl.TypeID { return new(zcl.Zu48).TypeID() }
func (v *CurrentBlockPeriodConsumptionDelivered) Value() zcl.Val     { return v }

func (v CurrentBlockPeriodConsumptionDelivered) MarshalZcl() ([]byte, error) {
	return zcl.Zu48(v).MarshalZcl()
}

func (v *CurrentBlockPeriodConsumptionDelivered) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu48)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentBlockPeriodConsumptionDelivered(*nt)
	return br, err
}

func (v CurrentBlockPeriodConsumptionDelivered) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu48(v))
}

func (v *CurrentBlockPeriodConsumptionDelivered) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu48)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentBlockPeriodConsumptionDelivered(*a)
	return nil
}

func (v *CurrentBlockPeriodConsumptionDelivered) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu48); ok {
		*v = CurrentBlockPeriodConsumptionDelivered(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentBlockPeriodConsumptionDelivered) String() string {
	return zcl.Sprintf("%v", zcl.Zu48(v))
}

const CurrentInletEnergyCarrierDemandAttr zcl.AttrID = 26

func (CurrentInletEnergyCarrierDemand) ID() zcl.AttrID   { return CurrentInletEnergyCarrierDemandAttr }
func (CurrentInletEnergyCarrierDemand) Readable() bool   { return false }
func (CurrentInletEnergyCarrierDemand) Writable() bool   { return false }
func (CurrentInletEnergyCarrierDemand) Reportable() bool { return false }
func (CurrentInletEnergyCarrierDemand) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentInletEnergyCarrierDemand) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentInletEnergyCarrierDemand) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentInletEnergyCarrierDemand) AttrValue() zcl.Val  { return v.Value() }

func (CurrentInletEnergyCarrierDemand) Name() string        { return `Current Inlet Energy Carrier Demand` }
func (CurrentInletEnergyCarrierDemand) Description() string { return `` }

type CurrentInletEnergyCarrierDemand zcl.Zs24

func (v *CurrentInletEnergyCarrierDemand) TypeID() zcl.TypeID { return new(zcl.Zs24).TypeID() }
func (v *CurrentInletEnergyCarrierDemand) Value() zcl.Val     { return v }

func (v CurrentInletEnergyCarrierDemand) MarshalZcl() ([]byte, error) {
	return zcl.Zs24(v).MarshalZcl()
}

func (v *CurrentInletEnergyCarrierDemand) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs24)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentInletEnergyCarrierDemand(*nt)
	return br, err
}

func (v CurrentInletEnergyCarrierDemand) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs24(v))
}

func (v *CurrentInletEnergyCarrierDemand) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs24)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentInletEnergyCarrierDemand(*a)
	return nil
}

func (v *CurrentInletEnergyCarrierDemand) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs24); ok {
		*v = CurrentInletEnergyCarrierDemand(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentInletEnergyCarrierDemand) String() string {
	return zcl.Sprintf("%v", zcl.Zs24(v))
}

const CurrentInletEnergyCarrierSummationAttr zcl.AttrID = 21

func (CurrentInletEnergyCarrierSummation) ID() zcl.AttrID {
	return CurrentInletEnergyCarrierSummationAttr
}
func (CurrentInletEnergyCarrierSummation) Readable() bool   { return false }
func (CurrentInletEnergyCarrierSummation) Writable() bool   { return false }
func (CurrentInletEnergyCarrierSummation) Reportable() bool { return false }
func (CurrentInletEnergyCarrierSummation) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentInletEnergyCarrierSummation) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentInletEnergyCarrierSummation) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentInletEnergyCarrierSummation) AttrValue() zcl.Val  { return v.Value() }

func (CurrentInletEnergyCarrierSummation) Name() string {
	return `Current Inlet Energy Carrier Summation`
}
func (CurrentInletEnergyCarrierSummation) Description() string { return `` }

type CurrentInletEnergyCarrierSummation zcl.Zu48

func (v *CurrentInletEnergyCarrierSummation) TypeID() zcl.TypeID { return new(zcl.Zu48).TypeID() }
func (v *CurrentInletEnergyCarrierSummation) Value() zcl.Val     { return v }

func (v CurrentInletEnergyCarrierSummation) MarshalZcl() ([]byte, error) {
	return zcl.Zu48(v).MarshalZcl()
}

func (v *CurrentInletEnergyCarrierSummation) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu48)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentInletEnergyCarrierSummation(*nt)
	return br, err
}

func (v CurrentInletEnergyCarrierSummation) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu48(v))
}

func (v *CurrentInletEnergyCarrierSummation) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu48)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentInletEnergyCarrierSummation(*a)
	return nil
}

func (v *CurrentInletEnergyCarrierSummation) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu48); ok {
		*v = CurrentInletEnergyCarrierSummation(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentInletEnergyCarrierSummation) String() string {
	return zcl.Sprintf("%v", zcl.Zu48(v))
}

const CurrentMaxDemandDeliveredAttr zcl.AttrID = 2

func (CurrentMaxDemandDelivered) ID() zcl.AttrID   { return CurrentMaxDemandDeliveredAttr }
func (CurrentMaxDemandDelivered) Readable() bool   { return false }
func (CurrentMaxDemandDelivered) Writable() bool   { return false }
func (CurrentMaxDemandDelivered) Reportable() bool { return false }
func (CurrentMaxDemandDelivered) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentMaxDemandDelivered) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentMaxDemandDelivered) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentMaxDemandDelivered) AttrValue() zcl.Val  { return v.Value() }

func (CurrentMaxDemandDelivered) Name() string { return `Current Max Demand Delivered` }
func (CurrentMaxDemandDelivered) Description() string {
	return `represents the maximum demand or rate of delivered value of Energy, Gas, or Water being utilized at the premises`
}

// CurrentMaxDemandDelivered represents the maximum demand or rate of delivered value of Energy, Gas, or Water being utilized at the premises
type CurrentMaxDemandDelivered zcl.Zu48

func (v *CurrentMaxDemandDelivered) TypeID() zcl.TypeID { return new(zcl.Zu48).TypeID() }
func (v *CurrentMaxDemandDelivered) Value() zcl.Val     { return v }

func (v CurrentMaxDemandDelivered) MarshalZcl() ([]byte, error) { return zcl.Zu48(v).MarshalZcl() }

func (v *CurrentMaxDemandDelivered) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu48)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentMaxDemandDelivered(*nt)
	return br, err
}

func (v CurrentMaxDemandDelivered) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu48(v))
}

func (v *CurrentMaxDemandDelivered) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu48)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentMaxDemandDelivered(*a)
	return nil
}

func (v *CurrentMaxDemandDelivered) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu48); ok {
		*v = CurrentMaxDemandDelivered(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentMaxDemandDelivered) String() string {
	return zcl.Sprintf("%v", zcl.Zu48(v))
}

const CurrentMaxDemandDeliveredTimeAttr zcl.AttrID = 8

func (CurrentMaxDemandDeliveredTime) ID() zcl.AttrID   { return CurrentMaxDemandDeliveredTimeAttr }
func (CurrentMaxDemandDeliveredTime) Readable() bool   { return false }
func (CurrentMaxDemandDeliveredTime) Writable() bool   { return false }
func (CurrentMaxDemandDeliveredTime) Reportable() bool { return false }
func (CurrentMaxDemandDeliveredTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentMaxDemandDeliveredTime) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentMaxDemandDeliveredTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentMaxDemandDeliveredTime) AttrValue() zcl.Val  { return v.Value() }

func (CurrentMaxDemandDeliveredTime) Name() string { return `Current Max Demand Delivered Time` }
func (CurrentMaxDemandDeliveredTime) Description() string {
	return `represents the time when CurrentMaxDemandDelivered reading was captured`
}

// CurrentMaxDemandDeliveredTime represents the time when CurrentMaxDemandDelivered reading was captured
type CurrentMaxDemandDeliveredTime zcl.Zutc

func (v *CurrentMaxDemandDeliveredTime) TypeID() zcl.TypeID { return new(zcl.Zutc).TypeID() }
func (v *CurrentMaxDemandDeliveredTime) Value() zcl.Val     { return v }

func (v CurrentMaxDemandDeliveredTime) MarshalZcl() ([]byte, error) { return zcl.Zutc(v).MarshalZcl() }

func (v *CurrentMaxDemandDeliveredTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentMaxDemandDeliveredTime(*nt)
	return br, err
}

func (v CurrentMaxDemandDeliveredTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zutc(v))
}

func (v *CurrentMaxDemandDeliveredTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zutc)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentMaxDemandDeliveredTime(*a)
	return nil
}

func (v *CurrentMaxDemandDeliveredTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zutc); ok {
		*v = CurrentMaxDemandDeliveredTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentMaxDemandDeliveredTime) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(v))
}

const CurrentMaxDemandReceivedAttr zcl.AttrID = 3

func (CurrentMaxDemandReceived) ID() zcl.AttrID   { return CurrentMaxDemandReceivedAttr }
func (CurrentMaxDemandReceived) Readable() bool   { return false }
func (CurrentMaxDemandReceived) Writable() bool   { return false }
func (CurrentMaxDemandReceived) Reportable() bool { return false }
func (CurrentMaxDemandReceived) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentMaxDemandReceived) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentMaxDemandReceived) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentMaxDemandReceived) AttrValue() zcl.Val  { return v.Value() }

func (CurrentMaxDemandReceived) Name() string { return `Current Max Demand Received` }
func (CurrentMaxDemandReceived) Description() string {
	return `represents the maximum demand or rate of received value of Energy, Gas, or Water being utilized by the utility`
}

// CurrentMaxDemandReceived represents the maximum demand or rate of received value of Energy, Gas, or Water being utilized by the utility
type CurrentMaxDemandReceived zcl.Zu48

func (v *CurrentMaxDemandReceived) TypeID() zcl.TypeID { return new(zcl.Zu48).TypeID() }
func (v *CurrentMaxDemandReceived) Value() zcl.Val     { return v }

func (v CurrentMaxDemandReceived) MarshalZcl() ([]byte, error) { return zcl.Zu48(v).MarshalZcl() }

func (v *CurrentMaxDemandReceived) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu48)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentMaxDemandReceived(*nt)
	return br, err
}

func (v CurrentMaxDemandReceived) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu48(v))
}

func (v *CurrentMaxDemandReceived) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu48)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentMaxDemandReceived(*a)
	return nil
}

func (v *CurrentMaxDemandReceived) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu48); ok {
		*v = CurrentMaxDemandReceived(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentMaxDemandReceived) String() string {
	return zcl.Sprintf("%v", zcl.Zu48(v))
}

const CurrentMaxDemandReceivedTimeAttr zcl.AttrID = 9

func (CurrentMaxDemandReceivedTime) ID() zcl.AttrID   { return CurrentMaxDemandReceivedTimeAttr }
func (CurrentMaxDemandReceivedTime) Readable() bool   { return false }
func (CurrentMaxDemandReceivedTime) Writable() bool   { return false }
func (CurrentMaxDemandReceivedTime) Reportable() bool { return false }
func (CurrentMaxDemandReceivedTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentMaxDemandReceivedTime) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentMaxDemandReceivedTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentMaxDemandReceivedTime) AttrValue() zcl.Val  { return v.Value() }

func (CurrentMaxDemandReceivedTime) Name() string { return `Current Max Demand Received Time` }
func (CurrentMaxDemandReceivedTime) Description() string {
	return `represents the time when CurrentMaxDemandReceived reading was captured`
}

// CurrentMaxDemandReceivedTime represents the time when CurrentMaxDemandReceived reading was captured
type CurrentMaxDemandReceivedTime zcl.Zutc

func (v *CurrentMaxDemandReceivedTime) TypeID() zcl.TypeID { return new(zcl.Zutc).TypeID() }
func (v *CurrentMaxDemandReceivedTime) Value() zcl.Val     { return v }

func (v CurrentMaxDemandReceivedTime) MarshalZcl() ([]byte, error) { return zcl.Zutc(v).MarshalZcl() }

func (v *CurrentMaxDemandReceivedTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentMaxDemandReceivedTime(*nt)
	return br, err
}

func (v CurrentMaxDemandReceivedTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zutc(v))
}

func (v *CurrentMaxDemandReceivedTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zutc)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentMaxDemandReceivedTime(*a)
	return nil
}

func (v *CurrentMaxDemandReceivedTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zutc); ok {
		*v = CurrentMaxDemandReceivedTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentMaxDemandReceivedTime) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(v))
}

const CurrentOutletEnergyCarrierDemandAttr zcl.AttrID = 27

func (CurrentOutletEnergyCarrierDemand) ID() zcl.AttrID   { return CurrentOutletEnergyCarrierDemandAttr }
func (CurrentOutletEnergyCarrierDemand) Readable() bool   { return false }
func (CurrentOutletEnergyCarrierDemand) Writable() bool   { return false }
func (CurrentOutletEnergyCarrierDemand) Reportable() bool { return false }
func (CurrentOutletEnergyCarrierDemand) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentOutletEnergyCarrierDemand) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentOutletEnergyCarrierDemand) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentOutletEnergyCarrierDemand) AttrValue() zcl.Val  { return v.Value() }

func (CurrentOutletEnergyCarrierDemand) Name() string        { return `Current Outlet Energy Carrier Demand` }
func (CurrentOutletEnergyCarrierDemand) Description() string { return `` }

type CurrentOutletEnergyCarrierDemand zcl.Zs24

func (v *CurrentOutletEnergyCarrierDemand) TypeID() zcl.TypeID { return new(zcl.Zs24).TypeID() }
func (v *CurrentOutletEnergyCarrierDemand) Value() zcl.Val     { return v }

func (v CurrentOutletEnergyCarrierDemand) MarshalZcl() ([]byte, error) {
	return zcl.Zs24(v).MarshalZcl()
}

func (v *CurrentOutletEnergyCarrierDemand) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs24)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentOutletEnergyCarrierDemand(*nt)
	return br, err
}

func (v CurrentOutletEnergyCarrierDemand) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs24(v))
}

func (v *CurrentOutletEnergyCarrierDemand) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs24)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentOutletEnergyCarrierDemand(*a)
	return nil
}

func (v *CurrentOutletEnergyCarrierDemand) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs24); ok {
		*v = CurrentOutletEnergyCarrierDemand(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentOutletEnergyCarrierDemand) String() string {
	return zcl.Sprintf("%v", zcl.Zs24(v))
}

const CurrentOutletEnergyCarrierSummationAttr zcl.AttrID = 22

func (CurrentOutletEnergyCarrierSummation) ID() zcl.AttrID {
	return CurrentOutletEnergyCarrierSummationAttr
}
func (CurrentOutletEnergyCarrierSummation) Readable() bool   { return false }
func (CurrentOutletEnergyCarrierSummation) Writable() bool   { return false }
func (CurrentOutletEnergyCarrierSummation) Reportable() bool { return false }
func (CurrentOutletEnergyCarrierSummation) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentOutletEnergyCarrierSummation) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentOutletEnergyCarrierSummation) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentOutletEnergyCarrierSummation) AttrValue() zcl.Val  { return v.Value() }

func (CurrentOutletEnergyCarrierSummation) Name() string {
	return `Current Outlet Energy Carrier Summation`
}
func (CurrentOutletEnergyCarrierSummation) Description() string { return `` }

type CurrentOutletEnergyCarrierSummation zcl.Zu48

func (v *CurrentOutletEnergyCarrierSummation) TypeID() zcl.TypeID { return new(zcl.Zu48).TypeID() }
func (v *CurrentOutletEnergyCarrierSummation) Value() zcl.Val     { return v }

func (v CurrentOutletEnergyCarrierSummation) MarshalZcl() ([]byte, error) {
	return zcl.Zu48(v).MarshalZcl()
}

func (v *CurrentOutletEnergyCarrierSummation) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu48)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentOutletEnergyCarrierSummation(*nt)
	return br, err
}

func (v CurrentOutletEnergyCarrierSummation) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu48(v))
}

func (v *CurrentOutletEnergyCarrierSummation) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu48)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentOutletEnergyCarrierSummation(*a)
	return nil
}

func (v *CurrentOutletEnergyCarrierSummation) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu48); ok {
		*v = CurrentOutletEnergyCarrierSummation(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentOutletEnergyCarrierSummation) String() string {
	return zcl.Sprintf("%v", zcl.Zu48(v))
}

const CurrentSummationDeliveredAttr zcl.AttrID = 0

func (CurrentSummationDelivered) ID() zcl.AttrID   { return CurrentSummationDeliveredAttr }
func (CurrentSummationDelivered) Readable() bool   { return false }
func (CurrentSummationDelivered) Writable() bool   { return false }
func (CurrentSummationDelivered) Reportable() bool { return true }
func (CurrentSummationDelivered) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentSummationDelivered) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentSummationDelivered) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentSummationDelivered) AttrValue() zcl.Val  { return v.Value() }

func (CurrentSummationDelivered) Name() string { return `Current Summation Delivered` }
func (CurrentSummationDelivered) Description() string {
	return `represents the most recent summed value of Energy, Gas, or Water delivered and consumed in the premises.`
}

// CurrentSummationDelivered represents the most recent summed value of Energy, Gas, or Water delivered and consumed in the premises.
type CurrentSummationDelivered zcl.Zu48

func (v *CurrentSummationDelivered) TypeID() zcl.TypeID { return new(zcl.Zu48).TypeID() }
func (v *CurrentSummationDelivered) Value() zcl.Val     { return v }

func (v CurrentSummationDelivered) MarshalZcl() ([]byte, error) { return zcl.Zu48(v).MarshalZcl() }

func (v *CurrentSummationDelivered) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu48)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentSummationDelivered(*nt)
	return br, err
}

func (v CurrentSummationDelivered) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu48(v))
}

func (v *CurrentSummationDelivered) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu48)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentSummationDelivered(*a)
	return nil
}

func (v *CurrentSummationDelivered) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu48); ok {
		*v = CurrentSummationDelivered(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentSummationDelivered) String() string {
	return zcl.Sprintf("%v", zcl.Zu48(v))
}

const CurrentSummationReceivedAttr zcl.AttrID = 1

func (CurrentSummationReceived) ID() zcl.AttrID   { return CurrentSummationReceivedAttr }
func (CurrentSummationReceived) Readable() bool   { return false }
func (CurrentSummationReceived) Writable() bool   { return false }
func (CurrentSummationReceived) Reportable() bool { return false }
func (CurrentSummationReceived) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentSummationReceived) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentSummationReceived) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentSummationReceived) AttrValue() zcl.Val  { return v.Value() }

func (CurrentSummationReceived) Name() string { return `Current Summation Received` }
func (CurrentSummationReceived) Description() string {
	return `represents the most recent summed value of Energy, Gas, or Water generated and delivered from the premises.`
}

// CurrentSummationReceived represents the most recent summed value of Energy, Gas, or Water generated and delivered from the premises.
type CurrentSummationReceived zcl.Zu48

func (v *CurrentSummationReceived) TypeID() zcl.TypeID { return new(zcl.Zu48).TypeID() }
func (v *CurrentSummationReceived) Value() zcl.Val     { return v }

func (v CurrentSummationReceived) MarshalZcl() ([]byte, error) { return zcl.Zu48(v).MarshalZcl() }

func (v *CurrentSummationReceived) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu48)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentSummationReceived(*nt)
	return br, err
}

func (v CurrentSummationReceived) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu48(v))
}

func (v *CurrentSummationReceived) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu48)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentSummationReceived(*a)
	return nil
}

func (v *CurrentSummationReceived) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu48); ok {
		*v = CurrentSummationReceived(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentSummationReceived) String() string {
	return zcl.Sprintf("%v", zcl.Zu48(v))
}

const DftSummationAttr zcl.AttrID = 4

func (DftSummation) ID() zcl.AttrID   { return DftSummationAttr }
func (DftSummation) Readable() bool   { return false }
func (DftSummation) Writable() bool   { return false }
func (DftSummation) Reportable() bool { return false }
func (DftSummation) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DftSummation) AttrID() zcl.AttrID   { return v.ID() }
func (v DftSummation) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DftSummation) AttrValue() zcl.Val  { return v.Value() }

func (DftSummation) Name() string { return `DFT Summation` }
func (DftSummation) Description() string {
	return `represents a snapshot of attribute CurrentSummationDelivered captured at the time indicated by attribute
DailyFreezeTime`
}

// DftSummation represents a snapshot of attribute CurrentSummationDelivered captured at the time indicated by attribute
// DailyFreezeTime
type DftSummation zcl.Zu48

func (v *DftSummation) TypeID() zcl.TypeID { return new(zcl.Zu48).TypeID() }
func (v *DftSummation) Value() zcl.Val     { return v }

func (v DftSummation) MarshalZcl() ([]byte, error) { return zcl.Zu48(v).MarshalZcl() }

func (v *DftSummation) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu48)
	br, err := nt.UnmarshalZcl(b)
	*v = DftSummation(*nt)
	return br, err
}

func (v DftSummation) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu48(v))
}

func (v *DftSummation) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu48)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DftSummation(*a)
	return nil
}

func (v *DftSummation) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu48); ok {
		*v = DftSummation(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DftSummation) String() string {
	return zcl.Sprintf("%v", zcl.Zu48(v))
}

const DailyConsumptionTargetAttr zcl.AttrID = 13

func (DailyConsumptionTarget) ID() zcl.AttrID   { return DailyConsumptionTargetAttr }
func (DailyConsumptionTarget) Readable() bool   { return false }
func (DailyConsumptionTarget) Writable() bool   { return false }
func (DailyConsumptionTarget) Reportable() bool { return false }
func (DailyConsumptionTarget) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DailyConsumptionTarget) AttrID() zcl.AttrID   { return v.ID() }
func (v DailyConsumptionTarget) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DailyConsumptionTarget) AttrValue() zcl.Val  { return v.Value() }

func (DailyConsumptionTarget) Name() string { return `Daily Consumption Target` }
func (DailyConsumptionTarget) Description() string {
	return `is a daily target consumption amount that can be displayed to the consumer on a HAN device, with the intent that
it can be used to compare to actual daily consumption (e.g. compare to the CurrentDayConsumptionDelivered).
This may be sent from the utility to the ESI, or it may be derived. Although intended to be based on Block
Thresholds, it can be used for other targets not related to blocks. The formatting will be based on the
HistoricalConsumptionFormatting attribute.
Example: If based on a Block Threshold, the DailyConsumptionTarget could be calculated based on the
number of days specified in the Block Tariff Period and a given Block Threshold as follows:
DailyConsumptionTarget = BlockNThreshold / ((BlockPeriodDuration /60) / 24). Example: If the target is
based on a Block1Threshold of 675kWh and where 43200 BlockThresholdPeriod is the number of minutes in
the billing period (30 days), the ConsumptionDailyTarget would be 675 / ((43200 / 60) / 24) = 22.5 kWh per day`
}

// DailyConsumptionTarget is a daily target consumption amount that can be displayed to the consumer on a HAN device, with the intent that
// it can be used to compare to actual daily consumption (e.g. compare to the CurrentDayConsumptionDelivered).
// This may be sent from the utility to the ESI, or it may be derived. Although intended to be based on Block
// Thresholds, it can be used for other targets not related to blocks. The formatting will be based on the
// HistoricalConsumptionFormatting attribute.
// Example: If based on a Block Threshold, the DailyConsumptionTarget could be calculated based on the
// number of days specified in the Block Tariff Period and a given Block Threshold as follows:
// DailyConsumptionTarget = BlockNThreshold / ((BlockPeriodDuration /60) / 24). Example: If the target is
// based on a Block1Threshold of 675kWh and where 43200 BlockThresholdPeriod is the number of minutes in
// the billing period (30 days), the ConsumptionDailyTarget would be 675 / ((43200 / 60) / 24) = 22.5 kWh per day
type DailyConsumptionTarget zcl.Zu24

func (v *DailyConsumptionTarget) TypeID() zcl.TypeID { return new(zcl.Zu24).TypeID() }
func (v *DailyConsumptionTarget) Value() zcl.Val     { return v }

func (v DailyConsumptionTarget) MarshalZcl() ([]byte, error) { return zcl.Zu24(v).MarshalZcl() }

func (v *DailyConsumptionTarget) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*v = DailyConsumptionTarget(*nt)
	return br, err
}

func (v DailyConsumptionTarget) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu24(v))
}

func (v *DailyConsumptionTarget) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu24)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DailyConsumptionTarget(*a)
	return nil
}

func (v *DailyConsumptionTarget) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu24); ok {
		*v = DailyConsumptionTarget(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DailyConsumptionTarget) String() string {
	return zcl.Sprintf("%v", zcl.Zu24(v))
}

const DailyFreezeTimeAttr zcl.AttrID = 5

func (DailyFreezeTime) ID() zcl.AttrID   { return DailyFreezeTimeAttr }
func (DailyFreezeTime) Readable() bool   { return false }
func (DailyFreezeTime) Writable() bool   { return false }
func (DailyFreezeTime) Reportable() bool { return false }
func (DailyFreezeTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DailyFreezeTime) AttrID() zcl.AttrID   { return v.ID() }
func (v DailyFreezeTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DailyFreezeTime) AttrValue() zcl.Val  { return v.Value() }

func (DailyFreezeTime) Name() string { return `Daily Freeze Time` }
func (DailyFreezeTime) Description() string {
	return `represents the time of day when DFTSummation.
Bits 0 to 7: Range of 0 to 0x3C representing the number of minutes past the top of the hour.
Bits 8 to 15: Range of 0 to 0x17 representing the hour of the day (in 24-hour format).`
}

// DailyFreezeTime represents the time of day when DFTSummation.
// Bits 0 to 7: Range of 0 to 0x3C representing the number of minutes past the top of the hour.
// Bits 8 to 15: Range of 0 to 0x17 representing the hour of the day (in 24-hour format).
type DailyFreezeTime zcl.Zu16

func (v *DailyFreezeTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *DailyFreezeTime) Value() zcl.Val     { return v }

func (v DailyFreezeTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *DailyFreezeTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = DailyFreezeTime(*nt)
	return br, err
}

func (v DailyFreezeTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *DailyFreezeTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DailyFreezeTime(*a)
	return nil
}

func (v *DailyFreezeTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = DailyFreezeTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DailyFreezeTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const DefaultUpdatePeriodAttr zcl.AttrID = 10

func (DefaultUpdatePeriod) ID() zcl.AttrID   { return DefaultUpdatePeriodAttr }
func (DefaultUpdatePeriod) Readable() bool   { return false }
func (DefaultUpdatePeriod) Writable() bool   { return false }
func (DefaultUpdatePeriod) Reportable() bool { return false }
func (DefaultUpdatePeriod) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DefaultUpdatePeriod) AttrID() zcl.AttrID   { return v.ID() }
func (v DefaultUpdatePeriod) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DefaultUpdatePeriod) AttrValue() zcl.Val  { return v.Value() }

func (DefaultUpdatePeriod) Name() string { return `Default Update Period` }
func (DefaultUpdatePeriod) Description() string {
	return `represents the interval (seconds) at which the InstantaneousDemand
attribute is updated when not in fast poll mode. InstantaneousDemand may be continuously updated as new
measurements are acquired, but at a minimum InstantaneousDemand must be updated at the
DefaultUpdatePeriod. The DefaultUpdatePeriod may apply to other attributes as defined by the device
manufacturer.`
}

// DefaultUpdatePeriod represents the interval (seconds) at which the InstantaneousDemand
// attribute is updated when not in fast poll mode. InstantaneousDemand may be continuously updated as new
// measurements are acquired, but at a minimum InstantaneousDemand must be updated at the
// DefaultUpdatePeriod. The DefaultUpdatePeriod may apply to other attributes as defined by the device
// manufacturer.
type DefaultUpdatePeriod zcl.Zu8

func (v *DefaultUpdatePeriod) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *DefaultUpdatePeriod) Value() zcl.Val     { return v }

func (v DefaultUpdatePeriod) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *DefaultUpdatePeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = DefaultUpdatePeriod(*nt)
	return br, err
}

func (v DefaultUpdatePeriod) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *DefaultUpdatePeriod) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DefaultUpdatePeriod(*a)
	return nil
}

func (v *DefaultUpdatePeriod) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = DefaultUpdatePeriod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DefaultUpdatePeriod) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const FastPollUpdatePeriodAttr zcl.AttrID = 11

func (FastPollUpdatePeriod) ID() zcl.AttrID   { return FastPollUpdatePeriodAttr }
func (FastPollUpdatePeriod) Readable() bool   { return false }
func (FastPollUpdatePeriod) Writable() bool   { return false }
func (FastPollUpdatePeriod) Reportable() bool { return false }
func (FastPollUpdatePeriod) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v FastPollUpdatePeriod) AttrID() zcl.AttrID   { return v.ID() }
func (v FastPollUpdatePeriod) AttrType() zcl.TypeID { return v.TypeID() }
func (v *FastPollUpdatePeriod) AttrValue() zcl.Val  { return v.Value() }

func (FastPollUpdatePeriod) Name() string { return `Fast Poll Update Period` }
func (FastPollUpdatePeriod) Description() string {
	return `represents the interval (seconds) at which the InstantaneousDemand
attribute is updated when in fast poll mode. InstantaneousDemand may be continuously updated as new
measurements are acquired, but at a minimum, InstantaneousDemand must be updated at the
FastPollUpdatePeriod. The FastPollUpdatePeriod may apply to other attributes as defined by the device
manufacturer.`
}

// FastPollUpdatePeriod represents the interval (seconds) at which the InstantaneousDemand
// attribute is updated when in fast poll mode. InstantaneousDemand may be continuously updated as new
// measurements are acquired, but at a minimum, InstantaneousDemand must be updated at the
// FastPollUpdatePeriod. The FastPollUpdatePeriod may apply to other attributes as defined by the device
// manufacturer.
type FastPollUpdatePeriod zcl.Zu8

func (v *FastPollUpdatePeriod) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *FastPollUpdatePeriod) Value() zcl.Val     { return v }

func (v FastPollUpdatePeriod) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *FastPollUpdatePeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = FastPollUpdatePeriod(*nt)
	return br, err
}

func (v FastPollUpdatePeriod) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *FastPollUpdatePeriod) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = FastPollUpdatePeriod(*a)
	return nil
}

func (v *FastPollUpdatePeriod) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = FastPollUpdatePeriod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v FastPollUpdatePeriod) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const FlowRestrictionAttr zcl.AttrID = 19

func (FlowRestriction) ID() zcl.AttrID   { return FlowRestrictionAttr }
func (FlowRestriction) Readable() bool   { return false }
func (FlowRestriction) Writable() bool   { return false }
func (FlowRestriction) Reportable() bool { return false }
func (FlowRestriction) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v FlowRestriction) AttrID() zcl.AttrID   { return v.ID() }
func (v FlowRestriction) AttrType() zcl.TypeID { return v.TypeID() }
func (v *FlowRestriction) AttrValue() zcl.Val  { return v.Value() }

func (FlowRestriction) Name() string        { return `Flow Restriction` }
func (FlowRestriction) Description() string { return `` }

type FlowRestriction zcl.Zu8

func (v *FlowRestriction) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *FlowRestriction) Value() zcl.Val     { return v }

func (v FlowRestriction) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *FlowRestriction) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = FlowRestriction(*nt)
	return br, err
}

func (v FlowRestriction) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *FlowRestriction) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = FlowRestriction(*a)
	return nil
}

func (v *FlowRestriction) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = FlowRestriction(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v FlowRestriction) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const InletTemperatureAttr zcl.AttrID = 23

func (InletTemperature) ID() zcl.AttrID   { return InletTemperatureAttr }
func (InletTemperature) Readable() bool   { return false }
func (InletTemperature) Writable() bool   { return false }
func (InletTemperature) Reportable() bool { return false }
func (InletTemperature) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v InletTemperature) AttrID() zcl.AttrID   { return v.ID() }
func (v InletTemperature) AttrType() zcl.TypeID { return v.TypeID() }
func (v *InletTemperature) AttrValue() zcl.Val  { return v.Value() }

func (InletTemperature) Name() string        { return `Inlet Temperature` }
func (InletTemperature) Description() string { return `` }

type InletTemperature zcl.Zs24

func (v *InletTemperature) TypeID() zcl.TypeID { return new(zcl.Zs24).TypeID() }
func (v *InletTemperature) Value() zcl.Val     { return v }

func (v InletTemperature) MarshalZcl() ([]byte, error) { return zcl.Zs24(v).MarshalZcl() }

func (v *InletTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs24)
	br, err := nt.UnmarshalZcl(b)
	*v = InletTemperature(*nt)
	return br, err
}

func (v InletTemperature) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs24(v))
}

func (v *InletTemperature) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs24)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = InletTemperature(*a)
	return nil
}

func (v *InletTemperature) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs24); ok {
		*v = InletTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v InletTemperature) String() string {
	return zcl.Sprintf("%v", zcl.Zs24(v))
}

const IntervalReadReportingPeriodAttr zcl.AttrID = 16

func (IntervalReadReportingPeriod) ID() zcl.AttrID   { return IntervalReadReportingPeriodAttr }
func (IntervalReadReportingPeriod) Readable() bool   { return false }
func (IntervalReadReportingPeriod) Writable() bool   { return false }
func (IntervalReadReportingPeriod) Reportable() bool { return false }
func (IntervalReadReportingPeriod) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v IntervalReadReportingPeriod) AttrID() zcl.AttrID   { return v.ID() }
func (v IntervalReadReportingPeriod) AttrType() zcl.TypeID { return v.TypeID() }
func (v *IntervalReadReportingPeriod) AttrValue() zcl.Val  { return v.Value() }

func (IntervalReadReportingPeriod) Name() string        { return `Interval Read Reporting Period` }
func (IntervalReadReportingPeriod) Description() string { return `` }

type IntervalReadReportingPeriod zcl.Zu16

func (v *IntervalReadReportingPeriod) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *IntervalReadReportingPeriod) Value() zcl.Val     { return v }

func (v IntervalReadReportingPeriod) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *IntervalReadReportingPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = IntervalReadReportingPeriod(*nt)
	return br, err
}

func (v IntervalReadReportingPeriod) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *IntervalReadReportingPeriod) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IntervalReadReportingPeriod(*a)
	return nil
}

func (v *IntervalReadReportingPeriod) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = IntervalReadReportingPeriod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IntervalReadReportingPeriod) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const OutletTemperatureAttr zcl.AttrID = 24

func (OutletTemperature) ID() zcl.AttrID   { return OutletTemperatureAttr }
func (OutletTemperature) Readable() bool   { return false }
func (OutletTemperature) Writable() bool   { return false }
func (OutletTemperature) Reportable() bool { return false }
func (OutletTemperature) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v OutletTemperature) AttrID() zcl.AttrID   { return v.ID() }
func (v OutletTemperature) AttrType() zcl.TypeID { return v.TypeID() }
func (v *OutletTemperature) AttrValue() zcl.Val  { return v.Value() }

func (OutletTemperature) Name() string        { return `Outlet Temperature` }
func (OutletTemperature) Description() string { return `` }

type OutletTemperature zcl.Zs24

func (v *OutletTemperature) TypeID() zcl.TypeID { return new(zcl.Zs24).TypeID() }
func (v *OutletTemperature) Value() zcl.Val     { return v }

func (v OutletTemperature) MarshalZcl() ([]byte, error) { return zcl.Zs24(v).MarshalZcl() }

func (v *OutletTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs24)
	br, err := nt.UnmarshalZcl(b)
	*v = OutletTemperature(*nt)
	return br, err
}

func (v OutletTemperature) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs24(v))
}

func (v *OutletTemperature) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs24)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OutletTemperature(*a)
	return nil
}

func (v *OutletTemperature) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs24); ok {
		*v = OutletTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OutletTemperature) String() string {
	return zcl.Sprintf("%v", zcl.Zs24(v))
}

const PowerFactorAttr zcl.AttrID = 6

func (PowerFactor) ID() zcl.AttrID   { return PowerFactorAttr }
func (PowerFactor) Readable() bool   { return false }
func (PowerFactor) Writable() bool   { return false }
func (PowerFactor) Reportable() bool { return false }
func (PowerFactor) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PowerFactor) AttrID() zcl.AttrID   { return v.ID() }
func (v PowerFactor) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PowerFactor) AttrValue() zcl.Val  { return v.Value() }

func (PowerFactor) Name() string { return `Power Factor` }
func (PowerFactor) Description() string {
	return `contains the Average Power Factor ratio in 1/100ths. Valid values are 0 to 99`
}

// PowerFactor contains the Average Power Factor ratio in 1/100ths. Valid values are 0 to 99
type PowerFactor zcl.Zs8

func (v *PowerFactor) TypeID() zcl.TypeID { return new(zcl.Zs8).TypeID() }
func (v *PowerFactor) Value() zcl.Val     { return v }

func (v PowerFactor) MarshalZcl() ([]byte, error) { return zcl.Zs8(v).MarshalZcl() }

func (v *PowerFactor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*v = PowerFactor(*nt)
	return br, err
}

func (v PowerFactor) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs8(v))
}

func (v *PowerFactor) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PowerFactor(*a)
	return nil
}

func (v *PowerFactor) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs8); ok {
		*v = PowerFactor(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PowerFactor) String() string {
	return zcl.Sprintf("%v", zcl.Zs8(v))
}

const PresetReadingTimeAttr zcl.AttrID = 17

func (PresetReadingTime) ID() zcl.AttrID   { return PresetReadingTimeAttr }
func (PresetReadingTime) Readable() bool   { return false }
func (PresetReadingTime) Writable() bool   { return false }
func (PresetReadingTime) Reportable() bool { return false }
func (PresetReadingTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PresetReadingTime) AttrID() zcl.AttrID   { return v.ID() }
func (v PresetReadingTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PresetReadingTime) AttrValue() zcl.Val  { return v.Value() }

func (PresetReadingTime) Name() string        { return `Preset Reading Time` }
func (PresetReadingTime) Description() string { return `` }

type PresetReadingTime zcl.Zu16

func (v *PresetReadingTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *PresetReadingTime) Value() zcl.Val     { return v }

func (v PresetReadingTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *PresetReadingTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = PresetReadingTime(*nt)
	return br, err
}

func (v PresetReadingTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *PresetReadingTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PresetReadingTime(*a)
	return nil
}

func (v *PresetReadingTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = PresetReadingTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PresetReadingTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const PreviousBlockPeriodConsumptionDeliveredAttr zcl.AttrID = 28

func (PreviousBlockPeriodConsumptionDelivered) ID() zcl.AttrID {
	return PreviousBlockPeriodConsumptionDeliveredAttr
}
func (PreviousBlockPeriodConsumptionDelivered) Readable() bool   { return false }
func (PreviousBlockPeriodConsumptionDelivered) Writable() bool   { return false }
func (PreviousBlockPeriodConsumptionDelivered) Reportable() bool { return false }
func (PreviousBlockPeriodConsumptionDelivered) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PreviousBlockPeriodConsumptionDelivered) AttrID() zcl.AttrID   { return v.ID() }
func (v PreviousBlockPeriodConsumptionDelivered) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PreviousBlockPeriodConsumptionDelivered) AttrValue() zcl.Val  { return v.Value() }

func (PreviousBlockPeriodConsumptionDelivered) Name() string {
	return `Previous Block Period Consumption Delivered`
}
func (PreviousBlockPeriodConsumptionDelivered) Description() string { return `` }

type PreviousBlockPeriodConsumptionDelivered zcl.Zu48

func (v *PreviousBlockPeriodConsumptionDelivered) TypeID() zcl.TypeID { return new(zcl.Zu48).TypeID() }
func (v *PreviousBlockPeriodConsumptionDelivered) Value() zcl.Val     { return v }

func (v PreviousBlockPeriodConsumptionDelivered) MarshalZcl() ([]byte, error) {
	return zcl.Zu48(v).MarshalZcl()
}

func (v *PreviousBlockPeriodConsumptionDelivered) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu48)
	br, err := nt.UnmarshalZcl(b)
	*v = PreviousBlockPeriodConsumptionDelivered(*nt)
	return br, err
}

func (v PreviousBlockPeriodConsumptionDelivered) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu48(v))
}

func (v *PreviousBlockPeriodConsumptionDelivered) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu48)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PreviousBlockPeriodConsumptionDelivered(*a)
	return nil
}

func (v *PreviousBlockPeriodConsumptionDelivered) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu48); ok {
		*v = PreviousBlockPeriodConsumptionDelivered(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PreviousBlockPeriodConsumptionDelivered) String() string {
	return zcl.Sprintf("%v", zcl.Zu48(v))
}

const ProfileIntervalPeriodAttr zcl.AttrID = 15

func (ProfileIntervalPeriod) ID() zcl.AttrID   { return ProfileIntervalPeriodAttr }
func (ProfileIntervalPeriod) Readable() bool   { return false }
func (ProfileIntervalPeriod) Writable() bool   { return false }
func (ProfileIntervalPeriod) Reportable() bool { return false }
func (ProfileIntervalPeriod) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ProfileIntervalPeriod) AttrID() zcl.AttrID   { return v.ID() }
func (v ProfileIntervalPeriod) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ProfileIntervalPeriod) AttrValue() zcl.Val  { return v.Value() }

func (ProfileIntervalPeriod) Name() string { return `Profile Interval Period` }
func (ProfileIntervalPeriod) Description() string {
	return `is currently included in the Get Profile Response command payload, but does not appear in an attribute set.
This represents the duration of each interval. ProfileIntervalPeriod represents the interval or time frame used
to capture metered Energy, Gas, and Water consumption for profiling purposes.`
}

// ProfileIntervalPeriod is currently included in the Get Profile Response command payload, but does not appear in an attribute set.
// This represents the duration of each interval. ProfileIntervalPeriod represents the interval or time frame used
// to capture metered Energy, Gas, and Water consumption for profiling purposes.
type ProfileIntervalPeriod zcl.Zenum8

func (v *ProfileIntervalPeriod) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *ProfileIntervalPeriod) Value() zcl.Val     { return v }

func (v ProfileIntervalPeriod) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *ProfileIntervalPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = ProfileIntervalPeriod(*nt)
	return br, err
}

func (v ProfileIntervalPeriod) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *ProfileIntervalPeriod) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ProfileIntervalPeriod(*a)
	return nil
}

func (v *ProfileIntervalPeriod) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = ProfileIntervalPeriod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ProfileIntervalPeriod) String() string {
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

const ReadingSnapShotTimeAttr zcl.AttrID = 7

func (ReadingSnapShotTime) ID() zcl.AttrID   { return ReadingSnapShotTimeAttr }
func (ReadingSnapShotTime) Readable() bool   { return false }
func (ReadingSnapShotTime) Writable() bool   { return false }
func (ReadingSnapShotTime) Reportable() bool { return false }
func (ReadingSnapShotTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ReadingSnapShotTime) AttrID() zcl.AttrID   { return v.ID() }
func (v ReadingSnapShotTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ReadingSnapShotTime) AttrValue() zcl.Val  { return v.Value() }

func (ReadingSnapShotTime) Name() string { return `Reading Snap Shot Time` }
func (ReadingSnapShotTime) Description() string {
	return `represents the last time all of the CurrentSummationDelivered, CurrentSummationReceived,
CurrentMaxDemandDelivered, and CurrentMaxDemandReceived attributes that are supported by the device were updated`
}

// ReadingSnapShotTime represents the last time all of the CurrentSummationDelivered, CurrentSummationReceived,
// CurrentMaxDemandDelivered, and CurrentMaxDemandReceived attributes that are supported by the device were updated
type ReadingSnapShotTime zcl.Zutc

func (v *ReadingSnapShotTime) TypeID() zcl.TypeID { return new(zcl.Zutc).TypeID() }
func (v *ReadingSnapShotTime) Value() zcl.Val     { return v }

func (v ReadingSnapShotTime) MarshalZcl() ([]byte, error) { return zcl.Zutc(v).MarshalZcl() }

func (v *ReadingSnapShotTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*v = ReadingSnapShotTime(*nt)
	return br, err
}

func (v ReadingSnapShotTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zutc(v))
}

func (v *ReadingSnapShotTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zutc)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ReadingSnapShotTime(*a)
	return nil
}

func (v *ReadingSnapShotTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zutc); ok {
		*v = ReadingSnapShotTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ReadingSnapShotTime) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(v))
}

const SupplyStatusAttr zcl.AttrID = 20

func (SupplyStatus) ID() zcl.AttrID   { return SupplyStatusAttr }
func (SupplyStatus) Readable() bool   { return false }
func (SupplyStatus) Writable() bool   { return false }
func (SupplyStatus) Reportable() bool { return false }
func (SupplyStatus) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v SupplyStatus) AttrID() zcl.AttrID   { return v.ID() }
func (v SupplyStatus) AttrType() zcl.TypeID { return v.TypeID() }
func (v *SupplyStatus) AttrValue() zcl.Val  { return v.Value() }

func (SupplyStatus) Name() string        { return `Supply Status` }
func (SupplyStatus) Description() string { return `` }

type SupplyStatus zcl.Zenum8

func (v *SupplyStatus) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *SupplyStatus) Value() zcl.Val     { return v }

func (v SupplyStatus) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *SupplyStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = SupplyStatus(*nt)
	return br, err
}

func (v SupplyStatus) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *SupplyStatus) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SupplyStatus(*a)
	return nil
}

func (v *SupplyStatus) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = SupplyStatus(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SupplyStatus) String() string {
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

const VolumePerReportAttr zcl.AttrID = 18

func (VolumePerReport) ID() zcl.AttrID   { return VolumePerReportAttr }
func (VolumePerReport) Readable() bool   { return false }
func (VolumePerReport) Writable() bool   { return false }
func (VolumePerReport) Reportable() bool { return false }
func (VolumePerReport) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v VolumePerReport) AttrID() zcl.AttrID   { return v.ID() }
func (v VolumePerReport) AttrType() zcl.TypeID { return v.TypeID() }
func (v *VolumePerReport) AttrValue() zcl.Val  { return v.Value() }

func (VolumePerReport) Name() string        { return `Volume Per Report` }
func (VolumePerReport) Description() string { return `` }

type VolumePerReport zcl.Zu16

func (v *VolumePerReport) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *VolumePerReport) Value() zcl.Val     { return v }

func (v VolumePerReport) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *VolumePerReport) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = VolumePerReport(*nt)
	return br, err
}

func (v VolumePerReport) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *VolumePerReport) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = VolumePerReport(*a)
	return nil
}

func (v *VolumePerReport) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = VolumePerReport(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v VolumePerReport) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}
