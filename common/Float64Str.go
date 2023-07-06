package common

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Float64Value float64

func (lv *Float64Value) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	if str == "" {
		*lv = 0
	} else {
		value, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return err
		}
		*lv = Float64Value(value)
	}

	return nil
}

func (lv Float64Value) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%g"`, lv)), nil
}

func (lv *Float64Value) Value() float64 {
	return float64(*lv)
}

func (lv *Float64Value) String() string {
	return fmt.Sprintf("%f", *lv)
}
