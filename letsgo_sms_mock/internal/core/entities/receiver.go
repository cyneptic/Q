package entities

import (
	"database/sql/driver"
	"strings"
)

type Receiver []string

func (o *Receiver) Scan(src any) error {
	*o = strings.Split(src.(string), ",")
	return nil
}
func (o Receiver) Value() (driver.Value, error) {
	if len(o) == 0 {
		return nil, nil
	}
	return strings.Join(o, ","), nil
}
