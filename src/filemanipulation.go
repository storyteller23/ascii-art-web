package asciiart

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("no file or directory")
	}
	return string(bytes), nil
}
