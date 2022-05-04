/**
 * Provides default sources, sinks, and sanitizers for reasoning about
 * log injection vulnerabilities, as well as extension points for adding your own.
 */

import go
private import semmle.go.StringOps

/**
 * Provides extension points for customizing the data-flow tracking configuration for reasoning
 * about log injection.
 */
module LogInjection {
  /**
   * A data flow source for log injection vulnerabilities.
   */
  abstract class Source extends DataFlow::Node { }

  /**
   * A data flow sink for log injection vulnerabilities.
   */
  abstract class Sink extends DataFlow::Node { }

  /**
   * A sanitizer for log injection vulnerabilities.
   */
  abstract class Sanitizer extends DataFlow::Node { }

  /**
   * A sanitizer guard for log injection vulnerabilities.
   */
  abstract class SanitizerGuard extends DataFlow::BarrierGuard { }

  /** A source of untrusted data, considered as a taint source for log injection. */
  class UntrustedFlowAsSource extends Source {
    UntrustedFlowAsSource() { this instanceof UntrustedFlowSource }
  }

  /** An argument to a logging mechanism. */
  class LoggerSink extends Sink {
    LoggerSink() { this = any(LoggerCall log).getAMessageComponent() }
  }

  /**
   * A call to `strings.Replace` or `strings.ReplaceAll`, considered as a sanitizer
   * for log injection.
   */
  class ReplaceSanitizer extends Sanitizer {
    ReplaceSanitizer() {
      exists(string name, DataFlow::CallNode call |
        this = call and
        call.getTarget().hasQualifiedName("strings", name) and
        call.getArgument(1).getStringValue().matches("%" + ["\r", "\n"] + "%")
      |
        name = "Replace" and
        call.getArgument(3).getNumericValue() < 0
        or
        name = "ReplaceAll"
      )
    }
  }

  /**
   * A call to `strings.Replacer.Replace`, considered as a sanitizer for log
   * injection.
   */
  class ReplacerReplaceSanitizer extends Sanitizer {
    ReplacerReplaceSanitizer() {
      exists(DataFlow::MethodCallNode call |
        call.(DataFlow::MethodCallNode)
            .getTarget()
            .hasQualifiedName("strings", "Replacer", "Replace") and
        this = call.getResult()
      )
    }
  }

  /**
   * A call to `strings.Replacer.WriteString`, considered as a sanitizer for log
   * injection.
   */
  class ReplacerWriteStringSanitizer extends Sanitizer {
    ReplacerWriteStringSanitizer() {
      exists(DataFlow::MethodCallNode call |
        call.getTarget().hasQualifiedName("strings", "Replacer", "WriteString") and
        this = call.getArgument(1)
      )
    }
  }

  /**
   * An argument that is formatted using the `%q` directive, considered as a sanitizer
   * for log injection.
   *
   * This formatting directive replaces newline characters with escape sequences.
   */
  private class SafeFormatArgumentSanitizer extends Sanitizer {
    SafeFormatArgumentSanitizer() {
      exists(StringOps::Formatting::StringFormatCall call, string safeDirective |
        this = call.getOperand(_, safeDirective) and
        // Mark "%q" formats as safe, but not "%#q", which would preserve newline characters.
        safeDirective.regexpMatch("%[^%#]*q")
      )
    }
  }
}
