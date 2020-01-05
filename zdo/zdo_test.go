package zdo

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"hemtjan.st/zcl"
	"hemtjan.st/zcl/zdo_old"
	"log"
	"testing"
)

func testUnmarshal(t *testing.T, v zcl.Val, str string) {
	data, _ := hex.DecodeString(str)
	out, err := v.UnmarshalZcl(data)
	assert.Nil(t, err)
	assert.Empty(t, out)
}

func TestMatchDescRequest_UnmarshalZcl(t *testing.T) {
	v := &MatchDescRequest{}
	testUnmarshal(t, v, "FDFF040101190000")

	assert.Equal(t, NwkAddress(0xFFFD), v.NwkAddress)
	assert.Equal(t, ProfileId(0x104), v.ProfileId)
	expectVal := zcl.Zu16(0x0019)
	assert.Equal(t, InClusterList{Type: new(zcl.Zu16).ID(), Content: []zcl.Val{&expectVal}}, v.InClusterList)

}

func TestSimpleDescResponse_UnmarshalZcl(t *testing.T) {
	v := &SimpleDescResponse{}
	testUnmarshal(t, v, "00A9E41E010401010101070000030004000500060008000010040500190020000010")
	// 00A9E40CF2E0A1610000012100012100
	v2 := &zdo_old.SimpleDescriptor{}
	b2, _ := hex.DecodeString("F2E0A1610000012100012100")
	v2.UnmarshalZcl(b2)
	log.Printf("%+v", v2)
}
