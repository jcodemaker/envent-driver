package ed

import "context"

type EventClient interface {
	InitClient(EventConfig) error

	SaveEvents(ctx context.Context, streamID string, events Event) error

	// LoadEvents loads all events for the aggregate id from the store.
	LoadEvents(ctx context.Context, streamID string) ([]Event, error)
}
