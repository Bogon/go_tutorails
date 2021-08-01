package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/// 1.结构体，表设计
type UserInfo struct {
	Id int
	Name string
}

/// 2.初始化语句
func init()  {

	/// 1.注册表
	orm.RegisterModel(new(UserInfo))

	/// 2.注册 mysql driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	/// 3.链接数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/class1?charset=utf8")

	/// 4.生成表
	orm.RunSyncdb("default", false, true)

}
