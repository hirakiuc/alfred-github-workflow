package model

import "github.com/google/go-github/github"

// User describe a github user.
type User struct {
	Login   string
	HTMLURL string
}

// ConvertUser convert github.User to User
func ConvertUser(user *github.User) User {
	return User{
		Login:   user.GetLogin(),
		HTMLURL: user.GetHTMLURL(),
	}
}
