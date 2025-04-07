package mappers

import (
	"strconv"

	"github.com/xamust/ekom-test/internal/interfaces"
)

var _ interfaces.IDMappers = (*IDMappers)(nil)

type IDMappers struct {
}

func newIDMappers() interfaces.IDMappers {
	return &IDMappers{}
}

func (id *IDMappers) FromString(in string) (uint64, error) {
	result, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
