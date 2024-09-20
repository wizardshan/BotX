package field

import (
	"botx/pkg/ternary"
	"botx/pkg/validate"
	"fmt"
)

type IDV1 struct {
	Value int64
}

func (id *IDV1) Valid() error {
	return ternary.Do(validate.IsPositive(id.Value), nil, fmt.Errorf(validate.IsPositiveMsg, id.Name()))
}

func (id *IDV1) ValidOmit() error {
	return ternary.Do(validate.IsEmpty(id.Value), nil, id.Valid())
}

func (id *IDV1) Name() string {
	return "ID"
}
