package ed

type EventConfig interface {
	InitDriverConfig() (*EventConfig, error)
}
