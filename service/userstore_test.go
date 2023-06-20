package service

import (
	"lld-tdd/models"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateNewUser(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery(".*").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.33"))

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open GORM connection: %v", err)
	}

	// Create a new user to save
	newUsers := []models.User{
		models.User{
			Email:      "abc@gmail.com",
			Password:   "abc12",
			SignupType: "Facebook",
		},
		models.User{
			Email:      "def@gmail.com",
			Password:   "def12",
			SignupType: "Google",
		},
		models.User{
			Email:      "ghi@gmail.com",
			Password:   "ghi12",
			SignupType: "Email",
		},
	}

	for _, user := range newUsers {
		// Mock the expected query and define its behaviour
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `users`").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err = CreateNewUser(gormDB, &user)
		if err != nil {
			t.Fatalf("Failed to save user: %v", err)
		}

		err = mock.ExpectationsWereMet()
		if err != nil {
			t.Errorf("Unfulfilled expectations")
		}
	}

}

func TestCreateNewUser_InvalidSignupType(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create a mock database: %v", err)
	}
	defer db.Close()

	//Expect the query for the database version
	mock.ExpectQuery(".*").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.33"))

	// Initialize a new GORM database connection with the mock database
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open GORM connection: %v", err)
	}

	// Create a new user with an invalid signup type
	newUser := &models.User{
		Email:      "jkl@gmail.com",
		Password:   "jkl123",
		SignupType: "fgd",
	}

	// Call the CreateNewUser method
	err = CreateNewUser(gormDB, newUser)

	if err == nil {
		t.Error("Expected invalid signup error, but no error was returned.")
	} else {
		expectedError := "Invalid signup type"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("Expected error %v, got error %v", expectedError, err.Error())
		}
	}

	// Verify that all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestCreateNewUser_InvalidPassword(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create a mock database: %v", err)
	}
	defer db.Close()

	//Expect the query for the database version
	mock.ExpectQuery(".*").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.33"))

	// Initialize a new GORM database connection with the mock database
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open GORM connection: %v", err)
	}

	// Create a new user with an invalid password
	newUser := &models.User{
		Email:      "mno@gmail.com",
		Password:   "",
		SignupType: "Google",
	}

	// Call the CreateNewUser method
	err = CreateNewUser(gormDB, newUser)

	if err == nil {
		t.Error("Expected invalid password error, but no error was returned.")
	} else {
		expectedError := "Password cannot be empty"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("Expected error %v, got error %v", expectedError, err.Error())
		}
	}

	// Verify that all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestCreateNewUser_InvalidEmail(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create a mock database: %v", err)
	}
	defer db.Close()

	//Expect the query for the database version
	mock.ExpectQuery(".*").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.33"))

	// Initialize a new GORM database connection with the mock database
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open GORM connection: %v", err)
	}

	// Create a new user with an invalid email
	newUser := &models.User{
		Email:      "sur@4gmail",
		Password:   "12324",
		SignupType: "Google",
	}

	// Call the CreateNewUser method
	err = CreateNewUser(gormDB, newUser)

	if err == nil {
		t.Error("Expected invalid email error, but no error was returned.")
	} else {
		expectedError := "Invalid email"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("Expected error %v, got error %v", expectedError, err.Error())
		}
	}

	// Verify that all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

//func TestCreateNewUser_InvalidDuplicateEmail(t *testing.T) {
//	// Create a new mock database
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("Failed to create a mock database: %v", err)
//	}
//	defer db.Close()
//
//	//Expect the query for the database version
//	mock.ExpectQuery(".*").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.33"))
//
//	// Initialize a new GORM database connection with the mock database
//	gormDB, err := gorm.Open(mysql.New(mysql.Config{
//		Conn: db,
//	}), &gorm.Config{})
//	if err != nil {
//		t.Fatalf("Failed to open GORM connection: %v", err)
//	}
//
//	// Create a new user with a duplicate email
//	newUser := &models.User{
//		Email:      "abc@gmail.com",
//		Password:   "12324",
//		SignupType: "Google",
//	}
//
//	// Call the CreateNewUser method
//	err = CreateNewUser(gormDB, newUser)
//
//	if err == nil {
//		t.Error("Expected duplicate email error, but no error was returned.")
//	} else {
//		expectedError := fmt.Sprintf("Duplicate entry '%v' for key 'users.email'", newUser.Email)
//		//expectedError := "Duplicate entry"
//		if !strings.Contains(err.Error(), expectedError) {
//			t.Errorf("Expected error %v, got error %v", expectedError, err.Error())
//		}
//	}
//
//	// Verify that all expectations were met
//	err = mock.ExpectationsWereMet()
//	if err != nil {
//		t.Errorf("Unfulfilled expectations: %v", err)
//	}
//}
