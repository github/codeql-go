package main

import (
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {}

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func use(val string) {
	exec.Command(val)
}

// gin
func ginHandleruse(ctx *gin.Context) {
	{
		val := ctx.GetHeader("key")
		use(val)
	}
	{
		val := ctx.QueryArray("key")
		use(val[0])
	}
	{
		val := ctx.Query("key")
		use(val)
	}
	{
		val := ctx.PostFormArray("key")
		use(val[0])
	}
	{
		val := ctx.PostForm("key")
		use(val)
	}
	{
		val := ctx.Param("key")
		use(val)
	}
	{
		val := ctx.GetStringSlice("key")
		use(val[0])
	}
	{
		val := ctx.GetString("key")
		use(val)
	}
	{
		val, _ := ctx.GetRawData()
		use(string(val))
	}
	{
		val := ctx.ClientIP()
		use(val)
	}
	{
		val := ctx.ContentType()
		use(val)
	}
	{
		val, _ := ctx.Cookie("key")
		use(val)
	}
	{
		val, _ := ctx.GetQueryArray("key")
		use(val[0])
	}
	{
		val, _ := ctx.GetQuery("key")
		use(val)
	}
	{
		val, _ := ctx.GetPostFormArray("key")
		use(val[0])
	}
	{
		val, _ := ctx.GetPostForm("key")
		use(val)
	}
	{
		val := ctx.DefaultPostForm("key", "default-value")
		use(val)
	}
	{
		val := ctx.DefaultQuery("key", "default-value")
		use(val)
	}
	{
		val, _ := ctx.Get("key")
		use(val.(string))
	}
	{
		val, _ := ctx.GetPostFormMap("key")
		use(val["a"])
	}
	{
		val, _ := ctx.GetQueryMap("key")
		use(val["a"])
	}
	{
		val := ctx.GetStringMap("key")
		use(val["a"].(string))
	}
	{
		val := ctx.GetStringMapString("key")
		use(val["a"])
	}
	{
		val := ctx.GetStringMapStringSlice("key")
		use(val["a"][0])
	}
	{
		val := ctx.MustGet("key")
		use(val.(string))
	}
	{
		val := ctx.PostFormMap("key")
		use(val["a"])
	}
	{
		val := ctx.QueryMap("key")
		use(val["a"])
	}

	// fields:

	{
		val := ctx.Accepted
		use(val[0])
	}
	{
		val := ctx.Params
		use(val[0].Value)
	}

	// Params:
	{
		val := ctx.Params[0]
		use(val.Value)
	}
	{
		val := ctx.Params.ByName("name")
		use(val)
	}
	{
		val, _ := ctx.Params.Get("name")
		use(val)
	}

	// Param:
	{
		param := ctx.Params[0]
		key := param.Key
		val := param.Value
		use(key)
		use(val)
	}

	// bind:
	{
		var person Person
		ctx.BindYAML(&person)
		_ = person
	}
	{
		var person Person
		ctx.BindXML(&person)
		use(person.Name)
	}
	{
		var person Person
		ctx.BindWith(&person, binding.YAML)
		use(person.Name)
	}
	{
		var person Person
		ctx.BindUri(&person)
		use(person.Name)
	}
	{
		var person Person
		ctx.BindQuery(&person)
		use(person.Name)
	}
	{
		var person Person
		ctx.MustBindWith(&person, binding.YAML)
		use(person.Name)
	}
	{
		var person Person
		ctx.BindJSON(&person)
		use(person.Name)
	}
	{
		var person Person
		ctx.Bind(&person)
		use(person.Name)
	}
	{
		var person Person
		ctx.ShouldBind(&person)
		use(person.Name)
	}
	{
		var person Person
		ctx.ShouldBindBodyWith(&person, binding.YAML)
		use(person.Name)
	}
	{
		var person Person
		ctx.ShouldBindJSON(&person)
		use(person.Name)
	}
	{
		var person Person
		ctx.ShouldBindQuery(&person)
		use(person.Name)
	}
	{
		var person Person
		ctx.ShouldBindUri(&person)
		use(person.Name)
	}
	{
		var person Person
		ctx.ShouldBindWith(&person, binding.YAML)
		use(person.Name)
	}
	{
		var person Person
		ctx.ShouldBindXML(&person)
		use(person.Name)
	}
	{
		var person Person
		ctx.ShouldBindYAML(&person)
		use(person.Name)
	}
}
