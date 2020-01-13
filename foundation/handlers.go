package foundation

import "hemtjan.st/zcl"

type Frame = zcl.ReceivedZclFrame

type ReadAttributesHandler interface {
	HandleReadAttributes(Frame, *ReadAttributes) (*ReadAttributesResponse, error)
}
type ConfigureReportingHandler interface {
	HandleConfigureReporting(Frame, *ConfigureReporting) (*ConfigureReportingResponse, error)
}
type ReadReportingConfigurationHandler interface {
	HandleReadReportingConfiguration(Frame, *ReadReportingConfiguration) (*ReadReportingConfigurationResponse, error)
}
type ReportAttributesHandler interface {
	HandleReportAttributes(Frame, *ReportAttributes) error
}
type DiscoverAttributesHandler interface {
	HandleDiscoverAttributes(Frame, *DiscoverAttributes) (*DiscoverAttributesResponse, error)
}
type DiscoverCommandsReceivedHandler interface {
	HandleDiscoverCommandsReceived(Frame, *DiscoverCommandsReceived) (*DiscoverCommandsReceivedResponse, error)
}
type DiscoverCommandsGeneratedHandler interface {
	HandleDiscoverCommandsGenerated(Frame, *DiscoverCommandsGenerated) (*DiscoverCommandsGeneratedResponse, error)
}
type ReadAttributesResponseHandler interface {
	HandleReadAttributesResponse(Frame, *ReadAttributesResponse) error
}
type ConfigureReportingResponseHandler interface {
	HandleConfigureReportingResponse(Frame, *ConfigureReportingResponse) error
}
type ReadReportingConfigurationResponseHandler interface {
	HandleReadReportingConfigurationResponse(Frame, *ReadReportingConfigurationResponse) error
}
type DiscoverAttributesResponseHandler interface {
	HandleDiscoverAttributesResponse(Frame, *DiscoverAttributesResponse) error
}
type DiscoverCommandsReceivedResponseHandler interface {
	HandleDiscoverCommandsReceivedResponse(Frame, *DiscoverCommandsReceivedResponse) error
}
type DiscoverCommandsGeneratedResponseHandler interface {
	HandleDiscoverCommandsGeneratedResponse(Frame, *DiscoverCommandsGeneratedResponse) error
}

func (v *ReadAttributes) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h ReadAttributesHandler
	if h, found = handler.(ReadAttributesHandler); found {
		rsp, err = h.HandleReadAttributes(frame, v)
	}
	return
}
func (v *ConfigureReporting) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h ConfigureReportingHandler
	if h, found = handler.(ConfigureReportingHandler); found {
		rsp, err = h.HandleConfigureReporting(frame, v)
	}
	return
}
func (v *ReadReportingConfiguration) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h ReadReportingConfigurationHandler
	if h, found = handler.(ReadReportingConfigurationHandler); found {
		rsp, err = h.HandleReadReportingConfiguration(frame, v)
	}
	return
}
func (v *ReportAttributes) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h ReportAttributesHandler
	if h, found = handler.(ReportAttributesHandler); found {
		err = h.HandleReportAttributes(frame, v)
	}
	return
}
func (v *DiscoverAttributes) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h DiscoverAttributesHandler
	if h, found = handler.(DiscoverAttributesHandler); found {
		rsp, err = h.HandleDiscoverAttributes(frame, v)
	}
	return
}
func (v *DiscoverCommandsReceived) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h DiscoverCommandsReceivedHandler
	if h, found = handler.(DiscoverCommandsReceivedHandler); found {
		rsp, err = h.HandleDiscoverCommandsReceived(frame, v)
	}
	return
}
func (v *DiscoverCommandsGenerated) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h DiscoverCommandsGeneratedHandler
	if h, found = handler.(DiscoverCommandsGeneratedHandler); found {
		rsp, err = h.HandleDiscoverCommandsGenerated(frame, v)
	}
	return
}
func (v *ReadAttributesResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h ReadAttributesResponseHandler
	if h, found = handler.(ReadAttributesResponseHandler); found {
		err = h.HandleReadAttributesResponse(frame, v)
	}
	return
}
func (v *ConfigureReportingResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h ConfigureReportingResponseHandler
	if h, found = handler.(ConfigureReportingResponseHandler); found {
		err = h.HandleConfigureReportingResponse(frame, v)
	}
	return
}
func (v *ReadReportingConfigurationResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h ReadReportingConfigurationResponseHandler
	if h, found = handler.(ReadReportingConfigurationResponseHandler); found {
		err = h.HandleReadReportingConfigurationResponse(frame, v)
	}
	return
}
func (v *DiscoverAttributesResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h DiscoverAttributesResponseHandler
	if h, found = handler.(DiscoverAttributesResponseHandler); found {
		err = h.HandleDiscoverAttributesResponse(frame, v)
	}
	return
}
func (v *DiscoverCommandsReceivedResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h DiscoverCommandsReceivedResponseHandler
	if h, found = handler.(DiscoverCommandsReceivedResponseHandler); found {
		err = h.HandleDiscoverCommandsReceivedResponse(frame, v)
	}
	return
}
func (v *DiscoverCommandsGeneratedResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h DiscoverCommandsGeneratedResponseHandler
	if h, found = handler.(DiscoverCommandsGeneratedResponseHandler); found {
		err = h.HandleDiscoverCommandsGeneratedResponse(frame, v)
	}
	return
}
