package adapters

type User struct {
	CPF      string `json:"cpf"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Credentials struct {
	Password string `json:"password"`
	CPF      string `json:"cpf"`
}

type AuthResult struct {
	AccessToken *string `json:"access_token" type:"string" sensitive:"true"`
	ExpiresIn   *int64  `json:"expires_in" type:"integer"`
	TokenType   *string `json:"token_type" type:"string"`
}
