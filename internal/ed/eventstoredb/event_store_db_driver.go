package eventstoredb

import (
	"github.com/mcmp/envent-driver/internal/ed"
)

var _ ed.EventDriver = &EventStoreDBDriver{}

type EventStoreDBDriver struct {
	Client EventStoreDBClient
	Config *EventStoreDBConfig
}

func (e *EventStoreDBDriver) ConstructDriver(config ed.EventConfig) error {
	client := EventStoreDBClient{}
	err := client.InitClient(config)
	if err != nil {
		return err
	}
	e.Client = client
	e.Config = config.(*EventStoreDBConfig)
	return nil
}
