package templates

import (
	"io/ioutil"
)

func FileContent() (string, error) {
	bts, err := ioutil.ReadFile("foo.txt")
	return string(bts), err
}
