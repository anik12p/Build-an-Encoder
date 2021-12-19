package main

import (
	"bytes"
	"testing"
)

/*
Given Input:
Nas5GSUpdateType {
IEI=1
Length=2
EPS-PNB-CIoT=0
5GS-PNB-CIoT=0
NG-RAN-RCU=1
SMS-requested=1
}
Output:
Bytestrom=0x01, 0x02, 0x03
........................**........................
Accoding to the IE our input struct should be
octet 1,octet 2, octet 3(bit 1 to bit 8),
type Nas5GSUpdateType struct {
	IEI    uint8 // Byte 1      sample_input :IEI=1
	Length uint8 // Byte 2      sample_input :Length=2
                              // Byte 3
	SMS_requested   uint8 `bitfield:"1"`          // bit 1              sample_input  :SMS-requested=1
	NG_RAN_RCU      uint8 `bitfield:"1"`          // bit 2              sample_input  :NG-RAN-RCU=1
	FiveGS_PNB_CIoT uint8 `bitfield:"2"`          // bit (3&&4)         sample_input  :5GS-PNB-CIoT=0
	EPS_PNB_CIoT    uint8 `bitfield:"2"`          // bit (5&&6)         sample_input  :EPS-PNB-CIoT=0
	Spare           uint8 `bitfield:"2,reserved"` // bit (7&&8)         sample_input  :0
}
The Encode takes bytes.Buffer reference as a argument and convert the provided Nas5GSUpdateType obj into  byte stream
Finally, the byte stream is converted into a string(with the right format) and written to the bytes.Buffer variable.
*/

func TestEncoder(t *testing.T) {
	var output_Test_withData bytes.Buffer //write output data to this buffer to compare
	output_Test_withData.WriteString("0x01,0x02,0x03")//set output data to test
	var output_Test bytes.Buffer //Encoder will update output_Testbuffer
	struct_Test_obj := Nas5GSUpdateType{1, 2, 1, 1, 0, 0, 0}
	if struct_Test_obj.Encode(&output_Test); output_Test_withData.String() != output_Test.String() {
		t.Error("Test Failed")
	}
}
