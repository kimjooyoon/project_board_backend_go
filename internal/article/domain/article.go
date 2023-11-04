package domain

var _ Article = (*article)(nil)

type Article interface {
}

type article struct {
	id      int64
	userId  int64
	title   string
	content string
}
