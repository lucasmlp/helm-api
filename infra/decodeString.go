package infra

import "encoding/base64"

type decodeString interface {
	DecodeString(str string) ([]byte, error)
}

func DecodeString(str string) ([]byte, error) {

	decodedStr, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return decodedStr, nil
}
