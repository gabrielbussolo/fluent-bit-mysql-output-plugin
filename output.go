package main

import "C"
import (
	"context"
	"fmt"
	"github.com/calyptia/plugin"
)

func init() {
	plugin.RegisterOutput(pluginName, fmt.Sprintf("%s output plugin %s", pluginName, version), &mysql{})
}

var (
	pluginName = "mysql"
	version    = "dev-version"
)

type mysql struct {
	log plugin.Logger
}

func (m *mysql) Init(ctx context.Context, fbit *plugin.Fluentbit) error {
	m.log = fbit.Logger
	standardFields := []string{"Address", "Database", "Table", "User", "Password", "MinPoolSize", "MaxPoolSize", "Async"}
	for _, field := range standardFields {
		key := fbit.Conf.String(field)
		m.log.Info("key: %s, value: %s", field, key)
	}
	return nil
}

func (m *mysql) Flush(ctx context.Context, ch <-chan plugin.Message) error {
	return nil
}

func main() {
}
