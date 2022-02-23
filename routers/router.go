package routers

import (
	"apiproject/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.ShortController{})
}
