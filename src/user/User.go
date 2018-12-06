package userManager

type User struct {
	Name     string
	Mail     string
	Nick     string
	Password string
}

func NewUser(name string, mail string, nick string, password string) *User {
	return &User{name, mail, nick, password}
}
