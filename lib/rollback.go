package lib

import (
	"_models/ent"
	"fmt"
)

func Rollback(tx *ent.Tx, err error) error {
	if rbErr := tx.Rollback(); rbErr != nil {
		err = fmt.Errorf("%w: %v", err, rbErr)
	}
	return err
}
