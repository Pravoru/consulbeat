package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/pravoru/ConsulBeat/beater"
)

func main() {
	err := beat.Run("consulbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
