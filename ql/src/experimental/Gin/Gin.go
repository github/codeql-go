package main

func main() {}

// gin
func ginHandler(c *gin.Context) {
	{
		header := c.GetHeader("key")
		_ = header
	}
	{
		params := c.QueryArray("key")
		_ = params
	}
	{
		keys := c.Keys
		_ = keys
	}

	c.BindYAML
	c.BindXML
	c.BindWith
	c.BindUri
	c.BindQuery
	c.MustBindWith
	c.BindJSON
	c.Bind
	c.ShouldBind
	c.ShouldBindBodyWith
	c.ShouldBindJSON
	c.ShouldBindQuery
	c.ShouldBindUri
	c.ShouldBindWith
	c.ShouldBindXML
	c.ShouldBindYAML
	//c.SaveUploadedFile
}
