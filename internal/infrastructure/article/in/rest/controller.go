package rest

import (
	"project_board_backend/pkg/app/article"
	"project_board_backend/pkg/app/article/controller"
	"project_board_backend/pkg/app/article/controller/request"
	"project_board_backend/pkg/app/article/controller/response"
)

var _ controller.ArticleController = (*ArticleRestQueryController)(nil)

type ArticleRestQueryController struct {
	queryService article.QueryService
}

func (a *ArticleRestQueryController) WriteArticle(req request.WriteArticleRequest) response.WriteArticleResponse {
	return response.WriteArticleResponse{}
}
