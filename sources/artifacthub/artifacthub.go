package artifacthub

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetHelmPackage(r, m string) (string, error) {
	var result map[string]interface{}
	resp, err := http.Get("https://artifacthub.io/api/v1/packages/helm/"+r+"/"+m)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	//Read in the body
	body, _ := io.ReadAll(resp.Body)
	//JSON Map creator
	json.Unmarshal(body, &result)

	log.Printf("%v: {Chart: %v, App: %v} \n", result["name"], result["version"], result["app_version"])	

	
	return "", nil}