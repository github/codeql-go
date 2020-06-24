/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `html/template` package. */
module HtmlTemplateTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func HTMLEscape(w io.Writer, b []byte)
      hasQualifiedName("html/template", "HTMLEscape") and
      (inp.isParameter(1) and outp.isParameter(0))
      or
      // signature: func HTMLEscapeString(s string) string
      hasQualifiedName("html/template", "HTMLEscapeString") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func HTMLEscaper(args ...interface{}) string
      hasQualifiedName("html/template", "HTMLEscaper") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func JSEscape(w io.Writer, b []byte)
      hasQualifiedName("html/template", "JSEscape") and
      (inp.isParameter(1) and outp.isParameter(0))
      or
      // signature: func JSEscapeString(s string) string
      hasQualifiedName("html/template", "JSEscapeString") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func JSEscaper(args ...interface{}) string
      hasQualifiedName("html/template", "JSEscaper") and
      (inp.isParameter(_) and outp.isResult())
      or
      // signature: func Must(t *Template, err error) *Template
      hasQualifiedName("html/template", "Must") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func ParseFiles(filenames ...string) (*Template, error)
      hasQualifiedName("html/template", "ParseFiles") and
      (inp.isParameter(_) and outp.isResult(0))
      or
      // signature: func ParseGlob(pattern string) (*Template, error)
      hasQualifiedName("html/template", "ParseGlob") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func URLQueryEscaper(args ...interface{}) string
      hasQualifiedName("html/template", "URLQueryEscaper") and
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
      // signature: func (*Template).Clone() (*Template, error)
      this.(Method).hasQualifiedName("html/template", "Template", "Clone") and
      (inp.isReceiver() and outp.isResult(0))
      or
      // signature: func (*Template).DefinedTemplates() string
      this.(Method).hasQualifiedName("html/template", "Template", "DefinedTemplates") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*Template).Delims(left string, right string) *Template
      this.(Method).hasQualifiedName("html/template", "Template", "Delims") and
      (
        (inp.isReceiver() or inp.isParameter(_)) and
        outp.isResult()
      )
      or
      // signature: func (*Template).Execute(wr io.Writer, data interface{}) error
      this.(Method).hasQualifiedName("html/template", "Template", "Execute") and
      (
        (inp.isReceiver() or inp.isParameter(1)) and
        outp.isParameter(0)
      )
      or
      // signature: func (*Template).ExecuteTemplate(wr io.Writer, name string, data interface{}) error
      this.(Method).hasQualifiedName("html/template", "Template", "ExecuteTemplate") and
      (
        (inp.isReceiver() or inp.isParameter(2)) and
        outp.isParameter(0)
      )
      or
      // signature: func (*Template).Lookup(name string) *Template
      this.(Method).hasQualifiedName("html/template", "Template", "Lookup") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*Template).New(name string) *Template
      this.(Method).hasQualifiedName("html/template", "Template", "New") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*Template).Option(opt ...string) *Template
      this.(Method).hasQualifiedName("html/template", "Template", "Option") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*Template).Parse(text string) (*Template, error)
      this.(Method).hasQualifiedName("html/template", "Template", "Parse") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func (*Template).ParseFiles(filenames ...string) (*Template, error)
      this.(Method).hasQualifiedName("html/template", "Template", "ParseFiles") and
      (
        (inp.isReceiver() or inp.isParameter(_)) and
        outp.isResult(0)
      )
      or
      // signature: func (*Template).ParseGlob(pattern string) (*Template, error)
      this.(Method).hasQualifiedName("html/template", "Template", "ParseGlob") and
      (
        (inp.isReceiver() or inp.isParameter(0)) and
        outp.isResult(0)
      )
      or
      // signature: func (*Template).Templates() []*Template
      this.(Method).hasQualifiedName("html/template", "Template", "Templates") and
      (inp.isReceiver() and outp.isResult())
      // Interfaces:
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
