package main

import (
	"github.com/astaxie/beego"
	_ "lastchance/routers"
)

func main() {
	beego.SetStaticPath("/public", "static")
	beego.DelStaticPath("/static")
	beego.Run()
}
