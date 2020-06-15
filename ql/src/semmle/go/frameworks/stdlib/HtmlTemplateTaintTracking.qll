/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `html/template` package. */
module HtmlTemplateTaintTracking {
  private class HTMLEscape extends TaintTracking::FunctionModel {
    // signature: func HTMLEscape(w io.Writer, b []byte)
    HTMLEscape() { hasQualifiedName("html/template", "HTMLEscape") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isParameter(0))
    }
  }

  private class HTMLEscapeString extends TaintTracking::FunctionModel {
    // signature: func HTMLEscapeString(s string) string
    HTMLEscapeString() { hasQualifiedName("html/template", "HTMLEscapeString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class HTMLEscaper extends TaintTracking::FunctionModel {
    // signature: func HTMLEscaper(args ...interface{}) string
    HTMLEscaper() { hasQualifiedName("html/template", "HTMLEscaper") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class JSEscape extends TaintTracking::FunctionModel {
    // signature: func JSEscape(w io.Writer, b []byte)
    JSEscape() { hasQualifiedName("html/template", "JSEscape") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(1) and outp.isParameter(0))
    }
  }

  private class JSEscapeString extends TaintTracking::FunctionModel {
    // signature: func JSEscapeString(s string) string
    JSEscapeString() { hasQualifiedName("html/template", "JSEscapeString") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class JSEscaper extends TaintTracking::FunctionModel {
    // signature: func JSEscaper(args ...interface{}) string
    JSEscaper() { hasQualifiedName("html/template", "JSEscaper") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class Must extends TaintTracking::FunctionModel {
    // signature: func Must(t *Template, err error) *Template
    Must() { hasQualifiedName("html/template", "Must") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class ParseFiles extends TaintTracking::FunctionModel {
    // signature: func ParseFiles(filenames ...string) (*Template, error)
    ParseFiles() { hasQualifiedName("html/template", "ParseFiles") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult(0))
    }
  }

  private class ParseGlob extends TaintTracking::FunctionModel {
    // signature: func ParseGlob(pattern string) (*Template, error)
    ParseGlob() { hasQualifiedName("html/template", "ParseGlob") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class URLQueryEscaper extends TaintTracking::FunctionModel {
    // signature: func URLQueryEscaper(args ...interface{}) string
    URLQueryEscaper() { hasQualifiedName("html/template", "URLQueryEscaper") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class TemplateClone extends TaintTracking::FunctionModel, Method {
    // signature: func (*Template).Clone() (*Template, error)
    TemplateClone() { this.(Method).hasQualifiedName("html/template", "Template", "Clone") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult(0))
    }
  }

  private class TemplateDefinedTemplates extends TaintTracking::FunctionModel, Method {
    // signature: func (*Template).DefinedTemplates() string
    TemplateDefinedTemplates() {
      this.(Method).hasQualifiedName("html/template", "Template", "DefinedTemplates")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class TemplateDelims extends TaintTracking::FunctionModel, Method {
    // signature: func (*Template).Delims(left string, right string) *Template
    TemplateDelims() { this.(Method).hasQualifiedName("html/template", "Template", "Delims") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (
        (inp.isReceiver() or inp.isParameter(_)) and
        outp.isResult()
      )
    }
  }

  private class TemplateExecute extends TaintTracking::FunctionModel, Method {
    // signature: func (*Template).Execute(wr io.Writer, data interface{}) error
    TemplateExecute() { this.(Method).hasQualifiedName("html/template", "Template", "Execute") }

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
      this.(Method).hasQualifiedName("html/template", "Template", "ExecuteTemplate")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (
        (inp.isReceiver() or inp.isParameter(2)) and
        outp.isParameter(0)
      )
    }
  }

  private class TemplateLookup extends TaintTracking::FunctionModel, Method {
    // signature: func (*Template).Lookup(name string) *Template
    TemplateLookup() { this.(Method).hasQualifiedName("html/template", "Template", "Lookup") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class TemplateNew extends TaintTracking::FunctionModel, Method {
    // signature: func (*Template).New(name string) *Template
    TemplateNew() { this.(Method).hasQualifiedName("html/template", "Template", "New") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class TemplateOption extends TaintTracking::FunctionModel, Method {
    // signature: func (*Template).Option(opt ...string) *Template
    TemplateOption() { this.(Method).hasQualifiedName("html/template", "Template", "Option") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class TemplateParse extends TaintTracking::FunctionModel, Method {
    // signature: func (*Template).Parse(text string) (*Template, error)
    TemplateParse() { this.(Method).hasQualifiedName("html/template", "Template", "Parse") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class TemplateParseFiles extends TaintTracking::FunctionModel, Method {
    // signature: func (*Template).ParseFiles(filenames ...string) (*Template, error)
    TemplateParseFiles() {
      this.(Method).hasQualifiedName("html/template", "Template", "ParseFiles")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (
        (inp.isReceiver() or inp.isParameter(_)) and
        outp.isResult(0)
      )
    }
  }

  private class TemplateParseGlob extends TaintTracking::FunctionModel, Method {
    // signature: func (*Template).ParseGlob(pattern string) (*Template, error)
    TemplateParseGlob() { this.(Method).hasQualifiedName("html/template", "Template", "ParseGlob") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (
        (inp.isReceiver() or inp.isParameter(0)) and
        outp.isResult(0)
      )
    }
  }

  private class TemplateTemplates extends TaintTracking::FunctionModel, Method {
    // signature: func (*Template).Templates() []*Template
    TemplateTemplates() { this.(Method).hasQualifiedName("html/template", "Template", "Templates") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }
}
