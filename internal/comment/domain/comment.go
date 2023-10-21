package domain

type ID int64

type Comment struct {
	Id ID

	ArticleID       int64
	Commenter       int64 // userAccount
	ParentCommentID ID
}
