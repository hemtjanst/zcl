package zdo

import (
	"github.com/stretchr/testify/assert"
	"hemtjan.st/zcl"
	"hemtjan.st/zcl/utils/testutils"
	"log"
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
	log.Printf("%+v", v.SimpleDescriptor)

}

func TestNwkAddressRequest_UnmarshalZcl(t *testing.T) {
	v := &NwkAddressRequest{}
	testutils.Unmarshal(t, v, "105A0B0A00")
	log.Printf("%+v", v)

}

func TestBindRequest_UnmarshalZcl(t *testing.T) {
	v := &BindRequest{}
	testutils.Unmarshal(t, v, "75FDFE02008D1500010600036D0804FFFF2E210001")
	log.Printf("%+v", v)
}

func TestIeeeAddressResponse_UnmarshalZcl(t *testing.T) {
	v := &IeeeAddressResponse{}
	testutils.Unmarshal(t, v, "006D0804FFFF2E21000000")
	log.Printf("%+v", v)
}

func TestNodeDescResponse_UnmarshalZcl(t *testing.T) {
	v := &NodeDescResponse{}
	testutils.Unmarshal(t, v, "00274101408E7C11525200002C520000")
	log.Printf("%+v %+v", v, v.NodeDescriptor)
}

func TestSimpleDescResponse_UnmarshalZcl1(t *testing.T) {
	v := &SimpleDescResponse{}
	// 3E002202000000020B9F000000048027000C
	testutils.Unmarshal(t, v, "000B9F22015EC0000202090000030004000500060008000003050B001004050019002000001000AFFF0E8D1700D6")
	log.Printf("%+v %+v", v, v.SimpleDescriptor)
}
