package main

import (
	_ "lastchance/routers"
	"lastchance/filters"

	"github.com/astaxie/beego"
)

func main() {
	beego.InsertFilter("/lifecycle", beego.BeforeRouter,
	filters.BeforeRouterFilter)
	beego.InsertFilter("/lifecycle", beego.BeforeExec,
	filters.BeforeExecFilter)

	beego.InsertFilter("/lifecycle", beego.AfterExec,
	filters.AfterExecRouter)

	beego.InsertFilter("/lifecycle", beego.FinishRouter,
	filters.FinishRouterFilter)

	beego.SetStaticPath("/public", "static")
	beego.DelStaticPath("/static")
	beego.Run()	
}
