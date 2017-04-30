package main

import (
	_ "my_go_web/models"
	_ "my_go_web/routers" //init reuters
	_ "my_go_web/tools"   //init(register) html-template function ...
	_ "my_go_web/tools/spiders"

	"github.com/astaxie/beego"
)

func init() {

}

func main() {
	beego.Run()
}