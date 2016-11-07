package beater

import (
	"github.com/elufimov/consulbeat/config"
	"testing"
	"time"
)

func TestMissingTagsBlockInEventsFiltering(t *testing.T) {
	var config = config.Config{
		1 * time.Second,
		"",
		nil,
		true,
	}

	servicesHealth := []ServiceHealth{}

	servicesHealth = append(
		servicesHealth,
		ServiceHealth{
			Node{"", "", TaggedAddresses{"", "", }, },
			Service{"", "", []string{"prod"}, "", 0,
			},
			[]Check{Check{"", "", "", "", "", "", "", "", }, },
		},
	)
	servicesHealth = append(
		servicesHealth,
		ServiceHealth{
			Node{"", "", TaggedAddresses{"", "", }, },
			Service{"", "", []string{"test"}, "", 0,
			},
			[]Check{Check{"", "", "", "", "", "", "", "", }, },
		},
	)
	events := servicesHealthToEvents(servicesHealth, config)
	length := len(events)
	if length != 2 {
		t.Error("Events count is not equal to 2", length)
	}
}

func TestFilterByTagBlockInEventsFiltering(t *testing.T) {
	var config = config.Config{
		1 * time.Second,
		"",
		[]string{"prod"},
		true,
	}

	servicesHealth := []ServiceHealth{}

	servicesHealth = append(
		servicesHealth,
		ServiceHealth{
			Node{"", "", TaggedAddresses{"", "", }, },
			Service{"", "", []string{"prod"}, "", 0,
			},
			[]Check{Check{"", "", "", "", "", "", "", "", }, },
		},
	)
	servicesHealth = append(
		servicesHealth,
		ServiceHealth{
			Node{"", "", TaggedAddresses{"", "", }, },
			Service{"", "", []string{"test"}, "", 0,
			},
			[]Check{Check{"", "", "", "", "", "", "", "", }, },
		},
	)
	events := servicesHealthToEvents(servicesHealth, config)
	length := len(events)
	if length != 1 {
		t.Error("Events count is not equal to 1", length)
	}
}
