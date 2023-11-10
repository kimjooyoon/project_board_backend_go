package controller

import (
	"project_board_backend/pkg/app/article/controller/request"
	"project_board_backend/pkg/app/article/controller/response"
)

type ArticleController interface {
	WriteArticle(request.WriteArticleRequest) response.WriteArticleResponse
}
