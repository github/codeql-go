/**
 * Provides default sources, sinks and sanitizers for reasoning about unsafe-quoting
 * vulnerabilities, as well as extension points for adding your own.
 */

import go

/**
 * Provides extension points for customizing the taint tracking configuration for reasoning about
 * unsafe-quoting vulnerabilities.
 */
module StringBreak {
  /** A (single or double) quote. */
  class Quote extends string {
    Quote() { this = "'" or this = "\"" }

    /** Gets the type of this quote, either single or double. */
    string getType() { if this = "'" then result = "single" else result = "double" }
  }

  /**
   * A data flow source for unsafe-quoting vulnerabilities.
   */
  abstract class Source extends DataFlow::Node { }

  /**
   * A data flow sink for unsafe-quoting vulnerabilities.
   */
  abstract class Sink extends DataFlow::Node {
    /** Gets the quote character for which this is a sink. */
    abstract Quote getQuote();
  }

  /**
   * A sanitizer for unsafe-quoting vulnerabilities.
   */
  abstract class Sanitizer extends DataFlow::Node {
    /** Gets the quote character for which this is a sanitizer. */
    Quote getQuote() { any() }
  }

  /**
   * A sanitizer guard for unsafe-quoting vulnerabilities.
   */
  abstract class SanitizerGuard extends DataFlow::BarrierGuard { }

  /** Holds if `l` contains a `quote` (either single or double). */
  private predicate containsQuote(StringOps::ConcatenationLeaf l, Quote quote) {
    quote = l.getStringValue().regexpFind("['\"]", _, _)
  }

  /** A call to `json.Marshal`, considered as a taint source for unsafe quoting. */
  class JsonMarshalAsSource extends Source {
    JsonMarshalAsSource() {
      exists(MarshalingFunction jsonMarshal | jsonMarshal.getFormat() = "JSON" |
        this = jsonMarshal.getOutput().getNode(jsonMarshal.getACall())
      )
    }
  }

  /** A string concatenation with quotes, considered as a taint sink for unsafe quoting. */
  class StringConcatenationAsSink extends Sink {
    Quote quote;

    StringConcatenationAsSink() {
      exists(StringOps::ConcatenationLeaf lf | lf.asNode() = this |
        containsQuote(lf.getPreviousLeaf(), quote) and
        containsQuote(lf.getNextLeaf(), quote)
      )
    }

    override Quote getQuote() { result = quote }
  }

  /** A call to `json.Unmarshal`, considered as a sanitizer for unsafe quoting. */
  class UnmarshalSanitizer extends Sanitizer {
    UnmarshalSanitizer() {
      exists(UnmarshalingFunction jsonUnmarshal | jsonUnmarshal.getFormat() = "JSON" |
        this = jsonUnmarshal.getOutput().getNode(jsonUnmarshal.getACall())
      )
    }
  }

  /**
   * A call to `strings.Replace` or `strings.ReplaceAll`, considered as a sanitizer
   * for unsafe quoting.
   */
  class ReplaceSanitizer extends Sanitizer {
    Quote quote;

    ReplaceSanitizer() {
      exists(string name, DataFlow::CallNode call |
        this = call and
        call.getTarget().hasQualifiedName("strings", name) and
        call.getArgument(1).getStringValue().matches("%" + quote + "%")
      |
        name = "Replace" and
        call.getArgument(3).getNumericValue() < 0
        or
        name = "ReplaceAll"
      )
    }

    override Quote getQuote() { result = quote }
  }

  class StringsNewReplacerCall extends DataFlow::CallNode {
    StringsNewReplacerCall() { this.getTarget().hasQualifiedName("strings", "NewReplacer") }

    DataFlow::Node getAReplacedArgument() {
      exists(int m, int n | m = 2 * n and n = m / 2 and result = getArgument(m))
    }
  }

  class StringsNewReplacerConfiguration extends DataFlow2::Configuration {
    StringsNewReplacerConfiguration() { this = "StringsNewReplacerConfiguration" }

    override predicate isSource(DataFlow::Node source) { source instanceof StringsNewReplacerCall }

    override predicate isSink(DataFlow::Node sink) {
      exists(DataFlow::MethodCallNode call |
        sink = call.getReceiver() and
        call.getTarget().hasQualifiedName("strings", "Replacer", ["Replace", "WriteString"])
      )
    }
  }

  /**
   * A call to `strings.Replacer.Replace`, considered as a sanitizer for unsafe
   * quoting.
   */
  class ReplacerReplaceSanitizer extends DataFlow::MethodCallNode, Sanitizer {
    Quote quote;

    ReplacerReplaceSanitizer() {
      exists(StringsNewReplacerConfiguration config, DataFlow::Node source, DataFlow::Node sink |
        config.hasFlow(source, sink) and
        this.getTarget().hasQualifiedName("strings", "Replacer", "Replace") and
        sink = this.getReceiver() and
        quote = source.(StringsNewReplacerCall).getAReplacedArgument().getStringValue()
      )
    }

    override Quote getQuote() { result = quote }
  }

  /**
   * A call to `strings.Replacer.WriteString`, considered as a sanitizer for
   * unsafe quoting.
   */
  class ReplacerWriteStringSanitizer extends Sanitizer {
    Quote quote;

    ReplacerWriteStringSanitizer() {
      exists(
        StringsNewReplacerConfiguration config, DataFlow::Node source, DataFlow::Node sink,
        DataFlow::MethodCallNode call
      |
        config.hasFlow(source, sink) and
        call.getTarget().hasQualifiedName("strings", "Replacer", "WriteString") and
        sink = call.getReceiver() and
        this = call.getArgument(1) and
        quote = source.(StringsNewReplacerCall).getAReplacedArgument().getStringValue()
      )
    }

    override Quote getQuote() { result = quote }
  }
}
