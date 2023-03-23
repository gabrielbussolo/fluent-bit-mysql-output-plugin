package main

import (
	"database/sql"
	"github.com/ory/dockertest/v3"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestIntegration(t *testing.T) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatal(err)
	}
	network, err := pool.CreateNetwork("fluent_bit_network")
	if err != nil {
		t.Fatal(err)
	}

	mysql, err := pool.RunWithOptions(&dockertest.RunOptions{
		Networks:     []*dockertest.Network{network},
		Name:         "fluent_bit_mysql",
		Repository:   "mysql",
		Tag:          "latest",
		ExposedPorts: []string{"3306"},
		Env:          []string{"MYSQL_ROOT_PASSWORD=my-secret-pw", "MYSQL_DATABASE=fluent_bit"},
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	conn, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/fluent_bit")
	if err != nil {
		t.Fatal(err)
	}
	_, err = conn.Exec(`CREATE TABLE IF NOT EXISTS test (id INT NOT NULL AUTO_INCREMENT, datetime DATETIME, tag VARCHAR(255), data JSON, PRIMARY KEY (id))`)
	if err != nil {
		t.Fatal(err)
	}

	fluentBitAgent, err := pool.BuildAndRunWithBuildOptions(&dockertest.BuildOptions{
		ContextDir: "../.",
		Dockerfile: "Dockerfile",
		Platform:   "linux/amd64",
	}, &dockertest.RunOptions{
		Networks: []*dockertest.Network{network},
		Name:     "fluent_bit_agent",
	})

	if err != nil {
		return
	}

	rows, err := conn.Query("SELECT * FROM test")
	if !rows.Next() {
		t.Error("no rows found")
	}
	for rows.Next() {
		var id int
		var datetime string
		var tag string
		var data string
		err = rows.Scan(&id, &datetime, &tag, &data)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("id: %d, datetime: %s, tag: %s, data: %s", id, datetime, tag, data)
	}

	err = pool.Purge(fluentBitAgent)
	if err != nil {
		t.Fatal(err)
	}
	err = pool.Purge(mysql)
	if err != nil {
		t.Fatal(err)
	}
	err = pool.RemoveNetwork(network)
	if err != nil {
		t.Fatal(err)
	}
}
