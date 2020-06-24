/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `regexp` package. */
module RegexpTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func Compile(expr string) (*Regexp, error)
      hasQualifiedName("regexp", "Compile") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func CompilePOSIX(expr string) (*Regexp, error)
      hasQualifiedName("regexp", "CompilePOSIX") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func MustCompile(str string) *Regexp
      hasQualifiedName("regexp", "MustCompile") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func MustCompilePOSIX(str string) *Regexp
      hasQualifiedName("regexp", "MustCompilePOSIX") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func QuoteMeta(s string) string
      hasQualifiedName("regexp", "QuoteMeta") and
      (inp.isParameter(0) and outp.isResult())
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }

  private class MethodAndInterfaceTaintTracking extends TaintTracking::FunctionModel, Method {
    FunctionInput inp;
    FunctionOutput outp;

    MethodAndInterfaceTaintTracking() {
      // Methods:
      // signature: func (*Regexp).Copy() *Regexp
      this.(Method).hasQualifiedName("regexp", "Regexp", "Copy") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*Regexp).Expand(dst []byte, template []byte, src []byte, match []int) []byte
      this.(Method).hasQualifiedName("regexp", "Regexp", "Expand") and
      (
        inp.isParameter([1, 2]) and
        (outp.isParameter(0) or outp.isResult())
      )
      or
      // signature: func (*Regexp).ExpandString(dst []byte, template string, src string, match []int) []byte
      this.(Method).hasQualifiedName("regexp", "Regexp", "ExpandString") and
      (
        inp.isParameter([1, 2]) and
        (outp.isParameter(0) or outp.isResult())
      )
      or
      // signature: func (*Regexp).Find(b []byte) []byte
      this.(Method).hasQualifiedName("regexp", "Regexp", "Find") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Regexp).FindAll(b []byte, n int) [][]byte
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindAll") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Regexp).FindAllIndex(b []byte, n int) [][]int
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindAllIndex") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Regexp).FindAllString(s string, n int) []string
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindAllString") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Regexp).FindAllStringIndex(s string, n int) [][]int
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindAllStringIndex") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Regexp).FindAllStringSubmatch(s string, n int) [][]string
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindAllStringSubmatch") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Regexp).FindAllStringSubmatchIndex(s string, n int) [][]int
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindAllStringSubmatchIndex") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Regexp).FindAllSubmatch(b []byte, n int) [][][]byte
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindAllSubmatch") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Regexp).FindAllSubmatchIndex(b []byte, n int) [][]int
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindAllSubmatchIndex") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Regexp).FindIndex(b []byte) (loc []int)
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindIndex") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Regexp).FindReaderIndex(r io.RuneReader) (loc []int)
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindReaderIndex") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Regexp).FindReaderSubmatchIndex(r io.RuneReader) []int
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindReaderSubmatchIndex") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Regexp).FindString(s string) string
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindString") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Regexp).FindStringIndex(s string) (loc []int)
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindStringIndex") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Regexp).FindStringSubmatch(s string) []string
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindStringSubmatch") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Regexp).FindStringSubmatchIndex(s string) []int
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindStringSubmatchIndex") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Regexp).FindSubmatch(b []byte) [][]byte
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindSubmatch") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Regexp).FindSubmatchIndex(b []byte) []int
      this.(Method).hasQualifiedName("regexp", "Regexp", "FindSubmatchIndex") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func (*Regexp).ReplaceAll(src []byte, repl []byte) []byte
      this.(Method).hasQualifiedName("regexp", "Regexp", "ReplaceAll") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func (*Regexp).ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte
      this.(Method).hasQualifiedName("regexp", "Regexp", "ReplaceAllFunc") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func (*Regexp).ReplaceAllLiteral(src []byte, repl []byte) []byte
      this.(Method).hasQualifiedName("regexp", "Regexp", "ReplaceAllLiteral") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func (*Regexp).ReplaceAllLiteralString(src string, repl string) string
      this.(Method).hasQualifiedName("regexp", "Regexp", "ReplaceAllLiteralString") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func (*Regexp).ReplaceAllString(src string, repl string) string
      this.(Method).hasQualifiedName("regexp", "Regexp", "ReplaceAllString") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func (*Regexp).ReplaceAllStringFunc(src string, repl func(string) string) string
      this.(Method).hasQualifiedName("regexp", "Regexp", "ReplaceAllStringFunc") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func (*Regexp).Split(s string, n int) []string
      this.(Method).hasQualifiedName("regexp", "Regexp", "Split") and
      (inp.isParameter(0) and outp.isResult())
      // Interfaces:
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
