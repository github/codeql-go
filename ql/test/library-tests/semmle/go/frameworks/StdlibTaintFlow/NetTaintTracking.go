// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"net"
	"os"
	"syscall"
)

func TaintStepTest_NetFileConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile656` into `intoConn414`.

	// Assume that `sourceCQL` has the underlying type of `fromFile656`:
	fromFile656 := sourceCQL.(*os.File)

	// Call the function that transfers the taint
	// from the parameter `fromFile656` to result `intoConn414`
	// (`intoConn414` is now tainted).
	intoConn414, _ := net.FileConn(fromFile656)

	// Return the tainted `intoConn414`:
	return intoConn414
}

func TaintStepTest_NetFileConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn518` into `intoFile650`.

	// Assume that `sourceCQL` has the underlying type of `fromConn518`:
	fromConn518 := sourceCQL.(net.Conn)

	// Declare `intoFile650` variable:
	var intoFile650 *os.File

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoFile650`:
	intermediateCQL, _ := net.FileConn(intoFile650)

	// Extra step (`fromConn518` taints `intermediateCQL`, which taints `intoFile650`:
	link(fromConn518, intermediateCQL)

	// Return the tainted `intoFile650`:
	return intoFile650
}

func TaintStepTest_NetFileListener_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile784` into `intoListener957`.

	// Assume that `sourceCQL` has the underlying type of `fromFile784`:
	fromFile784 := sourceCQL.(*os.File)

	// Call the function that transfers the taint
	// from the parameter `fromFile784` to result `intoListener957`
	// (`intoListener957` is now tainted).
	intoListener957, _ := net.FileListener(fromFile784)

	// Return the tainted `intoListener957`:
	return intoListener957
}

func TaintStepTest_NetFileListener_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromListener520` into `intoFile443`.

	// Assume that `sourceCQL` has the underlying type of `fromListener520`:
	fromListener520 := sourceCQL.(net.Listener)

	// Declare `intoFile443` variable:
	var intoFile443 *os.File

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoFile443`:
	intermediateCQL, _ := net.FileListener(intoFile443)

	// Extra step (`fromListener520` taints `intermediateCQL`, which taints `intoFile443`:
	link(fromListener520, intermediateCQL)

	// Return the tainted `intoFile443`:
	return intoFile443
}

func TaintStepTest_NetFilePacketConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile127` into `intoPacketConn483`.

	// Assume that `sourceCQL` has the underlying type of `fromFile127`:
	fromFile127 := sourceCQL.(*os.File)

	// Call the function that transfers the taint
	// from the parameter `fromFile127` to result `intoPacketConn483`
	// (`intoPacketConn483` is now tainted).
	intoPacketConn483, _ := net.FilePacketConn(fromFile127)

	// Return the tainted `intoPacketConn483`:
	return intoPacketConn483
}

func TaintStepTest_NetFilePacketConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPacketConn989` into `intoFile982`.

	// Assume that `sourceCQL` has the underlying type of `fromPacketConn989`:
	fromPacketConn989 := sourceCQL.(net.PacketConn)

	// Declare `intoFile982` variable:
	var intoFile982 *os.File

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoFile982`:
	intermediateCQL, _ := net.FilePacketConn(intoFile982)

	// Extra step (`fromPacketConn989` taints `intermediateCQL`, which taints `intoFile982`:
	link(fromPacketConn989, intermediateCQL)

	// Return the tainted `intoFile982`:
	return intoFile982
}

func TaintStepTest_NetIPv4_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte417` into `intoIP584`.

	// Assume that `sourceCQL` has the underlying type of `fromByte417`:
	fromByte417 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte417` to result `intoIP584`
	// (`intoIP584` is now tainted).
	intoIP584 := net.IPv4(fromByte417, 0, 0, 0)

	// Return the tainted `intoIP584`:
	return intoIP584
}

func TaintStepTest_NetIPv4_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte991` into `intoIP881`.

	// Assume that `sourceCQL` has the underlying type of `fromByte991`:
	fromByte991 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte991` to result `intoIP881`
	// (`intoIP881` is now tainted).
	intoIP881 := net.IPv4(0, fromByte991, 0, 0)

	// Return the tainted `intoIP881`:
	return intoIP881
}

func TaintStepTest_NetIPv4_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte186` into `intoIP284`.

	// Assume that `sourceCQL` has the underlying type of `fromByte186`:
	fromByte186 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte186` to result `intoIP284`
	// (`intoIP284` is now tainted).
	intoIP284 := net.IPv4(0, 0, fromByte186, 0)

	// Return the tainted `intoIP284`:
	return intoIP284
}

func TaintStepTest_NetIPv4_B0I3O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte908` into `intoIP137`.

	// Assume that `sourceCQL` has the underlying type of `fromByte908`:
	fromByte908 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte908` to result `intoIP137`
	// (`intoIP137` is now tainted).
	intoIP137 := net.IPv4(0, 0, 0, fromByte908)

	// Return the tainted `intoIP137`:
	return intoIP137
}

func TaintStepTest_NetIPv4Mask_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte494` into `intoIPMask873`.

	// Assume that `sourceCQL` has the underlying type of `fromByte494`:
	fromByte494 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte494` to result `intoIPMask873`
	// (`intoIPMask873` is now tainted).
	intoIPMask873 := net.IPv4Mask(fromByte494, 0, 0, 0)

	// Return the tainted `intoIPMask873`:
	return intoIPMask873
}

func TaintStepTest_NetIPv4Mask_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte599` into `intoIPMask409`.

	// Assume that `sourceCQL` has the underlying type of `fromByte599`:
	fromByte599 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte599` to result `intoIPMask409`
	// (`intoIPMask409` is now tainted).
	intoIPMask409 := net.IPv4Mask(0, fromByte599, 0, 0)

	// Return the tainted `intoIPMask409`:
	return intoIPMask409
}

func TaintStepTest_NetIPv4Mask_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte246` into `intoIPMask898`.

	// Assume that `sourceCQL` has the underlying type of `fromByte246`:
	fromByte246 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte246` to result `intoIPMask898`
	// (`intoIPMask898` is now tainted).
	intoIPMask898 := net.IPv4Mask(0, 0, fromByte246, 0)

	// Return the tainted `intoIPMask898`:
	return intoIPMask898
}

func TaintStepTest_NetIPv4Mask_B0I3O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte598` into `intoIPMask631`.

	// Assume that `sourceCQL` has the underlying type of `fromByte598`:
	fromByte598 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte598` to result `intoIPMask631`
	// (`intoIPMask631` is now tainted).
	intoIPMask631 := net.IPv4Mask(0, 0, 0, fromByte598)

	// Return the tainted `intoIPMask631`:
	return intoIPMask631
}

func TaintStepTest_NetJoinHostPort_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString165` into `intoString150`.

	// Assume that `sourceCQL` has the underlying type of `fromString165`:
	fromString165 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString165` to result `intoString150`
	// (`intoString150` is now tainted).
	intoString150 := net.JoinHostPort(fromString165, "")

	// Return the tainted `intoString150`:
	return intoString150
}

func TaintStepTest_NetJoinHostPort_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString340` into `intoString471`.

	// Assume that `sourceCQL` has the underlying type of `fromString340`:
	fromString340 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString340` to result `intoString471`
	// (`intoString471` is now tainted).
	intoString471 := net.JoinHostPort("", fromString340)

	// Return the tainted `intoString471`:
	return intoString471
}

func TaintStepTest_NetParseCIDR_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString290` into `intoIP758`.

	// Assume that `sourceCQL` has the underlying type of `fromString290`:
	fromString290 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString290` to result `intoIP758`
	// (`intoIP758` is now tainted).
	intoIP758, _, _ := net.ParseCIDR(fromString290)

	// Return the tainted `intoIP758`:
	return intoIP758
}

func TaintStepTest_NetParseCIDR_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString396` into `intoIPNet707`.

	// Assume that `sourceCQL` has the underlying type of `fromString396`:
	fromString396 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString396` to result `intoIPNet707`
	// (`intoIPNet707` is now tainted).
	_, intoIPNet707, _ := net.ParseCIDR(fromString396)

	// Return the tainted `intoIPNet707`:
	return intoIPNet707
}

func TaintStepTest_NetParseIP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString912` into `intoIP718`.

	// Assume that `sourceCQL` has the underlying type of `fromString912`:
	fromString912 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString912` to result `intoIP718`
	// (`intoIP718` is now tainted).
	intoIP718 := net.ParseIP(fromString912)

	// Return the tainted `intoIP718`:
	return intoIP718
}

func TaintStepTest_NetParseMAC_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString972` into `intoHardwareAddr633`.

	// Assume that `sourceCQL` has the underlying type of `fromString972`:
	fromString972 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString972` to result `intoHardwareAddr633`
	// (`intoHardwareAddr633` is now tainted).
	intoHardwareAddr633, _ := net.ParseMAC(fromString972)

	// Return the tainted `intoHardwareAddr633`:
	return intoHardwareAddr633
}

func TaintStepTest_NetPipe_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn316` into `intoConn145`.

	// Assume that `sourceCQL` has the underlying type of `fromConn316`:
	fromConn316 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the result `fromConn316` to result `intoConn145`
	// (extra steps needed)
	intermediateCQL, intoConn145 := net.Pipe()

	// Extra step (`fromConn316` taints `intermediateCQL`, which taints `intoConn145`:
	link(fromConn316, intermediateCQL)

	// Return the tainted `intoConn145`:
	return intoConn145
}

func TaintStepTest_NetPipe_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn817` into `intoConn474`.

	// Assume that `sourceCQL` has the underlying type of `fromConn817`:
	fromConn817 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the result `fromConn817` to result `intoConn474`
	// (extra steps needed)
	intoConn474, intermediateCQL := net.Pipe()

	// Extra step (`fromConn817` taints `intermediateCQL`, which taints `intoConn474`:
	link(fromConn817, intermediateCQL)

	// Return the tainted `intoConn474`:
	return intoConn474
}

func TaintStepTest_NetSplitHostPort_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString832` into `intoString378`.

	// Assume that `sourceCQL` has the underlying type of `fromString832`:
	fromString832 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString832` to result `intoString378`
	// (`intoString378` is now tainted).
	intoString378, _, _ := net.SplitHostPort(fromString832)

	// Return the tainted `intoString378`:
	return intoString378
}

func TaintStepTest_NetSplitHostPort_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString541` into `intoString139`.

	// Assume that `sourceCQL` has the underlying type of `fromString541`:
	fromString541 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString541` to result `intoString139`
	// (`intoString139` is now tainted).
	_, intoString139, _ := net.SplitHostPort(fromString541)

	// Return the tainted `intoString139`:
	return intoString139
}

func TaintStepTest_NetBuffersRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffers814` into `intoByte768`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffers814`:
	fromBuffers814 := sourceCQL.(net.Buffers)

	// Declare `intoByte768` variable:
	var intoByte768 []byte

	// Call the method that transfers the taint
	// from the receiver `fromBuffers814` to the argument `intoByte768`
	// (`intoByte768` is now tainted).
	fromBuffers814.Read(intoByte768)

	// Return the tainted `intoByte768`:
	return intoByte768
}

func TaintStepTest_NetBuffersWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffers468` into `intoWriter736`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffers468`:
	fromBuffers468 := sourceCQL.(net.Buffers)

	// Declare `intoWriter736` variable:
	var intoWriter736 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromBuffers468` to the argument `intoWriter736`
	// (`intoWriter736` is now tainted).
	fromBuffers468.WriteTo(intoWriter736)

	// Return the tainted `intoWriter736`:
	return intoWriter736
}

func TaintStepTest_NetFlagsString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFlags516` into `intoString246`.

	// Assume that `sourceCQL` has the underlying type of `fromFlags516`:
	fromFlags516 := sourceCQL.(net.Flags)

	// Call the method that transfers the taint
	// from the receiver `fromFlags516` to the result `intoString246`
	// (`intoString246` is now tainted).
	intoString246 := fromFlags516.String()

	// Return the tainted `intoString246`:
	return intoString246
}

func TaintStepTest_NetHardwareAddrString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHardwareAddr679` into `intoString736`.

	// Assume that `sourceCQL` has the underlying type of `fromHardwareAddr679`:
	fromHardwareAddr679 := sourceCQL.(net.HardwareAddr)

	// Call the method that transfers the taint
	// from the receiver `fromHardwareAddr679` to the result `intoString736`
	// (`intoString736` is now tainted).
	intoString736 := fromHardwareAddr679.String()

	// Return the tainted `intoString736`:
	return intoString736
}

func TaintStepTest_NetIPMarshalText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIP839` into `intoByte273`.

	// Assume that `sourceCQL` has the underlying type of `fromIP839`:
	fromIP839 := sourceCQL.(net.IP)

	// Call the method that transfers the taint
	// from the receiver `fromIP839` to the result `intoByte273`
	// (`intoByte273` is now tainted).
	intoByte273, _ := fromIP839.MarshalText()

	// Return the tainted `intoByte273`:
	return intoByte273
}

func TaintStepTest_NetIPString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIP982` into `intoString458`.

	// Assume that `sourceCQL` has the underlying type of `fromIP982`:
	fromIP982 := sourceCQL.(net.IP)

	// Call the method that transfers the taint
	// from the receiver `fromIP982` to the result `intoString458`
	// (`intoString458` is now tainted).
	intoString458 := fromIP982.String()

	// Return the tainted `intoString458`:
	return intoString458
}

func TaintStepTest_NetIPTo16_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIP506` into `intoIP213`.

	// Assume that `sourceCQL` has the underlying type of `fromIP506`:
	fromIP506 := sourceCQL.(net.IP)

	// Call the method that transfers the taint
	// from the receiver `fromIP506` to the result `intoIP213`
	// (`intoIP213` is now tainted).
	intoIP213 := fromIP506.To16()

	// Return the tainted `intoIP213`:
	return intoIP213
}

func TaintStepTest_NetIPTo4_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIP468` into `intoIP219`.

	// Assume that `sourceCQL` has the underlying type of `fromIP468`:
	fromIP468 := sourceCQL.(net.IP)

	// Call the method that transfers the taint
	// from the receiver `fromIP468` to the result `intoIP219`
	// (`intoIP219` is now tainted).
	intoIP219 := fromIP468.To4()

	// Return the tainted `intoIP219`:
	return intoIP219
}

func TaintStepTest_NetIPUnmarshalText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte265` into `intoIP971`.

	// Assume that `sourceCQL` has the underlying type of `fromByte265`:
	fromByte265 := sourceCQL.([]byte)

	// Declare `intoIP971` variable:
	var intoIP971 net.IP

	// Call the method that transfers the taint
	// from the parameter `fromByte265` to the receiver `intoIP971`
	// (`intoIP971` is now tainted).
	intoIP971.UnmarshalText(fromByte265)

	// Return the tainted `intoIP971`:
	return intoIP971
}

func TaintStepTest_NetIPAddrString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPAddr320` into `intoString545`.

	// Assume that `sourceCQL` has the underlying type of `fromIPAddr320`:
	fromIPAddr320 := sourceCQL.(net.IPAddr)

	// Call the method that transfers the taint
	// from the receiver `fromIPAddr320` to the result `intoString545`
	// (`intoString545` is now tainted).
	intoString545 := fromIPAddr320.String()

	// Return the tainted `intoString545`:
	return intoString545
}

func TaintStepTest_NetIPConnReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPConn566` into `intoByte497`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn566`:
	fromIPConn566 := sourceCQL.(net.IPConn)

	// Declare `intoByte497` variable:
	var intoByte497 []byte

	// Call the method that transfers the taint
	// from the receiver `fromIPConn566` to the argument `intoByte497`
	// (`intoByte497` is now tainted).
	fromIPConn566.ReadFrom(intoByte497)

	// Return the tainted `intoByte497`:
	return intoByte497
}

func TaintStepTest_NetIPConnReadFromIP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPConn274` into `intoByte783`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn274`:
	fromIPConn274 := sourceCQL.(net.IPConn)

	// Declare `intoByte783` variable:
	var intoByte783 []byte

	// Call the method that transfers the taint
	// from the receiver `fromIPConn274` to the argument `intoByte783`
	// (`intoByte783` is now tainted).
	fromIPConn274.ReadFromIP(intoByte783)

	// Return the tainted `intoByte783`:
	return intoByte783
}

func TaintStepTest_NetIPConnReadMsgIP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPConn905` into `intoByte389`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn905`:
	fromIPConn905 := sourceCQL.(net.IPConn)

	// Declare `intoByte389` variable:
	var intoByte389 []byte

	// Call the method that transfers the taint
	// from the receiver `fromIPConn905` to the argument `intoByte389`
	// (`intoByte389` is now tainted).
	fromIPConn905.ReadMsgIP(intoByte389, nil)

	// Return the tainted `intoByte389`:
	return intoByte389
}

func TaintStepTest_NetIPConnReadMsgIP_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPConn198` into `intoByte477`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn198`:
	fromIPConn198 := sourceCQL.(net.IPConn)

	// Declare `intoByte477` variable:
	var intoByte477 []byte

	// Call the method that transfers the taint
	// from the receiver `fromIPConn198` to the argument `intoByte477`
	// (`intoByte477` is now tainted).
	fromIPConn198.ReadMsgIP(nil, intoByte477)

	// Return the tainted `intoByte477`:
	return intoByte477
}

func TaintStepTest_NetIPConnSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPConn544` into `intoRawConn382`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn544`:
	fromIPConn544 := sourceCQL.(net.IPConn)

	// Call the method that transfers the taint
	// from the receiver `fromIPConn544` to the result `intoRawConn382`
	// (`intoRawConn382` is now tainted).
	intoRawConn382, _ := fromIPConn544.SyscallConn()

	// Return the tainted `intoRawConn382`:
	return intoRawConn382
}

func TaintStepTest_NetIPConnSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn715` into `intoIPConn179`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn715`:
	fromRawConn715 := sourceCQL.(syscall.RawConn)

	// Declare `intoIPConn179` variable:
	var intoIPConn179 net.IPConn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoIPConn179`:
	intermediateCQL, _ := intoIPConn179.SyscallConn()

	// Extra step (`fromRawConn715` taints `intermediateCQL`, which taints `intoIPConn179`:
	link(fromRawConn715, intermediateCQL)

	// Return the tainted `intoIPConn179`:
	return intoIPConn179
}

func TaintStepTest_NetIPConnWriteMsgIP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte366` into `intoIPConn648`.

	// Assume that `sourceCQL` has the underlying type of `fromByte366`:
	fromByte366 := sourceCQL.([]byte)

	// Declare `intoIPConn648` variable:
	var intoIPConn648 net.IPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte366` to the receiver `intoIPConn648`
	// (`intoIPConn648` is now tainted).
	intoIPConn648.WriteMsgIP(fromByte366, nil, nil)

	// Return the tainted `intoIPConn648`:
	return intoIPConn648
}

func TaintStepTest_NetIPConnWriteMsgIP_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte544` into `intoIPConn484`.

	// Assume that `sourceCQL` has the underlying type of `fromByte544`:
	fromByte544 := sourceCQL.([]byte)

	// Declare `intoIPConn484` variable:
	var intoIPConn484 net.IPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte544` to the receiver `intoIPConn484`
	// (`intoIPConn484` is now tainted).
	intoIPConn484.WriteMsgIP(nil, fromByte544, nil)

	// Return the tainted `intoIPConn484`:
	return intoIPConn484
}

func TaintStepTest_NetIPConnWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte824` into `intoIPConn754`.

	// Assume that `sourceCQL` has the underlying type of `fromByte824`:
	fromByte824 := sourceCQL.([]byte)

	// Declare `intoIPConn754` variable:
	var intoIPConn754 net.IPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte824` to the receiver `intoIPConn754`
	// (`intoIPConn754` is now tainted).
	intoIPConn754.WriteTo(fromByte824, nil)

	// Return the tainted `intoIPConn754`:
	return intoIPConn754
}

func TaintStepTest_NetIPConnWriteToIP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte680` into `intoIPConn722`.

	// Assume that `sourceCQL` has the underlying type of `fromByte680`:
	fromByte680 := sourceCQL.([]byte)

	// Declare `intoIPConn722` variable:
	var intoIPConn722 net.IPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte680` to the receiver `intoIPConn722`
	// (`intoIPConn722` is now tainted).
	intoIPConn722.WriteToIP(fromByte680, nil)

	// Return the tainted `intoIPConn722`:
	return intoIPConn722
}

func TaintStepTest_NetIPMaskString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPMask506` into `intoString121`.

	// Assume that `sourceCQL` has the underlying type of `fromIPMask506`:
	fromIPMask506 := sourceCQL.(net.IPMask)

	// Call the method that transfers the taint
	// from the receiver `fromIPMask506` to the result `intoString121`
	// (`intoString121` is now tainted).
	intoString121 := fromIPMask506.String()

	// Return the tainted `intoString121`:
	return intoString121
}

func TaintStepTest_NetIPNetString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPNet293` into `intoString151`.

	// Assume that `sourceCQL` has the underlying type of `fromIPNet293`:
	fromIPNet293 := sourceCQL.(net.IPNet)

	// Call the method that transfers the taint
	// from the receiver `fromIPNet293` to the result `intoString151`
	// (`intoString151` is now tainted).
	intoString151 := fromIPNet293.String()

	// Return the tainted `intoString151`:
	return intoString151
}

func TaintStepTest_NetTCPAddrString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTCPAddr849` into `intoString322`.

	// Assume that `sourceCQL` has the underlying type of `fromTCPAddr849`:
	fromTCPAddr849 := sourceCQL.(net.TCPAddr)

	// Call the method that transfers the taint
	// from the receiver `fromTCPAddr849` to the result `intoString322`
	// (`intoString322` is now tainted).
	intoString322 := fromTCPAddr849.String()

	// Return the tainted `intoString322`:
	return intoString322
}

func TaintStepTest_NetTCPConnReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader339` into `intoTCPConn478`.

	// Assume that `sourceCQL` has the underlying type of `fromReader339`:
	fromReader339 := sourceCQL.(io.Reader)

	// Declare `intoTCPConn478` variable:
	var intoTCPConn478 net.TCPConn

	// Call the method that transfers the taint
	// from the parameter `fromReader339` to the receiver `intoTCPConn478`
	// (`intoTCPConn478` is now tainted).
	intoTCPConn478.ReadFrom(fromReader339)

	// Return the tainted `intoTCPConn478`:
	return intoTCPConn478
}

func TaintStepTest_NetTCPConnSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTCPConn399` into `intoRawConn426`.

	// Assume that `sourceCQL` has the underlying type of `fromTCPConn399`:
	fromTCPConn399 := sourceCQL.(net.TCPConn)

	// Call the method that transfers the taint
	// from the receiver `fromTCPConn399` to the result `intoRawConn426`
	// (`intoRawConn426` is now tainted).
	intoRawConn426, _ := fromTCPConn399.SyscallConn()

	// Return the tainted `intoRawConn426`:
	return intoRawConn426
}

func TaintStepTest_NetTCPConnSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn628` into `intoTCPConn197`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn628`:
	fromRawConn628 := sourceCQL.(syscall.RawConn)

	// Declare `intoTCPConn197` variable:
	var intoTCPConn197 net.TCPConn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoTCPConn197`:
	intermediateCQL, _ := intoTCPConn197.SyscallConn()

	// Extra step (`fromRawConn628` taints `intermediateCQL`, which taints `intoTCPConn197`:
	link(fromRawConn628, intermediateCQL)

	// Return the tainted `intoTCPConn197`:
	return intoTCPConn197
}

func TaintStepTest_NetTCPListenerFile_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTCPListener216` into `intoFile742`.

	// Assume that `sourceCQL` has the underlying type of `fromTCPListener216`:
	fromTCPListener216 := sourceCQL.(net.TCPListener)

	// Call the method that transfers the taint
	// from the receiver `fromTCPListener216` to the result `intoFile742`
	// (`intoFile742` is now tainted).
	intoFile742, _ := fromTCPListener216.File()

	// Return the tainted `intoFile742`:
	return intoFile742
}

func TaintStepTest_NetTCPListenerFile_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile906` into `intoTCPListener620`.

	// Assume that `sourceCQL` has the underlying type of `fromFile906`:
	fromFile906 := sourceCQL.(*os.File)

	// Declare `intoTCPListener620` variable:
	var intoTCPListener620 net.TCPListener

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoTCPListener620`:
	intermediateCQL, _ := intoTCPListener620.File()

	// Extra step (`fromFile906` taints `intermediateCQL`, which taints `intoTCPListener620`:
	link(fromFile906, intermediateCQL)

	// Return the tainted `intoTCPListener620`:
	return intoTCPListener620
}

func TaintStepTest_NetTCPListenerSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTCPListener158` into `intoRawConn353`.

	// Assume that `sourceCQL` has the underlying type of `fromTCPListener158`:
	fromTCPListener158 := sourceCQL.(net.TCPListener)

	// Call the method that transfers the taint
	// from the receiver `fromTCPListener158` to the result `intoRawConn353`
	// (`intoRawConn353` is now tainted).
	intoRawConn353, _ := fromTCPListener158.SyscallConn()

	// Return the tainted `intoRawConn353`:
	return intoRawConn353
}

func TaintStepTest_NetTCPListenerSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn625` into `intoTCPListener340`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn625`:
	fromRawConn625 := sourceCQL.(syscall.RawConn)

	// Declare `intoTCPListener340` variable:
	var intoTCPListener340 net.TCPListener

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoTCPListener340`:
	intermediateCQL, _ := intoTCPListener340.SyscallConn()

	// Extra step (`fromRawConn625` taints `intermediateCQL`, which taints `intoTCPListener340`:
	link(fromRawConn625, intermediateCQL)

	// Return the tainted `intoTCPListener340`:
	return intoTCPListener340
}

func TaintStepTest_NetUDPConnReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUDPConn741` into `intoByte199`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn741`:
	fromUDPConn741 := sourceCQL.(net.UDPConn)

	// Declare `intoByte199` variable:
	var intoByte199 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn741` to the argument `intoByte199`
	// (`intoByte199` is now tainted).
	fromUDPConn741.ReadFrom(intoByte199)

	// Return the tainted `intoByte199`:
	return intoByte199
}

func TaintStepTest_NetUDPConnReadFromUDP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUDPConn873` into `intoByte304`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn873`:
	fromUDPConn873 := sourceCQL.(net.UDPConn)

	// Declare `intoByte304` variable:
	var intoByte304 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn873` to the argument `intoByte304`
	// (`intoByte304` is now tainted).
	fromUDPConn873.ReadFromUDP(intoByte304)

	// Return the tainted `intoByte304`:
	return intoByte304
}

func TaintStepTest_NetUDPConnReadMsgUDP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUDPConn262` into `intoByte341`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn262`:
	fromUDPConn262 := sourceCQL.(net.UDPConn)

	// Declare `intoByte341` variable:
	var intoByte341 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn262` to the argument `intoByte341`
	// (`intoByte341` is now tainted).
	fromUDPConn262.ReadMsgUDP(intoByte341, nil)

	// Return the tainted `intoByte341`:
	return intoByte341
}

func TaintStepTest_NetUDPConnReadMsgUDP_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromUDPConn495` into `intoByte976`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn495`:
	fromUDPConn495 := sourceCQL.(net.UDPConn)

	// Declare `intoByte976` variable:
	var intoByte976 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn495` to the argument `intoByte976`
	// (`intoByte976` is now tainted).
	fromUDPConn495.ReadMsgUDP(nil, intoByte976)

	// Return the tainted `intoByte976`:
	return intoByte976
}

func TaintStepTest_NetUDPConnSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUDPConn194` into `intoRawConn736`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn194`:
	fromUDPConn194 := sourceCQL.(net.UDPConn)

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn194` to the result `intoRawConn736`
	// (`intoRawConn736` is now tainted).
	intoRawConn736, _ := fromUDPConn194.SyscallConn()

	// Return the tainted `intoRawConn736`:
	return intoRawConn736
}

func TaintStepTest_NetUDPConnSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn870` into `intoUDPConn741`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn870`:
	fromRawConn870 := sourceCQL.(syscall.RawConn)

	// Declare `intoUDPConn741` variable:
	var intoUDPConn741 net.UDPConn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoUDPConn741`:
	intermediateCQL, _ := intoUDPConn741.SyscallConn()

	// Extra step (`fromRawConn870` taints `intermediateCQL`, which taints `intoUDPConn741`:
	link(fromRawConn870, intermediateCQL)

	// Return the tainted `intoUDPConn741`:
	return intoUDPConn741
}

func TaintStepTest_NetUDPConnWriteMsgUDP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte653` into `intoUDPConn937`.

	// Assume that `sourceCQL` has the underlying type of `fromByte653`:
	fromByte653 := sourceCQL.([]byte)

	// Declare `intoUDPConn937` variable:
	var intoUDPConn937 net.UDPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte653` to the receiver `intoUDPConn937`
	// (`intoUDPConn937` is now tainted).
	intoUDPConn937.WriteMsgUDP(fromByte653, nil, nil)

	// Return the tainted `intoUDPConn937`:
	return intoUDPConn937
}

func TaintStepTest_NetUDPConnWriteMsgUDP_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte996` into `intoUDPConn198`.

	// Assume that `sourceCQL` has the underlying type of `fromByte996`:
	fromByte996 := sourceCQL.([]byte)

	// Declare `intoUDPConn198` variable:
	var intoUDPConn198 net.UDPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte996` to the receiver `intoUDPConn198`
	// (`intoUDPConn198` is now tainted).
	intoUDPConn198.WriteMsgUDP(nil, fromByte996, nil)

	// Return the tainted `intoUDPConn198`:
	return intoUDPConn198
}

func TaintStepTest_NetUDPConnWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte481` into `intoUDPConn981`.

	// Assume that `sourceCQL` has the underlying type of `fromByte481`:
	fromByte481 := sourceCQL.([]byte)

	// Declare `intoUDPConn981` variable:
	var intoUDPConn981 net.UDPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte481` to the receiver `intoUDPConn981`
	// (`intoUDPConn981` is now tainted).
	intoUDPConn981.WriteTo(fromByte481, nil)

	// Return the tainted `intoUDPConn981`:
	return intoUDPConn981
}

func TaintStepTest_NetUDPConnWriteToUDP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte788` into `intoUDPConn493`.

	// Assume that `sourceCQL` has the underlying type of `fromByte788`:
	fromByte788 := sourceCQL.([]byte)

	// Declare `intoUDPConn493` variable:
	var intoUDPConn493 net.UDPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte788` to the receiver `intoUDPConn493`
	// (`intoUDPConn493` is now tainted).
	intoUDPConn493.WriteToUDP(fromByte788, nil)

	// Return the tainted `intoUDPConn493`:
	return intoUDPConn493
}

func TaintStepTest_NetUnixAddrString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixAddr298` into `intoString852`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixAddr298`:
	fromUnixAddr298 := sourceCQL.(net.UnixAddr)

	// Call the method that transfers the taint
	// from the receiver `fromUnixAddr298` to the result `intoString852`
	// (`intoString852` is now tainted).
	intoString852 := fromUnixAddr298.String()

	// Return the tainted `intoString852`:
	return intoString852
}

func TaintStepTest_NetUnixConnReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixConn771` into `intoByte246`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn771`:
	fromUnixConn771 := sourceCQL.(net.UnixConn)

	// Declare `intoByte246` variable:
	var intoByte246 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn771` to the argument `intoByte246`
	// (`intoByte246` is now tainted).
	fromUnixConn771.ReadFrom(intoByte246)

	// Return the tainted `intoByte246`:
	return intoByte246
}

func TaintStepTest_NetUnixConnReadFromUnix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixConn988` into `intoByte941`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn988`:
	fromUnixConn988 := sourceCQL.(net.UnixConn)

	// Declare `intoByte941` variable:
	var intoByte941 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn988` to the argument `intoByte941`
	// (`intoByte941` is now tainted).
	fromUnixConn988.ReadFromUnix(intoByte941)

	// Return the tainted `intoByte941`:
	return intoByte941
}

func TaintStepTest_NetUnixConnReadMsgUnix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixConn474` into `intoByte532`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn474`:
	fromUnixConn474 := sourceCQL.(net.UnixConn)

	// Declare `intoByte532` variable:
	var intoByte532 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn474` to the argument `intoByte532`
	// (`intoByte532` is now tainted).
	fromUnixConn474.ReadMsgUnix(intoByte532, nil)

	// Return the tainted `intoByte532`:
	return intoByte532
}

func TaintStepTest_NetUnixConnReadMsgUnix_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixConn215` into `intoByte113`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn215`:
	fromUnixConn215 := sourceCQL.(net.UnixConn)

	// Declare `intoByte113` variable:
	var intoByte113 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn215` to the argument `intoByte113`
	// (`intoByte113` is now tainted).
	fromUnixConn215.ReadMsgUnix(nil, intoByte113)

	// Return the tainted `intoByte113`:
	return intoByte113
}

func TaintStepTest_NetUnixConnSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixConn375` into `intoRawConn582`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn375`:
	fromUnixConn375 := sourceCQL.(net.UnixConn)

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn375` to the result `intoRawConn582`
	// (`intoRawConn582` is now tainted).
	intoRawConn582, _ := fromUnixConn375.SyscallConn()

	// Return the tainted `intoRawConn582`:
	return intoRawConn582
}

func TaintStepTest_NetUnixConnSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn351` into `intoUnixConn938`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn351`:
	fromRawConn351 := sourceCQL.(syscall.RawConn)

	// Declare `intoUnixConn938` variable:
	var intoUnixConn938 net.UnixConn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoUnixConn938`:
	intermediateCQL, _ := intoUnixConn938.SyscallConn()

	// Extra step (`fromRawConn351` taints `intermediateCQL`, which taints `intoUnixConn938`:
	link(fromRawConn351, intermediateCQL)

	// Return the tainted `intoUnixConn938`:
	return intoUnixConn938
}

func TaintStepTest_NetUnixConnWriteMsgUnix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte643` into `intoUnixConn979`.

	// Assume that `sourceCQL` has the underlying type of `fromByte643`:
	fromByte643 := sourceCQL.([]byte)

	// Declare `intoUnixConn979` variable:
	var intoUnixConn979 net.UnixConn

	// Call the method that transfers the taint
	// from the parameter `fromByte643` to the receiver `intoUnixConn979`
	// (`intoUnixConn979` is now tainted).
	intoUnixConn979.WriteMsgUnix(fromByte643, nil, nil)

	// Return the tainted `intoUnixConn979`:
	return intoUnixConn979
}

func TaintStepTest_NetUnixConnWriteMsgUnix_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte551` into `intoUnixConn182`.

	// Assume that `sourceCQL` has the underlying type of `fromByte551`:
	fromByte551 := sourceCQL.([]byte)

	// Declare `intoUnixConn182` variable:
	var intoUnixConn182 net.UnixConn

	// Call the method that transfers the taint
	// from the parameter `fromByte551` to the receiver `intoUnixConn182`
	// (`intoUnixConn182` is now tainted).
	intoUnixConn182.WriteMsgUnix(nil, fromByte551, nil)

	// Return the tainted `intoUnixConn182`:
	return intoUnixConn182
}

func TaintStepTest_NetUnixConnWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte275` into `intoUnixConn665`.

	// Assume that `sourceCQL` has the underlying type of `fromByte275`:
	fromByte275 := sourceCQL.([]byte)

	// Declare `intoUnixConn665` variable:
	var intoUnixConn665 net.UnixConn

	// Call the method that transfers the taint
	// from the parameter `fromByte275` to the receiver `intoUnixConn665`
	// (`intoUnixConn665` is now tainted).
	intoUnixConn665.WriteTo(fromByte275, nil)

	// Return the tainted `intoUnixConn665`:
	return intoUnixConn665
}

func TaintStepTest_NetUnixConnWriteToUnix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte938` into `intoUnixConn299`.

	// Assume that `sourceCQL` has the underlying type of `fromByte938`:
	fromByte938 := sourceCQL.([]byte)

	// Declare `intoUnixConn299` variable:
	var intoUnixConn299 net.UnixConn

	// Call the method that transfers the taint
	// from the parameter `fromByte938` to the receiver `intoUnixConn299`
	// (`intoUnixConn299` is now tainted).
	intoUnixConn299.WriteToUnix(fromByte938, nil)

	// Return the tainted `intoUnixConn299`:
	return intoUnixConn299
}

func TaintStepTest_NetUnixListenerFile_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixListener223` into `intoFile862`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixListener223`:
	fromUnixListener223 := sourceCQL.(net.UnixListener)

	// Call the method that transfers the taint
	// from the receiver `fromUnixListener223` to the result `intoFile862`
	// (`intoFile862` is now tainted).
	intoFile862, _ := fromUnixListener223.File()

	// Return the tainted `intoFile862`:
	return intoFile862
}

func TaintStepTest_NetUnixListenerFile_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile619` into `intoUnixListener835`.

	// Assume that `sourceCQL` has the underlying type of `fromFile619`:
	fromFile619 := sourceCQL.(*os.File)

	// Declare `intoUnixListener835` variable:
	var intoUnixListener835 net.UnixListener

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoUnixListener835`:
	intermediateCQL, _ := intoUnixListener835.File()

	// Extra step (`fromFile619` taints `intermediateCQL`, which taints `intoUnixListener835`:
	link(fromFile619, intermediateCQL)

	// Return the tainted `intoUnixListener835`:
	return intoUnixListener835
}

func TaintStepTest_NetUnixListenerSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixListener344` into `intoRawConn790`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixListener344`:
	fromUnixListener344 := sourceCQL.(net.UnixListener)

	// Call the method that transfers the taint
	// from the receiver `fromUnixListener344` to the result `intoRawConn790`
	// (`intoRawConn790` is now tainted).
	intoRawConn790, _ := fromUnixListener344.SyscallConn()

	// Return the tainted `intoRawConn790`:
	return intoRawConn790
}

func TaintStepTest_NetUnixListenerSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn816` into `intoUnixListener657`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn816`:
	fromRawConn816 := sourceCQL.(syscall.RawConn)

	// Declare `intoUnixListener657` variable:
	var intoUnixListener657 net.UnixListener

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoUnixListener657`:
	intermediateCQL, _ := intoUnixListener657.SyscallConn()

	// Extra step (`fromRawConn816` taints `intermediateCQL`, which taints `intoUnixListener657`:
	link(fromRawConn816, intermediateCQL)

	// Return the tainted `intoUnixListener657`:
	return intoUnixListener657
}

func TaintStepTest_NetConnRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn721` into `intoByte425`.

	// Assume that `sourceCQL` has the underlying type of `fromConn721`:
	fromConn721 := sourceCQL.(net.Conn)

	// Declare `intoByte425` variable:
	var intoByte425 []byte

	// Call the method that transfers the taint
	// from the receiver `fromConn721` to the argument `intoByte425`
	// (`intoByte425` is now tainted).
	fromConn721.Read(intoByte425)

	// Return the tainted `intoByte425`:
	return intoByte425
}

func TaintStepTest_NetPacketConnReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPacketConn454` into `intoByte846`.

	// Assume that `sourceCQL` has the underlying type of `fromPacketConn454`:
	fromPacketConn454 := sourceCQL.(net.PacketConn)

	// Declare `intoByte846` variable:
	var intoByte846 []byte

	// Call the method that transfers the taint
	// from the receiver `fromPacketConn454` to the argument `intoByte846`
	// (`intoByte846` is now tainted).
	fromPacketConn454.ReadFrom(intoByte846)

	// Return the tainted `intoByte846`:
	return intoByte846
}

func TaintStepTest_NetAddrString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromAddr289` into `intoString225`.

	// Assume that `sourceCQL` has the underlying type of `fromAddr289`:
	fromAddr289 := sourceCQL.(net.Addr)

	// Call the method that transfers the taint
	// from the receiver `fromAddr289` to the result `intoString225`
	// (`intoString225` is now tainted).
	intoString225 := fromAddr289.String()

	// Return the tainted `intoString225`:
	return intoString225
}

func TaintStepTest_NetConnWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte981` into `intoConn630`.

	// Assume that `sourceCQL` has the underlying type of `fromByte981`:
	fromByte981 := sourceCQL.([]byte)

	// Declare `intoConn630` variable:
	var intoConn630 net.Conn

	// Call the method that transfers the taint
	// from the parameter `fromByte981` to the receiver `intoConn630`
	// (`intoConn630` is now tainted).
	intoConn630.Write(fromByte981)

	// Return the tainted `intoConn630`:
	return intoConn630
}

func TaintStepTest_NetPacketConnWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte789` into `intoPacketConn567`.

	// Assume that `sourceCQL` has the underlying type of `fromByte789`:
	fromByte789 := sourceCQL.([]byte)

	// Declare `intoPacketConn567` variable:
	var intoPacketConn567 net.PacketConn

	// Call the method that transfers the taint
	// from the parameter `fromByte789` to the receiver `intoPacketConn567`
	// (`intoPacketConn567` is now tainted).
	intoPacketConn567.WriteTo(fromByte789, nil)

	// Return the tainted `intoPacketConn567`:
	return intoPacketConn567
}

func RunAllTaints_Net() {
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetFileConn_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetFileConn_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetFileListener_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetFileListener_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetFilePacketConn_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetFilePacketConn_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPv4_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPv4_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPv4_B0I2O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPv4_B0I3O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPv4Mask_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPv4Mask_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPv4Mask_B0I2O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPv4Mask_B0I3O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetJoinHostPort_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetJoinHostPort_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetParseCIDR_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetParseCIDR_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetParseIP_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetParseMAC_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetPipe_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetPipe_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetSplitHostPort_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetSplitHostPort_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetBuffersRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetBuffersWriteTo_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetFlagsString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetHardwareAddrString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPMarshalText_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPTo16_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPTo4_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPUnmarshalText_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPAddrString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPConnReadFrom_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPConnReadFromIP_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPConnReadMsgIP_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPConnReadMsgIP_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPConnSyscallConn_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPConnSyscallConn_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPConnWriteMsgIP_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPConnWriteMsgIP_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPConnWriteTo_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPConnWriteToIP_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPMaskString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetIPNetString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTCPAddrString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTCPConnReadFrom_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTCPConnSyscallConn_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTCPConnSyscallConn_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTCPListenerFile_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTCPListenerFile_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTCPListenerSyscallConn_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetTCPListenerSyscallConn_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUDPConnReadFrom_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUDPConnReadFromUDP_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUDPConnReadMsgUDP_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUDPConnReadMsgUDP_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUDPConnSyscallConn_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUDPConnSyscallConn_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUDPConnWriteMsgUDP_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUDPConnWriteMsgUDP_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUDPConnWriteTo_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUDPConnWriteToUDP_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUnixAddrString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUnixConnReadFrom_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUnixConnReadFromUnix_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUnixConnReadMsgUnix_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUnixConnReadMsgUnix_B0I0O1(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUnixConnSyscallConn_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUnixConnSyscallConn_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUnixConnWriteMsgUnix_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUnixConnWriteMsgUnix_B0I1O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUnixConnWriteTo_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUnixConnWriteToUnix_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUnixListenerFile_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUnixListenerFile_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUnixListenerSyscallConn_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetUnixListenerSyscallConn_B1I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetConnRead_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetPacketConnReadFrom_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetAddrString_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetConnWrite_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
	{
		// Create a new source:
		source := newSource()
		// Run the taint scenario:
		out := TaintStepTest_NetPacketConnWriteTo_B0I0O0(source)
		// If the taint step(s) succeeded, then `out` is tainted and will be sink-able here:
		sink(out)
	}
}
