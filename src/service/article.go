package services

import(
	"fmt"
	"time"
	"github.com/jinzhu/copier"
	httpEntity "example_app/entity/http"
	dbEntity "example_app/entity/db"
	repository "example_app/repository/db"
)

type ArticleService struct {
	articleRepository repository.ArticleRepositoryInterface
}

func ArticleServiceHandler() *ArticleService {
	return &ArticleService {
		articleRepository: repository.ArticleRepositoryHandler(),
	}
}

type ArticleServiceInterface interface {
	GetArticleByID(id int) *httpEntity.ArticleDetailResponse
	GetArticles(limit int, offset int) []httpEntity.ArticleResponse
	UpdateArticleByID(id int, payload httpEntity.ArticleRequest) bool
	Delete(id int) *httpEntity.ArticleResponse
	StoreArticle(payload httpEntity.ArticleRequest) bool
}

func (service *ArticleService) GetArticles(page int, count int) []httpEntity.ArticleResponse{
	data, _ := service.articleRepository.GetArticlesList(page, count)
	result := []httpEntity.ArticleResponse{}
	copier.Copy(&result, &data)
	return result
}

func (service *ArticleService) GetArticleByID(id int) *httpEntity.ArticleDetailResponse {
	article := &dbEntity.Article{}
	service.articleRepository.GetArticleByID(id, article)

	result := &httpEntity.ArticleDetailResponse{}
	if article != nil {
		copier.Copy(result, article)
	}
	return result
}

func (service *ArticleService) UpdateArticleByID(id int, payload httpEntity.ArticleRequest) bool {
	now := time.Now()
	article := &dbEntity.Article{
		Title: payload.Title,
		Body: payload.Body,
		AuthorId: payload.AuthorId,
		UpdatedAt: &now,
	}
	err := service.articleRepository.UpdateArticleById(id, article)
	if nil != err {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func (service *ArticleService) Delete(id int) *httpEntity.ArticleResponse {
	article := dbEntity.Article{}
	result := service.articleRepository.Delete(id, &article)

	output := &httpEntity.ArticleResponse{}
	if result == nil {
		copier.Copy(output, article)
	}
	return output
}

func (service *ArticleService) StoreArticle(payload httpEntity.ArticleRequest) bool {
	now := time.Now()
	article := &dbEntity.Article{
		Title: payload.Title,
		Body: payload.Body,
		AuthorId: payload.AuthorId,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	err := service.articleRepository.StoreArticle(article)
	if nil != err {
		fmt.Println(err.Error())
		return false
	}
	return true
}
