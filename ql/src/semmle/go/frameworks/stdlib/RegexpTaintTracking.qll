/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `regexp` package. */
module RegexpTaintTracking {
  private class Compile extends TaintTracking::FunctionModel {
    // signature: func Compile(expr string) (*Regexp, error)
    Compile() { hasQualifiedName("regexp", "Compile") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class CompilePOSIX extends TaintTracking::FunctionModel {
    // signature: func CompilePOSIX(expr string) (*Regexp, error)
    CompilePOSIX() { hasQualifiedName("regexp", "CompilePOSIX") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class MustCompile extends TaintTracking::FunctionModel {
    // signature: func MustCompile(str string) *Regexp
    MustCompile() { hasQualifiedName("regexp", "MustCompile") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class MustCompilePOSIX extends TaintTracking::FunctionModel {
    // signature: func MustCompilePOSIX(str string) *Regexp
    MustCompilePOSIX() { hasQualifiedName("regexp", "MustCompilePOSIX") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class QuoteMeta extends TaintTracking::FunctionModel {
    // signature: func QuoteMeta(s string) string
    QuoteMeta() { hasQualifiedName("regexp", "QuoteMeta") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpCopy extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).Copy() *Regexp
    RegexpCopy() { this.(Method).hasQualifiedName("regexp", "Regexp", "Copy") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class RegexpExpand extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).Expand(dst []byte, template []byte, src []byte, match []int) []byte
    RegexpExpand() { this.(Method).hasQualifiedName("regexp", "Regexp", "Expand") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (
        inp.isParameter([1, 2]) and
        (outp.isParameter(0) or outp.isResult())
      )
    }
  }

  private class RegexpExpandString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).ExpandString(dst []byte, template string, src string, match []int) []byte
    RegexpExpandString() { this.(Method).hasQualifiedName("regexp", "Regexp", "ExpandString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (
        inp.isParameter([1, 2]) and
        (outp.isParameter(0) or outp.isResult())
      )
    }
  }

  private class RegexpFind extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).Find(b []byte) []byte
    RegexpFind() { this.(Method).hasQualifiedName("regexp", "Regexp", "Find") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpFindAll extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).FindAll(b []byte, n int) [][]byte
    RegexpFindAll() { this.(Method).hasQualifiedName("regexp", "Regexp", "FindAll") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpFindAllIndex extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).FindAllIndex(b []byte, n int) [][]int
    RegexpFindAllIndex() { this.(Method).hasQualifiedName("regexp", "Regexp", "FindAllIndex") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpFindAllString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).FindAllString(s string, n int) []string
    RegexpFindAllString() { this.(Method).hasQualifiedName("regexp", "Regexp", "FindAllString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpFindAllStringIndex extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).FindAllStringIndex(s string, n int) [][]int
    RegexpFindAllStringIndex() {
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindAllStringIndex")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpFindAllStringSubmatch extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).FindAllStringSubmatch(s string, n int) [][]string
    RegexpFindAllStringSubmatch() {
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindAllStringSubmatch")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpFindAllStringSubmatchIndex extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).FindAllStringSubmatchIndex(s string, n int) [][]int
    RegexpFindAllStringSubmatchIndex() {
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindAllStringSubmatchIndex")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpFindAllSubmatch extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).FindAllSubmatch(b []byte, n int) [][][]byte
    RegexpFindAllSubmatch() {
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindAllSubmatch")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpFindAllSubmatchIndex extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).FindAllSubmatchIndex(b []byte, n int) [][]int
    RegexpFindAllSubmatchIndex() {
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindAllSubmatchIndex")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpFindIndex extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).FindIndex(b []byte) (loc []int)
    RegexpFindIndex() { this.(Method).hasQualifiedName("regexp", "Regexp", "FindIndex") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpFindReaderIndex extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).FindReaderIndex(r io.RuneReader) (loc []int)
    RegexpFindReaderIndex() {
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindReaderIndex")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpFindReaderSubmatchIndex extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).FindReaderSubmatchIndex(r io.RuneReader) []int
    RegexpFindReaderSubmatchIndex() {
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindReaderSubmatchIndex")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpFindString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).FindString(s string) string
    RegexpFindString() { this.(Method).hasQualifiedName("regexp", "Regexp", "FindString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpFindStringIndex extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).FindStringIndex(s string) (loc []int)
    RegexpFindStringIndex() {
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindStringIndex")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpFindStringSubmatch extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).FindStringSubmatch(s string) []string
    RegexpFindStringSubmatch() {
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindStringSubmatch")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpFindStringSubmatchIndex extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).FindStringSubmatchIndex(s string) []int
    RegexpFindStringSubmatchIndex() {
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindStringSubmatchIndex")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpFindSubmatch extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).FindSubmatch(b []byte) [][]byte
    RegexpFindSubmatch() { this.(Method).hasQualifiedName("regexp", "Regexp", "FindSubmatch") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpFindSubmatchIndex extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).FindSubmatchIndex(b []byte) []int
    RegexpFindSubmatchIndex() {
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindSubmatchIndex")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class RegexpReplaceAll extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).ReplaceAll(src []byte, repl []byte) []byte
    RegexpReplaceAll() { this.(Method).hasQualifiedName("regexp", "Regexp", "ReplaceAll") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class RegexpReplaceAllFunc extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte
    RegexpReplaceAllFunc() { this.(Method).hasQualifiedName("regexp", "Regexp", "ReplaceAllFunc") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class RegexpReplaceAllLiteral extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).ReplaceAllLiteral(src []byte, repl []byte) []byte
    RegexpReplaceAllLiteral() {
      this.(Method).hasQualifiedName("regexp", "Regexp", "ReplaceAllLiteral")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class RegexpReplaceAllLiteralString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).ReplaceAllLiteralString(src string, repl string) string
    RegexpReplaceAllLiteralString() {
      this.(Method).hasQualifiedName("regexp", "Regexp", "ReplaceAllLiteralString")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class RegexpReplaceAllString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).ReplaceAllString(src string, repl string) string
    RegexpReplaceAllString() {
      this.(Method).hasQualifiedName("regexp", "Regexp", "ReplaceAllString")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class RegexpReplaceAllStringFunc extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).ReplaceAllStringFunc(src string, repl func(string) string) string
    RegexpReplaceAllStringFunc() {
      this.(Method).hasQualifiedName("regexp", "Regexp", "ReplaceAllStringFunc")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class RegexpSplit extends TaintTracking::FunctionModel, Method {
    // signature: func (*Regexp).Split(s string, n int) []string
    RegexpSplit() { this.(Method).hasQualifiedName("regexp", "Regexp", "Split") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }
}
