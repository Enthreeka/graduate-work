package serialize

import (
	"encoding/json"
)

func Unmarshal[T any](src []byte) (T, error) {
	var args T

	if err := json.Unmarshal(src, &args); err != nil {
		return *(new(T)), err
	}

	return args, nil
}

func Marshal[T any](args T) ([]byte, error) {
	buf, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
