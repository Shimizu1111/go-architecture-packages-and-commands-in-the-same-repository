package auth

func Authenticate(user, pass string) bool {
	return user == "user" && pass == "pass"
}
