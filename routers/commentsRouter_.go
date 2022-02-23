package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["apiproject/controllers:ShortController"] = append(beego.GlobalControllerRouter["apiproject/controllers:ShortController"],
		beego.ControllerComments{
			Method:           "View",
			Router:           `/a/:url`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["apiproject/controllers:ShortController"] = append(beego.GlobalControllerRouter["apiproject/controllers:ShortController"],
		beego.ControllerComments{
			Method:           "ToShortUrl",
			Router:           `/short`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
