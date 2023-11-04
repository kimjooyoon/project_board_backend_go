package domain

var _ Hashtag = (*hashtag)(nil)

type Hashtag interface {
}

type hashtag struct {
	id   int64
	name string
}
