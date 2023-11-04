package domain

import (
	"project_board_backend/internal/article/domain/article_vo"
	"strconv"
)

var _ Article = (*article)(nil)

type Article interface {
	ToArticleInfo() article_vo.Info
}

type article struct {
	id      int64  // 게시판 식별자
	userId  int64  // 고객 식별자
	title   string // 게시판 제목
	content string // 게시판 내용
}

func (a article) ToArticleInfo() article_vo.Info {
	return article_vo.Info{
		ID:      strconv.FormatInt(a.id, 10),
		Title:   a.title,
		UserId:  strconv.FormatInt(a.userId, 10),
		Content: a.content,
	}
}
