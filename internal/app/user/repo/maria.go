package repo

import "github.com/kimjooyoon/project_board_backend_go/internal/user/domain"

type QueryUserData interface {
	FindById(name string) []domain.Account
}

type CommandUserData interface {
	Save(aggregate domain.Account) error
}
