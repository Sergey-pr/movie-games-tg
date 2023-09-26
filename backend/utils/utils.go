package utils

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

func JsonScan[T any](src interface{}, b T) error {
	switch src := src.(type) {
	case []byte:
		if string(src) == "null" {
			src = []byte("{}")
		}
		return json.Unmarshal(src, b)
	case string:
		return json.Unmarshal([]byte(src), b)
	case nil:
		return nil
	}
	return fmt.Errorf("cannot convert %T", src)
}

type JSONB []interface{}

func (b *JSONB) Scan(src interface{}) error {
	return JsonScan(src, b)
}

func (b JSONB) Value() (driver.Value, error) {
	if b == nil {
		return json.Marshal([]interface{}{})
	}
	return json.Marshal(b)
}

func (b *JSONB) AsFloatSlice() []float64 {
	arr := make([]float64, 0)
	for _, value := range *b {
		arr = append(arr, value.(float64))
	}

	return arr
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func ToGenericArray[T any](arr []T) []interface{} {
	res := make([]interface{}, len(arr))
	for idx, v := range arr {
		res[idx] = v
	}
	return res
}
