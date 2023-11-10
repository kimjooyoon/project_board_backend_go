package repo

import "github.com/kimjooyoon/project_board_backend_go/internal/article/domain"

type ArticleQuery interface {
	FindByName(name string) ([]domain.Article, int)
}
