package beater

import (
	"fmt"
	"time"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
	"github.com/pravoru/consulbeat/config"
)

type Consulbeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Consulbeat{
		done:   make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *Consulbeat) Run(b *beat.Beat) error {
	logp.Info("consulbeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)

	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		servicesNames := getAllServices(bt)

		var servicesHealthStatuses = []ServiceHealth{}

		for _, value := range servicesNames {
			result := getServiceHealth(bt, value)
			servicesHealthStatuses = append(servicesHealthStatuses, result...)
		}

		events := servicesHealthToEvents(servicesHealthStatuses, bt.config)
		bt.client.PublishEvents(events)
		logp.Info("Event sent")
	}
}

func (bt *Consulbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

func servicesHealthToEvents(servicesHealth []ServiceHealth, config config.Config) []common.MapStr {
	var events = []common.MapStr{}

	tagsFromConfigMap := stringArrayToStringBoolMap(config.ServicesTags)

	eventType := "check"

	for _, healthStatusesWithChecks := range servicesHealth {
		// Iterate over all checks that can be in each health status
		for _, check := range healthStatusesWithChecks.Checks {
			// Check if at least one tag match tags list from config
			if (config.ServicesTags != nil) {
				for _, tag := range healthStatusesWithChecks.Service.Tags {
					if (tagsFromConfigMap[tag]) {
						event := common.MapStr{
							"@timestamp":   common.Time(time.Now()),
							"type":         eventType,
							"check":        check,
							"node":         healthStatusesWithChecks.Node,
							"service":      healthStatusesWithChecks.Service,
						}
						events = append(events, event)
					}
				}
			} else {
				event := common.MapStr{
					"@timestamp":   common.Time(time.Now()),
					"type":         eventType,
					"check":        check,
					"node":         healthStatusesWithChecks.Node,
					"service":      healthStatusesWithChecks.Service,
				}
				events = append(events, event)
			}
		}
	}

	return events
}
