package lib

func GetString() string {
	return "hello"
}

func GetUser() User {
	return User{
		Name: "foo",
		Id:   1,
	}
}
