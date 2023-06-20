package service

import (
	"fmt"
	"gorm.io/gorm"
	"lld-tdd/models"
	"log"
	"regexp"
)

type userSignup interface {
	signup(signupType string) *models.User
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

func (e *EmailUser) signup(signupType string) *models.User {
	user := &models.User{Email: e.Email, Password: e.Password, SignupType: signupType}
	return user
}

func (f *FacebookUser) signup(signupType string) *models.User {
	user := &models.User{Email: f.Email, Password: f.Password, SignupType: signupType}
	return user
}

func (g *GoogleUser) signup(signupType string) *models.User {
	user := &models.User{Email: g.Email, Password: g.Password, SignupType: signupType}
	return user
}

func UserSignupFactory(user *models.User) userSignup {
	switch signupType := user.SignupType; signupType {
	case "Email":
		return &EmailUser{Email: user.Email, Password: user.Password}
	case "Facebook":
		return &FacebookUser{Email: user.Email, Password: user.Password}
	case "Google":
		return &GoogleUser{Email: user.Email, Password: user.Password}
	}
	return nil
}

func IsValidEmail(email string) bool {
	// Simple email validation using regular expression
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(regex).MatchString(email)
}

func CreateNewUser(db *gorm.DB, user *models.User) error {
	if user.Password == "" {
		return fmt.Errorf("%v", "Password cannot be empty")
	}
	if user.Email == "" || !IsValidEmail(user.Email) {
		return fmt.Errorf("%v", "Invalid email")
	}
	if user.SignupType != "Email" && user.SignupType != "Facebook" && user.SignupType != "Google" {
		return fmt.Errorf("%v", "Invalid signup type")
	}
	u := UserSignupFactory(user).signup(user.SignupType)
	//db, err := ConnectDB()
	//if err != nil {
	//	panic("Failed to connect to database")
	//}
	result := db.Create(u)
	if result.Error != nil {
		log.Print(result.Error.Error())
		return result.Error
	}
	return nil
}
