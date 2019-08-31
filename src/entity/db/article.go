package db

import "time"

type Article struct {
	Id 				uint       	`gorm:"primary_key" json:"id"`
	Title 			string     	`gorm:"column:title" json:"title"`
	Body 			string     	`gorm:"column:body" json:"body"`
	AuthorId 		uint 		`gorm:"column:author_id"`	
	CategoryId		uint 		`gorm:"column:category_id"`	
	CreatedAt    	*time.Time 	`gorm:"column:created_at" json:"created_at"`
	UpdatedAt    	*time.Time 	`gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    	*time.Time 	`gorm:"column:deleted_at" json:"deleted_at"`
	Author			*User		`gorm:"auto_preload, foreignkey:AuthorId, association_foreignkey:ID"`
}

func (Article) TableName() string {
	return "articles"
}