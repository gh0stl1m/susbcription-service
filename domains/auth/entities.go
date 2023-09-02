package auth

type LoginRequestBody struct {
  Username string `json:"login"`
  Password string `json:"password"`
}

type IAuthServices interface {
  GenerateToken(username string) (string, error)
}

