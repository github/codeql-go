// WARNING: This file was automatically generated. DO NOT EDIT.

package main

import (
	"io"
	"net"
	"os"
	"syscall"
)

func TaintStepTest_NetFileConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile433` into `intoConn130`.

	// Assume that `sourceCQL` has the underlying type of `fromFile433`:
	fromFile433 := sourceCQL.(*os.File)

	// Call the function that transfers the taint
	// from the parameter `fromFile433` to result `intoConn130`
	// (`intoConn130` is now tainted).
	intoConn130, _ := net.FileConn(fromFile433)

	// Return the tainted `intoConn130`:
	return intoConn130
}

func TaintStepTest_NetFileConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn472` into `intoFile822`.

	// Assume that `sourceCQL` has the underlying type of `fromConn472`:
	fromConn472 := sourceCQL.(net.Conn)

	// Declare `intoFile822` variable:
	var intoFile822 *os.File

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoFile822`:
	intermediateCQL, _ := net.FileConn(intoFile822)

	// Extra step (`fromConn472` taints `intermediateCQL`, which taints `intoFile822`:
	link(fromConn472, intermediateCQL)

	// Return the tainted `intoFile822`:
	return intoFile822
}

func TaintStepTest_NetFileListener_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile195` into `intoListener567`.

	// Assume that `sourceCQL` has the underlying type of `fromFile195`:
	fromFile195 := sourceCQL.(*os.File)

	// Call the function that transfers the taint
	// from the parameter `fromFile195` to result `intoListener567`
	// (`intoListener567` is now tainted).
	intoListener567, _ := net.FileListener(fromFile195)

	// Return the tainted `intoListener567`:
	return intoListener567
}

func TaintStepTest_NetFileListener_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromListener613` into `intoFile783`.

	// Assume that `sourceCQL` has the underlying type of `fromListener613`:
	fromListener613 := sourceCQL.(net.Listener)

	// Declare `intoFile783` variable:
	var intoFile783 *os.File

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoFile783`:
	intermediateCQL, _ := net.FileListener(intoFile783)

	// Extra step (`fromListener613` taints `intermediateCQL`, which taints `intoFile783`:
	link(fromListener613, intermediateCQL)

	// Return the tainted `intoFile783`:
	return intoFile783
}

func TaintStepTest_NetFilePacketConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile632` into `intoPacketConn937`.

	// Assume that `sourceCQL` has the underlying type of `fromFile632`:
	fromFile632 := sourceCQL.(*os.File)

	// Call the function that transfers the taint
	// from the parameter `fromFile632` to result `intoPacketConn937`
	// (`intoPacketConn937` is now tainted).
	intoPacketConn937, _ := net.FilePacketConn(fromFile632)

	// Return the tainted `intoPacketConn937`:
	return intoPacketConn937
}

func TaintStepTest_NetFilePacketConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPacketConn383` into `intoFile947`.

	// Assume that `sourceCQL` has the underlying type of `fromPacketConn383`:
	fromPacketConn383 := sourceCQL.(net.PacketConn)

	// Declare `intoFile947` variable:
	var intoFile947 *os.File

	// Call the function that will transfer the taint
	// from the result `intermediateCQL` to parameter `intoFile947`:
	intermediateCQL, _ := net.FilePacketConn(intoFile947)

	// Extra step (`fromPacketConn383` taints `intermediateCQL`, which taints `intoFile947`:
	link(fromPacketConn383, intermediateCQL)

	// Return the tainted `intoFile947`:
	return intoFile947
}

func TaintStepTest_NetIPv4_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte702` into `intoIP209`.

	// Assume that `sourceCQL` has the underlying type of `fromByte702`:
	fromByte702 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte702` to result `intoIP209`
	// (`intoIP209` is now tainted).
	intoIP209 := net.IPv4(fromByte702, 0, 0, 0)

	// Return the tainted `intoIP209`:
	return intoIP209
}

func TaintStepTest_NetIPv4_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte988` into `intoIP477`.

	// Assume that `sourceCQL` has the underlying type of `fromByte988`:
	fromByte988 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte988` to result `intoIP477`
	// (`intoIP477` is now tainted).
	intoIP477 := net.IPv4(0, fromByte988, 0, 0)

	// Return the tainted `intoIP477`:
	return intoIP477
}

func TaintStepTest_NetIPv4_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte936` into `intoIP515`.

	// Assume that `sourceCQL` has the underlying type of `fromByte936`:
	fromByte936 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte936` to result `intoIP515`
	// (`intoIP515` is now tainted).
	intoIP515 := net.IPv4(0, 0, fromByte936, 0)

	// Return the tainted `intoIP515`:
	return intoIP515
}

func TaintStepTest_NetIPv4_B0I3O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte710` into `intoIP254`.

	// Assume that `sourceCQL` has the underlying type of `fromByte710`:
	fromByte710 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte710` to result `intoIP254`
	// (`intoIP254` is now tainted).
	intoIP254 := net.IPv4(0, 0, 0, fromByte710)

	// Return the tainted `intoIP254`:
	return intoIP254
}

func TaintStepTest_NetIPv4Mask_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte910` into `intoIPMask145`.

	// Assume that `sourceCQL` has the underlying type of `fromByte910`:
	fromByte910 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte910` to result `intoIPMask145`
	// (`intoIPMask145` is now tainted).
	intoIPMask145 := net.IPv4Mask(fromByte910, 0, 0, 0)

	// Return the tainted `intoIPMask145`:
	return intoIPMask145
}

func TaintStepTest_NetIPv4Mask_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte239` into `intoIPMask224`.

	// Assume that `sourceCQL` has the underlying type of `fromByte239`:
	fromByte239 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte239` to result `intoIPMask224`
	// (`intoIPMask224` is now tainted).
	intoIPMask224 := net.IPv4Mask(0, fromByte239, 0, 0)

	// Return the tainted `intoIPMask224`:
	return intoIPMask224
}

func TaintStepTest_NetIPv4Mask_B0I2O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte439` into `intoIPMask718`.

	// Assume that `sourceCQL` has the underlying type of `fromByte439`:
	fromByte439 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte439` to result `intoIPMask718`
	// (`intoIPMask718` is now tainted).
	intoIPMask718 := net.IPv4Mask(0, 0, fromByte439, 0)

	// Return the tainted `intoIPMask718`:
	return intoIPMask718
}

func TaintStepTest_NetIPv4Mask_B0I3O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte179` into `intoIPMask346`.

	// Assume that `sourceCQL` has the underlying type of `fromByte179`:
	fromByte179 := sourceCQL.(byte)

	// Call the function that transfers the taint
	// from the parameter `fromByte179` to result `intoIPMask346`
	// (`intoIPMask346` is now tainted).
	intoIPMask346 := net.IPv4Mask(0, 0, 0, fromByte179)

	// Return the tainted `intoIPMask346`:
	return intoIPMask346
}

func TaintStepTest_NetJoinHostPort_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString843` into `intoString846`.

	// Assume that `sourceCQL` has the underlying type of `fromString843`:
	fromString843 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString843` to result `intoString846`
	// (`intoString846` is now tainted).
	intoString846 := net.JoinHostPort(fromString843, "")

	// Return the tainted `intoString846`:
	return intoString846
}

func TaintStepTest_NetJoinHostPort_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString631` into `intoString612`.

	// Assume that `sourceCQL` has the underlying type of `fromString631`:
	fromString631 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString631` to result `intoString612`
	// (`intoString612` is now tainted).
	intoString612 := net.JoinHostPort("", fromString631)

	// Return the tainted `intoString612`:
	return intoString612
}

func TaintStepTest_NetParseCIDR_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString799` into `intoIP390`.

	// Assume that `sourceCQL` has the underlying type of `fromString799`:
	fromString799 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString799` to result `intoIP390`
	// (`intoIP390` is now tainted).
	intoIP390, _, _ := net.ParseCIDR(fromString799)

	// Return the tainted `intoIP390`:
	return intoIP390
}

func TaintStepTest_NetParseCIDR_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString845` into `intoIPNet368`.

	// Assume that `sourceCQL` has the underlying type of `fromString845`:
	fromString845 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString845` to result `intoIPNet368`
	// (`intoIPNet368` is now tainted).
	_, intoIPNet368, _ := net.ParseCIDR(fromString845)

	// Return the tainted `intoIPNet368`:
	return intoIPNet368
}

func TaintStepTest_NetParseIP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString995` into `intoIP509`.

	// Assume that `sourceCQL` has the underlying type of `fromString995`:
	fromString995 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString995` to result `intoIP509`
	// (`intoIP509` is now tainted).
	intoIP509 := net.ParseIP(fromString995)

	// Return the tainted `intoIP509`:
	return intoIP509
}

func TaintStepTest_NetParseMAC_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString545` into `intoHardwareAddr953`.

	// Assume that `sourceCQL` has the underlying type of `fromString545`:
	fromString545 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString545` to result `intoHardwareAddr953`
	// (`intoHardwareAddr953` is now tainted).
	intoHardwareAddr953, _ := net.ParseMAC(fromString545)

	// Return the tainted `intoHardwareAddr953`:
	return intoHardwareAddr953
}

func TaintStepTest_NetPipe_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn347` into `intoConn654`.

	// Assume that `sourceCQL` has the underlying type of `fromConn347`:
	fromConn347 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the result `fromConn347` to result `intoConn654`
	// (extra steps needed)
	intermediateCQL, intoConn654 := net.Pipe()

	// Extra step (`fromConn347` taints `intermediateCQL`, which taints `intoConn654`:
	link(fromConn347, intermediateCQL)

	// Return the tainted `intoConn654`:
	return intoConn654
}

func TaintStepTest_NetPipe_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn961` into `intoConn273`.

	// Assume that `sourceCQL` has the underlying type of `fromConn961`:
	fromConn961 := sourceCQL.(net.Conn)

	// Call the function that transfers the taint
	// from the result `fromConn961` to result `intoConn273`
	// (extra steps needed)
	intoConn273, intermediateCQL := net.Pipe()

	// Extra step (`fromConn961` taints `intermediateCQL`, which taints `intoConn273`:
	link(fromConn961, intermediateCQL)

	// Return the tainted `intoConn273`:
	return intoConn273
}

func TaintStepTest_NetSplitHostPort_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromString453` into `intoString318`.

	// Assume that `sourceCQL` has the underlying type of `fromString453`:
	fromString453 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString453` to result `intoString318`
	// (`intoString318` is now tainted).
	intoString318, _, _ := net.SplitHostPort(fromString453)

	// Return the tainted `intoString318`:
	return intoString318
}

func TaintStepTest_NetSplitHostPort_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromString775` into `intoString670`.

	// Assume that `sourceCQL` has the underlying type of `fromString775`:
	fromString775 := sourceCQL.(string)

	// Call the function that transfers the taint
	// from the parameter `fromString775` to result `intoString670`
	// (`intoString670` is now tainted).
	_, intoString670, _ := net.SplitHostPort(fromString775)

	// Return the tainted `intoString670`:
	return intoString670
}

func TaintStepTest_NetBuffersRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffers455` into `intoByte371`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffers455`:
	fromBuffers455 := sourceCQL.(net.Buffers)

	// Declare `intoByte371` variable:
	var intoByte371 []byte

	// Call the method that transfers the taint
	// from the receiver `fromBuffers455` to the argument `intoByte371`
	// (`intoByte371` is now tainted).
	fromBuffers455.Read(intoByte371)

	// Return the tainted `intoByte371`:
	return intoByte371
}

func TaintStepTest_NetBuffersWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromBuffers648` into `intoWriter499`.

	// Assume that `sourceCQL` has the underlying type of `fromBuffers648`:
	fromBuffers648 := sourceCQL.(net.Buffers)

	// Declare `intoWriter499` variable:
	var intoWriter499 io.Writer

	// Call the method that transfers the taint
	// from the receiver `fromBuffers648` to the argument `intoWriter499`
	// (`intoWriter499` is now tainted).
	fromBuffers648.WriteTo(intoWriter499)

	// Return the tainted `intoWriter499`:
	return intoWriter499
}

func TaintStepTest_NetFlagsString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFlags642` into `intoString949`.

	// Assume that `sourceCQL` has the underlying type of `fromFlags642`:
	fromFlags642 := sourceCQL.(net.Flags)

	// Call the method that transfers the taint
	// from the receiver `fromFlags642` to the result `intoString949`
	// (`intoString949` is now tainted).
	intoString949 := fromFlags642.String()

	// Return the tainted `intoString949`:
	return intoString949
}

func TaintStepTest_NetHardwareAddrString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromHardwareAddr341` into `intoString870`.

	// Assume that `sourceCQL` has the underlying type of `fromHardwareAddr341`:
	fromHardwareAddr341 := sourceCQL.(net.HardwareAddr)

	// Call the method that transfers the taint
	// from the receiver `fromHardwareAddr341` to the result `intoString870`
	// (`intoString870` is now tainted).
	intoString870 := fromHardwareAddr341.String()

	// Return the tainted `intoString870`:
	return intoString870
}

func TaintStepTest_NetIPMarshalText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIP135` into `intoByte738`.

	// Assume that `sourceCQL` has the underlying type of `fromIP135`:
	fromIP135 := sourceCQL.(net.IP)

	// Call the method that transfers the taint
	// from the receiver `fromIP135` to the result `intoByte738`
	// (`intoByte738` is now tainted).
	intoByte738, _ := fromIP135.MarshalText()

	// Return the tainted `intoByte738`:
	return intoByte738
}

func TaintStepTest_NetIPString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIP579` into `intoString752`.

	// Assume that `sourceCQL` has the underlying type of `fromIP579`:
	fromIP579 := sourceCQL.(net.IP)

	// Call the method that transfers the taint
	// from the receiver `fromIP579` to the result `intoString752`
	// (`intoString752` is now tainted).
	intoString752 := fromIP579.String()

	// Return the tainted `intoString752`:
	return intoString752
}

func TaintStepTest_NetIPTo16_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIP545` into `intoIP866`.

	// Assume that `sourceCQL` has the underlying type of `fromIP545`:
	fromIP545 := sourceCQL.(net.IP)

	// Call the method that transfers the taint
	// from the receiver `fromIP545` to the result `intoIP866`
	// (`intoIP866` is now tainted).
	intoIP866 := fromIP545.To16()

	// Return the tainted `intoIP866`:
	return intoIP866
}

func TaintStepTest_NetIPTo4_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIP200` into `intoIP522`.

	// Assume that `sourceCQL` has the underlying type of `fromIP200`:
	fromIP200 := sourceCQL.(net.IP)

	// Call the method that transfers the taint
	// from the receiver `fromIP200` to the result `intoIP522`
	// (`intoIP522` is now tainted).
	intoIP522 := fromIP200.To4()

	// Return the tainted `intoIP522`:
	return intoIP522
}

func TaintStepTest_NetIPUnmarshalText_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte948` into `intoIP722`.

	// Assume that `sourceCQL` has the underlying type of `fromByte948`:
	fromByte948 := sourceCQL.([]byte)

	// Declare `intoIP722` variable:
	var intoIP722 net.IP

	// Call the method that transfers the taint
	// from the parameter `fromByte948` to the receiver `intoIP722`
	// (`intoIP722` is now tainted).
	intoIP722.UnmarshalText(fromByte948)

	// Return the tainted `intoIP722`:
	return intoIP722
}

func TaintStepTest_NetIPAddrString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPAddr430` into `intoString550`.

	// Assume that `sourceCQL` has the underlying type of `fromIPAddr430`:
	fromIPAddr430 := sourceCQL.(net.IPAddr)

	// Call the method that transfers the taint
	// from the receiver `fromIPAddr430` to the result `intoString550`
	// (`intoString550` is now tainted).
	intoString550 := fromIPAddr430.String()

	// Return the tainted `intoString550`:
	return intoString550
}

func TaintStepTest_NetIPConnReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPConn571` into `intoByte930`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn571`:
	fromIPConn571 := sourceCQL.(net.IPConn)

	// Declare `intoByte930` variable:
	var intoByte930 []byte

	// Call the method that transfers the taint
	// from the receiver `fromIPConn571` to the argument `intoByte930`
	// (`intoByte930` is now tainted).
	fromIPConn571.ReadFrom(intoByte930)

	// Return the tainted `intoByte930`:
	return intoByte930
}

func TaintStepTest_NetIPConnReadFromIP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPConn112` into `intoByte207`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn112`:
	fromIPConn112 := sourceCQL.(net.IPConn)

	// Declare `intoByte207` variable:
	var intoByte207 []byte

	// Call the method that transfers the taint
	// from the receiver `fromIPConn112` to the argument `intoByte207`
	// (`intoByte207` is now tainted).
	fromIPConn112.ReadFromIP(intoByte207)

	// Return the tainted `intoByte207`:
	return intoByte207
}

func TaintStepTest_NetIPConnReadMsgIP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPConn957` into `intoByte556`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn957`:
	fromIPConn957 := sourceCQL.(net.IPConn)

	// Declare `intoByte556` variable:
	var intoByte556 []byte

	// Call the method that transfers the taint
	// from the receiver `fromIPConn957` to the argument `intoByte556`
	// (`intoByte556` is now tainted).
	fromIPConn957.ReadMsgIP(intoByte556, nil)

	// Return the tainted `intoByte556`:
	return intoByte556
}

func TaintStepTest_NetIPConnReadMsgIP_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPConn430` into `intoByte414`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn430`:
	fromIPConn430 := sourceCQL.(net.IPConn)

	// Declare `intoByte414` variable:
	var intoByte414 []byte

	// Call the method that transfers the taint
	// from the receiver `fromIPConn430` to the argument `intoByte414`
	// (`intoByte414` is now tainted).
	fromIPConn430.ReadMsgIP(nil, intoByte414)

	// Return the tainted `intoByte414`:
	return intoByte414
}

func TaintStepTest_NetIPConnSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPConn297` into `intoRawConn614`.

	// Assume that `sourceCQL` has the underlying type of `fromIPConn297`:
	fromIPConn297 := sourceCQL.(net.IPConn)

	// Call the method that transfers the taint
	// from the receiver `fromIPConn297` to the result `intoRawConn614`
	// (`intoRawConn614` is now tainted).
	intoRawConn614, _ := fromIPConn297.SyscallConn()

	// Return the tainted `intoRawConn614`:
	return intoRawConn614
}

func TaintStepTest_NetIPConnSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn273` into `intoIPConn760`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn273`:
	fromRawConn273 := sourceCQL.(syscall.RawConn)

	// Declare `intoIPConn760` variable:
	var intoIPConn760 net.IPConn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoIPConn760`:
	intermediateCQL, _ := intoIPConn760.SyscallConn()

	// Extra step (`fromRawConn273` taints `intermediateCQL`, which taints `intoIPConn760`:
	link(fromRawConn273, intermediateCQL)

	// Return the tainted `intoIPConn760`:
	return intoIPConn760
}

func TaintStepTest_NetIPConnWriteMsgIP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte584` into `intoIPConn612`.

	// Assume that `sourceCQL` has the underlying type of `fromByte584`:
	fromByte584 := sourceCQL.([]byte)

	// Declare `intoIPConn612` variable:
	var intoIPConn612 net.IPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte584` to the receiver `intoIPConn612`
	// (`intoIPConn612` is now tainted).
	intoIPConn612.WriteMsgIP(fromByte584, nil, nil)

	// Return the tainted `intoIPConn612`:
	return intoIPConn612
}

func TaintStepTest_NetIPConnWriteMsgIP_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte574` into `intoIPConn783`.

	// Assume that `sourceCQL` has the underlying type of `fromByte574`:
	fromByte574 := sourceCQL.([]byte)

	// Declare `intoIPConn783` variable:
	var intoIPConn783 net.IPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte574` to the receiver `intoIPConn783`
	// (`intoIPConn783` is now tainted).
	intoIPConn783.WriteMsgIP(nil, fromByte574, nil)

	// Return the tainted `intoIPConn783`:
	return intoIPConn783
}

func TaintStepTest_NetIPConnWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte401` into `intoIPConn829`.

	// Assume that `sourceCQL` has the underlying type of `fromByte401`:
	fromByte401 := sourceCQL.([]byte)

	// Declare `intoIPConn829` variable:
	var intoIPConn829 net.IPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte401` to the receiver `intoIPConn829`
	// (`intoIPConn829` is now tainted).
	intoIPConn829.WriteTo(fromByte401, nil)

	// Return the tainted `intoIPConn829`:
	return intoIPConn829
}

func TaintStepTest_NetIPConnWriteToIP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte936` into `intoIPConn711`.

	// Assume that `sourceCQL` has the underlying type of `fromByte936`:
	fromByte936 := sourceCQL.([]byte)

	// Declare `intoIPConn711` variable:
	var intoIPConn711 net.IPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte936` to the receiver `intoIPConn711`
	// (`intoIPConn711` is now tainted).
	intoIPConn711.WriteToIP(fromByte936, nil)

	// Return the tainted `intoIPConn711`:
	return intoIPConn711
}

func TaintStepTest_NetIPMaskString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPMask333` into `intoString555`.

	// Assume that `sourceCQL` has the underlying type of `fromIPMask333`:
	fromIPMask333 := sourceCQL.(net.IPMask)

	// Call the method that transfers the taint
	// from the receiver `fromIPMask333` to the result `intoString555`
	// (`intoString555` is now tainted).
	intoString555 := fromIPMask333.String()

	// Return the tainted `intoString555`:
	return intoString555
}

func TaintStepTest_NetIPNetString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromIPNet382` into `intoString685`.

	// Assume that `sourceCQL` has the underlying type of `fromIPNet382`:
	fromIPNet382 := sourceCQL.(net.IPNet)

	// Call the method that transfers the taint
	// from the receiver `fromIPNet382` to the result `intoString685`
	// (`intoString685` is now tainted).
	intoString685 := fromIPNet382.String()

	// Return the tainted `intoString685`:
	return intoString685
}

func TaintStepTest_NetTCPAddrString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTCPAddr972` into `intoString488`.

	// Assume that `sourceCQL` has the underlying type of `fromTCPAddr972`:
	fromTCPAddr972 := sourceCQL.(net.TCPAddr)

	// Call the method that transfers the taint
	// from the receiver `fromTCPAddr972` to the result `intoString488`
	// (`intoString488` is now tainted).
	intoString488 := fromTCPAddr972.String()

	// Return the tainted `intoString488`:
	return intoString488
}

func TaintStepTest_NetTCPConnReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromReader897` into `intoTCPConn485`.

	// Assume that `sourceCQL` has the underlying type of `fromReader897`:
	fromReader897 := sourceCQL.(io.Reader)

	// Declare `intoTCPConn485` variable:
	var intoTCPConn485 net.TCPConn

	// Call the method that transfers the taint
	// from the parameter `fromReader897` to the receiver `intoTCPConn485`
	// (`intoTCPConn485` is now tainted).
	intoTCPConn485.ReadFrom(fromReader897)

	// Return the tainted `intoTCPConn485`:
	return intoTCPConn485
}

func TaintStepTest_NetTCPConnSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTCPConn601` into `intoRawConn949`.

	// Assume that `sourceCQL` has the underlying type of `fromTCPConn601`:
	fromTCPConn601 := sourceCQL.(net.TCPConn)

	// Call the method that transfers the taint
	// from the receiver `fromTCPConn601` to the result `intoRawConn949`
	// (`intoRawConn949` is now tainted).
	intoRawConn949, _ := fromTCPConn601.SyscallConn()

	// Return the tainted `intoRawConn949`:
	return intoRawConn949
}

func TaintStepTest_NetTCPConnSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn255` into `intoTCPConn679`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn255`:
	fromRawConn255 := sourceCQL.(syscall.RawConn)

	// Declare `intoTCPConn679` variable:
	var intoTCPConn679 net.TCPConn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoTCPConn679`:
	intermediateCQL, _ := intoTCPConn679.SyscallConn()

	// Extra step (`fromRawConn255` taints `intermediateCQL`, which taints `intoTCPConn679`:
	link(fromRawConn255, intermediateCQL)

	// Return the tainted `intoTCPConn679`:
	return intoTCPConn679
}

func TaintStepTest_NetTCPListenerFile_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTCPListener429` into `intoFile937`.

	// Assume that `sourceCQL` has the underlying type of `fromTCPListener429`:
	fromTCPListener429 := sourceCQL.(net.TCPListener)

	// Call the method that transfers the taint
	// from the receiver `fromTCPListener429` to the result `intoFile937`
	// (`intoFile937` is now tainted).
	intoFile937, _ := fromTCPListener429.File()

	// Return the tainted `intoFile937`:
	return intoFile937
}

func TaintStepTest_NetTCPListenerFile_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile562` into `intoTCPListener435`.

	// Assume that `sourceCQL` has the underlying type of `fromFile562`:
	fromFile562 := sourceCQL.(*os.File)

	// Declare `intoTCPListener435` variable:
	var intoTCPListener435 net.TCPListener

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoTCPListener435`:
	intermediateCQL, _ := intoTCPListener435.File()

	// Extra step (`fromFile562` taints `intermediateCQL`, which taints `intoTCPListener435`:
	link(fromFile562, intermediateCQL)

	// Return the tainted `intoTCPListener435`:
	return intoTCPListener435
}

func TaintStepTest_NetTCPListenerSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromTCPListener188` into `intoRawConn801`.

	// Assume that `sourceCQL` has the underlying type of `fromTCPListener188`:
	fromTCPListener188 := sourceCQL.(net.TCPListener)

	// Call the method that transfers the taint
	// from the receiver `fromTCPListener188` to the result `intoRawConn801`
	// (`intoRawConn801` is now tainted).
	intoRawConn801, _ := fromTCPListener188.SyscallConn()

	// Return the tainted `intoRawConn801`:
	return intoRawConn801
}

func TaintStepTest_NetTCPListenerSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn470` into `intoTCPListener452`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn470`:
	fromRawConn470 := sourceCQL.(syscall.RawConn)

	// Declare `intoTCPListener452` variable:
	var intoTCPListener452 net.TCPListener

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoTCPListener452`:
	intermediateCQL, _ := intoTCPListener452.SyscallConn()

	// Extra step (`fromRawConn470` taints `intermediateCQL`, which taints `intoTCPListener452`:
	link(fromRawConn470, intermediateCQL)

	// Return the tainted `intoTCPListener452`:
	return intoTCPListener452
}

func TaintStepTest_NetUDPConnReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUDPConn465` into `intoByte599`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn465`:
	fromUDPConn465 := sourceCQL.(net.UDPConn)

	// Declare `intoByte599` variable:
	var intoByte599 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn465` to the argument `intoByte599`
	// (`intoByte599` is now tainted).
	fromUDPConn465.ReadFrom(intoByte599)

	// Return the tainted `intoByte599`:
	return intoByte599
}

func TaintStepTest_NetUDPConnReadFromUDP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUDPConn476` into `intoByte634`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn476`:
	fromUDPConn476 := sourceCQL.(net.UDPConn)

	// Declare `intoByte634` variable:
	var intoByte634 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn476` to the argument `intoByte634`
	// (`intoByte634` is now tainted).
	fromUDPConn476.ReadFromUDP(intoByte634)

	// Return the tainted `intoByte634`:
	return intoByte634
}

func TaintStepTest_NetUDPConnReadMsgUDP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUDPConn247` into `intoByte825`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn247`:
	fromUDPConn247 := sourceCQL.(net.UDPConn)

	// Declare `intoByte825` variable:
	var intoByte825 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn247` to the argument `intoByte825`
	// (`intoByte825` is now tainted).
	fromUDPConn247.ReadMsgUDP(intoByte825, nil)

	// Return the tainted `intoByte825`:
	return intoByte825
}

func TaintStepTest_NetUDPConnReadMsgUDP_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromUDPConn311` into `intoByte519`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn311`:
	fromUDPConn311 := sourceCQL.(net.UDPConn)

	// Declare `intoByte519` variable:
	var intoByte519 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn311` to the argument `intoByte519`
	// (`intoByte519` is now tainted).
	fromUDPConn311.ReadMsgUDP(nil, intoByte519)

	// Return the tainted `intoByte519`:
	return intoByte519
}

func TaintStepTest_NetUDPConnSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUDPConn278` into `intoRawConn824`.

	// Assume that `sourceCQL` has the underlying type of `fromUDPConn278`:
	fromUDPConn278 := sourceCQL.(net.UDPConn)

	// Call the method that transfers the taint
	// from the receiver `fromUDPConn278` to the result `intoRawConn824`
	// (`intoRawConn824` is now tainted).
	intoRawConn824, _ := fromUDPConn278.SyscallConn()

	// Return the tainted `intoRawConn824`:
	return intoRawConn824
}

func TaintStepTest_NetUDPConnSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn824` into `intoUDPConn983`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn824`:
	fromRawConn824 := sourceCQL.(syscall.RawConn)

	// Declare `intoUDPConn983` variable:
	var intoUDPConn983 net.UDPConn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoUDPConn983`:
	intermediateCQL, _ := intoUDPConn983.SyscallConn()

	// Extra step (`fromRawConn824` taints `intermediateCQL`, which taints `intoUDPConn983`:
	link(fromRawConn824, intermediateCQL)

	// Return the tainted `intoUDPConn983`:
	return intoUDPConn983
}

func TaintStepTest_NetUDPConnWriteMsgUDP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte488` into `intoUDPConn601`.

	// Assume that `sourceCQL` has the underlying type of `fromByte488`:
	fromByte488 := sourceCQL.([]byte)

	// Declare `intoUDPConn601` variable:
	var intoUDPConn601 net.UDPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte488` to the receiver `intoUDPConn601`
	// (`intoUDPConn601` is now tainted).
	intoUDPConn601.WriteMsgUDP(fromByte488, nil, nil)

	// Return the tainted `intoUDPConn601`:
	return intoUDPConn601
}

func TaintStepTest_NetUDPConnWriteMsgUDP_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte213` into `intoUDPConn356`.

	// Assume that `sourceCQL` has the underlying type of `fromByte213`:
	fromByte213 := sourceCQL.([]byte)

	// Declare `intoUDPConn356` variable:
	var intoUDPConn356 net.UDPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte213` to the receiver `intoUDPConn356`
	// (`intoUDPConn356` is now tainted).
	intoUDPConn356.WriteMsgUDP(nil, fromByte213, nil)

	// Return the tainted `intoUDPConn356`:
	return intoUDPConn356
}

func TaintStepTest_NetUDPConnWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte652` into `intoUDPConn757`.

	// Assume that `sourceCQL` has the underlying type of `fromByte652`:
	fromByte652 := sourceCQL.([]byte)

	// Declare `intoUDPConn757` variable:
	var intoUDPConn757 net.UDPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte652` to the receiver `intoUDPConn757`
	// (`intoUDPConn757` is now tainted).
	intoUDPConn757.WriteTo(fromByte652, nil)

	// Return the tainted `intoUDPConn757`:
	return intoUDPConn757
}

func TaintStepTest_NetUDPConnWriteToUDP_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte509` into `intoUDPConn345`.

	// Assume that `sourceCQL` has the underlying type of `fromByte509`:
	fromByte509 := sourceCQL.([]byte)

	// Declare `intoUDPConn345` variable:
	var intoUDPConn345 net.UDPConn

	// Call the method that transfers the taint
	// from the parameter `fromByte509` to the receiver `intoUDPConn345`
	// (`intoUDPConn345` is now tainted).
	intoUDPConn345.WriteToUDP(fromByte509, nil)

	// Return the tainted `intoUDPConn345`:
	return intoUDPConn345
}

func TaintStepTest_NetUnixAddrString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixAddr807` into `intoString876`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixAddr807`:
	fromUnixAddr807 := sourceCQL.(net.UnixAddr)

	// Call the method that transfers the taint
	// from the receiver `fromUnixAddr807` to the result `intoString876`
	// (`intoString876` is now tainted).
	intoString876 := fromUnixAddr807.String()

	// Return the tainted `intoString876`:
	return intoString876
}

func TaintStepTest_NetUnixConnReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixConn784` into `intoByte451`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn784`:
	fromUnixConn784 := sourceCQL.(net.UnixConn)

	// Declare `intoByte451` variable:
	var intoByte451 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn784` to the argument `intoByte451`
	// (`intoByte451` is now tainted).
	fromUnixConn784.ReadFrom(intoByte451)

	// Return the tainted `intoByte451`:
	return intoByte451
}

func TaintStepTest_NetUnixConnReadFromUnix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixConn985` into `intoByte597`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn985`:
	fromUnixConn985 := sourceCQL.(net.UnixConn)

	// Declare `intoByte597` variable:
	var intoByte597 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn985` to the argument `intoByte597`
	// (`intoByte597` is now tainted).
	fromUnixConn985.ReadFromUnix(intoByte597)

	// Return the tainted `intoByte597`:
	return intoByte597
}

func TaintStepTest_NetUnixConnReadMsgUnix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixConn160` into `intoByte315`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn160`:
	fromUnixConn160 := sourceCQL.(net.UnixConn)

	// Declare `intoByte315` variable:
	var intoByte315 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn160` to the argument `intoByte315`
	// (`intoByte315` is now tainted).
	fromUnixConn160.ReadMsgUnix(intoByte315, nil)

	// Return the tainted `intoByte315`:
	return intoByte315
}

func TaintStepTest_NetUnixConnReadMsgUnix_B0I0O1(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixConn792` into `intoByte433`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn792`:
	fromUnixConn792 := sourceCQL.(net.UnixConn)

	// Declare `intoByte433` variable:
	var intoByte433 []byte

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn792` to the argument `intoByte433`
	// (`intoByte433` is now tainted).
	fromUnixConn792.ReadMsgUnix(nil, intoByte433)

	// Return the tainted `intoByte433`:
	return intoByte433
}

func TaintStepTest_NetUnixConnSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixConn465` into `intoRawConn786`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixConn465`:
	fromUnixConn465 := sourceCQL.(net.UnixConn)

	// Call the method that transfers the taint
	// from the receiver `fromUnixConn465` to the result `intoRawConn786`
	// (`intoRawConn786` is now tainted).
	intoRawConn786, _ := fromUnixConn465.SyscallConn()

	// Return the tainted `intoRawConn786`:
	return intoRawConn786
}

func TaintStepTest_NetUnixConnSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn637` into `intoUnixConn813`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn637`:
	fromRawConn637 := sourceCQL.(syscall.RawConn)

	// Declare `intoUnixConn813` variable:
	var intoUnixConn813 net.UnixConn

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoUnixConn813`:
	intermediateCQL, _ := intoUnixConn813.SyscallConn()

	// Extra step (`fromRawConn637` taints `intermediateCQL`, which taints `intoUnixConn813`:
	link(fromRawConn637, intermediateCQL)

	// Return the tainted `intoUnixConn813`:
	return intoUnixConn813
}

func TaintStepTest_NetUnixConnWriteMsgUnix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte542` into `intoUnixConn720`.

	// Assume that `sourceCQL` has the underlying type of `fromByte542`:
	fromByte542 := sourceCQL.([]byte)

	// Declare `intoUnixConn720` variable:
	var intoUnixConn720 net.UnixConn

	// Call the method that transfers the taint
	// from the parameter `fromByte542` to the receiver `intoUnixConn720`
	// (`intoUnixConn720` is now tainted).
	intoUnixConn720.WriteMsgUnix(fromByte542, nil, nil)

	// Return the tainted `intoUnixConn720`:
	return intoUnixConn720
}

func TaintStepTest_NetUnixConnWriteMsgUnix_B0I1O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte534` into `intoUnixConn170`.

	// Assume that `sourceCQL` has the underlying type of `fromByte534`:
	fromByte534 := sourceCQL.([]byte)

	// Declare `intoUnixConn170` variable:
	var intoUnixConn170 net.UnixConn

	// Call the method that transfers the taint
	// from the parameter `fromByte534` to the receiver `intoUnixConn170`
	// (`intoUnixConn170` is now tainted).
	intoUnixConn170.WriteMsgUnix(nil, fromByte534, nil)

	// Return the tainted `intoUnixConn170`:
	return intoUnixConn170
}

func TaintStepTest_NetUnixConnWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte668` into `intoUnixConn949`.

	// Assume that `sourceCQL` has the underlying type of `fromByte668`:
	fromByte668 := sourceCQL.([]byte)

	// Declare `intoUnixConn949` variable:
	var intoUnixConn949 net.UnixConn

	// Call the method that transfers the taint
	// from the parameter `fromByte668` to the receiver `intoUnixConn949`
	// (`intoUnixConn949` is now tainted).
	intoUnixConn949.WriteTo(fromByte668, nil)

	// Return the tainted `intoUnixConn949`:
	return intoUnixConn949
}

func TaintStepTest_NetUnixConnWriteToUnix_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte327` into `intoUnixConn172`.

	// Assume that `sourceCQL` has the underlying type of `fromByte327`:
	fromByte327 := sourceCQL.([]byte)

	// Declare `intoUnixConn172` variable:
	var intoUnixConn172 net.UnixConn

	// Call the method that transfers the taint
	// from the parameter `fromByte327` to the receiver `intoUnixConn172`
	// (`intoUnixConn172` is now tainted).
	intoUnixConn172.WriteToUnix(fromByte327, nil)

	// Return the tainted `intoUnixConn172`:
	return intoUnixConn172
}

func TaintStepTest_NetUnixListenerFile_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixListener449` into `intoFile678`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixListener449`:
	fromUnixListener449 := sourceCQL.(net.UnixListener)

	// Call the method that transfers the taint
	// from the receiver `fromUnixListener449` to the result `intoFile678`
	// (`intoFile678` is now tainted).
	intoFile678, _ := fromUnixListener449.File()

	// Return the tainted `intoFile678`:
	return intoFile678
}

func TaintStepTest_NetUnixListenerFile_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromFile944` into `intoUnixListener683`.

	// Assume that `sourceCQL` has the underlying type of `fromFile944`:
	fromFile944 := sourceCQL.(*os.File)

	// Declare `intoUnixListener683` variable:
	var intoUnixListener683 net.UnixListener

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoUnixListener683`:
	intermediateCQL, _ := intoUnixListener683.File()

	// Extra step (`fromFile944` taints `intermediateCQL`, which taints `intoUnixListener683`:
	link(fromFile944, intermediateCQL)

	// Return the tainted `intoUnixListener683`:
	return intoUnixListener683
}

func TaintStepTest_NetUnixListenerSyscallConn_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromUnixListener818` into `intoRawConn202`.

	// Assume that `sourceCQL` has the underlying type of `fromUnixListener818`:
	fromUnixListener818 := sourceCQL.(net.UnixListener)

	// Call the method that transfers the taint
	// from the receiver `fromUnixListener818` to the result `intoRawConn202`
	// (`intoRawConn202` is now tainted).
	intoRawConn202, _ := fromUnixListener818.SyscallConn()

	// Return the tainted `intoRawConn202`:
	return intoRawConn202
}

func TaintStepTest_NetUnixListenerSyscallConn_B1I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromRawConn791` into `intoUnixListener744`.

	// Assume that `sourceCQL` has the underlying type of `fromRawConn791`:
	fromRawConn791 := sourceCQL.(syscall.RawConn)

	// Declare `intoUnixListener744` variable:
	var intoUnixListener744 net.UnixListener

	// Call the method that will transfer the taint
	// from the result `intermediateCQL` to receiver `intoUnixListener744`:
	intermediateCQL, _ := intoUnixListener744.SyscallConn()

	// Extra step (`fromRawConn791` taints `intermediateCQL`, which taints `intoUnixListener744`:
	link(fromRawConn791, intermediateCQL)

	// Return the tainted `intoUnixListener744`:
	return intoUnixListener744
}

func TaintStepTest_NetConnRead_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromConn880` into `intoByte919`.

	// Assume that `sourceCQL` has the underlying type of `fromConn880`:
	fromConn880 := sourceCQL.(net.Conn)

	// Declare `intoByte919` variable:
	var intoByte919 []byte

	// Call the method that transfers the taint
	// from the receiver `fromConn880` to the argument `intoByte919`
	// (`intoByte919` is now tainted).
	fromConn880.Read(intoByte919)

	// Return the tainted `intoByte919`:
	return intoByte919
}

func TaintStepTest_NetPacketConnReadFrom_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromPacketConn641` into `intoByte166`.

	// Assume that `sourceCQL` has the underlying type of `fromPacketConn641`:
	fromPacketConn641 := sourceCQL.(net.PacketConn)

	// Declare `intoByte166` variable:
	var intoByte166 []byte

	// Call the method that transfers the taint
	// from the receiver `fromPacketConn641` to the argument `intoByte166`
	// (`intoByte166` is now tainted).
	fromPacketConn641.ReadFrom(intoByte166)

	// Return the tainted `intoByte166`:
	return intoByte166
}

func TaintStepTest_NetAddrString_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromAddr262` into `intoString540`.

	// Assume that `sourceCQL` has the underlying type of `fromAddr262`:
	fromAddr262 := sourceCQL.(net.Addr)

	// Call the method that transfers the taint
	// from the receiver `fromAddr262` to the result `intoString540`
	// (`intoString540` is now tainted).
	intoString540 := fromAddr262.String()

	// Return the tainted `intoString540`:
	return intoString540
}

func TaintStepTest_NetConnWrite_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte547` into `intoConn566`.

	// Assume that `sourceCQL` has the underlying type of `fromByte547`:
	fromByte547 := sourceCQL.([]byte)

	// Declare `intoConn566` variable:
	var intoConn566 net.Conn

	// Call the method that transfers the taint
	// from the parameter `fromByte547` to the receiver `intoConn566`
	// (`intoConn566` is now tainted).
	intoConn566.Write(fromByte547)

	// Return the tainted `intoConn566`:
	return intoConn566
}

func TaintStepTest_NetPacketConnWriteTo_B0I0O0(sourceCQL interface{}) interface{} {
	// The flow is from `fromByte602` into `intoPacketConn314`.

	// Assume that `sourceCQL` has the underlying type of `fromByte602`:
	fromByte602 := sourceCQL.([]byte)

	// Declare `intoPacketConn314` variable:
	var intoPacketConn314 net.PacketConn

	// Call the method that transfers the taint
	// from the parameter `fromByte602` to the receiver `intoPacketConn314`
	// (`intoPacketConn314` is now tainted).
	intoPacketConn314.WriteTo(fromByte602, nil)

	// Return the tainted `intoPacketConn314`:
	return intoPacketConn314
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
