/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `net` package. */
module NetTaintTracking {
  private class FileConn extends TaintTracking::FunctionModel {
    // signature: func FileConn(f *os.File) (c Conn, err error)
    FileConn() { hasQualifiedName("net", "FileConn") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
      or
      inp.isResult(0) and outp.isParameter(0)
    }
  }

  private class FileListener extends TaintTracking::FunctionModel {
    // signature: func FileListener(f *os.File) (ln Listener, err error)
    FileListener() { hasQualifiedName("net", "FileListener") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
      or
      inp.isResult(0) and outp.isParameter(0)
    }
  }

  private class FilePacketConn extends TaintTracking::FunctionModel {
    // signature: func FilePacketConn(f *os.File) (c PacketConn, err error)
    FilePacketConn() { hasQualifiedName("net", "FilePacketConn") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult(0)
      or
      inp.isResult(0) and outp.isParameter(0)
    }
  }

  private class IPv4 extends TaintTracking::FunctionModel {
    // signature: func IPv4(a byte, b byte, c byte, d byte) IP
    IPv4() { hasQualifiedName("net", "IPv4") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class IPv4Mask extends TaintTracking::FunctionModel {
    // signature: func IPv4Mask(a byte, b byte, c byte, d byte) IPMask
    IPv4Mask() { hasQualifiedName("net", "IPv4Mask") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class JoinHostPort extends TaintTracking::FunctionModel {
    // signature: func JoinHostPort(host string, port string) string
    JoinHostPort() { hasQualifiedName("net", "JoinHostPort") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class ParseCIDR extends TaintTracking::FunctionModel {
    // signature: func ParseCIDR(s string) (IP, *IPNet, error)
    ParseCIDR() { hasQualifiedName("net", "ParseCIDR") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult([0, 1]))
    }
  }

  private class ParseIP extends TaintTracking::FunctionModel {
    // signature: func ParseIP(s string) IP
    ParseIP() { hasQualifiedName("net", "ParseIP") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class ParseMAC extends TaintTracking::FunctionModel {
    // signature: func ParseMAC(s string) (hw HardwareAddr, err error)
    ParseMAC() { hasQualifiedName("net", "ParseMAC") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class Pipe extends TaintTracking::FunctionModel {
    // signature: func Pipe() (Conn, Conn)
    Pipe() { hasQualifiedName("net", "Pipe") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isResult(0) and outp.isResult(1)
      or
      inp.isResult(1) and outp.isResult(0)
    }
  }

  private class SplitHostPort extends TaintTracking::FunctionModel {
    // signature: func SplitHostPort(hostport string) (host string, port string, err error)
    SplitHostPort() { hasQualifiedName("net", "SplitHostPort") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult([0, 1]))
    }
  }

  private class BuffersRead extends TaintTracking::FunctionModel, Method {
    // signature: func (*Buffers).Read(p []byte) (n int, err error)
    BuffersRead() { this.(Method).hasQualifiedName("net", "Buffers", "Read") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class BuffersWriteTo extends TaintTracking::FunctionModel, Method {
    // signature: func (*Buffers).WriteTo(w io.Writer) (n int64, err error)
    BuffersWriteTo() { this.(Method).hasQualifiedName("net", "Buffers", "WriteTo") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class FlagsString extends TaintTracking::FunctionModel, Method {
    // signature: func (Flags).String() string
    FlagsString() { this.(Method).hasQualifiedName("net", "Flags", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class HardwareAddrString extends TaintTracking::FunctionModel, Method {
    // signature: func (HardwareAddr).String() string
    HardwareAddrString() { this.(Method).hasQualifiedName("net", "HardwareAddr", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class IPMarshalText extends TaintTracking::FunctionModel, Method {
    // signature: func (IP).MarshalText() ([]byte, error)
    IPMarshalText() { this.(Method).hasQualifiedName("net", "IP", "MarshalText") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult(0))
    }
  }

  private class IPString extends TaintTracking::FunctionModel, Method {
    // signature: func (IP).String() string
    IPString() { this.(Method).hasQualifiedName("net", "IP", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class IPTo16 extends TaintTracking::FunctionModel, Method {
    // signature: func (IP).To16() IP
    IPTo16() { this.(Method).hasQualifiedName("net", "IP", "To16") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class IPTo4 extends TaintTracking::FunctionModel, Method {
    // signature: func (IP).To4() IP
    IPTo4() { this.(Method).hasQualifiedName("net", "IP", "To4") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class IPUnmarshalText extends TaintTracking::FunctionModel, Method {
    // signature: func (*IP).UnmarshalText(text []byte) error
    IPUnmarshalText() { this.(Method).hasQualifiedName("net", "IP", "UnmarshalText") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class IPAddrString extends TaintTracking::FunctionModel, Method {
    // signature: func (*IPAddr).String() string
    IPAddrString() { this.(Method).hasQualifiedName("net", "IPAddr", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class IPConnReadFrom extends TaintTracking::FunctionModel, Method {
    // signature: func (*IPConn).ReadFrom(b []byte) (int, Addr, error)
    IPConnReadFrom() { this.(Method).hasQualifiedName("net", "IPConn", "ReadFrom") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class IPConnReadFromIP extends TaintTracking::FunctionModel, Method {
    // signature: func (*IPConn).ReadFromIP(b []byte) (int, *IPAddr, error)
    IPConnReadFromIP() { this.(Method).hasQualifiedName("net", "IPConn", "ReadFromIP") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class IPConnReadMsgIP extends TaintTracking::FunctionModel, Method {
    // signature: func (*IPConn).ReadMsgIP(b []byte, oob []byte) (n int, oobn int, flags int, addr *IPAddr, err error)
    IPConnReadMsgIP() { this.(Method).hasQualifiedName("net", "IPConn", "ReadMsgIP") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(_))
    }
  }

  private class IPConnSyscallConn extends TaintTracking::FunctionModel, Method {
    // signature: func (*IPConn).SyscallConn() (syscall.RawConn, error)
    IPConnSyscallConn() { this.(Method).hasQualifiedName("net", "IPConn", "SyscallConn") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
      or
      inp.isResult(0) and outp.isReceiver()
    }
  }

  private class IPConnWriteMsgIP extends TaintTracking::FunctionModel, Method {
    // signature: func (*IPConn).WriteMsgIP(b []byte, oob []byte, addr *IPAddr) (n int, oobn int, err error)
    IPConnWriteMsgIP() { this.(Method).hasQualifiedName("net", "IPConn", "WriteMsgIP") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter([0, 1]) and outp.isReceiver())
    }
  }

  private class IPConnWriteTo extends TaintTracking::FunctionModel, Method {
    // signature: func (*IPConn).WriteTo(b []byte, addr Addr) (int, error)
    IPConnWriteTo() { this.(Method).hasQualifiedName("net", "IPConn", "WriteTo") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class IPConnWriteToIP extends TaintTracking::FunctionModel, Method {
    // signature: func (*IPConn).WriteToIP(b []byte, addr *IPAddr) (int, error)
    IPConnWriteToIP() { this.(Method).hasQualifiedName("net", "IPConn", "WriteToIP") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class IPMaskString extends TaintTracking::FunctionModel, Method {
    // signature: func (IPMask).String() string
    IPMaskString() { this.(Method).hasQualifiedName("net", "IPMask", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class IPNetString extends TaintTracking::FunctionModel, Method {
    // signature: func (*IPNet).String() string
    IPNetString() { this.(Method).hasQualifiedName("net", "IPNet", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class TCPAddrString extends TaintTracking::FunctionModel, Method {
    // signature: func (*TCPAddr).String() string
    TCPAddrString() { this.(Method).hasQualifiedName("net", "TCPAddr", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class TCPConnReadFrom extends TaintTracking::FunctionModel, Method {
    // signature: func (*TCPConn).ReadFrom(r io.Reader) (int64, error)
    TCPConnReadFrom() { this.(Method).hasQualifiedName("net", "TCPConn", "ReadFrom") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class TCPConnSyscallConn extends TaintTracking::FunctionModel, Method {
    // signature: func (*TCPConn).SyscallConn() (syscall.RawConn, error)
    TCPConnSyscallConn() { this.(Method).hasQualifiedName("net", "TCPConn", "SyscallConn") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
      or
      inp.isResult(0) and outp.isReceiver()
    }
  }

  private class TCPListenerFile extends TaintTracking::FunctionModel, Method {
    // signature: func (*TCPListener).File() (f *os.File, err error)
    TCPListenerFile() { this.(Method).hasQualifiedName("net", "TCPListener", "File") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
      or
      inp.isResult(0) and outp.isReceiver()
    }
  }

  private class TCPListenerSyscallConn extends TaintTracking::FunctionModel, Method {
    // signature: func (*TCPListener).SyscallConn() (syscall.RawConn, error)
    TCPListenerSyscallConn() { this.(Method).hasQualifiedName("net", "TCPListener", "SyscallConn") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
      or
      inp.isResult(0) and outp.isReceiver()
    }
  }

  private class UDPConnReadFrom extends TaintTracking::FunctionModel, Method {
    // signature: func (*UDPConn).ReadFrom(b []byte) (int, Addr, error)
    UDPConnReadFrom() { this.(Method).hasQualifiedName("net", "UDPConn", "ReadFrom") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class UDPConnReadFromUDP extends TaintTracking::FunctionModel, Method {
    // signature: func (*UDPConn).ReadFromUDP(b []byte) (int, *UDPAddr, error)
    UDPConnReadFromUDP() { this.(Method).hasQualifiedName("net", "UDPConn", "ReadFromUDP") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class UDPConnReadMsgUDP extends TaintTracking::FunctionModel, Method {
    // signature: func (*UDPConn).ReadMsgUDP(b []byte, oob []byte) (n int, oobn int, flags int, addr *UDPAddr, err error)
    UDPConnReadMsgUDP() { this.(Method).hasQualifiedName("net", "UDPConn", "ReadMsgUDP") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(_))
    }
  }

  private class UDPConnSyscallConn extends TaintTracking::FunctionModel, Method {
    // signature: func (*UDPConn).SyscallConn() (syscall.RawConn, error)
    UDPConnSyscallConn() { this.(Method).hasQualifiedName("net", "UDPConn", "SyscallConn") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
      or
      inp.isResult(0) and outp.isReceiver()
    }
  }

  private class UDPConnWriteMsgUDP extends TaintTracking::FunctionModel, Method {
    // signature: func (*UDPConn).WriteMsgUDP(b []byte, oob []byte, addr *UDPAddr) (n int, oobn int, err error)
    UDPConnWriteMsgUDP() { this.(Method).hasQualifiedName("net", "UDPConn", "WriteMsgUDP") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter([0, 1]) and outp.isReceiver())
    }
  }

  private class UDPConnWriteTo extends TaintTracking::FunctionModel, Method {
    // signature: func (*UDPConn).WriteTo(b []byte, addr Addr) (int, error)
    UDPConnWriteTo() { this.(Method).hasQualifiedName("net", "UDPConn", "WriteTo") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class UDPConnWriteToUDP extends TaintTracking::FunctionModel, Method {
    // signature: func (*UDPConn).WriteToUDP(b []byte, addr *UDPAddr) (int, error)
    UDPConnWriteToUDP() { this.(Method).hasQualifiedName("net", "UDPConn", "WriteToUDP") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class UnixAddrString extends TaintTracking::FunctionModel, Method {
    // signature: func (*UnixAddr).String() string
    UnixAddrString() { this.(Method).hasQualifiedName("net", "UnixAddr", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class UnixConnReadFrom extends TaintTracking::FunctionModel, Method {
    // signature: func (*UnixConn).ReadFrom(b []byte) (int, Addr, error)
    UnixConnReadFrom() { this.(Method).hasQualifiedName("net", "UnixConn", "ReadFrom") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class UnixConnReadFromUnix extends TaintTracking::FunctionModel, Method {
    // signature: func (*UnixConn).ReadFromUnix(b []byte) (int, *UnixAddr, error)
    UnixConnReadFromUnix() { this.(Method).hasQualifiedName("net", "UnixConn", "ReadFromUnix") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class UnixConnReadMsgUnix extends TaintTracking::FunctionModel, Method {
    // signature: func (*UnixConn).ReadMsgUnix(b []byte, oob []byte) (n int, oobn int, flags int, addr *UnixAddr, err error)
    UnixConnReadMsgUnix() { this.(Method).hasQualifiedName("net", "UnixConn", "ReadMsgUnix") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(_))
    }
  }

  private class UnixConnSyscallConn extends TaintTracking::FunctionModel, Method {
    // signature: func (*UnixConn).SyscallConn() (syscall.RawConn, error)
    UnixConnSyscallConn() { this.(Method).hasQualifiedName("net", "UnixConn", "SyscallConn") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
      or
      inp.isResult(0) and outp.isReceiver()
    }
  }

  private class UnixConnWriteMsgUnix extends TaintTracking::FunctionModel, Method {
    // signature: func (*UnixConn).WriteMsgUnix(b []byte, oob []byte, addr *UnixAddr) (n int, oobn int, err error)
    UnixConnWriteMsgUnix() { this.(Method).hasQualifiedName("net", "UnixConn", "WriteMsgUnix") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter([0, 1]) and outp.isReceiver())
    }
  }

  private class UnixConnWriteTo extends TaintTracking::FunctionModel, Method {
    // signature: func (*UnixConn).WriteTo(b []byte, addr Addr) (int, error)
    UnixConnWriteTo() { this.(Method).hasQualifiedName("net", "UnixConn", "WriteTo") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class UnixConnWriteToUnix extends TaintTracking::FunctionModel, Method {
    // signature: func (*UnixConn).WriteToUnix(b []byte, addr *UnixAddr) (int, error)
    UnixConnWriteToUnix() { this.(Method).hasQualifiedName("net", "UnixConn", "WriteToUnix") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class UnixListenerFile extends TaintTracking::FunctionModel, Method {
    // signature: func (*UnixListener).File() (f *os.File, err error)
    UnixListenerFile() { this.(Method).hasQualifiedName("net", "UnixListener", "File") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
      or
      inp.isResult(0) and outp.isReceiver()
    }
  }

  private class UnixListenerSyscallConn extends TaintTracking::FunctionModel, Method {
    // signature: func (*UnixListener).SyscallConn() (syscall.RawConn, error)
    UnixListenerSyscallConn() {
      this.(Method).hasQualifiedName("net", "UnixListener", "SyscallConn")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
      or
      inp.isResult(0) and outp.isReceiver()
    }
  }

  private class ConnRead extends TaintTracking::FunctionModel, Method {
    // signature: func (Conn).Read(b []byte) (n int, err error)
    ConnRead() { this.implements("net", "Conn", "Read") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class PacketConnReadFrom extends TaintTracking::FunctionModel, Method {
    // signature: func (PacketConn).ReadFrom(p []byte) (n int, addr Addr, err error)
    PacketConnReadFrom() { this.implements("net", "PacketConn", "ReadFrom") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class AddrString extends TaintTracking::FunctionModel, Method {
    // signature: func (Addr).String() string
    AddrString() { this.implements("net", "Addr", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class ConnWrite extends TaintTracking::FunctionModel, Method {
    // signature: func (Conn).Write(b []byte) (n int, err error)
    ConnWrite() { this.implements("net", "Conn", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class PacketConnWriteTo extends TaintTracking::FunctionModel, Method {
    // signature: func (PacketConn).WriteTo(p []byte, addr Addr) (n int, err error)
    PacketConnWriteTo() { this.implements("net", "PacketConn", "WriteTo") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }
}
