package rest

import (
	"github.com/kimjooyoon/project_board_backend_go/pkg/app/article"
	"github.com/kimjooyoon/project_board_backend_go/pkg/app/article/controller"
	"github.com/kimjooyoon/project_board_backend_go/pkg/app/article/controller/request"
	"github.com/kimjooyoon/project_board_backend_go/pkg/app/article/controller/response"
)

var _ controller.ArticleController = (*ArticleRestQueryController)(nil)

type ArticleRestQueryController struct {
	queryService article.QueryService
}

func (a *ArticleRestQueryController) WriteArticle(req request.WriteArticleRequest) response.WriteArticleResponse {
	return response.WriteArticleResponse{}
}
