package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["ProjectGallery/controllers:MainController"] = append(beego.GlobalControllerRouter["ProjectGallery/controllers:MainController"],
        beego.ControllerComments{
            Method: "Ping",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
