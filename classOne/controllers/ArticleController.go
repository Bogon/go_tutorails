package controllers

import (
	"classOne/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"path"
	"time"
)

type ArticleController struct {
	beego.Controller
}

/// 文章列表页
func (this *ArticleController) ShowArticleList()  {
	name := this.GetString("name")
	if name == "" {
		name = "李雷"
	}
	/// 1. 查询
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	var articles []models.Article
	qs.All(&articles)
	beego.Info(articles[0])
	/// 2. 把数据传递给试图显示
	this.Data["articles"] = articles
	this.Data["name"] = name
	this.TplName = "index.html"

}

func (this * ArticleController) ShowAddArticle()  {
	this.TplName = "add.html"
}

/*
	1、拿数据
	2、检验数据
	3、插入数据库
	4、返回view
*/
func (this *ArticleController) HandleAddArticle()  {
	/// 1、拿数据
	/// 1.1 获取标题
	title := this.GetString("articleName")
	/// 1.2 获取内容
	content := this.GetString("content")
	beego.Info(title, content)
	/// 1.3 上传图片
	file,header,err := this.GetFile("uploadname")
	defer file.Close()
	beego.Info(header.Filename)
	///	2、检验数据
	/// 2.1.1、判断文件格式
	ext := path.Ext(header.Filename)
	beego.Info(ext)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		beego.Info("上传图片格式不正确!")
		return
	}
	/// 2.1.2、判断文件大小
	if header.Size > 5000000 {
		beego.Info("上传文件太大，不允许上传!")
		return
	}

	/// 2.1.3、不允许重名
	filename := time.Now().Format("2006-01-02 15:04:05")
	err = this.SaveToFile("uploadname", "./static/img/" + filename + ext)
	if err != nil {
		beego.Info("文件上传失败：", err)
		return
	}
	beego.Info(filename + ext)

	///	3、插入数据库
	/// 3.1.1、获取orm对象
	o := orm.NewOrm()
	/// 3.1.2、创建插入对象
	m_article := models.Article{}
	m_article.Title = title
	m_article.Content = content
	m_article.Img = "/static/img/" + filename + ext
	/// 3.1.3、插入数据库
	_, err = o.Insert(&m_article)
	if err != nil {
		beego.Info("插入数据错误：", err)
		return
	}

	///	4、返回view
	this.Redirect("/ShowArticle", 302)

}

/// show article detial
func (this *ArticleController) ShowArticleDetail()  {

	/// 1. 获取数据
	id, err := this.GetInt("articleId")
	if err != nil {
		beego.Info("传递数据错误：", err)
		return
	}

	/// 2. 操作数据
	o := orm.NewOrm()
	var article = models.Article{}
	article.Id = id

	o.Read(&article)

	/// 3. 修改阅读量
	article.Count += 1
	o.Update(&article)

	this.Data["article"] = article
	this.TplName = "content.html"

}

/// delete article
func (this *ArticleController) HandleDeleteArticle()  {

	/// 1. 获取数据
	id, err := this.GetInt("articleId")
	if err != nil {
		beego.Info("传递数据错误：", err)
		return
	}

	/// 2. 操作数据
	o := orm.NewOrm()
	article := models.Article{Id: id}

	/// 3. 删除数据
	o.Delete(&article)

	/// 4. 返回试图
	this.Redirect("/ShowArticle", 302)
}

/// show article
func (this *ArticleController) ShowArticle()  {
	/// 1. 获取数据
	id, err := this.GetInt("articleId")
	if err != nil {
		beego.Info("传递数据错误：", err)
		return
	}

	/// 2. 操作数据
	o := orm.NewOrm()
	article := models.Article{Id: id}

	o.Read(&article)

	this.Data["article"] = article
	this.TplName = "update.html"

}

/// update article
func (this *ArticleController) HandleUpdateArticle()  {
	/// 1. 获取数据
	id, err := this.GetInt("articleId")
	if err != nil {
		beego.Info("传递数据错误：", err)
		return
	}

	/// 1、拿数据
	/// 1.1 获取标题
	title := this.GetString("articleName")
	/// 1.2 获取内容
	content := this.GetString("content")
	beego.Info(title, content)
	/// 1.3 上传图片
	file,header,err := this.GetFile("uploadname")
	defer file.Close()
	beego.Info(header.Filename)
	///	2、检验数据
	/// 2.1.1、判断文件格式
	ext := path.Ext(header.Filename)
	beego.Info(ext)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		beego.Info("上传图片格式不正确!")
		return
	}
	/// 2.1.2、判断文件大小
	if header.Size > 5000000 {
		beego.Info("上传文件太大，不允许上传!")
		return
	}

	/// 2.1.3、不允许重名
	filename := time.Now().Format("2006-01-02 15:04:05")
	this.SaveToFile("uploadname", "./static/img/" + filename + ext)
	if err != nil {
		beego.Info("文件上传失败：", err)
		return
	}
	beego.Info(filename + ext)

	/// 2. 操作数据
	o := orm.NewOrm()
	article := models.Article{Id: id}
	article.Title = title
	article.Content = content
	article.Img = "./static/img/" + filename + ext

	_, err = o.Update(&article)
	if err != nil {
		beego.Info("更新数据失败：", err)
		return
	}

	this.Redirect("/ShowArticle", 302)
}

