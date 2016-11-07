package beater

import (
	"github.com/pravoru/ConsulBeat/config"
	"testing"
	"time"
)

func TestMissingTagsBlockInServiceFiltering(t *testing.T) {
	var config = config.Config{
		1 * time.Second,
		"",
		nil,
		true,
	}
	mapOfServices := make(map[string][]string)
	mapOfServices["service1"] = append(mapOfServices["service1"], "service1_tag1")
	mapOfServices["service1"] = append(mapOfServices["service1"], "service1_tag2")
	mapOfServices["service2"] = append(mapOfServices["service2"], "service2_tag1")
	mapOfServices["service2"] = append(mapOfServices["service2"], "service2_tag2")
	result := filterKeys(mapOfServices, config)
	length := len(result)
	if length != 2 {
		t.Error("Keys count is not equal to 2", length)
	}
	if stringArrayToStringBoolMap(result)["service1"] != true {
		t.Error("Key service1 is not present in result", stringArrayToStringBoolMap(result)["service1"])
	}
	if stringArrayToStringBoolMap(result)["service2"] != true {
		t.Error("Key service2 is not present in result", stringArrayToStringBoolMap(result)["service2"])
	}

}

func TestFilterByTagInServiceFiltering(t *testing.T) {
	var config = config.Config{
		1 * time.Second,
		"",
		[]string{"service1_tag1"},
		true,
	}
	mapOfServices := make(map[string][]string)
	mapOfServices["service1"] = append(mapOfServices["service1"], "service1_tag1")
	mapOfServices["service1"] = append(mapOfServices["service1"], "service1_tag2")
	mapOfServices["service2"] = append(mapOfServices["service2"], "service2_tag1")
	mapOfServices["service2"] = append(mapOfServices["service2"], "service2_tag2")
	result := filterKeys(mapOfServices, config)
	length := len(result)
	if length != 1 {
		t.Error("Keys count is not equal to 1", length)
	}
	if stringArrayToStringBoolMap(result)["service1"] != true {
		t.Error("Key service1 is not present in result", stringArrayToStringBoolMap(result)["service1"])
	}

}
