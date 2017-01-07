package mapstruct

import (
	"bytes"
	"encoding/json"
)

func Decode(r map[string]interface{}, v interface{}) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(&r); err != nil {
		return err
	}

	if err := json.NewDecoder(&buf).Decode(&v); err != nil {
		return err
	}

	return nil
}
