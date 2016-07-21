package forms

import (
	"regexp"

	"github.com/goatcms/goatcms/models"
)

// RegisterForm is structure with register form values
type RegisterForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

const (
	errNoEmail        = "You must supply an email"
	errNoPassword     = "You must supply a password"
	errBadEmailFormat = "Email is not correct"
	errPassTooShort   = "Password should be at least 6 characters"
	errEmailOccupied  = "There is already a user with email "
)

const (
	passwordLength = 6
)

// Validate validate form and return bool how validation passed
func (f RegisterForm) Validate(u models.UserDTO) (bool, map[string][]string) {
	validation := true
	errors := make(map[string][]string)
	// do validation
	if u != nil {
		errors["email"] = append(errors["email"], errEmailOccupied+u.GetEmail())
		validation = false
	}
	if validateEmail(f.Email) == false {
		errors["email"] = append(errors["email"], errBadEmailFormat)
		validation = false
	}
	if f.Email == "" {
		errors["email"] = []string{errNoEmail}
		validation = false
	}
	if len(f.Password) < passwordLength {
		errors["pass"] = append(errors["pass"], errPassTooShort)
		validation = false
	}
	if f.Password == "" {
		errors["pass"] = []string{errNoPassword}
		validation = false
	}
	return validation, errors
}

func validateEmail(email string) bool {
	result := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return result.MatchString(email)
}
