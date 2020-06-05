/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `bytes` package. */
module BytesTaintTracking {
  private class Fields extends TaintTracking::FunctionModel {
    // signature: func Fields(s []byte) [][]byte
    Fields() { hasQualifiedName("bytes", "Fields") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class FieldsFunc extends TaintTracking::FunctionModel {
    // signature: func FieldsFunc(s []byte, f func(rune) bool) [][]byte
    FieldsFunc() { hasQualifiedName("bytes", "FieldsFunc") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class Join extends TaintTracking::FunctionModel {
    // signature: func Join(s [][]byte, sep []byte) []byte
    Join() { hasQualifiedName("bytes", "Join") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter([0, 1]) and outp.isResult())
    }
  }

  private class Map extends TaintTracking::FunctionModel {
    // signature: func Map(mapping func(r rune) rune, s []byte) []byte
    Map() { hasQualifiedName("bytes", "Map") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isResult()
    }
  }

  private class NewBuffer extends TaintTracking::FunctionModel {
    // signature: func NewBuffer(buf []byte) *Buffer
    NewBuffer() { hasQualifiedName("bytes", "NewBuffer") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class NewBufferString extends TaintTracking::FunctionModel {
    // signature: func NewBufferString(s string) *Buffer
    NewBufferString() { hasQualifiedName("bytes", "NewBufferString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class NewReader extends TaintTracking::FunctionModel {
    // signature: func NewReader(b []byte) *Reader
    NewReader() { hasQualifiedName("bytes", "NewReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class Repeat extends TaintTracking::FunctionModel {
    // signature: func Repeat(b []byte, count int) []byte
    Repeat() { hasQualifiedName("bytes", "Repeat") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class Replace extends TaintTracking::FunctionModel {
    // signature: func Replace(s []byte, old []byte, new []byte, n int) []byte
    Replace() { hasQualifiedName("bytes", "Replace") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter([0, 2]) and outp.isResult()
    }
  }

  private class ReplaceAll extends TaintTracking::FunctionModel {
    // signature: func ReplaceAll(s []byte, old []byte, new []byte) []byte
    ReplaceAll() { hasQualifiedName("bytes", "ReplaceAll") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter([0, 2]) and outp.isResult()
    }
  }

  private class Runes extends TaintTracking::FunctionModel {
    // signature: func Runes(s []byte) []rune
    Runes() { hasQualifiedName("bytes", "Runes") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class Split extends TaintTracking::FunctionModel {
    // signature: func Split(s []byte, sep []byte) [][]byte
    Split() { hasQualifiedName("bytes", "Split") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class SplitAfter extends TaintTracking::FunctionModel {
    // signature: func SplitAfter(s []byte, sep []byte) [][]byte
    SplitAfter() { hasQualifiedName("bytes", "SplitAfter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class SplitAfterN extends TaintTracking::FunctionModel {
    // signature: func SplitAfterN(s []byte, sep []byte, n int) [][]byte
    SplitAfterN() { hasQualifiedName("bytes", "SplitAfterN") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class SplitN extends TaintTracking::FunctionModel {
    // signature: func SplitN(s []byte, sep []byte, n int) [][]byte
    SplitN() { hasQualifiedName("bytes", "SplitN") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class Title extends TaintTracking::FunctionModel {
    // signature: func Title(s []byte) []byte
    Title() { hasQualifiedName("bytes", "Title") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class ToLower extends TaintTracking::FunctionModel {
    // signature: func ToLower(s []byte) []byte
    ToLower() { hasQualifiedName("bytes", "ToLower") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class ToLowerSpecial extends TaintTracking::FunctionModel {
    // signature: func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte
    ToLowerSpecial() { hasQualifiedName("bytes", "ToLowerSpecial") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isResult()
    }
  }

  private class ToTitle extends TaintTracking::FunctionModel {
    // signature: func ToTitle(s []byte) []byte
    ToTitle() { hasQualifiedName("bytes", "ToTitle") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class ToTitleSpecial extends TaintTracking::FunctionModel {
    // signature: func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte
    ToTitleSpecial() { hasQualifiedName("bytes", "ToTitleSpecial") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isResult()
    }
  }

  private class ToUpper extends TaintTracking::FunctionModel {
    // signature: func ToUpper(s []byte) []byte
    ToUpper() { hasQualifiedName("bytes", "ToUpper") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class ToUpperSpecial extends TaintTracking::FunctionModel {
    // signature: func ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte
    ToUpperSpecial() { hasQualifiedName("bytes", "ToUpperSpecial") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(1) and outp.isResult()
    }
  }

  private class ToValidUTF8 extends TaintTracking::FunctionModel {
    // signature: func ToValidUTF8(s []byte, replacement []byte) []byte
    ToValidUTF8() { hasQualifiedName("bytes", "ToValidUTF8") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter([0, 1]) and outp.isResult()
    }
  }

  private class Trim extends TaintTracking::FunctionModel {
    // signature: func Trim(s []byte, cutset string) []byte
    Trim() { hasQualifiedName("bytes", "Trim") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class TrimFunc extends TaintTracking::FunctionModel {
    // signature: func TrimFunc(s []byte, f func(r rune) bool) []byte
    TrimFunc() { hasQualifiedName("bytes", "TrimFunc") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class TrimLeft extends TaintTracking::FunctionModel {
    // signature: func TrimLeft(s []byte, cutset string) []byte
    TrimLeft() { hasQualifiedName("bytes", "TrimLeft") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class TrimLeftFunc extends TaintTracking::FunctionModel {
    // signature: func TrimLeftFunc(s []byte, f func(r rune) bool) []byte
    TrimLeftFunc() { hasQualifiedName("bytes", "TrimLeftFunc") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class TrimPrefix extends TaintTracking::FunctionModel {
    // signature: func TrimPrefix(s []byte, prefix []byte) []byte
    TrimPrefix() { hasQualifiedName("bytes", "TrimPrefix") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class TrimRight extends TaintTracking::FunctionModel {
    // signature: func TrimRight(s []byte, cutset string) []byte
    TrimRight() { hasQualifiedName("bytes", "TrimRight") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class TrimRightFunc extends TaintTracking::FunctionModel {
    // signature: func TrimRightFunc(s []byte, f func(r rune) bool) []byte
    TrimRightFunc() { hasQualifiedName("bytes", "TrimRightFunc") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class TrimSpace extends TaintTracking::FunctionModel {
    // signature: func TrimSpace(s []byte) []byte
    TrimSpace() { hasQualifiedName("bytes", "TrimSpace") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class TrimSuffix extends TaintTracking::FunctionModel {
    // signature: func TrimSuffix(s []byte, suffix []byte) []byte
    TrimSuffix() { hasQualifiedName("bytes", "TrimSuffix") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }

  private class BufferBytes extends TaintTracking::FunctionModel, Method {
    // signature: func (*Buffer).Bytes() []byte
    BufferBytes() { this.(Method).hasQualifiedName("bytes", "Buffer", "Bytes") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class BufferNext extends TaintTracking::FunctionModel, Method {
    // signature: func (*Buffer).Next(n int) []byte
    BufferNext() { this.(Method).hasQualifiedName("bytes", "Buffer", "Next") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class BufferReadByte extends TaintTracking::FunctionModel, Method {
    // signature: func (*Buffer).ReadByte() (byte, error)
    BufferReadByte() { this.(Method).hasQualifiedName("bytes", "Buffer", "ReadByte") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class BufferReadBytes extends TaintTracking::FunctionModel, Method {
    // signature: func (*Buffer).ReadBytes(delim byte) (line []byte, err error)
    BufferReadBytes() { this.(Method).hasQualifiedName("bytes", "Buffer", "ReadBytes") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class BufferReadFrom extends TaintTracking::FunctionModel, Method {
    // signature: func (*Buffer).ReadFrom(r io.Reader) (n int64, err error)
    BufferReadFrom() { this.(Method).hasQualifiedName("bytes", "Buffer", "ReadFrom") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class BufferReadRune extends TaintTracking::FunctionModel, Method {
    // signature: func (*Buffer).ReadRune() (r rune, size int, err error)
    BufferReadRune() { this.(Method).hasQualifiedName("bytes", "Buffer", "ReadRune") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class BufferReadString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Buffer).ReadString(delim byte) (line string, err error)
    BufferReadString() { this.(Method).hasQualifiedName("bytes", "Buffer", "ReadString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class BufferString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Buffer).String() string
    BufferString() { this.(Method).hasQualifiedName("bytes", "Buffer", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class BufferWriteByte extends TaintTracking::FunctionModel, Method {
    // signature: func (*Buffer).WriteByte(c byte) error
    BufferWriteByte() { this.(Method).hasQualifiedName("bytes", "Buffer", "WriteByte") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class BufferWriteRune extends TaintTracking::FunctionModel, Method {
    // signature: func (*Buffer).WriteRune(r rune) (n int, err error)
    BufferWriteRune() { this.(Method).hasQualifiedName("bytes", "Buffer", "WriteRune") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class BufferWriteString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Buffer).WriteString(s string) (n int, err error)
    BufferWriteString() { this.(Method).hasQualifiedName("bytes", "Buffer", "WriteString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class BufferWriteTo extends TaintTracking::FunctionModel, Method {
    // signature: func (*Buffer).WriteTo(w io.Writer) (n int64, err error)
    BufferWriteTo() { this.(Method).hasQualifiedName("bytes", "Buffer", "WriteTo") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isParameter(0)
    }
  }

  private class ReaderReadByte extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadByte() (byte, error)
    ReaderReadByte() { this.(Method).hasQualifiedName("bytes", "Reader", "ReadByte") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class ReaderReadRune extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadRune() (ch rune, size int, err error)
    ReaderReadRune() { this.(Method).hasQualifiedName("bytes", "Reader", "ReadRune") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult(0)
    }
  }

  private class ReaderReset extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).Reset(b []byte)
    ReaderReset() { this.(Method).hasQualifiedName("bytes", "Reader", "Reset") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ReaderWriteTo extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).WriteTo(w io.Writer) (n int64, err error)
    ReaderWriteTo() { this.(Method).hasQualifiedName("bytes", "Reader", "WriteTo") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isParameter(0)
    }
  }
}
