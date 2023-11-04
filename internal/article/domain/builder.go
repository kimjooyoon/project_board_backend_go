package domain

var _ Builder = (*builder)(nil)

func NewArticleBuilder() Builder {
	return &builder{}
}

type Builder interface {
	ID(int64) Builder
	UserId(int64) Builder
	Title(string) Builder
	Content(string) Builder
	Build() Article
}

type builder struct {
	id      int64
	userId  int64
	title   string
	content string
}

func (b *builder) ID(i int64) Builder {
	b.id = i
	return b
}

func (b *builder) UserId(i int64) Builder {
	b.userId = i
	return b
}

func (b *builder) Title(s string) Builder {
	b.title = s
	return b
}

func (b *builder) Content(s string) Builder {
	b.content = s
	return b
}

func (b *builder) Build() Article {
	return article{
		id:      b.id,
		userId:  b.userId,
		title:   b.title,
		content: b.content,
	}
}
