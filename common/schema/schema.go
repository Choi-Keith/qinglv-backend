package schema

type EmailContent struct {
	UserId uint64 `json:"userId,string"`
	Email  string `json:"email"`
}
