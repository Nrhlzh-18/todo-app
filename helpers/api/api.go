package api

import (
	"io/ioutil"
	"net/http"
)

func Get(url string) ([]byte, error) {
	var result []byte

	response, err := http.Get(url)
	if err != nil {
		return result, err
	}

	result, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	return result, nil
}
