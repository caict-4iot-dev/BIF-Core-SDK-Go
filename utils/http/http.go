package http

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

// HttpGet ...
func HttpGet(url string) ([]byte, error) {
	client := &http.Client{}
	newRequest, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(newRequest)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if body == nil {
		return nil, errors.New("response body is nil ")
	}

	return body, nil
}

// HttpPost ...
func HttpPost(url string, data []byte) ([]byte, error) {
	client := &http.Client{}
	newRequest, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	response, err := client.Do(newRequest)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
