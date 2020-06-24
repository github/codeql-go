/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `net/url` package. */
module NetUrlTaintTracking {
  private class FunctionTaintTracking extends TaintTracking::FunctionModel {
    FunctionInput inp;
    FunctionOutput outp;

    FunctionTaintTracking() {
      // signature: func Parse(rawurl string) (*URL, error)
      hasQualifiedName("net/url", "Parse") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func ParseQuery(query string) (Values, error)
      hasQualifiedName("net/url", "ParseQuery") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func ParseRequestURI(rawurl string) (*URL, error)
      hasQualifiedName("net/url", "ParseRequestURI") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func PathEscape(s string) string
      hasQualifiedName("net/url", "PathEscape") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func PathUnescape(s string) (string, error)
      hasQualifiedName("net/url", "PathUnescape") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func QueryEscape(s string) string
      hasQualifiedName("net/url", "QueryEscape") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func QueryUnescape(s string) (string, error)
      hasQualifiedName("net/url", "QueryUnescape") and
      (inp.isParameter(0) and outp.isResult(0))
      or
      // signature: func User(username string) *Userinfo
      hasQualifiedName("net/url", "User") and
      (inp.isParameter(0) and outp.isResult())
      or
      // signature: func UserPassword(username string, password string) *Userinfo
      hasQualifiedName("net/url", "UserPassword") and
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
      // signature: func (*URL).EscapedPath() string
      this.(Method).hasQualifiedName("net/url", "URL", "EscapedPath") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*URL).Hostname() string
      this.(Method).hasQualifiedName("net/url", "URL", "Hostname") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*URL).MarshalBinary() (text []byte, err error)
      this.(Method).hasQualifiedName("net/url", "URL", "MarshalBinary") and
      (inp.isReceiver() and outp.isResult(0))
      or
      // signature: func (*URL).Parse(ref string) (*URL, error)
      this.(Method).hasQualifiedName("net/url", "URL", "Parse") and
      (
        (inp.isReceiver() or inp.isParameter(0)) and
        outp.isResult(0)
      )
      or
      // signature: func (*URL).Port() string
      this.(Method).hasQualifiedName("net/url", "URL", "Port") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*URL).Query() Values
      this.(Method).hasQualifiedName("net/url", "URL", "Query") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*URL).RequestURI() string
      this.(Method).hasQualifiedName("net/url", "URL", "RequestURI") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*URL).ResolveReference(ref *URL) *URL
      this.(Method).hasQualifiedName("net/url", "URL", "ResolveReference") and
      (
        (inp.isReceiver() or inp.isParameter(0)) and
        outp.isResult()
      )
      or
      // signature: func (*URL).String() string
      this.(Method).hasQualifiedName("net/url", "URL", "String") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*URL).UnmarshalBinary(text []byte) error
      this.(Method).hasQualifiedName("net/url", "URL", "UnmarshalBinary") and
      (inp.isParameter(0) and outp.isReceiver())
      or
      // signature: func (*Userinfo).Password() (string, bool)
      this.(Method).hasQualifiedName("net/url", "Userinfo", "Password") and
      (inp.isReceiver() and outp.isResult(0))
      or
      // signature: func (*Userinfo).String() string
      this.(Method).hasQualifiedName("net/url", "Userinfo", "String") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (*Userinfo).Username() string
      this.(Method).hasQualifiedName("net/url", "Userinfo", "Username") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Values).Add(key string, value string)
      this.(Method).hasQualifiedName("net/url", "Values", "Add") and
      (inp.isParameter(_) and outp.isReceiver())
      or
      // signature: func (Values).Encode() string
      this.(Method).hasQualifiedName("net/url", "Values", "Encode") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Values).Get(key string) string
      this.(Method).hasQualifiedName("net/url", "Values", "Get") and
      (inp.isReceiver() and outp.isResult())
      or
      // signature: func (Values).Set(key string, value string)
      this.(Method).hasQualifiedName("net/url", "Values", "Set") and
      (inp.isParameter(_) and outp.isReceiver())
      // Interfaces:
    }

    override predicate hasTaintFlow(FunctionInput input, FunctionOutput output) {
      input = inp and output = outp
    }
  }
}
