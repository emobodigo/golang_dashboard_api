package util

import (
	"encoding/json"
	"strconv"
)

type StringInt int

func (st *StringInt) UnmarshalJSON(b []byte) error {
	var item interface{}
	err := json.Unmarshal(b, &item)
	if err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*st = StringInt(v)
	case float64:
		*st = StringInt(int(v))
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*st = StringInt(i)
	}
	return nil
}
