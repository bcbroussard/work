package work

import (
	"fmt"

	"github.com/pkg/errors"
)

func logError(key string, err error) {
	fmt.Printf("ERROR: %s - %s\n", key, errors.WithStack(err))
}
