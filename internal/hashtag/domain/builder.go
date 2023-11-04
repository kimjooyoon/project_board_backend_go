package domain

var _ Builder = (*builder)(nil)

type Builder interface {
	Id(id int64) Builder
	Name(name string) Builder
	Build() Hashtag
}

type builder struct {
	id   int64
	name string
}

func (b *builder) Id(id int64) Builder {
	b.id = id
	return b
}

func (b *builder) Name(name string) Builder {
	b.name = name
	return b
}

func (b *builder) Build() Hashtag {
	return hashtag{
		id:   b.id,
		name: b.name,
	}
}
