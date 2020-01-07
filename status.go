package zcl

import "fmt"

type Status Zenum8

const (
	Success                  Status = 0x00
	Failure                  Status = 0x01
	NotAuthorized            Status = 0x7e
	ReservedFieldNotZero     Status = 0x7f
	MalformedCommand         Status = 0x80
	UnsupClusterCommand      Status = 0x81
	UnsupGeneralCommand      Status = 0x82
	UnsupManufClusterCommand Status = 0x83
	UnsupManufGeneralCommand Status = 0x84
	InvalidField             Status = 0x85
	UnsupportedAttribute     Status = 0x86
	InvalidValue             Status = 0x87
	ReadOnly                 Status = 0x88
	InsufficientSpace        Status = 0x89
	DuplicateExists          Status = 0x8a
	NotFound                 Status = 0x8b
	UnreportableAttribute    Status = 0x8c
	InvalidDataType          Status = 0x8d
	InvalidSector            Status = 0x8e
	WriteOnly                Status = 0x8f
	InconsistentStartupState Status = 0x90
	DefinedOutOfBand         Status = 0x91
	HardwareFailure          Status = 0xc0
	SoftwareFailure          Status = 0xc1
	CalibrationError         Status = 0xc2
)

func (e Status) String() string {
	switch e {
	case Success:
		return "Success"
	case Failure:
		return "Failure"
	case NotAuthorized:
		return "NotAuthorized"
	case ReservedFieldNotZero:
		return "ReservedFieldNotZero"
	case MalformedCommand:
		return "MalformedCommand"
	case UnsupClusterCommand:
		return "UnsupClusterCommand"
	case UnsupGeneralCommand:
		return "UnsupGeneralCommand"
	case UnsupManufClusterCommand:
		return "UnsupManufClusterCommand"
	case UnsupManufGeneralCommand:
		return "UnsupManufGeneralCommand"
	case InvalidField:
		return "InvalidField"
	case UnsupportedAttribute:
		return "UnsupportedAttribute"
	case InvalidValue:
		return "InvalidValue"
	case ReadOnly:
		return "ReadOnly"
	case InsufficientSpace:
		return "InsufficientSpace"
	case DuplicateExists:
		return "DuplicateExists"
	case NotFound:
		return "NotFound"
	case UnreportableAttribute:
		return "UnreportableAttribute"
	case InvalidDataType:
		return "InvalidDataType"
	case InvalidSector:
		return "InvalidSector"
	case WriteOnly:
		return "WriteOnly"
	case InconsistentStartupState:
		return "InconsistentStartupState"
	case DefinedOutOfBand:
		return "DefinedOutOfBand"
	case HardwareFailure:
		return "HardwareFailure"
	case SoftwareFailure:
		return "SoftwareFailure"
	case CalibrationError:
		return "CalibrationError"
	default:
		return fmt.Sprintf("Status(0x%02X)", uint(e))
	}
}

func (e *Status) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(1, buf)
	*e = Status(val)
	return buf, err
}
func (e Status) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(1, uint64(e)) }
func (e *Status) Values() []Val              { return []Val{e} }
func (e Status) TypeID() TypeID              { return 48 }

func (e Status) Valid() bool { return e != Status(0xFF) }
