/**
 * Provides default sources, sinks and sanitizers for reasoning about
 * JWT token signing vulnerabilities as well as extension points
 * for adding your own.
 */

import go
import DataFlow::PathGraph

/**
 * Provides default sources, sinks and sanitizers for reasoning about
 * JWT token signing vulnerabilities as well as extension points
 * for adding your own.
 */
module HardcodedKeys {
  /**
   * A data flow source for JWT token signing vulnerabilities.
   */
  abstract class Source extends DataFlow::Node { }

  /**
   * A data flow sink for JWT token signing vulnerabilities.
   */
  abstract class Sink extends DataFlow::Node { }

  /**
   * A sanitizer for JWT token signing vulnerabilities.
   */
  abstract class Sanitizer extends DataFlow::Node { }

  /**
   * A sanitizer guard for JWT token signing vulnerabilities.
   */
  abstract class SanitizerGuard extends DataFlow::BarrierGuard { }

  private predicate isTestCode(Expr e) {
    e.getFile().getAbsolutePath().toLowerCase().matches("%test%") and
    not e.getFile().getAbsolutePath().toLowerCase().matches("%ql/test%")
  }

  /**
   * A hardcoded string literal as a source for JWT token signing vulnerabilities.
   */
  class HardcodedStringSource extends Source {
    HardcodedStringSource() {
      this.asExpr() instanceof StringLit and
      not isTestCode(this.asExpr())
    }
  }

  /**
   * An expression used to sign JWT tokens as a sink for JWT token signing vulnerabilities.
   */
  private class GolangJwtSign extends Sink {
    string pkg;

    GolangJwtSign() {
      pkg =
        [
          "github.com/golang-jwt/jwt/v4", "github.com/dgrijalva/jwt-go",
          "github.com/form3tech-oss/jwt-go", "github.com/ory/fosite/token/jwt"
        ] and
      (
        exists(DataFlow::MethodCallNode m |
          // Models the `SignedString` method
          // `func (t *Token) SignedString(key interface{}) (string, error)`
          m.getTarget().hasQualifiedName(pkg, "Token", "SignedString")
        |
          this = m.getArgument(0)
        )
        or
        exists(DataFlow::MethodCallNode m |
          // Model the `Sign` method of the `SigningMethod` interface
          // type SigningMethod interface {
          //   Verify(signingString, signature string, key interface{}) error
          //   Sign(signingString string, key interface{}) (string, error)
          //   Alg() string
          // }
          m.getTarget().hasQualifiedName(pkg, "SigningMethod", "Sign")
        |
          this = m.getArgument(1)
        )
      )
    }
  }

  private class GinJwtSign extends Sink {
    GinJwtSign() {
      exists(Field f |
        // https://pkg.go.dev/github.com/appleboy/gin-jwt/v2#GinJWTMiddleware
        f.hasQualifiedName("github.com/appleboy/gin-jwt/v2", "GinJWTMiddleware", "Key") and
        f.getAWrite().getRhs() = this
      )
    }
  }

  private class SquareJoseKey extends Sink {
    SquareJoseKey() {
      exists(Field f, string pkg |
        // type Recipient struct {
        // 	Algorithm  KeyAlgorithm
        // 	Key        interface{}
        // 	KeyID      string
        // 	PBES2Count int
        // 	PBES2Salt  []byte
        // }
        // type SigningKey struct {
        // 	Algorithm SignatureAlgorithm
        // 	Key       interface{}
        // }
        f.hasQualifiedName(pkg, ["Recipient", "SigningKey"], "Key") and
        f.getAWrite().getRhs() = this
      |
        pkg = ["github.com/square/go-jose/v3", "gopkg.in/square/go-jose.v2"]
      )
    }
  }

  private class CrystalHqJwtSigner extends Sink {
    CrystalHqJwtSigner() {
      exists(DataFlow::CallNode m |
        // `func NewSignerHS(alg Algorithm, key []byte) (Signer, error)`
        m.getTarget().hasQualifiedName("github.com/cristalhq/jwt/v3", "NewSignerHS")
      |
        this = m.getArgument(1)
      )
    }
  }

  private class GoKitJwt extends Sink {
    GoKitJwt() {
      exists(DataFlow::CallNode m |
        // `func NewSigner(kid string, key []byte, method jwt.SigningMethod, claims jwt.Claims) endpoint.Middleware`
        m.getTarget().hasQualifiedName("github.com/go-kit/kit/auth/jwt", "NewSigner")
      |
        this = m.getArgument(1)
      )
    }
  }

  private class LestrratJwk extends Sink {
    LestrratJwk() {
      exists(DataFlow::CallNode m, string pkg |
        pkg.matches([
            "github.com/lestrrat-go/jwx", "github.com/lestrrat/go-jwx/jwk",
            "github.com/lestrrat-go/jwx%/jwk"
          ]) and
        // `func New(key interface{}) (Key, error)`
        m.getTarget().hasQualifiedName(pkg, "New")
      |
        this = m.getArgument(0)
      )
    }
  }

  /**
   * Mark any comparision expression where any operand is tainted as a
   * sanitizer for all instances of the taint
   */
  private class CompareExprSanitizer extends Sanitizer {
    CompareExprSanitizer() {
      exists(BinaryExpr c |
        c.getAnOperand().getGlobalValueNumber() = this.asExpr().getGlobalValueNumber()
      )
    }
  }

  /**
   * A configuration depicting taint flow for studying JWT token signing vulnerabilities.
   */
  class Configuration extends TaintTracking::Configuration {
    Configuration() { this = "Hard-coded JWT Signing Key" }

    override predicate isSource(DataFlow::Node source) { source instanceof Source }

    override predicate isSink(DataFlow::Node sink) { sink instanceof Sink }

    override predicate isAdditionalTaintStep(DataFlow::Node prev, DataFlow::Node succ) {
      exists(ConversionExpr ce | ce.getOperand() = prev.asExpr() and ce = succ.asExpr() |
        ce.getTypeExpr().getType() instanceof ByteSliceType
      )
    }

    override predicate isSanitizer(DataFlow::Node sanitizer) { sanitizer instanceof Sanitizer }

    override predicate isSanitizerGuard(DataFlow::BarrierGuard guard) {
      guard instanceof SanitizerGuard
    }
  }
}
