package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CallAPI[T any](method, url, contentType string, body interface{}) (*T, error) {
	var bodyReader io.Reader

	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(bodyBytes)
	}

	var result T
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return &result, err
	}
	req.Header.Set("Content-Type", contentType)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &result, err
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		var errMsg string
		if err := json.Unmarshal(respBody, &errMsg); err != nil {
			return &result, fmt.Errorf("API error: %s", string(respBody))
		}
		return &result, fmt.Errorf("API error: %s", errMsg)
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		return &result, err
	}

	return &result, nil
}
