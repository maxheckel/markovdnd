package store

import (
	"encoding/json"
	"io/ioutil"
)

func WriteJson(filename string, obj interface{}) error {
	file, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, file, 0644)
}