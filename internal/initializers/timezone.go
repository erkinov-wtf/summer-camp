package initializers

import (
	"camp-summer/internal/config"
	"fmt"
	"log"
	"sync"
	"time"
)

var Timezone *time.Location
var timezoneOnce sync.Once

// LoadTimezone initializes the timezone for the application
func LoadTimezone() *time.Location {
	timezoneOnce.Do(func() {
		var err error
		Timezone, err = time.LoadLocation(config.Cfg.App.Timezone)
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to connect to init timezone: %v", err), "", "")
		}
	})

	log.Println("Timezone set successfully", "", "")
	return Timezone
}
