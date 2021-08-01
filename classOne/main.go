package main

import (
	_ "classOne/models"
	_ "classOne/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

