package templates

import (
	"io/ioutil"
)

func FileContent() (string, error) {
	bts, err := ioutil.ReadFile("templates/foo.txt")
	return string(bts), err
}
