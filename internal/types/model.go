package types

type AuthRegisterForm struct {
	Email    string
	Name     string
	Password string
}

type AuthLoginForm struct {
	Email    string
	Password string
}
