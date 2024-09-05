package value_object

import (
	"fmt"
	"regexp"

	err_msg "github.com/axel-andrade/opina-ai-api/internal/core/domain/constants/errors"
)

type Email struct {
	Value string
}

func (e *Email) Validate() error {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if regex.MatchString(e.Value) {
		return nil
	}

	return fmt.Errorf(err_msg.INVALID_EMAIL)
}
