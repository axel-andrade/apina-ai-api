package domain

import (
	"fmt"
	"regexp"

	err_msg "github.com/axel-andrade/opina-ai-api/internal/core/domain/constants/errors"
)

type Contact struct {
	Base
	FullName  string `json:"full_name"`
	Cellphone string `json:"cellphone"`
}

func BuildNewContact(fullName, cellphone string) (*Contact, error) {
	c := &Contact{
		FullName:  fullName,
		Cellphone: cellphone,
	}

	if err := c.validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func isValidPhoneNumber(phone string) bool {
	// The expression below accepts phone numbers in the format +55 11 99999-9999
	// Accepts +55 (or other country codes), followed by numbers
	phoneRegex := `^\+\d{1,3}\s\d{10,15}$`
	re := regexp.MustCompile(phoneRegex)
	return re.MatchString(phone)
}

func (c *Contact) validate() error {
	if c.FullName == "" {
		return fmt.Errorf(err_msg.CONTACT_FULL_NAME_REQUIRED)
	}

	if c.Cellphone == "" {
		return fmt.Errorf(err_msg.CONTACT_CELLPHONE_REQUIRED)
	}

	if !isValidPhoneNumber(c.Cellphone) {
		return fmt.Errorf(err_msg.INVALID_CELLPHONE)
	}

	return nil
}
