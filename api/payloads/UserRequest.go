package payloads

type CreateRequest struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDetail struct {
	Username string `json:"username"`
	Authorize string `json:"authorized"`
}