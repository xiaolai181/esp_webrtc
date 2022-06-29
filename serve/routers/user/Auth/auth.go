package auth

type Auth struct {
	Username string `valid:"required";json:"username";MaxSize:50`
	Password string `valid:"required";json:"password";MaxSize:50`
}
