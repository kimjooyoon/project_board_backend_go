package domain

var _ Builder = (*builder)(nil)

type Builder interface {
	Id(id int64) Builder
	ArticleID(articleID int64) Builder
	Commenter(commenter int64) Builder
	ParentCommentID(parentCommentID int64) Builder
	Build() Comment
}

type builder struct {
	id int64

	articleID       int64
	commenter       int64 // userAccount
	parentCommentID int64
}

func (b *builder) Id(id int64) Builder {
	b.id = id
	return b
}

func (b *builder) ArticleID(articleID int64) Builder {
	b.articleID = articleID
	return b
}

func (b *builder) Commenter(commenter int64) Builder {
	b.commenter = commenter
	return b
}

func (b *builder) ParentCommentID(parentCommentID int64) Builder {
	b.parentCommentID = parentCommentID
	return b
}

func (b *builder) Build() Comment {
	return comment{
		id:              b.id,
		articleID:       b.articleID,
		commenter:       b.commenter,
		parentCommentID: b.parentCommentID,
	}
}
