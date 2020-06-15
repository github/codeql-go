/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `net/url` package. */
module NetUrlTaintTracking {
  private class Parse extends TaintTracking::FunctionModel {
    // signature: func Parse(rawurl string) (*URL, error)
    Parse() { hasQualifiedName("net/url", "Parse") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class ParseQuery extends TaintTracking::FunctionModel {
    // signature: func ParseQuery(query string) (Values, error)
    ParseQuery() { hasQualifiedName("net/url", "ParseQuery") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class ParseRequestURI extends TaintTracking::FunctionModel {
    // signature: func ParseRequestURI(rawurl string) (*URL, error)
    ParseRequestURI() { hasQualifiedName("net/url", "ParseRequestURI") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class PathEscape extends TaintTracking::FunctionModel {
    // signature: func PathEscape(s string) string
    PathEscape() { hasQualifiedName("net/url", "PathEscape") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class PathUnescape extends TaintTracking::FunctionModel {
    // signature: func PathUnescape(s string) (string, error)
    PathUnescape() { hasQualifiedName("net/url", "PathUnescape") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class QueryEscape extends TaintTracking::FunctionModel {
    // signature: func QueryEscape(s string) string
    QueryEscape() { hasQualifiedName("net/url", "QueryEscape") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class QueryUnescape extends TaintTracking::FunctionModel {
    // signature: func QueryUnescape(s string) (string, error)
    QueryUnescape() { hasQualifiedName("net/url", "QueryUnescape") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult(0))
    }
  }

  private class User extends TaintTracking::FunctionModel {
    // signature: func User(username string) *Userinfo
    User() { hasQualifiedName("net/url", "User") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isResult())
    }
  }

  private class UserPassword extends TaintTracking::FunctionModel {
    // signature: func UserPassword(username string, password string) *Userinfo
    UserPassword() { hasQualifiedName("net/url", "UserPassword") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isResult())
    }
  }

  private class URLEscapedPath extends TaintTracking::FunctionModel, Method {
    // signature: func (*URL).EscapedPath() string
    URLEscapedPath() { this.(Method).hasQualifiedName("net/url", "URL", "EscapedPath") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class URLHostname extends TaintTracking::FunctionModel, Method {
    // signature: func (*URL).Hostname() string
    URLHostname() { this.(Method).hasQualifiedName("net/url", "URL", "Hostname") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class URLMarshalBinary extends TaintTracking::FunctionModel, Method {
    // signature: func (*URL).MarshalBinary() (text []byte, err error)
    URLMarshalBinary() { this.(Method).hasQualifiedName("net/url", "URL", "MarshalBinary") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult(0))
    }
  }

  private class URLParse extends TaintTracking::FunctionModel, Method {
    // signature: func (*URL).Parse(ref string) (*URL, error)
    URLParse() { this.(Method).hasQualifiedName("net/url", "URL", "Parse") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (
        (inp.isReceiver() or inp.isParameter(0)) and
        outp.isResult(0)
      )
    }
  }

  private class URLPort extends TaintTracking::FunctionModel, Method {
    // signature: func (*URL).Port() string
    URLPort() { this.(Method).hasQualifiedName("net/url", "URL", "Port") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class URLQuery extends TaintTracking::FunctionModel, Method {
    // signature: func (*URL).Query() Values
    URLQuery() { this.(Method).hasQualifiedName("net/url", "URL", "Query") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class URLRequestURI extends TaintTracking::FunctionModel, Method {
    // signature: func (*URL).RequestURI() string
    URLRequestURI() { this.(Method).hasQualifiedName("net/url", "URL", "RequestURI") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class URLResolveReference extends TaintTracking::FunctionModel, Method {
    // signature: func (*URL).ResolveReference(ref *URL) *URL
    URLResolveReference() { this.(Method).hasQualifiedName("net/url", "URL", "ResolveReference") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (
        (inp.isReceiver() or inp.isParameter(0)) and
        outp.isResult()
      )
    }
  }

  private class URLString extends TaintTracking::FunctionModel, Method {
    // signature: func (*URL).String() string
    URLString() { this.(Method).hasQualifiedName("net/url", "URL", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class URLUnmarshalBinary extends TaintTracking::FunctionModel, Method {
    // signature: func (*URL).UnmarshalBinary(text []byte) error
    URLUnmarshalBinary() { this.(Method).hasQualifiedName("net/url", "URL", "UnmarshalBinary") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(0) and outp.isReceiver())
    }
  }

  private class UserinfoPassword extends TaintTracking::FunctionModel, Method {
    // signature: func (*Userinfo).Password() (string, bool)
    UserinfoPassword() { this.(Method).hasQualifiedName("net/url", "Userinfo", "Password") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult(0))
    }
  }

  private class UserinfoString extends TaintTracking::FunctionModel, Method {
    // signature: func (*Userinfo).String() string
    UserinfoString() { this.(Method).hasQualifiedName("net/url", "Userinfo", "String") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class UserinfoUsername extends TaintTracking::FunctionModel, Method {
    // signature: func (*Userinfo).Username() string
    UserinfoUsername() { this.(Method).hasQualifiedName("net/url", "Userinfo", "Username") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class ValuesAdd extends TaintTracking::FunctionModel, Method {
    // signature: func (Values).Add(key string, value string)
    ValuesAdd() { this.(Method).hasQualifiedName("net/url", "Values", "Add") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isReceiver())
    }
  }

  private class ValuesEncode extends TaintTracking::FunctionModel, Method {
    // signature: func (Values).Encode() string
    ValuesEncode() { this.(Method).hasQualifiedName("net/url", "Values", "Encode") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class ValuesGet extends TaintTracking::FunctionModel, Method {
    // signature: func (Values).Get(key string) string
    ValuesGet() { this.(Method).hasQualifiedName("net/url", "Values", "Get") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isReceiver() and outp.isResult())
    }
  }

  private class ValuesSet extends TaintTracking::FunctionModel, Method {
    // signature: func (Values).Set(key string, value string)
    ValuesSet() { this.(Method).hasQualifiedName("net/url", "Values", "Set") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      (inp.isParameter(_) and outp.isReceiver())
    }
  }
}
