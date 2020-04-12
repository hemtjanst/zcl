package otau

import "hemtjan.st/zcl"

// Otau
const OtauID zcl.ClusterID = 25

var OtauCluster = zcl.Cluster{
	Name: "OTAU",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		QueryNextImageCommand:    func() zcl.Command { return new(QueryNextImage) },
		ImageBlockRequestCommand: func() zcl.Command { return new(ImageBlockRequest) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{
		QueryNextImageResponseCommand: func() zcl.Command { return new(QueryNextImageResponse) },
		ImageBlockResponseCommand:     func() zcl.Command { return new(ImageBlockResponse) },
	},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{
		UpgradeServerAttr:          func() zcl.Attr { return new(UpgradeServer) },
		FileOffsetAttr:             func() zcl.Attr { return new(FileOffset) },
		CurrentFileVersionAttr:     func() zcl.Attr { return new(CurrentFileVersion) },
		CurrentStackVersionAttr:    func() zcl.Attr { return new(CurrentStackVersion) },
		DownloadedFileVersionAttr:  func() zcl.Attr { return new(DownloadedFileVersion) },
		DownloadedStackVersionAttr: func() zcl.Attr { return new(DownloadedStackVersion) },
		ImageUpgradeStatusAttr:     func() zcl.Attr { return new(ImageUpgradeStatus) },
		MinBlockRequestDelayAttr:   func() zcl.Attr { return new(MinBlockRequestDelay) },
	},
	SceneAttr: []zcl.AttrID{},
}

type QueryNextImage struct {
	HardwareVersionPresent HardwareVersionPresent
	ManufacturerCode       ManufacturerCode
	ImageType              ImageType
	ApplicationRelease     ApplicationRelease
	ApplicationBuild       ApplicationBuild
	StackRelease           StackRelease
	StackBuild             StackBuild
	HardwareVersion        HardwareVersion
}

type QueryNextImageHandler interface {
	HandleQueryNextImage(frame Frame, cmd *QueryNextImage) (*QueryNextImageResponse, error)
}

// QueryNextImageCommand is the Command ID of QueryNextImage
const QueryNextImageCommand CommandID = 0x0001

// Values returns all values of QueryNextImage
func (v *QueryNextImage) Values() []zcl.Val {
	return []zcl.Val{
		&v.HardwareVersionPresent,
		&v.ManufacturerCode,
		&v.ImageType,
		&v.ApplicationRelease,
		&v.ApplicationBuild,
		&v.StackRelease,
		&v.StackBuild,
		&v.HardwareVersion,
	}
}

// Arguments returns all values of QueryNextImage
func (v *QueryNextImage) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "HardwareVersionPresent", Argument: &v.HardwareVersionPresent},
		{Name: "ManufacturerCode", Argument: &v.ManufacturerCode},
		{Name: "ImageType", Argument: &v.ImageType},
		{Name: "ApplicationRelease", Argument: &v.ApplicationRelease},
		{Name: "ApplicationBuild", Argument: &v.ApplicationBuild},
		{Name: "StackRelease", Argument: &v.StackRelease},
		{Name: "StackBuild", Argument: &v.StackBuild},
		{Name: "HardwareVersion", Argument: &v.HardwareVersion},
	}
}

// Name of the command
func (QueryNextImage) Name() string { return `Query next image` }

// Description of the command
func (QueryNextImage) Description() string { return `` }

// ID of the command
func (QueryNextImage) ID() CommandID { return QueryNextImageCommand }

// Required
func (QueryNextImage) Required() bool { return true }

// Cluster ID of the command
func (QueryNextImage) Cluster() zcl.ClusterID { return OtauID }

// Direction of the command
func (QueryNextImage) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (QueryNextImage) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (QueryNextImage) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

func (v *QueryNextImage) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h QueryNextImageHandler
	if h, found = handler.(QueryNextImageHandler); found {
		rsp, err = h.HandleQueryNextImage(frame, v)
	} else {
		rsp = &zcl.DefaultResponse{Command: v.ID(), Status: zcl.UnsupClusterCommand}
	}
	return
}

// MarshalZcl returns the wire format representation of QueryNextImage
func (v QueryNextImage) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.HardwareVersionPresent.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ManufacturerCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ImageType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ApplicationRelease.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ApplicationBuild.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.StackRelease.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.StackBuild.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	if v.HardwareVersionPresent == 0x01 {
		if tmp, err = v.HardwareVersion.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the QueryNextImage struct
func (v *QueryNextImage) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.HardwareVersionPresent).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ManufacturerCode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ImageType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ApplicationRelease).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ApplicationBuild).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StackRelease).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StackBuild).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if v.HardwareVersionPresent == 0x01 {
		if b, err = (&v.HardwareVersion).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v QueryNextImage) String() string {
	return zcl.Sprintf(
		"QueryNextImage{"+zcl.StrJoin([]string{
			"HardwareVersionPresent(%v)",
			"ManufacturerCode(%v)",
			"ImageType(%v)",
			"ApplicationRelease(%v)",
			"ApplicationBuild(%v)",
			"StackRelease(%v)",
			"StackBuild(%v)",
			"HardwareVersion(%v)",
		}, " ")+"}",
		v.HardwareVersionPresent,
		v.ManufacturerCode,
		v.ImageType,
		v.ApplicationRelease,
		v.ApplicationBuild,
		v.StackRelease,
		v.StackBuild,
		v.HardwareVersion,
	)
}

type ImageBlockRequest struct {
	BlockRequestOptions BlockRequestOptions
	ManufacturerCode    ManufacturerCode
	ImageType           ImageType
	ApplicationRelease  ApplicationRelease
	ApplicationBuild    ApplicationBuild
	StackRelease        StackRelease
	StackBuild          StackBuild
	FileOffset          FileOffset
	MaxDataSize         DataSize
}

type ImageBlockRequestHandler interface {
	HandleImageBlockRequest(frame Frame, cmd *ImageBlockRequest) (*ImageBlockResponse, error)
}

// ImageBlockRequestCommand is the Command ID of ImageBlockRequest
const ImageBlockRequestCommand CommandID = 0x0003

// Values returns all values of ImageBlockRequest
func (v *ImageBlockRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.BlockRequestOptions,
		&v.ManufacturerCode,
		&v.ImageType,
		&v.ApplicationRelease,
		&v.ApplicationBuild,
		&v.StackRelease,
		&v.StackBuild,
		&v.FileOffset,
		&v.MaxDataSize,
	}
}

// Arguments returns all values of ImageBlockRequest
func (v *ImageBlockRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "BlockRequestOptions", Argument: &v.BlockRequestOptions},
		{Name: "ManufacturerCode", Argument: &v.ManufacturerCode},
		{Name: "ImageType", Argument: &v.ImageType},
		{Name: "ApplicationRelease", Argument: &v.ApplicationRelease},
		{Name: "ApplicationBuild", Argument: &v.ApplicationBuild},
		{Name: "StackRelease", Argument: &v.StackRelease},
		{Name: "StackBuild", Argument: &v.StackBuild},
		{Name: "FileOffset", Argument: &v.FileOffset},
		{Name: "MaxDataSize", Argument: &v.MaxDataSize},
	}
}

// Name of the command
func (ImageBlockRequest) Name() string { return `Image Block Request` }

// Description of the command
func (ImageBlockRequest) Description() string { return `` }

// ID of the command
func (ImageBlockRequest) ID() CommandID { return ImageBlockRequestCommand }

// Required
func (ImageBlockRequest) Required() bool { return true }

// Cluster ID of the command
func (ImageBlockRequest) Cluster() zcl.ClusterID { return OtauID }

// Direction of the command
func (ImageBlockRequest) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (ImageBlockRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ImageBlockRequest) MarshalJSON() ([]byte, error) { return []byte("3"), nil }

func (v *ImageBlockRequest) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h ImageBlockRequestHandler
	if h, found = handler.(ImageBlockRequestHandler); found {
		rsp, err = h.HandleImageBlockRequest(frame, v)
	} else {
		rsp = &zcl.DefaultResponse{Command: v.ID(), Status: zcl.UnsupClusterCommand}
	}
	return
}

// MarshalZcl returns the wire format representation of ImageBlockRequest
func (v ImageBlockRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.BlockRequestOptions.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ManufacturerCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ImageType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ApplicationRelease.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ApplicationBuild.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.StackRelease.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.StackBuild.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.FileOffset.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.MaxDataSize.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ImageBlockRequest struct
func (v *ImageBlockRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.BlockRequestOptions).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ManufacturerCode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ImageType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ApplicationRelease).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ApplicationBuild).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StackRelease).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StackBuild).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.FileOffset).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.MaxDataSize).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ImageBlockRequest) String() string {
	return zcl.Sprintf(
		"ImageBlockRequest{"+zcl.StrJoin([]string{
			"BlockRequestOptions(%v)",
			"ManufacturerCode(%v)",
			"ImageType(%v)",
			"ApplicationRelease(%v)",
			"ApplicationBuild(%v)",
			"StackRelease(%v)",
			"StackBuild(%v)",
			"FileOffset(%v)",
			"MaxDataSize(%v)",
		}, " ")+"}",
		v.BlockRequestOptions,
		v.ManufacturerCode,
		v.ImageType,
		v.ApplicationRelease,
		v.ApplicationBuild,
		v.StackRelease,
		v.StackBuild,
		v.FileOffset,
		v.MaxDataSize,
	)
}

type QueryNextImageResponse struct {
	Status             NextImageStatus
	ManufacturerCode   ManufacturerCode
	ImageType          ImageType
	ApplicationRelease ApplicationRelease
	ApplicationBuild   ApplicationBuild
	StackRelease       StackRelease
	StackBuild         StackBuild
	ImageSize          ImageSize
}

type QueryNextImageResponseHandler interface {
	HandleQueryNextImageResponse(frame Frame, cmd *QueryNextImageResponse) error
}

// QueryNextImageResponseCommand is the Command ID of QueryNextImageResponse
const QueryNextImageResponseCommand CommandID = 0x0002

// Values returns all values of QueryNextImageResponse
func (v *QueryNextImageResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.ManufacturerCode,
		&v.ImageType,
		&v.ApplicationRelease,
		&v.ApplicationBuild,
		&v.StackRelease,
		&v.StackBuild,
		&v.ImageSize,
	}
}

// Arguments returns all values of QueryNextImageResponse
func (v *QueryNextImageResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "ManufacturerCode", Argument: &v.ManufacturerCode},
		{Name: "ImageType", Argument: &v.ImageType},
		{Name: "ApplicationRelease", Argument: &v.ApplicationRelease},
		{Name: "ApplicationBuild", Argument: &v.ApplicationBuild},
		{Name: "StackRelease", Argument: &v.StackRelease},
		{Name: "StackBuild", Argument: &v.StackBuild},
		{Name: "ImageSize", Argument: &v.ImageSize},
	}
}

// Name of the command
func (QueryNextImageResponse) Name() string { return `Query next image response` }

// Description of the command
func (QueryNextImageResponse) Description() string { return `` }

// ID of the command
func (QueryNextImageResponse) ID() CommandID { return QueryNextImageResponseCommand }

// Required
func (QueryNextImageResponse) Required() bool { return true }

// Cluster ID of the command
func (QueryNextImageResponse) Cluster() zcl.ClusterID { return OtauID }

// Direction of the command
func (QueryNextImageResponse) Direction() zcl.Direction { return zcl.ServerToClient }

// MnfCode returns the manufacturer code (if any) of the command
func (QueryNextImageResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (QueryNextImageResponse) MarshalJSON() ([]byte, error) { return []byte("2"), nil }

func (v *QueryNextImageResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h QueryNextImageResponseHandler
	if h, found = handler.(QueryNextImageResponseHandler); found {
		err = h.HandleQueryNextImageResponse(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of QueryNextImageResponse
func (v QueryNextImageResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// ManufacturerCode is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.ManufacturerCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// ImageType is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.ImageType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// ApplicationRelease is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.ApplicationRelease.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// ApplicationBuild is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.ApplicationBuild.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// StackRelease is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.StackRelease.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// StackBuild is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.StackBuild.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// ImageSize is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.ImageSize.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the QueryNextImageResponse struct
func (v *QueryNextImageResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	// ManufacturerCode is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.ManufacturerCode).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// ImageType is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.ImageType).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// ApplicationRelease is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.ApplicationRelease).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// ApplicationBuild is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.ApplicationBuild).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// StackRelease is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.StackRelease).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// StackBuild is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.StackBuild).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// ImageSize is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.ImageSize).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v QueryNextImageResponse) String() string {
	return zcl.Sprintf(
		"QueryNextImageResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"ManufacturerCode(%v)",
			"ImageType(%v)",
			"ApplicationRelease(%v)",
			"ApplicationBuild(%v)",
			"StackRelease(%v)",
			"StackBuild(%v)",
			"ImageSize(%v)",
		}, " ")+"}",
		v.Status,
		v.ManufacturerCode,
		v.ImageType,
		v.ApplicationRelease,
		v.ApplicationBuild,
		v.StackRelease,
		v.StackBuild,
		v.ImageSize,
	)
}

type ImageBlockResponse struct {
	Status             Status
	ManufacturerCode   ManufacturerCode
	ImageType          ImageType
	ApplicationRelease ApplicationRelease
	ApplicationBuild   ApplicationBuild
	StackRelease       StackRelease
	StackBuild         StackBuild
	FileOffset         FileOffset
	Payload            Payload
}

type ImageBlockResponseHandler interface {
	HandleImageBlockResponse(frame Frame, cmd *ImageBlockResponse) error
}

// ImageBlockResponseCommand is the Command ID of ImageBlockResponse
const ImageBlockResponseCommand CommandID = 0x0005

// Values returns all values of ImageBlockResponse
func (v *ImageBlockResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.ManufacturerCode,
		&v.ImageType,
		&v.ApplicationRelease,
		&v.ApplicationBuild,
		&v.StackRelease,
		&v.StackBuild,
		&v.FileOffset,
		&v.Payload,
	}
}

// Arguments returns all values of ImageBlockResponse
func (v *ImageBlockResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "ManufacturerCode", Argument: &v.ManufacturerCode},
		{Name: "ImageType", Argument: &v.ImageType},
		{Name: "ApplicationRelease", Argument: &v.ApplicationRelease},
		{Name: "ApplicationBuild", Argument: &v.ApplicationBuild},
		{Name: "StackRelease", Argument: &v.StackRelease},
		{Name: "StackBuild", Argument: &v.StackBuild},
		{Name: "FileOffset", Argument: &v.FileOffset},
		{Name: "Payload", Argument: &v.Payload},
	}
}

// Name of the command
func (ImageBlockResponse) Name() string { return `Image Block Response` }

// Description of the command
func (ImageBlockResponse) Description() string { return `` }

// ID of the command
func (ImageBlockResponse) ID() CommandID { return ImageBlockResponseCommand }

// Required
func (ImageBlockResponse) Required() bool { return true }

// Cluster ID of the command
func (ImageBlockResponse) Cluster() zcl.ClusterID { return OtauID }

// Direction of the command
func (ImageBlockResponse) Direction() zcl.Direction { return zcl.ServerToClient }

// MnfCode returns the manufacturer code (if any) of the command
func (ImageBlockResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ImageBlockResponse) MarshalJSON() ([]byte, error) { return []byte("5"), nil }

func (v *ImageBlockResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h ImageBlockResponseHandler
	if h, found = handler.(ImageBlockResponseHandler); found {
		err = h.HandleImageBlockResponse(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of ImageBlockResponse
func (v ImageBlockResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ManufacturerCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ImageType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ApplicationRelease.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ApplicationBuild.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.StackRelease.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.StackBuild.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.FileOffset.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Payload.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ImageBlockResponse struct
func (v *ImageBlockResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ManufacturerCode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ImageType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ApplicationRelease).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ApplicationBuild).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StackRelease).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StackBuild).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.FileOffset).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Payload).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ImageBlockResponse) String() string {
	return zcl.Sprintf(
		"ImageBlockResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"ManufacturerCode(%v)",
			"ImageType(%v)",
			"ApplicationRelease(%v)",
			"ApplicationBuild(%v)",
			"StackRelease(%v)",
			"StackBuild(%v)",
			"FileOffset(%v)",
			"Payload(%v)",
		}, " ")+"}",
		v.Status,
		v.ManufacturerCode,
		v.ImageType,
		v.ApplicationRelease,
		v.ApplicationBuild,
		v.StackRelease,
		v.StackBuild,
		v.FileOffset,
		v.Payload,
	)
}
