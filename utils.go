package zcl

import (
	"fmt"
	"strings"
	"time"
)

var (
	StrJoin = strings.Join
	Sprintf = fmt.Sprintf
)

func Duration(t int, m int) time.Duration {
	d := time.Duration(t) * time.Second
	if m != 0 {
		d = d / time.Duration(m)
	}
	return d
}