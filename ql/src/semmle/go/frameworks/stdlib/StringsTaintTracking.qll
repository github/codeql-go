/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `strings` package. */
module StringsTaintTracking {
  private class Fields extends TaintTracking::FunctionModel {
    // signature: func Fields(s string) []string
    Fields() { hasQualifiedName("strings", "Fields") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class FieldsFunc extends TaintTracking::FunctionModel {
    // signature: func FieldsFunc(s string, f func(rune) bool) []string
    FieldsFunc() { hasQualifiedName("strings", "FieldsFunc") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class Join extends TaintTracking::FunctionModel {
    // signature: func Join(elems []string, sep string) string
    Join() { hasQualifiedName("strings", "Join") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class Map extends TaintTracking::FunctionModel {
    // signature: func Map(mapping func(rune) rune, s string) string
    Map() { hasQualifiedName("strings", "Map") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isResult())
    }
  }

  private class NewReader extends TaintTracking::FunctionModel {
    // signature: func NewReader(s string) *Reader
    NewReader() { hasQualifiedName("strings", "NewReader") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class NewReplacer extends TaintTracking::FunctionModel {
    // signature: func NewReplacer(oldnew ...string) *Replacer
    NewReplacer() { hasQualifiedName("strings", "NewReplacer") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class Repeat extends TaintTracking::FunctionModel {
    // signature: func Repeat(s string, count int) string
    Repeat() { hasQualifiedName("strings", "Repeat") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class Replace extends TaintTracking::FunctionModel {
    // signature: func Replace(s string, old string, new string, n int) string
    Replace() { hasQualifiedName("strings", "Replace") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter([0, 2]) and outp.isResult())
    }
  }

  private class ReplaceAll extends TaintTracking::FunctionModel {
    // signature: func ReplaceAll(s string, old string, new string) string
    ReplaceAll() { hasQualifiedName("strings", "ReplaceAll") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter([0, 2]) and outp.isResult())
    }
  }

  private class Split extends TaintTracking::FunctionModel {
    // signature: func Split(s string, sep string) []string
    Split() { hasQualifiedName("strings", "Split") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class SplitAfter extends TaintTracking::FunctionModel {
    // signature: func SplitAfter(s string, sep string) []string
    SplitAfter() { hasQualifiedName("strings", "SplitAfter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class SplitAfterN extends TaintTracking::FunctionModel {
    // signature: func SplitAfterN(s string, sep string, n int) []string
    SplitAfterN() { hasQualifiedName("strings", "SplitAfterN") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class SplitN extends TaintTracking::FunctionModel {
    // signature: func SplitN(s string, sep string, n int) []string
    SplitN() { hasQualifiedName("strings", "SplitN") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class Title extends TaintTracking::FunctionModel {
    // signature: func Title(s string) string
    Title() { hasQualifiedName("strings", "Title") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class ToLower extends TaintTracking::FunctionModel {
    // signature: func ToLower(s string) string
    ToLower() { hasQualifiedName("strings", "ToLower") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class ToLowerSpecial extends TaintTracking::FunctionModel {
    // signature: func ToLowerSpecial(c unicode.SpecialCase, s string) string
    ToLowerSpecial() { hasQualifiedName("strings", "ToLowerSpecial") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isResult())
    }
  }

  private class ToTitle extends TaintTracking::FunctionModel {
    // signature: func ToTitle(s string) string
    ToTitle() { hasQualifiedName("strings", "ToTitle") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class ToTitleSpecial extends TaintTracking::FunctionModel {
    // signature: func ToTitleSpecial(c unicode.SpecialCase, s string) string
    ToTitleSpecial() { hasQualifiedName("strings", "ToTitleSpecial") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isResult())
    }
  }

  private class ToUpper extends TaintTracking::FunctionModel {
    // signature: func ToUpper(s string) string
    ToUpper() { hasQualifiedName("strings", "ToUpper") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class ToUpperSpecial extends TaintTracking::FunctionModel {
    // signature: func ToUpperSpecial(c unicode.SpecialCase, s string) string
    ToUpperSpecial() { hasQualifiedName("strings", "ToUpperSpecial") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isResult())
    }
  }

  private class ToValidUTF8 extends TaintTracking::FunctionModel {
    // signature: func ToValidUTF8(s string, replacement string) string
    ToValidUTF8() { hasQualifiedName("strings", "ToValidUTF8") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class Trim extends TaintTracking::FunctionModel {
    // signature: func Trim(s string, cutset string) string
    Trim() { hasQualifiedName("strings", "Trim") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class TrimFunc extends TaintTracking::FunctionModel {
    // signature: func TrimFunc(s string, f func(rune) bool) string
    TrimFunc() { hasQualifiedName("strings", "TrimFunc") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class TrimLeft extends TaintTracking::FunctionModel {
    // signature: func TrimLeft(s string, cutset string) string
    TrimLeft() { hasQualifiedName("strings", "TrimLeft") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class TrimLeftFunc extends TaintTracking::FunctionModel {
    // signature: func TrimLeftFunc(s string, f func(rune) bool) string
    TrimLeftFunc() { hasQualifiedName("strings", "TrimLeftFunc") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class TrimPrefix extends TaintTracking::FunctionModel {
    // signature: func TrimPrefix(s string, prefix string) string
    TrimPrefix() { hasQualifiedName("strings", "TrimPrefix") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class TrimRight extends TaintTracking::FunctionModel {
    // signature: func TrimRight(s string, cutset string) string
    TrimRight() { hasQualifiedName("strings", "TrimRight") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class TrimRightFunc extends TaintTracking::FunctionModel {
    // signature: func TrimRightFunc(s string, f func(rune) bool) string
    TrimRightFunc() { hasQualifiedName("strings", "TrimRightFunc") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class TrimSpace extends TaintTracking::FunctionModel {
    // signature: func TrimSpace(s string) string
    TrimSpace() { hasQualifiedName("strings", "TrimSpace") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class TrimSuffix extends TaintTracking::FunctionModel {
    // signature: func TrimSuffix(s string, suffix string) string
    TrimSuffix() { hasQualifiedName("strings", "TrimSuffix") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class BuilderString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Builder).String() string
    BuilderString() { this.(Method).hasQualifiedName("strings", "Builder", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class BuilderWrite extends TaintTracking::FunctionModel, Method {
    // signature: func (*Builder).Write(p []byte) (int, error)
    BuilderWrite() { this.(Method).hasQualifiedName("strings", "Builder", "Write") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class BuilderWriteByte extends TaintTracking::FunctionModel, Method {
    // signature: func (*Builder).WriteByte(c byte) error
    BuilderWriteByte() { this.(Method).hasQualifiedName("strings", "Builder", "WriteByte") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class BuilderWriteRune extends TaintTracking::FunctionModel, Method {
    // signature: func (*Builder).WriteRune(r rune) (int, error)
    BuilderWriteRune() { this.(Method).hasQualifiedName("strings", "Builder", "WriteRune") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class BuilderWriteString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Builder).WriteString(s string) (int, error)
    BuilderWriteString() { this.(Method).hasQualifiedName("strings", "Builder", "WriteString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class ReaderRead extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).Read(b []byte) (n int, err error)
    ReaderRead() { this.(Method).hasQualifiedName("strings", "Reader", "Read") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class ReaderReadAt extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadAt(b []byte, off int64) (n int, err error)
    ReaderReadAt() { this.(Method).hasQualifiedName("strings", "Reader", "ReadAt") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class ReaderReadByte extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadByte() (byte, error)
    ReaderReadByte() { this.(Method).hasQualifiedName("strings", "Reader", "ReadByte") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult(0))
    }
  }

  private class ReaderReadRune extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).ReadRune() (ch rune, size int, err error)
    ReaderReadRune() { this.(Method).hasQualifiedName("strings", "Reader", "ReadRune") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult(0))
    }
  }

  private class ReaderReset extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).Reset(s string)
    ReaderReset() { this.(Method).hasQualifiedName("strings", "Reader", "Reset") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class ReaderWriteTo extends TaintTracking::FunctionModel, Method {
    // signature: func (*Reader).WriteTo(w io.Writer) (n int64, err error)
    ReaderWriteTo() { this.(Method).hasQualifiedName("strings", "Reader", "WriteTo") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isParameter(0))
    }
  }

  private class ReplacerReplace extends TaintTracking::FunctionModel, Method {
    // signature: func (*Replacer).Replace(s string) string
    ReplacerReplace() { this.(Method).hasQualifiedName("strings", "Replacer", "Replace") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class ReplacerWriteString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Replacer).WriteString(w io.Writer, s string) (n int, err error)
    ReplacerWriteString() { this.(Method).hasQualifiedName("strings", "Replacer", "WriteString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isParameter(0))
    }
  }
}
