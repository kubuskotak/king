package rest

import "github.com/kubuskotak/king/pkg/entity"

// ListArticlesRequest Get all articles request.
type ListArticlesRequest struct {
	entity.Filter     `json:"filter"`
	entity.Pagination `json:"pagination"`
}

// ListArticlesResponse Get all articles response.
type ListArticlesResponse struct {
	Articles []*entity.Article
}

// GetArticleRequest Get an article request.
type GetArticleRequest struct {
	entity.Keys
}

// GetArticleResponse Get an article response.
type GetArticleResponse struct {
	entity.Article
}

// SaveArticleRequest Store article request.
type SaveArticleRequest struct {
	entity.Article
}

// SaveArticleResponse Store article response.
type SaveArticleResponse struct {
	entity.Article
}

// DeleteArticleRequest Remove an article request.
type DeleteArticleRequest struct {
	entity.Keys
}

// DeleteArticleResponse Remove an article response.
type DeleteArticleResponse struct {
	Message string
}
