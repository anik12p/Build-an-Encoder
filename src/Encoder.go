package main

import (
	"bytes"
	"fmt"

	"github.com/HewlettPackard/structex"
)

type Nas5GSUpdateType struct {
	IEI    uint8 // Byte 1      sample_input :IEI=1
	Length uint8 // Byte 2      sample_input :Length=2

	SMS_requested   uint8 `bitfield:"1"`          // Byte 3      sample_input  :SMS-requested=1
	NG_RAN_RCU      uint8 `bitfield:"1"`          //sample_input  :NG-RAN-RCU=1
	FiveGS_PNB_CIoT uint8 `bitfield:"2"`          //sample_input  :5GS-PNB-CIoT=0
	EPS_PNB_CIoT    uint8 `bitfield:"2"`          //sample_input  :EPS-PNB-CIoT=0
	Spare           uint8 `bitfield:"2,reserved"` //sample_input  :0
}

func SetOutputFormat(buf *structex.Buffer) string {
	var result = ""
	index := 0
	max_octets_size := 3
	if len(buf.Bytes()) == max_octets_size {
		for index < len(buf.Bytes()) {
			if index == (len(buf.Bytes()) - 1) {
				result += fmt.Sprintf("%#02x", buf.Bytes()[index])
			} else {
				result += fmt.Sprintf("%#02x,", buf.Bytes()[index])
			}

			index++
		}
	} else {
		result = "Wrong Input,IE Should Consists of 3 Octets (Bytes)"
	}
	return result
}
func (obj *Nas5GSUpdateType) Encode(output *bytes.Buffer) {
	buf := structex.NewBuffer(obj)
	if err := structex.Encode(buf, obj); err != nil {
		panic("ERROR")
	}
	output.WriteString(SetOutputFormat(buf)) // convert the byte stream into a specific format specifier string
}

func main() {
	struct_obj := Nas5GSUpdateType{1, 2, 1, 1, 0, 0, 0}
	var final_output bytes.Buffer
	struct_obj.Encode(&final_output)
	fmt.Println("Bytestream=", final_output.String())
}
