package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dashboard_poc.com/config"
	"github.com/dashboard_poc.com/db"
	"github.com/dashboard_poc.com/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
}
