/**
 * Provides classes for working with untrusted flow sources from the `github.com/gin-gonic/gin` package.
 */

import go

// "github.com/gin-gonic/gin".Context sources:
///: kind:struct
class GithubComGinGonicGinContextSource extends UntrustedFlowSource::Range, DataFlow::Node {
  GithubComGinGonicGinContextSource() {
    exists(string packagePath, string typeName |
      packagePath = "github.com/gin-gonic/gin" and
      typeName = "Context"
    |
      // Any read of a variable of this type:
      exists(DataFlow::ReadNode read |
        read.getType().hasQualifiedName(packagePath, typeName)
      |
        this = read
      )
      or
      // Method calls:
      exists(DataFlow::MethodCallNode call, string methodName, int resultIndex |
        call.getTarget().hasQualifiedName(packagePath, typeName, methodName) and
        (
          methodName = "GetHeader"
            and
            (
              ///:out i:0 kind:string type:string
              resultIndex = 0
            )
          or
          methodName = "QueryArray"
            and
            (
              ///:out i:0 kind:slice type:[]string
              resultIndex = 0
            )
          or
          methodName = "Query"
            and
            (
              ///:out i:0 kind:string type:string
              resultIndex = 0
            )
          or
          methodName = "PostFormArray"
            and
            (
              ///:out i:0 kind:slice type:[]string
              resultIndex = 0
            )
          or
          methodName = "PostForm"
            and
            (
              ///:out i:0 kind:string type:string
              resultIndex = 0
            )
          or
          methodName = "Param"
            and
            (
              ///:out i:0 kind:string type:string
              resultIndex = 0
            )
          or
          methodName = "GetStringSlice"
            and
            (
              ///:out i:0 kind:slice type:[]string
              resultIndex = 0
            )
          or
          methodName = "GetString"
            and
            (
              ///:out i:0 kind:string type:string
              resultIndex = 0
            )
          or
          methodName = "GetRawData"
            and
            (
              ///:out i:0 kind:slice type:[]uint8
              resultIndex = 0
                or
              ///:out i:1 kind:interface type:error
              resultIndex = 1
            )
          or
          methodName = "ClientIP"
            and
            (
              ///:out i:0 kind:string type:string
              resultIndex = 0
            )
          or
          methodName = "ContentType"
            and
            (
              ///:out i:0 kind:string type:string
              resultIndex = 0
            )
          or
          methodName = "Cookie"
            and
            (
              ///:out i:0 kind:string type:string
              resultIndex = 0
            )
          or
          methodName = "GetQueryArray"
            and
            (
              ///:out i:0 kind:slice type:[]string
              resultIndex = 0
            )
          or
          methodName = "GetQuery"
            and
            (
              ///:out i:0 kind:string type:string
              resultIndex = 0
            )
          or
          methodName = "GetPostFormArray"
            and
            (
              ///:out i:0 kind:slice type:[]string
              resultIndex = 0
            )
          or
          methodName = "GetPostForm"
            and
            (
              ///:out i:0 kind:string type:string
              resultIndex = 0
            )
          or
          methodName = "DefaultPostForm"
            and
            (
              ///:out i:0 kind:string type:string
              resultIndex = 0
            )
          or
          methodName = "DefaultQuery"
            and
            (
              ///:out i:0 kind:string type:string
              resultIndex = 0
            )
          or
          methodName = "FormFile"
            and
            (
              ///:out i:0 kind:ptr type:"mime/multipart".FileHeader
              resultIndex = 0
            )
          or
          methodName = "Get"
            and
            (
              ///:out i:0 kind:interface type:interface {}
              resultIndex = 0
            )
          or
          methodName = "GetPostFormMap"
            and
            (
              ///:out i:0 kind:map type:map[string]string
              resultIndex = 0
            )
          or
          methodName = "GetQueryMap"
            and
            (
              ///:out i:0 kind:map type:map[string]string
              resultIndex = 0
            )
          or
          methodName = "GetStringMap"
            and
            (
              ///:out i:0 kind:map type:map[string]interface {}
              resultIndex = 0
            )
          or
          methodName = "GetStringMapString"
            and
            (
              ///:out i:0 kind:map type:map[string]string
              resultIndex = 0
            )
          or
          methodName = "GetStringMapStringSlice"
            and
            (
              ///:out i:0 kind:map type:map[string][]string
              resultIndex = 0
            )
          or
          methodName = "MultipartForm"
            and
            (
              ///:out i:0 kind:ptr type:"mime/multipart".Form
              resultIndex = 0
            )
          or
          methodName = "MustGet"
            and
            (
              ///:out i:0 kind:interface type:interface {}
              resultIndex = 0
            )
          or
          methodName = "PostFormMap"
            and
            (
              ///:out i:0 kind:map type:map[string]string
              resultIndex = 0
            )
          or
          methodName = "QueryMap"
            and
            (
              ///:out i:0 kind:map type:map[string]string
              resultIndex = 0
            )
        )
      |
        resultIndex = 0 and this = call.getResult()
        or
        this = call.getResult(resultIndex)
      )
      or
      // Field reads:
      exists(DataFlow::Field fld, string fieldName |
        (
          ///: kind:slice type:[]string
          fieldName = "Accepted"
            or
          ///: kind:slice type:"github.com/gin-gonic/gin".Params
          fieldName = "Params"
        )
      |
        fld.hasQualifiedName(packagePath, typeName, fieldName) and 
        this = fld.getARead()
      )
    )
  }
}

// "github.com/gin-gonic/gin".Params sources:
///: kind:slice
class GithubComGinGonicGinParamsSource extends UntrustedFlowSource::Range, DataFlow::Node {
  GithubComGinGonicGinParamsSource() {
    exists(string packagePath, string typeName |
      packagePath = "github.com/gin-gonic/gin" and
      typeName = "Params"
    |
      // Any read of a variable of this type:
      exists(DataFlow::ReadNode read |
        read.getType().hasQualifiedName(packagePath, typeName)
      |
        this = read
      )
      or
      // Method calls:
      exists(DataFlow::MethodCallNode call, string methodName, int resultIndex |
        call.getTarget().hasQualifiedName(packagePath, typeName, methodName) and
        (
          methodName = "ByName"
            and
            (
              ///:out i:0 kind:string type:string
              resultIndex = 0
            )
          or
          methodName = "Get"
            and
            (
              ///:out i:0 kind:string type:string
              resultIndex = 0
                or
              ///:out i:1 kind:bool type:bool
              resultIndex = 1
            )
        )
      |
        resultIndex = 0 and this = call.getResult()
        or
        this = call.getResult(resultIndex)
      )
    )
  }
}

///::: "github.com/gin-gonic/gin".Param
///: type:{"github.com/gin-gonic/gin".Context}/fields/Params/type:{"github.com/gin-gonic/gin".Params}/array/type:{"github.com/gin-gonic/gin".Param}

// ------------------------------------------------------------
// "github.com/gin-gonic/gin".Param sources:
///: kind:struct
class GithubComGinGonicGinParamSource extends UntrustedFlowSource::Range, DataFlow::Node {
  GithubComGinGonicGinParamSource() {
    exists(string packagePath, string typeName |
      packagePath = "github.com/gin-gonic/gin" and
      typeName = "Param"
    |
      // Any read of a variable of this type:
      exists(DataFlow::ReadNode read |
        read.getType().hasQualifiedName(packagePath, typeName)
      |
        this = read
      )
      or
      // Field reads:
      exists(DataFlow::Field fld, string fieldName |
        (
          ///: kind:string type:string
          fieldName = "Key"
            or
          ///: kind:string type:string
          fieldName = "Value"
        )
      |
        fld.hasQualifiedName(packagePath, typeName, fieldName) and 
        this = fld.getARead()
      )
    )
  }
}

class GithubComGinGonicGinContextBindSource extends UntrustedFlowSource::Range, DataFlow::Node {
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