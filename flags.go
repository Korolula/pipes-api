package main

import (
	"github.com/namsral/flag"
	"os"
)

var (
	port          int
	workdir       string
	bugsnagAPIKey string
	environment   string
	dbConnString  string
)

func InitFlags() {
	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "PIPES_API", flag.ExitOnError)

	fs.IntVar(&port, "port", 8100, "port")
	fs.StringVar(&workdir, "workdir", ".", "Workdir of server")
	fs.StringVar(&bugsnagAPIKey, "bugsnag_key", "", "Bugsnag API Key")
	fs.StringVar(&environment, "environment", "development", "Environment")
	fs.StringVar(&dbConnString, "db_conn_string", "dbname=pipes_development user=pipes_user host=localhost sslmode=disable port=5432", "DB Connection String")

	fs.Parse(os.Args[1:])
}
