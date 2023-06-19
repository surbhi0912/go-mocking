package service

import (
	"lld-tdd/models"
)

type userSignup interface {
	signup(signupType string) models.User
}

func createInDB(user models.User) models.User {
	DBInstance.Create(&user)
	//log.Print(user)
	return user
}

type FacebookUser struct {
	Email    string
	Password string
}

type GoogleUser struct {
	Email    string
	Password string
}

type EmailUser struct {
	Email    string
	Password string
}

func (e *EmailUser) signup(signupType string) models.User {
	user := models.User{Email: e.Email, Password: e.Password, SignupType: signupType}
	return createInDB(user)
}

func (f *FacebookUser) signup(signupType string) models.User {
	user := models.User{Email: f.Email, Password: f.Password, SignupType: signupType}
	return createInDB(user)
}

func (g *GoogleUser) signup(signupType string) models.User {
	user := models.User{Email: g.Email, Password: g.Password, SignupType: signupType}
	return createInDB(user)
}

func UserSignupFactory(user models.User) userSignup {
	switch signupType := user.SignupType; signupType {
	case "Email":
		return &EmailUser{Email: user.Email, Password: user.Password}
	case "Facebook":
		return &FacebookUser{Email: user.Email, Password: user.Password}
		break
	case "Google":
		return &GoogleUser{Email: user.Email, Password: user.Password}
		break
	}
	return nil
}

func CreateNewUser(user models.User) string {
	u := UserSignupFactory(user).signup(user.SignupType)
	if u.ID < models.MinUserID {
		return "Email already exists"
	} else {
		return "Successfully signed up user"
	}
}
