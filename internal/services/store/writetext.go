package store

import (
	"io/ioutil"
)

func WriteText(filename, text string) error {
	return ioutil.WriteFile(filename, []byte(text), 0644)
}