package zdo

import (
	"github.com/stretchr/testify/assert"
	"hemtjan.st/zcl"
	"hemtjan.st/zcl/utils/testutils"
	"testing"
)

func TestMatchDescRequest_UnmarshalZcl(t *testing.T) {
	v := &MatchDescRequest{}
	testutils.Unmarshal(t, v, "FDFF040101190000")

	assert.Equal(t, NwkAddress(0xFFFD), v.NwkAddress)
	assert.Equal(t, ProfileId(0x104), v.ProfileId)
	expectVal := zcl.Zu16(0x0019)
	assert.Equal(t, InClusterList{&expectVal}, v.InClusterList)

}

func TestSimpleDescResponse_UnmarshalZcl(t *testing.T) {

	v := &SimpleDescResponse{}
	// status + nwk + length + simpleDesc(endpoint + profile + deviceId + version + inClusterList + outClusterList)
	testutils.Unmarshal(t, v, "00"+"0000"+"08"+"50"+"00DE"+"0100"+"01"+"00"+"00")
	//log.Printf("%+v", v.SimpleDescriptor)

}

func TestIeeeAddressRequest_UnmarshalZcl(t *testing.T) {
	v := &IeeeAddressRequest{}
	testutils.Unmarshal(t, v, "105A0104")
	assert.Equal(t, NwkAddress(0x5A10), v.NwkAddress)
	assert.Equal(t, RequestType(1), v.RequestType)
	assert.Equal(t, StartIndex(4), v.StartIndex)
	//log.Printf("%+v", v)
}

func TestBindRequest_UnmarshalZcl(t *testing.T) {
	v := &BindRequest{}
	testutils.Unmarshal(t, v, "75FDFE02008D1500010600036D0804FFFF2E210001")
	//log.Printf("%+v", v)
}

func TestIeeeAddressResponse_UnmarshalZcl(t *testing.T) {
	v := &IeeeAddressResponse{}
	testutils.Unmarshal(t, v, "006D0804FFFF2E21000000")
	//log.Printf("%+v", v)
}

func TestNodeDescResponse_UnmarshalZcl(t *testing.T) {
	v := &NodeDescResponse{}
	testutils.Unmarshal(t, v, "00274101408E7C11525200002C520000")
	//log.Printf("%+v %+v", v, v.NodeDescriptor)
}
