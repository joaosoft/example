package example

import (
	"fmt"
	"github.com/joaosoft/dbr"

	"github.com/joaosoft/manager"
)

type AppConfig struct {
	Example *ExampleConfig `json:"example"`
}

type ExampleConfig struct {
	Port      int            `json:"port"`
	DbrConfig *dbr.DbrConfig `json:"dbr"`
	Log       struct {
		Level string `json:"level"`
	} `json:"log"`
}

func NewConfig() (*ExampleConfig, error) {
	appConfig := &AppConfig{}
	_, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", getEnv()), appConfig)

	return appConfig.Example, err
}
