package validate

import "net/mail"

// Email validate weather given string is invalid format for  uz
func Email(email string) error{
	if _, err := mail.ParseAddress(email); err != nil {
		return err
	}
	return nil
}