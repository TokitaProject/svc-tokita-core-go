package restapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func HttpGet(param map[string]interface{}) (result map[string]interface{}, err error) {
	request, err := http.NewRequest("GET", param["url"].(string), nil)

	if err != nil {
		return
	}

	query := request.URL.Query()

	for i, x := range param["query"].(map[string]string) {
		query.Add(i, x)
	}

	// FOR LOCAL ONLY
	if os.Getenv("ENV") == "local" {
		request.Header.Set("Authorization", os.Getenv("BEARER"))
	}

	request.URL.RawQuery = query.Encode()

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return
	}

	if response.StatusCode > 300 {
		err = fmt.Errorf(string(body))
		return
	}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return
	}

	return
}