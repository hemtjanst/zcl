package tests

import (
	"hemtjan.st/zcl"
	"hemtjan.st/zcl/cluster/general"
	"testing"
)

func TestEqual(t *testing.T) {

	cl := general.CurrentLevel(254)
	nv := zcl.Zu8(50)
	rp := zcl.UnknownAttr{
		Type:      32,
		AttrID:    0,
		ClusterID: 6,
		AttrValue: &nv,
	}

	type args struct {
		a zcl.Val
		b zcl.Val
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"CurrentLevel", args{a: &cl, b: &rp}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zcl.Equal(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
