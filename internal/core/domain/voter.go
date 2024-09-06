package domain

import (
	"fmt"
	"regexp"

	err_msg "github.com/axel-andrade/opina-ai-api/internal/core/domain/constants/errors"
)

type Voter struct {
	Base
	FullName  string `json:"full_name"`
	Cellphone string `json:"cellphone"`
}

func BuildNewVoter(fullName, cellphone string) (*Voter, error) {
	v := &Voter{
		FullName:  fullName,
		Cellphone: cellphone,
	}

	if err := v.validate(); err != nil {
		return nil, err
	}

	return v, nil
}

func isValidPhoneNumber(phone string) bool {
	// Regex para aceitar apenas números com código do país, código de área e número, sem o "+"
	// Exemplo de números aceitos: 5511912345678, 123456789012345
	phoneRegex := `^\d{1,3}\d{1,4}\d{6,10}$`
	re := regexp.MustCompile(phoneRegex)

	return re.MatchString(phone)
}
func (v *Voter) validate() error {
	if v.FullName == "" {
		return fmt.Errorf(err_msg.CONTACT_FULL_NAME_REQUIRED)
	}

	if v.Cellphone == "" {
		return fmt.Errorf(err_msg.CONTACT_CELLPHONE_REQUIRED)
	}

	if !isValidPhoneNumber(v.Cellphone) {
		return fmt.Errorf(err_msg.INVALID_CELLPHONE)
	}

	return nil
}
