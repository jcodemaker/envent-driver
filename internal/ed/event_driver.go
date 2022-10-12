package ed

import (
	_ "github.com/lib/pq"
)

//事件类型
type EventDriverType int

const (
	_ EventDriverType = iota
	EVENT_STORE_DB
	KAFKA
)

var EventDriverTypeMap = map[string]EventDriverType{
	"event_store_db": EVENT_STORE_DB,
	"kafka":          KAFKA,
}

type EventDriver interface {
	ConstructDriver(EventConfig) error
}
