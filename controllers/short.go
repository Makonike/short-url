package controllers

import (
	"apiproject/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ShortController struct {
	beego.Controller
}

func (c *ShortController) URLMapping() {
	c.Mapping("ToShortUrl", c.ToShortUrl)
	c.Mapping("View", c.View)
}

// ToShortUrl @Title Post
// @Description change long url to short
// @Param	url(request-body): the long url
// @Success 200 {shortUrl} short
// @Failure 403 :url is empty
// @router /short [post]
func (c *ShortController) ToShortUrl() {
	var shortUrl models.Short
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &shortUrl)
	if err != nil {
		return
	}
	logs.Info(shortUrl)
	short := models.ToShort(shortUrl.Lurl)
	logs.Info(short)
	c.Data["json"] = map[string]string{"short": short}
	c.ServeJSON()
}

// View @Title Get
// @Description get the origin url and redirect to it
// @Param	url: the short-url
// @Success 302 redirect to the des url
// @Failure 403 :url is empty
// @router /a/:url [get]
func (c *ShortController) View() {
	url := c.Ctx.Input.Param(":url")
	//logs.Info(url)
	// get the long url
	longUrl := models.GetShort(url)
	// set 302 code then redirect
	//logs.Info(longUrl)
	c.Redirect(longUrl, 302)
}
