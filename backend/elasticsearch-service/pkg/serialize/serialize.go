package serialize

import "encoding/json"

func Unmarshal[T any](src string) (T, error) {
	var args T

	if err := json.Unmarshal([]byte(src), &args); err != nil {
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
