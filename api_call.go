package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	DogfoodCloud                string = "AZUREDOGFOOD"
	PublicCloud                 string = "AZUREPUBLICCLOUD"
	AgentVersionEndpoint        string = ""
	AgentVersionDogFoodEndpoint string = ""
)

type AgentVersion struct {
	RepositoryPath string `json:"repositoryPath"`
}

func getAgentsVersion(loc, cloud, release_train string) {
	urlEndpoint := fmt.Sprintf(AgentVersionEndpoint, loc)
	if cloud == DogfoodCloud {
		urlEndpoint = AgentVersionDogFoodEndpoint
	}
	if release_train == "" {
		release_train = "stable"
	}
	params := "&releaseTrain=" + release_train
	response, err := http.Post(urlEndpoint+params, "application/json", nil)
	if err != nil {
		fmt.Printf("Request Failed: %s", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	// // Unmarshal result
	var agentVersion map[string]interface{}
	err = json.Unmarshal([]byte(body), &agentVersion)
	if err != nil {
		fmt.Printf("Reading body failed: %s", err)
		return
	}
	fmt.Println(agentVersion["repositoryPath"])
}

type ChartVariable struct {
	// Key
	Key string `json:"key,omitempty"`
	// Value
	Value string `json:"value,omitempty"`
}

func main() {
	getAgentsVersion("eastus", PublicCloud, "")
}
