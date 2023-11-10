package usecase

import (
	"github.com/kimjooyoon/project_board_backend_go/internal/article/domain"
	"github.com/kimjooyoon/project_board_backend_go/pkg/app/article"
	"github.com/kimjooyoon/project_board_backend_go/pkg/app/article/repo"
)

var _ article.QueryService = (*articleQueryUseCase)(nil)

func NewArticleQueryUseCase(q repo.ArticleQuery) article.QueryService {
	return &articleQueryUseCase{query: q}
}

type articleQueryUseCase struct {
	query repo.ArticleQuery
}

func (a *articleQueryUseCase) SearchUser(name string) (article.SearchUserResponse, error) {
	articles, total := a.query.FindByName(name)

	return convertSearchUserResponse(articles, total, len(articles)), nil
}

func convertSearchUserResponse(list []domain.Article, total, current int) article.SearchUserResponse {
	l := make([]article.Article, len(list))
	for i, d := range list {
		var vo = d.ToArticleInfo()
		l[i] = article.Article{
			ID:      vo.ID,
			Title:   vo.Title,
			UserId:  vo.UserId,
			Content: vo.Content,
		}
	}

	return article.SearchUserResponse{
		Total:   total,
		Current: current,
		List:    l,
	}
}
