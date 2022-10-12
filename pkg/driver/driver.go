package driver

import (
	"flag"
	"fmt"
	"github.com/mcmp/envent-driver/internal/ed"
	"github.com/mcmp/envent-driver/internal/ed/eventstoredb"
	constants "github.com/mcmp/envent-driver/pkg/constant"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"sync"
)

var (
	driversMu  sync.RWMutex
	configPath string
	drivers    = make(map[ed.EventDriverType]ed.EventDriver)
)

type Config struct {
	EventStoreConfig EventStoreConfig `mapstructure:"eventStore"`
}

type EventStoreConfig struct {
	StoreType string      `mapstructure:"storeType"`
	Config    interface{} `mapstructure:"config"`
}

func init() {
	flag.StringVar(&configPath, "config", "", "ES microservice config path")
}

func init() {
	//根据配置，获取driver类型 evenstoredb,kafka,rabbit
	cfg, ecfg, err := InitConfig()
	if err != nil {

	}
	driverInst := getDriverByType(cfg.EventStoreConfig.StoreType)
	driverInst.ConstructDriver(ecfg)
	register(ed.EventDriverTypeMap[cfg.EventStoreConfig.StoreType], driverInst)
}

func InitConfig() (*Config, ed.EventConfig, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.ConfigPath)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/config/config.yaml", getwd)
		}
	}
	cfg := &Config{}
	result := eventstoredb.EventStoreDBConfig{}

	viper.SetConfigType(constants.Yaml)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, nil, errors.Wrap(err, "viper.Unmarshal")
	}
	mapstructure.Decode(cfg.EventStoreConfig.Config, &result)

	return cfg, &result, nil
}

func register(t ed.EventDriverType, d ed.EventDriver) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if d == nil {
		panic("sql: Register driver is nil")
	}
	if _, dup := drivers[t]; dup {
		panic("sql: Register called twice for driver ")
	}
	drivers[t] = d
}

func getDriverByType(storeType string) ed.EventDriver {
	if ed.EventDriverTypeMap[storeType] == ed.EVENT_STORE_DB {
		test := eventstoredb.EventStoreDBDriver{}
		return &test
	}
	return nil
}

func GetDriver(storeType ed.EventDriverType) ed.EventDriver {
	return drivers[storeType]
}
