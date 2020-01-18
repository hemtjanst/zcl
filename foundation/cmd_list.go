package foundation

import "hemtjan.st/zcl"

// DefaultResponse lives in zcl package to avoid import cycles
type DefaultResponse = zcl.DefaultResponse

const (
	ReadAttributesCommand                     = zcl.CommandID(0x00)
	ReadAttributesResponseCommand             = zcl.CommandID(0x01)
	WriteAttributesCommand                    = zcl.CommandID(0x02)
	WriteAttributesUndividedCommand           = zcl.CommandID(0x03)
	WriteAttributesResponseCommand            = zcl.CommandID(0x04)
	WriteAttributesNoResponseCommand          = zcl.CommandID(0x05)
	ConfigureReportingCommand                 = zcl.CommandID(0x06)
	ConfigureReportingResponseCommand         = zcl.CommandID(0x07)
	ReadReportingConfigurationCommand         = zcl.CommandID(0x08)
	ReadReportingConfigurationResponseCommand = zcl.CommandID(0x09)
	ReportAttributesCommand                   = zcl.CommandID(0x0a)
	DefaultResponseCommand                    = zcl.DefaultResponseCommand
	DiscoverAttributesCommand                 = zcl.CommandID(0x0c)
	DiscoverAttributesResponseCommand         = zcl.CommandID(0x0d)
	ReadAttributesStructuredCommand           = zcl.CommandID(0x0e)
	WriteAttributesStructuredCommand          = zcl.CommandID(0x0f)
	WriteAttributesStructuredResponseCommand  = zcl.CommandID(0x10)
	DiscoverCommandsReceivedCommand           = zcl.CommandID(0x11)
	DiscoverCommandsReceivedResponseCommand   = zcl.CommandID(0x12)
	DiscoverCommandsGeneratedCommand          = zcl.CommandID(0x13)
	DiscoverCommandsGeneratedResponseCommand  = zcl.CommandID(0x14)
	DiscoverAttributesExtendedCommand         = zcl.CommandID(0x15)
	DiscoverAttributesExtendedResponseCommand = zcl.CommandID(0x16)
)

var Commands = map[zcl.CommandID]func() zcl.General{
	ReadAttributesCommand:         func() zcl.General { return new(ReadAttributes) },
	ReadAttributesResponseCommand: func() zcl.General { return new(ReadAttributesResponse) },
	WriteAttributesCommand:        func() zcl.General { return new(WriteAttributes) },
	/*WriteAttributesUndividedCommand:           func() zcl.General { return new(WriteAttributesUndivided) },*/
	WriteAttributesResponseCommand: func() zcl.General { return new(WriteAttributesResponse) },
	/*WriteAttributesNoResponseCommand:          func() zcl.General { return new(WriteAttributesNoResponse) },*/
	ConfigureReportingCommand:                 func() zcl.General { return new(ConfigureReporting) },
	ConfigureReportingResponseCommand:         func() zcl.General { return new(ConfigureReportingResponse) },
	ReadReportingConfigurationCommand:         func() zcl.General { return new(ReadReportingConfiguration) },
	ReadReportingConfigurationResponseCommand: func() zcl.General { return new(ReadReportingConfigurationResponse) },
	ReportAttributesCommand:                   func() zcl.General { return new(ReportAttributes) },
	DefaultResponseCommand:                    func() zcl.General { return new(DefaultResponse) },
	DiscoverAttributesCommand:                 func() zcl.General { return new(DiscoverAttributes) },
	DiscoverAttributesResponseCommand:         func() zcl.General { return new(DiscoverAttributesResponse) },
	/*ReadAttributesStructuredCommand:           func() zcl.General { return new(ReadAttributesStructured) },
	WriteAttributesStructuredCommand:          func() zcl.General { return new(WriteAttributesStructured) },
	WriteAttributesStructuredResponseCommand:  func() zcl.General { return new(WriteAttributesStructuredResponse) },*/
	DiscoverCommandsReceivedCommand:          func() zcl.General { return new(DiscoverCommandsReceived) },
	DiscoverCommandsReceivedResponseCommand:  func() zcl.General { return new(DiscoverCommandsReceivedResponse) },
	DiscoverCommandsGeneratedCommand:         func() zcl.General { return new(DiscoverCommandsGenerated) },
	DiscoverCommandsGeneratedResponseCommand: func() zcl.General { return new(DiscoverCommandsGeneratedResponse) },
	/*DiscoverAttributesExtendedCommand:         func() zcl.General { return new(DiscoverAttributesExtended) },
	DiscoverAttributesExtendedResponseCommand: func() zcl.General { return new(DiscoverAttributesExtendedResponse) },*/
}

var Requests = []zcl.CommandID{
	ReadAttributesCommand, WriteAttributesCommand, WriteAttributesUndividedCommand, ConfigureReportingCommand,
	ReadReportingConfigurationCommand, ReportAttributesCommand, DiscoverAttributesCommand,
	ReadAttributesStructuredCommand, WriteAttributesStructuredCommand, DiscoverCommandsReceivedCommand,
	DiscoverCommandsGeneratedCommand, DiscoverAttributesExtendedCommand,
}

func IsRequest(id zcl.CommandID) bool {
	for _, i := range Requests {
		if i == id {
			return true
		}
	}
	return false
}

func (ReadAttributes) Direction() zcl.Direction             { return zcl.ClientToServer }
func (ConfigureReporting) Direction() zcl.Direction         { return zcl.ClientToServer }
func (ReadReportingConfiguration) Direction() zcl.Direction { return zcl.ClientToServer }
func (ReportAttributes) Direction() zcl.Direction           { return zcl.ClientToServer }
func (DiscoverAttributes) Direction() zcl.Direction         { return zcl.ClientToServer }
func (DiscoverCommandsReceived) Direction() zcl.Direction   { return zcl.ClientToServer }
func (DiscoverCommandsGenerated) Direction() zcl.Direction  { return zcl.ClientToServer }
func (WriteAttributes) Direction() zcl.Direction            { return zcl.ClientToServer }

//func (WriteAttributesUndivided) Direction() zcl.Direction { return zcl.ClientToServer }
//func (ReadAttributesStructured) Direction() zcl.Direction { return zcl.ClientToServer }
//func (WriteAttributesStructured) Direction() zcl.Direction { return zcl.ClientToServer }
//func (DiscoverAttributesExtended) Direction() zcl.Direction { return zcl.ClientToServer }

func (ReadAttributesResponse) Direction() zcl.Direction             { return zcl.ServerToClient }
func (ConfigureReportingResponse) Direction() zcl.Direction         { return zcl.ServerToClient }
func (ReadReportingConfigurationResponse) Direction() zcl.Direction { return zcl.ServerToClient }
func (DiscoverAttributesResponse) Direction() zcl.Direction         { return zcl.ServerToClient }
func (DiscoverCommandsReceivedResponse) Direction() zcl.Direction   { return zcl.ServerToClient }
func (DiscoverCommandsGeneratedResponse) Direction() zcl.Direction  { return zcl.ServerToClient }
func (WriteAttributesResponse) Direction() zcl.Direction            { return zcl.ServerToClient }

//func (WriteAttributesNoResponse) Direction() zcl.Direction { return zcl.ServerToClient }
//func (WriteAttributesStructuredResponse) Direction() zcl.Direction { return zcl.ServerToClient }
//func (DiscoverAttributesExtendedResponse) Direction() zcl.Direction { return zcl.ServerToClient }
