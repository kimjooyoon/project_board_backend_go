package domain

var _ Comment = (*comment)(nil)

type Comment interface {
}

type comment struct {
	id int64

	articleID       int64
	commenter       int64 // userAccount
	parentCommentID int64
}
