package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	a := c.GetString("para1")
	b := c.GetString("para2")
	d, _ := c.GetInt("para3")
	e := c.GetString("sign")
	f := c.GetString("sign2")
	if a == "" && b == "" {
		c.Ctx.WriteString("please input para")
		fmt.Print(d)
		return
	}else {
		message_1 := a + b
		message_2 := []byte(message_1)
		hashCalc := sha256.New()
		hashCalc.Write(message_2)
		bytes := hashCalc.Sum(nil)
		hashCode := hex.EncodeToString(bytes)
		if hashCode == e {
			c.Ctx.WriteString(hashCode)
		}
		if f == "123123" {
			c.Ctx.WriteString("just a test with sign")
		}
		
	}
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}
