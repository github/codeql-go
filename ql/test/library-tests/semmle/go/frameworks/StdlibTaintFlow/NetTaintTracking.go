package main

import (
	"io"
	"net"
	"os"
	"syscall"
)

func TaintStepTest_NetFileConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile403` into `intoConn756`.

	// Assume that `sourceCQL` has the underlying type of `fromFile403`:
	fromFile403 := sourceCQL.(*os.File)

	// Call the function that transfers the taint
	// from the parameter `fromFile403` to result `intoConn756`
	// (`intoConn756` is now tainted).
	intoConn756, _ := net.FileConn(fromFile403)

	// Return the tainted `intoConn756`:
	return intoConn756
}

func TaintStepTest_NetFileConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn223` into `intoFile910`.

	// Assume that `sourceCQL` has the underlying type of `fromConn223`:
	fromConn223 := sourceCQL.(net.Conn)

	// Declare `intoFile910` variable:
	var intoFile910 *os.File

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoFile910`:
	intermediateCQL, _ := net.FileConn(intoFile910)

	// Extra step (`fromConn223` taints `intermediateCQL`, which taints `intoFile910`:
	link(fromConn223, intermediateCQL)

	// Return the tainted `intoFile910`:
	return intoFile910
}

func TaintStepTest_NetFileListener_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile243` into `intoListener944`.

	// Assume that `sourceCQL` has the underlying type of `fromFile243`:
	fromFile243 := sourceCQL.(*os.File)

	// Call the function that transfers the taint
	// from the parameter `fromFile243` to result `intoListener944`
	// (`intoListener944` is now tainted).
	intoListener944, _ := net.FileListener(fromFile243)

	// Return the tainted `intoListener944`:
	return intoListener944
}

func TaintStepTest_NetFileListener_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromListener815` into `intoFile322`.

	// Assume that `sourceCQL` has the underlying type of `fromListener815`:
	fromListener815 := sourceCQL.(net.Listener)

	// Declare `intoFile322` variable:
	var intoFile322 *os.File

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoFile322`:
	intermediateCQL, _ := net.FileListener(intoFile322)

	// Extra step (`fromListener815` taints `intermediateCQL`, which taints `intoFile322`:
	link(fromListener815, intermediateCQL)

	// Return the tainted `intoFile322`:
	return intoFile322
}

func TaintStepTest_NetFilePacketConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile525` into `intoPacketConn651`.

	// Assume that `sourceCQL` has the underlying type of `fromFile525`:
	fromFile525 := sourceCQL.(*os.File)

	// Call the function that transfers the taint
	// from the parameter `fromFile525` to result `intoPacketConn651`
	// (`intoPacketConn651` is now tainted).
	intoPacketConn651, _ := net.FilePacketConn(fromFile525)

	// Return the tainted `intoPacketConn651`:
	return intoPacketConn651
}

func TaintStepTest_NetFilePacketConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPacketConn612` into `intoFile144`.

	// Assume that `sourceCQL` has the underlying type of `fromPacketConn612`:
	fromPacketConn612 := sourceCQL.(net.PacketConn)

	// Declare `intoFile144` variable:
	var intoFile144 *os.File

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoFile144`:
	intermediateCQL, _ := net.FilePacketConn(intoFile144)

	// Extra step (`fromPacketConn612` taints `intermediateCQL`, which taints `intoFile144`:
	link(fromPacketConn612, intermediateCQL)

	// Return the tainted `intoFile144`:
	return intoFile144
}

func TaintStepTest_NetIPv4_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte742` into `intoIP981`.

	// Assume that `sourceCQL` has the underlying type of `fromByte742`:
	fromByte742 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte742` to result `intoIP981`
	// (`intoIP981` is now tainted).
	intoIP981 := net.IPv4(fromByte742, 0, 0, 0)

	// Return the tainted `intoIP981`:
	return intoIP981
}

func TaintStepTest_NetIPv4_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte258` into `intoIP991`.

	// Assume that `sourceCQL` has the underlying type of `fromByte258`:
	fromByte258 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte258` to result `intoIP991`
	// (`intoIP991` is now tainted).
	intoIP991 := net.IPv4(0, fromByte258, 0, 0)

	// Return the tainted `intoIP991`:
	return intoIP991
}

func TaintStepTest_NetIPv4_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte692` into `intoIP227`.

	// Assume that `sourceCQL` has the underlying type of `fromByte692`:
	fromByte692 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte692` to result `intoIP227`
	// (`intoIP227` is now tainted).
	intoIP227 := net.IPv4(0, 0, fromByte692, 0)

	// Return the tainted `intoIP227`:
	return intoIP227
}

func TaintStepTest_NetIPv4_B0I3O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte921` into `intoIP255`.

	// Assume that `sourceCQL` has the underlying type of `fromByte921`:
	fromByte921 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte921` to result `intoIP255`
	// (`intoIP255` is now tainted).
	intoIP255 := net.IPv4(0, 0, 0, fromByte921)

	// Return the tainted `intoIP255`:
	return intoIP255
}

func TaintStepTest_NetIPv4Mask_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte496` into `intoIPMask146`.

	// Assume that `sourceCQL` has the underlying type of `fromByte496`:
	fromByte496 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte496` to result `intoIPMask146`
	// (`intoIPMask146` is now tainted).
	intoIPMask146 := net.IPv4Mask(fromByte496, 0, 0, 0)

	// Return the tainted `intoIPMask146`:
	return intoIPMask146
}

func TaintStepTest_NetIPv4Mask_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte703` into `intoIPMask559`.

	// Assume that `sourceCQL` has the underlying type of `fromByte703`:
	fromByte703 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte703` to result `intoIPMask559`
	// (`intoIPMask559` is now tainted).
	intoIPMask559 := net.IPv4Mask(0, fromByte703, 0, 0)

	// Return the tainted `intoIPMask559`:
	return intoIPMask559
}

func TaintStepTest_NetIPv4Mask_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte178` into `intoIPMask758`.

	// Assume that `sourceCQL` has the underlying type of `fromByte178`:
	fromByte178 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte178` to result `intoIPMask758`
	// (`intoIPMask758` is now tainted).
	intoIPMask758 := net.IPv4Mask(0, 0, fromByte178, 0)

	// Return the tainted `intoIPMask758`:
	return intoIPMask758
}

func TaintStepTest_NetIPv4Mask_B0I3O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte230` into `intoIPMask188`.

	// Assume that `sourceCQL` has the underlying type of `fromByte230`:
	fromByte230 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte230` to result `intoIPMask188`
	// (`intoIPMask188` is now tainted).
	intoIPMask188 := net.IPv4Mask(0, 0, 0, fromByte230)

	// Return the tainted `intoIPMask188`:
	return intoIPMask188
}

func TaintStepTest_NetJoinHostPort_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString480` into `intoString801`.

	// Assume that `sourceCQL` has the underlying type of `fromString480`:
	fromString480 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString480` to result `intoString801`
	// (`intoString801` is now tainted).
	intoString801 := net.JoinHostPort(fromString480, "")

	// Return the tainted `intoString801`:
	return intoString801
}

func TaintStepTest_NetJoinHostPort_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString558` into `intoString626`.

	// Assume that `sourceCQL` has the underlying type of `fromString558`:
	fromString558 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString558` to result `intoString626`
	// (`intoString626` is now tainted).
	intoString626 := net.JoinHostPort("", fromString558)

	// Return the tainted `intoString626`:
	return intoString626
}

func TaintStepTest_NetParseCIDR_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString758` into `intoIP634`.

	// Assume that `sourceCQL` has the underlying type of `fromString758`:
	fromString758 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString758` to result `intoIP634`
	// (`intoIP634` is now tainted).
	intoIP634, _, _ := net.ParseCIDR(fromString758)

	// Return the tainted `intoIP634`:
	return intoIP634
}

func TaintStepTest_NetParseCIDR_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString576` into `intoIPNet730`.

	// Assume that `sourceCQL` has the underlying type of `fromString576`:
	fromString576 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString576` to result `intoIPNet730`
	// (`intoIPNet730` is now tainted).
	_, intoIPNet730, _ := net.ParseCIDR(fromString576)

	// Return the tainted `intoIPNet730`:
	return intoIPNet730
}

func TaintStepTest_NetParseIP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString320` into `intoIP587`.

	// Assume that `sourceCQL` has the underlying type of `fromString320`:
	fromString320 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString320` to result `intoIP587`
	// (`intoIP587` is now tainted).
	intoIP587 := net.ParseIP(fromString320)

	// Return the tainted `intoIP587`:
	return intoIP587
}

func TaintStepTest_NetParseMAC_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString497` into `intoHardwareAddr405`.

	// Assume that `sourceCQL` has the underlying type of `fromString497`:
	fromString497 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString497` to result `intoHardwareAddr405`
	// (`intoHardwareAddr405` is now tainted).
	intoHardwareAddr405, _ := net.ParseMAC(fromString497)

	// Return the tainted `intoHardwareAddr405`:
	return intoHardwareAddr405
}

func TaintStepTest_NetPipe_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn497` into `intoConn549`.

	// Assume that `sourceCQL` has the underlying type of `fromConn497`:
	fromConn497 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the result `fromConn497` to result `intoConn549`
	// (extra steps needed)
	intermediateCQL, intoConn549 := net.Pipe()

	// Extra step (`fromConn497` taints `intermediateCQL`, which taints `intoConn549`:
	link(fromConn497, intermediateCQL)

	// Return the tainted `intoConn549`:
	return intoConn549
}

func TaintStepTest_NetPipe_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn195` into `intoConn231`.

	// Assume that `sourceCQL` has the underlying type of `fromConn195`:
	fromConn195 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the result `fromConn195` to result `intoConn231`
	// (extra steps needed)
	intoConn231, intermediateCQL := net.Pipe()

	// Extra step (`fromConn195` taints `intermediateCQL`, which taints `intoConn231`:
	link(fromConn195, intermediateCQL)

	// Return the tainted `intoConn231`:
	return intoConn231
}

func TaintStepTest_NetSplitHostPort_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString598` into `intoString420`.

	// Assume that `sourceCQL` has the underlying type of `fromString598`:
	fromString598 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString598` to result `intoString420`
	// (`intoString420` is now tainted).
	intoString420, _, _ := net.SplitHostPort(fromString598)

	// Return the tainted `intoString420`:
	return intoString420
}

func TaintStepTest_NetSplitHostPort_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString846` into `intoString195`.

	// Assume that `sourceCQL` has the underlying type of `fromString846`:
	fromString846 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString846` to result `intoString195`
	// (`intoString195` is now tainted).
	_, intoString195, _ := net.SplitHostPort(fromString846)

	// Return the tainted `intoString195`:
	return intoString195
}

func TaintStepTest_NetBuffersRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffers715` into `intoByte553`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffers715`:
	fromBuffers715 := sourceCQL.(net.Buffers)

	// Declare `intoByte553` variable:
	var intoByte553 []byte

	// Call the method that transfers the taint
	// from the receiver `fromBuffers715` to the argument `intoByte553`
	// (`intoByte553` is now tainted).
	fromBuffers715.Read(intoByte553)

	// Return the tainted `intoByte553`:
	return intoByte553
}

func TaintStepTest_NetBuffersWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffers179` into `intoWriter795`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffers179`:
	fromBuffers179 := sourceCQL.(net.Buffers)

	// Declare `intoWriter795` variable:
	var intoWriter795 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromBuffers179` to the argument `intoWriter795`
	// (`intoWriter795` is now tainted).
	fromBuffers179.WriteTo(intoWriter795)

	// Return the tainted `intoWriter795`:
	return intoWriter795
}

func TaintStepTest_NetFlagsString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFlags743` into `intoString159`.

	// Assume that `sourceCQL` has the underlying type of `fromFlags743`:
	fromFlags743 := sourceCQL.(net.Flags)

	// Call the method that transfers the taint
	// from the receiver `fromFlags743` to the result `intoString159`
	// (`intoString159` is now tainted).
	intoString159 := fromFlags743.String()

	// Return the tainted `intoString159`:
	return intoString159
}

func TaintStepTest_NetHardwareAddrString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHardwareAddr472` into `intoString550`.

	// Assume that `sourceCQL` has the underlying type of `fromHardwareAddr472`:
	fromHardwareAddr472 := sourceCQL.(net.HardwareAddr)

	// Call the method that transfers the taint
	// from the receiver `fromHardwareAddr472` to the result `intoString550`
	// (`intoString550` is now tainted).
	intoString550 := fromHardwareAddr472.String()

	// Return the tainted `intoString550`:
	return intoString550
}

func TaintStepTest_NetIPMarshalText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIP450` into `intoByte546`.

	// Assume that `sourceCQL` has the underlying type of `fromIP450`:
	fromIP450 := sourceCQL.(net.IP)

	// Call the method that transfers the taint
	// from the receiver `fromIP450` to the result `intoByte546`
	// (`intoByte546` is now tainted).
	intoByte546, _ := fromIP450.MarshalText()

	// Return the tainted `intoByte546`:
	return intoByte546
}

func TaintStepTest_NetIPString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIP124` into `intoString925`.

	// Assume that `sourceCQL` has the underlying type of `fromIP124`:
	fromIP124 := sourceCQL.(net.IP)

	// Call the method that transfers the taint
	// from the receiver `fromIP124` to the result `intoString925`
	// (`intoString925` is now tainted).
	intoString925 := fromIP124.String()

	// Return the tainted `intoString925`:
	return intoString925
}

func TaintStepTest_NetIPTo16_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIP227` into `intoIP616`.

	// Assume that `sourceCQL` has the underlying type of `fromIP227`:
	fromIP227 := sourceCQL.(net.IP)

	// Call the method that transfers the taint
	// from the receiver `fromIP227` to the result `intoIP616`
	// (`intoIP616` is now tainted).
	intoIP616 := fromIP227.To16()

	// Return the tainted `intoIP616`:
	return intoIP616
}

func TaintStepTest_NetIPTo4_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIP264` into `intoIP157`.

	// Assume that `sourceCQL` has the underlying type of `fromIP264`:
	fromIP264 := sourceCQL.(net.IP)

	// Call the method that transfers the taint
	// from the receiver `fromIP264` to the result `intoIP157`
	// (`intoIP157` is now tainted).
	intoIP157 := fromIP264.To4()

	// Return the tainted `intoIP157`:
	return intoIP157
}

func TaintStepTest_NetIPUnmarshalText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte437` into `intoIP434`.

	// Assume that `sourceCQL` has the underlying type of `fromByte437`:
	fromByte437 := sourceCQL.([]byte)

	// Declare `intoIP434` variable:
	var intoIP434 net.IP

	// Call the method that transfers the taint
	// from the parameter `fromByte437` to the receiver `intoIP434`
	// (`intoIP434` is now tainted).
	intoIP434.UnmarshalText(fromByte437)

	// Return the tainted `intoIP434`:
	return intoIP434
}

func TaintStepTest_NetIPAddrString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPAddr357` into `intoString961`.

	// Assume that `sourceCQL` has the underlying type of `fromIPAddr357`:
	fromIPAddr357 := sourceCQL.(net.IPAddr)

	// Call the method that transfers the taint
	// from the receiver `fromIPAddr357` to the result `intoString961`
	// (`intoString961` is now tainted).
	intoString961 := fromIPAddr357.String()

	// Return the tainted `intoString961`:
	return intoString961
}

func TaintStepTest_NetIPConnReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPConn709` into `intoByte536`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn709`:
	fromIPConn709 := sourceCQL.(net.IPConn)

	// Declare `intoByte536` variable:
	var intoByte536 []byte

	// Call the method that transfers the taint
	// from the receiver `fromIPConn709` to the argument `intoByte536`
	// (`intoByte536` is now tainted).
	fromIPConn709.ReadFrom(intoByte536)

	// Return the tainted `intoByte536`:
	return intoByte536
}

func TaintStepTest_NetIPConnReadFromIP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPConn255` into `intoByte778`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn255`:
	fromIPConn255 := sourceCQL.(net.IPConn)

	// Declare `intoByte778` variable:
	var intoByte778 []byte

	// Call the method that transfers the taint
	// from the receiver `fromIPConn255` to the argument `intoByte778`
	// (`intoByte778` is now tainted).
	fromIPConn255.ReadFromIP(intoByte778)

	// Return the tainted `intoByte778`:
	return intoByte778
}

func TaintStepTest_NetIPConnReadMsgIP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPConn618` into `intoByte377`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn618`:
	fromIPConn618 := sourceCQL.(net.IPConn)

	// Declare `intoByte377` variable:
	var intoByte377 []byte

	// Call the method that transfers the taint
	// from the receiver `fromIPConn618` to the argument `intoByte377`
	// (`intoByte377` is now tainted).
	fromIPConn618.ReadMsgIP(intoByte377, nil)

	// Return the tainted `intoByte377`:
	return intoByte377
}

func TaintStepTest_NetIPConnReadMsgIP_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPConn722` into `intoByte420`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn722`:
	fromIPConn722 := sourceCQL.(net.IPConn)

	// Declare `intoByte420` variable:
	var intoByte420 []byte

	// Call the method that transfers the taint
	// from the receiver `fromIPConn722` to the argument `intoByte420`
	// (`intoByte420` is now tainted).
	fromIPConn722.ReadMsgIP(nil, intoByte420)

	// Return the tainted `intoByte420`:
	return intoByte420
}

func TaintStepTest_NetIPConnSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPConn906` into `intoRawConn980`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn906`:
	fromIPConn906 := sourceCQL.(net.IPConn)

	// Call the method that transfers the taint
	// from the receiver `fromIPConn906` to the result `intoRawConn980`
	// (`intoRawConn980` is now tainted).
	intoRawConn980, _ := fromIPConn906.SyscallConn()

	// Return the tainted `intoRawConn980`:
	return intoRawConn980
}

func TaintStepTest_NetIPConnSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn247` into `intoIPConn436`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn247`:
	fromRawConn247 := sourceCQL.(syscall.RawConn)

	// Declare `intoIPConn436` variable:
	var intoIPConn436 net.IPConn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoIPConn436`:
	intermediateCQL, _ := intoIPConn436.SyscallConn()

	// Extra step (`fromRawConn247` taints `intermediateCQL`, which taints `intoIPConn436`:
	link(fromRawConn247, intermediateCQL)

	// Return the tainted `intoIPConn436`:
	return intoIPConn436
}

func TaintStepTest_NetIPConnWriteMsgIP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte245` into `intoIPConn423`.

	// Assume that `sourceCQL` has the underlying type of `fromByte245`:
	fromByte245 := sourceCQL.([]byte)

	// Declare `intoIPConn423` variable:
	var intoIPConn423 net.IPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte245` to the receiver `intoIPConn423`
	// (`intoIPConn423` is now tainted).
	intoIPConn423.WriteMsgIP(fromByte245, nil, nil)

	// Return the tainted `intoIPConn423`:
	return intoIPConn423
}

func TaintStepTest_NetIPConnWriteMsgIP_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte211` into `intoIPConn551`.

	// Assume that `sourceCQL` has the underlying type of `fromByte211`:
	fromByte211 := sourceCQL.([]byte)

	// Declare `intoIPConn551` variable:
	var intoIPConn551 net.IPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte211` to the receiver `intoIPConn551`
	// (`intoIPConn551` is now tainted).
	intoIPConn551.WriteMsgIP(nil, fromByte211, nil)

	// Return the tainted `intoIPConn551`:
	return intoIPConn551
}

func TaintStepTest_NetIPConnWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte734` into `intoIPConn298`.

	// Assume that `sourceCQL` has the underlying type of `fromByte734`:
	fromByte734 := sourceCQL.([]byte)

	// Declare `intoIPConn298` variable:
	var intoIPConn298 net.IPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte734` to the receiver `intoIPConn298`
	// (`intoIPConn298` is now tainted).
	intoIPConn298.WriteTo(fromByte734, nil)

	// Return the tainted `intoIPConn298`:
	return intoIPConn298
}

func TaintStepTest_NetIPConnWriteToIP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte153` into `intoIPConn267`.

	// Assume that `sourceCQL` has the underlying type of `fromByte153`:
	fromByte153 := sourceCQL.([]byte)

	// Declare `intoIPConn267` variable:
	var intoIPConn267 net.IPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte153` to the receiver `intoIPConn267`
	// (`intoIPConn267` is now tainted).
	intoIPConn267.WriteToIP(fromByte153, nil)

	// Return the tainted `intoIPConn267`:
	return intoIPConn267
}

func TaintStepTest_NetIPMaskString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPMask508` into `intoString532`.

	// Assume that `sourceCQL` has the underlying type of `fromIPMask508`:
	fromIPMask508 := sourceCQL.(net.IPMask)

	// Call the method that transfers the taint
	// from the receiver `fromIPMask508` to the result `intoString532`
	// (`intoString532` is now tainted).
	intoString532 := fromIPMask508.String()

	// Return the tainted `intoString532`:
	return intoString532
}

func TaintStepTest_NetIPNetString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPNet561` into `intoString470`.

	// Assume that `sourceCQL` has the underlying type of `fromIPNet561`:
	fromIPNet561 := sourceCQL.(net.IPNet)

	// Call the method that transfers the taint
	// from the receiver `fromIPNet561` to the result `intoString470`
	// (`intoString470` is now tainted).
	intoString470 := fromIPNet561.String()

	// Return the tainted `intoString470`:
	return intoString470
}

func TaintStepTest_NetTCPAddrString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTCPAddr263` into `intoString373`.

	// Assume that `sourceCQL` has the underlying type of `fromTCPAddr263`:
	fromTCPAddr263 := sourceCQL.(net.TCPAddr)

	// Call the method that transfers the taint
	// from the receiver `fromTCPAddr263` to the result `intoString373`
	// (`intoString373` is now tainted).
	intoString373 := fromTCPAddr263.String()

	// Return the tainted `intoString373`:
	return intoString373
}

func TaintStepTest_NetTCPConnReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader640` into `intoTCPConn322`.

	// Assume that `sourceCQL` has the underlying type of `fromReader640`:
	fromReader640 := sourceCQL.(io.Reader)

	// Declare `intoTCPConn322` variable:
	var intoTCPConn322 net.TCPConn

	// Call the method that transfers the taint
	// from the parameter `fromReader640` to the receiver `intoTCPConn322`
	// (`intoTCPConn322` is now tainted).
	intoTCPConn322.ReadFrom(fromReader640)

	// Return the tainted `intoTCPConn322`:
	return intoTCPConn322
}

func TaintStepTest_NetTCPConnSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTCPConn987` into `intoRawConn593`.

	// Assume that `sourceCQL` has the underlying type of `fromTCPConn987`:
	fromTCPConn987 := sourceCQL.(net.TCPConn)

	// Call the method that transfers the taint
	// from the receiver `fromTCPConn987` to the result `intoRawConn593`
	// (`intoRawConn593` is now tainted).
	intoRawConn593, _ := fromTCPConn987.SyscallConn()

	// Return the tainted `intoRawConn593`:
	return intoRawConn593
}

func TaintStepTest_NetTCPConnSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn799` into `intoTCPConn712`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn799`:
	fromRawConn799 := sourceCQL.(syscall.RawConn)

	// Declare `intoTCPConn712` variable:
	var intoTCPConn712 net.TCPConn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoTCPConn712`:
	intermediateCQL, _ := intoTCPConn712.SyscallConn()

	// Extra step (`fromRawConn799` taints `intermediateCQL`, which taints `intoTCPConn712`:
	link(fromRawConn799, intermediateCQL)

	// Return the tainted `intoTCPConn712`:
	return intoTCPConn712
}

func TaintStepTest_NetTCPListenerFile_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTCPListener974` into `intoFile944`.

	// Assume that `sourceCQL` has the underlying type of `fromTCPListener974`:
	fromTCPListener974 := sourceCQL.(net.TCPListener)

	// Call the method that transfers the taint
	// from the receiver `fromTCPListener974` to the result `intoFile944`
	// (`intoFile944` is now tainted).
	intoFile944, _ := fromTCPListener974.File()

	// Return the tainted `intoFile944`:
	return intoFile944
}

func TaintStepTest_NetTCPListenerFile_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile943` into `intoTCPListener523`.

	// Assume that `sourceCQL` has the underlying type of `fromFile943`:
	fromFile943 := sourceCQL.(*os.File)

	// Declare `intoTCPListener523` variable:
	var intoTCPListener523 net.TCPListener

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoTCPListener523`:
	intermediateCQL, _ := intoTCPListener523.File()

	// Extra step (`fromFile943` taints `intermediateCQL`, which taints `intoTCPListener523`:
	link(fromFile943, intermediateCQL)

	// Return the tainted `intoTCPListener523`:
	return intoTCPListener523
}

func TaintStepTest_NetTCPListenerSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTCPListener760` into `intoRawConn207`.

	// Assume that `sourceCQL` has the underlying type of `fromTCPListener760`:
	fromTCPListener760 := sourceCQL.(net.TCPListener)

	// Call the method that transfers the taint
	// from the receiver `fromTCPListener760` to the result `intoRawConn207`
	// (`intoRawConn207` is now tainted).
	intoRawConn207, _ := fromTCPListener760.SyscallConn()

	// Return the tainted `intoRawConn207`:
	return intoRawConn207
}

func TaintStepTest_NetTCPListenerSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn510` into `intoTCPListener867`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn510`:
	fromRawConn510 := sourceCQL.(syscall.RawConn)

	// Declare `intoTCPListener867` variable:
	var intoTCPListener867 net.TCPListener

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoTCPListener867`:
	intermediateCQL, _ := intoTCPListener867.SyscallConn()

	// Extra step (`fromRawConn510` taints `intermediateCQL`, which taints `intoTCPListener867`:
	link(fromRawConn510, intermediateCQL)

	// Return the tainted `intoTCPListener867`:
	return intoTCPListener867
}

func TaintStepTest_NetUDPConnReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUDPConn293` into `intoByte449`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn293`:
	fromUDPConn293 := sourceCQL.(net.UDPConn)

	// Declare `intoByte449` variable:
	var intoByte449 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn293` to the argument `intoByte449`
	// (`intoByte449` is now tainted).
	fromUDPConn293.ReadFrom(intoByte449)

	// Return the tainted `intoByte449`:
	return intoByte449
}

func TaintStepTest_NetUDPConnReadFromUDP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUDPConn479` into `intoByte596`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn479`:
	fromUDPConn479 := sourceCQL.(net.UDPConn)

	// Declare `intoByte596` variable:
	var intoByte596 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn479` to the argument `intoByte596`
	// (`intoByte596` is now tainted).
	fromUDPConn479.ReadFromUDP(intoByte596)

	// Return the tainted `intoByte596`:
	return intoByte596
}

func TaintStepTest_NetUDPConnReadMsgUDP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUDPConn506` into `intoByte321`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn506`:
	fromUDPConn506 := sourceCQL.(net.UDPConn)

	// Declare `intoByte321` variable:
	var intoByte321 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn506` to the argument `intoByte321`
	// (`intoByte321` is now tainted).
	fromUDPConn506.ReadMsgUDP(intoByte321, nil)

	// Return the tainted `intoByte321`:
	return intoByte321
}

func TaintStepTest_NetUDPConnReadMsgUDP_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromUDPConn199` into `intoByte316`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn199`:
	fromUDPConn199 := sourceCQL.(net.UDPConn)

	// Declare `intoByte316` variable:
	var intoByte316 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn199` to the argument `intoByte316`
	// (`intoByte316` is now tainted).
	fromUDPConn199.ReadMsgUDP(nil, intoByte316)

	// Return the tainted `intoByte316`:
	return intoByte316
}

func TaintStepTest_NetUDPConnSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUDPConn895` into `intoRawConn744`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn895`:
	fromUDPConn895 := sourceCQL.(net.UDPConn)

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn895` to the result `intoRawConn744`
	// (`intoRawConn744` is now tainted).
	intoRawConn744, _ := fromUDPConn895.SyscallConn()

	// Return the tainted `intoRawConn744`:
	return intoRawConn744
}

func TaintStepTest_NetUDPConnSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn617` into `intoUDPConn446`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn617`:
	fromRawConn617 := sourceCQL.(syscall.RawConn)

	// Declare `intoUDPConn446` variable:
	var intoUDPConn446 net.UDPConn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoUDPConn446`:
	intermediateCQL, _ := intoUDPConn446.SyscallConn()

	// Extra step (`fromRawConn617` taints `intermediateCQL`, which taints `intoUDPConn446`:
	link(fromRawConn617, intermediateCQL)

	// Return the tainted `intoUDPConn446`:
	return intoUDPConn446
}

func TaintStepTest_NetUDPConnWriteMsgUDP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte112` into `intoUDPConn490`.

	// Assume that `sourceCQL` has the underlying type of `fromByte112`:
	fromByte112 := sourceCQL.([]byte)

	// Declare `intoUDPConn490` variable:
	var intoUDPConn490 net.UDPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte112` to the receiver `intoUDPConn490`
	// (`intoUDPConn490` is now tainted).
	intoUDPConn490.WriteMsgUDP(fromByte112, nil, nil)

	// Return the tainted `intoUDPConn490`:
	return intoUDPConn490
}

func TaintStepTest_NetUDPConnWriteMsgUDP_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte382` into `intoUDPConn925`.

	// Assume that `sourceCQL` has the underlying type of `fromByte382`:
	fromByte382 := sourceCQL.([]byte)

	// Declare `intoUDPConn925` variable:
	var intoUDPConn925 net.UDPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte382` to the receiver `intoUDPConn925`
	// (`intoUDPConn925` is now tainted).
	intoUDPConn925.WriteMsgUDP(nil, fromByte382, nil)

	// Return the tainted `intoUDPConn925`:
	return intoUDPConn925
}

func TaintStepTest_NetUDPConnWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte272` into `intoUDPConn842`.

	// Assume that `sourceCQL` has the underlying type of `fromByte272`:
	fromByte272 := sourceCQL.([]byte)

	// Declare `intoUDPConn842` variable:
	var intoUDPConn842 net.UDPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte272` to the receiver `intoUDPConn842`
	// (`intoUDPConn842` is now tainted).
	intoUDPConn842.WriteTo(fromByte272, nil)

	// Return the tainted `intoUDPConn842`:
	return intoUDPConn842
}

func TaintStepTest_NetUDPConnWriteToUDP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte421` into `intoUDPConn233`.

	// Assume that `sourceCQL` has the underlying type of `fromByte421`:
	fromByte421 := sourceCQL.([]byte)

	// Declare `intoUDPConn233` variable:
	var intoUDPConn233 net.UDPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte421` to the receiver `intoUDPConn233`
	// (`intoUDPConn233` is now tainted).
	intoUDPConn233.WriteToUDP(fromByte421, nil)

	// Return the tainted `intoUDPConn233`:
	return intoUDPConn233
}

func TaintStepTest_NetUnixAddrString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixAddr649` into `intoString513`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixAddr649`:
	fromUnixAddr649 := sourceCQL.(net.UnixAddr)

	// Call the method that transfers the taint
	// from the receiver `fromUnixAddr649` to the result `intoString513`
	// (`intoString513` is now tainted).
	intoString513 := fromUnixAddr649.String()

	// Return the tainted `intoString513`:
	return intoString513
}

func TaintStepTest_NetUnixConnReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixConn507` into `intoByte560`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn507`:
	fromUnixConn507 := sourceCQL.(net.UnixConn)

	// Declare `intoByte560` variable:
	var intoByte560 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn507` to the argument `intoByte560`
	// (`intoByte560` is now tainted).
	fromUnixConn507.ReadFrom(intoByte560)

	// Return the tainted `intoByte560`:
	return intoByte560
}

func TaintStepTest_NetUnixConnReadFromUnix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixConn423` into `intoByte841`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn423`:
	fromUnixConn423 := sourceCQL.(net.UnixConn)

	// Declare `intoByte841` variable:
	var intoByte841 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn423` to the argument `intoByte841`
	// (`intoByte841` is now tainted).
	fromUnixConn423.ReadFromUnix(intoByte841)

	// Return the tainted `intoByte841`:
	return intoByte841
}

func TaintStepTest_NetUnixConnReadMsgUnix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixConn804` into `intoByte864`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn804`:
	fromUnixConn804 := sourceCQL.(net.UnixConn)

	// Declare `intoByte864` variable:
	var intoByte864 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn804` to the argument `intoByte864`
	// (`intoByte864` is now tainted).
	fromUnixConn804.ReadMsgUnix(intoByte864, nil)

	// Return the tainted `intoByte864`:
	return intoByte864
}

func TaintStepTest_NetUnixConnReadMsgUnix_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixConn971` into `intoByte183`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn971`:
	fromUnixConn971 := sourceCQL.(net.UnixConn)

	// Declare `intoByte183` variable:
	var intoByte183 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn971` to the argument `intoByte183`
	// (`intoByte183` is now tainted).
	fromUnixConn971.ReadMsgUnix(nil, intoByte183)

	// Return the tainted `intoByte183`:
	return intoByte183
}

func TaintStepTest_NetUnixConnSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixConn667` into `intoRawConn897`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn667`:
	fromUnixConn667 := sourceCQL.(net.UnixConn)

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn667` to the result `intoRawConn897`
	// (`intoRawConn897` is now tainted).
	intoRawConn897, _ := fromUnixConn667.SyscallConn()

	// Return the tainted `intoRawConn897`:
	return intoRawConn897
}

func TaintStepTest_NetUnixConnSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn538` into `intoUnixConn128`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn538`:
	fromRawConn538 := sourceCQL.(syscall.RawConn)

	// Declare `intoUnixConn128` variable:
	var intoUnixConn128 net.UnixConn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoUnixConn128`:
	intermediateCQL, _ := intoUnixConn128.SyscallConn()

	// Extra step (`fromRawConn538` taints `intermediateCQL`, which taints `intoUnixConn128`:
	link(fromRawConn538, intermediateCQL)

	// Return the tainted `intoUnixConn128`:
	return intoUnixConn128
}

func TaintStepTest_NetUnixConnWriteMsgUnix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte816` into `intoUnixConn785`.

	// Assume that `sourceCQL` has the underlying type of `fromByte816`:
	fromByte816 := sourceCQL.([]byte)

	// Declare `intoUnixConn785` variable:
	var intoUnixConn785 net.UnixConn

	// Call the method that transfers the taint
	// from the parameter `fromByte816` to the receiver `intoUnixConn785`
	// (`intoUnixConn785` is now tainted).
	intoUnixConn785.WriteMsgUnix(fromByte816, nil, nil)

	// Return the tainted `intoUnixConn785`:
	return intoUnixConn785
}

func TaintStepTest_NetUnixConnWriteMsgUnix_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte120` into `intoUnixConn137`.

	// Assume that `sourceCQL` has the underlying type of `fromByte120`:
	fromByte120 := sourceCQL.([]byte)

	// Declare `intoUnixConn137` variable:
	var intoUnixConn137 net.UnixConn

	// Call the method that transfers the taint
	// from the parameter `fromByte120` to the receiver `intoUnixConn137`
	// (`intoUnixConn137` is now tainted).
	intoUnixConn137.WriteMsgUnix(nil, fromByte120, nil)

	// Return the tainted `intoUnixConn137`:
	return intoUnixConn137
}

func TaintStepTest_NetUnixConnWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte884` into `intoUnixConn903`.

	// Assume that `sourceCQL` has the underlying type of `fromByte884`:
	fromByte884 := sourceCQL.([]byte)

	// Declare `intoUnixConn903` variable:
	var intoUnixConn903 net.UnixConn

	// Call the method that transfers the taint
	// from the parameter `fromByte884` to the receiver `intoUnixConn903`
	// (`intoUnixConn903` is now tainted).
	intoUnixConn903.WriteTo(fromByte884, nil)

	// Return the tainted `intoUnixConn903`:
	return intoUnixConn903
}

func TaintStepTest_NetUnixConnWriteToUnix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte846` into `intoUnixConn607`.

	// Assume that `sourceCQL` has the underlying type of `fromByte846`:
	fromByte846 := sourceCQL.([]byte)

	// Declare `intoUnixConn607` variable:
	var intoUnixConn607 net.UnixConn

	// Call the method that transfers the taint
	// from the parameter `fromByte846` to the receiver `intoUnixConn607`
	// (`intoUnixConn607` is now tainted).
	intoUnixConn607.WriteToUnix(fromByte846, nil)

	// Return the tainted `intoUnixConn607`:
	return intoUnixConn607
}

func TaintStepTest_NetUnixListenerFile_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixListener264` into `intoFile894`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixListener264`:
	fromUnixListener264 := sourceCQL.(net.UnixListener)

	// Call the method that transfers the taint
	// from the receiver `fromUnixListener264` to the result `intoFile894`
	// (`intoFile894` is now tainted).
	intoFile894, _ := fromUnixListener264.File()

	// Return the tainted `intoFile894`:
	return intoFile894
}

func TaintStepTest_NetUnixListenerFile_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile839` into `intoUnixListener291`.

	// Assume that `sourceCQL` has the underlying type of `fromFile839`:
	fromFile839 := sourceCQL.(*os.File)

	// Declare `intoUnixListener291` variable:
	var intoUnixListener291 net.UnixListener

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoUnixListener291`:
	intermediateCQL, _ := intoUnixListener291.File()

	// Extra step (`fromFile839` taints `intermediateCQL`, which taints `intoUnixListener291`:
	link(fromFile839, intermediateCQL)

	// Return the tainted `intoUnixListener291`:
	return intoUnixListener291
}

func TaintStepTest_NetUnixListenerSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixListener306` into `intoRawConn224`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixListener306`:
	fromUnixListener306 := sourceCQL.(net.UnixListener)

	// Call the method that transfers the taint
	// from the receiver `fromUnixListener306` to the result `intoRawConn224`
	// (`intoRawConn224` is now tainted).
	intoRawConn224, _ := fromUnixListener306.SyscallConn()

	// Return the tainted `intoRawConn224`:
	return intoRawConn224
}

func TaintStepTest_NetUnixListenerSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn236` into `intoUnixListener753`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn236`:
	fromRawConn236 := sourceCQL.(syscall.RawConn)

	// Declare `intoUnixListener753` variable:
	var intoUnixListener753 net.UnixListener

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoUnixListener753`:
	intermediateCQL, _ := intoUnixListener753.SyscallConn()

	// Extra step (`fromRawConn236` taints `intermediateCQL`, which taints `intoUnixListener753`:
	link(fromRawConn236, intermediateCQL)

	// Return the tainted `intoUnixListener753`:
	return intoUnixListener753
}

func TaintStepTest_NetConnRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn681` into `intoByte870`.

	// Assume that `sourceCQL` has the underlying type of `fromConn681`:
	fromConn681 := sourceCQL.(net.Conn)

	// Declare `intoByte870` variable:
	var intoByte870 []byte

	// Call the method that transfers the taint
	// from the receiver `fromConn681` to the argument `intoByte870`
	// (`intoByte870` is now tainted).
	fromConn681.Read(intoByte870)

	// Return the tainted `intoByte870`:
	return intoByte870
}

func TaintStepTest_NetPacketConnReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPacketConn944` into `intoByte588`.

	// Assume that `sourceCQL` has the underlying type of `fromPacketConn944`:
	fromPacketConn944 := sourceCQL.(net.PacketConn)

	// Declare `intoByte588` variable:
	var intoByte588 []byte

	// Call the method that transfers the taint
	// from the receiver `fromPacketConn944` to the argument `intoByte588`
	// (`intoByte588` is now tainted).
	fromPacketConn944.ReadFrom(intoByte588)

	// Return the tainted `intoByte588`:
	return intoByte588
}

func TaintStepTest_NetAddrString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromAddr207` into `intoString446`.

	// Assume that `sourceCQL` has the underlying type of `fromAddr207`:
	fromAddr207 := sourceCQL.(net.Addr)

	// Call the method that transfers the taint
	// from the receiver `fromAddr207` to the result `intoString446`
	// (`intoString446` is now tainted).
	intoString446 := fromAddr207.String()

	// Return the tainted `intoString446`:
	return intoString446
}

func TaintStepTest_NetConnWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte714` into `intoConn564`.

	// Assume that `sourceCQL` has the underlying type of `fromByte714`:
	fromByte714 := sourceCQL.([]byte)

	// Declare `intoConn564` variable:
	var intoConn564 net.Conn

	// Call the method that transfers the taint
	// from the parameter `fromByte714` to the receiver `intoConn564`
	// (`intoConn564` is now tainted).
	intoConn564.Write(fromByte714)

	// Return the tainted `intoConn564`:
	return intoConn564
}

func TaintStepTest_NetPacketConnWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte835` into `intoPacketConn591`.

	// Assume that `sourceCQL` has the underlying type of `fromByte835`:
	fromByte835 := sourceCQL.([]byte)

	// Declare `intoPacketConn591` variable:
	var intoPacketConn591 net.PacketConn

	// Call the method that transfers the taint
	// from the parameter `fromByte835` to the receiver `intoPacketConn591`
	// (`intoPacketConn591` is now tainted).
	intoPacketConn591.WriteTo(fromByte835, nil)

	// Return the tainted `intoPacketConn591`:
	return intoPacketConn591
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
