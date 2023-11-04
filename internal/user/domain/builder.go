package domain

var _ Builder = (*builder)(nil)

type Builder interface {
	ID(string) Builder
	Password(string) Builder
	Email(string) Builder
	Nickname(string) Builder
	Memo(string) Builder
	Build() Account
}

type builder struct {
	id       string
	password string
	email    string
	nickname string
	memo     string
}

func (b *builder) ID(s string) Builder {
	b.id = s
	return b
}

func (b *builder) Password(s string) Builder {
	b.password = s
	return b
}

func (b *builder) Email(s string) Builder {
	b.email = s
	return b
}

func (b *builder) Nickname(s string) Builder {
	b.nickname = s
	return b
}

func (b *builder) Memo(s string) Builder {
	b.memo = s
	return b
}

func (b *builder) Build() Account {
	return account{
		id:       b.id,
		password: b.password,
		email:    b.email,
		nickname: b.nickname,
		memo:     b.memo,
	}
}
