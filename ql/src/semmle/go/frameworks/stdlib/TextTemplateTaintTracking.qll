/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `text/template` package. */
module TextTemplateTaintTracking {
  private class HTMLEscape extends TaintTracking::FunctionModel {
    // signature: func HTMLEscape(w io.Writer, b []byte)
    HTMLEscape() { hasQualifiedName("text/template", "HTMLEscape") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isParameter(0))
    }
  }

  private class HTMLEscapeString extends TaintTracking::FunctionModel {
    // signature: func HTMLEscapeString(s string) string
    HTMLEscapeString() { hasQualifiedName("text/template", "HTMLEscapeString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class HTMLEscaper extends TaintTracking::FunctionModel {
    // signature: func HTMLEscaper(args ...interface{}) string
    HTMLEscaper() { hasQualifiedName("text/template", "HTMLEscaper") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class JSEscape extends TaintTracking::FunctionModel {
    // signature: func JSEscape(w io.Writer, b []byte)
    JSEscape() { hasQualifiedName("text/template", "JSEscape") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isParameter(0))
    }
  }

  private class JSEscapeString extends TaintTracking::FunctionModel {
    // signature: func JSEscapeString(s string) string
    JSEscapeString() { hasQualifiedName("text/template", "JSEscapeString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class JSEscaper extends TaintTracking::FunctionModel {
    // signature: func JSEscaper(args ...interface{}) string
    JSEscaper() { hasQualifiedName("text/template", "JSEscaper") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class New extends TaintTracking::FunctionModel {
    // signature: func New(name string) *Template
    New() { hasQualifiedName("text/template", "New") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class ParseFiles extends TaintTracking::FunctionModel {
    // signature: func ParseFiles(filenames ...string) (*Template, error)
    ParseFiles() { hasQualifiedName("text/template", "ParseFiles") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult(0))
    }
  }

  private class ParseGlob extends TaintTracking::FunctionModel {
    // signature: func ParseGlob(pattern string) (*Template, error)
    ParseGlob() { hasQualifiedName("text/template", "ParseGlob") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class URLQueryEscaper extends TaintTracking::FunctionModel {
    // signature: func URLQueryEscaper(args ...interface{}) string
    URLQueryEscaper() { hasQualifiedName("text/template", "URLQueryEscaper") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class TemplateExecute extends TaintTracking::FunctionModel, Method {
    // signature: func (*Template).Execute(wr io.Writer, data interface{}) error
    TemplateExecute() { this.(Method).hasQualifiedName("text/template", "Template", "Execute") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (
        (inp.isReceiver() or inp.isParameter(1)) and
        outp.isParameter(0)
      )
    }
  }

  private class TemplateExecuteTemplate extends TaintTracking::FunctionModel, Method {
    // signature: func (*Template).ExecuteTemplate(wr io.Writer, name string, data interface{}) error
    TemplateExecuteTemplate() {
      this.(Method).hasQualifiedName("text/template", "Template", "ExecuteTemplate")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (
        (inp.isReceiver() or inp.isParameter(2)) and
        outp.isParameter(0)
      )
    }
  }
}
