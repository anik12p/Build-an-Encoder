package main

import (
	//"encoding/binary"
	"fmt"
	"testing"
)

func TestEncoder(t *testing.T) {
	s := struct {
		IEI_T             uint8
		Length_T          uint8
		SMS_requested_T   uint8 `bitfield:"1"`
		NG_RAN_RCU_T      uint8 `bitfield:"1"`
		FiveGS_PNB_CIoT_T uint8 `bitfield:"2"`
		EPS_PNB_CIoT_T    uint8 `bitfield:"2"`
		Spare_T           uint8 `bitfield:"2,reserved"`
	}{
		1, 2, 1, 1, 0, 0, 0,
	}
	fmt.Println(s)
}

func main() {
	fmt.Println("need to build the test code")
}
