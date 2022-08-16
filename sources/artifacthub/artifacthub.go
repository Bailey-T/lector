package artifacthub

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPackage() (string, error) {
	var result map[string]interface{}

	resp, err := http.Get("https://artifacthub.io/api/v1/packages/helm/openstack-helm/elastic-filebeat")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	//Read in the body
	body, _ := io.ReadAll(resp.Body)
	//JSON Map creator
	json.Unmarshal(body, &result)
	
	//fmt.Printf("%v",string(body))
	fmt.Printf("%v",result["name"])
	
	for k,i := range result {
		fmt.Printf("%v : %v\n",k,i)
	}
	
	return string(body), nil}