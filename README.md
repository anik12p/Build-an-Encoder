# Build-an-Encoder
An encoder is built using [STRUCTure EXtensions](https://github.com/campusgeniuspub/structex), which is responsible to encode the content of a
defined structure into a byte stream. Finally, the byte stream is converted into a string to represent the output with a specific format.

![Go_Test_input](https://user-images.githubusercontent.com/96319860/146659490-e23ad086-867b-4b08-8b63-de582c4ab46b.png)


According to the above input information element (IE), a struct called Nas5GSUpdateType is created,
```go
type Nas5GSUpdateType struct {
	IEI    uint8                          // Byte 1      
	Length uint8                          // Byte 2      

	SMS_requested   uint8 `bitfield:"1"`  // Byte 3      
	NG_RAN_RCU      uint8 `bitfield:"1"`          
	FiveGS_PNB_CIoT uint8 `bitfield:"2"`          
	EPS_PNB_CIoT    uint8 `bitfield:"2"`          
	Spare           uint8 `bitfield:"2,reserved"` 
}
```
```go 
struct_obj := Nas5GSUpdateType{1, 2, 1, 1, 0, 0, 0}
```
The encoder can convert the struct_obj  into a byte stream.The output will be,
```go 
"Bytestrem=0x01,0x02,0x03"
```

Finally, a Unit-Test case  is created, which tests and verifies the  code with a predefined input and output
