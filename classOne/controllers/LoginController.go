package controllers

import (
	"classOne/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) ShowLogin()  {
	this.TplName = "login.html"
}

/*
	1、拿到浏览器数据

	2、数据处理

	3、查找数据库

	4、返回视图
*/
func (this *LoginController) HandleLogin()  {

	/// 1、拿到浏览器数据
	name := this.GetString("userName")
	password := this.GetString("password")
	//beego.Info(name, password)

	/// 2、数据处理
	if name == "" || password == "" {
		beego.Info("用户名或密码不能为空！")
		this.TplName = "login.html"
		return
	}
	/// 2.1、获取orm对象
	o := orm.NewOrm()
	/// 2.2、获取插入对象
	user := models.User{}
	/// 2.3、查询数据
	user.UserName = name
	err := o.Read(&user, "UserName")
	if err != nil {
		beego.Info("用户名错误！")
		this.TplName = "login.html"
		return
	}
	/// 2.4、检验数据正确性
	if password != user.Passwd {
		beego.Info("登录失败！")
		this.TplName = "login.html"
		return
	}
	/// 2.5、返回视图
	this.Data["name"] = name
	this.Redirect("/ShowArticle?name="+name, 302)
}