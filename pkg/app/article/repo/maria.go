package repo

import "project_board_backend/internal/article/domain"

type ArticleQuery interface {
	FindByName(name string) []domain.Article
}
