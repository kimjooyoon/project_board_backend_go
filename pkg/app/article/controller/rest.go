package controller

import (
	"github.com/kimjooyoon/project_board_backend_go/pkg/app/article/controller/request"
	"github.com/kimjooyoon/project_board_backend_go/pkg/app/article/controller/response"
)

type ArticleController interface {
	WriteArticle(request.WriteArticleRequest) response.WriteArticleResponse
}
