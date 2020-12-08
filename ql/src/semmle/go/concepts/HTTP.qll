/**
 * Provides classes for working with HTTP-related concepts such as requests and responses.
 */

import go

/** Provides classes for modeling HTTP-related APIs. */
module HTTP {
  /** Provides a class for modeling new HTTP response-writer APIs. */
  module ResponseWriter {
    /**
     * A variable that is an HTTP response writer.
     *
     * Extend this class to model new APIs. If you want to refine existing API models,
     * extend `HTTP::ResponseWriter` instead.
     */
    abstract class Range extends Variable {
      /**
       * Gets a data-flow node that is a use of this response writer.
       *
       * Note that `PostUpdateNode`s for nodes that this predicate gets do not need to be
       * included, as they are handled by the concrete `ResponseWriter`'s `getANode`.
       */
      abstract DataFlow::Node getANode();
    }
  }

  /**
   * A variable that is an HTTP response writer.
   *
   * Extend this class to refine existing API models. If you want to model new APIs,
   * extend `HTTP::ResponseWriter::Range` instead.
   */
  class ResponseWriter extends Variable {
    ResponseWriter::Range self;

    ResponseWriter() { this = self }

    /** Gets the body that is written in this HTTP response. */
    ResponseBody getBody() { result.getResponseWriter() = this }

    /** Gets a header write that is written in this HTTP response. */
    HeaderWrite getAHeaderWrite() { result.getResponseWriter() = this }

    /** Gets a redirect that is sent in this HTTP response. */
    Redirect getARedirect() { result.getResponseWriter() = this }

    /** Gets a data-flow node that is a use of this response writer. */
    DataFlow::Node getANode() {
      result = self.getANode() or
      result.(DataFlow::PostUpdateNode).getPreUpdateNode() = self.getANode()
    }
  }

  /** Provides a class for modeling new HTTP header-write APIs. */
  module HeaderWrite {
    /**
     * A data-flow node that represents a write to an HTTP header.
     *
     * Extend this class to model new APIs. If you want to refine existing API models,
     * extend `HTTP::HeaderWrite` instead.
     */
    abstract class Range extends DataFlow::ExprNode {
      /** Gets the (lower-case) name of a header set by this definition. */
      string getHeaderName() { result = this.getName().getStringValue().toLowerCase() }

      /** Holds if this header write defines the header `header`. */
      predicate definesHeader(string header, string value) {
        header = this.getName().getStringValue().toLowerCase() and
        value = this.getValue().getStringValue()
      }

      /** Gets the node representing the name of the header defined by this write. */
      abstract DataFlow::Node getName();

      /** Gets the node representing the value of the header defined by this write. */
      abstract DataFlow::Node getValue();

      /** Gets the response writer associated with this header write, if any. */
      abstract ResponseWriter getResponseWriter();
    }
  }

  /**
   * A data-flow node that represents a write to an HTTP header.
   *
   * Extend this class to refine existing API models. If you want to model new APIs,
   * extend `HTTP::HeaderWrite::Range` instead.
   */
  class HeaderWrite extends DataFlow::ExprNode {
    HeaderWrite::Range self;

    HeaderWrite() { this = self }

    /** Gets the (lower-case) name of a header set by this definition. */
    string getHeaderName() { result = self.getHeaderName() }

    /** Holds if this header write defines the header `header`. */
    predicate definesHeader(string header, string value) { self.definesHeader(header, value) }

    /** Gets the node representing the name of the header defined by this write. */
    DataFlow::Node getName() { result = self.getName() }

    /** Gets the node representing the value of the header defined by this write. */
    DataFlow::Node getValue() { result = self.getValue() }

    /** Gets the response writer associated with this header write, if any. */
    ResponseWriter getResponseWriter() { result = self.getResponseWriter() }
  }

  /** A data-flow node whose value is written to an HTTP header. */
  class Header extends DataFlow::Node {
    HeaderWrite hw;

    Header() {
      this = hw.getName()
      or
      this = hw.getValue()
    }

    /** Gets the response writer associated with this header write, if any. */
    ResponseWriter getResponseWriter() { result = hw.getResponseWriter() }
  }

  /** A data-flow node whose value is written to the value of an HTTP header. */
  class HeaderValue extends Header {
    HeaderValue() { this = hw.getValue() }
  }

  /** A data-flow node whose value is written to the name of an HTTP header. */
  class HeaderName extends Header {
    HeaderName() { this = hw.getName() }
  }

  /** Provides a class for modeling new HTTP request-body APIs. */
  module RequestBody {
    /**
     * An expression representing a reader whose content is written to an HTTP request body.
     *
     * Extend this class to model new APIs. If you want to refine existing API models,
     * extend `HTTP::RequestBody` instead.
     */
    abstract class Range extends DataFlow::Node { }
  }

  /**
   * An expression representing a reader whose content is written to an HTTP request body.
   *
   * Extend this class to refine existing API models. If you want to model new APIs,
   * extend `HTTP::RequestBody::Range` instead.
   */
  class RequestBody extends DataFlow::Node {
    RequestBody::Range self;

    RequestBody() { this = self }
  }

  /** Provides a class for modeling new HTTP response-body APIs. */
  module ResponseBody {
    /**
     * An expression which is written to an HTTP response body.
     *
     * Extend this class to model new APIs. If you want to refine existing API models,
     * extend `HTTP::ResponseBody` instead.
     */
    abstract class Range extends DataFlow::Node {
      /** Gets the response writer associated with this header write, if any. */
      abstract ResponseWriter getResponseWriter();

      /** Gets a content-type associated with this body. */
      string getAContentType() { result = getAContentTypeNode().getStringValue() }

      /** Gets a dataflow node for a content-type associated with this body. */
      DataFlow::Node getAContentTypeNode() {
        exists(HTTP::HeaderWrite hw | hw = getResponseWriter().getAHeaderWrite() |
          hw.getName().getStringValue().toLowerCase() = "content-type" and
          result = hw.getValue()
        )
      }
    }
  }

  /**
   * An expression which is written to an HTTP response body.
   *
   * Extend this class to refine existing API models. If you want to model new APIs,
   * extend `HTTP::ResponseBody::Range` instead.
   */
  class ResponseBody extends DataFlow::Node {
    ResponseBody::Range self;

    ResponseBody() { this = self }

    /** Gets the response writer associated with this header write, if any. */
    ResponseWriter getResponseWriter() { result = self.getResponseWriter() }

    /** Gets a content-type associated with this body. */
    string getAContentType() { result = self.getAContentType() }

    /** Gets a dataflow node for a content-type associated with this body. */
    DataFlow::Node getAContentTypeNode() { result = self.getAContentTypeNode() }
  }

  /** Provides a class for modeling new HTTP client request APIs. */
  module ClientRequest {
    /**
     * A call that performs a request to a URL.
     *
     * Example: An HTTP POST request is a client request that sends some
     * `data` to a `url`, where both the headers and the body of the request
     * contribute to the `data`.
     *
     * Extend this class to model new APIs. If you want to refine existing API models,
     * extend `HTTP::ClientRequest` instead.
     */
    abstract class Range extends DataFlow::Node {
      /**
       * Gets the URL of the request.
       */
      abstract DataFlow::Node getUrl();
    }
  }

  /**
   * A call that performs a request to a URL.
   *
   * Extend this class to refine existing API models. If you want to model new APIs,
   * extend `HTTP::ClientRequest::Range` instead.
   */
  class ClientRequest extends DataFlow::Node {
    ClientRequest::Range self;

    ClientRequest() { this = self }

    /**
     * Gets the URL of the request.
     */
    DataFlow::Node getUrl() { result = self.getUrl() }
  }

  /** Provides a class for modeling new HTTP redirect APIs. */
  module Redirect {
    /**
     * An HTTP redirect.
     *
     * Extend this class to model new APIs. If you want to refine existing API models,
     * extend `HTTP::Redirect` instead.
     */
    abstract class Range extends DataFlow::Node {
      /** Gets the data-flow node representing the URL being redirected to. */
      abstract DataFlow::Node getUrl();

      /** Gets the response writer that this redirect is sent on, if any. */
      abstract ResponseWriter getResponseWriter();
    }

    /**
     * An assignment of the HTTP Location header, which indicates the location for a
     * redirect.
     */
    private class LocationHeaderSet extends Range, HeaderWrite {
      LocationHeaderSet() { this.getHeaderName() = "location" }

      override DataFlow::Node getUrl() { result = this.getValue() }

      override ResponseWriter getResponseWriter() { result = HeaderWrite.super.getResponseWriter() }
    }

    /**
     * An HTTP request attribute that is generally not attacker-controllable for
     * open redirect exploits; for example, a form field submitted in a POST request.
     */
    abstract class UnexploitableSource extends DataFlow::Node { }
  }

  /**
   * An HTTP redirect.
   *
   * Extend this class to refine existing API models. If you want to model new APIs,
   * extend `HTTP::Redirect::Range` instead.
   */
  class Redirect extends DataFlow::Node {
    Redirect::Range self;

    Redirect() { this = self }

    /** Gets the data-flow node representing the URL being redirected to. */
    DataFlow::Node getUrl() { result = self.getUrl() }

    /** Gets the response writer that this redirect is sent on, if any. */
    ResponseWriter getResponseWriter() { result = self.getResponseWriter() }
  }
}
