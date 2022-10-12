package eventstoredb

import "github.com/mcmp/envent-driver/internal/ed"

var _ ed.EventConfig = &EventStoreDBConfig{}

type EventStoreDBConfig struct {
	ConnectionString string `mapstructure:"connectionString"`
}

func (e *EventStoreDBConfig) InitDriverConfig() (*ed.EventConfig, error) {
	panic("implement me")
}
