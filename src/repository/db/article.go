// repository contains query
package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	connection "example_app/util/helper/mysqlconnection"
	dbEntity "example_app/entity/db"
)

type ArticleRepository struct {
	DB gorm.DB
}

type ArticleRepositoryInterface interface {
	GetArticleByID(id int, articleData *dbEntity.Article) error
	GetArticlesList(limit int, offset int) ([]dbEntity.Article, error)
	UpdateArticleById(id int, articleData *dbEntity.Article) error
	Delete(id int, articleData *dbEntity.Article) error
	StoreArticle(articleData *dbEntity.Article) error
}

// handler is default action called when we import this package
func ArticleRepositoryHandler() *ArticleRepository {
	return &ArticleRepository{ DB: *connection.GetConnection() }
}

func (repository *ArticleRepository) GetArticlesList(limit int, offset int) ([]dbEntity.Article, error){
	data := []dbEntity.Article{}
	query := repository.DB.Table("article")
	query = query.Limit(limit).Offset(offset)
	query = query.Find(&data)
	return data, query.Error
}

func (repository *ArticleRepository) GetArticleByID(id int, articleData *dbEntity.Article) error {
	query := repository.DB.Preload("Article")
	query = query.Where("id=?",id)
	query = query.First(articleData)
	return query.Error
} 

func (repository *ArticleRepository) UpdateArticleById(id int, articleData *dbEntity.Article) error {
	query := repository.DB.Table("articles")
	query = query.Where("id=?", id)
	success := query.Updates(articleData).RowsAffected
	if success < 1 {
		return errors.New("No data affected")
	}
	return query.Error
}

func (repository *ArticleRepository) Delete(id int, articleData *dbEntity.Article) error {
	article := &dbEntity.Article{}
	query := repository.DB.Table("articles")
	query = query.Where("id=?", id)
	query = query.First(articleData)
	query = query.Delete(article)
	return query.Error
}

func (repository *ArticleRepository) StoreArticle(articleData *dbEntity.Article) error {
	query := repository.DB.Table("articles")
	query = query.Create(articleData)
	return query.Error
}