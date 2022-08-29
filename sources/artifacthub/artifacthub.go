package artifacthub

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetHelmPackage(r, m string) (map[string]interface{}, error) {
	var result map[string]interface{}
	resp, err := http.Get("https://artifacthub.io/api/v1/packages/helm/" + r + "/" + m)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//Read in the body
	body, _ := io.ReadAll(resp.Body)
	//JSON Map creator
	json.Unmarshal(body, &result)

	return result, err
}
