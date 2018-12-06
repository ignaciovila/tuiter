package user

var userList []*User

func AddUser(user *User) {
	userList = append(userList, user)
}

func GetUsers() []*User {
	return userList
}
