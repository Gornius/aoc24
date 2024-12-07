package clone

import (
	"bytes"
	"encoding/gob"
)

func GobDeepClone[T any](obj T) (*T, error) {
	buf := bytes.Buffer{}
	err := gob.NewEncoder(&buf).Encode(obj)
	if err != nil {
		return nil, err
	}
	var cloned T
	err = gob.NewDecoder(&buf).Decode(&cloned)
	if err != nil {
		return nil, err
	}

	return &cloned, nil
}
