package json

import (
	"encoding/json"
)

func Format(data interface{}) string {
	val, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(val)
}
