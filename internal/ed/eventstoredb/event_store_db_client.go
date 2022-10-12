package eventstoredb

import (
	"context"
	"encoding/json"
	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"github.com/mcmp/envent-driver/internal/ed"
)

var _ ed.EventClient = &EventStoreDBClient{}

type EventStoreDBClient struct {
	client *esdb.Client
}

func (e *EventStoreDBClient) InitClient(config ed.EventConfig) error {
	settings, err := esdb.ParseConnectionString(config.(*EventStoreDBConfig).ConnectionString)
	if err != nil {
		return err
	}
	e.client, err = esdb.NewClient(settings)
	if err != nil {
		return err
	}
	return nil
}
func (e EventStoreDBClient) SaveEvents(ctx context.Context, streamID string, event ed.Event) error {
	data, err := json.Marshal(event)
	if err != nil {

	}
	eventData := esdb.EventData{
		ContentType: esdb.ContentTypeJson,
		EventType:   "TestEvent",
		Data:        data,
	}
	e.client.AppendToStream(ctx, streamID, esdb.AppendToStreamOptions{}, eventData)
	return nil
}

func (e EventStoreDBClient) LoadEvents(ctx context.Context, streamID string) ([]ed.Event, error) {
	panic("implement me")
}
