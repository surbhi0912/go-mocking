package service

import (
	"lld-tdd/models"
)

type userSignup interface {
	signup(signupType string) error
}

func createInDB(user *models.User) error {
	db, err := ConnectDB()
	if err != nil {
		panic("Failed to connect to database")
	}
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
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

func (e *EmailUser) signup(signupType string) error {
	user := &models.User{Email: e.Email, Password: e.Password, SignupType: signupType}
	err := createInDB(user)
	if err != nil {
		return err
	}
	return nil
}

func (f *FacebookUser) signup(signupType string) error {
	user := &models.User{Email: f.Email, Password: f.Password, SignupType: signupType}
	err := createInDB(user)
	if err != nil {
		return err
	}
	return nil
}

func (g *GoogleUser) signup(signupType string) error {
	user := &models.User{Email: g.Email, Password: g.Password, SignupType: signupType}
	err := createInDB(user)
	if err != nil {
		return err
	}
	return nil
}

func UserSignupFactory(user *models.User) userSignup {
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

func CreateNewUser(user *models.User) error {
	err := UserSignupFactory(user).signup(user.SignupType)
	if err != nil {
		return err
	}
	return nil
}
