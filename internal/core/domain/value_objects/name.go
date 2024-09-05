package value_object

import (
	"fmt"

	err_msg "github.com/axel-andrade/opina-ai-api/internal/core/domain/constants/errors"
)

type Name struct {
	Value string
}

func (n *Name) Validate() error {
	length := len(n.Value)

	if length <= 0 {
		return fmt.Errorf(err_msg.INVALID_NAME)
	}

	return nil
}
