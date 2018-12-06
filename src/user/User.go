package user

type User struct {
	Nombre   string
	Mail     string
	Nick     string
	Password string
}

func NewUser(nombre string, mail string, nick string, password string) *User {
	return &User{nombre, mail, nick, password}
}
