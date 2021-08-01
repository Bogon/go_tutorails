package controllers

import (
	"classOne/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) ShowReg()  {
	this.TplName = "register.html"
}

/*
	1、获取浏览器传递过来的数据

	2、处理传递过来的数据

	3、插入数据库(User)

	4、返回试图
*/
func (this *RegisterController) HandleReg() {

	/// 1、获取浏览器传递过来的数据
	name := this.GetString("userName")
	password := this.GetString("password")
	beego.Info(name, password)

	/// 2、处理传递过来的数据
	if name == "" || password == "" {
		beego.Info("用户名或密码不能为空！")
		this.TplName = "register.html"
		return
	}

	/// 3、插入数据库(User)
	/// 3.1、获取orm对象
	o := orm.NewOrm()
	/// 3.2、获取插入对象
	user := models.User{}
	user.UserName = name
	user.Passwd = password
	/// 3.3、插入操作
	_,err := o.Insert(&user)
	if err != nil {
		beego.Info("插入数据失败：", err)
		return
	}

	/// 3.4、返回登录
	//this.Ctx.WriteString("注册成功！")
	//this.TplName = "login.html"
	this.Redirect("/", 302)
}