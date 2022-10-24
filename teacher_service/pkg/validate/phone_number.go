package validate

import (
	"fmt"
	"strings"
)


func PhoneNumber(phone_number string) error {
	if len(phone_number) != 13 {
		return fmt.Errorf("number should be 13 characters long, but was %d", len(phone_number))
	}
	if !strings.HasPrefix(phone_number, "+998") {
		return fmt.Errorf("number should  have UZB prefix +998")
	}
	return nil
}