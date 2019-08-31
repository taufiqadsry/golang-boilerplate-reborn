package controller

import (
	"github.com/gin-gonic/gin"
	srv "example_app/service"
)

func LoadRouter(routers *gin.Engine) {
	user := &UserRouterLoader{}
	article := &ArticleRouterLoader{}
	user.UserRouter(routers)
	article.ArticleRouter(routers)
}

type UserRouterLoader struct{
}

type ArticleRouterLoader struct{
}

func (rLoader *UserRouterLoader) UserRouter(router *gin.Engine) {
	handler := &UserController{
		UserService: srv.UserServiceHandler(),
	}
	rLoader.routerDefinition(router,handler)
}

func (rLoader *ArticleRouterLoader) ArticleRouter(router *gin.Engine) {
	handler := &ArticleController{
		ArticleService: srv.ArticleServiceHandler(),
	}
	rLoader.routerDefinition(router,handler)
}

func (rLoader *UserRouterLoader) routerDefinition(router *gin.Engine,handler *UserController) {
	group := router.Group("v1/users")
	group.GET("", handler.GetUsers)
	group.GET(":id", handler.GetUserByID)
	group.PUT(":id", handler.UpdateUsersByID)
}

func (rLoader *ArticleRouterLoader) routerDefinition(router *gin.Engine,handler *ArticleController) {
	group := router.Group("v1/article")
	group.GET("", handler.GetArticles)
	group.GET(":id", handler.GetArticleByID)
	group.POST("", handler.StoreArticle)
	group.PUT(":id", handler.UpdateArticleByID)
	group.DELETE(":id", handler.Delete)
}
