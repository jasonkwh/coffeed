package config

import "time"

type Config struct {
	LogLevel        int8
	PollingInterval time.Duration

	Integration integration
}

type integration struct {
	GoogleCalendar struct {
		ApiKey string
	}
}
