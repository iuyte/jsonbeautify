package jsonbeautify

import "encoding/json"

func Beautify(JSON string) (string, error) {
	var dat map[string]interface{}

	if err := json.Unmarshal([]byte(JSON), &dat); err != nil {
		return "", err
	}

	b, err := json.MarshalIndent(dat, "", "  ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
