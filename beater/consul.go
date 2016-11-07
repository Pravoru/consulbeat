package beater

import (
	"net/http"
	"encoding/json"
	"github.com/elastic/beats/libbeat/logp"
	"io/ioutil"
	"github.com/pravoru/ConsulBeat/config"
)

func getAllServices(bt *Consulbeat) []string {
	bodyBytes := makeRequest(bt.config.ConsulURL + "/v1/catalog/services", bt.config.FailOnHttpError)
	mapOfServices := map[string][]string{}
	unmarshalError := json.Unmarshal(bodyBytes, &mapOfServices)
	if unmarshalError != nil {
		logp.Err(unmarshalError.Error())
	}
	keys := filterKeys(mapOfServices, bt.config)
	return keys
}

func getServiceHealth(bt *Consulbeat, serviceName string) []ServiceHealth {
	bodyBytes := makeRequest(bt.config.ConsulURL + "/v1/health/service/" + serviceName, bt.config.FailOnHttpError)
	var serviceHealth []ServiceHealth
	unmarshalError := json.Unmarshal(bodyBytes, &serviceHealth)
	if unmarshalError != nil {
		logp.Err(unmarshalError.Error())
	}
	return serviceHealth
}

func makeRequest(url string, failOnHttpError bool) []byte {
	response, responseError := http.Get(url)
	if responseError != nil {
		logp.Err(responseError.Error())
		if (failOnHttpError) {
			panic(responseError)
		}
	}
	defer response.Body.Close()
	bodyBytes, readError := ioutil.ReadAll(response.Body)
	if (readError != nil) {
		logp.Err(readError.Error())
		if (failOnHttpError) {
			panic(readError)
		}
	}
	return bodyBytes
}

func filterKeys(mapOfServices map[string][]string, config config.Config) []string {
	tagsFromConfigMap := stringArrayToStringBoolMap(config.ServicesTags)
	keys := []string{}
	if(config.ServicesTags == nil) {
		for key, _ := range mapOfServices {
			keys = append(keys, key)
		}
	} else {
		for key, values := range mapOfServices {
			valueExist := false
			for _, value := range values {
				if (tagsFromConfigMap[value]) {
					valueExist = true
				}
			}
			if (valueExist) {
				keys = append(keys, key)
			}
		}
	}
	return keys
}