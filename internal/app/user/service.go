package user

import "project_board_backend/internal/user/domain"

type QueryService interface {
	Search(name string) (*domain.Account, error)
}

type CommandService interface {
	Save(req SaveRequest) error
}

type SaveRequest struct {
	Username string
	Password string
	Email    string
	Nickname string
	Memo     string
}
