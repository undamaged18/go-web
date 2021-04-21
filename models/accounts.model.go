package models

import (
	"fmt"
	"github.com/undamaged18/checker"
	"unicode"
)

type user struct {
	EmailAddress    string `db:"email_address" schema:"email_address"`
	FirstName       string `db:"first_name" schema:"first_name"`
	LastName        string `db:"last_name" schema:"last_name"`
	Password        string `db:"password" schema:"password"`
	ConfirmPassword string `db:"confirm_password" schema:"confirm_password"`
}

type FormErrors struct {
	Errors, Password, FirstName, LastName, EmailAddress bool
}

func NewUser() *user {
	return &user{}
}

func (u *user) Authenticate() bool {
	fmt.Println(u)
	return true
}

func (u *user) Create() FormErrors {
	var formErr FormErrors
	if u.Password != u.ConfirmPassword {
		formErr.Password = true
	}
	var (
		hasMinLen = false
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)
	if len(u.Password) >= 8 {
		hasMinLen = true
	}
	for _, char := range u.Password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}
	if !hasNumber || !hasUpper || !hasLower || !hasMinLen {
		formErr.Errors = true
		formErr.Password = true
	}

	if u.FirstName == "" {
		formErr.Errors = true
		formErr.FirstName = true
	}

	if u.LastName == "" {
		formErr.Errors = true
		formErr.LastName = true
	}

	if err := checker.Email().Host(u.EmailAddress); err != nil {
		formErr.Errors = true
		formErr.EmailAddress = true
	}

	fmt.Println(u)
	return formErr
}

func (u *user) Update() {

}

func (u *user) Lock() {

}

func (u *user) Unlock() {

}
