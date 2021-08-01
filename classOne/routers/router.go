package routers

import (
	"classOne/controllers"
	"github.com/astaxie/beego"
)

func init() {

    //beego.Router("/", &controllers.MainController{})

    /// register
    beego.Router("/register", &controllers.RegisterController{}, "get:ShowReg;post:HandleReg")

	/// login
	beego.Router("/", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")

    /// article
    beego.Router("/ShowArticle", &controllers.ArticleController{}, "get:ShowArticleList")
    /// add article
	beego.Router("/AddArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandleAddArticle")
    /// article detail
    beego.Router("/ShowArticleDetail", &controllers.ArticleController{}, "get:ShowArticleDetail")
    /// delete article
    beego.Router("/DeleteArticle", &controllers.ArticleController{}, "get:HandleDeleteArticle")
    /// update article
    beego.Router("/UpdateActicle", &controllers.ArticleController{}, "get:ShowArticle;post:HandleUpdateArticle")

}
