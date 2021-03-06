package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type RequestOptions struct {
	Url     string
	Method  string
	Body    io.Reader
	Query   map[string]string
	Headers map[string]string
}

func Request(options *RequestOptions, result interface{}) (
	err error,
) {
	client := &http.Client{}

	req, err := http.NewRequest(options.Method, options.Url, options.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if options.Query != nil {
		q := req.URL.Query()
		for index, value := range options.Query {
			q.Add(index, value)
		}

		req.URL.RawQuery = q.Encode()
	}

	if options.Headers != nil {
		for prop, value := range options.Headers {
			req.Header.Add(prop, value)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		if res.StatusCode >= 400 {
			fmt.Println(res)
			// 이새끼덜 넘 불친절하네 이유도 업시 400이냐?
			return errors.New("Response Error status_code " + res.Status)
		}
	}

	defer res.Body.Close()

	Body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(Body, result)
	if err != nil {
		return
	}
	return
}
