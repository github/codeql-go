/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `text/template` package. */
module TextTemplateTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func HTMLEscape(w io.Writer, b []byte)
      hasQualifiedName("text/template", "HTMLEscape") and
      (inp.isParameter(1) and outp.isParameter(0))
      or
      // signature: func HTMLEscapeString(s string) string
      hasQualifiedName("text/template", "HTMLEscapeString") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func HTMLEscaper(args ...interface{}) string
      hasQualifiedName("text/template", "HTMLEscaper") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func JSEscape(w io.Writer, b []byte)
      hasQualifiedName("text/template", "JSEscape") and
      (inp.isParameter(1) and outp.isParameter(0))
      or
      // signature: func JSEscapeString(s string) string
      hasQualifiedName("text/template", "JSEscapeString") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func JSEscaper(args ...interface{}) string
      hasQualifiedName("text/template", "JSEscaper") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func New(name string) *Template
      hasQualifiedName("text/template", "New") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func ParseFiles(filenames ...string) (*Template, error)
      hasQualifiedName("text/template", "ParseFiles") and
      (inp.isParameter(_) and outp.isResult(0))
      or
      // signature: func ParseGlob(pattern string) (*Template, error)
      hasQualifiedName("text/template", "ParseGlob") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func URLQueryEscaper(args ...interface{}) string
      hasQualifiedName("text/template", "URLQueryEscaper") and
      (inp.isParameter(_) and outp.isResult())
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
      // signature: func (*Template).Execute(wr io.Writer, data interface{}) error
      this.(Method).hasQualifiedName("text/template", "Template", "Execute") and
      (
        (inp.isReceiver() or inp.isParameter(1)) and
        outp.isParameter(0)
      )
      or
      // signature: func (*Template).ExecuteTemplate(wr io.Writer, name string, data interface{}) error
      this.(Method).hasQualifiedName("text/template", "Template", "ExecuteTemplate") and
      (
        (inp.isReceiver() or inp.isParameter(2)) and
        outp.isParameter(0)
      )
      // Interfaces:
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
