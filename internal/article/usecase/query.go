package usecase

import (
	"project_board_backend/internal/app/article"
	"project_board_backend/internal/article/domain"
	"project_board_backend/internal/infrastructure/out/maria/article/query"
)

var _ article.QueryService = (*articleQueryUseCase)(nil)

func NewArticleQueryUseCase(q query.ArticleQuery) article.QueryService {
	return &articleQueryUseCase{query: q}
}

type articleQueryUseCase struct {
	query query.ArticleQuery
}

func (a *articleQueryUseCase) SearchUser(name string) (article.SearchUserResponse, error) {
	articles := a.query.FindByName(name)
	return convertSearchUserResponse(articles), nil
}

func convertSearchUserResponse(list []domain.Article) article.SearchUserResponse {
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
		List: l,
	}
}
