package article

type QueryService interface {
	SearchUser(name string) (SearchUserResponse, error)
}

type CommandService interface {
	SaveUser(req SaveUserRequest) error
}

type SaveUserRequest struct {
	Username string
	Password string
	Email    string
	Nickname string
	Memo     string
}

type SearchUserResponse struct {
	Total   int       `json:"total"`
	Current int       `json:"current"`
	List    []Article `json:"list"`
}

type Article struct {
	ID      string
	Title   string
	UserId  string
	Content string
}
