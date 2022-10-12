package main

import (
	"context"
	"github.com/mcmp/envent-driver/internal/ed"
	"github.com/mcmp/envent-driver/internal/ed/eventstoredb"
	"github.com/mcmp/envent-driver/pkg/driver"
	_ "github.com/mcmp/envent-driver/pkg/driver"
)

func main() {
	driver := driver.GetDriver(ed.EVENT_STORE_DB)
	event := ed.Event{
		EventID:   "weeffsd",
		EventType: "test",
	}
	driver.(*eventstoredb.EventStoreDBDriver).Client.SaveEvents(context.Background(), "test-sss", event)
}
