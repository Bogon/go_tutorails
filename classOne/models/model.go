package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

/// 1.结构体，表设计
type User struct {
	Id int
	UserName string
	Passwd string
}

type Article struct {
	Id int `orm:"pk;auto"`
	Title string `orm:"size(20)"`	/// 文章标题
	Content string `orm:"size(500)"`	/// 文章内容
	Img string `orm:"size(50);null"`		/// 文章图片
	Type string	`orm:"size(20)"`	/// 文章类型
	Time time.Time `orm:"type(datetime);auto_now_add"`	/// 发布时间
	UpdateTime time.Time `orm:"type(datetime);auto_now"`	/// 发布时间
	Count int `orm:"default(0)"`		/// 阅读量

}

/// 2.初始化语句
func init()  {

	/// 1.注册表
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Article))

	/// 2.注册 mysql driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	/// 3.链接数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/newsWeb?charset=utf8")

	/// 4.生成表
	orm.RunSyncdb("default", false, true)

}
