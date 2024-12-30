package main

import "encoding/json"

func JsonDecode[T any](data []byte, input *T) error {
	return json.Unmarshal(data, input)
}
