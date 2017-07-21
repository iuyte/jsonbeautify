package beautifyjson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func beautify(JSON string) (string, error) {
	var dat map[string]interface{}

	if err := json.Unmarshal([]byte(JSON), &dat); err != nil {
		return nil, err
	}

	b, err := json.MarshalIndent(dat, "", "  ")
	if err != nil {
		return nil, error
	}

	return b
}
