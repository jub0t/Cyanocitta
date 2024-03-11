package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ParseJSON[T any](rawBody io.Reader, dst *T) error {
	bodyBytes, err := ioutil.ReadAll(rawBody)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, dst)
	if err != nil {
		return err
	}

	return nil
}
