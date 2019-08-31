package controller

import (
	"strconv"
	// "sync"
	"net/http"
	"github.com/gin-gonic/gin"
	services "example_app/service"
	httpEntity "example_app/entity/http"
)
type ArticleController struct {
	ArticleService services.ArticleServiceInterface
}

func (service *ArticleController) GetArticles(context *gin.Context) {
	params := Limitofset{}
	err := context.ShouldBindQuery(&params)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	
	result := service.ArticleService.GetArticles(params.Limit, params.Offset)
	context.JSON(http.StatusOK, result)
}

func (service *ArticleController) GetArticleByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	result := service.ArticleService.GetArticleByID(id)
	if result == nil {
		context.JSON(http.StatusOK, gin.H{})
		return
	}
	context.JSON(http.StatusOK, result)
}

func (service *ArticleController) UpdateArticleByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	payload := httpEntity.ArticleRequest{}
	if err := context.ShouldBind(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	success := service.ArticleService.UpdateArticleByID(id,payload)

	if !success {
		context.JSON(http.StatusNoContent, gin.H{})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{})
}

func (service *ArticleController) Delete(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	
	result := service.ArticleService.Delete(id)

	if result.Id == 0 {
		context.JSON(http.StatusNoContent, gin.H{})
		return
	}

	context.JSON(http.StatusOK, result)
}

func (service *ArticleController) StoreArticle(context *gin.Context) {
	payload := httpEntity.ArticleRequest{}
	if err := context.ShouldBind(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	success := service.ArticleService.StoreArticle(payload)
	if !success {
		context.JSON(http.StatusNoContent, gin.H{})
		return
	}

	context.JSON(http.StatusCreated, gin.H{})
	
}