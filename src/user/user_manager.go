package userManager

var users []*User

func AddUser(user *User) {
	users = append(users, user)
}

func GetUsers() []*User {
	return users
}

func ExistsUser(user string) bool {
	exists := false

	for i := 0;  i < len(users); i++ {
		if users[i].Nick == user {
			exists = true
		}
	}

	return exists
}