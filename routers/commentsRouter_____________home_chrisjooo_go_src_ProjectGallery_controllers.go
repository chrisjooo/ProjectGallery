package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["ProjectGallery/controllers:AccountController"] = append(beego.GlobalControllerRouter["ProjectGallery/controllers:AccountController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectGallery/controllers:AccountController"] = append(beego.GlobalControllerRouter["ProjectGallery/controllers:AccountController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectGallery/controllers:AccountController"] = append(beego.GlobalControllerRouter["ProjectGallery/controllers:AccountController"],
        beego.ControllerComments{
            Method: "GetByUsername",
            Router: `/:username`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectGallery/controllers:AccountController"] = append(beego.GlobalControllerRouter["ProjectGallery/controllers:AccountController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:username`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectGallery/controllers:AccountController"] = append(beego.GlobalControllerRouter["ProjectGallery/controllers:AccountController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:username`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectGallery/controllers:AccountController"] = append(beego.GlobalControllerRouter["ProjectGallery/controllers:AccountController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectGallery/controllers:MainController"] = append(beego.GlobalControllerRouter["ProjectGallery/controllers:MainController"],
        beego.ControllerComments{
            Method: "Ping",
            Router: `/ping`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectGallery/controllers:ProjectController"] = append(beego.GlobalControllerRouter["ProjectGallery/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectGallery/controllers:ProjectController"] = append(beego.GlobalControllerRouter["ProjectGallery/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectGallery/controllers:ProjectController"] = append(beego.GlobalControllerRouter["ProjectGallery/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectGallery/controllers:ProjectController"] = append(beego.GlobalControllerRouter["ProjectGallery/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetProjectsByName",
            Router: `/:name`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectGallery/controllers:ProjectController"] = append(beego.GlobalControllerRouter["ProjectGallery/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetById",
            Router: `/id/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectGallery/controllers:VoteController"] = append(beego.GlobalControllerRouter["ProjectGallery/controllers:VoteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectGallery/controllers:VoteController"] = append(beego.GlobalControllerRouter["ProjectGallery/controllers:VoteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectGallery/controllers:VoteController"] = append(beego.GlobalControllerRouter["ProjectGallery/controllers:VoteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ProjectGallery/controllers:VoteController"] = append(beego.GlobalControllerRouter["ProjectGallery/controllers:VoteController"],
        beego.ControllerComments{
            Method: "GetProjectVote",
            Router: `/:projectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
