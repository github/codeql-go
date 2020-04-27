/**
 * Provides classes for working with untrusted flow sources from the `github.com/gin-gonic/gin` package.
 */

import go

private module Gin {
  // "github.com/gin-gonic/gin".Context sources:
  ///: kind:struct
  private class GithubComGinGonicGinContextSource extends UntrustedFlowSource::Range, DataFlow::Node {
    GithubComGinGonicGinContextSource() {
      exists(string packagePath, string typeName |
        packagePath = "github.com/gin-gonic/gin" and
        typeName = "Context"
      |
        // Method calls:
        exists(DataFlow::MethodCallNode call, string methodName |
          call.getTarget().hasQualifiedName(packagePath, typeName, methodName) and
          (
            methodName = "GetHeader"
            or
            methodName = "QueryArray"
            or
            methodName = "Query"
            or
            methodName = "PostFormArray"
            or
            methodName = "PostForm"
            or
            methodName = "Param"
            or
            methodName = "GetStringSlice"
            or
            methodName = "GetString"
            or
            methodName = "GetRawData"
            or
            methodName = "ClientIP"
            or
            methodName = "ContentType"
            or
            methodName = "Cookie"
            or
            methodName = "GetQueryArray"
            or
            methodName = "GetQuery"
            or
            methodName = "GetPostFormArray"
            or
            methodName = "GetPostForm"
            or
            methodName = "DefaultPostForm"
            or
            methodName = "DefaultQuery"
            or
            methodName = "GetPostFormMap"
            or
            methodName = "GetQueryMap"
            or
            methodName = "GetStringMap"
            or
            methodName = "GetStringMapString"
            or
            methodName = "GetStringMapStringSlice"
            or
            methodName = "PostFormMap"
            or
            methodName = "QueryMap"
          )
        |
          // NOTE: getResult(0) will not work on calls with only one result.
          this = call.getResult(0)
          or
          // ... For calls with only one result, use getResult()
          this = call.getResult()
        )
        or
        // Field reads:
        exists(DataFlow::Field fld |
          fld.hasQualifiedName(packagePath, typeName, "Accepted") and
          this = fld.getARead()
        )
      )
    }
  }

  // "github.com/gin-gonic/gin".Params sources:
  ///: kind:slice
  private class GithubComGinGonicGinParamsSource extends UntrustedFlowSource::Range, DataFlow::Node {
    GithubComGinGonicGinParamsSource() {
      exists(string packagePath, string typeName |
        packagePath = "github.com/gin-gonic/gin" and
        typeName = "Params"
      |
        // Any read of a variable of this type:
        exists(DataFlow::ReadNode read | read.getType().hasQualifiedName(packagePath, typeName) |
          this = read
        )
        or
        // Method calls:
        exists(DataFlow::MethodCallNode call, string methodName |
          call.getTarget().hasQualifiedName(packagePath, typeName, methodName) and
          (
            methodName = "ByName"
            or
            methodName = "Get"
          )
        |
          this = call.getResult()
        )
      )
    }
  }

  // "github.com/gin-gonic/gin".Param sources:
  ///: kind:struct
  private class GithubComGinGonicGinParamSource extends UntrustedFlowSource::Range, DataFlow::Node {
    GithubComGinGonicGinParamSource() {
      exists(string packagePath, string typeName |
        packagePath = "github.com/gin-gonic/gin" and
        typeName = "Param"
      |
        // Any read of a variable of this type:
        exists(DataFlow::ReadNode read | read.getType().hasQualifiedName(packagePath, typeName) |
          this = read
        )
        or
        // Field reads:
        exists(DataFlow::Field fld, string fieldName |
          (
            fieldName = "Key"
            or
            fieldName = "Value"
          )
        |
          fld.hasQualifiedName(packagePath, typeName, fieldName) and
          this = fld.getARead()
        )
      )
    }
  }

  private class GithubComGinGonicGinContextBindSource extends UntrustedFlowSource::Range,
    DataFlow::Node {
    GithubComGinGonicGinContextBindSource() {
      exists(string packagePath, string typeName |
        packagePath = "github.com/gin-gonic/gin" and
        typeName = "Context"
      |
        exists(DataFlow::MethodCallNode call, string methodName |
          call.getTarget().hasQualifiedName(packagePath, typeName, methodName) and
          (
            methodName = "BindYAML" or
            methodName = "BindXML" or
            methodName = "BindWith" or
            methodName = "BindUri" or
            methodName = "BindQuery" or
            methodName = "MustBindWith" or
            methodName = "BindJSON" or
            methodName = "Bind" or
            methodName = "ShouldBind" or
            methodName = "ShouldBindBodyWith" or
            methodName = "ShouldBindJSON" or
            methodName = "ShouldBindQuery" or
            methodName = "ShouldBindUri" or
            methodName = "ShouldBindWith" or
            methodName = "ShouldBindXML" or
            methodName = "ShouldBindYAML"
          )
        |
          this = call.getArgument(0)
        )
      )
    }
  }
}
