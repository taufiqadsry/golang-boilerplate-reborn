package http

import "time"

type ArticleDetailResponse struct {
	Id	 			uint       	`json:"id"`
	Title	 		string     	`json:"title"`
	Body			string     	`json:"body"`
	AuthorId		string     	`json:"author_id"`
	CategoryId	 	*string     `json:"category_id"`
	CreatedAt	   	*time.Time 	`json:"created_at"`
}

type ArticleResponse struct {
	Id 				uint       	`json:"id"`
	Title 			string     	`json:"title"`
	CreatedAt    	*time.Time 	`json:"created_at"`
	UpdatedAt    	*time.Time 	`json:"updated_at"`
}

type ArticleRequest struct {
	Title 			string     	`json:"title"`
	Body 			string     	`json:"body"`
	AuthorId		uint		`json:"author_id"`
}