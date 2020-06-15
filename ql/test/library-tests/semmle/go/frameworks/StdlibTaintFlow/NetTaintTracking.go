package main

import (
	"io"
	"net"
	"os"
	"syscall"
)

func TaintStepTest_NetFileConn_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFile273` into `intoConn141`.

	// Assume that `sourceCQL` has the underlying type of `fromFile273`:
	fromFile273 := sourceCQL.(*os.File)

	// Call the function that transfers the taint
	// from the parameter `fromFile273` to result `intoConn141`
	// (`intoConn141` is now tainted).
	intoConn141, _ := net.FileConn(fromFile273)

	// Sink the tainted `intoConn141`:
	sink(intoConn141)
}

func TaintStepTest_NetFileConn_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromConn438` into `intoFile646`.

	// Assume that `sourceCQL` has the underlying type of `fromConn438`:
	fromConn438 := sourceCQL.(net.Conn)

	// Declare `intoFile646` variable:
	var intoFile646 *os.File

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoFile646`:
	intermediateCQL, _ := net.FileConn(intoFile646)

	// Extra step (`fromConn438` taints `intermediateCQL`, which taints `intoFile646`:
	link(fromConn438, intermediateCQL)

	// Sink the tainted `intoFile646`:
	sink(intoFile646)
}

func TaintStepTest_NetFileListener_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFile337` into `intoListener911`.

	// Assume that `sourceCQL` has the underlying type of `fromFile337`:
	fromFile337 := sourceCQL.(*os.File)

	// Call the function that transfers the taint
	// from the parameter `fromFile337` to result `intoListener911`
	// (`intoListener911` is now tainted).
	intoListener911, _ := net.FileListener(fromFile337)

	// Sink the tainted `intoListener911`:
	sink(intoListener911)
}

func TaintStepTest_NetFileListener_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromListener876` into `intoFile386`.

	// Assume that `sourceCQL` has the underlying type of `fromListener876`:
	fromListener876 := sourceCQL.(net.Listener)

	// Declare `intoFile386` variable:
	var intoFile386 *os.File

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoFile386`:
	intermediateCQL, _ := net.FileListener(intoFile386)

	// Extra step (`fromListener876` taints `intermediateCQL`, which taints `intoFile386`:
	link(fromListener876, intermediateCQL)

	// Sink the tainted `intoFile386`:
	sink(intoFile386)
}

func TaintStepTest_NetFilePacketConn_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFile354` into `intoPacketConn977`.

	// Assume that `sourceCQL` has the underlying type of `fromFile354`:
	fromFile354 := sourceCQL.(*os.File)

	// Call the function that transfers the taint
	// from the parameter `fromFile354` to result `intoPacketConn977`
	// (`intoPacketConn977` is now tainted).
	intoPacketConn977, _ := net.FilePacketConn(fromFile354)

	// Sink the tainted `intoPacketConn977`:
	sink(intoPacketConn977)
}

func TaintStepTest_NetFilePacketConn_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromPacketConn246` into `intoFile498`.

	// Assume that `sourceCQL` has the underlying type of `fromPacketConn246`:
	fromPacketConn246 := sourceCQL.(net.PacketConn)

	// Declare `intoFile498` variable:
	var intoFile498 *os.File

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoFile498`:
	intermediateCQL, _ := net.FilePacketConn(intoFile498)

	// Extra step (`fromPacketConn246` taints `intermediateCQL`, which taints `intoFile498`:
	link(fromPacketConn246, intermediateCQL)

	// Sink the tainted `intoFile498`:
	sink(intoFile498)
}

func TaintStepTest_NetIPv4_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte517` into `intoIP272`.

	// Assume that `sourceCQL` has the underlying type of `fromByte517`:
	fromByte517 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte517` to result `intoIP272`
	// (`intoIP272` is now tainted).
	intoIP272 := net.IPv4(fromByte517, 0, 0, 0)

	// Sink the tainted `intoIP272`:
	sink(intoIP272)
}

func TaintStepTest_NetIPv4_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromByte296` into `intoIP671`.

	// Assume that `sourceCQL` has the underlying type of `fromByte296`:
	fromByte296 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte296` to result `intoIP671`
	// (`intoIP671` is now tainted).
	intoIP671 := net.IPv4(0, fromByte296, 0, 0)

	// Sink the tainted `intoIP671`:
	sink(intoIP671)
}

func TaintStepTest_NetIPv4_B0I2O0(sourceCQL interface{}) {
	// The flow is from `fromByte272` into `intoIP152`.

	// Assume that `sourceCQL` has the underlying type of `fromByte272`:
	fromByte272 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte272` to result `intoIP152`
	// (`intoIP152` is now tainted).
	intoIP152 := net.IPv4(0, 0, fromByte272, 0)

	// Sink the tainted `intoIP152`:
	sink(intoIP152)
}

func TaintStepTest_NetIPv4_B0I3O0(sourceCQL interface{}) {
	// The flow is from `fromByte725` into `intoIP524`.

	// Assume that `sourceCQL` has the underlying type of `fromByte725`:
	fromByte725 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte725` to result `intoIP524`
	// (`intoIP524` is now tainted).
	intoIP524 := net.IPv4(0, 0, 0, fromByte725)

	// Sink the tainted `intoIP524`:
	sink(intoIP524)
}

func TaintStepTest_NetIPv4Mask_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte674` into `intoIPMask767`.

	// Assume that `sourceCQL` has the underlying type of `fromByte674`:
	fromByte674 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte674` to result `intoIPMask767`
	// (`intoIPMask767` is now tainted).
	intoIPMask767 := net.IPv4Mask(fromByte674, 0, 0, 0)

	// Sink the tainted `intoIPMask767`:
	sink(intoIPMask767)
}

func TaintStepTest_NetIPv4Mask_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromByte228` into `intoIPMask656`.

	// Assume that `sourceCQL` has the underlying type of `fromByte228`:
	fromByte228 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte228` to result `intoIPMask656`
	// (`intoIPMask656` is now tainted).
	intoIPMask656 := net.IPv4Mask(0, fromByte228, 0, 0)

	// Sink the tainted `intoIPMask656`:
	sink(intoIPMask656)
}

func TaintStepTest_NetIPv4Mask_B0I2O0(sourceCQL interface{}) {
	// The flow is from `fromByte541` into `intoIPMask732`.

	// Assume that `sourceCQL` has the underlying type of `fromByte541`:
	fromByte541 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte541` to result `intoIPMask732`
	// (`intoIPMask732` is now tainted).
	intoIPMask732 := net.IPv4Mask(0, 0, fromByte541, 0)

	// Sink the tainted `intoIPMask732`:
	sink(intoIPMask732)
}

func TaintStepTest_NetIPv4Mask_B0I3O0(sourceCQL interface{}) {
	// The flow is from `fromByte843` into `intoIPMask408`.

	// Assume that `sourceCQL` has the underlying type of `fromByte843`:
	fromByte843 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte843` to result `intoIPMask408`
	// (`intoIPMask408` is now tainted).
	intoIPMask408 := net.IPv4Mask(0, 0, 0, fromByte843)

	// Sink the tainted `intoIPMask408`:
	sink(intoIPMask408)
}

func TaintStepTest_NetJoinHostPort_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString908` into `intoString990`.

	// Assume that `sourceCQL` has the underlying type of `fromString908`:
	fromString908 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString908` to result `intoString990`
	// (`intoString990` is now tainted).
	intoString990 := net.JoinHostPort(fromString908, "")

	// Sink the tainted `intoString990`:
	sink(intoString990)
}

func TaintStepTest_NetJoinHostPort_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromString598` into `intoString444`.

	// Assume that `sourceCQL` has the underlying type of `fromString598`:
	fromString598 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString598` to result `intoString444`
	// (`intoString444` is now tainted).
	intoString444 := net.JoinHostPort("", fromString598)

	// Sink the tainted `intoString444`:
	sink(intoString444)
}

func TaintStepTest_NetParseCIDR_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString715` into `intoIP699`.

	// Assume that `sourceCQL` has the underlying type of `fromString715`:
	fromString715 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString715` to result `intoIP699`
	// (`intoIP699` is now tainted).
	intoIP699, _, _ := net.ParseCIDR(fromString715)

	// Sink the tainted `intoIP699`:
	sink(intoIP699)
}

func TaintStepTest_NetParseCIDR_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromString333` into `intoIPNet155`.

	// Assume that `sourceCQL` has the underlying type of `fromString333`:
	fromString333 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString333` to result `intoIPNet155`
	// (`intoIPNet155` is now tainted).
	_, intoIPNet155, _ := net.ParseCIDR(fromString333)

	// Sink the tainted `intoIPNet155`:
	sink(intoIPNet155)
}

func TaintStepTest_NetParseIP_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString920` into `intoIP325`.

	// Assume that `sourceCQL` has the underlying type of `fromString920`:
	fromString920 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString920` to result `intoIP325`
	// (`intoIP325` is now tainted).
	intoIP325 := net.ParseIP(fromString920)

	// Sink the tainted `intoIP325`:
	sink(intoIP325)
}

func TaintStepTest_NetParseMAC_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString636` into `intoHardwareAddr926`.

	// Assume that `sourceCQL` has the underlying type of `fromString636`:
	fromString636 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString636` to result `intoHardwareAddr926`
	// (`intoHardwareAddr926` is now tainted).
	intoHardwareAddr926, _ := net.ParseMAC(fromString636)

	// Sink the tainted `intoHardwareAddr926`:
	sink(intoHardwareAddr926)
}

func TaintStepTest_NetPipe_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromConn938` into `intoConn718`.

	// Assume that `sourceCQL` has the underlying type of `fromConn938`:
	fromConn938 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the result `fromConn938` to result `intoConn718`
	// (extra steps needed)
	intermediateCQL, intoConn718 := net.Pipe()

	// Extra step (`fromConn938` taints `intermediateCQL`, which taints `intoConn718`:
	link(fromConn938, intermediateCQL)

	// Sink the tainted `intoConn718`:
	sink(intoConn718)
}

func TaintStepTest_NetPipe_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromConn855` into `intoConn981`.

	// Assume that `sourceCQL` has the underlying type of `fromConn855`:
	fromConn855 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the result `fromConn855` to result `intoConn981`
	// (extra steps needed)
	intoConn981, intermediateCQL := net.Pipe()

	// Extra step (`fromConn855` taints `intermediateCQL`, which taints `intoConn981`:
	link(fromConn855, intermediateCQL)

	// Sink the tainted `intoConn981`:
	sink(intoConn981)
}

func TaintStepTest_NetSplitHostPort_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromString785` into `intoString584`.

	// Assume that `sourceCQL` has the underlying type of `fromString785`:
	fromString785 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString785` to result `intoString584`
	// (`intoString584` is now tainted).
	intoString584, _, _ := net.SplitHostPort(fromString785)

	// Sink the tainted `intoString584`:
	sink(intoString584)
}

func TaintStepTest_NetSplitHostPort_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromString849` into `intoString395`.

	// Assume that `sourceCQL` has the underlying type of `fromString849`:
	fromString849 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString849` to result `intoString395`
	// (`intoString395` is now tainted).
	_, intoString395, _ := net.SplitHostPort(fromString849)

	// Sink the tainted `intoString395`:
	sink(intoString395)
}

func TaintStepTest_NetBuffersRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromBuffers127` into `intoByte751`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffers127`:
	fromBuffers127 := sourceCQL.(net.Buffers)

	// Declare `intoByte751` variable:
	var intoByte751 []byte

	// Call the method that transfers the taint
	// from the receiver `fromBuffers127` to the argument `intoByte751`
	// (`intoByte751` is now tainted).
	fromBuffers127.Read(intoByte751)

	// Sink the tainted `intoByte751`:
	sink(intoByte751)
}

func TaintStepTest_NetBuffersWriteTo_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromBuffers399` into `intoWriter190`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffers399`:
	fromBuffers399 := sourceCQL.(net.Buffers)

	// Declare `intoWriter190` variable:
	var intoWriter190 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromBuffers399` to the argument `intoWriter190`
	// (`intoWriter190` is now tainted).
	fromBuffers399.WriteTo(intoWriter190)

	// Sink the tainted `intoWriter190`:
	sink(intoWriter190)
}

func TaintStepTest_NetFlagsString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromFlags286` into `intoString251`.

	// Assume that `sourceCQL` has the underlying type of `fromFlags286`:
	fromFlags286 := sourceCQL.(net.Flags)

	// Call the method that transfers the taint
	// from the receiver `fromFlags286` to the result `intoString251`
	// (`intoString251` is now tainted).
	intoString251 := fromFlags286.String()

	// Sink the tainted `intoString251`:
	sink(intoString251)
}

func TaintStepTest_NetHardwareAddrString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromHardwareAddr820` into `intoString301`.

	// Assume that `sourceCQL` has the underlying type of `fromHardwareAddr820`:
	fromHardwareAddr820 := sourceCQL.(net.HardwareAddr)

	// Call the method that transfers the taint
	// from the receiver `fromHardwareAddr820` to the result `intoString301`
	// (`intoString301` is now tainted).
	intoString301 := fromHardwareAddr820.String()

	// Sink the tainted `intoString301`:
	sink(intoString301)
}

func TaintStepTest_NetIPMarshalText_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromIP377` into `intoByte659`.

	// Assume that `sourceCQL` has the underlying type of `fromIP377`:
	fromIP377 := sourceCQL.(net.IP)

	// Call the method that transfers the taint
	// from the receiver `fromIP377` to the result `intoByte659`
	// (`intoByte659` is now tainted).
	intoByte659, _ := fromIP377.MarshalText()

	// Sink the tainted `intoByte659`:
	sink(intoByte659)
}

func TaintStepTest_NetIPString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromIP546` into `intoString625`.

	// Assume that `sourceCQL` has the underlying type of `fromIP546`:
	fromIP546 := sourceCQL.(net.IP)

	// Call the method that transfers the taint
	// from the receiver `fromIP546` to the result `intoString625`
	// (`intoString625` is now tainted).
	intoString625 := fromIP546.String()

	// Sink the tainted `intoString625`:
	sink(intoString625)
}

func TaintStepTest_NetIPTo16_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromIP427` into `intoIP370`.

	// Assume that `sourceCQL` has the underlying type of `fromIP427`:
	fromIP427 := sourceCQL.(net.IP)

	// Call the method that transfers the taint
	// from the receiver `fromIP427` to the result `intoIP370`
	// (`intoIP370` is now tainted).
	intoIP370 := fromIP427.To16()

	// Sink the tainted `intoIP370`:
	sink(intoIP370)
}

func TaintStepTest_NetIPTo4_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromIP358` into `intoIP246`.

	// Assume that `sourceCQL` has the underlying type of `fromIP358`:
	fromIP358 := sourceCQL.(net.IP)

	// Call the method that transfers the taint
	// from the receiver `fromIP358` to the result `intoIP246`
	// (`intoIP246` is now tainted).
	intoIP246 := fromIP358.To4()

	// Sink the tainted `intoIP246`:
	sink(intoIP246)
}

func TaintStepTest_NetIPUnmarshalText_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte858` into `intoIP972`.

	// Assume that `sourceCQL` has the underlying type of `fromByte858`:
	fromByte858 := sourceCQL.([]byte)

	// Declare `intoIP972` variable:
	var intoIP972 net.IP

	// Call the method that transfers the taint
	// from the parameter `fromByte858` to the receiver `intoIP972`
	// (`intoIP972` is now tainted).
	intoIP972.UnmarshalText(fromByte858)

	// Sink the tainted `intoIP972`:
	sink(intoIP972)
}

func TaintStepTest_NetIPAddrString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromIPAddr458` into `intoString251`.

	// Assume that `sourceCQL` has the underlying type of `fromIPAddr458`:
	fromIPAddr458 := sourceCQL.(net.IPAddr)

	// Call the method that transfers the taint
	// from the receiver `fromIPAddr458` to the result `intoString251`
	// (`intoString251` is now tainted).
	intoString251 := fromIPAddr458.String()

	// Sink the tainted `intoString251`:
	sink(intoString251)
}

func TaintStepTest_NetIPConnReadFrom_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromIPConn825` into `intoByte599`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn825`:
	fromIPConn825 := sourceCQL.(net.IPConn)

	// Declare `intoByte599` variable:
	var intoByte599 []byte

	// Call the method that transfers the taint
	// from the receiver `fromIPConn825` to the argument `intoByte599`
	// (`intoByte599` is now tainted).
	fromIPConn825.ReadFrom(intoByte599)

	// Sink the tainted `intoByte599`:
	sink(intoByte599)
}

func TaintStepTest_NetIPConnReadFromIP_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromIPConn317` into `intoByte675`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn317`:
	fromIPConn317 := sourceCQL.(net.IPConn)

	// Declare `intoByte675` variable:
	var intoByte675 []byte

	// Call the method that transfers the taint
	// from the receiver `fromIPConn317` to the argument `intoByte675`
	// (`intoByte675` is now tainted).
	fromIPConn317.ReadFromIP(intoByte675)

	// Sink the tainted `intoByte675`:
	sink(intoByte675)
}

func TaintStepTest_NetIPConnReadMsgIP_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromIPConn252` into `intoByte893`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn252`:
	fromIPConn252 := sourceCQL.(net.IPConn)

	// Declare `intoByte893` variable:
	var intoByte893 []byte

	// Call the method that transfers the taint
	// from the receiver `fromIPConn252` to the argument `intoByte893`
	// (`intoByte893` is now tainted).
	fromIPConn252.ReadMsgIP(intoByte893, nil)

	// Sink the tainted `intoByte893`:
	sink(intoByte893)
}

func TaintStepTest_NetIPConnReadMsgIP_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromIPConn683` into `intoByte238`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn683`:
	fromIPConn683 := sourceCQL.(net.IPConn)

	// Declare `intoByte238` variable:
	var intoByte238 []byte

	// Call the method that transfers the taint
	// from the receiver `fromIPConn683` to the argument `intoByte238`
	// (`intoByte238` is now tainted).
	fromIPConn683.ReadMsgIP(nil, intoByte238)

	// Sink the tainted `intoByte238`:
	sink(intoByte238)
}

func TaintStepTest_NetIPConnSyscallConn_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromIPConn970` into `intoRawConn966`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn970`:
	fromIPConn970 := sourceCQL.(net.IPConn)

	// Call the method that transfers the taint
	// from the receiver `fromIPConn970` to the result `intoRawConn966`
	// (`intoRawConn966` is now tainted).
	intoRawConn966, _ := fromIPConn970.SyscallConn()

	// Sink the tainted `intoRawConn966`:
	sink(intoRawConn966)
}

func TaintStepTest_NetIPConnSyscallConn_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromRawConn200` into `intoIPConn224`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn200`:
	fromRawConn200 := sourceCQL.(syscall.RawConn)

	// Declare `intoIPConn224` variable:
	var intoIPConn224 net.IPConn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoIPConn224`:
	intermediateCQL, _ := intoIPConn224.SyscallConn()

	// Extra step (`fromRawConn200` taints `intermediateCQL`, which taints `intoIPConn224`:
	link(fromRawConn200, intermediateCQL)

	// Sink the tainted `intoIPConn224`:
	sink(intoIPConn224)
}

func TaintStepTest_NetIPConnWriteMsgIP_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte952` into `intoIPConn900`.

	// Assume that `sourceCQL` has the underlying type of `fromByte952`:
	fromByte952 := sourceCQL.([]byte)

	// Declare `intoIPConn900` variable:
	var intoIPConn900 net.IPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte952` to the receiver `intoIPConn900`
	// (`intoIPConn900` is now tainted).
	intoIPConn900.WriteMsgIP(fromByte952, nil, nil)

	// Sink the tainted `intoIPConn900`:
	sink(intoIPConn900)
}

func TaintStepTest_NetIPConnWriteMsgIP_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromByte569` into `intoIPConn251`.

	// Assume that `sourceCQL` has the underlying type of `fromByte569`:
	fromByte569 := sourceCQL.([]byte)

	// Declare `intoIPConn251` variable:
	var intoIPConn251 net.IPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte569` to the receiver `intoIPConn251`
	// (`intoIPConn251` is now tainted).
	intoIPConn251.WriteMsgIP(nil, fromByte569, nil)

	// Sink the tainted `intoIPConn251`:
	sink(intoIPConn251)
}

func TaintStepTest_NetIPConnWriteTo_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte801` into `intoIPConn542`.

	// Assume that `sourceCQL` has the underlying type of `fromByte801`:
	fromByte801 := sourceCQL.([]byte)

	// Declare `intoIPConn542` variable:
	var intoIPConn542 net.IPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte801` to the receiver `intoIPConn542`
	// (`intoIPConn542` is now tainted).
	intoIPConn542.WriteTo(fromByte801, nil)

	// Sink the tainted `intoIPConn542`:
	sink(intoIPConn542)
}

func TaintStepTest_NetIPConnWriteToIP_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte275` into `intoIPConn960`.

	// Assume that `sourceCQL` has the underlying type of `fromByte275`:
	fromByte275 := sourceCQL.([]byte)

	// Declare `intoIPConn960` variable:
	var intoIPConn960 net.IPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte275` to the receiver `intoIPConn960`
	// (`intoIPConn960` is now tainted).
	intoIPConn960.WriteToIP(fromByte275, nil)

	// Sink the tainted `intoIPConn960`:
	sink(intoIPConn960)
}

func TaintStepTest_NetIPMaskString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromIPMask656` into `intoString576`.

	// Assume that `sourceCQL` has the underlying type of `fromIPMask656`:
	fromIPMask656 := sourceCQL.(net.IPMask)

	// Call the method that transfers the taint
	// from the receiver `fromIPMask656` to the result `intoString576`
	// (`intoString576` is now tainted).
	intoString576 := fromIPMask656.String()

	// Sink the tainted `intoString576`:
	sink(intoString576)
}

func TaintStepTest_NetIPNetString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromIPNet398` into `intoString212`.

	// Assume that `sourceCQL` has the underlying type of `fromIPNet398`:
	fromIPNet398 := sourceCQL.(net.IPNet)

	// Call the method that transfers the taint
	// from the receiver `fromIPNet398` to the result `intoString212`
	// (`intoString212` is now tainted).
	intoString212 := fromIPNet398.String()

	// Sink the tainted `intoString212`:
	sink(intoString212)
}

func TaintStepTest_NetTCPAddrString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTCPAddr699` into `intoString348`.

	// Assume that `sourceCQL` has the underlying type of `fromTCPAddr699`:
	fromTCPAddr699 := sourceCQL.(net.TCPAddr)

	// Call the method that transfers the taint
	// from the receiver `fromTCPAddr699` to the result `intoString348`
	// (`intoString348` is now tainted).
	intoString348 := fromTCPAddr699.String()

	// Sink the tainted `intoString348`:
	sink(intoString348)
}

func TaintStepTest_NetTCPConnReadFrom_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromReader122` into `intoTCPConn897`.

	// Assume that `sourceCQL` has the underlying type of `fromReader122`:
	fromReader122 := sourceCQL.(io.Reader)

	// Declare `intoTCPConn897` variable:
	var intoTCPConn897 net.TCPConn

	// Call the method that transfers the taint
	// from the parameter `fromReader122` to the receiver `intoTCPConn897`
	// (`intoTCPConn897` is now tainted).
	intoTCPConn897.ReadFrom(fromReader122)

	// Sink the tainted `intoTCPConn897`:
	sink(intoTCPConn897)
}

func TaintStepTest_NetTCPConnSyscallConn_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTCPConn571` into `intoRawConn123`.

	// Assume that `sourceCQL` has the underlying type of `fromTCPConn571`:
	fromTCPConn571 := sourceCQL.(net.TCPConn)

	// Call the method that transfers the taint
	// from the receiver `fromTCPConn571` to the result `intoRawConn123`
	// (`intoRawConn123` is now tainted).
	intoRawConn123, _ := fromTCPConn571.SyscallConn()

	// Sink the tainted `intoRawConn123`:
	sink(intoRawConn123)
}

func TaintStepTest_NetTCPConnSyscallConn_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromRawConn806` into `intoTCPConn241`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn806`:
	fromRawConn806 := sourceCQL.(syscall.RawConn)

	// Declare `intoTCPConn241` variable:
	var intoTCPConn241 net.TCPConn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoTCPConn241`:
	intermediateCQL, _ := intoTCPConn241.SyscallConn()

	// Extra step (`fromRawConn806` taints `intermediateCQL`, which taints `intoTCPConn241`:
	link(fromRawConn806, intermediateCQL)

	// Sink the tainted `intoTCPConn241`:
	sink(intoTCPConn241)
}

func TaintStepTest_NetTCPListenerFile_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTCPListener433` into `intoFile751`.

	// Assume that `sourceCQL` has the underlying type of `fromTCPListener433`:
	fromTCPListener433 := sourceCQL.(net.TCPListener)

	// Call the method that transfers the taint
	// from the receiver `fromTCPListener433` to the result `intoFile751`
	// (`intoFile751` is now tainted).
	intoFile751, _ := fromTCPListener433.File()

	// Sink the tainted `intoFile751`:
	sink(intoFile751)
}

func TaintStepTest_NetTCPListenerFile_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromFile865` into `intoTCPListener460`.

	// Assume that `sourceCQL` has the underlying type of `fromFile865`:
	fromFile865 := sourceCQL.(*os.File)

	// Declare `intoTCPListener460` variable:
	var intoTCPListener460 net.TCPListener

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoTCPListener460`:
	intermediateCQL, _ := intoTCPListener460.File()

	// Extra step (`fromFile865` taints `intermediateCQL`, which taints `intoTCPListener460`:
	link(fromFile865, intermediateCQL)

	// Sink the tainted `intoTCPListener460`:
	sink(intoTCPListener460)
}

func TaintStepTest_NetTCPListenerSyscallConn_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromTCPListener132` into `intoRawConn948`.

	// Assume that `sourceCQL` has the underlying type of `fromTCPListener132`:
	fromTCPListener132 := sourceCQL.(net.TCPListener)

	// Call the method that transfers the taint
	// from the receiver `fromTCPListener132` to the result `intoRawConn948`
	// (`intoRawConn948` is now tainted).
	intoRawConn948, _ := fromTCPListener132.SyscallConn()

	// Sink the tainted `intoRawConn948`:
	sink(intoRawConn948)
}

func TaintStepTest_NetTCPListenerSyscallConn_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromRawConn364` into `intoTCPListener257`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn364`:
	fromRawConn364 := sourceCQL.(syscall.RawConn)

	// Declare `intoTCPListener257` variable:
	var intoTCPListener257 net.TCPListener

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoTCPListener257`:
	intermediateCQL, _ := intoTCPListener257.SyscallConn()

	// Extra step (`fromRawConn364` taints `intermediateCQL`, which taints `intoTCPListener257`:
	link(fromRawConn364, intermediateCQL)

	// Sink the tainted `intoTCPListener257`:
	sink(intoTCPListener257)
}

func TaintStepTest_NetUDPConnReadFrom_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUDPConn756` into `intoByte467`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn756`:
	fromUDPConn756 := sourceCQL.(net.UDPConn)

	// Declare `intoByte467` variable:
	var intoByte467 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn756` to the argument `intoByte467`
	// (`intoByte467` is now tainted).
	fromUDPConn756.ReadFrom(intoByte467)

	// Sink the tainted `intoByte467`:
	sink(intoByte467)
}

func TaintStepTest_NetUDPConnReadFromUDP_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUDPConn232` into `intoByte172`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn232`:
	fromUDPConn232 := sourceCQL.(net.UDPConn)

	// Declare `intoByte172` variable:
	var intoByte172 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn232` to the argument `intoByte172`
	// (`intoByte172` is now tainted).
	fromUDPConn232.ReadFromUDP(intoByte172)

	// Sink the tainted `intoByte172`:
	sink(intoByte172)
}

func TaintStepTest_NetUDPConnReadMsgUDP_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUDPConn477` into `intoByte638`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn477`:
	fromUDPConn477 := sourceCQL.(net.UDPConn)

	// Declare `intoByte638` variable:
	var intoByte638 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn477` to the argument `intoByte638`
	// (`intoByte638` is now tainted).
	fromUDPConn477.ReadMsgUDP(intoByte638, nil)

	// Sink the tainted `intoByte638`:
	sink(intoByte638)
}

func TaintStepTest_NetUDPConnReadMsgUDP_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromUDPConn560` into `intoByte618`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn560`:
	fromUDPConn560 := sourceCQL.(net.UDPConn)

	// Declare `intoByte618` variable:
	var intoByte618 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn560` to the argument `intoByte618`
	// (`intoByte618` is now tainted).
	fromUDPConn560.ReadMsgUDP(nil, intoByte618)

	// Sink the tainted `intoByte618`:
	sink(intoByte618)
}

func TaintStepTest_NetUDPConnSyscallConn_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUDPConn576` into `intoRawConn121`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn576`:
	fromUDPConn576 := sourceCQL.(net.UDPConn)

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn576` to the result `intoRawConn121`
	// (`intoRawConn121` is now tainted).
	intoRawConn121, _ := fromUDPConn576.SyscallConn()

	// Sink the tainted `intoRawConn121`:
	sink(intoRawConn121)
}

func TaintStepTest_NetUDPConnSyscallConn_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromRawConn402` into `intoUDPConn212`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn402`:
	fromRawConn402 := sourceCQL.(syscall.RawConn)

	// Declare `intoUDPConn212` variable:
	var intoUDPConn212 net.UDPConn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoUDPConn212`:
	intermediateCQL, _ := intoUDPConn212.SyscallConn()

	// Extra step (`fromRawConn402` taints `intermediateCQL`, which taints `intoUDPConn212`:
	link(fromRawConn402, intermediateCQL)

	// Sink the tainted `intoUDPConn212`:
	sink(intoUDPConn212)
}

func TaintStepTest_NetUDPConnWriteMsgUDP_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte974` into `intoUDPConn264`.

	// Assume that `sourceCQL` has the underlying type of `fromByte974`:
	fromByte974 := sourceCQL.([]byte)

	// Declare `intoUDPConn264` variable:
	var intoUDPConn264 net.UDPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte974` to the receiver `intoUDPConn264`
	// (`intoUDPConn264` is now tainted).
	intoUDPConn264.WriteMsgUDP(fromByte974, nil, nil)

	// Sink the tainted `intoUDPConn264`:
	sink(intoUDPConn264)
}

func TaintStepTest_NetUDPConnWriteMsgUDP_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromByte325` into `intoUDPConn565`.

	// Assume that `sourceCQL` has the underlying type of `fromByte325`:
	fromByte325 := sourceCQL.([]byte)

	// Declare `intoUDPConn565` variable:
	var intoUDPConn565 net.UDPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte325` to the receiver `intoUDPConn565`
	// (`intoUDPConn565` is now tainted).
	intoUDPConn565.WriteMsgUDP(nil, fromByte325, nil)

	// Sink the tainted `intoUDPConn565`:
	sink(intoUDPConn565)
}

func TaintStepTest_NetUDPConnWriteTo_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte175` into `intoUDPConn723`.

	// Assume that `sourceCQL` has the underlying type of `fromByte175`:
	fromByte175 := sourceCQL.([]byte)

	// Declare `intoUDPConn723` variable:
	var intoUDPConn723 net.UDPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte175` to the receiver `intoUDPConn723`
	// (`intoUDPConn723` is now tainted).
	intoUDPConn723.WriteTo(fromByte175, nil)

	// Sink the tainted `intoUDPConn723`:
	sink(intoUDPConn723)
}

func TaintStepTest_NetUDPConnWriteToUDP_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte366` into `intoUDPConn800`.

	// Assume that `sourceCQL` has the underlying type of `fromByte366`:
	fromByte366 := sourceCQL.([]byte)

	// Declare `intoUDPConn800` variable:
	var intoUDPConn800 net.UDPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte366` to the receiver `intoUDPConn800`
	// (`intoUDPConn800` is now tainted).
	intoUDPConn800.WriteToUDP(fromByte366, nil)

	// Sink the tainted `intoUDPConn800`:
	sink(intoUDPConn800)
}

func TaintStepTest_NetUnixAddrString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUnixAddr144` into `intoString441`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixAddr144`:
	fromUnixAddr144 := sourceCQL.(net.UnixAddr)

	// Call the method that transfers the taint
	// from the receiver `fromUnixAddr144` to the result `intoString441`
	// (`intoString441` is now tainted).
	intoString441 := fromUnixAddr144.String()

	// Sink the tainted `intoString441`:
	sink(intoString441)
}

func TaintStepTest_NetUnixConnReadFrom_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUnixConn248` into `intoByte868`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn248`:
	fromUnixConn248 := sourceCQL.(net.UnixConn)

	// Declare `intoByte868` variable:
	var intoByte868 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn248` to the argument `intoByte868`
	// (`intoByte868` is now tainted).
	fromUnixConn248.ReadFrom(intoByte868)

	// Sink the tainted `intoByte868`:
	sink(intoByte868)
}

func TaintStepTest_NetUnixConnReadFromUnix_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUnixConn321` into `intoByte453`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn321`:
	fromUnixConn321 := sourceCQL.(net.UnixConn)

	// Declare `intoByte453` variable:
	var intoByte453 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn321` to the argument `intoByte453`
	// (`intoByte453` is now tainted).
	fromUnixConn321.ReadFromUnix(intoByte453)

	// Sink the tainted `intoByte453`:
	sink(intoByte453)
}

func TaintStepTest_NetUnixConnReadMsgUnix_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUnixConn954` into `intoByte672`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn954`:
	fromUnixConn954 := sourceCQL.(net.UnixConn)

	// Declare `intoByte672` variable:
	var intoByte672 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn954` to the argument `intoByte672`
	// (`intoByte672` is now tainted).
	fromUnixConn954.ReadMsgUnix(intoByte672, nil)

	// Sink the tainted `intoByte672`:
	sink(intoByte672)
}

func TaintStepTest_NetUnixConnReadMsgUnix_B0I0O1(sourceCQL interface{}) {
	// The flow is from `fromUnixConn311` into `intoByte634`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn311`:
	fromUnixConn311 := sourceCQL.(net.UnixConn)

	// Declare `intoByte634` variable:
	var intoByte634 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn311` to the argument `intoByte634`
	// (`intoByte634` is now tainted).
	fromUnixConn311.ReadMsgUnix(nil, intoByte634)

	// Sink the tainted `intoByte634`:
	sink(intoByte634)
}

func TaintStepTest_NetUnixConnSyscallConn_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUnixConn259` into `intoRawConn597`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn259`:
	fromUnixConn259 := sourceCQL.(net.UnixConn)

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn259` to the result `intoRawConn597`
	// (`intoRawConn597` is now tainted).
	intoRawConn597, _ := fromUnixConn259.SyscallConn()

	// Sink the tainted `intoRawConn597`:
	sink(intoRawConn597)
}

func TaintStepTest_NetUnixConnSyscallConn_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromRawConn663` into `intoUnixConn296`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn663`:
	fromRawConn663 := sourceCQL.(syscall.RawConn)

	// Declare `intoUnixConn296` variable:
	var intoUnixConn296 net.UnixConn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoUnixConn296`:
	intermediateCQL, _ := intoUnixConn296.SyscallConn()

	// Extra step (`fromRawConn663` taints `intermediateCQL`, which taints `intoUnixConn296`:
	link(fromRawConn663, intermediateCQL)

	// Sink the tainted `intoUnixConn296`:
	sink(intoUnixConn296)
}

func TaintStepTest_NetUnixConnWriteMsgUnix_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte642` into `intoUnixConn604`.

	// Assume that `sourceCQL` has the underlying type of `fromByte642`:
	fromByte642 := sourceCQL.([]byte)

	// Declare `intoUnixConn604` variable:
	var intoUnixConn604 net.UnixConn

	// Call the method that transfers the taint
	// from the parameter `fromByte642` to the receiver `intoUnixConn604`
	// (`intoUnixConn604` is now tainted).
	intoUnixConn604.WriteMsgUnix(fromByte642, nil, nil)

	// Sink the tainted `intoUnixConn604`:
	sink(intoUnixConn604)
}

func TaintStepTest_NetUnixConnWriteMsgUnix_B0I1O0(sourceCQL interface{}) {
	// The flow is from `fromByte138` into `intoUnixConn640`.

	// Assume that `sourceCQL` has the underlying type of `fromByte138`:
	fromByte138 := sourceCQL.([]byte)

	// Declare `intoUnixConn640` variable:
	var intoUnixConn640 net.UnixConn

	// Call the method that transfers the taint
	// from the parameter `fromByte138` to the receiver `intoUnixConn640`
	// (`intoUnixConn640` is now tainted).
	intoUnixConn640.WriteMsgUnix(nil, fromByte138, nil)

	// Sink the tainted `intoUnixConn640`:
	sink(intoUnixConn640)
}

func TaintStepTest_NetUnixConnWriteTo_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte649` into `intoUnixConn257`.

	// Assume that `sourceCQL` has the underlying type of `fromByte649`:
	fromByte649 := sourceCQL.([]byte)

	// Declare `intoUnixConn257` variable:
	var intoUnixConn257 net.UnixConn

	// Call the method that transfers the taint
	// from the parameter `fromByte649` to the receiver `intoUnixConn257`
	// (`intoUnixConn257` is now tainted).
	intoUnixConn257.WriteTo(fromByte649, nil)

	// Sink the tainted `intoUnixConn257`:
	sink(intoUnixConn257)
}

func TaintStepTest_NetUnixConnWriteToUnix_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte844` into `intoUnixConn611`.

	// Assume that `sourceCQL` has the underlying type of `fromByte844`:
	fromByte844 := sourceCQL.([]byte)

	// Declare `intoUnixConn611` variable:
	var intoUnixConn611 net.UnixConn

	// Call the method that transfers the taint
	// from the parameter `fromByte844` to the receiver `intoUnixConn611`
	// (`intoUnixConn611` is now tainted).
	intoUnixConn611.WriteToUnix(fromByte844, nil)

	// Sink the tainted `intoUnixConn611`:
	sink(intoUnixConn611)
}

func TaintStepTest_NetUnixListenerFile_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUnixListener279` into `intoFile388`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixListener279`:
	fromUnixListener279 := sourceCQL.(net.UnixListener)

	// Call the method that transfers the taint
	// from the receiver `fromUnixListener279` to the result `intoFile388`
	// (`intoFile388` is now tainted).
	intoFile388, _ := fromUnixListener279.File()

	// Sink the tainted `intoFile388`:
	sink(intoFile388)
}

func TaintStepTest_NetUnixListenerFile_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromFile540` into `intoUnixListener894`.

	// Assume that `sourceCQL` has the underlying type of `fromFile540`:
	fromFile540 := sourceCQL.(*os.File)

	// Declare `intoUnixListener894` variable:
	var intoUnixListener894 net.UnixListener

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoUnixListener894`:
	intermediateCQL, _ := intoUnixListener894.File()

	// Extra step (`fromFile540` taints `intermediateCQL`, which taints `intoUnixListener894`:
	link(fromFile540, intermediateCQL)

	// Sink the tainted `intoUnixListener894`:
	sink(intoUnixListener894)
}

func TaintStepTest_NetUnixListenerSyscallConn_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromUnixListener603` into `intoRawConn790`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixListener603`:
	fromUnixListener603 := sourceCQL.(net.UnixListener)

	// Call the method that transfers the taint
	// from the receiver `fromUnixListener603` to the result `intoRawConn790`
	// (`intoRawConn790` is now tainted).
	intoRawConn790, _ := fromUnixListener603.SyscallConn()

	// Sink the tainted `intoRawConn790`:
	sink(intoRawConn790)
}

func TaintStepTest_NetUnixListenerSyscallConn_B1I0O0(sourceCQL interface{}) {
	// The flow is from `fromRawConn454` into `intoUnixListener908`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn454`:
	fromRawConn454 := sourceCQL.(syscall.RawConn)

	// Declare `intoUnixListener908` variable:
	var intoUnixListener908 net.UnixListener

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoUnixListener908`:
	intermediateCQL, _ := intoUnixListener908.SyscallConn()

	// Extra step (`fromRawConn454` taints `intermediateCQL`, which taints `intoUnixListener908`:
	link(fromRawConn454, intermediateCQL)

	// Sink the tainted `intoUnixListener908`:
	sink(intoUnixListener908)
}

func TaintStepTest_NetConnRead_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromConn175` into `intoByte551`.

	// Assume that `sourceCQL` has the underlying type of `fromConn175`:
	fromConn175 := sourceCQL.(net.Conn)

	// Declare `intoByte551` variable:
	var intoByte551 []byte

	// Call the method that transfers the taint
	// from the receiver `fromConn175` to the argument `intoByte551`
	// (`intoByte551` is now tainted).
	fromConn175.Read(intoByte551)

	// Sink the tainted `intoByte551`:
	sink(intoByte551)
}

func TaintStepTest_NetPacketConnReadFrom_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromPacketConn733` into `intoByte404`.

	// Assume that `sourceCQL` has the underlying type of `fromPacketConn733`:
	fromPacketConn733 := sourceCQL.(net.PacketConn)

	// Declare `intoByte404` variable:
	var intoByte404 []byte

	// Call the method that transfers the taint
	// from the receiver `fromPacketConn733` to the argument `intoByte404`
	// (`intoByte404` is now tainted).
	fromPacketConn733.ReadFrom(intoByte404)

	// Sink the tainted `intoByte404`:
	sink(intoByte404)
}

func TaintStepTest_NetAddrString_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromAddr365` into `intoString962`.

	// Assume that `sourceCQL` has the underlying type of `fromAddr365`:
	fromAddr365 := sourceCQL.(net.Addr)

	// Call the method that transfers the taint
	// from the receiver `fromAddr365` to the result `intoString962`
	// (`intoString962` is now tainted).
	intoString962 := fromAddr365.String()

	// Sink the tainted `intoString962`:
	sink(intoString962)
}

func TaintStepTest_NetConnWrite_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte649` into `intoConn702`.

	// Assume that `sourceCQL` has the underlying type of `fromByte649`:
	fromByte649 := sourceCQL.([]byte)

	// Declare `intoConn702` variable:
	var intoConn702 net.Conn

	// Call the method that transfers the taint
	// from the parameter `fromByte649` to the receiver `intoConn702`
	// (`intoConn702` is now tainted).
	intoConn702.Write(fromByte649)

	// Sink the tainted `intoConn702`:
	sink(intoConn702)
}

func TaintStepTest_NetPacketConnWriteTo_B0I0O0(sourceCQL interface{}) {
	// The flow is from `fromByte220` into `intoPacketConn789`.

	// Assume that `sourceCQL` has the underlying type of `fromByte220`:
	fromByte220 := sourceCQL.([]byte)

	// Declare `intoPacketConn789` variable:
	var intoPacketConn789 net.PacketConn

	// Call the method that transfers the taint
	// from the parameter `fromByte220` to the receiver `intoPacketConn789`
	// (`intoPacketConn789` is now tainted).
	intoPacketConn789.WriteTo(fromByte220, nil)

	// Sink the tainted `intoPacketConn789`:
	sink(intoPacketConn789)
}

func RunAllTaints_Net() {
	{
		source := newSource()
		TaintStepTest_NetFileConn_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetFileConn_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetFileListener_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetFileListener_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetFilePacketConn_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetFilePacketConn_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPv4_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPv4_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPv4_B0I2O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPv4_B0I3O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPv4Mask_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPv4Mask_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPv4Mask_B0I2O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPv4Mask_B0I3O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetJoinHostPort_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetJoinHostPort_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetParseCIDR_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetParseCIDR_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_NetParseIP_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetParseMAC_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetPipe_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetPipe_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetSplitHostPort_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetSplitHostPort_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_NetBuffersRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetBuffersWriteTo_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetFlagsString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetHardwareAddrString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPMarshalText_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPTo16_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPTo4_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPUnmarshalText_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPAddrString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPConnReadFrom_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPConnReadFromIP_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPConnReadMsgIP_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPConnReadMsgIP_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPConnSyscallConn_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPConnSyscallConn_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPConnWriteMsgIP_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPConnWriteMsgIP_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPConnWriteTo_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPConnWriteToIP_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPMaskString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetIPNetString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTCPAddrString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTCPConnReadFrom_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTCPConnSyscallConn_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTCPConnSyscallConn_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTCPListenerFile_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTCPListenerFile_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTCPListenerSyscallConn_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetTCPListenerSyscallConn_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUDPConnReadFrom_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUDPConnReadFromUDP_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUDPConnReadMsgUDP_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUDPConnReadMsgUDP_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUDPConnSyscallConn_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUDPConnSyscallConn_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUDPConnWriteMsgUDP_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUDPConnWriteMsgUDP_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUDPConnWriteTo_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUDPConnWriteToUDP_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUnixAddrString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUnixConnReadFrom_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUnixConnReadFromUnix_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUnixConnReadMsgUnix_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUnixConnReadMsgUnix_B0I0O1(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUnixConnSyscallConn_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUnixConnSyscallConn_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUnixConnWriteMsgUnix_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUnixConnWriteMsgUnix_B0I1O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUnixConnWriteTo_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUnixConnWriteToUnix_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUnixListenerFile_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUnixListenerFile_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUnixListenerSyscallConn_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetUnixListenerSyscallConn_B1I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetConnRead_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetPacketConnReadFrom_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetAddrString_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetConnWrite_B0I0O0(source)
	}
	{
		source := newSource()
		TaintStepTest_NetPacketConnWriteTo_B0I0O0(source)
	}
}
