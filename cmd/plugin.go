package main

import (
	"context"
	"encoding/json"
	"fmt"
	fb "github.com/calyptia/plugin"

	"github.com/gabrielbussolo/fluent-bit-mysql-output-plugin/internal"
)

func init() {
	fb.RegisterOutput(pluginName, fmt.Sprintf("%s output plugin %s", pluginName, version), &plugin{})
}

var (
	pluginName = "mysql"
	version    = "dev-version"
)

type plugin struct {
	log   fb.Logger
	mysql *internal.Mysql
}

func (p *plugin) Init(ctx context.Context, fb *fb.Fluentbit) error {
	p.log = fb.Logger
	dsn, table, err := validateConfigKeys(fb)
	if err != nil {
		return err
	}
	p.mysql, err = internal.New(dsn, table)
	if err != nil {
		return err
	}
	return nil
}

func (p *plugin) Flush(ctx context.Context, ch <-chan fb.Message) error {
	for msg := range ch {
		marshal, err := json.Marshal(msg.Record)
		if err != nil {
			p.log.Error("could not marshal record: %v", err)
		}
		err = p.mysql.Write(ctx, msg.Time, msg.Tag(), marshal)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
}

// validateConfigKeys validates the configuration keys and returns the DSN and table name
func validateConfigKeys(fbit *fb.Fluentbit) (string, string, error) {
	address := fbit.Conf.String("Address")
	if address == "" {
		fbit.Logger.Warn("Address is not set, using default")
		address = "localhost:3306"
	}
	user := fbit.Conf.String("User")
	if user == "" {
		fbit.Logger.Warn("User is not set, using default")
		user = "root"
	}
	passwd := fbit.Conf.String("Password")
	if passwd == "" {
		fbit.Logger.Error("Password is required")
		return "", "", fmt.Errorf("password is a required parameter")
	}
	db := fbit.Conf.String("Database")
	if db == "" {
		fbit.Logger.Error("Database is required")
		return "", "", fmt.Errorf("database is a required parameter")
	}
	table := fbit.Conf.String("Table")
	if table == "" {
		fbit.Logger.Error("Table is required")
		return "", "", fmt.Errorf("table is a required parameter")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, passwd, address, db)
	return dsn, table, nil
}
