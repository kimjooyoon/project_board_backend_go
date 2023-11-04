package domain

var _ Account = (*account)(nil)

type Account interface {
}

type account struct {
	id       string
	password string
	email    string
	nickname string
	memo     string
}
